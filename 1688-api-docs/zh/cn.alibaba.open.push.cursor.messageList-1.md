# 游标式获取失败的消息列表

API: `cn.alibaba.open:push.cursor.messageList:1` · Category: 消息  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=cn.alibaba.open:push.cursor.messageList-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/cn.alibaba.open/push.cursor.messageList/{appKey}`  
无需授权 · 需要签名

游标式获取失败的消息列表，获取的消息会自动消费成功的确认。所以下次以相同条件调用获取的是剩下的数据，直至返回数据为空

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 否 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `createStartTime` | java.util.Date | 否 | 消息创建时间查找范围开始 | 20130417000000000+0800 |
| `createEndTime` | java.util.Date | 否 | 消息创建时间查找范围结束 | 20130417000000000+0800 |
| `quantity` | int | 否 | 每次取的数据量，范围20-200，默认20 | 20 |
| `type` | String | 否 | 消息类型 | ORDER_BUYER_MAKER |
| `userInfo` | String | 否 | 用户Id | b2b-4137495171f2513 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `pushMessageList` | [message:PushMessage[]](#m-pushmessage[]) | 是 | 推送消息列表 | ["",""] |

<a id="m-pushmessage[]"></a>
#### PushMessage[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `msgId` | long | 是 | 消息唯一id |  |
| `type` | String | 是 | 消息类型 |  |
| `userInfo` | String | 是 | 消息关联的用户memberId |  |
| `data` | java.util.Map | 是 | 消息内容 |  |
| `gmtBorn` | long | 是 | 消息创建的时间戳，单位毫秒 |  |

## 示例

**json返回结果**

```
{
    "pushMessageList": [
        {
            "gmtBorn": 1399182274000,
            "topicGroup": "CAIGOU",
            "data": {
                "buyOfferId": 123,
                "subUserId": 123
            },
            "msgId": 123456,
            "type": "CAIGOU_MSG_BUYER_PUBLISH_BUYOFFER",
            "userInfo": "memberId",
            "appKey": "1688",
            "topicName": "MSG_BUYER_PUBLISH_BUYOFFER"
        }        
    ]
}
```
