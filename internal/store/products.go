package store

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/jackc/pgx/v5"
)

// productCols lists exactly the columns that map onto the Product struct.
//
// SELECT * cannot be used here: the table also carries raw (the full gateway
// payload) and the generated search_en tsvector, and the row mapper rejects any
// column it cannot place. Naming the columns keeps that failure at compile-time
// review rather than at runtime.
const productCols = `offer_id, seller_open_id, company_name, subject, subject_trans,
	description_trans, keywords, images, white_image, category_id, top_category_id,
	second_category_id, third_category_id, category_name, status, min_order_quantity,
	batch_number, quote_type, unit_trans, mix_general, mix_amount_fen, mix_number,
	amount_on_sale, month_sold, trade_score, identities, weight_g, length_mm, width_mm,
	height_mm, china_freight_fen, freight_free, price_min_fen, price_max_fen,
	sell_min_satang, visible, synced_at`

// productColsP is the same list qualified for a query that aliases products as p.
const productColsP = `p.offer_id, p.seller_open_id, p.company_name, p.subject, p.subject_trans,
	p.description_trans, p.keywords, p.images, p.white_image, p.category_id, p.top_category_id,
	p.second_category_id, p.third_category_id, p.category_name, p.status, p.min_order_quantity,
	p.batch_number, p.quote_type, p.unit_trans, p.mix_general, p.mix_amount_fen, p.mix_number,
	p.amount_on_sale, p.month_sold, p.trade_score, p.identities, p.weight_g, p.length_mm, p.width_mm,
	p.height_mm, p.china_freight_fen, p.freight_free, p.price_min_fen, p.price_max_fen,
	p.sell_min_satang, p.visible, p.synced_at`

// UpsertProduct writes a product and replaces its SKUs and price tiers in one
// transaction, so the catalogue is never half-updated.
func (db *DB) UpsertProduct(ctx context.Context, fp FullProduct, raw []byte) error {
	tx, err := db.Pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer func() { _ = tx.Rollback(ctx) }()

	p := fp.Product
	if len(raw) == 0 {
		raw = []byte("{}")
	}
	_, err = tx.Exec(ctx, `
		INSERT INTO products (
			offer_id, seller_open_id, company_name, subject, subject_trans, description_trans,
			keywords, images, white_image, category_id, top_category_id, second_category_id,
			third_category_id, category_name, status, min_order_quantity, batch_number, quote_type,
			unit_trans, mix_general, mix_amount_fen, mix_number, amount_on_sale, month_sold,
			trade_score, identities, weight_g, length_mm, width_mm, height_mm, china_freight_fen,
			freight_free, price_min_fen, price_max_fen, sell_min_satang, visible, raw, synced_at)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21,$22,
		        $23,$24,$25,$26,$27,$28,$29,$30,$31,$32,$33,$34,$35,$36,$37, now())
		ON CONFLICT (offer_id) DO UPDATE SET
			seller_open_id=EXCLUDED.seller_open_id, company_name=EXCLUDED.company_name,
			subject=EXCLUDED.subject, subject_trans=EXCLUDED.subject_trans,
			description_trans=EXCLUDED.description_trans, keywords=EXCLUDED.keywords,
			images=EXCLUDED.images, white_image=EXCLUDED.white_image,
			category_id=EXCLUDED.category_id, top_category_id=EXCLUDED.top_category_id,
			second_category_id=EXCLUDED.second_category_id, third_category_id=EXCLUDED.third_category_id,
			category_name=EXCLUDED.category_name, status=EXCLUDED.status,
			min_order_quantity=EXCLUDED.min_order_quantity, batch_number=EXCLUDED.batch_number,
			quote_type=EXCLUDED.quote_type, unit_trans=EXCLUDED.unit_trans,
			mix_general=EXCLUDED.mix_general, mix_amount_fen=EXCLUDED.mix_amount_fen,
			mix_number=EXCLUDED.mix_number, amount_on_sale=EXCLUDED.amount_on_sale,
			month_sold=EXCLUDED.month_sold, trade_score=EXCLUDED.trade_score,
			identities=EXCLUDED.identities, weight_g=EXCLUDED.weight_g,
			length_mm=EXCLUDED.length_mm, width_mm=EXCLUDED.width_mm, height_mm=EXCLUDED.height_mm,
			china_freight_fen=EXCLUDED.china_freight_fen, freight_free=EXCLUDED.freight_free,
			price_min_fen=EXCLUDED.price_min_fen, price_max_fen=EXCLUDED.price_max_fen,
			sell_min_satang=EXCLUDED.sell_min_satang, raw=EXCLUDED.raw, synced_at=now()`,
		p.OfferID, p.SellerOpenID, p.CompanyName, p.Subject, p.SubjectTrans, p.DescriptionTrans,
		p.Keywords, p.Images, p.WhiteImage, p.CategoryID, p.TopCategoryID, p.SecondCategoryID,
		p.ThirdCategoryID, p.CategoryName, p.Status, p.MinOrderQuantity, p.BatchNumber, p.QuoteType,
		p.UnitTrans, p.MixGeneral, p.MixAmountFen, p.MixNumber, p.AmountOnSale, p.MonthSold,
		p.TradeScore, p.Identities, p.WeightG, p.LengthMM, p.WidthMM, p.HeightMM, p.ChinaFreightFen,
		p.FreightFree, p.PriceMinFen, p.PriceMaxFen, p.SellMinSatang, p.Visible, raw)
	if err != nil {
		return fmt.Errorf("upsert product %d: %w", p.OfferID, err)
	}

	if _, err = tx.Exec(ctx, `DELETE FROM product_skus WHERE offer_id=$1`, p.OfferID); err != nil {
		return err
	}
	for _, s := range fp.SKUs {
		attrs := s.Attrs
		if len(attrs) == 0 {
			attrs = []byte("[]")
		}
		if _, err = tx.Exec(ctx, `
			INSERT INTO product_skus (sku_id, offer_id, spec_id, attrs, label, price_fen,
				promo_price_fen, amount_on_sale, cargo_number, image_url, weight_g)
			VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)`,
			s.SkuID, p.OfferID, s.SpecID, attrs, s.Label, s.PriceFen,
			s.PromoPriceFen, s.AmountOnSale, s.CargoNumber, s.ImageURL, s.WeightG); err != nil {
			return fmt.Errorf("upsert sku %d: %w", s.SkuID, err)
		}
	}

	if _, err = tx.Exec(ctx, `DELETE FROM price_tiers WHERE offer_id=$1`, p.OfferID); err != nil {
		return err
	}
	for _, t := range fp.Tiers {
		if _, err = tx.Exec(ctx, `
			INSERT INTO price_tiers (offer_id, start_quantity, price_fen, promo_price_fen)
			VALUES ($1,$2,$3,$4) ON CONFLICT (offer_id, start_quantity) DO UPDATE
			SET price_fen=EXCLUDED.price_fen, promo_price_fen=EXCLUDED.promo_price_fen`,
			p.OfferID, t.StartQuantity, t.PriceFen, t.PromoPriceFen); err != nil {
			return err
		}
	}
	return tx.Commit(ctx)
}

// GetProduct returns one product with its SKUs and tiers.
func (db *DB) GetProduct(ctx context.Context, offerID int64) (FullProduct, error) {
	var fp FullProduct
	rows, err := db.Pool.Query(ctx, `SELECT `+productCols+` FROM products WHERE offer_id=$1`, offerID)
	if err != nil {
		return fp, err
	}
	p, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByNameLax[Product])
	if err != nil {
		return fp, err
	}
	fp.Product = p

	if fp.SKUs, err = db.skusFor(ctx, offerID); err != nil {
		return fp, err
	}
	if fp.Tiers, err = db.tiersFor(ctx, offerID); err != nil {
		return fp, err
	}
	return fp, nil
}

// GetProducts loads several products at once, which is what cart pricing needs.
func (db *DB) GetProducts(ctx context.Context, offerIDs []int64) (map[int64]FullProduct, error) {
	out := map[int64]FullProduct{}
	if len(offerIDs) == 0 {
		return out, nil
	}
	rows, err := db.Pool.Query(ctx, `SELECT `+productCols+` FROM products WHERE offer_id = ANY($1)`, offerIDs)
	if err != nil {
		return nil, err
	}
	ps, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[Product])
	if err != nil {
		return nil, err
	}
	for _, p := range ps {
		out[p.OfferID] = FullProduct{Product: p}
	}

	srows, err := db.Pool.Query(ctx, `SELECT * FROM product_skus WHERE offer_id = ANY($1) ORDER BY sku_id`, offerIDs)
	if err != nil {
		return nil, err
	}
	sks, err := pgx.CollectRows(srows, pgx.RowToStructByNameLax[SKU])
	if err != nil {
		return nil, err
	}
	for _, s := range sks {
		fp := out[s.OfferID]
		fp.SKUs = append(fp.SKUs, s)
		out[s.OfferID] = fp
	}

	trows, err := db.Pool.Query(ctx, `SELECT * FROM price_tiers WHERE offer_id = ANY($1) ORDER BY start_quantity`, offerIDs)
	if err != nil {
		return nil, err
	}
	ts, err := pgx.CollectRows(trows, pgx.RowToStructByNameLax[Tier])
	if err != nil {
		return nil, err
	}
	for _, t := range ts {
		fp := out[t.OfferID]
		fp.Tiers = append(fp.Tiers, t)
		out[t.OfferID] = fp
	}
	return out, nil
}

func (db *DB) skusFor(ctx context.Context, offerID int64) ([]SKU, error) {
	rows, err := db.Pool.Query(ctx, `SELECT * FROM product_skus WHERE offer_id=$1 ORDER BY sku_id`, offerID)
	if err != nil {
		return nil, err
	}
	return pgx.CollectRows(rows, pgx.RowToStructByNameLax[SKU])
}

func (db *DB) tiersFor(ctx context.Context, offerID int64) ([]Tier, error) {
	rows, err := db.Pool.Query(ctx, `SELECT * FROM price_tiers WHERE offer_id=$1 ORDER BY start_quantity`, offerID)
	if err != nil {
		return nil, err
	}
	return pgx.CollectRows(rows, pgx.RowToStructByNameLax[Tier])
}

// SearchQuery is one storefront search. Empty fields mean "no filter".
type SearchQuery struct {
	Q          string
	CategoryID int64
	MinSatang  int64
	MaxSatang  int64
	MaxMOQ     int
	InStock    bool
	Sort       string // relevance | price_asc | price_desc | sales | newest
	Page       int
	Size       int
	IncludeAll bool // admin: ignore visibility and published status
}

// SearchResult is one page of products.
type SearchResult struct {
	Total int
	Page  int
	Size  int
	Items []Product
}

// hasCJK reports whether s contains Han, Hiragana, Katakana or Hangul. Postgres
// cannot tokenise those with any bundled parser, so such queries take the
// trigram path instead of full-text search.
func hasCJK(s string) bool {
	for _, r := range s {
		if unicode.Is(unicode.Han, r) || unicode.Is(unicode.Hiragana, r) ||
			unicode.Is(unicode.Katakana, r) || unicode.Is(unicode.Hangul, r) {
			return true
		}
	}
	return false
}

// SearchProducts runs one storefront query. English goes through the generated
// tsvector; anything containing CJK goes through trigram similarity on both the
// Chinese and English titles.
func (db *DB) SearchProducts(ctx context.Context, q SearchQuery) (SearchResult, error) {
	if q.Page < 1 {
		q.Page = 1
	}
	if q.Size < 1 || q.Size > 100 {
		q.Size = 24
	}

	var (
		where []string
		args  []any
	)
	// next appends a bind value and returns its placeholder, so conditions can be
	// composed in any order without counting parameters by hand.
	next := func(v any) string {
		args = append(args, v)
		return "$" + strconv.Itoa(len(args))
	}

	if !q.IncludeAll {
		where = append(where, "p.visible", "p.status = 'published'")
	}
	if q.CategoryID > 0 {
		c := next(q.CategoryID)
		where = append(where, "("+c+" IN (p.top_category_id, p.second_category_id, p.third_category_id, p.category_id))")
	}
	if q.MinSatang > 0 {
		where = append(where, "p.sell_min_satang >= "+next(q.MinSatang))
	}
	if q.MaxSatang > 0 {
		where = append(where, "p.sell_min_satang <= "+next(q.MaxSatang))
	}
	if q.MaxMOQ > 0 {
		where = append(where, "p.min_order_quantity <= "+next(q.MaxMOQ))
	}
	if q.InStock {
		where = append(where, "p.amount_on_sale > 0")
	}

	orderBy := "p.month_sold DESC, p.offer_id"
	term := strings.TrimSpace(q.Q)
	if term != "" {
		if hasCJK(term) {
			t := next(term)
			where = append(where, "(p.subject ILIKE '%'||"+t+"||'%' OR p.subject_trans ILIKE '%'||"+t+"||'%')")
			orderBy = "GREATEST(similarity(p.subject, " + t + "), similarity(p.subject_trans, " + t + ")) DESC, p.month_sold DESC"
		} else {
			t := next(term)
			where = append(where, "p.search_en @@ websearch_to_tsquery('english', "+t+")")
			orderBy = "ts_rank(p.search_en, websearch_to_tsquery('english', " + t + ")) DESC, p.month_sold DESC"
		}
	}

	switch q.Sort {
	case "price_asc":
		orderBy = "p.sell_min_satang ASC, p.offer_id"
	case "price_desc":
		orderBy = "p.sell_min_satang DESC, p.offer_id"
	case "sales":
		orderBy = "p.month_sold DESC, p.offer_id"
	case "newest":
		orderBy = "p.synced_at DESC, p.offer_id"
	}

	clause := ""
	if len(where) > 0 {
		clause = " WHERE " + strings.Join(where, " AND ")
	}

	var total int
	if err := db.Pool.QueryRow(ctx, `SELECT count(*) FROM products p`+clause, args...).Scan(&total); err != nil {
		return SearchResult{}, fmt.Errorf("search count: %w", err)
	}

	limit := next(q.Size)
	offset := next((q.Page - 1) * q.Size)
	sql := `SELECT ` + productColsP + ` FROM products p` + clause + ` ORDER BY ` + orderBy + ` LIMIT ` + limit + ` OFFSET ` + offset

	rows, err := db.Pool.Query(ctx, sql, args...)
	if err != nil {
		return SearchResult{}, fmt.Errorf("search: %w", err)
	}
	items, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[Product])
	if err != nil {
		return SearchResult{}, err
	}
	return SearchResult{Total: total, Page: q.Page, Size: q.Size, Items: items}, nil
}

// CategoryCount is one row of the storefront category rail.
type CategoryCount struct {
	ID    int64  `db:"id"`
	Name  string `db:"name"`
	Count int    `db:"count"`
}

// TopCategories returns the top-level categories that actually have visible
// products, with their counts.
func (db *DB) TopCategories(ctx context.Context) ([]CategoryCount, error) {
	rows, err := db.Pool.Query(ctx, `
		SELECT c.id, coalesce(nullif(c.name_trans,''), c.name) AS name, count(p.offer_id)::int AS count
		  FROM categories c
		  JOIN products p ON p.top_category_id = c.id
		 WHERE p.visible AND p.status = 'published'
		 GROUP BY c.id, name
		 ORDER BY count DESC, name`)
	if err != nil {
		return nil, err
	}
	return pgx.CollectRows(rows, pgx.RowToStructByName[CategoryCount])
}

// UpsertCategory writes one category node.
func (db *DB) UpsertCategory(ctx context.Context, c Category) error {
	_, err := db.Pool.Exec(ctx, `
		INSERT INTO categories (id, parent_id, name, name_trans, level, is_leaf, visible)
		VALUES ($1,$2,$3,$4,$5,$6,$7)
		ON CONFLICT (id) DO UPDATE SET parent_id=EXCLUDED.parent_id, name=EXCLUDED.name,
			name_trans=EXCLUDED.name_trans, level=EXCLUDED.level, is_leaf=EXCLUDED.is_leaf`,
		c.ID, c.ParentID, c.Name, c.NameTrans, c.Level, c.IsLeaf, c.Visible)
	return err
}

// SetProductVisible flips an admin visibility override.
func (db *DB) SetProductVisible(ctx context.Context, offerID int64, visible bool) error {
	_, err := db.Pool.Exec(ctx, `UPDATE products SET visible=$2 WHERE offer_id=$1`, offerID, visible)
	return err
}

// SetSellMinSatang stores the denormalised grid price for one product.
func (db *DB) SetSellMinSatang(ctx context.Context, offerID, satang int64) error {
	_, err := db.Pool.Exec(ctx, `UPDATE products SET sell_min_satang=$2 WHERE offer_id=$1`, offerID, satang)
	return err
}

// AllOfferIDs lists every imported offer, used by the reprice job.
func (db *DB) AllOfferIDs(ctx context.Context) ([]int64, error) {
	rows, err := db.Pool.Query(ctx, `SELECT offer_id FROM products ORDER BY offer_id`)
	if err != nil {
		return nil, err
	}
	return pgx.CollectRows(rows, pgx.RowTo[int64])
}

// CountProducts returns the total and published counts, for the admin dashboard.
func (db *DB) CountProducts(ctx context.Context) (total, published int, err error) {
	err = db.Pool.QueryRow(ctx, `
		SELECT count(*)::int,
		       count(*) FILTER (WHERE status='published' AND visible)::int
		  FROM products`).Scan(&total, &published)
	return
}
