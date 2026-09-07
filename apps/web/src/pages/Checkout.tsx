import { useState } from "react";
import type { FormEvent } from "react";
import { Link, useNavigate } from "react-router-dom";
import { useCart } from "../lib/cart";
import { api, ApiError, errorText } from "../lib/api";
import { Field, FieldRow } from "../components/Field";
import { Money } from "../components/Money";
import type { CheckoutRequest, CheckoutResult, Issue, Money as MoneyType } from "../lib/types";
import { pluralise } from "../lib/fmt";
import styles from "./Checkout.module.css";

const BLANK: CheckoutRequest = {
  email: "",
  name: "",
  phone: "",
  address: { line1: "", city: "", province: "", postcode: "" },
};

export function Checkout() {
  const navigate = useNavigate();
  const { cart, loading, reload } = useCart();

  const [form, setForm] = useState<CheckoutRequest>(BLANK);
  const [fieldErrors, setFieldErrors] = useState<Record<string, string>>({});
  const [issues, setIssues] = useState<Issue[]>([]);
  const [error, setError] = useState<string | undefined>();
  const [placing, setPlacing] = useState(false);
  const [placed, setPlaced] = useState<CheckoutResult | null>(null);
  const [paying, setPaying] = useState(false);

  const set = (patch: Partial<CheckoutRequest>) => setForm((f) => ({ ...f, ...patch }));
  const setAddr = (patch: Partial<CheckoutRequest["address"]>) =>
    setForm((f) => ({ ...f, address: { ...f.address, ...patch } }));

  function validate(): boolean {
    const errs: Record<string, string> = {};
    if (!form.email.trim()) errs.email = "Required — we send order updates here.";
    else if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(form.email.trim())) errs.email = "That does not look like an email address.";
    if (!form.name.trim()) errs.name = "Required.";
    if (!form.phone.trim()) errs.phone = "Required for delivery.";
    if (!form.address.line1.trim()) errs.line1 = "Required.";
    if (!form.address.city.trim()) errs.city = "Required.";
    if (!form.address.province.trim()) errs.province = "Required.";
    if (!form.address.postcode.trim()) errs.postcode = "Required.";
    setFieldErrors(errs);
    return Object.keys(errs).length === 0;
  }

  async function placeOrder(e: FormEvent) {
    e.preventDefault();
    if (!validate()) return;
    setPlacing(true);
    setError(undefined);
    setIssues([]);
    try {
      const result = await api.post<CheckoutResult>("/api/checkout", form);
      setPlaced(result);
    } catch (err) {
      // 409 means the cart went stale between rendering and submitting: price
      // moved, stock gone, item delisted. Show the issues and refresh the cart.
      if (err instanceof ApiError && err.status === 409) {
        setIssues(err.issues);
        setError("Your cart changed. Review the problems below and try again.");
        void reload();
      } else {
        setError(errorText(err));
      }
    } finally {
      setPlacing(false);
    }
  }

  async function pay() {
    if (!placed) return;
    setPaying(true);
    setError(undefined);
    try {
      await api.post(`/api/orders/${encodeURIComponent(placed.orderId)}/pay?t=${encodeURIComponent(placed.token)}`);
      navigate(`/orders/${encodeURIComponent(placed.orderId)}?t=${encodeURIComponent(placed.token)}`);
    } catch (err) {
      setError(errorText(err));
      setPaying(false);
    }
  }

  if (placed) {
    return (
      <div className={styles.narrow}>
        <div className="stack">
          <div>
            <h1>Order placed</h1>
            <p className="muted small">
              Order <span className="mono">{placed.orderId}</span> is reserved. It is not confirmed with
              the suppliers until payment goes through.
            </p>
          </div>

          <div className="card">
            <div className="rowBetween">
              <span className="strong">Amount due</span>
              <Money value={placed.total} size="lg" />
            </div>
            <hr className="divider" />
            <p className="small muted">Ships as {pluralise(placed.parcels || 1, "parcel")}.</p>
          </div>

          {error && <p className="notice noticeError">{error}</p>}

          <button className="btn btnPrimary btnLg btnBlock" onClick={pay} disabled={paying}>
            {paying ? "Processing payment…" : "Pay now"}
          </button>

          {placed.payUrl && (
            <p className="small center">
              <a href={placed.payUrl} rel="noreferrer">
                Pay on the payment provider's page instead
              </a>
            </p>
          )}

          <p className="small center">
            <Link to={`/orders/${encodeURIComponent(placed.orderId)}?t=${encodeURIComponent(placed.token)}`}>
              View order status
            </Link>
          </p>
          <p className="xs faint center">
            Keep this link — it is the only way back to this order.
          </p>
        </div>
      </div>
    );
  }

  if (loading && !cart) {
    return <div className="skeleton" style={{ height: 300, borderRadius: 12 }} />;
  }

  const lineCount = (cart?.groups ?? []).reduce((n, g) => n + g.lines.length, 0);
  if (lineCount === 0) {
    return (
      <div className="empty">
        <p>There is nothing to check out.</p>
        <p style={{ marginTop: 12 }}>
          <Link className="btn" to="/search">
            Browse products
          </Link>
        </p>
      </div>
    );
  }

  const parcels = (cart?.groups ?? []).reduce((n, g) => n + (g.parcels || 1), 0);

  return (
    <div className="stack">
      <h1>Checkout</h1>

      {error && <p className="notice noticeError">{error}</p>}
      {issues.length > 0 && (
        <ul className="stackSm">
          {issues.map((issue, i) => (
            <li key={i} className="notice noticeWarn">
              {issue.message}
              {issue.code && <span className="xs faint"> ({issue.code})</span>}
            </li>
          ))}
        </ul>
      )}

      <div className={styles.layout}>
        <form className={styles.form} onSubmit={placeOrder} noValidate>
          <section className="card">
            <h2 className={styles.sectionTitle}>Contact</h2>
            <div className="stack">
              <Field label="Email" required error={fieldErrors.email} hint="Order updates and the tracking link go here.">
                {(id) => (
                  <input
                    id={id}
                    className="input"
                    type="email"
                    autoComplete="email"
                    value={form.email}
                    onChange={(e) => set({ email: e.target.value })}
                  />
                )}
              </Field>
              <FieldRow>
                <Field label="Full name" required error={fieldErrors.name}>
                  {(id) => (
                    <input
                      id={id}
                      className="input"
                      autoComplete="name"
                      value={form.name}
                      onChange={(e) => set({ name: e.target.value })}
                    />
                  )}
                </Field>
                <Field label="Phone" required error={fieldErrors.phone}>
                  {(id) => (
                    <input
                      id={id}
                      className="input"
                      type="tel"
                      autoComplete="tel"
                      value={form.phone}
                      onChange={(e) => set({ phone: e.target.value })}
                    />
                  )}
                </Field>
              </FieldRow>
            </div>
          </section>

          <section className="card">
            <h2 className={styles.sectionTitle}>Delivery address</h2>
            <div className="stack">
              <Field label="Address" required error={fieldErrors.line1}>
                {(id) => (
                  <input
                    id={id}
                    className="input"
                    autoComplete="address-line1"
                    value={form.address.line1}
                    onChange={(e) => setAddr({ line1: e.target.value })}
                  />
                )}
              </Field>
              <FieldRow>
                <Field label="City / district" required error={fieldErrors.city}>
                  {(id) => (
                    <input
                      id={id}
                      className="input"
                      autoComplete="address-level2"
                      value={form.address.city}
                      onChange={(e) => setAddr({ city: e.target.value })}
                    />
                  )}
                </Field>
                <Field label="Province" required error={fieldErrors.province}>
                  {(id) => (
                    <input
                      id={id}
                      className="input"
                      autoComplete="address-level1"
                      value={form.address.province}
                      onChange={(e) => setAddr({ province: e.target.value })}
                    />
                  )}
                </Field>
                <Field label="Postcode" required error={fieldErrors.postcode}>
                  {(id) => (
                    <input
                      id={id}
                      className="input"
                      inputMode="numeric"
                      autoComplete="postal-code"
                      value={form.address.postcode}
                      onChange={(e) => setAddr({ postcode: e.target.value })}
                    />
                  )}
                </Field>
              </FieldRow>
            </div>
          </section>

          <button className="btn btnPrimary btnLg" type="submit" disabled={placing || !cart?.checkoutable}>
            {placing ? "Placing order…" : "Place order"}
          </button>
          {!cart?.checkoutable && (
            <p className="notice noticeWarn">
              Your cart has unresolved problems. <Link to="/cart">Go back to the cart</Link> to fix them.
            </p>
          )}
        </form>

        <aside className={styles.summary}>
          <h2 className={styles.sectionTitle}>Summary</h2>
          <dl className={styles.totals}>
            <Row label="Goods" value={cart?.totals.goods} />
            <Row label="China freight" value={cart?.totals.chinaFreight} />
            <Row label="International shipping" value={cart?.totals.intl} />
            <Row label="Platform fee" value={cart?.totals.fee} />
          </dl>
          <div className={styles.grand}>
            <span>Total</span>
            <Money value={cart?.totals.total} size="lg" />
          </div>
          <p className="xs faint">
            {pluralise(lineCount, "line")} · ships as {pluralise(parcels, "parcel")}. Nothing is added at
            payment.
          </p>
          <Link className="small" to="/cart">
            Edit cart
          </Link>
        </aside>
      </div>
    </div>
  );
}

function Row({ label, value }: { label: string; value: MoneyType | undefined }) {
  return (
    <div className={styles.totalLine}>
      <dt>{label}</dt>
      <dd>
        <Money value={value} size="sm" muted />
      </dd>
    </div>
  );
}
