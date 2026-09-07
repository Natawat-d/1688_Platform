# Buyer confirms receipt

Original name: 买家确认收货  
API: `com.alibaba.trade:trade.receivegoods.confirm:1` · Category: Orders  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:trade.receivegoods.confirm-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/trade.receivegoods.confirm/{appKey}`  
Requires user authorization (access_token) · Requires signature

Buyer confirms receipt of goods.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `orderId` | java.lang.Long | yes | Order ID | 56623232655125698 |
| `orderEntryIds` | java.lang.Long[] | yes | Sub-order ID | 562356635566365512 |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.ocean.openplatform.biz.trade.result.TradeConfirmReceiptResult](#m-alibaba-ocean-openplatform-biz-trade-result-tradeconfirmreceiptresult) | yes |  |  |

<a id="m-alibaba-ocean-openplatform-biz-trade-result-tradeconfirmreceiptresult"></a>
#### alibaba.ocean.openplatform.biz.trade.result.TradeConfirmReceiptResult

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `success` | boolean | yes |  |  |
| `errorInfo` | java.lang.String | yes |  |  |
| `errorCode` | java.lang.String | yes |  |  |
