# Unfollow a product (cross-border)

Original name: 取消跨境关注商品  
API: `com.alibaba.fenxiao.crossborder:product.kjdistribute.removeRelation:1` · Category: Products  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:product.kjdistribute.removeRelation-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/product.kjdistribute.removeRelation/{appKey}`  
Requires user authorization (access_token) · Requires signature

Remove a cross-border followed product.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `offerId` | Long | yes | Product ID | 232324312 |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.national.common.result.ResultModel](#m-alibaba-national-common-result-resultmodel) | yes | Return result | {"msg":"错误信息","success":true,"traceId":"32327748dcadwwe923934fcdcrfr33"} |

<a id="m-alibaba-national-common-result-resultmodel"></a>
#### alibaba.national.common.result.ResultModel

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `model` | java.lang.String[] | yes |  |  |
| `msg` | java.lang.String | yes |  |  |
| `traceId` | java.lang.String | yes |  |  |
| `success` | java.lang.Boolean | yes |  |  |
