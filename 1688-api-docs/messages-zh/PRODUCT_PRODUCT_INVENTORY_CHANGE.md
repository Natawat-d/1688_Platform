# 1688商品库存变更消息（关系用户视角）

Topic: `PRODUCT_PRODUCT_INVENTORY_CHANGE` · Group: PRODUCT (商品消息)  
Doc page: https://open.1688.com/doc/topicDetail.htm?id=PRODUCT_PRODUCT_INVENTORY_CHANGE

商品库存变更消息，如下情况会导致触发商品库存变更：商家编辑商品、交易下单、订单关闭，回补库存、小二修改库存、API修改库存。接收某商品库存变更的前提，需先调用关注商品API关注某个商品，然后才能接收该商品库存变更的消息。

## 消息字段

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `OfferInventoryChangeList` | Object[] | 否 | 库存变更列表 |  |
| &nbsp;&nbsp;↳ `offerId` | Number | 是 | 1688商品id | 1234567890 |
| &nbsp;&nbsp;↳ `offerOnSale` | Number | 是 | 在线可售offer数量 | 100 |
| &nbsp;&nbsp;↳ `skuId` | Number | 否 | 商品skuId | 1234567890 |
| &nbsp;&nbsp;↳ `skuOnSale` | Number | 否 | 在线可售sku数量 | 20 |
| &nbsp;&nbsp;↳ `quantity` | Number | 是 | 该offer整体库存变化数 | -10 |
| &nbsp;&nbsp;↳ `bizTime` | Date | 是 | 库存变更时间 | 1564984329147 |

## 消息示例

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
