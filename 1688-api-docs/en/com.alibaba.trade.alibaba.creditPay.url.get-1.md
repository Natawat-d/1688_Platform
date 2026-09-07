# Get payment link for Cheng-e-She credit pay

Original name: 获取使用诚e赊支付的支付链接  
API: `com.alibaba.trade:alibaba.creditPay.url.get:1` · Category: Payment  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.creditPay.url.get-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.creditPay.url.get/{appKey}`  
Requires user authorization (access_token) · Requires signature

Get a payment link for paying with Cheng-e-She (buy-now-pay-later credit).

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `orderIdList` | Long[] | yes | Order Id list, maximum 30 orders per batch; too many orders will cause timeout, 10 orders at a time is recommended | [111111,22222333] |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `success` | String | yes | Whether successful | true |
| `errorCode` | String | yes | Error code | 400_1 |
| `errorMsg` | String | yes | Error description |   |
| `payUrl` | String | yes | Cashier payment link | https://trade.1688.com/order/cashier.htm?orderId=15405143260 |
| `cantPayOrderList` | Long[] | yes | List of orders that cannot be paid in batch due to credit limit or risk control reasons | [123123,23123123] |

## Samples

**Return example**

```
{
  "payUrl": "https://trade.1688.com/order/cashier.htm?orderId=151923545498520",
  "cantPayOrderList":[12123123,12312222222]
  "success": true
}
```
