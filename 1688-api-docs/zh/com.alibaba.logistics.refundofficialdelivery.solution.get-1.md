# 官方退货上门揽方案查询

API: `com.alibaba.logistics:refundofficialdelivery.solution.get:1` · Category: 官方退上门取件  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.logistics:refundofficialdelivery.solution.get-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.logistics/refundofficialdelivery.solution.get/{appKey}`  
需要授权 (access_token) · 需要签名

官方退货上门揽方案查询

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `param` | [message:alibaba.ocean.openplatform.biz.logistics.param.OfficialDeliverySolutionQueryParam](#m-alibaba-ocean-openplatform-biz-logistics-param-officialdeliverysolutionqueryparam) | 是 |  |   |

<a id="m-alibaba-ocean-openplatform-biz-logistics-param-officialdeliverysolutionqueryparam"></a>
#### alibaba.ocean.openplatform.biz.logistics.param.OfficialDeliverySolutionQueryParam

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `receiverInfo` | [message:alibaba.ocean.openplatform.biz.trade.param.AddressParam](#m-alibaba-ocean-openplatform-biz-trade-param-addressparam) | 否 | 收件人，即卖家收货地址信息 | {} |
| `senderInfo` | [message:alibaba.ocean.openplatform.biz.trade.param.AddressParam](#m-alibaba-ocean-openplatform-biz-trade-param-addressparam) | 否 | 发件人，即买家上门揽地址 | {} |
| `receiverInfoStr` | java.lang.String | 否 | 收件人，即卖家收货地址,文本 | {\"townName\":\"东华门街道\",\"townCode\":\"110101001\",\"cityCode\":\"110100\",\"contactorName\":\"234234\",\"provinceCode\":\"110000\",\"mobile\":\"12342344324\",\"addressId\":111028,\"areaCode\":\"110101\",\"cityName\":\"北京市\",\"post\":\"123123\",\"areaName\":\"东城区\",\"detailAddress\":\"接到\",\"provinceName\":\"北京\"} |
| `senderInfoStr` | java.lang.String | 否 | 发件人，即买家上门揽地址，文本 | {\"townName\":\"东华门街道\",\"townCode\":\"110101001\",\"cityCode\":\"110100\",\"contactorName\":\"234234\",\"provinceCode\":\"110000\",\"mobile\":\"12342344324\",\"addressId\":111028,\"areaCode\":\"110101\",\"cityName\":\"北京市\",\"post\":\"123123\",\"areaName\":\"东城区\",\"detailAddress\":\"接到\",\"provinceName\":\"北京\"} |
| `refundId` | java.lang.String | 是 | 退款单Id | TQ284019625557092048 |
| `packageCount` | java.lang.Integer | 是 | 包裹数 | 1 |
| `totalVolume` | java.lang.String | 否 | 总体积cm^3 | 1 |
| `totalWeight` | java.lang.String | 否 | 总重量g | 1 |
| `couponId` | String | 否 | 优惠券id |   |

<a id="m-alibaba-ocean-openplatform-biz-trade-param-addressparam"></a>
#### alibaba.ocean.openplatform.biz.trade.param.AddressParam

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `address` | java.lang.String | 是 |  |    |
| `addressId` | java.lang.Long | 是 |  |    |
| `areaCode` | java.lang.String | 是 |  |    |
| `areaText` | java.lang.String | 是 |  |    |
| `cityCode` | java.lang.String | 是 |  |    |
| `cityText` | java.lang.String | 是 |  |    |
| `districtCode` | java.lang.String | 是 |  |     |
| `fullName` | java.lang.String | 是 |  |    |
| `mobile` | java.lang.String | 是 |  |    |
| `phone` | java.lang.String | 是 |  |    |
| `postCode` | java.lang.String | 是 |  |    |
| `provinceCode` | java.lang.String | 是 |  |    |
| `provinceText` | java.lang.String | 是 |  |    |
| `townCode` | java.lang.String | 是 |  |    |
| `townText` | java.lang.String | 是 |  |    |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:alibaba.openapi.shared.common.refundofficialdelivery.solution.get.ResultModel](#m-alibaba-openapi-shared-common-refundofficialdelivery-solution-get-resultmodel) | 是 |  |   |

<a id="m-alibaba-openapi-shared-common-refundofficialdelivery-solution-get-resultmodel"></a>
#### alibaba.openapi.shared.common.refundofficialdelivery.solution.get.ResultModel

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `success` | boolean | 是 | 是否成功 | true |
| `code` | java.lang.String | 是 | 错误码 | "" |
| `message` | java.lang.String | 是 | 错误信息 | "" |
| `result` | [message:alibaba.ocean.openplatform.biz.logistics.result.OfficialDeliverySolutionResult](#m-alibaba-ocean-openplatform-biz-logistics-result-officialdeliverysolutionresult) | 是 | 返回值 | {} |

<a id="m-alibaba-ocean-openplatform-biz-logistics-result-officialdeliverysolutionresult"></a>
#### alibaba.ocean.openplatform.biz.logistics.result.OfficialDeliverySolutionResult

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `height` | java.lang.Long | 是 | 高度 cm | 3 |
| `length` | java.lang.Long | 是 | 长度 cm | 3 |
| `offerInfoDTOList` | [message:alibaba.ocean.openplatform.biz.logistics.result.OfficialDeliveryOfferInfo[]](#m-alibaba-ocean-openplatform-biz-logistics-result-officialdeliveryofferinfo[]) | 是 | 商品信息 | "" |
| `orderEntryIdList` | java.lang.Long[] | 是 | 子订单 | 87929773894 |
| `orderId` | java.lang.Long | 是 | 子订单 | 87929773894 |
| `receiverInfoDTO` | [message:alibaba.ocean.openplatform.biz.trade.param.AddressParam](#m-alibaba-ocean-openplatform-biz-trade-param-addressparam) | 是 | 收件人 | {} |
| `sceneCode` | java.lang.String | 是 | 场景code | "" |
| `senderInfoDTO` | [message:alibaba.ocean.openplatform.biz.trade.param.AddressParam](#m-alibaba-ocean-openplatform-biz-trade-param-addressparam) | 是 | 发件人 | {} |
| `signErrorMsg` | java.lang.String | 是 | 错误信息 | "" |
| `solutionDTOList` | [message:alibaba.ocean.openplatform.biz.logistics.result.OfficialDeliverySolution[]](#m-alibaba-ocean-openplatform-biz-logistics-result-officialdeliverysolution[]) | 是 | 方案列表 | {} |
| `totalVolume` | java.lang.Long | 是 | 总体积 cm^3 | 12 |
| `totalWeight` | java.lang.Long | 是 | 总重量 g | 12 |
| `userAlipaySignUrl` | java.lang.String | 是 | 用户支付宝sign地址 | http:// |
| `userSettled` | java.lang.Boolean | 是 | 是否settled | true |
| `vipUser` | java.lang.Boolean | 是 | 是否vip用户 | true |
| `width` | java.lang.Long | 是 | 宽度 | 2 |

<a id="m-alibaba-ocean-openplatform-biz-logistics-result-officialdeliveryofferinfo[]"></a>
#### alibaba.ocean.openplatform.biz.logistics.result.OfficialDeliveryOfferInfo[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `offerId` | java.lang.Long | 是 | 商品id | 1 |
| `offerName` | java.lang.String | 是 | 商品名称 | "xxx商品" |

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

<a id="m-alibaba-ocean-openplatform-biz-logistics-result-officialdeliverysolution[]"></a>
#### alibaba.ocean.openplatform.biz.logistics.result.OfficialDeliverySolution[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `appointmentTimeDTO` | [message:alibaba.ocean.openplatform.biz.logistics.result.OfficialDeliverySolutionDTO.OfficialAppointmentTime](#m-alibaba-ocean-openplatform-biz-logistics-result-officialdeliverysolutiondto-officialappointmenttime) | 是 | 预约配送时间 | "" |
| `chargeDetailDTO` | [message:alibaba.ocean.openplatform.biz.logistics.result.OfficialDeliverySolutionDTO.OfficialChargeDetail](#m-alibaba-ocean-openplatform-biz-logistics-result-officialdeliverysolutiondto-officialchargedetail) | 是 | 计费详情 | "" |
| `cpCode` | java.lang.String | 是 | 物流服务商编码 | "YTO" |
| `cpName` | java.lang.String | 是 | 物流服务商名称 | "圆通速递" |
| `deliveryTimeDTO` | [message:alibaba.ocean.openplatform.biz.logistics.result.OfficialDeliverySolutionDTO.OfficialDeliveryTime](#m-alibaba-ocean-openplatform-biz-logistics-result-officialdeliverysolutiondto-officialdeliverytime) | 是 | 配送时间 | "" |
| `priority` | java.lang.Integer | 是 | 优先级，值越小优先级越高 | 1 |
| `remark` | java.lang.String | 是 | 备注 | "圆通快递价格优惠覆盖广" |
| `serviceType` | java.lang.String | 是 | 解决方案类型 | "1" |
| `serviceTypeDesc` | java.lang.String | 是 | 解决方案类型描述 | "圆通快递当日取" |
| `serviceTypeName` | java.lang.String | 是 | 解决方案类型名称 | "当日上门" |
| `solutionCode` | java.lang.String | 是 | 物流解决方案编码 | "422" |
| `solutionName` | java.lang.String | 是 | 物流解决方案名称 | "圆通快递当日取" |

<a id="m-alibaba-ocean-openplatform-biz-logistics-result-officialdeliverysolutiondto-officialappointmenttime"></a>
#### alibaba.ocean.openplatform.biz.logistics.result.OfficialDeliverySolutionDTO.OfficialAppointmentTime

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `timeSlotDTOList` | [message:alibaba.ocean.openplatform.biz.logistics.result.OfficialDeliverySolutionDTO.OfficialAppointmentTimeSlot[]](#m-alibaba-ocean-openplatform-biz-logistics-result-officialdeliverysolutiondto-officialappointmenttimeslot[]) | 是 | 时间段列表 | xx |

<a id="m-alibaba-ocean-openplatform-biz-logistics-result-officialdeliverysolutiondto-officialappointmenttimeslot[]"></a>
#### alibaba.ocean.openplatform.biz.logistics.result.OfficialDeliverySolutionDTO.OfficialAppointmentTimeSlot[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `dateStr` | java.lang.String | 是 | 日期 | "2025-06-18" |
| `gmtEnd` | java.lang.String | 是 | 截止时间 | "17:00" |
| `gmtStart` | java.lang.String | 是 | 开始时间 | "15:00" |

<a id="m-alibaba-ocean-openplatform-biz-logistics-result-officialdeliverysolutiondto-officialchargedetail"></a>
#### alibaba.ocean.openplatform.biz.logistics.result.OfficialDeliverySolutionDTO.OfficialChargeDetail

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `deliveryFee` | java.lang.Long | 是 | 配送费，单位分 | 1000 |
| `deliveryTotalFee` | java.lang.Long | 是 | 包含配送费+超区费等费用的运输费，单位分 | 1000 |
| `firstFee` | java.lang.Long | 是 | 首重费，单位分 | 1000 |
| `firstTotalCost` | java.lang.Long | 是 | 首重总费用，单位分 | 1000 |
| `firstUnit` | java.lang.Long | 是 | 首重单位，单位g | 1000 |
| `nextFee` | java.lang.Long | 是 | 续重费，单位分 | 1000 |
| `nextTotalCost` | java.lang.Long | 是 | 续重总费用，单位分 | 1000 |
| `nextUnit` | java.lang.Long | 是 | 续重单位，单位g | 1000 |
| `originTotalCost` | java.lang.Long | 是 | 原总费用（非会员价），单位分 | 1200 |
| `totalCost` | java.lang.Long | 是 | 总费用 = 运费 + 增值服务费 - 红包金额，单位分 | 1000 |
| `vasDetailDTOList` | [message:alibaba.ocean.openplatform.biz.logistics.result.OfficialDeliverySolutionDTO.OfficialChargeVasDetail[]](#m-alibaba-ocean-openplatform-biz-logistics-result-officialdeliverysolutiondto-officialchargevasdetail[]) | 是 | 增值服务费列表 | "" |
| `vasTotalFee` | java.lang.Long | 是 | 增值服务费总费用，单位分 | 1000 |
| `officialCouponDTO` | [message:alibaba.ocean.openplatform.biz.logistics.result.OfficialDeliverySolutionDTO.OfficialCoupon](#m-alibaba-ocean-openplatform-biz-logistics-result-officialdeliverysolutiondto-officialcoupon) | 是 | 优惠券 |   |
| `fromOverAreaFee` | Long | 是 | 揽收超区费，单位分 |   |
| `toOverAreaFee` | Long | 是 | 签收超区费，单位分 |   |
| `overDistanceFee` | Long | 是 | 超长费，单位分 |   |

<a id="m-alibaba-ocean-openplatform-biz-logistics-result-officialdeliverysolutiondto-officialchargevasdetail[]"></a>
#### alibaba.ocean.openplatform.biz.logistics.result.OfficialDeliverySolutionDTO.OfficialChargeVasDetail[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `cost` | java.lang.Long | 是 | 增值服务费用，单位分 | 1000 |
| `dialogTextContent` | java.lang.String | 是 | 增值服务弹窗内容 | xx |
| `dialogTitle` | java.lang.String | 是 | 增值服务弹窗标题 | xx |
| `dialogUrl` | java.lang.String | 是 | 增值服务弹窗链接 | xx |
| `operateType` | java.lang.String | 是 | 操作类型 | UN_SELECTABLE：不可操作 |
| `selectedValue` | java.lang.String | 是 | 增值服务选中情况 | xx |
| `vasCode` | java.lang.String | 是 | 增值服务code | xx |
| `vasDesc` | java.lang.String | 是 | 增值服务描述 | xx |
| `vasName` | java.lang.String | 是 | 增值服务名称 | xx |

<a id="m-alibaba-ocean-openplatform-biz-logistics-result-officialdeliverysolutiondto-officialcoupon"></a>
#### alibaba.ocean.openplatform.biz.logistics.result.OfficialDeliverySolutionDTO.OfficialCoupon

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `id` | String | 是 | 优惠券id |   |
| `name` | String | 是 | 名称 |   |
| `description` | String | 是 | 描述 |   |
| `amount` | Long | 是 | 优惠券金额（分） |   |
| `couponType` | String | 是 | 优惠券类型 |   |
| `status` | String | 是 | 优惠券状态 |   |
| `gmtEffect` | Date | 是 | 优惠券生效时间 |   |
| `gmtExpired` | Date | 是 | 优惠券过期时间 |   |

<a id="m-alibaba-ocean-openplatform-biz-logistics-result-officialdeliverysolutiondto-officialdeliverytime"></a>
#### alibaba.ocean.openplatform.biz.logistics.result.OfficialDeliverySolutionDTO.OfficialDeliveryTime

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `deliveryText` | java.lang.String | 是 | x | x |
| `deliveryTime` | java.lang.String | 是 | x | x |
| `remark` | java.lang.String | 是 | x | x |
