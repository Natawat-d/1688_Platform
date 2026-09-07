# Get basic info of a non-authorized user (cross-border)

Original name: 跨境场景获取非授权用户的基本信息  
API: `com.alibaba.account:alibaba.account.agent.crossBasic:1` · Category: Membership & Accounts  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.account:alibaba.account.agent.crossBasic-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.account/alibaba.account.agent.crossBasic/{appKey}`  
Requires user authorization (access_token) · Requires signature

View another user's basic information. For use in cross-border scenarios.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `loginId` | String | no | The user's loginId; the input parameters cannot both be empty | alitestforisv01 |
| `domain` | String | no | WangPu (storefront) domain name; input parameters cannot all be empty at the same time | trgm66666.1688.com |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.account.simpleAccountInfo](#m-alibaba-account-simpleaccountinfo) | yes | Member information | {} |
| `errorCode` | String | yes | Error code | 001 |
| `errorMessage` | String | yes | Error message |   |
| `success` | Boolean | yes | Whether successful | true |

<a id="m-alibaba-account-simpleaccountinfo"></a>
#### alibaba.account.simpleAccountInfo

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `loginId` | java.lang.String | yes | Login name | alitestforis01 |
| `categoryName` | java.lang.String | yes | Category name of the main industry | 连衣裙 |
| `companyName` | java.lang.String | yes | Company name | 测试公司1 |
| `shopUrl` | java.lang.String | yes | WangPu (shop) homepage URL | https://alitest.1688.com	 |
| `supplierName` | java.lang.String | yes | Supplier name | alitestforis01 |
| `kuaJingBao` | Boolean | yes | Whether it is Kuajingbao (Cross-Border Pay) | true |

## Samples

**Output result example**

```
{"result":{"loginId":"alitestforisv01","categoryName":"办公、文教","companyName":"AOP对外测试账号01","shopUrl":"https://trgm66666.1688.com"}}
```
