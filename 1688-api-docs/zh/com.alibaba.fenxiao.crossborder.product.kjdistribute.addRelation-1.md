# 增加跨境关注商品

API: `com.alibaba.fenxiao.crossborder:product.kjdistribute.addRelation:1` · Category: 商品  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:product.kjdistribute.addRelation-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/product.kjdistribute.addRelation/{appKey}`  
无需授权 · 需要签名

增加跨境关注商品

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 否 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `offerId` | java.lang.Long | 是 | 商品id |  |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:alibaba.national.common.result.ResultModel](#m-alibaba-national-common-result-resultmodel) | 是 |  |  |

<a id="m-alibaba-national-common-result-resultmodel"></a>
#### alibaba.national.common.result.ResultModel

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `msg` | java.lang.String | 是 |  |  |
| `success` | java.lang.Boolean | 是 |  |  |
| `traceId` | java.lang.String | 是 |  |  |
