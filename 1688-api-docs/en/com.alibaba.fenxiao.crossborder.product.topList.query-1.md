# Query ranking lists

Original name: 查询榜单列表  
API: `com.alibaba.fenxiao.crossborder:product.topList.query:1` · Category: Market Insights  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:product.topList.query-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/product.topList.query/{appKey}`  
Requires user authorization (access_token) · Requires signature

Query ranking (top) lists.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `rankQueryParams` | [message:product.topList.query.RankQueryParams](#m-product-toplist-query-rankqueryparams) | yes |  |  |

<a id="m-product-toplist-query-rankqueryparams"></a>
#### product.topList.query.RankQueryParams

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `rankId` | java.lang.String | yes | Ranking list ID; a category ID can be passed. Category rankings are currently supported | 1111 |
| `rankType` | java.lang.String | yes | Ranking list type: complex (comprehensive ranking), hot (best-seller ranking), goodPrice (best-price ranking) | complex |
| `limit` | java.lang.Integer | yes | Number of products on the ranking list, maximum 20 | 10 |
| `language` | java.lang.String | yes | Ranking list product language | en |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:product.topList.query.ResultModel](#m-product-toplist-query-resultmodel) | yes |  |  |

<a id="m-product-toplist-query-resultmodel"></a>
#### product.topList.query.ResultModel

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `success` | java.lang.Boolean | yes | Whether successful | true |
| `code` | java.lang.String | yes | Error code | S0000 |
| `message` | java.lang.String | yes | Error description | 成功 |
| `result` | [message:product.topList.query.RankModel](#m-product-toplist-query-rankmodel) | yes | Result | 结果 |

<a id="m-product-toplist-query-rankmodel"></a>
#### product.topList.query.RankModel

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `rankId` | java.lang.String | yes | Ranking list ID | 111 |
| `rankName` | java.lang.String | yes | Ranking list name | Comprehensive List |
| `rankType` | java.lang.String | yes | Ranking list type | complex |
| `rankProductModels` | [message:product.topList.query.RankProductModel[]](#m-product-toplist-query-rankproductmodel[]) | yes | Ranking list result | 如下 |

<a id="m-product-toplist-query-rankproductmodel[]"></a>
#### product.topList.query.RankProductModel[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `itemId` | java.lang.Long | yes | Product ID | 699490252651 |
| `title` | String | yes | Product title in Chinese | 2023厚底女士凉鞋软底拖鞋夏季现货踩屎感注塑鞋出口家居防滑凉鞋 |
| `translateTitle` | String | yes | Translated product title | 2023 thick-soled ladies sandals soft-soled slippers summer spot poop injection shoes export home non-slip sandals |
| `imgUrl` | java.lang.String | yes | Product image | http://img.china.alibaba.com/img/ibank/O1CN01p4SIPo1D6Wx4c0xs8_!!2201053890167-0-cib.search.jpg |
| `sort` | java.lang.Integer | yes | Product ranking | 1 |
| `serviceList` | java.lang.String[] | yes | Services included with the product: 24-hour shipping sendGoods24H, 48-hour shipping sendGoods48H | ["sendGoods48H"] |
| `buyerNum` | java.lang.Integer | yes | Number of buyers in the last 30 days | 1334 |
| `soldOut` | java.lang.Integer | yes | Number of product units sold in the last 30 days | 433454 |
| `goodsScore` | String | yes | Product transaction rating | 5 |
