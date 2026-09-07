# Update order memo

Original name: 修改订单备忘  
API: `com.alibaba.trade:alibaba.order.memoAdd:1` · Category: Orders  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.order.memoAdd-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.order.memoAdd/{appKey}`  
Requires user authorization (access_token) · Requires signature

If the authorized user is the seller, updates the seller memo; if the buyer, updates the buyer memo. Note: this interface can be called repeatedly, and the memo overwrites the content of the previous call.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `orderId` | java.lang.Long | yes | Order ID | 1234567 |
| `memo` | java.lang.String | yes | Memo info | 订单备忘详情 |
| `remarkIcon` | String | yes | Memo icon, currently only supports numbers. 1 is red icon, 2 is blue icon, 3 is green icon, 4 is yellow icon | 2 |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `success` | java.lang.Boolean | yes | Whether successful | true |
| `errorCode` | java.lang.String | yes | Error code | 400_4 |
| `errorMsg` | java.lang.String | yes | Error message | 修改失败 |

## Samples

**Example output parameter for order does not exist**

```
{
  "errorCode": "400_4",
  "errorMsg": "修改备注信息出错[ORDER_NOT_EXIST]，order[1198263322216969811]is not exsits",
  "success": true
}
```
