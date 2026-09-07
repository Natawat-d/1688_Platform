import { useEffect, useMemo, useState } from "react";
import { Link, useNavigate, useParams } from "react-router-dom";
import { useQuery } from "../lib/useQuery";
import { api, errorText, getAdminToken } from "../lib/api";
import { useCart } from "../lib/cart";
import { Gallery } from "../components/Gallery";
import { SkuPicker } from "../components/SkuPicker";
import { QtyStepper, snapQuantity } from "../components/QtyStepper";
import { Money } from "../components/Money";
import { PriceBreakdown } from "../components/PriceBreakdown";
import { Readout } from "../components/Field";
import { StatusBadge } from "../components/StatusBadge";
import type { CategoryPathNode, PriceTier, Product as ProductType, QuotePreview, Sku } from "../lib/types";
import { dateTimeText, dimensionsText, fenText, num, weightText } from "../lib/fmt";
import styles from "./Product.module.css";

/** The tier in force at `quantity`: the highest startQuantity not above it. */
export function activeTier(tiers: PriceTier[] | undefined, quantity: number): PriceTier | undefined {
  if (!tiers || tiers.length === 0) return undefined;
  const sorted = [...tiers].sort((a, b) => a.startQuantity - b.startQuantity);
  let found: PriceTier | undefined;
  for (const t of sorted) {
    if (quantity >= t.startQuantity) found = t;
    else break;
  }
  return found ?? sorted[0];
}

/** The next tier up, so we can nudge "add N more to reach ¥x". */
function nextTier(tiers: PriceTier[] | undefined, quantity: number): PriceTier | undefined {
  if (!tiers || tiers.length === 0) return undefined;
  return [...tiers].sort((a, b) => a.startQuantity - b.startQuantity).find((t) => t.startQuantity > quantity);
}

function pathName(node: CategoryPathNode): string {
  return typeof node === "string" ? node : node.name;
}
function pathId(node: CategoryPathNode): string | undefined {
  return typeof node === "string" ? undefined : node.id;
}

export function Product() {
  const { offerId = "" } = useParams();
  const navigate = useNavigate();
  const cart = useCart();

  const { data: product, error, loading, reload } = useQuery<ProductType>(
    offerId ? `/api/products/${encodeURIComponent(offerId)}` : null,
  );

  const [skuId, setSkuId] = useState<string | null>(null);
  const [quantity, setQuantity] = useState(1);
  const [showZh, setShowZh] = useState(false);
  const [adding, setAdding] = useState(false);
  const [addError, setAddError] = useState<string | undefined>();
  const [added, setAdded] = useState(false);

  const skus = product?.skus ?? [];
  const moq = Math.max(1, product?.moq ?? 1);
  const batch = Math.max(1, product?.batchNumber ?? 1);

  // First in-stock SKU, and a quantity that already satisfies MOQ and batch.
  useEffect(() => {
    if (!product) return;
    const first = skus.find((s) => s.available !== false && (s.stock ?? 0) > 0) ?? skus[0];
    setSkuId(first ? first.skuId : null);
    setQuantity(snapQuantity(moq, moq, batch));
    setAdded(false);
    setAddError(undefined);
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [product?.offerId]);

  const selectedSku: Sku | undefined = useMemo(
    () => skus.find((s) => s.skuId === skuId),
    [skus, skuId],
  );

  const tier = activeTier(product?.tiers, quantity);
  const upcoming = nextTier(product?.tiers, quantity);

  // Price shown to the shopper: the quantity tier wins when the offer has
  // tiers, otherwise the selected variant's own price. The cart is always the
  // authority for the line total; nothing here multiplies money.
  const unitPrice = tier?.price ?? selectedSku?.price ?? product?.priceRange?.min;

  const stock = selectedSku ? selectedSku.stock : (product?.stock ?? 0);
  const sellable = product?.sellable !== false;
  const needsSku = skus.length > 0;
  const skuChosen = !needsSku || Boolean(selectedSku);
  const outOfStock = stock <= 0;
  const canAdd = Boolean(product) && sellable && skuChosen && !outOfStock && !adding && !cart.busy;

  async function addToCart() {
    if (!product) return;
    setAdding(true);
    setAddError(undefined);
    try {
      await cart.add(product.offerId, selectedSku?.skuId ?? "", quantity);
      setAdded(true);
    } catch (err) {
      setAddError(errorText(err));
    } finally {
      setAdding(false);
    }
  }

  if (error) {
    return (
      <div className="notice noticeError">
        {error}{" "}
        <button className="btn btnGhost" onClick={reload}>
          Retry
        </button>
      </div>
    );
  }

  if (loading && !product) {
    return (
      <div className={styles.layout}>
        <div className={`skeleton ${styles.gallerySkeleton}`} />
        <div className="stack">
          <div className="skeleton" style={{ height: 28, width: "80%" }} />
          <div className="skeleton" style={{ height: 44, width: "40%" }} />
          <div className="skeleton" style={{ height: 120 }} />
          <div className="skeleton" style={{ height: 90 }} />
        </div>
      </div>
    );
  }

  if (!product) return <div className="empty">Product not found.</div>;

  const unit = product.unit || "piece";

  return (
    <div className="stack">
      {product.categoryPath?.length > 0 && (
        <nav className={styles.crumbs} aria-label="Breadcrumb">
          <Link to="/search">All</Link>
          {product.categoryPath.map((node, i) => {
            const id = pathId(node);
            return (
              <span key={`${pathName(node)}-${i}`}>
                <span className={styles.crumbSep} aria-hidden="true">
                  ›
                </span>
                {id ? <Link to={`/search?cat=${encodeURIComponent(id)}`}>{pathName(node)}</Link> : <span>{pathName(node)}</span>}
              </span>
            );
          })}
        </nav>
      )}

      <div className={styles.layout}>
        <div className={styles.galleryCol}>
          <Gallery
            images={product.images ?? []}
            alt={product.title}
            pinned={selectedSku?.attrs?.find((a) => a.image)?.image}
          />
        </div>

        <div className={styles.buyCol}>
          <div className={styles.titleBlock}>
            <h1 className={styles.title}>{showZh ? product.titleZh || product.title : product.title}</h1>
            {product.titleZh && product.titleZh !== product.title && (
              <button type="button" className="btn btnGhost" onClick={() => setShowZh((v) => !v)}>
                {showZh ? "Show English" : "Show original 中文"}
              </button>
            )}
          </div>

          <div className={styles.priceBlock}>
            <Money value={unitPrice} size="xl" />
            <span className={styles.per}>per {unit}</span>
            {tier && product.tiers.length > 1 && (
              <span className="badge badgeInfo">{num(tier.startQuantity)}+ tier</span>
            )}
          </div>
          {product.priceRange && product.priceRange.min?.text !== product.priceRange.max?.text && (
            <p className="xs faint">
              Range across variants and tiers: <Money value={product.priceRange.min} size="sm" muted /> –{" "}
              <Money value={product.priceRange.max} size="sm" muted />
            </p>
          )}
          {upcoming && (
            <p className={styles.tierNudge}>
              Order {num(upcoming.startQuantity)}+ for <Money value={upcoming.price} size="sm" /> per {unit}
            </p>
          )}

          {!sellable && (
            <p className="notice noticeWarn">
              This product is not currently sellable{product.status ? ` (${product.status})` : ""}.
            </p>
          )}

          <SkuPicker skus={skus} selectedSkuId={skuId} onSelect={setSkuId} />

          <div className={styles.qtyBlock}>
            <span className={styles.qtyLabel}>Quantity</span>
            <QtyStepper
              value={quantity}
              onChange={setQuantity}
              moq={moq}
              batch={batch}
              max={stock > 0 ? stock : undefined}
              unit={unit}
            />
          </div>

          <div className={styles.actions}>
            <button className="btn btnPrimary btnLg" onClick={addToCart} disabled={!canAdd}>
              {adding ? "Adding…" : outOfStock ? "Out of stock" : "Add to cart"}
            </button>
            {added && (
              <button className="btn btnLg" onClick={() => navigate("/cart")}>
                Go to cart
              </button>
            )}
          </div>
          {addError && <p className="notice noticeError">{addError}</p>}
          {added && !addError && <p className="notice noticeOk">Added. The cart shows the final price with shipping and fee.</p>}

          <PricePanel offerId={product.offerId} skuId={selectedSku?.skuId ?? ""} quantity={quantity} />
        </div>
      </div>

      <div className={styles.detailGrid}>
        {product.tiers?.length > 0 && (
          <section className="card">
            <h2 className={styles.panelTitle}>Quantity pricing</h2>
            <table className={styles.tierTable}>
              <thead>
                <tr>
                  <th scope="col">Quantity</th>
                  <th scope="col">Price per {unit}</th>
                </tr>
              </thead>
              <tbody>
                {[...product.tiers]
                  .sort((a, b) => a.startQuantity - b.startQuantity)
                  .map((t) => (
                    <tr key={t.startQuantity} className={tier?.startQuantity === t.startQuantity ? styles.tierActive : undefined}>
                      <td>{num(t.startQuantity)}+</td>
                      <td>
                        <Money value={t.price} size="sm" />
                      </td>
                    </tr>
                  ))}
              </tbody>
            </table>
          </section>
        )}

        <section className="card">
          <h2 className={styles.panelTitle}>Order rules</h2>
          <Readout label="Minimum order">
            {num(moq)} {unit}
          </Readout>
          {batch > 1 && <Readout label="Order multiples of">{num(batch)}</Readout>}
          <Readout label="Stock">{stock > 0 ? `${num(stock)} ${unit}` : "None"}</Readout>
          <Readout label="Status">
            <StatusBadge status={product.status} />
          </Readout>
          <Readout label="Mixed batch">
            {product.mix?.general
              ? `Allowed · min ${fenText(product.mix.amountFen)} or ${num(product.mix?.number ?? 0)} pieces`
              : "Not offered"}
          </Readout>
          {product.quoteType !== undefined && product.quoteType !== "" && (
            <Readout label="Quote type">{String(product.quoteType)}</Readout>
          )}
        </section>

        <section className="card">
          <h2 className={styles.panelTitle}>Shipping estimate</h2>
          <Readout label="Unit weight">{weightText(product.shipping?.weightG)}</Readout>
          <Readout label="Dimensions">
            {dimensionsText(product.shipping?.lengthMm, product.shipping?.widthMm, product.shipping?.heightMm)}
          </Readout>
          <p className={styles.panelNote}>
            China freight and the international leg are already in the price you see. Weight and
            volume come from the supplier's listing and decide the international rate.
          </p>
        </section>

        <section className="card">
          <h2 className={styles.panelTitle}>Supplier</h2>
          <Readout label="Name">{product.seller?.name || "Unknown"}</Readout>
          {product.seller?.score !== undefined && product.seller.score !== "" && (
            <Readout label="Trade score">{String(product.seller.score)}</Readout>
          )}
          <Readout label="Supplier ID" mono>
            {product.seller?.openId || "—"}
          </Readout>
          <Readout label="Offer ID" mono>
            {product.offerId}
          </Readout>
          <Readout label="Catalogue synced">{dateTimeText(product.syncedAt)}</Readout>
        </section>
      </div>

      {product.description && (
        <section className="card">
          <h2 className={styles.panelTitle}>Description</h2>
          <p className={styles.description}>{product.description}</p>
        </section>
      )}
    </div>
  );
}

/**
 * "How this price is made".
 *
 * The storefront API has no per-line quote endpoint, so the real numbers only
 * exist once a line is in the cart. When an admin token is present we pull the
 * genuine figures from the fee-rule preview, which turns this panel into the
 * pricing debugger it is meant to be; otherwise we explain the formula.
 */
function PricePanel({ offerId, skuId, quantity }: { offerId: string; skuId: string; quantity: number }) {
  const [quote, setQuote] = useState<QuotePreview | null>(null);
  const isAdmin = Boolean(getAdminToken());

  useEffect(() => {
    if (!isAdmin || !offerId) {
      setQuote(null);
      return;
    }
    let live = true;
    api
      .post<QuotePreview>("/api/admin/fee-rules/preview", { offerId, skuId, quantity })
      .then((q) => {
        if (live) setQuote(q);
      })
      .catch(() => {
        if (live) setQuote(null);
      });
    return () => {
      live = false;
    };
  }, [isAdmin, offerId, skuId, quantity]);

  if (quote?.breakdown) {
    return (
      <PriceBreakdown
        breakdown={quote.breakdown}
        quantity={quantity}
        unit={quote.unit}
        total={quote.total}
        title="How this price is made (live preview)"
      />
    );
  }

  return (
    <details className={styles.formula}>
      <summary className={styles.formulaSummary}>How this price is made</summary>
      <div className={styles.formulaBody}>
        <p>Every price on this site is one number with nothing added at checkout:</p>
        <ol className={styles.formulaSteps}>
          <li>
            <strong>Supplier price</strong> — the 1688 price for your quantity tier.
          </li>
          <li>
            <strong>China freight</strong> — supplier to our consolidation point.
          </li>
          <li>
            <strong>International shipping</strong> — our rate card, from the listed weight and volume.
          </li>
          <li>
            <strong>Platform fee</strong> — our margin, from the fee rule that applies to this product.
          </li>
          <li>
            <strong>FX</strong> — one conversion from CNY to THB at the rate recorded with your order.
          </li>
        </ol>
        <p className="xs faint">
          Add the item to your cart to see the exact figures for each step at your quantity.
        </p>
      </div>
    </details>
  );
}
