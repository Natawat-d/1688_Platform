# 复购合约查询

API: `com.alibaba.trade:repurchase.contract.get:1` · Category: 复购  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:repurchase.contract.get-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/repurchase.contract.get/{appKey}`  
需要授权 (access_token) · 需要签名

查询商品是否支持复购合约，复购合约是买卖家签署的复购优惠合约，卖家会给与复购用户特殊的价格、保障。复购合约需要通过特定的交易flow进行下单。

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `param` | [message:alibaba.openplatform.biz.trade.param.TradeRepurchaseQueryParam](#m-alibaba-openplatform-biz-trade-param-traderepurchasequeryparam) | 是 | 查询参数 | {"offerIds": [ 867442475017]} |

<a id="m-alibaba-openplatform-biz-trade-param-traderepurchasequeryparam"></a>
#### alibaba.openplatform.biz.trade.param.TradeRepurchaseQueryParam

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `offerIds` | java.lang.Long[] | 是 | 商品id列表 | [867442475017] |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:alibaba.openapi.shared.common.repurchase.contract.get.ResultModel](#m-alibaba-openapi-shared-common-repurchase-contract-get-resultmodel) | 是 | 返回值 | {} |

<a id="m-alibaba-openapi-shared-common-repurchase-contract-get-resultmodel"></a>
#### alibaba.openapi.shared.common.repurchase.contract.get.ResultModel

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `success` | boolean | 是 | 是否成功 | true |
| `code` | java.lang.String | 是 | 错误码 | "" |
| `message` | java.lang.String | 是 | 错误信息 | "" |
| `result` | [message:alibaba.openplatform.biz.trade.result.TradeRepurchaseQueryResult](#m-alibaba-openplatform-biz-trade-result-traderepurchasequeryresult) | 是 | 复购合约模型 | {} |

<a id="m-alibaba-openplatform-biz-trade-result-traderepurchasequeryresult"></a>
#### alibaba.openplatform.biz.trade.result.TradeRepurchaseQueryResult

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `offerRepurchases` | [message:alibaba.openplatform.biz.trade.model.RepurchaseOfferModel[]](#m-alibaba-openplatform-biz-trade-model-repurchaseoffermodel[]) | 是 | 商品复购合约 | {} |

<a id="m-alibaba-openplatform-biz-trade-model-repurchaseoffermodel[]"></a>
#### alibaba.openplatform.biz.trade.model.RepurchaseOfferModel[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `offerId` | java.lang.Long | 是 | 商品ID | 867442475017 |
| `contractId` | java.lang.Long | 是 | 合约ID | 1379269006 |
| `freightFee` | java.lang.Long | 是 | 运费 分 | 0 |
| `freightDiscount` | java.lang.Double | 是 | 运费折扣 | null |
| `freightStrategy` | java.lang.String | 是 | 运费策略 | CALC_BY_OUT_SPECIFIED_FREIGHT |
| `freeAddressCode` | java.lang.String | 是 | 免运费地址码 | 110101 |
| `includeTax` | boolean | 是 | 是否含税 | false |
| `invoiceTaxRate` | java.lang.String | 是 | 发票税率 | 0.01 |
| `invoiceRequired` | boolean | 是 | 是否需要发票 | true |
| `deliveryDays` | java.lang.Integer | 是 | 发货时效 天 | 2 |
| `deliveryDateIsDefault` | java.lang.Boolean | 是 | 是否默认发货时效 | false |
| `skuRepurchases` | [message:alibaba.openplatform.biz.trade.model.RepurchaseSkuModel[]](#m-alibaba-openplatform-biz-trade-model-repurchaseskumodel[]) | 是 | 支持复购合约的sku | {} |

<a id="m-alibaba-openplatform-biz-trade-model-repurchaseskumodel[]"></a>
#### alibaba.openplatform.biz.trade.model.RepurchaseSkuModel[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `skuId` | java.lang.Long | 是 | skuid | 5861839338331 |
| `priceRanges` | [message:alibaba.openplatform.biz.trade.model.RepurchasePriceRangeModel[]](#m-alibaba-openplatform-biz-trade-model-repurchasepricerangemodel[]) | 是 | 阶梯价 | {                 "price": 1,                 "end": null,                 "begin": 300               } |

<a id="m-alibaba-openplatform-biz-trade-model-repurchasepricerangemodel[]"></a>
#### alibaba.openplatform.biz.trade.model.RepurchasePriceRangeModel[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `begin` | java.lang.Integer | 是 | 订购量起 | 1 |
| `end` | java.lang.Integer | 是 | 订购量止 | 200 |
| `price` | java.lang.Long | 是 | 订购价 | 100 |
