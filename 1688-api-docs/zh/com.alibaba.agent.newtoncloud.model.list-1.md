# 牛顿云-查询可用模型档位列表

API: `com.alibaba.agent:newtoncloud.model.list:1` · Category: 询盘  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.agent:newtoncloud.model.list-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.agent/newtoncloud.model.list/{appKey}`  
需要授权 (access_token) · 需要签名

查询牛顿云当前可用的模型档位列表。返回各档位的档位码（id）、展示名称（displayName）等信息，档位码可作为 newtoncloud.task.create 的 model参数取值。无业务入参，仅需 access_token 与签名。

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `success` | Boolean | 是 | 请求是否成功 | true |
| `error` | String | 是 | 失败时的错误描述 | 查询模型档位列表时发生内部错误，请稍后重试 |
| `models` | [message:model[]](#m-model[]) | 是 | 可用模型档位列表。返回内容由平台配置驱动，仅包含对外开放的档位 | [{"id":"flagship","displayName":"旗舰","isDefault":false}] |

<a id="m-model[]"></a>
#### model[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `id` | String | 是 | 档位码，即创建任务（newtoncloud.task.create）时 model 参数的取值 | flagship |
| `displayName` | String | 是 | 档位展示名称 | 旗舰 |
| `description` | String | 是 | 档位描述 | 能力最强 |
| `recommended` | Boolean | 是 | 是否推荐档位 | false |
| `iconUrl` | String | 是 | 档位图标地址 | xxx |
| `isDefault` | Boolean | 是 | 是否默认档位 | true |
| `sortOrder` | Integer | 是 | 排序权重（升序） | 1 |
