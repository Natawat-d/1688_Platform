# Query refund/return reasons (for creating a request)

Original name: 查询退款退货原因（用于创建退款退货）  
API: `com.alibaba.trade:alibaba.trade.getRefundReasonList:1` · Category: Returns & Refunds  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.trade.getRefundReasonList-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.trade.getRefundReasonList/{appKey}`  
Requires user authorization (access_token) · Requires signature

Query the refund/return reasons (used when creating a refund or return request).

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `orderId` | Long | yes | Main order id |   |
| `orderEntryIds` | Long[] | yes | Sub-order ID |   |
| `goodsStatus` | String | yes | Cargo status | 售中等待买家发货:”refundWaitSellerSend"; 售中等待买家收货:"refundWaitBuyerReceive"; 售中已收货（未确认完成交易）:"refundBuyerReceived" 售后未收货:"aftersaleBuyerNotReceived"; 售后已收到货:"aftersaleBuyerReceived" |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.ocean.openplatform.common.OrderRefundReasonListResult](#m-alibaba-ocean-openplatform-common-orderrefundreasonlistresult) | yes | Return result |   |

<a id="m-alibaba-ocean-openplatform-common-orderrefundreasonlistresult"></a>
#### alibaba.ocean.openplatform.common.OrderRefundReasonListResult

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `code` | String | yes | Error code |   |
| `message` | String | yes | Error message |   |
| `result` | [message:alibaba.ocean.openplatform.biz.trade.result.OrderRefundReasonListResult](#m-alibaba-ocean-openplatform-biz-trade-result-orderrefundreasonlistresult) | yes | Result |   |
| `success` | Boolean | yes | Whether successful |   |

<a id="m-alibaba-ocean-openplatform-biz-trade-result-orderrefundreasonlistresult"></a>
#### alibaba.ocean.openplatform.biz.trade.result.OrderRefundReasonListResult

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `reasons` | [message:alibaba.ocean.openplatform.biz.trade.result.OrderRefundReasonModel[]](#m-alibaba-ocean-openplatform-biz-trade-result-orderrefundreasonmodel[]) | yes | List of reasons |   |

<a id="m-alibaba-ocean-openplatform-biz-trade-result-orderrefundreasonmodel[]"></a>
#### alibaba.ocean.openplatform.biz.trade.result.OrderRefundReasonModel[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `id` | java.lang.Long | yes | Reason ID |   |
| `name` | java.lang.String | yes | Reason |   |
| `needVoucher` | java.lang.Boolean | yes | Whether the voucher must be uploaded | "true"表示必须要上传凭证 |
| `noRefundCarriage` | java.lang.Boolean | yes | Whether return shipping fee refund is supported | “true" 表示不支持退运费 |
| `tip` | java.lang.String | yes | Prompt |   |

## Error codes

| Error | Symptom | How to fix |
|---|---|---|
| 5xxx | Caller error | Please confirm whether the input parameters are valid.<br>Goods status parameter example; sub-orders require array input... |
| 4xxx | Service-side internal error |  |

## Samples

**Input example**

```
orderId:586683458994743215
orderEntryIds:[586683458997743215]
goodsStatus:aftersaleBuyerNotReceived

```
