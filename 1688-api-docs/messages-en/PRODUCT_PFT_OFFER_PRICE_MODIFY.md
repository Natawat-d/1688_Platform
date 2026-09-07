# Curated-supply product price changed

Original name: 精选货源商品价格变动消息  
Topic: `PRODUCT_PFT_OFFER_PRICE_MODIFY` · Group: PRODUCT (Product Messages)  
Doc page: https://open.1688.com/doc/topicDetail.htm?id=PRODUCT_PFT_OFFER_PRICE_MODIFY

Curated supply product price change message

## Payload fields

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `offerId` | Number | yes | Product ID | 679304426617 |
| `minPrice` | Number | yes | Price change | 2100 |
| `activeTime` | Number | yes | Effective time of the price change | 1662426000000 |
| `skuUpdateInfos` | Object | no | SKU price change details |  |
| &nbsp;&nbsp;↳ `skuId` | Number | yes | skuId | 4851072600824 |
| &nbsp;&nbsp;↳ `historyRetailPrice` | Number | no | Changed price, in cents (fen) | 1550 |

## Sample message

```json
{
  "offerId": 679304426617,
  "minPrice": 2100,
  "activeTime": 1662426000000,
  "skuUpdateInfos": {
    "skuId": 4851072600824,
    "historyRetailPrice": 1550
  }
}
```
