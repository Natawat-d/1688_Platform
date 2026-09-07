# Order partially shipped (buyer view)

Original name: 1688订单部分发货（买家视角）/Partial delivery of 1688 order (buyer view)  
Topic: `ORDER_BUYER_VIEW_PART_PART_SENDGOODS` · Group: ORDER (订单消息)  
Doc page: https://open.1688.com/doc/topicDetail.htm?id=ORDER_BUYER_VIEW_PART_PART_SENDGOODS

Partial delivery of 1688 order (buyer view). Status changes sequentially, but messages are sent asynchronously, so a later-sent message may arrive before an earlier one.

## Payload fields

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `orderId` | Number | yes | Order ID | 167539019420540000 |
| `currentStatus` | String | yes | Current order status; the status value is waitsellersend | waitsellersend |
| `msgSendTime` | String | yes | Message sending time | 2018-05-30 19:34:27 |
| `buyerMemberId` | String | yes | Buyer's member ID on the Chinese site | b2b-665170100 |
| `sellerMemberId` | String | yes | Seller's member ID on the Chinese site | b2b-1676547900b7bb3 |

## Sample message

```json
{
  "orderId": 167539019420540000,
  "currentStatus": "waitsellersend",
  "msgSendTime": "2018-05-30 19:34:27",
  "buyerMemberId": "b2b-665170100",
  "sellerMemberId": "b2b-1676547900b7bb3"
}
```
