# 采购单/合同拒绝消息通知

Topic: `SYT_CONTRACT_REJECT` · Group: SYT (生意通)  
Doc page: https://open.1688.com/doc/topicDetail.htm?id=SYT_CONTRACT_REJECT

生意通采购单/合同拒绝消息通知

## 消息字段

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `draftNo` | String | 是 | 合同号或采购单号 | 88SYT20230113169004 |
| `contractCurrentStatus` | String | 是 | 合同当前状态 | SIGN_REJECT |
| `eventType` | String | 是 | 当前事件类型 | CONTRACT_SIGN_REJECT_EVENT |
| `signRejectReason` | String | 是 | 拒绝原因 | 无 |
| `signRejectTime` | String | 是 | 拒绝事件毫秒戳 | 1778505510000 |

## 消息示例

```json
{
  "draftNo": "88SYT20230113169004",
  "contractCurrentStatus": "SIGN_REJECT",
  "eventType": "CONTRACT_SIGN_REJECT_EVENT",
  "signRejectReason": "无",
  "signRejectTime": "1778505510000"
}
```
