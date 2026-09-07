package ali

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"net/url"
	"sort"
	"strings"
)

// ============================ UNVERIFIED =====================================
//
// This signing algorithm is reconstructed from our own notes. It is NOT taken
// from an official 1688 source: neither the archived documentation in
// 1688-api-docs/ nor anything else we hold describes it. The authoritative text
// is https://open.1688.com/doc/apiInvoke.htm, which needs to be read the day
// real API access is granted.
//
// The stub gateway validates signatures with this same function, so a wrong
// algorithm is symmetric and cannot be detected locally. That is accepted and
// expected. When live calls start rejecting with an invalid-signature error,
// THIS FILE is the only thing that needs to change.
//
//	signature = UPPER(HEX(HMAC_SHA1(appSecret,
//	    "param2/{version}/{namespace}/{name}/{appKey}" +
//	    concat(key + value for key in sorted(params)))))
//
// The signature parameter itself is excluded from the input.
//
// =============================================================================

// SignatureParam is the request field carrying the signature.
const SignatureParam = "_aop_signature"

// TimestampParam is the optional request timestamp field.
const TimestampParam = "_aop_timestamp"

// TokenParam is the request field carrying the user authorization token.
const TokenParam = "access_token"

// Signer produces the value of _aop_signature for a request.
type Signer interface {
	Sign(signPath string, params url.Values) string
}

// HMACSHA1Signer implements the algorithm described above.
type HMACSHA1Signer struct{ Secret string }

func (s HMACSHA1Signer) Sign(signPath string, params url.Values) string {
	keys := make([]string, 0, len(params))
	for k := range params {
		if k == SignatureParam {
			continue
		}
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var b strings.Builder
	b.WriteString(signPath)
	for _, k := range keys {
		b.WriteString(k)
		b.WriteString(params.Get(k))
	}

	mac := hmac.New(sha1.New, []byte(s.Secret))
	mac.Write([]byte(b.String()))
	return strings.ToUpper(hex.EncodeToString(mac.Sum(nil)))
}

// NoopSigner returns an empty signature. Selected with STUB_SIGN=off on both
// sides when debugging, so a signature mismatch can be ruled out quickly.
type NoopSigner struct{}

func (NoopSigner) Sign(string, url.Values) string { return "" }

// Verify reports whether the signature carried in params matches. The stub uses
// it; our client never does.
func Verify(s Signer, signPath string, params url.Values) bool {
	if _, ok := s.(NoopSigner); ok {
		return true
	}
	want := s.Sign(signPath, params)
	got := params.Get(SignatureParam)
	return hmac.Equal([]byte(strings.ToUpper(got)), []byte(want))
}
