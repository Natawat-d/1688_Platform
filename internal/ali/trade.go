package ali

import (
	"context"
	"encoding/json"
	"errors"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// The order pipeline: preview, create, pay, read, list, cancel.
//
// MONEY UNITS DIFFER BY FAMILY AND THE DIFFERENCE IS NOT COSMETIC.
//
//	trade creation (preview, create, pay)  fen, integer      -> Fen
//	order reading  (detail, list)          yuan, decimal     -> YuanFen
//	baseInfo.discount, baseInfo.refundPayment                -> Fen, see below
//
// Each of these endpoints has its own response envelope, so none of them share a
// decoder: preview, create, pay url, list and cancel are all bespoke, and only
// trade.get.buyerView fits family B.

// WebSite1688 is the site selector every trade API demands. The alternative,
// "alibaba", is the international site and is not what this platform buys from.
const WebSite1688 = "1688"

// Documented cancelReason values for alibaba.trade.cancel.
const (
	CancelBuyerCancel     = "buyerCancel"
	CancelSellerGoodsLack = "sellerGoodsLack"
	CancelOther           = "other"
)

// Documented flow values. The preview picks the cheapest when flow is empty;
// createCrossOrder requires one, so copy flowFlag from the preview.
const (
	FlowGeneral         = "general"
	FlowFenxiao         = "fenxiao"
	FlowPaired          = "paired"
	FlowRepurchase      = "repurchase"
	FlowSaleProxy       = "saleproxy"
	FlowBoutiqueFenxiao = "boutiquefenxiao"
	FlowBoutiquePifa    = "boutiquepifa"
)

// Order status values from baseInfo.status, the ones this platform maps.
const (
	StatusWaitBuyerPay     = "waitbuyerpay"
	StatusWaitSellerSend   = "waitsellersend"
	StatusWaitBuyerReceive = "waitbuyerreceive"
	StatusConfirmGoods     = "confirm_goods"
	StatusSuccess          = "success"
	StatusCancel           = "cancel"
	StatusTerminated       = "terminated"
)

// Timestamp is a 1688 date field. The gateway uses four different spellings for
// these — "20180614101942000+0800", the same without an offset, a naive
// "2018-07-24 21:55:33" that means Shanghai, and epoch milliseconds — so the
// text is kept verbatim and Time() does the parsing through ParseTime.
type Timestamp string

func (t Timestamp) String() string { return string(t) }

func (t Timestamp) IsZero() bool { return strings.TrimSpace(string(t)) == "" }

// Time parses the timestamp. A zero Time is returned for an empty field.
func (t Timestamp) Time() (time.Time, error) { return ParseTime(string(t)) }

// UnmarshalJSON accepts a quoted timestamp or a bare epoch-millisecond number.
func (t *Timestamp) UnmarshalJSON(b []byte) error {
	s := strings.TrimSpace(string(b))
	if s == "null" {
		*t = ""
		return nil
	}
	if len(s) > 0 && s[0] == '"' {
		var str string
		if err := json.Unmarshal(b, &str); err != nil {
			return err
		}
		*t = Timestamp(str)
		return nil
	}
	*t = Timestamp(s)
	return nil
}

// ------------------------------------------------------------ request objects

// Address is alibaba.trade.fast.address, shared by preview and create.
//
// The gateway resolves it in three tiers and stops at the first that is filled:
// addressId, then districtCode plus street, then province/city/area text.
type Address struct {
	AddressID    ID     `json:"addressId,omitempty"`
	FullName     string `json:"fullName,omitempty"`
	Mobile       string `json:"mobile,omitempty"`
	Phone        string `json:"phone,omitempty"`
	PostCode     string `json:"postCode,omitempty"`
	CityText     string `json:"cityText,omitempty"`
	ProvinceText string `json:"provinceText,omitempty"`
	AreaText     string `json:"areaText,omitempty"`
	TownText     string `json:"townText,omitempty"`
	Address      string `json:"address,omitempty"`
	DistrictCode string `json:"districtCode,omitempty"`
	AddressCode  string `json:"addressCode,omitempty"`
}

// Cargo is alibaba.trade.fast.cargo: one sku of one offer.
//
// quantity is documented as a Double because 1688 sells by the ton as well as by
// the piece, so it is a json.Number here: whole counts and fractional units both
// go out as bare numbers and no float ever touches the value.
type Cargo struct {
	OfferID     ID          `json:"offerId,omitempty"`
	SpecID      string      `json:"specId,omitempty"`
	Quantity    json.Number `json:"quantity,omitempty"`
	OpenOfferID string      `json:"openOfferId,omitempty"`
	OutMemberID string      `json:"outMemberId,omitempty"`
}

// NewCargo builds a cargo line for a whole number of units.
func NewCargo(offerID ID, specID string, quantity int64) Cargo {
	return Cargo{OfferID: offerID, SpecID: specID, Quantity: json.Number(strconv.FormatInt(quantity, 10))}
}

// Invoice is alibaba.trade.fast.invoice. invoiceType 0 is a regular invoice and
// 1 a VAT invoice.
type Invoice struct {
	InvoiceType        int    `json:"invoiceType"`
	ProvinceText       string `json:"provinceText,omitempty"`
	CityText           string `json:"cityText,omitempty"`
	AreaText           string `json:"areaText,omitempty"`
	TownText           string `json:"townText,omitempty"`
	PostCode           string `json:"postCode,omitempty"`
	Address            string `json:"address,omitempty"`
	FullName           string `json:"fullName,omitempty"`
	Phone              string `json:"phone,omitempty"`
	Mobile             string `json:"mobile,omitempty"`
	CompanyName        string `json:"companyName,omitempty"`
	TaxpayerIdentifier string `json:"taxpayerIdentifier,omitempty"`
	BankAndAccount     string `json:"bankAndAccount,omitempty"`
	LocalInvoiceID     string `json:"localInvoiceId,omitempty"`
}

// ---------------------------------------------------------- createOrder.preview

// PreviewRequest is the parameter set of alibaba.createOrder.preview. Everything
// but the address and the cargo list is optional; leaving Flow empty asks 1688
// to compare the channels and return the cheapest as flowFlag.
type PreviewRequest struct {
	Address                        Address
	Cargo                          []Cargo
	Invoice                        *Invoice
	Flow                           string
	InstanceID                     string
	ProxySettleRecordID            string
	InventoryMode                  string
	OutOrderID                     string
	PickupService                  string
	CrossBorderLogisticsSolutionID string
	UseBorderLogisticsSolution     bool
	IsvBizType                     string
}

// Params renders the request as gateway parameters. Objects and arrays are
// JSON-stringified by the client.
func (r PreviewRequest) Params() Params {
	p := Params{
		"addressParam":   r.Address,
		"cargoParamList": r.Cargo,
	}
	if r.Invoice != nil {
		p["invoiceParam"] = r.Invoice
	}
	setStr(p, "flow", r.Flow)
	setStr(p, "instanceId", r.InstanceID)
	setStr(p, "proxySettleRecordId", r.ProxySettleRecordID)
	setStr(p, "inventoryMode", r.InventoryMode)
	setStr(p, "outOrderId", r.OutOrderID)
	setStr(p, "pickupService", r.PickupService)
	setStr(p, "crossBorderLogisticsSolutionId", r.CrossBorderLogisticsSolutionID)
	setStr(p, "isvBizType", r.IsvBizType)
	if r.UseBorderLogisticsSolution {
		p["useBorderLogisticsSolution"] = true
	}
	return p
}

// Preview is the whole response of alibaba.createOrder.preview.
//
// orderPreviewResuslt is spelled exactly like that in the documentation and on
// the wire. Do not fix the typo.
type Preview struct {
	Results                            []PreviewedOrder `json:"orderPreviewResuslt"`
	Success                            FlexBool         `json:"success"`
	ErrorCode                          string           `json:"errorCode"`
	ErrorMsg                           string           `json:"errorMsg"` // errorMsg, not errorMessage
	PostFeeByDescOfferList             IDs              `json:"postFeeByDescOfferList"`
	ConsignOfferList                   IDs              `json:"consignOfferList"`
	UnsupportedCrossBorderPayOfferList IDs              `json:"unsupportedCrossBorderPayOfferList"`
	ExtPairList                        []ExtPair        `json:"extPairList"`
}

// PreviewedOrder is alibaba.createOrder.preview.result.model: one order as it
// would be created. Several entries mean 1688 would split the basket.
//
// All the amounts here are FEN.
type PreviewedOrder struct {
	DiscountFee               Fen                 `json:"discountFee"`
	TradeModeNameList         []string            `json:"tradeModeNameList"`
	Status                    FlexBool            `json:"status"`
	TaoSampleSinglePromotion  FlexBool            `json:"taoSampleSinglePromotion"`
	SumPayment                Fen                 `json:"sumPayment"`
	Message                   string              `json:"message"`
	SumCarriage               Fen                 `json:"sumCarriage"`
	ResultCode                string              `json:"resultCode"`
	SumPaymentNoCarriage      Fen                 `json:"sumPaymentNoCarriage"`
	AdditionalFee             Fen                 `json:"additionalFee"`
	FlowFlag                  string              `json:"flowFlag"`
	CargoList                 []PreviewedCargo    `json:"cargoList"`
	ShopPromotionList         []Promotion         `json:"shopPromotionList"`
	TradeModelList            []TradeMode         `json:"tradeModelList"`
	PayChannelInfos           []PayChannel        `json:"payChannelInfos"`
	TradeServiceList          []TradeServiceGroup `json:"tradeServiceList"`
	ExtPairList               []ExtPair           `json:"extPairList"`
	OrderGroup                string              `json:"orderGroup"`
	CanUseOfficialSolution    FlexBool            `json:"canUseOfficialSolution"`
	OfficialSolutionModelList []OfficialSolution  `json:"officialSolutionModelList"`
	TotalFundUsageAmount      Fen                 `json:"totalFundUsageAmount"`
	TotalPostFundUsageAmount  Fen                 `json:"totalPostFundUsageAmount"`
}

// TradeTypes are the tradeType values this preview allows the order to be
// created with. An empty list means the order cannot be placed through the API
// at all and has to go through the 1688 web checkout.
func (p PreviewedOrder) TradeTypes() []string { return p.TradeModeNameList }

// PreviewedCargo is alibaba.createOrder.preview.resultCargo.model. Unlike its
// parent, amount and finalUnitPrice are YUAN decimals.
type PreviewedCargo struct {
	Amount                   YuanFen             `json:"amount"`
	Message                  string              `json:"message"`
	FinalUnitPrice           YuanFen             `json:"finalUnitPrice"`
	SpecID                   string              `json:"specId"`
	SkuID                    ID                  `json:"skuId"`
	ResultCode               string              `json:"resultCode"`
	OfferID                  ID                  `json:"offerId"`
	OpenOfferID              string              `json:"openOfferId"`
	CargoPromotionList       []Promotion         `json:"cargoPromotionList"`
	ExtPairList              []ExtPair           `json:"extPairList"`
	TotalFundUsageAmount     Fen                 `json:"totalFundUsageAmount"`
	TotalPostFundUsageAmount Fen                 `json:"totalPostFundUsageAmount"`
	TradeServiceList         []TradeServiceGroup `json:"tradeServiceList"`
}

// Promotion is alibaba.trade.promotion.model, a shop or item level discount.
type Promotion struct {
	PromotionID string   `json:"promotionId"`
	Selected    FlexBool `json:"selected"`
	Text        string   `json:"text"`
	Desc        string   `json:"desc"`
	FreePostage FlexBool `json:"freePostage"`
	DiscountFee Fen      `json:"discountFee"`
}

// TradeMode is tradeModelExtensionList. Only entries with opSupport true may be
// passed as tradeType to createCrossOrder.
type TradeMode struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	TradeType   string   `json:"tradeType"`
	OpSupport   FlexBool `json:"opSupport"`
}

// PayChannel is payChaneelList — the documented spelling of the field, typo and
// all. amountLimit is fen.
type PayChannel struct {
	Name        string `json:"name"`
	AmountLimit Fen    `json:"amountLimit"`
}

// TradeServiceGroup is a group of value-added services offered with the order.
type TradeServiceGroup struct {
	OutID             string         `json:"outId"`
	GroupTitle        string         `json:"groupTitle"`
	GroupCode         string         `json:"groupCode"`
	GroupDescLink     string         `json:"groupDescLink"`
	TradeServiceList  []TradeService `json:"tradeServiceList"`
	ExtraMapStr       string         `json:"extraMapStr"`
	CombinedPrice     Fen            `json:"combinedPrice"`
	ShowCombinedPrice string         `json:"showCombinedPrice"`
}

// TradeService is one value-added service. unitPrice is fen.
type TradeService struct {
	SourceType  string  `json:"sourceType"`
	ServiceType string  `json:"serviceType"`
	SourceID    string  `json:"sourceId"`
	Title       string  `json:"title"`
	UnitPrice   Fen     `json:"unitPrice"`
	Quantity    Decimal `json:"quantity"`
	ServiceTips string  `json:"serviceTips"`
	ServiceURL  string  `json:"serviceUrl"`
	GroupTitle  string  `json:"groupTitle"`
	GroupCode   string  `json:"groupCode"`
	SponsorDesc string  `json:"sponsorDesc"`
	Code        string  `json:"code"`
	ExtMapStr   string  `json:"extMapStr"`
}

// OfficialSolution is one official door-to-door pickup plan. totalCost is fen.
type OfficialSolution struct {
	SolutionCode string `json:"solutionCode"`
	SolutionName string `json:"solutionName"`
	TotalCost    Fen    `json:"totalCost"`
}

// ExtPair is the extension key/value pair that turns up in half the trade
// payloads. value is itself a JSON document more often than not.
type ExtPair struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Desc  string `json:"desc"`
}

// PreviewOrder validates a basket against 1688 and returns the prices, the
// discounts and the trade modes the order could be created with. It is also the
// only place the documented 500_00x refusals surface, so its error is worth
// showing to a human verbatim.
func (c *Client) PreviewOrder(ctx context.Context, r PreviewRequest) (Preview, error) {
	if len(r.Cargo) == 0 {
		return Preview{}, errors.New("ali: createOrder.preview: cargoParamList is empty")
	}
	body, err := c.Call(ctx, OrderPreview, r.Params())
	if err != nil {
		return Preview{}, err
	}
	var out Preview
	if err := DecodeFlat(OrderPreview.Key(), body, &out); err != nil {
		return Preview{}, err
	}
	if out.ErrorCode != "" || (len(out.Results) == 0 && !out.Success.Bool()) {
		return out, &APIError{
			API:     OrderPreview.Key(),
			Code:    firstNonEmpty(out.ErrorCode, "PREVIEW_FAILED"),
			Message: out.ErrorMsg,
			Body:    truncate(body),
		}
	}
	return out, nil
}

// ------------------------------------------------------------ createCrossOrder

// CreateOrderRequest is the parameter set of alibaba.trade.createCrossOrder.
// Flow is required here even though the preview treats it as optional: pass the
// flowFlag the preview returned.
type CreateOrderRequest struct {
	Flow                            string
	Address                         Address
	Cargo                           []Cargo
	Invoice                         *Invoice
	Message                         string
	TradeType                       string
	ShopPromotionID                 string
	AnonymousBuyer                  bool
	FenxiaoChannel                  string
	InventoryMode                   string
	OutOrderID                      string
	PickupService                   string
	WarehouseCode                   string
	PreSelectPayChannel             string
	SmallProcurement                string
	UseRedEnvelope                  string
	Dropshipping                    string
	AddedService                    string
	CrossBorderLogisticsSolutionID  string
	UseCrossBorderLogisticsSolution bool
	UseOfficialSolution             bool
	UseOfficialSolutionModelList    []UseOfficialSolution
	IsvBizType                      string
	FromAgent                       string
}

// UseOfficialSolution names the pickup plan chosen for one order group.
type UseOfficialSolution struct {
	OrderGroup              string `json:"orderGroup"`
	UseOfficialSolutionCode string `json:"useOfficialSolutionCode"`
}

// Params renders the request as gateway parameters.
func (r CreateOrderRequest) Params() Params {
	p := Params{
		"flow":           r.Flow,
		"addressParam":   r.Address,
		"cargoParamList": r.Cargo,
	}
	if r.Invoice != nil {
		p["invoiceParam"] = r.Invoice
	}
	setStr(p, "message", r.Message)
	setStr(p, "tradeType", r.TradeType)
	setStr(p, "shopPromotionId", r.ShopPromotionID)
	setStr(p, "fenxiaoChannel", r.FenxiaoChannel)
	setStr(p, "inventoryMode", r.InventoryMode)
	setStr(p, "outOrderId", r.OutOrderID)
	setStr(p, "pickupService", r.PickupService)
	setStr(p, "warehouseCode", r.WarehouseCode)
	setStr(p, "preSelectPayChannel", r.PreSelectPayChannel)
	setStr(p, "smallProcurement", r.SmallProcurement)
	setStr(p, "useRedEnvelope", r.UseRedEnvelope)
	setStr(p, "dropshipping", r.Dropshipping)
	setStr(p, "addedService", r.AddedService)
	setStr(p, "crossBorderLogisticsSolutionId", r.CrossBorderLogisticsSolutionID)
	setStr(p, "isvBizType", r.IsvBizType)
	setStr(p, "fromAgent", r.FromAgent)
	if r.AnonymousBuyer {
		p["anonymousBuyer"] = true
	}
	if r.UseCrossBorderLogisticsSolution {
		p["useCrossBorderLogisticsSolution"] = true
	}
	if r.UseOfficialSolution {
		p["useOfficialSolution"] = true
	}
	if len(r.UseOfficialSolutionModelList) > 0 {
		p["useOfficialSolutionModelList"] = r.UseOfficialSolutionModelList
	}
	return p
}

// CreateResult is alibaba.trade.cross.result, the payload of createCrossOrder.
//
// It arrives two ways: wrapped in a "result" object beside an outer
// success/code/message, and flat at the top level. DecodeFlat accepts both.
// When several orders are created at once, orderId, postFee and
// totalSuccessAmount are empty and everything is in orderList instead.
type CreateResult struct {
	TotalSuccessAmount Fen            `json:"totalSuccessAmount"`
	OrderID            ID             `json:"orderId"` // a STRING on the wire
	Success            FlexBool       `json:"success"`
	Code               string         `json:"code"`
	Message            string         `json:"message"`
	AccountPeriod      *AccountPeriod `json:"accountPeriod"`
	FailedOfferList    []FailedOffer  `json:"failedOfferList"`
	PostFee            Fen            `json:"postFee"`
	OrderList          []CreatedOrder `json:"orderList"`
}

// OrderIDs is every 1688 order this call created, whether it answered with one
// order or with a list.
func (r CreateResult) OrderIDs() IDs {
	if len(r.OrderList) > 0 {
		out := make(IDs, 0, len(r.OrderList))
		for _, o := range r.OrderList {
			if o.OrderID != 0 {
				out = append(out, o.OrderID)
			}
		}
		return out
	}
	if r.OrderID != 0 {
		return IDs{r.OrderID}
	}
	return nil
}

// AccountPeriod is alibaba.trade.cross.period, present only for account-period
// payment. tapType 1 monthly, 3 bi-monthly, 6 quarterly, 5 by receipt date.
type AccountPeriod struct {
	TapType    int `json:"tapType"`
	TapDate    int `json:"tapDate"`
	TapOverdue int `json:"tapOverdue"`
}

// FailedOffer is alibaba.trade.fast.offer: a line 1688 refused.
type FailedOffer struct {
	OfferID      ID     `json:"offerId"`
	SpecID       string `json:"specId"`
	ErrorCode    string `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
}

// CreatedOrder is alibaba.tradeResult.BizSimpleOrder, one entry of orderList.
//
// orderAmmount is the documented spelling. The last four fields are undocumented
// but present in the multi-order sample, and the order relay wants them: they
// are the only place a per-order discount shows up.
type CreatedOrder struct {
	PostFee                        Fen      `json:"postFee"`
	OrderAmmount                   Fen      `json:"orderAmmount"`
	Message                        string   `json:"message"`
	ResultCode                     string   `json:"resultCode"`
	Success                        FlexBool `json:"success"`
	OrderID                        ID       `json:"orderId"`
	PayChannel                     string   `json:"payChannel"`
	Discount                       Fen      `json:"discount"`
	SumPaymentNoCarriageFromClient Fen      `json:"sumPaymentNoCarriageFromClient"`
	MergePay                       FlexBool `json:"mergePay"`
	ChooseFreeFreight              FlexBool `json:"chooseFreeFreight"`
}

// CreateOrder places one cross-border order. All the cargo must belong to one
// supplier and there may be at most fifty skus; split the basket first.
//
// outOrderId is 1688's idempotency key and is searchable through
// getBuyerOrderList, so always send our own order id in it.
func (c *Client) CreateOrder(ctx context.Context, r CreateOrderRequest) (CreateResult, error) {
	if r.Flow == "" {
		return CreateResult{}, errors.New("ali: createCrossOrder: flow is required")
	}
	if len(r.Cargo) == 0 {
		return CreateResult{}, errors.New("ali: createCrossOrder: cargoParamList is empty")
	}
	body, err := c.Call(ctx, CreateCrossOrder, r.Params())
	if err != nil {
		return CreateResult{}, err
	}

	var out CreateResult
	if err := DecodeFlat(CreateCrossOrder.Key(), body, &out); err != nil {
		return CreateResult{}, err
	}
	// The wrapped spelling carries its own status beside the result object.
	var outer struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}
	_ = json.Unmarshal(body, &outer)

	// success is false in two of the documented samples that plainly succeeded,
	// so the test is whether an order came back, not what the flag says.
	if len(out.OrderIDs()) == 0 {
		return out, &APIError{
			API:     CreateCrossOrder.Key(),
			Code:    firstNonEmpty(out.Code, outer.Code, "CREATE_ORDER_FAILED"),
			Message: firstNonEmpty(out.Message, outer.Message),
			Body:    truncate(body),
		}
	}
	return out, nil
}

// ---------------------------------------------------------------- alipay.url.get

// PayURL is the response of alibaba.alipay.url.get. erroMsg is the documented
// spelling; it is not errorMsg and not errorMessage.
type PayURL struct {
	ErroMsg             string   `json:"erroMsg"`
	PayURL              string   `json:"payUrl"`
	Success             FlexBool `json:"success"`
	ErrorCode           string   `json:"errorCode"`
	PayFailureOrderList IDs      `json:"payFailureOrderList"`
}

// AlipayURL returns a cashier link for one or more orders: the 1688 cashier for
// a single order, the Alipay cashier for a batch. At most 100 orders, or 30 for
// Kuajingbao.
//
// orderIdList is a Long[], which is why ID marshals as a bare number. Quoting
// the ids breaks this call.
func (c *Client) AlipayURL(ctx context.Context, orderIDs IDs) (PayURL, error) {
	if len(orderIDs) == 0 {
		return PayURL{}, errors.New("ali: alipay.url.get: orderIdList is empty")
	}
	body, err := c.Call(ctx, AlipayURLGet, Params{"orderIdList": orderIDs})
	if err != nil {
		return PayURL{}, err
	}
	var out PayURL
	if err := DecodeFlat(AlipayURLGet.Key(), body, &out); err != nil {
		return PayURL{}, err
	}
	if out.PayURL == "" {
		return out, &APIError{
			API:     AlipayURLGet.Key(),
			Code:    firstNonEmpty(out.ErrorCode, "NO_PAY_URL"),
			Message: out.ErroMsg,
			Body:    truncate(body),
		}
	}
	return out, nil
}

// ------------------------------------------------------------ trade.get.buyerView

// TradeInfo is alibaba.openplatform.trade.model.TradeInfo, returned by both the
// order detail and the order list APIs.
//
// The documented model runs to some six hundred fields. What is modelled here is
// what the order pipeline and the customer timeline read; everything else lands
// in Extra as raw JSON, so nothing is silently lost and adding a field later is
// a one-line change.
type TradeInfo struct {
	BaseInfo        OrderBaseInfo   `json:"baseInfo"`
	TradeTerms      []TradeTerm     `json:"tradeTerms"`
	ProductItems    []ProductItem   `json:"productItems"`
	NativeLogistics NativeLogistics `json:"nativeLogistics"`
	GuaranteesTerms []GuaranteeTerm `json:"guaranteesTerms"`
	ExtAttributes   []KeyValuePair  `json:"extAttributes"`

	// Extra holds every top-level member that is not modelled above.
	Extra map[string]json.RawMessage `json:"-"`
}

func (t *TradeInfo) UnmarshalJSON(b []byte) error {
	type raw TradeInfo
	var v raw
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	*t = TradeInfo(v)
	t.Extra = unknownFields(b, v)
	return nil
}

// OrderBaseInfo is alibaba.openplatform.trade.model.OrderBaseInfo.
//
// MONEY: totalAmount, sumProductPayment, shippingFee, couponFee and refund are
// YUAN decimals. discount and refundPayment are FEN integers — the same object
// mixes both units, and reading discount as yuan overstates it a hundredfold.
type OrderBaseInfo struct {
	ID ID `json:"id"`
	// IDOfStr is the id as text. It is FlexString because the documented
	// getBuyerOrderList sample sends it as a bare number while the order
	// detail sample quotes it; and it is worth reading because the same
	// detail sample shows id rounded to ...944 beside idOfStr "...941".
	IDOfStr           FlexString `json:"idOfStr"`
	PreOrderID        ID         `json:"preOrderId"`
	OutOrderID        string     `json:"outOrderId"`
	Status            string     `json:"status"`
	RefundStatus      string     `json:"refundStatus"`
	RefundStatusForAs string     `json:"refundStatusForAs"`
	RefundID          string     `json:"refundId"`
	BusinessType      string     `json:"businessType"`
	TradeType         string     `json:"tradeType"`
	TradeTypeCode     string     `json:"tradeTypeCode"`
	TradeTypeDesc     string     `json:"tradeTypeDesc"`
	FlowTemplateCode  string     `json:"flowTemplateCode"`

	// Amounts in yuan.
	TotalAmount       YuanFen `json:"totalAmount"`
	SumProductPayment YuanFen `json:"sumProductPayment"`
	ShippingFee       YuanFen `json:"shippingFee"`
	CouponFee         YuanFen `json:"couponFee"`
	Refund            YuanFen `json:"refund"`

	// Amounts in fen, despite every sibling above being yuan.
	Discount      Fen `json:"discount"`
	RefundPayment Fen `json:"refundPayment"`

	// Times, all in the compact "20180614101942000+0800" spelling.
	CreateTime       Timestamp `json:"createTime"`
	ModifyTime       Timestamp `json:"modifyTime"`
	PayTime          Timestamp `json:"payTime"`
	AllDeliveredTime Timestamp `json:"allDeliveredTime"`
	ReceivingTime    Timestamp `json:"receivingTime"`
	ConfirmedTime    Timestamp `json:"confirmedTime"`
	CompleteTime     Timestamp `json:"completeTime"`

	// Parties.
	BuyerID           string        `json:"buyerID"`
	BuyerUserID       ID            `json:"buyerUserId"`
	BuyerLoginID      string        `json:"buyerLoginId"`
	SubBuyerLoginID   string        `json:"subBuyerLoginId"`
	BuyerAlipayID     string        `json:"buyerAlipayId"`
	BuyerContact      Contact       `json:"buyerContact"`
	SellerID          string        `json:"sellerID"`
	SellerUserID      ID            `json:"sellerUserId"`
	SellerLoginID     string        `json:"sellerLoginId"`
	SellerAlipayID    string        `json:"sellerAlipayId"`
	SellerCreditLevel string        `json:"sellerCreditLevel"`
	SellerContact     SellerContact `json:"sellerContact"`
	ReceiverInfo      ReceiverInfo  `json:"receiverInfo"`

	// Closing and messages.
	CloseReason      string   `json:"closeReason"`
	CloseOperateType string   `json:"closeOperateType"`
	CloseRemark      string   `json:"closeRemark"`
	Remark           string   `json:"remark"`
	BuyerFeedback    string   `json:"buyerFeedback"`
	BuyerMemo        string   `json:"buyerMemo"`
	BuyerRemarkIcon  string   `json:"buyerRemarkIcon"`
	AlipayTradeID    string   `json:"alipayTradeId"`
	SellerOrder      FlexBool `json:"sellerOrder"`
	OverSeaOrder     FlexBool `json:"overSeaOrder"`
	StepPayAll       FlexBool `json:"stepPayAll"`

	// Payment and fulfilment.
	PayChannelList          []string `json:"payChannelList"`
	PayChannelCodeList      []string `json:"payChannelCodeList"`
	PayTimeout              int64    `json:"payTimeout"`     // seconds
	PayTimeoutType          int      `json:"payTimeoutType"` // 0 fixed length, 1 fixed time
	InventoryMode           string   `json:"inventoryMode"`
	OfficialSolutionOrderID string   `json:"officialSolutionOrderId"`
	OfficialSolutionCost    Fen      `json:"officialSolutionCost"`

	// Extra holds every top-level member of baseInfo that is not modelled above,
	// including the staged-payment lists.
	Extra map[string]json.RawMessage `json:"-"`
}

func (b *OrderBaseInfo) UnmarshalJSON(data []byte) error {
	type raw OrderBaseInfo
	var v raw
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	*b = OrderBaseInfo(v)
	b.Extra = unknownFields(data, v)
	return nil
}

// Contact is alibaba.trade.tradeContact.
type Contact struct {
	Phone        string `json:"phone"`
	Fax          string `json:"fax"`
	Email        string `json:"email"`
	IMInPlatform string `json:"imInPlatform"`
	Name         string `json:"name"`
	Mobile       string `json:"mobile"`
	CompanyName  string `json:"companyName"`
}

// SellerContact is alibaba.trade.tradeSellerContact: the buyer contact plus the
// shop name and the distribution sender identity.
type SellerContact struct {
	Phone         string `json:"phone"`
	Fax           string `json:"fax"`
	Email         string `json:"email"`
	IMInPlatform  string `json:"imInPlatform"`
	Name          string `json:"name"`
	Mobile        string `json:"mobile"`
	CompanyName   string `json:"companyName"`
	WgSenderName  string `json:"wgSenderName"`
	WgSenderPhone string `json:"wgSenderPhone"`
	ShopName      string `json:"shopName"`
}

// ReceiverInfo is alibaba.trade.orderReceiverInfo. toArea is the whole address
// as one line; toDivisionCode is the district code.
type ReceiverInfo struct {
	ToFullName     string `json:"toFullName"`
	ToDivisionCode string `json:"toDivisionCode"`
	ToMobile       string `json:"toMobile"`
	ToPhone        string `json:"toPhone"`
	ToPost         string `json:"toPost"`
	ToTownCode     string `json:"toTownCode"`
	ToArea         string `json:"toArea"`
}

// TradeTerm is alibaba.openplatform.trade.model.TradeTermsInfo, one payment
// stage. On 1688 payStatus is a number in a string: 1 unpaid, 2 paid, 4 fully
// refunded, 6 collected, 8 cancelled before payment, 9 in progress.
type TradeTerm struct {
	PayStatus      string    `json:"payStatus"`
	PayTime        Timestamp `json:"payTime"`
	PayWay         string    `json:"payWay"`
	PayWayDesc     string    `json:"payWayDesc"`
	PhasAmount     YuanFen   `json:"phasAmount"` // documented spelling
	Phase          ID        `json:"phase"`
	PhaseCondition string    `json:"phaseCondition"`
	PhaseDate      string    `json:"phaseDate"`
	CardPay        FlexBool  `json:"cardPay"`
	ExpressPay     FlexBool  `json:"expressPay"`
}

// ProductItem is alibaba.openplatform.trade.model.ProductItemInfo, one order
// line. itemAmount, price, refund and sharePostage are YUAN; entryDiscount is
// fen.
type ProductItem struct {
	SubItemID          ID              `json:"subItemID"`
	SubItemIDString    FlexString      `json:"subItemIDString"` // see OrderBaseInfo.IDOfStr
	ProductID          ID              `json:"productID"`
	SkuID              ID              `json:"skuID"`
	SpecID             string          `json:"specId"`
	Name               string          `json:"name"`
	Description        string          `json:"description"`
	CargoNumber        string          `json:"cargoNumber"`
	ProductCargoNumber string          `json:"productCargoNumber"`
	ProductImgURL      []string        `json:"productImgUrl"`
	ProductSnapshotURL string          `json:"productSnapshotUrl"`
	ItemAmount         YuanFen         `json:"itemAmount"`
	Price              YuanFen         `json:"price"`
	Refund             YuanFen         `json:"refund"`
	SharePostage       YuanFen         `json:"sharePostage"`
	EntryDiscount      Fen             `json:"entryDiscount"`
	Quantity           Decimal         `json:"quantity"`
	QuantityFactor     Decimal         `json:"quantityFactor"`
	Unit               string          `json:"unit"`
	Weight             string          `json:"weight"`
	WeightUnit         string          `json:"weightUnit"`
	Sort               int             `json:"sort"`
	Type               string          `json:"type"`
	Status             string          `json:"status"`
	StatusStr          string          `json:"statusStr"`
	RefundStatus       string          `json:"refundStatus"`
	RefundID           string          `json:"refundId"`
	RefundIDForAs      string          `json:"refundIdForAs"`
	CloseReason        string          `json:"closeReason"`
	LogisticsStatus    int             `json:"logisticsStatus"`
	SkuInfos           []SkuItemDesc   `json:"skuInfos"`
	GuaranteesTerms    []GuaranteeTerm `json:"guaranteesTerms"`
	GmtCreate          Timestamp       `json:"gmtCreate"`
	GmtModified        Timestamp       `json:"gmtModified"`
	GmtCompleted       Timestamp       `json:"gmtCompleted"`
	GmtPayExpireTime   Timestamp       `json:"gmtPayExpireTime"`
}

// SkuItemDesc is alibaba.trade.SkuItemDesc: the sale attributes as text.
type SkuItemDesc struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// GuaranteeTerm is alibaba.openplatform.trade.model.GuaranteeTermsInfo.
type GuaranteeTerm struct {
	AssuranceInfo        string `json:"assuranceInfo"`
	AssuranceType        string `json:"assuranceType"`
	QualityAssuranceType string `json:"qualityAssuranceType"`
	Value                string `json:"value"`
}

// NativeLogistics is alibaba.openplatform.trade.model.NativeLogisticsInfo: the
// domestic shipping address and the waybills raised against it.
type NativeLogistics struct {
	Address        string          `json:"address"`
	Area           string          `json:"area"`
	AreaCode       string          `json:"areaCode"`
	City           string          `json:"city"`
	Province       string          `json:"province"`
	Town           string          `json:"town"`
	TownCode       string          `json:"townCode"`
	Zip            string          `json:"zip"`
	ContactPerson  string          `json:"contactPerson"`
	Mobile         string          `json:"mobile"`
	Telephone      string          `json:"telephone"`
	Fax            string          `json:"fax"`
	LogisticsItems []LogisticsItem `json:"logisticsItems"`
}

// LogisticsItem is alibaba.openplatform.trade.model.NativeLogisticsItemsInfo,
// one waybill. type is "0" self delivery, "1" online delivery, "2" no logistics.
// carriage is yuan.
type LogisticsItem struct {
	ID                   ID        `json:"id"`
	LogisticsCode        string    `json:"logisticsCode"`
	LogisticsBillNo      string    `json:"logisticsBillNo"`
	LogisticsCompanyID   ID        `json:"logisticsCompanyId"`
	LogisticsCompanyNo   string    `json:"logisticsCompanyNo"`
	LogisticsCompanyName string    `json:"logisticsCompanyName"`
	Type                 string    `json:"type"`
	Status               string    `json:"status"`
	Carriage             YuanFen   `json:"carriage"`
	SubItemIDs           string    `json:"subItemIds"`
	DeliveredTime        Timestamp `json:"deliveredTime"`
	GmtCreate            Timestamp `json:"gmtCreate"`
	GmtModified          Timestamp `json:"gmtModified"`
	FromProvince         string    `json:"fromProvince"`
	FromCity             string    `json:"fromCity"`
	FromArea             string    `json:"fromArea"`
	FromAddress          string    `json:"fromAddress"`
	FromPhone            string    `json:"fromPhone"`
	FromMobile           string    `json:"fromMobile"`
	FromPost             string    `json:"fromPost"`
	ToProvince           string    `json:"toProvince"`
	ToCity               string    `json:"toCity"`
	ToArea               string    `json:"toArea"`
	ToAddress            string    `json:"toAddress"`
	ToPhone              string    `json:"toPhone"`
	ToMobile             string    `json:"toMobile"`
	ToPost               string    `json:"toPost"`
	NoLogisticsName      string    `json:"noLogisticsName"`
	NoLogisticsTel       string    `json:"noLogisticsTel"`
	NoLogisticsBillNo    string    `json:"noLogisticsBillNo"`
	NoLogisticsCondition string    `json:"noLogisticsCondition"`
}

// KeyValuePair is alibaba.openplatform.trade.KeyValuePair, used by extAttributes.
type KeyValuePair struct {
	Key         string `json:"key"`
	Value       string `json:"value"`
	Description string `json:"description"`
}

// OrderDetail reads one order. includeFields selects the optional sections;
// with none given 1688 returns GuaranteesTerms, NativeLogistics and
// OrderInvoice, which is what the tracking pipeline needs.
//
// This is the family B endpoint where success arrives as the STRING "true".
func (c *Client) OrderDetail(ctx context.Context, orderID ID, includeFields ...string) (TradeInfo, error) {
	if orderID == 0 {
		return TradeInfo{}, errors.New("ali: trade.get.buyerView: orderId is required")
	}
	p := Params{"webSite": WebSite1688, "orderId": orderID}
	if len(includeFields) > 0 {
		p["includeFields"] = strings.Join(includeFields, ",")
	}
	body, err := c.Call(ctx, OrderBuyerView, p)
	if err != nil {
		return TradeInfo{}, err
	}
	return DecodeB[TradeInfo](OrderBuyerView.Key(), body)
}

// ---------------------------------------------------------- getBuyerOrderList

// OrderListQuery is the parameter set of alibaba.trade.getBuyerOrderList. Every
// field is optional; the polling job filters on ModifyStartTime, and the relay
// looks an order up again by OutOrderID.
type OrderListQuery struct {
	Page                     int
	PageSize                 int
	CreateStartTime          time.Time
	CreateEndTime            time.Time
	ModifyStartTime          time.Time
	ModifyEndTime            time.Time
	OrderStatus              string
	RefundStatus             string
	BizTypes                 []string
	SellerMemberID           string
	SellerLoginID            string
	SellerRateStatus         int
	TradeType                string
	ProductName              string
	OutOrderID               string
	OrderIDs                 IDs
	IsHis                    bool
	NeedBuyerAddressAndPhone bool
	NeedMemoInfo             bool
	NeedInvoicingSetting     bool
}

// Params renders the query as gateway parameters. Times go out in the compact
// layout the order APIs use.
func (q OrderListQuery) Params() Params {
	p := Params{}
	if q.Page > 0 {
		p["page"] = q.Page
	}
	if q.PageSize > 0 {
		p["pageSize"] = q.PageSize
	}
	setTime(p, "createStartTime", q.CreateStartTime)
	setTime(p, "createEndTime", q.CreateEndTime)
	setTime(p, "modifyStartTime", q.ModifyStartTime)
	setTime(p, "modifyEndTime", q.ModifyEndTime)
	setStr(p, "orderStatus", q.OrderStatus)
	setStr(p, "refundStatus", q.RefundStatus)
	setStr(p, "sellerMemberId", q.SellerMemberID)
	setStr(p, "sellerLoginId", q.SellerLoginID)
	setStr(p, "tradeType", q.TradeType)
	setStr(p, "productName", q.ProductName)
	setStr(p, "outOrderId", q.OutOrderID)
	if len(q.BizTypes) > 0 {
		p["bizTypes"] = q.BizTypes
	}
	if len(q.OrderIDs) > 0 {
		p["orderIds"] = q.OrderIDs
	}
	if q.SellerRateStatus > 0 {
		p["sellerRateStatus"] = q.SellerRateStatus
	}
	if q.IsHis {
		p["isHis"] = true
	}
	if q.NeedBuyerAddressAndPhone {
		p["needBuyerAddressAndPhone"] = true
	}
	if q.NeedMemoInfo {
		p["needMemoInfo"] = true
	}
	if q.NeedInvoicingSetting {
		p["needInvoicingSetting"] = true
	}
	return p
}

// OrderListResult is the response of alibaba.trade.getBuyerOrderList. It has no
// success field at all: an empty errorCode is the only sign of success.
type OrderListResult struct {
	Orders       []TradeInfo `json:"result"`
	ErrorCode    string      `json:"errorCode"`
	ErrorMessage string      `json:"errorMessage"`
	TotalRecord  int64       `json:"totalRecord"`
}

// BuyerOrders lists the authorized buyer's orders.
func (c *Client) BuyerOrders(ctx context.Context, q OrderListQuery) (OrderListResult, error) {
	body, err := c.Call(ctx, BuyerOrderList, q.Params())
	if err != nil {
		return OrderListResult{}, err
	}
	var out OrderListResult
	if err := DecodeFlat(BuyerOrderList.Key(), body, &out); err != nil {
		return OrderListResult{}, err
	}
	if out.ErrorCode != "" {
		return out, &APIError{
			API:     BuyerOrderList.Key(),
			Code:    out.ErrorCode,
			Message: out.ErrorMessage,
			Body:    truncate(body),
		}
	}
	return out, nil
}

// ----------------------------------------------------------------- trade.cancel

// CancelResult is the response of alibaba.trade.cancel.
type CancelResult struct {
	Success      FlexBool `json:"success"`
	ErrorCode    string   `json:"errorCode"`
	ErrorMessage string   `json:"errorMessage"`
}

// CancelOrder closes an unpaid order. reason is one of CancelBuyerCancel,
// CancelSellerGoodsLack or CancelOther.
//
// The parameter is tradeID with a capital ID — this endpoint alone spells it
// that way — and cancelling within ten seconds of creation fails with
// CLOSE_ORDER_TOO_FAST.
func (c *Client) CancelOrder(ctx context.Context, orderID ID, reason, remark string) (CancelResult, error) {
	if orderID == 0 {
		return CancelResult{}, errors.New("ali: trade.cancel: tradeID is required")
	}
	if reason == "" {
		reason = CancelOther
	}
	p := Params{"webSite": WebSite1688, "tradeID": orderID, "cancelReason": reason}
	setStr(p, "remark", remark)

	body, err := c.Call(ctx, TradeCancel, p)
	if err != nil {
		return CancelResult{}, err
	}
	var out CancelResult
	if err := DecodeFlat(TradeCancel.Key(), body, &out); err != nil {
		return CancelResult{}, err
	}
	if !out.Success.Bool() {
		return out, &APIError{
			API:     TradeCancel.Key(),
			Code:    firstNonEmpty(out.ErrorCode, "CANCEL_FAILED"),
			Message: out.ErrorMessage,
			Body:    truncate(body),
		}
	}
	return out, nil
}

// ---------------------------------------------------------------------- helpers

// setStr adds a parameter only when it carries a value, so the signature is
// computed over exactly the fields that were meant to be sent.
func setStr(p Params, key, val string) {
	if val != "" {
		p[key] = val
	}
}

// setTime adds a time parameter in the compact layout, skipping the zero time.
func setTime(p Params, key string, t time.Time) {
	if !t.IsZero() {
		p[key] = FormatCompact(t)
	}
}

// unknownFields returns the members of b that are not json tags of v, which is
// how the order model keeps the six hundred fields it does not model.
func unknownFields(b []byte, v any) map[string]json.RawMessage {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(b, &raw); err != nil || len(raw) == 0 {
		return nil
	}
	t := reflect.TypeOf(v)
	for t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return nil
	}
	for i := 0; i < t.NumField(); i++ {
		name, _, _ := strings.Cut(t.Field(i).Tag.Get("json"), ",")
		if name == "" {
			name = t.Field(i).Name
		}
		delete(raw, name)
	}
	if len(raw) == 0 {
		return nil
	}
	return raw
}
