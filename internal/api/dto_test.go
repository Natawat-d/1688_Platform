package api

import (
	"encoding/json"
	"regexp"
	"testing"
	"time"

	"marketplace/internal/ali"
	"marketplace/internal/order"
	"marketplace/internal/pricing"
	"marketplace/internal/store"
)

// bareBigInt matches an unquoted integer of sixteen or more digits anywhere in a
// JSON document. JavaScript's Number loses precision past 2^53, so any such
// value reaching the browser would be silently corrupted. 1688 order ids are
// nineteen digits; this is the test that keeps them out of our own API as raw
// numbers.
var bareBigInt = regexp.MustCompile(`[^"\d]\d{16,}[^"\d]`)

func assertNoBareBigInts(t *testing.T, name string, v any) {
	t.Helper()
	b, err := json.Marshal(v)
	if err != nil {
		t.Fatalf("%s: marshal: %v", name, err)
	}
	if m := bareBigInt.Find(b); m != nil {
		t.Errorf("%s: response carries a bare large integer %q; identifiers must be strings", name, m)
	}
}

func TestDTOsNeverEmitBareLargeIntegers(t *testing.T) {
	const big = int64(2958624554509662976) // a real-sized 1688 order id
	now := time.Now()

	p := store.Product{
		OfferID: big, SellerOpenID: "seller", CompanyName: "Co", SubjectTrans: "Thing",
		SellMinSatang: 123456, MinOrderQuantity: 5, Images: []string{"x"},
	}
	assertNoBareBigInts(t, "ProductCardDTO", cardDTO(p))

	line := order.QuotedLine{
		Item:    store.CartItem{ID: big, OfferID: big, SkuID: big, Quantity: 2},
		Product: p,
		SKU:     store.SKU{SkuID: big, OfferID: big, PriceFen: 1850, AmountOnSale: 9},
		Quote: pricing.LineQuote{
			OfferID: big, SkuID: big, Quantity: 2, BaseFen: 1850, FeeRuleID: big,
			FXPpm: 4_900_000, UnitSatang: 9999, LineSatang: 19998,
		},
	}
	cart := order.CartQuote{
		CartID: "abc",
		Groups: []order.Group{{SellerOpenID: "s", SellerName: "S", Parcels: 1, Lines: []order.QuotedLine{line},
			Issues: []pricing.Issue{{Code: "500_005", OfferID: big, SkuID: big, Message: "moq"}}}},
		Checkoutable: false,
	}
	assertNoBareBigInts(t, "CartDTO", cartDTO(cart))

	cbu := big
	o := store.Order{ID: big, PublicID: "MK-2026-000001", CreatedAt: now, PaidAt: &now, TotalSatang: 100}
	sos := []store.SupplierOrder{{ID: big, OrderID: big, CbuOrderID: &cbu, Status: order.StatusShippedChina, StatusAt: now, CreatedAt: now}}
	items := []store.OrderItem{{ID: big, OrderID: big, SupplierOrderID: &cbu, OfferID: big, SkuID: big, Quantity: 1, FeeRuleID: &cbu, FxPpm: 4_900_000}}
	ships := []store.Shipment{{ID: big, SupplierOrderID: big, Leg: "china", MailNo: "SF123"}}
	evs := []store.TrackingEvent{{ID: big, ShipmentID: big, Source: "message", EventAt: now, Code: ali.TraceConsign, Remark: "shipped"}}
	assertNoBareBigInts(t, "OrderDTO", orderDTO(o, items, sos, ships, evs))

	assertNoBareBigInts(t, "IssueDTO", issueDTO(pricing.Issue{Code: "500_004", OfferID: big, SkuID: big}))
}

func TestMoneyIsStringMinorPlusText(t *testing.T) {
	m := thb(ali.Satang(123456))
	if m.Minor != "123456" || m.Currency != "THB" || m.Text != "฿1,234.56" {
		t.Errorf("thb(123456) = %+v", m)
	}
	c := cny(ali.Fen(-1850))
	if c.Minor != "-1850" || c.Text != "¥-18.50" && c.Text != "-¥18.50" {
		t.Errorf("cny(-1850) = %+v", c)
	}
}
