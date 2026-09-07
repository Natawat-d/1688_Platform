# 新增交易存证通知消息

Topic: `SYT_ADD_CONTRACT_CONTENT` · Group: SYT (生意通)  
Doc page: https://open.1688.com/doc/topicDetail.htm?id=SYT_ADD_CONTRACT_CONTENT

可用于：对手方在采购单/合同详情页面，新增物流凭证、发货凭证、付款凭证等信息时会做消息通知

## 消息字段

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `draftNo` | String | 是 | 合同号/采购单号 | 88SYT20240113169004 |
| `content` | String | 否 | 增加的具体内容，文本内容 | 无 |
| `eventType` | String | 是 | 当前事件类型 | CONTRACT_CONTENT_ADD_SUCCESS |
| `submitterRole` | String | 是 | 交易存证提交人角色 | PART_A:买家;PART_B:商家 |
| `addTime` | Number | 是 | 添加时间事件毫秒戳 | 1768288592321 |
| `attachments` | Object[] | 否 | 文件附件 | 无 |
| &nbsp;&nbsp;↳ `fileName` | String | 是 | 附件名称 | a.png |
| &nbsp;&nbsp;↳ `filePath` | String | 是 | 附件文件地址 | https://www.xxx.xxx/xxx.pdf |

## 消息示例

```json
{
  "draftNo": "88SYT20240113169004",
  "content": "无",
  "eventType": "CONTRACT_CONTENT_ADD_SUCCESS",
  "submitterRole": "PART_A:买家;PART_B:商家",
  "addTime": 1768288592321,
  "attachments": [
    {
      "fileName": "a.png",
      "filePath": "https://www.xxx.xxx/xxx.pdf"
    }
  ]
}
```
