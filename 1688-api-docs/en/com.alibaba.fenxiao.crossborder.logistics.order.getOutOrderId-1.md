# Query external order ID by waybill number or unclaimed-parcel code

Original name: 根据运单号或无主件码查询外部订单ID  
API: `com.alibaba.fenxiao.crossborder:logistics.order.getOutOrderId:1` · Category: Logistics  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:logistics.order.getOutOrderId-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/logistics.order.getOutOrderId/{appKey}`  
Requires user authorization (access_token) · Requires signature

Query the external order ID by waybill number or unclaimed-parcel code.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `shipmentId` | java.lang.String | no | Waybill number | 1 |
| `noMainPartCode` | java.lang.String | no | Unclaimed item code | 1 |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:logistics.order.getOutOrderId.model.ResultModel](#m-logistics-order-getoutorderid-model-resultmodel) | yes | Result | object |

<a id="m-logistics-order-getoutorderid-model-resultmodel"></a>
#### logistics.order.getOutOrderId.model.ResultModel

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `success` | java.lang.Boolean | yes | Success | true |
| `msg` | java.lang.String | yes | Info | 成功 |
| `model` | [message:alibaba.cbu.sc.label.model.ScOutOrderIdQueryModel](#m-alibaba-cbu-sc-label-model-scoutorderidquerymodel) | yes | Return value | object |

<a id="m-alibaba-cbu-sc-label-model-scoutorderidquerymodel"></a>
#### alibaba.cbu.sc.label.model.ScOutOrderIdQueryModel

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `outOrderId` | java.lang.String | yes | External order number | 1 |
| `orderId` | java.lang.String | yes | Order number | 1 |

## Samples

**Introduction to the warehouse label printing solution**

```
功能介绍：当贴有运单号的1688包裹到仓库时，客户需要贴上你的箱唛和货品标签运到海外。 

我们提供了两种能力 ：

1. 基础版：扫描运单号，可以查询你订单号，你通过订单号打印箱唛和货品标签。本api可以帮助你通过运单号或者无主件码查出你平台订单号和1688订单号。

需要你做的事情：
在下单的时候传两个参数：outOrderId = 你平台订单号  和 dropshipping = y


2. 升级版：在基础版基础上，提前与你平台打印箱唛和货品标签接口对接，提供操作界面连接你仓库打印机。

打印标签的网页地址：https://air.1688.com/app/channel-fe/chain-work/printwaybill.html

需要你平台做的事情：在基础版的基础上，额外提供通过你平台订单号查询箱唛和货品标签的接口，返回是pdf。 1688侧接入开发。升级版方案需要联系开发同学。
```
