import { useState } from "react";
import { useQuery } from "../lib/useQuery";
import { api, errorText } from "../lib/api";
import { DataTable } from "../components/DataTable";
import type { Column } from "../components/DataTable";
import { Field, FieldRow, Readout } from "../components/Field";
import { PriceBreakdown } from "../components/PriceBreakdown";
import { Money } from "../components/Money";
import type { FeeRule, FeeRuleScope, QuotePreview } from "../lib/types";
import { bpsText, dateText, fenText, num, titleCase } from "../lib/fmt";
import shell from "./Admin.module.css";
import styles from "./FeeRules.module.css";

const SCOPES: FeeRuleScope[] = ["global", "category", "supplier", "product"];

const BLANK: FeeRule = {
  id: "",
  name: "",
  scope: "global",
  scopeValue: "",
  feeBps: 0,
  feeFixedFen: "0",
  minFeeFen: "0",
  priority: 100,
  effectiveFrom: "",
  effectiveTo: "",
  active: true,
};

export function FeeRules() {
  const rules = useQuery<FeeRule[]>("/api/admin/fee-rules");
  const [editing, setEditing] = useState<FeeRule | null>(null);
  const [saving, setSaving] = useState(false);
  const [error, setError] = useState<string | undefined>();

  async function save(rule: FeeRule) {
    setSaving(true);
    setError(undefined);
    try {
      // Spread the original record first so any server-side field this UI does
      // not know about survives the round-trip untouched.
      const body = { ...rule, feeBps: Number(rule.feeBps) || 0, priority: Number(rule.priority) || 0 };
      if (rule.id) await api.put(`/api/admin/fee-rules/${encodeURIComponent(rule.id)}`, body);
      else await api.post("/api/admin/fee-rules", body);
      setEditing(null);
      rules.reload();
    } catch (err) {
      setError(errorText(err));
    } finally {
      setSaving(false);
    }
  }

  async function remove(rule: FeeRule) {
    if (!window.confirm(`Delete fee rule "${rule.name || rule.id}"? Orders already placed keep the fee they were charged.`)) {
      return;
    }
    setError(undefined);
    try {
      await api.del(`/api/admin/fee-rules/${encodeURIComponent(rule.id)}`);
      rules.reload();
    } catch (err) {
      setError(errorText(err));
    }
  }

  const columns: Column<FeeRule>[] = [
    {
      key: "name",
      header: "Rule",
      render: (r) => (
        <div>
          <div className="strong">{r.name || "Untitled"}</div>
          <div className="xs faint mono">{r.id}</div>
        </div>
      ),
    },
    {
      key: "scope",
      header: "Scope",
      render: (r) => (
        <div>
          <span className="badge">{titleCase(r.scope)}</span>
          {r.scopeValue && <div className="xs faint mono">{r.scopeValue}</div>}
        </div>
      ),
    },
    { key: "pct", header: "Percent", numeric: true, render: (r) => bpsText(r.feeBps) },
    { key: "fixed", header: "Fixed", numeric: true, render: (r) => fenText(r.feeFixedFen) },
    { key: "min", header: "Minimum", numeric: true, secondary: true, render: (r) => fenText(r.minFeeFen) },
    { key: "priority", header: "Priority", numeric: true, secondary: true, render: (r) => num(r.priority) },
    {
      key: "window",
      header: "Effective",
      secondary: true,
      render: (r) =>
        r.effectiveFrom || r.effectiveTo ? (
          <span className="xs">
            {r.effectiveFrom ? dateText(r.effectiveFrom) : "—"} → {r.effectiveTo ? dateText(r.effectiveTo) : "open"}
          </span>
        ) : (
          <span className="xs faint">Always</span>
        ),
    },
    {
      key: "active",
      header: "State",
      render: (r) => <span className={r.active ? "badge badgeOk" : "badge"}>{r.active ? "Active" : "Off"}</span>,
    },
    {
      key: "actions",
      header: "",
      render: (r) => (
        <div className={styles.rowActions}>
          <button className="btn btnGhost" onClick={() => setEditing({ ...r })}>
            Edit
          </button>
          <button className="btn btnGhost btnDanger" onClick={() => void remove(r)}>
            Delete
          </button>
        </div>
      ),
    },
  ];

  return (
    <div className="stack">
      <div className={shell.panelHead}>
        <div>
          <h2 className={shell.panelTitle}>Fee rules</h2>
          <p className={shell.panelNote}>
            The highest-priority rule matching a product decides the platform fee. Every order line
            records the rule it used, so editing a rule never rewrites history.
          </p>
        </div>
        <button className="btn btnPrimary" onClick={() => setEditing({ ...BLANK })}>
          New rule
        </button>
      </div>

      {error && <p className="notice noticeError">{error}</p>}
      {rules.error && <p className="notice noticeError">{rules.error}</p>}

      <DataTable
        rows={rules.data}
        columns={columns}
        loading={rules.loading}
        rowKey={(r, i) => r.id || `new-${i}`}
        empty="No fee rules yet. Add a global rule so every product has a fee."
      />

      {editing && (
        <RuleEditor
          rule={editing}
          saving={saving}
          onChange={setEditing}
          onCancel={() => setEditing(null)}
          onSave={() => void save(editing)}
        />
      )}

      <Preview />
    </div>
  );
}

function RuleEditor({
  rule,
  saving,
  onChange,
  onCancel,
  onSave,
}: {
  rule: FeeRule;
  saving: boolean;
  onChange: (r: FeeRule) => void;
  onCancel: () => void;
  onSave: () => void;
}) {
  const set = (patch: Partial<FeeRule>) => onChange({ ...rule, ...patch });
  const digits = (v: string) => v.replace(/[^0-9]/g, "");

  return (
    <section className="card">
      <h3 className={shell.panelTitle} style={{ marginBottom: 16 }}>
        {rule.id ? "Edit rule" : "New rule"}
      </h3>
      <form
        className="stack"
        onSubmit={(e) => {
          e.preventDefault();
          onSave();
        }}
      >
        <FieldRow>
          <Field label="Name" required>
            {(id) => (
              <input id={id} className="input" value={rule.name} onChange={(e) => set({ name: e.target.value })} />
            )}
          </Field>
          <Field label="Scope">
            {(id) => (
              <select
                id={id}
                className="select"
                value={rule.scope}
                onChange={(e) => set({ scope: e.target.value as FeeRuleScope, scopeValue: "" })}
              >
                {SCOPES.map((s) => (
                  <option key={s} value={s}>
                    {titleCase(s)}
                  </option>
                ))}
              </select>
            )}
          </Field>
          <Field
            label={
              rule.scope === "category"
                ? "Category ID"
                : rule.scope === "supplier"
                  ? "Supplier open ID"
                  : rule.scope === "product"
                    ? "Offer ID"
                    : "Applies to"
            }
            hint={rule.scope === "global" ? "Global rules match everything." : "Ids are strings; paste them exactly."}
          >
            {(id) => (
              <input
                id={id}
                className="input"
                value={rule.scopeValue}
                disabled={rule.scope === "global"}
                placeholder={rule.scope === "global" ? "Everything" : ""}
                onChange={(e) => set({ scopeValue: e.target.value.trim() })}
              />
            )}
          </Field>
        </FieldRow>

        <FieldRow>
          <Field label="Percent of subtotal" hint={`Basis points. ${bpsText(rule.feeBps)}`}>
            {(id) => (
              <input
                id={id}
                className="input"
                inputMode="numeric"
                value={String(rule.feeBps ?? 0)}
                onChange={(e) => set({ feeBps: Number(digits(e.target.value) || "0") })}
              />
            )}
          </Field>
          <Field label="Fixed fee" hint={`CNY fen. ${fenText(rule.feeFixedFen)}`}>
            {(id) => (
              <input
                id={id}
                className="input"
                inputMode="numeric"
                value={String(rule.feeFixedFen ?? "0")}
                onChange={(e) => set({ feeFixedFen: digits(e.target.value) || "0" })}
              />
            )}
          </Field>
          <Field label="Minimum fee" hint={`CNY fen. ${fenText(rule.minFeeFen)}`}>
            {(id) => (
              <input
                id={id}
                className="input"
                inputMode="numeric"
                value={String(rule.minFeeFen ?? "0")}
                onChange={(e) => set({ minFeeFen: digits(e.target.value) || "0" })}
              />
            )}
          </Field>
          <Field label="Priority" hint="Higher wins.">
            {(id) => (
              <input
                id={id}
                className="input"
                inputMode="numeric"
                value={String(rule.priority ?? 0)}
                onChange={(e) => set({ priority: Number(digits(e.target.value) || "0") })}
              />
            )}
          </Field>
        </FieldRow>

        <FieldRow>
          <Field label="Effective from" hint="Blank means immediately.">
            {(id) => (
              <input
                id={id}
                className="input"
                type="date"
                value={(rule.effectiveFrom || "").slice(0, 10)}
                onChange={(e) => set({ effectiveFrom: e.target.value })}
              />
            )}
          </Field>
          <Field label="Effective to" hint="Blank means open-ended.">
            {(id) => (
              <input
                id={id}
                className="input"
                type="date"
                value={(rule.effectiveTo || "").slice(0, 10)}
                onChange={(e) => set({ effectiveTo: e.target.value })}
              />
            )}
          </Field>
        </FieldRow>

        <label className="checkRow">
          <input type="checkbox" checked={Boolean(rule.active)} onChange={(e) => set({ active: e.target.checked })} />
          Active
        </label>

        <div className={shell.toolbar}>
          <button className="btn btnPrimary" type="submit" disabled={saving || !rule.name.trim()}>
            {saving ? "Saving…" : "Save rule"}
          </button>
          <button className="btn" type="button" onClick={onCancel} disabled={saving}>
            Cancel
          </button>
        </div>
      </form>
    </section>
  );
}

/** "What will this product cost" — the fee editor's sanity check. */
function Preview() {
  const [offerId, setOfferId] = useState("");
  const [skuId, setSkuId] = useState("");
  const [quantity, setQuantity] = useState("1");
  const [quote, setQuote] = useState<QuotePreview | null>(null);
  const [error, setError] = useState<string | undefined>();
  const [busy, setBusy] = useState(false);

  async function run() {
    setBusy(true);
    setError(undefined);
    try {
      const q = await api.post<QuotePreview>("/api/admin/fee-rules/preview", {
        offerId: offerId.trim(),
        skuId: skuId.trim(),
        quantity: Number(quantity.replace(/[^0-9]/g, "")) || 1,
      });
      setQuote(q);
    } catch (err) {
      setQuote(null);
      setError(errorText(err));
    } finally {
      setBusy(false);
    }
  }

  return (
    <section className="card">
      <h3 className={shell.panelTitle} style={{ marginBottom: 4 }}>
        Price preview
      </h3>
      <p className={shell.panelNote} style={{ marginBottom: 16 }}>
        Runs the pricing engine against a real offer and shows every component, so you can see what a
        rule change actually does before shoppers do.
      </p>

      <form
        className="stack"
        onSubmit={(e) => {
          e.preventDefault();
          void run();
        }}
      >
        <FieldRow>
          <Field label="Offer ID" required>
            {(id) => (
              <input
                id={id}
                className="input"
                value={offerId}
                placeholder="671234567890"
                onChange={(e) => setOfferId(e.target.value)}
              />
            )}
          </Field>
          <Field label="SKU ID" hint="Blank uses the offer's default variant.">
            {(id) => <input id={id} className="input" value={skuId} onChange={(e) => setSkuId(e.target.value)} />}
          </Field>
          <Field label="Quantity">
            {(id) => (
              <input
                id={id}
                className="input"
                inputMode="numeric"
                value={quantity}
                onChange={(e) => setQuantity(e.target.value.replace(/[^0-9]/g, ""))}
              />
            )}
          </Field>
        </FieldRow>
        <div>
          <button className="btn btnPrimary" type="submit" disabled={busy || !offerId.trim()}>
            {busy ? "Pricing…" : "Preview price"}
          </button>
        </div>
      </form>

      {error && (
        <p className="notice noticeError" style={{ marginTop: 12 }}>
          {error}
        </p>
      )}

      {quote && (
        <div className={styles.previewResult}>
          <div className={styles.previewHeadline}>
            <div>
              <span className="xs faint">Per unit</span>
              <div>
                <Money value={quote.unit} size="lg" />
              </div>
            </div>
            <div>
              <span className="xs faint">Total for {num(quote.quantity)}</span>
              <div>
                <Money value={quote.total} size="lg" />
              </div>
            </div>
          </div>

          {quote.feeRule && (
            <Readout label="Fee rule applied">
              {quote.feeRule.name} <span className="xs faint">({quote.feeRule.scope})</span>
            </Readout>
          )}

          <PriceBreakdown
            breakdown={quote.breakdown}
            quantity={quote.quantity}
            unit={quote.unit}
            total={quote.total}
            collapsible={false}
            title="Components"
          />

          {(quote.issues ?? []).length > 0 && (
            <ul className="stackSm" style={{ marginTop: 12 }}>
              {quote.issues!.map((issue, i) => (
                <li key={i} className="notice noticeWarn">
                  {issue.message}
                </li>
              ))}
            </ul>
          )}
        </div>
      )}
    </section>
  );
}
