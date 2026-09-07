# Query maximum refundable amount when applying

Original name: 申请退款时查询最大可退费用  
API: `com.alibaba.trade:alibaba.trade.getMaxRefundFee:1` · Category: Returns & Refunds  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.trade.getMaxRefundFee-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.trade.getMaxRefundFee/{appKey}`  
Requires user authorization (access_token) · Requires signature

Query the maximum refundable amount when applying for a refund.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `input` | [message:alibaba.ocean.openplatform.biz.trade.param.OpMaxRefundFeeGetParam](#m-alibaba-ocean-openplatform-biz-trade-param-opmaxrefundfeegetparam) | yes |  |   |

<a id="m-alibaba-ocean-openplatform-biz-trade-param-opmaxrefundfeegetparam"></a>
#### alibaba.ocean.openplatform.biz.trade.param.OpMaxRefundFeeGetParam

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `goodsStatus` | java.lang.String | yes | Cargo status |  售中等待卖家发货:"refundWaitSellerSend"; 售中等待买家收货:"refundWaitBuyerReceive"; 售中已收货（未确认完成交易）:"refundBuyerReceived" 售后未收货:"aftersaleBuyerNotReceived"; 售后已收到货:"aftersaleBuyerReceived" |
| `orderId` | java.lang.Long | yes | Order ID | 123 |
| `refundId` | java.lang.String | no | The refund order must be in refunding status; this can be omitted | TQ123 |
| `refundGoodsCountList` | [message:alibaba.ocean.openplatform.biz.trade.common.model.OrderEntryCountModel[]](#m-alibaba-ocean-openplatform-biz-trade-common-model-orderentrycountmodel[]) | yes | Return quantity | [{1: 1}] |

<a id="m-alibaba-ocean-openplatform-biz-trade-common-model-orderentrycountmodel[]"></a>
#### alibaba.ocean.openplatform.biz.trade.common.model.OrderEntryCountModel[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `id` | Long | yes | Sub-order ID | 1 |
| `count` | Integer | yes | Quantity of products purchased in the sub-order | 1 |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.ocean.openplatform.common.OpMaxRefundFeeResultModel](#m-alibaba-ocean-openplatform-common-opmaxrefundfeeresultmodel) | yes |  |   |

<a id="m-alibaba-ocean-openplatform-common-opmaxrefundfeeresultmodel"></a>
#### alibaba.ocean.openplatform.common.OpMaxRefundFeeResultModel

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `code` | java.lang.String | yes |  |   |
| `message` | java.lang.String | yes |  |   |
| `result` | [message:alibaba.ocean.openplatform.biz.trade.common.model.OpMaxRefundFeeModel](#m-alibaba-ocean-openplatform-biz-trade-common-model-opmaxrefundfeemodel) | yes |  |   |
| `retCodes` | java.lang.String[] | yes |  |   |
| `subCode` | java.lang.String | yes |  |   |
| `subMessage` | java.lang.String | yes |  |   |
| `success` | boolean | yes |  |   |

<a id="m-alibaba-ocean-openplatform-biz-trade-common-model-opmaxrefundfeemodel"></a>
#### alibaba.ocean.openplatform.biz.trade.common.model.OpMaxRefundFeeModel

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `maxGoodsFee` | java.lang.Long | yes | Maximum refundable product amount, in cents | 7 |
| `maxPostFee` | java.lang.Long | yes | Maximum refundable shipping fee, in cents | 2 |
| `maxRefundFee` | java.lang.Long | yes | Maximum refundable amount, i.e. the upper limit of freight + product amount, unit: cents (fen) | 9 |
| `maxSubstituteFetchFee` | java.lang.Long | yes | Maximum refundable official pickup-on-behalf service fee, unit: cents (fen) | 0 |
