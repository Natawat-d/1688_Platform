# 88生意通确认交易完成

API: `com.alibaba.syt:syt.contract.confirm:1` · Category: 88生意通  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.syt:syt.contract.confirm-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.syt/syt.contract.confirm/{appKey}`  
需要授权 (access_token) · 需要签名

88生意通确认交易完成

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `request` | [message:alibaba.china.share.creditpay.gateway.api.request.ContractConfirmApiRequest](#m-alibaba-china-share-creditpay-gateway-api-request-contractconfirmapirequest) | 是 | 请求参数 | {"draftNo":"88SYT20251204069031"} |

<a id="m-alibaba-china-share-creditpay-gateway-api-request-contractconfirmapirequest"></a>
#### alibaba.china.share.creditpay.gateway.api.request.ContractConfirmApiRequest

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `draftNo` | java.lang.String | 是 | 合同号 | 88SYT20251204069031 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:alibaba.china.share.creditpay.gateway.api.response.ContractConfirmApiResponse](#m-alibaba-china-share-creditpay-gateway-api-response-contractconfirmapiresponse) | 是 | 返回响应 | {   "success": true,   "code": "200",   "message": "合同确认成功",   "draftNo": "DRAFT20251230001",   "contractCurrentStatus": "CONFIRMED" } |

<a id="m-alibaba-china-share-creditpay-gateway-api-response-contractconfirmapiresponse"></a>
#### alibaba.china.share.creditpay.gateway.api.response.ContractConfirmApiResponse

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `draftNo` | java.lang.String | 是 | 合同号 | 88SYT20251204069031 |
| `contractCurrentStatus` | java.lang.String | 是 | 当前最新合同状态 | CONFIRMED |
| `success` | Boolean | 是 | 调用结果 | true |
| `errorCode` | String | 是 | 错误 code | SUCCESS |
| `errorMsg` | String | 是 | 错误消息 | 调用成功 |
| `traceId` | String | 是 | 调用 traceId | 123333 |
