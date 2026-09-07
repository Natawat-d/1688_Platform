# 查询退款退货原因（用于创建退款退货）

API: `com.alibaba.trade:alibaba.trade.getRefundReasonList:1` · Category: 退货退款  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.trade.getRefundReasonList-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.trade.getRefundReasonList/{appKey}`  
需要授权 (access_token) · 需要签名

查询退款退货原因（用于创建退款退货）

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `orderId` | Long | 是 | 主订单id |   |
| `orderEntryIds` | Long[] | 是 | 子订单id |   |
| `goodsStatus` | String | 是 | 货物状态 | 售中等待买家发货:”refundWaitSellerSend"; 售中等待买家收货:"refundWaitBuyerReceive"; 售中已收货（未确认完成交易）:"refundBuyerReceived" 售后未收货:"aftersaleBuyerNotReceived"; 售后已收到货:"aftersaleBuyerReceived" |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:alibaba.ocean.openplatform.common.OrderRefundReasonListResult](#m-alibaba-ocean-openplatform-common-orderrefundreasonlistresult) | 是 | 返回结果 |   |

<a id="m-alibaba-ocean-openplatform-common-orderrefundreasonlistresult"></a>
#### alibaba.ocean.openplatform.common.OrderRefundReasonListResult

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `code` | String | 是 | 错误码 |   |
| `message` | String | 是 | 错误信息 |   |
| `result` | [message:alibaba.ocean.openplatform.biz.trade.result.OrderRefundReasonListResult](#m-alibaba-ocean-openplatform-biz-trade-result-orderrefundreasonlistresult) | 是 | 结果 |   |
| `success` | Boolean | 是 | 是否成功 |   |

<a id="m-alibaba-ocean-openplatform-biz-trade-result-orderrefundreasonlistresult"></a>
#### alibaba.ocean.openplatform.biz.trade.result.OrderRefundReasonListResult

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `reasons` | [message:alibaba.ocean.openplatform.biz.trade.result.OrderRefundReasonModel[]](#m-alibaba-ocean-openplatform-biz-trade-result-orderrefundreasonmodel[]) | 是 | 原因列表 |   |

<a id="m-alibaba-ocean-openplatform-biz-trade-result-orderrefundreasonmodel[]"></a>
#### alibaba.ocean.openplatform.biz.trade.result.OrderRefundReasonModel[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `id` | java.lang.Long | 是 | 原因id |   |
| `name` | java.lang.String | 是 | 原因 |   |
| `needVoucher` | java.lang.Boolean | 是 | 凭证是否必须上传 | "true"表示必须要上传凭证 |
| `noRefundCarriage` | java.lang.Boolean | 是 | 是否支持退运费 | “true" 表示不支持退运费 |
| `tip` | java.lang.String | 是 | 提示 |   |

## 错误码

| 错误码 | 现象 | 解决方案 |
|---|---|---|
| 5xxx | 调用者错误 | 请确认输入参数是否合法。<br>货物状态参数示例，子订单需要数组输入... |
| 4xxx | 服务方内部错误 |  |

## 示例

**输入示例**

```
orderId:586683458994743215
orderEntryIds:[586683458997743215]
goodsStatus:aftersaleBuyerNotReceived

```
