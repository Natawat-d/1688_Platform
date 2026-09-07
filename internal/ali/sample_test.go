package ali_test

// Decodes the documented response samples through the same decoders the
// client uses, and pins concrete values that the documentation actually
// contains. Samples are HTML-escaped in the archive and some are prose rather
// than JSON; those are logged and skipped, never failed.

import (
	"encoding/json"
	"errors"
	"strings"
	"testing"
	"time"

	"marketplace/internal/ali"
)

// decoders maps each API onto the decode path its Client method takes.
var decoders = map[string]func(body []byte) (any, error){
	ali.AccountBasic.DocID():     func(b []byte) (any, error) { return ali.DecodeB[ali.SimpleAccountInfo](ali.AccountBasic.Key(), b) },
	ali.KeywordQuery.DocID():     func(b []byte) (any, error) { return ali.DecodeA[ali.OfferPage](ali.KeywordQuery.Key(), b) },
	ali.KeywordSN.DocID():        func(b []byte) (any, error) { return ali.DecodeA[[]ali.SearchNav](ali.KeywordSN.Key(), b) },
	ali.ProductDetail.DocID():    func(b []byte) (any, error) { return ali.DecodeA[ali.OfferDetail](ali.ProductDetail.Key(), b) },
	ali.CategoryByID.DocID():     func(b []byte) (any, error) { return ali.DecodeA[ali.Category](ali.CategoryByID.Key(), b) },
	ali.FreightEstimate.DocID():  func(b []byte) (any, error) { return ali.DecodeA[ali.ProductFreight](ali.FreightEstimate.Key(), b) },
	ali.OrderPreview.DocID():     flat[ali.Preview](ali.OrderPreview),
	ali.CreateCrossOrder.DocID(): flat[ali.CreateResult](ali.CreateCrossOrder),
	ali.AlipayURLGet.DocID():     flat[ali.PayURL](ali.AlipayURLGet),
	ali.OrderBuyerView.DocID():   func(b []byte) (any, error) { return ali.DecodeB[ali.TradeInfo](ali.OrderBuyerView.Key(), b) },
	ali.BuyerOrderList.DocID():   flat[ali.OrderListResult](ali.BuyerOrderList),
	ali.TradeCancel.DocID():      flat[ali.CancelResult](ali.TradeCancel),
	ali.LogisticsTrace.DocID():   flat[ali.LogisticsTraceResult](ali.LogisticsTrace),
	ali.PushCursorList.DocID(): flat[struct {
		PushMessageList []ali.PushMessage `json:"pushMessageList"`
	}](ali.PushCursorList),
	ali.PushQueryList.DocID(): flat[struct {
		PushMessagePage ali.PushMessagePage `json:"pushMessagePage"`
	}](ali.PushQueryList),
	ali.PushConfirm.DocID(): flat[struct {
		IsSuccess ali.FlexBool `json:"isSuccess"`
	}](ali.PushConfirm),
}

func flat[T any](api ali.API) func([]byte) (any, error) {
	return func(b []byte) (any, error) {
		var out T
		err := ali.DecodeFlat(api.Key(), b, &out)
		return out, err
	}
}

// decodeSample unescapes and decodes one documented sample through the
// API's decoder into out. It reports false, after logging, when the sample is
// absent or not JSON.
func decodeSample(t *testing.T, d *rawDoc, api ali.API, nameSubstr string, out any) bool {
	t.Helper()
	s, ok := sampleByName(d, nameSubstr)
	if !ok {
		t.Logf("%s: no sample named *%s*; skipping its assertions", api.DocID(), nameSubstr)
		return false
	}
	body := []byte(unescapeSample(s.Sample))
	if !json.Valid(body) {
		t.Logf("%s: sample %q is not JSON; skipping its assertions", api.DocID(), s.Name)
		return false
	}
	if err := ali.DecodeFlat(api.Key(), body, out); err != nil {
		t.Fatalf("%s: sample %q: %v", api.DocID(), s.Name, err)
	}
	return true
}

// TestDocSamplesDecode runs every documented response sample through the
// decoder its Client method uses, then pins concrete documented values.
func TestDocSamplesDecode(t *testing.T) {
	dir := docsDir(t)
	docs := map[string]*rawDoc{}
	for _, api := range ali.All() {
		docs[api.DocID()] = loadRawDoc(t, dir, api)
	}

	// Every response sample of every API must at least decode. A documented
	// placeholder status ("code": "状态码") surfaces as an APIError, which is the
	// decoder doing its job on a sample that never claimed to be a success.
	for _, api := range ali.All() {
		d := docs[api.DocID()]
		dec, ok := decoders[api.DocID()]
		if !ok {
			t.Errorf("%s: no decoder registered in sample_test.go", api.DocID())
			continue
		}
		n := 0
		for _, s := range d.Samples {
			if !isResponseSample(s.Name) {
				continue
			}
			body := []byte(unescapeSample(s.Sample))
			if !json.Valid(body) {
				t.Logf("%s: response sample %q is not JSON (prose or key:value lines); skipped", api.DocID(), s.Name)
				continue
			}
			n++
			_, err := dec(body)
			var ae *ali.APIError
			switch {
			case err == nil:
			case errors.As(err, &ae):
				t.Logf("%s: sample %q decoded but its status fields read as a refusal (%s: %s); shape is fine", api.DocID(), s.Name, ae.Code, ae.Message)
			default:
				t.Errorf("%s: sample %q does not decode through our decoder: %v\n    body: %s", api.DocID(), s.Name, err, truncateBody(body))
			}
		}
		if n == 0 {
			t.Logf("%s: no JSON response sample in the archive", api.DocID())
		}
	}

	t.Run("preview", func(t *testing.T) {
		var p ali.Preview
		if !decodeSample(t, docs[ali.OrderPreview.DocID()], ali.OrderPreview, "出参", &p) {
			return
		}
		if len(p.Results) != 1 {
			t.Fatalf("orderPreviewResuslt has %d entries, want 1", len(p.Results))
		}
		r := p.Results[0]
		if r.SumPayment != 4400 {
			t.Errorf("sumPayment = %d fen, want 4400", r.SumPayment)
		}
		if r.SumCarriage != 800 {
			t.Errorf("sumCarriage = %d fen, want 800", r.SumCarriage)
		}
		if r.SumPaymentNoCarriage != 3600 {
			t.Errorf("sumPaymentNoCarriage = %d fen, want 3600", r.SumPaymentNoCarriage)
		}
		if len(r.TradeModeNameList) != 1 || r.TradeModeNameList[0] != "nzassure" {
			t.Errorf("tradeModeNameList = %v, want [nzassure]", r.TradeModeNameList)
		}
		if len(r.CargoList) != 1 {
			t.Fatalf("cargoList has %d entries, want 1", len(r.CargoList))
		}
		if r.CargoList[0].SkuID != 3332085412530 {
			t.Errorf("cargoList[0].skuId = %d, want 3332085412530", r.CargoList[0].SkuID)
		}
		if r.CargoList[0].OfferID != 548818919805 {
			t.Errorf("cargoList[0].offerId = %d, want 548818919805", r.CargoList[0].OfferID)
		}
		// amount 36 and finalUnitPrice 18 are YUAN in the cargo model.
		if r.CargoList[0].Amount.Fen() != 3600 || r.CargoList[0].FinalUnitPrice.Fen() != 1800 {
			t.Errorf("cargoList[0] amount/finalUnitPrice = %d/%d fen, want 3600/1800", r.CargoList[0].Amount.Fen(), r.CargoList[0].FinalUnitPrice.Fen())
		}
		if r.FlowFlag != ali.FlowGeneral || !r.Status.Bool() || !p.Success.Bool() {
			t.Errorf("flowFlag=%q status=%v success=%v, want general/true/true", r.FlowFlag, r.Status, p.Success)
		}
	})

	t.Run("createCrossOrder", func(t *testing.T) {
		d := docs[ali.CreateCrossOrder.DocID()]

		// Wrapped, multi-order: everything is in result.orderList.
		var multi ali.CreateResult
		if decodeSample(t, d, ali.CreateCrossOrder, "多个订单", &multi) {
			if len(multi.OrderList) != 2 {
				t.Fatalf("multi-order sample: orderList has %d entries, want 2", len(multi.OrderList))
			}
			if multi.OrderList[0].OrderID != 105581756010628640 {
				t.Errorf("orderList[0].orderId = %d, want 105581756010628640", multi.OrderList[0].OrderID)
			}
			if multi.OrderList[0].OrderAmmount != 4848000 {
				t.Errorf("orderList[0].orderAmmount = %d fen, want 4848000", multi.OrderList[0].OrderAmmount)
			}
			if multi.OrderList[0].PostFee != 48000 {
				t.Errorf("orderList[0].postFee = %d fen, want 48000", multi.OrderList[0].PostFee)
			}
			if multi.OrderList[1].OrderID != 105637283010628640 {
				t.Errorf("orderList[1].orderId = %d, want 105637283010628640", multi.OrderList[1].OrderID)
			}
			if got := multi.OrderIDs(); len(got) != 2 {
				t.Errorf("OrderIDs() = %v, want two ids", got)
			}
			if multi.OrderID != 0 {
				t.Errorf("multi-order sample carries a top-level orderId %d; documented empty", multi.OrderID)
			}
		}

		// Flat, single-order: no result wrapper at all.
		var single ali.CreateResult
		if decodeSample(t, d, ali.CreateCrossOrder, "支付宝方式下单", &single) {
			if single.OrderID != 87407346014789305 {
				t.Errorf("flat sample orderId = %d, want 87407346014789305", single.OrderID)
			}
			if single.TotalSuccessAmount != 156800 {
				t.Errorf("flat sample totalSuccessAmount = %d fen, want 156800", single.TotalSuccessAmount)
			}
			if got := single.OrderIDs(); len(got) != 1 || got[0] != 87407346014789305 {
				t.Errorf("OrderIDs() = %v, want [87407346014789305]", got)
			}
		}

		var period ali.CreateResult
		if decodeSample(t, d, ali.CreateCrossOrder, "账期支付", &period) {
			if period.AccountPeriod == nil || period.AccountPeriod.TapType != 5 || period.AccountPeriod.TapDate != 360 {
				t.Errorf("account-period sample accountPeriod = %+v, want tapType 5 tapDate 360", period.AccountPeriod)
			}
		}
	})

	t.Run("logisticsTrace", func(t *testing.T) {
		var r ali.LogisticsTraceResult
		if !decodeSample(t, docs[ali.LogisticsTrace.DocID()], ali.LogisticsTrace, "出参", &r) {
			return
		}
		if len(r.LogisticsTrace) != 1 {
			t.Fatalf("logisticsTrace has %d entries, want 1", len(r.LogisticsTrace))
		}
		tr := r.LogisticsTrace[0]
		if tr.OrderID != 188983797838441800 {
			t.Errorf("orderId = %d, want 188983797838441800 (must survive as int64, not float64)", tr.OrderID)
		}
		if tr.LogisticsID != "LP00106397027178" || tr.LogisticsBillNo != "3832890717253" {
			t.Errorf("logisticsId/logisticsBillNo = %q/%q", tr.LogisticsID, tr.LogisticsBillNo)
		}
		if len(tr.LogisticsSteps) != 9 {
			t.Fatalf("logisticsSteps has %d entries, want 9", len(tr.LogisticsSteps))
		}
		got, err := ali.ParseTime(tr.LogisticsSteps[0].AcceptTime.String())
		if err != nil {
			t.Fatalf("ParseTime(%q): %v", tr.LogisticsSteps[0].AcceptTime, err)
		}
		want := time.Date(2018, 7, 24, 21, 55, 33, 0, ali.CST)
		if !got.Equal(want) {
			t.Errorf("logisticsSteps[0].acceptTime = %s, want %s", got, want)
		}
		if _, off := got.Zone(); off != 8*3600 {
			t.Errorf("acceptTime zone offset = %d, want +08:00: the naive layout means Asia/Shanghai", off)
		}
		if got.Hour() != 21 {
			t.Errorf("acceptTime hour = %d in its own zone, want 21 (an eight-hour shift means it was read as UTC)", got.Hour())
		}
	})

	t.Run("alipayURL", func(t *testing.T) {
		d := docs[ali.AlipayURLGet.DocID()]
		var multi ali.PayURL
		if decodeSample(t, d, ali.AlipayURLGet, "多个订单", &multi) {
			const want = "orderId=154051432607498520;151923545459498520"
			if !strings.Contains(multi.PayURL, want) {
				t.Errorf("payUrl = %q, want it to contain %q", multi.PayURL, want)
			}
			if !multi.Success.Bool() {
				t.Errorf("success = %v, want true", multi.Success)
			}
		}
		var single ali.PayURL
		if decodeSample(t, d, ali.AlipayURLGet, "单个订单", &single) {
			if !strings.HasSuffix(single.PayURL, "orderId=151923545459498520") {
				t.Errorf("single payUrl = %q", single.PayURL)
			}
		}
	})

	t.Run("accountBasic", func(t *testing.T) {
		d := docs[ali.AccountBasic.DocID()]
		s, ok := sampleByName(d, "出参")
		if !ok {
			t.Logf("%s: no 出参 sample; skipping", ali.AccountBasic.DocID())
			return
		}
		body := []byte(unescapeSample(s.Sample))
		acct, err := ali.DecodeB[ali.SimpleAccountInfo](ali.AccountBasic.Key(), body)
		if err != nil {
			t.Fatalf("DecodeB: %v", err)
		}
		if acct.LoginID != "alitestforisv01" {
			t.Errorf("loginId = %q, want alitestforisv01", acct.LoginID)
		}
		// The archived sample's memberId is b2b-1623492085 (its userId is
		// 1623492085); b2b-2248544159 is the buyerID of the order-detail sample
		// and the stub's buyer, not this account.
		if acct.MemberID != "b2b-1623492085" {
			t.Errorf("memberId = %q, want b2b-1623492085", acct.MemberID)
		}
		if acct.UserID != 1623492085 {
			t.Errorf("userId = %d, want 1623492085", acct.UserID)
		}
		if !acct.EnterpriseAccount.Bool() || acct.PersonAccount.Bool() {
			t.Errorf("enterpriseAccount/personAccount = %v/%v, want true/false", acct.EnterpriseAccount, acct.PersonAccount)
		}
		created, err := acct.CreateDate.Time()
		if err != nil || !created.Equal(time.Date(2013, 3, 12, 16, 8, 28, 0, ali.CST)) {
			t.Errorf("createDate = %q -> %s (%v), want 2013-03-12 16:08:28 +08:00", acct.CreateDate, created, err)
		}
	})

	t.Run("orderBuyerView", func(t *testing.T) {
		d := docs[ali.OrderBuyerView.DocID()]
		s, ok := sampleByName(d, "出参")
		if !ok {
			t.Logf("%s: no 出参 sample; skipping", ali.OrderBuyerView.DocID())
			return
		}
		info, err := ali.DecodeB[ali.TradeInfo](ali.OrderBuyerView.Key(), []byte(unescapeSample(s.Sample)))
		if err != nil {
			t.Fatalf("DecodeB: %v", err)
		}
		if info.BaseInfo.BuyerID != "b2b-2248544159" {
			t.Errorf("baseInfo.buyerID = %q, want b2b-2248544159", info.BaseInfo.BuyerID)
		}
		// The archived sample was rendered by a tool that rounds large numbers:
		// baseInfo.id reads 58218860983545944 beside idOfStr "58218860983545941",
		// and subItemID 128403042259997710 beside subItemIDString
		// "128403042259997715". That is the whole reason the string spellings
		// exist, and this pins that we keep them as text.
		if info.BaseInfo.IDOfStr.String() != "58218860983545941" {
			t.Errorf("baseInfo.idOfStr = %q, want 58218860983545941", info.BaseInfo.IDOfStr)
		}
		if info.BaseInfo.ID != 58218860983545944 {
			t.Errorf("baseInfo.id = %d, want the sample's 58218860983545944", info.BaseInfo.ID)
		}
		if len(info.ProductItems) == 0 {
			t.Fatal("productItems is empty")
		}
		if got := info.ProductItems[0].SubItemIDString.String(); got != "128403042259997715" {
			t.Errorf("productItems[0].subItemIDString = %q, want 128403042259997715", got)
		}
		if _, err := info.BaseInfo.CreateTime.Time(); err != nil {
			t.Errorf("baseInfo.createTime %q: %v", info.BaseInfo.CreateTime, err)
		}
		if len(info.Extra) == 0 {
			t.Errorf("TradeInfo.Extra is empty: the sample's unmodelled members (orderBizInfo, orderRateInfo, ...) should land there")
		}
	})

	t.Run("buyerOrderList", func(t *testing.T) {
		var list ali.OrderListResult
		if !decodeSample(t, docs[ali.BuyerOrderList.DocID()], ali.BuyerOrderList, "返回参数", &list) {
			return
		}
		if len(list.Orders) == 0 {
			t.Fatal("result is empty")
		}
		o := list.Orders[0]
		// The list sample sends idOfStr as a BARE NUMBER, unlike the order
		// detail sample, which quotes it. Both must decode to the same text.
		if o.BaseInfo.IDOfStr.String() != "196965465451498520" {
			t.Errorf("result[0].baseInfo.idOfStr = %q, want 196965465451498520 (sent unquoted in the sample)", o.BaseInfo.IDOfStr)
		}
		if o.BaseInfo.ID != 196965465451498520 {
			t.Errorf("result[0].baseInfo.id = %d, want 196965465451498520", o.BaseInfo.ID)
		}
		if o.BaseInfo.BuyerID != "b2b-1623492085" {
			t.Errorf("result[0].baseInfo.buyerID = %q, want b2b-1623492085", o.BaseInfo.BuyerID)
		}
		if list.ErrorCode != "" {
			t.Errorf("errorCode = %q on a successful sample", list.ErrorCode)
		}
	})

	t.Run("pushReplay", func(t *testing.T) {
		var cur struct {
			PushMessageList []ali.PushMessage `json:"pushMessageList"`
		}
		if decodeSample(t, docs[ali.PushCursorList.DocID()], ali.PushCursorList, "返回结果", &cur) {
			if len(cur.PushMessageList) != 1 || cur.PushMessageList[0].MsgID != 123456 {
				t.Errorf("cursor sample: %+v", cur.PushMessageList)
			} else if m := cur.PushMessageList[0]; m.Topic() != "CAIGOU_MSG_BUYER_PUBLISH_BUYOFFER" || ali.TopicType(m.TopicGroup, m.TopicName) != m.Type {
				t.Errorf("cursor sample: type=%q topicGroup=%q topicName=%q disagree", m.Type, m.TopicGroup, m.TopicName)
			}
		}
		var q struct {
			PushMessagePage ali.PushMessagePage `json:"pushMessagePage"`
		}
		if decodeSample(t, docs[ali.PushQueryList.DocID()], ali.PushQueryList, "返回结果", &q) {
			if q.PushMessagePage.TotalCount != 1 || len(q.PushMessagePage.Datas) != 1 || q.PushMessagePage.Datas[0].MsgID != 68891027 {
				t.Errorf("query sample: %+v", q.PushMessagePage)
			}
		}
	})
}

func truncateBody(b []byte) string {
	if len(b) > 300 {
		return string(b[:300]) + "…"
	}
	return string(b)
}
