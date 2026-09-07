package stub

import (
	"encoding/json"
	"net/http"
	"strconv"

	"marketplace/internal/ali"
	"marketplace/internal/stub/gen"
)

// The control plane. None of this is 1688: it is the set of levers a demo or an
// integration test needs — force an order forward, break an endpoint on purpose,
// stop the clock, start over. Everything here is INVENTED, which is why it all
// lives under /_control and is guarded by a token.

// control wraps a control handler with the shared token check.
func (s *Server) control(h func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if s.cfg.ControlToken != "" && r.Header.Get("X-Control-Token") != s.cfg.ControlToken {
			writeJSON(w, http.StatusForbidden, map[string]any{"error": "X-Control-Token required"})
			return
		}
		h(w, r)
	}
}

// decodeBody reads a JSON request body, tolerating an empty one.
func decodeBody(r *http.Request, v any) error {
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(v); err != nil && err.Error() != "EOF" {
		return err
	}
	return nil
}

// GET /_control/state — everything worth looking at in one place.
func (s *Server) controlState(w http.ResponseWriter, r *http.Request) {
	running, speed, autoPay := s.store.ClockState()

	orders := s.store.All()
	byStatus := map[string]int{}
	list := make([]any, 0, len(orders))
	for _, o := range orders {
		byStatus[o.Status]++
		entry := map[string]any{
			"orderId":    strconv.FormatInt(o.ID, 10),
			"outOrderId": o.OutOrderID,
			"status":     o.Status,
			"total":      o.Total.String(),
			"postFee":    o.PostFee.String(),
			"lines":      len(o.Lines),
			"seller":     o.SellerMemberID,
			"createTime": ali.FormatNaive(o.CreateTime),
			"modifyTime": ali.FormatNaive(o.ModifyTime),
		}
		if !o.NextAt.IsZero() {
			entry["nextAt"] = ali.FormatNaive(o.NextAt)
		}
		if o.Logistics != nil {
			entry["logisticsId"] = o.Logistics.LogisticsID
			entry["mailNo"] = o.Logistics.MailNo
			entry["cpCode"] = o.Logistics.CPCode
			entry["traceSteps"] = len(o.Logistics.Steps)
		}
		list = append(list, entry)
	}

	s.mu.Lock()
	faults := make([]any, 0, len(s.faults))
	for _, f := range s.faults {
		faults = append(faults, map[string]any{"api": f.API, "mode": f.Mode, "code": f.Code, "rate": f.Rate, "calls": f.count})
	}
	base, n := s.catalogueTop, s.catalogueN
	s.mu.Unlock()

	writeJSON(w, http.StatusOK, map[string]any{
		"now": ali.FormatNaive(s.clock.Now()),
		"clock": map[string]any{
			"running": running,
			"speedX":  speed,
			"autoPay": autoPay,
		},
		"catalogue": map[string]any{"base": base, "count": n},
		"orders":    list,
		"byStatus":  byStatus,
		"push": map[string]any{
			"pending": s.queue.Pending(),
			"sink":    sinkName(s.cfg.PushURL),
			"jitter":  s.cfg.PushJitter,
		},
		"switches": map[string]any{
			"chaos":      s.cfg.Chaos,
			"createFlat": s.cfg.CreateFlat,
			"sign":       s.cfg.Sign,
		},
		"faults": faults,
	})
}

func sinkName(pushURL string) string {
	if pushURL == "" {
		return "queue"
	}
	return "http:" + pushURL
}

// GET /_control/catalogue — the generated products, so a demo knows what to buy.
func (s *Server) controlCatalogue(w http.ResponseWriter, r *http.Request) {
	ids := s.catalogue()
	out := make([]any, 0, len(ids))
	for _, id := range ids {
		o := s.offer(id)
		skus := make([]any, 0, len(o.SKUs))
		for _, sk := range o.SKUs {
			skus = append(skus, map[string]any{
				"skuId":  sk.SkuID,
				"specId": sk.SpecID,
				"price":  sk.Price.String(),
				"stock":  sk.Stock,
			})
		}
		out = append(out, map[string]any{
			"offerId":          strconv.FormatInt(o.ID, 10),
			"subject":          o.Subject,
			"subjectTrans":     o.SubjectTrans,
			"archetype":        o.Archetype,
			"status":           o.Status,
			"quoteType":        o.QuoteType,
			"minOrderQuantity": o.MOQ,
			"batchNumber":      o.BatchNumber,
			"seller":           o.Seller.MemberID,
			"price":            o.MinPrice().String(),
			"tradeModes":       tradeModes(o.ID),
			"skus":             skus,
		})
	}
	writeJSON(w, http.StatusOK, map[string]any{"count": len(out), "offers": out})
}

// POST /_control/clock {"running":true,"speedX":60}
func (s *Server) controlClock(w http.ResponseWriter, r *http.Request) {
	body := struct {
		Running *bool   `json:"running"`
		SpeedX  float64 `json:"speedX"`
	}{}
	if err := decodeBody(r, &body); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
		return
	}
	running, speed, _ := s.store.ClockState()
	if body.Running != nil {
		running = *body.Running
	}
	if body.SpeedX > 0 {
		speed = body.SpeedX
	}
	s.store.SetClockRunning(running, speed)
	writeJSON(w, http.StatusOK, map[string]any{"running": running, "speedX": speed})
}

// POST /_control/orders/{id}/advance {"to":"waitbuyerreceive"}
func (s *Server) controlAdvance(w http.ResponseWriter, r *http.Request) {
	id, ok := parseInt64(r.PathValue("id"))
	if !ok {
		writeJSON(w, http.StatusBadRequest, map[string]any{"error": "bad order id"})
		return
	}
	body := struct {
		To string `json:"to"`
	}{}
	if err := decodeBody(r, &body); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
		return
	}
	if body.To == "" {
		body.To = StatusSuccess
	}

	events, err := s.store.AdvanceTo(id, body.To)
	if err != nil {
		writeJSON(w, http.StatusConflict, map[string]any{"error": err.Error()})
		return
	}
	s.Publish(r.Context(), events)

	o, _ := s.store.Get(id)
	writeJSON(w, http.StatusOK, map[string]any{
		"orderId":  strconv.FormatInt(id, 10),
		"status":   o.Status,
		"messages": len(events),
	})
}

// POST /_control/orders/{id}/trace {"statusChanged":"TRANSPORT","remark":"..."}
func (s *Server) controlTrace(w http.ResponseWriter, r *http.Request) {
	id, ok := parseInt64(r.PathValue("id"))
	if !ok {
		writeJSON(w, http.StatusBadRequest, map[string]any{"error": "bad order id"})
		return
	}
	body := struct {
		StatusChanged string `json:"statusChanged"`
		Remark        string `json:"remark"`
	}{}
	if err := decodeBody(r, &body); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
		return
	}
	if body.StatusChanged == "" {
		body.StatusChanged = "TRANSPORT"
	}
	if body.Remark == "" {
		body.Remark = "运输中"
	}

	events, err := s.store.AddTrace(id, body.StatusChanged, body.Remark)
	if err != nil {
		writeJSON(w, http.StatusConflict, map[string]any{"error": err.Error()})
		return
	}
	s.Publish(r.Context(), events)
	writeJSON(w, http.StatusOK, map[string]any{"orderId": strconv.FormatInt(id, 10), "added": body.StatusChanged})
}

// POST /_control/faults {"api":"com.alibaba.trade:alibaba.createOrder.preview:1",
//
//	"mode":"error","code":"500_004","rate":1}
//
// api "*" faults everything. rate 0 removes the fault.
func (s *Server) controlFaults(w http.ResponseWriter, r *http.Request) {
	body := fault{}
	if err := decodeBody(r, &body); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
		return
	}
	if body.API == "" {
		writeJSON(w, http.StatusBadRequest, map[string]any{"error": "api is required, or \"*\" for everything"})
		return
	}
	switch body.Mode {
	case "", "error", "timeout", "500":
	default:
		writeJSON(w, http.StatusBadRequest, map[string]any{"error": "mode must be error, timeout or 500"})
		return
	}

	s.mu.Lock()
	if body.Rate <= 0 {
		delete(s.faults, body.API)
	} else {
		if body.Rate > 1 {
			body.Rate = 1
		}
		f := body
		s.faults[body.API] = &f
	}
	n := len(s.faults)
	s.mu.Unlock()

	writeJSON(w, http.StatusOK, map[string]any{"api": body.API, "mode": body.Mode, "rate": body.Rate, "active": n})
}

// POST /_control/push {"topic":"ORDER_BUYER_VIEW_ORDER_PAY","data":{...}}
// Publishes one message verbatim, for testing a consumer without an order.
func (s *Server) controlPush(w http.ResponseWriter, r *http.Request) {
	body := struct {
		Topic string         `json:"topic"`
		User  string         `json:"userInfo"`
		Data  map[string]any `json:"data"`
	}{}
	if err := decodeBody(r, &body); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
		return
	}
	if body.Topic == "" {
		writeJSON(w, http.StatusBadRequest, map[string]any{"error": "topic is required"})
		return
	}
	if body.Data == nil {
		body.Data = map[string]any{}
	}
	if body.User == "" {
		body.User = buyer.MemberID
	}

	msg := s.envelope(body.Topic, body.User, body.Data, s.clock.Now())
	msgs := s.jitter([]map[string]any{msg})
	if err := s.sink.Deliver(r.Context(), msgs); err != nil {
		s.log.Debug("control push failed, queued", "err", err)
	}
	writeJSON(w, http.StatusOK, map[string]any{"published": len(msgs), "msgId": msg["msgId"], "type": msg["type"]})
}

// POST /_control/reset {"seed":3,"products":500}
// Clears every order and, optionally, moves the catalogue to a different band of
// offer ids so a fresh demo has fresh products.
func (s *Server) controlReset(w http.ResponseWriter, r *http.Request) {
	body := struct {
		Seed     int64 `json:"seed"`
		Products int   `json:"products"`
	}{}
	if err := decodeBody(r, &body); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
		return
	}

	s.store.Reset()
	_ = s.store.Save()

	s.mu.Lock()
	s.catalogueTop = gen.CatalogueBase + body.Seed*1000000
	if body.Products > 0 {
		s.catalogueN = body.Products
	}
	base, n := s.catalogueTop, s.catalogueN
	s.faults = map[string]*fault{}
	s.mu.Unlock()

	writeJSON(w, http.StatusOK, map[string]any{
		"orders":         0,
		"catalogueBase":  base,
		"catalogueCount": n,
	})
}
