# Sync downstream sales orders

Original name: 下游销售订单同步  
API: `com.alibaba.fenxiao.crossborder:trade.cross.orderSync:1` · Category: Data Write-back  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:trade.cross.orderSync-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/trade.cross.orderSync/{appKey}`  
Requires user authorization (access_token) · Requires signature

Sync downstream sales orders.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `orderParam` | [message:com.alibaba.cbu.order.param.OrderParamV](#m-com-alibaba-cbu-order-param-orderparamv) | yes | Order details |  |

<a id="m-com-alibaba-cbu-order-param-orderparamv"></a>
#### com.alibaba.cbu.order.param.OrderParamV

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `orderId` | String | yes | Order id | 1234 |
| `productId` | String | no | Product ID; required for sub-orders or combined main-sub orders | 12 |
| `productName` | String | no | Product name; required for sub-orders or merged main/sub orders | 手套 |
| `skuId` | String | no | skuId; required for sub-orders or merged main/sub orders | adsds1213 |
| `skuName` | String | no | SKU name; required for sub-orders or combined main-sub orders | 绿色 |
| `buyAmount` | Long | no | Purchase quantity | 2 |
| `createTime` | String | yes | Creation time | 2023-10-01 10:10:10 |
| `outMemberId` | String | yes | Downstream user ID | 1212113 |
| `payTime` | String | yes | Payment time | 2023-10-01 10:10:10 |
| `endTime` | String | no | Completion time | 2023-10-01 10:10:10 |
| `paidFee` | Long | yes | Actual payment amount, in cents | 200 |
| `refundStatus` | String | no | Refund status 0 - not refunded (refund closed, not applied) 1 - refunded | 0 |
| `status` | String | yes | payed - paid, success - transaction successful, close - transaction closed | success |
| `subOrderParamList` | [message:com.alibaba.cbu.order.param.SubOrderParam[]](#m-com-alibaba-cbu-order-param-suborderparam[]) | no | For one main order with multiple sub-orders, sub-order info must be passed | 子单信息 |

<a id="m-com-alibaba-cbu-order-param-suborderparam[]"></a>
#### com.alibaba.cbu.order.param.SubOrderParam[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `orderId` | String | yes | Sub-order ID | 2 |
| `productId` | String | yes | 1688 product ID, must be in plaintext | 1 |
| `productName` | String | yes | 1688 product name, must be in plaintext | 袜子 |
| `skuId` | String | yes | 1688 skuId must be plaintext | 1 |
| `skuName` | String | yes | 1688 SKU name, must be in plaintext | 绿色 |
| `buyAmount` | Long | yes | Purchase quantity of the sub-order | 1 |
| `paidFee` | Long | yes | Actual amount paid, in cents (fen) | 100 |
| `refundStatus` | String | yes | Refund status: 0 - not refunded or refund closed, 1 - refunded | 0 |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.openapi.shared.out.common.ResultModel](#m-alibaba-openapi-shared-out-common-resultmodel) | yes |  |  |

<a id="m-alibaba-openapi-shared-out-common-resultmodel"></a>
#### alibaba.openapi.shared.out.common.ResultModel

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `success` | boolean | yes |  |  |
| `code` | java.lang.String | yes |  |  |
| `message` | java.lang.String | yes |  |  |
