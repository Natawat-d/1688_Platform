# 88 ShengYiTong: buyer payment, get cashier URL

Original name:  88生意通 买家支付获取收银台 URL  
API: `com.alibaba.syt:syt.contract.pay:1` · Category: 88 ShengYiTong (Business Link)  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.syt:syt.contract.pay-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.syt/syt.contract.pay/{appKey}`  
Requires user authorization (access_token) · Requires signature

88 ShengYiTong solution.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `request` | [message:alibaba.china.share.creditpay.gateway.api.request.ContractPayApiRequest](#m-alibaba-china-share-creditpay-gateway-api-request-contractpayapirequest) | yes | Payment request parameters | {   "draftNo": "DRAFT20251230001",   "userId": "user123456",   "payChannel": "ALIPAY",   "amount": 100000 } |

<a id="m-alibaba-china-share-creditpay-gateway-api-request-contractpayapirequest"></a>
#### alibaba.china.share.creditpay.gateway.api.request.ContractPayApiRequest

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `draftNo` | java.lang.String | yes | Contract draft | 88SYT20251204069031 |
| `amount` | BigDecimal | yes | Payment amount, unit: yuan, currency: RMB (CNY) | 1 |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.china.share.creditpay.gateway.api.response.ContractPayApiResponse](#m-alibaba-china-share-creditpay-gateway-api-response-contractpayapiresponse) | yes | Response | {   "success": true,   "code": "200",   "message": "获取支付托管账号成功",   "escrowAccount": {     "accountNumber": "6222021234567890123",     "bankName": "中国工商银行",     "bankBranchCode": "102290000123"   },   "validPayTime": "2026-01-30 16:49:00",   "payOrderNo": "PAY20251230001" } |

<a id="m-alibaba-china-share-creditpay-gateway-api-response-contractpayapiresponse"></a>
#### alibaba.china.share.creditpay.gateway.api.response.ContractPayApiResponse

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `success` | java.lang.Boolean | yes | Whether successful | true |
| `errorCode` | java.lang.String | yes | Error code | SUCCESS |
| `errorMsg` | java.lang.String | yes | Error description | 成功 |
| `cashierUrl` | String | yes | Cashier URL | https://syt.1688.com/page/SYT/buyer-cashier?draftNo=88SYT20260x |
