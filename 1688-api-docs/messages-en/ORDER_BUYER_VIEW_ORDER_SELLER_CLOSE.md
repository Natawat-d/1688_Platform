# Order closed by seller (buyer view)

Original name: 1688卖家关闭订单（买家视角）/seller closing order (buyer view)  
Topic: `ORDER_BUYER_VIEW_ORDER_SELLER_CLOSE` · Group: ORDER (订单消息)  
Doc page: https://open.1688.com/doc/topicDetail.htm?id=ORDER_BUYER_VIEW_ORDER_SELLER_CLOSE

1688 seller closing order (buyer view). This message is also sent when a refund causes the transaction to close.

## Payload fields

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `orderId` | Number | yes | Order ID | 167539019420540000 |
| `currentStatus` | String | yes | Current order status; the status value is cancel | cancel |
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
