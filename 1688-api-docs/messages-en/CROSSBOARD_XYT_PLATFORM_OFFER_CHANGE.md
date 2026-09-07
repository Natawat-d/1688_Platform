# XunYuanTong workbench product change

Original name: 寻源通工作台商品变更信息  
Topic: `CROSSBOARD_XYT_PLATFORM_OFFER_CHANGE` · Group: CROSSBOARD (跨境消息)  
Doc page: https://open.1688.com/doc/topicDetail.htm?id=CROSSBOARD_XYT_PLATFORM_OFFER_CHANGE

Product change information from the XunYuanTong (sourcing) workbench

## Payload fields

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `productIds` | String | yes | Set of product IDs, at least one, separated by commas | 30798998397239 |
| `status` | String | yes | Message type, which can specifically be add (added) or delete (deleted) | add |

## Sample message

```json
{
  "productIds": "30798998397239",
  "status": "add"
}
```
