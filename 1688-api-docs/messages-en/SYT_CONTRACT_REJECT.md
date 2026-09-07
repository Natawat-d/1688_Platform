# Purchase order / contract rejected

Original name: 采购单/合同拒绝消息通知  
Topic: `SYT_CONTRACT_REJECT` · Group: SYT (生意通)  
Doc page: https://open.1688.com/doc/topicDetail.htm?id=SYT_CONTRACT_REJECT

ShengYiTong purchase order/contract rejection message notification

## Payload fields

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `draftNo` | String | yes | Contract number or purchase order number | 88SYT20230113169004 |
| `contractCurrentStatus` | String | yes | Current status of the contract | SIGN_REJECT |
| `eventType` | String | yes | Current event type | CONTRACT_SIGN_REJECT_EVENT |
| `signRejectReason` | String | yes | Reason for rejection | 无 |
| `signRejectTime` | String | yes | Millisecond timestamp of the rejection event | 1778505510000 |

## Sample message

```json
{
  "draftNo": "88SYT20230113169004",
  "contractCurrentStatus": "SIGN_REJECT",
  "eventType": "CONTRACT_SIGN_REJECT_EVENT",
  "signRejectReason": "无",
  "signRejectTime": "1778505510000"
}
```
