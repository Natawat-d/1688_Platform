# 跨境设为货源

Topic: `CROSSBOARD_CROSSBOARD_ADD_SUPPLY` · Group: CROSSBOARD (跨境消息)  
Doc page: https://open.1688.com/doc/topicDetail.htm?id=CROSSBOARD_CROSSBOARD_ADD_SUPPLY

当用户在 ISV 中访问产品开发工具同款页面并将 1688 商品设为货源时，系统将生成一条关于货源的记录，并通过消息通知到开放平台，由开放平台将此消息推送至下游 ISV；ISV 接受到此消息时，将产品和 1688 货源进行关联。

## 消息字段

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `offerId` | String | 是 | 1688商品ID | 537321696540 |
| `clientId` | String | 否 | ISV中的用户ID | b2b-26800439112ddd0 |
| `productId` | String | 否 | ISV中的商品ID | 560604911564 |
| `productUri` | String | 否 | ISV中的商品URL | http:123test.com |

## 消息示例

```json
{
  "offerId": "537321696540",
  "clientId": "b2b-26800439112ddd0",
  "productId": "560604911564",
  "productUri": "http:123test.com"
}
```
