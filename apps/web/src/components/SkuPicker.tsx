import { useMemo } from "react";
import type { Sku } from "../lib/types";
import { Money } from "./Money";
import { num } from "../lib/fmt";
import styles from "./SkuPicker.module.css";

/**
 * Derives the attribute axes from skus[].attrs and disables any value that
 * cannot combine with the current selection.
 *
 * A value is offered when at least one SKU carries it, agrees with every OTHER
 * currently-chosen axis, and is in stock. That "every other axis" rule is what
 * lets a shopper change one axis at a time without first clearing the rest.
 */

export type Axis = {
  name: string;
  values: Array<{ value: string; image?: string }>;
};

export function deriveAxes(skus: Sku[]): Axis[] {
  const order: string[] = [];
  const seen = new Map<string, Map<string, string | undefined>>();
  for (const sku of skus ?? []) {
    for (const attr of sku.attrs ?? []) {
      if (!attr || !attr.name) continue;
      if (!seen.has(attr.name)) {
        seen.set(attr.name, new Map());
        order.push(attr.name);
      }
      const values = seen.get(attr.name)!;
      if (!values.has(attr.value)) values.set(attr.value, attr.image);
      else if (!values.get(attr.value) && attr.image) values.set(attr.value, attr.image);
    }
  }
  return order.map((name) => ({
    name,
    values: Array.from(seen.get(name)!.entries()).map(([value, image]) => ({ value, image })),
  }));
}

function attrValue(sku: Sku, name: string): string | undefined {
  return (sku.attrs ?? []).find((a) => a.name === name)?.value;
}

export function skuFor(skus: Sku[], selection: Record<string, string>, axes: Axis[]): Sku | undefined {
  if (axes.length === 0) return undefined;
  return (skus ?? []).find((sku) => axes.every((ax) => attrValue(sku, ax.name) === selection[ax.name]));
}

function inStock(sku: Sku): boolean {
  return sku.available !== false && (sku.stock ?? 0) > 0;
}

export function SkuPicker({
  skus,
  selectedSkuId,
  onSelect,
}: {
  skus: Sku[];
  selectedSkuId: string | null;
  onSelect: (skuId: string) => void;
}) {
  const axes = useMemo(() => deriveAxes(skus), [skus]);
  const selected = useMemo(() => (skus ?? []).find((s) => s.skuId === selectedSkuId), [skus, selectedSkuId]);

  // Selection derived from the chosen SKU, so the two can never disagree.
  const selection = useMemo(() => {
    const sel: Record<string, string> = {};
    if (selected) for (const ax of axes) sel[ax.name] = attrValue(selected, ax.name) ?? "";
    return sel;
  }, [selected, axes]);

  if (!skus || skus.length === 0) {
    return <p className="muted small">This product has no variants listed.</p>;
  }

  // No attribute data: fall back to a plain list of SKU labels.
  if (axes.length === 0) {
    return (
      <fieldset className={styles.group}>
        <legend className={styles.legend}>Variant</legend>
        <div className={styles.values}>
          {skus.map((sku) => (
            <button
              key={sku.skuId}
              type="button"
              className={valueClass(sku.skuId === selectedSkuId, !inStock(sku))}
              disabled={!inStock(sku)}
              onClick={() => onSelect(sku.skuId)}
            >
              <span className={styles.valueText}>{sku.label || sku.specId || sku.skuId}</span>
              <span className={styles.valuePrice}>
                <Money value={sku.price} size="sm" />
              </span>
            </button>
          ))}
        </div>
      </fieldset>
    );
  }

  /** Is `value` on `axisName` reachable given every other chosen axis? */
  const reachable = (axisName: string, value: string): boolean =>
    skus.some((sku) => {
      if (attrValue(sku, axisName) !== value) return false;
      for (const ax of axes) {
        if (ax.name === axisName) continue;
        const chosen = selection[ax.name];
        if (chosen && attrValue(sku, ax.name) !== chosen) return false;
      }
      return inStock(sku);
    });

  /** Pick the SKU to switch to when this value is clicked. */
  const pick = (axisName: string, value: string) => {
    const candidates = skus.filter((sku) => attrValue(sku, axisName) === value);
    // Prefer one that keeps as many of the other axes as possible.
    let best: Sku | undefined;
    let bestScore = -1;
    for (const sku of candidates) {
      if (!inStock(sku)) continue;
      let score = 0;
      for (const ax of axes) {
        if (ax.name === axisName) continue;
        if (selection[ax.name] && attrValue(sku, ax.name) === selection[ax.name]) score++;
      }
      if (score > bestScore) {
        bestScore = score;
        best = sku;
      }
    }
    if (best) onSelect(best.skuId);
  };

  return (
    <div className={styles.picker}>
      {axes.map((axis) => (
        <fieldset key={axis.name} className={styles.group}>
          <legend className={styles.legend}>
            {axis.name}
            {selection[axis.name] && <span className={styles.chosen}>{selection[axis.name]}</span>}
          </legend>
          <div className={styles.values}>
            {axis.values.map(({ value, image }) => {
              const isSelected = selection[axis.name] === value;
              const ok = reachable(axis.name, value);
              return (
                <button
                  key={value}
                  type="button"
                  className={valueClass(isSelected, !ok)}
                  disabled={!ok && !isSelected}
                  aria-pressed={isSelected}
                  title={ok ? value : `${value} — unavailable with the current selection`}
                  onClick={() => pick(axis.name, value)}
                >
                  {image && <img className={styles.swatch} src={image} alt="" loading="lazy" />}
                  <span className={styles.valueText}>{value}</span>
                </button>
              );
            })}
          </div>
        </fieldset>
      ))}

      {selected && (
        <p className={styles.stockLine}>
          {inStock(selected) ? (
            <>
              <span className="badge badgeOk">In stock</span>
              <span className="muted small">
                {num(selected.stock)} available · {selected.label || selected.specId}
              </span>
            </>
          ) : (
            <span className="badge badgeDanger">Out of stock</span>
          )}
        </p>
      )}
    </div>
  );
}

function valueClass(selected: boolean, disabled: boolean): string {
  return [styles.value, selected ? styles.selected : "", disabled ? styles.disabled : ""].filter(Boolean).join(" ");
}
