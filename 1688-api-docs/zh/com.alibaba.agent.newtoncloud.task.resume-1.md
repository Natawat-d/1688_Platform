# 牛顿云-恢复任务

API: `com.alibaba.agent:newtoncloud.task.resume:1` · Category: 询盘  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.agent:newtoncloud.task.resume-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.agent/newtoncloud.task.resume/{appKey}`  
需要授权 (access_token) · 需要签名

牛顿云-恢复任务

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `sessionId` | String | 是 | sessionId | chat_03575ef4 |
| `scene` | String | 是 | scene | open_api |
| `taskId` | String | 是 | taskId | 68c83ae1-334a-46e2-9d26-e2db37cd1ba4 |
| `eventField` | [message:resumeEventObject](#m-resumeeventobject) | 否 | 用户选择内容,与selectedData二选一，建议使用selectedData | {"type":"EXTERNAL_EXECUTION_RESULT","executionResults":[{"type":"tool_result","id":"call_33420b70e2654b77aadb61cb","name":"show_interaction","output":[{"type":"text","text":"{\"data\":[{\"question\":\"请问您想找什么材质的袜子？\",\"selected\":\"纯棉\"}]}","id":"ir_1784631153063"}],"state":"completed"}]} |
| `selectedData` | [message:questionAndSelected[]](#m-questionandselected[]) | 否 | 用户选择内容 | [{"question":"请问您需要采购多少件西装？","selected":"10件"},{"question":"请问您的预算范围是多少？","selected":"50-100元/件"},{"question":"请问您对西装的材质/风格有什么要求？","selected":"棉麻/休闲"}] |
| `skipped` | Boolean | 否 | 是否跳过，true为跳过 | true |
| `userInput` | String | 否 | 用户输入需求 | 我想要便宜一些的 |

<a id="m-resumeeventobject"></a>
#### resumeEventObject

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `type` | String | 是 | type | EXTERNAL_EXECUTION_RESULT |
| `executionResults` | [message:eventObject.executionResults[]](#m-eventobject-executionresults[]) | 是 | 用户选择内容 | [{"type":"tool_result","id":"call_fab53346e8fe4103b726da70","name":"show_interaction","output":[{"type":"text","text":"{\"data\":[{\"question\":\"请问您需要采购多少双袜子？\",\"selected\":\"200双\"},{\"question\":\"您的预算大概是多少？\",\"selected\":\"30-80元\"},{\"question\":\"您有什么服务偏好吗？（可多选）\",\"selected\":\"无特别要求\"}]}","id":"ir_1784531665377"}],"state":"completed"}] |

<a id="m-eventobject-executionresults[]"></a>
#### eventObject.executionResults[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `type` | String | 是 | tool_result | tool_result |
| `id` | String | 是 | 结果id | call_9f0a309daaa045b2844ee563 |
| `name` | String | 是 | show_interaction | show_interaction |
| `state` | String | 是 | 状态 | interrupted |
| `output` | [message:output[]](#m-output[]) | 是 | 选择内容 | [{"type":"text","text":"{\"skipped\":true}","id":"ir_1784020010765"}] |

<a id="m-output[]"></a>
#### output[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `type` | String | 是 | 类型 | text |
| `text` | String | 是 | 内容 | {\"skipped\":true} |
| `id` | String | 是 | 选项id | ir_1784020010765 |

<a id="m-questionandselected[]"></a>
#### questionAndSelected[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `question` | String | 是 | 问题 | 请问您需要采购多少件西装？ |
| `selected` | String | 是 | 选择内容 | 10件 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:result](#m-result) | 是 | 返回结果 |     {       "success": true,       "data": {         "taskId": "858271d2-0817-4fbc-b019-32c60753cb18",         "status": "RUNNING",         "message": "task resumed"       },       "eagleTraceId": "2150980a17839485645261052e87dc"     }   |

<a id="m-result"></a>
#### result

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `success` | java.lang.Boolean | 是 | 是否成功 | true |
| `taskId` | java.lang.String | 是 | 任务ID | 858271d2-0817-4fbc-b019-32c60753cb18 |
| `status` | java.lang.String | 是 | 恢复后状态，正常为 RUNNING | RUNNING |
| `message` | java.lang.String | 是 | 结果描述 | task resumed |
| `error` | java.lang.String | 是 | 失败时的错误描述 | task status is END, expected WAIT_SKILL or WAIT_USER |
| `eagleTraceId` | java.lang.String | 是 | 链路追踪ID | 2150980a17839485645261052e87dc |
| `errorCode` | String | 是 | 失败错误码。任务仍在排队中（status=QUEUED）时调用本接口会返回 INVALID_REQUEST，此时应先轮询 task.get 等待任务进入 RUNNING 后再 resume | INVALID_REQUEST |

## 示例

**入参示例**

```
{
  "sessionId": "chat_03575ef4",
  "scene": "open_api",
  "taskId": "68c83ae1-334a-46e2-9d26-e2db37cd1ba4",
  "eventField": {
    "type": "EXTERNAL_EXECUTION_RESULT",
    "executionResults": [
      {
        "type": "tool_result",
        "id": "call_9f0a309daaa045b2844ee563",
        "name": "show_interaction",
        "state": "interrupted",
        "output": [
          {
            "type": "text",
            "text": "{\\\"skipped\\\":true}",
            "id": "ir_1784020010765"
          }
        ]
      }
    ]
  },
  "selectedData": [
    {
      "question": "请问您需要采购多少件西装？",
      "selected": "10件"
    }
  ],
  "skipped": true,
  "userInput": "我想要便宜一些的"
}
```

**出参示例**

```
{
  "result": {
    "success": true,
    "taskId": "858271d2-0817-4fbc-b019-32c60753cb18",
    "status": "RUNNING",
    "message": "task resumed",
    "error": "task status is END, expected WAIT_SKILL or WAIT_USER",
    "eagleTraceId": "2150980a17839485645261052e87dc"
  }
}
```
