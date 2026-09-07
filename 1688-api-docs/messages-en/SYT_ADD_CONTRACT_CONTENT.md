# New transaction evidence record added

Original name: 新增交易存证通知消息  
Topic: `SYT_ADD_CONTRACT_CONTENT` · Group: SYT (生意通)  
Doc page: https://open.1688.com/doc/topicDetail.htm?id=SYT_ADD_CONTRACT_CONTENT

Can be used for: a message notification is sent when the counterparty adds logistics proof, shipment proof, payment proof, or other information on the purchase order/contract detail page

## Payload fields

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `draftNo` | String | yes | Contract number/purchase order number | 88SYT20240113169004 |
| `content` | String | no | Specific content added, text content | 无 |
| `eventType` | String | yes | Current event type | CONTRACT_CONTENT_ADD_SUCCESS |
| `submitterRole` | String | yes | Role of the submitter of the transaction evidence record | PART_A:买家;PART_B:商家 |
| `addTime` | Number | yes | Millisecond timestamp of the add-time event | 1768288592321 |
| `attachments` | Object[] | no | File attachment | 无 |
| &nbsp;&nbsp;↳ `fileName` | String | yes | Attachment name | a.png |
| &nbsp;&nbsp;↳ `filePath` | String | yes | Attachment file address | https://www.xxx.xxx/xxx.pdf |

## Sample message

```json
{
  "draftNo": "88SYT20240113169004",
  "content": "无",
  "eventType": "CONTRACT_CONTENT_ADD_SUCCESS",
  "submitterRole": "PART_A:买家;PART_B:商家",
  "addTime": 1768288592321,
  "attachments": [
    {
      "fileName": "a.png",
      "filePath": "https://www.xxx.xxx/xxx.pdf"
    }
  ]
}
```
