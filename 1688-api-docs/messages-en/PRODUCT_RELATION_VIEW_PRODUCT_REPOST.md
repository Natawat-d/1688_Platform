# Product relisted (related-user view)

Original name: 1688产品上架（关系用户视角）  
Topic: `PRODUCT_RELATION_VIEW_PRODUCT_REPOST` · Group: PRODUCT (Product Messages)  
Doc page: https://open.1688.com/doc/topicDetail.htm?id=PRODUCT_RELATION_VIEW_PRODUCT_REPOST

1688 product listed (goes live); visible only to related users (including cross-border, distribution, and other relationships)

## Payload fields

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `productIds` | String | yes | Set of product IDs, at least one, separated by commas | 107680826 |
| `memberId` | String | yes | 1688 member ID | shyxsscl |
| `status` | String | yes | Message type, which can specifically be RELATION_VIEW_PRODUCT_EXPIRE, RELATION_VIEW_PRODUCT_NEW_OR_MODIFY, RELATION_VIEW_PRODUCT_DELETE, or RELATION_VIEW_PRODUCT_REPOST | RELATION_VIEW_PRODUCT_REPOST |
| `msgSendTime` | String | yes | Message sending time | 2018-05-30 20:29:37 |

## Sample message

```json
{
  "productIds": "107680826",
  "memberId": "shyxsscl",
  "status": "RELATION_VIEW_PRODUCT_REPOST",
  "msgSendTime": "2018-05-30 20:29:37"
}
```
