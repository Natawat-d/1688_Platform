# Openuid转换解密为旺旺昵称接口（仅可使用于用户唤起旺旺）

API: `com.alibaba.account:wangwangnick.openuid.decrypt:1` · Category: 工具  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.account:wangwangnick.openuid.decrypt-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.account/wangwangnick.openuid.decrypt/{appKey}`  
需要授权 (access_token) · 需要签名

Openuid转换解密为旺旺昵称接口。该接口进行风控，仅可使用于用户需要唤起旺旺，不允许自动化批量操作，不允许作为解密接口将明文展示给用户。旺旺支持加密后回收，请勿使用于其他场景。

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 是 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `openUid` | java.lang.String | 是 | 待解密的openUid |   |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `wangwangNick` | java.lang.String | 是 | 用以唤起旺旺的旺旺昵称 |   |
