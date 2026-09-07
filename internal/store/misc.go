package store

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
)

// ------------------------------------------------------------------ settings --

// Settings reads every setting as a map.
func (db *DB) Settings(ctx context.Context) (map[string]string, error) {
	rows, err := db.Pool.Query(ctx, `SELECT key, value FROM settings`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := map[string]string{}
	for rows.Next() {
		var k, v string
		if err := rows.Scan(&k, &v); err != nil {
			return nil, err
		}
		out[k] = v
	}
	return out, rows.Err()
}

// SetSetting writes one setting.
func (db *DB) SetSetting(ctx context.Context, key, value string) error {
	_, err := db.Pool.Exec(ctx, `
		INSERT INTO settings (key, value) VALUES ($1,$2)
		ON CONFLICT (key) DO UPDATE SET value=EXCLUDED.value`, key, value)
	return err
}

// ----------------------------------------------------------------- fee rules --

// FeeRules lists every rule, newest and highest priority first.
func (db *DB) FeeRules(ctx context.Context) ([]FeeRule, error) {
	rows, err := db.Pool.Query(ctx, `SELECT * FROM fee_rules ORDER BY priority DESC, id DESC`)
	if err != nil {
		return nil, err
	}
	return pgx.CollectRows(rows, pgx.RowToStructByNameLax[FeeRule])
}

// CreateFeeRule inserts a rule and returns it.
func (db *DB) CreateFeeRule(ctx context.Context, r FeeRule) (FeeRule, error) {
	rows, err := db.Pool.Query(ctx, `
		INSERT INTO fee_rules (scope, scope_value, fee_bps, fee_fixed_fen, min_fee_fen, priority,
			effective_from, effective_to, note)
		VALUES ($1,$2,$3,$4,$5,$6, coalesce($7, now()), $8, $9)
		RETURNING *`,
		r.Scope, r.ScopeValue, r.FeeBps, r.FeeFixedFen, r.MinFeeFen, r.Priority,
		nullTime(r.EffectiveFrom), r.EffectiveTo, r.Note)
	if err != nil {
		return FeeRule{}, err
	}
	return pgx.CollectExactlyOneRow(rows, pgx.RowToStructByNameLax[FeeRule])
}

// UpdateFeeRule replaces a rule's editable fields.
func (db *DB) UpdateFeeRule(ctx context.Context, r FeeRule) (FeeRule, error) {
	rows, err := db.Pool.Query(ctx, `
		UPDATE fee_rules SET scope=$2, scope_value=$3, fee_bps=$4, fee_fixed_fen=$5,
			min_fee_fen=$6, priority=$7, effective_from=coalesce($8, effective_from),
			effective_to=$9, note=$10
		 WHERE id=$1 RETURNING *`,
		r.ID, r.Scope, r.ScopeValue, r.FeeBps, r.FeeFixedFen, r.MinFeeFen, r.Priority,
		nullTime(r.EffectiveFrom), r.EffectiveTo, r.Note)
	if err != nil {
		return FeeRule{}, err
	}
	fr, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByNameLax[FeeRule])
	if errors.Is(err, pgx.ErrNoRows) {
		return FeeRule{}, ErrNotFound
	}
	return fr, err
}

// DeleteFeeRule removes a rule.
func (db *DB) DeleteFeeRule(ctx context.Context, id int64) error {
	_, err := db.Pool.Exec(ctx, `DELETE FROM fee_rules WHERE id=$1`, id)
	return err
}

func nullTime(t time.Time) *time.Time {
	if t.IsZero() {
		return nil
	}
	return &t
}

// ---------------------------------------------------------------------- cart --

// EnsureCart creates the cart row if it is new.
func (db *DB) EnsureCart(ctx context.Context, id string) error {
	_, err := db.Pool.Exec(ctx, `
		INSERT INTO carts (id) VALUES ($1)
		ON CONFLICT (id) DO UPDATE SET updated_at=now()`, id)
	return err
}

// CartItems lists one cart's lines.
func (db *DB) CartItems(ctx context.Context, cartID string) ([]CartItem, error) {
	rows, err := db.Pool.Query(ctx, `SELECT * FROM cart_items WHERE cart_id=$1 ORDER BY added_at, id`, cartID)
	if err != nil {
		return nil, err
	}
	return pgx.CollectRows(rows, pgx.RowToStructByNameLax[CartItem])
}

// AddCartItem adds a line, or increases the quantity when the SKU is already there.
func (db *DB) AddCartItem(ctx context.Context, it CartItem) error {
	_, err := db.Pool.Exec(ctx, `
		INSERT INTO cart_items (cart_id, offer_id, sku_id, spec_id, quantity)
		VALUES ($1,$2,$3,$4,$5)
		ON CONFLICT (cart_id, sku_id) DO UPDATE SET quantity = cart_items.quantity + EXCLUDED.quantity`,
		it.CartID, it.OfferID, it.SkuID, it.SpecID, it.Quantity)
	return err
}

// SetCartItemQty changes one line's quantity.
func (db *DB) SetCartItemQty(ctx context.Context, cartID string, itemID int64, qty int) error {
	_, err := db.Pool.Exec(ctx,
		`UPDATE cart_items SET quantity=$3 WHERE id=$2 AND cart_id=$1`, cartID, itemID, qty)
	return err
}

// DeleteCartItem removes one line.
func (db *DB) DeleteCartItem(ctx context.Context, cartID string, itemID int64) error {
	_, err := db.Pool.Exec(ctx, `DELETE FROM cart_items WHERE id=$2 AND cart_id=$1`, cartID, itemID)
	return err
}

// ClearCart empties a cart, which happens once its order is placed.
func (db *DB) ClearCart(ctx context.Context, cartID string) error {
	_, err := db.Pool.Exec(ctx, `DELETE FROM cart_items WHERE cart_id=$1`, cartID)
	return err
}

// ------------------------------------------------------------------ messages --

// RecordMessage stores an inbound push message. It reports whether the row was
// new: the primary key is 1688's own msgId, so duplicate delivery is free to
// ignore and needs no separate dedupe table.
func (db *DB) RecordMessage(ctx context.Context, msgID int64, typ string, gmtBorn time.Time, payload json.RawMessage) (bool, error) {
	if len(payload) == 0 {
		payload = []byte("{}")
	}
	if gmtBorn.IsZero() {
		gmtBorn = time.Now()
	}
	tag, err := db.Pool.Exec(ctx, `
		INSERT INTO message_events (msg_id, type, gmt_born, payload)
		VALUES ($1,$2,$3,$4)
		ON CONFLICT (msg_id) DO NOTHING`, msgID, typ, gmtBorn, payload)
	if err != nil {
		return false, err
	}
	return tag.RowsAffected() > 0, nil
}

// MarkMessageProcessed closes off one message, recording any handler error.
func (db *DB) MarkMessageProcessed(ctx context.Context, msgID int64, handlerErr string) error {
	_, err := db.Pool.Exec(ctx,
		`UPDATE message_events SET processed_at=now(), error=$2 WHERE msg_id=$1`, msgID, trim(handlerErr, 1000))
	return err
}

// PendingMessages lists messages recorded but not yet processed, which is how a
// crash between the insert and the handler is recovered.
func (db *DB) PendingMessages(ctx context.Context, limit int) ([]MessageEvent, error) {
	if limit <= 0 || limit > 500 {
		limit = 100
	}
	rows, err := db.Pool.Query(ctx, `
		SELECT * FROM message_events WHERE processed_at IS NULL ORDER BY received_at LIMIT $1`, limit)
	if err != nil {
		return nil, err
	}
	return pgx.CollectRows(rows, pgx.RowToStructByNameLax[MessageEvent])
}

// RecentMessages is the admin message log.
func (db *DB) RecentMessages(ctx context.Context, limit int) ([]MessageEvent, error) {
	if limit <= 0 || limit > 500 {
		limit = 100
	}
	rows, err := db.Pool.Query(ctx,
		`SELECT * FROM message_events ORDER BY received_at DESC LIMIT $1`, limit)
	if err != nil {
		return nil, err
	}
	return pgx.CollectRows(rows, pgx.RowToStructByNameLax[MessageEvent])
}

// ----------------------------------------------------------------- api calls --

// LogAPICall records one gateway call for the admin integration screen.
func (db *DB) LogAPICall(ctx context.Context, c APICall) error {
	req := c.Req
	if len(req) == 0 {
		req = []byte("{}")
	}
	_, err := db.Pool.Exec(ctx, `
		INSERT INTO api_calls (api, ms, ok, code, message, req, resp)
		VALUES ($1,$2,$3,$4,$5,$6,$7)`,
		c.API, c.MS, c.OK, c.Code, trim(c.Message, 500), req, trim(c.Resp, 4000))
	return err
}

// RecentAPICalls is the admin call log.
func (db *DB) RecentAPICalls(ctx context.Context, limit int) ([]APICall, error) {
	if limit <= 0 || limit > 500 {
		limit = 100
	}
	rows, err := db.Pool.Query(ctx, `SELECT * FROM api_calls ORDER BY created_at DESC LIMIT $1`, limit)
	if err != nil {
		return nil, err
	}
	return pgx.CollectRows(rows, pgx.RowToStructByNameLax[APICall])
}

// TrimAPICalls keeps the log bounded.
func (db *DB) TrimAPICalls(ctx context.Context, keep int) error {
	_, err := db.Pool.Exec(ctx, `
		DELETE FROM api_calls WHERE id < (
			SELECT coalesce(min(id),0) FROM (
				SELECT id FROM api_calls ORDER BY id DESC LIMIT $1
			) t)`, keep)
	return err
}
