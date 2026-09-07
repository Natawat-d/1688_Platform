# 创建退款退货申请

API: `com.alibaba.trade:alibaba.trade.createRefund:1` · Category: 退货退款  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.trade.createRefund-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.trade.createRefund/{appKey}`  
需要授权 (access_token) · 需要签名

创建退款退货申请

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `orderId` | Long | 是 | 主订单 |    |
| `orderEntryIds` | Long[] | 是 | 子订单 |   |
| `disputeRequest` | String | 是 | 退款/退款退货。只有已收到货，才可以选择退款退货。 | 退款:"refund"; 退款退货:"returnRefund" |
| `applyPayment` | Long | 是 | 退款金额（单位：分）。不大于实际付款金额；等待卖家发货时，必须为商品的实际付款金额。 |   |
| `applyCarriage` | Long | 是 | 退运费金额（单位：分）。 |   |
| `applyReasonId` | Long | 是 | 退款原因id（从API getRefundReasonList获取） |   |
| `description` | String | 是 | 退款申请理由，2-150字 |   |
| `goodsStatus` | String | 是 | 货物状态 |  售中等待卖家发货:"refundWaitSellerSend"; 售中等待买家收货:"refundWaitBuyerReceive"; 售中已收货（未确认完成交易）:"refundBuyerReceived" 售后未收货:"aftersaleBuyerNotReceived"; 售后已收到货:"aftersaleBuyerReceived" |
| `vouchers` | String[] | 否 | 凭证图片URLs。1-5张，必须使用API uploadRefundVoucher返回的“图片域名/相对路径” |  [https://cbu01.alicdn.com/img/ibank/2019/901/930/11848039109.jpg] |
| `orderEntryCountList` | [message:alibaba.ocean.openplatform.biz.trade.common.model.OrderEntryCountModel[]](#m-alibaba-ocean-openplatform-biz-trade-common-model-orderentrycountmodel[]) | 否 | 子订单退款数量。仅在售中买家已收货（退款退货）时，可指定退货数量；默认，全部退货。 |  [{"id":586683458996743215,"count":1}] |
| `customRefund` | Boolean | 否 | 是否定制退货运费 | true：订单退款运费按照isv回传为准，不做优化；false（默认）：订单退款金额按照平台规则为准，需要做优化； |
| `refundRemark` | String | 否 | 退款说明，用于接收下游退款单的实际描述 | 克重不对，偷工减料了 |

<a id="m-alibaba-ocean-openplatform-biz-trade-common-model-orderentrycountmodel[]"></a>
#### alibaba.ocean.openplatform.biz.trade.common.model.OrderEntryCountModel[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `id` | Long | 是 | 子订单id | 1 |
| `count` | Integer | 是 | 子订单购买商品数量 | 1 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:alibaba.ocean.openplatform.common.OrderRefundCreateResult](#m-alibaba-ocean-openplatform-common-orderrefundcreateresult) | 是 | 返回结果 | {"result":{"result":{"refundId":"TQ34931008034741532"},"success":true}} |

<a id="m-alibaba-ocean-openplatform-common-orderrefundcreateresult"></a>
#### alibaba.ocean.openplatform.common.OrderRefundCreateResult

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `code` | String | 是 | 错误码 |   |
| `message` | String | 是 | 错误信息 |   |
| `result` | [message:alibaba.ocean.openplatform.biz.trade.result.OrderRefundCreateResult](#m-alibaba-ocean-openplatform-biz-trade-result-orderrefundcreateresult) | 是 | 结果 |   |
| `success` | boolean | 是 | 是否成功 |   |

<a id="m-alibaba-ocean-openplatform-biz-trade-result-orderrefundcreateresult"></a>
#### alibaba.ocean.openplatform.biz.trade.result.OrderRefundCreateResult

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `refundId` | java.lang.String | 是 | 创建成功，退款id | 1 |
| `innerReasonId` | Long | 是 | 1688退款原因ID | 1 |
| `innerReason` | String | 是 | 1688退款原因 | 不 |
| `innerPostFee` | Long | 是 | 转换后退运费金额（分） | 123 |
| `isVoucher` | String | 是 | 是否写入留言凭证 1-是 0-否 | 1 |
| `refundOfficialSolutionCost` | Long | 是 | 退官方物流提货订单费用（分） | 770 |

## 错误码

| 错误码 | 现象 | 解决方案 |
|---|---|---|
| 5001 | 订单不存在 |  |
| 5002 | 订单买家与用户身份不一致 |  |
| 5003 | 货物状态状态不合法 | 参考goodsStatus示例 |
| 5004 | 退款退货原因错误，请确认原因id | 使用API alibaba.trade.getRefundReasonListI查找合法的参数 |
| 5005 | 需要上传凭证 | 部分退款原因需要强制上传凭证 |
| 5006 | 退运费大于0 | 部分退款原因不退运费 |
| 5010 | 退货状态错误 | 请确认订单是否支持能够退款，以及输入参数是否合法：<br>1.退款退货（已收到货的情况下，才能退款退货）<br>2.货物状态是否与订单状态一致（是否发货、是否售中）<br>3.检查子订单id、订单状态、是否重复提交等 |
| 5011 | 退款金额错误 | 不大于实际付款金额 |
| 5012 | 退款描述长度错误 |  |
| 5013 | 子订单退货数量 | 仅在售中买家已收货退款退货时，可以指定小于全部。在全部退货的情况下，推荐不填写，使用默认值。 |
| 5100 | 其他未知错误 | 请确认订单是否可申请退款，按照示例填写参数 |
| 4xxx | 服务方内部错误 |  |

## 示例

**输入示例**

```
orderId:506683458994743225
orderEntryIds:[506683458996743225]
disputeRequest:returnRefund
applyPayment:1
applyCarriage:0
applyReasonId:20028
description:测试
goodsStatus:refundBuyerReceived
vouchers:
orderEntryCountList:[{"id":506683458996743225,"count":1}]

```

****

```
售后、售中的区别：
交易完成前，为售中；交易完成后，为售后。一般情况，点击确认收货，交易完成。特殊情况，如：赊账、物流未签收等，视为售中或者不支持退款退货。

售中收到货、售后未收到货：
售中收到货的典型场景是卖家已发货买家未确认收货。此时，买家退货可以选择售中收到货，才能支持退款。
售后未收到货是防止买家确认收货，但实际未收货的情况。
```
