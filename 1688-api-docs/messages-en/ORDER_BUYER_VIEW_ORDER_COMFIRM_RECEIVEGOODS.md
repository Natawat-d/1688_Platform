# Receipt of goods confirmed (buyer view)

Original name: 1688订单确认收货（买家视角）/order receipt confirmation (buyer view)  
Topic: `ORDER_BUYER_VIEW_ORDER_COMFIRM_RECEIVEGOODS` · Group: ORDER (订单消息)  
Doc page: https://open.1688.com/doc/topicDetail.htm?id=ORDER_BUYER_VIEW_ORDER_COMFIRM_RECEIVEGOODS

1688 order receipt confirmation (buyer view). This message is sent every time receipt is confirmed, including partial receipt confirmations.

## Payload fields

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `orderId` | Number | yes | Order ID | 167539019420540000 |
| `currentStatus` | String | yes | Current order status; the status value is confirm_goods_and_has_subsidy | confirm_goods_and_has_subsidy |
| `msgSendTime` | String | yes | Message sending time | 2018-05-30 19:34:27 |
| `buyerMemberId` | String | yes | Buyer's member ID on the Chinese site | b2b-665170100 |
| `sellerMemberId` | String | yes | Seller's member ID on the Chinese site | b2b-1676547900b7bb3 |

## Sample message

```json
{
  "orderId": 167539019420540000,
  "currentStatus": "confirm_goods_and_has_subsidy",
  "msgSendTime": "2018-05-30 19:34:27",
  "buyerMemberId": "b2b-665170100",
  "sellerMemberId": "b2b-1676547900b7bb3"
}
```
