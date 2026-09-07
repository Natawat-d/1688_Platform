# Query repurchase contract

Original name: 复购合约查询  
API: `com.alibaba.trade:repurchase.contract.get:1` · Category: Repurchase  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:repurchase.contract.get-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/repurchase.contract.get/{appKey}`  
Requires user authorization (access_token) · Requires signature

Query whether a product supports a repurchase contract. A repurchase contract is a repeat-purchase discount agreement signed between buyer and seller, under which the seller gives repurchasing users special prices and guarantees. Orders under a repurchase contract must be placed through a specific trade flow.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `param` | [message:alibaba.openplatform.biz.trade.param.TradeRepurchaseQueryParam](#m-alibaba-openplatform-biz-trade-param-traderepurchasequeryparam) | yes | Query parameter | {"offerIds": [ 867442475017]} |

<a id="m-alibaba-openplatform-biz-trade-param-traderepurchasequeryparam"></a>
#### alibaba.openplatform.biz.trade.param.TradeRepurchaseQueryParam

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `offerIds` | java.lang.Long[] | yes | List of product IDs | [867442475017] |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.openapi.shared.common.repurchase.contract.get.ResultModel](#m-alibaba-openapi-shared-common-repurchase-contract-get-resultmodel) | yes | Return value | {} |

<a id="m-alibaba-openapi-shared-common-repurchase-contract-get-resultmodel"></a>
#### alibaba.openapi.shared.common.repurchase.contract.get.ResultModel

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `success` | boolean | yes | Whether successful | true |
| `code` | java.lang.String | yes | Error code | "" |
| `message` | java.lang.String | yes | Error message | "" |
| `result` | [message:alibaba.openplatform.biz.trade.result.TradeRepurchaseQueryResult](#m-alibaba-openplatform-biz-trade-result-traderepurchasequeryresult) | yes | Repurchase contract model | {} |

<a id="m-alibaba-openplatform-biz-trade-result-traderepurchasequeryresult"></a>
#### alibaba.openplatform.biz.trade.result.TradeRepurchaseQueryResult

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `offerRepurchases` | [message:alibaba.openplatform.biz.trade.model.RepurchaseOfferModel[]](#m-alibaba-openplatform-biz-trade-model-repurchaseoffermodel[]) | yes | Product repurchase contract | {} |

<a id="m-alibaba-openplatform-biz-trade-model-repurchaseoffermodel[]"></a>
#### alibaba.openplatform.biz.trade.model.RepurchaseOfferModel[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `offerId` | java.lang.Long | yes | Product ID | 867442475017 |
| `contractId` | java.lang.Long | yes | Contract ID | 1379269006 |
| `freightFee` | java.lang.Long | yes | Shipping fee, cents | 0 |
| `freightDiscount` | java.lang.Double | yes | Freight discount | null |
| `freightStrategy` | java.lang.String | yes | Shipping fee policy | CALC_BY_OUT_SPECIFIED_FREIGHT |
| `freeAddressCode` | java.lang.String | yes | Free-shipping address code | 110101 |
| `includeTax` | boolean | yes | Whether tax is included | false |
| `invoiceTaxRate` | java.lang.String | yes | Invoice tax rate | 0.01 |
| `invoiceRequired` | boolean | yes | Whether an invoice is required | true |
| `deliveryDays` | java.lang.Integer | yes | Shipping timeliness, days | 2 |
| `deliveryDateIsDefault` | java.lang.Boolean | yes | Whether it is the default shipping time | false |
| `skuRepurchases` | [message:alibaba.openplatform.biz.trade.model.RepurchaseSkuModel[]](#m-alibaba-openplatform-biz-trade-model-repurchaseskumodel[]) | yes | SKU that supports repurchase contracts | {} |

<a id="m-alibaba-openplatform-biz-trade-model-repurchaseskumodel[]"></a>
#### alibaba.openplatform.biz.trade.model.RepurchaseSkuModel[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `skuId` | java.lang.Long | yes | skuid | 5861839338331 |
| `priceRanges` | [message:alibaba.openplatform.biz.trade.model.RepurchasePriceRangeModel[]](#m-alibaba-openplatform-biz-trade-model-repurchasepricerangemodel[]) | yes | Tiered price | {                 "price": 1,                 "end": null,                 "begin": 300               } |

<a id="m-alibaba-openplatform-biz-trade-model-repurchasepricerangemodel[]"></a>
#### alibaba.openplatform.biz.trade.model.RepurchasePriceRangeModel[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `begin` | java.lang.Integer | yes | Order quantity threshold | 1 |
| `end` | java.lang.Integer | yes | Order quantity upper limit | 200 |
| `price` | java.lang.Long | yes | Order price | 100 |
