# Product added or modified (related-user view)

Original name: 1688产品新增或修改（关系用户视角）  
Topic: `PRODUCT_RELATION_VIEW_PRODUCT_NEW_OR_MODIFY` · Group: PRODUCT (Product Messages)  
Doc page: https://open.1688.com/doc/topicDetail.htm?id=PRODUCT_RELATION_VIEW_PRODUCT_NEW_OR_MODIFY

1688 product added or modified; visible only to related users (including cross-border, distribution, and other relationships)

## Payload fields

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `productIds` | String | yes | Set of product IDs, at least one, separated by commas | 547096780502 |
| `memberId` | String | yes | 1688 member ID | b2b-28308412336ca4e7 |
| `status` | String | yes | Message type, which can specifically be RELATION_VIEW_PRODUCT_EXPIRE, RELATION_VIEW_PRODUCT_NEW_OR_MODIFY, RELATION_VIEW_PRODUCT_DELETE, or RELATION_VIEW_PRODUCT_REPOST | RELATION_VIEW_PRODUCT_NEW_OR_MODIFY |
| `msgSendTime` | String | yes | Message sending time | 2018-05-30 20:28:42 |

## Sample message

```json
{
  "productIds": "547096780502",
  "memberId": "b2b-28308412336ca4e7",
  "status": "RELATION_VIEW_PRODUCT_NEW_OR_MODIFY",
  "msgSendTime": "2018-05-30 20:28:42"
}
```
