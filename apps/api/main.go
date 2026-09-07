// Command api serves the storefront, the admin console and the background
// workers that relay orders onto 1688 and track them coming back.
package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"sort"
	"strings"
	"syscall"
	"time"

	"marketplace/internal/ali"
	"marketplace/internal/api"
	"marketplace/internal/catalog"
	"marketplace/internal/jobs"
	"marketplace/internal/order"
	"marketplace/internal/store"
)

type config struct {
	Env         string
	Addr        string
	DatabaseURL string
	AdminToken  string

	GatewayURL  string
	AppKey      string
	AppSecret   string
	AccessToken string
	StubBase    string
}

// devPlaceholders are the values shipped in .env.example and docker-compose.yml
// so the project runs with no setup. Every one of them is public knowledge.
var devPlaceholders = map[string]bool{
	"dev": true, "devsecret": true, "devtoken": true, "DEVKEY": true,
	"changeme": true, "password": true, "postgres": true,
}

// checkProductionConfig refuses to start a production process that is still
// carrying development credentials.
//
// The admin token guards every settings, pricing and catalogue route, and it
// defaults to the literal string "dev". A deploy that forgets one environment
// variable would otherwise come up quietly and be wide open, which is exactly
// the failure that never gets noticed until it matters.
func checkProductionConfig(cfg config) error {
	if !strings.EqualFold(cfg.Env, "production") {
		return nil
	}
	var bad []string
	for name, value := range map[string]string{
		"ADMIN_TOKEN":      cfg.AdminToken,
		"ALI_APP_SECRET":   cfg.AppSecret,
		"ALI_ACCESS_TOKEN": cfg.AccessToken,
	} {
		if value == "" || devPlaceholders[value] {
			bad = append(bad, name)
		}
	}
	if len(bad) > 0 {
		sort.Strings(bad)
		return fmt.Errorf("ENV=production but these still hold a development value: %s",
			strings.Join(bad, ", "))
	}
	if strings.Contains(cfg.DatabaseURL, "://app:app@") {
		return errors.New("ENV=production but DATABASE_URL still uses the development password")
	}
	return nil
}

func load() config {
	return config{
		Env:         env("ENV", "dev"),
		Addr:        env("API_ADDR", ":8787"),
		DatabaseURL: env("DATABASE_URL", "postgres://app:app@localhost:5455/market?sslmode=disable"),
		AdminToken:  env("ADMIN_TOKEN", "dev"),
		GatewayURL:  env("ALI_BASE_URL", "http://localhost:8788"),
		AppKey:      env("ALI_APP_KEY", "DEVKEY"),
		AppSecret:   env("ALI_APP_SECRET", "devsecret"),
		AccessToken: env("ALI_ACCESS_TOKEN", "devtoken"),
		StubBase:    env("STUB_BASE", "http://localhost:8788"),
	}
}

func main() {
	smoke := flag.Bool("smoke", false, "make one signed gateway call and exit")
	flag.Parse()

	log := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: level()}))
	slog.SetDefault(log)
	cfg := load()

	if *smoke {
		if err := runSmoke(cfg, log); err != nil {
			log.Error("smoke test failed", "err", err)
			os.Exit(1)
		}
		return
	}

	if err := run(cfg, log); err != nil {
		log.Error("fatal", "err", err)
		os.Exit(1)
	}
}

func run(cfg config, log *slog.Logger) error {
	if err := checkProductionConfig(cfg); err != nil {
		return err
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	db, err := store.OpenWithRetry(ctx, cfg.DatabaseURL, 60*time.Second)
	if err != nil {
		return fmt.Errorf("database: %w (is Postgres running? try `make db-up`)", err)
	}
	defer db.Close()

	if err := db.Migrate(ctx); err != nil {
		return fmt.Errorf("migrate: %w", err)
	}
	log.Info("database ready")

	cli := ali.New(cfg.GatewayURL, cfg.AppKey, cfg.AppSecret, cfg.AccessToken)
	// Every gateway call lands in the admin integration log. The write is best
	// effort: losing an audit row must never fail the business operation.
	cli.Log = func(c ali.CallLog) {
		req, _ := json.Marshal(c.Params)
		code, msg, ok := "", "", c.Err == nil && c.Status < 400
		if c.Err != nil {
			msg = c.Err.Error()
		}
		if e := db.LogAPICall(context.WithoutCancel(ctx), store.APICall{
			API: c.API, MS: int32(c.Duration.Milliseconds()), OK: ok,
			Code: code, Message: msg, Req: req, Resp: c.Body,
		}); e != nil {
			log.Debug("api call log write failed", "err", e)
		}
	}

	importer := catalog.New(db, cli, log)
	orders := order.NewService(db, cli, log)
	if d, err := time.ParseDuration(env("POLL_INTERVAL", "5m")); err == nil && d > 0 {
		orders.PollInterval = d
	}

	runner := jobs.New(db, log)
	orders.Register(runner, importer)
	go runner.Run(ctx)

	// Singleton jobs keep themselves scheduled once started; the heartbeats make
	// sure they exist at all, including on a database that has just been created.
	go jobs.Heartbeat(ctx, db, store.JobPushGapfill, 5*time.Minute, log)
	go jobs.Heartbeat(ctx, db, store.JobMessageSweep, 2*time.Minute, log)
	go jobs.Heartbeat(ctx, db, store.JobOrderPollSweep, orders.PollInterval, log)
	go jobs.Heartbeat(ctx, db, store.JobTrimLogs, 24*time.Hour, log)

	srv := &api.Server{
		DB: db, Cli: cli, Orders: orders, Importer: importer, Log: log,
		AdminToken: cfg.AdminToken, StubBase: cfg.StubBase,
		// The gateway signs its pushes with the same secret it signs requests
		// with, so this is the one we verify against.
		PushSecret: cfg.AppSecret,
	}
	httpSrv := &http.Server{
		Addr:              cfg.Addr,
		Handler:           srv.Routes(),
		ReadHeaderTimeout: 10 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	errCh := make(chan error, 1)
	go func() {
		log.Info("listening", "addr", cfg.Addr, "gateway", cfg.GatewayURL)
		if err := httpSrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- err
		}
	}()

	select {
	case err := <-errCh:
		return err
	case <-ctx.Done():
		log.Info("shutting down")
	}

	shutdownCtx, cancel := context.WithTimeout(context.WithoutCancel(ctx), 15*time.Second)
	defer cancel()
	return httpSrv.Shutdown(shutdownCtx)
}

// runSmoke proves the signing and transport work end to end against whatever
// ALI_BASE_URL points at, without needing a database.
func runSmoke(cfg config, log *slog.Logger) error {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	cli := ali.New(cfg.GatewayURL, cfg.AppKey, cfg.AppSecret, cfg.AccessToken)
	cli.Log = func(c ali.CallLog) {
		log.Info("gateway call", "api", c.API, "status", c.Status, "ms", c.Duration.Milliseconds())
	}

	acct, err := cli.AccountBasic(ctx)
	if err != nil {
		return fmt.Errorf("account.basic: %w", err)
	}
	fmt.Printf("account: loginId=%s memberId=%s company=%s\n", acct.LoginID, acct.MemberID, acct.CompanyName)

	page, err := cli.SearchOffers(ctx, ali.OfferQuery{Keyword: "", BeginPage: 1, PageSize: 5, Country: "en"})
	if err != nil {
		return fmt.Errorf("keywordQuery: %w", err)
	}
	fmt.Printf("search: %d total, showing %d\n", page.TotalRecords, len(page.Data))
	for _, it := range page.Data {
		fmt.Printf("  %s  %-52s  %s\n", it.OfferID.String(), truncate(it.SubjectTrans, 52), it.PriceInfo.Price)
	}
	return nil
}

func truncate(s string, n int) string {
	r := []rune(s)
	if len(r) <= n {
		return s
	}
	return string(r[:n-1]) + "…"
}

func env(key, def string) string {
	if v := strings.TrimSpace(os.Getenv(key)); v != "" {
		return v
	}
	return def
}

func level() slog.Level {
	switch strings.ToLower(os.Getenv("LOG_LEVEL")) {
	case "debug":
		return slog.LevelDebug
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
