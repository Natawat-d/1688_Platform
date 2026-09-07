# Follow a product (cross-border)

Original name: 增加跨境关注商品  
API: `com.alibaba.fenxiao.crossborder:product.kjdistribute.addRelation:1` · Category: Products  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:product.kjdistribute.addRelation-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/product.kjdistribute.addRelation/{appKey}`  
No user authorization · Requires signature

Add a cross-border followed product.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | no | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `offerId` | java.lang.Long | yes | Product ID |  |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.national.common.result.ResultModel](#m-alibaba-national-common-result-resultmodel) | yes |  |  |

<a id="m-alibaba-national-common-result-resultmodel"></a>
#### alibaba.national.common.result.ResultModel

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `msg` | java.lang.String | yes |  |  |
| `success` | java.lang.Boolean | yes |  |  |
| `traceId` | java.lang.String | yes |  |  |
