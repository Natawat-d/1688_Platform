# Order created (buyer view)

Original name: 1688创建订单（买家视角）/order created (buyer view)  
Topic: `ORDER_BUYER_VIEW_BUYER_MAKE` · Group: ORDER (订单消息)  
Doc page: https://open.1688.com/doc/topicDetail.htm?id=ORDER_BUYER_VIEW_BUYER_MAKE

1688 order created (buyer view)

## Payload fields

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `orderId` | Number | yes | Order ID | 167539019420540000 |
| `currentStatus` | String | yes | Current order status; the status value is waitbuyerpay | waitbuyerpay |
| `msgSendTime` | String | yes | Message sending time | 2018-05-30 19:24:18 |
| `buyerMemberId` | String | yes | Buyer's member ID on the Chinese site | b2b-665170100 |
| `sellerMemberId` | String | yes | Seller's member ID on the Chinese site | b2b-1676547900b7bb3 |

## Sample message

```json
{
  "orderId": 167539019420540000,
  "currentStatus": "waitbuyerpay",
  "msgSendTime": "2018-05-30 19:24:18",
  "buyerMemberId": "b2b-665170100",
  "sellerMemberId": "b2b-1676547900b7bb3"
}
```
