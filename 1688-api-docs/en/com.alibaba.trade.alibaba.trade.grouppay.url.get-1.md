# Get combined-cashier URL

Original name: 组合收银台url获取  
API: `com.alibaba.trade:alibaba.trade.grouppay.url.get:1` · Category: Payment  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.trade.grouppay.url.get-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.trade.grouppay.url.get/{appKey}`  
Requires user authorization (access_token) · Requires signature

Get the combined-cashier URL.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `orderIds` | Long[] | yes | Order list | [123123413,1223234] |
| `payPlatformType` | String | no | PC or WIRELESS | PC |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `results` | [message:alibaba.ocean.openplatform.biz.trade.result.TradeCreateGroupPayUrlResult](#m-alibaba-ocean-openplatform-biz-trade-result-tradecreategrouppayurlresult) | yes |  |  |

<a id="m-alibaba-ocean-openplatform-biz-trade-result-tradecreategrouppayurlresult"></a>
#### alibaba.ocean.openplatform.biz.trade.result.TradeCreateGroupPayUrlResult

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `payUrl` | java.lang.String | yes | Return payurl | https://payment2.m.1688.com/page/cashier.html?orderId=154612245 |
| `success` | boolean | yes | Whether successful | true |
| `errorInfo` | java.lang.String | yes | Error message | null |
| `errorCode` | java.lang.String | yes | Error code | null |
