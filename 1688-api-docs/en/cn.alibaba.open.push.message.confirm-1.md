# Batch-confirm failed messages

Original name: 失败消息批量确认  
API: `cn.alibaba.open:push.message.confirm:1` · Category: Messaging  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=cn.alibaba.open:push.message.confirm-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/cn.alibaba.open/push.message.confirm/{appKey}`  
No user authorization · Requires signature

Manually call the confirmation API to confirm that messages have been consumed successfully. Only needed when using the query-style API to fetch failed messages.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | no | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `msgIdList` | java.util.List | no | List of message IDs pending confirmation | [123,456] |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `isSuccess` | boolean | yes | Whether the operation succeeded | true |
