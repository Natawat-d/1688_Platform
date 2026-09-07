import { useEffect, useState } from "react";
import { useQuery } from "../lib/useQuery";
import { api, errorText } from "../lib/api";
import { Field } from "../components/Field";
import type { Settings as SettingsType } from "../lib/types";
import { labelise } from "../lib/fmt";
import shell from "./Admin.module.css";
import styles from "./Settings.module.css";

/**
 * The form is generated from whatever the server sends rather than from a
 * schema hard-coded here. New settings appear in the UI the moment the backend
 * returns them, and this page never silently drops a key it does not know.
 *
 * Nested objects and arrays fall back to a JSON textarea, which is honest about
 * what is being edited instead of pretending to understand the shape.
 */

type Draft = Record<string, string | boolean>;

function toDraft(settings: SettingsType): Draft {
  const out: Draft = {};
  for (const [k, v] of Object.entries(settings ?? {})) {
    if (typeof v === "boolean") out[k] = v;
    else if (v === null || v === undefined) out[k] = "";
    else if (typeof v === "object") out[k] = JSON.stringify(v, null, 2);
    else out[k] = String(v);
  }
  return out;
}

/** Rebuild the payload with each value's original JSON type. */
function fromDraft(draft: Draft, original: SettingsType): SettingsType {
  const out: SettingsType = {};
  for (const [k, raw] of Object.entries(draft)) {
    const was = (original ?? {})[k];
    if (typeof raw === "boolean") {
      out[k] = raw;
    } else if (typeof was === "number") {
      const n = Number(raw);
      out[k] = Number.isFinite(n) ? n : raw;
    } else if (was !== null && typeof was === "object") {
      try {
        out[k] = JSON.parse(raw);
      } catch {
        out[k] = was; // keep the last good value rather than send broken JSON
      }
    } else {
      out[k] = raw;
    }
  }
  return out;
}

function isJsonValue(v: unknown): boolean {
  return v !== null && typeof v === "object";
}

export function Settings() {
  const { data, error, loading, reload } = useQuery<SettingsType>("/api/admin/settings");
  const [draft, setDraft] = useState<Draft>({});
  const [saving, setSaving] = useState(false);
  const [saveError, setSaveError] = useState<string | undefined>();
  const [saved, setSaved] = useState(false);

  useEffect(() => {
    if (data) setDraft(toDraft(data));
  }, [data]);

  async function save() {
    if (!data) return;
    setSaving(true);
    setSaveError(undefined);
    setSaved(false);
    try {
      await api.put("/api/admin/settings", fromDraft(draft, data));
      setSaved(true);
      reload();
    } catch (err) {
      setSaveError(errorText(err));
    } finally {
      setSaving(false);
    }
  }

  const keys = Object.keys(draft).sort();

  return (
    <div className="stack">
      <div className={shell.panelHead}>
        <div>
          <h2 className={shell.panelTitle}>Settings</h2>
          <p className={shell.panelNote}>
            Platform configuration: FX rate, rounding, snapshot window, consolidation address and the
            rest. Fields are generated from the server's own payload.
          </p>
        </div>
        <button className="btn" onClick={reload} disabled={loading}>
          Reload
        </button>
      </div>

      {error && <p className="notice noticeError">{error}</p>}
      {saveError && <p className="notice noticeError">{saveError}</p>}
      {saved && !saveError && <p className="notice noticeOk">Settings saved.</p>}

      {loading && !data ? (
        <div className="card">
          <div className="stack">
            {[0, 1, 2, 3, 4].map((i) => (
              <div key={i} className="skeleton" style={{ height: 56 }} />
            ))}
          </div>
        </div>
      ) : keys.length === 0 ? (
        <div className="empty">The server returned no settings.</div>
      ) : (
        <section className="card">
          <form
            className={styles.grid}
            onSubmit={(e) => {
              e.preventDefault();
              void save();
            }}
          >
            {keys.map((key) => {
              const value = draft[key];
              const wasJson = isJsonValue((data ?? {})[key]);

              if (typeof value === "boolean") {
                return (
                  <label key={key} className={`checkRow ${styles.full}`}>
                    <input
                      type="checkbox"
                      checked={value}
                      onChange={(e) => setDraft((d) => ({ ...d, [key]: e.target.checked }))}
                    />
                    {labelise(key)}
                    <code className={styles.keyName}>{key}</code>
                  </label>
                );
              }

              if (wasJson) {
                return (
                  <div key={key} className={styles.full}>
                    <Field label={labelise(key)} hint={`${key} — JSON`}>
                      {(id) => (
                        <textarea
                          id={id}
                          className="textarea"
                          value={value}
                          spellCheck={false}
                          onChange={(e) => setDraft((d) => ({ ...d, [key]: e.target.value }))}
                        />
                      )}
                    </Field>
                  </div>
                );
              }

              return (
                <Field key={key} label={labelise(key)} hint={key}>
                  {(id) => (
                    <input
                      id={id}
                      className="input"
                      value={value}
                      spellCheck={false}
                      onChange={(e) => setDraft((d) => ({ ...d, [key]: e.target.value }))}
                    />
                  )}
                </Field>
              );
            })}

            <div className={`${styles.full} ${shell.toolbar}`}>
              <button className="btn btnPrimary" type="submit" disabled={saving}>
                {saving ? "Saving…" : "Save settings"}
              </button>
              <button
                className="btn"
                type="button"
                disabled={saving || !data}
                onClick={() => data && setDraft(toDraft(data))}
              >
                Revert
              </button>
            </div>
          </form>
        </section>
      )}
    </div>
  );
}
