# Urge seller to ship

Original name: 催卖家发货  
API: `com.alibaba.logistics:logistics.delivery.urge:1` · Category: Logistics  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.logistics:logistics.delivery.urge-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.logistics/logistics.delivery.urge/{appKey}`  
Requires user authorization (access_token) · Requires signature

Urge the seller to ship. The order must still be in a not-yet-shipped status. Limited to once per 24 hours.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `orderId` | String | yes | Order id | 12898772891323 |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.openapi.shared.common.logistics.delivery.urge.ResultModel](#m-alibaba-openapi-shared-common-logistics-delivery-urge-resultmodel) | yes |  |  |

<a id="m-alibaba-openapi-shared-common-logistics-delivery-urge-resultmodel"></a>
#### alibaba.openapi.shared.common.logistics.delivery.urge.ResultModel

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `success` | boolean | yes |  |  |
| `code` | java.lang.String | yes |  |  |
| `message` | java.lang.String | yes |  |  |
| `result` | java.lang.Boolean | yes |  |  |
