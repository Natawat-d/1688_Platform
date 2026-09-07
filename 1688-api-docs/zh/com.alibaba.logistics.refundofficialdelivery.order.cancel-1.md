# 取消退货官方物流上门揽订单

API: `com.alibaba.logistics:refundofficialdelivery.order.cancel:1` · Category: 官方退上门取件  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.logistics:refundofficialdelivery.order.cancel-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.logistics/refundofficialdelivery.order.cancel/{appKey}`  
需要授权 (access_token) · 需要签名

取消退货官方物流上门揽订单

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `param` | [message:alibaba.ocean.openplatform.biz.logistics.param.OfficialDeliveryOrderCancelParam](#m-alibaba-ocean-openplatform-biz-logistics-param-officialdeliveryordercancelparam) | 是 |  |  |

<a id="m-alibaba-ocean-openplatform-biz-logistics-param-officialdeliveryordercancelparam"></a>
#### alibaba.ocean.openplatform.biz.logistics.param.OfficialDeliveryOrderCancelParam

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `cancelText` | java.lang.String | 是 | 取消原因 | 计划有变，暂时不需要寄了 |
| `cancelType` | java.lang.String | 是 | 取消原因类型  INFORMATION_INCORRECT(&quot;信息填错了（需要修改是时间/地址等）&quot;),     WANT_TO_SEND_MYSELF(&quot;想去附近的服务点自寄&quot;),     PLAN_CHANGED(&quot;计划有变，暂时不需要寄了&quot;),     WANT_TO_CHANGE_PICKUP_TIME(&quot;我想换个上门取件时间&quot;),     PRICE_TOO_HIGH(&quot;我觉得价格有点贵&quot;), 	ATE(&quot;快递员未准时上门取件&quot;),     COURIER_NOT_COMING(&quot;快递员不上门&quot;),     COURIER_BAD_ATTITUDE(&quot;快递员服务态度不好&quot;),     ITEM_CANNOT_BE_SHIPPED(&quot;物品类型无法邮寄&quot;),     COURIER_TOO_BUSY(&quot;快递员反馈因运力紧张暂无法揽收&quot;),     OTHER(&quot;其它&quot;), | PLAN_CHANGED |
| `officialDeliveryOrderId` | java.lang.Long | 是 | 上门揽订单id | 7210001 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:alibaba.openapi.shared.common.refundofficialdelivery.order.cancel.ResultModel](#m-alibaba-openapi-shared-common-refundofficialdelivery-order-cancel-resultmodel) | 是 |  |  |

<a id="m-alibaba-openapi-shared-common-refundofficialdelivery-order-cancel-resultmodel"></a>
#### alibaba.openapi.shared.common.refundofficialdelivery.order.cancel.ResultModel

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `success` | boolean | 是 | 是否成功 | true |
| `code` | java.lang.String | 是 | 错误码 | "" |
| `message` | java.lang.String | 是 | 错误信息 | "" |
| `result` | java.lang.Boolean | 是 | 订单取消业务是否成功 | true |
