# Get official return-pickup logistics order details

Original name: 退货官方物流上门揽订单详情获取  
API: `com.alibaba.logistics:refundofficialdelivery.order.get:1` · Category: Official Return Pickup  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.logistics:refundofficialdelivery.order.get-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.logistics/refundofficialdelivery.order.get/{appKey}`  
Requires user authorization (access_token) · Requires signature

Get the details of an official-logistics door-to-door pickup order for a return.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `officialDeliveryOrderId` | String | yes | Official logistics door-to-door pickup order ID | 1988772 |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.openapi.shared.common.refundofficialdelivery.order.get.ResultModel](#m-alibaba-openapi-shared-common-refundofficialdelivery-order-get-resultmodel) | yes | Return value | {} |

<a id="m-alibaba-openapi-shared-common-refundofficialdelivery-order-get-resultmodel"></a>
#### alibaba.openapi.shared.common.refundofficialdelivery.order.get.ResultModel

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `success` | boolean | yes | Whether successful | true |
| `code` | java.lang.String | yes | Error code | "" |
| `message` | java.lang.String | yes | Error message | "" |
| `result` | [message:alibaba.ocean.openplatform.biz.logistics.result.OfficialDeliveryOrderInfo](#m-alibaba-ocean-openplatform-biz-logistics-result-officialdeliveryorderinfo) | yes | Order details | {} |

<a id="m-alibaba-ocean-openplatform-biz-logistics-result-officialdeliveryorderinfo"></a>
#### alibaba.ocean.openplatform.biz.logistics.result.OfficialDeliveryOrderInfo

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `bizId` | java.lang.String | yes | Refund order ID | TQ284160529017092048 |
| `bizType` | java.lang.String | yes | Business type | buyer_refund |
| `dateStr` | java.lang.String | yes | Door-to-door pickup date | 2024-07-11 |
| `features` | java.util.Map | yes | Order features | {} |
| `gmtAcceptStart` | java.lang.String | yes | Door-to-door pickup time start | 2024-07-11 10:00 |
| `gmtAcceptEnd` | java.lang.String | yes | Door-to-door pickup time end | 2024-07-11 22:00 |
| `gmtModified` | java.util.Date | yes | Modification time | 20240710173724000+0800 |
| `id` | java.lang.Long | yes | Door-to-door pickup order ID | 7210001 |
| `lastTraceDetail` | java.lang.String | yes | Latest logistics tracking detail | "" |
| `lastTraceStatus` | java.lang.String | yes | Latest logistics tracking status | "" |
| `lastTraceStatusName` | java.lang.String | yes | Name of the latest logistics tracking status | "" |
| `logisticsStatus` | java.lang.String | yes | Logistics status | "" |
| `mailNo` | java.lang.String | yes | Logistics waybill number | sp981920934 |
| `orderId` | java.lang.String | yes | Order id | "" |
| `outBizId` | java.lang.String | yes | External business id | "" |
| `payRecordDTO` | [message:alibaba.ocean.openplatform.biz.logistics.result.OfficialDeliveryOrderInfo.OfficialPayRecordDTO](#m-alibaba-ocean-openplatform-biz-logistics-result-officialdeliveryorderinfo-officialpayrecorddto) | yes | Payment information | {} |
| `receiverInfo` | [message:alibaba.ocean.openplatform.biz.trade.param.AddressParam](#m-alibaba-ocean-openplatform-biz-trade-param-addressparam) | yes | Recipient address information | {} |
| `refundId` | java.lang.String | yes | Refund order ID | TQ284160529017092048 |
| `senderInfo` | [message:alibaba.ocean.openplatform.biz.trade.param.AddressParam](#m-alibaba-ocean-openplatform-biz-trade-param-addressparam) | yes | Shipping address information | {} |
| `signOfficial` | java.lang.Boolean | yes | Whether signed | false |
| `solutionCode` | java.lang.String | yes | Door-to-door pickup scheme code | 434 |
| `status` | java.lang.String | yes | Order status | activated |

<a id="m-alibaba-ocean-openplatform-biz-logistics-result-officialdeliveryorderinfo-officialpayrecorddto"></a>
#### alibaba.ocean.openplatform.biz.logistics.result.OfficialDeliveryOrderInfo.OfficialPayRecordDTO

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `amount` | java.lang.Long | yes | Quantity | 1 |
| `bizId` | java.lang.String | yes | Refund order ID | TQ284160529017092048 |
| `bizType` | java.lang.String | yes | Business type | buyer_refund |
| `failReason` | java.lang.String | yes | Error reason | "" |
| `gmtCreate` | java.util.Date | yes | Creation time | 20240710173724000+0800 |
| `gmtModified` | java.util.Date | yes | Modification time | 20240710173724000+0800 |
| `id` | java.lang.Long | yes | Door-to-door pickup order id | 7210001 |
| `outBizId` | java.lang.String | yes | External business id | "" |
| `payChannel` | java.lang.String | yes | Payment channel | "" |
| `payType` | java.lang.String | yes | Payment type | "" |
| `status` | java.lang.String | yes | Status | activated |
| `feature` | [message:alibaba.ocean.openplatform.biz.logistics.result.OfficialDeliveryOrderInfo.FeatureDTO](#m-alibaba-ocean-openplatform-biz-logistics-result-officialdeliveryorderinfo-featuredto) | yes |  |   |

<a id="m-alibaba-ocean-openplatform-biz-logistics-result-officialdeliveryorderinfo-featuredto"></a>
#### alibaba.ocean.openplatform.biz.logistics.result.OfficialDeliveryOrderInfo.FeatureDTO

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `firstWeightUnit` | String | yes | First-weight unit, unit: kg |   |
| `firstWeightPrice` | String | yes | First-weight price, in yuan |   |
| `bufferedFirstWeightPrice` | String | yes | First-weight price after adding buffer, unit: yuan |   |
| `nextWeightUnit` | String | yes | Additional weight unit, in kg |   |
| `nextWeightPrice` | String | yes | Additional weight price, in yuan |   |
| `fromOverAreaFee` | String | yes | Pickup out-of-zone fee, in yuan |   |
| `toOverAreaFee` | String | yes | Delivery out-of-zone fee, in yuan |   |
| `overDistanceFee` | String | yes | Extra-length fee, in yuan |   |
| `deliveryTotalFee` | String | yes | Total forward delivery fee |   |
| `nextTotalFee` | String | yes | Total additional weight fee, unit: yuan |   |
| `firstTotalFee` | String | yes | First-weight total fee, in yuan |   |
| `totalFee` | String | yes | Total fee, unit: yuan |   |
| `chargeType` | String | yes | Billing method, by weight / by volume |   |
| `chargeWeight` | String | yes | Chargeable weight, in kg |   |
| `chargeVolume` | String | yes | Chargeable volume, in cubic meters |   |
| `bubbleWeight` | String | yes | Volumetric weight unit kg -- returned when the chargeable weight is greater than the actual product weight; otherwise not returned |   |
| `isBubble` | String | yes | Whether volumetric weight is counted |   |
| `originTotalFee` | String | yes | Original total fee, amount calculated by Cainiao, in yuan |   |
| `refundId` | String | yes | Refund order number |   |
| `couponId` | String | yes | Coupon id |   |
| `couponFee` | String | yes | Coupon amount, in yuan |   |
| `couponUsedFee` | String | yes | Coupon redemption amount, in yuan |   |
| `couponName` | String | yes | Coupon model |   |
| `couponDesc` | String | yes | Coupon description |   |
| `couponSource` | String | yes | Coupon source |   |
| `vasList` | String | yes | List of value-added service fees |   |
| `mailNo` | String | yes | Waybill number |   |
| `settlementList` | String | yes | Settlement information |   |
| `orderSourceType` | String | yes | Order source |   |

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
