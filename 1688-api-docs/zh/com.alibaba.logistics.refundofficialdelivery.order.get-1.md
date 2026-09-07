# 退货官方物流上门揽订单详情获取

API: `com.alibaba.logistics:refundofficialdelivery.order.get:1` · Category: 官方退上门取件  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.logistics:refundofficialdelivery.order.get-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.logistics/refundofficialdelivery.order.get/{appKey}`  
需要授权 (access_token) · 需要签名

退货官方物流上门揽订单详情获取

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `officialDeliveryOrderId` | String | 是 | 官方物流上门揽订单id | 1988772 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:alibaba.openapi.shared.common.refundofficialdelivery.order.get.ResultModel](#m-alibaba-openapi-shared-common-refundofficialdelivery-order-get-resultmodel) | 是 | 返回值 | {} |

<a id="m-alibaba-openapi-shared-common-refundofficialdelivery-order-get-resultmodel"></a>
#### alibaba.openapi.shared.common.refundofficialdelivery.order.get.ResultModel

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `success` | boolean | 是 | 是否成功 | true |
| `code` | java.lang.String | 是 | 错误码 | "" |
| `message` | java.lang.String | 是 | 错误信息 | "" |
| `result` | [message:alibaba.ocean.openplatform.biz.logistics.result.OfficialDeliveryOrderInfo](#m-alibaba-ocean-openplatform-biz-logistics-result-officialdeliveryorderinfo) | 是 | 订单明细 | {} |

<a id="m-alibaba-ocean-openplatform-biz-logistics-result-officialdeliveryorderinfo"></a>
#### alibaba.ocean.openplatform.biz.logistics.result.OfficialDeliveryOrderInfo

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `bizId` | java.lang.String | 是 | 退款单id | TQ284160529017092048 |
| `bizType` | java.lang.String | 是 | 业务类型 | buyer_refund |
| `dateStr` | java.lang.String | 是 | 上门揽日期 | 2024-07-11 |
| `features` | java.util.Map | 是 | 订单features | {} |
| `gmtAcceptStart` | java.lang.String | 是 | 上门揽件时间起 | 2024-07-11 10:00 |
| `gmtAcceptEnd` | java.lang.String | 是 | 上门揽件时间止 | 2024-07-11 22:00 |
| `gmtModified` | java.util.Date | 是 | 修改时间 | 20240710173724000+0800 |
| `id` | java.lang.Long | 是 | 上门揽订单id | 7210001 |
| `lastTraceDetail` | java.lang.String | 是 | 最后物流跟踪明细 | "" |
| `lastTraceStatus` | java.lang.String | 是 | 最后物流跟踪状态 | "" |
| `lastTraceStatusName` | java.lang.String | 是 | 最后物流跟踪状态名称 | "" |
| `logisticsStatus` | java.lang.String | 是 | 物流状态 | "" |
| `mailNo` | java.lang.String | 是 | 物流运单号 | sp981920934 |
| `orderId` | java.lang.String | 是 | 订单id | "" |
| `outBizId` | java.lang.String | 是 | 外部业务id | "" |
| `payRecordDTO` | [message:alibaba.ocean.openplatform.biz.logistics.result.OfficialDeliveryOrderInfo.OfficialPayRecordDTO](#m-alibaba-ocean-openplatform-biz-logistics-result-officialdeliveryorderinfo-officialpayrecorddto) | 是 | 支付信息 | {} |
| `receiverInfo` | [message:alibaba.ocean.openplatform.biz.trade.param.AddressParam](#m-alibaba-ocean-openplatform-biz-trade-param-addressparam) | 是 | 收件人地址信息 | {} |
| `refundId` | java.lang.String | 是 | 退款单id | TQ284160529017092048 |
| `senderInfo` | [message:alibaba.ocean.openplatform.biz.trade.param.AddressParam](#m-alibaba-ocean-openplatform-biz-trade-param-addressparam) | 是 | 发货地址信息 | {} |
| `signOfficial` | java.lang.Boolean | 是 | 是否签约 | false |
| `solutionCode` | java.lang.String | 是 | 上门揽方案code | 434 |
| `status` | java.lang.String | 是 | 订单状态 | activated |

<a id="m-alibaba-ocean-openplatform-biz-logistics-result-officialdeliveryorderinfo-officialpayrecorddto"></a>
#### alibaba.ocean.openplatform.biz.logistics.result.OfficialDeliveryOrderInfo.OfficialPayRecordDTO

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `amount` | java.lang.Long | 是 | 数量 | 1 |
| `bizId` | java.lang.String | 是 | 退款单id | TQ284160529017092048 |
| `bizType` | java.lang.String | 是 | 业务类型 | buyer_refund |
| `failReason` | java.lang.String | 是 | 错误原因 | "" |
| `gmtCreate` | java.util.Date | 是 | 创建时间 | 20240710173724000+0800 |
| `gmtModified` | java.util.Date | 是 | 修改时间 | 20240710173724000+0800 |
| `id` | java.lang.Long | 是 | 上满揽订单id | 7210001 |
| `outBizId` | java.lang.String | 是 | 外部业务id | "" |
| `payChannel` | java.lang.String | 是 | 支付渠道 | "" |
| `payType` | java.lang.String | 是 | 支付类型 | "" |
| `status` | java.lang.String | 是 | 状态 | activated |
| `feature` | [message:alibaba.ocean.openplatform.biz.logistics.result.OfficialDeliveryOrderInfo.FeatureDTO](#m-alibaba-ocean-openplatform-biz-logistics-result-officialdeliveryorderinfo-featuredto) | 是 |  |   |

<a id="m-alibaba-ocean-openplatform-biz-logistics-result-officialdeliveryorderinfo-featuredto"></a>
#### alibaba.ocean.openplatform.biz.logistics.result.OfficialDeliveryOrderInfo.FeatureDTO

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `firstWeightUnit` | String | 是 | 首重单位，单位kg |   |
| `firstWeightPrice` | String | 是 | 首重价格，单位元 |   |
| `bufferedFirstWeightPrice` | String | 是 | 加buffer价后首重价格，单位元 |   |
| `nextWeightUnit` | String | 是 | 续重重单位，单位kg |   |
| `nextWeightPrice` | String | 是 | 续重价格，单位元 |   |
| `fromOverAreaFee` | String | 是 | 揽收超区费，单位元 |   |
| `toOverAreaFee` | String | 是 | 送达超区费，单位元 |   |
| `overDistanceFee` | String | 是 | 超长费，单位元 |   |
| `deliveryTotalFee` | String | 是 | 正向配送总费用 |   |
| `nextTotalFee` | String | 是 | 续重总费用，单位元 |   |
| `firstTotalFee` | String | 是 | 首重总费用，单位元 |   |
| `totalFee` | String | 是 | 总费用，单位元 |   |
| `chargeType` | String | 是 | 计费方式，按重/按体积 |   |
| `chargeWeight` | String | 是 | 计费重量，单位kg |   |
| `chargeVolume` | String | 是 | 计费体积，单位立方米 |   |
| `bubbleWeight` | String | 是 | 泡量单位kg--计费重量大于货品重量时，返回泡重；否则不反回 |   |
| `isBubble` | String | 是 | 是否记泡 |   |
| `originTotalFee` | String | 是 | 原总费用，菜鸟计算金额，单位元 |   |
| `refundId` | String | 是 | 退款单号 |   |
| `couponId` | String | 是 | 优惠券id |   |
| `couponFee` | String | 是 | 优惠券金额，单位元 |   |
| `couponUsedFee` | String | 是 | 优惠券核销金额，单位元 |   |
| `couponName` | String | 是 | 优惠券模型 |   |
| `couponDesc` | String | 是 | 优惠券描述 |   |
| `couponSource` | String | 是 | 优惠券来源 |   |
| `vasList` | String | 是 | 增值服务费列表 |   |
| `mailNo` | String | 是 | 运单号 |   |
| `settlementList` | String | 是 | 结算信息 |   |
| `orderSourceType` | String | 是 | 订单来源 |   |

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
