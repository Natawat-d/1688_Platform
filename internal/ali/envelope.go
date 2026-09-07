package ali

import (
	"encoding/json"
	"fmt"
)

// 1688 wraps responses in four different ways across the endpoints this build
// uses. Two of them are regular enough to be generic; the rest are one-offs and
// stay one-offs in trade.go and logistics.go, because unifying six bespoke
// shapes costs more than it saves.
//
//	Family A  {"result":{"success":b,"code":s,"message":s,"result":<payload>}}
//	          keyword search, product detail, SN query, category, freight, refund
//	Family B  {"result":<payload>,"errorCode":s,"errorMessage":s,"success":b|"b"}
//	          order detail (success is a String there), account basic
//	Family C  bespoke: preview, create, pay url, order list, trace, cancel
//
// Both generic decoders also accept the "flat" spelling, where the status fields
// sit beside the payload instead of wrapping it. That is not us being defensive
// for its own sake: the documented sample for category.translation.getById is
// flat while its own parameter table says nested, and two of the four samples
// for createCrossOrder are flat while its table says wrapped.

type statusFields struct {
	// Success is a pointer so an absent field is distinguishable from an
	// explicit false. Several documented samples omit it entirely, and treating
	// "absent" as "failed" would reject perfectly good responses.
	Success *FlexBool `json:"success"`
	Code    string    `json:"code"`
	Message string    `json:"message"`
	RetCode string    `json:"retCode"` // keywordSNQuery spells it this way
	RetMsg  string    `json:"retMsg"`
}

// ok reports whether these status fields describe a successful call.
func (s statusFields) ok() bool {
	if !okCode(s.code()) {
		return false
	}
	if s.Success != nil && !bool(*s.Success) {
		return false
	}
	return true
}

func (s statusFields) code() string {
	if s.Code != "" {
		return s.Code
	}
	return s.RetCode
}

func (s statusFields) message() string {
	if s.Message != "" {
		return s.Message
	}
	return s.RetMsg
}

// okCode reports whether a success code means success. 1688 uses "S0000" and
// "200" for success and leaves it empty when it has nothing to say.
func okCode(c string) bool {
	switch c {
	case "", "0", "200", "S0000", "success", "SUCCESS":
		return true
	}
	return false
}

// DecodeA reads a Family A response and returns the payload.
func DecodeA[T any](api string, body []byte) (T, error) {
	var zero T

	var top map[string]json.RawMessage
	if err := json.Unmarshal(body, &top); err != nil {
		return zero, fmt.Errorf("%s: decode: %w", api, err)
	}

	var st statusFields
	_ = json.Unmarshal(body, &st)

	raw, ok := top["result"]
	if !ok {
		return zero, &APIError{API: api, Code: "NO_RESULT", Message: "response carried no result field", Body: truncate(body)}
	}

	// Nested spelling: the inner object carries both a payload and its own status.
	if inner, isNested := nested(raw); isNested {
		var ist statusFields
		_ = json.Unmarshal(raw, &ist)
		if ist.code() != "" || ist.Success != nil {
			st = ist
		}
		raw = inner
	}

	if !st.ok() {
		return zero, &APIError{API: api, Code: firstNonEmpty(st.code(), "FAILED"), Message: st.message(), Body: truncate(body)}
	}

	var out T
	if len(raw) > 0 && string(raw) != "null" {
		if err := json.Unmarshal(raw, &out); err != nil {
			return zero, fmt.Errorf("%s: decode payload: %w", api, err)
		}
	}
	return out, nil
}

// DecodeB reads a Family B response and returns the payload. Members of this
// family sometimes omit success entirely — getBuyerOrderList has no such field —
// so success is inferred from an empty errorCode.
func DecodeB[T any](api string, body []byte) (T, error) {
	var zero T

	var env struct {
		Result       json.RawMessage `json:"result"`
		ErrorCode    string          `json:"errorCode"`
		ErrorMessage string          `json:"errorMessage"`
		ErrorMsg     string          `json:"errorMsg"` // preview spells it this way
		ErroMsg      string          `json:"erroMsg"`  // pay url spells it this way
	}
	if err := json.Unmarshal(body, &env); err != nil {
		return zero, fmt.Errorf("%s: decode: %w", api, err)
	}

	if env.ErrorCode != "" {
		return zero, &APIError{
			API:     api,
			Code:    env.ErrorCode,
			Message: firstNonEmpty(env.ErrorMessage, env.ErrorMsg, env.ErroMsg),
			Body:    truncate(body),
		}
	}

	var out T
	if len(env.Result) > 0 && string(env.Result) != "null" {
		if err := json.Unmarshal(env.Result, &out); err != nil {
			return zero, fmt.Errorf("%s: decode payload: %w", api, err)
		}
	}
	return out, nil
}

// DecodeFlat reads a bespoke Family C response into out, unwrapping a "result"
// object when one is present. createCrossOrder needs exactly this: its
// documentation wraps, two of its four samples do not.
func DecodeFlat(api string, body []byte, out any) error {
	raw := json.RawMessage(body)
	var top map[string]json.RawMessage
	if err := json.Unmarshal(body, &top); err != nil {
		return fmt.Errorf("%s: decode: %w", api, err)
	}
	if inner, ok := top["result"]; ok && looksLikeObject(inner) {
		raw = inner
	}
	if err := json.Unmarshal(raw, out); err != nil {
		return fmt.Errorf("%s: decode payload: %w", api, err)
	}
	return nil
}

// nested reports whether raw is an object carrying both a "result" key and at
// least one status key, which is what distinguishes the wrapping object from a
// payload that merely happens to have a result field of its own.
func nested(raw json.RawMessage) (json.RawMessage, bool) {
	if !looksLikeObject(raw) {
		return nil, false
	}
	var m map[string]json.RawMessage
	if err := json.Unmarshal(raw, &m); err != nil {
		return nil, false
	}
	inner, hasResult := m["result"]
	if !hasResult {
		return nil, false
	}
	for _, k := range []string{"success", "code", "message", "retCode", "retMsg"} {
		if _, ok := m[k]; ok {
			return inner, true
		}
	}
	return nil, false
}

func looksLikeObject(raw json.RawMessage) bool {
	for _, c := range raw {
		switch c {
		case ' ', '\t', '\n', '\r':
			continue
		case '{':
			return true
		default:
			return false
		}
	}
	return false
}

func firstNonEmpty(vals ...string) string {
	for _, v := range vals {
		if v != "" {
			return v
		}
	}
	return ""
}

func truncate(b []byte) string {
	const max = 2000
	if len(b) <= max {
		return string(b)
	}
	return string(b[:max]) + "…"
}
