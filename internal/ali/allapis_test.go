package ali_test

// Registry and request-side contract tests against 1688-api-docs/ALL-APIS.json:
// the endpoint flags and URLs in api.go, and the required parameters every
// request builder must emit.

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"testing"
	"time"

	"marketplace/internal/ali"
)

type allAPIEntry struct {
	API           string `json:"api"`
	NameEN        string `json:"name_en"`
	Category      string `json:"category"`
	URL           string `json:"url"`
	NeedAuth      bool   `json:"needAuth"`
	NeedSignature bool   `json:"needSignature"`
	SystemParams  []struct {
		Name     string `json:"name"`
		Required bool   `json:"required"`
	} `json:"systemParams"`
	Body []struct {
		Field    string `json:"field"`
		Type     string `json:"type"`
		Required bool   `json:"required"`
	} `json:"body"`
}

// loadAllAPIs reads ALL-APIS.json keyed by "ns:name:ver".
func loadAllAPIs(t *testing.T, dir string) map[string]allAPIEntry {
	t.Helper()
	path := filepath.Join(dir, "ALL-APIS.json")
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read %s: %v", path, err)
	}
	var list []allAPIEntry
	if err := json.Unmarshal(b, &list); err != nil {
		t.Fatalf("parse %s: %v", path, err)
	}
	out := make(map[string]allAPIEntry, len(list))
	for _, e := range list {
		out[e.API] = e
	}
	return out
}

// requiredBodyFields is every top-level body field documented required: the
// entries whose field name carries no dot (dotted names are members of an
// object parameter, and the object is what the caller sends).
func requiredBodyFields(e allAPIEntry) []string {
	var out []string
	for _, f := range e.Body {
		if f.Required && !strings.Contains(f.Field, ".") {
			out = append(out, f.Field)
		}
	}
	sort.Strings(out)
	return out
}

// TestRegistryMatchesDocs asserts that every ali.API's flags and URL agree with
// ALL-APIS.json, and that the timestamp flag agrees with the per-API document.
func TestRegistryMatchesDocs(t *testing.T) {
	dir := docsDir(t)
	all := loadAllAPIs(t, dir)

	for _, api := range ali.All() {
		e, ok := all[api.Key()]
		if !ok {
			t.Errorf("%s: not present in ALL-APIS.json", api.Key())
			continue
		}
		if api.NeedAuth != e.NeedAuth {
			t.Errorf("%s: registry NeedAuth=%v, ALL-APIS.json needAuth=%v", api.Key(), api.NeedAuth, e.NeedAuth)
		}
		if api.NeedSig != e.NeedSignature {
			t.Errorf("%s: registry NeedSig=%v, ALL-APIS.json needSignature=%v", api.Key(), api.NeedSig, e.NeedSignature)
		}

		want := strings.ReplaceAll(e.URL, "{appKey}", "APPKEY")
		if !strings.Contains(e.URL, "{appKey}") {
			t.Errorf("%s: documented url %q has no {appKey} placeholder", api.Key(), e.URL)
		}
		if got := api.URL("https://gw.open.1688.com", "APPKEY"); got != want {
			t.Errorf("%s: URL()\n  got  %s\n  want %s", api.Key(), got, want)
		}
		if got := api.URL("https://gw.open.1688.com/", "APPKEY"); got != want {
			t.Errorf("%s: URL() with a trailing slash on the base\n  got  %s\n  want %s", api.Key(), got, want)
		}

		// The system parameters must line up with the flags: a call that needs
		// auth must be documented with access_token required, and so on.
		sys := map[string]bool{}
		for _, sp := range e.SystemParams {
			sys[sp.Name] = sp.Required
		}
		if sys[ali.TokenParam] != api.NeedAuth {
			t.Errorf("%s: systemParams %s required=%v but registry NeedAuth=%v", api.Key(), ali.TokenParam, sys[ali.TokenParam], api.NeedAuth)
		}
		if sys[ali.SignatureParam] != api.NeedSig {
			t.Errorf("%s: systemParams %s required=%v but registry NeedSig=%v", api.Key(), ali.SignatureParam, sys[ali.SignatureParam], api.NeedSig)
		}
		if sys[ali.TimestampParam] != api.NeedTS {
			t.Errorf("%s: systemParams %s required=%v but registry NeedTS=%v", api.Key(), ali.TimestampParam, sys[ali.TimestampParam], api.NeedTS)
		}

		// The per-API document carries the same three flags.
		raw := loadRawDoc(t, dir, api)
		if raw.NeedAuth != api.NeedAuth || raw.NeedSignature != api.NeedSig {
			t.Errorf("%s: raw doc needAuth=%v needSignature=%v, registry NeedAuth=%v NeedSig=%v", api.Key(), raw.NeedAuth, raw.NeedSignature, api.NeedAuth, api.NeedSig)
		}
		if raw.NeedTimestamp != api.NeedTS {
			t.Errorf("%s: raw doc needTimestamp=%v, registry NeedTS=%v", api.Key(), raw.NeedTimestamp, api.NeedTS)
		}
	}
}

// formCapture records the form every request posted, keyed by request path.
type formCapture struct {
	mu    sync.Mutex
	forms map[string]url.Values
}

func (c *formCapture) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	c.mu.Lock()
	c.forms[r.URL.Path] = r.PostForm
	c.mu.Unlock()
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(`{}`))
}

// TestRequestParamsComplete drives every Client method with a canned, valid
// input against a capturing HTTP server and asserts that each request carries
// every top-level body field ALL-APIS.json documents as required, plus the
// required system parameters, and that it was posted to the documented path.
// The three exported Params() builders are checked directly as well.
func TestRequestParamsComplete(t *testing.T) {
	dir := docsDir(t)
	all := loadAllAPIs(t, dir)

	cap := &formCapture{forms: map[string]url.Values{}}
	srv := httptest.NewServer(cap)
	defer srv.Close()

	client := ali.New(srv.URL, "APPKEY", "SECRET", "TOKEN")
	client.HTTP = srv.Client()
	client.Attempts = 1

	address := ali.Address{
		FullName: "张三", Mobile: "15251667788", Phone: "0571-88990077", PostCode: "310052",
		ProvinceText: "浙江省", CityText: "杭州市", AreaText: "滨江区", Address: "网商路699号", DistrictCode: "330108",
	}
	cargo := []ali.Cargo{ali.NewCargo(900000000000, "eb81c61de14f4adb405ffcc2c8a4a3fb", 2)}
	preview := ali.PreviewRequest{Address: address, Cargo: cargo}
	create := ali.CreateOrderRequest{Flow: ali.FlowGeneral, Address: address, Cargo: cargo, OutOrderID: "SHOP-1"}
	freight := ali.FreightQuery{OfferID: 900000000000, ToProvinceCode: "330000", ToCityCode: "330100", ToCountryCode: "330108", TotalNum: 2}
	freight.AddSku(3332085412530, 2)
	since := time.Date(2026, 3, 1, 0, 0, 0, 0, ali.CST)

	ctx := context.Background()
	// Every method returns an error here because the server answers {}; only
	// the captured request matters.
	calls := []struct {
		api  ali.API
		call func() error
	}{
		{ali.AccountBasic, func() error { _, err := client.AccountBasic(ctx); return err }},
		{ali.KeywordQuery, func() error {
			_, err := client.SearchOffers(ctx, ali.OfferQuery{Keyword: "dress", Country: "en", BeginPage: 1, PageSize: 20})
			return err
		}},
		{ali.KeywordSN, func() error {
			_, err := client.SearchNav(ctx, ali.SNQuery{Keyword: "dress", Language: "en_US", Region: "US", Currency: "USD"})
			return err
		}},
		{ali.ProductDetail, func() error {
			_, err := client.OfferDetail(ctx, ali.OfferDetailQuery{OfferID: 900000000000, Country: "en"})
			return err
		}},
		{ali.CategoryByID, func() error { _, err := client.CategoryByID(ctx, "en", 0, 0); return err }},
		{ali.FreightEstimate, func() error { _, err := client.EstimateFreight(ctx, freight); return err }},
		{ali.OrderPreview, func() error { _, err := client.PreviewOrder(ctx, preview); return err }},
		{ali.CreateCrossOrder, func() error { _, err := client.CreateOrder(ctx, create); return err }},
		{ali.AlipayURLGet, func() error { _, err := client.AlipayURL(ctx, ali.IDs{105581756010628640}); return err }},
		{ali.OrderBuyerView, func() error { _, err := client.OrderDetail(ctx, 105581756010628640); return err }},
		{ali.BuyerOrderList, func() error {
			_, err := client.BuyerOrders(ctx, ali.OrderListQuery{Page: 1, PageSize: 20, ModifyStartTime: since})
			return err
		}},
		{ali.TradeCancel, func() error {
			_, err := client.CancelOrder(ctx, 105581756010628640, ali.CancelBuyerCancel, "改主意了")
			return err
		}},
		{ali.LogisticsTrace, func() error { _, err := client.TraceOrder(ctx, 105581756010628640, ""); return err }},
		{ali.PushCursorList, func() error {
			_, err := client.CursorMessages(ctx, ali.PushQuery{CreateStartTime: since, Quantity: 50})
			return err
		}},
		{ali.PushQueryList, func() error {
			_, err := client.QueryMessages(ctx, ali.PushQuery{CreateStartTime: since, Page: 1, PageSize: 50})
			return err
		}},
		{ali.PushConfirm, func() error { _, err := client.ConfirmMessages(ctx, ali.IDs{123456}); return err }},
	}

	covered := map[string]bool{}
	for _, c := range calls {
		covered[c.api.Key()] = true
		e, ok := all[c.api.Key()]
		if !ok {
			t.Errorf("%s: not present in ALL-APIS.json", c.api.Key())
			continue
		}
		_ = c.call()

		wantPath := strings.TrimPrefix(strings.ReplaceAll(e.URL, "{appKey}", "APPKEY"), "https://gw.open.1688.com")
		cap.mu.Lock()
		form, posted := cap.forms[wantPath]
		var seen []string
		for p := range cap.forms {
			seen = append(seen, p)
		}
		cap.mu.Unlock()
		if !posted {
			sort.Strings(seen)
			t.Errorf("%s: nothing was posted to the documented path %s; paths seen: %v", c.api.Key(), wantPath, seen)
			continue
		}

		assertKeys(t, c.api.Key()+" (posted form)", formKeys(form), requiredBodyFields(e))
		for _, sp := range e.SystemParams {
			if sp.Required && form.Get(sp.Name) == "" {
				t.Errorf("%s: required system parameter %s missing from the posted form", c.api.Key(), sp.Name)
			}
		}
		if !e.NeedAuth && form.Get(ali.TokenParam) != "" {
			t.Errorf("%s: access_token was sent to an endpoint documented needAuth=false", c.api.Key())
		}
	}
	for _, api := range ali.All() {
		if !covered[api.Key()] {
			t.Errorf("%s is in ali.All() but no Client method exercises it above", api.Key())
		}
	}

	// The exported builders, independent of the transport.
	assertKeys(t, ali.OrderPreview.Key()+" PreviewRequest.Params()", paramKeys(preview.Params()), requiredBodyFields(all[ali.OrderPreview.Key()]))
	assertKeys(t, ali.CreateCrossOrder.Key()+" CreateOrderRequest.Params()", paramKeys(create.Params()), requiredBodyFields(all[ali.CreateCrossOrder.Key()]))
	assertKeys(t, ali.BuyerOrderList.Key()+" OrderListQuery.Params()", paramKeys(ali.OrderListQuery{}.Params()), requiredBodyFields(all[ali.BuyerOrderList.Key()]))
}

func assertKeys(t *testing.T, what string, have map[string]bool, required []string) {
	t.Helper()
	var missing []string
	for _, f := range required {
		if !have[f] {
			missing = append(missing, f)
		}
	}
	if len(missing) > 0 {
		var got []string
		for k := range have {
			got = append(got, k)
		}
		sort.Strings(got)
		t.Errorf("%s: documented required field(s) missing: %v\n    sent: %v", what, missing, got)
	}
}

func formKeys(form url.Values) map[string]bool {
	out := map[string]bool{}
	for k, vs := range form {
		if len(vs) > 0 && vs[0] != "" {
			out[k] = true
		}
	}
	return out
}

func paramKeys(p ali.Params) map[string]bool {
	out := map[string]bool{}
	for k, v := range p {
		if v != nil {
			out[k] = true
		}
	}
	return out
}
