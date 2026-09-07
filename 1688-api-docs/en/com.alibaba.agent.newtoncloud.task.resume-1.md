# Newton Cloud: resume task

Original name: 牛顿云-恢复任务  
API: `com.alibaba.agent:newtoncloud.task.resume:1` · Category: Inquiries (Newton Cloud)  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.agent:newtoncloud.task.resume-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.agent/newtoncloud.task.resume/{appKey}`  
Requires user authorization (access_token) · Requires signature

Resume a Newton Cloud task.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `sessionId` | String | yes | sessionId | chat_03575ef4 |
| `scene` | String | yes | scene | open_api |
| `taskId` | String | yes | taskId | 68c83ae1-334a-46e2-9d26-e2db37cd1ba4 |
| `eventField` | [message:resumeEventObject](#m-resumeeventobject) | no | User-selected content; choose either this or selectedData — selectedData is recommended | {"type":"EXTERNAL_EXECUTION_RESULT","executionResults":[{"type":"tool_result","id":"call_33420b70e2654b77aadb61cb","name":"show_interaction","output":[{"type":"text","text":"{\"data\":[{\"question\":\"请问您想找什么材质的袜子？\",\"selected\":\"纯棉\"}]}","id":"ir_1784631153063"}],"state":"completed"}]} |
| `selectedData` | [message:questionAndSelected[]](#m-questionandselected[]) | no | User selection content | [{"question":"请问您需要采购多少件西装？","selected":"10件"},{"question":"请问您的预算范围是多少？","selected":"50-100元/件"},{"question":"请问您对西装的材质/风格有什么要求？","selected":"棉麻/休闲"}] |
| `skipped` | Boolean | no | Whether to skip; true means skip | true |
| `userInput` | String | no | User input requirement | 我想要便宜一些的 |

<a id="m-resumeeventobject"></a>
#### resumeEventObject

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `type` | String | yes | type | EXTERNAL_EXECUTION_RESULT |
| `executionResults` | [message:eventObject.executionResults[]](#m-eventobject-executionresults[]) | yes | User selection content | [{"type":"tool_result","id":"call_fab53346e8fe4103b726da70","name":"show_interaction","output":[{"type":"text","text":"{\"data\":[{\"question\":\"请问您需要采购多少双袜子？\",\"selected\":\"200双\"},{\"question\":\"您的预算大概是多少？\",\"selected\":\"30-80元\"},{\"question\":\"您有什么服务偏好吗？（可多选）\",\"selected\":\"无特别要求\"}]}","id":"ir_1784531665377"}],"state":"completed"}] |

<a id="m-eventobject-executionresults[]"></a>
#### eventObject.executionResults[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `type` | String | yes | tool_result | tool_result |
| `id` | String | yes | Result id | call_9f0a309daaa045b2844ee563 |
| `name` | String | yes | show_interaction | show_interaction |
| `state` | String | yes | Status | interrupted |
| `output` | [message:output[]](#m-output[]) | yes | Selected content | [{"type":"text","text":"{\"skipped\":true}","id":"ir_1784020010765"}] |

<a id="m-output[]"></a>
#### output[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `type` | String | yes | Type | text |
| `text` | String | yes | Content | {\"skipped\":true} |
| `id` | String | yes | Option ID | ir_1784020010765 |

<a id="m-questionandselected[]"></a>
#### questionAndSelected[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `question` | String | yes | Question | 请问您需要采购多少件西装？ |
| `selected` | String | yes | Selected content | 10件 |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:result](#m-result) | yes | Return result |     {       "success": true,       "data": {         "taskId": "858271d2-0817-4fbc-b019-32c60753cb18",         "status": "RUNNING",         "message": "task resumed"       },       "eagleTraceId": "2150980a17839485645261052e87dc"     }   |

<a id="m-result"></a>
#### result

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `success` | java.lang.Boolean | yes | Whether successful | true |
| `taskId` | java.lang.String | yes | Task ID | 858271d2-0817-4fbc-b019-32c60753cb18 |
| `status` | java.lang.String | yes | Status after resuming; normally RUNNING | RUNNING |
| `message` | java.lang.String | yes | Result description | task resumed |
| `error` | java.lang.String | yes | Error description on failure | task status is END, expected WAIT_SKILL or WAIT_USER |
| `eagleTraceId` | java.lang.String | yes | Trace ID | 2150980a17839485645261052e87dc |
| `errorCode` | String | yes | Failure error code. If this API is called while the task is still queued (status=QUEUED), it returns INVALID_REQUEST. In this case, poll task.get first and wait for the task to enter RUNNING before calling resume. | INVALID_REQUEST |

## Samples

**Input parameter example**

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

**Output parameter example**

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
