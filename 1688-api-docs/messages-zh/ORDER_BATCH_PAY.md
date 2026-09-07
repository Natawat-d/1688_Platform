# 1688订单批量支付状态同步消息

Topic: `ORDER_BATCH_PAY` · Group: ORDER (订单消息)  
Doc page: https://open.1688.com/doc/topicDetail.htm?id=ORDER_BATCH_PAY

订单批量支付状态同步消息，能返回批量支付订单各个订单的支付状态。

## 消息字段

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `batchPay` | Object[] | 是 |  |  |
| &nbsp;&nbsp;↳ `orderId` | String | 是 | 订单id | 558880835194545941 |
| &nbsp;&nbsp;↳ `status` | String | 是 | 订单支付状态，可为successed（支付成功）、ACCOUNT_BALANCE_NOT_ENOUGH（余额不足）、ACCOUNT_NOT_EXIST（跨境宝2.0场景下可能签约但是在ipay没有开户）、ACCOUNT_FROZEN（账户冻结）、PARAM_ILLEGAL（参数非法） | successed |

## 消息示例

```json
{
  "batchPay": [
    {
      "orderId": "558880835194545941",
      "status": "successed"
    }
  ]
}
```
