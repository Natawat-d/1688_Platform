# Get simple product info from a previously purchased supplier

Original name: 获取已购买过商家的商品简单信息  
API: `com.alibaba.product:alibaba.product.simple.get:1` · Category: Products  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.product:alibaba.product.simple.get-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.product/alibaba.product.simple.get/{appKey}`  
Requires user authorization (access_token) · Requires signature

Get product details by product ID. This interface returns simple information for products of suppliers you have already purchased from. Access to this interface is paid. It returns only basic information and is mainly intended for data association in ERP systems.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `productID` | Long | yes | Product ID | 565507182121 |
| `webSite` | String | yes | Site information, specifies whether the called API belongs to the international site (alibaba) or the 1688 website (1688). | 1688 |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `productInfo` | [message:alibaba.product.ProductInfo](#m-alibaba-product-productinfo) | yes | Product detailed information | {} |
| `bizGroupInfos` | [message:alibaba.product.ProductBizGroupInfo[]](#m-alibaba-product-productbizgroupinfo[]) | yes | Support information for the product business; support = false means it is not supported. | [] |
| `errMsg` | String | yes | Return error info | 商品[57053081292]不存在 |

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
| `qualityLevel` | Integer | yes | Quality star rating (1-7) | 5 |
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
| `internationalTradeInfo` | [message:alibaba.product.ProductInternationalTradeInfo](#m-alibaba-product-productinternationaltradeinfo) | yes | Product international trade information; this field does not need to be handled for 1688 |  |
| `bizGroupInfos` | [message:alibaba.product.ProductBizGroupInfo[]](#m-alibaba-product-productbizgroupinfo[]) | yes | Support information for the product business; support = false means it is not supported. | [] |
| `sellerLoginId` | String | yes | Seller's WangWang ID | 卖家旺旺ID |
| `intelligentInfo` | [message:com.alibaba.ocean.openplatform.biz.product.common.model.ProductIntelligentInfo](#m-com-alibaba-ocean-openplatform-biz-product-common-model-productintelligentinfo) | yes | Product algorithm smart-rewrite info, includes the algorithm-optimized product title and image info; if not rewritten, the original title and original image are returned directly | [] |
| `processing` | [message:alibaba.product.ProductProcessing](#m-alibaba-product-productprocessing) | no | Processing/customization information; passing this parameter indicates a processed/customized product | 默认 |
| `productProcessing` | [message:alibaba.product.ProductProcessing](#m-alibaba-product-productprocessing) | yes | Processing/customization information | 默认 |
| `attributeChanges` | String[] | yes | Specific attributes modified for the product | [image] |
| `productLineId` | Long | no | Product line | -1 |
| `processingOfferId` | Long | yes | The custom product id associated with the spot (ready-stock) product | 6388 |
| `sevenDaysRefunds` | Boolean | yes | Whether 7-day no-reason return is supported | true |
| `privateChannelInfo` | [message:com.alibaba.forGetcommon.model.PrivateChannelInfo](#m-com-alibaba-forgetcommon-model-privatechannelinfo) | yes | Product private price information |   |
| `reserveInfo` | [message:alibaba.product.ReserveInfo](#m-alibaba-product-reserveinfo) | yes | Order data | {} |
| `sellStartTime` | Date | yes | Presale start time | 20181219175505000+0800 |
| `productOfficialLogisticsModel` | [message:com.alibaba.ocean.openplatform.biz.product.common.model.ProductOfficialLogisticsModel](#m-com-alibaba-ocean-openplatform-biz-product-common-model-productofficiallogisticsmodel) | yes | Official logistics weight/dimension model | {} |

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
| `consignPrice` | Double | yes | Distribution base price |   |
| `takeSamplePrice` | Double | yes | SKU-level sample price; not applicable to the international site | 1.00 |
| `pftPrice` | Double | yes | Wholesale group price | 0.1 |
| `jxhyPrice` | Double | yes | Curated supply price | 0.1 |
| `jxhyPfPrice` | Double | yes | Curated supply wholesale price | 5.5 |
| `dailySpecialPrice` | Double | yes | Daily flash sale price | 5.4 |
| `eleXdChannelPrice` | Double | yes | Ele.me store-exclusive price | 5.4 |

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
| `attrType` | String | yes | Attribute type | 1 规格属性 2规格扩展属性 ，默认规格属性 |
| `attributeName` | String | yes | The display name corresponding to the SKU attribute ID, e.g. color, size |    |

<a id="m-alibaba-product-productpricerange[]"></a>
#### alibaba.product.ProductPriceRange[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `startQuantity` | Integer | yes | Minimum order quantity | 3 |
| `price` | Double | yes | Price | 445 |

<a id="m-alibaba-product-productsaleinfo"></a>
#### alibaba.product.ProductSaleInfo

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `supportOnlineTrade` | Boolean | yes | Whether online transactions are supported. true: supported, false: not supported. The international site does not need to be concerned with this field. | TRUE |
| `mixWholeSale` | Boolean | yes | Whether mixed batch is supported; the international site does not need to worry about this field | TRUE |
| `saleType` | String | yes | Sales method, sold by piece (normal) or by batch (batch); this field does not need to be considered for the 1688 site | normal |
| `priceAuth` | Boolean | yes | Whether the price is private information; this field does not need to be considered for the international site | TRUE |
| `priceRanges` | [message:alibaba.product.ProductPriceRange[]](#m-alibaba-product-productpricerange[]) | yes | Tiered price. A price set according to quantity range. | [] |
| `amountOnSale` | Double | yes | Available-for-sale quantity; not applicable to the international site | 108 |
| `unit` | String | yes | Unit of measure | 件 |
| `minOrderQuantity` | Integer | yes | Minimum order quantity, range 1-99999. | 3 |
| `batchNumber` | Integer | yes | Quantity per batch, default is empty or a non-zero value; when this attribute is not empty, sellunit is required | 1 |
| `retailprice` | Double | yes | Suggested retail price; not needed for the international site | 20600 |
| `tax` | String | yes | Tax rate related info, content is user-defined, not relevant for the international site |   |
| `sellunit` | String | yes | Selling unit; if items are sold in batches, this indicates the unit of sale. When this attribute is not empty, batchNumber is required, e.g. the &quot;手&quot; (bundle) in 1&quot;手&quot;=12“件&quot; (1 bundle = 12 pieces); not applicable to the international site | 1手等于1件 |
| `quoteType` | Integer | yes | Regular quote-FIXED_PRICE(&quot;0&quot;), SKU spec quote-SKU_PRICE(&quot;1&quot;), SKU range quote (product dimension)-SKU_PRICE_RANGE_FOR_OFFER(&quot;2&quot;), SKU range quote (SKU dimension)-SKU_PRICE_RANGE(&quot;3&quot;); not applicable to the international site | 1 |
| `consignPrice` | Double | yes | Distribution base price. Used in all consignment (distribution) scenarios. For products with SKUs, see consignPrice in skuInfo | 20600 |
| `deliveryLimit` | Integer | yes | Shipping time limit (not the buyer-guarantee shipping cycle), calculated from opening | 5 |
| `invReduceType` | String | yes | Stock deduction method: 1 deducts stock upon order placement, 2 deducts stock upon payment. | 1 |
| `pftPrice` | Double | yes | Wholesale group product price | 0.1 |
| `jxhyPrice` | Double | yes | Curated supply price with free shipping for a single item | 5.0 |
| `jxhyPfPrice` | Double | yes | Curated supply wholesale price | 5.8 |
| `dailySpecialPrice` | Double | yes | Daily flash sale price | 5.5 |
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
| `offerSuttleWeight` | Double | yes | OFFER net weight (kg) | 2 |
| `offerWidth` | Double | yes | Offer packaging width (cm) | 30 |
| `offerHeight` | Double | yes | OFFER package height (cm) | 30 |
| `offerLength` | Double | yes | OFFER packaging length (cm) | 30 |
| `freightTemplate` | [message:alibaba.product.FreightTemplate[]](#m-alibaba-product-freighttemplate[]) | yes | Product shipping fee rate | [] |
| `channelPriceFreePostage` | Boolean | yes | Whether the Changhuotong channel-exclusive price includes free shipping; must be considered together with non-free-shipping addresses — if the shipping address is in a non-free-shipping region, the product does not include free shipping. | true |
| `channelPriceExcludeAreaCodes` | [message:com.alibaba.ocean.openplatform.biz.trade.result.ProductAddressCode[]](#m-com-alibaba-ocean-openplatform-biz-trade-result-productaddresscode[]) | yes | Factory-direct channel exclusive price non-free-shipping regions (address info list, province info) | [{}] |

<a id="m-alibaba-product-freighttemplate[]"></a>
#### alibaba.product.FreightTemplate[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `addressCodeText` | String | yes | Text corresponding to the address region code (including province, city, district, separated by spaces) | 福建省 福州市 鼓楼区 |
| `fromAreaCode` | String | yes | Shipping address region code | 350102 |
| `id` | Long | yes | Address ID | 1234 |
| `name` | String | yes | Template name | 2019 |
| `remark` | String | yes | Remark | 2019 |
| `status` | Integer | yes | Status: 1 means valid, -1 means invalid | 1 |
| `expressSubTemplate` | [message:alibaba.product.DeliverySubTemplateDetailDTO](#m-alibaba-product-deliverysubtemplatedetaildto) | yes | Express delivery sub-template | {} |
| `logisticsSubTemplate` | [message:alibaba.product.DeliverySubTemplateDetailDTO](#m-alibaba-product-deliverysubtemplatedetaildto) | yes | Freight sub-template | {} |
| `codSubTemplate` | [message:alibaba.product.DeliverySubTemplateDetailDTO](#m-alibaba-product-deliverysubtemplatedetaildto) | yes | Cash-on-delivery sub-template | {} |

<a id="m-alibaba-product-deliverysubtemplatedetaildto"></a>
#### alibaba.product.DeliverySubTemplateDetailDTO

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `subTemplateDTO` | [message:alibaba.product.DeliverySubTemplateDTO](#m-alibaba-product-deliverysubtemplatedto) | yes | Sub-template | {} |
| `rateList` | [message:alibaba.product.DeliveryRateDetailDTO[]](#m-alibaba-product-deliveryratedetaildto[]) | yes | Rate | [] |

<a id="m-alibaba-product-deliverysubtemplatedto"></a>
#### alibaba.product.DeliverySubTemplateDTO

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `chargeType` | Integer | yes | Counting type. 0: weight 1: piece count 2: volume | 0 |
| `isSysTemplate` | Boolean | yes | Whether it is a system template | false |
| `serviceChargeType` | Integer | yes | Shipping fee bearer type. Seller bears it: 0; buyer bears it: 1. | 1 |
| `serviceType` | Integer | yes | Service type. 0: express delivery 1: freight 2: cash on delivery | 0 |
| `type` | Integer | yes | Sub-template type. 0: baseline; 1: value-added. Default 0. | 0 |

<a id="m-alibaba-product-deliveryratedetaildto[]"></a>
#### alibaba.product.DeliveryRateDetailDTO[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `isSysRate` | Boolean | yes | Whether it is a system template | false |
| `toAreaCodeText` | String | yes | Address code text, separated by 、 (Chinese enumeration comma). Example: Shanghai、Fujian Province、Guangdong Province | 上海、福建省、广东省 |
| `rateDTO` | [message:alibaba.product.DeliveryRateDTO](#m-alibaba-product-deliveryratedto) | yes | Regular sub-template rate | {} |
| `sysRateDTO` | [message:alibaba.product.DeliverySysRateDTO](#m-alibaba-product-deliverysysratedto) | yes | System sub-template rate | {} |

<a id="m-alibaba-product-deliveryratedto"></a>
#### alibaba.product.DeliveryRateDTO

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `firstUnit` | Long | yes | First weight unit (in grams) or first item unit (in pieces) | 1 |
| `firstUnitFee` | Long | yes | Price for the first weight unit or first item (unit: cents/fen) | 600 |
| `leastExpenses` | Long | yes | Minimum per shipment | 1 |
| `nextUnit` | Long | yes | Additional weight increment unit | 2 |
| `nextUnitFee` | Long | yes | Additional weight unit price (unit: cents/fen) | 100 |

<a id="m-alibaba-product-deliverysysratedto"></a>
#### alibaba.product.DeliverySysRateDTO

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `firstUnit` | Long | yes | First weight unit (in grams) or first item unit (in pieces) | 1 |
| `firstUnitFee` | Long | yes | Price for the first weight unit or first item (unit: cents/fen) | 600 |
| `leastExpenses` | Long | yes | Minimum per shipment | 1 |
| `nextUnit` | Long | yes | Additional weight (in grams) or additional item unit (in pieces) | 1 |
| `nextUnitFee` | Long | yes | Additional weight unit price (unit: cents/fen) | 100 |

<a id="m-com-alibaba-ocean-openplatform-biz-trade-result-productaddresscode[]"></a>
#### com.alibaba.ocean.openplatform.biz.trade.result.ProductAddressCode[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `code` | String | yes | Address code | 620000 |
| `name` | String | yes | Address name | 甘肃省 |

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

<a id="m-alibaba-product-productprocessing"></a>
#### alibaba.product.ProductProcessing

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `images` | String[] | yes | Customized product sample image | ["https://1.jpg","https://1.jpg"] |
| `customType` | String[] | yes | Customization method, optional values: processing, sku | ["processing"] |
| `minPrice` | Double | yes | Lowest price in the price range | 10.00 |
| `maxPrice` | Double | yes | Highest price in the price range | 100.00 |
| `minOrderQuantity` | Long | yes | Minimum order quantity | 50 |
| `stockOfferIds` | Long[] | yes | List of ids of associated spot (ready-stock) products | [6386,6387] |
| `reserveRanges` | [message:alibaba.product.ReserveRange[]](#m-alibaba-product-reserverange[]) | yes | Shipping cycle | [{"quantity":100,"period":10}] |
| `sampleInfo` | [message:alibaba.product.SampleInfo](#m-alibaba-product-sampleinfo) | yes | Sample-making information | 	{"isSupportSample":true,"price":10,"period":2} |

<a id="m-alibaba-product-reserverange[]"></a>
#### alibaba.product.ReserveRange[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `beginAmount` | Long | yes | Order quantity | 100 |
| `date` | Integer | yes | Estimated shipping time (days) | 10 |
| `endAmount` | Long | yes | Order cutoff quantity | 100 |
| `quantity` | Long | yes | Minimum order quantity | 1000 |
| `period` | Integer | yes | Estimated shipping time (days) | 10 |

<a id="m-alibaba-product-sampleinfo"></a>
#### alibaba.product.SampleInfo

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `price` | Double | yes | Sample-making price (unit: yuan) | 10.00 |
| `period` | Integer | yes | Sample-making lead time (unit: days) | 2 |
| `isSupportSample` | Boolean | yes | Whether sampling is supported | true |

<a id="m-com-alibaba-forgetcommon-model-privatechannelinfo"></a>
#### com.alibaba.forGetcommon.model.PrivateChannelInfo

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `offerPrivatePriceInfo` | [message:com.alibaba.forGet.common.model.OfferPrivatePriceInfo[]](#m-com-alibaba-forget-common-model-offerprivatepriceinfo[]) | yes | Private pricing information for private products | [] |

<a id="m-com-alibaba-forget-common-model-offerprivatepriceinfo[]"></a>
#### com.alibaba.forGet.common.model.OfferPrivatePriceInfo[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `skuId` | Long | yes | skuid; 0 for non-SKU-specific pricing | 0 |
| `privatePriceInfo` | String | yes | Specific membership-tier price (regular, advanced, VIP, supreme VIP, entry-level); note the membership order | [] |

<a id="m-alibaba-product-reserveinfo"></a>
#### alibaba.product.ReserveInfo

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `supportReserve` | Boolean | yes | Whether ordering is supported | false |
| `minQuantity` | Long | yes | Minimum order quantity | 10 |
| `maxQuantity` | Long | yes | Maximum order quantity | 1000 |
| `reserveRangeInfos` | [message:alibaba.product.ReserveRangeInfo[]](#m-alibaba-product-reserverangeinfo[]) | yes | Order quantity range information | [] |

<a id="m-alibaba-product-reserverangeinfo[]"></a>
#### alibaba.product.ReserveRangeInfo[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `quantity` | Long | yes | Order quantity | 1000 |
| `period` | Integer | yes | Shipping time (days) | 10 |

<a id="m-com-alibaba-ocean-openplatform-biz-product-common-model-productofficiallogisticsmodel"></a>
#### com.alibaba.ocean.openplatform.biz.product.common.model.ProductOfficialLogisticsModel

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `officialLogisticsInfoModel` | [message:com.alibaba.ocean.openplatform.biz.product.common.model.ProductOfficialLogisticsInfoModel](#m-com-alibaba-ocean-openplatform-biz-product-common-model-productofficiallogisticsinfomodel) | yes | When there is no SKU, the official product logistics model — length, width, height, etc. | {} |
| `officialLogisticsSkuInfoModels` | [message:com.alibaba.ocean.openplatform.biz.product.common.model.ProductOfficialLogisticsSkuInfoModel[]](#m-com-alibaba-ocean-openplatform-biz-product-common-model-productofficiallogisticsskuinfomodel[]) | yes | SKU's logistics model for the official product when a SKU exists, e.g. length, width, height | {} |

<a id="m-com-alibaba-ocean-openplatform-biz-product-common-model-productofficiallogisticsinfomodel"></a>
#### com.alibaba.ocean.openplatform.biz.product.common.model.ProductOfficialLogisticsInfoModel

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `length` | Double | yes | Length | 1 |
| `width` | Double | yes | Width | 1 |
| `height` | Double | yes | Height | 1 |
| `weight` | Long | yes | Weight | 100 |
| `volume` | Double | yes | Volume | 1 |

<a id="m-com-alibaba-ocean-openplatform-biz-product-common-model-productofficiallogisticsskuinfomodel[]"></a>
#### com.alibaba.ocean.openplatform.biz.product.common.model.ProductOfficialLogisticsSkuInfoModel[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `skuId` | Long | yes | skuid | 121315446134 |
| `specId` | String | yes | specid | 1o9KdfOwelir12934kDKFuikA87LIk |
| `length` | Double | yes | Length | 1 |
| `width` | Double | yes | Width | 1 |
| `height` | Double | yes | Height | 1 |
| `weight` | Long | yes | Weight | 100 |
| `volume` | Double | yes | Volume | 1 |

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
| Product [56550718212] does not exist | Product [56550718212] does not exist | Check whether the product ID is correct |

## Samples

**Output parameter example**

```
{
  "productInfo": {
    "productID": 568791940008,
    "categoryID": 122204003,
    "status": "published",
    "subject": "天天特价pzy测试请不要拍-奶疗素-1对2",
    "description": "<p>的范德萨发生</p>",
    "image": {
      "images": [
        "img/ibank/2015/027/083/2685380720_409073960.jpg"
      ]
    },
    "skuInfos": [
      {
        "attributes": [
          {
            "attributeID": 1654,
            "attributeValue": "200mL"
          }
        ],
        "cargoNumber": "",
        "amountOnSale": 6000,
        "skuCode": "3810851782509",
        "skuId": 3810851782509,
        "specId": "e59a659a891496dc81829ef0e5b34d12"
      },
      {
        "attributes": [
          {
            "attributeID": 1654,
            "attributeValue": "1000mL"
          }
        ],
        "cargoNumber": "",
        "amountOnSale": 500,
        "skuCode": "3810851782510",
        "skuId": 3810851782510,
        "specId": "a7776ca120e0ccf6cc14b191b6d68f2e"
      },
      {
        "attributes": [
          {
            "attributeID": 1654,
            "attributeValue": "500ml"
          }
        ],
        "cargoNumber": "",
        "amountOnSale": 36,
        "skuCode": "3810851782511",
        "skuId": 3810851782511,
        "specId": "2a8cfead182437b1723d9622e18a98f3"
      }
    ],
    "referencePrice": "66.0"
  },
  "bizGroupInfos": [
    {
      "support": false,
      "description": "零售通",
      "code": "isRetailOffer"
    },
    {
      "support": false,
      "description": "微供商品",
      "code": "isMicroSupply"
    },
    {
      "support": false,
      "description": "1688企业内部商城",
      "code": "isEMallOffer"
    },
    {
      "support": true,
      "description": "1688代销产品",
      "code": "isConsignMarketOffer"
    },
    {
      "support": false,
      "description": "代发业务产品",
      "code": "hasLinkedToSupplier"
    }
  ]
}
```
