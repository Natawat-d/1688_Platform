# 获取运单号集合

API: `com.alibaba.fenxiao.crossborder:bigcustomer.mailNo.query:1` · Category: 物流  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:bigcustomer.mailNo.query-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/bigcustomer.mailNo.query/{appKey}`  
需要授权 (access_token) · 需要签名

获取运单号集合

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `outOrderId` | java.lang.String | 是 | 订单id | 1234 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:alibaba.national.common.result.ResultModel](#m-alibaba-national-common-result-resultmodel) | 是 | 返回结果 | {} |

<a id="m-alibaba-national-common-result-resultmodel"></a>
#### alibaba.national.common.result.ResultModel

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `model` | java.util.Map | 是 | 返回的具体数据 | {} |
| `msg` | java.lang.String | 是 | msg | 请求成功 |
| `traceId` | java.lang.String | 是 | traceId | 2121212123344343423 |
| `success` | java.lang.Boolean | 是 | 是否成功 | true |
