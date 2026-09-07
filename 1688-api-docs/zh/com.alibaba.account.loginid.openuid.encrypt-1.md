# 用户loginId加密转换为Openuid接口

API: `com.alibaba.account:loginid.openuid.encrypt:1` · Category: 工具  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.account:loginid.openuid.encrypt-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.account/loginid.openuid.encrypt/{appKey}`  
需要授权 (access_token) · 需要签名

用户loginId加密转换为Openuid接口。该接口进行风控，不允许批量操作，仅允许商家手动触发，比如：搜索用户订单、设置规则等功能

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 是 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `loginId` | String | 是 | 用户登陆名 |   |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `openUid` | String | 是 | openUid |   |
