# Batch-query sub-account authorizations

Original name: 批量查询子账号授权  
API: `system.oauth2:subaccount.auth.list:1` · Category: Membership & Accounts  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=system.oauth2:subaccount.auth.list-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/system.oauth2/subaccount.auth.list/{appKey}`  
Requires user authorization (access_token) · Requires signature

Batch-query the authorization status of the sub-accounts under a main account.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `subUserIdentityList` | java.lang.String[] | yes | List of sub-account ids | "a:b","a:c" |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.ocean.auth.dto.AuthResultModelRelation](#m-alibaba-ocean-auth-dto-authresultmodelrelation) | yes | Query result | xx |

<a id="m-alibaba-ocean-auth-dto-authresultmodelrelation"></a>
#### alibaba.ocean.auth.dto.AuthResultModelRelation

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `errorCode` | java.lang.String | yes | Error code | param_error |
| `errorMessage` | java.lang.String | yes | Error description | lack of param |
| `returnValue` | [message:alibaba.ocean.auth.dto.AuthRelationDTO[]](#m-alibaba-ocean-auth-dto-authrelationdto[]) | yes | Return result | 结构体 |
| `success` | boolean | yes | Whether successful | true |

<a id="m-alibaba-ocean-auth-dto-authrelationdto[]"></a>
#### alibaba.ocean.auth.dto.AuthRelationDTO[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `accessToken` | java.lang.String | yes | Authorization credential | xx |
| `adminOwnerId` | java.lang.String | yes | Main account loginId | xx |
| `adminUserId` | java.lang.Long | yes | Main account userId | xx |
| `clientId` | java.lang.String | yes | appKey | xx |
| `clientName` | java.lang.String | yes | appName | xx |
| `gmtExpired` | java.util.Date | yes | Authorization expiration time | xx |
| `memberId` | java.lang.String | yes | The memberId of the authorized user | xx |
| `ownerId` | java.lang.String | yes | The loginId of the authorized user | xx |
| `resourceScopes` | java.lang.String | yes | Resource domain | xx |
| `site` | java.lang.String | yes | Authorized site | xx |
| `status` | java.lang.String | yes | Authorization status | xx |
| `subAuth` | java.lang.Boolean | yes | Whether it is sub-account authorization | xx |
| `subOwnerId` | java.lang.String | yes | Sub-account loginId | xx |
| `subUserId` | java.lang.Long | yes | Sub-account userId | xx |
| `userId` | java.lang.Long | yes | Authorized user's userId | xx |

## Samples

**Request example**

```
https://gw.open.1688.com/openapi/param2/1/system.oauth2/subaccount.auth.list/YOUR_APPKEY?access_token=xx&subUserIdentityList=a,b
```
