package pricing

import (
	"strconv"
	"testing"
	"time"

	"marketplace/internal/ali"
)

// fx is 4.9 THB per CNY in parts per million, the rate every expectation in
// this file was worked out by hand against.
const fx = 4_900_000

// plain is the neutral setting: real FX, no rounding step, no shipping.
var plain = Settings{FXPpm: fx, RoundingStepSatang: 1, RoundingMode: RoundNearest}

// promoOn is plain with promotional prices allowed, which is off by default.
var promoOn = Settings{FXPpm: fx, RoundingStepSatang: 1, RoundingMode: RoundNearest, UsePromoPrices: true}

// ---- fixtures -------------------------------------------------------------

func prod(offer int64, seller string, quoteType int, opts ...func(*Product)) Product {
	p := Product{OfferID: offer, SellerOpenID: seller, Status: StatusPublished, QuoteType: quoteType}
	for _, o := range opts {
		o(&p)
	}
	return p
}

func withTiers(t ...Tier) func(*Product) { return func(p *Product) { p.Tiers = t } }
func withMOQ(n int) func(*Product)       { return func(p *Product) { p.MinOrderQty = n } }
func withBatch(n int) func(*Product)     { return func(p *Product) { p.BatchNumber = n } }
func withStatus(s string) func(*Product) { return func(p *Product) { p.Status = s } }
func withWeight(g int) func(*Product)    { return func(p *Product) { p.WeightG = g } }
func withCategory(id int64) func(*Product) {
	return func(p *Product) { p.CategoryID = id }
}
func withFreight(f ali.Fen) func(*Product) {
	return func(p *Product) { p.ChinaFreightFen = f }
}
func withFreeFreight() func(*Product) {
	return func(p *Product) { p.FreightFree = true }
}
func withMix(number int, amount ali.Fen) func(*Product) {
	return func(p *Product) { p.MixGeneral, p.MixNumber, p.MixAmountFen = true, number, amount }
}

func mkSKU(id int64, price ali.Fen, stock int) SKU {
	return SKU{SkuID: id, SpecID: "spec-" + strconv.FormatInt(id, 10), PriceFen: price, Stock: stock}
}

func mkSKUPromo(id int64, price, promo ali.Fen, stock int) SKU {
	s := mkSKU(id, price, stock)
	s.PromoPriceFen = promo
	return s
}

func row(p Product, s SKU, qty int) Line { return Line{Product: p, SKU: s, Quantity: qty} }

func flat(bps int, fixed, min ali.Fen, id int64) func(Product) *FeeRule {
	r := FeeRule{ID: id, Scope: ScopeGlobal, FeeBps: bps, FeeFixedFen: fixed, MinFeeFen: min}
	return func(Product) *FeeRule { return &r }
}

// wantIssue is an Issue without its prose, so the table stays about codes and
// ids; the exact wording is asserted once in TestIssueMessages.
type wantIssue struct {
	Code    string
	OfferID int64
	SkuID   int64
}

// ---- shared products ------------------------------------------------------

// tiered is quoted by product quantity: 1000 fen a piece, 900 from fifty up.
func tiered(offer int64, seller string, quoteType int, opts ...func(*Product)) Product {
	all := append([]func(*Product){withTiers(
		Tier{StartQuantity: 10, PriceFen: 1000},
		Tier{StartQuantity: 50, PriceFen: 900},
	)}, opts...)
	return prod(offer, seller, quoteType, all...)
}

func TestQuote(t *testing.T) {
	cases := []struct {
		name       string
		lines      []Line
		totals     map[int64]int
		rule       func(Product) *FeeRule
		settings   Settings
		wantLines  []LineQuote
		wantTotals Totals
		wantIssues []wantIssue
	}{
		// -- quoteType 1: the SKU's own price ---------------------------------
		{
			name: "quoteType 1 uses the SKU price and ignores the tiers",
			lines: []Line{row(
				prod(1, "s1", QuoteBySKU, withTiers(Tier{StartQuantity: 1, PriceFen: 100000})),
				mkSKU(11, 1850, 100), 2)},
			totals:   map[int64]int{1: 2},
			settings: plain,
			wantLines: []LineQuote{{
				OfferID: 1, SkuID: 11, Quantity: 2, BaseFen: 1850, FXPpm: fx,
				UnitSatang: 9065, LineSatang: 18130,
			}},
			wantTotals: Totals{GoodsSatang: 18130, TotalSatang: 18130},
		},
		{
			name: "quoteType 1 takes a promo price below the base price",
			lines: []Line{row(prod(1, "s1", QuoteBySKU),
				mkSKUPromo(11, 1850, 1500, 100), 1)},
			totals:   map[int64]int{1: 1},
			settings: promoOn,
			wantLines: []LineQuote{{
				OfferID: 1, SkuID: 11, Quantity: 1, BaseFen: 1500, FXPpm: fx,
				UnitSatang: 7350, LineSatang: 7350,
			}},
			wantTotals: Totals{GoodsSatang: 7350, TotalSatang: 7350},
		},
		{
			name:     "a cheaper promo price is ignored while promotions are off",
			lines:    []Line{row(prod(1, "s1", QuoteBySKU), mkSKUPromo(11, 1850, 1500, 100), 1)},
			totals:   map[int64]int{1: 1},
			settings: plain,
			wantLines: []LineQuote{{
				OfferID: 1, SkuID: 11, Quantity: 1, BaseFen: 1850, FXPpm: fx,
				UnitSatang: 9065, LineSatang: 9065,
			}},
			wantTotals: Totals{GoodsSatang: 9065, TotalSatang: 9065},
		},
		{
			name: "quoteType 1 ignores a promo price above the base price",
			lines: []Line{row(prod(1, "s1", QuoteBySKU),
				mkSKUPromo(11, 1850, 2000, 100), 1)},
			totals:   map[int64]int{1: 1},
			settings: plain,
			wantLines: []LineQuote{{
				OfferID: 1, SkuID: 11, Quantity: 1, BaseFen: 1850, FXPpm: fx,
				UnitSatang: 9065, LineSatang: 9065,
			}},
			wantTotals: Totals{GoodsSatang: 9065, TotalSatang: 9065},
		},
		{
			name: "quoteType 1 ignores a promo price equal to the base price",
			lines: []Line{row(prod(1, "s1", QuoteBySKU),
				mkSKUPromo(11, 1850, 1850, 100), 1)},
			totals:   map[int64]int{1: 1},
			settings: plain,
			wantLines: []LineQuote{{
				OfferID: 1, SkuID: 11, Quantity: 1, BaseFen: 1850, FXPpm: fx,
				UnitSatang: 9065, LineSatang: 9065,
			}},
			wantTotals: Totals{GoodsSatang: 9065, TotalSatang: 9065},
		},

		// -- tier selection ---------------------------------------------------
		{
			name:     "quoteType 0 at the exact tier boundary",
			lines:    []Line{row(tiered(2, "s2", QuoteByProduct), SKU{}, 50)},
			totals:   map[int64]int{2: 50},
			settings: plain,
			wantLines: []LineQuote{{
				OfferID: 2, Quantity: 50, BaseFen: 900, FXPpm: fx,
				UnitSatang: 4410, LineSatang: 220500,
			}},
			wantTotals: Totals{GoodsSatang: 220500, TotalSatang: 220500},
		},
		{
			name:     "quoteType 0 one unit below the tier boundary",
			lines:    []Line{row(tiered(2, "s2", QuoteByProduct), SKU{}, 49)},
			totals:   map[int64]int{2: 49},
			settings: plain,
			wantLines: []LineQuote{{
				OfferID: 2, Quantity: 49, BaseFen: 1000, FXPpm: fx,
				UnitSatang: 4900, LineSatang: 240100,
			}},
			wantTotals: Totals{GoodsSatang: 240100, TotalSatang: 240100},
		},
		{
			name:     "below the first tier the first tier still applies",
			lines:    []Line{row(tiered(2, "s2", QuoteByProduct), SKU{}, 5)},
			totals:   map[int64]int{2: 5},
			settings: plain,
			wantLines: []LineQuote{{
				OfferID: 2, Quantity: 5, BaseFen: 1000, FXPpm: fx,
				UnitSatang: 4900, LineSatang: 24500,
			}},
			wantTotals: Totals{GoodsSatang: 24500, TotalSatang: 24500},
		},
		{
			name: "unsorted tiers pick the same winner",
			lines: []Line{row(prod(2, "s2", QuoteByProduct, withTiers(
				Tier{StartQuantity: 50, PriceFen: 900},
				Tier{StartQuantity: 10, PriceFen: 1000},
			)), SKU{}, 50)},
			totals:   map[int64]int{2: 50},
			settings: plain,
			wantLines: []LineQuote{{
				OfferID: 2, Quantity: 50, BaseFen: 900, FXPpm: fx,
				UnitSatang: 4410, LineSatang: 220500,
			}},
			wantTotals: Totals{GoodsSatang: 220500, TotalSatang: 220500},
		},
		{
			name: "a tier promo price wins when it is cheaper",
			lines: []Line{row(prod(2, "s2", QuoteByProduct,
				withTiers(Tier{StartQuantity: 1, PriceFen: 1000, PromoPriceFen: 800})), SKU{}, 1)},
			totals:   map[int64]int{2: 1},
			settings: promoOn,
			wantLines: []LineQuote{{
				OfferID: 2, Quantity: 1, BaseFen: 800, FXPpm: fx,
				UnitSatang: 3920, LineSatang: 3920,
			}},
			wantTotals: Totals{GoodsSatang: 3920, TotalSatang: 3920},
		},

		// -- quoteType 2: has SKUs, priced by product quantity -----------------
		{
			name: "quoteType 2 prices from the tier, never from the SKU",
			lines: []Line{row(
				prod(3, "s3", QuoteBySKUProduct, withTiers(Tier{StartQuantity: 1, PriceFen: 1200})),
				mkSKU(31, 5000, 100), 3)},
			totals:   map[int64]int{3: 3},
			settings: plain,
			wantLines: []LineQuote{{
				OfferID: 3, SkuID: 31, Quantity: 3, BaseFen: 1200, FXPpm: fx,
				UnitSatang: 5880, LineSatang: 17640,
			}},
			wantTotals: Totals{GoodsSatang: 17640, TotalSatang: 17640},
		},
		{
			name:     "quoteType 2 with no tiers falls back to the SKU price",
			lines:    []Line{row(prod(3, "s3", QuoteBySKUProduct), mkSKU(31, 5000, 100), 1)},
			totals:   map[int64]int{3: 1},
			settings: plain,
			wantLines: []LineQuote{{
				OfferID: 3, SkuID: 31, Quantity: 1, BaseFen: 5000, FXPpm: fx,
				UnitSatang: 24500, LineSatang: 24500,
			}},
			wantTotals: Totals{GoodsSatang: 24500, TotalSatang: 24500},
		},

		// -- rule 2: the offer total, not the line quantity --------------------
		{
			name: "two colours of one offer share the offer's tier",
			lines: []Line{
				row(prod(4, "s4", QuoteBySKUProduct, withTiers(
					Tier{StartQuantity: 1, PriceFen: 1500},
					Tier{StartQuantity: 12, PriceFen: 1000},
				)), mkSKU(41, 9999, 100), 6),
				row(prod(4, "s4", QuoteBySKUProduct, withTiers(
					Tier{StartQuantity: 1, PriceFen: 1500},
					Tier{StartQuantity: 12, PriceFen: 1000},
				)), mkSKU(42, 9999, 100), 6),
			},
			totals:   map[int64]int{4: 12},
			settings: plain,
			wantLines: []LineQuote{
				{OfferID: 4, SkuID: 41, Quantity: 6, BaseFen: 1000, FXPpm: fx, UnitSatang: 4900, LineSatang: 29400},
				{OfferID: 4, SkuID: 42, Quantity: 6, BaseFen: 1000, FXPpm: fx, UnitSatang: 4900, LineSatang: 29400},
			},
			wantTotals: Totals{GoodsSatang: 58800, TotalSatang: 58800},
		},
		{
			name: "six alone does not reach the twelve tier",
			lines: []Line{row(prod(4, "s4", QuoteBySKUProduct, withTiers(
				Tier{StartQuantity: 1, PriceFen: 1500},
				Tier{StartQuantity: 12, PriceFen: 1000},
			)), mkSKU(41, 9999, 100), 6)},
			totals:   map[int64]int{4: 6},
			settings: plain,
			wantLines: []LineQuote{
				{OfferID: 4, SkuID: 41, Quantity: 6, BaseFen: 1500, FXPpm: fx, UnitSatang: 7350, LineSatang: 44100},
			},
			wantTotals: Totals{GoodsSatang: 44100, TotalSatang: 44100},
		},
		{
			name:  "a caller quoting one line of a bigger cart keeps the cart's tier",
			lines: []Line{row(tiered(2, "s2", QuoteByProduct), SKU{}, 2)},
			// The cart holds 100 of this offer; we are only pricing 2 of them.
			totals:   map[int64]int{2: 100},
			settings: plain,
			wantLines: []LineQuote{{
				OfferID: 2, Quantity: 2, BaseFen: 900, FXPpm: fx,
				UnitSatang: 4410, LineSatang: 8820,
			}},
			wantTotals: Totals{GoodsSatang: 8820, TotalSatang: 8820},
		},
		{
			name: "a nil totals map falls back to the quantities on hand",
			lines: []Line{
				row(tiered(2, "s2", QuoteByProduct), SKU{}, 25),
				row(tiered(2, "s2", QuoteByProduct), SKU{}, 25),
			},
			totals:   nil,
			settings: plain,
			wantLines: []LineQuote{
				{OfferID: 2, Quantity: 25, BaseFen: 900, FXPpm: fx, UnitSatang: 4410, LineSatang: 110250},
				{OfferID: 2, Quantity: 25, BaseFen: 900, FXPpm: fx, UnitSatang: 4410, LineSatang: 110250},
			},
			wantTotals: Totals{GoodsSatang: 220500, TotalSatang: 220500},
		},
		{
			name: "only the offers the caller left out are filled in",
			lines: []Line{
				row(tiered(2, "s2", QuoteByProduct), SKU{}, 1),
				row(tiered(5, "s2", QuoteByProduct), SKU{}, 50),
			},
			totals:   map[int64]int{2: 50}, // offer 5 is absent, so its own 50 counts
			settings: plain,
			wantLines: []LineQuote{
				{OfferID: 2, Quantity: 1, BaseFen: 900, FXPpm: fx, UnitSatang: 4410, LineSatang: 4410},
				{OfferID: 5, Quantity: 50, BaseFen: 900, FXPpm: fx, UnitSatang: 4410, LineSatang: 220500},
			},
			wantTotals: Totals{GoodsSatang: 224910, TotalSatang: 224910},
		},

		// -- minimum order quantity and batch multiples -------------------------
		{
			name: "minimum order quantity met",
			lines: []Line{row(prod(6, "s6", QuoteByProduct, withMOQ(5),
				withTiers(Tier{StartQuantity: 1, PriceFen: 1000})), SKU{}, 5)},
			totals:   map[int64]int{6: 5},
			settings: plain,
			wantLines: []LineQuote{
				{OfferID: 6, Quantity: 5, BaseFen: 1000, FXPpm: fx, UnitSatang: 4900, LineSatang: 24500},
			},
			wantTotals: Totals{GoodsSatang: 24500, TotalSatang: 24500},
		},
		{
			name: "minimum order quantity not met",
			lines: []Line{row(prod(6, "s6", QuoteByProduct, withMOQ(5),
				withTiers(Tier{StartQuantity: 1, PriceFen: 1000})), SKU{}, 3)},
			totals:   map[int64]int{6: 3},
			settings: plain,
			wantLines: []LineQuote{
				{OfferID: 6, Quantity: 3, BaseFen: 1000, FXPpm: fx, UnitSatang: 4900, LineSatang: 14700},
			},
			wantTotals: Totals{GoodsSatang: 14700, TotalSatang: 14700},
			wantIssues: []wantIssue{{Code: ali.ErrBelowMOQ, OfferID: 6}},
		},
		{
			name: "the MOQ is reported once however many rows share the offer",
			lines: []Line{
				row(prod(6, "s6", QuoteBySKUProduct, withMOQ(10),
					withTiers(Tier{StartQuantity: 1, PriceFen: 1000})), mkSKU(61, 0, 100), 1),
				row(prod(6, "s6", QuoteBySKUProduct, withMOQ(10),
					withTiers(Tier{StartQuantity: 1, PriceFen: 1000})), mkSKU(62, 0, 100), 1),
			},
			totals:   map[int64]int{6: 2},
			settings: plain,
			wantLines: []LineQuote{
				{OfferID: 6, SkuID: 61, Quantity: 1, BaseFen: 1000, FXPpm: fx, UnitSatang: 4900, LineSatang: 4900},
				{OfferID: 6, SkuID: 62, Quantity: 1, BaseFen: 1000, FXPpm: fx, UnitSatang: 4900, LineSatang: 4900},
			},
			wantTotals: Totals{GoodsSatang: 9800, TotalSatang: 9800},
			wantIssues: []wantIssue{{Code: ali.ErrBelowMOQ, OfferID: 6}},
		},
		{
			name: "batch multiple satisfied",
			lines: []Line{row(prod(7, "s7", QuoteByProduct, withMOQ(200), withBatch(200),
				withTiers(Tier{StartQuantity: 1, PriceFen: 100})), SKU{}, 400)},
			totals:   map[int64]int{7: 400},
			settings: plain,
			wantLines: []LineQuote{
				{OfferID: 7, Quantity: 400, BaseFen: 100, FXPpm: fx, UnitSatang: 490, LineSatang: 196000},
			},
			wantTotals: Totals{GoodsSatang: 196000, TotalSatang: 196000},
		},
		{
			name: "batch multiple broken",
			lines: []Line{row(prod(7, "s7", QuoteByProduct, withMOQ(200), withBatch(200),
				withTiers(Tier{StartQuantity: 1, PriceFen: 100})), SKU{}, 250)},
			totals:   map[int64]int{7: 250},
			settings: plain,
			wantLines: []LineQuote{
				{OfferID: 7, Quantity: 250, BaseFen: 100, FXPpm: fx, UnitSatang: 490, LineSatang: 122500},
			},
			wantTotals: Totals{GoodsSatang: 122500, TotalSatang: 122500},
			wantIssues: []wantIssue{{Code: ali.ErrBelowMOQ, OfferID: 7}},
		},
		{
			name: "MOQ and batch can both fail on one offer",
			lines: []Line{row(prod(7, "s7", QuoteByProduct, withMOQ(200), withBatch(200),
				withTiers(Tier{StartQuantity: 1, PriceFen: 100})), SKU{}, 150)},
			totals:   map[int64]int{7: 150},
			settings: plain,
			wantLines: []LineQuote{
				{OfferID: 7, Quantity: 150, BaseFen: 100, FXPpm: fx, UnitSatang: 490, LineSatang: 73500},
			},
			wantTotals: Totals{GoodsSatang: 73500, TotalSatang: 73500},
			wantIssues: []wantIssue{
				{Code: ali.ErrBelowMOQ, OfferID: 7},
				{Code: ali.ErrBelowMOQ, OfferID: 7},
			},
		},

		// -- mixed batch: quantity OR amount, never both -------------------------
		{
			name: "mixed batch satisfied by quantity alone",
			lines: []Line{
				row(prod(81, "sm", QuoteByProduct, withMix(10, 5000),
					withTiers(Tier{StartQuantity: 1, PriceFen: 100})), SKU{}, 6),
				row(prod(82, "sm", QuoteByProduct, withMix(10, 5000),
					withTiers(Tier{StartQuantity: 1, PriceFen: 100})), SKU{}, 4),
			},
			totals:   map[int64]int{81: 6, 82: 4},
			settings: plain,
			wantLines: []LineQuote{
				{OfferID: 81, Quantity: 6, BaseFen: 100, FXPpm: fx, UnitSatang: 490, LineSatang: 2940},
				{OfferID: 82, Quantity: 4, BaseFen: 100, FXPpm: fx, UnitSatang: 490, LineSatang: 1960},
			},
			wantTotals: Totals{GoodsSatang: 4900, TotalSatang: 4900},
		},
		{
			name: "mixed batch satisfied by amount alone",
			lines: []Line{row(prod(81, "sm", QuoteByProduct, withMix(10, 5000),
				withTiers(Tier{StartQuantity: 1, PriceFen: 6000})), SKU{}, 1)},
			totals:   map[int64]int{81: 1},
			settings: plain,
			wantLines: []LineQuote{
				{OfferID: 81, Quantity: 1, BaseFen: 6000, FXPpm: fx, UnitSatang: 29400, LineSatang: 29400},
			},
			wantTotals: Totals{GoodsSatang: 29400, TotalSatang: 29400},
		},
		{
			name: "mixed batch satisfied by amount exactly at the threshold",
			lines: []Line{row(prod(81, "sm", QuoteByProduct, withMix(10, 5000),
				withTiers(Tier{StartQuantity: 1, PriceFen: 5000})), SKU{}, 1)},
			totals:   map[int64]int{81: 1},
			settings: plain,
			wantLines: []LineQuote{
				{OfferID: 81, Quantity: 1, BaseFen: 5000, FXPpm: fx, UnitSatang: 24500, LineSatang: 24500},
			},
			wantTotals: Totals{GoodsSatang: 24500, TotalSatang: 24500},
		},
		{
			name: "mixed batch failing both quantity and amount",
			lines: []Line{row(prod(81, "sm", QuoteBySKUProduct, withMix(10, 5000),
				withTiers(Tier{StartQuantity: 1, PriceFen: 1000})), mkSKU(811, 9999, 100), 1)},
			totals:   map[int64]int{81: 1},
			settings: plain,
			wantLines: []LineQuote{
				{OfferID: 81, SkuID: 811, Quantity: 1, BaseFen: 1000, FXPpm: fx, UnitSatang: 4900, LineSatang: 4900},
			},
			wantTotals: Totals{GoodsSatang: 4900, TotalSatang: 4900},
			wantIssues: []wantIssue{{Code: ali.ErrMixedBatch, OfferID: 81, SkuID: 811}},
		},
		{
			name: "a supplier with no mixed batch setting is never gated",
			lines: []Line{row(prod(83, "sn", QuoteByProduct,
				withTiers(Tier{StartQuantity: 1, PriceFen: 100})), SKU{}, 1)},
			totals:   map[int64]int{83: 1},
			settings: plain,
			wantLines: []LineQuote{
				{OfferID: 83, Quantity: 1, BaseFen: 100, FXPpm: fx, UnitSatang: 490, LineSatang: 490},
			},
			wantTotals: Totals{GoodsSatang: 490, TotalSatang: 490},
		},
		{
			name: "each supplier's mixed batch is judged on its own basket",
			lines: []Line{
				row(prod(84, "sx", QuoteByProduct, withMix(10, 500000),
					withTiers(Tier{StartQuantity: 1, PriceFen: 100})), SKU{}, 10),
				row(prod(85, "sy", QuoteByProduct, withMix(10, 500000),
					withTiers(Tier{StartQuantity: 1, PriceFen: 100})), SKU{}, 1),
			},
			totals:   map[int64]int{84: 10, 85: 1},
			settings: plain,
			wantLines: []LineQuote{
				{OfferID: 84, Quantity: 10, BaseFen: 100, FXPpm: fx, UnitSatang: 490, LineSatang: 4900},
				{OfferID: 85, Quantity: 1, BaseFen: 100, FXPpm: fx, UnitSatang: 490, LineSatang: 490},
			},
			wantTotals: Totals{GoodsSatang: 5390, TotalSatang: 5390},
			wantIssues: []wantIssue{{Code: ali.ErrMixedBatch, OfferID: 85}},
		},

		// -- availability and stock -------------------------------------------
		{
			name:     "an unpublished product is unavailable but still priced",
			lines:    []Line{row(prod(9, "s9", QuoteBySKU, withStatus("expired")), mkSKU(91, 1850, 100), 1)},
			totals:   map[int64]int{9: 1},
			settings: plain,
			wantLines: []LineQuote{
				{OfferID: 9, SkuID: 91, Quantity: 1, BaseFen: 1850, FXPpm: fx, UnitSatang: 9065, LineSatang: 9065},
			},
			wantTotals: Totals{GoodsSatang: 9065, TotalSatang: 9065},
			wantIssues: []wantIssue{{Code: CodeUnavailable, OfferID: 9, SkuID: 91}},
		},
		{
			name:     "a SKU with no stock at all",
			lines:    []Line{row(prod(9, "s9", QuoteBySKU), mkSKU(91, 1850, 0), 1)},
			totals:   map[int64]int{9: 1},
			settings: plain,
			wantLines: []LineQuote{
				{OfferID: 9, SkuID: 91, Quantity: 1, BaseFen: 1850, FXPpm: fx, UnitSatang: 9065, LineSatang: 9065},
			},
			wantTotals: Totals{GoodsSatang: 9065, TotalSatang: 9065},
			wantIssues: []wantIssue{{Code: ali.ErrNoStock, OfferID: 9, SkuID: 91}},
		},
		{
			name:     "more ordered than the SKU has left",
			lines:    []Line{row(prod(9, "s9", QuoteBySKU), mkSKU(91, 1850, 3), 5)},
			totals:   map[int64]int{9: 5},
			settings: plain,
			wantLines: []LineQuote{
				{OfferID: 9, SkuID: 91, Quantity: 5, BaseFen: 1850, FXPpm: fx, UnitSatang: 9065, LineSatang: 45325},
			},
			wantTotals: Totals{GoodsSatang: 45325, TotalSatang: 45325},
			wantIssues: []wantIssue{{Code: ali.ErrNoStock, OfferID: 9, SkuID: 91}},
		},
		{
			name:     "exactly the last unit in stock is fine",
			lines:    []Line{row(prod(9, "s9", QuoteBySKU), mkSKU(91, 1850, 5), 5)},
			totals:   map[int64]int{9: 5},
			settings: plain,
			wantLines: []LineQuote{
				{OfferID: 9, SkuID: 91, Quantity: 5, BaseFen: 1850, FXPpm: fx, UnitSatang: 9065, LineSatang: 45325},
			},
			wantTotals: Totals{GoodsSatang: 45325, TotalSatang: 45325},
		},
		{
			name:     "a product quoted without a SKU has no stock to check",
			lines:    []Line{row(tiered(2, "s2", QuoteByProduct), SKU{}, 50)},
			totals:   map[int64]int{2: 50},
			settings: plain,
			wantLines: []LineQuote{
				{OfferID: 2, Quantity: 50, BaseFen: 900, FXPpm: fx, UnitSatang: 4410, LineSatang: 220500},
			},
			wantTotals: Totals{GoodsSatang: 220500, TotalSatang: 220500},
		},
		{
			name:     "a price of zero is refused",
			lines:    []Line{row(prod(9, "s9", QuoteBySKU), mkSKU(91, 0, 100), 1)},
			totals:   map[int64]int{9: 1},
			settings: plain,
			wantLines: []LineQuote{
				{OfferID: 9, SkuID: 91, Quantity: 1, BaseFen: 0, FXPpm: fx},
			},
			wantTotals: Totals{},
			wantIssues: []wantIssue{{Code: ali.ErrZeroPrice, OfferID: 9, SkuID: 91}},
		},
		{
			name: "a promo price cannot rescue a base price of zero",
			// The documented rule is promo > 0 AND promo < price, and 0 is not
			// greater than a promo, so a zero base price stays zero.
			lines:    []Line{row(prod(9, "s9", QuoteBySKU), mkSKUPromo(91, 0, 900, 100), 1)},
			totals:   map[int64]int{9: 1},
			settings: plain,
			wantLines: []LineQuote{
				{OfferID: 9, SkuID: 91, Quantity: 1, BaseFen: 0, FXPpm: fx},
			},
			wantTotals: Totals{},
			wantIssues: []wantIssue{{Code: ali.ErrZeroPrice, OfferID: 9, SkuID: 91}},
		},

		// -- freight, international leg and fee ---------------------------------
		{
			name: "the full breakdown: goods, China freight, international, fee",
			lines: []Line{row(prod(10, "s10", QuoteBySKU, withFreight(300), withWeight(500)),
				mkSKU(101, 1850, 100), 2)},
			totals:   map[int64]int{10: 2},
			rule:     flat(1000, 0, 0, 7),
			settings: Settings{FXPpm: fx, RoundingStepSatang: 1, RoundingMode: RoundNearest, IntlRateSatangPerKg: 20000, IntlMinSatang: 5000},
			wantLines: []LineQuote{{
				OfferID: 10, SkuID: 101, Quantity: 2,
				BaseFen: 1850, FreightFen: 300, IntlFen: 2041, FeeFen: 419,
				FeeRuleID: 7, FXPpm: fx, UnitSatang: 22589, LineSatang: 45178,
			}},
			wantTotals: Totals{
				GoodsSatang: 18130, FreightSatang: 2940, IntlSatang: 20002,
				FeeSatang: 4106, TotalSatang: 45178,
			},
		},
		{
			name: "the international minimum binds on a single light unit",
			lines: []Line{row(prod(10, "s10", QuoteBySKU, withFreight(300), withWeight(100)),
				mkSKU(101, 1850, 100), 1)},
			totals:   map[int64]int{10: 1},
			settings: Settings{FXPpm: fx, RoundingStepSatang: 1, RoundingMode: RoundNearest, IntlRateSatangPerKg: 20000, IntlMinSatang: 5000},
			wantLines: []LineQuote{{
				OfferID: 10, SkuID: 101, Quantity: 1,
				BaseFen: 1850, FreightFen: 300, IntlFen: 1020,
				FXPpm: fx, UnitSatang: 15533, LineSatang: 15533,
			}},
			wantTotals: Totals{
				GoodsSatang: 9065, FreightSatang: 1470, IntlSatang: 4998, TotalSatang: 15533,
			},
		},
		{
			name: "the SKU weight beats the product weight",
			lines: []Line{func() Line {
				s := mkSKU(101, 1850, 100)
				s.WeightG = 500
				return row(prod(10, "s10", QuoteBySKU, withWeight(100)), s, 1)
			}()},
			totals:   map[int64]int{10: 1},
			settings: Settings{FXPpm: fx, RoundingStepSatang: 1, RoundingMode: RoundNearest, IntlRateSatangPerKg: 20000},
			wantLines: []LineQuote{{
				OfferID: 10, SkuID: 101, Quantity: 1,
				BaseFen: 1850, IntlFen: 2041, FXPpm: fx,
				UnitSatang: 19066, LineSatang: 19066,
			}},
			wantTotals: Totals{GoodsSatang: 9065, IntlSatang: 10001, TotalSatang: 19066},
		},
		{
			name: "a free shipping product carries no China freight",
			lines: []Line{row(prod(10, "s10", QuoteBySKU, withFreight(300), withFreeFreight()),
				mkSKU(101, 1850, 100), 1)},
			totals:   map[int64]int{10: 1},
			settings: plain,
			wantLines: []LineQuote{{
				OfferID: 10, SkuID: 101, Quantity: 1, BaseFen: 1850, FXPpm: fx,
				UnitSatang: 9065, LineSatang: 9065,
			}},
			wantTotals: Totals{GoodsSatang: 9065, TotalSatang: 9065},
		},
		{
			name:     "the minimum fee binds",
			lines:    []Line{row(prod(11, "s11", QuoteBySKU), mkSKU(111, 1000, 100), 1)},
			totals:   map[int64]int{11: 1},
			rule:     flat(100, 5, 100, 21),
			settings: plain,
			wantLines: []LineQuote{{
				OfferID: 11, SkuID: 111, Quantity: 1, BaseFen: 1000, FeeFen: 100,
				FeeRuleID: 21, FXPpm: fx, UnitSatang: 5390, LineSatang: 5390,
			}},
			wantTotals: Totals{GoodsSatang: 4900, FeeSatang: 490, TotalSatang: 5390},
		},
		{
			name:     "the computed fee clears the minimum",
			lines:    []Line{row(prod(11, "s11", QuoteBySKU), mkSKU(111, 1000, 100), 1)},
			totals:   map[int64]int{11: 1},
			rule:     flat(5000, 5, 100, 22),
			settings: plain,
			wantLines: []LineQuote{{
				OfferID: 11, SkuID: 111, Quantity: 1, BaseFen: 1000, FeeFen: 505,
				FeeRuleID: 22, FXPpm: fx, UnitSatang: 7375, LineSatang: 7375,
			}},
			wantTotals: Totals{GoodsSatang: 4900, FeeSatang: 2475, TotalSatang: 7375},
		},
		{
			name:     "a nil rule callback quotes at zero fee",
			lines:    []Line{row(prod(11, "s11", QuoteBySKU), mkSKU(111, 1000, 100), 1)},
			totals:   map[int64]int{11: 1},
			rule:     nil,
			settings: plain,
			wantLines: []LineQuote{{
				OfferID: 11, SkuID: 111, Quantity: 1, BaseFen: 1000, FXPpm: fx,
				UnitSatang: 4900, LineSatang: 4900,
			}},
			wantTotals: Totals{GoodsSatang: 4900, TotalSatang: 4900},
		},
		{
			name:     "a rule callback that finds nothing quotes at zero fee",
			lines:    []Line{row(prod(11, "s11", QuoteBySKU), mkSKU(111, 1000, 100), 1)},
			totals:   map[int64]int{11: 1},
			rule:     func(Product) *FeeRule { return nil },
			settings: plain,
			wantLines: []LineQuote{{
				OfferID: 11, SkuID: 111, Quantity: 1, BaseFen: 1000, FXPpm: fx,
				UnitSatang: 4900, LineSatang: 4900,
			}},
			wantTotals: Totals{GoodsSatang: 4900, TotalSatang: 4900},
		},
		{
			name: "Rules wires SelectFeeRule into Quote and records the winner",
			lines: []Line{row(prod(620201390233, "23tsdvcdsjngp3oj4j3i5", QuoteBySKU, withCategory(1031910)),
				mkSKU(111, 1000, 100), 1)},
			totals: map[int64]int{620201390233: 1},
			rule: Rules([]FeeRule{
				{ID: 1, Scope: ScopeGlobal, FeeBps: 2000, From: jan2025},
				{ID: 2, Scope: ScopeSupplier, ScopeValue: "23tsdvcdsjngp3oj4j3i5", FeeBps: 1000, From: jan2025},
			}, jun2026),
			settings: plain,
			wantLines: []LineQuote{{
				OfferID: 620201390233, SkuID: 111, Quantity: 1, BaseFen: 1000, FeeFen: 100,
				FeeRuleID: 2, FXPpm: fx, UnitSatang: 5390, LineSatang: 5390,
			}},
			wantTotals: Totals{GoodsSatang: 4900, FeeSatang: 490, TotalSatang: 5390},
		},

		// -- degenerate input ---------------------------------------------------
		{
			name:       "an empty cart quotes to nothing",
			lines:      nil,
			settings:   plain,
			wantLines:  nil,
			wantTotals: Totals{},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := Quote(tc.lines, tc.totals, tc.rule, tc.settings)

			if len(got.Lines) != len(tc.wantLines) {
				t.Fatalf("got %d line quotes, want %d: %+v", len(got.Lines), len(tc.wantLines), got.Lines)
			}
			for i := range tc.wantLines {
				if got.Lines[i] != tc.wantLines[i] {
					t.Errorf("line %d:\n got %+v\nwant %+v", i, got.Lines[i], tc.wantLines[i])
				}
			}
			if got.Totals != tc.wantTotals {
				t.Errorf("totals:\n got %+v\nwant %+v", got.Totals, tc.wantTotals)
			}
			// The invoice must always balance, whatever the rounding did.
			sum := got.Totals.GoodsSatang + got.Totals.FreightSatang + got.Totals.IntlSatang + got.Totals.FeeSatang
			if sum != got.Totals.TotalSatang {
				t.Errorf("totals do not balance: parts %d, total %d", sum, got.Totals.TotalSatang)
			}
			var lineSum ali.Satang
			for _, l := range got.Lines {
				lineSum += l.LineSatang
			}
			if lineSum != got.Totals.TotalSatang {
				t.Errorf("line sum %d != total %d", lineSum, got.Totals.TotalSatang)
			}

			if len(got.Issues) != len(tc.wantIssues) {
				t.Fatalf("got %d issues, want %d: %+v", len(got.Issues), len(tc.wantIssues), got.Issues)
			}
			for i, w := range tc.wantIssues {
				g := got.Issues[i]
				if g.Code != w.Code || g.OfferID != w.OfferID || g.SkuID != w.SkuID {
					t.Errorf("issue %d: got {%s %d %d}, want {%s %d %d}",
						i, g.Code, g.OfferID, g.SkuID, w.Code, w.OfferID, w.SkuID)
				}
				if g.Message == "" {
					t.Errorf("issue %d (%s) has no message for the shopper", i, g.Code)
				}
			}
			if got.OK() != (len(tc.wantIssues) == 0) {
				t.Errorf("OK() = %v with %d issues", got.OK(), len(got.Issues))
			}
		})
	}
}

// TestRounding pins both snap modes at steps 1, 100 and 1000, and proves the
// unit is rounded before it is multiplied.
func TestRounding(t *testing.T) {
	cases := []struct {
		name     string
		base     ali.Fen
		qty      int
		step     int64
		mode     string
		wantUnit ali.Satang
		wantLine ali.Satang
	}{
		// 898 fen at 4.9 is 4400.2 satang, which lands on 4400.
		{"step 0 does not snap", 898, 3, 0, RoundNearest, 4400, 13200},
		{"step 1 does not snap", 898, 3, 1, RoundNearest, 4400, 13200},
		{"step 1 up does not snap", 898, 3, 1, RoundUp, 4400, 13200},
		{"step 100 already on the step", 898, 3, 100, RoundNearest, 4400, 13200},
		{"step 100 up already on the step", 898, 3, 100, RoundUp, 4400, 13200},
		{"step 1000 nearest rounds down", 898, 3, 1000, RoundNearest, 4000, 12000},
		{"step 1000 up rounds up", 898, 3, 1000, RoundUp, 5000, 15000},
		// 1005 fen at 4.9 is 4924.5 satang, which rounds half up to 4925.
		{"step 1 keeps the half up conversion", 1005, 3, 1, RoundNearest, 4925, 14775},
		{"step 100 nearest rounds down", 1005, 3, 100, RoundNearest, 4900, 14700},
		{"step 100 up rounds up", 1005, 3, 100, RoundUp, 5000, 15000},
		{"step 1000 nearest rounds up", 1005, 3, 1000, RoundNearest, 5000, 15000},
		{"step 1000 up rounds up", 1005, 3, 1000, RoundUp, 5000, 15000},
		{"an unknown mode rounds to nearest", 1005, 3, 100, "sideways", 4900, 14700},
		{"an empty mode rounds to nearest", 1005, 3, 100, "", 4900, 14700},
		{"UP is matched case insensitively", 1005, 3, 100, "UP", 5000, 15000},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := Quote(
				[]Line{row(prod(1, "s1", QuoteBySKU), mkSKU(11, tc.base, 1000), tc.qty)},
				map[int64]int{1: tc.qty},
				nil,
				Settings{FXPpm: fx, RoundingStepSatang: tc.step, RoundingMode: tc.mode},
			)
			if len(got.Lines) != 1 {
				t.Fatalf("got %d lines", len(got.Lines))
			}
			if got.Lines[0].UnitSatang != tc.wantUnit {
				t.Errorf("unit = %d, want %d", got.Lines[0].UnitSatang, tc.wantUnit)
			}
			if got.Lines[0].LineSatang != tc.wantLine {
				t.Errorf("line = %d, want %d", got.Lines[0].LineSatang, tc.wantLine)
			}
		})
	}
}

// TestRoundsUnitThenMultiplies is the reason the previous test exists: a
// shopper who divides the line total by the quantity must get back exactly the
// unit price we printed next to it.
func TestRoundsUnitThenMultiplies(t *testing.T) {
	got := Quote(
		[]Line{row(prod(1, "s1", QuoteBySKU), mkSKU(11, 1005, 1000), 3)},
		map[int64]int{1: 3},
		nil,
		Settings{FXPpm: fx, RoundingStepSatang: 100, RoundingMode: RoundUp},
	)
	l := got.Lines[0]
	if l.UnitSatang != 5000 || l.LineSatang != 15000 {
		t.Fatalf("got unit %d line %d, want 5000 and 15000", l.UnitSatang, l.LineSatang)
	}
	// Rounding the line total instead would have produced 14800, and 14800/3
	// is not a price anyone can display.
	if l.LineSatang != l.UnitSatang*ali.Satang(l.Quantity) {
		t.Fatalf("line %d is not unit %d times %d", l.LineSatang, l.UnitSatang, l.Quantity)
	}
}

// TestQuoteIsPure checks that Quote leaves its inputs alone, which is what
// lets a caller quote the same cart under two fee rules and compare.
func TestQuoteIsPure(t *testing.T) {
	lines := []Line{
		row(tiered(2, "s2", QuoteByProduct), SKU{}, 25),
		row(tiered(2, "s2", QuoteByProduct), SKU{}, 25),
	}
	totals := map[int64]int{2: 50}

	first := Quote(lines, totals, nil, plain)
	if len(totals) != 1 || totals[2] != 50 {
		t.Fatalf("Quote modified the totals map: %v", totals)
	}
	if lines[0].Quantity != 25 || len(lines[0].Product.Tiers) != 2 {
		t.Fatalf("Quote modified its lines: %+v", lines[0])
	}
	second := Quote(lines, totals, nil, plain)
	if first.Totals != second.Totals {
		t.Fatalf("Quote is not deterministic: %+v then %+v", first.Totals, second.Totals)
	}
}

// TestIssueMessages pins the shopper-facing wording, once, so the table above
// can stay about codes and numbers.
func TestIssueMessages(t *testing.T) {
	cases := []struct {
		name  string
		lines []Line
		total map[int64]int
		want  string
	}{
		{
			name:  "unavailable",
			lines: []Line{row(prod(1, "s", QuoteBySKU, withStatus("expired")), mkSKU(11, 1000, 10), 1)},
			total: map[int64]int{1: 1},
			want:  "This product is no longer available",
		},
		{
			name:  "out of stock",
			lines: []Line{row(prod(1, "s", QuoteBySKU), mkSKU(11, 1000, 0), 1)},
			total: map[int64]int{1: 1},
			want:  "This option is out of stock",
		},
		{
			name:  "not enough stock",
			lines: []Line{row(prod(1, "s", QuoteBySKU), mkSKU(11, 1000, 3), 5)},
			total: map[int64]int{1: 5},
			want:  "Only 3 pieces left in stock, you asked for 5 pieces",
		},
		{
			name: "below the minimum order quantity",
			lines: []Line{row(prod(1, "s", QuoteByProduct, withMOQ(5),
				withTiers(Tier{StartQuantity: 1, PriceFen: 1000})), SKU{}, 1)},
			total: map[int64]int{1: 1},
			want:  "Minimum order is 5 pieces for this product, you have 1 piece",
		},
		{
			name: "not a batch multiple",
			lines: []Line{row(prod(1, "s", QuoteByProduct, withBatch(200),
				withTiers(Tier{StartQuantity: 1, PriceFen: 1000})), SKU{}, 250)},
			total: map[int64]int{1: 250},
			want:  "This product is sold in multiples of 200 pieces, you have 250 pieces",
		},
		{
			name: "mixed batch, both thresholds set",
			lines: []Line{row(prod(1, "s", QuoteByProduct, withMix(10, 5000),
				withTiers(Tier{StartQuantity: 1, PriceFen: 1000})), SKU{}, 1)},
			total: map[int64]int{1: 1},
			want:  "This supplier ships mixed batches of at least 10 pieces or ¥50.00, your basket from them is 1 piece (¥10.00)",
		},
		{
			name: "mixed batch, quantity only",
			lines: []Line{row(prod(1, "s", QuoteByProduct, withMix(10, 0),
				withTiers(Tier{StartQuantity: 1, PriceFen: 1000})), SKU{}, 1)},
			total: map[int64]int{1: 1},
			want:  "This supplier ships mixed batches of at least 10 pieces, your basket from them is 1 piece (¥10.00)",
		},
		{
			name: "mixed batch, amount only",
			lines: []Line{row(prod(1, "s", QuoteByProduct, withMix(0, 5000),
				withTiers(Tier{StartQuantity: 1, PriceFen: 1000})), SKU{}, 1)},
			total: map[int64]int{1: 1},
			want:  "This supplier ships mixed batches of at least ¥50.00, your basket from them is 1 piece (¥10.00)",
		},
		{
			name:  "zero price",
			lines: []Line{row(prod(1, "s", QuoteBySKU), mkSKU(11, 0, 10), 1)},
			total: map[int64]int{1: 1},
			want:  "This product has no price we can sell at right now",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := Quote(tc.lines, tc.total, nil, plain)
			if len(got.Issues) != 1 {
				t.Fatalf("got %d issues, want 1: %+v", len(got.Issues), got.Issues)
			}
			if got.Issues[0].Message != tc.want {
				t.Errorf("message:\n got %q\nwant %q", got.Issues[0].Message, tc.want)
			}
		})
	}
}

// TestMixedBatchTakesTheStricterSetting documents what happens when two
// products from one supplier disagree about the mixed batch thresholds.
func TestMixedBatchTakesTheStricterSetting(t *testing.T) {
	lines := []Line{
		row(prod(1, "s", QuoteByProduct, withMix(5, 1000),
			withTiers(Tier{StartQuantity: 1, PriceFen: 100})), SKU{}, 3),
		row(prod(2, "s", QuoteByProduct, withMix(10, 2000),
			withTiers(Tier{StartQuantity: 1, PriceFen: 100})), SKU{}, 4),
	}
	totals := map[int64]int{1: 3, 2: 4}

	// 7 pieces clears the looser 5 but not the stricter 10, and ¥7.00 clears
	// neither amount, so the cart is held.
	got := Quote(lines, totals, nil, plain)
	if len(got.Issues) != 1 || got.Issues[0].Code != ali.ErrMixedBatch {
		t.Fatalf("want one mixed batch issue, got %+v", got.Issues)
	}
	if got.Issues[0].OfferID != 1 {
		t.Errorf("issue hung on offer %d, want the first mixable offer 1", got.Issues[0].OfferID)
	}
}

// ---- SelectFeeRule --------------------------------------------------------

var (
	jan2025 = time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	jan2026 = time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	jun2026 = time.Date(2026, 6, 1, 0, 0, 0, 0, time.UTC)
	dec2026 = time.Date(2026, 12, 1, 0, 0, 0, 0, time.UTC)
)

func ptr(t time.Time) *time.Time { return &t }

func TestSelectFeeRule(t *testing.T) {
	subject := prod(620201390233, "23tsdvcdsjngp3oj4j3i5", QuoteBySKU, withCategory(1031910))

	global := FeeRule{ID: 1, Scope: ScopeGlobal, From: jan2025}
	category := FeeRule{ID: 2, Scope: ScopeCategory, ScopeValue: "1031910", From: jan2025}
	supplier := FeeRule{ID: 3, Scope: ScopeSupplier, ScopeValue: "23tsdvcdsjngp3oj4j3i5", From: jan2025}
	product := FeeRule{ID: 4, Scope: ScopeProduct, ScopeValue: "620201390233", From: jan2025}

	cases := []struct {
		name   string
		rules  []FeeRule
		now    time.Time
		wantID int64 // 0 means nil
	}{
		{"no rules at all", nil, jun2026, 0},
		{"a lone global rule", []FeeRule{global}, jun2026, 1},
		{"category beats global at equal priority", []FeeRule{global, category}, jun2026, 2},
		{"supplier beats category at equal priority", []FeeRule{category, supplier}, jun2026, 3},
		{"product beats supplier at equal priority", []FeeRule{supplier, product}, jun2026, 4},
		{"product beats everything at equal priority", []FeeRule{global, category, supplier, product}, jun2026, 4},
		{"the same set in reverse order picks the same rule", []FeeRule{product, supplier, category, global}, jun2026, 4},
		{
			name: "priority outranks specificity",
			rules: []FeeRule{
				{ID: 1, Scope: ScopeGlobal, Priority: 10, From: jan2025},
				{ID: 4, Scope: ScopeProduct, ScopeValue: "620201390233", Priority: 1, From: jan2025},
			},
			now: jun2026, wantID: 1,
		},
		{
			name: "the higher id breaks a tie between identical rules",
			rules: []FeeRule{
				{ID: 5, Scope: ScopeGlobal, From: jan2025},
				{ID: 9, Scope: ScopeGlobal, From: jan2025},
				{ID: 7, Scope: ScopeGlobal, From: jan2025},
			},
			now: jun2026, wantID: 9,
		},
		{
			name: "the higher id breaks a tie between identical product rules",
			rules: []FeeRule{
				{ID: 3, Scope: ScopeProduct, ScopeValue: "620201390233", From: jan2025},
				{ID: 4, Scope: ScopeProduct, ScopeValue: "620201390233", From: jan2025},
			},
			now: jun2026, wantID: 4,
		},
		{
			name:  "a category rule for another category does not match",
			rules: []FeeRule{{ID: 2, Scope: ScopeCategory, ScopeValue: "999", From: jan2025}},
			now:   jun2026, wantID: 0,
		},
		{
			name:  "a supplier rule for another supplier does not match",
			rules: []FeeRule{{ID: 3, Scope: ScopeSupplier, ScopeValue: "someone-else", From: jan2025}},
			now:   jun2026, wantID: 0,
		},
		{
			name:  "a product rule for another offer does not match",
			rules: []FeeRule{{ID: 4, Scope: ScopeProduct, ScopeValue: "620201390234", From: jan2025}},
			now:   jun2026, wantID: 0,
		},
		{
			name:  "an unknown scope never matches",
			rules: []FeeRule{{ID: 8, Scope: "region", ScopeValue: "TH", From: jan2025}},
			now:   jun2026, wantID: 0,
		},
		{
			name: "an expired rule is skipped for an open one",
			rules: []FeeRule{
				{ID: 9, Scope: ScopeProduct, ScopeValue: "620201390233", Priority: 99, From: jan2025, To: ptr(jan2026)},
				{ID: 1, Scope: ScopeGlobal, From: jan2025},
			},
			now: jun2026, wantID: 1,
		},
		{
			name:  "a rule that has not started yet is skipped",
			rules: []FeeRule{{ID: 1, Scope: ScopeGlobal, From: dec2026}},
			now:   jun2026, wantID: 0,
		},
		{
			name:  "a rule starting exactly now is live",
			rules: []FeeRule{{ID: 1, Scope: ScopeGlobal, From: jun2026}},
			now:   jun2026, wantID: 1,
		},
		{
			name:  "a rule ending exactly now is already over",
			rules: []FeeRule{{ID: 1, Scope: ScopeGlobal, From: jan2025, To: ptr(jun2026)}},
			now:   jun2026, wantID: 0,
		},
		{
			name:  "an open ended rule never expires",
			rules: []FeeRule{{ID: 1, Scope: ScopeGlobal, From: jan2025, To: nil}},
			now:   dec2026, wantID: 1,
		},
		{
			name:  "scope names are matched case insensitively",
			rules: []FeeRule{{ID: 6, Scope: "Supplier", ScopeValue: "23tsdvcdsjngp3oj4j3i5", From: jan2025}},
			now:   jun2026, wantID: 6,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := SelectFeeRule(tc.rules, subject, tc.now)
			switch {
			case tc.wantID == 0 && got != nil:
				t.Fatalf("got rule %d, want none", got.ID)
			case tc.wantID == 0:
				return
			case got == nil:
				t.Fatalf("got no rule, want %d", tc.wantID)
			case got.ID != tc.wantID:
				t.Fatalf("got rule %d, want %d", got.ID, tc.wantID)
			}
		})
	}
}

// TestSelectFeeRuleReturnsACopy makes sure a caller cannot edit the fee table
// through the pointer it gets back.
func TestSelectFeeRuleReturnsACopy(t *testing.T) {
	rules := []FeeRule{{ID: 1, Scope: ScopeGlobal, FeeBps: 1000, From: jan2025}}
	got := SelectFeeRule(rules, prod(1, "s", QuoteBySKU), jun2026)
	if got == nil {
		t.Fatal("no rule selected")
	}
	got.FeeBps = 9999
	if rules[0].FeeBps != 1000 {
		t.Fatalf("SelectFeeRule handed back a pointer into the caller's slice: %d", rules[0].FeeBps)
	}
}
