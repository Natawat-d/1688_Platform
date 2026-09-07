# Fetch failed messages (cursor style)

Original name: 游标式获取失败的消息列表  
API: `cn.alibaba.open:push.cursor.messageList:1` · Category: Messaging  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=cn.alibaba.open:push.cursor.messageList-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/cn.alibaba.open/push.cursor.messageList/{appKey}`  
No user authorization · Requires signature

Cursor-style retrieval of failed messages. Retrieved messages are automatically confirmed as consumed, so the next call with the same conditions returns the remaining data, until the result is empty.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | no | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `createStartTime` | java.util.Date | no | Start of the message creation time search range | 20130417000000000+0800 |
| `createEndTime` | java.util.Date | no | End of the message creation time search range | 20130417000000000+0800 |
| `quantity` | int | no | Number of records fetched per request, range 20-200, default 20 | 20 |
| `type` | String | no | Message type | ORDER_BUYER_MAKER |
| `userInfo` | String | no | User Id | b2b-4137495171f2513 |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `pushMessageList` | [message:PushMessage[]](#m-pushmessage[]) | yes | List of push messages | ["",""] |

<a id="m-pushmessage[]"></a>
#### PushMessage[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `msgId` | long | yes | Unique message id |  |
| `type` | String | yes | Message type |  |
| `userInfo` | String | yes | The memberId of the user associated with the message |  |
| `data` | java.util.Map | yes | Message content |  |
| `gmtBorn` | long | yes | Timestamp when the message was created, in milliseconds |  |

## Samples

**json return result**

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
