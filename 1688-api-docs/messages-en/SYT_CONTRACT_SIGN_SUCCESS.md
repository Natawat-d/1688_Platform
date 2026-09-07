# Contract / purchase order signed successfully by both parties

Original name: 合同/采购单双方签署成功通知  
Topic: `SYT_CONTRACT_SIGN_SUCCESS` · Group: SYT (生意通)  
Doc page: https://open.1688.com/doc/topicDetail.htm?id=SYT_CONTRACT_SIGN_SUCCESS

A message notification is sent after both parties successfully sign the contract/purchase order

## Payload fields

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `draftNo` | String | yes | Contract number or purchase order number | 88SYT20230113169004 |
| `contractCurrentStatus` | String | yes | Current status of the contract/purchase order | SIGN_SUCCESS |
| `eventType` | String | yes | Current event type | CONTRACT_SIGN_SUCCESS_EVENT |
| `signSuccessTime` | String | yes | Millisecond timestamp of successful contract signing | 1778505510000 |

## Sample message

```json
{
  "draftNo": "88SYT20230113169004",
  "contractCurrentStatus": "SIGN_SUCCESS",
  "eventType": "CONTRACT_SIGN_SUCCESS_EVENT",
  "signSuccessTime": "1778505510000"
}
```
