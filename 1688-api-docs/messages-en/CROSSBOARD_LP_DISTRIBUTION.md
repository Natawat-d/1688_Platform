# Cross-border purchasing assistant one-click listing

Original name: 跨境采购助手一键铺货消息  
Topic: `CROSSBOARD_LP_DISTRIBUTION` · Group: CROSSBOARD (跨境消息)  
Doc page: https://open.1688.com/doc/topicDetail.htm?id=CROSSBOARD_LP_DISTRIBUTION

The cross-border procurement assistant page lists 1688 products to the connected ERP system with one click

## Payload fields

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `customerId` | String | yes | Customer ID accessing the product-selection page | testIsv1 |
| `userKey` | String | yes | Encrypted user ID within the ISV system, remains unchanged | qwerty |
| `appkey` | String | yes | Open Platform appkey | 123 |
| `userToken` | String | yes | User token information within the ISV system (the ISV can parse its own user login state information from this) | asdfggjkl |
| `offerList` | Object[] | yes | List of listed products | "offerList":[{  "offerId":"1244567" ，   "pic":"商品图片"}] |
| &nbsp;&nbsp;↳ `offerId` | String | no | Product ID | 1244567 |
| &nbsp;&nbsp;↳ `pic` | String | no | Product main image URL | https://cbu01.alicdn.com/xxx.jpg |
| `sign` | String | yes | Signature information | ABCDEFG |
| `timestamp` | String | yes | Timestamp of the current listing operation (milliseconds) | 123454654645 |

## Sample message

```json
{
  "customerId": "testIsv1",
  "userKey": "qwerty",
  "appkey": "123",
  "userToken": "asdfggjkl",
  "offerList": [
    {
      "offerId": "1244567",
      "pic": "https://cbu01.alicdn.com/xxx.jpg"
    }
  ],
  "sign": "ABCDEFG",
  "timestamp": "123454654645"
}
```
