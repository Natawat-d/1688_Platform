package stub

import (
	"html"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"marketplace/internal/ali"
)

// The cashier. alibaba.alipay.url.get hands back a link and the real one opens
// 1688's payment page; ours opens this, which lists the orders and has a button
// that actually pays them. That closes the loop for a demo: place an order, get
// a pay url, click it, watch the push messages arrive.
//
// INVENTED: 1688 documents the shape of the cashier URL and nothing about the
// page behind it. Everything below the URL is ours.

func (s *Server) serveCashier(w http.ResponseWriter, r *http.Request) {
	ids := parseCashierIDs(rawQueryParam(r.URL.RawQuery, "orderId"))
	if len(ids) == 0 {
		http.Error(w, "orderId is required, e.g. /cashier?orderId=1;2", http.StatusBadRequest)
		return
	}

	var rows strings.Builder
	var total ali.Fen
	payable := 0
	for _, id := range ids {
		o, ok := s.store.Get(id)
		if !ok {
			rows.WriteString(`<tr class="missing"><td colspan="4">订单 ` + strconv.FormatInt(id, 10) + ` 不存在 · order not found</td></tr>`)
			continue
		}
		if o.Status == StatusWaitPay {
			payable++
			total += o.Total
		}
		var items []string
		for _, l := range o.Lines {
			items = append(items, html.EscapeString(l.Name)+" × "+strconv.FormatInt(l.Quantity, 10))
		}
		rows.WriteString(`<tr>` +
			`<td class="mono">` + strconv.FormatInt(o.ID, 10) + `</td>` +
			`<td>` + strings.Join(items, "<br>") + `</td>` +
			`<td><span class="pill s-` + html.EscapeString(o.Status) + `">` + html.EscapeString(statusText(o.Status)) + `</span></td>` +
			`<td class="mono right">` + o.Total.Text() + `</td>` +
			`</tr>`)
	}

	var form string
	if payable > 0 {
		var hidden strings.Builder
		for _, id := range ids {
			hidden.WriteString(`<input type="hidden" name="orderId" value="` + strconv.FormatInt(id, 10) + `">`)
		}
		form = `<form method="post" action="/cashier/pay">` + hidden.String() +
			`<button type="submit">立即支付 ` + total.Text() + ` · Pay now</button></form>`
	} else {
		form = `<p class="done">这些订单无需支付 · nothing here is awaiting payment.</p>`
	}

	writeHTML(w, "收银台 · Cashier", `
<h1>收银台 <span class="sub">Cashier</span></h1>
<table>
  <thead><tr><th>订单号 Order</th><th>商品 Items</th><th>状态 Status</th><th class="right">金额 Amount</th></tr></thead>
  <tbody>`+rows.String()+`</tbody>
</table>
`+form+`
<p class="note">This is the stub gateway's cashier. The real one lives at
trade.1688.com; this one exists so a demo can complete a payment without leaving
the machine.</p>`)
}

func (s *Server) serveCashierPay(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "bad form", http.StatusBadRequest)
		return
	}
	var paid, failed []string
	for _, raw := range r.PostForm["orderId"] {
		id, ok := parseInt64(raw)
		if !ok {
			continue
		}
		events, err := s.store.Pay(id)
		if err != nil {
			failed = append(failed, raw)
			continue
		}
		s.Publish(r.Context(), events)
		paid = append(paid, raw)
	}

	body := `<h1>支付完成 <span class="sub">Paid</span></h1>`
	if len(paid) > 0 {
		body += `<p class="done">已支付 · paid: <span class="mono">` + html.EscapeString(strings.Join(paid, ", ")) + `</span></p>`
	}
	if len(failed) > 0 {
		body += `<p class="warn">未处理 · not payable: <span class="mono">` + html.EscapeString(strings.Join(failed, ", ")) + `</span></p>`
	}
	body += `<p class="note">The order has moved to waitsellersend and an
ORDER_BUYER_VIEW_ORDER_PAY message has been published.</p>`
	writeHTML(w, "支付完成 · Paid", body)
}

// rawQueryParam reads one query parameter out of the RAW query string.
//
// r.URL.Query() cannot be used here. Since Go 1.17 the standard query parser
// treats a semicolon as invalid and silently DROPS the whole parameter that
// contains one — and the documented 1688 cashier url is
// cashier.htm?orderId=a;b. Reaching for Query() would make every multi-order
// pay link look like a missing parameter.
func rawQueryParam(rawQuery, key string) string {
	for _, pair := range strings.Split(rawQuery, "&") {
		k, v, _ := strings.Cut(pair, "=")
		if k != key {
			continue
		}
		if dec, err := url.QueryUnescape(v); err == nil {
			return dec
		}
		return v
	}
	return ""
}

// parseCashierIDs reads the documented semicolon-separated orderId list.
func parseCashierIDs(raw string) []int64 {
	var out []int64
	for _, part := range strings.FieldsFunc(raw, func(r rune) bool { return r == ';' || r == ',' }) {
		if id, ok := parseInt64(part); ok {
			out = append(out, id)
		}
	}
	return out
}

func writeHTML(w http.ResponseWriter, title, body string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(`<!doctype html><html lang="zh"><head><meta charset="utf-8">
<meta name="viewport" content="width=device-width,initial-scale=1">
<title>` + html.EscapeString(title) + `</title>
<style>
:root{color-scheme:light dark;--bg:#fbfbfa;--fg:#1a1a19;--mut:#6b6b66;--line:#e3e3df;--card:#fff;--acc:#c2410c}
@media (prefers-color-scheme:dark){:root{--bg:#191918;--fg:#eeeeec;--mut:#a1a19a;--line:#33332f;--card:#222220;--acc:#fb923c}}
*{box-sizing:border-box}
body{margin:0;background:var(--bg);color:var(--fg);font:15px/1.55 system-ui,-apple-system,"Segoe UI",Helvetica,Arial,sans-serif}
main{max-width:760px;margin:0 auto;padding:40px 20px 80px}
h1{font-size:26px;margin:0 0 24px;font-weight:650}
.sub{color:var(--mut);font-weight:400;font-size:17px}
table{width:100%;border-collapse:collapse;background:var(--card);border:1px solid var(--line);border-radius:10px;overflow:hidden}
th,td{padding:11px 13px;text-align:left;border-bottom:1px solid var(--line);vertical-align:top}
th{font-size:12px;letter-spacing:.04em;text-transform:uppercase;color:var(--mut);font-weight:600}
tr:last-child td{border-bottom:none}
.right{text-align:right}
.mono{font-family:ui-monospace,SFMono-Regular,Menlo,Consolas,monospace;font-size:13px}
.pill{display:inline-block;padding:2px 9px;border-radius:99px;background:var(--line);font-size:12px}
.s-waitbuyerpay{background:#fed7aa;color:#7c2d12}
.s-success,.s-confirm_goods{background:#bbf7d0;color:#14532d}
.s-cancel{background:#e5e5e5;color:#404040}
button{margin-top:22px;padding:12px 22px;font-size:16px;font-weight:600;color:#fff;background:var(--acc);border:0;border-radius:9px;cursor:pointer}
button:hover{filter:brightness(1.08)}
.note{color:var(--mut);font-size:13px;margin-top:32px}
.done{color:#15803d}.warn{color:#b45309}.missing td{color:var(--mut)}
</style></head><body><main>` + body + `</main></body></html>`))
}
