package stub

import (
	"net/url"
	"time"

	"marketplace/internal/ali"
)

// handleLogisticsTrace serves
// com.alibaba.logistics:alibaba.trade.getLogisticsTraceInfo.buyerView:1.
//
// The envelope is bespoke — logisticsTrace sits at the top level beside success
// — and the steps are {acceptTime, remark} pairs in ascending time, exactly as
// the documented sample shows them. acceptTime is the NAIVE layout, which means
// Asia/Shanghai with no zone marker.
func (s *Server) handleLogisticsTrace(form url.Values) any {
	id, ok := parseInt64(form.Get("orderId"))
	if !ok {
		return traceError(ali.ErrTraceNotFound, "无法找到对应的物流跟踪信息")
	}

	o, found := s.store.Get(id)
	if !found {
		return traceError(ali.ErrTraceNotFound, "无法找到对应的物流跟踪信息")
	}
	if o.BuyerMemberID != buyer.MemberID {
		return traceError(ali.ErrTraceNoPerm, "无权获取该订单详情（买家侧）")
	}
	// Documented: tracking for orders older than a year is not served.
	if s.clock.Now().Sub(o.CreateTime) > 365*24*time.Hour {
		return traceError(ali.ErrTraceTooOld, "不支持查询一年以前的物流跟踪信息")
	}
	if o.Logistics == nil || len(o.Logistics.Steps) == 0 {
		// Nothing has shipped yet: the same 404 the real gateway gives, since
		// the waybill genuinely does not exist.
		return traceError(ali.ErrTraceNotFound, "无法找到对应的物流跟踪信息")
	}
	// A logisticsId was supplied and does not match this order's waybill.
	if lid := form.Get("logisticsId"); lid != "" && lid != o.Logistics.LogisticsID {
		return traceError(ali.ErrTraceNotFound, "无法找到对应的物流跟踪信息")
	}

	steps := make([]any, 0, len(o.Logistics.Steps))
	for _, st := range o.Logistics.Steps {
		steps = append(steps, map[string]any{
			"acceptTime": ali.FormatNaive(st.At),
			"remark":     st.Remark,
		})
	}

	return map[string]any{
		"logisticsTrace": []any{map[string]any{
			"logisticsId":     o.Logistics.LogisticsID,
			"orderId":         o.ID,
			"logisticsBillNo": o.Logistics.MailNo,
			"logisticsSteps":  steps,
		}},
		"errorCode":                  "",
		"errorMessage":               "",
		"success":                    true,
		"crossPackageFulfillmentDTO": []any{},
	}
}

func traceError(code, msg string) map[string]any {
	return map[string]any{
		"logisticsTrace": []any{},
		"errorCode":      code,
		"errorMessage":   msg,
		"success":        false,
	}
}
