# Cancel transaction

Original name: 取消交易  
API: `com.alibaba.trade:alibaba.trade.cancel:1` · Category: Orders  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.trade.cancel-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.trade.cancel/{appKey}`  
Requires user authorization (access_token) · Requires signature

Buyer or seller cancels a transaction. Only transactions in specific statuses can be cancelled; on 1688 this is used to cancel unpaid orders. If an order is closed less than 10 seconds after it was created, the error CLOSE_ORDER_TOO_FAST is returned.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `webSite` | String | yes | Site information, specifies whether the called API belongs to the international site (alibaba) or the 1688 website (1688). | 1688 |
| `tradeID` | Long | yes | Transaction ID, order number | 123456 |
| `cancelReason` | String | yes | Reason description; buyerCancel: buyer cancelled the order; sellerGoodsLack: seller out of stock; other: other | other |
| `remark` | String | no | Remark | 备注 |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `success` | Boolean | yes | Whether processed successfully: true for success, false for failure; see error for the failure reason | true |
| `errorCode` | String | yes | Error code | ORDER_STATUS_ERROR |
| `errorMessage` | String | yes | Error message | 订单状态错误 |

## Error codes

| Error | Symptom | How to fix |
|---|---|---|
| ORDER_STATUS_ERROR | Order status error | Can only cancel orders pending payment; the order status needs to be confirmed first |
| 400_3 | No permission to cancel this order | Only the buyer and seller of the order can cancel it; confirm whether the authorized user is the buyer or seller of this order |
| ORDER_NOT_EXIST | Order does not exist | Confirm whether the order number is correct |

## Samples

**Request parameter example**

```
{"webSite":"1688","tradeID":"202711458975969812","cancelReason":"other","remark":"取消订单"}
```

**Example of return parameters**

```
{"success":true}
```
