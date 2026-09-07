# 1688跨境物流包裹消息

Topic: `LOGISTICS_GLOBAL_1688_PACKAGE_CHANGE` · Group: LOGISTICS (物流消息)  
Doc page: https://open.1688.com/doc/topicDetail.htm?id=LOGISTICS_GLOBAL_1688_PACKAGE_CHANGE

1688跨境物流包裹创建、变更

## 消息字段

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `packageId` | Number | 是 | 包裹唯一id | 12345 |
| `buyerUserId` | Number | 是 | 买家id | 123456 |
| `stageType` | String | 是 | 包裹所属阶段类型 | external |
| `mailNo` | String | 是 | 运单号 | EMS12345 |
| `carrierPartnerCode` | String | 是 | 物流公司code | EMS |
| `transportStatus` | Number | 是 | 包裹运输状态 | 900 |
| `warehouseStatus` | Number | 否 | 包裹仓内状态 | 900100 |
| `actionTime` | Number | 是 | 包裹变更发生时间，ms | 1739951291000 |
| `tradeOrderIds` | Number[] | 是 | 包裹所属的订单id列表 | [123, 456] |

## 消息示例

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
