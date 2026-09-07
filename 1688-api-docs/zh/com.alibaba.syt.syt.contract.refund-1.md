# 88生意通合同退款申请

API: `com.alibaba.syt:syt.contract.refund:1` · Category: 88生意通  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.syt:syt.contract.refund-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.syt/syt.contract.refund/{appKey}`  
需要授权 (access_token) · 需要签名

88生意通合同退款申请

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `request` | [message:alibaba.china.share.creditpay.gateway.api.request.ContractRefundApplyApiRequest](#m-alibaba-china-share-creditpay-gateway-api-request-contractrefundapplyapirequest) | 是 | 退款申请参数 | {   "draftNo": "DRAFT20251230001",   "buyerId": "buyer123456",   "amount": 100000,   "requestNo": "REQ20251230001",   "payOrderNo": "PAY20251230001",   "refundReason": "商品质量问题，申请退款" } |

<a id="m-alibaba-china-share-creditpay-gateway-api-request-contractrefundapplyapirequest"></a>
#### alibaba.china.share.creditpay.gateway.api.request.ContractRefundApplyApiRequest

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `draftNo` | java.lang.String | 是 | 合同号 | 88SYT20251201169011 |
| `amount` | BigDecimal | 是 | 退款金额，单位元，币种人民币 CNY | 1 |
| `requestNo` | java.lang.String | 是 | 退款请求幂等号，标识唯一一次请求 | requestNo123abc |
| `payOrderNo` | java.lang.String | 是 | 支付单号，来自支付接口 | 8899x |
| `refundReason` | java.lang.String | 是 | 退款申请原因 | 已经线下沟通退款10元 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:alibaba.china.share.creditpay.gateway.api.response.ContractRefundApplyApiResponse](#m-alibaba-china-share-creditpay-gateway-api-response-contractrefundapplyapiresponse) | 是 | 响应 | {   "success": true,   "code": "200",   "message": "退款申请成功",   "refundOrderNo": "REFUND20251230001" } |

<a id="m-alibaba-china-share-creditpay-gateway-api-response-contractrefundapplyapiresponse"></a>
#### alibaba.china.share.creditpay.gateway.api.response.ContractRefundApplyApiResponse

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `refundOrderNo` | java.lang.String | 是 | 退款单号 | 889900 |
| `success` | Boolean | 是 | 调用是否成功，true 成功，false 失败 |  true |
| `errorMsg` | String | 是 | 调用描述 | 调用成功 |
| `errorCode` | String | 是 | 调用错误码 | SUCCESS |
| `traceId` | String | 是 | 调用链路 ID，排查问题 | 123 |
