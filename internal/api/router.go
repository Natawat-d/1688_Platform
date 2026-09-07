// Package api serves our own JSON API and the built storefront.
//
// Two conventions run through every response and are enforced by dto_test.go:
// identifiers are strings, and money is an object of minor units plus display
// text. Neither the browser nor a future mobile client should ever have to know
// what a fen is.
package api

import (
	"context"
	"crypto/rand"
	"crypto/subtle"
	"encoding/hex"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
	"time"

	"marketplace/internal/ali"
	"marketplace/internal/catalog"
	"marketplace/internal/order"
	"marketplace/internal/store"
)

// Server holds everything the handlers need.
type Server struct {
	DB         *store.DB
	Cli        *ali.Client
	Orders     *order.Service
	Importer   *catalog.Importer
	Log        *slog.Logger
	AdminToken string
	StubBase   string // when set, /api/admin/stub/* proxies here; empty at go-live

	// PushSecret verifies the signature on inbound gateway messages. Empty
	// accepts unsigned deliveries, which is the right default only until the
	// real gateway's signing scheme is known.
	PushSecret string
}

const cartCookie = "cart"

// Routes builds the mux.
func (s *Server) Routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/health", s.health)

	// Storefront.
	mux.HandleFunc("GET /api/categories", s.categories)
	mux.HandleFunc("GET /api/products", s.searchProducts)
	mux.HandleFunc("GET /api/products/{offerId}", s.getProduct)
	mux.HandleFunc("GET /api/cart", s.getCart)
	mux.HandleFunc("POST /api/cart/items", s.addCartItem)
	mux.HandleFunc("PATCH /api/cart/items/{id}", s.updateCartItem)
	mux.HandleFunc("DELETE /api/cart/items/{id}", s.deleteCartItem)
	mux.HandleFunc("POST /api/checkout", s.checkout)
	mux.HandleFunc("POST /api/orders/{publicId}/pay", s.payOrder)
	mux.HandleFunc("GET /api/orders/{publicId}", s.getOrder)
	mux.HandleFunc("POST /api/orders/{publicId}/cancel", s.cancelOrder)

	// The gateway pushes here. No bearer token: it is authenticated by the
	// signature header and made safe by being idempotent on message id.
	mux.HandleFunc("POST /api/hooks/1688", s.receivePush)

	// Admin.
	mux.Handle("GET /api/admin/settings", s.admin(s.getSettings))
	mux.Handle("PUT /api/admin/settings", s.admin(s.putSettings))
	mux.Handle("GET /api/admin/fee-rules", s.admin(s.listFeeRules))
	mux.Handle("POST /api/admin/fee-rules", s.admin(s.createFeeRule))
	mux.Handle("PUT /api/admin/fee-rules/{id}", s.admin(s.updateFeeRule))
	mux.Handle("DELETE /api/admin/fee-rules/{id}", s.admin(s.deleteFeeRule))
	mux.Handle("POST /api/admin/fee-rules/preview", s.admin(s.previewFee))
	mux.Handle("POST /api/admin/import", s.admin(s.startImport))
	mux.Handle("GET /api/admin/products", s.admin(s.adminProducts))
	mux.Handle("PATCH /api/admin/products/{offerId}", s.admin(s.setProductVisible))
	mux.Handle("GET /api/admin/orders", s.admin(s.adminOrders))
	mux.Handle("GET /api/admin/orders/{id}", s.admin(s.adminOrder))
	mux.Handle("POST /api/admin/supplier-orders/{id}/relay", s.admin(s.retryRelay))
	mux.Handle("POST /api/admin/supplier-orders/{id}/cancel", s.admin(s.adminCancelSupplier))
	mux.Handle("GET /api/admin/jobs", s.admin(s.listJobs))
	mux.Handle("POST /api/admin/jobs/{id}/retry", s.admin(s.retryJob))
	mux.Handle("GET /api/admin/api-calls", s.admin(s.listAPICalls))
	mux.Handle("GET /api/admin/messages", s.admin(s.listMessages))
	mux.Handle("/api/admin/stub/", s.admin(s.proxyStub))

	// The built single-page app, when it was embedded at build time.
	mux.Handle("/", s.spa())

	return logging(s.Log, mux)
}

func (s *Server) health(w http.ResponseWriter, r *http.Request) {
	out := map[string]any{"ok": true, "time": time.Now().UTC()}
	if err := s.DB.Ping(r.Context()); err != nil {
		out["ok"] = false
		out["db"] = err.Error()
		writeJSON(w, http.StatusServiceUnavailable, out)
		return
	}
	writeJSON(w, http.StatusOK, out)
}

// ------------------------------------------------------------- middleware --

func logging(log *slog.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		started := time.Now()
		rec := &statusRecorder{ResponseWriter: w, status: http.StatusOK}
		next.ServeHTTP(rec, r)
		if strings.HasPrefix(r.URL.Path, "/api/") {
			log.Info("http", "method", r.Method, "path", r.URL.Path,
				"status", rec.status, "ms", time.Since(started).Milliseconds())
		}
	})
}

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (r *statusRecorder) WriteHeader(code int) {
	r.status = code
	r.ResponseWriter.WriteHeader(code)
}

// admin guards a handler with the bearer token, compared in constant time.
func (s *Server) admin(h http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
		if s.AdminToken == "" || subtle.ConstantTimeCompare([]byte(token), []byte(s.AdminToken)) != 1 {
			writeErr(w, http.StatusUnauthorized, "unauthorized", "Admin token required.")
			return
		}
		h(w, r)
	})
}

// ---------------------------------------------------------------- session --

// cartID reads the shopper's cart cookie, minting one when they have none.
// Anonymous carts are the whole identity model here; there are no accounts.
func (s *Server) cartID(w http.ResponseWriter, r *http.Request) (string, error) {
	if c, err := r.Cookie(cartCookie); err == nil && len(c.Value) == 32 && isHex(c.Value) {
		return c.Value, s.DB.EnsureCart(r.Context(), c.Value)
	}
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	value := hex.EncodeToString(b)
	c := &http.Cookie{
		Name:     cartCookie,
		Value:    value,
		Path:     "/",
		HttpOnly: true,
		// Secure whenever the visitor arrived over HTTPS. Behind CloudFront or a
		// load balancer the connection to this process is plain HTTP, so the
		// forwarded protocol is the only honest signal; setting Secure
		// unconditionally would break local development over HTTP.
		Secure:   isHTTPS(r),
		SameSite: http.SameSiteLaxMode,
		MaxAge:   60 * 60 * 24 * 90,
	}
	http.SetCookie(w, c)
	// Attach the new cookie to the request as well, so a second lookup within
	// the same request finds it instead of minting another cart. Without this,
	// an add-to-cart that renders the cart afterwards wrote the line into one
	// cart and handed the browser the cookie of a different, empty one.
	r.AddCookie(c)
	return value, s.DB.EnsureCart(r.Context(), value)
}

func isHex(s string) bool {
	_, err := hex.DecodeString(s)
	return err == nil
}

// isHTTPS reports whether the visitor's own connection was encrypted, including
// when TLS was terminated by CloudFront or a load balancer in front of us.
func isHTTPS(r *http.Request) bool {
	if r.TLS != nil {
		return true
	}
	// X-Forwarded-Proto may be a list when several proxies are chained; the
	// first entry is the one the visitor used.
	proto, _, _ := strings.Cut(r.Header.Get("X-Forwarded-Proto"), ",")
	return strings.EqualFold(strings.TrimSpace(proto), "https")
}

// ------------------------------------------------------------------ replies --

func writeJSON(w http.ResponseWriter, code int, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	if err := enc.Encode(v); err != nil {
		// The status line has already gone out, so there is nothing to do but
		// leave a trace for the log.
		_, _ = w.Write([]byte("\n"))
	}
}

// ErrorResponse is the shape every failure takes.
type ErrorResponse struct {
	Error   string     `json:"error"`
	Message string     `json:"message"`
	Issues  []IssueDTO `json:"issues,omitempty"`
}

func writeErr(w http.ResponseWriter, code int, kind, msg string) {
	writeJSON(w, code, ErrorResponse{Error: kind, Message: msg})
}

// decode reads a JSON body, rejecting unknown fields so a typo in a request is
// reported rather than silently ignored.
func decode(r *http.Request, v any) error {
	dec := json.NewDecoder(http.MaxBytesReader(nil, r.Body, 1<<20))
	dec.DisallowUnknownFields()
	return dec.Decode(v)
}

func pathInt(r *http.Request, name string) (int64, error) {
	return strconv.ParseInt(r.PathValue(name), 10, 64)
}

func queryInt(r *http.Request, name string, def int) int {
	v, err := strconv.Atoi(r.URL.Query().Get(name))
	if err != nil {
		return def
	}
	return v
}

func queryInt64(r *http.Request, name string, def int64) int64 {
	v, err := strconv.ParseInt(r.URL.Query().Get(name), 10, 64)
	if err != nil {
		return def
	}
	return v
}

// notFound answers a missing row consistently.
func (s *Server) notFound(w http.ResponseWriter, err error, what string) bool {
	if errors.Is(err, store.ErrNotFound) {
		writeErr(w, http.StatusNotFound, "not_found", what+" was not found.")
		return true
	}
	return false
}

func (s *Server) fail(w http.ResponseWriter, r *http.Request, err error, msg string) {
	s.Log.Error(msg, "path", r.URL.Path, "err", err)
	writeErr(w, http.StatusInternalServerError, "internal", msg)
}

// ctxTimeout bounds a handler's work.
func ctxTimeout(r *http.Request, d time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(r.Context(), d)
}
