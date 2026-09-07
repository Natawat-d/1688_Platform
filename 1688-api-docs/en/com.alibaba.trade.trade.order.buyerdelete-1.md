# Buyer deletes a closed order

Original name: 买家删除已关闭的订单  
API: `com.alibaba.trade:trade.order.buyerdelete:1` · Category: Orders  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:trade.order.buyerdelete-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/trade.order.buyerdelete/{appKey}`  
Requires user authorization (access_token) · Requires signature

Buyer deletes an order that has been closed.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `orderId` | java.lang.Long | yes | Order id | 12213412341 |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.openapi.shared.common.trade.order.buyerdelete.ResultModel](#m-alibaba-openapi-shared-common-trade-order-buyerdelete-resultmodel) | yes | Return value | {} |

<a id="m-alibaba-openapi-shared-common-trade-order-buyerdelete-resultmodel"></a>
#### alibaba.openapi.shared.common.trade.order.buyerdelete.ResultModel

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `success` | java.lang.Boolean | yes | Whether successful | true |
| `code` | java.lang.String | yes | Error code | "" |
| `message` | java.lang.String | yes | Error message | "" |
| `result` | java.lang.Boolean | yes | Whether the business operation was successful | true |
