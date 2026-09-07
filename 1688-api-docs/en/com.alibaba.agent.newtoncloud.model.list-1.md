# Newton Cloud: list available model tiers

Original name: 牛顿云-查询可用模型档位列表  
API: `com.alibaba.agent:newtoncloud.model.list:1` · Category: Inquiries (Newton Cloud)  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.agent:newtoncloud.model.list-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.agent/newtoncloud.model.list/{appKey}`  
Requires user authorization (access_token) · Requires signature

Query the model tiers currently available on Newton Cloud. Returns each tier's code (id), display name (displayName) and other information. The tier code can be used as the model parameter of newtoncloud.task.create. No business parameters; only access_token and signature are required.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `success` | Boolean | yes | Whether the request was successful | true |
| `error` | String | yes | Error description on failure | 查询模型档位列表时发生内部错误，请稍后重试 |
| `models` | [message:model[]](#m-model[]) | yes | List of available model tiers. The returned content is driven by platform configuration and only includes tiers open to the public | [{"id":"flagship","displayName":"旗舰","isDefault":false}] |

<a id="m-model[]"></a>
#### model[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `id` | String | yes | Tier code, i.e. the value of the model parameter when creating a task (newtoncloud.task.create) | flagship |
| `displayName` | String | yes | Tier display name | 旗舰 |
| `description` | String | yes | Tier description | 能力最强 |
| `recommended` | Boolean | yes | Whether it is a recommended tier | false |
| `iconUrl` | String | yes | Tier icon URL | xxx |
| `isDefault` | Boolean | yes | Whether it is the default tier | true |
| `sortOrder` | Integer | yes | Sort weight (ascending order) | 1 |
