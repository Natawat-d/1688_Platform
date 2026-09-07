# 创建退货官方物流上门揽订单

API: `com.alibaba.logistics:refundofficialdelivery.order.create:1` · Category: 官方退上门取件  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.logistics:refundofficialdelivery.order.create-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.logistics/refundofficialdelivery.order.create/{appKey}`  
需要授权 (access_token) · 需要签名

创建退货官方物流上门揽订单

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `param` | [message:alibaba.ocean.openplatform.biz.logistics.param.OfficialDeliveryOrderCreateParam](#m-alibaba-ocean-openplatform-biz-logistics-param-officialdeliveryordercreateparam) | 是 |  |  |

<a id="m-alibaba-ocean-openplatform-biz-logistics-param-officialdeliveryordercreateparam"></a>
#### alibaba.ocean.openplatform.biz.logistics.param.OfficialDeliveryOrderCreateParam

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `couponId` | java.lang.String | 否 | 优惠券id | 1234123 |
| `dateStr` | java.lang.String | 是 | 上门揽时间yyyy-MM-dd | 2024-07-11 |
| `gmtAcceptStart` | java.lang.String | 是 | 上门揽时间起 HH:mm | 9:00 |
| `gmtAcceptEnd` | java.lang.String | 是 | 上门揽时间止 HH:mm | 21:00 |
| `goodsType` | java.lang.String | 是 | 商品类型 (&quot;DAILY_NECESSITIES&quot;, &quot;日用品&quot;); (&quot;FOOD&quot;, &quot;食品&quot;); (&quot;FURNITURE&quot;, &quot;家具&quot;); (&quot;METALS&quot;, &quot;五金&quot;); (&quot;COSMETICS&quot;, &quot;化妆品&quot;); (&quot;DIGITAL&quot;, &quot;数码&quot;); (&quot;GIFT&quot;, &quot;礼品&quot;); (&quot;OTHER&quot;, &quot;其他&quot;); | DAILY_NECESSITIES |
| `packageCount` | java.lang.String | 是 | 包裹数量 | 1 |
| `receiverInfo` | [message:alibaba.ocean.openplatform.biz.trade.param.AddressParam](#m-alibaba-ocean-openplatform-biz-trade-param-addressparam) | 否 | 收件人信息 | {} |
| `receiverInfoStr` | java.lang.String | 否 | 收件人信息文本 | "" |
| `refundId` | java.lang.String | 是 | 退款单id | TQ284160529017092048 |
| `senderInfo` | [message:alibaba.ocean.openplatform.biz.trade.param.AddressParam](#m-alibaba-ocean-openplatform-biz-trade-param-addressparam) | 否 | 发件人信息 | {} |
| `senderInfoStr` | java.lang.String | 否 | 发件人文本 | "" |
| `solutionCode` | java.lang.String | 是 | 上门揽方案code | 434 |
| `totalVolume` | java.lang.String | 否 | 总体积 单位cm^3 | 10 |
| `totalWeight` | java.lang.String | 否 | 总重量 单位g | 1000 |

<a id="m-alibaba-ocean-openplatform-biz-trade-param-addressparam"></a>
#### alibaba.ocean.openplatform.biz.trade.param.AddressParam

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `address` | java.lang.String | 是 |  |  |
| `addressId` | java.lang.Long | 是 |  |  |
| `areaCode` | java.lang.String | 是 |  |  |
| `areaText` | java.lang.String | 是 |  |  |
| `cityCode` | java.lang.String | 是 |  |  |
| `cityText` | java.lang.String | 是 |  |  |
| `districtCode` | java.lang.String | 是 |  |  |
| `fullName` | java.lang.String | 是 |  |  |
| `mobile` | java.lang.String | 是 |  |  |
| `phone` | java.lang.String | 是 |  |  |
| `postCode` | java.lang.String | 是 |  |  |
| `provinceCode` | java.lang.String | 是 |  |  |
| `provinceText` | java.lang.String | 是 |  |  |
| `townCode` | java.lang.String | 是 |  |  |
| `townText` | java.lang.String | 是 |  |  |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:alibaba.openapi.shared.common.refundofficialdelivery.order.create.ResultModel](#m-alibaba-openapi-shared-common-refundofficialdelivery-order-create-resultmodel) | 是 |  |  |

<a id="m-alibaba-openapi-shared-common-refundofficialdelivery-order-create-resultmodel"></a>
#### alibaba.openapi.shared.common.refundofficialdelivery.order.create.ResultModel

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `success` | boolean | 是 | 是否成功 | true |
| `code` | java.lang.String | 是 | 错误码 | "" |
| `message` | java.lang.String | 是 | 错误消息 | "" |
| `result` | Long | 是 | 返回的订单id | 19882 |
