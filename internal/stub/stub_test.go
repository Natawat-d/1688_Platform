package stub

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"marketplace/internal/ali"
	"marketplace/internal/stub/gen"
)

// The doc-driven contract tests live elsewhere. What is checked here is the part
// of the stub that has no documentation to check it against: the order state
// machine, and the idempotency rule that stops a retried create from buying
// everything twice.

func testDelays() Delays {
	return Delays{
		Pay:  2 * time.Minute,
		Ship: 10 * time.Minute,
		Step: 6 * time.Minute,
		Sign: 20 * time.Minute,
	}
}

func testParams() CreateParams {
	o := gen.Offer(gen.CatalogueBase)
	sku := o.SKUs[0]
	return CreateParams{
		Flow:      "general",
		TradeType: "assureTrade",
		Seller:    o.Seller,
		Buyer:     buyer,
		PostFee:   ali.MustParseFen("8.00"),
		Address: Address{
			FullName:     "张三",
			Mobile:       "15251667788",
			ProvinceText: "浙江省",
			CityText:     "杭州市",
			AreaText:     "滨江区",
			Address:      "网商路699号",
			DistrictCode: "330108",
		},
		Lines: []Line{{
			OfferID:   o.ID,
			SkuID:     sku.SkuID,
			SpecID:    sku.SpecID,
			Quantity:  3,
			UnitPrice: sku.Price,
			Amount:    sku.Price * 3,
			Name:      o.Subject,
			Unit:      o.UnitZH,
		}},
	}
}

// TestLifecycle walks one order from creation to success on an injected clock,
// checking the status, the timing and the message published at each step. No
// sleeping: the whole lifecycle happens in microseconds.
func TestLifecycle(t *testing.T) {
	clock := NewTestClock(time.Date(2026, 3, 1, 9, 0, 0, 0, ali.CST))
	store := NewStore(StoreConfig{Delays: testDelays(), Speed: 1, AutoPay: true}, clock)

	orders, events, replayed, err := store.Create(testParams())
	if err != nil {
		t.Fatalf("create: %v", err)
	}
	if replayed {
		t.Fatal("first create reported as a replay")
	}
	if len(orders) != 1 {
		t.Fatalf("want one order, got %d", len(orders))
	}
	o := orders[0]
	if o.Status != StatusWaitPay {
		t.Fatalf("new order status = %q, want %q", o.Status, StatusWaitPay)
	}
	if len(events) != 1 || events[0].Topic != TopicBuyerMake {
		t.Fatalf("create should publish %s, got %+v", TopicBuyerMake, events)
	}
	if events[0].CurrentStatus != StatusWaitPay {
		t.Fatalf("%s currentStatus = %q, want %q", TopicBuyerMake, events[0].CurrentStatus, StatusWaitPay)
	}

	// Nothing happens before the payment delay is up.
	clock.Advance(time.Minute)
	if ev := store.Tick(); len(ev) != 0 {
		t.Fatalf("order moved early: %+v", ev)
	}
	if o.Status != StatusWaitPay {
		t.Fatalf("status = %q after one minute, want %q", o.Status, StatusWaitPay)
	}

	// Payment.
	clock.Advance(2 * time.Minute)
	assertTopic(t, store.Tick(), TopicOrderPay)
	if o.Status != StatusWaitSend {
		t.Fatalf("after payment status = %q, want %q", o.Status, StatusWaitSend)
	}
	if o.PayTime.IsZero() {
		t.Fatal("payTime not recorded")
	}

	// Shipment: a waybill appears and the first tracking node with it.
	clock.Advance(10 * time.Minute)
	shipEvents := store.Tick()
	assertTopic(t, shipEvents, TopicAnnounceSendGoods)
	assertTopic(t, shipEvents, TopicLogisticsTrace)
	if o.Status != StatusWaitReceive {
		t.Fatalf("after shipment status = %q, want %q", o.Status, StatusWaitReceive)
	}
	if o.Logistics == nil {
		t.Fatal("shipment created no logistics record")
	}
	if o.Logistics.MailNo == "" || o.Logistics.CPCode == "" {
		t.Fatalf("incomplete waybill: %+v", o.Logistics)
	}
	if len(o.Logistics.Steps) != 1 || o.Logistics.Steps[0].Status != "CONSIGN" {
		t.Fatalf("first trace node = %+v, want one CONSIGN", o.Logistics.Steps)
	}

	// Three more nodes, then the signature, which confirms receipt.
	for i := 0; i < 3; i++ {
		clock.Advance(6 * time.Minute)
		store.Tick()
	}
	if got, want := len(o.Logistics.Steps), 4; got != want {
		t.Fatalf("after three steps there are %d trace nodes, want %d", got, want)
	}
	if o.Status != StatusWaitReceive {
		t.Fatalf("status = %q mid-delivery, want %q", o.Status, StatusWaitReceive)
	}

	clock.Advance(20 * time.Minute)
	assertTopic(t, store.Tick(), TopicConfirmReceive)
	if o.Status != StatusConfirmGoods {
		t.Fatalf("after signature status = %q, want %q", o.Status, StatusConfirmGoods)
	}
	if last := o.Logistics.Steps[len(o.Logistics.Steps)-1]; last.Status != "SIGN" {
		t.Fatalf("last trace node = %q, want SIGN", last.Status)
	}

	clock.Advance(6 * time.Minute)
	assertTopic(t, store.Tick(), TopicOrderSuccess)
	if o.Status != StatusSuccess {
		t.Fatalf("final status = %q, want %q", o.Status, StatusSuccess)
	}
	if !o.NextAt.IsZero() {
		t.Fatal("a finished order should have no next transition")
	}

	// And it stays finished.
	clock.Advance(24 * time.Hour)
	if ev := store.Tick(); len(ev) != 0 {
		t.Fatalf("finished order moved again: %+v", ev)
	}
}

// TestConfirmReceiveStatusIsDocumented pins the one topic whose currentStatus is
// not the order status: the docs say confirm_goods_and_has_subsidy.
func TestConfirmReceiveStatusIsDocumented(t *testing.T) {
	clock := NewTestClock(time.Date(2026, 3, 1, 9, 0, 0, 0, ali.CST))
	store := NewStore(StoreConfig{Delays: testDelays(), Speed: 1, AutoPay: false}, clock)

	orders, _, _, _ := store.Create(testParams())
	events, err := store.AdvanceTo(orders[0].ID, StatusConfirmGoods)
	if err != nil {
		t.Fatalf("advance: %v", err)
	}
	for _, ev := range events {
		if ev.Topic != TopicConfirmReceive {
			continue
		}
		if ev.CurrentStatus != "confirm_goods_and_has_subsidy" {
			t.Fatalf("currentStatus = %q, want confirm_goods_and_has_subsidy", ev.CurrentStatus)
		}
		return
	}
	t.Fatalf("no %s message among %d events", TopicConfirmReceive, len(events))
}

// TestNoAutoPay checks that STUB_AUTOPAY=0 parks an order at waitbuyerpay until
// somebody pays it, which is what makes the cashier page worth having.
func TestNoAutoPay(t *testing.T) {
	clock := NewTestClock(time.Date(2026, 3, 1, 9, 0, 0, 0, ali.CST))
	store := NewStore(StoreConfig{Delays: testDelays(), Speed: 1, AutoPay: false}, clock)

	orders, _, _, _ := store.Create(testParams())
	o := orders[0]

	clock.Advance(48 * time.Hour)
	if ev := store.Tick(); len(ev) != 0 {
		t.Fatalf("order moved without autopay: %+v", ev)
	}
	if o.Status != StatusWaitPay {
		t.Fatalf("status = %q, want %q", o.Status, StatusWaitPay)
	}

	events, err := store.Pay(o.ID)
	if err != nil {
		t.Fatalf("pay: %v", err)
	}
	assertTopic(t, events, TopicOrderPay)
	if o.Status != StatusWaitSend {
		t.Fatalf("after manual payment status = %q, want %q", o.Status, StatusWaitSend)
	}
	if _, err := store.Pay(o.ID); err != errStatus {
		t.Fatalf("paying twice returned %v, want errStatus", err)
	}
}

// TestSpeedScalesDelays checks that STUB_SPEED shortens the waits rather than
// changing the sequence.
func TestSpeedScalesDelays(t *testing.T) {
	clock := NewTestClock(time.Date(2026, 3, 1, 9, 0, 0, 0, ali.CST))
	store := NewStore(StoreConfig{Delays: testDelays(), Speed: 60, AutoPay: true}, clock)

	orders, _, _, _ := store.Create(testParams())
	o := orders[0]

	// Two minutes at 60x is two seconds.
	clock.Advance(time.Second)
	store.Tick()
	if o.Status != StatusWaitPay {
		t.Fatalf("status = %q after one second at 60x, want %q", o.Status, StatusWaitPay)
	}
	clock.Advance(2 * time.Second)
	store.Tick()
	if o.Status != StatusWaitSend {
		t.Fatalf("status = %q after three seconds at 60x, want %q", o.Status, StatusWaitSend)
	}
}

// TestOutOrderIDIsIdempotent is the rule that matters when a network retry hits
// an order that already exists: same external id, same orders, no second charge.
func TestOutOrderIDIsIdempotent(t *testing.T) {
	clock := NewTestClock(time.Date(2026, 3, 1, 9, 0, 0, 0, ali.CST))
	store := NewStore(StoreConfig{Delays: testDelays(), Speed: 1, AutoPay: false}, clock)

	p := testParams()
	p.OutOrderID = "SHOP-20260301-0001"

	first, firstEvents, replayed, err := store.Create(p)
	if err != nil {
		t.Fatalf("create: %v", err)
	}
	if replayed {
		t.Fatal("first create reported as a replay")
	}
	if len(firstEvents) != 1 {
		t.Fatalf("first create published %d messages, want 1", len(firstEvents))
	}

	// The retry arrives a minute later. Nothing about it should be new.
	clock.Advance(time.Minute)
	second, secondEvents, replayed, err := store.Create(p)
	if err != nil {
		t.Fatalf("repeat create: %v", err)
	}
	if !replayed {
		t.Fatal("repeat create was not reported as a replay")
	}
	if len(secondEvents) != 0 {
		t.Fatalf("repeat create published %d messages, want 0", len(secondEvents))
	}
	if len(second) != len(first) {
		t.Fatalf("repeat create returned %d orders, want %d", len(second), len(first))
	}
	for i := range first {
		if first[i].ID != second[i].ID {
			t.Fatalf("order %d: repeat create returned id %d, want %d", i, second[i].ID, first[i].ID)
		}
	}
	if got := len(store.All()); got != len(first) {
		t.Fatalf("store holds %d orders after the retry, want %d", got, len(first))
	}

	// A different external id is a different order.
	p.OutOrderID = "SHOP-20260301-0002"
	other, _, replayed, err := store.Create(p)
	if err != nil {
		t.Fatalf("create with a new outOrderId: %v", err)
	}
	if replayed {
		t.Fatal("a new outOrderId was treated as a replay")
	}
	if other[0].ID == first[0].ID {
		t.Fatal("a new outOrderId reused an existing order id")
	}
}

// TestCancelRules covers the three documented refusals of alibaba.trade.cancel.
func TestCancelRules(t *testing.T) {
	clock := NewTestClock(time.Date(2026, 3, 1, 9, 0, 0, 0, ali.CST))
	store := NewStore(StoreConfig{Delays: testDelays(), Speed: 1, AutoPay: false}, clock)

	if _, err := store.Cancel(1234567890123, "buyerCancel", ""); err != errNotExist {
		t.Fatalf("cancelling a missing order returned %v, want errNotExist", err)
	}

	orders, _, _, _ := store.Create(testParams())
	o := orders[0]

	// Under ten seconds old.
	clock.Advance(9 * time.Second)
	if _, err := store.Cancel(o.ID, "buyerCancel", ""); err != errTooFast {
		t.Fatalf("cancelling a nine-second-old order returned %v, want errTooFast", err)
	}

	clock.Advance(2 * time.Second)
	events, err := store.Cancel(o.ID, "buyerCancel", "改主意了")
	if err != nil {
		t.Fatalf("cancel: %v", err)
	}
	assertTopic(t, events, TopicBuyerClose)
	if o.Status != StatusCancel {
		t.Fatalf("status = %q after cancel, want %q", o.Status, StatusCancel)
	}
	if o.CloseOperateType != "CLOSE_TRADE_BY_BUYER" {
		t.Fatalf("closeOperateType = %q", o.CloseOperateType)
	}

	// Already cancelled: a status error, not a second cancellation.
	if _, err := store.Cancel(o.ID, "buyerCancel", ""); err != errStatus {
		t.Fatalf("cancelling twice returned %v, want errStatus", err)
	}

	// Paid orders cannot be cancelled either.
	orders2, _, _, _ := store.Create(testParams())
	paid := orders2[0]
	clock.Advance(30 * time.Second)
	if _, err := store.Pay(paid.ID); err != nil {
		t.Fatalf("pay: %v", err)
	}
	if _, err := store.Cancel(paid.ID, "buyerCancel", ""); err != errStatus {
		t.Fatalf("cancelling a paid order returned %v, want errStatus", err)
	}
}

// TestSnapshotRoundTrip checks that a restart mid-demo keeps the orders and the
// idempotency index.
func TestSnapshotRoundTrip(t *testing.T) {
	path := t.TempDir() + "/state.json"
	clock := NewTestClock(time.Date(2026, 3, 1, 9, 0, 0, 0, ali.CST))

	store := NewStore(StoreConfig{Delays: testDelays(), Speed: 1, AutoPay: true, StateFile: path}, clock)
	p := testParams()
	p.OutOrderID = "SHOP-RESTART-1"
	orders, _, _, _ := store.Create(p)
	clock.Advance(3 * time.Minute)
	store.Tick()
	if err := store.Save(); err != nil {
		t.Fatalf("save: %v", err)
	}

	reloaded := NewStore(StoreConfig{Delays: testDelays(), Speed: 1, AutoPay: true, StateFile: path}, clock)
	o, ok := reloaded.Get(orders[0].ID)
	if !ok {
		t.Fatal("order did not survive the restart")
	}
	if o.Status != StatusWaitSend {
		t.Fatalf("restored status = %q, want %q", o.Status, StatusWaitSend)
	}
	if _, ok := reloaded.GetByOut("SHOP-RESTART-1"); !ok {
		t.Fatal("the outOrderId index did not survive the restart")
	}
	// And the reloaded store keeps handing out fresh ids.
	more, _, _, _ := reloaded.Create(testParams())
	if more[0].ID == orders[0].ID {
		t.Fatal("a reloaded store reused an order id")
	}
}

// TestSignedCallEndToEnd drives one real request through the whole gateway:
// signature, system-parameter validation, dispatch, envelope.
func TestSignedCallEndToEnd(t *testing.T) {
	srv := testServer(t)
	ts := httptest.NewServer(srv.Handler())
	defer ts.Close()

	body := call(t, ts.URL, "1", "com.alibaba.account", "alibaba.account.basic", url.Values{})
	var env struct {
		Result struct {
			LoginID  string `json:"loginId"`
			MemberID string `json:"memberId"`
		} `json:"result"`
		Success bool `json:"success"`
	}
	if err := json.Unmarshal(body, &env); err != nil {
		t.Fatalf("decode: %v — %s", err, body)
	}
	if !env.Success || env.Result.MemberID != buyer.MemberID {
		t.Fatalf("unexpected account.basic response: %s", body)
	}
}

// TestSignatureIsChecked makes sure the gateway is not waving requests through.
func TestSignatureIsChecked(t *testing.T) {
	srv := testServer(t)
	ts := httptest.NewServer(srv.Handler())
	defer ts.Close()

	form := url.Values{}
	form.Set(ali.TokenParam, "devtoken")
	form.Set(ali.SignatureParam, "DEADBEEF")
	resp, err := http.PostForm(ts.URL+"/openapi/param2/1/com.alibaba.account/alibaba.account.basic/DEVKEY", form)
	if err != nil {
		t.Fatalf("post: %v", err)
	}
	defer resp.Body.Close()

	var out map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if out["error_code"] != "400" || out["success"] != false {
		t.Fatalf("a bad signature was accepted: %v", out)
	}
}

// TestUnknownAPI pins the invented not-found body.
func TestUnknownAPI(t *testing.T) {
	srv := testServer(t)
	ts := httptest.NewServer(srv.Handler())
	defer ts.Close()

	resp, err := http.PostForm(ts.URL+"/openapi/param2/1/com.alibaba.nope/does.not.exist/DEVKEY", url.Values{})
	if err != nil {
		t.Fatalf("post: %v", err)
	}
	defer resp.Body.Close()

	var out map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if out["error_code"] != "400" || out["error_message"] != "api not found" {
		t.Fatalf("unexpected not-found body: %v", out)
	}
}

// TestRequiredParameterIsEnforced proves the validation really is driven by
// ALL-APIS.json: cancel documents tradeID as required.
func TestRequiredParameterIsEnforced(t *testing.T) {
	srv := testServer(t)
	ts := httptest.NewServer(srv.Handler())
	defer ts.Close()

	form := url.Values{}
	form.Set("webSite", "1688")
	form.Set("cancelReason", "other")
	body := call(t, ts.URL, "1", "com.alibaba.trade", "alibaba.trade.cancel", form)

	var out map[string]any
	if err := json.Unmarshal(body, &out); err != nil {
		t.Fatalf("decode: %v", err)
	}
	msg, _ := out["error_message"].(string)
	if !strings.Contains(msg, "tradeID") {
		t.Fatalf("missing tradeID was not rejected: %v", out)
	}
}

// TestPreviewRefusesUnorderableOffer walks one documented business refusal all
// the way through the gateway, and checks the documented typo survives.
func TestPreviewRefusesUnorderableOffer(t *testing.T) {
	srv := testServer(t)
	ts := httptest.NewServer(srv.Handler())
	defer ts.Close()

	// Every 37th offer is not published.
	var target int64
	for _, id := range srv.catalogue() {
		if !gen.Offer(id).Orderable() {
			target = id
			break
		}
	}
	if target == 0 {
		t.Fatal("the generated catalogue contains no unorderable offer")
	}
	o := gen.Offer(target)

	form := url.Values{}
	form.Set("addressParam", `{"fullName":"张三","mobile":"15251667788","provinceText":"浙江省","cityText":"杭州市","areaText":"滨江区","address":"网商路699号"}`)
	form.Set("cargoParamList", `[{"offerId":`+itoa(target)+`,"specId":"`+o.SKUs[0].SpecID+`","quantity":`+itoa(int64(o.MOQ))+`}]`)
	body := call(t, ts.URL, "1", "com.alibaba.trade", "alibaba.createOrder.preview", form)

	var out map[string]any
	if err := json.Unmarshal(body, &out); err != nil {
		t.Fatalf("decode: %v — %s", err, body)
	}
	if out["errorCode"] != ali.ErrNoOnlineTrade {
		t.Fatalf("errorCode = %v, want %s — %s", out["errorCode"], ali.ErrNoOnlineTrade, body)
	}
	if _, ok := out["errorMsg"]; !ok {
		t.Fatalf("preview refusal used errorMessage, not the documented errorMsg: %s", body)
	}
	if _, ok := out["orderPreviewResuslt"]; !ok {
		t.Fatalf("the documented misspelling orderPreviewResuslt is missing: %s", body)
	}
}

// TestPushQueueReplay checks the documented cursor semantics: read once, and the
// next read with the same filter comes back empty.
func TestPushQueueReplay(t *testing.T) {
	srv := testServer(t)
	clock := NewTestClock(time.Date(2026, 3, 1, 9, 0, 0, 0, ali.CST))

	var msgs []map[string]any
	for i := 0; i < 3; i++ {
		msgs = append(msgs, srv.envelope(TopicOrderPay, buyer.MemberID, map[string]any{"orderId": int64(i)}, clock.Now()))
	}
	srv.queue.Add(msgs)

	first := srv.handlePushCursor(url.Values{}).(map[string]any)["pushMessageList"].([]map[string]any)
	if len(first) != 3 {
		t.Fatalf("cursor returned %d messages, want 3", len(first))
	}
	second := srv.handlePushCursor(url.Values{}).(map[string]any)["pushMessageList"].([]map[string]any)
	if len(second) != 0 {
		t.Fatalf("cursor returned %d messages on the second read, want 0", len(second))
	}
	if got := srv.queue.Pending(); got != 0 {
		t.Fatalf("%d messages still pending after a cursor read, want 0", got)
	}
}

// TestPushEnvelope pins the documented envelope, including type == topicGroup +
// "_" + topicName.
func TestPushEnvelope(t *testing.T) {
	srv := testServer(t)
	ev := Event{
		Topic:          TopicOrderPay,
		OrderID:        105581756010628737,
		CurrentStatus:  StatusWaitSend,
		BuyerMemberID:  buyer.MemberID,
		SellerMemberID: "b2b-2248564064",
		At:             time.Date(2026, 3, 1, 19, 34, 27, 0, ali.CST),
	}
	msg := srv.messageFor(ev)

	if msg["type"] != "ORDER_BUYER_VIEW_ORDER_PAY" {
		t.Fatalf("type = %v", msg["type"])
	}
	if msg["topicGroup"] != "ORDER" || msg["topicName"] != "BUYER_VIEW_ORDER_PAY" {
		t.Fatalf("group/name = %v/%v", msg["topicGroup"], msg["topicName"])
	}
	data := msg["data"].(map[string]any)
	if data["msgSendTime"] != "2026-03-01 19:34:27" {
		t.Fatalf("msgSendTime = %v, want the naive Shanghai layout", data["msgSendTime"])
	}
	if data["orderId"] != int64(105581756010628737) {
		t.Fatalf("orderId = %v, want a bare number", data["orderId"])
	}
}

// ------------------------------------------------------------------- helpers

func testServer(t *testing.T) *Server {
	t.Helper()
	srv, err := New(Config{
		Base:        "http://stub.test",
		AppKey:      "DEVKEY",
		AccessToken: "devtoken",
		AppSecret:   "devsecret",
		Products:    120,
		DocsDir:     "../../1688-api-docs",
		Speed:       1,
		Sign:        "on",
		StateFile:   t.TempDir() + "/state.json",
		Delays:      testDelays(),
	}, slog.New(slog.NewTextHandler(discard{}, &slog.HandlerOptions{Level: slog.LevelError})),
		NewTestClock(time.Date(2026, 3, 1, 9, 0, 0, 0, ali.CST)))
	if err != nil {
		t.Fatalf("new server: %v", err)
	}
	return srv
}

// call signs and posts one request the way internal/ali does.
func call(t *testing.T, base, version, namespace, name string, form url.Values) []byte {
	t.Helper()
	form.Set(ali.TokenParam, "devtoken")
	signPath := "param2/" + version + "/" + namespace + "/" + name + "/DEVKEY"
	form.Set(ali.SignatureParam, ali.HMACSHA1Signer{Secret: "devsecret"}.Sign(signPath, form))

	resp, err := http.PostForm(base+"/openapi/param2/"+version+"/"+namespace+"/"+name+"/DEVKEY", form)
	if err != nil {
		t.Fatalf("post: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("status = %d, want 200", resp.StatusCode)
	}
	buf := make([]byte, 0, 4096)
	tmp := make([]byte, 4096)
	for {
		n, err := resp.Body.Read(tmp)
		buf = append(buf, tmp[:n]...)
		if err != nil {
			break
		}
	}
	return buf
}

func itoa(v int64) string {
	return strings.TrimSpace(jsonNumber(v))
}

func jsonNumber(v int64) string {
	b, _ := json.Marshal(v)
	return string(b)
}

// discard is an io.Writer that throws logs away during tests.
type discard struct{}

func (discard) Write(p []byte) (int, error) { return len(p), nil }

func assertTopic(t *testing.T, events []Event, topic string) {
	t.Helper()
	for _, ev := range events {
		if ev.Topic == topic {
			return
		}
	}
	got := make([]string, 0, len(events))
	for _, ev := range events {
		got = append(got, ev.Topic)
	}
	t.Fatalf("no %s among %v", topic, got)
}
