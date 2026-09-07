# 牛顿云-查询批量询盘结果

API: `com.alibaba.agent:newtoncloud.batchInquiry.getResult:1` · Category: 询盘  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.agent:newtoncloud.batchInquiry.getResult-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.agent/newtoncloud.batchInquiry.getResult/{appKey}`  
需要授权 (access_token) · 需要签名

根据批量询盘任务ID（wwTaskId）查询询盘结果。ISV 收到 AGENT_NEWTON_CLOUD_TASK_NOTIFY通知（stage=BATCH_INQUIRY，stageDetail.phase=COMPLETE）后，使用通知中 stageDetail.wwTaskId调用本接口，同步返回结构化询盘结果，包含商家回复摘要、报价信息、推荐理由等

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `wwTaskId` | String | 否 | 批量询盘任务ID | 1784790418611 |
| `taskId` | String | 否 | 牛顿云任务ID，没有wwTaskId时传入，API根据牛顿云任务ID自动查询询盘结果 | e0e5e0c0-6911-4b5a-9082-140474666exxx |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `success` | Boolean | 是 | 请求是否成功 | true |
| `data` | String | 是 | 询盘结果JSON字符串 | {"taskId":"1784790418611",":"订单咨询","status":"SUCCESS","result":{"summary":[{"question":"是否支持一件代发","answer":"支持"}]}}]}]}status":"SUCCESS","questions":["是否支持一件代发"],"subTasks":[{"topics":[{"topicName" |
| `eagleTraceId` | String | 是 | 请求唯一ID，eagleTraceId用于链路排查 | 21089b3017847890685524618e1299 |
| `error` | String | 是 | 查询失败描述 | request.wwTaskId required |
| `inquiryStatus` | String | 是 | 批量询盘任务状态 | success |

## 示例

**入参示例**

```
{
  "wwTaskId": "1784790418611"
}
```

**出参示例**

```
{
  "success": true,
  "data": "{\"taskId\":\"1784790418611\",\":\"订单咨询\",\"status\":\"SUCCESS\",\"result\":{\"summary\":[{\"question\":\"是否支持一件代发\",\"answer\":\"支持\"}]}}]}]}status\":\"SUCCESS\",\"questions\":[\"是否支持一件代发\"],\"subTasks\":[{\"topics\":[{\"topicName\"",
  "eagleTraceId": "21089b3017847890685524618e1299",
  "error": "request.wwTaskId required"
}
```
