# Save the business line an account belongs to

Original name: 保存账号所属业务线  
API: `com.alibaba.fenxiao.crossborder:account.business.save:1` · Category: Data Write-back  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:account.business.save-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/account.business.save/{appKey}`  
Requires user authorization (access_token) · Requires signature

Save the business line that an account belongs to.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `accountPerformance` | [message:account.business.save.AccountPerformance](#m-account-business-save-accountperformance) | yes | Input parameter | 如下 |

<a id="m-account-business-save-accountperformance"></a>
#### account.business.save.AccountPerformance

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `account` | java.lang.String | yes | Account name | b测试账号009 或 b测试账号009:lctest |
| `business` | java.lang.String | yes | Business line it belongs to | 东南亚sea，南亚sa，日韩jk，港澳台hmt，中东me，北美na，拉美la，西欧we，泛俄ru，非洲af |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:account.business.save.AccountPerformanceResult](#m-account-business-save-accountperformanceresult) | yes | Result | 如下 |

<a id="m-account-business-save-accountperformanceresult"></a>
#### account.business.save.AccountPerformanceResult

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `saveResult` | Boolean | yes | Save result, true for success, false for failure | true |
