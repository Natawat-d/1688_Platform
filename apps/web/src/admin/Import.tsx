import { useEffect, useState } from "react";
import { api, errorText } from "../lib/api";
import { Field, Readout } from "../components/Field";
import { StatusBadge } from "../components/StatusBadge";
import type { ImportJobRequest, Job } from "../lib/types";
import { dateTimeText, num } from "../lib/fmt";
import shell from "./Admin.module.css";
import styles from "./Import.module.css";

type Mode = "keyword" | "ids" | "all";

const DONE = /^(done|complete|completed|success|succeeded|failed|error|cancelled|canceled)$/i;

export function Import() {
  const [mode, setMode] = useState<Mode>("keyword");
  const [keyword, setKeyword] = useState("");
  const [ids, setIds] = useState("");
  const [starting, setStarting] = useState(false);
  const [error, setError] = useState<string | undefined>();
  const [jobId, setJobId] = useState<string | null>(null);
  const [job, setJob] = useState<Job | null>(null);

  // Watch the job until it settles. Two seconds is short enough to feel live
  // and long enough not to hammer the API during a large catalogue import.
  useEffect(() => {
    if (!jobId) return;
    let live = true;
    const tick = () => {
      api
        .get<Job>(`/api/admin/import/${encodeURIComponent(jobId)}`)
        .then((j) => {
          if (!live) return;
          setJob(j);
          if (DONE.test(j.status ?? "")) window.clearInterval(id);
        })
        .catch((e) => {
          if (!live) return;
          setError(errorText(e));
          window.clearInterval(id);
        });
    };
    const id = window.setInterval(tick, 2000);
    tick();
    return () => {
      live = false;
      window.clearInterval(id);
    };
  }, [jobId]);

  /** Ids may be pasted comma-, space- or newline-separated. Never parsed to numbers. */
  const idList = ids
    .split(/[\s,;]+/)
    .map((s) => s.trim())
    .filter(Boolean);

  const valid = mode === "keyword" ? keyword.trim() !== "" : mode === "ids" ? idList.length > 0 : true;

  async function start() {
    setStarting(true);
    setError(undefined);
    setJob(null);
    setJobId(null);
    const body: ImportJobRequest =
      mode === "keyword" ? { keyword: keyword.trim() } : mode === "ids" ? { offerIds: idList } : { all: true };
    try {
      const res = await api.post<{ jobId: string }>("/api/admin/import", body);
      setJobId(res.jobId);
    } catch (err) {
      setError(errorText(err));
    } finally {
      setStarting(false);
    }
  }

  const pct =
    job && typeof job.total === "number" && job.total > 0 && typeof job.done === "number"
      ? Math.min(100, Math.round((job.done / job.total) * 100))
      : typeof job?.progress === "number"
        ? Math.min(100, Math.round(job.progress <= 1 ? job.progress * 100 : job.progress))
        : null;

  return (
    <div className="stack">
      <div className={shell.panelHead}>
        <div>
          <h2 className={shell.panelTitle}>Import catalogue</h2>
          <p className={shell.panelNote}>
            Pulls offers from 1688 into our local catalogue. Products only become visible to shoppers
            once they are imported.
          </p>
        </div>
      </div>

      <section className="card">
        <div className={styles.modes} role="radiogroup" aria-label="Import source">
          {(
            [
              ["keyword", "By keyword", "Runs a keyword search and imports the results."],
              ["ids", "By offer ID", "Imports specific offers you already know."],
              ["all", "Refresh all", "Re-syncs everything already in the catalogue."],
            ] as Array<[Mode, string, string]>
          ).map(([value, label, hint]) => (
            <button
              key={value}
              type="button"
              role="radio"
              aria-checked={mode === value}
              className={mode === value ? `${styles.mode} ${styles.modeActive}` : styles.mode}
              onClick={() => setMode(value)}
            >
              <span className="strong small">{label}</span>
              <span className="xs faint">{hint}</span>
            </button>
          ))}
        </div>

        <hr className="divider" />

        {mode === "keyword" && (
          <Field label="Keyword" hint="Chinese or English; the gateway translates the query.">
            {(id) => (
              <input
                id={id}
                className="input"
                value={keyword}
                placeholder="e.g. 手机壳 or phone case"
                onChange={(e) => setKeyword(e.target.value)}
              />
            )}
          </Field>
        )}

        {mode === "ids" && (
          <Field
            label="Offer IDs"
            hint={`Separated by commas, spaces or newlines. ${idList.length > 0 ? `${num(idList.length)} recognised.` : ""}`}
          >
            {(id) => (
              <textarea
                id={id}
                className="textarea"
                value={ids}
                placeholder="671234567890&#10;671234567891"
                onChange={(e) => setIds(e.target.value)}
              />
            )}
          </Field>
        )}

        {mode === "all" && (
          <p className="notice noticeWarn">
            Re-syncs every product in the catalogue against 1688. This can take a long time and uses a
            lot of API quota.
          </p>
        )}

        <div className={shell.toolbar} style={{ marginTop: 16 }}>
          <button className="btn btnPrimary" onClick={start} disabled={!valid || starting}>
            {starting ? "Starting…" : "Start import"}
          </button>
          {jobId && <span className="mono">{jobId}</span>}
        </div>

        {error && (
          <p className="notice noticeError" style={{ marginTop: 12 }}>
            {error}
          </p>
        )}
      </section>

      {job && (
        <section className="card">
          <div className="rowBetween" style={{ marginBottom: 12 }}>
            <h3 className={shell.panelTitle}>Job progress</h3>
            <StatusBadge status={job.status} />
          </div>

          {pct !== null && (
            <div className={styles.progress}>
              <div
                className={styles.progressTrack}
                role="progressbar"
                aria-valuenow={pct}
                aria-valuemin={0}
                aria-valuemax={100}
              >
                <div className={styles.progressBar} style={{ width: `${pct}%` }} />
              </div>
              <span className={styles.progressLabel}>{pct}%</span>
            </div>
          )}

          <Readout label="Job ID" mono>
            {job.id || jobId}
          </Readout>
          {job.kind && <Readout label="Kind">{job.kind}</Readout>}
          {typeof job.done === "number" && (
            <Readout label="Processed">
              {num(job.done)}
              {typeof job.total === "number" && job.total > 0 ? ` of ${num(job.total)}` : ""}
            </Readout>
          )}
          {job.startedAt && <Readout label="Started">{dateTimeText(job.startedAt)}</Readout>}
          {job.finishedAt && <Readout label="Finished">{dateTimeText(job.finishedAt)}</Readout>}
          {job.detail && <Readout label="Detail">{job.detail}</Readout>}
          {job.error && (
            <p className="notice noticeError" style={{ marginTop: 12 }}>
              {job.error}
            </p>
          )}
        </section>
      )}
    </div>
  );
}
