# 88 ShengYiTong: void contract

Original name: 88生意通合同作废  
API: `com.alibaba.syt:syt.contract.invalid:1` · Category: 88 ShengYiTong (Business Link)  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.syt:syt.contract.invalid-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.syt/syt.contract.invalid/{appKey}`  
Requires user authorization (access_token) · Requires signature

88 ShengYiTong: void (invalidate) a contract.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `request` | [message:alibaba.china.share.creditpay.gateway.api.request.ContractInvalidApiRequest](#m-alibaba-china-share-creditpay-gateway-api-request-contractinvalidapirequest) | yes | Request parameters | {   "draftNo": "DRAFT20251230001" } |

<a id="m-alibaba-china-share-creditpay-gateway-api-request-contractinvalidapirequest"></a>
#### alibaba.china.share.creditpay.gateway.api.request.ContractInvalidApiRequest

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `draftNo` | java.lang.String | yes | Contract number | 88SYT20251201169011 |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.china.share.creditpay.gateway.api.response.ContractInvalidApiResponse](#m-alibaba-china-share-creditpay-gateway-api-response-contractinvalidapiresponse) | yes | Response parameters | {   "success": true,   "code": "200",   "message": "合同失效成功",   "draftNo": "DRAFT20251230001",   "contractCurrentStatus": "INVALID" } |

<a id="m-alibaba-china-share-creditpay-gateway-api-response-contractinvalidapiresponse"></a>
#### alibaba.china.share.creditpay.gateway.api.response.ContractInvalidApiResponse

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `draftNo` | java.lang.String | yes | Contract number | 88SYT20251201169011 |
| `contractCurrentStatus` | java.lang.String | yes | Current contract status | INVALID |
| `success` | Boolean | yes | Call result | true |
| `errorCode` | String | yes | Error message | SUCCESS |
| `errorMsg` | String | yes | Error description | 错误描述 |
| `traceId` | String | yes | traceId. Must be recorded and provided to the platform for troubleshooting. | traceId |
