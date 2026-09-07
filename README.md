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

## Deploying to AWS

```bash
./deploy/deploy.sh            # build, push, deploy, seed — about fifteen minutes
./deploy/deploy.sh status     # URLs and the generated tokens
./deploy/deploy.sh logs       # tail the task
./deploy/deploy.sh destroy    # remove everything
```

One CloudFormation stack in `deploy/cloudformation.yml` holds the lot: an ECS Fargate task
running all three containers as siblings on localhost, exactly as Compose does locally, behind
an Application Load Balancer and CloudFront. Roughly $40 a month; `aws ecs update-service
--desired-count 0` between demos takes it to about $18.

The deploy runs in two passes on purpose. The stub stamps its own public origin into every
product image URL, and those URLs are written into the database when the catalogue is imported.
CloudFront's domain does not exist until the stack does, so the stack is created once to learn
the domain and updated once to hand it back, before anything is imported.

Credentials are generated on first deploy into one Secrets Manager entry and injected as task
secrets, never as plain environment values. `ENV=production` makes the backend refuse to start
if any of them is still a development placeholder.

The storefront is public. The admin console, the settings API and the stub's cashier are
restricted to the deploying machine's public address by a CloudFront function, because behind
CloudFront the load balancer sees only CloudFront addresses and cannot do it. The stub's
`/_control/*` plane is never routed from the internet at all, and is separately guarded by a
token.

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
