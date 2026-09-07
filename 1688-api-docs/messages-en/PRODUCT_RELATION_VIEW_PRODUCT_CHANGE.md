# Product changed (related-user view; covers every product change action)

Original name: 商品变更消息(关系用户视角、包含所有商品变更动作)  
Topic: `PRODUCT_RELATION_VIEW_PRODUCT_CHANGE` · Group: PRODUCT (Product Messages)  
Doc page: https://open.1688.com/doc/topicDetail.htm?id=PRODUCT_RELATION_VIEW_PRODUCT_CHANGE

Product change message (from the related-user view, includes all product change actions)

## Payload fields

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `memberId` | String | yes | User memberid | xxxx |
| `productIds` | String | yes | List of product IDs | 897022666243,897022666242 |
| `action` | String | yes | Operation type | 动作(new:用户新发,modify:用户修改,member_delete:用户删除,member_expired:用户下架,repost:用户重发,audit:小二审核,update:运营或运营工具修改,revised:商品订正,publish:上架,sku_new:sku发布,sku_modify:sku编辑,sku_delete:sku删除) |
| `msgSendTime` | String | yes | Message sending time | 2025-09-05 10:43:28 |

## Sample message

```json
{
  "memberId": "xxxx",
  "productIds": "897022666243,897022666242",
  "action": "动作(new:用户新发,modify:用户修改,member_delete:用户删除,member_expired:用户下架,repost:用户重发,audit:小二审核,update:运营或运营工具修改,revised:商品订正,publish:上架,sku_new:sku发布,sku_modify:sku编辑,sku_delete:sku删除)",
  "msgSendTime": "2025-09-05 10:43:28"
}
```
