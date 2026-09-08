# Marketplace on 1688

A Shopee/Lazada-style storefront whose catalogue mirrors goods from 1688.com. Shoppers
browse and buy here in THB; the platform re-orders on 1688 under one admin-owned account,
adds a configurable platform fee, and shows parcel progress back to the shopper.

Real 1688 access needs an ISV app key, an OAuth-authorised account and per-API approval,
none of which exist yet. So this repo ships a **stub gateway** that speaks the documented
1688 protocol: same URL shape, same signed form-encoded requests, the same four response
envelopes, the same field names including their real typos, the same push messages. The
backend only ever speaks real protocol. Going live means pointing one environment variable
at `gw.open.1688.com` and adding credentials.

```
apps/web  (React, Vite)  ──►  apps/api (Go :8787)  ──►  Postgres 17 (:5455)
                                    │  signed 1688 protocol
                                    ▼
                              apps/stubgw (Go :8788)  ── push messages ──►  api /api/hooks/1688
```

## Layout

| Path | What |
|---|---|
| `apps/api` | backend: storefront + admin JSON API, relay worker, tracking |
| `apps/stubgw` | the fake 1688 gateway: generated catalogue, order state machine, logistics simulator, push, fault injection |
| `apps/web` | React storefront and admin console |
| `internal/ali` | the 1688 protocol layer: signing, envelopes, money, ids, typed calls, doc-driven contract tests |
| `internal/stub` | the stub's implementation |
| `internal/store` | Postgres: schema (embedded migrations) and every SQL statement |
| `internal/pricing` | pure quote engine: tiers, MOQ, mixed batch, fee rules, FX, rounding |
| `internal/catalog` | import from the gateway into our tables, reprice |
| `internal/order` | checkout, relay, message handling, polling, timeline |
| `internal/jobs` | background jobs on a Postgres table (`FOR UPDATE SKIP LOCKED`) |
| `internal/api` | HTTP handlers and DTOs |
| `deploy/` | one CloudFormation stack and the script that drives it — see [docs/deployment.md](docs/deployment.md) |
| `docs/` | deployment guide |
| `1688-api-docs/` | the archived 1688 documentation (143 APIs, 39 topics), consumed by the stub and by the contract tests |
| `PLATFORM-PLAN.md` | the architecture plan this implements |
| `scripts/orderflow.py` | drives one order through the whole pipeline; doubles as an end-to-end check |

## Run it

Prerequisites: Docker Desktop (running), Go 1.27, Node 24.

```bash
cp .env.example .env         # defaults work as they are
make db-up                   # Postgres on 5455 (5432 is left alone; other projects use it)
make dev                     # stub :8788, api :8787, vite :5173 in the foreground
```

Then, in another shell:

```bash
make smoke                   # signed round trip against the stub
make seed                    # import the 300 generated products
python3 scripts/orderflow.py # place, relay, pay, ship and track one order
```

Open http://localhost:5173 for the shop and http://localhost:5173/admin (token `dev`) for
fee rules, settings, catalogue import, order and relay queues, logs, and the stub controls.

`make up` runs everything in containers with the web app embedded in the api binary; `make psql`
opens a database shell through the container (no local client needed).

## What the stub enforces

It validates `access_token`, the HMAC signature and every documented required parameter
mechanically from `1688-api-docs/ALL-APIS.json`, so forgetting `country` on a search fails
here exactly as it would on 1688. It refuses previews with the documented `500_00x` codes for
stock, MOQ, mixed batch, multi-seller and zero price; caps an order at 50 SKUs from one
supplier; honours `outOrderId` idempotently; returns `CLOSE_ORDER_TOO_FAST` for a cancel
inside ten seconds; and emits the documented trace nodes `CONSIGN → ACCEPT → TRANSPORT →
DELIVERING → SIGN`.

Knobs (see `.env.example`): `STUB_SPEED`, `STUB_AUTOPAY`, `STUB_CHAOS` (nulls for empty
arrays, unknown fields), `STUB_PUSH_JITTER` (reordered and duplicated messages),
`STUB_CREATE_FLAT` (the unwrapped create response two doc samples show). Setting `PUSH_URL`
empty turns pushes off entirely; an order must still complete through polling and the
documented replay APIs, and that is the acceptance test for the recovery path.

## The API

One Go binary serves everything on port 8787: the JSON API below, and the React app itself
from an embedded filesystem. Anything that is not `/api/*` and has no file extension returns
the app shell, so client-side routes like `/p/900000000175` work on a hard refresh.

Deployment lives in **[docs/deployment.md](docs/deployment.md)**.

### Two conventions that run through every response

**Money is an object, never a number.**

```json
{ "minor": "176100", "currency": "THB", "text": "฿1,761.00" }
```

`minor` is a string of minor units — satang for THB, fen for CNY — and `text` is preformatted
for display. The browser renders `text` and never does currency arithmetic. Nothing in the
system stores money as a float; the pricing engine has a test that fails the build if the word
`float64` appears in it.

**Every identifier is a string.** 1688 order ids reach nineteen digits and `JSON.parse` in a
browser silently corrupts any integer past 2^53. A test marshals every response type and fails
if a bare integer of sixteen digits or more ever appears.

### Storefront

No account is needed. A cart cookie identifies the shopper, and an order is addressed by its
public id plus a URL secret issued at checkout.

| Method | Path | Purpose |
|---|---|---|
| `GET` | `/api/health` | Liveness plus a real database ping. 503 when Postgres is unreachable. |
| `GET` | `/api/categories` | Top-level categories with product counts. |
| `GET` | `/api/products` | Search and browse. |
| `GET` | `/api/products/{offerId}` | One product with SKUs, quantity tiers and shipping. |
| `GET` | `/api/cart` | The cart, priced and grouped by supplier. |
| `POST` | `/api/cart/items` | `{offerId, skuId, quantity}`. Returns the whole cart. |
| `PATCH` | `/api/cart/items/{id}` | `{quantity}`. Zero removes the line. |
| `DELETE` | `/api/cart/items/{id}` | Remove a line. |
| `POST` | `/api/checkout` | `{email, name, phone, address}`. Validates against 1688 before creating anything. |
| `POST` | `/api/orders/{publicId}/pay?t=` | Records payment, which releases the relay. |
| `GET` | `/api/orders/{publicId}?t=` | Status, timeline, parcels and tracking events. |
| `POST` | `/api/orders/{publicId}/cancel?t=` | Requests cancellation of every parcel. |

**`GET /api/products`** takes `q`, `cat`, `min`, `max` (whole baht), `moq`, `instock`, `sort`
(`relevance`, `price_asc`, `price_desc`, `sales`, `newest`), `page` and `size`.

```json
{ "total": 13, "page": 1, "size": 1, "items": [ {
    "offerId": "900000000289",
    "title": "Affordable-luxury PU Leather Shoulder Bag Factory Direct Wholesale",
    "image": "https://…/img/900000000289/0.svg",
    "priceFrom": { "minor": "196400", "currency": "THB", "text": "฿1,964.00" },
    "moq": 5, "unit": "Piece", "monthSold": 10890,
    "seller": { "openId": "YU4432…", "name": "东莞市长安镇精密五金厂", "score": "4.8" } } ] }
```

An English query goes through Postgres full-text search; a query containing Chinese takes a
trigram path instead, because no bundled tokenizer segments Chinese.

**`GET /api/products/{offerId}`** adds `titleZh`, `images`, `skus` (each with `specId`, stock
and its own price), `tiers`, `mix`, `shipping` in grams and millimetres, `priceRange`,
`quoteType` and `sellable`. Prices come from the same engine checkout uses, quoted at the
minimum order quantity, so the shop cannot advertise a price it will not honour.

**`GET /api/cart`** is where the business rules become visible:

```json
{ "groups": [ { "sellerOpenId": "RT6674…", "sellerName": "温州市瓯海区文具制造厂",
      "parcels": 1,
      "lines": [ { "id": "12", "quantity": 5, "stock": 5822,
          "unit": { "minor": "176100", "currency": "THB", "text": "฿1,761.00" },
          "breakdown": { "base": {…}, "freight": {…}, "intl": {…}, "fee": {…},
                         "fxPpm": "4900000", "feeRuleId": "1" } } ],
      "issues": [ { "code": "500_005", "field": "900000000175",
                    "message": "Minimum order is 5 pieces for this product" } ],
      "subtotal": {…} } ],
  "totals": { "goods": {…}, "chinaFreight": {…}, "intl": {…}, "fee": {…}, "total": {…} },
  "checkoutable": false, "count": 5 }
```

Lines group by supplier because each group becomes one 1688 order, split further at fifty
SKUs, so `parcels` tells the shopper how many deliveries to expect **before** paying. The
`issues` codes are 1688's own — `500_004` stock, `500_005` minimum order, `500_006` mixed
batch — so our gate and the gateway's gate speak one vocabulary. `breakdown` is per unit and
in CNY: it is what the storefront's "how this price is made" panel renders.

**`POST /api/checkout`** calls the gateway's order preview for every supplier group **before**
anything is created or charged. A supplier that cannot accept API orders, a price that has
drifted beyond tolerance, or any stock or minimum-order problem comes back as `409` with the
same `issues` shape. On success:

```json
{ "orderId": "MK-2026-000001", "token": "46e4a335…", "parcels": 2,
  "total": {…}, "payUrl": "/order/MK-2026-000001?t=46e4a335…" }
```

Keep `token`: it is the only thing authorising later access to that order.

**`GET /api/orders/{publicId}?t=`** returns totals, items with their price breakdown, a fixed
eight-step `timeline` where each step carries `done`, `at` and sometimes `detail`
("1 of 2 parcels"), and a `parcels` array with the 1688 order id, carrier, tracking number and
every tracking event with its `source` — `message` when a push arrived, `poll` when the
periodic sweep found it. Both paths write the same events and converge on a dedupe key, so
duplicates are impossible and neither is required for correctness.

### Admin

Every route below needs `Authorization: Bearer <ADMIN_TOKEN>`, compared in constant time. On a
deployment they are additionally restricted by source address.

| Method | Path | Purpose |
|---|---|---|
| `GET` `PUT` | `/api/admin/settings` | FX rate, rounding, international rate, warehouse address, tolerances. A change that moves prices enqueues a reprice. |
| `GET` `POST` | `/api/admin/fee-rules` | List and create fee rules. |
| `PUT` `DELETE` | `/api/admin/fee-rules/{id}` | Edit and remove. |
| `POST` | `/api/admin/fee-rules/preview` | `{offerId, skuId, quantity}` → the full breakdown. The fastest way to see what a rule change does. |
| `POST` | `/api/admin/import` | `{all}`, `{keyword, pages}` or `{offerIds}`. Queues a catalogue import. |
| `GET` `PATCH` | `/api/admin/products[/{offerId}]` | Browse everything including hidden items; toggle visibility. |
| `GET` | `/api/admin/orders[/{id}]` | Orders with their supplier orders and raw gateway payloads. |
| `POST` | `/api/admin/supplier-orders/{id}/relay` | Retry a stalled relay or payment. |
| `POST` | `/api/admin/supplier-orders/{id}/cancel` | `{reason, remark}`. Scheduled past the ten-second window 1688 refuses inside. |
| `GET` `POST` | `/api/admin/jobs[/{id}/retry]` | The background queue, including dead jobs. |
| `GET` | `/api/admin/api-calls` | Every gateway call with timing and result. |
| `GET` | `/api/admin/messages` | Every inbound push, whether it applied, and why not. |
| `ANY` | `/api/admin/stub/*` | Proxied to the stub's control plane. 404s when no stub is configured, which is what happens at go-live. |

A fee rule is scoped `global`, `category`, `supplier` or `product`, carries `feeBps`,
`feeFixedFen`, `minFeeFen`, a `priority` and an effective window. The winner is chosen by
priority, then by scope specificity.

### The gateway webhook

`POST /api/hooks/1688` takes push messages. It accepts a bare envelope, a list, or either
wrapped in `pushMessageList`, as JSON or as a form field, because the documentation defines
the envelope and says nothing about the delivery.

Authenticity is checked first: the body must carry `X-Aop-Signature`, an HMAC-SHA1 under the
app secret, or the request is rejected with 401. Then message id, which is the primary key of
`message_events`, makes replays and duplicates cost one rejected insert. A reply of
`{"isSuccess":true}` acknowledges.

### Errors

Failures are `{"error": "<kind>", "message": "<something a person can act on>"}`, plus
`issues` on a 409 from checkout. `400` malformed, `401` bad or missing token, `404` unknown or
wrong URL secret, `409` the cart or the supplier refused, `500` ours, `503` the database is
unreachable.

## Where it is honest about guessing

* **Request signing is reconstructed from notes, not from an official source.** It lives in
  one file, `internal/ali/sign.go`, and the stub validates with the same function, so a wrong
  algorithm passes locally and fails on the first real call. Verifying it against
  https://open.1688.com/doc/apiInvoke.htm is the first task on the day access is granted.
* **OAuth is not documented locally at all.** The token comes from the environment behind
  `ali.TokenSource`.
* **The HTTP mechanics of push are undocumented.** The envelope is documented; the method,
  encoding and acknowledgement are invented and labelled as such. The receiver is lenient and
  idempotent on message id, and the documented replay APIs are the recovery path.

## Going live

Change `ALI_BASE_URL`, `ALI_APP_KEY`, `ALI_APP_SECRET` and the token source; fix `sign.go`
if verification finds it wrong; add a rate limiter in `internal/ali/client.go`; register the
push URL and adapt `internal/api/hooks.go` to the real contract; put the real China
consolidation-warehouse address with its district code in settings; replace the fake
payment endpoint with a real provider; delete `apps/stubgw`, `internal/stub`, the `stub`
Compose service, the `/api/admin/stub/*` proxy and the admin Stub tab. Nothing in the
storefront, the pricing engine or the schema should need to change.

## Pricing policy worth knowing

Quotes use the supplier's **wholesale** price and ignore `promotionPrice` unless the
`use_promo_prices` setting is on. The documentation never says an order preview honours a
promotion, and against the gateway it did not: the preview charged wholesale where we had
quoted the promotion, and that gap would have come out of our margin. Checkout compares our
goods total with the preview's `sumPaymentNoCarriage` and refuses the order when they differ
by more than `price_tolerance_bps`, so a real drift is caught before anyone is charged.

## Known gaps

* Chinese search runs on trigram similarity rather than a real tokenizer; `zhparser` would
  need a custom Postgres image.
* Category names come through as the Chinese `categoryName` the detail API returns; the
  category-translation API exists and could feed English names into the rail.
* No customer accounts: orders are anonymous, addressed by public id plus a URL secret.
* Refunds and returns are designed for but not built.
