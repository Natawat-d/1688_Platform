# Order batch-payment status sync

Original name: 1688订单批量支付状态同步消息  
Topic: `ORDER_BATCH_PAY` · Group: ORDER (订单消息)  
Doc page: https://open.1688.com/doc/topicDetail.htm?id=ORDER_BATCH_PAY

Order batch payment status sync message; returns the payment status of each order in a batch payment.

## Payload fields

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `batchPay` | Object[] | yes |  |  |
| &nbsp;&nbsp;↳ `orderId` | String | yes | Order id | 558880835194545941 |
| &nbsp;&nbsp;↳ `status` | String | yes | Order payment status, which can be successed (payment successful), ACCOUNT_BALANCE_NOT_ENOUGH (insufficient balance), ACCOUNT_NOT_EXIST (in the Cross-border Bao 2.0 scenario, the account may be signed up but not opened in ipay), ACCOUNT_FROZEN (account frozen), or PARAM_ILLEGAL (invalid parameter) | successed |

## Sample message

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
