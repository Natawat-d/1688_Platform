import { useEffect, useMemo, useState } from "react";
import { Link, useParams, useSearchParams } from "react-router-dom";
import { useQuery } from "../lib/useQuery";
import { Timeline } from "../components/Timeline";
import { StatusBadge } from "../components/StatusBadge";
import { Money } from "../components/Money";
import { Readout } from "../components/Field";
import type { Money as MoneyType, Order, Parcel } from "../lib/types";
import { dateTimeText, num, pluralise } from "../lib/fmt";
import styles from "./OrderStatus.module.css";

/** Once an order reaches one of these there is nothing left to poll for. */
const TERMINAL = /^(delivered|completed|complete|cancelled|canceled|closed|refunded|failed)$/i;

const POLL_MS = 10_000;

export function OrderStatus() {
  const { publicId = "" } = useParams();
  const [params] = useSearchParams();
  const token = params.get("t") ?? "";

  const url = useMemo(() => {
    if (!publicId) return null;
    return `/api/orders/${encodeURIComponent(publicId)}${token ? `?t=${encodeURIComponent(token)}` : ""}`;
  }, [publicId, token]);

  const { data: order, error, loading, reload } = useQuery<Order>(url);
  const terminal = TERMINAL.test(order?.status ?? "");

  // Poll while the order is still moving. Pause when the tab is hidden so a
  // forgotten background tab is not hitting the API every ten seconds forever.
  useEffect(() => {
    if (!order || terminal) return;
    const id = window.setInterval(() => {
      if (document.visibilityState === "visible") reload();
    }, POLL_MS);
    return () => window.clearInterval(id);
  }, [order, terminal, reload]);

  if (error) {
    return (
      <div className="stack">
        <h1>Order</h1>
        <p className="notice noticeError">{error}</p>
        <p className="small muted">
          Order pages open with the link from your confirmation. If you followed an old link, the token
          may have expired.
        </p>
        <p>
          <Link className="btn" to="/">
            Go to the home page
          </Link>
        </p>
      </div>
    );
  }

  if (loading && !order) {
    return (
      <div className="stack">
        <div className="skeleton" style={{ height: 90, borderRadius: 12 }} />
        <div className="skeleton" style={{ height: 220, borderRadius: 12 }} />
      </div>
    );
  }

  if (!order) return <div className="empty">Order not found.</div>;

  return (
    <div className="stack">
      <header className={styles.head}>
        <div>
          <p className="xs faint">Order</p>
          <h1 className={styles.orderId}>{order.orderId}</h1>
          <p className="small muted">
            Placed {dateTimeText(order.placedAt)}
            {order.paidAt ? ` · paid ${dateTimeText(order.paidAt)}` : ""}
          </p>
        </div>
        <div className={styles.headRight}>
          <StatusBadge status={order.status} label={order.statusLabel} />
          {!terminal && (
            <span className="xs faint" title={`Refreshes every ${POLL_MS / 1000} seconds`}>
              Live · updates automatically
            </span>
          )}
        </div>
      </header>

      <div className={styles.layout}>
        <div className={styles.main}>
          <section className="card">
            <h2 className={styles.panelTitle}>Progress</h2>
            <Timeline steps={order.timeline} />
          </section>

          <section>
            <div className="sectionTitle">
              <h2>{pluralise(order.parcels?.length ?? 0, "parcel")}</h2>
              <span className="xs faint">One per supplier</span>
            </div>
            <div className="stackSm">
              {(order.parcels ?? []).map((parcel, i) => (
                <ParcelPanel key={parcel.id || String(i)} parcel={parcel} defaultOpen={i === 0} />
              ))}
              {(order.parcels ?? []).length === 0 && (
                <div className="empty">Parcels appear once the suppliers accept the order.</div>
              )}
            </div>
          </section>
        </div>

        <aside className={styles.side}>
          <section className="card">
            <h2 className={styles.panelTitle}>Total</h2>
            <Readout label="Goods">
              <Money value={order.totals?.goods} size="sm" muted />
            </Readout>
            <Readout label="China freight">
              <Money value={order.totals?.chinaFreight} size="sm" muted />
            </Readout>
            <Readout label="International">
              <Money value={order.totals?.intl} size="sm" muted />
            </Readout>
            <Readout label="Platform fee">
              <Money value={order.totals?.fee} size="sm" muted />
            </Readout>
            <div className={styles.grand}>
              <span>Paid</span>
              <Money value={order.totals?.total} size="md" />
            </div>
            {order.email && <p className="xs faint">Updates sent to {order.email}</p>}
          </section>

          <section className="card">
            <h2 className={styles.panelTitle}>Items</h2>
            <ul className={styles.items}>
              {(order.items ?? []).map((item, i) => (
                <li key={`${item.offerId}-${item.skuId}-${i}`} className={styles.item}>
                  {item.image ? (
                    <img className={styles.itemThumb} src={item.image} alt="" loading="lazy" />
                  ) : (
                    <span className={styles.itemThumb} aria-hidden="true" />
                  )}
                  <div className={styles.itemBody}>
                    <Link className={styles.itemTitle} to={`/p/${encodeURIComponent(item.offerId)}`}>
                      {item.title}
                    </Link>
                    {item.skuLabel && <p className="xs faint">{item.skuLabel}</p>}
                    <p className="xs muted">
                      {num(item.quantity)} × <MoneyInline value={item.unit} />
                    </p>
                  </div>
                  <Money value={item.line} size="sm" />
                </li>
              ))}
            </ul>
          </section>
        </aside>
      </div>
    </div>
  );
}

function MoneyInline({ value }: { value: MoneyType | undefined }) {
  return <Money value={value} size="sm" muted />;
}

function ParcelPanel({ parcel, defaultOpen }: { parcel: Parcel; defaultOpen: boolean }) {
  const [open, setOpen] = useState(defaultOpen);

  return (
    <section className={styles.parcel}>
      <button
        type="button"
        className={styles.parcelHead}
        aria-expanded={open}
        onClick={() => setOpen((o) => !o)}
      >
        <span className={styles.parcelTitle}>
          <span className="strong">{parcel.sellerName || "Supplier"}</span>
          <StatusBadge status={parcel.status} />
        </span>
        <span className={styles.parcelMeta}>
          {parcel.trackingNo ? `${parcel.carrier || "Carrier"} · ${parcel.trackingNo}` : "No tracking yet"}
          <span className={open ? `${styles.chevron} ${styles.chevronOpen}` : styles.chevron} aria-hidden="true">
            ▾
          </span>
        </span>
      </button>

      {open && (
        <div className={styles.parcelBody}>
          <div className={styles.parcelFacts}>
            {parcel.cbuOrderId && (
              <Readout label="Supplier order" mono>
                {parcel.cbuOrderId}
              </Readout>
            )}
            {parcel.carrier && <Readout label="Carrier">{parcel.carrier}</Readout>}
            {parcel.trackingNo && (
              <Readout label="Tracking number" mono>
                {parcel.trackingNo}
              </Readout>
            )}
          </div>

          {(parcel.items ?? []).length > 0 && (
            <ul className={styles.parcelItems}>
              {parcel.items.map((item, i) => (
                <li key={`${item.offerId}-${i}`}>
                  <span className={styles.parcelItemQty}>{num(item.quantity)}×</span>
                  <span className={styles.parcelItemTitle}>{item.title}</span>
                  {item.skuLabel && <span className="xs faint">{item.skuLabel}</span>}
                </li>
              ))}
            </ul>
          )}

          <h3 className={styles.eventsTitle}>Tracking</h3>
          {(parcel.events ?? []).length === 0 ? (
            <p className="small faint">No tracking events yet.</p>
          ) : (
            <ol className={styles.events}>
              {parcel.events.map((ev, i) => (
                <li key={`${ev.at}-${i}`} className={styles.event}>
                  <time className={styles.eventAt}>{dateTimeText(ev.at)}</time>
                  <span className={styles.eventText}>{ev.text}</span>
                  {ev.source && <span className={styles.eventSource}>{ev.source}</span>}
                </li>
              ))}
            </ol>
          )}
        </div>
      )}
    </section>
  );
}
