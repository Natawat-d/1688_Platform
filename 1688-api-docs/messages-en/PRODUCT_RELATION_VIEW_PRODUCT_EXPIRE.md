# Product delisted (related-user view)

Original name: 1688产品下架（关系用户视角）  
Topic: `PRODUCT_RELATION_VIEW_PRODUCT_EXPIRE` · Group: PRODUCT (Product Messages)  
Doc page: https://open.1688.com/doc/topicDetail.htm?id=PRODUCT_RELATION_VIEW_PRODUCT_EXPIRE

1688 product delisted (taken offline); visible only to related users (including cross-border, distribution, and other relationships)

## Payload fields

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `productIds` | String | yes | Set of product IDs, at least one, separated by commas | 44179498967 |
| `memberId` | String | yes | 1688 member ID | xuxan |
| `status` | String | yes | Message type, which can specifically be RELATION_VIEW_PRODUCT_EXPIRE, RELATION_VIEW_PRODUCT_NEW_OR_MODIFY, RELATION_VIEW_PRODUCT_DELETE, or RELATION_VIEW_PRODUCT_REPOST | RELATION_VIEW_PRODUCT_EXPIRE |
| `msgSendTime` | String | yes | Message sending time | 2018-05-30 20:27:48 |

## Sample message

```json
{
  "productIds": "44179498967",
  "memberId": "xuxan",
  "status": "RELATION_VIEW_PRODUCT_EXPIRE",
  "msgSendTime": "2018-05-30 20:27:48"
}
```
