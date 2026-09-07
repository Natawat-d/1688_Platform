# Batch-get payment links for orders

Original name: 批量获取订单的支付链接  
API: `com.alibaba.trade:alibaba.alipay.url.get:1` · Category: Payment  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.alipay.url.get-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.alipay.url.get/{appKey}`  
Requires user authorization (access_token) · Requires signature

When paying through an ERP, use this API to get a cashier link for batch payment. A single order returns the 1688 cashier URL; multiple orders return the Alipay cashier URL. The ERP can redirect the user to the cashier link to complete payment. The user's 1688 login status is verified before payment.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `orderIdList` | Long[] | yes | Order ID list. Up to 100 orders per batch; for Kuajingbao, batches support a maximum of 30 orders. | [74321349391498520] |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `erroMsg` | String | yes | Error message |   |
| `payUrl` | String | yes | Payment link | https://a.b.com |
| `success` | Boolean | yes | Whether successful; may be partially successful, needs to be checked together with payFailureOrderList | true |
| `errorCode` | String | yes | Error code |   |
| `payFailureOrderList` | Long[] | yes | Partially failed order id | [1299871823,19798172783] |

## Error codes

| Error | Symptom | How to fix |
|---|---|---|
| Batch pay : not surport MANUAL-TRADE! | Batch pay : not surport MANUAL-TRADE! | Invitation orders without supplemented buyer receiving information cannot be combined for payment. Whether an order is an invitation order can be determined via the baseInfo.sellerOrder field in the order details |
| Failed to operate on stock: PRODUCT_TRADE_STAT_ERROR | inventoryErrorIds:[16397675**722128**] | There is an order that deducts stock on payment, and the stock deduction failed. |

## Samples

**Output parameter: payment links for multiple orders**

```
{
  "payUrl": "https://trade.1688.com/order/cashier.htm?orderId=154051432607498520;151923545459498520",
  "success": true
}
```

**Output parameter: payment link for a single order**

```
{
  "payUrl": "https://trade.1688.com/order/cashier.htm?orderId=151923545459498520",
  "success": true
}
```
