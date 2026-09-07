package stub

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"marketplace/internal/ali"
	"marketplace/internal/stub/gen"
)

// Config is the whole of the stub's behaviour. Everything here comes from an
// environment variable in apps/stubgw; nothing is read from a config file.
type Config struct {
	Addr         string
	Base         string // absolute base url this gateway is reachable at
	AppKey       string
	AccessToken  string
	AppSecret    string
	Products     int
	StateFile    string
	DocsDir      string
	PushURL      string
	Speed        float64
	AutoPay      bool
	Delays       Delays
	Chaos        bool
	PushJitter   bool
	CreateFlat   bool
	Sign         string // "on" (default) or "off"
	ControlToken string
}

// buyer is the single authorised user this gateway pretends to have
// authenticated. Everything that needs "who is calling" uses it.
var buyer = Buyer{
	MemberID: "b2b-2248544159",
	UserID:   2248544159,
	LoginID:  "alitestforisv01",
	Company:  "AOP对外测试账号01",
}

// fault is one injected failure, from /_control/faults.
type fault struct {
	API   string  `json:"api"`
	Mode  string  `json:"mode"` // error | timeout | 500
	Code  string  `json:"code"`
	Rate  float64 `json:"rate"`
	count int64
}

// Server is the whole gateway: dispatch, validation, catalogue, orders, push.
type Server struct {
	cfg   Config
	log   *slog.Logger
	specs map[string]APISpec

	store *Store
	queue *Queue
	sink  PushSink
	clock Clock

	mu           sync.Mutex
	faults       map[string]*fault
	catalogueTop int64
	catalogueN   int

	msgSeq    atomic.Int64
	jitterSeq atomic.Int64
	mux       *http.ServeMux
}

// handler is one endpoint. It receives the already-validated form and returns
// whatever should be marshalled as the response body.
type handler func(ctx context.Context, form url.Values) any

// New wires a gateway. It reads ALL-APIS.json once, here, and never again.
func New(cfg Config, log *slog.Logger, clock Clock) (*Server, error) {
	if log == nil {
		log = slog.Default()
	}
	if clock == nil {
		clock = SystemClock()
	}
	specs, err := LoadSpecs(cfg.DocsDir)
	if err != nil {
		return nil, err
	}

	q := NewQueue(5000)
	var sink PushSink = QueueSink{Queue: q}
	if cfg.PushURL != "" {
		sink = &HTTPSink{URL: cfg.PushURL, Secret: cfg.AppSecret, Queue: q, Log: log}
	}

	s := &Server{
		cfg:   cfg,
		log:   log,
		specs: specs,
		queue: q,
		sink:  sink,
		clock: clock,
		store: NewStore(StoreConfig{
			Delays:    cfg.Delays,
			Speed:     cfg.Speed,
			AutoPay:   cfg.AutoPay,
			StateFile: cfg.StateFile,
		}, clock),
		faults:       map[string]*fault{},
		catalogueTop: gen.CatalogueBase,
		catalogueN:   cfg.Products,
	}
	s.msgSeq.Store(clock.Now().UnixMilli())
	s.routes()
	return s, nil
}

// Handler is the http.Handler to serve.
func (s *Server) Handler() http.Handler { return s.mux }

// Store exposes the order store, for tests.
func (s *Server) Store() *Store { return s.store }

// Queue exposes the replay queue, for tests.
func (s *Server) Queue() *Queue { return s.queue }

func (s *Server) routes() {
	mux := http.NewServeMux()

	// One handler for the entire documented surface. Dispatch is on
	// namespace/name/version, exactly as the request path spells them.
	mux.HandleFunc("POST /openapi/param2/{version}/{namespace}/{name}/{appKey}", s.serveAPI)

	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Write([]byte("ok\n"))
	})
	mux.HandleFunc("GET /img/{offerId}/{file}", s.serveImage)
	mux.HandleFunc("GET /cashier", s.serveCashier)
	mux.HandleFunc("POST /cashier/pay", s.serveCashierPay)

	mux.HandleFunc("GET /_control/state", s.control(s.controlState))
	mux.HandleFunc("GET /_control/catalogue", s.control(s.controlCatalogue))
	mux.HandleFunc("POST /_control/clock", s.control(s.controlClock))
	mux.HandleFunc("POST /_control/orders/{id}/advance", s.control(s.controlAdvance))
	mux.HandleFunc("POST /_control/orders/{id}/trace", s.control(s.controlTrace))
	mux.HandleFunc("POST /_control/faults", s.control(s.controlFaults))
	mux.HandleFunc("POST /_control/push", s.control(s.controlPush))
	mux.HandleFunc("POST /_control/reset", s.control(s.controlReset))

	s.mux = mux
}

// Run drives the lifecycle: tick, publish, snapshot. It returns when ctx is done.
func (s *Server) Run(ctx context.Context) {
	go s.store.Flush(ctx, 500*time.Millisecond)

	t := time.NewTicker(250 * time.Millisecond)
	defer t.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-t.C:
			if events := s.store.Tick(); len(events) > 0 {
				s.Publish(ctx, events)
			}
		}
	}
}

// serveAPI is the single gateway entry point.
func (s *Server) serveAPI(w http.ResponseWriter, r *http.Request) {
	version := r.PathValue("version")
	namespace := r.PathValue("namespace")
	name := r.PathValue("name")
	appKey := r.PathValue("appKey")
	key := namespace + ":" + name + ":" + version

	if err := r.ParseForm(); err != nil {
		writeJSON(w, http.StatusOK, (&gatewayError{Code: "400", Message: "malformed request body"}).body())
		return
	}
	form := r.PostForm

	spec, ok := s.specs[key]
	if !ok {
		// INVENTED: the gateway's not-found body is not documented. See the
		// comment on gatewayError.
		s.log.Warn("unknown api", "api", key)
		writeJSON(w, http.StatusOK, map[string]any{
			"error_code":    "400",
			"error_message": "api not found",
			"success":       false,
		})
		return
	}

	h, ok := s.dispatch(key)
	if !ok {
		// Documented by 1688 but not implemented here: still a real endpoint,
		// so say so rather than pretending it does not exist.
		writeJSON(w, http.StatusOK, (&gatewayError{Code: "400", Message: "api not enabled on this gateway: " + key}).body())
		return
	}

	if gErr := s.validate(spec, appKey, form); gErr != nil {
		s.log.Info("rejected", "api", key, "code", gErr.Code, "message", gErr.Message)
		writeJSON(w, http.StatusOK, gErr.body())
		return
	}

	if resp, injected := s.injectFault(r.Context(), key, w); injected {
		if resp != nil {
			writeJSON(w, http.StatusOK, resp)
		}
		return
	}

	s.log.Debug("call", "api", key)
	writeJSON(w, http.StatusOK, h(r.Context(), form))
}

// dispatch maps an api key onto its handler.
func (s *Server) dispatch(key string) (handler, bool) {
	switch key {
	case "com.alibaba.account:alibaba.account.basic:1":
		return s.wrap(s.handleAccountBasic), true

	case "com.alibaba.fenxiao.crossborder:product.search.keywordQuery:1":
		return s.wrap(s.handleKeywordQuery), true
	case "com.alibaba.fenxiao.crossborder:product.search.keywordSNQuery:1":
		return s.wrap(s.handleKeywordSNQuery), true
	case "com.alibaba.fenxiao.crossborder:product.search.queryProductDetail:1":
		return s.wrap(s.handleProductDetail), true
	case "com.alibaba.fenxiao.crossborder:category.translation.getById:1":
		return s.wrap(s.handleCategoryByID), true
	case "com.alibaba.fenxiao.crossborder:product.freight.estimate:1":
		return s.wrap(s.handleFreightEstimate), true

	case "com.alibaba.trade:alibaba.createOrder.preview:1":
		return s.wrap(s.handlePreview), true
	case "com.alibaba.trade:alibaba.trade.createCrossOrder:1":
		return s.handleCreateCrossOrder, true
	case "com.alibaba.trade:alibaba.alipay.url.get:1":
		return s.wrap(s.handleAlipayURL), true
	case "com.alibaba.trade:alibaba.trade.get.buyerView:1":
		return s.wrap(s.handleOrderBuyerView), true
	case "com.alibaba.trade:alibaba.trade.getBuyerOrderList:1":
		return s.wrap(s.handleBuyerOrderList), true
	case "com.alibaba.trade:alibaba.trade.cancel:1":
		return s.handleTradeCancel, true

	case "com.alibaba.logistics:alibaba.trade.getLogisticsTraceInfo.buyerView:1":
		return s.wrap(s.handleLogisticsTrace), true

	case "cn.alibaba.open:push.cursor.messageList:1":
		return s.wrap(s.handlePushCursor), true
	case "cn.alibaba.open:push.query.messageList:1":
		return s.wrap(s.handlePushQuery), true
	case "cn.alibaba.open:push.message.confirm:1":
		return s.wrap(s.handlePushConfirm), true
	}
	return nil, false
}

// wrap adapts the handlers that need neither the context nor to publish.
func (s *Server) wrap(f func(url.Values) any) handler {
	return func(_ context.Context, form url.Values) any { return f(form) }
}

// injectFault applies /_control/faults. It returns (body, true) when the call
// should not reach its handler; a body of nil means the response was already
// written.
func (s *Server) injectFault(ctx context.Context, key string, w http.ResponseWriter) (any, bool) {
	s.mu.Lock()
	f := s.faults[key]
	if f == nil {
		f = s.faults["*"]
	}
	var mode, code string
	var hit bool
	if f != nil {
		// Deterministic rather than random: the nth call in every hundred is
		// faulted, so a failing test fails again.
		n := f.count
		f.count++
		hit = float64(n%100) < f.Rate*100
		mode, code = f.Mode, f.Code
	}
	s.mu.Unlock()
	if !hit {
		return nil, false
	}

	switch mode {
	case "timeout":
		select {
		case <-ctx.Done():
		case <-time.After(60 * time.Second):
		}
		return nil, true
	case "500":
		http.Error(w, "gateway error", http.StatusInternalServerError)
		return nil, true
	default: // "error"
		if code == "" {
			code = "500_001"
		}
		return (&gatewayError{Code: code, Message: "injected fault"}).body(), true
	}
}

// nextMsgID hands out message ids that look like the documented ones and never
// repeat within a run.
func (s *Server) nextMsgID() int64 { return s.msgSeq.Add(1) }

// catalogue is the current list of offer ids.
func (s *Server) catalogue() []int64 {
	s.mu.Lock()
	base, n := s.catalogueTop, s.catalogueN
	s.mu.Unlock()
	return gen.CatalogueFrom(base, n)
}

// offer looks up one product. Any id works: the generator is total.
func (s *Server) offer(id int64) gen.Product { return gen.Offer(id) }

// imageURL turns a generator-relative image path into an absolute one.
func (s *Server) imageURL(path string) string { return strings.TrimRight(s.cfg.Base, "/") + path }

func (s *Server) imageURLs(paths []string) []string {
	out := make([]string, 0, len(paths))
	for _, p := range paths {
		out = append(out, s.imageURL(p))
	}
	return out
}

// serveImage renders the deterministic SVG behind every image url in the
// catalogue, so a demo needs no network at all.
func (s *Server) serveImage(w http.ResponseWriter, r *http.Request) {
	offerID, err := strconv.ParseInt(r.PathValue("offerId"), 10, 64)
	if err != nil {
		http.Error(w, "bad offer id", http.StatusBadRequest)
		return
	}
	file := strings.TrimSuffix(r.PathValue("file"), ".svg")
	n, err := strconv.Atoi(file)
	if err != nil || n < 0 || n > 64 {
		http.Error(w, "bad image index", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "image/svg+xml; charset=utf-8")
	w.Header().Set("Cache-Control", "public, max-age=86400")
	w.Write([]byte(gen.SVG(offerID, n)))
}

// ------------------------------------------------------------------- plumbing

func writeJSON(w http.ResponseWriter, status int, v any) {
	b, err := json.Marshal(v)
	if err != nil {
		http.Error(w, "encode error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	w.Write(b)
}

func atoiDefault(s string, def int) int {
	if s == "" {
		return def
	}
	n, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		return def
	}
	return n
}

func parseInt64(s string) (int64, bool) {
	v, err := strconv.ParseInt(strings.Trim(strings.TrimSpace(s), `"`), 10, 64)
	return v, err == nil
}

// parseDocTime reads any of the timestamp spellings the docs use, and returns
// the zero time for anything unparseable rather than failing a whole call.
func parseDocTime(s string) time.Time {
	t, err := ali.ParseTime(s)
	if err != nil {
		return time.Time{}
	}
	return t
}

// jsonObject parses one object-shaped request parameter.
func jsonObject(form url.Values, field string) (map[string]any, bool) {
	raw := strings.TrimSpace(form.Get(field))
	if raw == "" {
		return nil, false
	}
	var m map[string]any
	if err := json.Unmarshal([]byte(raw), &m); err != nil {
		return nil, false
	}
	return m, true
}

// jsonArray parses one array-shaped request parameter.
func jsonArray(form url.Values, field string) ([]map[string]any, bool) {
	raw := strings.TrimSpace(form.Get(field))
	if raw == "" {
		return nil, false
	}
	var a []map[string]any
	if err := json.Unmarshal([]byte(raw), &a); err != nil {
		return nil, false
	}
	return a, true
}

// mapStr reads a string out of a decoded JSON object, tolerating numbers.
func mapStr(m map[string]any, key string) string {
	v, ok := m[key]
	if !ok || v == nil {
		return ""
	}
	switch t := v.(type) {
	case string:
		return t
	case json.Number:
		return t.String()
	case float64:
		return strconv.FormatFloat(t, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(t)
	}
	return ""
}

// mapInt64 reads an integer out of a decoded JSON object. JSON numbers arrive as
// float64, which loses precision above 2^53 — so a value that came in as a
// string is parsed as text, and only genuine numbers go through float64.
func mapInt64(m map[string]any, key string) int64 {
	v, ok := m[key]
	if !ok || v == nil {
		return 0
	}
	switch t := v.(type) {
	case string:
		n, _ := parseInt64(t)
		return n
	case json.Number:
		n, _ := t.Int64()
		return n
	case float64:
		return int64(t)
	}
	return 0
}

// chaosList implements half of STUB_CHAOS: on odd offer ids an empty collection
// comes back as null instead of []. Clients that assume "always an array" break,
// which is the point.
func (s *Server) chaosList(offerID int64, list []any) any {
	if list == nil {
		list = []any{}
	}
	if s.cfg.Chaos && offerID%2 == 1 && len(list) == 0 {
		return nil
	}
	return list
}

// chaosExtra implements the other half: undocumented extra fields on odd offer
// ids. A strict decoder that rejects unknown fields will trip over these.
func (s *Server) chaosExtra(offerID int64, m map[string]any) map[string]any {
	if !s.cfg.Chaos || offerID%2 == 0 {
		return m
	}
	m["_stubChaos"] = true
	m["experimentBucket"] = "B"
	m["serverRegion"] = "cn-hangzhou"
	return m
}
