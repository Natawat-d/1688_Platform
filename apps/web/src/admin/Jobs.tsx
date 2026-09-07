import { useState } from "react";
import { useQuery } from "../lib/useQuery";
import { DataTable } from "../components/DataTable";
import type { Column } from "../components/DataTable";
import { StatusBadge } from "../components/StatusBadge";
import type { ApiCall, Job, PushMessage } from "../lib/types";
import { dateTimeText, num, timeText } from "../lib/fmt";
import shell from "./Admin.module.css";
import styles from "./Jobs.module.css";

type Tab = "jobs" | "calls" | "messages";

const TABS: Array<[Tab, string]> = [
  ["jobs", "Jobs"],
  ["calls", "API calls"],
  ["messages", "Push messages"],
];

export function Jobs() {
  const [tab, setTab] = useState<Tab>("jobs");

  const jobs = useQuery<Job[]>(tab === "jobs" ? "/api/admin/jobs" : null);
  const calls = useQuery<ApiCall[]>(tab === "calls" ? "/api/admin/api-calls" : null);
  const messages = useQuery<PushMessage[]>(tab === "messages" ? "/api/admin/messages" : null);

  const active = tab === "jobs" ? jobs : tab === "calls" ? calls : messages;

  return (
    <div className="stack">
      <div className={shell.panelHead}>
        <div>
          <h2 className={shell.panelTitle}>Jobs and logs</h2>
          <p className={shell.panelNote}>
            Background work, every call we made to the 1688 gateway, and every push message we
            received. This is where an integration failure shows itself first.
          </p>
        </div>
        <button className="btn" onClick={active.reload} disabled={active.loading}>
          Refresh
        </button>
      </div>

      <div className={styles.switch} role="tablist">
        {TABS.map(([value, label]) => (
          <button
            key={value}
            role="tab"
            aria-selected={tab === value}
            className={tab === value ? `${styles.switchTab} ${styles.switchActive}` : styles.switchTab}
            onClick={() => setTab(value)}
          >
            {label}
          </button>
        ))}
      </div>

      {active.error && <p className="notice noticeError">{active.error}</p>}

      {tab === "jobs" && <JobsTable rows={jobs.data} loading={jobs.loading} />}
      {tab === "calls" && <CallsTable rows={calls.data} loading={calls.loading} />}
      {tab === "messages" && <MessagesTable rows={messages.data} loading={messages.loading} />}
    </div>
  );
}

function JobsTable({ rows, loading }: { rows: Job[] | undefined; loading: boolean }) {
  const columns: Column<Job>[] = [
    { key: "id", header: "ID", render: (j) => <span className="mono">{j.id}</span> },
    { key: "kind", header: "Kind", render: (j) => j.kind || "—" },
    { key: "status", header: "Status", render: (j) => <StatusBadge status={j.status} /> },
    {
      key: "progress",
      header: "Progress",
      numeric: true,
      render: (j) =>
        typeof j.done === "number"
          ? `${num(j.done)}${typeof j.total === "number" && j.total > 0 ? ` / ${num(j.total)}` : ""}`
          : "—",
    },
    { key: "started", header: "Started", secondary: true, render: (j) => <span className="xs">{dateTimeText(j.startedAt)}</span> },
    {
      key: "finished",
      header: "Finished",
      secondary: true,
      render: (j) => <span className="xs">{j.finishedAt ? dateTimeText(j.finishedAt) : "—"}</span>,
    },
    {
      key: "error",
      header: "Error",
      render: (j) => (j.error ? <span className="xs" style={{ color: "var(--danger)" }}>{j.error}</span> : <span className="xs faint">—</span>),
    },
  ];
  return <DataTable rows={rows} columns={columns} loading={loading} rowKey={(j, i) => j.id || String(i)} empty="No jobs have run." />;
}

function CallsTable({ rows, loading }: { rows: ApiCall[] | undefined; loading: boolean }) {
  const columns: Column<ApiCall>[] = [
    { key: "at", header: "Time", render: (c) => <span className="xs nums">{timeText(c.at)}</span> },
    { key: "api", header: "API", render: (c) => <span className="mono">{c.api}</span> },
    {
      key: "status",
      header: "HTTP",
      numeric: true,
      render: (c) => (
        <span className={c.status && c.status >= 400 ? "badge badgeDanger" : "badge badgeOk"}>{c.status ?? "—"}</span>
      ),
    },
    {
      key: "duration",
      header: "Time",
      numeric: true,
      secondary: true,
      render: (c) => (typeof c.durationMs === "number" ? `${num(c.durationMs)} ms` : "—"),
    },
    { key: "url", header: "URL", secondary: true, render: (c) => <span className="mono">{c.url || "—"}</span> },
    {
      key: "err",
      header: "Error",
      render: (c) => (c.err ? <span className="xs" style={{ color: "var(--danger)" }}>{c.err}</span> : <span className="xs faint">—</span>),
    },
  ];
  return (
    <DataTable
      rows={rows}
      columns={columns}
      loading={loading}
      rowKey={(c, i) => c.id || `${c.api}-${i}`}
      empty="No gateway calls logged yet."
    />
  );
}

function MessagesTable({ rows, loading }: { rows: PushMessage[] | undefined; loading: boolean }) {
  const [open, setOpen] = useState<string | null>(null);

  const columns: Column<PushMessage>[] = [
    { key: "topic", header: "Topic", render: (m) => <span className="mono">{m.topic}</span> },
    { key: "msgId", header: "Message ID", secondary: true, render: (m) => <span className="mono">{m.msgId || m.id || "—"}</span> },
    { key: "status", header: "Status", render: (m) => <StatusBadge status={m.status} /> },
    { key: "received", header: "Received", render: (m) => <span className="xs nums">{dateTimeText(m.receivedAt)}</span> },
    {
      key: "processed",
      header: "Processed",
      secondary: true,
      render: (m) => <span className="xs nums">{m.processedAt ? dateTimeText(m.processedAt) : "—"}</span>,
    },
    {
      key: "payload",
      header: "",
      render: (m) => {
        const key = m.msgId || m.id || "";
        return (
          <button className="btn btnGhost" onClick={() => setOpen((o) => (o === key ? null : key))}>
            {open === key ? "Hide" : "Payload"}
          </button>
        );
      },
    },
  ];

  const shown = rows ?? [];
  const detail = open ? shown.find((m) => (m.msgId || m.id) === open) : undefined;

  return (
    <>
      <DataTable
        rows={rows}
        columns={columns}
        loading={loading}
        rowKey={(m, i) => m.msgId || m.id || String(i)}
        empty="No push messages received."
      />
      {detail && (
        <section className="card" style={{ marginTop: 16 }}>
          <div className="rowBetween" style={{ marginBottom: 12 }}>
            <h3 className={shell.panelTitle}>
              <span className="mono">{detail.topic}</span>
            </h3>
            <button className="btn btnGhost" onClick={() => setOpen(null)}>
              Close
            </button>
          </div>
          <pre className="pre">{JSON.stringify(detail.payload ?? detail, null, 2)}</pre>
        </section>
      )}
    </>
  );
}
