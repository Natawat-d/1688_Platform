// Package catalog imports 1688 offers into our own tables and keeps the
// denormalised sell price in step with the fee rules.
//
// Nothing here calls the gateway during a shopper's request. Search and product
// pages read our database only, so the storefront stays fast and deterministic
// and a rate limit on the 1688 side can never take the shop down.
package catalog

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"strconv"
	"strings"
	"time"

	"marketplace/internal/ali"
	"marketplace/internal/pricing"
	"marketplace/internal/store"
)

// Importer pulls offers from the gateway and writes them to the database.
type Importer struct {
	DB      *store.DB
	Cli     *ali.Client
	Log     *slog.Logger
	Country string // language code for the multilingual endpoints, e.g. "en"
}

// New builds an importer.
func New(db *store.DB, cli *ali.Client, log *slog.Logger) *Importer {
	return &Importer{DB: db, Cli: cli, Log: log, Country: "en"}
}

// ImportOffers fetches and stores each offer, returning how many landed. One
// failure does not abandon the batch: the reason is logged and the rest continue,
// because a single delisted offer should not fail a three-hundred-item import.
func (im *Importer) ImportOffers(ctx context.Context, offerIDs []int64) (int, error) {
	n := 0
	for _, id := range offerIDs {
		if err := ctx.Err(); err != nil {
			return n, err
		}
		if err := im.ImportOffer(ctx, id); err != nil {
			im.Log.Warn("import offer failed", "offerId", id, "err", err)
			continue
		}
		n++
	}
	return n, nil
}

// ImportOffer fetches one offer's detail, estimates its China-side freight once,
// and upserts the product with its SKUs and price tiers.
func (im *Importer) ImportOffer(ctx context.Context, offerID int64) error {
	detail, err := im.Cli.OfferDetail(ctx, ali.OfferDetailQuery{
		OfferID: ali.ID(offerID),
		Country: im.Country,
	})
	if err != nil {
		return fmt.Errorf("detail: %w", err)
	}

	fp, raw := mapDetail(detail)

	// Freight is cached at import rather than quoted at checkout: it barely moves,
	// the estimate wants district codes, and this keeps checkout down to one
	// gateway call.
	if addr, ok := im.warehouse(ctx); ok && len(fp.SKUs) > 0 {
		fr, ferr := im.Cli.EstimateFreight(ctx, ali.FreightQuery{
			OfferID:        ali.ID(offerID),
			ToProvinceCode: addr.ProvinceCode,
			ToCityCode:     addr.CityCode,
			ToCountryCode:  addr.DistrictCode, // documented as toCountryCode; it is a district
			TotalNum:       1,
			LogisticsSkuNumModels: []ali.FreightSkuNum{
				{SkuID: ali.StringID(fp.SKUs[0].SkuID), Number: 1},
			},
		})
		switch {
		case ferr != nil:
			im.Log.Debug("freight estimate unavailable", "offerId", offerID, "err", ferr)
		case bool(fr.FreePostage):
			fp.Product.FreightFree = true
			fp.Product.ChinaFreightFen = 0
		default:
			fp.Product.ChinaFreightFen = int64(fr.Freight.Fen())
		}
	}

	if err := im.DB.UpsertProduct(ctx, fp, raw); err != nil {
		return err
	}
	im.upsertCategories(ctx, detail)
	return im.RepriceOffer(ctx, offerID)
}

// ImportSearch runs a keyword search against the gateway and imports every offer
// it returns, which is how the admin fills the catalogue.
func (im *Importer) ImportSearch(ctx context.Context, keyword string, pages, pageSize int) (int, error) {
	if pages <= 0 {
		pages = 1
	}
	if pageSize <= 0 || pageSize > 50 {
		pageSize = 20
	}

	var ids []int64
	for page := 1; page <= pages; page++ {
		res, err := im.Cli.SearchOffers(ctx, ali.OfferQuery{
			Keyword:   keyword,
			BeginPage: page,
			PageSize:  pageSize,
			Country:   im.Country,
		})
		if err != nil {
			return 0, fmt.Errorf("search page %d: %w", page, err)
		}
		for _, it := range res.Data {
			ids = append(ids, int64(it.OfferID))
		}
		if len(res.Data) < pageSize {
			break
		}
	}
	return im.ImportOffers(ctx, ids)
}

// warehouseAddr is the parsed China consolidation address, which supplies the
// district codes the freight estimate needs.
type warehouseAddr struct {
	ProvinceCode string
	CityCode     string
	DistrictCode string
}

func (im *Importer) warehouse(ctx context.Context) (warehouseAddr, bool) {
	settings, err := im.DB.Settings(ctx)
	if err != nil {
		return warehouseAddr{}, false
	}
	var a struct {
		DistrictCode string `json:"districtCode"`
	}
	if err := json.Unmarshal([]byte(settings["warehouse_address"]), &a); err != nil || len(a.DistrictCode) < 6 {
		return warehouseAddr{}, false
	}
	// 330108 is district 330108, city 330100, province 330000.
	return warehouseAddr{
		ProvinceCode: a.DistrictCode[:2] + "0000",
		CityCode:     a.DistrictCode[:4] + "00",
		DistrictCode: a.DistrictCode,
	}, true
}

// mapDetail converts a gateway product into our row shape. Everything the
// gateway sent is also kept verbatim in raw, so a field we did not model is
// never lost and a later migration can backfill from it.
func mapDetail(d ali.OfferDetail) (store.FullProduct, []byte) {
	p := store.Product{
		OfferID:          int64(d.OfferID),
		SellerOpenID:     d.SellerOpenID,
		CompanyName:      d.CompanyName,
		Subject:          d.Subject,
		SubjectTrans:     firstNonEmpty(d.SubjectTrans, d.Subject),
		DescriptionTrans: firstNonEmpty(d.DescriptionTrans, d.Description),
		WhiteImage:       d.ProductImage.WhiteImage,
		CategoryID:       int64(d.CategoryID),
		TopCategoryID:    int64(d.TopCategoryID),
		SecondCategoryID: int64(d.SecondCategoryID),
		ThirdCategoryID:  int64(d.ThirdCategoryID),
		CategoryName:     d.CategoryName,
		Status:           strings.TrimSpace(d.Status),
		MinOrderQuantity: int32(max64(d.MinOrderQuantity, 1)),
		BatchNumber:      int32(d.BatchNumber),
		QuoteType:        int16(d.ProductSaleInfo.QuoteType),
		UnitTrans:        firstNonEmpty(d.ProductSaleInfo.UnitInfo.TransUnit, d.ProductSaleInfo.UnitInfo.Unit),
		MixGeneral:       bool(d.SellerMixSetting.GeneralHunpi),
		MixAmountFen:     int64(d.SellerMixSetting.MixAmount),
		MixNumber:        int32(d.SellerMixSetting.MixNumber),
		AmountOnSale:     int32(d.ProductSaleInfo.AmountOnSale),
		TradeScore:       d.TradeScore,
		Identities:       d.OfferIdentities,
		Visible:          true,
		SyncedAt:         time.Now(),
	}
	if p.Status == "" {
		p.Status = "published"
	}
	if n, err := strconv.Atoi(strings.TrimSpace(d.SoldOut)); err == nil {
		p.MonthSold = int32(n)
	}

	images := d.ProductImage.Images
	if len(d.ProductImageTrans.Images) > 0 {
		images = d.ProductImageTrans.Images
	}
	p.Images = images
	if p.WhiteImage == "" && len(images) > 0 {
		p.WhiteImage = images[0]
	}
	if p.Identities == nil {
		p.Identities = []string{}
	}
	if p.Images == nil {
		p.Images = []string{}
	}

	// Weight arrives in kilograms and dimensions in centimetres; we store grams and
	// millimetres so the international rate card stays integral. Decimal.Milli
	// multiplies by a thousand without touching a float, so kilograms land as
	// grams directly and centimetres need one more division by a hundred.
	p.WeightG = int32(d.ProductShippingInfo.Weight.Milli())
	p.LengthMM = int32(d.ProductShippingInfo.Length.Milli() / 100)
	p.WidthMM = int32(d.ProductShippingInfo.Width.Milli() / 100)
	p.HeightMM = int32(d.ProductShippingInfo.Height.Milli() / 100)

	var keywords []string
	keywords = append(keywords, d.CategoryName)
	for _, a := range d.ProductAttribute {
		keywords = append(keywords, firstNonEmpty(a.ValueTrans, a.Value), firstNonEmpty(a.AttributeNameTrans, a.AttributeName))
	}
	keywords = append(keywords, d.SellingPoint...)
	p.Keywords = strings.Join(compact(keywords), " ")

	fp := store.FullProduct{Product: p}

	for _, s := range d.ProductSkuInfos {
		attrs := make([]store.SKUAttr, 0, len(s.SkuAttributes))
		labels := make([]string, 0, len(s.SkuAttributes))
		image := ""
		for _, a := range s.SkuAttributes {
			attrs = append(attrs, store.SKUAttr{
				Name:       a.AttributeName,
				NameTrans:  firstNonEmpty(a.AttributeNameTrans, a.AttributeName),
				Value:      a.Value,
				ValueTrans: firstNonEmpty(a.ValueTrans, a.Value),
				Image:      a.SkuImageURL,
			})
			labels = append(labels, firstNonEmpty(a.ValueTrans, a.Value))
			if image == "" {
				image = firstNonEmpty(a.SkuImageURLTrans, a.SkuImageURL)
			}
		}
		blob, _ := json.Marshal(attrs)
		fp.SKUs = append(fp.SKUs, store.SKU{
			SkuID:         int64(s.SkuID),
			OfferID:       p.OfferID,
			SpecID:        s.SpecID,
			Attrs:         blob,
			Label:         strings.Join(labels, " / "),
			PriceFen:      int64(s.Price.Fen()),
			PromoPriceFen: int64(s.PromotionPrice.Fen()),
			AmountOnSale:  int32(s.AmountOnSale),
			CargoNumber:   s.CargoNumber,
			ImageURL:      image,
			WeightG:       skuWeight(d, s.SpecID, p.WeightG),
		})
	}

	for _, t := range d.ProductSaleInfo.PriceRangeList {
		fp.Tiers = append(fp.Tiers, store.Tier{
			OfferID:       p.OfferID,
			StartQuantity: int32(t.StartQuantity),
			PriceFen:      int64(t.Price.Fen()),
			PromoPriceFen: int64(t.PromotionPrice.Fen()),
		})
	}

	fp.Product.PriceMinFen, fp.Product.PriceMaxFen = priceBounds(fp)
	if fp.Product.AmountOnSale == 0 {
		var total int64
		for _, s := range fp.SKUs {
			total += int64(s.AmountOnSale)
		}
		fp.Product.AmountOnSale = int32(total)
	}

	raw, _ := json.Marshal(d)
	return fp, raw
}

// skuWeight prefers a per-SKU weight when the gateway supplied one.
//
// Note the unit trap in the documentation: skuShippingInfoList carries weight in
// grams while skuShippingDetails carries it in kilograms. This reads the former,
// so the value is already grams.
func skuWeight(d ali.OfferDetail, specID string, fallback int32) int32 {
	for _, s := range d.ProductShippingInfo.SkuShippingInfoList {
		if s.SpecID == specID && s.Weight.Int64() > 0 {
			return int32(s.Weight.Int64())
		}
	}
	return fallback
}

func priceBounds(fp store.FullProduct) (min, max int64) {
	consider := func(v int64) {
		if v <= 0 {
			return
		}
		if min == 0 || v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	for _, s := range fp.SKUs {
		consider(s.PriceFen)
		consider(s.PromoPriceFen)
	}
	for _, t := range fp.Tiers {
		consider(t.PriceFen)
		consider(t.PromoPriceFen)
	}
	return min, max
}

// upsertCategories records the three category levels an offer names, so the
// storefront rail has something to group by without walking the whole 1688 tree.
func (im *Importer) upsertCategories(ctx context.Context, d ali.OfferDetail) {
	type node struct {
		id     int64
		parent int64
		level  int16
		leaf   bool
	}
	nodes := []node{
		{int64(d.TopCategoryID), 0, 1, false},
		{int64(d.SecondCategoryID), int64(d.TopCategoryID), 2, false},
		{int64(d.ThirdCategoryID), int64(d.SecondCategoryID), 3, true},
	}
	for _, n := range nodes {
		if n.id == 0 {
			continue
		}
		c := store.Category{ID: n.id, ParentID: n.parent, Level: n.level, IsLeaf: n.leaf, Visible: true}
		if n.level == 3 || (n.level == 1 && d.SecondCategoryID == 0) {
			c.Name = d.CategoryName
			c.NameTrans = d.CategoryName
		}
		if c.Name == "" {
			c.Name = "Category " + strconv.FormatInt(n.id, 10)
			c.NameTrans = c.Name
		}
		if err := im.DB.UpsertCategory(ctx, c); err != nil {
			im.Log.Debug("category upsert failed", "id", n.id, "err", err)
		}
	}
	// Give the top-level node the offer's own category name when we have nothing
	// better, so the rail is readable rather than numeric.
	if d.TopCategoryID != 0 && d.CategoryName != "" {
		_ = im.DB.UpsertCategory(ctx, store.Category{
			ID: int64(d.TopCategoryID), ParentID: 0, Level: 1,
			Name: d.CategoryName, NameTrans: d.CategoryName, Visible: true,
		})
	}
}

// RepriceOffer recomputes the denormalised grid price for one product. It runs
// after an import and whenever a fee rule or the exchange rate changes; without
// it the grid and the product page would disagree.
func (im *Importer) RepriceOffer(ctx context.Context, offerID int64) error {
	fp, err := im.DB.GetProduct(ctx, offerID)
	if err != nil {
		return err
	}
	settings, err := im.DB.Settings(ctx)
	if err != nil {
		return err
	}
	rules, err := im.DB.FeeRules(ctx)
	if err != nil {
		return err
	}

	sat := CheapestSatang(fp, rules, settings)
	return im.DB.SetSellMinSatang(ctx, offerID, int64(sat))
}

// RepriceAll recomputes every product's grid price.
func (im *Importer) RepriceAll(ctx context.Context) (int, error) {
	ids, err := im.DB.AllOfferIDs(ctx)
	if err != nil {
		return 0, err
	}
	settings, err := im.DB.Settings(ctx)
	if err != nil {
		return 0, err
	}
	rules, err := im.DB.FeeRules(ctx)
	if err != nil {
		return 0, err
	}

	n := 0
	for _, id := range ids {
		if err := ctx.Err(); err != nil {
			return n, err
		}
		fp, err := im.DB.GetProduct(ctx, id)
		if err != nil {
			continue
		}
		if err := im.DB.SetSellMinSatang(ctx, id, int64(CheapestSatang(fp, rules, settings))); err != nil {
			im.Log.Warn("reprice failed", "offerId", id, "err", err)
			continue
		}
		n++
	}
	return n, nil
}

// CheapestSatang is the lowest price a shopper could pay for one unit of this
// product, which is what the grid shows. It quotes the cheapest SKU at the
// product's minimum order quantity through the same engine checkout uses, so the
// grid can never drift from the real price.
func CheapestSatang(fp store.FullProduct, rules []store.FeeRule, settings map[string]string) ali.Satang {
	if len(fp.SKUs) == 0 {
		return 0
	}
	best := fp.SKUs[0]
	for _, s := range fp.SKUs[1:] {
		if s.PriceFen > 0 && (best.PriceFen == 0 || s.PriceFen < best.PriceFen) {
			best = s
		}
	}

	qty := int(fp.Product.MinOrderQuantity)
	if qty < 1 {
		qty = 1
	}
	line := pricing.Line{
		Product:  ToPricingProduct(fp),
		SKU:      ToPricingSKU(best),
		Quantity: qty,
	}
	res := pricing.Quote([]pricing.Line{line}, map[int64]int{fp.Product.OfferID: qty},
		pricing.Rules(ToPricingRules(rules), time.Now()), ToPricingSettings(settings))
	if len(res.Lines) == 0 {
		return 0
	}
	return res.Lines[0].UnitSatang
}

// --------------------------------------------------------------- conversions --

// ToPricingProduct converts a stored product into the pricing engine's input.
func ToPricingProduct(fp store.FullProduct) pricing.Product {
	p := pricing.Product{
		OfferID:         fp.Product.OfferID,
		SellerOpenID:    fp.Product.SellerOpenID,
		Status:          fp.Product.Status,
		QuoteType:       int(fp.Product.QuoteType),
		MinOrderQty:     int(fp.Product.MinOrderQuantity),
		BatchNumber:     int(fp.Product.BatchNumber),
		MixGeneral:      fp.Product.MixGeneral,
		MixAmountFen:    ali.Fen(fp.Product.MixAmountFen),
		MixNumber:       int(fp.Product.MixNumber),
		ChinaFreightFen: ali.Fen(fp.Product.ChinaFreightFen),
		FreightFree:     fp.Product.FreightFree,
		WeightG:         int(fp.Product.WeightG),
		CategoryID:      fp.Product.TopCategoryID,
	}
	for _, t := range fp.Tiers {
		p.Tiers = append(p.Tiers, pricing.Tier{
			StartQuantity: int(t.StartQuantity),
			PriceFen:      ali.Fen(t.PriceFen),
			PromoPriceFen: ali.Fen(t.PromoPriceFen),
		})
	}
	return p
}

// ToPricingSKU converts a stored SKU into the pricing engine's input.
func ToPricingSKU(s store.SKU) pricing.SKU {
	return pricing.SKU{
		SkuID:         s.SkuID,
		SpecID:        s.SpecID,
		PriceFen:      ali.Fen(s.PriceFen),
		PromoPriceFen: ali.Fen(s.PromoPriceFen),
		Stock:         int(s.AmountOnSale),
		WeightG:       int(s.WeightG),
	}
}

// ToPricingRules converts stored fee rules into the engine's form.
func ToPricingRules(rules []store.FeeRule) []pricing.FeeRule {
	out := make([]pricing.FeeRule, 0, len(rules))
	for _, r := range rules {
		out = append(out, pricing.FeeRule{
			ID:          r.ID,
			Scope:       r.Scope,
			ScopeValue:  r.ScopeValue,
			FeeBps:      int(r.FeeBps),
			FeeFixedFen: ali.Fen(r.FeeFixedFen),
			MinFeeFen:   ali.Fen(r.MinFeeFen),
			Priority:    int(r.Priority),
			From:        r.EffectiveFrom,
			To:          r.EffectiveTo,
		})
	}
	return out
}

// ToPricingSettings reads the engine's knobs out of the settings table.
func ToPricingSettings(s map[string]string) pricing.Settings {
	return pricing.Settings{
		FXPpm:               atoi64(s["fx_thb_per_cny_ppm"], 4_900_000),
		RoundingStepSatang:  atoi64(s["rounding_step_satang"], 100),
		RoundingMode:        orDefault(s["rounding_mode"], "up"),
		IntlRateSatangPerKg: atoi64(s["intl_rate_satang_per_kg"], 18_000),
		IntlMinSatang:       atoi64(s["intl_min_satang"], 5_000),
		UsePromoPrices:      isTrue(s["use_promo_prices"]),
	}
}

func isTrue(v string) bool {
	switch strings.ToLower(strings.TrimSpace(v)) {
	case "1", "true", "yes", "on":
		return true
	}
	return false
}

// ------------------------------------------------------------------- helpers --

func firstNonEmpty(vals ...string) string {
	for _, v := range vals {
		if strings.TrimSpace(v) != "" {
			return v
		}
	}
	return ""
}

func compact(in []string) []string {
	out := in[:0]
	seen := map[string]bool{}
	for _, v := range in {
		v = strings.TrimSpace(v)
		if v == "" || seen[v] {
			continue
		}
		seen[v] = true
		out = append(out, v)
	}
	return out
}

func max64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func atoi64(s string, def int64) int64 {
	v, err := strconv.ParseInt(strings.TrimSpace(s), 10, 64)
	if err != nil {
		return def
	}
	return v
}

func orDefault(s, def string) string {
	if strings.TrimSpace(s) == "" {
		return def
	}
	return s
}
