# Query official return-pickup plans

Original name: 官方退货上门揽方案查询  
API: `com.alibaba.logistics:refundofficialdelivery.solution.get:1` · Category: Official Return Pickup  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.logistics:refundofficialdelivery.solution.get-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.logistics/refundofficialdelivery.solution.get/{appKey}`  
Requires user authorization (access_token) · Requires signature

Query the available official door-to-door return-pickup plans.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `param` | [message:alibaba.ocean.openplatform.biz.logistics.param.OfficialDeliverySolutionQueryParam](#m-alibaba-ocean-openplatform-biz-logistics-param-officialdeliverysolutionqueryparam) | yes |  |   |

<a id="m-alibaba-ocean-openplatform-biz-logistics-param-officialdeliverysolutionqueryparam"></a>
#### alibaba.ocean.openplatform.biz.logistics.param.OfficialDeliverySolutionQueryParam

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `receiverInfo` | [message:alibaba.ocean.openplatform.biz.trade.param.AddressParam](#m-alibaba-ocean-openplatform-biz-trade-param-addressparam) | no | Recipient, i.e. the seller's shipping address information | {} |
| `senderInfo` | [message:alibaba.ocean.openplatform.biz.trade.param.AddressParam](#m-alibaba-ocean-openplatform-biz-trade-param-addressparam) | no | Sender, i.e., the buyer's door-to-door pickup address | {} |
| `receiverInfoStr` | java.lang.String | no | Recipient, i.e., the seller's receiving address; text | {\"townName\":\"东华门街道\",\"townCode\":\"110101001\",\"cityCode\":\"110100\",\"contactorName\":\"234234\",\"provinceCode\":\"110000\",\"mobile\":\"12342344324\",\"addressId\":111028,\"areaCode\":\"110101\",\"cityName\":\"北京市\",\"post\":\"123123\",\"areaName\":\"东城区\",\"detailAddress\":\"接到\",\"provinceName\":\"北京\"} |
| `senderInfoStr` | java.lang.String | no | Sender, i.e., the buyer's door-to-door pickup address; text | {\"townName\":\"东华门街道\",\"townCode\":\"110101001\",\"cityCode\":\"110100\",\"contactorName\":\"234234\",\"provinceCode\":\"110000\",\"mobile\":\"12342344324\",\"addressId\":111028,\"areaCode\":\"110101\",\"cityName\":\"北京市\",\"post\":\"123123\",\"areaName\":\"东城区\",\"detailAddress\":\"接到\",\"provinceName\":\"北京\"} |
| `refundId` | java.lang.String | yes | Refund order ID | TQ284019625557092048 |
| `packageCount` | java.lang.Integer | yes | Number of packages | 1 |
| `totalVolume` | java.lang.String | no | Total volume, cm^3 | 1 |
| `totalWeight` | java.lang.String | no | Total weight, g | 1 |
| `couponId` | String | no | Coupon id |   |

<a id="m-alibaba-ocean-openplatform-biz-trade-param-addressparam"></a>
#### alibaba.ocean.openplatform.biz.trade.param.AddressParam

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `address` | java.lang.String | yes |  |    |
| `addressId` | java.lang.Long | yes |  |    |
| `areaCode` | java.lang.String | yes |  |    |
| `areaText` | java.lang.String | yes |  |    |
| `cityCode` | java.lang.String | yes |  |    |
| `cityText` | java.lang.String | yes |  |    |
| `districtCode` | java.lang.String | yes |  |     |
| `fullName` | java.lang.String | yes |  |    |
| `mobile` | java.lang.String | yes |  |    |
| `phone` | java.lang.String | yes |  |    |
| `postCode` | java.lang.String | yes |  |    |
| `provinceCode` | java.lang.String | yes |  |    |
| `provinceText` | java.lang.String | yes |  |    |
| `townCode` | java.lang.String | yes |  |    |
| `townText` | java.lang.String | yes |  |    |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.openapi.shared.common.refundofficialdelivery.solution.get.ResultModel](#m-alibaba-openapi-shared-common-refundofficialdelivery-solution-get-resultmodel) | yes |  |   |

<a id="m-alibaba-openapi-shared-common-refundofficialdelivery-solution-get-resultmodel"></a>
#### alibaba.openapi.shared.common.refundofficialdelivery.solution.get.ResultModel

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `success` | boolean | yes | Whether successful | true |
| `code` | java.lang.String | yes | Error code | "" |
| `message` | java.lang.String | yes | Error message | "" |
| `result` | [message:alibaba.ocean.openplatform.biz.logistics.result.OfficialDeliverySolutionResult](#m-alibaba-ocean-openplatform-biz-logistics-result-officialdeliverysolutionresult) | yes | Return value | {} |

<a id="m-alibaba-ocean-openplatform-biz-logistics-result-officialdeliverysolutionresult"></a>
#### alibaba.ocean.openplatform.biz.logistics.result.OfficialDeliverySolutionResult

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `height` | java.lang.Long | yes | Height cm | 3 |
| `length` | java.lang.Long | yes | Length cm | 3 |
| `offerInfoDTOList` | [message:alibaba.ocean.openplatform.biz.logistics.result.OfficialDeliveryOfferInfo[]](#m-alibaba-ocean-openplatform-biz-logistics-result-officialdeliveryofferinfo[]) | yes | Product information | "" |
| `orderEntryIdList` | java.lang.Long[] | yes | Sub-order | 87929773894 |
| `orderId` | java.lang.Long | yes | Sub-order | 87929773894 |
| `receiverInfoDTO` | [message:alibaba.ocean.openplatform.biz.trade.param.AddressParam](#m-alibaba-ocean-openplatform-biz-trade-param-addressparam) | yes | Recipient | {} |
| `sceneCode` | java.lang.String | yes | Scenario code | "" |
| `senderInfoDTO` | [message:alibaba.ocean.openplatform.biz.trade.param.AddressParam](#m-alibaba-ocean-openplatform-biz-trade-param-addressparam) | yes | Sender | {} |
| `signErrorMsg` | java.lang.String | yes | Error message | "" |
| `solutionDTOList` | [message:alibaba.ocean.openplatform.biz.logistics.result.OfficialDeliverySolution[]](#m-alibaba-ocean-openplatform-biz-logistics-result-officialdeliverysolution[]) | yes | List of solutions | {} |
| `totalVolume` | java.lang.Long | yes | Total volume, cm^3 | 12 |
| `totalWeight` | java.lang.Long | yes | Total weight, g | 12 |
| `userAlipaySignUrl` | java.lang.String | yes | User's Alipay sign address | http:// |
| `userSettled` | java.lang.Boolean | yes | Whether settled | true |
| `vipUser` | java.lang.Boolean | yes | Whether the user is a VIP | true |
| `width` | java.lang.Long | yes | Width | 2 |

<a id="m-alibaba-ocean-openplatform-biz-logistics-result-officialdeliveryofferinfo[]"></a>
#### alibaba.ocean.openplatform.biz.logistics.result.OfficialDeliveryOfferInfo[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `offerId` | java.lang.Long | yes | Product ID | 1 |
| `offerName` | java.lang.String | yes | Product name | "xxx商品" |

<a id="m-alibaba-ocean-openplatform-biz-trade-param-addressparam"></a>
#### alibaba.ocean.openplatform.biz.trade.param.AddressParam

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `address` | java.lang.String | yes |  |  |
| `addressId` | java.lang.Long | yes |  |  |
| `areaCode` | java.lang.String | yes |  |  |
| `areaText` | java.lang.String | yes |  |  |
| `cityCode` | java.lang.String | yes |  |  |
| `cityText` | java.lang.String | yes |  |  |
| `districtCode` | java.lang.String | yes |  |  |
| `fullName` | java.lang.String | yes |  |  |
| `mobile` | java.lang.String | yes |  |  |
| `phone` | java.lang.String | yes |  |  |
| `postCode` | java.lang.String | yes |  |  |
| `provinceCode` | java.lang.String | yes |  |  |
| `provinceText` | java.lang.String | yes |  |  |
| `townCode` | java.lang.String | yes |  |  |
| `townText` | java.lang.String | yes |  |  |

<a id="m-alibaba-ocean-openplatform-biz-logistics-result-officialdeliverysolution[]"></a>
#### alibaba.ocean.openplatform.biz.logistics.result.OfficialDeliverySolution[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `appointmentTimeDTO` | [message:alibaba.ocean.openplatform.biz.logistics.result.OfficialDeliverySolutionDTO.OfficialAppointmentTime](#m-alibaba-ocean-openplatform-biz-logistics-result-officialdeliverysolutiondto-officialappointmenttime) | yes | Scheduled delivery time | "" |
| `chargeDetailDTO` | [message:alibaba.ocean.openplatform.biz.logistics.result.OfficialDeliverySolutionDTO.OfficialChargeDetail](#m-alibaba-ocean-openplatform-biz-logistics-result-officialdeliverysolutiondto-officialchargedetail) | yes | Billing details | "" |
| `cpCode` | java.lang.String | yes | Logistics service provider code | "YTO" |
| `cpName` | java.lang.String | yes | Logistics service provider name | "圆通速递" |
| `deliveryTimeDTO` | [message:alibaba.ocean.openplatform.biz.logistics.result.OfficialDeliverySolutionDTO.OfficialDeliveryTime](#m-alibaba-ocean-openplatform-biz-logistics-result-officialdeliverysolutiondto-officialdeliverytime) | yes | Delivery time | "" |
| `priority` | java.lang.Integer | yes | Priority; the smaller the value, the higher the priority | 1 |
| `remark` | java.lang.String | yes | Remark | "圆通快递价格优惠覆盖广" |
| `serviceType` | java.lang.String | yes | Solution type | "1" |
| `serviceTypeDesc` | java.lang.String | yes | Solution type description | "圆通快递当日取" |
| `serviceTypeName` | java.lang.String | yes | Solution type name | "当日上门" |
| `solutionCode` | java.lang.String | yes | Logistics solution code | "422" |
| `solutionName` | java.lang.String | yes | Logistics solution name | "圆通快递当日取" |

<a id="m-alibaba-ocean-openplatform-biz-logistics-result-officialdeliverysolutiondto-officialappointmenttime"></a>
#### alibaba.ocean.openplatform.biz.logistics.result.OfficialDeliverySolutionDTO.OfficialAppointmentTime

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `timeSlotDTOList` | [message:alibaba.ocean.openplatform.biz.logistics.result.OfficialDeliverySolutionDTO.OfficialAppointmentTimeSlot[]](#m-alibaba-ocean-openplatform-biz-logistics-result-officialdeliverysolutiondto-officialappointmenttimeslot[]) | yes | List of time periods | xx |

<a id="m-alibaba-ocean-openplatform-biz-logistics-result-officialdeliverysolutiondto-officialappointmenttimeslot[]"></a>
#### alibaba.ocean.openplatform.biz.logistics.result.OfficialDeliverySolutionDTO.OfficialAppointmentTimeSlot[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `dateStr` | java.lang.String | yes | Date | "2025-06-18" |
| `gmtEnd` | java.lang.String | yes | End time | "17:00" |
| `gmtStart` | java.lang.String | yes | Start time | "15:00" |

<a id="m-alibaba-ocean-openplatform-biz-logistics-result-officialdeliverysolutiondto-officialchargedetail"></a>
#### alibaba.ocean.openplatform.biz.logistics.result.OfficialDeliverySolutionDTO.OfficialChargeDetail

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `deliveryFee` | java.lang.Long | yes | Delivery fee, in cents (fen) | 1000 |
| `deliveryTotalFee` | java.lang.Long | yes | Shipping fee including delivery fee, out-of-zone fee, etc., in cents | 1000 |
| `firstFee` | java.lang.Long | yes | First weight fee, in cents (fen) | 1000 |
| `firstTotalCost` | java.lang.Long | yes | Total fee for the first weight unit, in cents (fen) | 1000 |
| `firstUnit` | java.lang.Long | yes | First weight unit, in g | 1000 |
| `nextFee` | java.lang.Long | yes | Additional weight fee, in cents (fen) | 1000 |
| `nextTotalCost` | java.lang.Long | yes | Total cost of additional weight, in cents | 1000 |
| `nextUnit` | java.lang.Long | yes | Additional weight unit, in g | 1000 |
| `originTotalCost` | java.lang.Long | yes | Original total cost (non-member price), in cents | 1200 |
| `totalCost` | java.lang.Long | yes | Total cost = shipping fee + value-added service fee - red packet amount, in cents | 1000 |
| `vasDetailDTOList` | [message:alibaba.ocean.openplatform.biz.logistics.result.OfficialDeliverySolutionDTO.OfficialChargeVasDetail[]](#m-alibaba-ocean-openplatform-biz-logistics-result-officialdeliverysolutiondto-officialchargevasdetail[]) | yes | List of value-added service fees | "" |
| `vasTotalFee` | java.lang.Long | yes | Total value-added service fee, in cents (fen) | 1000 |
| `officialCouponDTO` | [message:alibaba.ocean.openplatform.biz.logistics.result.OfficialDeliverySolutionDTO.OfficialCoupon](#m-alibaba-ocean-openplatform-biz-logistics-result-officialdeliverysolutiondto-officialcoupon) | yes | Coupon |   |
| `fromOverAreaFee` | Long | yes | Out-of-area pickup fee, in cents (fen) |   |
| `toOverAreaFee` | Long | yes | Out-of-area sign-for fee, in cents (fen) |   |
| `overDistanceFee` | Long | yes | Oversize fee, in cents (fen) |   |

<a id="m-alibaba-ocean-openplatform-biz-logistics-result-officialdeliverysolutiondto-officialchargevasdetail[]"></a>
#### alibaba.ocean.openplatform.biz.logistics.result.OfficialDeliverySolutionDTO.OfficialChargeVasDetail[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `cost` | java.lang.Long | yes | Value-added service fee, in cents (fen) | 1000 |
| `dialogTextContent` | java.lang.String | yes | Value-added service popup content | xx |
| `dialogTitle` | java.lang.String | yes | Value-added service popup title | xx |
| `dialogUrl` | java.lang.String | yes | Value-added service popup link | xx |
| `operateType` | java.lang.String | yes | Operation type | UN_SELECTABLE：不可操作 |
| `selectedValue` | java.lang.String | yes | Value-added service selection status | xx |
| `vasCode` | java.lang.String | yes | Value-added service code | xx |
| `vasDesc` | java.lang.String | yes | Value-added service description | xx |
| `vasName` | java.lang.String | yes | Value-added service name | xx |

<a id="m-alibaba-ocean-openplatform-biz-logistics-result-officialdeliverysolutiondto-officialcoupon"></a>
#### alibaba.ocean.openplatform.biz.logistics.result.OfficialDeliverySolutionDTO.OfficialCoupon

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `id` | String | yes | Coupon id |   |
| `name` | String | yes | Name |   |
| `description` | String | yes | Description |   |
| `amount` | Long | yes | Coupon amount (in cents/fen) |   |
| `couponType` | String | yes | Coupon type |   |
| `status` | String | yes | Coupon status |   |
| `gmtEffect` | Date | yes | Coupon effective time |   |
| `gmtExpired` | Date | yes | Coupon expiration time |   |

<a id="m-alibaba-ocean-openplatform-biz-logistics-result-officialdeliverysolutiondto-officialdeliverytime"></a>
#### alibaba.ocean.openplatform.biz.logistics.result.OfficialDeliverySolutionDTO.OfficialDeliveryTime

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `deliveryText` | java.lang.String | yes | x | x |
| `deliveryTime` | java.lang.String | yes | x | x |
| `remark` | java.lang.String | yes | x | x |
