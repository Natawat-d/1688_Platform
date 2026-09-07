# 合同或采购单待签署消息

Topic: `SYT_CONTRACT_WAIT_SIGN` · Group: SYT (生意通)  
Doc page: https://open.1688.com/doc/topicDetail.htm?id=SYT_CONTRACT_WAIT_SIGN

当合同或采购单需要签署确认时，会做消息通知

## 消息字段

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `draftNo` | String | 是 | 合同号或采购单号 | 88SYT20230113169004 |
| `contractCurrentStatus` | String | 是 | 合同或采购单当前状态 | SIGNING |
| `eventType` | String | 是 | 当前事件类型 | CONTRACT_WAIT_SIGN_EVENT |

## 消息示例

```json
{
  "draftNo": "88SYT20230113169004",
  "contractCurrentStatus": "SIGNING",
  "eventType": "CONTRACT_WAIT_SIGN_EVENT"
}
```
