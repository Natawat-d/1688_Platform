# 失败消息批量确认

API: `cn.alibaba.open:push.message.confirm:1` · Category: 消息  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=cn.alibaba.open:push.message.confirm-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/cn.alibaba.open/push.message.confirm/{appKey}`  
无需授权 · 需要签名

手动调用确认api，确认消息已经被消费成功。仅当使用查询式获取失败消息的api时，才需要使用

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 否 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `msgIdList` | java.util.List | 否 | 待确认的消息id列表 | [123,456] |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `isSuccess` | boolean | 是 | 操作是否成功 | true |
