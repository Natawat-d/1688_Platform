# 合同/采购单双方支付成功消息通知

Topic: `SYT_CONTRACT_PAY_SUCCESS` · Group: SYT (生意通)  
Doc page: https://open.1688.com/doc/topicDetail.htm?id=SYT_CONTRACT_PAY_SUCCESS

合同/采购单双方支付成功消息通知

## 消息字段

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `draftNo` | String | 是 | 合同号或采购单号 | 88SYT20230113169004 |
| `contractCurrentStatus` | String | 是 | 合同或采购单当前状态 | PAID |
| `eventType` | String | 是 | 当前事件类型 | CONTRACT_PAY_SUCCESS_EVENT |
| `payChannel` | String | 是 | 买家支付渠道 | BANK_TRANSFER |

## 消息示例

```json
{
  "draftNo": "88SYT20230113169004",
  "contractCurrentStatus": "PAID",
  "eventType": "CONTRACT_PAY_SUCCESS_EVENT",
  "payChannel": "BANK_TRANSFER"
}
```
