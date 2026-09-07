# 牛顿云-创建长程任务

API: `com.alibaba.agent:newtoncloud.task.create:1` · Category: 询盘  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.agent:newtoncloud.task.create-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.agent/newtoncloud.task.create/{appKey}`  
需要授权 (access_token) · 需要签名

创建一个牛顿云长程任务，异步执行，返回 taskId/sessionId/status

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `message` | String | 是 | 开启任务时，用户发起的问题 | 你有哪些skill供我调用？ |
| `sessionId` | String | 否 | 1688、用户生成的对话维度id，一个sessionId可挂接多个taskId | internal_a1b2c3d4-e5f6-7890-abcd-ef1234567890 |
| `taskId` | String | 否 | 1688、用户生成的任务维度id，一个taskId一定归属于一个对话消息 | fdd93e89-2f43-4d58-b5b4-356da6cda5e4 |
| `auto` | Boolean | 否 | 是否自动模式 | true/false |
| `fileUrls` | String[] | 否 | 用户上传文件的URL列表 | https://selleragent.1688.com/api/seller/knowledge/file/download/xxx |
| `fileNames` | String[] | 否 | 文件名称列表，与fileUrls一一对应 | 11111.jpg |
| `model` | String | 否 | 可选择的模型 | flagship |
| `workflowName` | String | 否 | 指定要直接执行的 Workflow 名称。传入后将跳过主 Agent 的意图识别和规划过程，直接执行对应 Workflow。名称不存在时任务执行失败。 | 1688-supplychain-procurement-search |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `success` | Boolean | 是 | 请求是否成功 | true |
| `taskId` | String | 是 | 1688、用户生成的任务维度id，一个taskId一定归属于一个对话消息 | fdd93e89-2f43-4d58-b5b4-356da6cda5e4 |
| `sessionId` | String | 是 | 1688、用户生成的对话维度id，一个sessionId可挂接多个taskId | internal_a1b2c3d4-e5f6-7890-abcd-ef1234567890 |
| `status` | String | 是 | 任务状态。任务共有 6 种状态：INIT（已创建，任务进入队列等待执行）、RUNNING（正在执行中，Agent 正进行推理和工具调用）、WAIT_SKILL（Agent 调用了异步技能，任务暂停等待技能执行结果回调）、WAIT_USER（任务需要用户确认或补充输入，暂停等待用户通过 resume 接口恢复）、END（任务正常执行完成，已产出最终结果）、KILL（任务被终止，可能是用户主动取消、执行超时或系统异常导致的强制终止）。其中 END 和 KILL 为终态，一旦进入则不可再变更。 | INIT |
| `error` | String | 是 | 任务失败描述 | 当前有 N 个任务正在执行，已达上限 Y，请等待任务完成后再提交 |
| `eagleTraceId` | String | 是 | 请求唯一ID ，eagleTraceId用于链路排查 | 2146a21517799488479778174e1539 |
| `errorCode` | String | 是 | 失败错误码，仅当 success=false 时返回。取值：RATE_LIMIT_CONCURRENT（并发已满）、RATE_LIMIT_PENDING_TOTAL（排队总量超限）、RATE_LIMIT_QUEUE_FULL（队列已满）、RATE_LIMIT_HOURLY（小时配额用尽）、RATE_LIMIT_DAILY（日配额用尽） | RATE_LIMIT_QUEUE_FULL |
