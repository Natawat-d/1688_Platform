# Query product count in a product pool

Original name: 查询商品池中商品总数  
API: `com.alibaba.fenxiao.crossborder:pool.product.total:1` · Category: Tools  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:pool.product.total-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/pool.product.total/{appKey}`  
No user authorization · Requires signature

Query the total number of products in a product pool.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | no | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `appKey` | java.lang.String | yes | appKey | 123 |
| `palletId` | java.lang.Long | yes | palletId | 123 |
| `categoryId` | java.lang.String | no | categoryId | 1 |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.national.pool.product.total.new.ResultModel](#m-alibaba-national-pool-product-total-new-resultmodel) | yes | Result | {} |

<a id="m-alibaba-national-pool-product-total-new-resultmodel"></a>
#### alibaba.national.pool.product.total.new.ResultModel

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `msg` | String | yes | msg | 1 |
| `code` | String | yes | code | 001 |
| `traceId` | String | yes | traceId | 12121212333443212sada |
| `success` | String | yes | Flag indicating whether the request was successful | true |
| `model` | Long | yes | Total count | 100 |
