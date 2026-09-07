# Newton Cloud: query task table

Original name: 牛顿云-查询任务表格服务  
API: `com.alibaba.agent:newtoncloud.task.fetch:1` · Category: Inquiries (Newton Cloud)  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.agent:newtoncloud.task.fetch-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.agent/newtoncloud.task.fetch/{appKey}`  
Requires user authorization (access_token) · Requires signature

Newton Cloud task-table query service.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `taskId` | String | yes | Original task ID | 80c230b3-0cd1-4094-a223-27cc257bc8e0 |
| `id` | String | yes | Table instance ID | 69ec9bdcc878431c873f68e953c30e31_2631391132_20260714211145 |
| `scene` | String | yes | Use the value from complex_table.content; in procurement scenarios this is usually newton | newton |
| `stage` | String | no | Stage, e.g. recall / inquiry | recall |
| `subScene` | String | yes | Use the value from complex_table.content; the procurement scenario is usually purchase | purchase |
| `pageNo` | Integer | no | Requested page number, defaults to the first page | 1 |
| `pageSize` | Integer | no | Requested page size, default 10 | 10 |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:result](#m-result) | yes | Service-side response body, including whether the request succeeded, table data, error information, etc. | {       "success": true,       "data": {         "result": [           {             "itemId": "652731660850",             "title": "简约白色陶瓷马克杯可印LOGO简约家用水杯咖啡杯牛奶杯礼品杯子",             "realtimeSinglePrice": 4,             "realtimeDeliveryPrice": 5,             "realtimeTotalPrice": 8.9,             "realtimeQuantity": 1,             "salesCount": 1828,             "imageUrl": "https://cbu01.alicdn.com/img/...",             "company": "醴陵市鸿博瓷业有限公司",             "customerStar": "4.5分",             "sourceFactoryInfo": "产销资质：源头制造直供；员工规模：51~100人。"           }         ],         "total": 79,         "fieldSpec": [           {"key": "itemId", "type": "string"},           {"key": "title", "type": "string"},           {"key": "realtimeSinglePrice", "type": "number"},           {"key": "imageUrl", "type": "string"},           {"key": "company", "type": "string"}         ],         "eagleTraceId": "2146a21517799488479778174e1539"       },       "eagleTraceId": "2146a21517799488479778174e1539"     } |

<a id="m-result"></a>
#### result

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `success` | java.lang.Boolean | yes | Whether the request succeeded; true for success, false for failure | true |
| `error` | java.lang.String | yes | Error code when the request fails. INVALID_REQUEST=missing parameter, TASK_NOT_FOUND=task does not exist or no permission, UNSUPPORTED_IN_OPENAPI=sub-scenario not yet supported, CREDENTIAL_UNAVAILABLE=failed to obtain credential, UPSTREAM_TIMEOUT=upstream timeout, UPSTREAM_ERROR=upstream exception, INTERNAL_ERROR=internal error. Detailed error description when the request fails | TASK_NOT_FOUND |
| `message` | java.lang.String | yes | Error detail description when the request fails |  task xxx not found or not owned by userId xxx |
| `data` | java.lang.String | yes | Table data JSON string, needs to be used after JSON.parse. Contains result (product array), total (total count), fieldSpec (field name and type definitions, can be used to dynamically render table columns) | {"result":[{"itemId":"652731660850","title":"简约白色陶瓷马克杯","realtimeSinglePrice":4,"imageUrl":"https://cbu01.alicdn.com/...","company":"醴陵市鸿博瓷业有限公司","salesCount":1828}],"total":79,"fieldSpec":[{"key":"itemId","type":"string"},{"key":"title","type":"string"},{"key":"realtimeSinglePrice","type":"number"}]} |
| `eagleTraceId` | java.lang.String | yes | Unique request ID, used for tracing | 2146a21517799488479778174e1539 |

## Samples

**Input parameter example**

```
{
  "request": {
    "taskId": "60c1f539-fe88-49cf-9eb7-89e645896c25",
    "scene": "newton",
    "subScenc": "purchase",
    "id": "eeb49347fee946699f78de41218ee9aa_2631391132_20260713152648",
    "stage": "recall"
  },
  "userId": 2631391132
}
```

**Output parameter example**

```
{
  "result": {
    "success": true,
    "error": "TASK_NOT_FOUND",
    "message": " task xxx not found or not owned by userId xxx",
    "data": "{\"result\":[{\"itemId\":\"652731660850\",\"title\":\"简约白色陶瓷马克杯\",\"realtimeSinglePrice\":4,\"imageUrl\":\"https://cbu01.alicdn.com/...\",\"company\":\"醴陵市鸿博瓷业有限公司\",\"salesCount\":1828}],\"total\":79,\"fieldSpec\":[{\"key\":\"itemId\",\"type\":\"string\"},{\"key\":\"title\",\"type\":\"string\"},{\"key\":\"realtimeSinglePrice\",\"type\":\"number\"}]}",
    "eagleTraceId": "2146a21517799488479778174e1539"
  }
}
```
