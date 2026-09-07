package stub

// Doc-driven contract test for the stub gateway: every handler is invoked with
// a canned, valid request, the map it returns is flattened to dotted paths,
// and those paths are compared with 1688-api-docs/raw — never with the Go
// response types in internal/ali, so a mistake made identically on both sides
// still fails here.

import (
	"context"
	"encoding/json"
	"fmt"
	"html"
	"io"
	"log/slog"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"testing"
	"time"

	"marketplace/internal/ali"
	"marketplace/internal/stub/gen"
)

// ------------------------------------------------------------ corpus loading

// contractDocsDir is DOCS_DIR, else ../../1688-api-docs relative to this
// package. A relative DOCS_DIR is tried as given and then against the module
// root, since go test runs each package in its own directory. The test is
// skipped when the corpus is absent so container builds stay green.
func contractDocsDir(t *testing.T) string {
	t.Helper()
	var candidates []string
	if env := os.Getenv("DOCS_DIR"); env != "" {
		candidates = append(candidates, env)
		if !filepath.IsAbs(env) {
			if root := moduleRoot(); root != "" {
				candidates = append(candidates, filepath.Join(root, env))
			}
		}
	} else {
		candidates = append(candidates, filepath.Join("..", "..", "1688-api-docs"))
	}
	for _, dir := range candidates {
		if st, err := os.Stat(filepath.Join(dir, "raw")); err == nil && st.IsDir() {
			return dir
		}
	}
	t.Skipf("documentation corpus not found (tried %v; set DOCS_DIR)", candidates)
	return ""
}

func moduleRoot() string {
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return ""
		}
		dir = parent
	}
}

type paramNode struct {
	Name     string      `json:"name"`
	Type     string      `json:"type"`
	TypeName string      `json:"typeName"`
	Required bool        `json:"required"`
	Fields   []paramNode `json:"fields"`
}

type rawDoc struct {
	Namespace    string      `json:"namespace"`
	Name         string      `json:"name"`
	Version      int         `json:"version"`
	ReturnParams []paramNode `json:"apiReturnParamVOList"`
	Samples      []struct {
		Name   string `json:"name"`
		Sample string `json:"sample"`
	} `json:"apiDocSampleVOList"`
}

func loadRawDoc(t *testing.T, dir string, api ali.API) *rawDoc {
	t.Helper()
	path := filepath.Join(dir, "raw", fmt.Sprintf("%s.%s-%d.json", api.Namespace, api.Name, api.Version))
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("%s: no archived documentation: %v", api.DocID(), err)
	}
	var d rawDoc
	if err := json.Unmarshal(b, &d); err != nil {
		t.Fatalf("%s: parse %s: %v", api.DocID(), path, err)
	}
	return &d
}

// docShape is the documented response as three path sets: every path, the
// required paths, and the paths whose documented type is an opaque map
// (java.util.Map), beneath which anything may appear.
type docShape struct {
	all, required, opaque map[string]bool
}

func shapeOf(nodes []paramNode) docShape {
	s := docShape{all: map[string]bool{}, required: map[string]bool{}, opaque: map[string]bool{}}
	var walk func(nodes []paramNode, prefix string, required bool)
	walk = func(nodes []paramNode, prefix string, required bool) {
		for _, n := range nodes {
			p := prefix + n.Name
			s.all[p] = true
			req := required && n.Required
			if req {
				s.required[p] = true
			}
			if strings.Contains(n.Type, "Map") && len(n.Fields) == 0 {
				s.opaque[p] = true
			}
			if len(n.Fields) > 0 {
				walk(n.Fields, p+".", req)
			}
		}
	}
	walk(nodes, "", true)
	return s
}

// samplePaths flattens every JSON response sample of a doc into dotted paths;
// a field shown in a documented sample is documented.
func samplePaths(d *rawDoc) map[string]bool {
	out := map[string]bool{}
	for _, s := range d.Samples {
		if strings.Contains(s.Name, "入参") || strings.Contains(s.Name, "请求") || strings.Contains(s.Name, "说明") {
			continue
		}
		var v any
		if err := json.Unmarshal([]byte(strings.TrimSpace(html.UnescapeString(s.Sample))), &v); err != nil {
			continue
		}
		flattenJSON(v, "", out, nil)
	}
	return out
}

// flattenJSON walks decoded JSON into dotted paths, dropping array indexes.
// When empties is non-nil, paths whose value is null, an empty array or an
// empty object are recorded there: an empty collection cannot exhibit the
// shape of its elements.
func flattenJSON(v any, prefix string, out, empties map[string]bool) {
	switch t := v.(type) {
	case map[string]any:
		for k, child := range t {
			p := k
			if prefix != "" {
				p = prefix + "." + k
			}
			out[p] = true
			if empties != nil && isEmptyJSON(child) {
				empties[p] = true
			}
			flattenJSON(child, p, out, empties)
		}
	case []any:
		for _, child := range t {
			flattenJSON(child, prefix, out, empties)
		}
	}
}

func isEmptyJSON(v any) bool {
	switch t := v.(type) {
	case nil:
		return true
	case []any:
		return len(t) == 0
	case map[string]any:
		return len(t) == 0
	}
	return false
}

// emittedPaths marshals each handler response and flattens the union. A path
// counts as empty only if no response showed anything beneath it.
func emittedPaths(t *testing.T, resps []any) (paths, empties map[string]bool) {
	t.Helper()
	paths, empties = map[string]bool{}, map[string]bool{}
	for _, v := range resps {
		b, err := json.Marshal(v)
		if err != nil {
			t.Fatalf("marshal handler response: %v", err)
		}
		var decoded any
		if err := json.Unmarshal(b, &decoded); err != nil {
			t.Fatalf("re-decode handler response: %v", err)
		}
		flattenJSON(decoded, "", paths, empties)
	}
	for e := range empties {
		for p := range paths {
			if strings.HasPrefix(p, e+".") {
				delete(empties, e)
				break
			}
		}
	}
	return paths, empties
}

func underAny(p string, prefixes map[string]bool) (string, bool) {
	for prefix := range prefixes {
		if strings.HasPrefix(p, prefix+".") {
			return prefix, true
		}
	}
	return "", false
}

type allow struct{ path, reason string }

func findAllow(entries []allow, path string) (allow, bool) {
	for _, a := range entries {
		if a.path == path {
			return a, true
		}
		if strings.HasSuffix(a.path, ".*") {
			prefix := strings.TrimSuffix(a.path, ".*")
			if path == prefix || strings.HasPrefix(path, prefix+".") {
				return a, true
			}
		}
	}
	return allow{}, false
}

func report(t *testing.T, api, direction string, paths []string) {
	t.Helper()
	if len(paths) == 0 {
		return
	}
	sort.Strings(paths)
	t.Errorf("%s: %s: %d path(s):\n    %s", api, direction, len(paths), strings.Join(paths, "\n    "))
}

// ------------------------------------------------------------------ fixtures

func contractServer(t *testing.T, dir string) (*Server, *TestClock) {
	t.Helper()
	clock := NewTestClock(time.Date(2026, 3, 1, 9, 0, 0, 0, ali.CST))
	srv, err := New(Config{
		Base:        "http://stub.test",
		AppKey:      "DEVKEY",
		AccessToken: "devtoken",
		AppSecret:   "devsecret",
		Products:    120,
		DocsDir:     dir,
		Speed:       1,
		Sign:        "on",
		StateFile:   filepath.Join(t.TempDir(), "state.json"),
		Delays:      Delays{Pay: 2 * time.Minute, Ship: 10 * time.Minute, Step: 6 * time.Minute, Sign: 20 * time.Minute},
	}, slog.New(slog.NewTextHandler(io.Discard, &slog.HandlerOptions{Level: slog.LevelError})), clock)
	if err != nil {
		t.Fatalf("new server: %v", err)
	}
	return srv, clock
}

// orderableOffer is the first generated offer the preview accepts outright:
// published, with trade modes, a priced first sku with enough stock, and a
// quantity that satisfies the minimum and the batch rule.
func orderableOffer(srv *Server) (gen.Product, gen.SKU, int64) {
	for _, id := range srv.catalogue() {
		o := gen.Offer(id)
		if !o.Orderable() || len(tradeModes(o.ID)) == 0 || len(o.SKUs) == 0 {
			continue
		}
		sku := o.SKUs[0]
		qty := int64(o.MOQ)
		if qty < 1 {
			qty = 1
		}
		if o.BatchNumber > 0 && qty%int64(o.BatchNumber) != 0 {
			continue
		}
		if sku.Price == 0 || sku.Stock < qty {
			continue
		}
		return o, sku, qty
	}
	panic("no orderable offer in the generated catalogue")
}

const addressParam = `{"fullName":"张三","mobile":"15251667788","phone":"0571-88990077","postCode":"310052","provinceText":"浙江省","cityText":"杭州市","areaText":"滨江区","townText":"长河街道","address":"网商路699号","districtCode":"330108"}`

func cargoParam(o gen.Product, sku gen.SKU, qty int64) string {
	return fmt.Sprintf(`[{"offerId":%d,"specId":%q,"quantity":%d}]`, o.ID, sku.SpecID, qty)
}

func createForm(o gen.Product, sku gen.SKU, qty int64, outOrderID string) url.Values {
	form := url.Values{}
	form.Set("flow", "general")
	form.Set("addressParam", addressParam)
	form.Set("cargoParamList", cargoParam(o, sku, qty))
	form.Set("outOrderId", outOrderID)
	return form
}

func createdOrderID(t *testing.T, resp any) int64 {
	t.Helper()
	m, _ := resp.(map[string]any)
	result, _ := m["result"].(map[string]any)
	id, ok := parseInt64(mapStr(result, "orderId"))
	if !ok || id == 0 {
		b, _ := json.Marshal(resp)
		t.Fatalf("createCrossOrder did not return an order id: %s", b)
	}
	return id
}

func orderForm(id int64) url.Values {
	return url.Values{"webSite": {"1688"}, "orderId": {strconv.FormatInt(id, 10)}}
}

// tradeInfoOmissions is the part of the documented TradeInfo the stub does not
// simulate, with the reason for each. forList drops the members that only the
// order-detail document renders.
func tradeInfoOmissions(forList bool) []allow {
	const staged = "staged (step) payment is not simulated"
	const refunds = "refunds are not simulated"
	out := []allow{
		{"result.baseInfo.stepAgreementPath", staged},
		{"result.baseInfo.stepOrderList.*", staged},
		{"result.tradeTerms.phaseCondition", staged},
		{"result.tradeTerms.phaseDate", staged},
		{"result.baseInfo.officialSolutionCost", "official door-to-door pickup is not simulated"},
		{"result.baseInfo.officialSolutionOrderId", "official door-to-door pickup is not simulated"},
		{"result.baseInfo.preOrderId", "pre-orders are not simulated"},
		{"result.baseInfo.subBuyerLoginId", "sub-accounts are not simulated"},
		{"result.baseInfo.refundId", refunds},
		{"result.productItems.refundId", refunds},
		{"result.productItems.refundIdForAs", refunds},
		{"result.customs.*", "customs declarations are not simulated: the forwarder handles customs"},
		{"result.invoicingSettingModel.*", "invoicing is not simulated"},
		{"result.orderInvoiceInfo.*", "invoicing is not simulated"},
		{"result.overseasExtraAddress.*", "overseas addresses are not simulated: orders ship to the domestic consolidation warehouse"},
		{"result.quoteList.*", "caigou (RFQ) quotes are not simulated"},
		{"result.overseaLogisticsInfo.*", "official overseas logistics is not simulated"},
		{"result.orderBizInfo.*", "only the three documented flags that describe a plain purchase are emitted: odsCyd, creditOrder, dropshipping"},
	}
	if forList {
		return append(out,
			allow{"result.nativeLogistics.*", "the stub lists orders without the nativeLogistics block; the order-detail call is where the address and waybills are read"},
			allow{"result.baseInfo.buyerSubID", "sub-accounts are not simulated (documented on the list model only)"},
			allow{"result.baseInfo.sellerSubID", "sub-accounts are not simulated (documented on the list model only)"},
			allow{"result.baseInfo.currency", "documented on the list model only; the order-reading family is CNY throughout the stub"},
			allow{"result.baseInfo.relatedCode", "documented on the list model only, meaning undocumented"},
			allow{"result.productItems.relatedCode", "documented on the list model only, meaning undocumented"},
		)
	}
	return append(out,
		allow{"result.baseInfo.newStepOrderList.*", staged},
		allow{"result.encryptOutOrderInfo.*", "encrypted (privacy) orders are not simulated"},
		allow{"result.productItems.gmtCompleted", "per-line completion and pay-expiry times are not tracked by the stub's order store"},
		allow{"result.productItems.gmtPayExpireTime", "per-line completion and pay-expiry times are not tracked by the stub's order store"},
		allow{"result.orderRateInfo.buyerRateList.*", "ratings are not simulated; the two status fields are emitted"},
		allow{"result.orderRateInfo.sellerRateList.*", "ratings are not simulated; the two status fields are emitted"},
		allow{"result.nativeLogistics.logisticsItems.noLogisticsName", "type-2 (no logistics) waybills are not simulated; every stub order ships online"},
		allow{"result.nativeLogistics.logisticsItems.noLogisticsTel", "type-2 (no logistics) waybills are not simulated; every stub order ships online"},
		allow{"result.nativeLogistics.logisticsItems.noLogisticsBillNo", "type-2 (no logistics) waybills are not simulated; every stub order ships online"},
		allow{"result.nativeLogistics.logisticsItems.noLogisticsCondition", "type-2 (no logistics) waybills are not simulated; every stub order ships online"},
	)
}

// ------------------------------------------------------------------- the test

// TestStubEmitsDocumentedShapes asserts, for every handler, that (a) every
// documented required response field is emitted, and (b) nothing is emitted
// that the documentation does not describe, bar the invented set listed with
// reasons below.
func TestStubEmitsDocumentedShapes(t *testing.T) {
	dir := contractDocsDir(t)
	srv, clock := contractServer(t, dir)
	ctx := context.Background()
	o, sku, qty := orderableOffer(srv)

	// Order fixtures, through the gateway's own handlers and state API. Five
	// creates: the first is shipped (its detail carries a waybill and its
	// trace has steps), the second stays unpaid for the pay url, the third is
	// cancelled, the fourth runs to success, and the fifth splits into two
	// orders, which is the documented multi-order response shape.
	var createSingle, createMulti any
	var shippedID, unpaidID, cancelID, doneID int64
	for i := 1; i <= 5; i++ {
		resp := srv.handleCreateCrossOrder(ctx, createForm(o, sku, qty, "CONTRACT-"+strconv.Itoa(i)))
		switch i {
		case 1:
			createSingle = resp
			shippedID = createdOrderID(t, resp)
		case 2:
			unpaidID = createdOrderID(t, resp)
		case 3:
			cancelID = createdOrderID(t, resp)
		case 4:
			doneID = createdOrderID(t, resp)
		case 5:
			createMulti = resp
		}
	}
	if m, _ := createMulti.(map[string]any); m != nil {
		if r, _ := m["result"].(map[string]any); r != nil {
			if list, _ := r["orderList"].([]any); len(list) != 2 {
				t.Fatalf("fifth create should have split into two orders, got %d", len(list))
			}
		}
	}
	for _, step := range []struct {
		id int64
		to string
	}{{shippedID, StatusWaitReceive}, {doneID, StatusSuccess}} {
		events, err := srv.store.AdvanceTo(step.id, step.to)
		if err != nil {
			t.Fatalf("advance %d to %s: %v", step.id, step.to, err)
		}
		srv.Publish(ctx, events)
	}
	clock.Advance(11 * time.Second) // past the CLOSE_ORDER_TOO_FAST window

	cancelForm := url.Values{}
	cancelForm.Set("webSite", "1688")
	cancelForm.Set("tradeID", strconv.FormatInt(cancelID, 10))
	cancelForm.Set("cancelReason", "buyerCancel")
	cancelForm.Set("remark", "改主意了")
	cancelResp := srv.handleTradeCancel(ctx, cancelForm)
	if m, _ := cancelResp.(map[string]any); m == nil || m["success"] != true {
		t.Fatalf("cancel fixture failed: %v", cancelResp)
	}

	// Push fixtures: query first (it does not confirm), then confirm one of
	// the ids it returned, then drain with the cursor call.
	pushQuery := srv.handlePushQuery(url.Values{})
	var confirmID int64
	if page, _ := pushQuery.(map[string]any)["pushMessagePage"].(map[string]any); page != nil {
		if datas, _ := page["datas"].([]map[string]any); len(datas) > 0 {
			confirmID, _ = datas[0]["msgId"].(int64)
		}
	}
	if confirmID == 0 {
		t.Fatal("no push messages were queued by the order fixtures")
	}
	confirmForm := url.Values{}
	confirmForm.Set("msgIdList", "["+strconv.FormatInt(confirmID, 10)+"]")

	leaf := gen.Archetypes()[0].LeafCategories[0].ID

	// also names another API whose document describes the same wire model,
	// for direction (b): the list returns TradeInfo[], and its document
	// renders a shorter TradeInfo than the order-detail document does.
	cases := []struct {
		api   ali.API
		name  string
		also  []ali.API
		resps []any
	}{
		{api: ali.AccountBasic, name: "account.basic", resps: []any{srv.handleAccountBasic(url.Values{})}},
		{api: ali.KeywordQuery, name: "keywordQuery", resps: []any{srv.handleKeywordQuery(url.Values{"offerQueryParam": {`{"keyword":"","country":"en","beginPage":1,"pageSize":10}`}})}},
		{api: ali.KeywordSN, name: "keywordSNQuery", resps: []any{srv.handleKeywordSNQuery(url.Values{"snParams": {`{"keyword":"dress","language":"en_US","region":"US","currency":"USD"}`}})}},
		{api: ali.ProductDetail, name: "queryProductDetail", resps: []any{
			srv.handleProductDetail(url.Values{"offerDetailParam": {fmt.Sprintf(`{"offerId":%d,"country":"en"}`, o.ID)}}),
			// Every 13th offer carries a certificate list.
			srv.handleProductDetail(url.Values{"offerDetailParam": {fmt.Sprintf(`{"offerId":%d,"country":"en"}`, gen.CatalogueBase+(13-gen.CatalogueBase%13)%13)}}),
		}},
		{api: ali.CategoryByID, name: "category", resps: []any{
			srv.handleCategoryByID(url.Values{"language": {"en"}, "categoryId": {"0"}}),
			srv.handleCategoryByID(url.Values{"language": {"en"}, "categoryId": {strconv.FormatInt(leaf, 10)}}),
		}},
		{api: ali.FreightEstimate, name: "freight.estimate", resps: []any{srv.handleFreightEstimate(url.Values{"productFreightQueryParamsNew": {fmt.Sprintf(`{"offerId":%d,"toProvinceCode":"330000","toCityCode":"330100","toCountryCode":"330108","totalNum":%d,"logisticsSkuNumModels":[{"skuId":"%d","number":%d}]}`, o.ID, qty, sku.SkuID, qty)}})}},
		{api: ali.OrderPreview, name: "createOrder.preview", resps: []any{srv.handlePreview(url.Values{"addressParam": {addressParam}, "cargoParamList": {cargoParam(o, sku, qty)}})}},
		{api: ali.CreateCrossOrder, name: "createCrossOrder", resps: []any{createSingle, createMulti}},
		{api: ali.AlipayURLGet, name: "alipay.url.get", resps: []any{srv.handleAlipayURL(url.Values{"orderIdList": {"[" + strconv.FormatInt(unpaidID, 10) + "]"}})}},
		{api: ali.OrderBuyerView, name: "trade.get.buyerView", resps: []any{
			srv.handleOrderBuyerView(orderForm(shippedID)),
			srv.handleOrderBuyerView(orderForm(cancelID)),
			srv.handleOrderBuyerView(orderForm(doneID)),
		}},
		{api: ali.BuyerOrderList, name: "getBuyerOrderList", also: []ali.API{ali.OrderBuyerView}, resps: []any{srv.handleBuyerOrderList(url.Values{})}},
		{api: ali.TradeCancel, name: "trade.cancel", resps: []any{cancelResp}},
		{api: ali.LogisticsTrace, name: "getLogisticsTraceInfo", resps: []any{srv.handleLogisticsTrace(orderForm(shippedID))}},
		{api: ali.PushQueryList, name: "push.query.messageList", resps: []any{pushQuery}},
		{api: ali.PushConfirm, name: "push.message.confirm", resps: []any{srv.handlePushConfirm(confirmForm)}},
		{api: ali.PushCursorList, name: "push.cursor.messageList", resps: []any{srv.handlePushCursor(url.Values{})}},
	}

	// (a) Documented required fields the stub deliberately does not emit.
	notEmitted := map[string][]allow{
		ali.ProductDetail.DocID(): {
			{"result.result.productSkuInfos.foreignCurrencyPrice", "foreign-currency quotes are returned only for a currency the caller asks for; the stub quotes CNY"},
			{"result.result.productSkuInfos.foreignCurrencyPromotionPrice", "foreign-currency quotes are returned only for a currency the caller asks for; the stub quotes CNY"},
			{"result.result.productSkuInfos.foreignCurrencyRetailPrice", "foreign-currency quotes are returned only for a currency the caller asks for; the stub quotes CNY"},
			{"result.result.productSaleInfo.foreignCurrencyPrice", "foreign-currency quotes are returned only for a currency the caller asks for; the stub quotes CNY"},
			{"result.result.productSaleInfo.foreignCurrencyRetailPrice", "foreign-currency quotes are returned only for a currency the caller asks for; the stub quotes CNY"},
			{"result.result.productSaleInfo.priceRangeList.foreignCurrencyPrice", "foreign-currency quotes are returned only for a currency the caller asks for; the stub quotes CNY"},
			{"result.result.productSaleInfo.priceRangeList.foreignCurrencyPromotionPrice", "foreign-currency quotes are returned only for a currency the caller asks for; the stub quotes CNY"},
		},
		ali.OrderBuyerView.DocID(): tradeInfoOmissions(false),
		ali.BuyerOrderList.DocID(): tradeInfoOmissions(true),
	}

	// (b) Emitted fields that are not documented anywhere: the invented set.
	invented := map[string][]allow{}

	covered := map[string]bool{}
	for _, c := range cases {
		covered[c.api.DocID()] = true
		t.Run(c.name, func(t *testing.T) {
			d := loadRawDoc(t, dir, c.api)
			shape := shapeOf(d.ReturnParams)
			samples := samplePaths(d)
			documented := map[string]string{} // path -> which API documents it
			for p := range shape.all {
				documented[p] = ""
			}
			for _, other := range c.also {
				for p := range shapeOf(loadRawDoc(t, dir, other).ReturnParams).all {
					if _, ok := documented[p]; !ok {
						documented[p] = other.DocID()
					}
				}
			}
			emitted, empties := emittedPaths(t, c.resps)

			// (a) required documented -> emitted.
			var missing []string
			usedOmit := map[string]bool{}
			vacuous := map[string]bool{}
			for p := range shape.required {
				if emitted[p] {
					continue
				}
				if prefix, ok := underAny(p, empties); ok {
					vacuous[prefix] = true
					continue
				}
				if a, ok := findAllow(notEmitted[c.api.DocID()], p); ok {
					usedOmit[a.path] = true
					continue
				}
				missing = append(missing, p)
			}
			for prefix := range vacuous {
				t.Logf("%s: %s is emitted empty, so its documented members cannot be checked from these responses", c.api.DocID(), prefix)
			}
			for _, a := range notEmitted[c.api.DocID()] {
				if !usedOmit[a.path] {
					t.Errorf("%s: stale notEmitted entry %q (%s): it matches no missing required path", c.api.DocID(), a.path, a.reason)
				}
			}
			report(t, c.api.DocID(), "documented required field not emitted by the stub", missing)

			// (b) emitted -> documented.
			var undocumented, viaOther []string
			usedInvented := map[string]bool{}
			for p := range emitted {
				if by, ok := documented[p]; ok {
					if by != "" {
						viaOther = append(viaOther, p+" ("+by+")")
					}
					continue
				}
				if _, ok := underAny(p, shape.opaque); ok {
					continue // beneath a java.util.Map: the message payload, documented per topic
				}
				if samples[p] {
					t.Logf("%s: %s is documented by a response sample only", c.api.DocID(), p)
					continue
				}
				if a, ok := findAllow(invented[c.api.DocID()], p); ok {
					usedInvented[a.path] = true
					t.Logf("%s: %s is invented: %s", c.api.DocID(), p, a.reason)
					continue
				}
				undocumented = append(undocumented, p)
			}
			if len(viaOther) > 0 {
				sort.Strings(viaOther)
				t.Logf("%s: %d emitted path(s) documented by a sibling API rather than this document:\n    %s", c.api.DocID(), len(viaOther), strings.Join(viaOther, "\n    "))
			}
			for _, a := range invented[c.api.DocID()] {
				if !usedInvented[a.path] {
					t.Errorf("%s: stale invented entry %q (%s): the stub no longer emits it", c.api.DocID(), a.path, a.reason)
				}
			}
			report(t, c.api.DocID(), "stub emits a field the documentation does not describe", undocumented)
		})
	}
	for _, api := range ali.All() {
		if !covered[api.DocID()] {
			t.Errorf("%s is in ali.All() but no stub handler is exercised above", api.DocID())
		}
	}
}
