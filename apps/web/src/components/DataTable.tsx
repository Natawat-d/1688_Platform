import type { ReactNode } from "react";
import styles from "./DataTable.module.css";

export type Column<T> = {
  key: string;
  header: string;
  render: (row: T, index: number) => ReactNode;
  /** Right-align and tabular-align a numeric column. */
  numeric?: boolean;
  /** Hide below 640px, for columns that are nice-to-have. */
  secondary?: boolean;
  width?: string;
};

/**
 * The admin console's one table. Scrolls horizontally inside its own box so a
 * wide log never makes the page scroll sideways.
 */
export function DataTable<T>({
  rows,
  columns,
  rowKey,
  empty = "Nothing to show.",
  loading = false,
  onRowClick,
}: {
  rows: T[] | undefined;
  columns: Column<T>[];
  rowKey: (row: T, index: number) => string;
  empty?: ReactNode;
  loading?: boolean;
  onRowClick?: (row: T) => void;
}) {
  if (loading && !rows) {
    return (
      <div className={styles.wrap}>
        <div className={styles.loading}>
          {[0, 1, 2, 3].map((i) => (
            <div key={i} className="skeleton" style={{ height: 32 }} />
          ))}
        </div>
      </div>
    );
  }

  if (!rows || rows.length === 0) {
    return <div className="empty">{empty}</div>;
  }

  return (
    <div className={styles.wrap}>
      <table className={styles.table}>
        <thead>
          <tr>
            {columns.map((c) => (
              <th
                key={c.key}
                scope="col"
                style={c.width ? { width: c.width } : undefined}
                className={[c.numeric ? styles.numeric : "", c.secondary ? styles.secondary : ""]
                  .filter(Boolean)
                  .join(" ")}
              >
                {c.header}
              </th>
            ))}
          </tr>
        </thead>
        <tbody>
          {rows.map((row, i) => (
            <tr
              key={rowKey(row, i)}
              onClick={onRowClick ? () => onRowClick(row) : undefined}
              className={onRowClick ? styles.clickable : undefined}
            >
              {columns.map((c) => (
                <td
                  key={c.key}
                  className={[c.numeric ? styles.numeric : "", c.secondary ? styles.secondary : ""]
                    .filter(Boolean)
                    .join(" ")}
                >
                  {c.render(row, i)}
                </td>
              ))}
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}
