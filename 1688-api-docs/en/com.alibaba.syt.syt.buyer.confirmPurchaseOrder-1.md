# 88 ShengYiTong: buyer confirms purchase order

Original name: 88生意通买家确认采购单  
API: `com.alibaba.syt:syt.buyer.confirmPurchaseOrder:1` · Category: 88 ShengYiTong (Business Link)  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.syt:syt.buyer.confirmPurchaseOrder-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.syt/syt.buyer.confirmPurchaseOrder/{appKey}`  
Requires user authorization (access_token) · Requires signature

88 ShengYiTong: the buyer confirms a purchase order.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `request` | [message:alibaba.china.share.creditpay.gateway.api.request.SignConfirmApiRequest](#m-alibaba-china-share-creditpay-gateway-api-request-signconfirmapirequest) | yes | Request parameters | {   "draftNo": "CT2026051100001" } |

<a id="m-alibaba-china-share-creditpay-gateway-api-request-signconfirmapirequest"></a>
#### alibaba.china.share.creditpay.gateway.api.request.SignConfirmApiRequest

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `draftNo` | java.lang.String | yes | Purchase order number | 88SYT123456 |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.china.share.creditpay.gateway.api.response.SignConfirmApiResponse](#m-alibaba-china-share-creditpay-gateway-api-response-signconfirmapiresponse) | yes | Response result | {   "success": true,   "errorCode": null,   "errorMsg": null,   "errorNumCode": null,   "draftNo": "CT2026051100001"} |

<a id="m-alibaba-china-share-creditpay-gateway-api-response-signconfirmapiresponse"></a>
#### alibaba.china.share.creditpay.gateway.api.response.SignConfirmApiResponse

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `draftNo` | java.lang.String | yes | Purchase order number | 88SYT123456 |
| `errorCode` | java.lang.String | yes | Call result code; relevant when isSuccess=true |  SUCCESS |
| `errorMsg` | java.lang.String | yes | Call result description; relevant when isSuccess=true | 调用成功 |
| `traceId` | String | yes | traceId | traceId |
| `isSuccess` | java.lang.Boolean | yes | true | true |
