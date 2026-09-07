# 1688产品上架（关系用户视角）

Topic: `PRODUCT_RELATION_VIEW_PRODUCT_REPOST` · Group: PRODUCT (商品消息)  
Doc page: https://open.1688.com/doc/topicDetail.htm?id=PRODUCT_RELATION_VIEW_PRODUCT_REPOST

1688产品上架，仅关系用户（包含跨境、分销等关系）可见

## 消息字段

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `productIds` | String | 是 | 商品ID集合，至少有一个，用逗号分割 | 107680826 |
| `memberId` | String | 是 | 1688会员ID | shyxsscl |
| `status` | String | 是 | 消息类型，具体可为RELATION_VIEW_PRODUCT_EXPIRE、RELATION_VIEW_PRODUCT_NEW_OR_MODIFY、RELATION_VIEW_PRODUCT_DELETE、RELATION_VIEW_PRODUCT_REPOST | RELATION_VIEW_PRODUCT_REPOST |
| `msgSendTime` | String | 是 | 消息发送时间 | 2018-05-30 20:29:37 |

## 消息示例

```json
{
  "productIds": "107680826",
  "memberId": "shyxsscl",
  "status": "RELATION_VIEW_PRODUCT_REPOST",
  "msgSendTime": "2018-05-30 20:29:37"
}
```
