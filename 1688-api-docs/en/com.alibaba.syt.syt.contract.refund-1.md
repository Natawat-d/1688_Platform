# 88 ShengYiTong: contract refund request

Original name: 88生意通合同退款申请  
API: `com.alibaba.syt:syt.contract.refund:1` · Category: 88 ShengYiTong (Business Link)  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.syt:syt.contract.refund-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.syt/syt.contract.refund/{appKey}`  
Requires user authorization (access_token) · Requires signature

88 ShengYiTong: apply for a contract refund.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `request` | [message:alibaba.china.share.creditpay.gateway.api.request.ContractRefundApplyApiRequest](#m-alibaba-china-share-creditpay-gateway-api-request-contractrefundapplyapirequest) | yes | Refund application parameters | {   "draftNo": "DRAFT20251230001",   "buyerId": "buyer123456",   "amount": 100000,   "requestNo": "REQ20251230001",   "payOrderNo": "PAY20251230001",   "refundReason": "商品质量问题，申请退款" } |

<a id="m-alibaba-china-share-creditpay-gateway-api-request-contractrefundapplyapirequest"></a>
#### alibaba.china.share.creditpay.gateway.api.request.ContractRefundApplyApiRequest

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `draftNo` | java.lang.String | yes | Contract number | 88SYT20251201169011 |
| `amount` | BigDecimal | yes | Refund amount, in yuan, currency RMB (CNY) | 1 |
| `requestNo` | java.lang.String | yes | Refund request idempotency key, identifies a unique request | requestNo123abc |
| `payOrderNo` | java.lang.String | yes | Payment order number, from the payment interface | 8899x |
| `refundReason` | java.lang.String | yes | Refund application reason | 已经线下沟通退款10元 |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.china.share.creditpay.gateway.api.response.ContractRefundApplyApiResponse](#m-alibaba-china-share-creditpay-gateway-api-response-contractrefundapplyapiresponse) | yes | Response | {   "success": true,   "code": "200",   "message": "退款申请成功",   "refundOrderNo": "REFUND20251230001" } |

<a id="m-alibaba-china-share-creditpay-gateway-api-response-contractrefundapplyapiresponse"></a>
#### alibaba.china.share.creditpay.gateway.api.response.ContractRefundApplyApiResponse

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `refundOrderNo` | java.lang.String | yes | Refund order number | 889900 |
| `success` | Boolean | yes | Whether the call succeeded, true for success, false for failure |  true |
| `errorMsg` | String | yes | Call description | 调用成功 |
| `errorCode` | String | yes | Call error code | SUCCESS |
| `traceId` | String | yes | Call trace ID, for troubleshooting | 123 |
