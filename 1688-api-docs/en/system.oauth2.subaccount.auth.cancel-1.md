# Batch-cancel sub-account authorizations

Original name: 批量取消子账号授权  
API: `system.oauth2:subaccount.auth.cancel:1` · Category: Membership & Accounts  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=system.oauth2:subaccount.auth.cancel-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/system.oauth2/subaccount.auth.cancel/{appKey}`  
No user authorization · Requires signature

Cancel the authorization of sub-accounts in batch.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | no | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `subUserIdentityList` | java.lang.String[] | yes | List of sub-account ids | a:b,a:c |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.ocean.auth.dto.AuthResultModelBoolean](#m-alibaba-ocean-auth-dto-authresultmodelboolean) | yes | Return result | xx |

<a id="m-alibaba-ocean-auth-dto-authresultmodelboolean"></a>
#### alibaba.ocean.auth.dto.AuthResultModelBoolean

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `errorCode` | java.lang.String | yes | Error code | param_error |
| `errorMessage` | java.lang.String | yes | Error description | invalid param |
| `returnValue` | java.lang.Boolean | yes | Return result | true |
| `success` | boolean | yes | Whether successful | true |

## Samples

**Request example**

```
https://gw.open.1688.com/openapi/param2/1/system.oauth2/subaccount.auth.cancel/YOUR_APPKEY?access_token=xx&subUserIdentityList=a,b
```
