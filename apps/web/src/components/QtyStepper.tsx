import { useEffect, useState } from "react";
import { num } from "../lib/fmt";
import styles from "./QtyStepper.module.css";

/**
 * Quantity input that only ever emits a quantity 1688 will accept.
 *
 * `moq` is minOrderQuantity and `batch` is batchNumber: an order must be at
 * least the MOQ and, when batchNumber > 1, a whole multiple of it. Free typing
 * is allowed while the field has focus; the value is snapped on blur so the
 * user is not fighting the input mid-keystroke.
 */
export function snapQuantity(value: number, moq: number, batch: number, max?: number): number {
  const step = batch > 1 ? batch : 1;
  const floor = Math.max(1, moq > 0 ? moq : 1);
  let q = Number.isFinite(value) ? Math.trunc(value) : floor;
  if (q < floor) q = floor;
  if (step > 1) {
    // Round up to the next whole batch, measured from zero.
    const rem = q % step;
    if (rem !== 0) q = q + (step - rem);
  }
  if (max !== undefined && max > 0 && q > max) {
    let capped = max;
    if (step > 1) capped = capped - (capped % step);
    if (capped >= floor) q = capped;
  }
  return q;
}

export function QtyStepper({
  value,
  onChange,
  moq = 1,
  batch = 1,
  max,
  unit = "piece",
  disabled = false,
  compact = false,
}: {
  value: number;
  onChange: (quantity: number) => void;
  moq?: number;
  batch?: number;
  max?: number;
  unit?: string;
  disabled?: boolean;
  compact?: boolean;
}) {
  const step = batch > 1 ? batch : 1;
  const [draft, setDraft] = useState(String(value));

  useEffect(() => {
    setDraft(String(value));
  }, [value]);

  const commit = (raw: string) => {
    const parsed = Number(raw.replace(/[^0-9]/g, ""));
    const next = snapQuantity(Number.isFinite(parsed) && parsed > 0 ? parsed : moq, moq, batch, max);
    setDraft(String(next));
    if (next !== value) onChange(next);
  };

  const bump = (delta: number) => {
    const next = snapQuantity(value + delta, moq, batch, max);
    if (next !== value) onChange(next);
    else setDraft(String(value));
  };

  const atMin = value <= Math.max(1, moq);
  const atMax = max !== undefined && max > 0 && value + step > max;

  return (
    <div className={compact ? `${styles.stepper} ${styles.compact}` : styles.stepper}>
      <button
        type="button"
        className={styles.button}
        onClick={() => bump(-step)}
        disabled={disabled || atMin}
        aria-label={`Decrease by ${step}`}
      >
        −
      </button>
      <input
        className={styles.input}
        type="text"
        inputMode="numeric"
        value={draft}
        disabled={disabled}
        aria-label={`Quantity in ${unit}`}
        onChange={(e) => setDraft(e.target.value)}
        onBlur={(e) => commit(e.target.value)}
        onKeyDown={(e) => {
          if (e.key === "Enter") {
            e.preventDefault();
            commit((e.target as HTMLInputElement).value);
          }
        }}
      />
      <button
        type="button"
        className={styles.button}
        onClick={() => bump(step)}
        disabled={disabled || atMax}
        aria-label={`Increase by ${step}`}
      >
        +
      </button>
      {!compact && (
        <span className={styles.rule}>
          {unit}
          {moq > 1 && ` · min ${num(moq)}`}
          {step > 1 && ` · in ${num(step)}s`}
        </span>
      )}
    </div>
  );
}
