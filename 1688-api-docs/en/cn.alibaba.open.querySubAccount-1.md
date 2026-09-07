# Query sub-account info

Original name: 查询子账号信息  
API: `cn.alibaba.open:querySubAccount:1` · Category: Membership & Accounts  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=cn.alibaba.open:querySubAccount-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/cn.alibaba.open/querySubAccount/{appKey}`  
No user authorization · Requires signature

Query sub-account information.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | no | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `loginId` | String | yes | loginid | bonlientest:zhagnshan |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:ResultModel](#m-resultmodel) | yes | Return model | {"isSuccess":true,"data":{"userId":"3453","name":"张三"}} |

<a id="m-resultmodel"></a>
#### ResultModel

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `errorMsg` | java.lang.String | no | Error message | null |
| `resultCode` | java.lang.String | no | Result code | null |
| `isSuccess` | boolean | yes | Success flag | true |
| `data` | [message:CoopSubAccountModel](#m-coopsubaccountmodel) | yes | Data | {"departmentId":"352","userId":"3454234","name":"zhangsan"} |

<a id="m-coopsubaccountmodel"></a>
#### CoopSubAccountModel

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `userId` | java.lang.Long | yes | userId | 345345 |
| `loginId` | java.lang.String | yes | Login ID | bonlinetest:zhangsan |
| `employeeId` | java.lang.String | no | employeeid | 23453465 |
| `name` | java.lang.String | no | Name | 张三 |
| `sex` | java.lang.String | no | Gender | 男 |
| `personalPhone` | java.lang.String | no | Phone number | 13234567854 |
| `mail` | java.lang.String | no | Email | 234324343@qq.com |
| `departmentId` | java.lang.Long | no | Department ID | 36456 |
| `departmentName` | java.lang.String | no | Department name | 设计部 |

## Samples

****

```
入参
loginID：bonlinetest:zhangsan

返回
{"isSuccess":true,"data":{"userId":"3453","loginId":"bonlinetest:zhangsan","employeeId":"3453456","name":"张三"}}
```
