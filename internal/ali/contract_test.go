package ali_test

// Doc-driven contract tests for the response models in internal/ali.
//
// These compare our Go types against the ARCHIVED 1688 DOCUMENTATION in
// 1688-api-docs, never against the stub gateway: a mistake made identically in
// internal/ali and internal/stub is invisible to an end-to-end test, and is
// exactly what these catch.
//
// The helpers at the bottom of this file (corpus loading, path flattening,
// struct reflection) are shared by the other *_test.go files in this package.

import (
	"encoding/json"
	"fmt"
	"html"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"testing"
	"time"

	"marketplace/internal/ali"
)

// TestResponseFieldsModelled walks apiReturnParamVOList of every API in
// ali.All() into a set of dotted paths and reflects our response struct into
// the same shape from its json tags, then asserts both directions:
//
//	(a) every Go json tag path is documented — the direction that catches OUR
//	    typos (orderPreviewResult vs the documented orderPreviewResuslt, say);
//	(b) every documented path is modelled, or is listed below with a reason.
//
// ENVELOPE ROOTS. The docs describe the whole wire body; our structs describe
// the payload the generic decoders hand back. So the comparison starts at a
// documented root, chosen per API by reading both the docs and the Go type:
//
//	API                    Go type               root             why
//	----------------------------------------------------------------------------------------------
//	alibaba.account.basic  SimpleAccountInfo     result           family B: DecodeB unwraps result
//	keywordQuery           OfferPage             result.result    family A: DecodeA unwraps the
//	keywordSNQuery         SearchNav (elements)  result.result      inner result; the outer
//	queryProductDetail     OfferDetail           result.result      success/code/message wrapper
//	category.getById       Category              result.result      is read generically
//	product.freight.est.   ProductFreight        result.result
//	createOrder.preview    Preview               (top level)      bespoke: the body IS the payload
//	createCrossOrder       CreateResult          result           bespoke: docs wrap the payload in
//	                                                                result beside success/code/message;
//	                                                                DecodeFlat unwraps it
//	alipay.url.get         PayURL                (top level)      bespoke, flat
//	trade.get.buyerView    TradeInfo             result           family B
//	getBuyerOrderList      OrderListResult       (top level)      bespoke: result[] + totalRecord,
//	                                                                no success field at all
//	trade.cancel           CancelResult          (top level)      bespoke, flat
//	getLogisticsTraceInfo  LogisticsTraceResult  (top level)      bespoke: logisticsTrace beside success
//	push.cursor.messageList PushMessage (elems)  pushMessageList  the anonymous wrapper in
//	push.query.messageList PushMessagePage       pushMessagePage    CursorMessages holds the list
//	push.message.confirm   (none)                —                decoded into an anonymous struct
//	                                                                in ConfirmMessages; covered by
//	                                                                sample_test.go and the stub test
func TestResponseFieldsModelled(t *testing.T) {
	dir := docsDir(t)

	// alsoDoc names another API whose documented subtree describes the same
	// wire model, for direction (a) only: getBuyerOrderList returns
	// TradeInfo[], but its document renders a shorter TradeInfo than
	// trade.get.buyerView does (no logisticsItems, no cardPay, twelve fewer
	// baseInfo members), and the struct is one type because the wire model is.
	type alsoDoc struct {
		api    ali.API
		root   string // documented root in the other API
		prefix string // where that subtree sits among this API's paths
		reason string
	}
	type modelled struct {
		api  ali.API
		root string
		typ  reflect.Type // nil: no exported response struct, see note
		also []alsoDoc
		note string
	}
	table := []modelled{
		{api: ali.AccountBasic, root: "result", typ: reflect.TypeOf(ali.SimpleAccountInfo{})},
		{api: ali.KeywordQuery, root: "result.result", typ: reflect.TypeOf(ali.OfferPage{})},
		{api: ali.KeywordSN, root: "result.result", typ: reflect.TypeOf(ali.SearchNav{})},
		{api: ali.ProductDetail, root: "result.result", typ: reflect.TypeOf(ali.OfferDetail{})},
		{api: ali.CategoryByID, root: "result.result", typ: reflect.TypeOf(ali.Category{})},
		{api: ali.FreightEstimate, root: "result.result", typ: reflect.TypeOf(ali.ProductFreight{})},
		{api: ali.OrderPreview, root: "", typ: reflect.TypeOf(ali.Preview{})},
		{api: ali.CreateCrossOrder, root: "result", typ: reflect.TypeOf(ali.CreateResult{})},
		{api: ali.AlipayURLGet, root: "", typ: reflect.TypeOf(ali.PayURL{})},
		{api: ali.OrderBuyerView, root: "result", typ: reflect.TypeOf(ali.TradeInfo{})},
		{api: ali.BuyerOrderList, root: "", typ: reflect.TypeOf(ali.OrderListResult{}), also: []alsoDoc{
			{ali.OrderBuyerView, "result", "result.", "getBuyerOrderList returns TradeInfo[]; the full TradeInfo model is documented on trade.get.buyerView and the list document renders a shorter one"},
		}},
		{api: ali.TradeCancel, root: "", typ: reflect.TypeOf(ali.CancelResult{})},
		{api: ali.LogisticsTrace, root: "", typ: reflect.TypeOf(ali.LogisticsTraceResult{})},
		{api: ali.PushCursorList, root: "pushMessageList", typ: reflect.TypeOf(ali.PushMessage{})},
		{api: ali.PushQueryList, root: "pushMessagePage", typ: reflect.TypeOf(ali.PushMessagePage{})},
		{api: ali.PushConfirm, root: "", typ: nil, note: "decoded into an anonymous {isSuccess} struct inside ConfirmMessages; sample_test.go asserts isSuccess decodes and the stub test asserts it is emitted"},
	}

	// Direction (a) allowlist: Go json tags that are NOT in the documented
	// parameter table. Every entry needs a reason, must still be a real Go path
	// (a stale entry fails), and its leaf name must occur somewhere in the
	// documentation corpus — a table or a sample of any API in ali.All() — so a
	// misspelling cannot hide behind an allowlist entry.
	allowGo := map[string][]allow{
		ali.KeywordQuery.DocID(): {
			{"data.sellerDataInfo.collect30DayWithin48HPercent", "SellerDataInfo is one Go type shared by search and detail; this member is documented on queryProductDetail only"},
			{"data.sellerDataInfo.qualityRefundWithin30Day", "SellerDataInfo is one Go type shared by search and detail; this member is documented on queryProductDetail only"},
		},
		ali.ProductDetail.DocID(): {
			{"sellerDataInfo.tpYear", "SellerDataInfo is one Go type shared by search and detail; tpYear is documented on keywordQuery only"},
		},
		ali.KeywordSN.DocID(): {
			{"children.children", "one SearchNav type serves both navigation levels; the documented sub-navigation model has no children member (omitempty, never populated)"},
		},
		ali.CategoryByID.DocID(): {
			{"children.children", "one Category type serves both levels; the documented child model has no children member and the sample sends null"},
		},
		ali.CreateCrossOrder.DocID(): {
			{"orderList.discount", "absent from the BizSimpleOrder table but present in the documented multi-order sample; the order relay reads it"},
			{"orderList.sumPaymentNoCarriageFromClient", "absent from the BizSimpleOrder table but present in the documented multi-order sample"},
			{"orderList.mergePay", "absent from the BizSimpleOrder table but present in the documented multi-order sample"},
			{"orderList.chooseFreeFreight", "absent from the BizSimpleOrder table but present in the documented multi-order sample"},
		},
		ali.PushCursorList.DocID(): {
			{"topicGroup", "the PushMessage table lists five members; topicGroup/appKey/topicName appear in both documented replay samples and in every push envelope"},
			{"appKey", "the PushMessage table lists five members; topicGroup/appKey/topicName appear in both documented replay samples and in every push envelope"},
			{"topicName", "the PushMessage table lists five members; topicGroup/appKey/topicName appear in both documented replay samples and in every push envelope"},
		},
		ali.PushQueryList.DocID(): {
			{"datas.topicGroup", "the PushMessage table lists five members; topicGroup/appKey/topicName appear in both documented replay samples and in every push envelope"},
			{"datas.appKey", "the PushMessage table lists five members; topicGroup/appKey/topicName appear in both documented replay samples and in every push envelope"},
			{"datas.topicName", "the PushMessage table lists five members; topicGroup/appKey/topicName appear in both documented replay samples and in every push envelope"},
		},
	}

	// Direction (b) allowlist: documented paths that are deliberately not
	// modelled. "x.*" covers x and everything beneath it. Each entry must match
	// at least one currently-unmodelled documented path, or it is stale.
	allowDoc := map[string][]allow{
		ali.AccountBasic.DocID(): {
			{"saleRate", "seller-side rating counters; this platform is a buyer"},
			{"rateNum", "seller-side rating counters; this platform is a buyer"},
			{"rateSum", "seller-side rating counters; this platform is a buyer"},
			{"buyRate", "rating counter, unused"},
			{"maturity", "unused account metadata"},
			{"memo", "unused account metadata"},
			{"communityLevel", "unused account metadata"},
			{"domainInPlatforms", "unused account metadata"},
			{"industry", "unused account metadata"},
			{"product", "unused account metadata"},
			{"department", "unused account metadata"},
			{"addressLocation", "unused account metadata"},
			{"isPm", "internal 1688 flags with no documented meaning for a buyer"},
			{"pm", "internal 1688 flags with no documented meaning for a buyer"},
			{"fm", "internal 1688 flags with no documented meaning for a buyer"},
		},
		ali.OrderBuyerView.DocID(): {
			{"orderBizInfo.*", "kept verbatim in TradeInfo.Extra (json:\"-\"); nothing in the order pipeline reads it"},
			{"orderInvoiceInfo.*", "kept verbatim in TradeInfo.Extra; invoicing is not offered"},
			{"orderRateInfo.*", "kept verbatim in TradeInfo.Extra; ratings are not surfaced"},
			{"overseasExtraAddress.*", "kept verbatim in TradeInfo.Extra; orders ship to the domestic consolidation warehouse"},
			{"customs.*", "kept verbatim in TradeInfo.Extra; customs is handled by the forwarder"},
			{"quoteList.*", "kept verbatim in TradeInfo.Extra; caigou quotes are not used"},
			{"fromEncryptOrder", "kept verbatim in TradeInfo.Extra; encrypted orders are not used"},
			{"encryptOutOrderInfo.*", "kept verbatim in TradeInfo.Extra; encrypted orders are not used"},
			{"overseaLogisticsInfo.*", "kept verbatim in TradeInfo.Extra; official overseas logistics is not used"},
			{"invoicingSettingModel.*", "kept verbatim in TradeInfo.Extra; invoicing is not offered"},
			{"baseInfo.stepAgreementPath", "staged payment is not used; kept in OrderBaseInfo.Extra"},
			{"baseInfo.stepOrderList.*", "staged payment is not used; kept in OrderBaseInfo.Extra"},
			{"baseInfo.newStepOrderList.*", "staged payment is not used; kept in OrderBaseInfo.Extra"},
		},
		ali.BuyerOrderList.DocID(): {
			{"result.orderBizInfo.*", "kept verbatim in TradeInfo.Extra; see trade.get.buyerView"},
			{"result.orderInvoiceInfo.*", "kept verbatim in TradeInfo.Extra; see trade.get.buyerView"},
			{"result.orderRateInfo.*", "kept verbatim in TradeInfo.Extra; see trade.get.buyerView"},
			{"result.overseasExtraAddress.*", "kept verbatim in TradeInfo.Extra; see trade.get.buyerView"},
			{"result.customs.*", "kept verbatim in TradeInfo.Extra; see trade.get.buyerView"},
			{"result.quoteList.*", "kept verbatim in TradeInfo.Extra; see trade.get.buyerView"},
			{"result.overseaLogisticsInfo.*", "kept verbatim in TradeInfo.Extra; see trade.get.buyerView"},
			{"result.invoicingSettingModel.*", "kept verbatim in TradeInfo.Extra; see trade.get.buyerView"},
			{"result.baseInfo.stepAgreementPath", "staged payment is not used; kept in OrderBaseInfo.Extra"},
			{"result.baseInfo.stepOrderList.*", "staged payment is not used; kept in OrderBaseInfo.Extra"},
			{"result.baseInfo.buyerSubID", "documented on the list model only: sub-account ids, unused; kept in OrderBaseInfo.Extra"},
			{"result.baseInfo.sellerSubID", "documented on the list model only: sub-account ids, unused; kept in OrderBaseInfo.Extra"},
			{"result.baseInfo.currency", "documented on the list model only; the order-reading family is CNY and pricing converts itself; kept in OrderBaseInfo.Extra"},
			{"result.baseInfo.relatedCode", "documented on the list model only, meaning undocumented; kept in OrderBaseInfo.Extra"},
			{"result.productItems.relatedCode", "documented on the list model only, meaning undocumented; ProductItem has no Extra and nothing reads it"},
		},
	}

	// Leaf-name evidence for allowGo entries: every field name that occurs
	// anywhere in the corpus for the APIs we speak.
	corpusLeaves := map[string]bool{}
	docs := map[string]*rawDoc{}
	for _, api := range ali.All() {
		d := loadRawDoc(t, dir, api)
		docs[api.DocID()] = d
		for p := range docPathSet(d.ReturnParams, "") {
			corpusLeaves[lastSegment(p)] = true
		}
		for _, s := range responseSamples(d) {
			for p := range jsonPathSet(s.decoded) {
				corpusLeaves[lastSegment(p)] = true
			}
		}
	}

	covered := map[string]bool{}
	for _, m := range table {
		covered[m.api.DocID()] = true
		t.Run(m.api.Name, func(t *testing.T) {
			d := docs[m.api.DocID()]
			if m.typ == nil {
				t.Logf("%s: no exported response struct: %s", m.api.DocID(), m.note)
				return
			}
			rootNodes, ok := subtree(d.ReturnParams, m.root)
			if !ok {
				t.Fatalf("%s: documented root %q not found in apiReturnParamVOList", m.api.DocID(), m.root)
			}
			docSet := docPathSet(rootNodes, "")
			goSet := goPathSet(m.typ)

			alsoSet := map[string]alsoDoc{}
			for _, a := range m.also {
				nodes, ok := subtree(docs[a.api.DocID()].ReturnParams, a.root)
				if !ok {
					t.Fatalf("%s: documented root %q not found in %s", m.api.DocID(), a.root, a.api.DocID())
				}
				for p := range docPathSet(nodes, a.prefix) {
					alsoSet[p] = a
				}
			}

			// (a) Go -> docs.
			var missingInDocs, viaOther []string
			usedGo := map[string]bool{}
			for p := range goSet {
				if docSet[p] {
					continue
				}
				if _, ok := alsoSet[p]; ok {
					viaOther = append(viaOther, p)
					continue
				}
				if a, ok := findAllow(allowGo[m.api.DocID()], p); ok {
					usedGo[a.path] = true
					if !corpusLeaves[lastSegment(p)] {
						t.Errorf("%s: Go tag %q is allowlisted (%s) but its leaf %q occurs nowhere in the documentation corpus — a misspelling cannot be allowlisted", m.api.DocID(), p, a.reason, lastSegment(p))
					}
					t.Logf("%s: Go tag %q not in the parameter table, allowed: %s", m.api.DocID(), p, a.reason)
					continue
				}
				missingInDocs = append(missingInDocs, p)
			}
			for _, a := range allowGo[m.api.DocID()] {
				if !usedGo[a.path] {
					t.Errorf("%s: stale allowGo entry %q (%s): it is either not a Go tag path or it is now documented", m.api.DocID(), a.path, a.reason)
				}
			}
			if len(viaOther) > 0 {
				sort.Strings(viaOther)
				t.Logf("%s: %d Go tag(s) absent from this document but documented on %s (%s):\n    %s",
					m.api.DocID(), len(viaOther), alsoSet[viaOther[0]].api.DocID(), alsoSet[viaOther[0]].reason, strings.Join(viaOther, "\n    "))
			}
			reportPaths(t, m.api.DocID(), "Go json tag not documented under root "+quoteRoot(m.root)+" (our typo?)", missingInDocs)

			// (b) docs -> Go.
			var unmodelled []string
			usedDoc := map[string]bool{}
			for p := range docSet {
				if goSet[p] {
					continue
				}
				if a, ok := findAllow(allowDoc[m.api.DocID()], p); ok {
					usedDoc[a.path] = true
					continue
				}
				unmodelled = append(unmodelled, p)
			}
			for _, a := range allowDoc[m.api.DocID()] {
				if !usedDoc[a.path] {
					t.Errorf("%s: stale allowDoc entry %q (%s): it matches no unmodelled documented path", m.api.DocID(), a.path, a.reason)
				}
			}
			reportPaths(t, m.api.DocID(), "documented path not modelled by "+m.typ.String()+" and not allowlisted", unmodelled)
		})
	}
	for _, api := range ali.All() {
		if !covered[api.DocID()] {
			t.Errorf("%s is in ali.All() but has no entry in the contract table above", api.DocID())
		}
	}
}

// ---------------------------------------------------------------- allowlists

type allow struct{ path, reason string }

// findAllow returns the entry covering path: an exact match, or a "prefix.*"
// entry covering prefix itself and everything beneath it.
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

func reportPaths(t *testing.T, api, direction string, paths []string) {
	t.Helper()
	if len(paths) == 0 {
		return
	}
	sort.Strings(paths)
	t.Errorf("%s: %s: %d path(s):\n    %s", api, direction, len(paths), strings.Join(paths, "\n    "))
}

func quoteRoot(root string) string {
	if root == "" {
		return "(top level)"
	}
	return "\"" + root + "\""
}

func lastSegment(p string) string {
	if i := strings.LastIndexByte(p, '.'); i >= 0 {
		return p[i+1:]
	}
	return p
}

// ------------------------------------------------------------ corpus loading

// docsDir is DOCS_DIR, else ../../1688-api-docs relative to this package. A
// relative DOCS_DIR is tried as given and then against the module root, since
// go test runs each package in its own directory. The test is skipped when the
// corpus is absent so container builds stay green.
func docsDir(t *testing.T) string {
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

// moduleRoot walks up from the working directory to the directory holding go.mod.
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

// paramNode is one entry of apiAppParamVOList / apiReturnParamVOList. Fields
// is the resolved nested model; a type or typeName ending in "[]" is an array
// of that model, which the path walkers flatten.
type paramNode struct {
	Name            string      `json:"name"`
	Type            string      `json:"type"`
	TypeName        string      `json:"typeName"`
	Required        bool        `json:"required"`
	Description     string      `json:"description"`
	ExampleValue    string      `json:"exampleValue"`
	ComplexTypeFlag bool        `json:"complexTypeFlag"`
	Fields          []paramNode `json:"fields"`
}

type docSample struct {
	Name   string `json:"name"`
	Sample string `json:"sample"`
	Type   string `json:"type"`

	decoded any // set by responseSamples when the sample is valid JSON
}

type docErrorCode struct {
	Code     string `json:"code"`
	Desc     string `json:"desc"`
	HowToFix string `json:"howToFix"`
}

type rawDoc struct {
	Namespace     string         `json:"namespace"`
	Name          string         `json:"name"`
	Version       int            `json:"version"`
	NeedAuth      bool           `json:"needAuth"`
	NeedSignature bool           `json:"needSignature"`
	NeedTimestamp bool           `json:"needTimestamp"`
	AppParams     []paramNode    `json:"apiAppParamVOList"`
	ReturnParams  []paramNode    `json:"apiReturnParamVOList"`
	ErrorCodes    []docErrorCode `json:"apiErrorCodeVOList"`
	Samples       []docSample    `json:"apiDocSampleVOList"`
	ReturnExample string         `json:"returnExample"`
}

// loadRawDoc reads 1688-api-docs/raw/<namespace>.<name>-<version>.json.
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
	if d.Namespace != api.Namespace || d.Name != api.Name || d.Version != api.Version {
		t.Fatalf("%s: %s describes %s:%s-%d", api.DocID(), path, d.Namespace, d.Name, d.Version)
	}
	return &d
}

// unescapeSample undoes the HTML escaping the archive applied to every sample.
func unescapeSample(s string) string {
	return strings.TrimSpace(html.UnescapeString(s))
}

// isResponseSample classifies a sample by its name. Request samples (入参,
// 请求) and prose (说明) are excluded; anything that says 出参 / 返回 / 结果
// is a response.
func isResponseSample(name string) bool {
	if strings.Contains(name, "入参") || strings.Contains(name, "请求") || strings.Contains(name, "说明") {
		return false
	}
	return strings.Contains(name, "出参") || strings.Contains(name, "返回") || strings.Contains(name, "结果")
}

// responseSamples returns the documented response samples that are valid JSON,
// with decoded set. Samples that are not JSON are dropped here; callers that
// want to log them use the raw Samples list.
func responseSamples(d *rawDoc) []docSample {
	var out []docSample
	for _, s := range d.Samples {
		if !isResponseSample(s.Name) {
			continue
		}
		var v any
		if err := json.Unmarshal([]byte(unescapeSample(s.Sample)), &v); err != nil {
			continue
		}
		s.decoded = v
		out = append(out, s)
	}
	return out
}

// sampleByName returns the first sample whose name contains substr.
func sampleByName(d *rawDoc, substr string) (docSample, bool) {
	for _, s := range d.Samples {
		if strings.Contains(s.Name, substr) {
			return s, true
		}
	}
	return docSample{}, false
}

// ------------------------------------------------------------ path walkers

// docPathSet flattens a parameter tree into dotted paths. Arrays are flattened:
// the element model's members hang directly off the array's path.
func docPathSet(nodes []paramNode, prefix string) map[string]bool {
	out := map[string]bool{}
	var walk func(nodes []paramNode, prefix string)
	walk = func(nodes []paramNode, prefix string) {
		for _, n := range nodes {
			p := prefix + n.Name
			out[p] = true
			if len(n.Fields) > 0 {
				walk(n.Fields, p+".")
			}
		}
	}
	walk(nodes, prefix)
	return out
}

// subtree navigates a dotted root ("result.result") and returns the members
// beneath it. An empty root is the top level.
func subtree(nodes []paramNode, root string) ([]paramNode, bool) {
	if root == "" {
		return nodes, true
	}
	cur := nodes
	for _, seg := range strings.Split(root, ".") {
		var next []paramNode
		found := false
		for _, n := range cur {
			if n.Name == seg {
				next, found = n.Fields, true
				break
			}
		}
		if !found {
			return nil, false
		}
		cur = next
	}
	return cur, true
}

// requiredDocPathSet is docPathSet restricted to nodes whose every ancestor and
// self are documented required.
func requiredDocPathSet(nodes []paramNode, prefix string) map[string]bool {
	out := map[string]bool{}
	var walk func(nodes []paramNode, prefix string)
	walk = func(nodes []paramNode, prefix string) {
		for _, n := range nodes {
			if !n.Required {
				continue
			}
			p := prefix + n.Name
			out[p] = true
			if len(n.Fields) > 0 {
				walk(n.Fields, p+".")
			}
		}
	}
	walk(nodes, prefix)
	return out
}

// jsonPathSet flattens decoded JSON into dotted paths, dropping array indexes
// so every element contributes to the same paths.
func jsonPathSet(v any) map[string]bool {
	out := map[string]bool{}
	var walk func(v any, prefix string)
	walk = func(v any, prefix string) {
		switch t := v.(type) {
		case map[string]any:
			for k, child := range t {
				p := k
				if prefix != "" {
					p = prefix + "." + k
				}
				out[p] = true
				walk(child, p)
			}
		case []any:
			for _, child := range t {
				walk(child, prefix)
			}
		}
	}
	walk(v, "")
	return out
}

// scalarTypes are the types reflection stops at: they decode one JSON value.
var scalarTypes = []reflect.Type{
	reflect.TypeOf(time.Time{}),
	reflect.TypeOf(json.RawMessage{}),
	reflect.TypeOf(json.Number("")),
	reflect.TypeOf(ali.ID(0)),
	reflect.TypeOf(ali.IDs{}),
	reflect.TypeOf(ali.StringID(0)),
	reflect.TypeOf(ali.FlexBool(false)),
	reflect.TypeOf(ali.FlexInt64(0)),
	reflect.TypeOf(ali.FlexString("")),
	reflect.TypeOf(ali.Fen(0)),
	reflect.TypeOf(ali.YuanFen(0)),
	reflect.TypeOf(ali.PriceString("")),
	reflect.TypeOf(ali.Decimal("")),
	reflect.TypeOf(ali.Timestamp("")),
}

func isScalarType(t reflect.Type) bool {
	for _, s := range scalarTypes {
		if t == s {
			return true
		}
	}
	return false
}

// goPathSet reflects a struct type into dotted json-tag paths, recursing into
// struct fields, pointers, slices, arrays and map values. Fields tagged "-" or
// unexported are skipped; embedded structs are flattened like encoding/json
// does. A type may appear twice on the recursion stack (root and one nested
// level) so self-referential models such as Category.children still describe
// their child members without recursing forever.
func goPathSet(root reflect.Type) map[string]bool {
	out := map[string]bool{}
	var walk func(t reflect.Type, prefix string, stack []reflect.Type)
	walk = func(t reflect.Type, prefix string, stack []reflect.Type) {
		for !isScalarType(t) && (t.Kind() == reflect.Pointer || t.Kind() == reflect.Slice || t.Kind() == reflect.Array || t.Kind() == reflect.Map) {
			t = t.Elem()
		}
		if isScalarType(t) || t.Kind() != reflect.Struct {
			return
		}
		seen := 0
		for _, s := range stack {
			if s == t {
				seen++
			}
		}
		if seen >= 2 {
			return
		}
		stack = append(stack, t)
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			if f.PkgPath != "" && !f.Anonymous {
				continue
			}
			name, _, _ := strings.Cut(f.Tag.Get("json"), ",")
			if name == "-" {
				continue
			}
			if name == "" {
				if f.Anonymous {
					walk(f.Type, prefix, stack)
					continue
				}
				name = f.Name
			}
			p := prefix + name
			out[p] = true
			walk(f.Type, p+".", stack)
		}
	}
	walk(root, "", nil)
	return out
}
