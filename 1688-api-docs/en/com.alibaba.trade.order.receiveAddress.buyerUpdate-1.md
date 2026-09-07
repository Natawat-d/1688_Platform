# Buyer requests a shipping-address change

Original name: 买家申请修改收货地址  
API: `com.alibaba.trade:order.receiveAddress.buyerUpdate:1` · Category: Orders  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:order.receiveAddress.buyerUpdate-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/order.receiveAddress.buyerUpdate/{appKey}`  
Requires user authorization (access_token) · Requires signature

Buyer changes the shipping address. If the new address is in a remote area or the shipping fee must be recalculated, the seller must confirm the change and may reject it.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `param` | [message:alibaba.ocean.openplatform.biz.trade.param.BuyerDeliveryAddressModifyParam](#m-alibaba-ocean-openplatform-biz-trade-param-buyerdeliveryaddressmodifyparam) | yes | Modify address info object | {} |

<a id="m-alibaba-ocean-openplatform-biz-trade-param-buyerdeliveryaddressmodifyparam"></a>
#### alibaba.ocean.openplatform.biz.trade.param.BuyerDeliveryAddressModifyParam

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `orderId` | java.lang.Long | yes | Order id | 1231231123123 |
| `receiveAddress` | [message:alibaba.ocean.openplatform.biz.trade.param.AddressParam](#m-alibaba-ocean-openplatform-biz-trade-param-addressparam) | yes | Modify address | {} |

<a id="m-alibaba-ocean-openplatform-biz-trade-param-addressparam"></a>
#### alibaba.ocean.openplatform.biz.trade.param.AddressParam

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `addressId` | java.lang.Long | yes |  |  |
| `fullName` | java.lang.String | yes |  |  |
| `mobile` | java.lang.String | yes |  |  |
| `phone` | java.lang.String | yes |  |  |
| `postCode` | java.lang.String | yes |  |  |
| `cityText` | java.lang.String | yes |  |  |
| `cityCode` | java.lang.String | yes |  |  |
| `provinceText` | java.lang.String | yes |  |  |
| `provinceCode` | java.lang.String | yes |  |  |
| `areaText` | java.lang.String | yes |  |  |
| `areaCode` | java.lang.String | yes |  |  |
| `townText` | java.lang.String | yes |  |  |
| `townCode` | java.lang.String | yes |  |  |
| `address` | java.lang.String | yes |  |  |
| `districtCode` | java.lang.String | yes |  |  |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.openapi.shared.common.order.receiveAddress.buyerUpdate.ResultModel](#m-alibaba-openapi-shared-common-order-receiveaddress-buyerupdate-resultmodel) | yes | Return value | {} |

<a id="m-alibaba-openapi-shared-common-order-receiveaddress-buyerupdate-resultmodel"></a>
#### alibaba.openapi.shared.common.order.receiveAddress.buyerUpdate.ResultModel

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `success` | boolean | yes | Whether the gateway interface succeeded | true |
| `code` | java.lang.String | yes | Error code | "" |
| `message` | java.lang.String | yes | Error message | "" |
| `result` | Boolean | yes | Whether the address modification request was successful | true |
