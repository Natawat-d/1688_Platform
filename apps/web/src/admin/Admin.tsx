import { useState } from "react";
import { NavLink, Outlet } from "react-router-dom";
import { getAdminToken, setAdminToken } from "../lib/api";
import { Field } from "../components/Field";
import styles from "./Admin.module.css";

const TABS = [
  { to: "/admin/import", label: "Import" },
  { to: "/admin/fee-rules", label: "Fee rules" },
  { to: "/admin/settings", label: "Settings" },
  { to: "/admin/orders", label: "Orders" },
  { to: "/admin/jobs", label: "Jobs & logs" },
  { to: "/admin/stub", label: "Stub" },
];

/**
 * Token gate. The token the operator types is held in sessionStorage and sent
 * as a bearer header by apiFetch for every /api/admin path, so no admin page
 * has to think about auth. Closing the tab logs out.
 */
export function Admin() {
  const [token, setToken] = useState(() => getAdminToken());
  const [draft, setDraft] = useState("");

  if (!token) {
    return (
      <div className={styles.gate}>
        <div className="card">
          <h1 className={styles.gateTitle}>Admin console</h1>
          <p className="small muted" style={{ marginBottom: 16 }}>
            Enter the admin API token. It is kept for this browser tab only and is never written to
            disk.
          </p>
          <form
            className="stack"
            onSubmit={(e) => {
              e.preventDefault();
              const t = draft.trim();
              if (!t) return;
              setAdminToken(t);
              setToken(t);
              setDraft("");
            }}
          >
            <Field label="API token" required>
              {(id) => (
                <input
                  id={id}
                  className="input"
                  type="password"
                  autoComplete="off"
                  autoFocus
                  value={draft}
                  onChange={(e) => setDraft(e.target.value)}
                />
              )}
            </Field>
            <button className="btn btnPrimary" type="submit" disabled={!draft.trim()}>
              Unlock
            </button>
          </form>
        </div>
      </div>
    );
  }

  return (
    <div className="stack">
      <div className={styles.head}>
        <h1 className={styles.title}>Admin console</h1>
        <button
          className="btn btnGhost"
          onClick={() => {
            setAdminToken("");
            setToken("");
          }}
        >
          Lock
        </button>
      </div>

      <nav className={styles.tabs}>
        {TABS.map((tab) => (
          <NavLink
            key={tab.to}
            to={tab.to}
            className={({ isActive }) => (isActive ? `${styles.tab} ${styles.tabActive}` : styles.tab)}
          >
            {tab.label}
          </NavLink>
        ))}
      </nav>

      <Outlet />
    </div>
  );
}
