// Command stubgw is a stand-in for the 1688 open platform gateway.
//
// It speaks the documented protocol — the same paths, the same form encoding,
// the same signature, the same error envelopes, the same field names down to the
// typos — over a deterministic catalogue and a real order lifecycle. Our backend
// talks to it exactly as it would talk to gw.open.1688.com, so nothing in the
// client is written twice.
//
// Everything is configured by environment variable:
//
//	STUB_ADDR           listen address                    (default :8788)
//	STUB_BASE           absolute url this gateway is at   (default http://localhost:8788)
//	STUB_APP_KEY        app key echoed in push envelopes  (default DEVKEY)
//	STUB_ACCESS_TOKEN   the only token accepted           (default devtoken)
//	STUB_APP_SECRET     signing secret                    (default devsecret)
//	STUB_PRODUCTS       catalogue size                    (default 240)
//	STUB_STATE_FILE     order snapshot path               (default ./.stub-state.json)
//	DOCS_DIR            1688-api-docs checkout            (default ./1688-api-docs)
//	PUSH_URL            where to POST push messages       (default empty: queue only)
//	STUB_SPEED          lifecycle speed multiplier        (default 60)
//	STUB_AUTOPAY        1 pays orders automatically       (default 1)
//	STUB_DELAY_PAY      wait before autopay               (default 2m)
//	STUB_DELAY_SHIP     wait before shipment              (default 10m)
//	STUB_DELAY_STEP     wait between tracking nodes       (default 6m)
//	STUB_DELAY_SIGN     wait before signature             (default 20m)
//	STUB_CHAOS          1 emits nulls and extra fields    (default 0)
//	STUB_PUSH_JITTER    1 reorders and duplicates pushes  (default 0)
//	STUB_CREATE_FLAT    1 returns the unwrapped create    (default 0)
//	STUB_SIGN           "off" skips signature checking    (default on)
//	STUB_CONTROL_TOKEN  guards /_control/*                (default empty: open)
package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync/atomic"
	"syscall"
	"time"

	"marketplace/internal/stub"
)

// listenFailed records a bind failure so the process can exit non-zero.
var listenFailed atomic.Bool

func main() {
	log := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: logLevel()}))
	slog.SetDefault(log)

	cfg := stub.Config{
		Addr:        env("STUB_ADDR", ":8788"),
		Base:        strings.TrimRight(env("STUB_BASE", "http://localhost:8788"), "/"),
		AppKey:      env("STUB_APP_KEY", "DEVKEY"),
		AccessToken: env("STUB_ACCESS_TOKEN", "devtoken"),
		AppSecret:   env("STUB_APP_SECRET", "devsecret"),
		Products:    envInt("STUB_PRODUCTS", 240),
		StateFile:   env("STUB_STATE_FILE", "./.stub-state.json"),
		DocsDir:     env("DOCS_DIR", "./1688-api-docs"),
		PushURL:     env("PUSH_URL", ""),
		Speed:       envFloat("STUB_SPEED", 60),
		AutoPay:     envBool("STUB_AUTOPAY", true),
		Delays: stub.Delays{
			Pay:  envDuration("STUB_DELAY_PAY", 2*time.Minute),
			Ship: envDuration("STUB_DELAY_SHIP", 10*time.Minute),
			Step: envDuration("STUB_DELAY_STEP", 6*time.Minute),
			Sign: envDuration("STUB_DELAY_SIGN", 20*time.Minute),
		},
		Chaos:        envBool("STUB_CHAOS", false),
		PushJitter:   envBool("STUB_PUSH_JITTER", false),
		CreateFlat:   envBool("STUB_CREATE_FLAT", false),
		Sign:         env("STUB_SIGN", "on"),
		ControlToken: env("STUB_CONTROL_TOKEN", ""),
	}

	srv, err := stub.New(cfg, log, stub.SystemClock())
	if err != nil {
		log.Error("startup failed", "err", err)
		os.Exit(1)
	}

	ctx, stopSignals := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stopSignals()

	go srv.Run(ctx)

	httpSrv := &http.Server{
		Addr:              cfg.Addr,
		Handler:           srv.Handler(),
		ReadHeaderTimeout: 10 * time.Second,
	}

	go func() {
		log.Info("stub gateway listening",
			"addr", cfg.Addr,
			"base", cfg.Base,
			"products", cfg.Products,
			"docs", cfg.DocsDir,
			"sign", cfg.Sign,
			"autopay", cfg.AutoPay,
			"speed", cfg.Speed,
			"push", pushTarget(cfg.PushURL),
			"chaos", cfg.Chaos,
			"jitter", cfg.PushJitter,
			"createFlat", cfg.CreateFlat,
		)
		if err := httpSrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error("listen failed", "err", err)
			listenFailed.Store(true)
			stopSignals()
		}
	}()

	<-ctx.Done()
	log.Info("shutting down")
	// A bind failure is the one way this process ends without being asked to.
	// It must not look like a clean exit to whatever supervises it.
	defer func() {
		if listenFailed.Load() {
			os.Exit(1)
		}
	}()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := httpSrv.Shutdown(shutdownCtx); err != nil {
		log.Error("shutdown", "err", err)
	}
	// The lifecycle goroutine saves on its way out, but a demo that was killed
	// mid-transition deserves the last word written too.
	if err := srv.Store().Save(); err != nil {
		log.Error("saving state", "err", err)
	}
	log.Info("stopped")
}

func pushTarget(u string) string {
	if u == "" {
		return "queue-only"
	}
	return u
}

func env(key, def string) string {
	if v, ok := os.LookupEnv(key); ok && v != "" {
		return v
	}
	return def
}

func envInt(key string, def int) int {
	v, ok := os.LookupEnv(key)
	if !ok || v == "" {
		return def
	}
	n, err := strconv.Atoi(strings.TrimSpace(v))
	if err != nil {
		return def
	}
	return n
}

func envFloat(key string, def float64) float64 {
	v, ok := os.LookupEnv(key)
	if !ok || v == "" {
		return def
	}
	f, err := strconv.ParseFloat(strings.TrimSpace(v), 64)
	if err != nil || f <= 0 {
		return def
	}
	return f
}

// envBool reads 1/true/yes/on as true and 0/false/no/off as false.
func envBool(key string, def bool) bool {
	v, ok := os.LookupEnv(key)
	if !ok || v == "" {
		return def
	}
	switch strings.ToLower(strings.TrimSpace(v)) {
	case "1", "true", "yes", "on", "y":
		return true
	case "0", "false", "no", "off", "n":
		return false
	}
	return def
}

// envDuration accepts a Go duration ("90s", "10m") or a bare number of seconds.
func envDuration(key string, def time.Duration) time.Duration {
	v, ok := os.LookupEnv(key)
	if !ok || v == "" {
		return def
	}
	v = strings.TrimSpace(v)
	if d, err := time.ParseDuration(v); err == nil {
		return d
	}
	if n, err := strconv.Atoi(v); err == nil {
		return time.Duration(n) * time.Second
	}
	return def
}

func logLevel() slog.Level {
	switch strings.ToLower(os.Getenv("STUB_LOG_LEVEL")) {
	case "debug":
		return slog.LevelDebug
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	}
	return slog.LevelInfo
}
