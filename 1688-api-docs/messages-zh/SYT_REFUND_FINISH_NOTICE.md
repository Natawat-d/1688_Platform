# 合同或采购单退款结果消息

Topic: `SYT_REFUND_FINISH_NOTICE` · Group: SYT (生意通)  
Doc page: https://open.1688.com/doc/topicDetail.htm?id=SYT_REFUND_FINISH_NOTICE

合同或采购单退款结果消息

## 消息字段

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `draftNo` | String | 是 | 合同号/采购单号 | 88SYT20240113169004 |
| `eventType` | String | 是 | 当前事件类型 | CONTRACT_REFUND_FINISH |
| `refundOrderNo` | String | 是 | 退款单号 | 8896010RO10005525121800000017003 |
| `refundStatus` | String | 是 | 退款单状态 | REFUND_SUCCESS/FAILED |
| `rejectReason` | String | 否 | 退款拒绝原因 | 无 |

## 消息示例

```json
{
  "draftNo": "88SYT20240113169004",
  "eventType": "CONTRACT_REFUND_FINISH",
  "refundOrderNo": "8896010RO10005525121800000017003",
  "refundStatus": "REFUND_SUCCESS/FAILED",
  "rejectReason": "无"
}
```
