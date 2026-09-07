# Waybill number changed

Original name: 物流单号修改消息  
Topic: `LOGISTICS_MAIL_NO_CHANGE` · Group: LOGISTICS (Logistics Messages)  
Doc page: https://open.1688.com/doc/topicDetail.htm?id=LOGISTICS_MAIL_NO_CHANGE

For cases where a 1688 waybill number is modified, notifies upstream and downstream parties promptly

## Payload fields

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `MailNoChangeModel` | Object | no |  |  |
| &nbsp;&nbsp;↳ `logisticsId` | String | no | Logistics tracking number | 123456 |
| &nbsp;&nbsp;↳ `oldCpCode` | String | no | cp code before the change | 1234 |
| &nbsp;&nbsp;↳ `newCpCode` | String | no | cp code after the change | 12345 |
| &nbsp;&nbsp;↳ `oldMailNo` | String | no | Waybill number before the change | 123 |
| &nbsp;&nbsp;↳ `newMailNo` | String | no | Waybill number after the change | 1234 |
| &nbsp;&nbsp;↳ `eventTime` | Date | no | Time of occurrence |  |
| &nbsp;&nbsp;↳ `orderLogsItems` | Object[] | no | Order information associated with this logistics order |  |
| &nbsp;&nbsp;&nbsp;&nbsp;↳ `orderId` | Number | no | Main order ID of the transaction | 2958424554509662976 |
| &nbsp;&nbsp;&nbsp;&nbsp;↳ `orderEntryId` | Number | no | Sub-order ID of the transaction | 2958424554509662976 |

## Sample message

```json
{
  "MailNoChangeModel": {
    "logisticsId": "123456",
    "oldCpCode": "1234",
    "newCpCode": "12345",
    "oldMailNo": "123",
    "newMailNo": "1234",
    "eventTime": " ",
    "orderLogsItems": [
      {
        "orderId": 2958424554509662976,
        "orderEntryId": 2958424554509662976
      }
    ]
  }
}
```
