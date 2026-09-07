# 88生意通查询客户是否已认证

API: `com.alibaba.syt:syt.customer.queryUserAuthStatus:1` · Category: 88生意通  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.syt:syt.customer.queryUserAuthStatus-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.syt/syt.customer.queryUserAuthStatus/{appKey}`  
需要授权 (access_token) · 需要签名

88生意通查询客户是否已认证

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `request` | [message:alibaba.china.share.creditpay.gateway.api.request.UserAuthStatusQueryApiRequest](#m-alibaba-china-share-creditpay-gateway-api-request-userauthstatusqueryapirequest) | 是 | 输入参数 | {   "loginId": "user1688" } |

<a id="m-alibaba-china-share-creditpay-gateway-api-request-userauthstatusqueryapirequest"></a>
#### alibaba.china.share.creditpay.gateway.api.request.UserAuthStatusQueryApiRequest

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `loginId` | java.lang.String | 是 | 1688 loginId | ces测试002 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:alibaba.china.share.creditpay.gateway.api.response.UserAuthStatusQueryApiResponse](#m-alibaba-china-share-creditpay-gateway-api-response-userauthstatusqueryapiresponse) | 是 | 输出结果 | {   "success": true,   "errorCode": null,   "errorMsg": null,   "errorNumCode": null,   "authenticated": true } |

<a id="m-alibaba-china-share-creditpay-gateway-api-response-userauthstatusqueryapiresponse"></a>
#### alibaba.china.share.creditpay.gateway.api.response.UserAuthStatusQueryApiResponse

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `authenticated` | java.lang.Boolean | 是 | 是否已实名认证 | true |
| `errorCode` | java.lang.String | 是 | 错误code，isSuccess=true 时关注 |  SUCCESS |
| `errorMsg` | java.lang.String | 是 | 错误信息描述，isSuccess=true 时关注 | 调用成功 |
| `isSuccess` | java.lang.Boolean | 是 | 调用是否成功 | true |
| `traceId` | String | 是 | traceId | 12222 |
