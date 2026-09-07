import { useCallback, useEffect, useState } from "react";
import { apiFetch, errorText } from "./api";

export type QueryState<T> = {
  data: T | undefined;
  error: string | undefined;
  loading: boolean;
  reload: () => void;
};

/**
 * The only data-fetching mechanism in the app. Refetches when `url` or any
 * dependency changes, aborts the in-flight request on change or unmount, and
 * keeps the previous data visible while reloading so lists do not flash.
 *
 * Pass `url = null` to hold off (an unauthenticated admin page, a route param
 * that has not resolved yet).
 */
export function useQuery<T>(url: string | null, deps: unknown[] = []): QueryState<T> {
  const [data, setData] = useState<T | undefined>(undefined);
  const [error, setError] = useState<string | undefined>(undefined);
  const [loading, setLoading] = useState<boolean>(url !== null);
  const [nonce, setNonce] = useState(0);

  useEffect(() => {
    if (url === null) {
      setLoading(false);
      return;
    }
    const ac = new AbortController();
    setLoading(true);
    setError(undefined);
    apiFetch<T>(url, { signal: ac.signal })
      .then((d) => {
        if (ac.signal.aborted) return;
        setData(d);
        setLoading(false);
      })
      .catch((e) => {
        if (ac.signal.aborted || (e as { name?: string })?.name === "AbortError") return;
        setError(errorText(e));
        setLoading(false);
      });
    return () => ac.abort();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [url, nonce, ...deps]);

  const reload = useCallback(() => setNonce((n) => n + 1), []);
  return { data, error, loading, reload };
}
