# 88生意通合同作废

API: `com.alibaba.syt:syt.contract.invalid:1` · Category: 88生意通  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.syt:syt.contract.invalid-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.syt/syt.contract.invalid/{appKey}`  
需要授权 (access_token) · 需要签名

88生意通合同作废

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `request` | [message:alibaba.china.share.creditpay.gateway.api.request.ContractInvalidApiRequest](#m-alibaba-china-share-creditpay-gateway-api-request-contractinvalidapirequest) | 是 | 请求参数 | {   "draftNo": "DRAFT20251230001" } |

<a id="m-alibaba-china-share-creditpay-gateway-api-request-contractinvalidapirequest"></a>
#### alibaba.china.share.creditpay.gateway.api.request.ContractInvalidApiRequest

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `draftNo` | java.lang.String | 是 | 合同号 | 88SYT20251201169011 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:alibaba.china.share.creditpay.gateway.api.response.ContractInvalidApiResponse](#m-alibaba-china-share-creditpay-gateway-api-response-contractinvalidapiresponse) | 是 | 响应参数 | {   "success": true,   "code": "200",   "message": "合同失效成功",   "draftNo": "DRAFT20251230001",   "contractCurrentStatus": "INVALID" } |

<a id="m-alibaba-china-share-creditpay-gateway-api-response-contractinvalidapiresponse"></a>
#### alibaba.china.share.creditpay.gateway.api.response.ContractInvalidApiResponse

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `draftNo` | java.lang.String | 是 | 合同号 | 88SYT20251201169011 |
| `contractCurrentStatus` | java.lang.String | 是 | 当前合同状态 | INVALID |
| `success` | Boolean | 是 | 调用结果 | true |
| `errorCode` | String | 是 | 错误信息 | SUCCESS |
| `errorMsg` | String | 是 | 错误描述 | 错误描述 |
| `traceId` | String | 是 | traceId，必须记录，排查问题提供给平台 | traceId |
