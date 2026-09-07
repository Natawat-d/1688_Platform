import { useEffect, useState } from "react";
import { Link, NavLink, Outlet, useLocation, useNavigate, useSearchParams } from "react-router-dom";
import { useCart } from "./lib/cart";

/** Header search box. Seeded from ?q= so it stays in step with the URL. */
function HeaderSearch() {
  const navigate = useNavigate();
  const location = useLocation();
  const [params] = useSearchParams();
  const [q, setQ] = useState("");

  useEffect(() => {
    setQ(location.pathname === "/search" ? (params.get("q") ?? "") : "");
  }, [location.pathname, params]);

  return (
    <form
      className="headerSearch"
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
        name="q"
        placeholder="Search products"
        aria-label="Search products"
        value={q}
        onChange={(e) => setQ(e.target.value)}
      />
      <button className="btn" type="submit">
        Search
      </button>
    </form>
  );
}

export function App() {
  const { count } = useCart();
  const { pathname } = useLocation();

  // A fresh route should start at the top, the way a server-rendered page would.
  useEffect(() => {
    window.scrollTo(0, 0);
  }, [pathname]);

  return (
    <div className="app">
      <header className="header">
        <div className="headerInner">
          <Link className="brand" to="/">
            Marketplace
          </Link>
          <HeaderSearch />
          <nav className="headerNav">
            <NavLink className="navLink" to="/search">
              Browse
            </NavLink>
            <NavLink className="navLink" to="/admin">
              Admin
            </NavLink>
            <Link className="cartLink" to="/cart">
              Cart
              {count > 0 && (
                <span className="cartBadge" aria-label={`${count} items in cart`}>
                  {count}
                </span>
              )}
            </Link>
          </nav>
        </div>
      </header>

      <main className="main">
        <Outlet />
      </main>

      <footer className="footer">
        <div className="footerInner">
          <span>Goods sourced from 1688.com suppliers. Prices include shipping and platform fee.</span>
          <span>Prices in THB · settled in CNY</span>
        </div>
      </footer>
    </div>
  );
}
