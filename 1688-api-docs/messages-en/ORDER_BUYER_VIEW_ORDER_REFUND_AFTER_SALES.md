# After-sales refund on order (buyer view)

Original name: 1688订单售后退款（买家视角）  
Topic: `ORDER_BUYER_VIEW_ORDER_REFUND_AFTER_SALES` · Group: ORDER (订单消息)  
Doc page: https://open.1688.com/doc/topicDetail.htm?id=ORDER_BUYER_VIEW_ORDER_REFUND_AFTER_SALES

1688 order after-sales refund; the transaction has been completed, only the buyer can receive the refund message

## Payload fields

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `orderId` | Number | yes | Order ID | 161292821695900310 |
| `currentStatus` | String | yes | Current order status | waitbuyersend |
| `msgSendTime` | String | yes | Message sending time | 2018-05-30 19:37:28 |
| `buyerMemberId` | String | yes | Buyer's member ID on the Chinese site | b2b-346900403 |
| `refundAction` | String | no | Refund operation; the specific statuses are: BUYER_APPLY_REFUND (buyer applies for refund), BUYER_RECEIVE_CLOSE (buyer confirms receipt, closed), SELLER_SEND_GOODS_CLOSE (seller ships goods, closed), BUYER_CANCEL_REFUND_CLOSE (buyer cancels refund application, closed), BUYER_UPLOAD_BILL (buyer uploads proof), SELLER_UPLOAD_BILL (seller uploads proof), SELLER_REJECT_REFUND (seller rejects refund), SELLER_AGREE_REFUND (seller agrees to refund), SELLER_RECEIVE_GOODS (seller confirms receipt of returned goods), BUYER_SEND_GOODS (buyer declares goods shipped), BUYER_MODIFY_REFUND_PROTOCOL (buyer modifies refund agreement), BUYER_APPLY_SUPPORT (buyer requests customer service intervention), SELLER_APPLY_SUPPORT (seller requests customer service intervention), SYSTEM_AGREE_REFUND_PROTOCOL (system agrees to refund agreement on timeout), SYSTEM_AGREE_REFUND (system agrees to refund on timeout, i.e. refund succeeds), SYSTEM_SEND_GOODS (system-triggered return on timeout, return in the main transaction flow), SYSTEM_MODIFY_REFUND_PROTOCOL (system modifies agreement on timeout), SYSTEM_NOTIFY_APPLY_SUPPORT (system notifies that customer service intervention was requested), SELLER_AGREE_REFUND_PROCOTOL (seller agrees to refund agreement), SELLER_REJECT_REFUND_PROCOTOL (seller rejects refund agreement), CRM_APPLY_TIMEOUT_CLOSE (customer service intervention requested, closed on timeout; currently used only for after-sales business), CRM_APPLY_SUPPORT (CRM intervention requested), CRM_INTERVENE_TASK (CRM handles the case), CRM_DISMISS_TASK (CRM dismisses the ticket), CRM_FINISH_TASK (CRM closes the ticket), BUYER_STEP_PAY_ORDER_CLOSE (buyer paid, refund closed, staged-order scenario), BUYER_STEP_CONFIRM_CLOSE (buyer confirmed, refund closed, staged-order scenario), BUYER_CLOSE_TRADE_CLOSE (buyer terminated transaction, refund closed, staged-order scenario), SELLER_CONFIRM_ORDER_CLOSE (seller confirmed, refund closed, staged-order scenario), SELLER_STEP_PUSH_CLOSE (seller advanced, refund closed, staged-order scenario) | SELLER_AGREE_REFUND_PROCOTOL |
| `operator` | String | no | Initiator of the operation: buyer, seller, or system | seller |
| `sellerMemberId` | String | yes | Seller's member ID on the Chinese site | b2b-9161839253fcc5 |
| `refundId` | String | yes | Refund order ID | 1234567 |

## Sample message

```json
{
  "orderId": 161292821695900310,
  "currentStatus": "waitbuyersend",
  "msgSendTime": "2018-05-30 19:37:28",
  "buyerMemberId": "b2b-346900403",
  "refundAction": "SELLER_AGREE_REFUND_PROCOTOL",
  "operator": "seller",
  "sellerMemberId": "b2b-9161839253fcc5",
  "refundId": "1234567"
}
```
