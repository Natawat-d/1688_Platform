# 获取唤起旺旺聊天的链接

API: `com.alibaba.account:account.wangwangUrl.get:1` · Category: 工具  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.account:account.wangwangUrl.get-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.account/account.wangwangUrl.get/{appKey}`  
需要授权 (access_token) · 需要签名

获取唤起旺旺聊天的链接

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `toOpenUid` | String | 是 | 聊天对象openUid |   |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:alibaba.ocean.openplatform.common.ResultModelGTlfvn7n](#m-alibaba-ocean-openplatform-common-resultmodelgtlfvn7n) | 是 |  |  |

<a id="m-alibaba-ocean-openplatform-common-resultmodelgtlfvn7n"></a>
#### alibaba.ocean.openplatform.common.ResultModelGTlfvn7n

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `code` | java.lang.String | 是 |  |   |
| `message` | java.lang.String | 是 |  |   |
| `result` | java.lang.String | 是 | 旺旺聊天链接 |   |
| `retCodes` | java.lang.String[] | 是 |  |   |
| `subCode` | java.lang.String | 是 |  |   |
| `subMessage` | java.lang.String | 是 |  |   |
| `success` | boolean | 是 |  |   |
