# Product inventory changed (related-user view)

Original name: 1688商品库存变更消息（关系用户视角）  
Topic: `PRODUCT_PRODUCT_INVENTORY_CHANGE` · Group: PRODUCT (Product Messages)  
Doc page: https://open.1688.com/doc/topicDetail.htm?id=PRODUCT_PRODUCT_INVENTORY_CHANGE

Product stock change message. The following situations trigger a product stock change: the merchant edits the product, an order is placed, an order is closed and stock is replenished, platform staff modify the stock, or the stock is modified via the API. To receive stock change messages for a product, you must first call the follow-product API to follow that product before you can receive its stock change messages.

## Payload fields

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `OfferInventoryChangeList` | Object[] | no | List of stock changes |  |
| &nbsp;&nbsp;↳ `offerId` | Number | yes | 1688 product ID | 1234567890 |
| &nbsp;&nbsp;↳ `offerOnSale` | Number | yes | Number of offers currently available for sale online | 100 |
| &nbsp;&nbsp;↳ `skuId` | Number | no | Product skuId | 1234567890 |
| &nbsp;&nbsp;↳ `skuOnSale` | Number | no | Number of SKUs currently available for sale online | 20 |
| &nbsp;&nbsp;↳ `quantity` | Number | yes | Overall stock change quantity for the offer | -10 |
| &nbsp;&nbsp;↳ `bizTime` | Date | yes | Time of stock change | 1564984329147 |

## Sample message

```json
{
  "OfferInventoryChangeList": [
    {
      "offerId": 1234567890,
      "offerOnSale": 100,
      "skuId": 1234567890,
      "skuOnSale": 20,
      "quantity": -10,
      "bizTime": "1564984329147"
    }
  ]
}
```
