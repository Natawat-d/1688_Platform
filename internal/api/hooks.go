package api

import (
	"crypto/hmac"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"

	"marketplace/internal/ali"
)

// PushSignatureHeader carries the HMAC of the delivered body.
//
// INVENTED, like the rest of the push transport: the documentation defines the
// message envelope and says nothing about how a message is delivered. Our stub
// signs the body as if it were a single request parameter, reusing the same
// HMAC-SHA1 the gateway uses for ordinary requests, and this is the other half
// of that agreement.
const PushSignatureHeader = "X-Aop-Signature"

// verifyPushSignature reports whether body carries a valid signature.
//
// With no secret configured it returns true: at go-live the real gateway's
// signing scheme is unknown, and refusing every real message would be worse
// than accepting them. With a secret set, a bad signature is rejected.
func (s *Server) verifyPushSignature(r *http.Request, body []byte) bool {
	if s.PushSecret == "" {
		return true
	}
	got := strings.TrimSpace(r.Header.Get(PushSignatureHeader))
	if got == "" {
		return false
	}
	want := ali.HMACSHA1Signer{Secret: s.PushSecret}.
		Sign("push", url.Values{"pushMessageList": {string(body)}})
	return hmac.Equal([]byte(strings.ToUpper(got)), []byte(want))
}

// receivePush accepts messages from the gateway.
//
// The body shape is deliberately forgiving, because the documentation defines
// the message envelope but says nothing at all about the HTTP call that carries
// it: not the method, not the encoding, not what counts as an acknowledgement.
// Rather than guess narrowly and break on the first real delivery, this accepts
// a bare envelope, a list of envelopes, or either wrapped in the documented
// pushMessageList field, as a JSON body or as a form field.
//
// Authenticity is a separate matter and is checked first, against the shared
// secret. Idempotency, from message id being the primary key of message_events,
// then makes replays and duplicates cost one rejected insert. The replay
// endpoints remain the authoritative recovery path.
func (s *Server) receivePush(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(io.LimitReader(r.Body, 8<<20))
	if err != nil {
		writeErr(w, http.StatusBadRequest, "bad_request", "Could not read the message body.")
		return
	}

	// Idempotency stops replays; it does nothing about forgeries. Without this
	// check anyone who can reach this endpoint can drive an order's state.
	if !s.verifyPushSignature(r, body) {
		s.Log.Warn("rejected an unsigned or wrongly signed push", "bytes", len(body),
			"remote", r.Header.Get("X-Forwarded-For"))
		writeErr(w, http.StatusUnauthorized, "bad_signature", "Signature verification failed.")
		return
	}

	msgs := parsePush(body)
	if len(msgs) == 0 {
		// A form-encoded delivery: look for the payload in the usual fields.
		if err := r.ParseForm(); err == nil {
			for _, field := range []string{"message", "data", "pushMessageList", "body"} {
				if v := r.PostForm.Get(field); v != "" {
					msgs = parsePush([]byte(v))
					if len(msgs) > 0 {
						break
					}
				}
			}
		}
	}
	if len(msgs) == 0 {
		s.Log.Warn("push delivery carried no messages", "bytes", len(body))
		writeJSON(w, http.StatusOK, map[string]any{"isSuccess": true, "applied": 0})
		return
	}

	applied := 0
	for _, m := range msgs {
		if m.MsgID == 0 {
			continue
		}
		if err := s.Orders.ReceiveMessage(r.Context(), m); err != nil {
			// Answering with a failure would make the gateway redeliver, which is
			// what we want for a transient problem. The message is already
			// recorded, so the sweep job will retry it either way.
			s.Log.Warn("push message failed", "msgId", m.MsgID, "topic", m.Topic(), "err", err)
			continue
		}
		applied++
	}

	// isSuccess is the field name the documented confirm endpoint uses, which
	// makes it the least invented choice available for an acknowledgement.
	writeJSON(w, http.StatusOK, map[string]any{"isSuccess": true, "applied": applied})
}

// parsePush pulls messages out of any of the shapes a delivery might take.
func parsePush(body []byte) []ali.PushMessage {
	if len(body) == 0 {
		return nil
	}

	var wrapped struct {
		PushMessageList []ali.PushMessage `json:"pushMessageList"`
		Messages        []ali.PushMessage `json:"messages"`
	}
	if err := json.Unmarshal(body, &wrapped); err == nil {
		if len(wrapped.PushMessageList) > 0 {
			return wrapped.PushMessageList
		}
		if len(wrapped.Messages) > 0 {
			return wrapped.Messages
		}
	}

	var list []ali.PushMessage
	if err := json.Unmarshal(body, &list); err == nil && len(list) > 0 {
		return list
	}

	var one ali.PushMessage
	if err := json.Unmarshal(body, &one); err == nil && one.MsgID != 0 {
		return []ali.PushMessage{one}
	}
	return nil
}
