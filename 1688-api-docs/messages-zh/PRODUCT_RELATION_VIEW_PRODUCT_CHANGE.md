# 商品变更消息(关系用户视角、包含所有商品变更动作)

Topic: `PRODUCT_RELATION_VIEW_PRODUCT_CHANGE` · Group: PRODUCT (商品消息)  
Doc page: https://open.1688.com/doc/topicDetail.htm?id=PRODUCT_RELATION_VIEW_PRODUCT_CHANGE

商品变更消息(关系用户视角、包含所有商品变更动作)

## 消息字段

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `memberId` | String | 是 | 用户memberid | xxxx |
| `productIds` | String | 是 | 商品ID列表 | 897022666243,897022666242 |
| `action` | String | 是 | 操作类型 | 动作(new:用户新发,modify:用户修改,member_delete:用户删除,member_expired:用户下架,repost:用户重发,audit:小二审核,update:运营或运营工具修改,revised:商品订正,publish:上架,sku_new:sku发布,sku_modify:sku编辑,sku_delete:sku删除) |
| `msgSendTime` | String | 是 | 消息发送时间 | 2025-09-05 10:43:28 |

## 消息示例

```json
{
  "memberId": "xxxx",
  "productIds": "897022666243,897022666242",
  "action": "动作(new:用户新发,modify:用户修改,member_delete:用户删除,member_expired:用户下架,repost:用户重发,audit:小二审核,update:运营或运营工具修改,revised:商品订正,publish:上架,sku_new:sku发布,sku_modify:sku编辑,sku_delete:sku删除)",
  "msgSendTime": "2025-09-05 10:43:28"
}
```
