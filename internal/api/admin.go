package api

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"marketplace/internal/ali"
	"marketplace/internal/catalog"
	"marketplace/internal/order"
	"marketplace/internal/pricing"
	"marketplace/internal/store"
)

// ----------------------------------------------------------------- settings --

func (s *Server) getSettings(w http.ResponseWriter, r *http.Request) {
	settings, err := s.DB.Settings(r.Context())
	if err != nil {
		s.fail(w, r, err, "Could not load settings.")
		return
	}
	writeJSON(w, http.StatusOK, settings)
}

func (s *Server) putSettings(w http.ResponseWriter, r *http.Request) {
	var in map[string]string
	if err := decode(r, &in); err != nil {
		writeErr(w, http.StatusBadRequest, "bad_request", "That request could not be read.")
		return
	}
	repriceNeeded := false
	for k, v := range in {
		if err := s.DB.SetSetting(r.Context(), k, v); err != nil {
			s.fail(w, r, err, "Could not save settings.")
			return
		}
		switch k {
		case "fx_thb_per_cny_ppm", "rounding_step_satang", "rounding_mode",
			"intl_rate_satang_per_kg", "intl_min_satang":
			repriceNeeded = true
		}
	}
	// Grid prices are denormalised, so anything that moves a price has to trigger
	// a reprice or the grid and the product page will disagree.
	if repriceNeeded {
		_ = s.DB.Enqueue(r.Context(), store.JobReprice, "all", nil, time.Now())
	}
	s.getSettings(w, r)
}

// ---------------------------------------------------------------- fee rules --

type feeRuleDTO struct {
	ID            string     `json:"id"`
	Scope         string     `json:"scope"`
	ScopeValue    string     `json:"scopeValue"`
	FeeBps        int        `json:"feeBps"`
	FeeFixedFen   int64      `json:"feeFixedFen"`
	MinFeeFen     int64      `json:"minFeeFen"`
	Priority      int        `json:"priority"`
	EffectiveFrom time.Time  `json:"effectiveFrom"`
	EffectiveTo   *time.Time `json:"effectiveTo,omitempty"`
	Note          string     `json:"note"`
}

func feeRuleOut(r store.FeeRule) feeRuleDTO {
	return feeRuleDTO{
		ID: id(r.ID), Scope: r.Scope, ScopeValue: r.ScopeValue, FeeBps: int(r.FeeBps),
		FeeFixedFen: r.FeeFixedFen, MinFeeFen: r.MinFeeFen, Priority: int(r.Priority),
		EffectiveFrom: r.EffectiveFrom, EffectiveTo: r.EffectiveTo, Note: r.Note,
	}
}

func (s *Server) listFeeRules(w http.ResponseWriter, r *http.Request) {
	rules, err := s.DB.FeeRules(r.Context())
	if err != nil {
		s.fail(w, r, err, "Could not load fee rules.")
		return
	}
	out := make([]feeRuleDTO, 0, len(rules))
	for _, fr := range rules {
		out = append(out, feeRuleOut(fr))
	}
	writeJSON(w, http.StatusOK, out)
}

func (s *Server) createFeeRule(w http.ResponseWriter, r *http.Request) {
	in, ok := s.readFeeRule(w, r)
	if !ok {
		return
	}
	fr, err := s.DB.CreateFeeRule(r.Context(), in)
	if err != nil {
		s.fail(w, r, err, "Could not save the rule.")
		return
	}
	_ = s.DB.Enqueue(r.Context(), store.JobReprice, "all", nil, time.Now())
	writeJSON(w, http.StatusCreated, feeRuleOut(fr))
}

func (s *Server) updateFeeRule(w http.ResponseWriter, r *http.Request) {
	ruleID, err := pathInt(r, "id")
	if err != nil {
		writeErr(w, http.StatusBadRequest, "bad_request", "That rule id is not valid.")
		return
	}
	in, ok := s.readFeeRule(w, r)
	if !ok {
		return
	}
	in.ID = ruleID
	fr, err := s.DB.UpdateFeeRule(r.Context(), in)
	if s.notFound(w, err, "That rule") {
		return
	}
	if err != nil {
		s.fail(w, r, err, "Could not save the rule.")
		return
	}
	_ = s.DB.Enqueue(r.Context(), store.JobReprice, "all", nil, time.Now())
	writeJSON(w, http.StatusOK, feeRuleOut(fr))
}

func (s *Server) deleteFeeRule(w http.ResponseWriter, r *http.Request) {
	ruleID, err := pathInt(r, "id")
	if err != nil {
		writeErr(w, http.StatusBadRequest, "bad_request", "That rule id is not valid.")
		return
	}
	if err := s.DB.DeleteFeeRule(r.Context(), ruleID); err != nil {
		s.fail(w, r, err, "Could not delete the rule.")
		return
	}
	_ = s.DB.Enqueue(r.Context(), store.JobReprice, "all", nil, time.Now())
	writeJSON(w, http.StatusOK, map[string]bool{"ok": true})
}

func (s *Server) readFeeRule(w http.ResponseWriter, r *http.Request) (store.FeeRule, bool) {
	var in feeRuleDTO
	if err := decode(r, &in); err != nil {
		writeErr(w, http.StatusBadRequest, "bad_request", "That request could not be read.")
		return store.FeeRule{}, false
	}
	switch in.Scope {
	case "global", "category", "supplier", "product":
	default:
		writeErr(w, http.StatusBadRequest, "bad_request", "Scope must be global, category, supplier or product.")
		return store.FeeRule{}, false
	}
	if in.Scope != "global" && strings.TrimSpace(in.ScopeValue) == "" {
		writeErr(w, http.StatusBadRequest, "bad_request", "This scope needs a value to match on.")
		return store.FeeRule{}, false
	}
	return store.FeeRule{
		Scope: in.Scope, ScopeValue: in.ScopeValue, FeeBps: int32(in.FeeBps),
		FeeFixedFen: in.FeeFixedFen, MinFeeFen: in.MinFeeFen, Priority: int32(in.Priority),
		EffectiveFrom: in.EffectiveFrom, EffectiveTo: in.EffectiveTo, Note: in.Note,
	}, true
}

// previewFee prices one SKU at one quantity and returns the full breakdown. This
// is the fastest way to see what a rule change actually does.
func (s *Server) previewFee(w http.ResponseWriter, r *http.Request) {
	var in struct {
		OfferID  string `json:"offerId"`
		SkuID    string `json:"skuId"`
		Quantity int    `json:"quantity"`
	}
	if err := decode(r, &in); err != nil {
		writeErr(w, http.StatusBadRequest, "bad_request", "That request could not be read.")
		return
	}
	offerID, err := strconv.ParseInt(in.OfferID, 10, 64)
	if err != nil {
		writeErr(w, http.StatusBadRequest, "bad_request", "That product id is not valid.")
		return
	}
	if in.Quantity <= 0 {
		in.Quantity = 1
	}

	fp, err := s.DB.GetProduct(r.Context(), offerID)
	if s.notFound(w, err, "That product") {
		return
	}
	if err != nil {
		s.fail(w, r, err, "Could not load the product.")
		return
	}
	if len(fp.SKUs) == 0 {
		writeErr(w, http.StatusBadRequest, "bad_request", "That product has no options to price.")
		return
	}

	sku := fp.SKUs[0]
	if in.SkuID != "" {
		if want, err := strconv.ParseInt(in.SkuID, 10, 64); err == nil {
			for _, sk := range fp.SKUs {
				if sk.SkuID == want {
					sku = sk
				}
			}
		}
	}

	settings, err := s.DB.Settings(r.Context())
	if err != nil {
		s.fail(w, r, err, "Could not load settings.")
		return
	}
	rules, err := s.DB.FeeRules(r.Context())
	if err != nil {
		s.fail(w, r, err, "Could not load fee rules.")
		return
	}

	res := pricing.Quote(
		[]pricing.Line{{
			Product:  catalog.ToPricingProduct(fp),
			SKU:      catalog.ToPricingSKU(sku),
			Quantity: in.Quantity,
		}},
		map[int64]int{offerID: in.Quantity},
		pricing.Rules(catalog.ToPricingRules(rules), time.Now()),
		catalog.ToPricingSettings(settings),
	)

	out := map[string]any{
		"offerId":  id(offerID),
		"skuId":    id(sku.SkuID),
		"quantity": in.Quantity,
		"issues":   []IssueDTO{},
	}
	for _, is := range res.Issues {
		out["issues"] = append(out["issues"].([]IssueDTO), issueDTO(is))
	}
	if len(res.Lines) > 0 {
		l := res.Lines[0]
		out["breakdown"] = breakdownDTO(l)
		out["unit"] = thb(l.UnitSatang)
		out["line"] = thb(l.LineSatang)
		out["feeRuleId"] = id(l.FeeRuleID)
	}
	out["totals"] = CartTotalsDTO{
		Goods:        thb(res.Totals.GoodsSatang),
		ChinaFreight: thb(res.Totals.FreightSatang),
		Intl:         thb(res.Totals.IntlSatang),
		Fee:          thb(res.Totals.FeeSatang),
		Total:        thb(res.Totals.TotalSatang),
	}
	writeJSON(w, http.StatusOK, out)
}

// ---------------------------------------------------------------- catalogue --

func (s *Server) startImport(w http.ResponseWriter, r *http.Request) {
	var in struct {
		Keyword  string   `json:"keyword"`
		OfferIDs []string `json:"offerIds"`
		Pages    int      `json:"pages"`
		PageSize int      `json:"pageSize"`
		All      bool     `json:"all"`
	}
	if err := decode(r, &in); err != nil {
		writeErr(w, http.StatusBadRequest, "bad_request", "That request could not be read.")
		return
	}

	payload := map[string]any{
		"keyword": in.Keyword, "pages": in.Pages, "pageSize": in.PageSize, "all": in.All,
	}
	var ids []int64
	for _, raw := range in.OfferIDs {
		if v, err := strconv.ParseInt(strings.TrimSpace(raw), 10, 64); err == nil {
			ids = append(ids, v)
		}
	}
	if len(ids) > 0 {
		payload["offerIds"] = ids
	}

	key := "manual:" + strconv.FormatInt(time.Now().UnixNano(), 36)
	if err := s.DB.Enqueue(r.Context(), store.JobCatalogImport, key, payload, time.Now()); err != nil {
		s.fail(w, r, err, "Could not start the import.")
		return
	}
	writeJSON(w, http.StatusAccepted, map[string]any{"jobKey": key, "queued": true})
}

func (s *Server) adminProducts(w http.ResponseWriter, r *http.Request) {
	res, err := s.DB.SearchProducts(r.Context(), store.SearchQuery{
		Q:          strings.TrimSpace(r.URL.Query().Get("q")),
		Page:       queryInt(r, "page", 1),
		Size:       queryInt(r, "size", 50),
		IncludeAll: true,
	})
	if err != nil {
		s.fail(w, r, err, "Could not load products.")
		return
	}
	total, published, _ := s.DB.CountProducts(r.Context())
	items := make([]map[string]any, 0, len(res.Items))
	for _, p := range res.Items {
		items = append(items, map[string]any{
			"offerId":  id(p.OfferID),
			"title":    p.SubjectTrans,
			"status":   p.Status,
			"visible":  p.Visible,
			"seller":   p.CompanyName,
			"price":    thb(ali.Satang(p.SellMinSatang)),
			"stock":    p.AmountOnSale,
			"syncedAt": p.SyncedAt,
		})
	}
	writeJSON(w, http.StatusOK, map[string]any{
		"total": res.Total, "page": res.Page, "size": res.Size, "items": items,
		"counts": map[string]int{"all": total, "published": published},
	})
}

func (s *Server) setProductVisible(w http.ResponseWriter, r *http.Request) {
	offerID, err := pathInt(r, "offerId")
	if err != nil {
		writeErr(w, http.StatusBadRequest, "bad_request", "That product id is not valid.")
		return
	}
	var in struct {
		Visible bool `json:"visible"`
	}
	if err := decode(r, &in); err != nil {
		writeErr(w, http.StatusBadRequest, "bad_request", "That request could not be read.")
		return
	}
	if err := s.DB.SetProductVisible(r.Context(), offerID, in.Visible); err != nil {
		s.fail(w, r, err, "Could not update the product.")
		return
	}
	writeJSON(w, http.StatusOK, map[string]bool{"ok": true})
}

// ------------------------------------------------------------------- orders --

func (s *Server) adminOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := s.DB.ListOrders(r.Context(), r.URL.Query().Get("status"), queryInt(r, "limit", 100))
	if err != nil {
		s.fail(w, r, err, "Could not load orders.")
		return
	}
	out := make([]map[string]any, 0, len(orders))
	for _, o := range orders {
		sos, _ := s.DB.SupplierOrders(r.Context(), o.ID)
		parcels := make([]map[string]any, 0, len(sos))
		for _, so := range sos {
			p := map[string]any{
				"id": id(so.ID), "seller": so.SellerName, "status": so.Status,
				"outOrderId": so.OutOrderID, "payUrl": so.PayURL, "error": so.ErrorMessage,
			}
			if so.CbuOrderID != nil {
				p["cbuOrderId"] = id(*so.CbuOrderID)
			}
			parcels = append(parcels, p)
		}
		out = append(out, map[string]any{
			"orderId": o.PublicID, "status": order.RollUp(o, sos), "email": o.Email,
			"total": thb(ali.Satang(o.TotalSatang)), "placedAt": o.CreatedAt,
			"paidAt": o.PaidAt, "token": o.AccessToken, "parcels": parcels,
		})
	}
	writeJSON(w, http.StatusOK, out)
}

func (s *Server) adminOrder(w http.ResponseWriter, r *http.Request) {
	o, err := s.DB.GetOrderByPublicID(r.Context(), r.PathValue("id"))
	if s.notFound(w, err, "That order") {
		return
	}
	if err != nil {
		s.fail(w, r, err, "Could not load the order.")
		return
	}
	dto, err := s.orderView(r, o)
	if err != nil {
		s.fail(w, r, err, "Could not load the order.")
		return
	}
	writeJSON(w, http.StatusOK, dto)
}

func (s *Server) retryRelay(w http.ResponseWriter, r *http.Request) {
	soID, err := pathInt(r, "id")
	if err != nil {
		writeErr(w, http.StatusBadRequest, "bad_request", "That parcel id is not valid.")
		return
	}
	so, err := s.DB.GetSupplierOrder(r.Context(), soID)
	if s.notFound(w, err, "That parcel") {
		return
	}
	if err != nil {
		s.fail(w, r, err, "Could not load the parcel.")
		return
	}
	kind := store.JobRelayCreate
	if so.CbuOrderID != nil {
		kind = store.JobRelayPay
	}
	if err := s.DB.Enqueue(r.Context(), kind, "so:"+id(so.ID),
		map[string]int64{"supplierOrderId": so.ID}, time.Now()); err != nil {
		s.fail(w, r, err, "Could not queue the retry.")
		return
	}
	writeJSON(w, http.StatusAccepted, map[string]any{"queued": kind})
}

func (s *Server) adminCancelSupplier(w http.ResponseWriter, r *http.Request) {
	soID, err := pathInt(r, "id")
	if err != nil {
		writeErr(w, http.StatusBadRequest, "bad_request", "That parcel id is not valid.")
		return
	}
	var in struct {
		Reason string `json:"reason"`
		Remark string `json:"remark"`
	}
	_ = decode(r, &in)

	so, err := s.DB.GetSupplierOrder(r.Context(), soID)
	if s.notFound(w, err, "That parcel") {
		return
	}
	if err != nil {
		s.fail(w, r, err, "Could not load the parcel.")
		return
	}
	if in.Reason == "" {
		in.Reason = "buyerCancel"
	}
	if err := s.Orders.CancelSupplierOrder(r.Context(), so, in.Reason, in.Remark); err != nil {
		s.fail(w, r, err, "Could not queue the cancellation.")
		return
	}
	writeJSON(w, http.StatusAccepted, map[string]bool{"queued": true})
}

// --------------------------------------------------------------------- logs --

func (s *Server) listJobs(w http.ResponseWriter, r *http.Request) {
	list, err := s.DB.ListJobs(r.Context(), r.URL.Query().Get("state"), queryInt(r, "limit", 100))
	if err != nil {
		s.fail(w, r, err, "Could not load jobs.")
		return
	}
	out := make([]map[string]any, 0, len(list))
	for _, j := range list {
		out = append(out, map[string]any{
			"id": id(j.ID), "kind": j.Kind, "key": j.Key, "state": j.State,
			"attempts": j.Attempts, "runAt": j.RunAt, "lastError": j.LastError,
			"updatedAt": j.UpdatedAt,
		})
	}
	writeJSON(w, http.StatusOK, out)
}

func (s *Server) retryJob(w http.ResponseWriter, r *http.Request) {
	jobID, err := pathInt(r, "id")
	if err != nil {
		writeErr(w, http.StatusBadRequest, "bad_request", "That job id is not valid.")
		return
	}
	if err := s.DB.ReviveJob(r.Context(), jobID); err != nil {
		s.fail(w, r, err, "Could not retry the job.")
		return
	}
	writeJSON(w, http.StatusOK, map[string]bool{"ok": true})
}

func (s *Server) listAPICalls(w http.ResponseWriter, r *http.Request) {
	calls, err := s.DB.RecentAPICalls(r.Context(), queryInt(r, "limit", 100))
	if err != nil {
		s.fail(w, r, err, "Could not load the call log.")
		return
	}
	out := make([]map[string]any, 0, len(calls))
	for _, c := range calls {
		out = append(out, map[string]any{
			"id": id(c.ID), "api": c.API, "ms": c.MS, "ok": c.OK, "code": c.Code,
			"message": c.Message, "at": c.CreatedAt,
		})
	}
	writeJSON(w, http.StatusOK, out)
}

func (s *Server) listMessages(w http.ResponseWriter, r *http.Request) {
	msgs, err := s.DB.RecentMessages(r.Context(), queryInt(r, "limit", 100))
	if err != nil {
		s.fail(w, r, err, "Could not load messages.")
		return
	}
	out := make([]map[string]any, 0, len(msgs))
	for _, m := range msgs {
		out = append(out, map[string]any{
			"msgId": id(m.MsgID), "type": m.Type, "bornAt": m.GmtBorn,
			"receivedAt": m.ReceivedAt, "processedAt": m.ProcessedAt, "error": m.Error,
			"payload": json.RawMessage(m.Payload),
		})
	}
	writeJSON(w, http.StatusOK, out)
}

// proxyStub forwards the admin console's simulation controls to the stub
// gateway, so the browser talks to a single origin.
//
// At go-live STUB_BASE is unset and this route answers 404, which is the correct
// behaviour: there is no simulation to drive against the real 1688.
func (s *Server) proxyStub(w http.ResponseWriter, r *http.Request) {
	if s.StubBase == "" {
		writeErr(w, http.StatusNotFound, "no_stub", "No stub gateway is configured.")
		return
	}
	rest := strings.TrimPrefix(r.URL.Path, "/api/admin/stub")
	target := strings.TrimRight(s.StubBase, "/") + "/_control" + rest
	if q := r.URL.RawQuery; q != "" {
		target += "?" + q
	}
	if _, err := url.Parse(target); err != nil {
		writeErr(w, http.StatusBadRequest, "bad_request", "That control path is not valid.")
		return
	}

	body, _ := io.ReadAll(io.LimitReader(r.Body, 1<<20))
	req, err := http.NewRequestWithContext(r.Context(), r.Method, target, strings.NewReader(string(body)))
	if err != nil {
		s.fail(w, r, err, "Could not reach the stub gateway.")
		return
	}
	req.Header.Set("Content-Type", "application/json")
	if tok := r.Header.Get("X-Control-Token"); tok != "" {
		req.Header.Set("X-Control-Token", tok)
	}

	resp, err := (&http.Client{Timeout: 15 * time.Second}).Do(req)
	if err != nil {
		writeErr(w, http.StatusBadGateway, "stub_unreachable", "The stub gateway did not answer.")
		return
	}
	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(resp.StatusCode)
	_, _ = io.Copy(w, io.LimitReader(resp.Body, 4<<20))
}
