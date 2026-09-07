package store

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

// ErrNotFound is returned when a lookup finds nothing.
var ErrNotFound = errors.New("store: not found")

// NewOrder is everything needed to write one customer order and its supplier
// groups in a single transaction.
type NewOrder struct {
	AccessToken   string
	CartID        string
	Email         string
	ShipName      string
	ShipPhone     string
	ShipAddress   json.RawMessage
	FxPpm         int64
	GoodsSatang   int64
	FreightSatang int64
	IntlSatang    int64
	FeeSatang     int64
	TotalSatang   int64
	Groups        []NewSupplierGroup
}

// NewSupplierGroup is one 1688 order to be: a single supplier, at most fifty SKUs.
type NewSupplierGroup struct {
	SellerOpenID string
	SellerName   string
	GroupSeq     int32
	Items        []OrderItem
}

// CreateOrder writes the order, its supplier groups and its line items together.
// The public id comes from a sequence, and out_order_id is derived from it: that
// string is what 1688 receives as outOrderId and is our idempotency key.
func (db *DB) CreateOrder(ctx context.Context, in NewOrder) (Order, error) {
	var out Order

	tx, err := db.Pool.Begin(ctx)
	if err != nil {
		return out, err
	}
	defer func() { _ = tx.Rollback(ctx) }()

	var publicID string
	if err := tx.QueryRow(ctx, `
		SELECT 'MK-' || to_char(now(), 'YYYY') || '-' || lpad(nextval('order_public_seq')::text, 6, '0')`).
		Scan(&publicID); err != nil {
		return out, fmt.Errorf("order number: %w", err)
	}

	addr := in.ShipAddress
	if len(addr) == 0 {
		addr = []byte("{}")
	}
	var orderID int64
	if err := tx.QueryRow(ctx, `
		INSERT INTO orders (public_id, access_token, cart_id, email, ship_name, ship_phone,
			ship_address, status, fx_ppm, goods_satang, freight_satang, intl_satang, fee_satang,
			total_satang)
		VALUES ($1,$2,$3,$4,$5,$6,$7,'AWAITING_PAYMENT',$8,$9,$10,$11,$12,$13)
		RETURNING id`,
		publicID, in.AccessToken, in.CartID, in.Email, in.ShipName, in.ShipPhone, addr,
		in.FxPpm, in.GoodsSatang, in.FreightSatang, in.IntlSatang, in.FeeSatang, in.TotalSatang).
		Scan(&orderID); err != nil {
		return out, fmt.Errorf("insert order: %w", err)
	}

	for _, g := range in.Groups {
		outOrderID := fmt.Sprintf("%s-%d", publicID, g.GroupSeq)
		var soID int64
		if err := tx.QueryRow(ctx, `
			INSERT INTO supplier_orders (order_id, seller_open_id, seller_name, group_seq,
				out_order_id, status)
			VALUES ($1,$2,$3,$4,$5,'PENDING') RETURNING id`,
			orderID, g.SellerOpenID, g.SellerName, g.GroupSeq, outOrderID).Scan(&soID); err != nil {
			return out, fmt.Errorf("insert supplier order: %w", err)
		}

		for _, it := range g.Items {
			snap := it.Snapshot
			if len(snap) == 0 {
				snap = []byte("{}")
			}
			if _, err := tx.Exec(ctx, `
				INSERT INTO order_items (order_id, supplier_order_id, offer_id, sku_id, spec_id,
					title, image_url, sku_label, quantity, base_fen, freight_fen, intl_fen, fee_fen,
					fee_rule_id, fx_ppm, unit_satang, line_satang, snapshot, seller_open_id)
				VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19)`,
				orderID, soID, it.OfferID, it.SkuID, it.SpecID, it.Title, it.ImageURL, it.SkuLabel,
				it.Quantity, it.BaseFen, it.FreightFen, it.IntlFen, it.FeeFen, it.FeeRuleID,
				it.FxPpm, it.UnitSatang, it.LineSatang, snap, g.SellerOpenID); err != nil {
				return out, fmt.Errorf("insert order item: %w", err)
			}
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return out, err
	}
	return db.GetOrder(ctx, orderID)
}

// GetOrder loads one order by surrogate key.
func (db *DB) GetOrder(ctx context.Context, id int64) (Order, error) {
	rows, err := db.Pool.Query(ctx, `SELECT * FROM orders WHERE id=$1`, id)
	if err != nil {
		return Order{}, err
	}
	o, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByNameLax[Order])
	if errors.Is(err, pgx.ErrNoRows) {
		return Order{}, ErrNotFound
	}
	return o, err
}

// GetOrderByPublicID loads one order by its customer-facing number.
func (db *DB) GetOrderByPublicID(ctx context.Context, publicID string) (Order, error) {
	rows, err := db.Pool.Query(ctx, `SELECT * FROM orders WHERE public_id=$1`, publicID)
	if err != nil {
		return Order{}, err
	}
	o, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByNameLax[Order])
	if errors.Is(err, pgx.ErrNoRows) {
		return Order{}, ErrNotFound
	}
	return o, err
}

// OrderItems lists the lines of one order.
func (db *DB) OrderItems(ctx context.Context, orderID int64) ([]OrderItem, error) {
	rows, err := db.Pool.Query(ctx, `SELECT * FROM order_items WHERE order_id=$1 ORDER BY id`, orderID)
	if err != nil {
		return nil, err
	}
	return pgx.CollectRows(rows, pgx.RowToStructByNameLax[OrderItem])
}

// SupplierOrders lists the 1688-side orders belonging to one customer order.
func (db *DB) SupplierOrders(ctx context.Context, orderID int64) ([]SupplierOrder, error) {
	rows, err := db.Pool.Query(ctx, `SELECT * FROM supplier_orders WHERE order_id=$1 ORDER BY group_seq`, orderID)
	if err != nil {
		return nil, err
	}
	return pgx.CollectRows(rows, pgx.RowToStructByNameLax[SupplierOrder])
}

// GetSupplierOrder loads one supplier order.
func (db *DB) GetSupplierOrder(ctx context.Context, id int64) (SupplierOrder, error) {
	rows, err := db.Pool.Query(ctx, `SELECT * FROM supplier_orders WHERE id=$1`, id)
	if err != nil {
		return SupplierOrder{}, err
	}
	so, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByNameLax[SupplierOrder])
	if errors.Is(err, pgx.ErrNoRows) {
		return SupplierOrder{}, ErrNotFound
	}
	return so, err
}

// SupplierOrderByCbuID finds a supplier order by the 1688 order id, which is how
// an inbound push message is matched to our records.
func (db *DB) SupplierOrderByCbuID(ctx context.Context, cbuOrderID int64) (SupplierOrder, error) {
	rows, err := db.Pool.Query(ctx, `SELECT * FROM supplier_orders WHERE cbu_order_id=$1`, cbuOrderID)
	if err != nil {
		return SupplierOrder{}, err
	}
	so, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByNameLax[SupplierOrder])
	if errors.Is(err, pgx.ErrNoRows) {
		return SupplierOrder{}, ErrNotFound
	}
	return so, err
}

// MarkOrderPaid records the customer's payment and moves the order along.
func (db *DB) MarkOrderPaid(ctx context.Context, orderID int64) error {
	_, err := db.Pool.Exec(ctx, `
		UPDATE orders SET status='PAID', paid_at=coalesce(paid_at, now()), updated_at=now()
		 WHERE id=$1 AND paid_at IS NULL`, orderID)
	return err
}

// SetOrderStatus writes the rolled-up customer-facing status.
func (db *DB) SetOrderStatus(ctx context.Context, orderID int64, status string) error {
	_, err := db.Pool.Exec(ctx,
		`UPDATE orders SET status=$2, updated_at=now() WHERE id=$1 AND status<>$2`, orderID, status)
	return err
}

// SupplierUpdate carries the fields the relay worker writes after each step.
type SupplierUpdate struct {
	CbuOrderID    *int64
	Flow          string
	TradeType     string
	Status        string
	PayURL        string
	SumPaymentFen int64
	PostFeeFen    int64
	ErrorCode     string
	ErrorMessage  string
	Raw           json.RawMessage
}

// UpdateSupplierOrder writes relay progress. Only non-zero fields are applied, so
// one step never erases another's work.
func (db *DB) UpdateSupplierOrder(ctx context.Context, id int64, u SupplierUpdate) error {
	raw := u.Raw
	if len(raw) == 0 {
		raw = nil
	}
	_, err := db.Pool.Exec(ctx, `
		UPDATE supplier_orders SET
			cbu_order_id    = coalesce($2, cbu_order_id),
			flow            = coalesce(nullif($3,''), flow),
			trade_type      = coalesce(nullif($4,''), trade_type),
			status          = coalesce(nullif($5,''), status),
			pay_url         = coalesce(nullif($6,''), pay_url),
			sum_payment_fen = CASE WHEN $7::bigint <> 0 THEN $7 ELSE sum_payment_fen END,
			post_fee_fen    = CASE WHEN $8::bigint <> 0 THEN $8 ELSE post_fee_fen END,
			error_code      = $9,
			error_message   = $10,
			raw             = coalesce($11::jsonb, raw),
			updated_at      = now()
		 WHERE id=$1`,
		id, u.CbuOrderID, u.Flow, u.TradeType, u.Status, u.PayURL,
		u.SumPaymentFen, u.PostFeeFen, u.ErrorCode, u.ErrorMessage, raw)
	return err
}

// ApplySupplierStatus writes a status change only when it genuinely advances the
// order, which is what makes out-of-order push messages harmless.
//
// A message is applied when its rank is higher than what we hold, or when the
// rank ties and the event is newer. Terminal states (rank 9) always win, because
// a cancellation must not be undone by a stale in-flight event.
//
// It reports whether the update was applied; a false result is a stale event and
// is surfaced in the admin message log rather than treated as an error.
func (db *DB) ApplySupplierStatus(ctx context.Context, id int64, status1688, ourStatus string, rank int, at time.Time) (bool, error) {
	tag, err := db.Pool.Exec(ctx, `
		UPDATE supplier_orders
		   SET status_1688=$2, status=$3, status_rank=$4, status_at=$5, updated_at=now()
		 WHERE id=$1
		   AND ($4 >= 9 OR $4 > status_rank OR ($4 = status_rank AND $5 > status_at))`,
		id, status1688, ourStatus, rank, at)
	if err != nil {
		return false, err
	}
	return tag.RowsAffected() > 0, nil
}

// OpenSupplierOrders lists supplier orders that have not reached a terminal
// state, for the poller.
func (db *DB) OpenSupplierOrders(ctx context.Context, limit int) ([]SupplierOrder, error) {
	rows, err := db.Pool.Query(ctx, `
		SELECT * FROM supplier_orders
		 WHERE status NOT IN ('ARRIVED_WAREHOUSE','CANCELLED','FAILED')
		   AND cbu_order_id IS NOT NULL
		 ORDER BY updated_at
		 LIMIT $1`, limit)
	if err != nil {
		return nil, err
	}
	return pgx.CollectRows(rows, pgx.RowToStructByNameLax[SupplierOrder])
}

// ListOrders is the admin order list.
func (db *DB) ListOrders(ctx context.Context, status string, limit int) ([]Order, error) {
	if limit <= 0 || limit > 500 {
		limit = 100
	}
	sql := `SELECT * FROM orders`
	args := []any{}
	if status != "" {
		sql += ` WHERE status=$1`
		args = append(args, status)
	}
	sql += ` ORDER BY created_at DESC LIMIT ` + fmt.Sprint(limit)
	rows, err := db.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	return pgx.CollectRows(rows, pgx.RowToStructByNameLax[Order])
}

// ---------------------------------------------------------------- shipments --

// UpsertShipment creates or updates the parcel record for a supplier order.
func (db *DB) UpsertShipment(ctx context.Context, s Shipment) (int64, error) {
	var id int64
	err := db.Pool.QueryRow(ctx, `
		INSERT INTO shipments (supplier_order_id, leg, logistics_id, mail_no, cp_code,
			company_name, status, updated_at)
		VALUES ($1,$2,$3,$4,$5,$6,$7, now())
		ON CONFLICT (supplier_order_id, leg, logistics_id) DO UPDATE
		   SET mail_no      = coalesce(nullif(EXCLUDED.mail_no,''), shipments.mail_no),
		       cp_code      = coalesce(nullif(EXCLUDED.cp_code,''), shipments.cp_code),
		       company_name = coalesce(nullif(EXCLUDED.company_name,''), shipments.company_name),
		       status       = coalesce(nullif(EXCLUDED.status,''), shipments.status),
		       updated_at   = now()
		RETURNING id`,
		s.SupplierOrderID, s.Leg, s.LogisticsID, s.MailNo, s.CpCode, s.CompanyName, s.Status).Scan(&id)
	return id, err
}

// ShipmentsFor lists the parcels of several supplier orders.
func (db *DB) ShipmentsFor(ctx context.Context, supplierOrderIDs []int64) ([]Shipment, error) {
	if len(supplierOrderIDs) == 0 {
		return nil, nil
	}
	rows, err := db.Pool.Query(ctx,
		`SELECT * FROM shipments WHERE supplier_order_id = ANY($1) ORDER BY id`, supplierOrderIDs)
	if err != nil {
		return nil, err
	}
	return pgx.CollectRows(rows, pgx.RowToStructByNameLax[Shipment])
}

// AddTrackingEvent records one trace step. The unique dedupe key means push and
// polling can both report the same physical event and converge with no ordering
// logic at all; a duplicate is silently ignored.
func (db *DB) AddTrackingEvent(ctx context.Context, e TrackingEvent) (bool, error) {
	tag, err := db.Pool.Exec(ctx, `
		INSERT INTO tracking_events (shipment_id, source, event_at, code, remark, dedupe_key)
		VALUES ($1,$2,$3,$4,$5,$6)
		ON CONFLICT (shipment_id, dedupe_key) DO NOTHING`,
		e.ShipmentID, e.Source, e.EventAt, e.Code, e.Remark, e.DedupeKey)
	if err != nil {
		return false, err
	}
	return tag.RowsAffected() > 0, nil
}

// TrackingEventsFor lists trace steps for several parcels, oldest first.
func (db *DB) TrackingEventsFor(ctx context.Context, shipmentIDs []int64) ([]TrackingEvent, error) {
	if len(shipmentIDs) == 0 {
		return nil, nil
	}
	rows, err := db.Pool.Query(ctx,
		`SELECT * FROM tracking_events WHERE shipment_id = ANY($1) ORDER BY event_at, id`, shipmentIDs)
	if err != nil {
		return nil, err
	}
	return pgx.CollectRows(rows, pgx.RowToStructByNameLax[TrackingEvent])
}
