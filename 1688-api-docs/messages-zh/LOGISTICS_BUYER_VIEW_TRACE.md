# 物流单状态变更（买家视角）

Topic: `LOGISTICS_BUYER_VIEW_TRACE` · Group: LOGISTICS (物流消息)  
Doc page: https://open.1688.com/doc/topicDetail.htm?id=LOGISTICS_BUYER_VIEW_TRACE

1688物流单状态变更消息，包括发货、揽收、运输、派送、签收五个节点首次触发时的消息，仅买家或买家授权的用户能接收到

## 消息字段

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `OrderLogisticsTracingModel` | Object | 否 |  |  |
| &nbsp;&nbsp;↳ `logisticsId` | String | 否 | 物流编号 | 12345 |
| &nbsp;&nbsp;↳ `cpCode` | String | 否 | cp code | 123 |
| &nbsp;&nbsp;↳ `mailNo` | String | 否 | 运单号 | 123456 |
| &nbsp;&nbsp;↳ `statusChanged` | String | 否 | 物流单发生变化的状态，包括发货（CONSIGN）、揽收（ACCEPT）、运输（TRANSPORT）、派送（DELIVERING）、待取件（AGENT_SIGN）、签收（SIGN）、异常（FAILED） | CONSIGN |
| &nbsp;&nbsp;↳ `changeTime` | Date | 否 | 发生变化的时间 |  |
| &nbsp;&nbsp;↳ `orderLogsItems` | Object[] | 否 | 该物流单关联的订单信息 |  |
| &nbsp;&nbsp;&nbsp;&nbsp;↳ `orderId` | Number | 否 | 交易主单id | 2938624554509662976 |
| &nbsp;&nbsp;&nbsp;&nbsp;↳ `orderEntryId` | Number | 否 | 交易子单id | 2958624554509662976 |

## 消息示例

```json
{
  "OrderLogisticsTracingModel": {
    "logisticsId": "12345",
    "cpCode": "123",
    "mailNo": "123456",
    "statusChanged": "CONSIGN",
    "changeTime": " ",
    "orderLogsItems": [
      {
        "orderId": 2938624554509662976,
        "orderEntryId": 2958624554509662976
      }
    ]
  }
}
```
