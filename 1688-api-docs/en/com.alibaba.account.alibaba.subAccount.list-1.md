# Get sub-account list

Original name: 获取子账号列表  
API: `com.alibaba.account:alibaba.subAccount.list:1` · Category: Orders  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.account:alibaba.subAccount.list-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.account/alibaba.subAccount.list/{appKey}`  
Requires user authorization (access_token) · Requires signature

Get the user's main-account and sub-account information. If the API is authorized as a sub-account, only the main account that the sub-account belongs to is returned. If authorized as a main account, the list of all sub-accounts is returned.

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
| `mainUserId` | String | yes | Main account Userid | 12346 |
| `mainLoginId` | String | yes | Main account loginId | alitestforisv01 |
| `mainMemberId` | String | yes | Main account MemberId | b2b-123565 |
| `subAccountList` | [message:alibaba.account.simpleAccountInfo[]](#m-alibaba-account-simpleaccountinfo[]) | yes | Sub-account list | [] |
| `errorMsg` | String | yes | Error description | 获取帐号信息失败 |
| `errorCode` | String | yes | Error code | 500_001 |

<a id="m-alibaba-account-simpleaccountinfo[]"></a>
#### alibaba.account.simpleAccountInfo[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `loginId` | java.lang.String | yes | Sub-account login name | alitestforisv01:01 |
| `userId` | java.lang.Long | yes | Sub-account user ID | 12124134 |
| `memberId` | java.lang.String | yes | Sub-account memberId | b2b-12315435 |
| `status` | String | yes | Sub-account status | enabled |

## Samples

**Output parameter example**

```
{
  "mainLoginId": "b测试账号003",
  "mainMemberId": "b2b-22**4086",
  "mainUserId": "22**86",
  "subAccountList": [
    {
      "loginId": "b测试账号003:yms",
      "userId": 400137**76,
      "memberId": "b2b-22**4086"
    },
    {
      "loginId": "b测试账号003:yansen",
      "userId": 4000**0833,
      "memberId": "b2b-22**86"
    }
  ]
}

```
