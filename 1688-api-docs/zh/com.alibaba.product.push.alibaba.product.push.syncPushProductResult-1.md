# 同步铺货结果

API: `com.alibaba.product.push:alibaba.product.push.syncPushProductResult:1` · Category: 选品铺货  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.product.push:alibaba.product.push.syncPushProductResult-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.product.push/alibaba.product.push.syncPushProductResult/{appKey}`  
需要授权 (access_token) · 需要签名

同步铺货结果，在源平台(1688)经过ISV把商品铺货到目标平台(比如TAOBAO)时，ISV需要把铺货结果返回。铺货结果的状态描述必须和源平台(1688)定义的一致，同时该接口也支持下架等操作，这些操作都由铺货状态来表述

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `platformDefinition` | [message:alibaba.product.push.PlatformDefinition](#m-alibaba-product-push-platformdefinition) | 是 | 目标平台的定义 | {} |
| `pushRecordIdentity` | [message:alibaba.product.push.Identity](#m-alibaba-product-push-identity) | 否 | 在批量铺货时，源平台可能会为每次铺货产生一个批次传递给ISV，ISV可以在同步通知时返回该字段。该字段由平台传递给ISV，该字段不是必须。 | {} |
| `pushProductResults` | [message:alibaba.product.push.PushProductResult](#m-alibaba-product-push-pushproductresult) | 是 | 商品级别的铺货结果 | {} |

<a id="m-alibaba-product-push-platformdefinition"></a>
#### alibaba.product.push.PlatformDefinition

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `definitionId` | java.lang.String | 是 | 平台ID，由alibaba定义。比如淘宝为www.taobao.com。枚举值：亚马逊Amazon:AMAZON,速卖通:AE,Wish:WISH,Ebay易贝:EBAY,Lazada来赞达:LAZADA,淘宝Taobao:TAOBAO |  |

<a id="m-alibaba-product-push-identity"></a>
#### alibaba.product.push.Identity

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `content` | java.lang.String | 是 |  |  |

<a id="m-alibaba-product-push-pushproductresult"></a>
#### alibaba.product.push.PushProductResult

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `productIdInTargetPlatform` | java.lang.String | 是 | 在目标平台的商品ID，纯数字字符串 |  |
| `productIdInPartner` | java.lang.String | 是 | 在第三方平台的商品ID |  |
| `productIdInSource` | java.lang.String | 是 | 在源平台的商品ID，纯数字字符串 |  |
| `productPushStatus` | java.lang.String | 是 | 铺货状态，0：未成功，1：已成功 |  |
| `productInfoInTargetPlatform` | [message:alibaba.product.push.SimpleItemDesc](#m-alibaba-product-push-simpleitemdesc) | 是 | 商品级别的铺货结果 |  |
| `skus` | [message:alibaba.product.push.PushProductSKUResult[]](#m-alibaba-product-push-pushproductskuresult[]) | 是 | SKU数组 |  |
| `userShopIdInTargetPlatform` | java.lang.String | 是 | 用户在目标平台的店铺标志，纯数字字符串，对应后台的targetUserId。比如目标平台为淘宝，则需要传递客户在淘宝的用户Id |  |

<a id="m-alibaba-product-push-simpleitemdesc"></a>
#### alibaba.product.push.SimpleItemDesc

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `id` | java.lang.String | 是 | 唯一标识 |  |
| `price` | java.lang.Double | 是 | 价格 |  |
| `subject` | java.lang.String | 是 | 商品名称 |  |
| `description` | java.lang.String | 是 | 描述 |  |
| `url` | java.lang.String | 是 | 商品的URL |  |
| `priceRanges` | [message:alibaba.product.push.priceRanges[]](#m-alibaba-product-push-priceranges[]) | 是 | 价格区间 |  |

<a id="m-alibaba-product-push-priceranges[]"></a>
#### alibaba.product.push.priceRanges[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `price` | double | 是 |  |  |
| `startQuantity` | Integer | 是 |  |  |

<a id="m-alibaba-product-push-pushproductskuresult[]"></a>
#### alibaba.product.push.PushProductSKUResult[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `skuIdInSource` | java.lang.String | 是 | 在源平台的SKU标志 | 312312312312 |
| `skuIdInPartner` | java.lang.String | 否 | 在第三方平台的SKU标志 | 101123123 |
| `skuPushStatus` | java.lang.String | 是 | 铺货状态 | success |
| `skuIdInTargetPlatform` | java.lang.String | 是 | 在目标平台的SKU标志 | 10111 |
| `skuInfoInTargetPlatform` | [message:alibaba.product.push.SimpleItemDesc](#m-alibaba-product-push-simpleitemdesc) | 否 | SKU铺货结果 | {} |
| `mappingInfo` | [message:alibaba.product.push.attributeRelationMapping[]](#m-alibaba-product-push-attributerelationmapping[]) | 否 | 属性映射关系 | {} |

<a id="m-alibaba-product-push-attributerelationmapping[]"></a>
#### alibaba.product.push.attributeRelationMapping[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `propertyIdInSource` | java.lang.String | 是 | 在源平台(1688)的属性ID | 123 |
| `propertyTextInSource` | java.lang.String | 是 | 在源平台(1688)的属性ID文本 | 颜色 |
| `valueIdInSource` | java.lang.String | 是 | 在源平台(1688)的属性值ID | 1233 |
| `valueTextInSource` | java.lang.String | 是 | 在源平台(1688)的属性值文本 | 红色 |
| `propertyIdInTarget` | java.lang.String | 是 | 在目标平台(比如TAOBAO)的属性ID | 234 |
| `propertyTextInTarget` | java.lang.String | 是 | 在目标平台(比如TAOBAO)的属性ID文本 | 颜色 |
| `valueIdInTarget` | java.lang.String | 是 | 在目标平台(比如TAOBAO)的属性值的ID | 2341 |
| `valueTextInTarget` | java.lang.String | 是 | 在目标平台(比如TAOBAO)的属性值的文本 | 橙红色 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `errorCode` | String | 是 | 错误码 |   |
| `errorMessage` | String | 是 | 错误描述 |   |
| `success` | boolean | 是 | 是否成功 | true |

## 错误码

| 错误码 | 现象 | 解决方案 |
|---|---|---|
| APPKEY_ISV_CONFIG_IS_NOT_EXIST | 当前isv未在巴拿马平台注册 | 找负责的小二在巴拿马平台添加。 |
| getProductIdInSource_IS_EMPTY | 入参的源ProductId为空 | 检查入参的源productId是否正确 |
| getPlatformDefId_IS_EMPTY | 入参的PlatformDefId为空 | 检查PlatformDefId是否正确 |
| getProductIdInTargetPlatform_IS_EMPTY | 入参的目的ProductId为空 | 入参的目的ProductId是否正确 |

## 示例

**pushProductResults入参示例**

```
{"productIdInTargetPlatform":"565847490150","productIdInSource":"565831341592","productPushStatus":"1","userShopIdInTargetPlatform":"356038588"}
```

**platformDefinition入参示例**

```
{"definitionId":"TAOBAO"}
```
