# Newton Cloud: query batch-inquiry results

Original name: 牛顿云-查询批量询盘结果  
API: `com.alibaba.agent:newtoncloud.batchInquiry.getResult:1` · Category: Inquiries (Newton Cloud)  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.agent:newtoncloud.batchInquiry.getResult-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.agent/newtoncloud.batchInquiry.getResult/{appKey}`  
Requires user authorization (access_token) · Requires signature

Query inquiry results by batch-inquiry task ID (wwTaskId). After the ISV receives an AGENT_NEWTON_CLOUD_TASK_NOTIFY notification (stage=BATCH_INQUIRY, stageDetail.phase=COMPLETE), call this interface with the stageDetail.wwTaskId from the notification. It returns structured inquiry results synchronously, including supplier reply summaries, quotation details and recommendation reasons.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `wwTaskId` | String | no | Batch inquiry task ID | 1784790418611 |
| `taskId` | String | no | Newton Cloud task ID; pass this in when there is no wwTaskId, and the API automatically queries the inquiry result based on the Newton Cloud task ID | e0e5e0c0-6911-4b5a-9082-140474666exxx |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `success` | Boolean | yes | Whether the request was successful | true |
| `data` | String | yes | Inquiry result JSON string | {"taskId":"1784790418611",":"订单咨询","status":"SUCCESS","result":{"summary":[{"question":"是否支持一件代发","answer":"支持"}]}}]}]}status":"SUCCESS","questions":["是否支持一件代发"],"subTasks":[{"topics":[{"topicName" |
| `eagleTraceId` | String | yes | Unique request ID; eagleTraceId is used for trace troubleshooting | 21089b3017847890685524618e1299 |
| `error` | String | yes | Query failure description | request.wwTaskId required |
| `inquiryStatus` | String | yes | Batch inquiry task status | success |

## Samples

**Input parameter example**

```
{
  "wwTaskId": "1784790418611"
}
```

**Output parameter example**

```
{
  "success": true,
  "data": "{\"taskId\":\"1784790418611\",\":\"订单咨询\",\"status\":\"SUCCESS\",\"result\":{\"summary\":[{\"question\":\"是否支持一件代发\",\"answer\":\"支持\"}]}}]}]}status\":\"SUCCESS\",\"questions\":[\"是否支持一件代发\"],\"subTasks\":[{\"topics\":[{\"topicName\"",
  "eagleTraceId": "21089b3017847890685524618e1299",
  "error": "request.wwTaskId required"
}
```
