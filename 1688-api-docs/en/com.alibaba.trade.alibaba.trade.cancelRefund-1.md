# Cancel refund/return request

Original name: 取消退款退货申请   
API: `com.alibaba.trade:alibaba.trade.cancelRefund:1` · Category: Returns & Refunds  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.trade.cancelRefund-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.trade.cancelRefund/{appKey}`  
No user authorization · Requires signature

Cancel a refund or return request.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | no | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `refundId` | String | yes | Refund order ID | TQ267395256051660259 |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.ocean.openplatform.biz.trade.result.RefundCancelResult](#m-alibaba-ocean-openplatform-biz-trade-result-refundcancelresult) | yes |  |  |

<a id="m-alibaba-ocean-openplatform-biz-trade-result-refundcancelresult"></a>
#### alibaba.ocean.openplatform.biz.trade.result.RefundCancelResult

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `success` | boolean | yes |  |  |
| `errorInfo` | java.lang.String | yes |  |  |
| `errorCode` | java.lang.String | yes |  |  |
