# Contract / purchase order paid successfully by both parties

Original name: 合同/采购单双方支付成功消息通知  
Topic: `SYT_CONTRACT_PAY_SUCCESS` · Group: SYT (生意通)  
Doc page: https://open.1688.com/doc/topicDetail.htm?id=SYT_CONTRACT_PAY_SUCCESS

Message notification when both parties successfully pay for a contract/purchase order

## Payload fields

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `draftNo` | String | yes | Contract number or purchase order number | 88SYT20230113169004 |
| `contractCurrentStatus` | String | yes | Current status of the contract or purchase order | PAID |
| `eventType` | String | yes | Current event type | CONTRACT_PAY_SUCCESS_EVENT |
| `payChannel` | String | yes | Buyer's payment channel | BANK_TRANSFER |

## Sample message

```json
{
  "draftNo": "88SYT20230113169004",
  "contractCurrentStatus": "PAID",
  "eventType": "CONTRACT_PAY_SUCCESS_EVENT",
  "payChannel": "BANK_TRANSFER"
}
```
