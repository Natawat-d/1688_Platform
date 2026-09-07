# Warehouse receipt / shelving of goods

Original name: 仓库签收/上架商品  
API: `com.alibaba.fenxiao.crossborder:consignment.co.status.sync:1` · Category: Fully Managed (Consignment)  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:consignment.co.status.sync-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/consignment.co.status.sync/{appKey}`  
No user authorization · Requires signature

Open to Buffalo. Syncs the Buffalo warehouse's receipt or shelving of goods to the 1688 shipment-order status.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | no | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `bizInboundOrderId` | java.lang.String | yes | Inbound (warehouse receipt) order id |  |
| `status` | java.lang.String | yes | signed/onShelves |  |
| `processTime` | java.lang.Long | yes | Timestamp |  |
| `num` | String | yes | Quantity |  |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.national.common.result.ResultModel](#m-alibaba-national-common-result-resultmodel) | yes |  |  |

<a id="m-alibaba-national-common-result-resultmodel"></a>
#### alibaba.national.common.result.ResultModel

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `msg` | java.lang.String | yes | Error message | 参数为空 |
| `traceId` | java.lang.String | yes | traceId | 111 |
| `success` | java.lang.Boolean | yes | true/false | true |
| `model` | java.lang.Object | yes | null | null |
