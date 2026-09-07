// Package pricing turns a cart of 1688 offers into the price a shopper pays,
// as a pure function of its inputs. There is no database, no clock, no I/O and
// no package state here: the same inputs always produce the same Result, which
// is what lets an order line store its quote and be re-checked years later.
//
// All money is integer minor units. CNY amounts are ali.Fen, THB amounts are
// ali.Satang, and nofloat_test.go fails the build if a float ever appears.
package pricing

import (
	"strconv"
	"strings"

	"marketplace/internal/ali"
)

// Quote types, from productSaleInfo.quoteType in
// product.search.queryProductDetail:
//
//	0 - no SKU, quoted by product quantity
//	1 - quoted by SKU spec
//	2 - has SKUs but is still quoted by product quantity
//
// Only QuoteBySKU reads the SKU's own price. Type 2 is the trap: the product
// has SKUs, they carry prices, and those prices are not the ones you charge.
const (
	QuoteByProduct    = 0
	QuoteBySKU        = 1
	QuoteBySKUProduct = 2
)

// StatusPublished is the only 1688 product status we will sell.
const StatusPublished = "published"

// CodeUnavailable is our own issue code for a product that is no longer on
// sale. Every other code in this package is one of 1688's, so that our gate
// and the gateway's gate reject a cart with the same vocabulary.
const CodeUnavailable = "UNAVAILABLE"

// Product is the part of a synced 1688 offer that affects price and eligibility.
type Product struct {
	OfferID      int64
	SellerOpenID string
	Status       string // "published" and nothing else is sellable
	QuoteType    int

	MinOrderQty  int  // minOrderQuantity, checked against the offer total
	BatchNumber  int  // batchNumber: the offer total must be a multiple of it
	MixGeneral   bool // sellerMixSetting.generalHunpi
	MixAmountFen ali.Fen
	MixNumber    int

	ChinaFreightFen ali.Fen // per unit, from product.freight.estimate at import time
	FreightFree     bool
	WeightG         int // fallback when the SKU carries no weight

	Tiers      []Tier
	CategoryID int64
}

// Tier is one entry of productSaleInfo.priceRangeList: from StartQuantity
// units up, the unit price is PriceFen.
type Tier struct {
	StartQuantity int
	PriceFen      ali.Fen
	PromoPriceFen ali.Fen
}

// SKU is one productSkuInfos entry.
type SKU struct {
	SkuID         int64
	SpecID        string
	PriceFen      ali.Fen
	PromoPriceFen ali.Fen
	Stock         int
	WeightG       int
}

// Line is one cart row.
type Line struct {
	Product  Product
	SKU      SKU
	Quantity int
}

// Settings are the platform knobs an admin turns.
//
// FXPpm is THB per CNY in parts per million, so 4.9 THB/CNY is 4_900_000.
// RoundingStepSatang snaps the *unit* price to a tidy number; a step of 0 or 1
// leaves it alone. RoundingMode is "up" (ceil to the step) or "nearest" (round
// half up to the step); anything else is treated as "nearest".
type Settings struct {
	FXPpm               int64
	RoundingStepSatang  int64
	RoundingMode        string
	IntlRateSatangPerKg int64
	IntlMinSatang       int64

	// UsePromoPrices lets a quote use a supplier's promotionPrice when it is
	// lower than the wholesale price. Off by default: the documentation never
	// promises that an order preview honours a promotion, and against the
	// gateway it did not, so quoting one would have come out of our margin.
	UsePromoPrices bool
}

// Rounding modes.
const (
	RoundUp      = "up"
	RoundNearest = "nearest"
)

// Issue is one reason a cart cannot be ordered as it stands. Message is
// written for a shopper, not for a log.
type Issue struct {
	Code    string
	Message string
	OfferID int64
	SkuID   int64
}

// LineQuote is the price breakdown for one cart row. Every Fen field is a
// per-unit amount, because that is what the shopper is shown and what the
// order line stores; LineSatang is UnitSatang times Quantity.
type LineQuote struct {
	OfferID    int64
	SkuID      int64
	Quantity   int
	BaseFen    ali.Fen
	FreightFen ali.Fen
	IntlFen    ali.Fen
	FeeFen     ali.Fen
	FeeRuleID  int64
	FXPpm      int64
	UnitSatang ali.Satang
	LineSatang ali.Satang
}

// Totals are the cart totals in THB. The four parts always add up to
// TotalSatang exactly: TotalSatang is the authoritative sum of the line
// totals, and GoodsSatang absorbs the difference the unit rounding step
// introduces, so an invoice built from these five numbers always balances.
type Totals struct {
	GoodsSatang   ali.Satang
	FreightSatang ali.Satang
	IntlSatang    ali.Satang
	FeeSatang     ali.Satang
	TotalSatang   ali.Satang
}

// Result is the whole answer: one LineQuote per input line in input order, the
// totals, and every issue found. Lines are priced even when they have issues,
// so the cart can still render a price next to the complaint.
type Result struct {
	Lines  []LineQuote
	Totals Totals
	Issues []Issue
}

// OK reports whether the cart can be ordered as it stands.
func (r Result) OK() bool { return len(r.Issues) == 0 }

// Quote prices lines and validates them.
//
// offerTotals maps an offer id to the total quantity of that offer across the
// WHOLE cart, which is what tier selection, the minimum order quantity and the
// batch multiple are all measured against. Quote never derives it from lines,
// because the caller may be quoting only part of a cart; an offer missing from
// the map falls back to the quantity present in lines, so a caller quoting an
// entire cart may pass nil.
//
// rule chooses the fee rule for a product and may return nil, or be nil
// itself, in which case the line is quoted with a zero fee and FeeRuleID 0.
//
// Issues come out in a stable order: for each line in input order, its offer's
// quantity problems (once per offer, at its first line) then that line's own
// problems, and finally one mixed batch issue per seller group that fails,
// which can only be judged once every line has a price.
func Quote(lines []Line, offerTotals map[int64]int, rule func(Product) *FeeRule, s Settings) Result {
	res := Result{Lines: make([]LineQuote, 0, len(lines))}
	totals := resolveOfferTotals(lines, offerTotals)

	groups := make(map[string]*sellerGroup, len(lines))
	groupOrder := make([]string, 0, len(lines))
	offerSeen := make(map[int64]bool, len(lines))

	var freightSat, intlSat, feeSat, totalSat ali.Satang

	for _, ln := range lines {
		p, sku, qty := ln.Product, ln.SKU, ln.Quantity
		offerTotal := totals[p.OfferID]

		// Offer-level quantity rules, reported once however many rows of the
		// same offer are in the cart.
		if !offerSeen[p.OfferID] {
			offerSeen[p.OfferID] = true
			if p.MinOrderQty > 0 && offerTotal < p.MinOrderQty {
				res.Issues = append(res.Issues, Issue{
					Code:    ali.ErrBelowMOQ,
					Message: "Minimum order is " + pieces(p.MinOrderQty) + " for this product, you have " + pieces(offerTotal),
					OfferID: p.OfferID,
				})
			}
			if p.BatchNumber > 0 && offerTotal%p.BatchNumber != 0 {
				res.Issues = append(res.Issues, Issue{
					Code:    ali.ErrBelowMOQ,
					Message: "This product is sold in multiples of " + pieces(p.BatchNumber) + ", you have " + pieces(offerTotal),
					OfferID: p.OfferID,
				})
			}
		}

		if !strings.EqualFold(strings.TrimSpace(p.Status), StatusPublished) {
			res.Issues = append(res.Issues, Issue{
				Code:    CodeUnavailable,
				Message: "This product is no longer available",
				OfferID: p.OfferID,
				SkuID:   sku.SkuID,
			})
		}

		// Stock lives on the SKU, so a product quoted without one (quoteType 0)
		// has nothing to check here.
		if sku.SkuID != 0 && qty > sku.Stock {
			msg := "Only " + pieces(sku.Stock) + " left in stock, you asked for " + pieces(qty)
			if sku.Stock <= 0 {
				msg = "This option is out of stock"
			}
			res.Issues = append(res.Issues, Issue{
				Code:    ali.ErrNoStock,
				Message: msg,
				OfferID: p.OfferID,
				SkuID:   sku.SkuID,
			})
		}

		base := unitBaseFen(p, sku, offerTotal, s.UsePromoPrices)
		if base <= 0 {
			res.Issues = append(res.Issues, Issue{
				Code:    ali.ErrZeroPrice,
				Message: "This product has no price we can sell at right now",
				OfferID: p.OfferID,
				SkuID:   sku.SkuID,
			})
		}

		var freight ali.Fen
		if !p.FreightFree {
			freight = p.ChinaFreightFen
		}
		intl := intlPerUnitSatang(s, weightG(p, sku), qty).ToFen(s.FXPpm)

		subtotal := base + freight + intl

		fr := FeeRule{}
		if rule != nil {
			if got := rule(p); got != nil {
				fr = *got
			}
		}
		fee := feeFen(subtotal, fr)

		unit := snap((subtotal + fee).ToSatang(s.FXPpm), s.RoundingStepSatang, s.RoundingMode)
		line := unit * ali.Satang(qty)

		res.Lines = append(res.Lines, LineQuote{
			OfferID:    p.OfferID,
			SkuID:      sku.SkuID,
			Quantity:   qty,
			BaseFen:    base,
			FreightFen: freight,
			IntlFen:    intl,
			FeeFen:     fee,
			FeeRuleID:  fr.ID,
			FXPpm:      s.FXPpm,
			UnitSatang: unit,
			LineSatang: line,
		})

		freightSat += (freight * ali.Fen(qty)).ToSatang(s.FXPpm)
		intlSat += (intl * ali.Fen(qty)).ToSatang(s.FXPpm)
		feeSat += (fee * ali.Fen(qty)).ToSatang(s.FXPpm)
		totalSat += line

		g := groups[p.SellerOpenID]
		if g == nil {
			g = &sellerGroup{}
			groups[p.SellerOpenID] = g
			groupOrder = append(groupOrder, p.SellerOpenID)
		}
		g.add(p, sku, qty, base)
	}

	for _, seller := range groupOrder {
		if iss, bad := groups[seller].check(); bad {
			res.Issues = append(res.Issues, iss)
		}
	}

	res.Totals = Totals{
		FreightSatang: freightSat,
		IntlSatang:    intlSat,
		FeeSatang:     feeSat,
		TotalSatang:   totalSat,
		GoodsSatang:   totalSat - freightSat - intlSat - feeSat,
	}
	return res
}

// resolveOfferTotals copies the caller's totals and fills in only the offers it
// did not mention.
func resolveOfferTotals(lines []Line, given map[int64]int) map[int64]int {
	out := make(map[int64]int, len(lines))
	for _, ln := range lines {
		if _, ok := given[ln.Product.OfferID]; ok {
			continue
		}
		out[ln.Product.OfferID] += ln.Quantity
	}
	for id, n := range given {
		out[id] = n
	}
	return out
}

// unitBaseFen is the 1688 unit price for this line, before any of our costs.
//
// quoteType 1 uses the SKU's own price. Everything else is quoted by product
// quantity: the tier whose StartQuantity is the largest one not above the
// offer's TOTAL quantity, so two colours of one product bought six at a time
// both price at the twelve tier. Below the cheapest tier's threshold the first
// tier applies, and an offer with no tiers at all falls back to the SKU price.
func unitBaseFen(p Product, sku SKU, offerTotal int, usePromo bool) ali.Fen {
	if p.QuoteType == QuoteBySKU {
		return effectiveFen(sku.PriceFen, sku.PromoPriceFen, usePromo)
	}
	if t, ok := pickTier(p.Tiers, offerTotal); ok {
		return effectiveFen(t.PriceFen, t.PromoPriceFen, usePromo)
	}
	return effectiveFen(sku.PriceFen, sku.PromoPriceFen, usePromo)
}

// pickTier returns the tier that governs qty. Tiers need not be sorted.
func pickTier(tiers []Tier, qty int) (Tier, bool) {
	if len(tiers) == 0 {
		return Tier{}, false
	}
	best, lowest := -1, 0
	for i := range tiers {
		if tiers[i].StartQuantity < tiers[lowest].StartQuantity {
			lowest = i
		}
		if tiers[i].StartQuantity > qty {
			continue
		}
		if best < 0 || tiers[i].StartQuantity >= tiers[best].StartQuantity {
			best = i
		}
	}
	if best < 0 {
		best = lowest // below the first tier: the first tier's price stands
	}
	return tiers[best], true
}

// effectiveFen prefers the promotional price when promotions are enabled, and
// only when there really is one and it really is cheaper. 1688 sends 0 for "no
// promotion" and, often enough, a promotionPrice above the base price for an
// expired campaign.
func effectiveFen(price, promo ali.Fen, usePromo bool) ali.Fen {
	if usePromo && promo > 0 && promo < price {
		return promo
	}
	return price
}

// weightG prefers the SKU's shipping weight and falls back to the product's.
func weightG(p Product, sku SKU) int {
	if sku.WeightG > 0 {
		return sku.WeightG
	}
	return p.WeightG
}

// intlPerUnitSatang is our own international leg, per unit: the weight based
// charge, or the minimum for the line spread over its units, whichever is
// larger. It is computed in THB because that is the currency of our rate card,
// and converted back to Fen by the caller so the breakdown reads in one
// currency.
func intlPerUnitSatang(s Settings, grams, qty int) ali.Satang {
	byWeight := ali.DivRoundHalfUp(int64(grams)*s.IntlRateSatangPerKg, 1000)
	minPerUnit := ali.DivRoundHalfUp(s.IntlMinSatang, int64(qty)) // qty <= 0 gives 0
	if minPerUnit > byWeight {
		return ali.Satang(minPerUnit)
	}
	return ali.Satang(byWeight)
}

// feeFen is our margin on one unit: a share of the landed cost plus a fixed
// part, never less than the floor.
func feeFen(subtotal ali.Fen, r FeeRule) ali.Fen {
	fee := ali.Fen(ali.DivRoundHalfUp(int64(subtotal)*int64(r.FeeBps), 10_000)) + r.FeeFixedFen
	if fee < r.MinFeeFen {
		return r.MinFeeFen
	}
	return fee
}

// snap rounds a unit price to a tidy step. We snap the UNIT and then multiply:
// snapping the line total instead makes the displayed unit price disagree with
// the total divided by the quantity, which shoppers notice and support tickets
// follow.
func snap(v ali.Satang, step int64, mode string) ali.Satang {
	if step <= 1 {
		return v
	}
	n := int64(v)
	if strings.EqualFold(strings.TrimSpace(mode), RoundUp) {
		if n >= 0 {
			return ali.Satang(((n + step - 1) / step) * step)
		}
		return ali.Satang(-((-n) / step) * step)
	}
	return ali.Satang(ali.DivRoundHalfUp(n, step) * step)
}

// sellerGroup accumulates one 1688 supplier's share of the cart, which is the
// unit the mixed batch rule is judged on because it is also the unit that
// becomes one 1688 order.
type sellerGroup struct {
	qty     int
	baseFen ali.Fen

	hasMix    bool
	mixNumber int
	mixAmount ali.Fen
	offerID   int64 // the first mix enabled offer, to hang the issue on
	skuID     int64
}

func (g *sellerGroup) add(p Product, sku SKU, qty int, base ali.Fen) {
	g.qty += qty
	g.baseFen += base * ali.Fen(qty)
	if !p.MixGeneral {
		return
	}
	if !g.hasMix {
		g.hasMix, g.offerID, g.skuID = true, p.OfferID, sku.SkuID
	}
	// Sellers normally publish one mixed batch setting across their catalogue,
	// but if two products in the group disagree we hold the cart to the
	// stricter of the two rather than let the looser one wave it through.
	if p.MixNumber > g.mixNumber {
		g.mixNumber = p.MixNumber
	}
	if p.MixAmountFen > g.mixAmount {
		g.mixAmount = p.MixAmountFen
	}
}

// check applies the mixed batch rule. It is satisfied by quantity OR by
// amount, never by both: a seller who asks for 20 pieces or 100 yuan is happy
// with either, and requiring both is the classic way to reject a valid cart.
func (g *sellerGroup) check() (Issue, bool) {
	if !g.hasMix || (g.mixNumber <= 0 && g.mixAmount <= 0) {
		return Issue{}, false
	}
	if g.mixNumber > 0 && g.qty >= g.mixNumber {
		return Issue{}, false
	}
	if g.mixAmount > 0 && g.baseFen >= g.mixAmount {
		return Issue{}, false
	}

	var need string
	switch {
	case g.mixNumber > 0 && g.mixAmount > 0:
		need = pieces(g.mixNumber) + " or " + g.mixAmount.Text()
	case g.mixNumber > 0:
		need = pieces(g.mixNumber)
	default:
		need = g.mixAmount.Text()
	}
	return Issue{
		Code:    ali.ErrMixedBatch,
		Message: "This supplier ships mixed batches of at least " + need + ", your basket from them is " + pieces(g.qty) + " (" + g.baseFen.Text() + ")",
		OfferID: g.offerID,
		SkuID:   g.skuID,
	}, true
}

func pieces(n int) string {
	if n == 1 {
		return "1 piece"
	}
	return strconv.Itoa(n) + " pieces"
}
