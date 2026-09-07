// Package order owns checkout, the relay of customer orders onto 1688, and the
// tracking pipeline that brings their progress back.
package order

import (
	"sort"
	"time"

	"marketplace/internal/store"
)

// Our supplier-order statuses. They map from 1688's own vocabulary and are what
// the customer-facing status is rolled up from.
const (
	StatusPending        = "PENDING"          // created here, not yet sent
	StatusRelayedUnpaid  = "RELAYED_UNPAID"   // exists on 1688, awaiting payment
	StatusPaidToSupplier = "PAID_TO_SUPPLIER" // paid, supplier preparing
	StatusShippedChina   = "SHIPPED_IN_CHINA" // in transit inside China
	StatusAtWarehouse    = "ARRIVED_WAREHOUSE"
	StatusCancelled      = "CANCELLED"
	StatusFailed         = "FAILED" // relay refused; needs a human
)

// Customer-facing order statuses.
const (
	OrderAwaitingPayment = "AWAITING_PAYMENT"
	OrderPaid            = "PAID"
	OrderProcessing      = "PROCESSING"
	OrderPreparing       = "PREPARING"
	OrderShippedChina    = "SHIPPED_IN_CHINA"
	OrderAtWarehouse     = "ARRIVED_WAREHOUSE"
	OrderCancelled       = "CANCELLED"
	OrderPartlyCancelled = "PARTIALLY_CANCELLED"
	OrderRelayFailed     = "RELAY_FAILED"
)

// StatusRank orders 1688's statuses so an out-of-order push message can be
// recognised as stale. The documentation is explicit that messages are
// asynchronous and may arrive out of sequence, so nothing may assign status
// directly: every write goes through this rank.
//
// Terminal states share rank 9 and always win, because a cancellation must never
// be undone by an in-flight event that was merely delayed.
func StatusRank(s1688 string) int {
	switch s1688 {
	case "waitbuyerpay":
		return 1
	case "waitsellersend", "waitlogisticstakein":
		return 2
	case "waitbuyerreceive", "waitbuyersign":
		return 3
	case "confirm_goods", "confirm_goods_and_has_subsidy", "signinsuccess":
		return 4
	case "success":
		return 5
	case "cancel", "terminated":
		return 9
	default:
		return 0
	}
}

// OurStatus translates a 1688 status into ours.
func OurStatus(s1688 string) string {
	switch s1688 {
	case "waitbuyerpay":
		return StatusRelayedUnpaid
	case "waitsellersend", "waitlogisticstakein":
		return StatusPaidToSupplier
	case "waitbuyerreceive", "waitbuyersign":
		return StatusShippedChina
	case "confirm_goods", "confirm_goods_and_has_subsidy", "signinsuccess", "success":
		return StatusAtWarehouse
	case "cancel", "terminated":
		return StatusCancelled
	default:
		return ""
	}
}

// supplierRank orders our own statuses for the roll-up.
func supplierRank(s string) int {
	switch s {
	case StatusPending:
		return 0
	case StatusRelayedUnpaid:
		return 1
	case StatusPaidToSupplier:
		return 2
	case StatusShippedChina:
		return 3
	case StatusAtWarehouse:
		return 4
	default:
		return 0
	}
}

// RollUp computes the customer-facing status from the supplier orders beneath it.
// The slowest parcel sets the headline, because telling someone their order has
// shipped while half of it is still being packed is worse than saying nothing.
func RollUp(o store.Order, sos []store.SupplierOrder) string {
	if len(sos) == 0 {
		if o.PaidAt != nil {
			return OrderPaid
		}
		return OrderAwaitingPayment
	}

	cancelled, failed, live := 0, 0, 0
	lowest := 99
	for _, so := range sos {
		switch so.Status {
		case StatusCancelled:
			cancelled++
		case StatusFailed:
			failed++
		default:
			live++
			if r := supplierRank(so.Status); r < lowest {
				lowest = r
			}
		}
	}

	switch {
	case failed > 0:
		return OrderRelayFailed
	case cancelled == len(sos):
		return OrderCancelled
	case live == 0:
		return OrderCancelled
	}

	status := OrderProcessing
	switch lowest {
	case 0:
		if o.PaidAt == nil {
			status = OrderAwaitingPayment
		} else {
			status = OrderProcessing
		}
	case 1:
		status = OrderProcessing
	case 2:
		status = OrderPreparing
	case 3:
		status = OrderShippedChina
	case 4:
		status = OrderAtWarehouse
	}
	if cancelled > 0 {
		return OrderPartlyCancelled
	}
	return status
}

// Label is the shopper-facing wording for a customer order status.
func Label(status string) string {
	switch status {
	case OrderAwaitingPayment:
		return "Awaiting payment"
	case OrderPaid:
		return "Payment received"
	case OrderProcessing:
		return "Processing"
	case OrderPreparing:
		return "Preparing"
	case OrderShippedChina:
		return "Shipped (China leg)"
	case OrderAtWarehouse:
		return "At consolidation warehouse"
	case OrderCancelled:
		return "Cancelled"
	case OrderPartlyCancelled:
		return "Partially cancelled"
	case OrderRelayFailed:
		return "Needs attention"
	default:
		return status
	}
}

// ---------------------------------------------------------------- timeline --

// Step is one node of the customer-facing progress timeline.
type Step struct {
	Key    string
	Label  string
	At     *time.Time
	Done   bool
	Detail string
}

// timelineSteps is the fixed spine of the timeline. Every order shows all of
// them, so the shopper can see what is still to come rather than only what has
// happened.
var timelineSteps = []struct{ key, label string }{
	{"placed", "Order placed"},
	{"paid", "Payment received"},
	{"relayed", "Sent to supplier"},
	{"supplier_paid", "Supplier paid"},
	{"shipped_cn", "Shipped in China"},
	{"at_warehouse", "At consolidation warehouse"},
	{"intl_shipped", "International shipping"},
	{"delivered", "Delivered"},
}

// BuildTimeline turns an order and its parcels into the progress list. It is a
// pure function of its inputs so it can be tested without a database.
func BuildTimeline(o store.Order, sos []store.SupplierOrder, ships []store.Shipment, evs []store.TrackingEvent) []Step {
	// The earliest time at which every live parcel had reached a given rank.
	reached := map[int]*time.Time{}
	live := 0
	for _, so := range sos {
		if so.Status == StatusCancelled || so.Status == StatusFailed {
			continue
		}
		live++
	}
	for rank := 1; rank <= 4; rank++ {
		count := 0
		var latest time.Time
		for _, so := range sos {
			if so.Status == StatusCancelled || so.Status == StatusFailed {
				continue
			}
			if supplierRank(so.Status) >= rank {
				count++
				if so.StatusAt.After(latest) {
					latest = so.StatusAt
				}
			}
		}
		if live > 0 && count == live && !latest.IsZero() {
			t := latest
			reached[rank] = &t
		}
	}

	partial := func(rank int) string {
		if live <= 1 {
			return ""
		}
		n := 0
		for _, so := range sos {
			if so.Status == StatusCancelled || so.Status == StatusFailed {
				continue
			}
			if supplierRank(so.Status) >= rank {
				n++
			}
		}
		if n > 0 && n < live {
			return itoa(n) + " of " + itoa(live) + " parcels"
		}
		return ""
	}

	relayedAt := (*time.Time)(nil)
	for _, so := range sos {
		if so.CbuOrderID != nil {
			t := so.CreatedAt
			if relayedAt == nil || t.Before(*relayedAt) {
				relayedAt = &t
			}
		}
	}

	// International progress comes from the forwarder, not from 1688: any event
	// recorded against an intl-leg shipment counts.
	var intlFirst, intlDelivered *time.Time
	intlShipments := map[int64]bool{}
	for _, s := range ships {
		if s.Leg == "intl" {
			intlShipments[s.ID] = true
		}
	}
	sorted := append([]store.TrackingEvent(nil), evs...)
	sort.Slice(sorted, func(i, j int) bool { return sorted[i].EventAt.Before(sorted[j].EventAt) })
	for _, e := range sorted {
		if !intlShipments[e.ShipmentID] {
			continue
		}
		if intlFirst == nil {
			t := e.EventAt
			intlFirst = &t
		}
		if e.Code == "SIGN" || e.Code == "DELIVERED" {
			t := e.EventAt
			intlDelivered = &t
		}
	}

	at := map[string]*time.Time{
		"placed":        &o.CreatedAt,
		"paid":          o.PaidAt,
		"relayed":       relayedAt,
		"supplier_paid": reached[2],
		"shipped_cn":    reached[3],
		"at_warehouse":  reached[4],
		"intl_shipped":  intlFirst,
		"delivered":     intlDelivered,
	}
	detail := map[string]string{
		"supplier_paid": partial(2),
		"shipped_cn":    partial(3),
		"at_warehouse":  partial(4),
	}

	out := make([]Step, 0, len(timelineSteps))
	for _, s := range timelineSteps {
		t := at[s.key]
		out = append(out, Step{
			Key:    s.key,
			Label:  s.label,
			At:     t,
			Done:   t != nil,
			Detail: detail[s.key],
		})
	}
	return out
}

func itoa(v int) string {
	if v == 0 {
		return "0"
	}
	var buf [20]byte
	i := len(buf)
	for v > 0 {
		i--
		buf[i] = byte('0' + v%10)
		v /= 10
	}
	return string(buf[i:])
}
