import { createContext, useCallback, useContext, useEffect, useMemo, useRef, useState } from "react";
import type { ReactNode } from "react";
import { api, errorText } from "./api";
import type { Cart } from "./types";

/**
 * The one piece of cross-page state that does not live in the URL.
 *
 * Every mutation is write-then-refetch: the server owns pricing, grouping,
 * MOQ validation and the issue list, so the client never predicts the result
 * of its own write.
 */

type CartContextValue = {
  cart: Cart | undefined;
  loading: boolean;
  busy: boolean;
  error: string | undefined;
  /** Total units across every line, for the header badge. */
  count: number;
  reload: () => Promise<void>;
  add: (offerId: string, skuId: string, quantity: number) => Promise<void>;
  setQty: (lineId: string, quantity: number) => Promise<void>;
  remove: (lineId: string) => Promise<void>;
};

const CartContext = createContext<CartContextValue | null>(null);

export function CartProvider({ children }: { children: ReactNode }) {
  const [cart, setCart] = useState<Cart | undefined>(undefined);
  const [loading, setLoading] = useState(true);
  const [busy, setBusy] = useState(false);
  const [error, setError] = useState<string | undefined>(undefined);
  const alive = useRef(true);

  useEffect(() => {
    alive.current = true;
    return () => {
      alive.current = false;
    };
  }, []);

  const reload = useCallback(async () => {
    try {
      const next = await api.get<Cart>("/api/cart");
      if (!alive.current) return;
      setCart(next);
      setError(undefined);
    } catch (err) {
      if (!alive.current) return;
      setError(errorText(err));
    } finally {
      if (alive.current) setLoading(false);
    }
  }, []);

  useEffect(() => {
    void reload();
  }, [reload]);

  // Every mutation funnels through here so `busy` and error handling are uniform.
  const mutate = useCallback(
    async (op: () => Promise<unknown>) => {
      setBusy(true);
      setError(undefined);
      try {
        await op();
        await reload();
      } catch (err) {
        if (alive.current) setError(errorText(err));
        throw err;
      } finally {
        if (alive.current) setBusy(false);
      }
    },
    [reload],
  );

  const add = useCallback(
    (offerId: string, skuId: string, quantity: number) =>
      mutate(() => api.post("/api/cart/items", { offerId, skuId, quantity })),
    [mutate],
  );

  const setQty = useCallback(
    (lineId: string, quantity: number) =>
      mutate(() => api.patch(`/api/cart/items/${encodeURIComponent(lineId)}`, { quantity })),
    [mutate],
  );

  const remove = useCallback(
    (lineId: string) => mutate(() => api.del(`/api/cart/items/${encodeURIComponent(lineId)}`)),
    [mutate],
  );

  const count = useMemo(() => {
    if (!cart) return 0;
    let n = 0;
    for (const g of cart.groups ?? []) for (const l of g.lines ?? []) n += l.quantity || 0;
    return n;
  }, [cart]);

  const value = useMemo<CartContextValue>(
    () => ({ cart, loading, busy, error, count, reload, add, setQty, remove }),
    [cart, loading, busy, error, count, reload, add, setQty, remove],
  );

  return <CartContext.Provider value={value}>{children}</CartContext.Provider>;
}

export function useCart(): CartContextValue {
  const ctx = useContext(CartContext);
  if (!ctx) throw new Error("useCart must be used inside <CartProvider>");
  return ctx;
}
