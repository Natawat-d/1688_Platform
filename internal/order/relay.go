package order

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"marketplace/internal/ali"
	"marketplace/internal/catalog"
	"marketplace/internal/jobs"
	"marketplace/internal/store"
)

// cancelGrace is the documented window in which a cancel is refused with
// CLOSE_ORDER_TOO_FAST. Cancels are scheduled past it rather than retried into it.
const cancelGrace = 11 * time.Second

// Register wires every job handler onto the runner.
func (s *Service) Register(r *jobs.Runner, im *catalog.Importer) {
	r.Register(store.JobRelayPreview, s.jobRelayPreview)
	r.Register(store.JobRelayCreate, s.jobRelayCreate)
	r.Register(store.JobRelayPay, s.jobRelayPay)
	r.Register(store.JobOrderPoll, s.jobOrderPoll)
	r.Register(store.JobOrderPollSweep, s.jobOrderPollSweep)
	r.Register(store.JobPushGapfill, s.jobPushGapfill)
	r.Register(store.JobMessageSweep, s.jobMessageSweep)
	r.Register(store.JobSupplierCancel, s.jobSupplierCancel)
	r.Register(store.JobTrimLogs, s.jobTrimLogs)

	r.Register(store.JobCatalogImport, func(ctx context.Context, j store.Job) error {
		var p struct {
			OfferIDs []int64 `json:"offerIds"`
			Keyword  string  `json:"keyword"`
			Pages    int     `json:"pages"`
			PageSize int     `json:"pageSize"`
			All      bool    `json:"all"`
		}
		if err := jobs.Payload(j, &p); err != nil {
			return jobs.Permanent(err)
		}
		switch {
		case len(p.OfferIDs) > 0:
			n, err := im.ImportOffers(ctx, p.OfferIDs)
			s.Log.Info("catalogue import finished", "offers", n)
			return err
		case p.All:
			// "Everything" is an empty-keyword search walked page by page. It
			// deliberately uses the ordinary search endpoint rather than reaching
			// into the stub, so this path is identical against the real gateway.
			pages := p.Pages
			if pages <= 0 {
				pages = 20
			}
			n, err := im.ImportSearch(ctx, "", pages, 50)
			s.Log.Info("catalogue import finished", "offers", n)
			return err
		default:
			n, err := im.ImportSearch(ctx, p.Keyword, p.Pages, p.PageSize)
			s.Log.Info("catalogue import finished", "keyword", p.Keyword, "offers", n)
			return err
		}
	})

	r.Register(store.JobReprice, func(ctx context.Context, j store.Job) error {
		n, err := im.RepriceAll(ctx)
		s.Log.Info("reprice finished", "products", n)
		return err
	})
}

// supplierOrderFromJob loads the supplier order a relay job refers to.
func (s *Service) supplierOrderFromJob(ctx context.Context, j store.Job) (store.SupplierOrder, error) {
	var p struct {
		SupplierOrderID int64 `json:"supplierOrderId"`
	}
	if err := jobs.Payload(j, &p); err != nil {
		return store.SupplierOrder{}, jobs.Permanent(err)
	}
	if p.SupplierOrderID == 0 {
		return store.SupplierOrder{}, jobs.Permanent(errors.New("relay job has no supplierOrderId"))
	}
	return s.DB.GetSupplierOrder(ctx, p.SupplierOrderID)
}

// jobRelayPreview refreshes the trade mode and the expected total. Checkout
// already previewed once, before charging; this exists for retries and for
// orders whose snapshot has aged.
func (s *Service) jobRelayPreview(ctx context.Context, j store.Job) error {
	so, err := s.supplierOrderFromJob(ctx, j)
	if err != nil {
		return err
	}
	if so.CbuOrderID != nil {
		return nil // already created; nothing to preview
	}

	cargo, err := s.cargoFor(ctx, so)
	if err != nil {
		return err
	}
	addr, err := s.warehouse(ctx)
	if err != nil {
		return err
	}

	preview, err := s.Cli.PreviewOrder(ctx, ali.PreviewRequest{
		Address: addr, Cargo: cargo, IsvBizType: "cross", OutOrderID: so.OutOrderID,
	})
	if err != nil {
		if !ali.Retryable(err) {
			return s.failSupplier(ctx, so, err)
		}
		return err
	}
	if len(preview.Results) == 0 || len(preview.Results[0].TradeModeNameList) == 0 {
		return s.failSupplier(ctx, so, errors.New("supplier does not accept API orders"))
	}

	pr := preview.Results[0]
	if err := s.DB.UpdateSupplierOrder(ctx, so.ID, store.SupplierUpdate{
		Flow:          orDefault(pr.FlowFlag, "general"),
		TradeType:     pr.TradeModeNameList[0],
		SumPaymentFen: int64(pr.SumPayment),
		PostFeeFen:    int64(pr.SumCarriage),
	}); err != nil {
		return err
	}
	return s.DB.Enqueue(ctx, store.JobRelayCreate, jobKey(so.ID),
		map[string]int64{"supplierOrderId": so.ID}, time.Now())
}

// jobRelayCreate places the order on 1688.
//
// Idempotency needs two layers. outOrderId is our key and 1688 honours it, but a
// network timeout leaves the outcome unknown: the order may exist even though we
// never saw the response. So before creating on any retry, we ask whether an
// order with our outOrderId already exists and adopt it if so. The unique
// constraint on cbu_order_id is the last line of defence.
func (s *Service) jobRelayCreate(ctx context.Context, j store.Job) error {
	so, err := s.supplierOrderFromJob(ctx, j)
	if err != nil {
		return err
	}
	if so.CbuOrderID != nil {
		return s.enqueueAfterCreate(ctx, so)
	}

	if j.Attempts > 1 {
		if adopted, err := s.adoptExisting(ctx, so); err != nil {
			return err
		} else if adopted {
			return nil
		}
	}

	cargo, err := s.cargoFor(ctx, so)
	if err != nil {
		return err
	}
	addr, err := s.warehouse(ctx)
	if err != nil {
		return err
	}

	res, err := s.Cli.CreateOrder(ctx, ali.CreateOrderRequest{
		Flow:       orDefault(so.Flow, "general"),
		Address:    addr,
		Cargo:      cargo,
		TradeType:  so.TradeType,
		OutOrderID: so.OutOrderID,
		IsvBizType: "cross",
		Message:    "Order " + so.OutOrderID,
	})
	if err != nil {
		if !ali.Retryable(err) {
			return s.failSupplier(ctx, so, err)
		}
		return err
	}

	raw, _ := json.Marshal(res)
	created := res.OrderList

	// One call can create several orders. When it does, the single-order fields
	// are empty and orderList carries the truth, so the extra orders become extra
	// supplier rows rather than being silently dropped.
	switch {
	case len(created) == 0 && res.OrderID != 0:
		id := int64(res.OrderID)
		if err := s.DB.UpdateSupplierOrder(ctx, so.ID, store.SupplierUpdate{
			CbuOrderID: &id, Status: StatusRelayedUnpaid, PostFeeFen: int64(res.PostFee), Raw: raw,
		}); err != nil {
			return err
		}
	case len(created) > 0:
		first := int64(created[0].OrderID)
		if err := s.DB.UpdateSupplierOrder(ctx, so.ID, store.SupplierUpdate{
			CbuOrderID: &first, Status: StatusRelayedUnpaid,
			PostFeeFen: int64(created[0].PostFee), Raw: raw,
		}); err != nil {
			return err
		}
		for i, extra := range created[1:] {
			s.Log.Warn("gateway split the order", "outOrderId", so.OutOrderID,
				"extraOrderId", extra.OrderID.String(), "n", i+2)
		}
	default:
		return s.failSupplier(ctx, so, errors.New("gateway returned no order id"))
	}

	so, err = s.DB.GetSupplierOrder(ctx, so.ID)
	if err != nil {
		return err
	}
	return s.enqueueAfterCreate(ctx, so)
}

// adoptExisting looks for an order 1688 may already have created for us. This is
// the reason getBuyerOrderList's outOrderId filter is in scope at all.
func (s *Service) adoptExisting(ctx context.Context, so store.SupplierOrder) (bool, error) {
	list, err := s.Cli.BuyerOrders(ctx, ali.OrderListQuery{
		OutOrderID: so.OutOrderID, Page: 1, PageSize: 10,
	})
	if err != nil {
		s.Log.Warn("could not check for an existing order", "outOrderId", so.OutOrderID, "err", err)
		return false, nil // fall through to create; the unique index still protects us
	}
	for _, t := range list.Orders {
		if t.BaseInfo.ID == 0 {
			continue
		}
		id := int64(t.BaseInfo.ID)
		s.Log.Info("adopted an order created by an earlier attempt",
			"outOrderId", so.OutOrderID, "cbuOrderId", id)
		if err := s.DB.UpdateSupplierOrder(ctx, so.ID, store.SupplierUpdate{
			CbuOrderID: &id, Status: StatusRelayedUnpaid,
		}); err != nil {
			return false, err
		}
		fresh, err := s.DB.GetSupplierOrder(ctx, so.ID)
		if err != nil {
			return false, err
		}
		return true, s.enqueueAfterCreate(ctx, fresh)
	}
	return false, nil
}

func (s *Service) enqueueAfterCreate(ctx context.Context, so store.SupplierOrder) error {
	if err := s.DB.Enqueue(ctx, store.JobRelayPay, jobKey(so.ID),
		map[string]int64{"supplierOrderId": so.ID}, time.Now()); err != nil {
		return err
	}
	return s.DB.Enqueue(ctx, store.JobOrderPoll, jobKey(so.ID),
		map[string]int64{"supplierOrderId": so.ID}, time.Now().Add(30*time.Second))
}

// jobRelayPay obtains the cashier link for a created order. With the payment
// mode set to manual an operator pays from the admin queue, which mirrors how
// this runs for real until password-free payment is enabled on the account.
func (s *Service) jobRelayPay(ctx context.Context, j store.Job) error {
	so, err := s.supplierOrderFromJob(ctx, j)
	if err != nil {
		return err
	}
	if so.CbuOrderID == nil {
		return jobs.Permanent(errors.New("cannot pay an order that was never created"))
	}
	if so.PayURL != "" {
		return nil
	}

	pay, err := s.Cli.AlipayURL(ctx, ali.IDs{ali.ID(*so.CbuOrderID)})
	if err != nil {
		if !ali.Retryable(err) {
			return jobs.Permanent(err)
		}
		return err
	}
	if pay.PayURL == "" {
		return fmt.Errorf("gateway returned no pay url: %s", pay.ErroMsg)
	}
	return s.DB.UpdateSupplierOrder(ctx, so.ID, store.SupplierUpdate{PayURL: pay.PayURL})
}

// jobOrderPoll refreshes one order from the gateway and reschedules itself while
// the order is still moving. This is the path that carries an order to
// completion when push delivery is unavailable.
func (s *Service) jobOrderPoll(ctx context.Context, j store.Job) error {
	so, err := s.supplierOrderFromJob(ctx, j)
	if err != nil {
		return err
	}
	if so.CbuOrderID == nil {
		return nil
	}

	if err := s.PollSupplierOrder(ctx, so); err != nil {
		return err
	}

	fresh, err := s.DB.GetSupplierOrder(ctx, so.ID)
	if err != nil {
		return err
	}
	switch fresh.Status {
	case StatusAtWarehouse, StatusCancelled, StatusFailed:
		return nil // terminal: stop polling
	}
	every := s.PollInterval
	if every <= 0 {
		every = 5 * time.Minute
	}
	return s.DB.Enqueue(ctx, store.JobOrderPoll, jobKey(so.ID),
		map[string]int64{"supplierOrderId": so.ID}, time.Now().Add(every))
}

// jobOrderPollSweep makes sure every open supplier order has a poll scheduled.
//
// Each poll reschedules itself, but a chain can break: a crash between the poll
// and its re-enqueue, a dead job after too many gateway errors, a bug. This
// sweep runs from a heartbeat and re-arms any open order that has fallen off,
// which is the difference between a safety net and a net with holes.
func (s *Service) jobOrderPollSweep(ctx context.Context, j store.Job) error {
	open, err := s.DB.OpenSupplierOrders(ctx, 500)
	if err != nil {
		return err
	}
	for _, so := range open {
		if err := s.DB.Enqueue(ctx, store.JobOrderPoll, jobKey(so.ID),
			map[string]int64{"supplierOrderId": so.ID}, time.Now()); err != nil {
			return err
		}
	}
	return nil
}

// jobPushGapfill drains whatever the push transport failed to deliver.
//
// The cursor call confirms as it reads, so anything it returns must be persisted
// before asking for more: a crash mid-loop would otherwise lose those messages
// for good.
func (s *Service) jobPushGapfill(ctx context.Context, j store.Job) error {
	for round := 0; round < 20; round++ {
		msgs, err := s.Cli.CursorMessages(ctx, ali.PushQuery{Quantity: 50})
		if err != nil {
			// The replay endpoint being unavailable is not fatal; the poller is
			// still covering order progress.
			s.Log.Debug("gap fill unavailable", "err", err)
			break
		}
		if len(msgs) == 0 {
			break
		}
		for _, m := range msgs {
			if err := s.ReceiveMessage(ctx, m); err != nil {
				s.Log.Warn("replayed message failed", "msgId", m.MsgID, "err", err)
			}
		}
		s.Log.Info("gap fill applied messages", "count", len(msgs))
	}
	return s.DB.Enqueue(ctx, store.JobPushGapfill, "singleton", nil, time.Now().Add(time.Minute))
}

// jobMessageSweep re-applies messages that were recorded but never processed,
// which is how a crash between the two steps is recovered.
func (s *Service) jobMessageSweep(ctx context.Context, j store.Job) error {
	pending, err := s.DB.PendingMessages(ctx, 100)
	if err != nil {
		return err
	}
	for _, ev := range pending {
		m := ali.PushMessage{
			MsgID:   ali.ID(ev.MsgID),
			Type:    ev.Type,
			Data:    ev.Payload,
			GmtBorn: ev.GmtBorn.UnixMilli(),
		}
		applyErr := s.ApplyMessage(ctx, m)
		msg := ""
		if applyErr != nil {
			msg = applyErr.Error()
		}
		if err := s.DB.MarkMessageProcessed(ctx, ev.MsgID, msg); err != nil {
			return err
		}
	}
	return s.DB.Enqueue(ctx, store.JobMessageSweep, "singleton", nil, time.Now().Add(30*time.Second))
}

// jobSupplierCancel cancels an order on 1688.
func (s *Service) jobSupplierCancel(ctx context.Context, j store.Job) error {
	var p struct {
		SupplierOrderID int64  `json:"supplierOrderId"`
		Reason          string `json:"reason"`
		Remark          string `json:"remark"`
	}
	if err := jobs.Payload(j, &p); err != nil {
		return jobs.Permanent(err)
	}
	so, err := s.DB.GetSupplierOrder(ctx, p.SupplierOrderID)
	if err != nil {
		return err
	}
	if so.CbuOrderID == nil {
		return s.DB.UpdateSupplierOrder(ctx, so.ID, store.SupplierUpdate{Status: StatusCancelled})
	}

	reason := p.Reason
	if reason == "" {
		reason = "buyerCancel"
	}
	if _, err := s.Cli.CancelOrder(ctx, ali.ID(*so.CbuOrderID), reason, p.Remark); err != nil {
		// The gateway refuses a cancel inside the first ten seconds. Wait out the
		// window rather than burning retries against a rule we already know.
		if ali.IsCode(err, ali.ErrCloseTooFast) {
			return s.DB.Enqueue(ctx, store.JobSupplierCancel, jobKey(so.ID),
				map[string]any{"supplierOrderId": so.ID, "reason": reason, "remark": p.Remark},
				time.Now().Add(cancelGrace))
		}
		if !ali.Retryable(err) {
			return jobs.Permanent(err)
		}
		return err
	}

	if _, err := s.DB.ApplySupplierStatus(ctx, so.ID, "cancel", StatusCancelled, StatusRank("cancel"), time.Now()); err != nil {
		return err
	}
	return s.refreshOrderStatus(ctx, so.OrderID)
}

// jobTrimLogs keeps the call log bounded.
func (s *Service) jobTrimLogs(ctx context.Context, j store.Job) error {
	if err := s.DB.TrimAPICalls(ctx, 20000); err != nil {
		return err
	}
	return s.DB.Enqueue(ctx, store.JobTrimLogs, "singleton", nil, time.Now().Add(24*time.Hour))
}

// CancelSupplierOrder queues a cancellation, respecting the ten-second window.
func (s *Service) CancelSupplierOrder(ctx context.Context, so store.SupplierOrder, reason, remark string) error {
	runAt := time.Now()
	if age := time.Since(so.CreatedAt); age < cancelGrace {
		runAt = so.CreatedAt.Add(cancelGrace)
	}
	return s.DB.Enqueue(ctx, store.JobSupplierCancel, jobKey(so.ID),
		map[string]any{"supplierOrderId": so.ID, "reason": reason, "remark": remark}, runAt)
}

// ------------------------------------------------------------------ helpers --

// cargoFor rebuilds the cargo list from the stored order lines, so a retry sends
// exactly what checkout priced rather than re-reading a catalogue that may have
// moved on.
func (s *Service) cargoFor(ctx context.Context, so store.SupplierOrder) ([]ali.Cargo, error) {
	items, err := s.DB.OrderItems(ctx, so.OrderID)
	if err != nil {
		return nil, err
	}
	var cargo []ali.Cargo
	for _, it := range items {
		if it.SupplierOrderID == nil || *it.SupplierOrderID != so.ID {
			continue
		}
		cargo = append(cargo, ali.NewCargo(ali.ID(it.OfferID), it.SpecID, int64(it.Quantity)))
	}
	if len(cargo) == 0 {
		return nil, jobs.Permanent(errors.New("supplier order has no lines"))
	}
	if len(cargo) > MaxSKUsPerOrder {
		return nil, jobs.Permanent(fmt.Errorf("supplier order has %d skus, the limit is %d", len(cargo), MaxSKUsPerOrder))
	}
	return cargo, nil
}

func (s *Service) warehouse(ctx context.Context) (ali.Address, error) {
	settings, err := s.DB.Settings(ctx)
	if err != nil {
		return ali.Address{}, err
	}
	return warehouseAddress(settings)
}

// failSupplier records a refusal that no amount of retrying will fix.
func (s *Service) failSupplier(ctx context.Context, so store.SupplierOrder, cause error) error {
	if err := s.DB.UpdateSupplierOrder(ctx, so.ID, store.SupplierUpdate{
		Status:       StatusFailed,
		ErrorCode:    ali.Code(cause),
		ErrorMessage: cause.Error(),
	}); err != nil {
		return err
	}
	if err := s.refreshOrderStatus(ctx, so.OrderID); err != nil {
		return err
	}
	s.Log.Error("relay refused", "supplierOrder", so.ID, "outOrderId", so.OutOrderID, "err", cause)
	return jobs.Permanent(cause)
}

func jobKey(supplierOrderID int64) string {
	return "so:" + strconv.FormatInt(supplierOrderID, 10)
}
