# Create cross-border order

Original name: 跨境订单创建  
API: `com.alibaba.trade:alibaba.trade.createCrossOrder:1` · Category: Orders  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.trade.createCrossOrder-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.trade.createCrossOrder/{appKey}`  
Requires user authorization (access_token) · Requires signature

Cross-border-only order creation. An order may contain at most 50 SKUs, all from the same supplier. For multiple suppliers or more than 50 SKUs, split the order yourself before submitting. In some special cases several orders are created at once and several order numbers are returned. Supports both the open-marketplace and distribution scenarios. Orders are placed under the main account or a sub-account depending on the currently authorized user.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `flow` | String | yes | general (create wholesale order), fenxiao (create distribution order), the saleproxy process will verify the distribution relationship, paired (Tiantian Texmai / Daily Deals), repurchase (repurchase contract order), boutiquefenxiao (curated supply distribution-price order; used when purchase quantity is 1, free shipping applies), boutiquepifa (curated supply wholesale-price order; used when purchase quantity is greater than 2). Different order-placing flow channels may cause differences in price and purchase protection services. | general |
| `message` | String | no | Buyer message | 留言 |
| `isvBizType` | String | no | Open platform business code, default is cross, representing the 1688 business side being integrated with. cross (cross-border business), cross_daigou (cross-border purchasing-agent business), cross_distribution (overseas distribution business); you can contact your business manager (Xiaoer) to confirm — this field does not affect order placement | cross |
| `addressParam` | [message:alibaba.trade.fast.address](#m-alibaba-trade-fast-address) | yes | Shipping address info | {"address":"网商路699号","phone": "0517-88990077","mobile": "15251667788","fullName": "张三","postCode": "000000","areaText": "滨江区","townText": "","cityText": "杭州市","provinceText": "浙江省"} |
| `cargoParamList` | [message:alibaba.trade.fast.cargo[]](#m-alibaba-trade-fast-cargo[]) | yes | Product information | [{"specId": "b266e0726506185beaf205cbae88530d","quantity": 5,"offerId": 554456348334},{"specId": "2ba3d63866a71fbae83909d9b4814f01","quantity": 6,"offerId": 554456348334}] |
| `invoiceParam` | [message:alibaba.trade.fast.invoice](#m-alibaba-trade-fast-invoice) | no | Invoice information | {"invoiceType":0,"cityText": "杭州市","provinceText": "浙江省","address": "网商路699号","phone": "0517-88990077","mobile": "15251667788","fullName": "张五","postCode": "000000","areaText": "滨江区","companyName": "测试公司","taxpayerIdentifier": "123455"} |
| `tradeType` | String | no | Since different products support different trade modes and no single trade mode applies universally, the trade mode usable for the current order must be obtained from tradeModeNameList in the order preview API. Trade mode type descriptions: assureTrade (Trade 4.0 universal escrow transaction), alipay (Alipay escrow transaction used generally in the open marketplace; currently being phased out and will be removed later), period (regular account period transaction), assure (the escrow transaction process required when a large buyer enterprise places an order via procurement inquiry/quote), creditBuy (Cheng-e-She), bank (bank transfer), 631staged (631 staged payment), 37staged (37 staged payment). If this field is not passed, the system defaults to selecting an available trade mode for the order; if Cheng-e-She is enabled the default is creditBuy (Cheng-e-She), otherwise the default is Alipay escrow transaction. | assureTrade |
| `shopPromotionId` | String | no | Store discount ID, obtained via the "preview data before order creation" API. If empty, the default discount is used | itemCoupon-5600812521_31032085284-398517001570 |
| `anonymousBuyer` | Boolean | no | Whether the order is anonymous |   |
| `fenxiaoChannel` | String | no | Downstream platform for reflow orders: Taobao-thyny, Tmall-tm, Taote-taote, Alibaba C2M-c2m, JD-jingdong, Pinduoduo-pinduoduo, WeChat-weixin, Cross-border-kuajing, Kuaishou-kuaishou, Youzan-youzan, Douyin-douyin, Siku-siku, Meituan Tuanhaohuo-meituan, Xiaohongshu-xiaohongshu, Dangdang-dangdang, Suning-suning, DaVdian-davdian, Xingyun-xingyun, Miya-miya, Boluopai Mall-boluo, Kuaituantuan-kuaituantuan, Other-other | douyin |
| `inventoryMode` | String | no | Stock mode, JIT (JIT mode) or NORMAL (warehouse fulfillment mode). Currently only available for AE. | JIT |
| `outOrderId` | String | no | External order number | 988129883123 |
| `pickupService` | String | no | Door-to-door pickup, currently available for AE supply, not yet enabled for other scenarios. y or n, default is n | n |
| `warehouseCode` | String | no | Door-to-door pickup warehouse code | any |
| `preSelectPayChannel` | String | no | Pre-selected payment channel, used for financial order routing. Returned by the order information query API as result.exAttributes.preSelectPayChannel; this value is the pre-selected payment channel flag passed in when creating the order via the create-order API. | alipay |
| `smallProcurement` | String | no | Whether it is a small-amount purchase. Currently available for AE supply. Value: y/n, default is n. | y |
| `useRedEnvelope` | String | no | Whether to use a red packet (coupon): y - use, n - do not use. Uses red packet by default | n |
| `dropshipping` | String | no | Whether it is a transshipment order. Value: y/n, default is n. | y |
| `addedService` | String | no | Value-added service, toB / toC | toB |
| `crossBorderLogisticsSolutionId` | String | no | The sourceId of the official cross-border logistics solution selected by the user | GLOBAL_CAINIAO_VN_TRANSIT_LAND |
| `useCrossBorderLogisticsSolution` | Boolean | no | Whether to use the cross-border logistics solution preview. Defaults to false. This parameter does not take effect for overseas addresses. It takes effect for Hong Kong, Macao, and Taiwan addresses. | true |
| `extendParam` | [message:com.alibaba.ocean.openplatform.biz.trade.param.ExtendParam](#m-com-alibaba-ocean-openplatform-biz-trade-param-extendparam) | no | Order placement extension information | {} |
| `useOfficialSolution` | Boolean | no | Delivery method uses official logistics pickup | false |
| `useOfficialSolutionModelList` | [message:com.alibaba.ocean.openplatform.biz.trade.common.model.UseOfficialSolutionModel[]](#m-com-alibaba-ocean-openplatform-biz-trade-common-model-useofficialsolutionmodel[]) | no | Official logistics pickup solution (required when useOfficialSolution is true) | [] |
| `fromAgent` | String | no | External agent | qoderwork |

<a id="m-alibaba-trade-fast-address"></a>
#### alibaba.trade.fast.address

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `addressId` | Long | yes | Shipping address ID | 1234 |
| `fullName` | String | yes | Recipient name | 张三 |
| `mobile` | String | yes | Mobile phone | 15251667788 |
| `phone` | String | yes | Phone number | 0517-88990077  注意: 只能填电话号码,不能加其他中文 |
| `postCode` | String | yes | Postal code | 000000 |
| `cityText` | String | yes | City text | 杭州市 |
| `provinceText` | String | yes | Province text | 浙江省 |
| `areaText` | String | yes | District text | 滨江区 |
| `townText` | String | yes | Town text | 长河镇 |
| `address` | String | yes | Street address | 网商路699号 |
| `districtCode` | String | yes | Address code | 310107 |
| `addressCode` | String | yes | Address code | 地址码 |

<a id="m-alibaba-trade-fast-cargo[]"></a>
#### alibaba.trade.fast.cargo[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `offerId` | Long | yes | offer id corresponding to the product | 554456348334 |
| `specId` | String | yes | Product SKU ID | b266e0726506185beaf205cbae88530d |
| `quantity` | Double | yes | Product quantity (used for amount calculation) | 5 |
| `openOfferId` | String | no | Encrypted offerId. When the search result returns only openOfferId, use openOfferId in place of offerId to place the order. | Wcv2w970KoL1BCpJGQp3vwKvcBThOPlpHnUtkK3T0n4= |
| `outMemberId` | String | no | External downstream member ID | 98928912-23 |

<a id="m-alibaba-trade-fast-invoice"></a>
#### alibaba.trade.fast.invoice

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `invoiceType` | Integer | yes | Invoice type<br>0: regular invoice, 1: VAT invoice | 0 |
| `provinceText` | String | yes | Province text | 浙江省 |
| `cityText` | String | yes | City text | 杭州市 |
| `areaText` | String | yes | Region text | 滨江区 |
| `townText` | String | yes | Town text | 长河镇 |
| `postCode` | String | yes | Postal code | 333333 |
| `address` | String | yes | Street | 网商路699号 |
| `fullName` | String | yes | Invoice recipient's name | 张三 |
| `phone` | String | yes | Phone number | 0517-88990077 |
| `mobile` | String | yes | Mobile phone | 15251667788 |
| `companyName` | String | yes | Purchasing company name (invoice title) | 测试公司 |
| `taxpayerIdentifier` | String | yes | Tax identification number | 12345 |
| `bankAndAccount` | String | yes | Bank name and account number | 网商银行 |
| `localInvoiceId` | String | yes | VAT local invoice number | 123123123 |

<a id="m-com-alibaba-ocean-openplatform-biz-trade-param-extendparam"></a>
#### com.alibaba.ocean.openplatform.biz.trade.param.ExtendParam

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `crossBorderLogisticsSelfPickupAddress` | [message:com.alibaba.ocean.openplatform.biz.trade.param.CrossBorderLogisticsSelfPickupAddress](#m-com-alibaba-ocean-openplatform-biz-trade-param-crossborderlogisticsselfpickupaddress) | yes | Cross-border consolidation warehouse pickup point address and contact person | {} |
| `extPairList` | [message:com.alibaba.ocean.openplatform.biz.trade.common.model.ExtrPair[]](#m-com-alibaba-ocean-openplatform-biz-trade-common-model-extrpair[]) | no | List of extended objects | [] |
| `offerIsvCargoFromList` | [message:alibaba.ocean.openplatform.common.KeyValuePair[]](#m-alibaba-ocean-openplatform-common-keyvaluepair[]) | no | Lightweight app closed-loop order placement, offer source info list | [] |
| `selectedTradeServiceList` | [message:com.alibaba.ocean.openplatform.biz.trade.param.TradeService[]](#m-com-alibaba-ocean-openplatform-biz-trade-param-tradeservice[]) | no | Trade service parameters | [] |

<a id="m-com-alibaba-ocean-openplatform-biz-trade-param-crossborderlogisticsselfpickupaddress"></a>
#### com.alibaba.ocean.openplatform.biz.trade.param.CrossBorderLogisticsSelfPickupAddress

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `addressText` | String | yes | Consolidation warehouse self-pickup address; province/city/district do not need to be filled in — they follow the default solution's consolidation warehouse | 网商路699号 |
| `mobile` | String | yes | Consolidation warehouse self-pickup mobile number | 18792983782 |
| `contractName` | String | yes | Consolidation warehouse self-pickup contact person | 无名 |

<a id="m-com-alibaba-ocean-openplatform-biz-trade-common-model-extrpair[]"></a>
#### com.alibaba.ocean.openplatform.biz.trade.common.model.ExtrPair[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `key` | String | yes | Key | doorPickup |
| `value` | String | yes | Value | {       "serviceCode": "mkd",       "price": 1200, // 人民币，分     } |
| `desc` | String | yes | Description | 是否使用上门揽 |

<a id="m-alibaba-ocean-openplatform-common-keyvaluepair[]"></a>
#### alibaba.ocean.openplatform.common.KeyValuePair[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `key` | java.lang.String | yes | offerId |    |
| `value` | java.lang.String | yes | Offer source information |    |

<a id="m-com-alibaba-ocean-openplatform-biz-trade-param-tradeservice[]"></a>
#### com.alibaba.ocean.openplatform.biz.trade.param.TradeService[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `code` | String | yes | Trade service code | vas |
| `sourceId` | String | yes | Trade service sourceId | 3C |

<a id="m-com-alibaba-ocean-openplatform-biz-trade-common-model-useofficialsolutionmodel[]"></a>
#### com.alibaba.ocean.openplatform.biz.trade.common.model.UseOfficialSolutionModel[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `orderGroup` | String | yes | Order group | 123 |
| `useOfficialSolutionCode` | String | yes | Official logistics pickup solution code | 123 |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.trade.cross.result](#m-alibaba-trade-cross-result) | yes | Order creation result | {} |
| `success` | Boolean | yes | Whether successful | true |
| `code` | String | yes | Error code | 400 |
| `message` | String | yes | Error information | AddressId invalid:AddressId invalid |

<a id="m-alibaba-trade-cross-result"></a>
#### alibaba.trade.cross.result

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `totalSuccessAmount` | Long | yes | Total order amount (in cents). This field is empty when multiple orders are created at once. | 100 |
| `orderId` | String | yes | Order ID; this field is empty when multiple orders are created at once | 111111111 |
| `success` | Boolean | yes | Whether successful | true |
| `code` | String | yes | Error code |  |
| `message` | String | yes | Error message |  |
| `accountPeriod` | [message:alibaba.trade.cross.period](#m-alibaba-trade-cross-period) | yes | Account period information; returns empty for orders not paid via account period |  |
| `failedOfferList` | [message:alibaba.trade.fast.offer[]](#m-alibaba-trade-fast-offer[]) | yes | Failed product information |  |
| `postFee` | Long | yes | Freight, unit: cents (fen); when creating multiple orders at once, this field is empty |  |
| `orderList` | [message:alibaba.tradeResult.BizSimpleOrder[]](#m-alibaba-traderesult-bizsimpleorder[]) | yes | Create multiple orders at once |  |

<a id="m-alibaba-trade-cross-period"></a>
#### alibaba.trade.cross.period

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `tapType` | Integer | yes | Account period type: 1: settled once on a specified date each month, 3: settled once on a specified date every two months, 6: settled once on a specified date every three months, 5: settled based on receipt time and account period date | 1 |
| `tapDate` | Integer | yes | Varies according to the account period type. For monthly settlement type, this value represents a specific day; for settlement based on receipt time, this value represents the settlement time period | 12 |
| `tapOverdue` | Integer | yes | Number of overdue occurrences | 0 |

<a id="m-alibaba-trade-fast-offer[]"></a>
#### alibaba.trade.fast.offer[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `offerId` | String | yes | Product for which order placement failed | 554456348334 |
| `specId` | String | yes | SKU ID of the product that failed to be ordered | b266e0726506185beaf205cbae88530d |
| `errorCode` | String | yes | Error code for order placement failure |  |
| `errorMessage` | String | yes | Error description for order placement failure |  |

<a id="m-alibaba-traderesult-bizsimpleorder[]"></a>
#### alibaba.tradeResult.BizSimpleOrder[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `postFee` | java.lang.Long | yes | Shipping fee | 10 |
| `orderAmmount` | java.lang.Long | yes | Actual amount paid for the order, in cents | 100 |
| `message` | java.lang.String | yes | Description information | null |
| `resultCode` | java.lang.String | yes | Return code | 200 |
| `success` | java.lang.Boolean | yes | Whether successful | true |
| `orderId` | java.lang.String | yes | Order number | 123124124 |
| `payChannel` | String | yes | Payment channel | shegou |

## Error codes

| Error | Symptom | How to fix |
|---|---|---|
| not support tradeType:【XXXX】 | Unsupported transaction type | The trade type currently supported for this transaction; must be obtained via the tradeModeNameList field from the preview API |

## Samples

**Example of order return for Alipay payment mode**

```
{"totalSuccessAmount":156800,"orderId":"87407346014789305","success":false}
```

**Example of order return for account period payment**

```
{"totalSuccessAmount":156800,"orderId":"87407346015789305","success":false,"accountPeriod":{"tapType":5,"tapDate":360,"tapOverdue":1}}
```

**Shipping address description**

```
1、使用保存的收货地址传参示例：{"addressId":593861699}，其中addressId是调用“买家获取保存的收货地址信息列表”接口获取；
2、使用地址编码传参示例：{"address":"桃浦镇 金达路888号贸易8楼","phone": "0517-88990077","mobile": "15251667788","fullName": "张三","postCode": "000000","districtCode": "310107"}，其中districtCode需要调用“根据地址解析地区码”接口获取；
3、直接使用文本地址示例：{"address":"网商路699号","phone": "0517-88990077","mobile": "15251667788","fullName": "张三","postCode": "000000","areaText": "滨江区","townText": "","cityText": "杭州市","provinceText": "浙江省"}，省市区要传文本；
4、优先级说明，如果同时传入以上参数，系统按从1至3的优先级获取地址，满足1的条件下不会使用上示2中的参数，满足2的条件下不会使用上示3中的参数；
```

**Return value when creating multiple orders at the same time**

```
{
    "result":{
        "success":true,
        "orderList":[
            {
                "postFee":48000,
                "orderAmmount":4848000,
                "discount":0,
                "sumPaymentNoCarriageFromClient":4800000,
                "mergePay":false,
                "orderId":"105581756010628640",
                "chooseFreeFreight":false
            },
            {
                "postFee":50000,
                "orderAmmount":5050000,
                "discount":0,
                "sumPaymentNoCarriageFromClient":5000000,
                "mergePay":false,
                "orderId":"105637283010628640",
                "chooseFreeFreight":false
            }
        ]
    },
    "success":true
}
```
