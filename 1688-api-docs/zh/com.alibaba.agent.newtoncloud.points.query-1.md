# 牛顿云-查询积分详情

API: `com.alibaba.agent:newtoncloud.points.query:1` · Category: 询盘  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.agent:newtoncloud.points.query-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.agent/newtoncloud.points.query/{appKey}`  
需要授权 (access_token) · 需要签名

查询当前授权账号的云牛顿积分信息，包括积分总额、已使用积分、可用积分、积分来源、积分到期信息及分页使用明细。用户身份由access_token 自动识别，无需传入 userId。

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `pageNo` | Integer | 否 | 积分消耗明细页码，从 1 开始 | 1 |
| `pageSize` | Integer | 否 | 每页记录数量，最大 100 | 20 |
| `startTime` | String | 否 | 查询开始时间，格式 yyyy-MM-dd HH:mm:ss | 2026-08-01 00:00:00 |
| `endTime` | String | 否 | 查询结束时间，格式 yyyy-MM-dd HH:mm:ss | 2026-08-12 23:59:59 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:pointsQueryResponse](#m-pointsqueryresponse) | 是 | 牛顿积分查询结果 | - |

<a id="m-pointsqueryresponse"></a>
#### pointsQueryResponse

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `success` | Boolean | 是 | 查询是否成功 | true |
| `availableCredits` | BigDecimal | 是 | 当前账户剩余积分 | 12230 |
| `pageNo` | Integer | 是 | 当前页码 | 1 |
| `pageSize` | Integer | 是 | 每页记录数 | 20 |
| `total` | Integer | 是 | 会话总数 | 223 |
| `records` | [message:pointsSessionUsage[]](#m-pointssessionusage[]) | 是 | 会话积分消耗记录 | - |
| `msgInfo` | String | 是 | 查询失败时的错误说明 | 查询异常！ |
| `eagleTraceId` | String | 是 | 请求链路 ID | eagletrace1213 |

<a id="m-pointssessionusage[]"></a>
#### pointsSessionUsage[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `sessionId` | String | 是 | 牛顿会话 ID | session-1234 |
| `sessionName` | String | 是 | 牛顿会话名称 | 查询商品服务 |
| `creditsConsumed` | BigDecimal | 是 | 会话累计消耗积分 | 37.5 |
| `totalTokens` | Long | 是 | 会话累计消耗 Token 数 | 37500 |

## 示例

**入参示例**

```
{
  "pageNo": 1,
  "pageSize": 20,
  "startTime": "2026-08-01 00:00:00",
  "endTime": "2026-08-12 23:59:59"
}
```

**出参示例**

```
{
  "result": {
    "success": true,
    "availableCredits": "12230",
    "pageNo": 1,
    "pageSize": 20,
    "total": 223,
    "records": [
      {
        "sessionId": "session-1234",
        "sessionName": "查询商品服务",
        "creditsConsumed": "37.5",
        "totalTokens": 37500
      }
    ],
    "msgInfo": "查询异常！",
    "eagleTraceId": "eagletrace1213"
  }
}
```
