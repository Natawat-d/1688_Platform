package ali

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

// CST is Asia/Shanghai, fixed at +08:00. It is defined here rather than loaded
// from the zone database so the binaries need no tzdata.
var CST = time.FixedZone("CST", 8*60*60)

// ErrBadTime is returned when no known 1688 layout matches.
var ErrBadTime = errors.New("ali: unrecognised timestamp")

// 1688 emits at least four timestamp spellings across this API set:
//
//	20180614101942000+0800   order detail, order list, account basic
//	20180614101942000        the same without an offset
//	2018-07-24 21:55:33      logistics trace steps, push msgSendTime — NAIVE, and
//	                         it means Asia/Shanghai. Reading it as UTC shifts every
//	                         timeline by eight hours, which is the kind of bug that
//	                         looks like a data problem for a week.
//	1564984329147            epoch milliseconds, sometimes quoted (inventory bizTime)
const (
	layoutCompactMS = "20060102150405"
	layoutNaive     = "2006-01-02 15:04:05"
	layoutDate      = "2006-01-02"
)

// ParseTime accepts every layout above and returns a zero Time for empty input.
func ParseTime(s string) (time.Time, error) {
	s = strings.TrimSpace(strings.Trim(strings.TrimSpace(s), `"`))
	if s == "" || s == "null" {
		return time.Time{}, nil
	}

	// 20180614101942000+0800 — Go cannot parse fractional seconds without a
	// separator, so lift the milliseconds out and rebuild the string.
	if len(s) >= 17 && allDigits(s[:14]) && allDigits(s[14:17]) {
		ms, _ := strconv.Atoi(s[14:17])
		rest := s[17:]
		base, err := time.ParseInLocation(layoutCompactMS, s[:14], CST)
		if err == nil {
			if rest != "" {
				if off, ok := parseOffset(rest); ok {
					base, err = time.ParseInLocation(layoutCompactMS, s[:14], off)
					if err != nil {
						return time.Time{}, err
					}
				}
			}
			return base.Add(time.Duration(ms) * time.Millisecond), nil
		}
	}

	if t, err := time.Parse(time.RFC3339, s); err == nil {
		return t, nil
	}
	if t, err := time.ParseInLocation(layoutNaive, s, CST); err == nil {
		return t, nil
	}
	if t, err := time.ParseInLocation(layoutDate, s, CST); err == nil {
		return t, nil
	}
	// Epoch milliseconds, quoted or bare.
	if allDigits(s) && len(s) >= 12 {
		if ms, err := strconv.ParseInt(s, 10, 64); err == nil {
			return time.UnixMilli(ms).In(CST), nil
		}
	}
	return time.Time{}, ErrBadTime
}

// parseOffset reads "+0800" or "-0700" into a fixed zone.
func parseOffset(s string) (*time.Location, bool) {
	if len(s) != 5 || (s[0] != '+' && s[0] != '-') || !allDigits(s[1:]) {
		return nil, false
	}
	h, _ := strconv.Atoi(s[1:3])
	m, _ := strconv.Atoi(s[3:5])
	secs := h*3600 + m*60
	if s[0] == '-' {
		secs = -secs
	}
	return time.FixedZone("", secs), true
}

// FormatNaive renders the layout used by push messages and trace steps:
// Shanghai wall-clock time with no zone marker.
func FormatNaive(t time.Time) string { return t.In(CST).Format(layoutNaive) }

// FormatCompact renders the layout used by the order-reading APIs.
func FormatCompact(t time.Time) string {
	t = t.In(CST)
	return t.Format(layoutCompactMS) + pad3(t.Nanosecond()/1e6) + t.Format("-0700")
}

func pad3(v int) string {
	s := strconv.Itoa(v)
	for len(s) < 3 {
		s = "0" + s
	}
	return s
}

// Millis is epoch milliseconds, the unit the push envelope uses for gmtBorn.
func Millis(t time.Time) int64 { return t.UnixMilli() }
