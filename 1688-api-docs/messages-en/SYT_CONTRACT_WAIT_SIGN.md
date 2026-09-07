# Contract or purchase order awaiting signature

Original name: 合同或采购单待签署消息  
Topic: `SYT_CONTRACT_WAIT_SIGN` · Group: SYT (生意通)  
Doc page: https://open.1688.com/doc/topicDetail.htm?id=SYT_CONTRACT_WAIT_SIGN

A message notification is sent when the contract or purchase order requires signing confirmation

## Payload fields

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `draftNo` | String | yes | Contract number or purchase order number | 88SYT20230113169004 |
| `contractCurrentStatus` | String | yes | Current status of the contract or purchase order | SIGNING |
| `eventType` | String | yes | Current event type | CONTRACT_WAIT_SIGN_EVENT |

## Sample message

```json
{
  "draftNo": "88SYT20230113169004",
  "contractCurrentStatus": "SIGNING",
  "eventType": "CONTRACT_WAIT_SIGN_EVENT"
}
```
