import type { ReactNode } from "react";
import type { ProductListItem } from "../lib/types";
import { ProductCard } from "./ProductCard";
import styles from "./ProductGrid.module.css";

export function ProductGrid({
  items,
  loading = false,
  skeletonCount = 12,
  empty = "No products matched.",
}: {
  items: ProductListItem[] | undefined;
  loading?: boolean;
  skeletonCount?: number;
  empty?: ReactNode;
}) {
  if (loading && !items) {
    return (
      <div className={styles.grid}>
        {Array.from({ length: skeletonCount }, (_, i) => (
          <div key={i} className={`skeleton ${styles.skeleton}`} />
        ))}
      </div>
    );
  }

  if (!items || items.length === 0) {
    return <div className="empty">{empty}</div>;
  }

  return (
    <div className={styles.grid} data-loading={loading || undefined}>
      {items.map((item) => (
        <ProductCard key={item.offerId} item={item} />
      ))}
    </div>
  );
}

/** Horizontal scroller used for the "popular" rail on the home page. */
export function ProductRail({ items, loading }: { items: ProductListItem[] | undefined; loading?: boolean }) {
  if (loading && !items) {
    return (
      <div className={styles.rail}>
        {Array.from({ length: 6 }, (_, i) => (
          <div key={i} className={`skeleton ${styles.railSkeleton}`} />
        ))}
      </div>
    );
  }
  if (!items || items.length === 0) return null;
  return (
    <div className={styles.rail}>
      {items.map((item) => (
        <div key={item.offerId} className={styles.railItem}>
          <ProductCard item={item} />
        </div>
      ))}
    </div>
  );
}
