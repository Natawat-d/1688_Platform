# Add products to the listing list (cross-border)

Original name: 跨境场景下将商品加入铺货列表  
API: `com.alibaba.product.push:alibaba.cross.syncProductListPushed:1` · Category: Product Selection & Listing  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.product.push:alibaba.cross.syncProductListPushed-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.product.push/alibaba.cross.syncProductListPushed/{appKey}`  
Requires user authorization (access_token) · Requires signature

Cross-border only. Adds products to the listing list (that is, creates listing relationships); at most 20 items per call. Only after a product is added can its details be queried through the product-detail interface. Contact the cross-border operations staff to configure permissions manually before calling.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `productIdList` | Long[] | yes | List of 1688 product IDs; the list length must not exceed 20. | [123456] |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.panama.commonResult](#m-alibaba-panama-commonresult) | yes | Sync result | { "success": false} |

<a id="m-alibaba-panama-commonresult"></a>
#### alibaba.panama.commonResult

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `errorCode` | java.lang.String | yes | Error code |  |
| `errorMsg` | java.lang.String | yes | Error message |  |
| `success` | boolean | yes | Whether successful |  |

## Error codes

| Error | Symptom | How to fix |
|---|---|---|
| OFFERIDS_IS_NULL | The product ID list parameter is empty | Check whether the input product exists |
| APPKEY_ISV_CONFIG_IS_NOT_EXIST | Platform appkey configuration does not exist | Contact operations staff to add the configuration on the platform |
| ISVUSERS_IS_NULL | The ISV-side user account list is empty | Sync ISV account information |

## Samples

**Input parameter example**

```
[526415838394]
```
