import styles from "./Pager.module.css";
import { num } from "../lib/fmt";

/** Windowed page numbers: 1 … 4 5 [6] 7 8 … 20 */
function pageWindow(page: number, pages: number): Array<number | "gap"> {
  if (pages <= 7) return Array.from({ length: pages }, (_, i) => i + 1);
  const out: Array<number | "gap"> = [1];
  const from = Math.max(2, page - 1);
  const to = Math.min(pages - 1, page + 1);
  if (from > 2) out.push("gap");
  for (let p = from; p <= to; p++) out.push(p);
  if (to < pages - 1) out.push("gap");
  out.push(pages);
  return out;
}

export function Pager({
  page,
  size,
  total,
  onPage,
}: {
  page: number;
  size: number;
  total: number;
  onPage: (page: number) => void;
}) {
  const pages = size > 0 ? Math.max(1, Math.ceil(total / size)) : 1;
  if (pages <= 1) {
    return total > 0 ? <p className={styles.summary}>{num(total)} results</p> : null;
  }

  const first = (page - 1) * size + 1;
  const last = Math.min(page * size, total);

  return (
    <nav className={styles.pager} aria-label="Pagination">
      <p className={styles.summary}>
        {num(first)}–{num(last)} of {num(total)}
      </p>
      <div className={styles.controls}>
        <button className="btn" onClick={() => onPage(page - 1)} disabled={page <= 1}>
          Previous
        </button>
        <ul className={styles.pages}>
          {pageWindow(page, pages).map((p, i) =>
            p === "gap" ? (
              <li key={`gap-${i}`} className={styles.gap} aria-hidden="true">
                …
              </li>
            ) : (
              <li key={p}>
                <button
                  className={p === page ? `${styles.page} ${styles.current}` : styles.page}
                  aria-current={p === page ? "page" : undefined}
                  onClick={() => onPage(p)}
                >
                  {p}
                </button>
              </li>
            ),
          )}
        </ul>
        <button className="btn" onClick={() => onPage(page + 1)} disabled={page >= pages}>
          Next
        </button>
      </div>
    </nav>
  );
}
