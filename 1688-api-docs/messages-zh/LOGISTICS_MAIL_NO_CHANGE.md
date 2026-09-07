# 物流单号修改消息

Topic: `LOGISTICS_MAIL_NO_CHANGE` · Group: LOGISTICS (物流消息)  
Doc page: https://open.1688.com/doc/topicDetail.htm?id=LOGISTICS_MAIL_NO_CHANGE

针对1688物流单号修改的情况，及时通知上下游

## 消息字段

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `MailNoChangeModel` | Object | 否 |  |  |
| &nbsp;&nbsp;↳ `logisticsId` | String | 否 | 物流单号 | 123456 |
| &nbsp;&nbsp;↳ `oldCpCode` | String | 否 | 更改前的cp code | 1234 |
| &nbsp;&nbsp;↳ `newCpCode` | String | 否 | 更改后的cp code | 12345 |
| &nbsp;&nbsp;↳ `oldMailNo` | String | 否 | 更改前的运单号 | 123 |
| &nbsp;&nbsp;↳ `newMailNo` | String | 否 | 更改后的运单号 | 1234 |
| &nbsp;&nbsp;↳ `eventTime` | Date | 否 | 发生时间 |  |
| &nbsp;&nbsp;↳ `orderLogsItems` | Object[] | 否 | 该物流单关联的订单信息 |  |
| &nbsp;&nbsp;&nbsp;&nbsp;↳ `orderId` | Number | 否 | 交易主单id | 2958424554509662976 |
| &nbsp;&nbsp;&nbsp;&nbsp;↳ `orderEntryId` | Number | 否 | 交易子单id | 2958424554509662976 |

## 消息示例

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
