package order

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"marketplace/internal/ali"
	"marketplace/internal/store"
)

// ReceiveMessage records an inbound push message and applies it.
//
// The message id is the primary key of message_events, so duplicate delivery
// costs one rejected insert and nothing else. A crash between recording and
// applying is picked up by the message sweep job, which is why the two steps are
// separate rather than one transaction.
func (s *Service) ReceiveMessage(ctx context.Context, m ali.PushMessage) error {
	born := time.UnixMilli(m.GmtBorn)
	if m.GmtBorn == 0 {
		born = time.Now()
	}
	fresh, err := s.DB.RecordMessage(ctx, int64(m.MsgID), m.Topic(), born, m.Data)
	if err != nil {
		return err
	}
	if !fresh {
		return nil // already seen; acknowledging again is harmless
	}

	applyErr := s.ApplyMessage(ctx, m)
	msg := ""
	if applyErr != nil {
		msg = applyErr.Error()
		s.Log.Warn("message handler failed", "msgId", m.MsgID, "topic", m.Topic(), "err", applyErr)
	}
	if err := s.DB.MarkMessageProcessed(ctx, int64(m.MsgID), msg); err != nil {
		return err
	}
	return applyErr
}

// ApplyMessage routes one message to its handler.
func (s *Service) ApplyMessage(ctx context.Context, m ali.PushMessage) error {
	switch m.Topic() {
	case ali.TopicOrderBuyerMake, ali.TopicOrderPay, ali.TopicOrderAnnounceSend,
		ali.TopicOrderPartSend, ali.TopicOrderConfirmReceive, ali.TopicOrderSuccess,
		ali.TopicOrderBuyerClose, ali.TopicOrderSellerClose:
		return s.applyOrderEvent(ctx, m)
	case ali.TopicLogisticsTrace:
		return s.applyTrace(ctx, m)
	case ali.TopicMailNoChange:
		return s.applyMailNoChange(ctx, m)
	case ali.TopicInventoryChange:
		return s.applyInventoryChange(ctx, m)
	case ali.TopicProductChange:
		return s.applyProductChange(ctx, m)
	default:
		s.Log.Debug("ignoring unsubscribed topic", "topic", m.Topic())
		return nil
	}
}

// applyOrderEvent moves a supplier order along, if the event is not stale.
//
// PART_PART_SENDGOODS deliberately carries waitsellersend rather than
// waitbuyerreceive: a partial shipment does not advance the order, and treating
// it as one would tell the shopper their whole order is on its way when half of
// it is still on a shelf.
func (s *Service) applyOrderEvent(ctx context.Context, m ali.PushMessage) error {
	var ev ali.OrderEvent
	if err := json.Unmarshal(m.Data, &ev); err != nil {
		return fmt.Errorf("decode order event: %w", err)
	}
	if ev.OrderID == 0 {
		return errors.New("order event carried no orderId")
	}

	so, err := s.DB.SupplierOrderByCbuID(ctx, int64(ev.OrderID))
	if errors.Is(err, store.ErrNotFound) {
		// An order that is not ours, or one whose create response has not landed
		// yet. Either way there is nothing to update.
		s.Log.Debug("order event for unknown order", "cbuOrderId", ev.OrderID)
		return nil
	}
	if err != nil {
		return err
	}

	status := ev.CurrentStatus
	if status == "" {
		status = ali.TopicStatus(m.Topic())
	}
	at, _ := ali.ParseTime(string(ev.MsgSendTime))
	if at.IsZero() {
		at = time.Now()
	}
	return s.applyStatus(ctx, so, status, at)
}

// applyStatus is the single place a supplier order's status is written.
func (s *Service) applyStatus(ctx context.Context, so store.SupplierOrder, status1688 string, at time.Time) error {
	ours := OurStatus(status1688)
	if ours == "" {
		s.Log.Debug("unmapped 1688 status", "status", status1688, "supplierOrder", so.ID)
		return nil
	}
	applied, err := s.DB.ApplySupplierStatus(ctx, so.ID, status1688, ours, StatusRank(status1688), at)
	if err != nil {
		return err
	}
	if !applied {
		// Worth a line at Info: it is the out-of-order guard doing its job, and
		// an operator reading the log should be able to see it happen.
		s.Log.Info("ignored stale status event", "supplierOrder", so.ID,
			"status", status1688, "held", so.Status1688, "at", at)
		return nil
	}
	return s.refreshOrderStatus(ctx, so.OrderID)
}

// refreshOrderStatus recomputes the customer-facing status from its parcels.
func (s *Service) refreshOrderStatus(ctx context.Context, orderID int64) error {
	o, err := s.DB.GetOrder(ctx, orderID)
	if err != nil {
		return err
	}
	sos, err := s.DB.SupplierOrders(ctx, orderID)
	if err != nil {
		return err
	}
	return s.DB.SetOrderStatus(ctx, orderID, RollUp(o, sos))
}

// applyTrace records one waybill node.
func (s *Service) applyTrace(ctx context.Context, m ali.PushMessage) error {
	var p ali.LogisticsTracePush
	if err := json.Unmarshal(m.Data, &p); err != nil {
		return fmt.Errorf("decode trace: %w", err)
	}
	model := p.Model
	at, _ := ali.ParseTime(string(model.ChangeTime))
	if at.IsZero() {
		at = time.UnixMilli(m.GmtBorn)
	}

	for _, item := range model.OrderLogsItems {
		so, err := s.DB.SupplierOrderByCbuID(ctx, int64(item.OrderID))
		if errors.Is(err, store.ErrNotFound) {
			continue
		}
		if err != nil {
			return err
		}

		shipID, err := s.DB.UpsertShipment(ctx, store.Shipment{
			SupplierOrderID: so.ID,
			Leg:             "china",
			LogisticsID:     model.LogisticsID,
			MailNo:          model.MailNo,
			CpCode:          model.CpCode,
			Status:          model.StatusChanged,
		})
		if err != nil {
			return err
		}
		if _, err := s.DB.AddTrackingEvent(ctx, store.TrackingEvent{
			ShipmentID: shipID,
			Source:     "message",
			EventAt:    at,
			Code:       model.StatusChanged,
			Remark:     traceRemark(model.StatusChanged),
			DedupeKey:  dedupeKey(at, model.StatusChanged),
		}); err != nil {
			return err
		}

		// A waybill reaching its first node means the parcel really has shipped,
		// which the order topics also announce. Applying it here as well means a
		// dropped ANNOUNCE_SENDGOODS does not leave the order looking unshipped.
		if model.StatusChanged == ali.TraceConsign {
			if err := s.applyStatus(ctx, so, "waitbuyerreceive", at); err != nil {
				return err
			}
		}
	}
	return nil
}

// applyMailNoChange corrects a waybill whose carrier or number was changed.
func (s *Service) applyMailNoChange(ctx context.Context, m ali.PushMessage) error {
	var p ali.MailNoChangePush
	if err := json.Unmarshal(m.Data, &p); err != nil {
		return fmt.Errorf("decode waybill change: %w", err)
	}
	at, _ := ali.ParseTime(string(p.Model.EventTime))
	if at.IsZero() {
		at = time.UnixMilli(m.GmtBorn)
	}

	for _, item := range p.Model.OrderLogsItems {
		so, err := s.DB.SupplierOrderByCbuID(ctx, int64(item.OrderID))
		if errors.Is(err, store.ErrNotFound) {
			continue
		}
		if err != nil {
			return err
		}
		shipID, err := s.DB.UpsertShipment(ctx, store.Shipment{
			SupplierOrderID: so.ID,
			Leg:             "china",
			LogisticsID:     p.Model.LogisticsID,
			MailNo:          p.Model.NewMailNo,
			CpCode:          p.Model.NewCpCode,
		})
		if err != nil {
			return err
		}
		if _, err := s.DB.AddTrackingEvent(ctx, store.TrackingEvent{
			ShipmentID: shipID,
			Source:     "message",
			EventAt:    at,
			Code:       "MAIL_NO_CHANGE",
			Remark:     "Tracking number updated to " + p.Model.NewMailNo,
			DedupeKey:  dedupeKey(at, "MAIL_NO_CHANGE"+p.Model.NewMailNo),
		}); err != nil {
			return err
		}
	}
	return nil
}

// applyInventoryChange marks an offer for re-sync when its stock moves. The
// message carries a signed delta rather than a level, so rather than trying to
// apply arithmetic to a possibly stale row, we refetch the offer.
func (s *Service) applyInventoryChange(ctx context.Context, m ali.PushMessage) error {
	var p ali.InventoryChangePush
	if err := json.Unmarshal(m.Data, &p); err != nil {
		return fmt.Errorf("decode inventory change: %w", err)
	}
	seen := map[int64]bool{}
	for _, c := range p.Changes {
		id := int64(c.OfferID)
		if id == 0 || seen[id] {
			continue
		}
		seen[id] = true
		if err := s.DB.Enqueue(ctx, store.JobCatalogImport, "offer:"+strconv.FormatInt(id, 10),
			map[string]any{"offerIds": []int64{id}}, time.Now()); err != nil {
			return err
		}
	}
	return nil
}

// applyProductChange re-syncs an edited offer, or hides a delisted one.
func (s *Service) applyProductChange(ctx context.Context, m ali.PushMessage) error {
	var p ali.ProductChange
	if err := json.Unmarshal(m.Data, &p); err != nil {
		return fmt.Errorf("decode product change: %w", err)
	}
	for _, id := range p.OfferIDs() {
		switch p.Action {
		case ali.ActionMemberDelete, ali.ActionMemberExpired:
			if err := s.DB.SetProductVisible(ctx, int64(id), false); err != nil {
				return err
			}
		default:
			if err := s.DB.Enqueue(ctx, store.JobCatalogImport, "offer:"+id.String(),
				map[string]any{"offerIds": []int64{int64(id)}}, time.Now()); err != nil {
				return err
			}
		}
	}
	return nil
}

// PollSupplierOrder refreshes one order straight from the gateway. This is the
// safety net under the push pipeline: with pushes switched off entirely, an
// order still completes on this path alone.
func (s *Service) PollSupplierOrder(ctx context.Context, so store.SupplierOrder) error {
	if so.CbuOrderID == nil {
		return nil
	}
	cbuID := ali.ID(*so.CbuOrderID)

	info, err := s.Cli.OrderDetail(ctx, cbuID)
	if err != nil {
		return err
	}
	if st := info.BaseInfo.Status; st != "" {
		at, _ := ali.ParseTime(string(info.BaseInfo.ModifyTime))
		if at.IsZero() {
			at = time.Now()
		}
		if err := s.applyStatus(ctx, so, st, at); err != nil {
			return err
		}
	}

	// Waybills come from the order detail; the per-step trace comes from its own
	// endpoint, and both write through the same dedupe key as the push path, so
	// the two converge with no ordering logic.
	for _, li := range info.NativeLogistics.LogisticsItems {
		if li.LogisticsBillNo == "" && li.ID == 0 {
			continue
		}
		if _, err := s.DB.UpsertShipment(ctx, store.Shipment{
			SupplierOrderID: so.ID,
			Leg:             "china",
			LogisticsID:     li.ID.String(),
			MailNo:          li.LogisticsBillNo,
			CpCode:          li.LogisticsCompanyNo,
			CompanyName:     li.LogisticsCompanyName,
			Status:          li.Status,
		}); err != nil {
			return err
		}
	}

	trace, err := s.Cli.TraceOrder(ctx, cbuID, "")
	if err != nil {
		// A parcel with no tracking yet is normal, not a failure.
		if ali.IsCode(err, ali.ErrTraceNotFound, ali.ErrTraceNoPerm, ali.ErrTraceTooOld) {
			return nil
		}
		return err
	}
	for _, t := range trace.LogisticsTrace {
		shipID, err := s.DB.UpsertShipment(ctx, store.Shipment{
			SupplierOrderID: so.ID,
			Leg:             "china",
			LogisticsID:     t.LogisticsID,
			MailNo:          t.LogisticsBillNo,
		})
		if err != nil {
			return err
		}
		for _, step := range t.LogisticsSteps {
			at, _ := ali.ParseTime(string(step.AcceptTime))
			if at.IsZero() {
				continue
			}
			if _, err := s.DB.AddTrackingEvent(ctx, store.TrackingEvent{
				ShipmentID: shipID,
				Source:     "poll",
				EventAt:    at,
				Remark:     step.Remark,
				DedupeKey:  dedupeKey(at, step.Remark),
			}); err != nil {
				return err
			}
		}
	}
	return nil
}

// dedupeKey identifies one physical tracking event. Push and polling describe
// the same movement in different words at the same instant, so the key is the
// timestamp plus the text: identical events collapse, genuinely different ones
// do not.
func dedupeKey(at time.Time, text string) string {
	sum := sha1.Sum([]byte(at.UTC().Format(time.RFC3339) + "|" + text))
	return hex.EncodeToString(sum[:])
}

// traceRemark is the shopper-facing wording for a trace node.
func traceRemark(code string) string {
	switch code {
	case ali.TraceConsign:
		return "Shipped by the supplier"
	case ali.TraceAccept:
		return "Picked up by the carrier"
	case ali.TraceTransport:
		return "In transit"
	case ali.TraceDelivering:
		return "Out for delivery"
	case ali.TraceAgentSign:
		return "Waiting at a collection point"
	case ali.TraceSign:
		return "Delivered to the warehouse"
	case ali.TraceFailed:
		return "Delivery exception reported"
	default:
		return code
	}
}
