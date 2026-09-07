# 1688创建订单（买家视角）/order created (buyer view)

Topic: `ORDER_BUYER_VIEW_BUYER_MAKE` · Group: ORDER (订单消息)  
Doc page: https://open.1688.com/doc/topicDetail.htm?id=ORDER_BUYER_VIEW_BUYER_MAKE

1688创建订单（买家视角）/order created (buyer view)

## 消息字段

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `orderId` | Number | 是 | 订单ID | 167539019420540000 |
| `currentStatus` | String | 是 | 当前订单状态，状态值为waitbuyerpay | waitbuyerpay |
| `msgSendTime` | String | 是 | 消息发送时间 | 2018-05-30 19:24:18 |
| `buyerMemberId` | String | 是 | 买家中文站会员ID | b2b-665170100 |
| `sellerMemberId` | String | 是 | 卖家中文站会员ID | b2b-1676547900b7bb3 |

## 消息示例

```json
{
  "orderId": 167539019420540000,
  "currentStatus": "waitbuyerpay",
  "msgSendTime": "2018-05-30 19:24:18",
  "buyerMemberId": "b2b-665170100",
  "sellerMemberId": "b2b-1676547900b7bb3"
}
```
