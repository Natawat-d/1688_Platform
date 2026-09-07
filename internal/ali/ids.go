package ali

import (
	"bytes"
	"encoding/json"
	"strconv"
	"strings"
)

// ID is a 1688 identifier: offer, order, sub-order, sku. They run to nineteen
// digits, so they are int64 here and never float64.
//
// TWO RULES THAT MUST NOT BE CONFLATED:
//
//  1. Outbound, ID marshals to a BARE NUMBER. alibaba.alipay.url.get takes
//     orderIdList as a Long[], and quoting the elements breaks the call.
//  2. Our own JSON API never exposes an ID at all. internal/api/dto.go declares
//     every identifier as a string, and dto_test.go fails if any response
//     carries a bare integer of sixteen digits or more, because JavaScript
//     silently mangles those.
//
// Inbound, 1688 is inconsistent — createCrossOrder returns orderId as a string,
// trade.get.buyerView returns it as a number — so unmarshalling accepts both.
type ID int64

func (i ID) String() string { return strconv.FormatInt(int64(i), 10) }

// Int64 is the value as a plain int64.
func (i ID) Int64() int64 { return int64(i) }

// MarshalJSON emits a bare number. See the type comment before changing this.
func (i ID) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(int64(i), 10)), nil
}

func (i *ID) UnmarshalJSON(b []byte) error {
	s := strings.Trim(strings.TrimSpace(string(b)), `"`)
	if s == "" || s == "null" {
		*i = 0
		return nil
	}
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		// Some payloads carry a decimal for an identifier field. Take the
		// integer part rather than failing the whole document.
		if whole, _, ok := strings.Cut(s, "."); ok {
			if v2, err2 := strconv.ParseInt(whole, 10, 64); err2 == nil {
				*i = ID(v2)
				return nil
			}
		}
		return err
	}
	*i = ID(v)
	return nil
}

// IDs marshals as a JSON array of bare numbers, which is what Long[] params want.
type IDs []ID

// FlexBool decodes true, "true", 1, "1", "y" and "Y" as true. 1688 types the
// same logical field as Boolean on one API and String on another —
// trade.get.buyerView returns success as "true" while account.basic returns it
// as true.
type FlexBool bool

func (f FlexBool) Bool() bool { return bool(f) }

func (f *FlexBool) UnmarshalJSON(b []byte) error {
	s := strings.ToLower(strings.Trim(strings.TrimSpace(string(b)), `"`))
	switch s {
	case "true", "1", "y", "yes":
		*f = true
	default:
		*f = false
	}
	return nil
}

func (f FlexBool) MarshalJSON() ([]byte, error) {
	if f {
		return []byte("true"), nil
	}
	return []byte("false"), nil
}

// FlexInt64 decodes a number or a quoted number, and tolerates "".
type FlexInt64 int64

func (f FlexInt64) Int64() int64 { return int64(f) }

func (f *FlexInt64) UnmarshalJSON(b []byte) error {
	s := strings.Trim(strings.TrimSpace(string(b)), `"`)
	if s == "" || s == "null" {
		*f = 0
		return nil
	}
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		fv, err2 := strconv.ParseFloat(s, 64)
		if err2 != nil {
			return err
		}
		v = int64(fv)
	}
	*f = FlexInt64(v)
	return nil
}

func (f FlexInt64) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(int64(f), 10)), nil
}

// FlexString decodes a JSON string, number or boolean as its text and
// tolerates null. 1688 documents idOfStr and subItemIDString as String — they
// exist so nineteen-digit ids survive tooling that rounds numbers — yet the
// documented getBuyerOrderList sample sends idOfStr as a bare number, and a
// string field would refuse the whole page.
type FlexString string

func (f FlexString) String() string { return string(f) }

func (f *FlexString) UnmarshalJSON(b []byte) error {
	s := strings.TrimSpace(string(b))
	if s == "null" {
		*f = ""
		return nil
	}
	if len(s) > 0 && s[0] == '"' {
		var str string
		if err := json.Unmarshal(b, &str); err != nil {
			return err
		}
		*f = FlexString(str)
		return nil
	}
	*f = FlexString(s)
	return nil
}

func (f FlexString) MarshalJSON() ([]byte, error) { return json.Marshal(string(f)) }

// YuanFen decodes a yuan amount that arrives as a JSON number or a quoted
// decimal, and stores it as Fen. Used for the order-reading family, where
// totalAmount and friends are decimals rather than cents.
type YuanFen Fen

func (y YuanFen) Fen() Fen { return Fen(y) }

func (y *YuanFen) UnmarshalJSON(b []byte) error {
	s := strings.Trim(strings.TrimSpace(string(b)), `"`)
	if s == "" || s == "null" {
		*y = 0
		return nil
	}
	v, err := ParseFen(s)
	if err != nil {
		return err
	}
	*y = YuanFen(v)
	return nil
}

func (y YuanFen) MarshalJSON() ([]byte, error) {
	return []byte(`"` + Fen(y).String() + `"`), nil
}

// PriceString is a catalogue price: 1688 sends these as decimal strings such as
// "18.50", and sometimes as "". Keeping the raw text lets contract tests compare
// against the docs while Fen() gives the value.
type PriceString string

func (p PriceString) Fen() Fen {
	v, err := ParseFen(string(p))
	if err != nil {
		return 0
	}
	return v
}

func (p PriceString) Empty() bool { return strings.TrimSpace(string(p)) == "" }

// UnmarshalJSON accepts a quoted decimal or a bare number, since the docs show
// both spellings for price-shaped fields.
func (p *PriceString) UnmarshalJSON(b []byte) error {
	b = bytes.TrimSpace(b)
	if bytes.Equal(b, []byte("null")) {
		*p = ""
		return nil
	}
	if len(b) > 0 && b[0] == '"' {
		var s string
		if err := json.Unmarshal(b, &s); err != nil {
			return err
		}
		*p = PriceString(s)
		return nil
	}
	*p = PriceString(string(b))
	return nil
}
