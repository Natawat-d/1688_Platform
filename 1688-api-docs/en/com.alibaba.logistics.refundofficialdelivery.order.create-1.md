# Create official return-pickup logistics order

Original name: 创建退货官方物流上门揽订单  
API: `com.alibaba.logistics:refundofficialdelivery.order.create:1` · Category: Official Return Pickup  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.logistics:refundofficialdelivery.order.create-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.logistics/refundofficialdelivery.order.create/{appKey}`  
Requires user authorization (access_token) · Requires signature

Create an official-logistics door-to-door pickup order for a return.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `param` | [message:alibaba.ocean.openplatform.biz.logistics.param.OfficialDeliveryOrderCreateParam](#m-alibaba-ocean-openplatform-biz-logistics-param-officialdeliveryordercreateparam) | yes |  |  |

<a id="m-alibaba-ocean-openplatform-biz-logistics-param-officialdeliveryordercreateparam"></a>
#### alibaba.ocean.openplatform.biz.logistics.param.OfficialDeliveryOrderCreateParam

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `couponId` | java.lang.String | no | Coupon id | 1234123 |
| `dateStr` | java.lang.String | yes | Door-to-door pickup time, format yyyy-MM-dd | 2024-07-11 |
| `gmtAcceptStart` | java.lang.String | yes | Door-to-door pickup start time, HH:mm | 9:00 |
| `gmtAcceptEnd` | java.lang.String | yes | Door-to-door pickup end time, HH:mm | 21:00 |
| `goodsType` | java.lang.String | yes | Product type (&quot;DAILY_NECESSITIES&quot;, &quot;Daily necessities&quot;); (&quot;FOOD&quot;, &quot;Food&quot;); (&quot;FURNITURE&quot;, &quot;Furniture&quot;); (&quot;METALS&quot;, &quot;Hardware&quot;); (&quot;COSMETICS&quot;, &quot;Cosmetics&quot;); (&quot;DIGITAL&quot;, &quot;Digital&quot;); (&quot;GIFT&quot;, &quot;Gift&quot;); (&quot;OTHER&quot;, &quot;Other&quot;); | DAILY_NECESSITIES |
| `packageCount` | java.lang.String | yes | Number of packages | 1 |
| `receiverInfo` | [message:alibaba.ocean.openplatform.biz.trade.param.AddressParam](#m-alibaba-ocean-openplatform-biz-trade-param-addressparam) | no | Recipient info | {} |
| `receiverInfoStr` | java.lang.String | no | Recipient info text | "" |
| `refundId` | java.lang.String | yes | Refund order ID | TQ284160529017092048 |
| `senderInfo` | [message:alibaba.ocean.openplatform.biz.trade.param.AddressParam](#m-alibaba-ocean-openplatform-biz-trade-param-addressparam) | no | Sender info | {} |
| `senderInfoStr` | java.lang.String | no | Sender text | "" |
| `solutionCode` | java.lang.String | yes | Door-to-door pickup scheme code | 434 |
| `totalVolume` | java.lang.String | no | Total volume, unit: cm^3 | 10 |
| `totalWeight` | java.lang.String | no | Total weight, in g | 1000 |

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

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.openapi.shared.common.refundofficialdelivery.order.create.ResultModel](#m-alibaba-openapi-shared-common-refundofficialdelivery-order-create-resultmodel) | yes |  |  |

<a id="m-alibaba-openapi-shared-common-refundofficialdelivery-order-create-resultmodel"></a>
#### alibaba.openapi.shared.common.refundofficialdelivery.order.create.ResultModel

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `success` | boolean | yes | Whether successful | true |
| `code` | java.lang.String | yes | Error code | "" |
| `message` | java.lang.String | yes | Error message | "" |
| `result` | Long | yes | Returned order ID | 19882 |
