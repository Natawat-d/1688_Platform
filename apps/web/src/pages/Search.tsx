import { useCallback, useMemo } from "react";
import { useSearchParams } from "react-router-dom";
import { useQuery } from "../lib/useQuery";
import { withQuery } from "../lib/api";
import { Filters } from "../components/Filters";
import type { FilterValues } from "../components/Filters";
import { ProductGrid } from "../components/ProductGrid";
import { Pager } from "../components/Pager";
import type { Category, ProductPage } from "../lib/types";
import { toInt, num } from "../lib/fmt";
import styles from "./Search.module.css";

const SORTS = [
  { value: "", label: "Relevance" },
  { value: "popular", label: "Most sold" },
  { value: "price_asc", label: "Price: low to high" },
  { value: "price_desc", label: "Price: high to low" },
  { value: "newest", label: "Newest" },
];

const PAGE_SIZE = 24;

export function Search() {
  const [params, setParams] = useSearchParams();

  // The URL is the single source of truth for every piece of search state.
  const q = params.get("q") ?? "";
  const cat = params.get("cat") ?? "";
  const min = params.get("min") ?? "";
  const max = params.get("max") ?? "";
  const moq = params.get("moq") ?? "";
  const instock = params.get("instock") === "1";
  const sort = params.get("sort") ?? "";
  const page = Math.max(1, toInt(params.get("page"), 1));
  const size = Math.max(1, toInt(params.get("size"), PAGE_SIZE));

  const url = useMemo(
    () =>
      withQuery("/api/products", {
        q,
        cat,
        min,
        max,
        moq,
        instock: instock ? "1" : "",
        sort,
        page,
        size,
      }),
    [q, cat, min, max, moq, instock, sort, page, size],
  );

  const results = useQuery<ProductPage>(url);
  const categories = useQuery<Category[]>("/api/categories");

  /** Writing a filter always resets to page 1; anything else is a confusing result set. */
  const patch = useCallback(
    (next: Record<string, string | number | boolean | undefined>, keepPage = false) => {
      setParams(
        (prev) => {
          const out = new URLSearchParams(prev);
          for (const [k, v] of Object.entries(next)) {
            const s = typeof v === "boolean" ? (v ? "1" : "") : v === undefined ? "" : String(v);
            if (s === "") out.delete(k);
            else out.set(k, s);
          }
          if (!keepPage) out.delete("page");
          return out;
        },
        { replace: false },
      );
    },
    [setParams],
  );

  const filterValues: FilterValues = { cat, min, max, moq, instock };

  const total = results.data?.total ?? 0;
  const heading = q ? `Results for “${q}”` : cat ? "Category" : "All products";

  return (
    <div className={styles.layout}>
      <Filters
        categories={categories.data}
        value={filterValues}
        loading={categories.loading}
        onChange={(p) => patch(p as Record<string, string | boolean>)}
        onClear={() => patch({ cat: "", min: "", max: "", moq: "", instock: false })}
      />

      <div className={styles.results}>
        <div className={styles.head}>
          <div>
            <h1 className={styles.title}>{heading}</h1>
            <p className="small muted">
              {results.loading && !results.data ? "Searching…" : `${num(total)} products`}
            </p>
          </div>
          <label className={styles.sort}>
            <span className="srOnly">Sort by</span>
            <select className="select" value={sort} onChange={(e) => patch({ sort: e.target.value })}>
              {SORTS.map((s) => (
                <option key={s.value} value={s.value}>
                  {s.label}
                </option>
              ))}
            </select>
          </label>
        </div>

        {results.error ? (
          <div className="notice noticeError">
            {results.error}{" "}
            <button className="btn btnGhost" onClick={results.reload}>
              Retry
            </button>
          </div>
        ) : (
          <>
            <ProductGrid
              items={results.data?.items}
              loading={results.loading}
              empty={
                q || cat || min || max || moq || instock
                  ? "No products matched those filters. Try widening the price range or clearing a filter."
                  : "The catalogue is empty. Import products from the admin console."
              }
            />
            <Pager
              page={results.data?.page ?? page}
              size={results.data?.size ?? size}
              total={total}
              onPage={(p) => patch({ page: p }, true)}
            />
          </>
        )}
      </div>
    </div>
  );
}
