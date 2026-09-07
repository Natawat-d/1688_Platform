#  88生意通 买家支付获取收银台 URL

API: `com.alibaba.syt:syt.contract.pay:1` · Category: 88生意通  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.syt:syt.contract.pay-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.syt/syt.contract.pay/{appKey}`  
需要授权 (access_token) · 需要签名

88生意通解决方案

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `request` | [message:alibaba.china.share.creditpay.gateway.api.request.ContractPayApiRequest](#m-alibaba-china-share-creditpay-gateway-api-request-contractpayapirequest) | 是 | 支付请求参数 | {   "draftNo": "DRAFT20251230001",   "userId": "user123456",   "payChannel": "ALIPAY",   "amount": 100000 } |

<a id="m-alibaba-china-share-creditpay-gateway-api-request-contractpayapirequest"></a>
#### alibaba.china.share.creditpay.gateway.api.request.ContractPayApiRequest

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `draftNo` | java.lang.String | 是 | 合同草稿 | 88SYT20251204069031 |
| `amount` | BigDecimal | 是 | 支付金额，单位元，币种人民币 CNY | 1 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:alibaba.china.share.creditpay.gateway.api.response.ContractPayApiResponse](#m-alibaba-china-share-creditpay-gateway-api-response-contractpayapiresponse) | 是 | 响应 | {   "success": true,   "code": "200",   "message": "获取支付托管账号成功",   "escrowAccount": {     "accountNumber": "6222021234567890123",     "bankName": "中国工商银行",     "bankBranchCode": "102290000123"   },   "validPayTime": "2026-01-30 16:49:00",   "payOrderNo": "PAY20251230001" } |

<a id="m-alibaba-china-share-creditpay-gateway-api-response-contractpayapiresponse"></a>
#### alibaba.china.share.creditpay.gateway.api.response.ContractPayApiResponse

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `success` | java.lang.Boolean | 是 | 是否成功 | true |
| `errorCode` | java.lang.String | 是 | 错误码 | SUCCESS |
| `errorMsg` | java.lang.String | 是 | 错误描述 | 成功 |
| `cashierUrl` | String | 是 | 收银台 url | https://syt.1688.com/page/SYT/buyer-cashier?draftNo=88SYT20260x |
