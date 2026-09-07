# Logistics order status changed (buyer view)

Original name: 物流单状态变更（买家视角）  
Topic: `LOGISTICS_BUYER_VIEW_TRACE` · Group: LOGISTICS (Logistics Messages)  
Doc page: https://open.1688.com/doc/topicDetail.htm?id=LOGISTICS_BUYER_VIEW_TRACE

1688 logistics order status change message, covering the first-time trigger of five nodes: shipment, pickup, transport, delivery, and signed for (delivered). Only the buyer or users authorized by the buyer can receive this.

## Payload fields

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `OrderLogisticsTracingModel` | Object | no |  |  |
| &nbsp;&nbsp;↳ `logisticsId` | String | no | Logistics number | 12345 |
| &nbsp;&nbsp;↳ `cpCode` | String | no | cp code | 123 |
| &nbsp;&nbsp;↳ `mailNo` | String | no | Waybill number | 123456 |
| &nbsp;&nbsp;↳ `statusChanged` | String | no | Status change of the logistics order, including shipment (CONSIGN), pickup (ACCEPT), transport (TRANSPORT), delivery (DELIVERING), awaiting pickup (AGENT_SIGN), signed for/delivered (SIGN), and exception (FAILED) | CONSIGN |
| &nbsp;&nbsp;↳ `changeTime` | Date | no | Time the change occurred |  |
| &nbsp;&nbsp;↳ `orderLogsItems` | Object[] | no | Order information associated with this logistics order |  |
| &nbsp;&nbsp;&nbsp;&nbsp;↳ `orderId` | Number | no | Main order ID of the transaction | 2938624554509662976 |
| &nbsp;&nbsp;&nbsp;&nbsp;↳ `orderEntryId` | Number | no | Sub-order ID of the transaction | 2958624554509662976 |

## Sample message

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
