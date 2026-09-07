# 精选货源商品价格变动消息

Topic: `PRODUCT_PFT_OFFER_PRICE_MODIFY` · Group: PRODUCT (商品消息)  
Doc page: https://open.1688.com/doc/topicDetail.htm?id=PRODUCT_PFT_OFFER_PRICE_MODIFY

精选货源商品价格变动消息

## 消息字段

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `offerId` | Number | 是 | 商品ID | 679304426617 |
| `minPrice` | Number | 是 | 价格变动 | 2100 |
| `activeTime` | Number | 是 | 价格变动生效时间 | 1662426000000 |
| `skuUpdateInfos` | Object | 否 | sku价格变动情况 |  |
| &nbsp;&nbsp;↳ `skuId` | Number | 是 | skuId | 4851072600824 |
| &nbsp;&nbsp;↳ `historyRetailPrice` | Number | 否 | 变更价格，分 | 1550 |

## 消息示例

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
