package ali

import (
	"errors"
	"strconv"
	"strings"
)

// Money in this codebase is ALWAYS an integer count of minor units. There is no
// float64 anywhere in the money path; internal/pricing has a test that fails the
// build if one appears.
//
// Which unit a 1688 field carries is not consistent, so it is written down here
// once and encoded in the struct tags of catalog.go / trade.go:
//
//	fen (integer)      sumPayment, sumCarriage, sumPaymentNoCarriage, discountFee,
//	                   additionalFee, postFee, totalSuccessAmount, orderAmmount,
//	                   applyPayment, applyCarriage, innerPostFee
//	fen (integer)      order detail baseInfo.discount and baseInfo.refundPayment
//	                   — yes, cents, while every sibling amount is yuan
//	yuan (decimal)     order detail totalAmount, sumProductPayment, shippingFee,
//	                   couponFee, refund, productItems[].price, productItems[].itemAmount
//	yuan (string)      product prices: productSkuInfos[].price, priceRangeList[].price,
//	                   search priceInfo.*, freight estimate freight/firstFee/nextFee
//
// Rule of thumb: the trade-creation family speaks fen, the order-reading family
// speaks yuan, and the catalogue speaks decimal strings.

// Fen is one hundredth of a CNY. 1850 is ¥18.50.
type Fen int64

// Satang is one hundredth of a THB. 12345 is ฿123.45.
type Satang int64

// ErrBadAmount is returned for input that is not a plain decimal number.
var ErrBadAmount = errors.New("ali: not a decimal amount")

// ParseFen converts a decimal yuan string to Fen exactly, without ever going
// through a float. "18.5" and "18.50" both give 1850. More than two decimal
// places round half away from zero. Empty input is 0, because 1688 routinely
// sends "" for absent prices.
func ParseFen(s string) (Fen, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0, nil
	}
	neg := false
	switch s[0] {
	case '-':
		neg, s = true, s[1:]
	case '+':
		s = s[1:]
	}
	if s == "" {
		return 0, ErrBadAmount
	}

	whole, frac, hasFrac := strings.Cut(s, ".")
	if whole == "" {
		whole = "0"
	}
	if !allDigits(whole) || (hasFrac && !allDigits(frac)) {
		return 0, ErrBadAmount
	}

	w, err := strconv.ParseInt(whole, 10, 64)
	if err != nil {
		return 0, ErrBadAmount
	}

	// Take two decimal places, rounding half away from zero on the third.
	var cents int64
	switch {
	case len(frac) == 0:
	case len(frac) == 1:
		cents = int64(frac[0]-'0') * 10
	default:
		cents = int64(frac[0]-'0')*10 + int64(frac[1]-'0')
		if len(frac) > 2 && frac[2] >= '5' {
			cents++
		}
	}

	v := w*100 + cents
	if neg {
		v = -v
	}
	return Fen(v), nil
}

// MustParseFen is ParseFen for fixtures and tests, where the input is a literal.
func MustParseFen(s string) Fen {
	v, err := ParseFen(s)
	if err != nil {
		panic("ali.MustParseFen(" + s + "): " + err.Error())
	}
	return v
}

func allDigits(s string) bool {
	if s == "" {
		return false
	}
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return true
}

// String renders Fen the way 1688 sends prices back: a plain decimal, two places.
func (f Fen) String() string { return minor(int64(f)) }

// String renders Satang as a plain decimal, two places, with no currency symbol.
func (s Satang) String() string { return minor(int64(s)) }

// Text renders Satang for display, e.g. "฿1,234.50".
func (s Satang) Text() string { return "฿" + group(minor(int64(s))) }

// Text renders Fen for display, e.g. "¥18.50".
func (f Fen) Text() string { return "¥" + group(minor(int64(f))) }

func minor(v int64) string {
	sign := ""
	if v < 0 {
		sign, v = "-", -v
	}
	return sign + strconv.FormatInt(v/100, 10) + "." + twoDigits(v%100)
}

func twoDigits(v int64) string {
	if v < 10 {
		return "0" + strconv.FormatInt(v, 10)
	}
	return strconv.FormatInt(v, 10)
}

// group inserts thousands separators into the whole part of a decimal string.
func group(s string) string {
	neg := strings.HasPrefix(s, "-")
	s = strings.TrimPrefix(s, "-")
	whole, frac, _ := strings.Cut(s, ".")
	var b strings.Builder
	if neg {
		b.WriteByte('-')
	}
	for i, c := range whole {
		if i > 0 && (len(whole)-i)%3 == 0 {
			b.WriteByte(',')
		}
		b.WriteRune(c)
	}
	b.WriteByte('.')
	b.WriteString(frac)
	return b.String()
}

// ToSatang converts CNY to THB using a rate in parts per million (THB per CNY
// times 1e6), rounding half up. Integer arithmetic throughout: the largest
// realistic product here is well inside int64.
func (f Fen) ToSatang(fxPPM int64) Satang {
	return Satang(divRoundHalfUp(int64(f)*fxPPM, 1_000_000))
}

// ToFen is the inverse of ToSatang, used to express an international shipping
// charge back in CNY so a price breakdown reads in a single currency.
func (s Satang) ToFen(fxPPM int64) Fen {
	if fxPPM == 0 {
		return 0
	}
	return Fen(divRoundHalfUp(int64(s)*1_000_000, fxPPM))
}

// divRoundHalfUp divides a by b, rounding halves away from zero. b must be > 0.
func divRoundHalfUp(a, b int64) int64 {
	if b <= 0 {
		return 0
	}
	if a >= 0 {
		return (a + b/2) / b
	}
	return -((-a + b/2) / b)
}

// DivRoundHalfUp is the exported form, used by the pricing engine.
func DivRoundHalfUp(a, b int64) int64 { return divRoundHalfUp(a, b) }
