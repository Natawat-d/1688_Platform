# 取消跨境关注商品

API: `com.alibaba.fenxiao.crossborder:product.kjdistribute.removeRelation:1` · Category: 商品  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:product.kjdistribute.removeRelation-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/product.kjdistribute.removeRelation/{appKey}`  
需要授权 (access_token) · 需要签名

取消跨境关注商品

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `offerId` | Long | 是 | 商品id | 232324312 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:alibaba.national.common.result.ResultModel](#m-alibaba-national-common-result-resultmodel) | 是 | 返回结果 | {"msg":"错误信息","success":true,"traceId":"32327748dcadwwe923934fcdcrfr33"} |

<a id="m-alibaba-national-common-result-resultmodel"></a>
#### alibaba.national.common.result.ResultModel

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `model` | java.lang.String[] | 是 |  |  |
| `msg` | java.lang.String | 是 |  |  |
| `traceId` | java.lang.String | 是 |  |  |
| `success` | java.lang.Boolean | 是 |  |  |
