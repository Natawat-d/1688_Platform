# 买家删除已关闭的订单

API: `com.alibaba.trade:trade.order.buyerdelete:1` · Category: 订单  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:trade.order.buyerdelete-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/trade.order.buyerdelete/{appKey}`  
需要授权 (access_token) · 需要签名

买家删除已关闭的订单

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `orderId` | java.lang.Long | 是 | 订单id | 12213412341 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:alibaba.openapi.shared.common.trade.order.buyerdelete.ResultModel](#m-alibaba-openapi-shared-common-trade-order-buyerdelete-resultmodel) | 是 | 返回值 | {} |

<a id="m-alibaba-openapi-shared-common-trade-order-buyerdelete-resultmodel"></a>
#### alibaba.openapi.shared.common.trade.order.buyerdelete.ResultModel

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `success` | java.lang.Boolean | 是 | 是否成功 | true |
| `code` | java.lang.String | 是 | 错误码 | "" |
| `message` | java.lang.String | 是 | 错误信息 | "" |
| `result` | java.lang.Boolean | 是 | 业务是否成功 | true |
