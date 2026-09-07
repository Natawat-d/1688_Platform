package api

import (
	"encoding/json"
	"strconv"
	"time"

	"marketplace/internal/ali"
	"marketplace/internal/order"
	"marketplace/internal/pricing"
	"marketplace/internal/store"
)

// Money is how every amount crosses the wire.
//
// minor is a STRING of minor units and text is preformatted for display, so the
// browser never does currency arithmetic and never sees a float. The client
// renders text; minor exists for sorting and for the price-breakdown panel.
type Money struct {
	Minor    string `json:"minor"`
	Currency string `json:"currency"`
	Text     string `json:"text"`
}

func thb(v ali.Satang) Money {
	return Money{Minor: strconv.FormatInt(int64(v), 10), Currency: "THB", Text: v.Text()}
}

func cny(v ali.Fen) Money {
	return Money{Minor: strconv.FormatInt(int64(v), 10), Currency: "CNY", Text: v.Text()}
}

// id renders a 1688 or database identifier as a string.
//
// This is the rule that keeps JavaScript from corrupting our data: 1688 order
// ids reach nineteen digits, and JSON.parse silently rounds any integer past
// 2^53. dto_test.go fails the build if a response ever carries a bare integer of
// sixteen digits or more.
func id(v int64) string { return strconv.FormatInt(v, 10) }

// ---------------------------------------------------------------- catalogue --

type CategoryDTO struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type SellerDTO struct {
	OpenID string `json:"openId"`
	Name   string `json:"name"`
	Score  string `json:"score"`
}

type ProductCardDTO struct {
	OfferID   string    `json:"offerId"`
	Title     string    `json:"title"`
	Image     string    `json:"image"`
	PriceFrom Money     `json:"priceFrom"`
	MOQ       int       `json:"moq"`
	Unit      string    `json:"unit"`
	MonthSold int       `json:"monthSold"`
	Seller    SellerDTO `json:"seller"`
}

type SkuAttrDTO struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	Image string `json:"image,omitempty"`
}

type SkuDTO struct {
	SkuID     string       `json:"skuId"`
	SpecID    string       `json:"specId"`
	Label     string       `json:"label"`
	Attrs     []SkuAttrDTO `json:"attrs"`
	Stock     int          `json:"stock"`
	Available bool         `json:"available"`
	Price     Money        `json:"price"`
}

type TierDTO struct {
	StartQuantity int   `json:"startQuantity"`
	Price         Money `json:"price"`
}

type ShippingDTO struct {
	WeightG  int `json:"weightG"`
	LengthMM int `json:"lengthMm"`
	WidthMM  int `json:"widthMm"`
	HeightMM int `json:"heightMm"`
}

type MixDTO struct {
	General bool   `json:"general"`
	Amount  Money  `json:"amount"`
	Number  int    `json:"number"`
	Note    string `json:"note,omitempty"`
}

type ProductDTO struct {
	OfferID      string        `json:"offerId"`
	Title        string        `json:"title"`
	TitleZh      string        `json:"titleZh"`
	Description  string        `json:"description"`
	Images       []string      `json:"images"`
	CategoryPath []CategoryDTO `json:"categoryPath"`
	Seller       SellerDTO     `json:"seller"`
	Status       string        `json:"status"`
	Sellable     bool          `json:"sellable"`
	MOQ          int           `json:"moq"`
	BatchNumber  int           `json:"batchNumber"`
	QuoteType    int           `json:"quoteType"`
	Unit         string        `json:"unit"`
	Mix          MixDTO        `json:"mix"`
	Stock        int           `json:"stock"`
	SKUs         []SkuDTO      `json:"skus"`
	Tiers        []TierDTO     `json:"tiers"`
	Shipping     ShippingDTO   `json:"shipping"`
	PriceRange   struct {
		Min Money `json:"min"`
		Max Money `json:"max"`
	} `json:"priceRange"`
	SyncedAt time.Time `json:"syncedAt"`
}

type SearchResponse struct {
	Total int              `json:"total"`
	Page  int              `json:"page"`
	Size  int              `json:"size"`
	Items []ProductCardDTO `json:"items"`
}

// --------------------------------------------------------------------- cart --

type BreakdownDTO struct {
	Base      Money  `json:"base"`
	Freight   Money  `json:"freight"`
	Intl      Money  `json:"intl"`
	Fee       Money  `json:"fee"`
	FxPpm     string `json:"fxPpm"`
	FeeRuleID string `json:"feeRuleId"`
}

type CartLineDTO struct {
	ID        string       `json:"id"`
	OfferID   string       `json:"offerId"`
	SkuID     string       `json:"skuId"`
	Title     string       `json:"title"`
	Image     string       `json:"image"`
	SkuLabel  string       `json:"skuLabel"`
	Quantity  int          `json:"quantity"`
	Stock     int          `json:"stock"`
	Unit      Money        `json:"unit"`
	Line      Money        `json:"line"`
	Breakdown BreakdownDTO `json:"breakdown"`
}

type IssueDTO struct {
	Code    string `json:"code"`
	Field   string `json:"field"`
	Message string `json:"message"`
}

type CartGroupDTO struct {
	SellerOpenID string        `json:"sellerOpenId"`
	SellerName   string        `json:"sellerName"`
	Parcels      int           `json:"parcels"`
	Lines        []CartLineDTO `json:"lines"`
	Issues       []IssueDTO    `json:"issues"`
	Subtotal     Money         `json:"subtotal"`
}

type CartTotalsDTO struct {
	Goods        Money `json:"goods"`
	ChinaFreight Money `json:"chinaFreight"`
	Intl         Money `json:"intl"`
	Fee          Money `json:"fee"`
	Total        Money `json:"total"`
}

type CartDTO struct {
	ID           string         `json:"id"`
	Groups       []CartGroupDTO `json:"groups"`
	Totals       CartTotalsDTO  `json:"totals"`
	Checkoutable bool           `json:"checkoutable"`
	Count        int            `json:"count"`
}

// ------------------------------------------------------------------- orders --

type TimelineStepDTO struct {
	Key    string     `json:"key"`
	Label  string     `json:"label"`
	At     *time.Time `json:"at,omitempty"`
	Done   bool       `json:"done"`
	Detail string     `json:"detail,omitempty"`
}

type TrackEventDTO struct {
	At     time.Time `json:"at"`
	Text   string    `json:"text"`
	Source string    `json:"source"`
	Code   string    `json:"code,omitempty"`
}

type ParcelDTO struct {
	ID          string          `json:"id"`
	SellerName  string          `json:"sellerName"`
	Status      string          `json:"status"`
	StatusLabel string          `json:"statusLabel"`
	CbuOrderID  string          `json:"cbuOrderId,omitempty"`
	Carrier     string          `json:"carrier,omitempty"`
	TrackingNo  string          `json:"trackingNo,omitempty"`
	PayURL      string          `json:"payUrl,omitempty"`
	Items       []string        `json:"items"`
	Events      []TrackEventDTO `json:"events"`
	Error       string          `json:"error,omitempty"`
}

type OrderItemDTO struct {
	OfferID   string       `json:"offerId"`
	SkuID     string       `json:"skuId"`
	Title     string       `json:"title"`
	Image     string       `json:"image"`
	SkuLabel  string       `json:"skuLabel"`
	Quantity  int          `json:"quantity"`
	Unit      Money        `json:"unit"`
	Line      Money        `json:"line"`
	Breakdown BreakdownDTO `json:"breakdown"`
}

type OrderDTO struct {
	OrderID     string            `json:"orderId"`
	Status      string            `json:"status"`
	StatusLabel string            `json:"statusLabel"`
	PlacedAt    time.Time         `json:"placedAt"`
	PaidAt      *time.Time        `json:"paidAt,omitempty"`
	Email       string            `json:"email"`
	Totals      CartTotalsDTO     `json:"totals"`
	Items       []OrderItemDTO    `json:"items"`
	Timeline    []TimelineStepDTO `json:"timeline"`
	Parcels     []ParcelDTO       `json:"parcels"`
	Token       string            `json:"token,omitempty"`
}

// ------------------------------------------------------------------ mapping --

func cardDTO(p store.Product) ProductCardDTO {
	return ProductCardDTO{
		OfferID:   id(p.OfferID),
		Title:     p.SubjectTrans,
		Image:     firstImage(p),
		PriceFrom: thb(ali.Satang(p.SellMinSatang)),
		MOQ:       int(p.MinOrderQuantity),
		Unit:      p.UnitTrans,
		MonthSold: int(p.MonthSold),
		Seller:    SellerDTO{OpenID: p.SellerOpenID, Name: p.CompanyName, Score: p.TradeScore},
	}
}

func firstImage(p store.Product) string {
	if p.WhiteImage != "" {
		return p.WhiteImage
	}
	if len(p.Images) > 0 {
		return p.Images[0]
	}
	return ""
}

func issueDTO(is pricing.Issue) IssueDTO {
	field := ""
	if is.OfferID != 0 {
		field = id(is.OfferID)
	}
	return IssueDTO{Code: is.Code, Field: field, Message: is.Message}
}

func breakdownDTO(q pricing.LineQuote) BreakdownDTO {
	return BreakdownDTO{
		Base:      cny(q.BaseFen),
		Freight:   cny(q.FreightFen),
		Intl:      cny(q.IntlFen),
		Fee:       cny(q.FeeFen),
		FxPpm:     strconv.FormatInt(q.FXPpm, 10),
		FeeRuleID: id(q.FeeRuleID),
	}
}

func cartDTO(q order.CartQuote) CartDTO {
	out := CartDTO{
		ID:           q.CartID,
		Checkoutable: q.Checkoutable,
		Totals: CartTotalsDTO{
			Goods:        thb(q.Totals.GoodsSatang),
			ChinaFreight: thb(q.Totals.FreightSatang),
			Intl:         thb(q.Totals.IntlSatang),
			Fee:          thb(q.Totals.FeeSatang),
			Total:        thb(q.Totals.TotalSatang),
		},
		Groups: []CartGroupDTO{},
	}
	for _, g := range q.Groups {
		gd := CartGroupDTO{
			SellerOpenID: g.SellerOpenID,
			SellerName:   g.SellerName,
			Parcels:      g.Parcels,
			Subtotal:     thb(g.SubtotalSatang),
			Lines:        []CartLineDTO{},
			Issues:       []IssueDTO{},
		}
		for _, l := range g.Lines {
			out.Count += int(l.Item.Quantity)
			gd.Lines = append(gd.Lines, CartLineDTO{
				ID:        id(l.Item.ID),
				OfferID:   id(l.Product.OfferID),
				SkuID:     id(l.SKU.SkuID),
				Title:     l.Product.SubjectTrans,
				Image:     lineImage(l),
				SkuLabel:  l.SKU.Label,
				Quantity:  int(l.Item.Quantity),
				Stock:     int(l.SKU.AmountOnSale),
				Unit:      thb(l.Quote.UnitSatang),
				Line:      thb(l.Quote.LineSatang),
				Breakdown: breakdownDTO(l.Quote),
			})
		}
		for _, is := range g.Issues {
			gd.Issues = append(gd.Issues, issueDTO(is))
		}
		out.Groups = append(out.Groups, gd)
	}
	return out
}

func lineImage(l order.QuotedLine) string {
	if l.SKU.ImageURL != "" {
		return l.SKU.ImageURL
	}
	return firstImage(l.Product)
}

func orderDTO(o store.Order, items []store.OrderItem, sos []store.SupplierOrder,
	ships []store.Shipment, evs []store.TrackingEvent) OrderDTO {

	status := order.RollUp(o, sos)
	out := OrderDTO{
		OrderID:     o.PublicID,
		Status:      status,
		StatusLabel: order.Label(status),
		PlacedAt:    o.CreatedAt,
		PaidAt:      o.PaidAt,
		Email:       o.Email,
		Totals: CartTotalsDTO{
			Goods:        thb(ali.Satang(o.GoodsSatang)),
			ChinaFreight: thb(ali.Satang(o.FreightSatang)),
			Intl:         thb(ali.Satang(o.IntlSatang)),
			Fee:          thb(ali.Satang(o.FeeSatang)),
			Total:        thb(ali.Satang(o.TotalSatang)),
		},
		Items:    []OrderItemDTO{},
		Timeline: []TimelineStepDTO{},
		Parcels:  []ParcelDTO{},
	}

	for _, it := range items {
		out.Items = append(out.Items, OrderItemDTO{
			OfferID:  id(it.OfferID),
			SkuID:    id(it.SkuID),
			Title:    it.Title,
			Image:    it.ImageURL,
			SkuLabel: it.SkuLabel,
			Quantity: int(it.Quantity),
			Unit:     thb(ali.Satang(it.UnitSatang)),
			Line:     thb(ali.Satang(it.LineSatang)),
			Breakdown: BreakdownDTO{
				Base:      cny(ali.Fen(it.BaseFen)),
				Freight:   cny(ali.Fen(it.FreightFen)),
				Intl:      cny(ali.Fen(it.IntlFen)),
				Fee:       cny(ali.Fen(it.FeeFen)),
				FxPpm:     strconv.FormatInt(it.FxPpm, 10),
				FeeRuleID: ruleIDString(it.FeeRuleID),
			},
		})
	}

	for _, st := range order.BuildTimeline(o, sos, ships, evs) {
		out.Timeline = append(out.Timeline, TimelineStepDTO{
			Key: st.Key, Label: st.Label, At: st.At, Done: st.Done, Detail: st.Detail,
		})
	}

	shipsBySupplier := map[int64][]store.Shipment{}
	for _, sh := range ships {
		shipsBySupplier[sh.SupplierOrderID] = append(shipsBySupplier[sh.SupplierOrderID], sh)
	}
	evsByShipment := map[int64][]store.TrackingEvent{}
	for _, e := range evs {
		evsByShipment[e.ShipmentID] = append(evsByShipment[e.ShipmentID], e)
	}
	itemsBySupplier := map[int64][]string{}
	for _, it := range items {
		if it.SupplierOrderID != nil {
			itemsBySupplier[*it.SupplierOrderID] = append(itemsBySupplier[*it.SupplierOrderID], id(it.SkuID))
		}
	}

	for _, so := range sos {
		p := ParcelDTO{
			ID:          id(so.ID),
			SellerName:  so.SellerName,
			Status:      so.Status,
			StatusLabel: parcelLabel(so.Status),
			PayURL:      so.PayURL,
			Items:       itemsBySupplier[so.ID],
			Events:      []TrackEventDTO{},
			Error:       so.ErrorMessage,
		}
		if so.CbuOrderID != nil {
			p.CbuOrderID = id(*so.CbuOrderID)
		}
		for _, sh := range shipsBySupplier[so.ID] {
			if p.TrackingNo == "" {
				p.TrackingNo = sh.MailNo
				p.Carrier = firstNonEmpty(sh.CompanyName, sh.CpCode)
			}
			for _, e := range evsByShipment[sh.ID] {
				p.Events = append(p.Events, TrackEventDTO{
					At: e.EventAt, Text: e.Remark, Source: e.Source, Code: e.Code,
				})
			}
		}
		if p.Items == nil {
			p.Items = []string{}
		}
		out.Parcels = append(out.Parcels, p)
	}
	return out
}

func parcelLabel(status string) string {
	switch status {
	case order.StatusPending:
		return "Queued"
	case order.StatusRelayedUnpaid:
		return "Awaiting supplier payment"
	case order.StatusPaidToSupplier:
		return "Supplier preparing"
	case order.StatusShippedChina:
		return "In transit in China"
	case order.StatusAtWarehouse:
		return "At warehouse"
	case order.StatusCancelled:
		return "Cancelled"
	case order.StatusFailed:
		return "Needs attention"
	default:
		return status
	}
}

func ruleIDString(v *int64) string {
	if v == nil {
		return "0"
	}
	return id(*v)
}

func firstNonEmpty(vals ...string) string {
	for _, v := range vals {
		if v != "" {
			return v
		}
	}
	return ""
}

func attrsOf(raw json.RawMessage) []SkuAttrDTO {
	var in []store.SKUAttr
	if len(raw) > 0 {
		_ = json.Unmarshal(raw, &in)
	}
	out := make([]SkuAttrDTO, 0, len(in))
	for _, a := range in {
		out = append(out, SkuAttrDTO{
			Name:  firstNonEmpty(a.NameTrans, a.Name),
			Value: firstNonEmpty(a.ValueTrans, a.Value),
			Image: a.Image,
		})
	}
	return out
}
