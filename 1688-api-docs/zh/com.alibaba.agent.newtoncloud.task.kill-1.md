# 牛顿云-终止任务

API: `com.alibaba.agent:newtoncloud.task.kill:1` · Category: 询盘  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.agent:newtoncloud.task.kill-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.agent/newtoncloud.task.kill/{appKey}`  
需要授权 (access_token) · 需要签名

终止指定 taskId 的任务，标记 KILL 状态，返回 killed

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `taskId` | String | 是 | 业务侧生成的任务维度id，一个任务id一定归属于一个对话id | 009e1847-d652-4fa1-be5e-346b9b09ac3 |
| `reason` | String | 是 | 终止任务原因，仅作数据记录 | 测试收集样例，主动终止 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `success` | Boolean | 是 | 接口是否成功 | true |
| `taskId` | String | 是 | 业务侧生成的任务维度id，一个任务id一定归属于一个对话id | 009e1847-d652-4fa1-be5e-346b9b09ac38 |
| `killed` | Boolean | 是 | 是否终止成功 | true |
| `error` | String | 是 | 任务失败描述 | 系统内部异常 |
| `eagleTraceId` | String | 是 | 请求唯一ID ，eagleTraceId用于链路排查 | 2146a21517799488479778174e1539 |
| `message` | String | 是 | 终止结果说明。排队中任务被取消时返回「排队中的任务已取消」，运行中任务被中断时返回「任务已终止」，失败时返回原因 | 排队中的任务已取消	 |
