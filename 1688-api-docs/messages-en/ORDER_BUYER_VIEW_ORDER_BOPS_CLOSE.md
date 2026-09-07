# Order closed by operations back-office (buyer view)

Original name: 1688运营后台关闭订单（买家视角）  
Topic: `ORDER_BUYER_VIEW_ORDER_BOPS_CLOSE` · Group: ORDER (订单消息)  
Doc page: https://open.1688.com/doc/topicDetail.htm?id=ORDER_BUYER_VIEW_ORDER_BOPS_CLOSE

1688 order closed by operations backend, closed due to payment timeout, including cases where the buyer placed an order but did not pay and the system automatically closed the order

## Payload fields

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `orderId` | Number | yes | Order ID | 167539019420540000 |
| `currentStatus` | String | yes | Current order status | cancel |
| `msgSendTime` | String | yes | Message sending time | 2018-05-30 19:34:27 |
| `buyerMemberId` | String | yes | Buyer's member ID on the Chinese site | b2b-665170100 |
| `sellerMemberId` | String | yes | Seller's member ID on the Chinese site | b2b-1676547900b7bb3 |

## Sample message

```json
{
  "orderId": 167539019420540000,
  "currentStatus": "cancel",
  "msgSendTime": "2018-05-30 19:34:27",
  "buyerMemberId": "b2b-665170100",
  "sellerMemberId": "b2b-1676547900b7bb3"
}
```
