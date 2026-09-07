package store

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
)

// Job kinds. Each one has a handler registered in internal/jobs.
const (
	JobRelayPreview   = "relay_preview"
	JobRelayCreate    = "relay_create"
	JobRelayPay       = "relay_pay"
	JobOrderPoll      = "order_poll"
	JobOrderPollSweep = "order_poll_sweep"
	JobPushGapfill    = "push_gapfill"
	JobMessageSweep   = "message_sweep"
	JobCatalogImport  = "catalog_import"
	JobReprice        = "reprice"
	JobSupplierCancel = "supplier_cancel"
	JobTrimLogs       = "trim_logs"
)

// Enqueue adds a job unless an identical one is already scheduled. The partial
// unique index on (kind, key) over pending rows makes this idempotent without a
// read-then-write race, while still letting a running job line up its own
// successor.
func (db *DB) Enqueue(ctx context.Context, kind, key string, payload any, runAt time.Time) error {
	body := []byte("{}")
	if payload != nil {
		b, err := json.Marshal(payload)
		if err != nil {
			return err
		}
		body = b
	}
	if runAt.IsZero() {
		runAt = time.Now()
	}
	_, err := db.Pool.Exec(ctx, `
		INSERT INTO jobs (kind, key, payload, run_at)
		VALUES ($1,$2,$3,$4)
		ON CONFLICT DO NOTHING`, kind, key, body, runAt)
	return err
}

// ClaimJob takes the next ready job. FOR UPDATE SKIP LOCKED lets several workers,
// in this process or in another container, claim concurrently with no lease
// bookkeeping: a crashed worker's row unlocks when its connection drops.
//
// It returns nil when there is nothing to do.
func (db *DB) ClaimJob(ctx context.Context) (*Job, error) {
	tx, err := db.Pool.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer func() { _ = tx.Rollback(ctx) }()

	rows, err := tx.Query(ctx, `
		SELECT * FROM jobs
		 WHERE state='pending' AND run_at <= now()
		 ORDER BY run_at
		 FOR UPDATE SKIP LOCKED
		 LIMIT 1`)
	if err != nil {
		return nil, err
	}
	job, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByNameLax[Job])
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	if _, err := tx.Exec(ctx, `
		UPDATE jobs SET state='running', attempts=attempts+1, updated_at=now() WHERE id=$1`,
		job.ID); err != nil {
		return nil, err
	}
	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}
	job.Attempts++
	job.State = "running"
	return &job, nil
}

// RequeueStaleJobs returns jobs that have sat in the running state for longer
// than olderThan back to pending, and reports how many it touched.
//
// The claim commits state='running' before the handler executes, so the row is
// durable rather than locked while the work happens. That is what lets several
// processes share the queue, but it also means a process killed mid-job leaves
// the row running forever, and because only one live job per (kind, key) may
// exist, a stranded singleton silently blocks every future enqueue of its kind.
// The runner calls this at startup and on a slow tick.
func (db *DB) RequeueStaleJobs(ctx context.Context, olderThan time.Duration) (int64, error) {
	tag, err := db.Pool.Exec(ctx, `
		UPDATE jobs SET state='pending', run_at=now(), updated_at=now(),
		       last_error='requeued: worker did not finish (' || last_error || ')'
		 WHERE state='running' AND updated_at < now() - $1::interval`, olderThan.String())
	if err != nil {
		return 0, err
	}
	return tag.RowsAffected(), nil
}

// FinishJob marks a claimed job done.
func (db *DB) FinishJob(ctx context.Context, id int64) error {
	_, err := db.Pool.Exec(ctx,
		`UPDATE jobs SET state='done', last_error='', updated_at=now() WHERE id=$1`, id)
	return err
}

// RetryJob puts a failed job back with a delay.
func (db *DB) RetryJob(ctx context.Context, id int64, cause string, in time.Duration) error {
	_, err := db.Pool.Exec(ctx, `
		UPDATE jobs SET state='pending', run_at=now()+$3::interval, last_error=$2, updated_at=now()
		 WHERE id=$1`, id, trim(cause, 2000), in.String())
	return err
}

// KillJob marks a job dead, which surfaces it in the admin jobs tab with a retry
// button rather than retrying forever.
func (db *DB) KillJob(ctx context.Context, id int64, cause string) error {
	_, err := db.Pool.Exec(ctx,
		`UPDATE jobs SET state='dead', last_error=$2, updated_at=now() WHERE id=$1`, id, trim(cause, 2000))
	return err
}

// ReviveJob puts a dead job back in the queue.
func (db *DB) ReviveJob(ctx context.Context, id int64) error {
	_, err := db.Pool.Exec(ctx, `
		UPDATE jobs SET state='pending', attempts=0, run_at=now(), last_error='', updated_at=now()
		 WHERE id=$1`, id)
	return err
}

// ListJobs is the admin jobs view.
func (db *DB) ListJobs(ctx context.Context, state string, limit int) ([]Job, error) {
	if limit <= 0 || limit > 500 {
		limit = 100
	}
	sql := `SELECT * FROM jobs`
	args := []any{}
	if state != "" {
		sql += ` WHERE state=$1`
		args = append(args, state)
	}
	sql += ` ORDER BY updated_at DESC LIMIT ` + itoa(limit)
	rows, err := db.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	return pgx.CollectRows(rows, pgx.RowToStructByNameLax[Job])
}

// GetJob loads one job.
func (db *DB) GetJob(ctx context.Context, id int64) (Job, error) {
	rows, err := db.Pool.Query(ctx, `SELECT * FROM jobs WHERE id=$1`, id)
	if err != nil {
		return Job{}, err
	}
	j, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByNameLax[Job])
	if errors.Is(err, pgx.ErrNoRows) {
		return Job{}, ErrNotFound
	}
	return j, err
}

func trim(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n] + "…"
}

func itoa(v int) string {
	const digits = "0123456789"
	if v == 0 {
		return "0"
	}
	neg := v < 0
	if neg {
		v = -v
	}
	var buf [20]byte
	i := len(buf)
	for v > 0 {
		i--
		buf[i] = digits[v%10]
		v /= 10
	}
	if neg {
		i--
		buf[i] = '-'
	}
	return string(buf[i:])
}
