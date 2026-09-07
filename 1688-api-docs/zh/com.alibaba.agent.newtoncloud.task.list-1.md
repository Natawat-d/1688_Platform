# 牛顿云-查询任务列表

API: `com.alibaba.agent:newtoncloud.task.list:1` · Category: 询盘  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.agent:newtoncloud.task.list-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.agent/newtoncloud.task.list/{appKey}`  
需要授权 (access_token) · 需要签名

查询当前用户的任务列表，返回每个任务的 taskId/sessionId/status/taskType/创建与完成时间

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `pageNo` | Integer | 否 | 分页页数，不传该参数时为全量查询 | 1 |
| `pageSize` | Integer | 否 | 分页每页数量 | 20 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `success` | Boolean | 是 | 接口是否成功 | true |
| `data` | [message:taskItem[]](#m-taskitem[]) | 是 | 任务列表，每一项为一个任务的详细信息。详情可参考 查询任务服务 | [         {             "taskId": "7fe59521-cf41-4df6-aae7-24e95c04dabc",             "sessionId": "87ea8842-3378-4081-905b-5ec06fe3aa56",             "status": "END",             "taskType": "LONG_RUNNING",             "createdAt": "2026-06-02T07:19:42.000Z",             "finishedAt": "2026-06-03T10:16:56.000Z"         } 	] |
| `error` | String | 是 | 任务失败描述 | 系统内部异常 |
| `eagleTraceId` | String | 是 | 请求唯一ID ，eagleTraceId用于链路排查 | 2146a21517799488479778174e1539 |
| `total` | Integer | 是 | 任务总数 | 100 |

<a id="m-taskitem[]"></a>
#### taskItem[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `taskId` | String | 是 | 业务侧生成的任务维度id，一个任务id一定归属于一个对话id | 009e1847-d652-4fa1-be5e-346b9b09ac38 |
| `sessionId` | String | 是 | 业务侧生成的对话维度id，一个对话id可挂接多个任务id | 2b6a7f8b-2c34-403c-9b9c-7e6121e50548 |
| `taskType` | String | 是 | 任务类型，LONG_RUNNING表示长程任务；NORMAL表示普通对话 | LONG_RUNNING |
| `status` | String | 是 | 任务状态。任务共有 6 种状态：INIT（已创建，任务进入队列等待执行）、RUNNING（正在执行中，Agent 正进行推理和工具调用）、WAIT_SKILL（Agent 调用了异步技能，任务暂停等待技能执行结果回调）、WAIT_USER（任务需要用户确认或补充输入，暂停等待用户通过 resume 接口恢复）、END（任务正常执行完成，已产出最终结果）、KILL（任务被终止，可能是用户主动取消、执行超时或系统异常导致的强制终止）。其中 END 和 KILL 为终态，一旦进入则不可再变更。 | RUNNING |
| `createdAt` | String | 是 | 任务创建时间 | 2026-06-09 20:32:26 |
| `finishedAt` | String | 是 | 任务开始时间 | 2026-06-09 20:43:57 |
| `extra` | String | 是 | 额外信息(json) | {"message":"你好"} |

## 示例

**入参示例**

```
{
  "pageNo": 1,
  "pageSize": 20
}
```

**出参示例**

```
{
  "success": true,
  "data": [
    {
      "taskId": "009e1847-d652-4fa1-be5e-346b9b09ac38",
      "sessionId": "2b6a7f8b-2c34-403c-9b9c-7e6121e50548",
      "taskType": "LONG_RUNNING",
      "status": "RUNNING",
      "createdAt": "2026-06-09 20:32:26",
      "finishedAt": "2026-06-09 20:43:57",
      "extra": "{\"message\":\"你好\"}"
    }
  ],
  "error": "系统内部异常",
  "eagleTraceId": "2146a21517799488479778174e1539",
  "total": 100
}
```
