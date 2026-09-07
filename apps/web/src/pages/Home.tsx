import { useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import { useQuery } from "../lib/useQuery";
import { ProductRail } from "../components/ProductGrid";
import type { Category, ProductPage } from "../lib/types";
import { num } from "../lib/fmt";
import styles from "./Home.module.css";

export function Home() {
  const navigate = useNavigate();
  const [q, setQ] = useState("");

  const categories = useQuery<Category[]>("/api/categories");
  const popular = useQuery<ProductPage>("/api/products?sort=popular&size=12&page=1");

  return (
    <div className="stack">
      <section className={styles.hero}>
        <h1 className={styles.heroTitle}>Wholesale goods from 1688, delivered to Thailand</h1>
        <p className={styles.heroText}>
          One price in baht with China freight, international shipping and our fee already included.
          We place the order with the supplier and follow every parcel for you.
        </p>
        <form
          className={styles.heroSearch}
          role="search"
          onSubmit={(e) => {
            e.preventDefault();
            const term = q.trim();
            navigate(term ? `/search?q=${encodeURIComponent(term)}` : "/search");
          }}
        >
          <input
            className="input"
            type="search"
            placeholder="Search for a product, e.g. phone case, tote bag"
            aria-label="Search products"
            value={q}
            onChange={(e) => setQ(e.target.value)}
          />
          <button className="btn btnPrimary btnLg" type="submit">
            Search
          </button>
        </form>
      </section>

      <section>
        <div className="sectionTitle">
          <h2>Categories</h2>
          <Link className="small" to="/search">
            Browse everything
          </Link>
        </div>
        {categories.error ? (
          <p className="notice noticeError">{categories.error}</p>
        ) : categories.loading && !categories.data ? (
          <div className={styles.tiles}>
            {Array.from({ length: 8 }, (_, i) => (
              <div key={i} className={`skeleton ${styles.tileSkeleton}`} />
            ))}
          </div>
        ) : (categories.data ?? []).length === 0 ? (
          <div className="empty">No categories yet. Import a catalogue from the admin console.</div>
        ) : (
          <ul className={styles.tiles}>
            {(categories.data ?? []).map((c) => (
              <li key={c.id}>
                <Link className={styles.tile} to={`/search?cat=${encodeURIComponent(c.id)}`}>
                  <span className={styles.tileName}>{c.name}</span>
                  <span className={styles.tileCount}>{num(c.count)} items</span>
                </Link>
              </li>
            ))}
          </ul>
        )}
      </section>

      <section>
        <div className="sectionTitle">
          <h2>Popular now</h2>
          <Link className="small" to="/search?sort=popular">
            See all
          </Link>
        </div>
        {popular.error ? (
          <p className="notice noticeError">{popular.error}</p>
        ) : (
          <ProductRail items={popular.data?.items} loading={popular.loading} />
        )}
      </section>
    </div>
  );
}
