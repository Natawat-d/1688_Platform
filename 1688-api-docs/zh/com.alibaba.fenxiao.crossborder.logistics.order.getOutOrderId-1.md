# 根据运单号或无主件码查询外部订单ID

API: `com.alibaba.fenxiao.crossborder:logistics.order.getOutOrderId:1` · Category: 物流  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:logistics.order.getOutOrderId-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/logistics.order.getOutOrderId/{appKey}`  
需要授权 (access_token) · 需要签名

根据运单号或无主件码查询外部订单ID

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `shipmentId` | java.lang.String | 否 | 运单号 | 1 |
| `noMainPartCode` | java.lang.String | 否 | 无主件码 | 1 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:logistics.order.getOutOrderId.model.ResultModel](#m-logistics-order-getoutorderid-model-resultmodel) | 是 | 结果 | object |

<a id="m-logistics-order-getoutorderid-model-resultmodel"></a>
#### logistics.order.getOutOrderId.model.ResultModel

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `success` | java.lang.Boolean | 是 | 成功 | true |
| `msg` | java.lang.String | 是 | 信息 | 成功 |
| `model` | [message:alibaba.cbu.sc.label.model.ScOutOrderIdQueryModel](#m-alibaba-cbu-sc-label-model-scoutorderidquerymodel) | 是 | 返回值 | object |

<a id="m-alibaba-cbu-sc-label-model-scoutorderidquerymodel"></a>
#### alibaba.cbu.sc.label.model.ScOutOrderIdQueryModel

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `outOrderId` | java.lang.String | 是 | 外部订单号 | 1 |
| `orderId` | java.lang.String | 是 | 订单号 | 1 |

## 示例

**仓库打印标签解决方案介绍**

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
