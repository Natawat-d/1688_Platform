package stub

import (
	"bytes"
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"marketplace/internal/ali"
)

// The message TOPICS and the message ENVELOPE are documented; the transport is
// not. 1688 documents which fields a push message carries and how to fetch the
// ones that failed, but nothing we hold says how a message reaches a
// subscriber's endpoint. So delivery sits behind PushSink and the HTTP flavour
// is clearly marked INVENTED, while the replay APIs it falls back to are exactly
// as documented.

// Documented topics, spelled exactly as the docs spell them — including
// COMFIRM, which is 1688's typo, not ours.
const (
	TopicBuyerMake         = "ORDER_BUYER_VIEW_BUYER_MAKE"
	TopicOrderPay          = "ORDER_BUYER_VIEW_ORDER_PAY"
	TopicAnnounceSendGoods = "ORDER_BUYER_VIEW_ANNOUNCE_SENDGOODS"
	TopicPartSendGoods     = "ORDER_BUYER_VIEW_PART_PART_SENDGOODS"
	TopicConfirmReceive    = "ORDER_BUYER_VIEW_ORDER_COMFIRM_RECEIVEGOODS"
	TopicOrderSuccess      = "ORDER_BUYER_VIEW_ORDER_SUCCESS"
	TopicSellerClose       = "ORDER_BUYER_VIEW_ORDER_SELLER_CLOSE"
	TopicBuyerClose        = "ORDER_BUYER_VIEW_ORDER_BUYER_CLOSE"
	TopicLogisticsTrace    = "LOGISTICS_BUYER_VIEW_TRACE"
)

// topicGroup splits a topic id into its group and its name. The envelope's type
// is group + "_" + name, which for these topics reproduces the topic id.
func topicGroup(topic string) (group, name string) {
	if i := strings.Index(topic, "_"); i > 0 {
		return topic[:i], topic[i+1:]
	}
	return topic, topic
}

// PushSink is where lifecycle messages go. Two implementations: one posts them
// somewhere and queues what it could not deliver, the other only queues.
type PushSink interface {
	Deliver(ctx context.Context, msgs []map[string]any) error
}

// queued is one message waiting in the replay queue.
type queued struct {
	MsgID int64
	Type  string
	User  string
	Born  time.Time
	Msg   map[string]any
	Done  bool
}

// Queue backs the three documented replay APIs. Messages that were never
// delivered — or never had anywhere to go — wait here until a client fetches
// them with push.cursor.messageList or push.query.messageList.
type Queue struct {
	mu   sync.Mutex
	msgs []*queued
	max  int
}

// NewQueue builds a queue that keeps at most max messages.
func NewQueue(max int) *Queue {
	if max <= 0 {
		max = 5000
	}
	return &Queue{max: max}
}

// Add parks messages for replay.
func (q *Queue) Add(msgs []map[string]any) {
	q.mu.Lock()
	defer q.mu.Unlock()
	for _, m := range msgs {
		id, _ := m["msgId"].(int64)
		typ, _ := m["type"].(string)
		user, _ := m["userInfo"].(string)
		born, _ := m["gmtBorn"].(int64)
		q.msgs = append(q.msgs, &queued{
			MsgID: id, Type: typ, User: user,
			Born: time.UnixMilli(born).In(ali.CST),
			Msg:  m,
		})
	}
	if len(q.msgs) > q.max {
		q.msgs = q.msgs[len(q.msgs)-q.max:]
	}
}

type queueFilter struct {
	Type  string
	User  string
	From  time.Time
	To    time.Time
	Limit int
	Page  int
}

func (q *queued) matches(f queueFilter) bool {
	if f.Type != "" && q.Type != f.Type {
		return false
	}
	if f.User != "" && q.User != f.User {
		return false
	}
	if !f.From.IsZero() && q.Born.Before(f.From) {
		return false
	}
	if !f.To.IsZero() && q.Born.After(f.To) {
		return false
	}
	return true
}

// Cursor returns up to limit unconfirmed messages AND confirms them, which is
// what "retrieved messages are automatically confirmed as consumed" means: call
// it repeatedly with the same filter until it comes back empty.
func (q *Queue) Cursor(f queueFilter) []map[string]any {
	q.mu.Lock()
	defer q.mu.Unlock()
	out := make([]map[string]any, 0, f.Limit)
	for _, m := range q.msgs {
		if len(out) >= f.Limit {
			break
		}
		if m.Done || !m.matches(f) {
			continue
		}
		m.Done = true
		out = append(out, m.Msg)
	}
	return out
}

// Query returns one page of unconfirmed messages WITHOUT confirming them, plus
// the total. The caller confirms with push.message.confirm.
func (q *Queue) Query(f queueFilter) ([]map[string]any, int) {
	q.mu.Lock()
	defer q.mu.Unlock()
	var all []map[string]any
	for _, m := range q.msgs {
		if m.Done || !m.matches(f) {
			continue
		}
		all = append(all, m.Msg)
	}
	total := len(all)
	page, size := f.Page, f.Limit
	if page < 1 {
		page = 1
	}
	from := (page - 1) * size
	if from >= total {
		return nil, total
	}
	to := from + size
	if to > total {
		to = total
	}
	return all[from:to], total
}

// Confirm marks messages consumed.
func (q *Queue) Confirm(ids []int64) int {
	q.mu.Lock()
	defer q.mu.Unlock()
	want := map[int64]bool{}
	for _, id := range ids {
		want[id] = true
	}
	n := 0
	for _, m := range q.msgs {
		if want[m.MsgID] && !m.Done {
			m.Done = true
			n++
		}
	}
	return n
}

// Pending is how many messages are still unconfirmed, for /_control/state.
func (q *Queue) Pending() int {
	q.mu.Lock()
	defer q.mu.Unlock()
	n := 0
	for _, m := range q.msgs {
		if !m.Done {
			n++
		}
	}
	return n
}

// QueueSink is the sink used when PUSH_URL is empty: nothing is delivered
// anywhere, everything waits for the replay APIs.
type QueueSink struct{ Queue *Queue }

func (s QueueSink) Deliver(_ context.Context, msgs []map[string]any) error {
	s.Queue.Add(msgs)
	return nil
}

// HTTPSink posts messages to a subscriber endpoint.
//
// INVENTED — the whole of this transport. The documentation describes the
// message envelope and the replay APIs and says nothing whatever about how a
// message is delivered. What we made up, and what our own receiver therefore
// implements:
//
//	POST {PUSH_URL}
//	Content-Type: application/json
//	X-Aop-Signature: HMAC-SHA1 of the body under the app secret, uppercase hex
//	body: {"pushMessageList":[ <envelope>, ... ]}
//	expected: HTTP 200 with {"isSuccess":true}
//
// Three attempts with a linear backoff; anything still undelivered falls into
// the queue and is served by the documented replay APIs, which is the one part
// of this that IS documented behaviour.
type HTTPSink struct {
	URL    string
	Secret string
	HTTP   *http.Client
	Queue  *Queue
	Log    *slog.Logger
}

func (s *HTTPSink) Deliver(ctx context.Context, msgs []map[string]any) error {
	if len(msgs) == 0 {
		return nil
	}
	body, err := json.Marshal(map[string]any{"pushMessageList": msgs})
	if err != nil {
		s.Queue.Add(msgs)
		return err
	}

	// Sign the body as if it were a single request parameter, reusing the same
	// HMAC-SHA1 the gateway uses for requests.
	sig := ali.HMACSHA1Signer{Secret: s.Secret}.Sign("push", url.Values{"pushMessageList": {string(body)}})

	client := s.HTTP
	if client == nil {
		client = &http.Client{Timeout: 10 * time.Second}
	}

	var lastErr error
	for attempt := 0; attempt < 3; attempt++ {
		if attempt > 0 {
			select {
			case <-ctx.Done():
				s.Queue.Add(msgs)
				return ctx.Err()
			case <-time.After(time.Duration(attempt) * 300 * time.Millisecond):
			}
		}
		req, err := http.NewRequestWithContext(ctx, http.MethodPost, s.URL, bytes.NewReader(body))
		if err != nil {
			lastErr = err
			break
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("X-Aop-Signature", sig)

		resp, err := client.Do(req)
		if err != nil {
			lastErr = err
			continue
		}
		var ack struct {
			IsSuccess ali.FlexBool `json:"isSuccess"`
		}
		dec := json.NewDecoder(resp.Body)
		_ = dec.Decode(&ack)
		resp.Body.Close()

		if resp.StatusCode == http.StatusOK && bool(ack.IsSuccess) {
			return nil
		}
		lastErr = &ali.APIError{API: "push", Code: "PUSH_NOT_ACKED", Message: resp.Status, HTTP: resp.StatusCode}
	}

	if s.Log != nil {
		s.Log.Warn("push delivery failed, queued for replay", "count", len(msgs), "err", lastErr)
	}
	s.Queue.Add(msgs)
	return lastErr
}

// ------------------------------------------------------------------ envelopes

// envelope builds the documented push envelope. type is topicGroup + "_" +
// topicName, which is how the samples spell it.
func (s *Server) envelope(topic, user string, data map[string]any, born time.Time) map[string]any {
	group, name := topicGroup(topic)
	return map[string]any{
		"msgId":      s.nextMsgID(),
		"type":       group + "_" + name,
		"userInfo":   user,
		"data":       data,
		"gmtBorn":    ali.Millis(born),
		"topicGroup": group,
		"appKey":     s.cfg.AppKey,
		"topicName":  name,
	}
}

// messageFor turns one lifecycle event into one push message.
func (s *Server) messageFor(ev Event) map[string]any {
	if ev.Topic == TopicLogisticsTrace {
		// The logistics topic carries a nested model, not the flat order shape.
		data := map[string]any{
			"OrderLogisticsTracingModel": map[string]any{
				"logisticsId":   ev.LogisticsID,
				"cpCode":        ev.CPCode,
				"mailNo":        ev.MailNo,
				"statusChanged": ev.StatusChanged,
				"changeTime":    ali.FormatNaive(ev.At),
				"orderLogsItems": []any{
					map[string]any{"orderId": ev.OrderID, "orderEntryId": ev.SubOrderID},
				},
			},
		}
		return s.envelope(ev.Topic, ev.BuyerMemberID, data, ev.At)
	}

	// Every ORDER topic carries the same flat five fields. orderId is a Number.
	data := map[string]any{
		"orderId":        ev.OrderID,
		"currentStatus":  ev.CurrentStatus,
		"msgSendTime":    ali.FormatNaive(ev.At),
		"buyerMemberId":  ev.BuyerMemberID,
		"sellerMemberId": ev.SellerMemberID,
	}
	return s.envelope(ev.Topic, ev.BuyerMemberID, data, ev.At)
}

// Publish delivers a batch of events, applying the jitter switch first.
func (s *Server) Publish(ctx context.Context, events []Event) {
	if len(events) == 0 {
		return
	}
	msgs := make([]map[string]any, 0, len(events))
	for _, ev := range events {
		msgs = append(msgs, s.messageFor(ev))
	}
	msgs = s.jitter(msgs)
	if err := s.sink.Deliver(ctx, msgs); err != nil {
		s.log.Debug("push sink reported failure", "err", err)
	}
}

// jitter reproduces the two ways a real message bus misbehaves: out-of-order
// delivery and at-least-once duplication. Deterministic rather than random, so a
// consumer bug found today is findable again tomorrow: consecutive pairs are
// swapped within the batch, and every 7th message published in this run is sent
// twice.
//
// The duplication counter is per RUN, not per batch. A lifecycle publishes one
// or two messages at a time, so a per-batch counter would never reach seven and
// the switch would do nothing.
func (s *Server) jitter(msgs []map[string]any) []map[string]any {
	if !s.cfg.PushJitter || len(msgs) == 0 {
		return msgs
	}
	out := make([]map[string]any, len(msgs))
	copy(out, msgs)
	for i := 0; i+1 < len(out); i += 2 {
		out[i], out[i+1] = out[i+1], out[i]
	}
	dup := make([]map[string]any, 0, len(out)+1)
	for _, m := range out {
		dup = append(dup, m)
		if s.jitterSeq.Add(1)%7 == 0 {
			dup = append(dup, m)
		}
	}
	return dup
}

// ------------------------------------------------------------- replay handlers

// handlePushCursor serves cn.alibaba.open:push.cursor.messageList:1.
func (s *Server) handlePushCursor(form url.Values) any {
	f := queueFilter{
		Type:  form.Get("type"),
		User:  form.Get("userInfo"),
		From:  parseDocTime(form.Get("createStartTime")),
		To:    parseDocTime(form.Get("createEndTime")),
		Limit: clampQuantity(form.Get("quantity")),
	}
	list := s.queue.Cursor(f)
	if list == nil {
		list = []map[string]any{}
	}
	return map[string]any{"pushMessageList": list}
}

// handlePushQuery serves cn.alibaba.open:push.query.messageList:1.
func (s *Server) handlePushQuery(form url.Values) any {
	f := queueFilter{
		Type:  form.Get("type"),
		User:  form.Get("userInfo"),
		From:  parseDocTime(form.Get("createStartTime")),
		To:    parseDocTime(form.Get("createEndTime")),
		Limit: clampQuantity(form.Get("pageSize")),
		Page:  atoiDefault(form.Get("page"), 1),
	}
	list, total := s.queue.Query(f)
	if list == nil {
		list = []map[string]any{}
	}
	return map[string]any{
		"pushMessagePage": map[string]any{
			"datas":      list,
			"totalCount": total,
		},
	}
}

// handlePushConfirm serves cn.alibaba.open:push.message.confirm:1.
func (s *Server) handlePushConfirm(form url.Values) any {
	var ids []int64
	raw := strings.TrimSpace(form.Get("msgIdList"))
	if raw != "" {
		var vals []ali.FlexInt64
		if err := json.Unmarshal([]byte(raw), &vals); err != nil {
			return map[string]any{"isSuccess": false}
		}
		for _, v := range vals {
			ids = append(ids, int64(v))
		}
	}
	s.queue.Confirm(ids)
	return map[string]any{"isSuccess": true}
}

// clampQuantity applies the documented 20-200 range with a default of 20.
func clampQuantity(s string) int {
	n := atoiDefault(s, 20)
	if n < 20 {
		n = 20
	}
	if n > 200 {
		n = 200
	}
	return n
}
