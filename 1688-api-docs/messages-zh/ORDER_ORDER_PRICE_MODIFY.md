# 1688修改订单价格（卖家视角）

Topic: `ORDER_ORDER_PRICE_MODIFY` · Group: ORDER (订单消息)  
Doc page: https://open.1688.com/doc/topicDetail.htm?id=ORDER_ORDER_PRICE_MODIFY

1688修改订单价格（包括关闭子订单），订单价格修改成功后会发送该消息，用户可以根据该消息调用API查询最新的订单价格。

## 消息字段

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `orderId` | Number | 是 | 订单ID | 168133548553170001 |
| `currentStatus` | String | 是 | 当前订单状态，状态值为waitbuyerpay | waitbuyerpay |
| `msgSendTime` | String | 是 | 消息发送时间 | 2018-05-30 19:34:27 |
| `buyerMemberId` | String | 否 | 买家中文站会员ID | b2b-665170100 |

## 消息示例

```json
{
  "orderId": 168133548553170001,
  "currentStatus": "waitbuyerpay",
  "msgSendTime": "2018-05-30 19:34:27",
  "buyerMemberId": "b2b-665170100"
}
```
