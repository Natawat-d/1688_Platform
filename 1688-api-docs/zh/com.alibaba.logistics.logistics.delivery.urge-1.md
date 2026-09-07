# 催卖家发货

API: `com.alibaba.logistics:logistics.delivery.urge:1` · Category: 物流  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.logistics:logistics.delivery.urge-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.logistics/logistics.delivery.urge/{appKey}`  
需要授权 (access_token) · 需要签名

催卖家发货，催发货的订单状态必须为未发货之前。24小时内限制最多一次

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `orderId` | String | 是 | 订单id | 12898772891323 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:alibaba.openapi.shared.common.logistics.delivery.urge.ResultModel](#m-alibaba-openapi-shared-common-logistics-delivery-urge-resultmodel) | 是 |  |  |

<a id="m-alibaba-openapi-shared-common-logistics-delivery-urge-resultmodel"></a>
#### alibaba.openapi.shared.common.logistics.delivery.urge.ResultModel

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `success` | boolean | 是 |  |  |
| `code` | java.lang.String | 是 |  |  |
| `message` | java.lang.String | 是 |  |  |
| `result` | java.lang.Boolean | 是 |  |  |
