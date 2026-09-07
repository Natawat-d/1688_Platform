# Get product details (cross-border)

Original name: 跨境场景获取商品详情  
API: `com.alibaba.product:alibaba.cross.productInfo:1` · Category: Product Selection & Listing  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.product:alibaba.cross.productInfo-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.product/alibaba.cross.productInfo/{appKey}`  
Requires user authorization (access_token) · Requires signature

Get product details in cross-border scenarios. A cross-border listing relationship must be established before details can be retrieved.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `productId` | java.lang.Long | yes | 1688 product ID | 573741401425 |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `productInfo` | [message:alibaba.product.ProductInfo](#m-alibaba-product-productinfo) | yes | Product details |  {} |
| `bizGroupInfos` | [message:alibaba.product.ProductBizGroupInfo[]](#m-alibaba-product-productbizgroupinfo[]) | yes | Support information for the product business; support = false means it is not supported. | {} |
| `success` | Boolean | yes | Whether successful | true |
| `message` | String | yes | Call information |   |

<a id="m-alibaba-product-productinfo"></a>
#### alibaba.product.ProductInfo

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `productID` | Long | yes | Product ID | 584051070147 |
| `productType` | String | yes | Product type: online wholesale product (wholesale) or inquiry product (sourcing). Defaults to wholesale on the 1688 website. | wholesale |
| `categoryID` | Long | yes | Category ID, identifies the category the product belongs to | 1048182 |
| `attributes` | [message:alibaba.product.ProductAttribute[]](#m-alibaba-product-productattribute[]) | yes | Product attributes and attribute values | [] |
| `groupID` | Long[] | yes | Group ID, determines the group the product belongs to. On 1688, multiple group IDs can be passed in; on the international site, a product can belong to only one group, so by default only the first one is taken. | [107331682] |
| `status` | String | yes | Product status. published: online status; member expired: revoked by member; auto expired: naturally expired; expired: expired (includes both manually and automatically expired); member deleted: deleted by member; modified: modified; new: newly published; deleted: deleted; TBD: to be delete; approved: approved; auditing: under review; untread: review not passed; | published |
| `subject` | String | yes | Product title, up to 128 characters | 高端气质OL韩版雪纺女装套头半高领长袖修身型蕾丝衫 |
| `description` | String | yes | Product detail description, may include image URLs from the image center | 高端气质OL韩版雪纺女装套头半高领长袖修身型蕾丝衫 |
| `language` | String | yes | Language; see the FAQ for language enum values. The 1688 website passes CHINESE by default | ENGLISH |
| `periodOfValidity` | Integer | yes | Information validity period, calculated in days; not applicable for the international site | 3650 |
| `bizType` | Integer | yes | Business type. 1: Product, 2: Processing, 3: Agency, 4: Cooperation, 5: Business service. The international site defaults to Product. | 1 |
| `pictureAuth` | Boolean | yes | Whether the image is private information; this field is invalid for the international site | false |
| `image` | [message:alibaba.product.ProductImageInfo](#m-alibaba-product-productimageinfo) | yes | Product main image | {} |
| `skuInfos` | [message:alibaba.product.ProductSKUInfo[]](#m-alibaba-product-productskuinfo[]) | yes | SKU information | [] |
| `saleInfo` | [message:alibaba.product.ProductSaleInfo](#m-alibaba-product-productsaleinfo) | yes | Product sales information | {} |
| `shippingInfo` | [message:alibaba.product.ProductShippingInfo](#m-alibaba-product-productshippinginfo) | yes | Product logistics information | {} |
| `extendInfos` | [message:alibaba.product.ProductExtendInfo[]](#m-alibaba-product-productextendinfo[]) | yes | Product extension info | [] |
| `supplierUserId` | java.lang.String | yes | Supplier user ID | 1234 |
| `qualityLevel` | Integer | yes | Quality star rating (0-5) | 5 |
| `supplierLoginId` | java.lang.String | yes | Supplier loginId | alitestforisv01 |
| `categoryName` | java.lang.String | yes | Category name | 连衣裙 |
| `mainVedio` | java.lang.String | yes | Main image video playback URL | https://cloud.video.taobao.com/play/u/1685/p/1/e/6/t/1/5224**.mp4 |
| `productCargoNumber` | java.lang.String | yes | Product model/article number, the article number in product attributes | 666 |
| `crossBorderOffer` | Boolean | yes | Whether it is overseas dropshipping | true |
| `referencePrice` | java.lang.String | yes | Reference price; returns a price range, may be empty | 500 |
| `createTime` | java.util.Date | yes | Creation time | 20181213201638000+0800 |
| `lastUpdateTime` | java.util.Date | yes | Last operation time | 20181219175505000+0800 |
| `expireTime` | java.util.Date | yes | Expiration time | 20281216175505000+0800 |
| `modifyTime` | java.util.Date | yes | Modification time | 20281216175505000+0800 |
| `approvedTime` | java.util.Date | yes | Review time | 20181219175505000+0800 |
| `lastRepostTime` | java.util.Date | yes | Last resend time | 20181217090842000+0800 |
| `bookedCount` | String | yes | Transaction volume | 1999 |
| `productLine` | String | yes | Product line | 默认 |
| `detailVedio` | String | yes | Detail video | https://cloud.video.taobao.com/play/u/1685/p/1/e/6/t/1/5224**.mp4 |
| `internationalTradeInfo` | [message:alibaba.product.ProductInternationalTradeInfo](#m-alibaba-product-productinternationaltradeinfo) | yes | Product international trade information; this field does not need to be handled for 1688 |   |
| `bizGroupInfos` | [message:alibaba.product.ProductBizGroupInfo[]](#m-alibaba-product-productbizgroupinfo[]) | yes | Support information for the product business; support = false means it is not supported. | [] |
| `sellerLoginId` | String | yes | Seller's WangWang ID | 卖家旺旺ID |
| `intelligentInfo` | [message:com.alibaba.ocean.openplatform.biz.product.common.model.ProductIntelligentInfo](#m-com-alibaba-ocean-openplatform-biz-product-common-model-productintelligentinfo) | yes | Product algorithm smart-rewrite info, includes the algorithm-optimized product title and image info; if not rewritten, the original title and original image are returned directly | [] |
| `sellStartTime` | Date | yes | Presale start time | 20181219175505000+0800 |
| `sevenDaysRefunds` | Boolean | yes | Whether 7-day no-reason return is supported | true |

<a id="m-alibaba-product-productattribute[]"></a>
#### alibaba.product.ProductAttribute[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `attributeID` | Long | yes | Attribute ID | 123456 |
| `attributeName` | String | yes | Attribute name | color |
| `valueID` | Long | yes | Attribute value ID | 123456 |
| `value` | String | yes | Attribute value | grey |
| `isCustom` | Boolean | yes | Whether it is a custom attribute; not applicable to the international site | true |

<a id="m-alibaba-product-productimageinfo"></a>
#### alibaba.product.ProductImageInfo

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `images` | String[] | yes | List of main images, using relative paths; the domain needs to be prepended: https://cbu01.alicdn.com/ | ["img/ibank/2014/766/624/1652426667_642119312.jpg","img/ibank/2014/656/624/1652426656_642119312.jpg","img/ibank/2014/236/624/1652426632_642119312.jpg"] |
| `isWatermark` | Boolean | yes | Whether to add a watermark, yes (true) or no (false). 1688 does not need to be concerned with this field; 1688's watermark information is handled when the image is uploaded. |  |
| `isWatermarkFrame` | Boolean | yes | Whether the watermark has a border, with border (true) or without border (false). 1688 does not need to worry about this field; 1688's watermark information is processed when the image is uploaded |  |
| `watermarkPosition` | String | yes | Watermark position, either center or bottom. 1688 does not need to use this field; 1688's watermark information is processed when the image is uploaded |  |

<a id="m-alibaba-product-productskuinfo[]"></a>
#### alibaba.product.ProductSKUInfo[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `attributes` | [message:alibaba.product.SKUAttrInfo[]](#m-alibaba-product-skuattrinfo[]) | yes | SKU attribute value; multiple sets of information may be entered | [] |
| `cargoNumber` | String | yes | Product number for the specified SKU; not applicable to the international site | cy小黄人 |
| `amountOnSale` | Integer | yes | Sellable quantity; not applicable for the international site. | 1200 |
| `retailPrice` | double | yes | Suggested retail price; not needed for the international site | 40 |
| `price` | Double | yes | The unit price of this SKU at the time of quotation. Note for the international site: this value is used when setting a specific price for an online wholesale product with SKU attributes; if tiered pricing is set, use priceRange instead. | 40 |
| `priceRange` | [message:alibaba.product.ProductPriceRange[]](#m-alibaba-product-productpricerange[]) | yes | Tiered quotation; not applicable to 1688 | [] |
| `skuCode` | String | yes | Product code; not applicable to 1688 | 3746778972330 |
| `skuId` | Long | yes | skuId; not applicable for the international site. | 3746778972330 |
| `specId` | String | yes | specId; not applicable to the international site | fb82997a5b64af3c729cea89bef44409 |
| `consignPrice` | Double | yes | Distribution base price |  0.2 |
| `pftPrice` | Double | yes | Wholesale group price | 0.1 |
| `jxhyPrice` | Double | yes | Curated supply price - single-item free-shipping price | 5.9 |
| `jxhyPfPrice` | Double | yes | Curated supply wholesale price | 5.5 |

<a id="m-alibaba-product-skuattrinfo[]"></a>
#### alibaba.product.SKUAttrInfo[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `attributeID` | Long | yes | SKU attribute ID | 3216 |
| `attValueID` | Long | yes | SKU value ID; not applicable on 1688. |   |
| `attributeValue` | String | yes | SKU value content; not applicable to the international site | 黑色 |
| `customValueName` | String | yes | Custom attribute value name; not applicable to 1688 |   |
| `skuImageUrl` | String | yes | SKU image | img/ibank/2018/221/909/9143909122_1606139362.jpg |
| `attributeDisplayName` | String | yes | The display name corresponding to the SKU attribute ID, e.g. color, size |   |

<a id="m-alibaba-product-productpricerange[]"></a>
#### alibaba.product.ProductPriceRange[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `startQuantity` | Integer | yes | Minimum order quantity | 1 |
| `price` | Double | yes | Product price | 40 |

<a id="m-alibaba-product-productsaleinfo"></a>
#### alibaba.product.ProductSaleInfo

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `supportOnlineTrade` | Boolean | yes | Whether online transactions are supported. true: supported, false: not supported. The international site does not need to be concerned with this field. | TRUE |
| `mixWholeSale` | Boolean | yes | Whether mixed batch is supported; the international site does not need to worry about this field | FALSE |
| `saleType` | String | yes | Sales method, sold by piece (normal) or by batch (batch); this field does not need to be considered for the 1688 site | batch |
| `priceAuth` | Boolean | yes | Whether the price is private information; this field does not need to be considered for the international site | FALSE |
| `priceRanges` | [message:alibaba.product.ProductPriceRange[]](#m-alibaba-product-productpricerange[]) | yes | Tiered price. A price set according to quantity range. | [] |
| `amountOnSale` | Double | yes | Available-for-sale quantity; not applicable to the international site | 24000 |
| `unit` | String | yes | Unit of measure | 件 |
| `minOrderQuantity` | Integer | yes | Minimum order quantity, range 1-99999. | 1 |
| `batchNumber` | Integer | yes | Quantity per batch, default is empty or a non-zero value; when this attribute is not empty, sellunit is required | 1 |
| `retailprice` | Double | yes | Suggested retail price; not needed for the international site | 40 |
| `tax` | String | yes | Tax rate related info, content is user-defined, not relevant for the international site |   |
| `sellunit` | String | yes | Selling unit; if items are sold in batches, this indicates the unit of sale. When this attribute is not empty, batchNumber is required, e.g. the &quot;手&quot; (bundle) in 1&quot;手&quot;=12“件&quot; (1 bundle = 12 pieces); not applicable to the international site | 1手等于1件 |
| `quoteType` | Integer | yes | Regular quote-FIXED_PRICE(&quot;0&quot;), SKU spec quote-SKU_PRICE(&quot;1&quot;), SKU range quote (product dimension)-SKU_PRICE_RANGE_FOR_OFFER(&quot;2&quot;), SKU range quote (SKU dimension)-SKU_PRICE_RANGE(&quot;3&quot;); not applicable to the international site | 1 |
| `consignPrice` | Double | yes | Distribution base price |  0.5 |
| `pftPrice` | Double | yes | Wholesale group product price | 0.1 |
| `jxhyPrice` | Double | yes | Curated supply price with free shipping for a single item | 5.0 |
| `jxhyPfPrice` | Double | yes | Curated supply wholesale price | 5.8 |
| `beginAmount` | Integer | yes | Wholesale minimum order quantity | 2 |
| `retailBeginAmount` | Long | yes | Retail minimum order quantity | 1 |

<a id="m-alibaba-product-productshippinginfo"></a>
#### alibaba.product.ProductShippingInfo

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `unitWeight` | Double | yes | Weight/gross weight, unit: kg/piece | 121 |
| `volume` | Integer | yes | Volume, in cubic centimeters, range 1-9999999. This field does not need to be considered for 1688. | 500 |
| `handlingTime` | Integer | yes | Preparation lead time, in days, range 1-60. 1688 does not need to process this field. | 12 |
| `freightTemplateID` | Long | yes | Freight template ID. 0 means freight description, 1 means seller bears the freight, other values mean using a freight template. This parameter can be obtained via freight-template-related APIs | 121133 |
| `suttleWeight` | Double | yes | Net weight, in kg per unit | 1001 |
| `sendGoodsAddressText` | java.lang.String | yes | Shipping origin description | asda |
| `width` | Double | yes | Width, in centimeters | 30 |
| `height` | Double | yes | Height, in centimeters | 20 |
| `length` | Double | yes | Length, unit: centimeters | 10 |
| `packageSize` | String | yes | Dimensions, in centimeters; length, width, and height each range from 1-9999999. This field does not need to be considered for 1688 | 10x20x50 |
| `sendGoodsAddressId` | Long | yes | Shipping origin address ID; not applicable to the international site | 124431 |

<a id="m-alibaba-product-productextendinfo[]"></a>
#### alibaba.product.ProductExtendInfo[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `key` | String | yes | key of the extension structure | 代销价格,consignPrice;<br>买家保障,buyerProtection; |
| `value` | String | yes | value of the extension structure | 代销价格,key为skuId，value为用户设置的代销价，<br>示例：31151771910:2088.0;31151771909:2088.0;31151771908:2088.0;31152339121:2088.0;<br>买家保障,string数组，value为买保全拼，<br>示例：["psbj","swtwlybt","swtbh","ssbxsfh"] |

<a id="m-alibaba-product-productinternationaltradeinfo"></a>
#### alibaba.product.ProductInternationalTradeInfo

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `fobCurrency` | String | yes | FOB price currency; see the FAQ for the currency enum values |  |
| `fobMinPrice` | String | yes | FOB minimum price |  |
| `fobMaxPrice` | String | yes | FOB maximum price |  |
| `fobUnitType` | String | yes | FOB unit of measure, see FAQ for unit-of-measure enum values |  |
| `paymentMethods` | String[] | yes | Payment method; see the FAQ for the payment method enum values |  |
| `minOrderQuantity` | Integer | yes | Minimum order quantity |  |
| `minOrderUnitType` | String | yes | Minimum order quantity unit of measure, see FAQ for unit-of-measure enum values |  |
| `supplyQuantity` | Integer | yes | supplyQuantity |  |
| `supplyUnitType` | String | yes | Supply capacity unit of measurement; see the FAQ for the unit of measurement enum values |  |
| `supplyPeriodType` | String | yes | Supply capability cycle; see the FAQ for the time period enum values |  |
| `deliveryPort` | String | yes | Shipping port |  |
| `deliveryTime` | String | yes | Shipping deadline |  |
| `consignmentDate` | Integer | yes | New shipping deadline |  |
| `packagingDesc` | String | yes | Standard packaging |  |

<a id="m-com-alibaba-ocean-openplatform-biz-product-common-model-productintelligentinfo"></a>
#### com.alibaba.ocean.openplatform.biz.product.common.model.ProductIntelligentInfo

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `title` | String | yes | Product title optimized by algorithm | 算法优化后的商品标题 |
| `images` | String[] | yes | Product image after algorithm optimization | ["https://cbu01.alicdn.com/img/ibank/2020/932/210/13529012239_321095253.jpg"] |
| `skuImages` | [message:com.alibaba.ocean.openplatform.biz.product.common.model.SkuIntelligentInfo[]](#m-com-alibaba-ocean-openplatform-biz-product-common-model-skuintelligentinfo[]) | yes | Algorithm-optimized spec image | [] |
| `descriptionImages` | String[] | yes | Detail image after algorithm optimization | ["https://cbu01.alicdn.com/img/ibank/2020/932/210/13529012239_321095253.jpg"] |

<a id="m-com-alibaba-ocean-openplatform-biz-product-common-model-skuintelligentinfo[]"></a>
#### com.alibaba.ocean.openplatform.biz.product.common.model.SkuIntelligentInfo[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `skuId` | Long | yes | skuId | 123433333 |
| `imageUrl` | String | yes | The image URL after algorithmic processing; if not processed, the original image URL is returned. | https://cbu01.alicdn.com/img/ibank/2020/932/210/13529012239_321095253.jpg |

<a id="m-alibaba-product-productbizgroupinfo[]"></a>
#### alibaba.product.ProductBizGroupInfo[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `support` | java.lang.Boolean | yes | Whether supported |  |
| `description` | java.lang.String | yes | Vertical market name, e.g. Weigong Market, Goods Market |  |
| `code` | java.lang.String | yes | Vertical market flag |  |

## Error codes

| Error | Symptom | How to fix |
|---|---|---|
| No permission to query this product | No permission to query this product | No listing relationship has been established; click one-click listing on the product page, or call the listing sync API. |
| Product [57374140142] does not exist | Product [57374140142] does not exist | Product ID error, please check the product ID |

## Samples

**Input parameter example**

```
{
  productID:573741401425
}
```
