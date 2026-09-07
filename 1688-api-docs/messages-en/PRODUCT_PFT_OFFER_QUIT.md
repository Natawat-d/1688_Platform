# Curated-supply product delisted

Original name: 精选货源商品下架消息  
Topic: `PRODUCT_PFT_OFFER_QUIT` · Group: PRODUCT (Product Messages)  
Doc page: https://open.1688.com/doc/topicDetail.htm?id=PRODUCT_PFT_OFFER_QUIT

Curated supply product delisted message

## Payload fields

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `offerId` | Number | yes | Product ID | 662112483578 |
| `type` | String | no | Group type | single_direct |
| `openUid` | String | no | openUid of the merchant that owns the offer | 1117283046 |

## Sample message

```json
{
  "offerId": 662112483578,
  "type": "single_direct",
  "openUid": "1117283046"
}
```
