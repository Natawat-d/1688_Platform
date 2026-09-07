# 分销价格变更

Topic: `FENXIAO_PRICE_CHANGE` · Group: FENXIAO (分销)  
Doc page: https://open.1688.com/doc/topicDetail.htm?id=FENXIAO_PRICE_CHANGE

分销价格变更（关系视角），需要关注商品

## 消息字段

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `productId` | String | 是 | 1688商品id | 738629661216 |
| `memberId` | String | 是 | 1688会员id | b2b-2248564064 |
| `msgSendTime` | String | 是 | 消息发送时间 | 2024-05-09 00:00:00 |

## 消息示例

```json
{
  "productId": "738629661216",
  "memberId": "b2b-2248564064",
  "msgSendTime": "2024-05-09 00:00:00"
}
```
