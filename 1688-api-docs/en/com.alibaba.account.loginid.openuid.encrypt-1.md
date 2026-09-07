# Encrypt a user loginId into an OpenUID

Original name: 用户loginId加密转换为Openuid接口  
API: `com.alibaba.account:loginid.openuid.encrypt:1` · Category: Tools  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.account:loginid.openuid.encrypt-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.account/loginid.openuid.encrypt/{appKey}`  
Requires user authorization (access_token) · Requires signature

Encrypts a user's loginId into an OpenUID. This interface is risk-controlled: batch operations are not allowed, and it may only be triggered manually by the merchant, for example when searching a user's orders or configuring rules.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | yes | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `loginId` | String | yes | User login name |   |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `openUid` | String | yes | openUid |   |
