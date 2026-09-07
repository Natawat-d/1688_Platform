# Create refund/return request

Original name: 创建退款退货申请  
API: `com.alibaba.trade:alibaba.trade.createRefund:1` · Category: Returns & Refunds  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.trade.createRefund-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.trade.createRefund/{appKey}`  
Requires user authorization (access_token) · Requires signature

Create a refund or return request.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `orderId` | Long | yes | Main order |    |
| `orderEntryIds` | Long[] | yes | Sub-order |   |
| `disputeRequest` | String | yes | Refund / refund and return. Refund and return can only be selected once the goods have been received. | 退款:"refund"; 退款退货:"returnRefund" |
| `applyPayment` | Long | yes | Refund amount (unit: cent). Must not exceed the actual payment amount; when waiting for the seller to ship, it must equal the product's actual payment amount. |   |
| `applyCarriage` | Long | yes | Refunded freight amount (unit: cents/fen). |   |
| `applyReasonId` | Long | yes | Refund reason ID (obtained from the API getRefundReasonList) |   |
| `description` | String | yes | Reason for refund application, 2-150 characters |   |
| `goodsStatus` | String | yes | Cargo status |  售中等待卖家发货:"refundWaitSellerSend"; 售中等待买家收货:"refundWaitBuyerReceive"; 售中已收货（未确认完成交易）:"refundBuyerReceived" 售后未收货:"aftersaleBuyerNotReceived"; 售后已收到货:"aftersaleBuyerReceived" |
| `vouchers` | String[] | no | Voucher image URLs. 1-5 images, must use the “image domain/relative path” returned by the uploadRefundVoucher API |  [https://cbu01.alicdn.com/img/ibank/2019/901/930/11848039109.jpg] |
| `orderEntryCountList` | [message:alibaba.ocean.openplatform.biz.trade.common.model.OrderEntryCountModel[]](#m-alibaba-ocean-openplatform-biz-trade-common-model-orderentrycountmodel[]) | no | Sub-order refund quantity. The return quantity can be specified only when the buyer has already received the goods during in-sale (refund and return); by default, all goods are returned. |  [{"id":586683458996743215,"count":1}] |
| `customRefund` | Boolean | no | Whether the return shipping fee is customized | true：订单退款运费按照isv回传为准，不做优化；false（默认）：订单退款金额按照平台规则为准，需要做优化； |
| `refundRemark` | String | no | Refund description, used to receive the actual description of the downstream refund order | 克重不对，偷工减料了 |

<a id="m-alibaba-ocean-openplatform-biz-trade-common-model-orderentrycountmodel[]"></a>
#### alibaba.ocean.openplatform.biz.trade.common.model.OrderEntryCountModel[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `id` | Long | yes | Sub-order ID | 1 |
| `count` | Integer | yes | Quantity of products purchased in the sub-order | 1 |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.ocean.openplatform.common.OrderRefundCreateResult](#m-alibaba-ocean-openplatform-common-orderrefundcreateresult) | yes | Return result | {"result":{"result":{"refundId":"TQ34931008034741532"},"success":true}} |

<a id="m-alibaba-ocean-openplatform-common-orderrefundcreateresult"></a>
#### alibaba.ocean.openplatform.common.OrderRefundCreateResult

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `code` | String | yes | Error code |   |
| `message` | String | yes | Error message |   |
| `result` | [message:alibaba.ocean.openplatform.biz.trade.result.OrderRefundCreateResult](#m-alibaba-ocean-openplatform-biz-trade-result-orderrefundcreateresult) | yes | Result |   |
| `success` | boolean | yes | Whether successful |   |

<a id="m-alibaba-ocean-openplatform-biz-trade-result-orderrefundcreateresult"></a>
#### alibaba.ocean.openplatform.biz.trade.result.OrderRefundCreateResult

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `refundId` | java.lang.String | yes | Created successfully, refund id | 1 |
| `innerReasonId` | Long | yes | 1688 refund reason ID | 1 |
| `innerReason` | String | yes | 1688 refund reason | 不 |
| `innerPostFee` | Long | yes | Converted return shipping fee amount (cents) | 123 |
| `isVoucher` | String | yes | Whether to write a message voucher/record 1 - yes 0 - no | 1 |
| `refundOfficialSolutionCost` | Long | yes | Fee for returning an official logistics pickup order (cents/fen) | 770 |

## Error codes

| Error | Symptom | How to fix |
|---|---|---|
| 5001 | Order does not exist |  |
| 5002 | The order buyer does not match the user identity |  |
| 5003 | Goods status is invalid | Refer to the goodsStatus example |
| 5004 | Incorrect refund/return reason; please confirm the reason id | Use the API alibaba.trade.getRefundReasonListI to find valid parameters |
| 5005 | Proof needs to be uploaded | Some refund reasons require mandatory proof upload |
| 5006 | Return shipping fee greater than 0 | Some refund reasons do not refund the shipping fee |
| 5010 | Return status error | Please confirm whether the order supports refund, and whether the input parameters are valid:<br>1. Refund and return (refund and return is only possible once the goods have been received)<br>2. Whether the goods status matches the order status (whether shipped, whether in-sale)<br>3. Check the sub-order id, order status, whether duplicately submitted, etc. |
| 5011 | Incorrect refund amount | Not greater than the actual payment amount |
| 5012 | Refund description length error |  |
| 5013 | Sub-order return quantity | Can only be specified as less than the total when the in-sale buyer has received the goods and applies for a refund and return. For a full return, it is recommended not to fill this in and use the default value. |
| 5100 | Other unknown error | Please confirm whether the order is eligible for a refund request, and fill in the parameters according to the example |
| 4xxx | Service-side internal error |  |

## Samples

**Input example**

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
