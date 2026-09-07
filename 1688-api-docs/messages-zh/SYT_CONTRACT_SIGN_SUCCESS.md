# 合同/采购单双方签署成功通知

Topic: `SYT_CONTRACT_SIGN_SUCCESS` · Group: SYT (生意通)  
Doc page: https://open.1688.com/doc/topicDetail.htm?id=SYT_CONTRACT_SIGN_SUCCESS

合同/采购单双方签署成功后，会做消息通知

## 消息字段

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `draftNo` | String | 是 | 合同号或采购单号 | 88SYT20230113169004 |
| `contractCurrentStatus` | String | 是 | 合同/采购单当前状态 | SIGN_SUCCESS |
| `eventType` | String | 是 | 当前事件类型 | CONTRACT_SIGN_SUCCESS_EVENT |
| `signSuccessTime` | String | 是 | 签约成功时间毫秒戳 | 1778505510000 |

## 消息示例

```json
{
  "draftNo": "88SYT20230113169004",
  "contractCurrentStatus": "SIGN_SUCCESS",
  "eventType": "CONTRACT_SIGN_SUCCESS_EVENT",
  "signSuccessTime": "1778505510000"
}
```
