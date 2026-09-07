# 精选货源商品下架消息

Topic: `PRODUCT_PFT_OFFER_QUIT` · Group: PRODUCT (商品消息)  
Doc page: https://open.1688.com/doc/topicDetail.htm?id=PRODUCT_PFT_OFFER_QUIT

精选货源商品下架消息

## 消息字段

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `offerId` | Number | 是 | 产品id | 662112483578 |
| `type` | String | 否 | 团类型 | single_direct |
| `openUid` | String | 否 | offer所属商家openUid | 1117283046 |

## 消息示例

```json
{
  "offerId": 662112483578,
  "type": "single_direct",
  "openUid": "1117283046"
}
```
