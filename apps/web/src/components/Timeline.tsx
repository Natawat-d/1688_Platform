import type { TimelineStep } from "../lib/types";
import { dateTimeText } from "../lib/fmt";
import styles from "./Timeline.module.css";

/**
 * Vertical order timeline. The first not-done step after a done one is the
 * "current" step and gets the emphasis; everything after it is pending.
 */
export function Timeline({ steps }: { steps: TimelineStep[] | undefined }) {
  if (!steps || steps.length === 0) return null;

  const currentIndex = steps.findIndex((s) => !s.done);

  return (
    <ol className={styles.timeline}>
      {steps.map((step, i) => {
        const state = step.done ? "done" : i === currentIndex ? "current" : "pending";
        return (
          <li key={step.key || `${i}`} className={styles.step} data-state={state}>
            <span className={styles.marker} aria-hidden="true" />
            <div className={styles.content}>
              <div className={styles.headline}>
                <span className={styles.label}>{step.label}</span>
                {step.at && <time className={styles.at}>{dateTimeText(step.at)}</time>}
              </div>
              {step.detail && <p className={styles.detail}>{step.detail}</p>}
            </div>
          </li>
        );
      })}
    </ol>
  );
}
