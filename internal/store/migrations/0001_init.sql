-- Marketplace schema.
--
-- Conventions, applied without exception:
--   * All money is BIGINT in minor units. No numeric, no float. Column names carry
--     the unit: _fen is CNY cents, _satang is THB satang.
--   * Rates are integers too: fee_bps is basis points, fx_ppm is THB per CNY times 1e6.
--   * 1688 identifiers are BIGINT (they reach nineteen digits). Our own primary keys
--     are BIGINT GENERATED ALWAYS AS IDENTITY.
--   * Raw gateway payloads are kept in JSONB so a field we did not model is never lost.
--   * Times are timestamptz.

CREATE EXTENSION IF NOT EXISTS pg_trgm;

-- ---------------------------------------------------------------- catalogue --

CREATE TABLE categories (
    id          BIGINT PRIMARY KEY,
    parent_id   BIGINT   NOT NULL DEFAULT 0,
    name        TEXT     NOT NULL DEFAULT '',
    name_trans  TEXT     NOT NULL DEFAULT '',
    level       SMALLINT NOT NULL DEFAULT 1,
    is_leaf     BOOLEAN  NOT NULL DEFAULT false,
    visible     BOOLEAN  NOT NULL DEFAULT true
);
CREATE INDEX categories_parent_idx ON categories (parent_id);

CREATE TABLE products (
    offer_id            BIGINT PRIMARY KEY,
    seller_open_id      TEXT    NOT NULL DEFAULT '',
    company_name        TEXT    NOT NULL DEFAULT '',
    subject             TEXT    NOT NULL DEFAULT '',   -- Chinese title
    subject_trans       TEXT    NOT NULL DEFAULT '',   -- English title, what we display
    description_trans   TEXT    NOT NULL DEFAULT '',
    keywords            TEXT    NOT NULL DEFAULT '',   -- category + attribute values, feeds search_en
    images              TEXT[]  NOT NULL DEFAULT '{}',
    white_image         TEXT    NOT NULL DEFAULT '',
    category_id         BIGINT  NOT NULL DEFAULT 0,
    top_category_id     BIGINT  NOT NULL DEFAULT 0,
    second_category_id  BIGINT  NOT NULL DEFAULT 0,
    third_category_id   BIGINT  NOT NULL DEFAULT 0,
    category_name       TEXT    NOT NULL DEFAULT '',

    -- Raw 1688 lifecycle value. Only 'published' is sellable; the rest are hidden.
    status              TEXT    NOT NULL DEFAULT 'published',
    min_order_quantity  INT     NOT NULL DEFAULT 1,
    batch_number        INT     NOT NULL DEFAULT 0,
    -- 0 = no SKU, quoted by product quantity; 1 = quoted per SKU; 2 = has SKUs but
    -- still quoted by product quantity. Decides which price the engine uses.
    quote_type          SMALLINT NOT NULL DEFAULT 0,
    unit_trans          TEXT    NOT NULL DEFAULT '',

    mix_general         BOOLEAN NOT NULL DEFAULT false,
    mix_amount_fen      BIGINT  NOT NULL DEFAULT 0,
    mix_number          INT     NOT NULL DEFAULT 0,

    amount_on_sale      INT     NOT NULL DEFAULT 0,
    month_sold          INT     NOT NULL DEFAULT 0,
    trade_score         TEXT    NOT NULL DEFAULT '',
    identities          TEXT[]  NOT NULL DEFAULT '{}',

    weight_g            INT     NOT NULL DEFAULT 0,
    length_mm           INT     NOT NULL DEFAULT 0,
    width_mm            INT     NOT NULL DEFAULT 0,
    height_mm           INT     NOT NULL DEFAULT 0,

    -- Cached at import from product.freight.estimate. Freight barely moves and the
    -- estimate needs district codes, so checkout does not call it.
    china_freight_fen   BIGINT  NOT NULL DEFAULT 0,
    freight_free        BOOLEAN NOT NULL DEFAULT false,

    price_min_fen       BIGINT  NOT NULL DEFAULT 0,
    price_max_fen       BIGINT  NOT NULL DEFAULT 0,
    -- Denormalised sell price for grid sorting. Rebuilt by the reprice job whenever a
    -- fee rule or an FX rate changes; a consistency test resamples it.
    sell_min_satang     BIGINT  NOT NULL DEFAULT 0,

    visible             BOOLEAN NOT NULL DEFAULT true,
    raw                 JSONB   NOT NULL DEFAULT '{}',
    synced_at           timestamptz NOT NULL DEFAULT now(),

    -- Latin text search. The two-argument to_tsvector is immutable, so it is allowed
    -- in a generated column and can never go stale.
    search_en tsvector GENERATED ALWAYS AS (
        to_tsvector('english',
            coalesce(subject_trans, '') || ' ' ||
            coalesce(company_name, '')  || ' ' ||
            coalesce(keywords, ''))
    ) STORED
);
CREATE INDEX products_search_en_idx ON products USING GIN (search_en);
-- Chinese does not tokenise under any bundled parser, so the Chinese title is matched
-- by trigram similarity instead. See internal/catalog/search.go for the dispatch.
CREATE INDEX products_zh_trgm_idx ON products USING GIN (subject gin_trgm_ops);
CREATE INDEX products_en_trgm_idx ON products USING GIN (subject_trans gin_trgm_ops);
CREATE INDEX products_browse_idx  ON products (top_category_id, sell_min_satang)
    WHERE visible AND status = 'published';
CREATE INDEX products_sold_idx    ON products (month_sold DESC)
    WHERE visible AND status = 'published';
CREATE INDEX products_seller_idx  ON products (seller_open_id);

CREATE TABLE product_skus (
    sku_id          BIGINT PRIMARY KEY,
    offer_id        BIGINT NOT NULL REFERENCES products(offer_id) ON DELETE CASCADE,
    spec_id         TEXT   NOT NULL DEFAULT '',   -- what order creation needs
    attrs           JSONB  NOT NULL DEFAULT '[]', -- [{name,nameTrans,value,valueTrans,image}]
    label           TEXT   NOT NULL DEFAULT '',   -- "Red / XL", precomputed for display
    price_fen       BIGINT NOT NULL DEFAULT 0,
    promo_price_fen BIGINT NOT NULL DEFAULT 0,    -- 0 means none
    amount_on_sale  INT    NOT NULL DEFAULT 0,
    cargo_number    TEXT   NOT NULL DEFAULT '',
    image_url       TEXT   NOT NULL DEFAULT '',
    weight_g        INT    NOT NULL DEFAULT 0,
    UNIQUE (offer_id, spec_id)
);
CREATE INDEX product_skus_offer_idx ON product_skus (offer_id);

CREATE TABLE price_tiers (
    offer_id        BIGINT NOT NULL REFERENCES products(offer_id) ON DELETE CASCADE,
    start_quantity  INT    NOT NULL,
    price_fen       BIGINT NOT NULL,
    promo_price_fen BIGINT NOT NULL DEFAULT 0,
    PRIMARY KEY (offer_id, start_quantity)
);

-- ------------------------------------------------------------------ pricing --

CREATE TABLE fee_rules (
    id             BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    scope          TEXT   NOT NULL CHECK (scope IN ('global', 'category', 'supplier', 'product')),
    scope_value    TEXT   NOT NULL DEFAULT '',    -- category id, seller open id, offer id, or ''
    fee_bps        INT    NOT NULL DEFAULT 0,     -- basis points: 1500 is 15%
    fee_fixed_fen  BIGINT NOT NULL DEFAULT 0,
    min_fee_fen    BIGINT NOT NULL DEFAULT 0,
    priority       INT    NOT NULL DEFAULT 0,
    effective_from timestamptz NOT NULL DEFAULT now(),
    effective_to   timestamptz,                   -- null means open ended
    note           TEXT   NOT NULL DEFAULT '',
    created_at     timestamptz NOT NULL DEFAULT now(),
    CHECK (scope = 'global' OR scope_value <> '')
);
CREATE INDEX fee_rules_lookup_idx ON fee_rules (scope, scope_value, priority DESC);

CREATE TABLE settings (
    key   TEXT PRIMARY KEY,
    value TEXT NOT NULL
);

-- ------------------------------------------------------------------ commerce --

CREATE TABLE carts (
    id         TEXT PRIMARY KEY,                  -- 128-bit hex, also the cookie value
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE cart_items (
    id       BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    cart_id  TEXT   NOT NULL REFERENCES carts(id) ON DELETE CASCADE,
    offer_id BIGINT NOT NULL,
    sku_id   BIGINT NOT NULL,
    spec_id  TEXT   NOT NULL DEFAULT '',
    quantity INT    NOT NULL CHECK (quantity > 0),
    added_at timestamptz NOT NULL DEFAULT now(),
    UNIQUE (cart_id, sku_id)
);
CREATE INDEX cart_items_cart_idx ON cart_items (cart_id);

-- Customer-facing order numbers come from their own sequence so they stay short
-- and readable, independent of the surrogate key.
CREATE SEQUENCE order_public_seq START 1;

CREATE TABLE orders (
    id             BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    public_id      TEXT NOT NULL UNIQUE,          -- MK-2026-000123; also the outOrderId prefix
    access_token   TEXT NOT NULL,                 -- URL secret; anonymous orders would be enumerable without it
    cart_id        TEXT NOT NULL DEFAULT '',
    email          TEXT NOT NULL DEFAULT '',
    ship_name      TEXT NOT NULL DEFAULT '',
    ship_phone     TEXT NOT NULL DEFAULT '',
    ship_address   JSONB NOT NULL DEFAULT '{}',   -- the customer's own address, display only in v1
    status         TEXT NOT NULL DEFAULT 'NEW',
    fx_ppm         BIGINT NOT NULL DEFAULT 0,     -- snapshot; history is never recomputed
    goods_satang   BIGINT NOT NULL DEFAULT 0,
    freight_satang BIGINT NOT NULL DEFAULT 0,
    intl_satang    BIGINT NOT NULL DEFAULT 0,
    fee_satang     BIGINT NOT NULL DEFAULT 0,
    total_satang   BIGINT NOT NULL DEFAULT 0,
    created_at     timestamptz NOT NULL DEFAULT now(),
    paid_at        timestamptz,
    updated_at     timestamptz NOT NULL DEFAULT now()
);
CREATE INDEX orders_status_idx ON orders (status, updated_at DESC);

CREATE TABLE supplier_orders (
    id              BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    order_id        BIGINT NOT NULL REFERENCES orders(id) ON DELETE CASCADE,
    seller_open_id  TEXT   NOT NULL DEFAULT '',
    seller_name     TEXT   NOT NULL DEFAULT '',
    group_seq       INT    NOT NULL DEFAULT 1,
    -- The idempotency key we hand 1688 as outOrderId, and the one we search on when a
    -- create times out and we need to find out whether it actually landed.
    out_order_id    TEXT   NOT NULL UNIQUE,
    cbu_order_id    BIGINT UNIQUE,                -- 1688's own order id
    flow            TEXT   NOT NULL DEFAULT 'general',
    trade_type      TEXT   NOT NULL DEFAULT '',
    status          TEXT   NOT NULL DEFAULT 'PENDING',
    status_1688     TEXT   NOT NULL DEFAULT '',
    -- Guards against out-of-order push messages: an event is applied only when it is
    -- more advanced, or equally advanced and newer.
    status_rank     SMALLINT NOT NULL DEFAULT 0,
    status_at       timestamptz NOT NULL DEFAULT 'epoch',
    pay_url         TEXT   NOT NULL DEFAULT '',
    sum_payment_fen BIGINT NOT NULL DEFAULT 0,
    post_fee_fen    BIGINT NOT NULL DEFAULT 0,
    error_code      TEXT   NOT NULL DEFAULT '',
    error_message   TEXT   NOT NULL DEFAULT '',
    raw             JSONB  NOT NULL DEFAULT '{}',
    created_at      timestamptz NOT NULL DEFAULT now(),
    updated_at      timestamptz NOT NULL DEFAULT now(),
    UNIQUE (order_id, group_seq)
);
CREATE INDEX supplier_orders_open_idx ON supplier_orders (status, updated_at)
    WHERE status NOT IN ('ARRIVED_WAREHOUSE', 'CANCELLED', 'FAILED');
CREATE INDEX supplier_orders_order_idx ON supplier_orders (order_id);

CREATE TABLE order_items (
    id                BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    order_id          BIGINT NOT NULL REFERENCES orders(id) ON DELETE CASCADE,
    supplier_order_id BIGINT REFERENCES supplier_orders(id) ON DELETE SET NULL,
    offer_id          BIGINT NOT NULL,
    sku_id            BIGINT NOT NULL,
    spec_id           TEXT   NOT NULL DEFAULT '',
    title             TEXT   NOT NULL DEFAULT '',
    image_url         TEXT   NOT NULL DEFAULT '',
    sku_label         TEXT   NOT NULL DEFAULT '',
    quantity          INT    NOT NULL,
    -- The price snapshot. A later fee or FX change must never rewrite history.
    base_fen          BIGINT NOT NULL DEFAULT 0,
    freight_fen       BIGINT NOT NULL DEFAULT 0,
    intl_fen          BIGINT NOT NULL DEFAULT 0,
    fee_fen           BIGINT NOT NULL DEFAULT 0,
    fee_rule_id       BIGINT REFERENCES fee_rules(id),
    fx_ppm            BIGINT NOT NULL DEFAULT 0,
    unit_satang       BIGINT NOT NULL DEFAULT 0,
    line_satang       BIGINT NOT NULL DEFAULT 0,
    snapshot          JSONB  NOT NULL DEFAULT '{}',
    seller_open_id    TEXT   NOT NULL DEFAULT ''
);
CREATE INDEX order_items_order_idx    ON order_items (order_id);
CREATE INDEX order_items_supplier_idx ON order_items (supplier_order_id);

-- ------------------------------------------------------------------ tracking --

CREATE TABLE shipments (
    id                BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    supplier_order_id BIGINT NOT NULL REFERENCES supplier_orders(id) ON DELETE CASCADE,
    leg               TEXT   NOT NULL DEFAULT 'china' CHECK (leg IN ('china', 'intl')),
    logistics_id      TEXT   NOT NULL DEFAULT '',
    mail_no           TEXT   NOT NULL DEFAULT '',
    cp_code           TEXT   NOT NULL DEFAULT '',
    company_name      TEXT   NOT NULL DEFAULT '',
    status            TEXT   NOT NULL DEFAULT '',
    updated_at        timestamptz NOT NULL DEFAULT now(),
    UNIQUE (supplier_order_id, leg, logistics_id)
);

CREATE TABLE tracking_events (
    id          BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    shipment_id BIGINT NOT NULL REFERENCES shipments(id) ON DELETE CASCADE,
    source      TEXT   NOT NULL CHECK (source IN ('message', 'poll', 'manual')),
    event_at    timestamptz NOT NULL,
    code        TEXT   NOT NULL DEFAULT '',
    remark      TEXT   NOT NULL DEFAULT '',
    -- Push and polling both report the same physical event. Deduping on a hash of
    -- time and text lets the two converge with no ordering logic at all.
    dedupe_key  TEXT   NOT NULL,
    UNIQUE (shipment_id, dedupe_key)
);
CREATE INDEX tracking_events_ship_idx ON tracking_events (shipment_id, event_at);

-- ------------------------------------------------------------------ plumbing --

CREATE TABLE message_events (
    msg_id       BIGINT PRIMARY KEY,             -- 1688 msgId; the primary key IS the dedupe
    type         TEXT   NOT NULL DEFAULT '',
    gmt_born     timestamptz NOT NULL DEFAULT now(),
    payload      JSONB  NOT NULL DEFAULT '{}',
    received_at  timestamptz NOT NULL DEFAULT now(),
    processed_at timestamptz,
    error        TEXT   NOT NULL DEFAULT ''
);
CREATE INDEX message_events_pending_idx ON message_events (received_at) WHERE processed_at IS NULL;

CREATE TABLE jobs (
    id         BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    kind       TEXT   NOT NULL,
    key        TEXT   NOT NULL DEFAULT '',
    payload    JSONB  NOT NULL DEFAULT '{}',
    state      TEXT   NOT NULL DEFAULT 'pending'
               CHECK (state IN ('pending', 'running', 'done', 'dead')),
    attempts   INT    NOT NULL DEFAULT 0,
    run_at     timestamptz NOT NULL DEFAULT now(),
    last_error TEXT   NOT NULL DEFAULT '',
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now()
);
-- One scheduled job per (kind, key); superseded by 0002, kept so a fresh database
-- replays history exactly.
CREATE UNIQUE INDEX jobs_key_idx   ON jobs (kind, key) WHERE state IN ('pending', 'running');
CREATE INDEX        jobs_ready_idx ON jobs (run_at) WHERE state = 'pending';

CREATE TABLE api_calls (
    id         BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    api        TEXT   NOT NULL,
    ms         INT    NOT NULL DEFAULT 0,
    ok         BOOLEAN NOT NULL DEFAULT true,
    code       TEXT   NOT NULL DEFAULT '',
    message    TEXT   NOT NULL DEFAULT '',
    req        JSONB  NOT NULL DEFAULT '{}',
    resp       TEXT   NOT NULL DEFAULT '',
    created_at timestamptz NOT NULL DEFAULT now()
);
CREATE INDEX api_calls_time_idx ON api_calls (created_at DESC);

-- --------------------------------------------------------------- seed values --

INSERT INTO settings (key, value) VALUES
    ('fx_thb_per_cny_ppm',      '4900000'),   -- 4.90 THB per CNY
    ('rounding_step_satang',    '100'),       -- price to the whole baht
    ('rounding_mode',           'up'),
    ('intl_rate_satang_per_kg', '18000'),     -- ฿180 per kg
    ('intl_min_satang',         '5000'),      -- ฿50 minimum per line
    ('price_snapshot_ttl_sec',  '3600'),
    ('price_tolerance_bps',     '500'),       -- 5% drift between our snapshot and the preview
    ('payment_mode',            'manual'),    -- manual cashier queue, or auto
    ('warehouse_address',       '{"fullName":"Marketplace Warehouse","mobile":"13800138000","phone":"0571-88888888","postCode":"310052","provinceText":"浙江省","cityText":"杭州市","areaText":"滨江区","townText":"长河街道","address":"网商路699号","districtCode":"330108"}')
ON CONFLICT (key) DO NOTHING;

INSERT INTO fee_rules (scope, scope_value, fee_bps, fee_fixed_fen, min_fee_fen, priority, note)
VALUES ('global', '', 1500, 0, 500, 0, 'Launch margin: 15% with a ¥5 floor')
ON CONFLICT DO NOTHING;
