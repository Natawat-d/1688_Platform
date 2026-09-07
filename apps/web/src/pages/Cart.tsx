import { Link, useNavigate } from "react-router-dom";
import { useCart } from "../lib/cart";
import { CartGroup } from "../components/CartGroup";
import { Money } from "../components/Money";
import type { Money as MoneyType } from "../lib/types";
import { pluralise } from "../lib/fmt";
import styles from "./Cart.module.css";

export function Cart() {
  const navigate = useNavigate();
  const { cart, loading, busy, error, setQty, remove, reload } = useCart();

  if (loading && !cart) {
    return (
      <div className="stack">
        <h1>Cart</h1>
        <div className="skeleton" style={{ height: 200, borderRadius: 12 }} />
      </div>
    );
  }

  const groups = cart?.groups ?? [];
  const empty = groups.length === 0;

  if (empty) {
    return (
      <div className="stack">
        <h1>Cart</h1>
        {error && <p className="notice noticeError">{error}</p>}
        <div className="empty">
          <p>Your cart is empty.</p>
          <p style={{ marginTop: 12 }}>
            <Link className="btn" to="/search">
              Browse products
            </Link>
          </p>
        </div>
      </div>
    );
  }

  const parcels = groups.reduce((n, g) => n + (g.parcels || 1), 0);
  const allIssues = groups.flatMap((g) => g.issues ?? []);

  return (
    <div className="stack">
      <h1>Cart</h1>

      {/* Splitting by supplier is not a UI choice: each group becomes one 1688
          order, so it is also how the goods physically arrive. */}
      <p className={styles.banner}>
        Ships as {pluralise(parcels, "parcel")} from {pluralise(groups.length, "supplier")}. Each
        supplier is ordered and shipped separately, so parcels can arrive on different days.
      </p>

      {error && <p className="notice noticeError">{error}</p>}

      <div className={styles.layout}>
        <div className={styles.groups}>
          {groups.map((group) => (
            <CartGroup
              key={group.sellerOpenId}
              group={group}
              busy={busy}
              onQty={(id, q) => void setQty(id, q).catch(() => undefined)}
              onRemove={(id) => void remove(id).catch(() => undefined)}
            />
          ))}
        </div>

        <aside className={styles.summary}>
          <h2 className={styles.summaryTitle}>Order summary</h2>
          <dl className={styles.totals}>
            <Line label="Goods" value={cart?.totals.goods} />
            <Line label="China freight" value={cart?.totals.chinaFreight} />
            <Line label="International shipping" value={cart?.totals.intl} />
            <Line label="Platform fee" value={cart?.totals.fee} />
          </dl>
          <div className={styles.grandTotal}>
            <span>Total</span>
            <Money value={cart?.totals.total} size="lg" />
          </div>
          <p className="xs faint">Duties on arrival, if any, are not included.</p>

          {!cart?.checkoutable && (
            <div className="notice noticeWarn">
              {allIssues.length > 0
                ? "Fix the problems listed with your items before checking out."
                : "This cart cannot be checked out yet."}
            </div>
          )}

          <button
            className="btn btnPrimary btnLg btnBlock"
            disabled={!cart?.checkoutable || busy}
            onClick={() => navigate("/checkout")}
          >
            Checkout
          </button>
          <button className="btn btnGhost btnBlock" onClick={() => void reload()} disabled={busy}>
            {busy ? "Updating…" : "Refresh prices"}
          </button>
        </aside>
      </div>
    </div>
  );
}

function Line({ label, value }: { label: string; value: MoneyType | undefined }) {
  return (
    <div className={styles.totalLine}>
      <dt>{label}</dt>
      <dd>
        <Money value={value} size="sm" muted />
      </dd>
    </div>
  );
}
