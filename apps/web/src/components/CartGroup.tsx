import { Link } from "react-router-dom";
import type { CartGroupData } from "../lib/types";
import { Money } from "./Money";
import { QtyStepper } from "./QtyStepper";
import { PriceBreakdown } from "./PriceBreakdown";
import { pluralise } from "../lib/fmt";
import styles from "./CartGroup.module.css";

/**
 * One supplier's worth of cart. Each group becomes one 1688 order, which is
 * why the parcel count and the issue list belong at this level rather than on
 * the cart as a whole.
 */
export function CartGroup({
  group,
  busy,
  onQty,
  onRemove,
}: {
  group: CartGroupData;
  busy: boolean;
  onQty: (lineId: string, quantity: number) => void;
  onRemove: (lineId: string) => void;
}) {
  const issuesByLine = new Map<string, string[]>();
  const groupIssues: string[] = [];
  for (const issue of group.issues ?? []) {
    if (issue.field && group.lines.some((l) => l.id === issue.field)) {
      const list = issuesByLine.get(issue.field) ?? [];
      list.push(issue.message);
      issuesByLine.set(issue.field, list);
    } else {
      groupIssues.push(issue.message);
    }
  }

  return (
    <section className={styles.group}>
      <header className={styles.head}>
        <div>
          <h3 className={styles.seller}>{group.sellerName || "Supplier"}</h3>
          <p className={styles.sellerMeta}>
            {pluralise(group.parcels || 1, "parcel")} · {pluralise(group.lines.length, "line")}
          </p>
        </div>
        <div className={styles.subtotal}>
          <span className="xs faint">Subtotal</span>
          <Money value={group.subtotal} size="md" />
        </div>
      </header>

      {groupIssues.length > 0 && (
        <ul className={styles.issues}>
          {groupIssues.map((m, i) => (
            <li key={i} className="notice noticeWarn">
              {m}
            </li>
          ))}
        </ul>
      )}

      <ul className={styles.lines}>
        {group.lines.map((line) => {
          const lineIssues = issuesByLine.get(line.id) ?? [];
          return (
            <li key={line.id} className={styles.line}>
              <Link className={styles.thumbLink} to={`/p/${encodeURIComponent(line.offerId)}`}>
                {line.image ? (
                  <img className={styles.thumb} src={line.image} alt="" loading="lazy" />
                ) : (
                  <span className={styles.thumbPlaceholder} aria-hidden="true" />
                )}
              </Link>

              <div className={styles.lineBody}>
                <Link className={styles.lineTitle} to={`/p/${encodeURIComponent(line.offerId)}`}>
                  {line.title}
                </Link>
                {line.skuLabel && <p className={styles.skuLabel}>{line.skuLabel}</p>}
                <div className={styles.controls}>
                  <QtyStepper
                    value={line.quantity}
                    onChange={(q) => onQty(line.id, q)}
                    max={line.stock}
                    disabled={busy}
                    compact
                  />
                  <button
                    type="button"
                    className="btn btnGhost btnDanger"
                    onClick={() => onRemove(line.id)}
                    disabled={busy}
                  >
                    Remove
                  </button>
                </div>
                {lineIssues.map((m, i) => (
                  <p key={i} className={styles.lineIssue}>
                    {m}
                  </p>
                ))}
                <div className={styles.breakdownSlot}>
                  <PriceBreakdown
                    breakdown={line.breakdown}
                    quantity={line.quantity}
                    unit={line.unit}
                    total={line.line}
                    title="Price detail"
                  />
                </div>
              </div>

              <div className={styles.lineMoney}>
                <Money value={line.line} size="md" />
                <span className="xs faint">
                  <Money value={line.unit} size="sm" muted /> each
                </span>
              </div>
            </li>
          );
        })}
      </ul>
    </section>
  );
}
