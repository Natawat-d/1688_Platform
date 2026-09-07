package store

import (
	"encoding/json"
	"time"
)

// Row structs. Field order does not matter; the db tags carry the column names
// so pgx.RowToStructByName can populate them.

type Category struct {
	ID        int64  `db:"id"`
	ParentID  int64  `db:"parent_id"`
	Name      string `db:"name"`
	NameTrans string `db:"name_trans"`
	Level     int16  `db:"level"`
	IsLeaf    bool   `db:"is_leaf"`
	Visible   bool   `db:"visible"`
}

type Product struct {
	OfferID          int64     `db:"offer_id"`
	SellerOpenID     string    `db:"seller_open_id"`
	CompanyName      string    `db:"company_name"`
	Subject          string    `db:"subject"`
	SubjectTrans     string    `db:"subject_trans"`
	DescriptionTrans string    `db:"description_trans"`
	Keywords         string    `db:"keywords"`
	Images           []string  `db:"images"`
	WhiteImage       string    `db:"white_image"`
	CategoryID       int64     `db:"category_id"`
	TopCategoryID    int64     `db:"top_category_id"`
	SecondCategoryID int64     `db:"second_category_id"`
	ThirdCategoryID  int64     `db:"third_category_id"`
	CategoryName     string    `db:"category_name"`
	Status           string    `db:"status"`
	MinOrderQuantity int32     `db:"min_order_quantity"`
	BatchNumber      int32     `db:"batch_number"`
	QuoteType        int16     `db:"quote_type"`
	UnitTrans        string    `db:"unit_trans"`
	MixGeneral       bool      `db:"mix_general"`
	MixAmountFen     int64     `db:"mix_amount_fen"`
	MixNumber        int32     `db:"mix_number"`
	AmountOnSale     int32     `db:"amount_on_sale"`
	MonthSold        int32     `db:"month_sold"`
	TradeScore       string    `db:"trade_score"`
	Identities       []string  `db:"identities"`
	WeightG          int32     `db:"weight_g"`
	LengthMM         int32     `db:"length_mm"`
	WidthMM          int32     `db:"width_mm"`
	HeightMM         int32     `db:"height_mm"`
	ChinaFreightFen  int64     `db:"china_freight_fen"`
	FreightFree      bool      `db:"freight_free"`
	PriceMinFen      int64     `db:"price_min_fen"`
	PriceMaxFen      int64     `db:"price_max_fen"`
	SellMinSatang    int64     `db:"sell_min_satang"`
	Visible          bool      `db:"visible"`
	SyncedAt         time.Time `db:"synced_at"`
}

type SKU struct {
	SkuID         int64           `db:"sku_id"`
	OfferID       int64           `db:"offer_id"`
	SpecID        string          `db:"spec_id"`
	Attrs         json.RawMessage `db:"attrs"`
	Label         string          `db:"label"`
	PriceFen      int64           `db:"price_fen"`
	PromoPriceFen int64           `db:"promo_price_fen"`
	AmountOnSale  int32           `db:"amount_on_sale"`
	CargoNumber   string          `db:"cargo_number"`
	ImageURL      string          `db:"image_url"`
	WeightG       int32           `db:"weight_g"`
}

type Tier struct {
	OfferID       int64 `db:"offer_id"`
	StartQuantity int32 `db:"start_quantity"`
	PriceFen      int64 `db:"price_fen"`
	PromoPriceFen int64 `db:"promo_price_fen"`
}

// SKUAttr is one attribute of a SKU, stored as JSON in product_skus.attrs.
type SKUAttr struct {
	Name       string `json:"name"`
	NameTrans  string `json:"nameTrans"`
	Value      string `json:"value"`
	ValueTrans string `json:"valueTrans"`
	Image      string `json:"image,omitempty"`
}

type FeeRule struct {
	ID            int64      `db:"id"`
	Scope         string     `db:"scope"`
	ScopeValue    string     `db:"scope_value"`
	FeeBps        int32      `db:"fee_bps"`
	FeeFixedFen   int64      `db:"fee_fixed_fen"`
	MinFeeFen     int64      `db:"min_fee_fen"`
	Priority      int32      `db:"priority"`
	EffectiveFrom time.Time  `db:"effective_from"`
	EffectiveTo   *time.Time `db:"effective_to"`
	Note          string     `db:"note"`
	CreatedAt     time.Time  `db:"created_at"`
}

type CartItem struct {
	ID       int64     `db:"id"`
	CartID   string    `db:"cart_id"`
	OfferID  int64     `db:"offer_id"`
	SkuID    int64     `db:"sku_id"`
	SpecID   string    `db:"spec_id"`
	Quantity int32     `db:"quantity"`
	AddedAt  time.Time `db:"added_at"`
}

type Order struct {
	ID            int64           `db:"id"`
	PublicID      string          `db:"public_id"`
	AccessToken   string          `db:"access_token"`
	CartID        string          `db:"cart_id"`
	Email         string          `db:"email"`
	ShipName      string          `db:"ship_name"`
	ShipPhone     string          `db:"ship_phone"`
	ShipAddress   json.RawMessage `db:"ship_address"`
	Status        string          `db:"status"`
	FxPpm         int64           `db:"fx_ppm"`
	GoodsSatang   int64           `db:"goods_satang"`
	FreightSatang int64           `db:"freight_satang"`
	IntlSatang    int64           `db:"intl_satang"`
	FeeSatang     int64           `db:"fee_satang"`
	TotalSatang   int64           `db:"total_satang"`
	CreatedAt     time.Time       `db:"created_at"`
	PaidAt        *time.Time      `db:"paid_at"`
	UpdatedAt     time.Time       `db:"updated_at"`
}

type OrderItem struct {
	ID              int64           `db:"id"`
	OrderID         int64           `db:"order_id"`
	SupplierOrderID *int64          `db:"supplier_order_id"`
	OfferID         int64           `db:"offer_id"`
	SkuID           int64           `db:"sku_id"`
	SpecID          string          `db:"spec_id"`
	Title           string          `db:"title"`
	ImageURL        string          `db:"image_url"`
	SkuLabel        string          `db:"sku_label"`
	Quantity        int32           `db:"quantity"`
	BaseFen         int64           `db:"base_fen"`
	FreightFen      int64           `db:"freight_fen"`
	IntlFen         int64           `db:"intl_fen"`
	FeeFen          int64           `db:"fee_fen"`
	FeeRuleID       *int64          `db:"fee_rule_id"`
	FxPpm           int64           `db:"fx_ppm"`
	UnitSatang      int64           `db:"unit_satang"`
	LineSatang      int64           `db:"line_satang"`
	Snapshot        json.RawMessage `db:"snapshot"`
	SellerOpenID    string          `db:"seller_open_id"`
}

type SupplierOrder struct {
	ID            int64           `db:"id"`
	OrderID       int64           `db:"order_id"`
	SellerOpenID  string          `db:"seller_open_id"`
	SellerName    string          `db:"seller_name"`
	GroupSeq      int32           `db:"group_seq"`
	OutOrderID    string          `db:"out_order_id"`
	CbuOrderID    *int64          `db:"cbu_order_id"`
	Flow          string          `db:"flow"`
	TradeType     string          `db:"trade_type"`
	Status        string          `db:"status"`
	Status1688    string          `db:"status_1688"`
	StatusRank    int16           `db:"status_rank"`
	StatusAt      time.Time       `db:"status_at"`
	PayURL        string          `db:"pay_url"`
	SumPaymentFen int64           `db:"sum_payment_fen"`
	PostFeeFen    int64           `db:"post_fee_fen"`
	ErrorCode     string          `db:"error_code"`
	ErrorMessage  string          `db:"error_message"`
	Raw           json.RawMessage `db:"raw"`
	CreatedAt     time.Time       `db:"created_at"`
	UpdatedAt     time.Time       `db:"updated_at"`
}

type Shipment struct {
	ID              int64     `db:"id"`
	SupplierOrderID int64     `db:"supplier_order_id"`
	Leg             string    `db:"leg"`
	LogisticsID     string    `db:"logistics_id"`
	MailNo          string    `db:"mail_no"`
	CpCode          string    `db:"cp_code"`
	CompanyName     string    `db:"company_name"`
	Status          string    `db:"status"`
	UpdatedAt       time.Time `db:"updated_at"`
}

type TrackingEvent struct {
	ID         int64     `db:"id"`
	ShipmentID int64     `db:"shipment_id"`
	Source     string    `db:"source"`
	EventAt    time.Time `db:"event_at"`
	Code       string    `db:"code"`
	Remark     string    `db:"remark"`
	DedupeKey  string    `db:"dedupe_key"`
}

type Job struct {
	ID        int64           `db:"id"`
	Kind      string          `db:"kind"`
	Key       string          `db:"key"`
	Payload   json.RawMessage `db:"payload"`
	State     string          `db:"state"`
	Attempts  int32           `db:"attempts"`
	RunAt     time.Time       `db:"run_at"`
	LastError string          `db:"last_error"`
	CreatedAt time.Time       `db:"created_at"`
	UpdatedAt time.Time       `db:"updated_at"`
}

type MessageEvent struct {
	MsgID       int64           `db:"msg_id"`
	Type        string          `db:"type"`
	GmtBorn     time.Time       `db:"gmt_born"`
	Payload     json.RawMessage `db:"payload"`
	ReceivedAt  time.Time       `db:"received_at"`
	ProcessedAt *time.Time      `db:"processed_at"`
	Error       string          `db:"error"`
}

type APICall struct {
	ID        int64           `db:"id"`
	API       string          `db:"api"`
	MS        int32           `db:"ms"`
	OK        bool            `db:"ok"`
	Code      string          `db:"code"`
	Message   string          `db:"message"`
	Req       json.RawMessage `db:"req"`
	Resp      string          `db:"resp"`
	CreatedAt time.Time       `db:"created_at"`
}

// FullProduct is a product with everything the storefront needs in one value.
type FullProduct struct {
	Product Product
	SKUs    []SKU
	Tiers   []Tier
}
