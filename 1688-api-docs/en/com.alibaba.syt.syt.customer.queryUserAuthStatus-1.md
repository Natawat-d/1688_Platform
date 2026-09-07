# 88 ShengYiTong: check whether customer is verified

Original name: 88生意通查询客户是否已认证  
API: `com.alibaba.syt:syt.customer.queryUserAuthStatus:1` · Category: 88 ShengYiTong (Business Link)  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.syt:syt.customer.queryUserAuthStatus-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.syt/syt.customer.queryUserAuthStatus/{appKey}`  
Requires user authorization (access_token) · Requires signature

88 ShengYiTong: check whether a customer has completed verification.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `request` | [message:alibaba.china.share.creditpay.gateway.api.request.UserAuthStatusQueryApiRequest](#m-alibaba-china-share-creditpay-gateway-api-request-userauthstatusqueryapirequest) | yes | Input parameters | {   "loginId": "user1688" } |

<a id="m-alibaba-china-share-creditpay-gateway-api-request-userauthstatusqueryapirequest"></a>
#### alibaba.china.share.creditpay.gateway.api.request.UserAuthStatusQueryApiRequest

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `loginId` | java.lang.String | yes | 1688 loginId | ces测试002 |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.china.share.creditpay.gateway.api.response.UserAuthStatusQueryApiResponse](#m-alibaba-china-share-creditpay-gateway-api-response-userauthstatusqueryapiresponse) | yes | Output result | {   "success": true,   "errorCode": null,   "errorMsg": null,   "errorNumCode": null,   "authenticated": true } |

<a id="m-alibaba-china-share-creditpay-gateway-api-response-userauthstatusqueryapiresponse"></a>
#### alibaba.china.share.creditpay.gateway.api.response.UserAuthStatusQueryApiResponse

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `authenticated` | java.lang.Boolean | yes | Whether real-name verification has been completed | true |
| `errorCode` | java.lang.String | yes | Error code, to be noted when isSuccess=true |  SUCCESS |
| `errorMsg` | java.lang.String | yes | Error message description; relevant when isSuccess=true | 调用成功 |
| `isSuccess` | java.lang.Boolean | yes | Whether the call succeeded | true |
| `traceId` | String | yes | traceId | 12222 |
