# Decrypt an OpenUID into a WangWang nickname (only for launching WangWang)

Original name: Openuid转换解密为旺旺昵称接口（仅可使用于用户唤起旺旺）  
API: `com.alibaba.account:wangwangnick.openuid.decrypt:1` · Category: Tools  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.account:wangwangnick.openuid.decrypt-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.account/wangwangnick.openuid.decrypt/{appKey}`  
Requires user authorization (access_token) · Requires signature

Decrypts an OpenUID into a WangWang nickname. This interface is risk-controlled: it may only be used when a user needs to launch WangWang. Automated batch operations are not allowed, and it must not be used as a decryption interface to show plaintext to users. WangWang supports recall after encryption; do not use it for any other scenario.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | yes | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `openUid` | java.lang.String | yes | openUid to be decrypted |   |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `wangwangNick` | java.lang.String | yes | WangWang nickname used to launch WangWang |   |
