# 牛顿云-查询任务表格服务

API: `com.alibaba.agent:newtoncloud.task.fetch:1` · Category: 询盘  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.agent:newtoncloud.task.fetch-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.agent/newtoncloud.task.fetch/{appKey}`  
需要授权 (access_token) · 需要签名

牛顿云-查询任务表格服务

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `taskId` | String | 是 | 原任务 ID | 80c230b3-0cd1-4094-a223-27cc257bc8e0 |
| `id` | String | 是 | 表格实例 ID | 69ec9bdcc878431c873f68e953c30e31_2631391132_20260714211145 |
| `scene` | String | 是 | 使用 complex_table.content 中的值，采购场景通常是 newton | newton |
| `stage` | String | 否 | 阶段，如 recall / inquiry | recall |
| `subScene` | String | 是 | 使用 complex_table.content 中的值，采购场景通常是 purchase | purchase |
| `pageNo` | Integer | 否 | 请求页数，默认第一页 | 1 |
| `pageSize` | Integer | 否 | 请求页大小，默认10 | 10 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:result](#m-result) | 是 | 务端响应体，包含请求是否成功、表格数据、错误信息等 | {       "success": true,       "data": {         "result": [           {             "itemId": "652731660850",             "title": "简约白色陶瓷马克杯可印LOGO简约家用水杯咖啡杯牛奶杯礼品杯子",             "realtimeSinglePrice": 4,             "realtimeDeliveryPrice": 5,             "realtimeTotalPrice": 8.9,             "realtimeQuantity": 1,             "salesCount": 1828,             "imageUrl": "https://cbu01.alicdn.com/img/...",             "company": "醴陵市鸿博瓷业有限公司",             "customerStar": "4.5分",             "sourceFactoryInfo": "产销资质：源头制造直供；员工规模：51~100人。"           }         ],         "total": 79,         "fieldSpec": [           {"key": "itemId", "type": "string"},           {"key": "title", "type": "string"},           {"key": "realtimeSinglePrice", "type": "number"},           {"key": "imageUrl", "type": "string"},           {"key": "company", "type": "string"}         ],         "eagleTraceId": "2146a21517799488479778174e1539"       },       "eagleTraceId": "2146a21517799488479778174e1539"     } |

<a id="m-result"></a>
#### result

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `success` | java.lang.Boolean | 是 | 请求是否成功，true 成功，false 失败 | true |
| `error` | java.lang.String | 是 | 请求失败时的错误码。INVALID_REQUEST=参数缺失、TASK_NOT_FOUND=任务不存在或无权限、UNSUPPORTED_IN_OPENAPI=子场景暂不支持、CREDENTIAL_UNAVAILABLE=凭证获取失败、UPSTREAM_TIMEOUT=上游超时、UPSTREAM_ERROR=上游异常、INTERNAL_ERROR=内部错误请求失败时的错误详情描述 | TASK_NOT_FOUND |
| `message` | java.lang.String | 是 | 请求失败时的错误详情描述 |  task xxx not found or not owned by userId xxx |
| `data` | java.lang.String | 是 | 表格数据 JSON 字符串，需 JSON.parse 后使用。内含result（商品数组）、total（总条数）、fieldSpec（字段名与类型定义，可用于动态渲染表格列） | {"result":[{"itemId":"652731660850","title":"简约白色陶瓷马克杯","realtimeSinglePrice":4,"imageUrl":"https://cbu01.alicdn.com/...","company":"醴陵市鸿博瓷业有限公司","salesCount":1828}],"total":79,"fieldSpec":[{"key":"itemId","type":"string"},{"key":"title","type":"string"},{"key":"realtimeSinglePrice","type":"number"}]} |
| `eagleTraceId` | java.lang.String | 是 | 请求唯一ID，用于链路排查 | 2146a21517799488479778174e1539 |

## 示例

**入参示例**

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

**出参示例**

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
