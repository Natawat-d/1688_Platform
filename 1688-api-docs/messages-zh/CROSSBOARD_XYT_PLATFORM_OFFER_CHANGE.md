# 寻源通工作台商品变更信息

Topic: `CROSSBOARD_XYT_PLATFORM_OFFER_CHANGE` · Group: CROSSBOARD (跨境消息)  
Doc page: https://open.1688.com/doc/topicDetail.htm?id=CROSSBOARD_XYT_PLATFORM_OFFER_CHANGE

寻源通工作台商品变更信息

## 消息字段

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `productIds` | String | 是 | 商品ID集合，至少有一个，用逗号分割 | 30798998397239 |
| `status` | String | 是 | 消息类型，具体可为add（新增）、delete（删除） | add |

## 消息示例

```json
{
  "productIds": "30798998397239",
  "status": "add"
}
```
