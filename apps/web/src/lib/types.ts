/**
 * Wire types for our own HTTP API.
 *
 * Two rules govern every type in this file:
 *
 *  1. Money is always the server-rendered `Money` envelope. The UI renders `.text`.
 *     `minor` is a decimal string of integer minor units and exists only so a
 *     component can compare or sort; nothing in TypeScript ever does money maths.
 *  2. Every identifier is a `string`. 1688 ids are 18-19 digit int64s which
 *     `number` cannot hold, so they are never parsed, only passed through.
 */

export type Money = {
  minor: string;
  currency: string;
  text: string;
};

/** Integer minor units on the wire (fen). String or number, never arithmetic. */
export type Minor = string | number;

export type Issue = {
  code: string;
  field?: string;
  message: string;
};

// ---------------------------------------------------------------- catalogue

export type Category = {
  id: string;
  name: string;
  count: number;
};

export type Seller = {
  openId: string;
  name: string;
  score?: number | string;
};

export type ProductListItem = {
  offerId: string;
  title: string;
  image: string;
  priceFrom: Money;
  moq: number;
  unit: string;
  monthSold: number;
  seller: Seller;
};

export type ProductPage = {
  total: number;
  page: number;
  size: number;
  items: ProductListItem[];
};

export type SkuAttr = {
  name: string;
  value: string;
  image?: string;
};

export type Sku = {
  skuId: string;
  specId: string;
  label: string;
  attrs: SkuAttr[];
  stock: number;
  available: boolean;
  price: Money;
};

export type PriceTier = {
  startQuantity: number;
  price: Money;
};

export type ShippingSpec = {
  weightG: number;
  lengthMm: number;
  widthMm: number;
  heightMm: number;
};

/** sellerMixSetting: whether the offer joins general mixed-batch, and the thresholds. */
export type MixSetting = {
  general: boolean;
  amountFen: Minor;
  number: number;
};

/** A category path entry may arrive as a plain name or as {id,name}. */
export type CategoryPathNode = string | { id: string; name: string };

export type Product = {
  offerId: string;
  title: string;
  titleZh: string;
  description: string;
  images: string[];
  categoryPath: CategoryPathNode[];
  seller: Seller;
  status: string;
  sellable: boolean;
  moq: number;
  batchNumber: number;
  quoteType: number | string;
  unit: string;
  mix: MixSetting;
  stock: number;
  skus: Sku[];
  tiers: PriceTier[];
  shipping: ShippingSpec;
  priceRange: { min: Money; max: Money };
  syncedAt: string;
};

// --------------------------------------------------------------------- cart

/**
 * How one line's sell price was built. Every *Fen field is CNY minor units;
 * fxPpm is CNY->THB parts per million. Rendered by <PriceBreakdown>, which is
 * both the customer trust panel and our fastest pricing-bug debugger.
 */
export type Breakdown = {
  baseFen: Minor;
  freightFen: Minor;
  intlFen: Minor;
  feeFen: Minor;
  fxPpm: Minor;
  feeRuleId: string;
};

export type CartLine = {
  id: string;
  offerId: string;
  skuId: string;
  title: string;
  image: string;
  skuLabel: string;
  quantity: number;
  stock: number;
  unit: Money;
  line: Money;
  breakdown: Breakdown;
};

export type CartGroupData = {
  sellerOpenId: string;
  sellerName: string;
  parcels: number;
  lines: CartLine[];
  issues: Issue[];
  subtotal: Money;
};

export type CartTotals = {
  goods: Money;
  chinaFreight: Money;
  intl: Money;
  fee: Money;
  total: Money;
};

export type Cart = {
  id: string;
  groups: CartGroupData[];
  totals: CartTotals;
  checkoutable: boolean;
};

// ----------------------------------------------------------------- checkout

export type Address = {
  line1: string;
  city: string;
  province: string;
  postcode: string;
};

export type CheckoutRequest = {
  email: string;
  name: string;
  phone: string;
  address: Address;
};

export type CheckoutResult = {
  orderId: string;
  token: string;
  total: Money;
  parcels: number;
  payUrl: string;
};

// ------------------------------------------------------------------- orders

export type TimelineStep = {
  key: string;
  label: string;
  at: string;
  done: boolean;
  detail: string;
};

export type TraceEvent = {
  at: string;
  text: string;
  source: string;
};

export type OrderItem = {
  offerId: string;
  skuId: string;
  title: string;
  image: string;
  skuLabel: string;
  quantity: number;
  unit: Money;
  line: Money;
};

export type Parcel = {
  id: string;
  sellerName: string;
  status: string;
  cbuOrderId: string;
  carrier: string;
  trackingNo: string;
  items: OrderItem[];
  events: TraceEvent[];
};

export type Order = {
  orderId: string;
  status: string;
  statusLabel: string;
  placedAt: string;
  paidAt: string;
  email: string;
  totals: CartTotals;
  items: OrderItem[];
  timeline: TimelineStep[];
  parcels: Parcel[];
};

// -------------------------------------------------------------------- admin

export type FeeRuleScope = "global" | "category" | "supplier" | "product";

export type FeeRule = {
  id: string;
  name: string;
  scope: FeeRuleScope;
  /** category id / seller openId / offerId; empty for the global rule. */
  scopeValue: string;
  /** Percentage of the CNY subtotal in basis points. 250 = 2.50%. */
  feeBps: number;
  feeFixedFen: Minor;
  minFeeFen: Minor;
  priority: number;
  effectiveFrom: string;
  effectiveTo: string;
  active: boolean;
  /** Unknown server-side fields survive an edit round-trip. */
  [extra: string]: unknown;
};

export type QuotePreview = {
  offerId: string;
  skuId: string;
  quantity: number;
  unit: Money;
  total: Money;
  breakdown: Breakdown;
  feeRule?: { id: string; name: string; scope: string };
  issues?: Issue[];
};

/** Settings are rendered from whatever keys the server sends; no fixed schema. */
export type Settings = Record<string, unknown>;

export type ImportJobRequest = {
  keyword?: string;
  offerIds?: string[];
  all?: boolean;
};

export type Job = {
  id: string;
  kind: string;
  status: string;
  progress?: number;
  total?: number;
  done?: number;
  error?: string;
  startedAt?: string;
  finishedAt?: string;
  detail?: string;
  [extra: string]: unknown;
};

export type AdminOrder = {
  orderId: string;
  publicId?: string;
  status: string;
  email?: string;
  placedAt?: string;
  total?: Money;
  supplierOrders?: SupplierOrder[];
  [extra: string]: unknown;
};

export type SupplierOrder = {
  id: string;
  sellerOpenId?: string;
  sellerName?: string;
  status: string;
  cbuOrderId?: string;
  error?: string;
  [extra: string]: unknown;
};

export type ApiCall = {
  id?: string;
  api: string;
  url?: string;
  status?: number;
  durationMs?: number;
  at?: string;
  err?: string;
  [extra: string]: unknown;
};

export type PushMessage = {
  id?: string;
  msgId?: string;
  topic: string;
  status?: string;
  receivedAt?: string;
  processedAt?: string;
  payload?: unknown;
  [extra: string]: unknown;
};
