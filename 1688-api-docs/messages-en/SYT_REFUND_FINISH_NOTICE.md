# Contract or purchase order refund result

Original name: 合同或采购单退款结果消息  
Topic: `SYT_REFUND_FINISH_NOTICE` · Group: SYT (生意通)  
Doc page: https://open.1688.com/doc/topicDetail.htm?id=SYT_REFUND_FINISH_NOTICE

Refund result message for the contract or purchase order

## Payload fields

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `draftNo` | String | yes | Contract number/purchase order number | 88SYT20240113169004 |
| `eventType` | String | yes | Current event type | CONTRACT_REFUND_FINISH |
| `refundOrderNo` | String | yes | Refund order number | 8896010RO10005525121800000017003 |
| `refundStatus` | String | yes | Refund order status | REFUND_SUCCESS/FAILED |
| `rejectReason` | String | no | Reason for refund rejection | 无 |

## Sample message

```json
{
  "draftNo": "88SYT20240113169004",
  "eventType": "CONTRACT_REFUND_FINISH",
  "refundOrderNo": "8896010RO10005525121800000017003",
  "refundStatus": "REFUND_SUCCESS/FAILED",
  "rejectReason": "无"
}
```
