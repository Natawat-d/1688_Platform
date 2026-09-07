# Newton Cloud task execution notification queue

Original name: 牛顿云任务执行通知队列  
Topic: `AGENT_NEWTON_CLOUD_TASK_NOTIFY` · Group: AGENT (智能体)  
Doc page: https://open.1688.com/doc/topicDetail.htm?id=AGENT_NEWTON_CLOUD_TASK_NOTIFY

Newton Cloud task execution notification queue; when the status of a long-running task started by the user changes, a message is delivered to this queue

## Payload fields

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `taskId` | String | yes | UUID that uniquely identifies the task | 0cabd613-7c43-4259-8230-834d7df7cdbf |
| `sessionId` | String | yes | UUID that uniquely identifies the session | f84b407a-c495-49a0-942b-a24c3985ca3f |
| `status` | String | yes | Current status of the task, e.g. END indicates completion | END |
| `createdAt` | String | yes | Timestamp when the task was created (ISO 8601 format) | 2026-06-02T13:17:34.000Z |
| `finishedAt` | String | no | Timestamp when the task was completed (ISO 8601 format) | 2026-06-02T13:17:38.000Z |
| `errorMessage` | String | no | If the task fails, contains a specific error description; null on success | 用户手动中断 |
| `stage` | String | no | Intermediate-state event identifier, e.g. WAIT_USER / TOOL_PROGRESS / COMPLEX_TABLE_READY / ERROR; this field is absent in final-state notifications | WAIT_USER |
| `stageDetail` | String | no | Intermediate-state details, flat key-value pairs in JSON format (≤256 bytes); this field is absent in final-state notifications | {"reason":"show_interaction","toolName":"search","toolCallId":"call_abc123","expectedEventType":"EXTERNAL_EXECUTION_RESULT"} |
| `payload` | String | no | Intermediate-state payload, serialized as a JSON object (≤1024 bytes); this field is absent in final-state notifications | {"toolCallId":"call_abc123","toolName":"search","expectedEventType":"EXTERNAL_EXECUTION_RESULT","message":"请提供搜索关键词"} |

## Sample message

```json
{"taskId":"0cabd613-7c43-4259-8230-834d7df7cdbf","sessionId":"f84b407a-c495-49a0-942b-a24c3985ca3f","status":"END","createdAt":"2026-06-02T13:17:34.000Z","finishedAt":"2026-06-02T13:17:38.000Z","errorMessage":"用户手动中断","stage":"WAIT_USER","stageDetail":"{"reason":"show_interaction","toolName":"search","toolCallId":"call_abc123","expectedEventType":"EXTERNAL_EXECUTION_RESULT"}","payload":"{"toolCallId":"call_abc123","toolName":"search","expectedEventType":"EXTERNAL_EXECUTION_RESULT","message":"请提供搜索关键词"}"}
```
