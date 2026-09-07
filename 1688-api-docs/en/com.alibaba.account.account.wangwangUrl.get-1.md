# Get a link that opens a WangWang chat

Original name: 获取唤起旺旺聊天的链接  
API: `com.alibaba.account:account.wangwangUrl.get:1` · Category: Tools  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.account:account.wangwangUrl.get-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.account/account.wangwangUrl.get/{appKey}`  
Requires user authorization (access_token) · Requires signature

Get a link that launches an AliWangWang chat.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `toOpenUid` | String | yes | Chat counterpart's openUid |   |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.ocean.openplatform.common.ResultModelGTlfvn7n](#m-alibaba-ocean-openplatform-common-resultmodelgtlfvn7n) | yes |  |  |

<a id="m-alibaba-ocean-openplatform-common-resultmodelgtlfvn7n"></a>
#### alibaba.ocean.openplatform.common.ResultModelGTlfvn7n

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `code` | java.lang.String | yes |  |   |
| `message` | java.lang.String | yes |  |   |
| `result` | java.lang.String | yes | WangWang chat link |   |
| `retCodes` | java.lang.String[] | yes |  |   |
| `subCode` | java.lang.String | yes |  |   |
| `subMessage` | java.lang.String | yes |  |   |
| `success` | boolean | yes |  |   |
