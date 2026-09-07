import type { Money as MoneyType } from "../lib/types";
import styles from "./Money.module.css";

/**
 * The only way money reaches the screen. The server has already rendered the
 * string; this component picks a size and never touches the number.
 */
export function Money({
  value,
  size = "md",
  muted = false,
  strike = false,
}: {
  value: MoneyType | null | undefined;
  size?: "sm" | "md" | "lg" | "xl";
  muted?: boolean;
  strike?: boolean;
}) {
  const text = value?.text ?? "—";
  const cls = [styles.money, styles[size], muted ? styles.muted : "", strike ? styles.strike : ""]
    .filter(Boolean)
    .join(" ");
  return (
    <span className={cls} title={value ? `${value.minor} minor units ${value.currency}` : undefined}>
      {text}
    </span>
  );
}

/** "from ¥12.00" style prefix used on cards and price ranges. */
export function MoneyFrom({ value, size = "md" }: { value: MoneyType | null | undefined; size?: "sm" | "md" | "lg" }) {
  return (
    <span className={styles.from}>
      <span className={styles.fromLabel}>from</span>
      <Money value={value} size={size} />
    </span>
  );
}
