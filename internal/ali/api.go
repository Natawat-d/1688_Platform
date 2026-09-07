package ali

import "strconv"

// API identifies one 1688 endpoint. The flags mirror needAuth / needSignature /
// needTimestamp in 1688-api-docs/ALL-APIS.json, and allapis_test.go asserts they
// still agree with the documentation.
type API struct {
	Namespace string
	Name      string
	Version   int
	NeedAuth  bool
	NeedSig   bool
	NeedTS    bool
}

// Key is the "namespace:name:version" form used by the docs and by our logs.
func (a API) Key() string {
	return a.Namespace + ":" + a.Name + ":" + strconv.Itoa(a.Version)
}

// DocID is the "namespace:name-version" form used in doc filenames and URLs.
func (a API) DocID() string {
	return a.Namespace + ":" + a.Name + "-" + strconv.Itoa(a.Version)
}

// Path is the request path, without a leading slash and without the app key.
func (a API) Path() string {
	return "openapi/param2/" + strconv.Itoa(a.Version) + "/" + a.Namespace + "/" + a.Name
}

// SignPath is the prefix of the signature input: the documented form starts at
// "param2/", not at "openapi/".
func (a API) SignPath(appKey string) string {
	return "param2/" + strconv.Itoa(a.Version) + "/" + a.Namespace + "/" + a.Name + "/" + appKey
}

// URL is the full endpoint for a given gateway base and app key.
func (a API) URL(base, appKey string) string {
	if len(base) > 0 && base[len(base)-1] == '/' {
		base = base[:len(base)-1]
	}
	return base + "/" + a.Path() + "/" + appKey
}

// The endpoints this build uses. Every one is documented needAuth and
// needSignature, except the message-replay pair, which needs only a signature.
var (
	AccountBasic = API{"com.alibaba.account", "alibaba.account.basic", 1, true, true, false}

	KeywordQuery    = API{"com.alibaba.fenxiao.crossborder", "product.search.keywordQuery", 1, true, true, false}
	KeywordSN       = API{"com.alibaba.fenxiao.crossborder", "product.search.keywordSNQuery", 1, true, true, false}
	ProductDetail   = API{"com.alibaba.fenxiao.crossborder", "product.search.queryProductDetail", 1, true, true, false}
	CategoryByID    = API{"com.alibaba.fenxiao.crossborder", "category.translation.getById", 1, true, true, false}
	FreightEstimate = API{"com.alibaba.fenxiao.crossborder", "product.freight.estimate", 1, true, true, false}

	OrderPreview     = API{"com.alibaba.trade", "alibaba.createOrder.preview", 1, true, true, false}
	CreateCrossOrder = API{"com.alibaba.trade", "alibaba.trade.createCrossOrder", 1, true, true, false}
	AlipayURLGet     = API{"com.alibaba.trade", "alibaba.alipay.url.get", 1, true, true, false}
	OrderBuyerView   = API{"com.alibaba.trade", "alibaba.trade.get.buyerView", 1, true, true, false}
	BuyerOrderList   = API{"com.alibaba.trade", "alibaba.trade.getBuyerOrderList", 1, true, true, false}
	TradeCancel      = API{"com.alibaba.trade", "alibaba.trade.cancel", 1, true, true, false}

	LogisticsTrace = API{"com.alibaba.logistics", "alibaba.trade.getLogisticsTraceInfo.buyerView", 1, true, true, false}

	PushCursorList = API{"cn.alibaba.open", "push.cursor.messageList", 1, false, true, false}
	PushQueryList  = API{"cn.alibaba.open", "push.query.messageList", 1, false, true, false}
	PushConfirm    = API{"cn.alibaba.open", "push.message.confirm", 1, false, true, false}
)

// All is every endpoint this build speaks, used by the contract tests and by the
// stub's dispatch table.
func All() []API {
	return []API{
		AccountBasic,
		KeywordQuery, KeywordSN, ProductDetail, CategoryByID, FreightEstimate,
		OrderPreview, CreateCrossOrder, AlipayURLGet, OrderBuyerView, BuyerOrderList, TradeCancel,
		LogisticsTrace,
		PushCursorList, PushQueryList, PushConfirm,
	}
}
