import { useState } from "react";
import type { Breakdown, Money as MoneyType } from "../lib/types";
import { Money } from "./Money";
import { fenText, ppmText, num } from "../lib/fmt";
import styles from "./PriceBreakdown.module.css";

/**
 * "How this price is made".
 *
 * Every CNY component is rendered from its integer fen value by string
 * surgery, so what is on screen is exactly what the pricing engine computed —
 * which makes this panel the fastest way to find a pricing bug, as well as the
 * trust feature it looks like.
 */
export function PriceBreakdown({
  breakdown,
  quantity,
  unit,
  total,
  open: openProp,
  collapsible = true,
  title = "How this price is made",
}: {
  breakdown: Breakdown | null | undefined;
  quantity?: number;
  unit?: MoneyType;
  total?: MoneyType;
  open?: boolean;
  collapsible?: boolean;
  title?: string;
}) {
  const [open, setOpen] = useState(Boolean(openProp));
  if (!breakdown) return null;

  const body = (
    <div className={styles.body}>
      <p className={styles.intro}>
        Built per unit in CNY on the 1688 side, then converted once to THB. Amounts shown are the
        integer values recorded with your order.
      </p>
      <dl className={styles.rows}>
        <Row label="1688 item price" value={fenText(breakdown.baseFen)} note="the SKU price at your quantity tier" />
        <Row label="China freight" value={fenText(breakdown.freightFen)} note="supplier to our consolidation point" />
        <Row label="International shipping" value={fenText(breakdown.intlFen)} note="our rate card, by weight and volume" />
        <Row label="Platform fee" value={fenText(breakdown.feeFen)} note={breakdown.feeRuleId ? `rule ${breakdown.feeRuleId}` : undefined} />
        <div className={styles.ruleRow}>
          <dt>FX rate applied</dt>
          <dd>
            <span className="nums">{ppmText(breakdown.fxPpm)}</span>
            <span className={styles.note}>THB per CNY</span>
          </dd>
        </div>
      </dl>
      {(unit || total) && (
        <div className={styles.totals}>
          {unit && (
            <div className={styles.totalRow}>
              <span>Per unit</span>
              <Money value={unit} size="sm" />
            </div>
          )}
          {total && (
            <div className={`${styles.totalRow} ${styles.grand}`}>
              <span>{quantity ? `Total for ${num(quantity)}` : "Total"}</span>
              <Money value={total} size="md" />
            </div>
          )}
        </div>
      )}
    </div>
  );

  if (!collapsible) {
    return (
      <section className={styles.panel}>
        <h3 className={styles.staticTitle}>{title}</h3>
        {body}
      </section>
    );
  }

  return (
    <section className={styles.panel}>
      <button
        type="button"
        className={styles.toggle}
        onClick={() => setOpen((o) => !o)}
        aria-expanded={open}
      >
        <span>{title}</span>
        <span className={open ? `${styles.chevron} ${styles.chevronOpen}` : styles.chevron} aria-hidden="true">
          ▾
        </span>
      </button>
      {open && body}
    </section>
  );
}

function Row({ label, value, note }: { label: string; value: string; note?: string }) {
  return (
    <div className={styles.ruleRow}>
      <dt>
        {label}
        {note && <span className={styles.note}>{note}</span>}
      </dt>
      <dd className="nums">{value}</dd>
    </div>
  );
}
