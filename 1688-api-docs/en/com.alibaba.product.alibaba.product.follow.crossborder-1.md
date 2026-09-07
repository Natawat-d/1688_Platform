# Follow a product

Original name: 关注商品  
API: `com.alibaba.product:alibaba.product.follow.crossborder:1` · Category: Product Selection & Listing  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.product:alibaba.product.follow.crossborder-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.product/alibaba.product.follow.crossborder/{appKey}`  
Requires user authorization (access_token) · Requires signature

Follow a product.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `productId` | Long | yes | Product ID | 52312121144 |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `code` | int | yes | code, 0 means success | 0 |
| `message` | String | yes | Description of the result | 成功 |

## Samples

**Output parameter example document**

```
{
  "code": 0,
  "message": "success"
}
```
