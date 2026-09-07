package ali

import (
	"errors"
	"fmt"
	"net"
	"strings"
)

// APIError is a business-level failure reported by the gateway: the HTTP call
// succeeded but 1688 refused the request. Transport failures are returned as
// plain wrapped errors instead.
type APIError struct {
	API     string // "com.alibaba.trade:alibaba.createOrder.preview:1"
	Code    string
	Message string
	HTTP    int
	Body    string // truncated, for the admin call log
}

func (e *APIError) Error() string {
	if e.Message == "" {
		return fmt.Sprintf("%s: %s", e.API, e.Code)
	}
	return fmt.Sprintf("%s: %s: %s", e.API, e.Code, e.Message)
}

// Code returns the gateway error code carried by err, or "".
func Code(err error) string {
	var ae *APIError
	if errors.As(err, &ae) {
		return ae.Code
	}
	return ""
}

// IsCode reports whether err is an APIError with one of the given codes.
func IsCode(err error, codes ...string) bool {
	got := Code(err)
	if got == "" {
		return false
	}
	for _, c := range codes {
		if strings.EqualFold(got, c) {
			return true
		}
	}
	return false
}

// Business codes documented for the endpoints this build calls. Every one of
// these means "do not retry, tell a human"; contract_test.go asserts that each
// code listed in the docs appears here.
const (
	// alibaba.createOrder.preview
	ErrNoOnlineTrade  = "500_001" // offer does not support online trade
	ErrMultiSeller    = "500_002" // several sellers in one order, or specId missing
	ErrSpecNotInOffer = "500_003"
	ErrNoStock        = "500_004"
	ErrBelowMOQ       = "500_005"
	ErrMixedBatch     = "500_006"
	ErrNoConsignRel   = "500_007"
	ErrZeroPrice      = "500_008"
	ErrBelowWholesale = "500_009"

	// alibaba.trade.cancel
	ErrCloseTooFast  = "CLOSE_ORDER_TOO_FAST" // cancelled under ten seconds after creation
	ErrOrderStatus   = "ORDER_STATUS_ERROR"
	ErrNoPermission  = "400_3"
	ErrOrderNotExist = "ORDER_NOT_EXIST"

	// alibaba.trade.getLogisticsTraceInfo.buyerView
	ErrTraceNotFound = "404"
	ErrTraceNoPerm   = "order.nopermission.buyer"
	ErrTraceTooOld   = "order.createtime.history"
)

// businessCodes is every code above, for classification and for the contract test.
func businessCodes() []string {
	return []string{
		ErrNoOnlineTrade, ErrMultiSeller, ErrSpecNotInOffer, ErrNoStock, ErrBelowMOQ,
		ErrMixedBatch, ErrNoConsignRel, ErrZeroPrice, ErrBelowWholesale,
		ErrCloseTooFast, ErrOrderStatus, ErrNoPermission, ErrOrderNotExist,
		ErrTraceNotFound, ErrTraceNoPerm, ErrTraceTooOld,
	}
}

// BusinessCodes is the exported view used by tests.
func BusinessCodes() []string { return businessCodes() }

// Retryable reports whether err is worth trying again. Network trouble and
// gateway 5xx are; a documented business refusal never is, because retrying a
// rejected order just produces the same rejection.
func Retryable(err error) bool {
	if err == nil {
		return false
	}
	var ae *APIError
	if errors.As(err, &ae) {
		if ae.HTTP >= 500 {
			return true
		}
		for _, c := range businessCodes() {
			if strings.EqualFold(ae.Code, c) {
				return false
			}
		}
		// An unrecognised code with a 2xx: assume it is business logic.
		return false
	}
	var nerr net.Error
	if errors.As(err, &nerr) {
		return true
	}
	return true // transport, DNS, connection reset and friends
}
