// Package store owns every SQL statement in this program. Nothing outside it
// builds a query, which is what keeps the database swappable and keeps the rest
// of the code readable.
package store

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// DB is a connection pool plus the queries defined across this package.
type DB struct {
	Pool *pgxpool.Pool
}

// Open connects, verifies the connection, and returns a pool.
func Open(ctx context.Context, dsn string) (*DB, error) {
	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("store: bad DATABASE_URL: %w", err)
	}
	cfg.MaxConns = 8
	cfg.MinConns = 1
	cfg.MaxConnIdleTime = 5 * time.Minute
	cfg.MaxConnLifetime = time.Hour

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("store: connect: %w", err)
	}
	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("store: ping: %w", err)
	}
	return &DB{Pool: pool}, nil
}

// OpenWithRetry waits for the database to accept connections, which matters when
// the api container starts alongside Postgres.
func OpenWithRetry(ctx context.Context, dsn string, wait time.Duration) (*DB, error) {
	deadline := time.Now().Add(wait)
	var lastErr error
	for {
		db, err := Open(ctx, dsn)
		if err == nil {
			return db, nil
		}
		lastErr = err
		if time.Now().After(deadline) {
			return nil, lastErr
		}
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(time.Second):
		}
	}
}

// Close releases the pool.
func (db *DB) Close() {
	if db != nil && db.Pool != nil {
		db.Pool.Close()
	}
}

// Ping reports whether the database is reachable.
func (db *DB) Ping(ctx context.Context) error { return db.Pool.Ping(ctx) }
