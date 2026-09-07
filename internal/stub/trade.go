package stub

import (
	"context"
	"encoding/json"
	"net/url"
	"strconv"
	"strings"

	"marketplace/internal/ali"
)

// The trade family. These are the endpoints with real rules behind them: the
// preview refuses nine documented ways, the create refuses more, and the cancel
// has a ten-second window. Getting refused is the interesting case, so all of it
// is implemented rather than waved through.

// cargo is one line of cargoParamList as submitted.
type cargo struct {
	OfferID  int64
	SpecID   string
	Quantity int64
}

func parseCargo(form url.Values) ([]cargo, bool) {
	raw, ok := jsonArray(form, "cargoParamList")
	if !ok {
		return nil, false
	}
	out := make([]cargo, 0, len(raw))
	for _, m := range raw {
		q := mapInt64(m, "quantity")
		if q <= 0 {
			q = 1
		}
		out = append(out, cargo{
			OfferID:  mapInt64(m, "offerId"),
			SpecID:   mapStr(m, "specId"),
			Quantity: q,
		})
	}
	return out, true
}

func parseAddress(form url.Values) Address {
	m, ok := jsonObject(form, "addressParam")
	if !ok {
		return Address{}
	}
	return Address{
		AddressID:    mapInt64(m, "addressId"),
		FullName:     mapStr(m, "fullName"),
		Mobile:       mapStr(m, "mobile"),
		Phone:        mapStr(m, "phone"),
		PostCode:     mapStr(m, "postCode"),
		ProvinceText: mapStr(m, "provinceText"),
		CityText:     mapStr(m, "cityText"),
		AreaText:     mapStr(m, "areaText"),
		TownText:     mapStr(m, "townText"),
		Address:      mapStr(m, "address"),
		DistrictCode: firstNonEmpty(mapStr(m, "districtCode"), mapStr(m, "addressCode")),
	}
}

func firstNonEmpty(vals ...string) string {
	for _, v := range vals {
		if v != "" {
			return v
		}
	}
	return ""
}

// tradeModes is the list of trade types an offer supports. Every 31st offer
// supports none, which is the documented "this transaction cannot be placed via
// the API and must be placed on the 1688 page" case.
func tradeModes(offerID int64) []string {
	if offerID%31 == 0 {
		return nil
	}
	return []string{"assureTrade", "alipay"}
}

// previewError is a documented refusal from alibaba.createOrder.preview.
func previewError(code, msg string) map[string]any {
	// Field names verbatim: errorMsg here, NOT errorMessage, and the result
	// field is orderPreviewResuslt.
	return map[string]any{
		"success":                            false,
		"errorCode":                          code,
		"errorMsg":                           msg,
		"orderPreviewResuslt":                []any{},
		"postFeeByDescOfferList":             []any{},
		"consignOfferList":                   []any{},
		"unsupportedCrossBorderPayOfferList": []any{},
		"extPairList":                        []any{},
	}
}

// handlePreview serves com.alibaba.trade:alibaba.createOrder.preview:1.
func (s *Server) handlePreview(form url.Values) any {
	cargos, ok := parseCargo(form)
	if !ok || len(cargos) == 0 {
		return previewError("400", "cargoParamList is required")
	}
	flow := form.Get("flow")

	if err := s.checkCargo(cargos, flow); err != nil {
		return previewError(err.code, err.msg)
	}

	first := s.offer(cargos[0].OfferID)
	lines, sumNoCarriage := s.priceLines(cargos)
	carriage := s.postFee(cargos)
	discount := s.discount(first.ID, sumNoCarriage)

	cargoList := make([]any, 0, len(lines))
	for _, l := range lines {
		cargoList = append(cargoList, map[string]any{
			"amount":                   yuan(l.Amount),
			"message":                  nil,
			"finalUnitPrice":           yuan(l.UnitPrice),
			"specId":                   l.SpecID,
			"skuId":                    l.SkuID,
			"resultCode":               "200",
			"offerId":                  l.OfferID,
			"openOfferId":              "",
			"cargoPromotionList":       []any{},
			"extPairList":              []any{},
			"totalFundUsageAmount":     0,
			"totalPostFundUsageAmount": 0,
			"tradeServiceList":         []any{},
		})
	}

	modes := tradeModes(first.ID)
	modeList := make([]any, 0, len(modes))
	for _, m := range modes {
		modeList = append(modeList, map[string]any{
			"tradeType":   m,
			"name":        tradeModeName(m),
			"description": "买家下单5天内全款支付。自卖家发货起，买家需在10天内确认收货，确认收货后打款给卖家。",
			"opSupport":   true,
		})
	}
	nameList := make([]any, 0, len(modes))
	for _, m := range modes {
		nameList = append(nameList, m)
	}

	result := map[string]any{
		"discountFee":              int64(discount),
		"tradeModeNameList":        nameList,
		"status":                   true,
		"taoSampleSinglePromotion": false,
		"sumPayment":               int64(sumNoCarriage + carriage - discount),
		"message":                  "",
		"sumCarriage":              int64(carriage),
		"resultCode":               "200",
		"sumPaymentNoCarriage":     int64(sumNoCarriage),
		"additionalFee":            int64(0),
		"flowFlag":                 firstNonEmpty(flow, "general"),
		"cargoList":                cargoList,
		"shopPromotionList":        []any{},
		"tradeModelList":           modeList,
		"payChannelInfos": []any{
			map[string]any{"name": "alipay", "amountLimit": 100000000},
			map[string]any{"name": "kjb", "amountLimit": 30000000},
		},
		"tradeServiceList":       []any{},
		"extPairList":            []any{},
		"orderGroup":             strconv.FormatInt(first.ID%100000, 10),
		"canUseOfficialSolution": first.ID%3 == 0,
		"officialSolutionModelList": []any{
			map[string]any{"solutionCode": "471", "solutionName": "特惠当日上门", "totalCost": 770},
		},
		"totalFundUsageAmount":     0,
		"totalPostFundUsageAmount": 0,
	}

	consign := make([]any, 0, len(cargos))
	for _, c := range cargos {
		consign = append(consign, c.OfferID)
	}

	return map[string]any{
		"orderPreviewResuslt":                []any{result},
		"success":                            true,
		"errorCode":                          "",
		"errorMsg":                           "",
		"postFeeByDescOfferList":             []any{},
		"consignOfferList":                   consign,
		"unsupportedCrossBorderPayOfferList": []any{},
		"extPairList":                        []any{},
	}
}

func tradeModeName(t string) string {
	switch t {
	case "assureTrade":
		return "担保交易"
	case "alipay":
		return "支付宝担保交易"
	}
	return t
}

// bizError is one documented business refusal.
type bizError struct{ code, msg string }

// checkCargo runs the documented preview validations, in the documented order.
// Every one of these is reachable from the generated catalogue: see gen.Offer.
func (s *Server) checkCargo(cargos []cargo, flow string) *bizError {
	var seller string
	perOffer := map[int64]int64{}
	for _, c := range cargos {
		perOffer[c.OfferID] += c.Quantity
	}

	for _, c := range cargos {
		id := strconv.FormatInt(c.OfferID, 10)
		o := s.offer(c.OfferID)

		if !o.Orderable() {
			return &bizError{ali.ErrNoOnlineTrade, "商品[" + id + "]不支持在线交易，无法下单。"}
		}
		if c.SpecID == "" {
			return &bizError{ali.ErrMultiSeller, "商品[" + id + "]不属于同一个卖家，或者没有指定specId。"}
		}
		if seller == "" {
			seller = o.Seller.MemberID
		} else if seller != o.Seller.MemberID {
			return &bizError{ali.ErrMultiSeller, "商品[" + id + "]不属于同一个卖家，或者没有指定specId。"}
		}

		sku, ok := o.SKUBySpec(c.SpecID)
		if !ok {
			return &bizError{ali.ErrSpecNotInOffer, "商品[" + id + "]不属于同一个卖家，或者规格[" + c.SpecID + "]不属于商品[" + id + "]。"}
		}
		if flow == "saleproxy" {
			return &bizError{ali.ErrNoConsignRel, "与供应商的代销关系不存在，无法通过saleproxy渠道下单。"}
		}
		if sku.Stock < c.Quantity {
			return &bizError{ali.ErrNoStock, "商品[" + id + "_" + c.SpecID + "]库存不足，请确认库存后再下单。"}
		}
		if perOffer[c.OfferID] < int64(o.MOQ) {
			return &bizError{ali.ErrBelowMOQ, "商品[" + id + "]的采购量不满足起订量限制。"}
		}
		if o.BatchNumber > 0 && perOffer[c.OfferID]%int64(o.BatchNumber) != 0 {
			return &bizError{ali.ErrMixedBatch, "商品[" + id + "]的采购量或者金额不满足混批限制。"}
		}
		if sku.Price == 0 {
			return &bizError{ali.ErrZeroPrice, "商品规格[" + id + "_" + c.SpecID + "]的价格为0，无法下单，请检查后重新提交。"}
		}
		// boutiquepifa is documented as the curated-supply WHOLESALE flow, used
		// when the purchase quantity is greater than two.
		if flow == "boutiquepifa" && perOffer[c.OfferID] < 2 {
			return &bizError{ali.ErrBelowWholesale, "商品[" + id + "]的采购量不满足批发起批量限制。"}
		}
	}
	return nil
}

// priceLines turns cargo into priced order lines.
func (s *Server) priceLines(cargos []cargo) ([]Line, ali.Fen) {
	perOffer := map[int64]int64{}
	for _, c := range cargos {
		perOffer[c.OfferID] += c.Quantity
	}

	var total ali.Fen
	lines := make([]Line, 0, len(cargos))
	for _, c := range cargos {
		o := s.offer(c.OfferID)
		sku, _ := o.SKUBySpec(c.SpecID)

		unit := sku.Price
		if o.QuoteType != 1 {
			// Quoted by product quantity: the wholesale tier decides the price.
			unit = o.TierPrice(perOffer[c.OfferID])
		}
		amount := unit * ali.Fen(c.Quantity)
		total += amount

		attrs := make([]NameValue, 0, len(sku.Attrs))
		for _, a := range sku.Attrs {
			attrs = append(attrs, NameValue{Name: a.NameZH, Value: a.ValueZH})
		}
		lines = append(lines, Line{
			OfferID:     o.ID,
			SkuID:       sku.SkuID,
			SpecID:      c.SpecID,
			Quantity:    c.Quantity,
			UnitPrice:   unit,
			Amount:      amount,
			Name:        o.Subject,
			NameTrans:   o.SubjectTrans,
			Unit:        o.UnitZH,
			CargoNumber: sku.CargoNumber,
			Images:      s.imageURLs(o.ImagePaths),
			SkuAttrs:    attrs,
		})
	}
	return lines, total
}

// postFee is the shipping charge, computed the same way freight.estimate does.
func (s *Server) postFee(cargos []cargo) ali.Fen {
	var grams int64
	var first ali.Fen
	for _, c := range cargos {
		o := s.offer(c.OfferID)
		if first == 0 {
			first = ali.Fen(500 + o.ID%800)
		}
		grams += o.WeightG * c.Quantity
	}
	kilos := (grams + 999) / 1000
	if kilos < 1 {
		kilos = 1
	}
	next := ali.Fen(200)
	return first + next*ali.Fen(kilos-1)
}

// discount is a small shop-level reduction on some offers, so the discountFee
// field is not always zero.
func (s *Server) discount(offerID int64, sum ali.Fen) ali.Fen {
	if offerID%4 != 0 {
		return 0
	}
	return ali.Fen(ali.DivRoundHalfUp(int64(sum), 20)) // 5%
}

// handleCreateCrossOrder serves com.alibaba.trade:alibaba.trade.createCrossOrder:1.
func (s *Server) handleCreateCrossOrder(ctx context.Context, form url.Values) any {
	cargos, ok := parseCargo(form)
	if !ok || len(cargos) == 0 {
		return createError("400", "cargoParamList is required")
	}
	// An order may carry at most fifty SKUs. The limit is documented; the code
	// for breaching it is not, so 400 it is.
	if len(cargos) > 50 {
		return createError("400", "an order may contain at most 50 SKUs, got "+strconv.Itoa(len(cargos)))
	}

	flow := form.Get("flow")
	if err := s.checkCargo(cargos, flow); err != nil {
		return createError(err.code, err.msg)
	}

	first := s.offer(cargos[0].OfferID)
	tradeType := form.Get("tradeType")
	modes := tradeModes(first.ID)
	if tradeType != "" && !contains(modes, tradeType) {
		// Documented verbatim, brackets and all.
		return createError("400", "not support tradeType:【"+tradeType+"】")
	}
	if tradeType == "" {
		if len(modes) == 0 {
			return createError("400", "not support tradeType:【】")
		}
		tradeType = modes[0]
	}

	lines, sumNoCarriage := s.priceLines(cargos)
	carriage := s.postFee(cargos)
	discount := s.discount(first.ID, sumNoCarriage)

	// Every fifth successful create splits into two orders, exercising the
	// documented "several orders are created at once" response shape.
	split := (s.store.Creates()+1)%5 == 0

	orders, events, replayed, err := s.store.Create(CreateParams{
		OutOrderID: form.Get("outOrderId"),
		Flow:       firstNonEmpty(flow, "general"),
		TradeType:  tradeType,
		Message:    form.Get("message"),
		Address:    parseAddress(form),
		Lines:      lines,
		PostFee:    carriage,
		Discount:   discount,
		Seller:     first.Seller,
		Buyer:      buyer,
		Split:      split,
	})
	if err != nil {
		return createError("500", err.Error())
	}
	if !replayed {
		s.Publish(ctx, events)
	}

	if len(orders) > 1 {
		list := make([]any, 0, len(orders))
		for _, o := range orders {
			list = append(list, map[string]any{
				"postFee": int64(o.PostFee),
				// orderAmmount: 1688's spelling. Two m's.
				"orderAmmount": int64(o.Total),
				"message":      nil,
				"resultCode":   "200",
				"success":      true,
				"orderId":      strconv.FormatInt(o.ID, 10),
				"payChannel":   o.PayChannel,
			})
		}
		result := map[string]any{
			"totalSuccessAmount": nil, // documented empty when several orders are created
			"orderId":            "",
			"success":            true,
			"code":               "",
			"message":            "",
			"accountPeriod":      nil,
			"failedOfferList":    []any{},
			"postFee":            nil,
			"orderList":          list,
		}
		if s.cfg.CreateFlat {
			return result
		}
		return map[string]any{"result": result, "success": true, "code": "", "message": ""}
	}

	o := orders[0]
	result := map[string]any{
		"totalSuccessAmount": int64(o.Total),
		"orderId":            strconv.FormatInt(o.ID, 10),
		"success":            true,
		"code":               "",
		"message":            "",
		"accountPeriod":      nil,
		"failedOfferList":    []any{},
		"postFee":            int64(o.PostFee),
		"orderList":          []any{},
	}
	if s.cfg.CreateFlat {
		// STUB_CREATE_FLAT: two of the four documented samples return the
		// creation result unwrapped, with no "result" object at all.
		return map[string]any{
			"totalSuccessAmount": int64(o.Total),
			"orderId":            strconv.FormatInt(o.ID, 10),
			"success":            true,
		}
	}
	return map[string]any{"result": result, "success": true, "code": "", "message": ""}
}

func createError(code, msg string) map[string]any {
	return map[string]any{
		"result":  nil,
		"success": false,
		"code":    code,
		"message": msg,
	}
}

func contains(list []string, v string) bool {
	for _, s := range list {
		if s == v {
			return true
		}
	}
	return false
}

// handleAlipayURL serves com.alibaba.trade:alibaba.alipay.url.get:1.
//
// The real gateway answers with https://trade.1688.com/order/cashier.htm?orderId=a;b
// — semicolon-separated. This one answers with its own cashier at the same
// shape, so the demo can actually click Pay.
func (s *Server) handleAlipayURL(form url.Values) any {
	raw := strings.TrimSpace(form.Get("orderIdList"))
	var ids []ali.FlexInt64
	if err := json.Unmarshal([]byte(raw), &ids); err != nil {
		return map[string]any{
			"success":             false,
			"errorCode":           "400",
			"erroMsg":             "orderIdList is not a JSON array of order ids",
			"payUrl":              "",
			"payFailureOrderList": []any{},
		}
	}
	if len(ids) == 0 {
		return map[string]any{
			"success":             false,
			"errorCode":           "400",
			"erroMsg":             "orderIdList is required",
			"payUrl":              "",
			"payFailureOrderList": []any{},
		}
	}
	// Documented: up to 100 orders per batch, 30 for Kuajingbao.
	if len(ids) > 100 {
		return map[string]any{
			"success":             false,
			"errorCode":           "400",
			"erroMsg":             "orderIdList supports at most 100 orders per batch",
			"payUrl":              "",
			"payFailureOrderList": []any{},
		}
	}

	var ok []string
	failed := []any{}
	for _, id := range ids {
		o, found := s.store.Get(int64(id))
		if !found || o.Status != StatusWaitPay {
			failed = append(failed, int64(id))
			continue
		}
		ok = append(ok, strconv.FormatInt(int64(id), 10))
	}
	if len(ok) == 0 {
		return map[string]any{
			"success":             false,
			"errorCode":           "ORDER_STATUS_ERROR",
			"erroMsg":             "no payable order in orderIdList",
			"payUrl":              "",
			"payFailureOrderList": failed,
		}
	}

	return map[string]any{
		"payUrl":              strings.TrimRight(s.cfg.Base, "/") + "/cashier?orderId=" + strings.Join(ok, ";"),
		"success":             len(failed) == 0,
		"errorCode":           "",
		"erroMsg":             "",
		"payFailureOrderList": failed,
	}
}

// handleTradeCancel serves com.alibaba.trade:alibaba.trade.cancel:1.
// tradeID — capital ID, this endpoint only.
func (s *Server) handleTradeCancel(ctx context.Context, form url.Values) any {
	id, ok := parseInt64(form.Get("tradeID"))
	if !ok {
		return map[string]any{"success": false, "errorCode": ali.ErrOrderNotExist, "errorMessage": "订单不存在"}
	}
	reason := form.Get("cancelReason")

	events, err := s.store.Cancel(id, reason, form.Get("remark"))
	switch err {
	case nil:
		s.Publish(ctx, events)
		return map[string]any{"success": true, "errorCode": "", "errorMessage": ""}
	case errNotExist:
		return map[string]any{"success": false, "errorCode": ali.ErrOrderNotExist, "errorMessage": "订单不存在"}
	case errTooFast:
		return map[string]any{"success": false, "errorCode": ali.ErrCloseTooFast, "errorMessage": "下单不足10秒，暂时无法关闭订单"}
	case errStatus:
		return map[string]any{"success": false, "errorCode": ali.ErrOrderStatus, "errorMessage": "订单状态错误"}
	}
	return map[string]any{"success": false, "errorCode": "500", "errorMessage": err.Error()}
}

// handleOrderBuyerView serves com.alibaba.trade:alibaba.trade.get.buyerView:1.
// Family B with success as a STRING.
func (s *Server) handleOrderBuyerView(form url.Values) any {
	var o *Order
	var ok bool
	if id, isNum := parseInt64(form.Get("orderId")); isNum {
		o, ok = s.store.Get(id)
	}
	if !ok && form.Get("outOrderId") != "" {
		o, ok = s.store.GetByOut(form.Get("outOrderId"))
	}
	if !ok {
		return famBStrErr(ali.ErrOrderNotExist, "订单不存在")
	}

	include := form.Get("includeFields")
	wantLogistics := include == "" || strings.Contains(include, "NativeLogistics")
	return famBStr(s.tradeInfo(o, wantLogistics))
}

// handleBuyerOrderList serves com.alibaba.trade:alibaba.trade.getBuyerOrderList:1.
// Note the envelope: result plus totalRecord, and no success field at all.
func (s *Server) handleBuyerOrderList(form url.Values) any {
	f := Filter{
		Status:      form.Get("orderStatus"),
		OutOrderID:  form.Get("outOrderId"),
		SellerID:    form.Get("sellerMemberId"),
		CreatedGTE:  parseDocTime(form.Get("createStartTime")),
		CreatedLTE:  parseDocTime(form.Get("createEndTime")),
		ModifiedGTE: parseDocTime(form.Get("modifyStartTime")),
		ModifiedLTE: parseDocTime(form.Get("modifyEndTime")),
	}
	if raw := strings.TrimSpace(form.Get("orderIds")); raw != "" {
		var ids []ali.FlexInt64
		if err := json.Unmarshal([]byte(raw), &ids); err == nil {
			for _, v := range ids {
				f.OrderIDs = append(f.OrderIDs, int64(v))
			}
		}
	}

	page := atoiDefault(form.Get("page"), 1)
	size := atoiDefault(form.Get("pageSize"), 20)
	orders, total := s.store.List(f, page, size)

	list := make([]any, 0, len(orders))
	for _, o := range orders {
		list = append(list, s.tradeInfo(o, false))
	}
	return map[string]any{
		"result":       list,
		"totalRecord":  total,
		"errorCode":    "",
		"errorMessage": "",
	}
}

// tradeInfo builds alibaba.openplatform.trade.model.TradeInfo.
//
// Read the amounts carefully: baseInfo.discount and baseInfo.refundPayment are
// documented in CENTS while every sibling amount on the same object is in YUAN.
// That inconsistency is real, and it is reproduced here on purpose.
func (s *Server) tradeInfo(o *Order, withLogistics bool) map[string]any {
	items := make([]any, 0, len(o.Lines))
	for _, l := range o.Lines {
		attrs := make([]any, 0, len(l.SkuAttrs))
		for _, a := range l.SkuAttrs {
			attrs = append(attrs, map[string]any{"name": a.Name, "value": a.Value})
		}
		items = append(items, map[string]any{
			"cargoNumber":        l.CargoNumber,
			"productCargoNumber": l.CargoNumber,
			"description":        "",
			"itemAmount":         yuan(l.Amount),
			"name":               l.Name,
			"price":              yuan(l.UnitPrice),
			"productID":          l.OfferID,
			"productImgUrl":      toAny(l.Images),
			"productSnapshotUrl": s.cfg.Base + "/order/offer_snapshot.htm?order_entry_id=" + strconv.FormatInt(l.SubItemID, 10),
			"quantity":           rawNumber(strconv.FormatInt(l.Quantity, 10)),
			"refund":             yuan(0),
			"skuID":              l.SkuID,
			"sort":               0,
			"status":             l.Status,
			"statusStr":          statusText(l.Status),
			"subItemID":          l.SubItemID,
			"subItemIDString":    strconv.FormatInt(l.SubItemID, 10),
			"type":               "common",
			"unit":               l.Unit,
			"weight":             "",
			"weightUnit":         "kg",
			"guaranteesTerms":    []any{},
			"skuInfos":           attrs,
			"entryDiscount":      0,
			"specId":             l.SpecID,
			"quantityFactor":     rawNumber("1"),
			"refundStatus":       "",
			"closeReason":        o.CloseReason,
			"logisticsStatus":    l.LogisticsStatus,
			"gmtCreate":          ali.FormatCompact(o.CreateTime),
			"gmtModified":        ali.FormatCompact(o.ModifyTime),
			"sharePostage":       yuan(0),
		})
	}

	baseInfo := map[string]any{
		"id":                o.ID,
		"idOfStr":           strconv.FormatInt(o.ID, 10),
		"status":            o.Status,
		"businessType":      "cb",
		"buyerID":           o.BuyerMemberID,
		"sellerID":          o.SellerMemberID,
		"buyerLoginId":      o.BuyerLoginID,
		"sellerLoginId":     o.SellerLoginID,
		"buyerUserId":       o.BuyerUserID,
		"sellerUserId":      o.SellerUserID,
		"buyerAlipayId":     "2088611489970483",
		"sellerAlipayId":    "2088611383470360",
		"alipayTradeId":     o.AlipayTradeID,
		"createTime":        ali.FormatCompact(o.CreateTime),
		"modifyTime":        ali.FormatCompact(o.ModifyTime),
		"totalAmount":       yuan(o.Total),
		"sumProductPayment": yuan(o.SumProduct()),
		"shippingFee":       yuan(o.PostFee),
		"couponFee":         yuan(0),
		"refund":            yuan(0),
		// discount and refundPayment are CENTS. See the doc comment above.
		"discount":           int64(o.Discount),
		"refundPayment":      int64(0),
		"refundStatus":       "",
		"refundStatusForAs":  "",
		"tradeType":          "50060",
		"tradeTypeCode":      o.TradeType,
		"tradeTypeDesc":      tradeModeName(o.TradeType),
		"flowTemplateCode":   o.Flow,
		"buyerFeedback":      o.Message,
		"remark":             o.Message,
		"buyerMemo":          "",
		"buyerRemarkIcon":    "",
		"sellerOrder":        false,
		"overSeaOrder":       true,
		"stepPayAll":         false,
		"sellerCreditLevel":  "L" + strconv.Itoa(1+int(o.SellerUserID%6)),
		"payTimeout":         43200,
		"payTimeoutType":     0,
		"inventoryMode":      "cang",
		"outOrderId":         o.OutOrderID,
		"payChannelList":     []any{"支付宝", "跨境宝"},
		"payChannelCodeList": []any{"alipay", "kjb"},
		"buyerContact": map[string]any{
			"phone":        "86-0571-81895955",
			"fax":          "",
			"email":        "buyer@example.test",
			"imInPlatform": o.BuyerLoginID,
			"name":         o.Address.FullName,
			"mobile":       o.Address.Mobile,
			"companyName":  buyer.Company,
		},
		"sellerContact": map[string]any{
			"phone":         "86-0571-88881888",
			"fax":           "",
			"email":         "seller@example.test",
			"imInPlatform":  o.SellerLoginID,
			"name":          "孟舒",
			"mobile":        "13312919596",
			"companyName":   o.SellerCompany,
			"shopName":      o.SellerCompany,
			"wgSenderName":  "",
			"wgSenderPhone": "",
		},
		"receiverInfo": map[string]any{
			"toFullName":     o.Address.FullName,
			"toDivisionCode": o.Address.DistrictCode,
			"toMobile":       o.Address.Mobile,
			"toPhone":        o.Address.Phone,
			"toPost":         o.Address.PostCode,
			"toTownCode":     o.Address.TownCode,
			"toArea":         strings.TrimSpace(o.Address.ProvinceText + " " + o.Address.CityText + " " + o.Address.AreaText + " " + o.Address.Address),
		},
	}
	if !o.PayTime.IsZero() {
		baseInfo["payTime"] = ali.FormatCompact(o.PayTime)
	}
	if !o.ShipTime.IsZero() {
		baseInfo["allDeliveredTime"] = ali.FormatCompact(o.ShipTime)
	}
	if !o.ReceiveTime.IsZero() {
		baseInfo["receivingTime"] = ali.FormatCompact(o.ReceiveTime)
		baseInfo["confirmedTime"] = ali.FormatCompact(o.ReceiveTime)
	}
	if !o.CompleteTime.IsZero() {
		baseInfo["completeTime"] = ali.FormatCompact(o.CompleteTime)
	}
	if o.Status == StatusCancel {
		baseInfo["closeReason"] = o.CloseReason
		baseInfo["closeOperateType"] = o.CloseOperateType
		baseInfo["closeRemark"] = o.CloseRemark
	}

	info := map[string]any{
		"baseInfo":        baseInfo,
		"productItems":    items,
		"tradeTerms":      s.tradeTerms(o),
		"extAttributes":   []any{},
		"guaranteesTerms": []any{},
		"orderRateInfo": map[string]any{
			"buyerRateStatus":  5,
			"sellerRateStatus": 5,
		},
		"fromEncryptOrder": false,
		// orderBizInfo carries the documented order-type flags. For a plain
		// purchase the three that apply are all false: not a caiyuanbao order,
		// not credit pay, not dropshipping. (The earlier odd/cross/crossBorderPay
		// members were not documented anywhere; the contract test caught them.)
		"orderBizInfo": map[string]any{
			"odsCyd":       false,
			"creditOrder":  false,
			"dropshipping": false,
		},
	}
	if withLogistics {
		info["nativeLogistics"] = s.nativeLogistics(o)
	}
	return info
}

func (s *Server) tradeTerms(o *Order) []any {
	if o.PayTime.IsZero() {
		return []any{map[string]any{
			"payStatus":  "1",
			"payWay":     "1",
			"payWayDesc": "支付宝",
			"phasAmount": yuan(o.Total),
			"phase":      o.ID + 1,
			"cardPay":    false,
			"expressPay": false,
		}}
	}
	return []any{map[string]any{
		"payStatus":  "2",
		"payTime":    ali.FormatCompact(o.PayTime),
		"payWay":     "1",
		"payWayDesc": "支付宝",
		"phasAmount": yuan(o.Total),
		"phase":      o.ID + 1,
		"cardPay":    false,
		"expressPay": true,
	}}
}

func (s *Server) nativeLogistics(o *Order) map[string]any {
	nl := map[string]any{
		"address":        o.Address.Address,
		"area":           o.Address.AreaText,
		"areaCode":       o.Address.DistrictCode,
		"city":           o.Address.CityText,
		"contactPerson":  o.Address.FullName,
		"fax":            "",
		"mobile":         o.Address.Mobile,
		"province":       o.Address.ProvinceText,
		"telephone":      o.Address.Phone,
		"zip":            o.Address.PostCode,
		"town":           o.Address.TownText,
		"townCode":       o.Address.TownCode,
		"logisticsItems": []any{},
	}
	if o.Logistics == nil {
		return nl
	}

	subItems := make([]string, 0, len(o.Lines))
	for _, l := range o.Lines {
		subItems = append(subItems, strconv.FormatInt(l.SubItemID, 10))
	}
	nl["logisticsItems"] = []any{map[string]any{
		"deliveredTime":        ali.FormatCompact(o.ShipTime),
		"logisticsCode":        o.Logistics.LogisticsID,
		"type":                 "1",
		"id":                   o.Logistics.CompanyID,
		"status":               "alreadysend",
		"gmtCreate":            ali.FormatCompact(o.ShipTime),
		"gmtModified":          ali.FormatCompact(o.ModifyTime),
		"carriage":             yuan(o.PostFee),
		"fromProvince":         o.SellerProvince,
		"fromCity":             o.SellerCity,
		"fromArea":             o.SellerDistrict,
		"fromAddress":          o.SellerCompany,
		"fromPhone":            "86-0571-88881888",
		"fromMobile":           "13312919596",
		"fromPost":             "310000",
		"logisticsCompanyId":   o.Logistics.CompanyID,
		"logisticsCompanyNo":   o.Logistics.CPCode,
		"logisticsCompanyName": o.Logistics.CompanyName,
		"logisticsBillNo":      o.Logistics.MailNo,
		"subItemIds":           strings.Join(subItems, ","),
		"toProvince":           o.Address.ProvinceText,
		"toCity":               o.Address.CityText,
		"toArea":               o.Address.AreaText,
		"toAddress":            o.Address.Address,
		"toPhone":              o.Address.Phone,
		"toMobile":             o.Address.Mobile,
		"toPost":               o.Address.PostCode,
	}}
	return nl
}

func statusText(status string) string {
	switch status {
	case StatusWaitPay:
		return "等待买家付款"
	case StatusWaitSend:
		return "等待卖家发货"
	case StatusWaitReceive:
		return "等待买家收货"
	case StatusConfirmGoods:
		return "已收货"
	case StatusSuccess:
		return "交易成功"
	case StatusCancel:
		return "交易取消"
	}
	return status
}
