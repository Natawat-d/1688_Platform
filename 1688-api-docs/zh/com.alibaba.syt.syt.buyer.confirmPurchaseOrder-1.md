# 88生意通买家确认采购单

API: `com.alibaba.syt:syt.buyer.confirmPurchaseOrder:1` · Category: 88生意通  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.syt:syt.buyer.confirmPurchaseOrder-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.syt/syt.buyer.confirmPurchaseOrder/{appKey}`  
需要授权 (access_token) · 需要签名

88生意通买家确认采购单

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `request` | [message:alibaba.china.share.creditpay.gateway.api.request.SignConfirmApiRequest](#m-alibaba-china-share-creditpay-gateway-api-request-signconfirmapirequest) | 是 | 请求参数 | {   "draftNo": "CT2026051100001" } |

<a id="m-alibaba-china-share-creditpay-gateway-api-request-signconfirmapirequest"></a>
#### alibaba.china.share.creditpay.gateway.api.request.SignConfirmApiRequest

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `draftNo` | java.lang.String | 是 | 采购单编号 | 88SYT123456 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:alibaba.china.share.creditpay.gateway.api.response.SignConfirmApiResponse](#m-alibaba-china-share-creditpay-gateway-api-response-signconfirmapiresponse) | 是 | 响应结果 | {   "success": true,   "errorCode": null,   "errorMsg": null,   "errorNumCode": null,   "draftNo": "CT2026051100001"} |

<a id="m-alibaba-china-share-creditpay-gateway-api-response-signconfirmapiresponse"></a>
#### alibaba.china.share.creditpay.gateway.api.response.SignConfirmApiResponse

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `draftNo` | java.lang.String | 是 | 采购单编号 | 88SYT123456 |
| `errorCode` | java.lang.String | 是 | 调用结果 code，isSuccess=true 时关注 |  SUCCESS |
| `errorMsg` | java.lang.String | 是 | 调用结果描述，isSuccess=true 时关注 | 调用成功 |
| `traceId` | String | 是 | traceId | traceId |
| `isSuccess` | java.lang.Boolean | 是 | true | true |
