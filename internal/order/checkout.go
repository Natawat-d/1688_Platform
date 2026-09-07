package order

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"sort"
	"strconv"
	"time"

	"marketplace/internal/ali"
	"marketplace/internal/catalog"
	"marketplace/internal/pricing"
	"marketplace/internal/store"
)

// MaxSKUsPerOrder is the documented ceiling for one createCrossOrder call: at
// most fifty SKUs, all from the same supplier. A group larger than this is split
// at checkout, so the shopper sees the real parcel count before paying rather
// than discovering it afterwards.
const MaxSKUsPerOrder = 50

// Service carries out checkout, relay and tracking.
type Service struct {
	DB  *store.DB
	Cli *ali.Client
	Log *slog.Logger

	// PollInterval is how often an open supplier order is refreshed from the
	// gateway. Push messages carry most updates; this is the safety net under
	// them, and it is also what completes an order when pushes are off entirely.
	// Five minutes is right for production; a local demo sets it to seconds.
	PollInterval time.Duration
}

// NewService builds the order service.
func NewService(db *store.DB, cli *ali.Client, log *slog.Logger) *Service {
	return &Service{DB: db, Cli: cli, Log: log, PollInterval: 5 * time.Minute}
}

// QuotedLine is one cart line with everything needed to price and display it.
type QuotedLine struct {
	Item    store.CartItem
	Product store.Product
	SKU     store.SKU
	Quote   pricing.LineQuote
}

// Group is one supplier's worth of a cart: exactly what becomes one 1688 order,
// or several when it exceeds the fifty-SKU ceiling.
type Group struct {
	SellerOpenID   string
	SellerName     string
	Parcels        int
	Lines          []QuotedLine
	Issues         []pricing.Issue
	SubtotalSatang ali.Satang
}

// CartQuote is a priced cart.
type CartQuote struct {
	CartID       string
	Groups       []Group
	Totals       pricing.Totals
	Issues       []pricing.Issue
	Checkoutable bool
	FxPpm        int64
}

// QuoteCart prices a cart through the same engine checkout uses, so what the
// shopper sees and what they are charged can never diverge.
func (s *Service) QuoteCart(ctx context.Context, cartID string) (CartQuote, error) {
	out := CartQuote{CartID: cartID}

	items, err := s.DB.CartItems(ctx, cartID)
	if err != nil {
		return out, err
	}
	if len(items) == 0 {
		return out, nil
	}

	offerIDs := make([]int64, 0, len(items))
	for _, it := range items {
		offerIDs = append(offerIDs, it.OfferID)
	}
	products, err := s.DB.GetProducts(ctx, offerIDs)
	if err != nil {
		return out, err
	}
	settings, err := s.DB.Settings(ctx)
	if err != nil {
		return out, err
	}
	rules, err := s.DB.FeeRules(ctx)
	if err != nil {
		return out, err
	}

	ps := catalog.ToPricingSettings(settings)
	out.FxPpm = ps.FXPpm
	ruleFor := pricing.Rules(catalog.ToPricingRules(rules), time.Now())

	// Tier prices are quoted on the total quantity of an offer across the whole
	// cart, so two colours of the same product share a tier. The engine takes
	// these totals as an argument rather than deriving them, which keeps it pure.
	offerTotals := map[int64]int{}
	var lines []pricing.Line
	index := make([]struct {
		item store.CartItem
		fp   store.FullProduct
		sku  store.SKU
	}, 0, len(items))

	for _, it := range items {
		fp, ok := products[it.OfferID]
		if !ok {
			continue
		}
		var sku store.SKU
		found := false
		for _, sk := range fp.SKUs {
			if sk.SkuID == it.SkuID {
				sku, found = sk, true
				break
			}
		}
		if !found {
			continue
		}
		offerTotals[it.OfferID] += int(it.Quantity)
		lines = append(lines, pricing.Line{
			Product:  catalog.ToPricingProduct(fp),
			SKU:      catalog.ToPricingSKU(sku),
			Quantity: int(it.Quantity),
		})
		index = append(index, struct {
			item store.CartItem
			fp   store.FullProduct
			sku  store.SKU
		}{it, fp, sku})
	}
	if len(lines) == 0 {
		return out, nil
	}

	res := pricing.Quote(lines, offerTotals, ruleFor, ps)
	out.Totals = res.Totals
	out.Issues = res.Issues

	// Group by supplier, because each supplier becomes its own 1688 order.
	bySeller := map[string]*Group{}
	var order []string
	for i, lq := range res.Lines {
		if i >= len(index) {
			break
		}
		e := index[i]
		seller := e.fp.Product.SellerOpenID
		g, ok := bySeller[seller]
		if !ok {
			g = &Group{SellerOpenID: seller, SellerName: e.fp.Product.CompanyName}
			bySeller[seller] = g
			order = append(order, seller)
		}
		g.Lines = append(g.Lines, QuotedLine{Item: e.item, Product: e.fp.Product, SKU: e.sku, Quote: lq})
		g.SubtotalSatang += lq.LineSatang
	}

	// Attach each issue to the group it belongs to, so the cart can show the
	// problem next to the line that caused it.
	sellerOfOffer := map[int64]string{}
	for _, e := range index {
		sellerOfOffer[e.fp.Product.OfferID] = e.fp.Product.SellerOpenID
	}
	for _, is := range res.Issues {
		if g, ok := bySeller[sellerOfOffer[is.OfferID]]; ok {
			g.Issues = append(g.Issues, is)
		}
	}

	for _, seller := range order {
		g := bySeller[seller]
		g.Parcels = (len(g.Lines) + MaxSKUsPerOrder - 1) / MaxSKUsPerOrder
		if g.Parcels == 0 {
			g.Parcels = 1
		}
		out.Groups = append(out.Groups, *g)
	}
	sort.Slice(out.Groups, func(i, j int) bool { return out.Groups[i].SellerName < out.Groups[j].SellerName })

	out.Checkoutable = len(out.Issues) == 0 && len(out.Groups) > 0
	return out, nil
}

// CheckoutInput is what the shopper supplies.
type CheckoutInput struct {
	Email    string
	Name     string
	Phone    string
	Address  map[string]string
	CartID   string
	Currency string
}

// ErrNotCheckoutable is returned when the cart still has blocking issues.
var ErrNotCheckoutable = errors.New("order: cart has unresolved issues")

// CheckoutResult carries the created order and anything the gateway objected to.
type CheckoutResult struct {
	Order  store.Order
	Issues []pricing.Issue
}

// Checkout validates the cart against 1688 and, only if it passes, creates our
// order.
//
// The preview call happens here, synchronously, before the customer is charged.
// That is deliberate: an empty tradeModeNameList means the offer cannot be
// ordered through the API at all, and discovering that after taking someone's
// money is the worst possible time to find out.
func (s *Service) Checkout(ctx context.Context, in CheckoutInput) (CheckoutResult, error) {
	var res CheckoutResult

	q, err := s.QuoteCart(ctx, in.CartID)
	if err != nil {
		return res, err
	}
	if !q.Checkoutable {
		res.Issues = q.Issues
		return res, ErrNotCheckoutable
	}

	settings, err := s.DB.Settings(ctx)
	if err != nil {
		return res, err
	}
	warehouse, err := warehouseAddress(settings)
	if err != nil {
		return res, err
	}
	toleranceBps := atoi(settings["price_tolerance_bps"], 500)

	// Split each supplier group at the documented fifty-SKU ceiling, then ask the
	// gateway to validate every resulting order.
	type plan struct {
		group     Group
		seq       int32
		lines     []QuotedLine
		tradeType string
		flow      string
		sumFen    ali.Fen
		postFen   ali.Fen
	}
	var plans []plan
	seq := int32(0)
	for _, g := range q.Groups {
		for start := 0; start < len(g.Lines); start += MaxSKUsPerOrder {
			end := min(start+MaxSKUsPerOrder, len(g.Lines))
			seq++
			plans = append(plans, plan{group: g, seq: seq, lines: g.Lines[start:end]})
		}
	}

	for i := range plans {
		p := &plans[i]
		cargo := make([]ali.Cargo, 0, len(p.lines))
		// ourFen is goods only, no freight: it is compared against the preview's
		// sumPaymentNoCarriage below, and that figure excludes carriage by name.
		var ourFen ali.Fen
		for _, l := range p.lines {
			cargo = append(cargo, ali.NewCargo(ali.ID(l.Product.OfferID), l.SKU.SpecID, int64(l.Item.Quantity)))
			ourFen += l.Quote.BaseFen * ali.Fen(l.Item.Quantity)
		}

		preview, err := s.Cli.PreviewOrder(ctx, ali.PreviewRequest{
			Address:    warehouse,
			Cargo:      cargo,
			IsvBizType: "cross",
		})
		if err != nil {
			// A documented refusal is the shopper's problem to fix, so it is
			// surfaced in their own vocabulary rather than as a server error.
			res.Issues = append(res.Issues, pricing.Issue{
				Code:    ali.Code(err),
				Message: previewMessage(err),
				OfferID: p.lines[0].Product.OfferID,
			})
			continue
		}
		if len(preview.Results) == 0 {
			res.Issues = append(res.Issues, pricing.Issue{
				Code:    "UNAVAILABLE",
				Message: "This supplier's items cannot be ordered right now.",
				OfferID: p.lines[0].Product.OfferID,
			})
			continue
		}

		pr := preview.Results[0]
		if len(pr.TradeModeNameList) == 0 {
			res.Issues = append(res.Issues, pricing.Issue{
				Code:    "NO_TRADE_MODE",
				Message: "This supplier does not accept orders through our channel.",
				OfferID: p.lines[0].Product.OfferID,
			})
			continue
		}
		p.tradeType = pr.TradeModeNameList[0]
		p.flow = orDefault(pr.FlowFlag, "general")
		p.sumFen = pr.SumPayment
		p.postFen = pr.SumCarriage

		// Guard against price drift between our snapshot and what 1688 will
		// actually charge. Compare fen against fen: the preview family speaks
		// cents while the order-reading family speaks yuan, and mixing them here
		// would silently pass or fail everything.
		if ourFen > 0 && pr.SumPaymentNoCarriage > 0 {
			diff := int64(pr.SumPaymentNoCarriage - ourFen)
			if diff < 0 {
				diff = -diff
			}
			if diff*10000 > int64(ourFen)*int64(toleranceBps) {
				s.Log.Warn("preview price drifted from snapshot",
					"seller", p.group.SellerOpenID, "oursFen", int64(ourFen),
					"previewFen", int64(pr.SumPaymentNoCarriage), "toleranceBps", toleranceBps)
				res.Issues = append(res.Issues, pricing.Issue{
					Code:    "PRICE_CHANGED",
					Message: "The supplier's price changed. Please review your cart.",
					OfferID: p.lines[0].Product.OfferID,
				})
			}
		}
	}
	if len(res.Issues) > 0 {
		return res, ErrNotCheckoutable
	}

	// Everything checks out: write the order.
	addr, _ := json.Marshal(in.Address)
	token, err := secret()
	if err != nil {
		return res, err
	}

	newOrder := store.NewOrder{
		AccessToken:   token,
		CartID:        in.CartID,
		Email:         in.Email,
		ShipName:      in.Name,
		ShipPhone:     in.Phone,
		ShipAddress:   addr,
		FxPpm:         q.FxPpm,
		GoodsSatang:   int64(q.Totals.GoodsSatang),
		FreightSatang: int64(q.Totals.FreightSatang),
		IntlSatang:    int64(q.Totals.IntlSatang),
		FeeSatang:     int64(q.Totals.FeeSatang),
		TotalSatang:   int64(q.Totals.TotalSatang),
	}
	for _, p := range plans {
		g := store.NewSupplierGroup{
			SellerOpenID: p.group.SellerOpenID,
			SellerName:   p.group.SellerName,
			GroupSeq:     p.seq,
		}
		for _, l := range p.lines {
			snap, _ := json.Marshal(map[string]any{
				"quote":     l.Quote,
				"tradeType": p.tradeType,
				"flow":      p.flow,
				"previewed": map[string]int64{"sumFen": int64(p.sumFen), "postFen": int64(p.postFen)},
				"product":   map[string]any{"status": l.Product.Status, "quoteType": l.Product.QuoteType},
			})
			g.Items = append(g.Items, store.OrderItem{
				OfferID:    l.Product.OfferID,
				SkuID:      l.SKU.SkuID,
				SpecID:     l.SKU.SpecID,
				Title:      l.Product.SubjectTrans,
				ImageURL:   firstImage(l.Product, l.SKU),
				SkuLabel:   l.SKU.Label,
				Quantity:   l.Item.Quantity,
				BaseFen:    int64(l.Quote.BaseFen),
				FreightFen: int64(l.Quote.FreightFen),
				IntlFen:    int64(l.Quote.IntlFen),
				FeeFen:     int64(l.Quote.FeeFen),
				FeeRuleID:  ruleID(l.Quote.FeeRuleID),
				FxPpm:      l.Quote.FXPpm,
				UnitSatang: int64(l.Quote.UnitSatang),
				LineSatang: int64(l.Quote.LineSatang),
				Snapshot:   snap,
			})
		}
		newOrder.Groups = append(newOrder.Groups, g)
	}

	o, err := s.DB.CreateOrder(ctx, newOrder)
	if err != nil {
		return res, err
	}

	// Carry the previewed trade mode onto each supplier order so the relay does
	// not have to preview again in the common path.
	sos, err := s.DB.SupplierOrders(ctx, o.ID)
	if err == nil {
		for _, so := range sos {
			for _, p := range plans {
				if p.seq == so.GroupSeq {
					_ = s.DB.UpdateSupplierOrder(ctx, so.ID, store.SupplierUpdate{
						Flow:          p.flow,
						TradeType:     p.tradeType,
						SumPaymentFen: int64(p.sumFen),
						PostFeeFen:    int64(p.postFen),
					})
				}
			}
		}
	}

	if err := s.DB.ClearCart(ctx, in.CartID); err != nil {
		s.Log.Warn("clear cart failed", "cart", in.CartID, "err", err)
	}

	res.Order = o
	return res, nil
}

// MarkPaid records the customer's payment and queues the relay. Payment here is
// our own provider; the 1688 side is paid separately by the relay worker.
func (s *Service) MarkPaid(ctx context.Context, o store.Order) error {
	if err := s.DB.MarkOrderPaid(ctx, o.ID); err != nil {
		return err
	}
	sos, err := s.DB.SupplierOrders(ctx, o.ID)
	if err != nil {
		return err
	}
	for _, so := range sos {
		if err := s.DB.Enqueue(ctx, store.JobRelayCreate, "so:"+strconv.FormatInt(so.ID, 10),
			map[string]int64{"supplierOrderId": so.ID}, time.Now()); err != nil {
			return err
		}
	}
	return s.DB.SetOrderStatus(ctx, o.ID, OrderProcessing)
}

// ------------------------------------------------------------------ helpers --

func warehouseAddress(settings map[string]string) (ali.Address, error) {
	var a ali.Address
	raw := settings["warehouse_address"]
	if raw == "" {
		return a, errors.New("order: warehouse_address is not configured")
	}
	if err := json.Unmarshal([]byte(raw), &a); err != nil {
		return a, fmt.Errorf("order: warehouse_address is not valid JSON: %w", err)
	}
	if a.DistrictCode == "" && a.AddressID == 0 {
		return a, errors.New("order: warehouse_address needs a districtCode")
	}
	return a, nil
}

// previewMessage turns a gateway refusal into something a shopper can act on.
func previewMessage(err error) string {
	switch ali.Code(err) {
	case ali.ErrNoStock:
		return "One of these items just went out of stock."
	case ali.ErrBelowMOQ, ali.ErrBelowWholesale:
		return "This supplier requires a larger minimum order."
	case ali.ErrMixedBatch:
		return "This supplier's mixed-batch minimum has not been met."
	case ali.ErrZeroPrice:
		return "This item has no valid price right now."
	case ali.ErrNoOnlineTrade:
		return "This item cannot be bought online."
	case ali.ErrSpecNotInOffer:
		return "The selected option is no longer available."
	default:
		return "The supplier rejected this basket. Please review your cart."
	}
}

func firstImage(p store.Product, s store.SKU) string {
	if s.ImageURL != "" {
		return s.ImageURL
	}
	if p.WhiteImage != "" {
		return p.WhiteImage
	}
	if len(p.Images) > 0 {
		return p.Images[0]
	}
	return ""
}

func ruleID(id int64) *int64 {
	if id == 0 {
		return nil
	}
	return &id
}

func secret() (string, error) {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func atoi(s string, def int) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		return def
	}
	return v
}

func orDefault(s, def string) string {
	if s == "" {
		return def
	}
	return s
}
