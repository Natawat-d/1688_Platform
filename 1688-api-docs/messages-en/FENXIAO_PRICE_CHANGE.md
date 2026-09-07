# Distribution price changed

Original name: 分销价格变更  
Topic: `FENXIAO_PRICE_CHANGE` · Group: FENXIAO (分销)  
Doc page: https://open.1688.com/doc/topicDetail.htm?id=FENXIAO_PRICE_CHANGE

Distribution price change (relationship view); requires following the product

## Payload fields

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `productId` | String | yes | 1688 product ID | 738629661216 |
| `memberId` | String | yes | 1688 member ID | b2b-2248564064 |
| `msgSendTime` | String | yes | Message sending time | 2024-05-09 00:00:00 |

## Sample message

```json
{
  "productId": "738629661216",
  "memberId": "b2b-2248564064",
  "msgSendTime": "2024-05-09 00:00:00"
}
```
