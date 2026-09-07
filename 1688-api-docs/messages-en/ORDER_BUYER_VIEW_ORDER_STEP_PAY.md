# Order stage payment (buyer view)

Original name: 1688订单阶段付款（买家视角）  
Topic: `ORDER_BUYER_VIEW_ORDER_STEP_PAY` · Group: ORDER (订单消息)  
Doc page: https://open.1688.com/doc/topicDetail.htm?id=ORDER_BUYER_VIEW_ORDER_STEP_PAY

1688 order staged payment (buyer view)

## Payload fields

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `orderId` | Number | yes | Order ID | 167539019420540000 |
| `currentStatus` | String | yes | Current order status; the status value is waitsellersend | waitsellersend |
| `msgSendTime` | String | yes | Message sending time | 2018-05-30 19:34:27 |
| `buyerMemberId` | String | yes | Buyer's member ID on the Chinese site | b2b-665170100 |
| `sellerMemberId` | String | no | Seller's member ID on the Chinese site | b2b-1676547900b7bb3 |

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
