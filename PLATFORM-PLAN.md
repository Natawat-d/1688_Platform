# Marketplace-on-1688 Platform Plan

A Shopee/Lazada-style storefront whose catalogue is pulled from 1688.com through the
"XunYuanTong cross-border key-account sourcing" solution (solution key 1703167397970).
Customers browse and buy on our platform; the platform re-orders on 1688 through one
admin-owned 1688 account, adds a configurable platform fee, and shows parcel progress.

API references used below are documented locally in `1688-api-docs/` (English in `en/`,
originals in `zh/`, push-message topics in `messages-en/`).

---

## 1. What the platform does

| Capability | Where it lives | 1688 APIs / messages involved |
|---|---|---|
| Browse / search / filter goods (queryable catalogue) | Storefront + Catalog service | `product.search.keywordQuery`, `product.search.imageQuery`, `product.search.keywordSNQuery`, `product.search.offerRecommend`, `product.related.recommend`, `category.translation.getById` / `getByKeyword`, `product.topList.query`, `product.search.topKeyword` |
| Product page with SKUs, tiered prices, stock, MOQ, shipping info | Catalog service (synced copy) | `product.search.queryProductDetail` (translated title `subjectTrans`, images, `productSkuInfos`, `productSaleInfo`, `productShippingInfo`, `minOrderQuantity`, `sellerMixSetting`, `status`) |
| Sell price = 1688 price + freight estimate + platform fee (adjustable) | Pricing engine | `product.freight.estimate`, `coupon.optimal.claim` (optional), `alibaba.trade.OpQueryMarketingMixConfig` |
| Customer checks out and pays us | Checkout + local payment gateway | none (our PSP) |
| Platform places the real order on 1688 under the admin account | Order relay worker | `alibaba.createOrder.preview` → `alibaba.trade.createCrossOrder` → `alibaba.alipay.url.get` or `alibaba.trade.pay.protocolPay.preparePay`, `alibaba.trade.payWay.query`, `alibaba.trade.cancel` |
| Order status mirrored to the customer | Order sync (push + poll) | Messages `ORDER_BUYER_VIEW_*`, `ORDER_BATCH_PAY`; polling `alibaba.trade.getBuyerOrderList`, `alibaba.trade.get.buyerView` |
| Parcel tracking (China leg) | Tracking service | Message `LOGISTICS_BUYER_VIEW_TRACE`, `LOGISTICS_MAIL_NO_CHANGE`; `alibaba.trade.getLogisticsInfos.buyerView`, `alibaba.trade.getLogisticsTraceInfo.buyerView`, `logistics.delivery.urge`, `alibaba.logistics.OpQueryLogisticCompanyList` |
| Parcel tracking (international leg) | Tracking service | forwarder / carrier API (outside 1688), or `crossPackageFulfillmentDTO` when 1688 official cross-border logistics is used |
| Returns and refunds | After-sales module | `alibaba.trade.getRefundReasonList`, `alibaba.trade.createRefund`, `alibaba.trade.uploadRefundVoucher`, `alibaba.trade.refund.returnGoods`, `alibaba.trade.refund.OpQueryOrderRefund`, `refund.address.get`, `ORDER_BUYER_VIEW_ORDER_*REFUND*` messages |
| Keep catalogue fresh (price, stock, delisting) | Catalog sync worker | Messages `PRODUCT_RELATION_VIEW_PRODUCT_CHANGE`, `PRODUCT_PRODUCT_INVENTORY_CHANGE`, `PRODUCT_RELATION_VIEW_PRODUCT_EXPIRE`, `FENXIAO_PRICE_CHANGE`; re-fetch with `queryProductDetail` |
| Connect 1688 via admin account | Admin → Integrations | OAuth2 authorization of the admin's 1688 account (`access_token`, refresh), `alibaba.account.basic`, `alibaba.subAccount.list` |
| Reliable message delivery | Integration layer | `push.cursor.messageList`, `push.query.messageList`, `push.message.confirm` |

Not in scope for v1: invoicing APIs, 88 ShengYiTong contracts, Newton Cloud inquiry agent,
fully-managed (Buffalo) consignment, WB listing callbacks. They can be added later.

---

## 2. Architecture

```
Customers ──> Storefront (web/mobile) ──┐
                                        ├──> API / BFF ──> Postgres (catalogue, orders, fees)
Admin ──────> Admin console ────────────┘         │            Redis (cache, queues)
                                                  │            Search index (Postgres FTS or Meilisearch)
                                                  ▼
                                   1688 Integration Layer
                                   ├─ Gateway client (signing, token refresh, rate limiting, retries)
                                   ├─ Catalog sync worker      (search/detail → products/skus/prices)
                                   ├─ Order relay worker       (preview → create → pay → confirm)
                                   ├─ Order/Tracking sync      (push messages + scheduled polling)
                                   └─ Message receiver (HTTPS)  (1688 push → queue → handlers)
                                                  │
                                                  ▼
                                   gw.open.1688.com (API)  ·  1688 message push  ·  Forwarder/carrier API
```

**Recommended stack.** Next.js (App Router) for storefront and admin, TypeScript API
(Next route handlers or a NestJS/Hono service), Postgres, Redis + BullMQ for queues and
scheduled jobs, S3-compatible object storage for cached product images, Meilisearch if
Postgres full-text search proves too slow past ~200k SKUs. Any equivalent stack works; the
integration layer is the only 1688-specific part.

### 2.1 1688 gateway client

* Request URL: `https://gw.open.1688.com/openapi/param2/{version}/{namespace}/{name}/{appKey}`
  (POST form-encoded).
* System params on every call: `access_token`, `_aop_signature`, optional `_aop_timestamp`.
* Signature: HMAC-SHA1 over `param2/{version}/{namespace}/{name}/{appKey}` followed by all
  request params concatenated as `key+value` in sorted key order, keyed by the app secret,
  hex upper-case. Verify against the official invocation guide
  (https://open.1688.com/doc/apiInvoke.htm) before relying on it.
* Auth: admin authorizes once through the 1688 OAuth flow
  (https://open.1688.com/doc/apiAuth.htm); store `access_token` + `refresh_token`, refresh
  ahead of expiry, alert admin if refresh fails. All calls run as this one account.
* Complex params are JSON strings (for example `cargoParamList`, `addressParam`).
* Central rate limiter + retry with backoff; log every call (API, latency, code, message).
* Response envelope is usually `{success, code, message, result}`; some older APIs return
  `{errorCode, errorMessage, success, ...}`. Normalise both.

### 2.2 Catalogue

**Ingestion.** Admin picks what to import: by keyword search, by 1688 category, by
XunYuanTong assortment (`productCollectionId`), by ranking list, or by pasting offer IDs.
Each hit is fetched with `product.search.queryProductDetail` (`country=en`, optionally
`currency`) and stored:

```
products        offer_id (PK), supplier_open_id, title_en (subjectTrans), title_zh,
                description_en (descriptionTrans), images[], white_image, video,
                category_id (1688), local_category_id, status (1688 status),
                moq (minOrderQuantity), batch_number, mix_setting, sold_out, trade_score,
                shipping_info (weight/size), tags[], raw_json, synced_at, visible (bool)
product_skus    sku_id, offer_id, spec_id, attributes[], price_cny, stock, sku_image
price_tiers     offer_id, min_qty, price_cny   (from productSaleInfo quantity/price ranges)
categories      1688 category tree (category.translation.getById from root 0) + local mapping
```

Keep `status` from 1688: only `published` items are visible. Everything else is hidden
automatically.

**Freshness.** Subscribe to the product message topics and re-fetch the affected offer.
Fallback: nightly re-fetch of every visible product (details API), and re-fetch on every
product-page view if the copy is older than N hours (configurable).

**Query layer.** The storefront queries our own DB, never 1688 live, so search and filters are
fast and deterministic: full-text on title, filter by local category, price range (sell price),
MOQ, supplier tags (`offerIdentities`), sort by price / sales / newest. Live 1688 search
(`keywordQuery`) is exposed only in the admin import screen and, optionally, as a
"more from 1688" fallback when local results are empty.

### 2.3 Pricing and platform fee

Sell price is computed by a pure function and recorded with every order line so fee changes
never rewrite history:

```
base_cny        = SKU price for the ordered quantity tier
china_freight   = product.freight.estimate (to our China consolidation address) / units, cached per offer
intl_shipping   = platform's own rate card (weight/volume from productShippingInfo)
subtotal_cny    = base_cny + china_freight + intl_shipping
platform_fee    = subtotal_cny * fee_pct + fee_fixed        (rule chosen by priority)
sell_price      = round((subtotal_cny + platform_fee) * fx_rate(CNY→local), rounding rule)
```

`fee_rules` table: scope (global / category / supplier / product), `fee_pct`, `fee_fixed`,
`min_fee`, `effective_from`, `effective_to`, `priority`; edited in the admin console. Every order
line stores `fee_rule_id`, `fee_amount`, `fx_rate`, and the 1688 price snapshot used.

### 2.4 Checkout and order relay

1. **Cart** groups lines by 1688 supplier (`sellerOpenId`); each group becomes one 1688 order.
   `createCrossOrder` accepts at most 50 SKUs from one supplier, so groups are split at 50.
2. **Validation before payment**: MOQ (`minOrderQuantity`, `batchNumber`, mixed-batch rules),
   stock, item still `published`, price snapshot not older than the configured window.
   Call `alibaba.createOrder.preview` per group to confirm 1688-side price, promotions and
   allowed `tradeModeNameList`; if the previewed total differs from our snapshot beyond a
   tolerance, refresh and ask the customer to confirm.
3. **Customer pays us** through the local PSP. Our order is `PAID`.
4. **Relay worker** (queue, idempotent by our order id):
   * `alibaba.trade.createCrossOrder` with `flow` from the preview, `outOrderId` = our order
     id (1688 uses it for idempotency and it is searchable in `getBuyerOrderList`),
     `addressParam` = our China consolidation warehouse address (or HK/MO/TW address with a
     `crossBorderLogisticsSolutionId` if using 1688 official cross-border logistics),
     `isvBizType=cross`, `message` = our order reference.
   * Store returned 1688 order id(s) in `supplier_orders`. Some calls return several ids.
   * **Pay the 1688 order.** Two modes, admin-selectable:
     - Automatic: `alibaba.trade.pay.protocolPay.isopen` then `preparePay` (needs Cheng-e-She
       or Alipay password-free payment enabled on the admin account; retry at most 3 times per
       the doc).
     - Manual: `alibaba.alipay.url.get` for a cashier URL; admin console shows a "Pay on 1688"
       queue; status flips when `ORDER_BUYER_VIEW_ORDER_PAY` arrives.
   * On failure (out of stock, price change, supplier closed): mark `RELAY_FAILED`, notify
     admin, offer customer refund or substitution.
5. **Cancellation**: customer cancel before 1688 payment → `alibaba.trade.cancel`; after
   payment → after-sales flow (refund request).

### 2.5 Order and parcel tracking

**Status model** (our order shows the customer-facing status; each supplier order keeps the
raw 1688 status):

| 1688 `status` | Our supplier-order status | Customer sees |
|---|---|---|
| `waitbuyerpay` | RELAYED_UNPAID | Processing |
| `waitsellersend` | PAID_TO_SUPPLIER | Preparing |
| `waitbuyerreceive` | SHIPPED_IN_CHINA | Shipped (China leg) |
| `confirm_goods` / `success` | ARRIVED_WAREHOUSE (after our warehouse scan) | At consolidation warehouse |
| `cancel` / `terminated` | CANCELLED | Cancelled |
| our forwarder events | INTL_IN_TRANSIT / DELIVERED | In transit / Delivered |

**Sources**, in priority order:

1. Push messages (`ORDER_BUYER_VIEW_ANNOUNCE_SENDGOODS`, `..._PART_PART_SENDGOODS`,
   `LOGISTICS_BUYER_VIEW_TRACE` with `statusChanged` CONSIGN/ACCEPT/TRANSPORT/DELIVERING/SIGN,
   `LOGISTICS_MAIL_NO_CHANGE`). Messages are asynchronous and may arrive out of order; apply
   only if the event time is newer than what we hold.
2. Scheduled polling every 15–30 min for open supplier orders: `getBuyerOrderList`
   (`modifyStartTime`, `outOrderId`) and `getLogisticsTraceInfo.buyerView(orderId,
   webSite=1688)` for step-level trace (`logisticsSteps[].acceptTime/remark`).
3. Gap fill: `push.cursor.messageList` on startup / after downtime, then
   `push.message.confirm` when using the query-style API.

**Timeline for customers**: China-leg steps from 1688 trace, then warehouse inbound scan,
then international carrier events (forwarder API or manual tracking-number entry by admin).
Admin can trigger `logistics.delivery.urge` (once per 24 h) on stalled orders.

### 2.6 Returns and refunds

Customer opens a request on our platform → admin decides whether to relay to 1688:
`getRefundReasonList` → `createRefund` (in-sale or after-sales) → optional
`uploadRefundVoucher` → when seller agrees, `refund.address.get` and `refund.returnGoods` with
the logistics company code from `OpQueryLogisticCompanyList`. Track with
`OpQueryOrderRefund` and the refund messages; refund the customer from our PSP when 1688
refund reaches `refundsuccess` (or earlier, per policy).

### 2.7 Admin console

* Integrations: connect 1688 account (OAuth), token health, message-push endpoint status,
  API call log and error rate.
* Catalogue: import (search / category / assortment / IDs), visibility, local categories,
  translation overrides, sync status, delisted items report.
* Pricing: fee rules editor with preview ("what will this product cost"), FX rates
  (manual or provider), international rate card, rounding.
* Orders: relay queue, pay-on-1688 queue, failures, manual re-relay, cancel, urge shipment,
  attach international tracking number.
* After-sales: refund requests, evidence upload, status.
* Reports: GMV, fee revenue, margin per order, supplier performance (`tradeScore`, on-time
  shipping), stock-out rate.

---

## 3. Data model (core tables)

```
users, addresses
products, product_skus, price_tiers, categories, category_map, product_sync_log
fee_rules, fx_rates, shipping_rate_cards
carts, cart_items
orders (id, user_id, status, currency, totals, fee_total, fx_rate, paid_at)
order_items (order_id, offer_id, sku_id, qty, base_cny, freight_cny, fee_amount, fee_rule_id, sell_price, snapshot_json)
supplier_orders (id, order_id, seller_open_id, cbu_order_id, out_order_id, status_1688, refund_status, paid_1688_at, raw_json)
shipments (supplier_order_id, logistics_id, mail_no, cp_code, leg: china|intl, status)
tracking_events (shipment_id, source: message|poll|forwarder, event_time, status, remark)
refund_requests (order_item_id, cbu_refund_id, status, reason, evidence[])
integration_tokens (provider, access_token, refresh_token, expires_at, member_id)
message_events (topic, msg_id, payload, received_at, processed_at, status)   -- idempotency
api_call_log
```

---

## 4. Delivery phases

| Phase | Outcome | Main work |
|---|---|---|
| 0. Access (1–2 wks) | App key/secret, solution ordered, permissions granted, admin account authorized, sandbox calls succeed | Apply through the solution page order link; join the business DingTalk group for permissions (per the solution page); implement gateway client + signing + token refresh; smoke-test `alibaba.account.basic`, `keywordQuery`, `queryProductDetail` |
| 1. Catalogue (2–3 wks) | Import products, browse/search storefront | Sync worker, schema, search index, category tree, image caching, product page |
| 2. Pricing & checkout (2 wks) | Fee rules, FX, cart, PSP payment | Pricing engine, fee admin UI, cart grouping by supplier, MOQ validation, `createOrder.preview` check |
| 3. Order relay (2–3 wks) | Paid orders become 1688 orders and get paid | Relay worker, `createCrossOrder`, payment mode (auto/manual), failure handling, admin order screens |
| 4. Tracking (2 wks) | Customer sees live progress | Message receiver endpoint, topic handlers, polling jobs, trace timeline, forwarder integration |
| 5. After-sales (1–2 wks) | Refund/return flow end to end | Refund APIs, evidence upload, customer + admin UI |
| 6. Hardening (ongoing) | Production readiness | Rate limiting, retries, monitoring, reconciliation reports (our orders vs `getBuyerOrderList`), backups, load test of search |

Rough total for a first production release: 10–14 weeks for a team of 2–3 developers.

---

## 5. Risks and decisions to make early

1. **Permissions and business approval.** Several APIs need approval by 1688 staff (the
   solution page says permission is granted through the business team's DingTalk group).
   Start this first; it gates everything.
2. **Paying 1688 orders.** Automatic payment needs password-free payment (Cheng-e-She or
   Alipay agreement) on the admin account. Without it, an operator must pay through the
   cashier URL. Decide the operating model and float/credit needed.
3. **Delivery address.** `createCrossOrder` expects a mainland-China address (or HK/MO/TW with
   an official cross-border logistics solution). Plan for a China consolidation warehouse or
   forwarder and its inbound scanning; international tracking comes from that partner.
4. **One supplier per 1688 order, max 50 SKUs.** Customer orders that span suppliers become
   several 1688 orders with separate shipping; show this clearly.
5. **MOQ and mixed-batch rules.** Many offers have `minOrderQuantity` > 1 or mixed-batch
   thresholds. Either enforce MOQ per SKU, or aggregate demand before relaying (changes the
   fulfilment model).
6. **Price and stock drift.** Snapshot prices at checkout, re-verify with the preview API
   before relaying, and subscribe to product/inventory messages.
7. **Rate limits and paid APIs.** `alibaba.product.suggest.crossBorder` is rate-limited;
   `alibaba.product.simple.get` is a paid API. Cache aggressively, poll gently.
8. **Translation quality.** `subjectTrans`/`descriptionTrans` come from 1688's own
   translation; allow admin overrides and consider a second-pass translation for top sellers.
9. **Message push endpoint.** 1688 pushes messages to a public HTTPS endpoint you register;
   it must be highly available and idempotent (dedupe by message id).
10. **Legal/tax.** Cross-border import duties, consumer protection and payment licensing in
    the selling country are outside the API but affect pricing and the refund policy.

---

## 6. First sprint checklist

- [ ] Register the ISV app on open.1688.com; order solution 1703167397970; request permissions.
- [ ] Admin authorizes the 1688 account; store tokens; verify `alibaba.account.basic`.
- [ ] Implement gateway client (signing, retries, logging) with contract tests against
      `1688-api-docs/en/*.md` samples.
- [ ] Import 200 products via `keywordQuery` + `queryProductDetail`; render product pages.
- [ ] Implement fee rules and price preview.
- [ ] Place one real end-to-end test order (small item) through preview → create → pay,
      and watch the `ORDER_BUYER_VIEW_*` and `LOGISTICS_BUYER_VIEW_TRACE` messages arrive.
