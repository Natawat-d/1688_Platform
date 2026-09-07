# 买家申请修改收货地址

API: `com.alibaba.trade:order.receiveAddress.buyerUpdate:1` · Category: 订单  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:order.receiveAddress.buyerUpdate-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/order.receiveAddress.buyerUpdate/{appKey}`  
需要授权 (access_token) · 需要签名

买家修改收货地址,当修改的地址为偏远地区或者需要重新计算运费时，需要卖家确认，卖家可拒绝修改地址。

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `param` | [message:alibaba.ocean.openplatform.biz.trade.param.BuyerDeliveryAddressModifyParam](#m-alibaba-ocean-openplatform-biz-trade-param-buyerdeliveryaddressmodifyparam) | 是 | 修改地址信息对象 | {} |

<a id="m-alibaba-ocean-openplatform-biz-trade-param-buyerdeliveryaddressmodifyparam"></a>
#### alibaba.ocean.openplatform.biz.trade.param.BuyerDeliveryAddressModifyParam

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `orderId` | java.lang.Long | 是 | 订单id | 1231231123123 |
| `receiveAddress` | [message:alibaba.ocean.openplatform.biz.trade.param.AddressParam](#m-alibaba-ocean-openplatform-biz-trade-param-addressparam) | 是 | 修改地址 | {} |

<a id="m-alibaba-ocean-openplatform-biz-trade-param-addressparam"></a>
#### alibaba.ocean.openplatform.biz.trade.param.AddressParam

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `addressId` | java.lang.Long | 是 |  |  |
| `fullName` | java.lang.String | 是 |  |  |
| `mobile` | java.lang.String | 是 |  |  |
| `phone` | java.lang.String | 是 |  |  |
| `postCode` | java.lang.String | 是 |  |  |
| `cityText` | java.lang.String | 是 |  |  |
| `cityCode` | java.lang.String | 是 |  |  |
| `provinceText` | java.lang.String | 是 |  |  |
| `provinceCode` | java.lang.String | 是 |  |  |
| `areaText` | java.lang.String | 是 |  |  |
| `areaCode` | java.lang.String | 是 |  |  |
| `townText` | java.lang.String | 是 |  |  |
| `townCode` | java.lang.String | 是 |  |  |
| `address` | java.lang.String | 是 |  |  |
| `districtCode` | java.lang.String | 是 |  |  |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:alibaba.openapi.shared.common.order.receiveAddress.buyerUpdate.ResultModel](#m-alibaba-openapi-shared-common-order-receiveaddress-buyerupdate-resultmodel) | 是 | 返回值 | {} |

<a id="m-alibaba-openapi-shared-common-order-receiveaddress-buyerupdate-resultmodel"></a>
#### alibaba.openapi.shared.common.order.receiveAddress.buyerUpdate.ResultModel

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `success` | boolean | 是 | 网关接口是否成功 | true |
| `code` | java.lang.String | 是 | 错误码 | "" |
| `message` | java.lang.String | 是 | 错误信息 | "" |
| `result` | Boolean | 是 | 申请修改地址是否成功 | true |
