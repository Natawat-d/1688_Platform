# Product audit result (related-user view)

Original name: 1688产品审核（关系用户视角）  
Topic: `PRODUCT_RELATION_VIEW_PRODUCT_AUDIT` · Group: PRODUCT (Product Messages)  
Doc page: https://open.1688.com/doc/topicDetail.htm?id=PRODUCT_RELATION_VIEW_PRODUCT_AUDIT

1688 product review (audit), from the perspective of related users (including cross-border, distribution, and other relationships)

## Payload fields

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `productIds` | String | yes | Set of product IDs, at least one, separated by commas | 570872343603 |
| `memberId` | String | yes | 1688 member ID | b2b-342071512394d025 |
| `status` | String | yes | Message type, which can specifically be RELATION_VIEW_PRODUCT_EXPIRE, RELATION_VIEW_PRODUCT_NEW_OR_MODIFY, RELATION_VIEW_PRODUCT_DELETE, or RELATION_VIEW_PRODUCT_REPOST | RELATION_VIEW_PRODUCT_AUDIT |
| `msgSendTime` | String | yes | Message sending time | 2018-05-30 20:26:41 |

## Sample message

```json
{
  "productIds": "570872343603",
  "memberId": "b2b-342071512394d025",
  "status": "RELATION_VIEW_PRODUCT_AUDIT",
  "msgSendTime": "2018-05-30 20:26:41"
}
```
