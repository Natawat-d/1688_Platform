# 牛顿云任务执行通知队列

Topic: `AGENT_NEWTON_CLOUD_TASK_NOTIFY` · Group: AGENT (智能体)  
Doc page: https://open.1688.com/doc/topicDetail.htm?id=AGENT_NEWTON_CLOUD_TASK_NOTIFY

牛顿云任务执行通知队列，用户开启的长程任务状态发生改变后，会往该队列投递消息

## 消息字段

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `taskId` | String | 是 | 唯一标识任务的UUID | 0cabd613-7c43-4259-8230-834d7df7cdbf |
| `sessionId` | String | 是 | 唯一标识会话的UUID | f84b407a-c495-49a0-942b-a24c3985ca3f |
| `status` | String | 是 | 任务当前状态，例如 END 表示结束 | END |
| `createdAt` | String | 是 | 任务创建的时间戳（ISO 8601格式） | 2026-06-02T13:17:34.000Z |
| `finishedAt` | String | 否 | 任务完成的时间戳（ISO 8601格式） | 2026-06-02T13:17:38.000Z |
| `errorMessage` | String | 否 | 如果任务失败，包含具体的错误描述；成功时为 null | 用户手动中断 |
| `stage` | String | 否 | 中间态事件标识，如 WAIT_USER / TOOL_PROGRESS / COMPLEX_TABLE_READY / ERROR；终态通知无此字段 | WAIT_USER |
| `stageDetail` | String | 否 | 中间态详情，JSON格式扁平kv（≤256字节）；终态通知无此字段 | {"reason":"show_interaction","toolName":"search","toolCallId":"call_abc123","expectedEventType":"EXTERNAL_EXECUTION_RESULT"} |
| `payload` | String | 否 | 中间态负载，JSON 对象序列化（≤1024字节）；终态通知无此字段 | {"toolCallId":"call_abc123","toolName":"search","expectedEventType":"EXTERNAL_EXECUTION_RESULT","message":"请提供搜索关键词"} |

## 消息示例

```json
{"taskId":"0cabd613-7c43-4259-8230-834d7df7cdbf","sessionId":"f84b407a-c495-49a0-942b-a24c3985ca3f","status":"END","createdAt":"2026-06-02T13:17:34.000Z","finishedAt":"2026-06-02T13:17:38.000Z","errorMessage":"用户手动中断","stage":"WAIT_USER","stageDetail":"{"reason":"show_interaction","toolName":"search","toolCallId":"call_abc123","expectedEventType":"EXTERNAL_EXECUTION_RESULT"}","payload":"{"toolCallId":"call_abc123","toolName":"search","expectedEventType":"EXTERNAL_EXECUTION_RESULT","message":"请提供搜索关键词"}"}
```
