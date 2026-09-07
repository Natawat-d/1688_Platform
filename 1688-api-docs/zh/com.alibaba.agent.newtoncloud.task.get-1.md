# 牛顿云-查询任务服务

API: `com.alibaba.agent:newtoncloud.task.get:1` · Category: 询盘  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.agent:newtoncloud.task.get-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.agent/newtoncloud.task.get/{appKey}`  
需要授权 (access_token) · 需要签名

牛顿云-查询任务服务

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `taskId` | String | 是 | 业务侧生成的任务维度id，一个任务id一定归属于一个对话id | 009e1847-d652-4fa1-be5e-346b9b09ac38 |
| `fromIndex` | Long | 否 | 增量输出起始下标，首次传 0，后续传上次返回的 nextIndex | 0 |
| `includeBlocks` | Boolean | 否 | 是否返回富文本/块结构，获取 chunks 时建议传 true | true |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `success` | Boolean | 是 | 接口请求是否成功，true成功，false失败 | true |
| `taskId` | String | 是 | 业务侧生成的任务维度id，一个任务id一定归属于一个对话id | 009e1847-d652-4fa1-be5e-346b9b09ac38 |
| `sessionId` | String | 是 | 业务侧生成的对话维度id，一个对话id可挂接多个任务id | 2b6a7f8b-2c34-403c-9b9c-7e6121e50548 |
| `taskType` | String | 是 | 任务类型，LONG_RUNNING表示长程任务；NORMAL表示普通对话 | LONG_RUNNING |
| `status` | String | 是 | 任务状态。任务共有 6 种状态：INIT（已创建，任务进入队列等待执行）、RUNNING（正在执行中，Agent 正进行推理和工具调用）、WAIT_SKILL（Agent 调用了异步技能，任务暂停等待技能执行结果回调）、WAIT_USER（任务需要用户确认或补充输入，暂停等待用户通过 resume 接口恢复）、END（任务正常执行完成，已产出最终结果）、KILL（任务被终止，可能是用户主动取消、执行超时或系统异常导致的强制终止）。其中 END 和 KILL 为终态，一旦进入则不可再变更。 | RUNNING |
| `createdAt` | String | 是 | 任务创建时间 | 2026-06-02T13:11:57.000Z |
| `startedAt` | String | 是 | 任务开始时间 | 2026-06-02T13:11:58.000Z |
| `content` | String | 是 | 当前流式对话的输出，仅在任务非终态返回 | <aside>使用商家数据自由查询 Skill 获取昨日店铺经营数据</aside>\n\n" |
| `messages` | [message:message[]](#m-message[]) | 是 | 当前任务的所有ai返回消息列表，仅在任务终态返回 | [             {                 "type": "text",                 "content": "你好，我是牛顿（Newton），一个 helpful AI 助手。我可以帮你处理数据分析、文件操作、电商运营等多种任务。请问有什么我可以帮你的吗？"             }         ] |
| `errorMessage` | String | 是 | 任务失败描述 | 超时终止：长程任务执行超过 600s |
| `error` | String | 是 | 本次请求失败描述 | 系统内部异常 |
| `eagleTraceId` | String | 是 | 请求唯一ID ，eagleTraceId用于链路排查 | 2146a21517799488479778174e1539 |
| `nextIndex` | Long | 是 | 下一次轮询应传入的 fromIndex，用于增量获取后续输出块，避免重复拉取。 | 327 |
| `outputStatus` | String | 是 | 当前任务输出流状态。producing 表示仍在生成，done 表示输出完成，error 表示输出异常，wait_user 表示等待用户交互。 | wait_user |
| `chunks` | String | 是 | 增量输出块列表，JSON 数组字符串。每个元素含    type（text/thinking/tool_call/tool_result/complex_table/done/error/wait_user）、content（内容文本，type=complex_table 时为 JSON 字符串，可解析出    scene、subScene、id、stage 等表格查询参数）、timestamp（毫秒时间戳）。首次轮询 fromIndex 传 0，后续传上次返回的 nextIndex 实现增量拉取。 | [{"type":"complex_table","content":"{\"requestId\":\"60eed89b-5cd9-4649-9f55-67f2733d59d7\",\"scene\":\"newton\",\"subScene\":\"purchase\",\"id\":\     "eeb49347fee946699f78de41218ee9aa_2631391132_20260713152648\",\"stage\":\"recall\",\"__dataFilled\":false}","timestamp":1783927640343}] |
| `queuePosition` | Long | 是 | 当前排队位次，从 1 开始。仅当 status=QUEUED 时返回；若任务已不在队列中则该字段不返回 | 3 |
| `queueLength` | Long | 是 | 当前排队队列总长度（含本任务）。仅当 status=QUEUED 时返回 | 5 |
| `queuedTotal` | Long | 是 | 当前账号处于排队中（QUEUED）的任务总数。仅当 status=QUEUED 时返回 | 5 |
| `runningTotal` | Long | 是 | 当前账号正在执行中（INIT/RUNNING/WAIT_SKILL）的任务总数。仅当 status=QUEUED 时返回 | 2 |
| `finishedTotal` | Long | 是 | 当前账号已结束（END/KILL/ERROR）的任务总数。仅当 status=QUEUED 时返回 | 128 |
| `finishedAt` | String | 是 | 描述「任务结束时间，格式 yyyy-MM-dd HH:mm:ss，未结束时为空字符串」 | 2026-08-21 10:12:33 |

<a id="m-message[]"></a>
#### message[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `type` | String | 是 | 消息类型 | text |
| `content` | String | 是 | 消息内容 | 你好，我是牛顿（Newton），一个 helpful AI 助手。我可以帮你处理数据分析、文件操作、电商运营等多种任务。请问有什么我可以帮你的吗？ |

## 示例

**入参示例**

```
{
  "taskId": "009e1847-d652-4fa1-be5e-346b9b09ac38"
}
```

**出参示例**

```
{
  "success": true,
  "taskId": "009e1847-d652-4fa1-be5e-346b9b09ac38",
  "sessionId": "2b6a7f8b-2c34-403c-9b9c-7e6121e50548",
  "taskType": "LONG_RUNNING",
  "status": "RUNNING",
  "createdAt": "2026-06-02T13:11:57.000Z",
  "startedAt": "2026-06-02T13:11:58.000Z",
  "content": "<aside>使用商家数据自由查询 Skill 获取昨日店铺经营数据</aside>\\n\\n\"",
  "messages": [
    {
      "type": "text",
      "content": "你好，我是牛顿（Newton），一个 helpful AI 助手。我可以帮你处理数据分析、文件操作、电商运营等多种任务。请问有什么我可以帮你的吗？"
    }
  ],
  "errorMessage": "超时终止：长程任务执行超过 600s",
  "error": "系统内部异常",
  "eagleTraceId": "2146a21517799488479778174e1539"
}
```
