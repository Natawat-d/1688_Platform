import { useState } from "react";
import { useQuery } from "../lib/useQuery";
import { api, errorText } from "../lib/api";
import { Field, FieldRow } from "../components/Field";
import shell from "./Admin.module.css";
import styles from "./StubControls.module.css";

/**
 * Controls for the stub gateway (cmd/stubgw), proxied through
 * /api/admin/stub/*. These drive the simulated world: how fast its clock runs,
 * where an order sits in its lifecycle, which call should fail next, and which
 * push message to fire.
 *
 * The named forms below assume the paths listed in each panel. The raw console
 * at the bottom talks to any stub path, so nothing here is a dead end if the
 * stub settles on different routes.
 */

const TOPICS = [
  "ORDER_BUYER_VIEW_ORDER_PAY",
  "ORDER_BUYER_VIEW_BUYER_MAKE",
  "ORDER_BUYER_VIEW_ANNOUNCE_SENDGOODS",
  "ORDER_BUYER_VIEW_PART_PART_SENDGOODS",
  "ORDER_BUYER_VIEW_ORDER_COMFIRM_RECEIVEGOODS",
  "ORDER_BUYER_VIEW_ORDER_SUCCESS",
  "ORDER_BUYER_VIEW_ORDER_BUYER_CLOSE",
  "ORDER_BUYER_VIEW_ORDER_SELLER_CLOSE",
  "ORDER_BUYER_VIEW_ORDER_PRICE_MODIFY",
  "ORDER_BATCH_PAY",
  "LOGISTICS_BUYER_VIEW_TRACE",
  "LOGISTICS_MAIL_NO_CHANGE",
  "PRODUCT_PRODUCT_INVENTORY_CHANGE",
  "PRODUCT_RELATION_VIEW_PRODUCT_CHANGE",
  "PRODUCT_RELATION_VIEW_PRODUCT_EXPIRE",
  "FENXIAO_PRICE_CHANGE",
];

const FAULT_APIS = [
  "alibaba.createOrder.preview",
  "alibaba.trade.createCrossOrder",
  "alibaba.alipay.url.get",
  "alibaba.trade.get.buyerView",
  "alibaba.trade.getBuyerOrderList",
  "alibaba.trade.cancel",
  "alibaba.trade.getLogisticsTraceInfo.buyerView",
  "product.search.queryProductDetail",
  "product.search.keywordQuery",
  "product.freight.estimate",
];

const FAULT_CODES = [
  "500_001",
  "500_002",
  "500_003",
  "500_004",
  "500_005",
  "500_006",
  "500_007",
  "500_008",
  "500_009",
  "CLOSE_ORDER_TOO_FAST",
  "ORDER_STATUS_ERROR",
  "ORDER_NOT_EXIST",
  "400_3",
  "404",
  "order.nopermission.buyer",
  "order.createtime.history",
];

const ORDER_STEPS = [
  "waitbuyerpay",
  "waitsellersend",
  "waitbuyerreceive",
  "success",
  "cancel",
];

export function StubControls() {
  const state = useQuery<Record<string, unknown>>("/api/admin/stub/state");
  const [log, setLog] = useState<Array<{ at: string; line: string; ok: boolean }>>([]);

  function record(line: string, ok: boolean) {
    setLog((l) => [{ at: new Date().toLocaleTimeString(), line, ok }, ...l].slice(0, 40));
  }

  async function send(path: string, body: unknown, label: string, method: "POST" | "PUT" | "DELETE" = "POST") {
    try {
      const res =
        method === "DELETE"
          ? await api.del<unknown>(path)
          : method === "PUT"
            ? await api.put<unknown>(path, body)
            : await api.post<unknown>(path, body);
      record(`${label} → ${res && typeof res === "object" ? JSON.stringify(res) : String(res ?? "ok")}`, true);
      state.reload();
    } catch (err) {
      record(`${label} → ${errorText(err)}`, false);
    }
  }

  return (
    <div className="stack">
      <div className={shell.panelHead}>
        <div>
          <h2 className={shell.panelTitle}>Stub controls</h2>
          <p className={shell.panelNote}>
            Drives the simulated 1688 gateway. Nothing here reaches the real 1688.
          </p>
        </div>
        <button className="btn" onClick={state.reload} disabled={state.loading}>
          Refresh state
        </button>
      </div>

      <div className={styles.grid}>
        <ClockPanel onSend={send} />
        <AdvancePanel onSend={send} />
        <FaultPanel onSend={send} />
        <PushPanel onSend={send} />
      </div>

      <section className="card">
        <h3 className={shell.panelTitle} style={{ marginBottom: 12 }}>
          Stub state
        </h3>
        {state.error ? (
          <p className="notice noticeWarn">
            {state.error} — the stub may not expose <span className="mono">GET /api/admin/stub/state</span>.
          </p>
        ) : state.loading && !state.data ? (
          <div className="skeleton" style={{ height: 120 }} />
        ) : (
          <pre className="pre">{JSON.stringify(state.data ?? {}, null, 2)}</pre>
        )}
      </section>

      <RawConsole onRecord={record} />

      {log.length > 0 && (
        <section className="card">
          <div className="rowBetween" style={{ marginBottom: 12 }}>
            <h3 className={shell.panelTitle}>Activity</h3>
            <button className="btn btnGhost" onClick={() => setLog([])}>
              Clear
            </button>
          </div>
          <ul className={styles.log}>
            {log.map((entry, i) => (
              <li key={i} className={entry.ok ? styles.logOk : styles.logErr}>
                <span className={styles.logAt}>{entry.at}</span>
                <span className={styles.logLine}>{entry.line}</span>
              </li>
            ))}
          </ul>
        </section>
      )}
    </div>
  );
}

type Sender = (path: string, body: unknown, label: string, method?: "POST" | "PUT" | "DELETE") => Promise<void>;

function ClockPanel({ onSend }: { onSend: Sender }) {
  const [speed, setSpeed] = useState("1");
  return (
    <section className="card">
      <h3 className={styles.title}>Clock speed</h3>
      <p className={styles.path}>POST /api/admin/stub/clock</p>
      <p className="xs faint" style={{ marginBottom: 12 }}>
        Multiplier on the stub's simulated time. Raise it to watch an order run its whole lifecycle in
        seconds instead of days.
      </p>
      <div className={styles.controls}>
        <Field label="Multiplier">
          {(id) => (
            <input
              id={id}
              className="input"
              inputMode="numeric"
              value={speed}
              onChange={(e) => setSpeed(e.target.value.replace(/[^0-9.]/g, ""))}
            />
          )}
        </Field>
        <div className={styles.presets}>
          {["1", "10", "60", "600"].map((s) => (
            <button key={s} className="btn btnGhost" onClick={() => setSpeed(s)}>
              ×{s}
            </button>
          ))}
        </div>
      </div>
      <button
        className="btn btnPrimary"
        style={{ marginTop: 12 }}
        onClick={() => void onSend("/api/admin/stub/clock", { speed: Number(speed) || 1 }, `clock ×${speed}`)}
      >
        Set clock
      </button>
    </section>
  );
}

function AdvancePanel({ onSend }: { onSend: Sender }) {
  const [orderId, setOrderId] = useState("");
  const [step, setStep] = useState("");
  return (
    <section className="card">
      <h3 className={styles.title}>Advance an order</h3>
      <p className={styles.path}>POST /api/admin/stub/orders/{"{cbuOrderId}"}/advance</p>
      <p className="xs faint" style={{ marginBottom: 12 }}>
        Pushes a supplier order to its next state, or straight to a chosen one.
      </p>
      <div className="stack">
        <Field label="1688 order ID" required>
          {(id) => (
            <input
              id={id}
              className="input"
              value={orderId}
              placeholder="2891234567890123"
              onChange={(e) => setOrderId(e.target.value.trim())}
            />
          )}
        </Field>
        <Field label="Target state" hint="Blank advances by one step.">
          {(id) => (
            <select id={id} className="select" value={step} onChange={(e) => setStep(e.target.value)}>
              <option value="">Next step</option>
              {ORDER_STEPS.map((s) => (
                <option key={s} value={s}>
                  {s}
                </option>
              ))}
            </select>
          )}
        </Field>
        <button
          className="btn btnPrimary"
          disabled={!orderId}
          onClick={() =>
            void onSend(
              `/api/admin/stub/orders/${encodeURIComponent(orderId)}/advance`,
              step ? { status: step } : {},
              `advance ${orderId}${step ? ` → ${step}` : ""}`,
            )
          }
        >
          Advance
        </button>
      </div>
    </section>
  );
}

function FaultPanel({ onSend }: { onSend: Sender }) {
  const [apiName, setApiName] = useState(FAULT_APIS[0]);
  const [code, setCode] = useState(FAULT_CODES[0]);
  const [message, setMessage] = useState("");
  const [count, setCount] = useState("1");
  return (
    <section className="card">
      <h3 className={styles.title}>Inject a fault</h3>
      <p className={styles.path}>POST /api/admin/stub/faults</p>
      <p className="xs faint" style={{ marginBottom: 12 }}>
        Makes the next N calls to one API fail with a documented business code, so the relay's error
        handling can be exercised on purpose.
      </p>
      <div className="stack">
        <Field label="API">
          {(id) => (
            <select id={id} className="select" value={apiName} onChange={(e) => setApiName(e.target.value)}>
              {FAULT_APIS.map((a) => (
                <option key={a} value={a}>
                  {a}
                </option>
              ))}
            </select>
          )}
        </Field>
        <FieldRow>
          <Field label="Error code">
            {(id) => (
              <select id={id} className="select" value={code} onChange={(e) => setCode(e.target.value)}>
                {FAULT_CODES.map((c) => (
                  <option key={c} value={c}>
                    {c}
                  </option>
                ))}
              </select>
            )}
          </Field>
          <Field label="Times">
            {(id) => (
              <input
                id={id}
                className="input"
                inputMode="numeric"
                value={count}
                onChange={(e) => setCount(e.target.value.replace(/[^0-9]/g, ""))}
              />
            )}
          </Field>
        </FieldRow>
        <Field label="Message" hint="Blank uses the documented text for the code.">
          {(id) => <input id={id} className="input" value={message} onChange={(e) => setMessage(e.target.value)} />}
        </Field>
        <div className={shell.toolbar}>
          <button
            className="btn btnPrimary"
            onClick={() =>
              void onSend(
                "/api/admin/stub/faults",
                { api: apiName, code, message, count: Number(count) || 1 },
                `fault ${apiName} ${code} ×${count || 1}`,
              )
            }
          >
            Arm fault
          </button>
          <button className="btn" onClick={() => void onSend("/api/admin/stub/faults", null, "clear faults", "DELETE")}>
            Clear all
          </button>
        </div>
      </div>
    </section>
  );
}

function PushPanel({ onSend }: { onSend: Sender }) {
  const [topic, setTopic] = useState(TOPICS[0]);
  const [target, setTarget] = useState("");
  return (
    <section className="card">
      <h3 className={styles.title}>Fire a push message</h3>
      <p className={styles.path}>POST /api/admin/stub/push</p>
      <p className="xs faint" style={{ marginBottom: 12 }}>
        Emits one of the 39 documented topics at our receiver, without waiting for the stub's own
        lifecycle to reach that point.
      </p>
      <div className="stack">
        <Field label="Topic">
          {(id) => (
            <select id={id} className="select" value={topic} onChange={(e) => setTopic(e.target.value)}>
              {TOPICS.map((t) => (
                <option key={t} value={t}>
                  {t}
                </option>
              ))}
            </select>
          )}
        </Field>
        <Field label="Order or offer ID" hint="The subject of the message.">
          {(id) => (
            <input id={id} className="input" value={target} onChange={(e) => setTarget(e.target.value.trim())} />
          )}
        </Field>
        <button
          className="btn btnPrimary"
          onClick={() => void onSend("/api/admin/stub/push", { topic, id: target }, `push ${topic}`)}
        >
          Fire message
        </button>
      </div>
    </section>
  );
}

/** Escape hatch: any stub path, any method, any body. */
function RawConsole({ onRecord }: { onRecord: (line: string, ok: boolean) => void }) {
  const [method, setMethod] = useState<"GET" | "POST" | "PUT" | "DELETE">("POST");
  const [path, setPath] = useState("/api/admin/stub/");
  const [body, setBody] = useState("{}");
  const [result, setResult] = useState<string | null>(null);
  const [busy, setBusy] = useState(false);

  async function run() {
    setBusy(true);
    setResult(null);
    let parsed: unknown = undefined;
    if (method !== "GET" && method !== "DELETE") {
      try {
        parsed = JSON.parse(body || "{}");
      } catch {
        setResult("Request body is not valid JSON.");
        onRecord(`${method} ${path} → invalid JSON body`, false);
        setBusy(false);
        return;
      }
    }
    try {
      const res =
        method === "GET"
          ? await api.get<unknown>(path)
          : method === "DELETE"
            ? await api.del<unknown>(path)
            : method === "PUT"
              ? await api.put<unknown>(path, parsed)
              : await api.post<unknown>(path, parsed);
      setResult(JSON.stringify(res ?? null, null, 2));
      onRecord(`${method} ${path} → ok`, true);
    } catch (err) {
      setResult(errorText(err));
      onRecord(`${method} ${path} → ${errorText(err)}`, false);
    } finally {
      setBusy(false);
    }
  }

  return (
    <section className="card">
      <h3 className={shell.panelTitle} style={{ marginBottom: 4 }}>
        Raw stub request
      </h3>
      <p className={shell.panelNote} style={{ marginBottom: 16 }}>
        Everything under <span className="mono">/api/admin/stub/</span> is proxied straight to the stub
        gateway's control endpoints.
      </p>
      <div className="stack">
        <div className={styles.rawRow}>
          <select
            className="select"
            style={{ width: 110 }}
            value={method}
            onChange={(e) => setMethod(e.target.value as typeof method)}
          >
            <option>GET</option>
            <option>POST</option>
            <option>PUT</option>
            <option>DELETE</option>
          </select>
          <input className="input" value={path} onChange={(e) => setPath(e.target.value)} aria-label="Path" />
          <button className="btn btnPrimary" onClick={() => void run()} disabled={busy || !path.startsWith("/api/admin/stub")}>
            {busy ? "…" : "Send"}
          </button>
        </div>
        {method !== "GET" && method !== "DELETE" && (
          <Field label="JSON body">
            {(id) => (
              <textarea
                id={id}
                className="textarea"
                value={body}
                spellCheck={false}
                onChange={(e) => setBody(e.target.value)}
              />
            )}
          </Field>
        )}
        {result !== null && <pre className="pre">{result}</pre>}
      </div>
    </section>
  );
}
