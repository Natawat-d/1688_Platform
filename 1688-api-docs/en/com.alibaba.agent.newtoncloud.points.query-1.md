# Newton Cloud: query points details

Original name: 牛顿云-查询积分详情  
API: `com.alibaba.agent:newtoncloud.points.query:1` · Category: Inquiries (Newton Cloud)  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.agent:newtoncloud.points.query-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.agent/newtoncloud.points.query/{appKey}`  
Requires user authorization (access_token) · Requires signature

Query the Newton Cloud points information of the currently authorized account, including total points, used points, available points, points sources, expiry information and paginated usage details. The user is identified automatically from the access_token; no userId is needed.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `pageNo` | Integer | no | Page number for points consumption details, starting from 1 | 1 |
| `pageSize` | Integer | no | Number of records per page, maximum 100 | 20 |
| `startTime` | String | no | Query start time, format yyyy-MM-dd HH:mm:ss | 2026-08-01 00:00:00 |
| `endTime` | String | no | Query end time, format yyyy-MM-dd HH:mm:ss | 2026-08-12 23:59:59 |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:pointsQueryResponse](#m-pointsqueryresponse) | yes | Newton points query result | - |

<a id="m-pointsqueryresponse"></a>
#### pointsQueryResponse

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `success` | Boolean | yes | Whether the query was successful | true |
| `availableCredits` | BigDecimal | yes | Remaining points in the current account | 12230 |
| `pageNo` | Integer | yes | Current page number | 1 |
| `pageSize` | Integer | yes | Records per page | 20 |
| `total` | Integer | yes | Total number of conversations | 223 |
| `records` | [message:pointsSessionUsage[]](#m-pointssessionusage[]) | yes | Conversation points consumption record | - |
| `msgInfo` | String | yes | Error description when the query fails | 查询异常！ |
| `eagleTraceId` | String | yes | Request trace ID | eagletrace1213 |

<a id="m-pointssessionusage[]"></a>
#### pointsSessionUsage[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `sessionId` | String | yes | Newton Cloud session ID | session-1234 |
| `sessionName` | String | yes | Newton Cloud session name | 查询商品服务 |
| `creditsConsumed` | BigDecimal | yes | Cumulative points consumed in the conversation | 37.5 |
| `totalTokens` | Long | yes | Cumulative number of tokens consumed in the session | 37500 |

## Samples

**Input parameter example**

```
{
  "pageNo": 1,
  "pageSize": 20,
  "startTime": "2026-08-01 00:00:00",
  "endTime": "2026-08-12 23:59:59"
}
```

**Output parameter example**

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
