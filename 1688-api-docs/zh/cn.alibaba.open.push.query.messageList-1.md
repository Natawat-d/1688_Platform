# 查询式获取失败的消息列表

API: `cn.alibaba.open:push.query.messageList:1` · Category: 消息  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=cn.alibaba.open:push.query.messageList-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/cn.alibaba.open/push.query.messageList/{appKey}`  
无需授权 · 需要签名

查询式获取发送的消息列表，获取的消息不会自动确认，需要调用方手动调用确认api来确认消费状态。需注意，确认后，会影响分页返回的数据

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 否 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `createStartTime` | java.util.Date | 否 | 消息创建时间查找开始范围 | 20130417000000000+0800 |
| `createEndTime` | java.util.Date | 否 | 消息创建时间查找结束范围 | 20130417000000000+0800 |
| `page` | int | 否 | 当前数据页，默认为1 | 1 |
| `pageSize` | int | 否 | 每次分页取的数据量，范围20-200，默认20 | 20 |
| `type` | String | 否 | 消息类型 | ORDER_BUYER_MAKER |
| `userInfo` | String | 否 | 用户Id | b2b-4137495171f2513 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `pushMessagePage` | [message:PushMessagePage](#m-pushmessagepage) | 是 | 分页数据 | ["",""] |

<a id="m-pushmessagepage"></a>
#### PushMessagePage

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `datas` | [message:PushMessage[]](#m-pushmessage[]) | 是 | 分页的消息数据列表 |  |
| `totalCount` | int | 是 | 消息总数 |  |

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
