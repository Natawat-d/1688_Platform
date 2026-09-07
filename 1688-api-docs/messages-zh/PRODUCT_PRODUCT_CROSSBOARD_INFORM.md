# 一键铺货消息

Topic: `PRODUCT_PRODUCT_CROSSBOARD_INFORM` · Group: PRODUCT (商品消息)  
Doc page: https://open.1688.com/doc/topicDetail.htm?id=PRODUCT_PRODUCT_CROSSBOARD_INFORM

用户触发一键铺货

## 消息字段

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `offerId` | String | 是 | 商品ID | 570046582711 |
| `userInfo` | String | 是 | 分销商CBU的会员ID | 2804951212 |
| `action` | String | 是 | 执行动作，值为distribution | distribution |

## 消息示例

```json
{
  "offerId": "570046582711",
  "userInfo": "2804951212",
  "action": "distribution"
}
```
