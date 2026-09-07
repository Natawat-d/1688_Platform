# Get waybill-number set

Original name: 获取运单号集合  
API: `com.alibaba.fenxiao.crossborder:bigcustomer.mailNo.query:1` · Category: Logistics  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:bigcustomer.mailNo.query-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/bigcustomer.mailNo.query/{appKey}`  
Requires user authorization (access_token) · Requires signature

Get the set of waybill numbers.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `outOrderId` | java.lang.String | yes | Order id | 1234 |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.national.common.result.ResultModel](#m-alibaba-national-common-result-resultmodel) | yes | Return result | {} |

<a id="m-alibaba-national-common-result-resultmodel"></a>
#### alibaba.national.common.result.ResultModel

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `model` | java.util.Map | yes | The specific data returned | {} |
| `msg` | java.lang.String | yes | msg | 请求成功 |
| `traceId` | java.lang.String | yes | traceId | 2121212123344343423 |
| `success` | java.lang.Boolean | yes | Whether successful | true |
