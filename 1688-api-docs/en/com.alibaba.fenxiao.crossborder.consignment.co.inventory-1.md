# Update warehouse product inventory

Original name: 仓库商品库存更新  
API: `com.alibaba.fenxiao.crossborder:consignment.co.inventory:1` · Category: Fully Managed (Consignment)  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:consignment.co.inventory-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/consignment.co.inventory/{appKey}`  
No user authorization · Requires signature

Update the inventory of products in the warehouse.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | no | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `bizSkuId` | java.lang.String | yes | External skuId | 12212122 |
| `sourceItemId` | java.lang.String | yes | Warehouse product ID | 12312323123 |
| `outboundQuantity` | java.lang.String | yes | Outbound quantity | 12 |
| `bizId` | java.lang.String | yes | Business ID, used to ensure idempotency of stock deduction | 12321312 |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.national.common.result.ResultModel](#m-alibaba-national-common-result-resultmodel) | yes |  |  |

<a id="m-alibaba-national-common-result-resultmodel"></a>
#### alibaba.national.common.result.ResultModel

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `msg` | java.lang.String | yes | Info | 成功 |
| `traceId` | java.lang.String | yes | Trace id | 12312312313 |
| `success` | java.lang.Boolean | yes | Success/Failure | true |
| `model` | java.lang.Object | yes | Return result | null |
