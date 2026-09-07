import { Navigate, Route, Routes } from "react-router-dom";
import { App } from "./App";
import { Home } from "./pages/Home";
import { Search } from "./pages/Search";
import { Product } from "./pages/Product";
import { Cart } from "./pages/Cart";
import { Checkout } from "./pages/Checkout";
import { OrderStatus } from "./pages/OrderStatus";
import { NotFound } from "./pages/NotFound";
import { Admin } from "./admin/Admin";
import { Import } from "./admin/Import";
import { FeeRules } from "./admin/FeeRules";
import { Settings } from "./admin/Settings";
import { Orders } from "./admin/Orders";
import { Jobs } from "./admin/Jobs";
import { StubControls } from "./admin/StubControls";

export function AppRoutes() {
  return (
    <Routes>
      <Route element={<App />}>
        <Route path="/" element={<Home />} />
        <Route path="/search" element={<Search />} />
        <Route path="/p/:offerId" element={<Product />} />
        <Route path="/cart" element={<Cart />} />
        <Route path="/checkout" element={<Checkout />} />
        <Route path="/orders/:publicId" element={<OrderStatus />} />

        <Route path="/admin" element={<Admin />}>
          <Route index element={<Navigate to="import" replace />} />
          <Route path="import" element={<Import />} />
          <Route path="fee-rules" element={<FeeRules />} />
          <Route path="settings" element={<Settings />} />
          <Route path="orders" element={<Orders />} />
          <Route path="jobs" element={<Jobs />} />
          <Route path="stub" element={<StubControls />} />
        </Route>

        <Route path="*" element={<NotFound />} />
      </Route>
    </Routes>
  );
}
