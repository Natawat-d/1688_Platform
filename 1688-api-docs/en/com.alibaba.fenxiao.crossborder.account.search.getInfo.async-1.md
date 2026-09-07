# Run AI supplier-search task (async)

Original name: 异步执行AI找商任务  
API: `com.alibaba.fenxiao.crossborder:account.search.getInfo.async:1` · Category: Suppliers  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:account.search.getInfo.async-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/account.search.getInfo.async/{appKey}`  
Requires user authorization (access_token) · Requires signature

Execute an AI supplier-search task asynchronously.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `request` | [message:com.alibaba.global1688.common.request.SpApiExecuteRequest](#m-com-alibaba-global1688-common-request-spapiexecuterequest) | yes | Request parameters | {} |

<a id="m-com-alibaba-global1688-common-request-spapiexecuterequest"></a>
#### com.alibaba.global1688.common.request.SpApiExecuteRequest

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `query` | String | yes | Query statement | 毛巾 |
| `spFilterCondition` | [message:com.alibaba.global1688.common.request.SpFilterCondition](#m-com-alibaba-global1688-common-request-spfiltercondition) | no | Filter condition | 过滤条件 |

<a id="m-com-alibaba-global1688-common-request-spfiltercondition"></a>
#### com.alibaba.global1688.common.request.SpFilterCondition

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `businessModel` | String | no | Business mode; only supports passing factory (indicates a source factory) | factory |
| `establishedYearsMin` | Integer | no | Minimum years established (filters for merchants established for ≥ this many years) | 3 |
| `repeatPurchaseRateMin` | Integer | no | Minimum 30-day repeat purchase rate, range [0, 100] | 50 |
| `supportCustomization` | Boolean | no | Whether processing/customization is supported; only true can be passed | true |
| `afterSalesScoreMin` | Double | no | Minimum after-sales experience score, range [0, 5] | 4.0 |
| `productQualityScoreMin` | Double | no | Minimum product experience score, range [0, 5] | 4.0 |
| `consultationResponseScoreMin` | Double | no | Minimum inquiry experience score, range [0, 5] | 4.0 |
| `logisticsTimelinessScoreMin` | Double | no | Minimum logistics experience score, range [0, 5] | 4.0 |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:com.alibaba.global1688.common.response.ApiResultExecute](#m-com-alibaba-global1688-common-response-apiresultexecute) | yes | Return message | {} |

<a id="m-com-alibaba-global1688-common-response-apiresultexecute"></a>
#### com.alibaba.global1688.common.response.ApiResultExecute

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `success` | Boolean | yes | Whether the AI product-sourcing task was successfully initiated | true |
| `result` | String | yes | Unique task identifier; required when querying task status | xxxxxx |
| `code` | String | yes | Status code | SUCCESS |
| `message` | String | yes | Exception info | error_msg_xxxxx |

## Samples

**Input parameter example**

```
{
    "spFilterCondition": {
        "repeatPurchaseRateMin": 70,
        "consultationResponseScoreMin": 3.5,
        "establishedYearsMin": 6,
        "productQualityScoreMin": 3.5,
        "businessModel": "factory",
        "afterSalesScoreMin": 3.5,
        "logisticsTimelinessScoreMin": 3.5
    },
    "query": "连衣裙"
}
```

**Output parameter example**

```
{
  "result": "taskIdxxxxxxxxxx",
  "code": "SUCCESS",
  "success": true
}
```
