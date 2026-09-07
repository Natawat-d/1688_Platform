import { Link } from "react-router-dom";
import type { ProductListItem } from "../lib/types";
import { MoneyFrom } from "./Money";
import { num } from "../lib/fmt";
import styles from "./ProductCard.module.css";

export function ProductCard({ item }: { item: ProductListItem }) {
  const unit = item.unit || "piece";
  return (
    <article className={styles.card}>
      {/* The whole card is one link; the ::after overlay makes the image and
          text share a single hit target without nesting interactive elements. */}
      <div className={styles.imageBox}>
        {item.image ? (
          <img className={styles.image} src={item.image} alt="" loading="lazy" decoding="async" />
        ) : (
          <div className={styles.imagePlaceholder} aria-hidden="true" />
        )}
      </div>
      <div className={styles.body}>
        <h3 className={styles.title}>
          <Link className={styles.link} to={`/p/${encodeURIComponent(item.offerId)}`}>
            {item.title || "Untitled"}
          </Link>
        </h3>
        <div className={styles.price}>
          <MoneyFrom value={item.priceFrom} size="lg" />
          <span className={styles.per}>/ {unit}</span>
        </div>
        <dl className={styles.meta}>
          <div>
            <dt>MOQ</dt>
            <dd>
              {num(item.moq)} {unit}
            </dd>
          </div>
          {item.monthSold > 0 && (
            <div>
              <dt>Sold</dt>
              <dd>{num(item.monthSold)}/mo</dd>
            </div>
          )}
        </dl>
        <p className={styles.seller} title={item.seller?.name}>
          {item.seller?.name || "Unknown supplier"}
        </p>
      </div>
    </article>
  );
}
