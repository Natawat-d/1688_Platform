# Sync listing results

Original name: 同步铺货结果  
API: `com.alibaba.product.push:alibaba.product.push.syncPushProductResult:1` · Category: Product Selection & Listing  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.product.push:alibaba.product.push.syncPushProductResult-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.product.push/alibaba.product.push.syncPushProductResult/{appKey}`  
Requires user authorization (access_token) · Requires signature

Sync listing results. When an ISV lists products from the source platform (1688) onto a target platform (for example TAOBAO), the ISV must return the listing result. The listing-status descriptions must match those defined by the source platform (1688). This interface also supports operations such as delisting, all of which are expressed through the listing status.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `platformDefinition` | [message:alibaba.product.push.PlatformDefinition](#m-alibaba-product-push-platformdefinition) | yes | Definition of the target platform | {} |
| `pushRecordIdentity` | [message:alibaba.product.push.Identity](#m-alibaba-product-push-identity) | no | During bulk listing, the source platform may generate a batch for each listing operation and pass it to the ISV. The ISV can return this field in the sync notification. This field is passed from the platform to the ISV and is not required. | {} |
| `pushProductResults` | [message:alibaba.product.push.PushProductResult](#m-alibaba-product-push-pushproductresult) | yes | Product-level listing (distribution) result | {} |

<a id="m-alibaba-product-push-platformdefinition"></a>
#### alibaba.product.push.PlatformDefinition

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `definitionId` | java.lang.String | yes | Platform ID, defined by Alibaba. For example, Taobao is www.taobao.com. Enum values: Amazon: AMAZON, AliExpress: AE, Wish: WISH, eBay: EBAY, Lazada: LAZADA, Taobao: TAOBAO |  |

<a id="m-alibaba-product-push-identity"></a>
#### alibaba.product.push.Identity

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `content` | java.lang.String | yes |  |  |

<a id="m-alibaba-product-push-pushproductresult"></a>
#### alibaba.product.push.PushProductResult

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `productIdInTargetPlatform` | java.lang.String | yes | Product ID on the target platform, a numeric-only string |  |
| `productIdInPartner` | java.lang.String | yes | Product ID on the third-party platform |  |
| `productIdInSource` | java.lang.String | yes | Product ID on the source platform, a pure numeric string |  |
| `productPushStatus` | java.lang.String | yes | Listing status. 0: not successful, 1: successful |  |
| `productInfoInTargetPlatform` | [message:alibaba.product.push.SimpleItemDesc](#m-alibaba-product-push-simpleitemdesc) | yes | Product-level listing (distribution) result |  |
| `skus` | [message:alibaba.product.push.PushProductSKUResult[]](#m-alibaba-product-push-pushproductskuresult[]) | yes | SKU array |  |
| `userShopIdInTargetPlatform` | java.lang.String | yes | The user's store identifier on the target platform, a pure numeric string, corresponding to targetUserId on the backend. For example, if the target platform is Taobao, pass the customer's user ID on Taobao. |  |

<a id="m-alibaba-product-push-simpleitemdesc"></a>
#### alibaba.product.push.SimpleItemDesc

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `id` | java.lang.String | yes | Unique identifier |  |
| `price` | java.lang.Double | yes | Price |  |
| `subject` | java.lang.String | yes | Product name |  |
| `description` | java.lang.String | yes | Description |  |
| `url` | java.lang.String | yes | Product URL |  |
| `priceRanges` | [message:alibaba.product.push.priceRanges[]](#m-alibaba-product-push-priceranges[]) | yes | Price range |  |

<a id="m-alibaba-product-push-priceranges[]"></a>
#### alibaba.product.push.priceRanges[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `price` | double | yes |  |  |
| `startQuantity` | Integer | yes |  |  |

<a id="m-alibaba-product-push-pushproductskuresult[]"></a>
#### alibaba.product.push.PushProductSKUResult[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `skuIdInSource` | java.lang.String | yes | SKU identifier on the source platform | 312312312312 |
| `skuIdInPartner` | java.lang.String | no | SKU identifier on the third-party platform | 101123123 |
| `skuPushStatus` | java.lang.String | yes | Listing status | success |
| `skuIdInTargetPlatform` | java.lang.String | yes | SKU identifier on the target platform | 10111 |
| `skuInfoInTargetPlatform` | [message:alibaba.product.push.SimpleItemDesc](#m-alibaba-product-push-simpleitemdesc) | no | SKU listing result | {} |
| `mappingInfo` | [message:alibaba.product.push.attributeRelationMapping[]](#m-alibaba-product-push-attributerelationmapping[]) | no | Attribute mapping relationship | {} |

<a id="m-alibaba-product-push-attributerelationmapping[]"></a>
#### alibaba.product.push.attributeRelationMapping[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `propertyIdInSource` | java.lang.String | yes | The attribute ID on the source platform (1688) | 123 |
| `propertyTextInSource` | java.lang.String | yes | The attribute ID text on the source platform (1688) | 颜色 |
| `valueIdInSource` | java.lang.String | yes | Attribute value ID on the source platform (1688) | 1233 |
| `valueTextInSource` | java.lang.String | yes | The attribute value text on the source platform (1688) | 红色 |
| `propertyIdInTarget` | java.lang.String | yes | The attribute ID on the target platform (e.g. TAOBAO) | 234 |
| `propertyTextInTarget` | java.lang.String | yes | Attribute ID text on the target platform (e.g. TAOBAO) | 颜色 |
| `valueIdInTarget` | java.lang.String | yes | The attribute value ID on the target platform (e.g. TAOBAO) | 2341 |
| `valueTextInTarget` | java.lang.String | yes | Attribute value text on the target platform (e.g. TAOBAO) | 橙红色 |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `errorCode` | String | yes | Error code |   |
| `errorMessage` | String | yes | Error description |   |
| `success` | boolean | yes | Whether successful | true |

## Error codes

| Error | Symptom | How to fix |
|---|---|---|
| APPKEY_ISV_CONFIG_IS_NOT_EXIST | The current ISV is not registered on the Panama platform | Find the responsible Xiaoer (staff) to add it on the Panama platform. |
| getProductIdInSource_IS_EMPTY | The source ProductId in the input parameter is empty | Check whether the input source productId is correct |
| getPlatformDefId_IS_EMPTY | The input parameter PlatformDefId is empty | Check whether PlatformDefId is correct |
| getProductIdInTargetPlatform_IS_EMPTY | The input parameter destination ProductId is empty | Whether the destination ProductId in the input parameters is correct |

## Samples

**Example of pushProductResults input parameters**

```
{"productIdInTargetPlatform":"565847490150","productIdInSource":"565831341592","productPushStatus":"1","userShopIdInTargetPlatform":"356038588"}
```

**Example of the platformDefinition input parameter**

```
{"definitionId":"TAOBAO"}
```
