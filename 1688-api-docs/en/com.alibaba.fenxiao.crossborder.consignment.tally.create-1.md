# Warehouse creates discrepancy tally sheet

Original name: 仓库创建差异理货单  
API: `com.alibaba.fenxiao.crossborder:consignment.tally.create:1` · Category: Fully Managed (Consignment)  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:consignment.tally.create-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/consignment.tally.create/{appKey}`  
No user authorization · Requires signature

The Buffalo warehouse creates a discrepancy tally sheet.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | no | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `bizInboundOrderId` | java.lang.String | yes | Inbound (warehouse receipt) order id |  |
| `status` | java.lang.String | yes | tallysheet |  |
| `processTime` | java.lang.Long | yes | Timestamp |  |
| `tallySheetId` | java.lang.String | yes | Warehouse discrepancy order ID |  |
| `tallySheetReason` | java.lang.String | yes | Reason for tally discrepancy |  |
| `inboundCount` | java.lang.String | yes | Inbound quantity |  |
| `varianceCount` | java.lang.String | yes | Discrepancy quantity |  |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.national.common.result.ResultModel](#m-alibaba-national-common-result-resultmodel) | yes |  |  |

<a id="m-alibaba-national-common-result-resultmodel"></a>
#### alibaba.national.common.result.ResultModel

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `msg` | java.lang.String | yes | Error message | 参数为空 |
| `traceId` | java.lang.String | yes | traceId | 111111 |
| `success` | java.lang.String | yes | true/false | true |
| `model` | java.lang.Object | yes | null | null |
