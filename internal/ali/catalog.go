package ali

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"
	"strings"
)

// The catalogue endpoints: multilingual keyword search, search navigation,
// product detail, category translation, and the domestic freight estimate.
// All five answer in envelope family A, so DecodeA does the unwrapping.
//
// Every catalogue price is a decimal STRING ("18.50", sometimes ""), which is
// why they are PriceString and not Fen: keeping the text lets the contract test
// compare against the documentation, and .Fen() converts exactly when the
// pricing engine needs a number.

// Decimal is a non-money numeric that 1688 sends as a JSON number or a quoted
// decimal: quantities, weights, package dimensions. The text is kept verbatim so
// nothing is rounded on the way in, and Int64/Milli convert without ever going
// through a float.
type Decimal string

func (d Decimal) String() string { return string(d) }

// Empty reports whether the field carried no value.
func (d Decimal) Empty() bool { return strings.TrimSpace(string(d)) == "" }

// Int64 is the value rounded to a whole number, half away from zero.
func (d Decimal) Int64() int64 { v, _ := scale10(string(d), 0); return v }

// Milli is the value times one thousand, rounded half away from zero: grams from
// a weight in kg, millimetres from a length in cm.
func (d Decimal) Milli() int64 { v, _ := scale10(string(d), 3); return v }

func (d *Decimal) UnmarshalJSON(b []byte) error {
	s := strings.TrimSpace(string(b))
	if s == "null" {
		*d = ""
		return nil
	}
	if len(s) > 0 && s[0] == '"' {
		var str string
		if err := json.Unmarshal(b, &str); err != nil {
			return err
		}
		*d = Decimal(str)
		return nil
	}
	*d = Decimal(s)
	return nil
}

// MarshalJSON emits a bare number, which is what the Double-typed request fields
// expect. An empty Decimal becomes 0 rather than invalid JSON.
func (d Decimal) MarshalJSON() ([]byte, error) {
	if d.Empty() {
		return []byte("0"), nil
	}
	return []byte(d), nil
}

// scale10 converts a decimal string to an integer scaled by 10^digits, rounding
// half away from zero. Text arithmetic only, so a weight of 0.001 kg survives.
func scale10(s string, digits int) (int64, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0, nil
	}
	neg := false
	switch s[0] {
	case '-':
		neg, s = true, s[1:]
	case '+':
		s = s[1:]
	}
	whole, frac, _ := strings.Cut(s, ".")
	if whole == "" {
		whole = "0"
	}
	if !allDigits(whole) || (frac != "" && !allDigits(frac)) {
		return 0, ErrBadAmount
	}
	for len(frac) < digits+1 {
		frac += "0"
	}
	v, err := strconv.ParseInt(whole+frac[:digits], 10, 64)
	if err != nil {
		return 0, ErrBadAmount
	}
	if frac[digits] >= '5' {
		v++
	}
	if neg {
		v = -v
	}
	return v, nil
}

// StringID is an ID that travels as a JSON string. 1688 types skuId as String on
// the freight estimate and in skuShippingDetails while typing the same value as
// Long everywhere else, so both spellings need a home.
type StringID ID

func (s StringID) ID() ID { return ID(s) }

func (s StringID) String() string { return ID(s).String() }

func (s StringID) MarshalJSON() ([]byte, error) { return []byte(`"` + ID(s).String() + `"`), nil }

func (s *StringID) UnmarshalJSON(b []byte) error {
	var i ID
	if err := i.UnmarshalJSON(b); err != nil {
		return err
	}
	*s = StringID(i)
	return nil
}

// checkResultStatus catches the one family A failure DecodeA cannot see.
//
// DecodeA recognises the wrapping object by it having BOTH a status field and a
// "result" of its own. A failed call has no inner result to return, so the
// wrapper stops looking like a wrapper and its {success:false,code,message} is
// decoded as though it were the payload — a refusal would arrive as an empty
// page. Checking the inner status first closes that hole; when the inner result
// is present DecodeA reaches the same verdict and this is a no-op.
func checkResultStatus(api string, body []byte) error {
	var env struct {
		Result statusFields `json:"result"`
	}
	if err := json.Unmarshal(body, &env); err != nil {
		return nil // let DecodeA report the malformed document
	}
	if env.Result.ok() {
		return nil
	}
	return &APIError{
		API:     api,
		Code:    firstNonEmpty(env.Result.code(), "FAILED"),
		Message: env.Result.message(),
		Body:    truncate(body),
	}
}

// SortBy renders the sort parameter of keywordQuery, which is a JSON string
// inside the query object: SortBy("price", "asc") is {"price":"asc"}.
func SortBy(field, direction string) string {
	b, err := json.Marshal(map[string]string{field: direction})
	if err != nil {
		return ""
	}
	return string(b)
}

// ---------------------------------------------------------------- keywordQuery

// OfferQuery is product.search.keywordQuery.param.OfferQueryParam, sent as the
// single object parameter offerQueryParam.
type OfferQuery struct {
	Keyword             string       `json:"keyword"`
	BeginPage           int          `json:"beginPage"`
	PageSize            int          `json:"pageSize"`
	Country             string       `json:"country"` // required: language code, "en"
	Filter              string       `json:"filter,omitempty"`
	Sort                string       `json:"sort,omitempty"` // JSON string, see SortBy
	OutMemberID         string       `json:"outMemberId,omitempty"`
	PriceStart          string       `json:"priceStart,omitempty"`
	PriceEnd            string       `json:"priceEnd,omitempty"`
	CategoryID          ID           `json:"categoryId,omitempty"`
	CategoryIDList      string       `json:"categoryIdList,omitempty"` // "2,45"
	RegionOpp           string       `json:"regionOpp,omitempty"`
	ProductCollectionID string       `json:"productCollectionId,omitempty"`
	SnID                string       `json:"snId,omitempty"` // "978" or "978:1352"
	KeywordTranslate    bool         `json:"keywordTranslate,omitempty"`
	SaleFilterList      []SaleFilter `json:"saleFilterList,omitempty"`
}

// SaleFilter is com.alibaba.cbu.offer.param.SaleFilterParam. saleType is one of
// sales7, sales14, sales30, totalSales; several entries intersect.
type SaleFilter struct {
	SaleType  string `json:"saleType,omitempty"`
	SaleStart string `json:"saleStart,omitempty"`
	SaleEnd   string `json:"saleEnd,omitempty"`
}

// OfferPage is product.search.keywordQuery.model.PageInfoV3.
type OfferPage struct {
	TotalRecords int            `json:"totalRecords"`
	TotalPage    int            `json:"totalPage"`
	PageSize     int            `json:"pageSize"`
	CurrentPage  int            `json:"currentPage"`
	Data         []OfferSummary `json:"data"`
}

// OfferSummary is product.search.keywordQuery.model.ProductInfoModelV2, one hit
// in the search result. Note promotionURL: the search API capitalises the URL
// while the detail API spells the same idea promotionUrl.
type OfferSummary struct {
	ImageURL                  string           `json:"imageUrl"`
	AigcImageURL              string           `json:"aigcImageUrl"`
	Subject                   string           `json:"subject"`
	SubjectTrans              string           `json:"subjectTrans"`
	OfferID                   ID               `json:"offerId"`
	IsJxhy                    FlexBool         `json:"isJxhy"`
	PriceInfo                 PriceInfo        `json:"priceInfo"`
	RepurchaseRate            string           `json:"repurchaseRate"`
	MonthSold                 int64            `json:"monthSold"`
	TraceInfo                 string           `json:"traceInfo"`
	IsOnePsale                FlexBool         `json:"isOnePsale"`
	SellerIdentities          []string         `json:"sellerIdentities"`
	OfferIdentities           []string         `json:"offerIdentities"`
	TradeScore                string           `json:"tradeScore"`
	WhiteImage                string           `json:"whiteImage"`
	PromotionModel            PromotionModel   `json:"promotionModel"`
	TopCategoryID             ID               `json:"topCategoryId"`
	SecondCategoryID          ID               `json:"secondCategoryId"`
	ThirdCategoryID           ID               `json:"thirdCategoryId"`
	IsPatentProduct           FlexBool         `json:"isPatentProduct"`
	CreateDate                Timestamp        `json:"createDate"`
	ModifyDate                Timestamp        `json:"modifyDate"`
	IsSelect                  FlexBool         `json:"isSelect"`
	MinOrderQuantity          int64            `json:"minOrderQuantity"`
	SellerDataInfo            SellerDataInfo   `json:"sellerDataInfo"`
	ProductSimpleShippingInfo SimpleShipping   `json:"productSimpleShippingInfo"`
	Token                     string           `json:"token"`
	PromotionURL              string           `json:"promotionURL"`
	Sales7d                   string           `json:"sales7d"`
	ProductTradeInfo          ProductTradeInfo `json:"productTradeInfo"`
	InvoiceInfo               InvoiceInfo      `json:"invoiceInfo"`
}

// PriceInfo is product.search.keywordQuery.model.PriceInfoV2. Yuan decimals.
type PriceInfo struct {
	Price          PriceString `json:"price"`
	JxhyPrice      PriceString `json:"jxhyPrice"`
	PfJxhyPrice    PriceString `json:"pfJxhyPrice"`
	ConsignPrice   PriceString `json:"consignPrice"`
	PromotionPrice PriceString `json:"promotionPrice"`
}

// PromotionModel is com.alibaba.cbu.offer.model.PromotionModel.
type PromotionModel struct {
	HasPromotion  FlexBool `json:"hasPromotion"`
	PromotionType string   `json:"promotionType"`
}

// SellerDataInfo is the union of the search and detail spellings: the search
// model carries tpYear, the detail model carries the two 30-day rates.
type SellerDataInfo struct {
	TradeMedalLevel              string `json:"tradeMedalLevel"`
	CompositeServiceScore        string `json:"compositeServiceScore"`
	LogisticsExperienceScore     string `json:"logisticsExperienceScore"`
	DisputeComplaintScore        string `json:"disputeComplaintScore"`
	OfferExperienceScore         string `json:"offerExperienceScore"`
	AfterSalesExperienceScore    string `json:"afterSalesExperienceScore"`
	ConsultingExperienceScore    string `json:"consultingExperienceScore"`
	RepeatPurchasePercent        string `json:"repeatPurchasePercent"`
	TpYear                       int    `json:"tpYear"`
	Collect30DayWithin48HPercent string `json:"collect30DayWithin48HPercent"`
	QualityRefundWithin30Day     string `json:"qualityRefundWithin30Day"`
}

// SimpleShipping is product.search.keywordQuery.model.ProductSimpleShippingInfo.
type SimpleShipping struct {
	ShippingTimeGuarantee     string `json:"shippingTimeGuarantee"`
	PerfectFulfillmentRate30d string `json:"perfectFulfillmentRate30d"`
	PickupWithin24hRate30d    string `json:"pickupWithin24hRate30d"`
	QualityReturnRate30d      string `json:"qualityReturnRate30d"`
	PerfectFulfillmentRate7d  string `json:"perfectFulfillmentRate7d"`
	PickupWithin24hRate7d     string `json:"pickupWithin24hRate7d"`
	DelayedShippingRate7d     string `json:"delayedShippingRate7d"`
}

// ProductTradeInfo is com.alibaba.cbu.offer.model.ProductTradeInfo.
type ProductTradeInfo struct {
	AddCartCount7d   string `json:"addCartCount7d"`
	PayBuyerCount7d  string `json:"payBuyerCount7d"`
	AddCartCount30d  string `json:"addCartCount30d"`
	PayBuyerCount30d string `json:"payBuyerCount30d"`
}

// InvoiceInfo is product.search.queryProductDetail.model.InvoiceInfo, returned
// by both the search and the detail API.
type InvoiceInfo struct {
	SupportOnlineInvoice FlexBool `json:"supportOnlineInvoice"`
	SupportFastInvoice   FlexBool `json:"supportFastInvoice"`
	InvoiceTypes         []string `json:"invoiceTypes"`
	TaxpayerType         string   `json:"taxpayerType"`
}

// SearchOffers runs the multilingual keyword search. Country is the language
// code and is mandatory; page defaults to the documented first page.
func (c *Client) SearchOffers(ctx context.Context, q OfferQuery) (OfferPage, error) {
	if q.Country == "" {
		return OfferPage{}, errors.New("ali: keywordQuery: country is required")
	}
	if q.BeginPage < 1 {
		q.BeginPage = 1
	}
	if q.PageSize < 1 {
		q.PageSize = 20
	}
	body, err := c.Call(ctx, KeywordQuery, Params{"offerQueryParam": q})
	if err != nil {
		return OfferPage{}, err
	}
	if err := checkResultStatus(KeywordQuery.Key(), body); err != nil {
		return OfferPage{}, err
	}
	return DecodeA[OfferPage](KeywordQuery.Key(), body)
}

// -------------------------------------------------------------- keywordSNQuery

// SNQuery is product.search.keywordSNQuery.KeywordSNQueryParams, sent as the
// object parameter snParams. All four fields are documented as required.
type SNQuery struct {
	Keyword  string `json:"keyword"`
	Language string `json:"language"` // "en_US"
	Region   string `json:"region"`   // "US"
	Currency string `json:"currency"` // "USD"
}

// SearchNav is one search-navigation node. The sub-navigation model has the same
// three fields without children, so one type serves both levels.
type SearchNav struct {
	ID            string      `json:"id"` // "973" or "973:28105"
	Name          string      `json:"name"`
	TranslateName string      `json:"translateName"`
	Children      []SearchNav `json:"children,omitempty"`
}

// SearchNav fetches the navigation facets for a keyword. This endpoint reports
// its status as retCode/retMsg rather than code/message.
func (c *Client) SearchNav(ctx context.Context, q SNQuery) ([]SearchNav, error) {
	if q.Keyword == "" || q.Language == "" || q.Region == "" || q.Currency == "" {
		return nil, errors.New("ali: keywordSNQuery: keyword, language, region and currency are all required")
	}
	body, err := c.Call(ctx, KeywordSN, Params{"snParams": q})
	if err != nil {
		return nil, err
	}
	if err := checkResultStatus(KeywordSN.Key(), body); err != nil {
		return nil, err
	}
	return DecodeA[[]SearchNav](KeywordSN.Key(), body)
}

// ---------------------------------------------------------- queryProductDetail

// OfferDetailQuery is product.search.queryProductDetail.param.OfferDetailParam,
// sent as the object parameter offerDetailParam.
type OfferDetailQuery struct {
	OfferID     ID     `json:"offerId"`
	Country     string `json:"country"` // required: language code
	OutMemberID string `json:"outMemberId,omitempty"`
	Currency    string `json:"currency,omitempty"`
}

// OfferDetail is product.search.queryProductDetail.model.ProductDetailModel:
// everything the storefront product page renders.
type OfferDetail struct {
	OfferID             ID              `json:"offerId"`
	CategoryID          ID              `json:"categoryId"`
	CategoryName        string          `json:"categoryName"`
	Subject             string          `json:"subject"`
	SubjectTrans        string          `json:"subjectTrans"`
	Description         string          `json:"description"`
	MainVideo           string          `json:"mainVideo"`
	DetailVideo         string          `json:"detailVideo"`
	ProductImage        ProductImage    `json:"productImage"`
	ProductAttribute    []ProductAttr   `json:"productAttribute"`
	ProductSkuInfos     []SkuInfo       `json:"productSkuInfos"`
	ProductSaleInfo     ProductSaleInfo `json:"productSaleInfo"`
	ProductShippingInfo ShippingInfo    `json:"productShippingInfo"`
	IsJxhy              FlexBool        `json:"isJxhy"`
	SellerOpenID        string          `json:"sellerOpenId"`
	MinOrderQuantity    int64           `json:"minOrderQuantity"`
	BatchNumber         int64           `json:"batchNumber"`
	Status              string          `json:"status"` // only "published" is sellable
	TagInfoList         []ProductTag    `json:"tagInfoList"`
	TraceInfo           string          `json:"traceInfo"`
	SellerMixSetting    MixSetting      `json:"sellerMixSetting"`
	ProductCargoNumber  string          `json:"productCargoNumber"`
	SellerDataInfo      SellerDataInfo  `json:"sellerDataInfo"`
	SoldOut             string          `json:"soldOut"`
	ChannelPrice        ChannelPrice    `json:"channelPrice"`
	PromotionModel      PromotionModel  `json:"promotionModel"`
	TradeScore          string          `json:"tradeScore"`
	TopCategoryID       ID              `json:"topCategoryId"`
	SecondCategoryID    ID              `json:"secondCategoryId"`
	ThirdCategoryID     ID              `json:"thirdCategoryId"`
	SellingPoint        []string        `json:"sellingPoint"`
	OfferIdentities     []string        `json:"offerIdentities"`
	CreateDate          Timestamp       `json:"createDate"`
	IsSelect            FlexBool        `json:"isSelect"` // String here, Boolean in search
	CertificateList     []Certificate   `json:"certificateList"`
	PromotionURL        string          `json:"promotionUrl"` // lower-case url, unlike search
	DescriptionTrans    string          `json:"descriptionTrans"`
	ProductImageTrans   ProductImage    `json:"productImageTrans"`
	CompanyName         string          `json:"companyName"`
	InvoiceInfo         InvoiceInfo     `json:"invoiceInfo"`
}

// ProductImage is product.search.queryProductDetail.model.ProductImage.
type ProductImage struct {
	Images     []string `json:"images"`
	WhiteImage string   `json:"whiteImage"`
}

// ProductAttr is product.search.queryProductDetail.model.ProductAttribute.
type ProductAttr struct {
	AttributeID        string `json:"attributeId"`
	AttributeName      string `json:"attributeName"`
	Value              string `json:"value"`
	AttributeNameTrans string `json:"attributeNameTrans"`
	ValueTrans         string `json:"valueTrans"`
}

// SkuInfo is product.search.queryProductDetail.model.SkuInfo. specId is what the
// order APIs want in cargoParamList, not skuId.
type SkuInfo struct {
	AmountOnSale                  int64            `json:"amountOnSale"`
	Price                         PriceString      `json:"price"`
	JxhyPrice                     PriceString      `json:"jxhyPrice"`
	SkuID                         ID               `json:"skuId"`
	SpecID                        string           `json:"specId"`
	SkuAttributes                 []SkuAttribute   `json:"skuAttributes"`
	PfJxhyPrice                   PriceString      `json:"pfJxhyPrice"`
	ConsignPrice                  PriceString      `json:"consignPrice"`
	CargoNumber                   string           `json:"cargoNumber"`
	PromotionPrice                PriceString      `json:"promotionPrice"`
	FenxiaoPriceInfo              FenxiaoPriceInfo `json:"fenxiaoPriceInfo"`
	ForeignCurrencyPrice          PriceString      `json:"foreignCurrencyPrice"`
	ForeignCurrencyPromotionPrice PriceString      `json:"foreignCurrencyPromotionPrice"`
	RetailPrice                   PriceString      `json:"retailPrice"`
	ForeignCurrencyRetailPrice    PriceString      `json:"foreignCurrencyRetailPrice"`
}

// SkuAttribute is product.search.queryProductDetail.model.SkuAttribute.
type SkuAttribute struct {
	AttributeID        ID     `json:"attributeId"`
	AttributeName      string `json:"attributeName"`
	AttributeNameTrans string `json:"attributeNameTrans"`
	Value              string `json:"value"`
	ValueTrans         string `json:"valueTrans"`
	SkuImageURL        string `json:"skuImageUrl"`
	SkuImageURLTrans   string `json:"skuImageUrlTrans"`
}

// FenxiaoPriceInfo is com.alibaba.cbu.offer.model.FenxiaoPriceInfo.
type FenxiaoPriceInfo struct {
	OnePiecePrice PriceString `json:"onePiecePrice"`
	OfferPrice    PriceString `json:"offerPrice"`
}

// ProductSaleInfo is product.search.queryProductDetail.model.ProductSaleInfo.
// quoteType: 0 no sku, 1 quote by sku, 2 has sku but quoted by product quantity.
type ProductSaleInfo struct {
	AmountOnSale               int64           `json:"amountOnSale"`
	PriceRangeList             []PriceRange    `json:"priceRangeList"`
	ForeignCurrencyPrice       PriceString     `json:"foreignCurrencyPrice"`
	QuoteType                  int             `json:"quoteType"`
	ConsignPrice               PriceString     `json:"consignPrice"`
	JxhyPrice                  PriceString     `json:"jxhyPrice"`
	UnitInfo                   UnitInfo        `json:"unitInfo"`
	FenxiaoSaleInfo            FenxiaoSaleInfo `json:"fenxiaoSaleInfo"`
	RetailPrice                PriceString     `json:"retailPrice"`
	ForeignCurrencyRetailPrice PriceString     `json:"foreignCurrencyRetailPrice"`
}

// PriceRange is one quantity tier: buy startQuantity or more, pay price.
type PriceRange struct {
	StartQuantity                 int64       `json:"startQuantity"`
	Price                         PriceString `json:"price"`
	PromotionPrice                PriceString `json:"promotionPrice"`
	ForeignCurrencyPrice          PriceString `json:"foreignCurrencyPrice"`
	ForeignCurrencyPromotionPrice PriceString `json:"foreignCurrencyPromotionPrice"`
}

// UnitInfo is com.alibaba.cbu.offer.model.UnitInfo.
type UnitInfo struct {
	Unit      string `json:"unit"`
	TransUnit string `json:"transUnit"`
}

// FenxiaoSaleInfo is com.alibaba.cbu.offer.model.FenxiaoSaleInfo.
type FenxiaoSaleInfo struct {
	OnePieceFreePostage FlexBool    `json:"onePieceFreePostage"`
	StartQuantity       int64       `json:"startQuantity"`
	OnePiecePrice       PriceString `json:"onePiecePrice"`
	OfferPrice          PriceString `json:"offerPrice"`
}

// ShippingInfo is product.search.queryProductDetail.model.ProductShippingInfo.
// Weights are kilograms and dimensions centimetres; both are Decimal so the
// international rate card can work in integer grams and millimetres.
type ShippingInfo struct {
	SendGoodsAddressText  string              `json:"sendGoodsAddressText"`
	Weight                Decimal             `json:"weight"`
	Width                 Decimal             `json:"width"`
	Height                Decimal             `json:"height"`
	Length                Decimal             `json:"length"`
	SkuShippingInfoList   []SkuShipping       `json:"skuShippingInfoList"`
	ShippingTimeGuarantee string              `json:"shippingTimeGuarantee"`
	SkuShippingDetails    []SkuShippingDetail `json:"skuShippingDetails"`
	PkgSizeSource         string              `json:"pkgSizeSource"`
	OfficialLength        Decimal             `json:"officialLength"`
	OfficialWidth         Decimal             `json:"officialWidth"`
	OfficialHeight        Decimal             `json:"officialHeight"`
	OfficialWeight        Decimal             `json:"officialWeight"`
}

// SkuShipping is com.alibaba.cbu.offer.model.SkuShippingInfo. Its weight is in
// GRAMS, unlike every other weight in this API, which is in kilograms.
type SkuShipping struct {
	SpecID string  `json:"specId"`
	SkuID  ID      `json:"skuId"`
	Width  Decimal `json:"width"`
	Length Decimal `json:"length"`
	Height Decimal `json:"height"`
	Weight Decimal `json:"weight"`
}

// SkuShippingDetail is com.alibaba.cbu.offer.model.SkuShippingDetail, where the
// skuId arrives as a string and the weight is back in kilograms.
type SkuShippingDetail struct {
	SkuID            StringID `json:"skuId"`
	Width            Decimal  `json:"width"`
	Length           Decimal  `json:"length"`
	Height           Decimal  `json:"height"`
	Weight           Decimal  `json:"weight"`
	PkgSizeSource    string   `json:"pkgSizeSource"`
	OfficialLength   Decimal  `json:"officialLength"`
	OfficialWidth    Decimal  `json:"officialWidth"`
	OfficialHeight   Decimal  `json:"officialHeight"`
	OfficialWeight   Decimal  `json:"officialWeight"`
	AiWeight         Decimal  `json:"aiWeight"`
	AiWeightAccuracy string   `json:"aiWeightAccuracy"`
}

// ProductTag is com.alibaba.cbu.offer.model.out.ProductTagInfo: isOnePsale,
// select and friends.
type ProductTag struct {
	Key   string   `json:"key"`
	Value FlexBool `json:"value"`
}

// MixSetting is product.search.queryProductDetail.model.SellerMixSetting, the
// mixed-batch rule that preview enforces with error 500_006.
type MixSetting struct {
	GeneralHunpi FlexBool `json:"generalHunpi"`
	MixAmount    int64    `json:"mixAmount"`
	MixNumber    int64    `json:"mixNumber"`
}

// ChannelPrice is product.search.queryProductDetail.model.ChannelPrice.
type ChannelPrice struct {
	ChannelSkuPriceList []ChannelSkuPrice `json:"channelSkuPriceList"`
}

// ChannelSkuPrice is one channel price for one sku.
type ChannelSkuPrice struct {
	SkuID        ID          `json:"skuId"`
	CurrentPrice PriceString `json:"currentPrice"`
}

// Certificate is com.alibaba.cbu.offer.model.OfferCertificateModel.
type Certificate struct {
	CertificateName      string   `json:"certificateName"`
	CertificateCode      string   `json:"certificateCode"`
	CertificatePhotoList []string `json:"certificatePhotoList"`
}

// OfferDetail fetches one product in one language.
func (c *Client) OfferDetail(ctx context.Context, q OfferDetailQuery) (OfferDetail, error) {
	if q.OfferID == 0 {
		return OfferDetail{}, errors.New("ali: queryProductDetail: offerId is required")
	}
	if q.Country == "" {
		return OfferDetail{}, errors.New("ali: queryProductDetail: country is required")
	}
	body, err := c.Call(ctx, ProductDetail, Params{"offerDetailParam": q})
	if err != nil {
		return OfferDetail{}, err
	}
	if err := checkResultStatus(ProductDetail.Key(), body); err != nil {
		return OfferDetail{}, err
	}
	return DecodeA[OfferDetail](ProductDetail.Key(), body)
}

// ------------------------------------------------- category.translation.getById

// Category is category.translation.getById.Category. The child model carries the
// same fields without children of its own, so one type covers both levels.
//
// leaf and parentCateId are typed String on the parent and Boolean/Long on the
// children, which is why they are FlexBool and ID here.
type Category struct {
	CategoryID     ID         `json:"categoryId"`
	ChineseName    string     `json:"chineseName"`
	TranslatedName string     `json:"translatedName"`
	Language       string     `json:"language"`
	Leaf           FlexBool   `json:"leaf"`
	Level          string     `json:"level"`
	ParentCateID   ID         `json:"parentCateId"`
	FromCache      FlexBool   `json:"fromCache"`
	Children       []Category `json:"children"`
}

// CategoryByID reads one category and its children in the requested language.
// Category 0 is the root, which is how the catalogue tree is walked. Unlike its
// siblings this endpoint takes flat parameters, not one object.
func (c *Client) CategoryByID(ctx context.Context, language string, categoryID, parentCateID ID) (Category, error) {
	if language == "" {
		return Category{}, errors.New("ali: category.translation.getById: language is required")
	}
	p := Params{"language": language, "categoryId": categoryID}
	if parentCateID != 0 {
		p["parentCateId"] = parentCateID
	}
	body, err := c.Call(ctx, CategoryByID, p)
	if err != nil {
		return Category{}, err
	}
	if err := checkResultStatus(CategoryByID.Key(), body); err != nil {
		return Category{}, err
	}
	return DecodeA[Category](CategoryByID.Key(), body)
}

// ------------------------------------------------------ product.freight.estimate

// FreightQuery is product.freight.estimate.ProductFreightQueryParamsNew, sent as
// the object parameter productFreightQueryParamsNew.
//
// toCountryCode is NOT a country: it is the Chinese DISTRICT code, e.g. 330108
// for Binjiang. The three codes together are province/city/district.
type FreightQuery struct {
	OfferID               ID              `json:"offerId"`
	ToProvinceCode        string          `json:"toProvinceCode"`
	ToCityCode            string          `json:"toCityCode"`
	ToCountryCode         string          `json:"toCountryCode"`
	TotalNum              int64           `json:"totalNum"`
	LogisticsSkuNumModels []FreightSkuNum `json:"logisticsSkuNumModels"`
}

// AddSku appends one sku quantity to the query.
func (q *FreightQuery) AddSku(skuID ID, number int64) {
	q.LogisticsSkuNumModels = append(q.LogisticsSkuNumModels, FreightSkuNum{SkuID: StringID(skuID), Number: number})
}

// FreightSkuNum is product.freight.estimate.LogisticsSkuNumModel. skuId is
// documented as a String, so it goes out quoted.
type FreightSkuNum struct {
	SkuID  StringID `json:"skuId"`
	Number int64    `json:"number"`
}

// ProductFreight is product.freight.estimate.ProductFreightModel. Fees are yuan
// decimal strings; templateType 1 means the seller pays.
type ProductFreight struct {
	OfferID                     ID               `json:"offerId"`
	Freight                     PriceString      `json:"freight"`
	TemplateID                  ID               `json:"templateId"`
	SingleProductWeight         Decimal          `json:"singleProductWeight"`
	SingleProductWidth          Decimal          `json:"singleProductWidth"`
	SingleProductHeight         Decimal          `json:"singleProductHeight"`
	SingleProductLength         Decimal          `json:"singleProductLength"`
	TemplateType                int              `json:"templateType"`
	TemplateName                string           `json:"templateName"`
	SubTemplateType             int              `json:"subTemplateType"`
	SubTemplateName             string           `json:"subTemplateName"`
	FirstFee                    PriceString      `json:"firstFee"`
	FirstUnit                   string           `json:"firstUnit"`
	NextFee                     PriceString      `json:"nextFee"`
	NextUnit                    string           `json:"nextUnit"`
	Discount                    string           `json:"discount"`
	ChargeType                  string           `json:"chargeType"` // 0 weight, 1 quantity, 2 volume
	FreePostage                 FlexBool         `json:"freePostage"`
	ProductFreightSkuInfoModels []FreightSkuInfo `json:"productFreightSkuInfoModels"`
	SizeValueType               int              `json:"sizeValueType"`
}

// FreightSkuInfo is product.freight.estimate.ProductFreightSkuInfoModel.
type FreightSkuInfo struct {
	SkuID           StringID `json:"skuId"`
	SingleSkuWeight Decimal  `json:"singleSkuWeight"`
	SingleSkuWidth  Decimal  `json:"singleSkuWidth"`
	SingleSkuHeight Decimal  `json:"singleSkuHeight"`
	SingleSkuLength Decimal  `json:"singleSkuLength"`
}

// EstimateFreight quotes the domestic leg to a mainland China address, which for
// this platform is the consolidation warehouse.
func (c *Client) EstimateFreight(ctx context.Context, q FreightQuery) (ProductFreight, error) {
	if q.OfferID == 0 {
		return ProductFreight{}, errors.New("ali: product.freight.estimate: offerId is required")
	}
	if q.TotalNum < 1 {
		q.TotalNum = 1
	}
	body, err := c.Call(ctx, FreightEstimate, Params{"productFreightQueryParamsNew": q})
	if err != nil {
		return ProductFreight{}, err
	}
	if err := checkResultStatus(FreightEstimate.Key(), body); err != nil {
		return ProductFreight{}, err
	}
	return DecodeA[ProductFreight](FreightEstimate.Key(), body)
}
