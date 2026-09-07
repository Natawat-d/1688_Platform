# Pull products from a product pool

Original name: 拉取商品池中商品数据  
API: `com.alibaba.fenxiao.crossborder:pool.product.pull:1` · Category: Tools  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:pool.product.pull-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/pool.product.pull/{appKey}`  
Requires user authorization (access_token) · Requires signature

Batch-pull product data directly from a product pool by pool ID.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `offerPoolQueryParam` | [message:pool.product.pull.OfferPoolQueryParam](#m-pool-product-pull-offerpoolqueryparam) | yes |  |  |

<a id="m-pool-product-pull-offerpoolqueryparam"></a>
#### pool.product.pull.OfferPoolQueryParam

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `offerPoolId` | java.lang.Long | yes | Product pool ID (business-customized and permission-controlled; obtain it from the integrated business - passing an arbitrary value will cause an error. For Xunyuantong purchasing-on-behalf, using the keyword search interface is recommended) | 111 |
| `cateId` | java.lang.Long | no | Category ID | 11 |
| `taskId` | java.lang.String | yes | Query task ID. For example, if an assortment has 10,000 products, querying 1,000 per page for 10 queries retrieves all 10,000 products; the same taskId must be passed for all 10 queries. The organization needs to fix and retain one taskId during paginated queries, then pass the same taskId each time when paging through this API. | 1 |
| `language` | java.lang.String | no | Language | en |
| `pageNo` | java.lang.Integer | yes | Page number | 1 |
| `pageSize` | java.lang.Integer | yes | Number per page | 10 |
| `sortField` | String | no | Sort field | order1m/buyer1m（order1m：最近1个月销售额排序；buyer1m：最近1个月买家数） |
| `sortType` | String | no | Sorting rule | ASC/DESC |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:pool.product.pull.ResultModel](#m-pool-product-pull-resultmodel) | yes |  |  |

<a id="m-pool-product-pull-resultmodel"></a>
#### pool.product.pull.ResultModel

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `success` | java.lang.String | yes | Whether successful | true |
| `code` | java.lang.String | yes | Error code | S0000 |
| `message` | java.lang.String | yes | Error description | 成功 |
| `result` | [message:pool.product.pull.ProductPoolModel[]](#m-pool-product-pull-productpoolmodel[]) | yes | Result | 结果 |

<a id="m-pool-product-pull-productpoolmodel[]"></a>
#### pool.product.pull.ProductPoolModel[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `offerId` | java.lang.Long | yes | Product ID | 111111 |
| `bizCategoryId` | java.lang.String | yes | Institution's category ID | 111111 |
| `offerPoolTotal` | Integer | yes | Total number of products in the pool (returned for every offer) | 122211 |
