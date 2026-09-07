import type { Category } from "../lib/types";
import { Field } from "./Field";
import { num } from "../lib/fmt";
import styles from "./Filters.module.css";

export type FilterValues = {
  cat: string;
  min: string;
  max: string;
  moq: string;
  instock: boolean;
};

/**
 * Every control writes straight back to the URL through `onChange`, so a
 * filtered result page is shareable and the back button walks the filter
 * history exactly as a user expects.
 */
export function Filters({
  categories,
  value,
  onChange,
  onClear,
  loading,
}: {
  categories: Category[] | undefined;
  value: FilterValues;
  onChange: (patch: Partial<FilterValues>) => void;
  onClear: () => void;
  loading?: boolean;
}) {
  const dirty = Boolean(value.cat || value.min || value.max || value.moq || value.instock);

  return (
    <aside className={styles.panel} aria-label="Filters">
      <div className={styles.head}>
        <h2 className={styles.title}>Filters</h2>
        {dirty && (
          <button type="button" className="btn btnGhost" onClick={onClear}>
            Clear
          </button>
        )}
      </div>

      <Field label="Category">
        {(id) => (
          <select
            id={id}
            className="select"
            value={value.cat}
            disabled={loading && !categories}
            onChange={(e) => onChange({ cat: e.target.value })}
          >
            <option value="">All categories</option>
            {(categories ?? []).map((c) => (
              <option key={c.id} value={c.id}>
                {c.name}
                {c.count > 0 ? ` (${num(c.count)})` : ""}
              </option>
            ))}
          </select>
        )}
      </Field>

      <fieldset className={styles.group}>
        <legend className={styles.legend}>Price (THB)</legend>
        <div className={styles.range}>
          <input
            className="input"
            type="text"
            inputMode="numeric"
            placeholder="Min"
            aria-label="Minimum price in THB"
            value={value.min}
            onChange={(e) => onChange({ min: e.target.value.replace(/[^0-9]/g, "") })}
          />
          <span className={styles.dash} aria-hidden="true">
            –
          </span>
          <input
            className="input"
            type="text"
            inputMode="numeric"
            placeholder="Max"
            aria-label="Maximum price in THB"
            value={value.max}
            onChange={(e) => onChange({ max: e.target.value.replace(/[^0-9]/g, "") })}
          />
        </div>
        <p className={styles.hint}>Whole baht, applied to the lowest price of each product.</p>
      </fieldset>

      <Field label="Maximum MOQ" hint="Hide products with a higher minimum order.">
        {(id) => (
          <input
            id={id}
            className="input"
            type="text"
            inputMode="numeric"
            placeholder="Any"
            value={value.moq}
            onChange={(e) => onChange({ moq: e.target.value.replace(/[^0-9]/g, "") })}
          />
        )}
      </Field>

      <label className="checkRow">
        <input
          type="checkbox"
          checked={value.instock}
          onChange={(e) => onChange({ instock: e.target.checked })}
        />
        In stock only
      </label>
    </aside>
  );
}
