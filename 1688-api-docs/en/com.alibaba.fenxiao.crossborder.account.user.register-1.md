# 1688 member registration

Original name: 1688会员注册  
API: `com.alibaba.fenxiao.crossborder:account.user.register:1` · Category: Membership & Accounts  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:account.user.register-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/account.user.register/{appKey}`  
Requires user authorization (access_token) · Requires signature

Register a 1688 member.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `countryAccount` | [message:account.user.register.param.CountryAccount](#m-account-user-register-param-countryaccount) | yes | Account registration input parameters | {     "country": "japan",     "site": "sniff",     "outLoginId": "18899993333",     "outMemberId": "c5b9e8a658554771852063f3a44d4e3d",     "mobile": "18899993333",     "mobileArea": "JP",     "ip": "11.11.11.11",     "email": "123@163.com" } |

<a id="m-account-user-register-param-countryaccount"></a>
#### account.user.register.param.CountryAccount

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `country` | java.lang.String | yes | Country; for supported parameters see the country enum. | japan |
| `site` | java.lang.String | yes | Site; please pass in the company name or e-commerce product name. Must contain no spaces and be lowercase. No more than 16 characters. | sniff |
| `outLoginId` | java.lang.String | yes | The user's login name within the organization, no more than 60 characters. | 1234test |
| `outMemberId` | java.lang.String | yes | The user's unique identifier within the organization, no more than 60 characters. | 1232fdsf |
| `email` | java.lang.String | yes | Email; format will be validated, no more than 60 characters. | 123@163.com |
| `mobile` | java.lang.String | yes | Mobile number; format will be validated, no more than 30 characters. | 1234567890 |
| `mobileArea` | java.lang.String | yes | The region the mobile number belongs to; see the mobileArea enum for supported parameters. | JP |
| `ip` | java.lang.String | yes | IP address; format will be validated. | 11.11.11.11 |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:account.user.register.result.ResultModel](#m-account-user-register-result-resultmodel) | yes |  |  |

<a id="m-account-user-register-result-resultmodel"></a>
#### account.user.register.result.ResultModel

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `success` | boolean | yes | Whether the call succeeded, true for success, false for failure. | true |
| `code` | java.lang.String | yes | Error code, e.g. S0000 indicates success. | S0000 |
| `message` | java.lang.String | yes | Error message, e.g. success. | 成功 |
| `result` | java.lang.Boolean | yes | Return result. | true |

## Samples

**Input parameter example**

```
{
    "country": "japan",
    "outLoginId": "1234test",
    "outUserId": "3333333333",
    "site": "test",
    "mobileArea": "JP",
    "ip": "11.11.11.11",
    "email": "122323@234.com",
    "mobile": "1234"
}
```

**Output parameter example**

```
{
    "success": true,
    "code": "S0000",
    "message": "success",
    "result": true
}
```
