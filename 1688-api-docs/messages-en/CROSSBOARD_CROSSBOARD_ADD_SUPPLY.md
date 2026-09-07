# Product set as cross-border supply source

Original name: 跨境设为货源  
Topic: `CROSSBOARD_CROSSBOARD_ADD_SUPPLY` · Group: CROSSBOARD (跨境消息)  
Doc page: https://open.1688.com/doc/topicDetail.htm?id=CROSSBOARD_CROSSBOARD_ADD_SUPPLY

When a user accesses the matching-product development tool page within the ISV and sets a 1688 product as the supply source, the system generates a supply-source record and notifies the Open Platform via a message; the Open Platform then pushes this message to the downstream ISV. When the ISV receives this message, it associates the product with the 1688 supply source.

## Payload fields

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `offerId` | String | yes | 1688 product ID | 537321696540 |
| `clientId` | String | no | User ID in the ISV | b2b-26800439112ddd0 |
| `productId` | String | no | Product ID in the ISV | 560604911564 |
| `productUri` | String | no | Product URL in the ISV | http:123test.com |

## Sample message

```json
{
  "offerId": "537321696540",
  "clientId": "b2b-26800439112ddd0",
  "productId": "560604911564",
  "productUri": "http:123test.com"
}
```
