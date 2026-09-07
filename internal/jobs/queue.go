// Package jobs runs background work off a Postgres table. There is no Redis and
// no broker: a row is claimed with FOR UPDATE SKIP LOCKED, which gives several
// workers safe concurrency and needs no lease bookkeeping, because a crashed
// worker's row unlocks the moment its connection drops.
package jobs

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"sync"
	"time"

	"marketplace/internal/store"
)

// Handler runs one job. Returning an error schedules a retry with backoff until
// MaxAttempts, after which the job is marked dead and shown in the admin queue.
//
// A handler that returns ErrPermanent is not retried at all, which is what a
// documented business refusal deserves: retrying a rejected order only produces
// the same rejection.
type Handler func(ctx context.Context, job store.Job) error

// ErrPermanent wraps a failure that must not be retried.
type ErrPermanent struct{ Err error }

func (e *ErrPermanent) Error() string { return "permanent: " + e.Err.Error() }
func (e *ErrPermanent) Unwrap() error { return e.Err }

// Permanent marks an error as not worth retrying.
func Permanent(err error) error { return &ErrPermanent{Err: err} }

// Runner claims and executes jobs.
type Runner struct {
	DB          *store.DB
	Log         *slog.Logger
	Concurrency int
	Tick        time.Duration
	MaxAttempts int

	mu       sync.RWMutex
	handlers map[string]Handler
}

// New builds a runner with defaults.
func New(db *store.DB, log *slog.Logger) *Runner {
	return &Runner{
		DB:          db,
		Log:         log,
		Concurrency: 3,
		Tick:        time.Second,
		MaxAttempts: 8,
		handlers:    map[string]Handler{},
	}
}

// Register attaches a handler to a job kind.
func (r *Runner) Register(kind string, h Handler) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.handlers == nil {
		r.handlers = map[string]Handler{}
	}
	r.handlers[kind] = h
}

func (r *Runner) handler(kind string) (Handler, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	h, ok := r.handlers[kind]
	return h, ok
}

// Run works the queue until the context is cancelled. It blocks, so callers
// usually start it in a goroutine.
func (r *Runner) Run(ctx context.Context) {
	n := r.Concurrency
	if n < 1 {
		n = 1
	}
	tick := r.Tick
	if tick <= 0 {
		tick = time.Second
	}

	// Anything still marked running belongs to a process that is no longer
	// here: at startup that is certain, and later it means a worker died. A
	// singleton left in that state would otherwise block its own kind forever.
	r.requeueStale(ctx, 0)
	go func() {
		t := time.NewTicker(time.Minute)
		defer t.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-t.C:
				r.requeueStale(ctx, 10*time.Minute)
			}
		}
	}()

	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(worker int) {
			defer wg.Done()
			t := time.NewTicker(tick)
			defer t.Stop()
			for {
				select {
				case <-ctx.Done():
					return
				case <-t.C:
					// Drain rather than taking one job per tick, so a burst is
					// worked through immediately.
					for {
						did, err := r.step(ctx)
						if err != nil {
							r.Log.Error("job step failed", "worker", worker, "err", err)
							break
						}
						if !did || ctx.Err() != nil {
							break
						}
					}
				}
			}
		}(i)
	}
	wg.Wait()
}

// requeueStale hands abandoned running jobs back to the queue.
func (r *Runner) requeueStale(ctx context.Context, olderThan time.Duration) {
	n, err := r.DB.RequeueStaleJobs(ctx, olderThan)
	if err != nil {
		r.Log.Error("requeue stale jobs failed", "err", err)
		return
	}
	if n > 0 {
		r.Log.Warn("requeued abandoned jobs", "count", n, "olderThan", olderThan.String())
	}
}

// step claims and runs at most one job. It reports whether it did any work.
func (r *Runner) step(ctx context.Context) (bool, error) {
	job, err := r.DB.ClaimJob(ctx)
	if err != nil || job == nil {
		return false, err
	}

	h, ok := r.handler(job.Kind)
	if !ok {
		_ = r.DB.KillJob(ctx, job.ID, "no handler registered for kind "+job.Kind)
		r.Log.Error("job has no handler", "kind", job.Kind, "id", job.ID)
		return true, nil
	}

	started := time.Now()
	err = r.runOne(ctx, h, *job)
	if err == nil {
		if e := r.DB.FinishJob(ctx, job.ID); e != nil {
			r.Log.Error("job finish failed", "id", job.ID, "err", e)
		}
		r.Log.Debug("job done", "kind", job.Kind, "key", job.Key, "ms", time.Since(started).Milliseconds())
		return true, nil
	}

	var perm *ErrPermanent
	permanent := asPermanent(err, &perm)
	if permanent || int(job.Attempts) >= r.MaxAttempts {
		if e := r.DB.KillJob(ctx, job.ID, err.Error()); e != nil {
			r.Log.Error("job kill failed", "id", job.ID, "err", e)
		}
		r.Log.Error("job dead", "kind", job.Kind, "key", job.Key, "attempts", job.Attempts,
			"permanent", permanent, "err", err)
		return true, nil
	}

	delay := backoff(int(job.Attempts))
	if e := r.DB.RetryJob(ctx, job.ID, err.Error(), delay); e != nil {
		r.Log.Error("job retry failed", "id", job.ID, "err", e)
	}
	r.Log.Warn("job retrying", "kind", job.Kind, "key", job.Key,
		"attempts", job.Attempts, "in", delay.String(), "err", err)
	return true, nil
}

// runOne isolates a handler panic so one bad job cannot take the worker down.
func (r *Runner) runOne(ctx context.Context, h Handler, job store.Job) (err error) {
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("panic: %v", p)
		}
	}()
	return h(ctx, job)
}

func asPermanent(err error, target **ErrPermanent) bool {
	for e := err; e != nil; {
		if p, ok := e.(*ErrPermanent); ok {
			*target = p
			return true
		}
		u, ok := e.(interface{ Unwrap() error })
		if !ok {
			return false
		}
		e = u.Unwrap()
	}
	return false
}

// backoff grows five seconds per attempt, doubling, capped at ten minutes.
func backoff(attempt int) time.Duration {
	d := 5 * time.Second
	for i := 1; i < attempt && d < 10*time.Minute; i++ {
		d *= 2
	}
	if d > 10*time.Minute {
		d = 10 * time.Minute
	}
	return d
}

// Payload decodes a job's payload into v.
func Payload(job store.Job, v any) error {
	if len(job.Payload) == 0 {
		return nil
	}
	return json.Unmarshal(job.Payload, v)
}

// Heartbeat keeps a singleton job scheduled. The partial unique index means the
// enqueue is a no-op while one is already pending, so this is safe to call on a
// timer from every instance.
func Heartbeat(ctx context.Context, db *store.DB, kind string, every time.Duration, log *slog.Logger) {
	t := time.NewTicker(every)
	defer t.Stop()
	for {
		if err := db.Enqueue(ctx, kind, "singleton", nil, time.Now()); err != nil {
			log.Error("heartbeat enqueue failed", "kind", kind, "err", err)
		}
		select {
		case <-ctx.Done():
			return
		case <-t.C:
		}
	}
}
