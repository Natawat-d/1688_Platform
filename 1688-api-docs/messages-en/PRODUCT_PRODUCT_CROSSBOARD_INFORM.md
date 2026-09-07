# One-click listing notification

Original name: 一键铺货消息  
Topic: `PRODUCT_PRODUCT_CROSSBOARD_INFORM` · Group: PRODUCT (Product Messages)  
Doc page: https://open.1688.com/doc/topicDetail.htm?id=PRODUCT_PRODUCT_CROSSBOARD_INFORM

User triggers one-click listing

## Payload fields

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `offerId` | String | yes | Product ID | 570046582711 |
| `userInfo` | String | yes | Member ID of the distributor's CBU | 2804951212 |
| `action` | String | yes | Action performed; the value is distribution | distribution |

## Sample message

```json
{
  "offerId": "570046582711",
  "userInfo": "2804951212",
  "action": "distribution"
}
```
