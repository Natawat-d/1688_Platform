import { useEffect, useMemo, useState } from "react";
import styles from "./Gallery.module.css";

/**
 * Main image plus thumbnails. `pinned` lets the SKU picker push the image of
 * the currently selected variant to the front without losing the user's own
 * thumbnail choice for the rest of the gallery.
 */
export function Gallery({ images, alt, pinned }: { images: string[]; alt: string; pinned?: string }) {
  const list = useMemo(() => (Array.isArray(images) ? images.filter(Boolean) : []), [images]);
  const [active, setActive] = useState(0);

  useEffect(() => {
    if (!pinned) return;
    const i = list.indexOf(pinned);
    if (i >= 0) setActive(i);
  }, [pinned, list]);

  // A shorter image list (a different product) must not leave `active` past the end.
  useEffect(() => {
    setActive((a) => (a < list.length ? a : 0));
  }, [list.length]);

  const src = pinned && !list.includes(pinned) ? pinned : list[active];

  return (
    <div className={styles.gallery}>
      <div className={styles.main}>
        {src ? (
          <img className={styles.mainImage} src={src} alt={alt} />
        ) : (
          <div className={styles.placeholder} aria-hidden="true" />
        )}
      </div>
      {list.length > 1 && (
        <ul className={styles.thumbs}>
          {list.map((img, i) => (
            <li key={`${img}-${i}`}>
              <button
                type="button"
                className={i === active && !(pinned && !list.includes(pinned)) ? `${styles.thumb} ${styles.thumbActive}` : styles.thumb}
                onClick={() => setActive(i)}
                aria-label={`Image ${i + 1} of ${list.length}`}
                aria-pressed={i === active}
              >
                <img src={img} alt="" loading="lazy" decoding="async" />
              </button>
            </li>
          ))}
        </ul>
      )}
    </div>
  );
}
