package ali

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// Params is one request's application-level parameters. Values may be strings,
// numbers, booleans, or anything JSON-encodable: objects and arrays are sent as
// JSON strings, which is what the gateway expects.
type Params map[string]any

// CallLog is one gateway call, handed to Client.Log for the api_calls table and
// the admin integration screen. Token and signature are already redacted.
type CallLog struct {
	API      string
	URL      string
	Params   map[string]string
	Status   int
	Duration time.Duration
	Body     string
	Err      error
}

// Client speaks the 1688 protocol. It has no idea whether it is talking to the
// real gateway or to our stub: that is entirely BaseURL.
type Client struct {
	BaseURL string
	AppKey  string
	Signer  Signer
	Token   TokenSource
	HTTP    *http.Client
	Log     func(CallLog)

	// Attempts is how many times a transport failure is retried. Business
	// refusals are never retried here; the job runner owns that policy.
	Attempts int
}

// New builds a client with sensible defaults.
func New(baseURL, appKey, appSecret, accessToken string) *Client {
	return &Client{
		BaseURL:  baseURL,
		AppKey:   appKey,
		Signer:   HMACSHA1Signer{Secret: appSecret},
		Token:    StaticToken(accessToken),
		HTTP:     &http.Client{Timeout: 30 * time.Second},
		Attempts: 3,
	}
}

// Call performs one request and returns the raw response body. Decoding is left
// to the typed wrappers, because the envelope differs per endpoint.
//
// Going live adds a rate limiter here. Against the stub it would only slow tests
// down; against 1688 it is mandatory, and several of these endpoints are
// documented as rate limited or billed per call.
func (c *Client) Call(ctx context.Context, a API, p Params) ([]byte, error) {
	form := url.Values{}
	for k, v := range p {
		if v == nil {
			continue
		}
		s, err := encodeParam(v)
		if err != nil {
			return nil, fmt.Errorf("%s: parameter %s: %w", a.Key(), k, err)
		}
		form.Set(k, s)
	}

	if a.NeedAuth {
		tok, err := c.Token.Token(ctx)
		if err != nil {
			return nil, fmt.Errorf("%s: token: %w", a.Key(), err)
		}
		form.Set(TokenParam, tok)
	}
	if a.NeedTS {
		form.Set(TimestampParam, strconv.FormatInt(time.Now().UnixMilli(), 10))
	}
	if a.NeedSig && c.Signer != nil {
		form.Set(SignatureParam, c.Signer.Sign(a.SignPath(c.AppKey), form))
	}

	endpoint := a.URL(c.BaseURL, c.AppKey)
	attempts := c.Attempts
	if attempts < 1 {
		attempts = 1
	}

	var lastErr error
	for attempt := 0; attempt < attempts; attempt++ {
		if attempt > 0 {
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			case <-time.After(time.Duration(attempt) * 400 * time.Millisecond):
			}
		}

		started := time.Now()
		body, status, err := c.do(ctx, endpoint, form)
		entry := CallLog{
			API:      a.Key(),
			URL:      endpoint,
			Params:   redact(form),
			Status:   status,
			Duration: time.Since(started),
			Body:     truncate(body),
			Err:      err,
		}
		if c.Log != nil {
			c.Log(entry)
		}

		if err != nil {
			lastErr = err
			continue // transport trouble: worth another go
		}
		if status >= 500 {
			lastErr = &APIError{API: a.Key(), Code: strconv.Itoa(status), Message: "gateway error", HTTP: status, Body: truncate(body)}
			continue
		}
		if status >= 400 {
			return nil, &APIError{API: a.Key(), Code: strconv.Itoa(status), Message: http.StatusText(status), HTTP: status, Body: truncate(body)}
		}
		return body, nil
	}
	return nil, fmt.Errorf("%s: %w", a.Key(), lastErr)
}

func (c *Client) do(ctx context.Context, endpoint string, form url.Values) ([]byte, int, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, 0, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")

	client := c.HTTP
	if client == nil {
		client = http.DefaultClient
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(io.LimitReader(resp.Body, 8<<20))
	if err != nil {
		return nil, resp.StatusCode, err
	}
	return body, resp.StatusCode, nil
}

// encodeParam renders one parameter value the way the gateway expects it:
// scalars as their plain text, everything else as a JSON string.
func encodeParam(v any) (string, error) {
	switch t := v.(type) {
	case string:
		return t, nil
	case []byte:
		return string(t), nil
	case json.RawMessage:
		return string(t), nil
	case bool:
		return strconv.FormatBool(t), nil
	case time.Time:
		return FormatCompact(t), nil
	}

	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(rv.Int(), 10), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(rv.Uint(), 10), nil
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(rv.Float(), 'f', -1, 64), nil
	case reflect.String:
		return rv.String(), nil
	}

	b, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// redact copies the form for logging with the secret-bearing fields masked.
func redact(form url.Values) map[string]string {
	out := make(map[string]string, len(form))
	for k, vs := range form {
		v := ""
		if len(vs) > 0 {
			v = vs[0]
		}
		switch k {
		case TokenParam, SignatureParam:
			v = "***"
		}
		if len(v) > 500 {
			v = v[:500] + "…"
		}
		out[k] = v
	}
	return out
}
