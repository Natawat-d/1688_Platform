# 88 ShengYiTong: confirm transaction complete

Original name: 88生意通确认交易完成  
API: `com.alibaba.syt:syt.contract.confirm:1` · Category: 88 ShengYiTong (Business Link)  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.syt:syt.contract.confirm-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.syt/syt.contract.confirm/{appKey}`  
Requires user authorization (access_token) · Requires signature

88 ShengYiTong: confirm that the transaction is complete.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `request` | [message:alibaba.china.share.creditpay.gateway.api.request.ContractConfirmApiRequest](#m-alibaba-china-share-creditpay-gateway-api-request-contractconfirmapirequest) | yes | Request parameters | {"draftNo":"88SYT20251204069031"} |

<a id="m-alibaba-china-share-creditpay-gateway-api-request-contractconfirmapirequest"></a>
#### alibaba.china.share.creditpay.gateway.api.request.ContractConfirmApiRequest

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `draftNo` | java.lang.String | yes | Contract number | 88SYT20251204069031 |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.china.share.creditpay.gateway.api.response.ContractConfirmApiResponse](#m-alibaba-china-share-creditpay-gateway-api-response-contractconfirmapiresponse) | yes | Return response | {   "success": true,   "code": "200",   "message": "合同确认成功",   "draftNo": "DRAFT20251230001",   "contractCurrentStatus": "CONFIRMED" } |

<a id="m-alibaba-china-share-creditpay-gateway-api-response-contractconfirmapiresponse"></a>
#### alibaba.china.share.creditpay.gateway.api.response.ContractConfirmApiResponse

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `draftNo` | java.lang.String | yes | Contract number | 88SYT20251204069031 |
| `contractCurrentStatus` | java.lang.String | yes | Current latest contract status | CONFIRMED |
| `success` | Boolean | yes | Call result | true |
| `errorCode` | String | yes | Error code | SUCCESS |
| `errorMsg` | String | yes | Error message | 调用成功 |
| `traceId` | String | yes | Call traceId | 123333 |
