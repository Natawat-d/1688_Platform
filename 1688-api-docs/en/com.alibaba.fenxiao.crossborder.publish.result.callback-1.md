# Publish result callback

Original name: 发布结果回调  
API: `com.alibaba.fenxiao.crossborder:publish.result.callback:1` · Category: Listing (Publishing)  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:publish.result.callback-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/publish.result.callback/{appKey}`  
Requires user authorization (access_token) · Requires signature

WB-1688 integration Method 3: publish-result callback interface.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `request` | [message:alibaba.global1688.silicon.user.api.param.WbPublishCallbackRequestDTO](#m-alibaba-global1688-silicon-user-api-param-wbpublishcallbackrequestdto) | yes | Callback request parameters | {"offerId":123456789,"imtId":"imt123","status":"SUCCESS","requestId":"req123","cards":[]} |

<a id="m-alibaba-global1688-silicon-user-api-param-wbpublishcallbackrequestdto"></a>
#### alibaba.global1688.silicon.user.api.param.WbPublishCallbackRequestDTO

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `offerId` | java.lang.Long | yes | 1688 product ID | 123456789 |
| `imtId` | java.lang.String | yes | WB product IMT ID | imt123 |
| `status` | java.lang.String | yes | Publish status | SUCCESS |
| `requestId` | java.lang.String | yes | Request ID | req123 |
| `cards` | [message:alibaba.global1688.silicon.user.api.param.WbPublishCallbackRequestDTO.WbCardResult[]](#m-alibaba-global1688-silicon-user-api-param-wbpublishcallbackrequestdto-wbcardresult[]) | yes | Card result list | [] |

<a id="m-alibaba-global1688-silicon-user-api-param-wbpublishcallbackrequestdto-wbcardresult[]"></a>
#### alibaba.global1688.silicon.user.api.param.WbPublishCallbackRequestDTO.WbCardResult[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:ResultGT52f0rf](#m-resultgt52f0rf) | yes | Callback result | {"success":true,"duplicate":false,"requestId":"req123"} |

<a id="m-resultgt52f0rf"></a>
#### ResultGT52f0rf

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `success` | java.lang.Boolean | yes | Whether successful | true |
| `result` | [message:alibaba.global1688.silicon.user.api.result.WbPublishCallbackResponseDTO](#m-alibaba-global1688-silicon-user-api-result-wbpublishcallbackresponsedto) | yes | Return result | {} |
| `code` | java.lang.String | no | Error code |  |
| `permissionName` | java.lang.String | no | Permission name |  |
| `message` | java.lang.String | no | Error message |  |

<a id="m-alibaba-global1688-silicon-user-api-result-wbpublishcallbackresponsedto"></a>
#### alibaba.global1688.silicon.user.api.result.WbPublishCallbackResponseDTO

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `success` | java.lang.Boolean | yes | Whether successful | true |
| `duplicate` | java.lang.Boolean | yes | Whether it is a duplicate request | false |
| `requestId` | java.lang.String | yes | Request ID | req123 |

## Samples

**Input parameter example**

```
{"request":{"offerId":123456789,"imtId":"imt123","status":"SUCCESS","requestId":"req123","cards":[]}}
```

**Output parameter example**

```
{"result":{"success":true,"duplicate":false,"requestId":"req123"}}
```
