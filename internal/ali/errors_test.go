package ali_test

// Every business error code the documentation lists for the endpoints whose
// refusals we classify must appear in ali.BusinessCodes(), so Retryable never
// retries a documented refusal.

import (
	"strings"
	"testing"

	"marketplace/internal/ali"
)

func TestErrorCodesClassified(t *testing.T) {
	dir := docsDir(t)

	known := map[string]bool{}
	for _, c := range ali.BusinessCodes() {
		known[strings.ToLower(c)] = true
	}

	apis := []ali.API{ali.OrderPreview, ali.CreateCrossOrder, ali.TradeCancel, ali.LogisticsTrace}
	total := 0
	for _, api := range apis {
		d := loadRawDoc(t, dir, api)
		if len(d.ErrorCodes) == 0 {
			t.Logf("%s: no apiErrorCodeVOList in the archive", api.DocID())
			continue
		}
		for _, e := range d.ErrorCodes {
			code := strings.TrimSpace(e.Code) // "order.nopermission.buyer\t" is documented with a trailing tab
			if code == "" {
				continue
			}
			// A "code" with spaces is a prose message, not a code: the alipay
			// doc lists "Batch pay : not surport MANUAL-TRADE!" and the create
			// doc lists "not support tradeType:【XXXX】".
			if strings.ContainsAny(code, " \t") {
				t.Logf("%s: documented error %q is prose, not a code; skipped", api.DocID(), code)
				continue
			}
			total++
			if !known[strings.ToLower(code)] {
				t.Errorf("%s: documented error code %q (%s) is not in ali.BusinessCodes()", api.DocID(), code, strings.TrimSpace(e.Desc))
			}
		}
	}
	if total == 0 {
		t.Fatalf("no error codes found for %d APIs; the corpus looks wrong", len(apis))
	}
	t.Logf("%d documented codes checked against %d classified", total, len(known))
}
