package stub

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"marketplace/internal/ali"
)

// The gateway's own validation is driven entirely by ALL-APIS.json rather than
// by anything hand-written here. That is the point: if the docs say a parameter
// is required and our client forgets it, the stub refuses the call, and the
// contract tests notice. Nothing in this file knows what any endpoint does.

// APISpec is one endpoint as the documentation describes it.
type APISpec struct {
	API           string // "com.alibaba.trade:alibaba.trade.cancel:1"
	Namespace     string
	Name          string
	Version       string
	NeedAuth      bool
	NeedSignature bool

	// Required is every top-level body field documented as required. Nested
	// fields (those whose name carries a dot) are not checked: the docs describe
	// them as members of an object parameter, and the object itself is what the
	// caller sends.
	Required []string
	// Objects is every top-level body field whose documented type is not a
	// scalar, i.e. one that must arrive as JSON text.
	Objects map[string]bool
	// SystemRequired is the required system parameters, e.g. _aop_signature.
	SystemRequired []string
}

type rawAPI struct {
	API           string `json:"api"`
	NameEN        string `json:"name_en"`
	Category      string `json:"category"`
	URL           string `json:"url"`
	NeedAuth      bool   `json:"needAuth"`
	NeedSignature bool   `json:"needSignature"`
	SystemParams  []struct {
		Name     string `json:"name"`
		Required bool   `json:"required"`
	} `json:"systemParams"`
	Body []struct {
		Field    string `json:"field"`
		Type     string `json:"type"`
		Required bool   `json:"required"`
	} `json:"body"`
}

// scalarTypes are the documented types that arrive as plain text rather than as
// JSON. Everything else — including every type ending in "[]", java.util.List
// and java.util.Map — has to parse as JSON.
var scalarTypes = map[string]bool{
	"String": true, "java.lang.String": true,
	"Long": true, "java.lang.Long": true,
	"Integer": true, "java.lang.Integer": true, "int": true, "long": true,
	"Boolean": true, "java.lang.Boolean": true, "boolean": true,
	"Double": true, "java.lang.Double": true, "double": true,
	"Float": true, "java.lang.Float": true, "float": true,
	"BigDecimal": true, "java.math.BigDecimal": true,
	"BigInteger": true, "java.math.BigInteger": true,
	"Date": true, "java.util.Date": true,
	"Byte": true, "java.lang.Byte": true,
}

// LoadSpecs reads ALL-APIS.json once at startup. Every call thereafter is
// validated against the map it returns.
func LoadSpecs(docsDir string) (map[string]APISpec, error) {
	path := filepath.Join(docsDir, "ALL-APIS.json")
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("stub: read %s: %w", path, err)
	}
	var raws []rawAPI
	if err := json.Unmarshal(b, &raws); err != nil {
		return nil, fmt.Errorf("stub: parse %s: %w", path, err)
	}

	specs := make(map[string]APISpec, len(raws))
	for _, r := range raws {
		parts := strings.Split(r.API, ":")
		if len(parts) != 3 {
			return nil, fmt.Errorf("stub: %s: malformed api key %q", path, r.API)
		}
		s := APISpec{
			API:           r.API,
			Namespace:     parts[0],
			Name:          parts[1],
			Version:       parts[2],
			NeedAuth:      r.NeedAuth,
			NeedSignature: r.NeedSignature,
			Objects:       map[string]bool{},
		}
		for _, sp := range r.SystemParams {
			if sp.Required {
				s.SystemRequired = append(s.SystemRequired, sp.Name)
			}
		}
		for _, f := range r.Body {
			if strings.Contains(f.Field, ".") {
				continue // a member of an object parameter, not a parameter
			}
			if f.Required {
				s.Required = append(s.Required, f.Field)
			}
			if !scalarTypes[f.Type] {
				s.Objects[f.Field] = true
			}
		}
		specs[r.API] = s
	}
	if len(specs) == 0 {
		return nil, fmt.Errorf("stub: %s described no APIs", path)
	}
	return specs, nil
}

// validate runs the mechanical checks in the documented order: authorization,
// signature, required parameters, then parameter shape. It returns a gateway
// error to send back, or nil when the request may reach its handler.
func (s *Server) validate(spec APISpec, appKey string, form url.Values) *gatewayError {
	if spec.NeedAuth {
		tok := form.Get(ali.TokenParam)
		if tok == "" {
			return &gatewayError{Code: "400", Message: "missing required system parameter: access_token"}
		}
		if s.cfg.AccessToken != "" && tok != s.cfg.AccessToken {
			return &gatewayError{Code: "401", Message: "invalid access_token"}
		}
	}

	if spec.NeedSignature && !strings.EqualFold(s.cfg.Sign, "off") {
		if form.Get(ali.SignatureParam) == "" {
			return &gatewayError{Code: "400", Message: "missing required system parameter: _aop_signature"}
		}
		signPath := "param2/" + spec.Version + "/" + spec.Namespace + "/" + spec.Name + "/" + appKey
		if !ali.Verify(ali.HMACSHA1Signer{Secret: s.cfg.AppSecret}, signPath, form) {
			return &gatewayError{Code: "400", Message: "Invalid signature"}
		}
	}

	for _, f := range spec.Required {
		if strings.TrimSpace(form.Get(f)) == "" {
			return &gatewayError{Code: "400", Message: "missing required parameter: " + f}
		}
	}

	for f := range spec.Objects {
		v := strings.TrimSpace(form.Get(f))
		if v == "" {
			continue // absent is fine unless it was also required, checked above
		}
		if !json.Valid([]byte(v)) {
			return &gatewayError{Code: "400", Message: "parameter " + f + " is not valid JSON"}
		}
	}
	return nil
}

// gatewayError is a transport-level refusal, as opposed to a business refusal.
//
// INVENTED: the gateway's error body for a bad signature, a missing parameter or
// an unknown api is not described anywhere in 1688-api-docs. This shape
// (error_code / error_message / success, HTTP 200) is our own invention. It is
// deliberately NOT one of the documented business envelopes, so that a client
// which mistakes it for one is obviously wrong.
type gatewayError struct {
	Code    string
	Message string
}

func (e *gatewayError) body() map[string]any {
	return map[string]any{
		"error_code":    e.Code,
		"error_message": e.Message,
		"success":       false,
	}
}
