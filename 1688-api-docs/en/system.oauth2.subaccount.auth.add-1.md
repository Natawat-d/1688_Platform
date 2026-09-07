# Batch-add sub-account authorizations

Original name: 批量添加子账号授权  
API: `system.oauth2:subaccount.auth.add:1` · Category: Membership & Accounts  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=system.oauth2:subaccount.auth.add-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/system.oauth2/subaccount.auth.add/{appKey}`  
Requires user authorization (access_token) · Requires signature

Batch-add authorization for the sub-accounts under a main account. The main account must already be authorized.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `subUserIdentityList` | java.lang.String[] | yes | List of sub-account ids | abb:test1,abb:test2 |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.ocean.auth.dto.AuthResultModelMap](#m-alibaba-ocean-auth-dto-authresultmodelmap) | yes | Return result object | {} |

<a id="m-alibaba-ocean-auth-dto-authresultmodelmap"></a>
#### alibaba.ocean.auth.dto.AuthResultModelMap

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `errorCode` | java.lang.String | yes | Error code | param_error |
| `errorMessage` | java.lang.String | yes | Error message | invalid param |
| `returnValue` | java.util.Map | yes | Return result | map |
| `success` | boolean | yes | Whether successful | true |

## Samples

**Request example**

```
https://gw.open.1688.com/openapi/param2/1/system.oauth2/subaccount.auth.add/YOUR_APPKEY?access_token=xx&subUserIdentityList=a,b
```
