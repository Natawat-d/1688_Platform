import { useMemo, useState } from "react";
import { Link, useSearchParams } from "react-router-dom";
import { useQuery } from "../lib/useQuery";
import { api, errorText, withQuery } from "../lib/api";
import { DataTable } from "../components/DataTable";
import type { Column } from "../components/DataTable";
import { StatusBadge } from "../components/StatusBadge";
import { Money } from "../components/Money";
import type { AdminOrder, SupplierOrder } from "../lib/types";
import { dateTimeText } from "../lib/fmt";
import shell from "./Admin.module.css";
import styles from "./Orders.module.css";

const STATUSES = ["", "pending", "paid", "relaying", "relayed", "shipped", "delivered", "cancelled", "failed"];

type Action = "relay" | "pay" | "cancel";

export function Orders() {
  const [params, setParams] = useSearchParams();
  const status = params.get("status") ?? "";
  const url = useMemo(() => withQuery("/api/admin/orders", { status }), [status]);
  const orders = useQuery<AdminOrder[]>(url);

  const [busyId, setBusyId] = useState<string | null>(null);
  const [error, setError] = useState<string | undefined>();
  const [expanded, setExpanded] = useState<string | null>(null);

  async function act(supplierOrderId: string, action: Action) {
    if (action === "cancel" && !window.confirm("Cancel this supplier order on 1688?")) return;
    setBusyId(supplierOrderId);
    setError(undefined);
    try {
      await api.post(`/api/admin/supplier-orders/${encodeURIComponent(supplierOrderId)}/${action}`);
      orders.reload();
    } catch (err) {
      setError(errorText(err));
    } finally {
      setBusyId(null);
    }
  }

  const columns: Column<AdminOrder>[] = [
    {
      key: "order",
      header: "Order",
      render: (o) => {
        const pid = o.publicId || o.orderId;
        return (
          <div>
            <button className={styles.linkish} onClick={() => setExpanded((e) => (e === o.orderId ? null : o.orderId))}>
              <span className="mono">{o.orderId}</span>
            </button>
            {pid !== o.orderId && <div className="xs faint mono">{pid}</div>}
          </div>
        );
      },
    },
    { key: "status", header: "Status", render: (o) => <StatusBadge status={o.status} /> },
    { key: "email", header: "Customer", secondary: true, render: (o) => <span className="small">{o.email || "—"}</span> },
    {
      key: "placed",
      header: "Placed",
      secondary: true,
      render: (o) => <span className="xs">{dateTimeText(o.placedAt)}</span>,
    },
    { key: "total", header: "Total", numeric: true, render: (o) => <Money value={o.total} size="sm" /> },
    {
      key: "suppliers",
      header: "Supplier orders",
      render: (o) => {
        const subs = o.supplierOrders ?? [];
        if (subs.length === 0) return <span className="xs faint">None yet</span>;
        return (
          <div className={styles.subs}>
            {subs.map((s) => (
              <SupplierRow key={s.id} sub={s} busy={busyId === s.id} onAction={act} />
            ))}
          </div>
        );
      },
    },
  ];

  const shown = orders.data ?? [];
  const detail = expanded ? shown.find((o) => o.orderId === expanded) : undefined;

  return (
    <div className="stack">
      <div className={shell.panelHead}>
        <div>
          <h2 className={shell.panelTitle}>Orders</h2>
          <p className={shell.panelNote}>
            Relay queue and supplier orders. Each supplier order is one 1688 order; relay creates it,
            pay settles it from the platform account.
          </p>
        </div>
        <div className={shell.toolbar}>
          <label className="row">
            <span className="srOnly">Filter by status</span>
            <select
              className="select"
              value={status}
              onChange={(e) => {
                const next = new URLSearchParams(params);
                if (e.target.value) next.set("status", e.target.value);
                else next.delete("status");
                setParams(next);
              }}
            >
              {STATUSES.map((s) => (
                <option key={s} value={s}>
                  {s === "" ? "All statuses" : s}
                </option>
              ))}
            </select>
          </label>
          <button className="btn" onClick={orders.reload} disabled={orders.loading}>
            Refresh
          </button>
        </div>
      </div>

      {error && <p className="notice noticeError">{error}</p>}
      {orders.error && <p className="notice noticeError">{orders.error}</p>}

      <DataTable
        rows={orders.data}
        columns={columns}
        loading={orders.loading}
        rowKey={(o) => o.orderId}
        empty={status ? `No orders with status "${status}".` : "No orders yet."}
      />

      {detail && (
        <section className="card">
          <div className="rowBetween" style={{ marginBottom: 12 }}>
            <h3 className={shell.panelTitle}>
              Raw order <span className="mono">{detail.orderId}</span>
            </h3>
            <div className={shell.toolbar}>
              <Link className="btn btnGhost" to={`/orders/${encodeURIComponent(detail.publicId || detail.orderId)}`}>
                Customer view
              </Link>
              <button className="btn btnGhost" onClick={() => setExpanded(null)}>
                Close
              </button>
            </div>
          </div>
          <pre className="pre">{JSON.stringify(detail, null, 2)}</pre>
        </section>
      )}
    </div>
  );
}

function SupplierRow({
  sub,
  busy,
  onAction,
}: {
  sub: SupplierOrder;
  busy: boolean;
  onAction: (id: string, action: Action) => void;
}) {
  return (
    <div className={styles.sub}>
      <div className={styles.subHead}>
        <StatusBadge status={sub.status} />
        <span className="xs">{sub.sellerName || sub.sellerOpenId || "Supplier"}</span>
        {sub.cbuOrderId && <span className="xs faint mono">{sub.cbuOrderId}</span>}
      </div>
      {sub.error && <p className="xs" style={{ color: "var(--danger)" }}>{sub.error}</p>}
      <div className={styles.subActions}>
        <button className="btn btnGhost" disabled={busy} onClick={() => onAction(sub.id, "relay")}>
          Relay
        </button>
        <button className="btn btnGhost" disabled={busy} onClick={() => onAction(sub.id, "pay")}>
          Pay
        </button>
        <button className="btn btnGhost btnDanger" disabled={busy} onClick={() => onAction(sub.id, "cancel")}>
          Cancel
        </button>
      </div>
    </div>
  );
}
