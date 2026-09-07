# Order price modified (seller view)

Original name: 1688修改订单价格（卖家视角）  
Topic: `ORDER_ORDER_PRICE_MODIFY` · Group: ORDER (订单消息)  
Doc page: https://open.1688.com/doc/topicDetail.htm?id=ORDER_ORDER_PRICE_MODIFY

1688 order price modification (including closing sub-orders). This message is sent after the order price is successfully modified; users can call the API based on this message to query the latest order price.

## Payload fields

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `orderId` | Number | yes | Order ID | 168133548553170001 |
| `currentStatus` | String | yes | Current order status; the status value is waitbuyerpay | waitbuyerpay |
| `msgSendTime` | String | yes | Message sending time | 2018-05-30 19:34:27 |
| `buyerMemberId` | String | no | Buyer's member ID on the Chinese site | b2b-665170100 |

## Sample message

```json
{
  "orderId": 168133548553170001,
  "currentStatus": "waitbuyerpay",
  "msgSendTime": "2018-05-30 19:34:27",
  "buyerMemberId": "b2b-665170100"
}
```
