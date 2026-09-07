# Fetch failed messages (query style)

Original name: 查询式获取失败的消息列表  
API: `cn.alibaba.open:push.query.messageList:1` · Category: Messaging  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=cn.alibaba.open:push.query.messageList-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/cn.alibaba.open/push.query.messageList/{appKey}`  
No user authorization · Requires signature

Query-style retrieval of sent messages. Retrieved messages are not confirmed automatically; the caller must call the confirmation API to confirm the consumption status. Note that confirming affects the data returned by pagination.

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
| `page` | int | no | Current data page, default is 1 | 1 |
| `pageSize` | int | no | Number of records retrieved per page, range 20-200, default 20 | 20 |
| `type` | String | no | Message type | ORDER_BUYER_MAKER |
| `userInfo` | String | no | User Id | b2b-4137495171f2513 |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `pushMessagePage` | [message:PushMessagePage](#m-pushmessagepage) | yes | Paginated data | ["",""] |

<a id="m-pushmessagepage"></a>
#### PushMessagePage

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `datas` | [message:PushMessage[]](#m-pushmessage[]) | yes | Paginated list of message data |  |
| `totalCount` | int | yes | Total number of messages |  |

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
    "pushMessagePage": {
        "datas": [
            {
                "gmtBorn": 1399452484000,
                "topicGroup": "CAIGOU1",
                "data": {
                    "quotationId": 123,
                    "buyOfferId": 123,
                    "supplierMemberId": "memberId"
                },
                "msgId": 68891027,
                "type": "CAIGOU_MSG_BUYER_MARK_QUOTATION",
                "userInfo": "memberId",
                "appKey": "123456",
                "topicName": "MSG_BUYER_MARK_QUOTATION"
            }                
        ],
        "totalCount": 1
    }
}
```
