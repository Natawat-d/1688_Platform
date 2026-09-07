package ali

import (
	"context"
	"errors"
)

// Logistics tracking, buyer view. One order can carry several waybills, so the
// response is a list of traces, each with its own steps.
//
// The steps are the China leg of the journey. Everything after our consolidation
// warehouse comes from the forwarder, not from here.

// LogisticsTraceResult is the response of
// alibaba.trade.getLogisticsTraceInfo.buyerView.
type LogisticsTraceResult struct {
	LogisticsTrace             []LogisticsTraceInfo      `json:"logisticsTrace"`
	ErrorCode                  string                    `json:"errorCode"`
	ErrorMessage               string                    `json:"errorMessage"`
	Success                    FlexBool                  `json:"success"`
	CrossPackageFulfillmentDTO []CrossPackageFulfillment `json:"crossPackageFulfillmentDTO"`
}

// LogisticsTraceInfo is alibaba.logistics.OpenPlatformLogisticsTrace: one
// waybill and its scan history. logisticsId is 1688's own id for the waybill;
// logisticsBillNo is the carrier's tracking number.
type LogisticsTraceInfo struct {
	LogisticsID     string          `json:"logisticsId"`
	OrderID         ID              `json:"orderId"`
	LogisticsBillNo string          `json:"logisticsBillNo"`
	LogisticsSteps  []LogisticsStep `json:"logisticsSteps"`
}

// LogisticsStep is alibaba.logistics.OpenPlatformLogisticsStep, one scan.
//
// acceptTime is the naive "2018-07-24 21:55:33" spelling, which means Shanghai
// local time. Reading it as UTC shifts every customer timeline by eight hours.
type LogisticsStep struct {
	AcceptTime Timestamp `json:"acceptTime"`
	Remark     string    `json:"remark"`
}

// CrossPackageFulfillment is the overseas leg 1688 knows about, present only for
// orders shipped through official cross-border logistics.
type CrossPackageFulfillment struct {
	CrossMailNo              string              `json:"crossMailNo"`
	OrderID                  ID                  `json:"orderId"`
	CrossPackageTraceDTOList []CrossPackageTrace `json:"crossPackageTraceDTOList"`
}

// CrossPackageTrace is one node of the overseas shipment timeline.
type CrossPackageTrace struct {
	BuyerUserID        string         `json:"buyerUserId"`
	CarrierPartnerCode string         `json:"carrierPartnerCode"`
	OutTraceNodeCode   string         `json:"outTraceNodeCode"`
	OpCode             string         `json:"opCode"`
	NodeActionTime     Timestamp      `json:"nodeActionTime"`
	NodeDetail         string         `json:"nodeDetail"`
	StageType          string         `json:"stageType"`
	TraceNode          TraceNode      `json:"traceNode"`
	TraceException     ErrorTraceInfo `json:"traceException"`
}

// TraceNode names one node of the overseas timeline.
type TraceNode struct {
	TraceNodeName string `json:"traceNodeName"`
	TraceNodeCode string `json:"traceNodeCode"`
}

// ErrorTraceInfo is an exception raised against an overseas node.
type ErrorTraceInfo struct {
	ExceptionCode string `json:"exceptionCode"`
	TraceNodeCode string `json:"traceNodeCode"`
	ExceptionDesc string `json:"exceptionDesc"`
}

// TraceOrder reads the logistics tracking for one order. logisticsId is
// optional and narrows the answer to a single waybill.
//
// The documented refusals are worth handling rather than retrying: 404 when
// there is no trace yet, order.nopermission.buyer when the order belongs to
// another account, and order.createtime.history past one year.
func (c *Client) TraceOrder(ctx context.Context, orderID ID, logisticsID string) (LogisticsTraceResult, error) {
	if orderID == 0 {
		return LogisticsTraceResult{}, errors.New("ali: getLogisticsTraceInfo.buyerView: orderId is required")
	}
	p := Params{"orderId": orderID, "webSite": WebSite1688}
	setStr(p, "logisticsId", logisticsID)

	body, err := c.Call(ctx, LogisticsTrace, p)
	if err != nil {
		return LogisticsTraceResult{}, err
	}
	var out LogisticsTraceResult
	if err := DecodeFlat(LogisticsTrace.Key(), body, &out); err != nil {
		return LogisticsTraceResult{}, err
	}
	if out.ErrorCode != "" {
		return out, &APIError{
			API:     LogisticsTrace.Key(),
			Code:    out.ErrorCode,
			Message: out.ErrorMessage,
			Body:    truncate(body),
		}
	}
	return out, nil
}
