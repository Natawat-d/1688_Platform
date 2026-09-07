# 1688 cross-border logistics package update

Original name: 1688跨境物流包裹消息  
Topic: `LOGISTICS_GLOBAL_1688_PACKAGE_CHANGE` · Group: LOGISTICS (Logistics Messages)  
Doc page: https://open.1688.com/doc/topicDetail.htm?id=LOGISTICS_GLOBAL_1688_PACKAGE_CHANGE

1688 cross-border logistics package created or changed

## Payload fields

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `packageId` | Number | yes | Unique package ID | 12345 |
| `buyerUserId` | Number | yes | Buyer ID | 123456 |
| `stageType` | String | yes | Stage type the package belongs to | external |
| `mailNo` | String | yes | Waybill number | EMS12345 |
| `carrierPartnerCode` | String | yes | Logistics company code | EMS |
| `transportStatus` | Number | yes | Package shipment status | 900 |
| `warehouseStatus` | Number | no | Package status within the warehouse | 900100 |
| `actionTime` | Number | yes | Time the package change occurred, in ms | 1739951291000 |
| `tradeOrderIds` | Number[] | yes | List of order IDs the package belongs to | [123, 456] |

## Sample message

```json
{
  "packageId": 12345,
  "buyerUserId": 123456,
  "stageType": "external",
  "mailNo": "EMS12345",
  "carrierPartnerCode": "EMS",
  "transportStatus": 900,
  "warehouseStatus": 900100,
  "actionTime": 1739951291000,
  "tradeOrderIds": [
    123,
    456
  ]
}
```
