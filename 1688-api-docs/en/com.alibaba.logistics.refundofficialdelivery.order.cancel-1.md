# Cancel official return-pickup logistics order

Original name: 取消退货官方物流上门揽订单  
API: `com.alibaba.logistics:refundofficialdelivery.order.cancel:1` · Category: Official Return Pickup  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.logistics:refundofficialdelivery.order.cancel-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.logistics/refundofficialdelivery.order.cancel/{appKey}`  
Requires user authorization (access_token) · Requires signature

Cancel an official-logistics door-to-door pickup order for a return.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `param` | [message:alibaba.ocean.openplatform.biz.logistics.param.OfficialDeliveryOrderCancelParam](#m-alibaba-ocean-openplatform-biz-logistics-param-officialdeliveryordercancelparam) | yes |  |  |

<a id="m-alibaba-ocean-openplatform-biz-logistics-param-officialdeliveryordercancelparam"></a>
#### alibaba.ocean.openplatform.biz.logistics.param.OfficialDeliveryOrderCancelParam

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `cancelText` | java.lang.String | yes | Cancellation reason | 计划有变，暂时不需要寄了 |
| `cancelType` | java.lang.String | yes | Cancellation reason type: INFORMATION_INCORRECT(&quot;Information entered incorrectly (time/address etc. needs to be changed)&quot;), WANT_TO_SEND_MYSELF(&quot;Want to go to a nearby service point to send it myself&quot;), PLAN_CHANGED(&quot;Plans changed, no need to ship for now&quot;), WANT_TO_CHANGE_PICKUP_TIME(&quot;Want to change the door-to-door pickup time&quot;), PRICE_TOO_HIGH(&quot;I think the price is a bit high&quot;), 	ATE(&quot;Courier did not arrive for pickup on time&quot;), COURIER_NOT_COMING(&quot;Courier does not come for pickup&quot;), COURIER_BAD_ATTITUDE(&quot;Courier's service attitude is poor&quot;), ITEM_CANNOT_BE_SHIPPED(&quot;Item type cannot be shipped&quot;), COURIER_TOO_BUSY(&quot;Courier reported unable to pick up due to tight capacity&quot;), OTHER(&quot;Other&quot;), | PLAN_CHANGED |
| `officialDeliveryOrderId` | java.lang.Long | yes | Door-to-door pickup order ID | 7210001 |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.openapi.shared.common.refundofficialdelivery.order.cancel.ResultModel](#m-alibaba-openapi-shared-common-refundofficialdelivery-order-cancel-resultmodel) | yes |  |  |

<a id="m-alibaba-openapi-shared-common-refundofficialdelivery-order-cancel-resultmodel"></a>
#### alibaba.openapi.shared.common.refundofficialdelivery.order.cancel.ResultModel

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `success` | boolean | yes | Whether successful | true |
| `code` | java.lang.String | yes | Error code | "" |
| `message` | java.lang.String | yes | Error message | "" |
| `result` | java.lang.Boolean | yes | Whether the order cancellation was successful | true |
