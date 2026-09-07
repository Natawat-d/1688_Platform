# Multilingual product detail

Original name: 多语言商详  
API: `com.alibaba.fenxiao.crossborder:product.search.queryProductDetail:1` · Category: Products  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:product.search.queryProductDetail-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/product.search.queryProductDetail/{appKey}`  
Requires user authorization (access_token) · Requires signature

Multilingual product detail.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `offerDetailParam` | [message:product.search.queryProductDetail.param.OfferDetailParam](#m-product-search-queryproductdetail-param-offerdetailparam) | yes | Parameter | {"offerId":1,"country":"en"} |

<a id="m-product-search-queryproductdetail-param-offerdetailparam"></a>
#### product.search.queryProductDetail.param.OfferDetailParam

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `offerId` | java.lang.Long | yes | Product ID | 1 |
| `country` | java.lang.String | yes | Language | ja-日语 en-英语 |
| `outMemberId` | java.lang.String | no | External user ID | 1 |
| `currency` | String | no | Currency code | HKD |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:product.search.queryProductDetail.resut.ResultModel](#m-product-search-queryproductdetail-resut-resultmodel) | yes | Result | {} |

<a id="m-product-search-queryproductdetail-resut-resultmodel"></a>
#### product.search.queryProductDetail.resut.ResultModel

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `success` | boolean | yes | Whether successful | true |
| `code` | java.lang.String | yes | Return code | 200 |
| `message` | java.lang.String | yes | Prompt | 成功 |
| `result` | [message:product.search.queryProductDetail.model.ProductDetailModel](#m-product-search-queryproductdetail-model-productdetailmodel) | yes | Result | 1 |

<a id="m-product-search-queryproductdetail-model-productdetailmodel"></a>
#### product.search.queryProductDetail.model.ProductDetailModel

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `offerId` | java.lang.Long | yes | Product ID | 34513534353434 |
| `categoryId` | java.lang.Long | yes | Category ID | 32 |
| `categoryName` | java.lang.String | yes | Category name | 裙子 |
| `subject` | java.lang.String | yes | Chinese title | 猫砂盆超大号防外溅猫厕所加高巨无霸开放式宠物用品厂家直销批发 |
| `subjectTrans` | java.lang.String | yes | Translated title | Cat litter box extra large splash-proof cat toilet heightened giant open pet supplies factory direct sales wholesale |
| `description` | java.lang.String | yes | Detail description | 1 |
| `mainVideo` | java.lang.String | yes | Main video | 1 |
| `detailVideo` | java.lang.String | yes | Detail video | 1 |
| `productImage` | [message:product.search.queryProductDetail.model.ProductImage](#m-product-search-queryproductdetail-model-productimage) | yes | Image model | 1 |
| `productAttribute` | [message:product.search.queryProductDetail.model.ProductAttribute[]](#m-product-search-queryproductdetail-model-productattribute[]) | yes | Product CPV attributes | 1 |
| `productSkuInfos` | [message:product.search.queryProductDetail.model.SkuInfo[]](#m-product-search-queryproductdetail-model-skuinfo[]) | yes | Product SKU | 1 |
| `productSaleInfo` | [message:product.search.queryProductDetail.model.ProductSaleInfo](#m-product-search-queryproductdetail-model-productsaleinfo) | yes | Product sales information | 1 |
| `productShippingInfo` | [message:product.search.queryProductDetail.model.ProductShippingInfo](#m-product-search-queryproductdetail-model-productshippinginfo) | yes | Data related to product package delivery | 1 |
| `isJxhy` | boolean | yes | Whether it is curated supply | true |
| `sellerOpenId` | java.lang.String | yes | Merchant's encrypted ID | 23tsdvcdsjngp3oj4j3i5 |
| `minOrderQuantity` | java.lang.Integer | yes | Minimum order quantity model | 1 |
| `batchNumber` | Integer | yes | Quantity per lot | 200 |
| `status` | String | yes | Product status. published: online status; member expired: revoked by member; auto expired: naturally expired; expired: expired (includes both manually and automatically expired); member deleted: deleted by member; modified: modified; new: newly published; deleted: deleted; TBD: to be delete; approved: approved; auditing: under review; untread: review not passed; | published |
| `tagInfoList` | [message:com.alibaba.cbu.offer.model.out.ProductTagInfo[]](#m-com-alibaba-cbu-offer-model-out-producttaginfo[]) | yes | Product service tag | 1 |
| `traceInfo` | String | yes | Tracking point info, used for reporting tracking data to 1688 | object_id@620201390233^object_type@offer |
| `sellerMixSetting` | [message:product.search.queryProductDetail.model.SellerMixSetting](#m-product-search-queryproductdetail-model-sellermixsetting) | yes | Seller mixed batch configuration | 1 |
| `productCargoNumber` | String | yes | Product item number | 1 |
| `sellerDataInfo` | [message:product.search.queryProductDetail.model.SellerDataInfo](#m-product-search-queryproductdetail-model-sellerdatainfo) | yes | Merchant attribute data | {} |
| `soldOut` | String | yes | Product sales volume | 12342 |
| `channelPrice` | [message:product.search.queryProductDetail.model.ChannelPrice](#m-product-search-queryproductdetail-model-channelprice) | yes | Channel price data | 如下 |
| `promotionModel` | [message:com.alibaba.cbu.offer.model.PromotionModel](#m-com-alibaba-cbu-offer-model-promotionmodel) | yes | Marketing | 营销 |
| `tradeScore` | String | yes | Product transaction rating | 5.0 |
| `topCategoryId` | Long | yes | Top-level category | 1 |
| `secondCategoryId` | Long | yes | Secondary category | 2 |
| `thirdCategoryId` | Long | yes | Third-level category | 3 |
| `sellingPoint` | String[] | yes | Multilingual selling points | 卖点:卖点描述 |
| `offerIdentities` | String[] | yes | Merchant identity | 超级工厂,实力商家，诚信通 |
| `createDate` | String | yes | Creation time | 2024-05-24 00:00:00 |
| `isSelect` | String | yes | Cross-border select assortment | true |
| `certificateList` | [message:com.alibaba.cbu.offer.model.OfferCertificateModel[]](#m-com-alibaba-cbu-offer-model-offercertificatemodel[]) | yes | Product certificate list | [             {                 "certificateCode": "ZL 2018 3 0658542.7",       "certificatePhotoList": [      "https://cbu01.alicdn.com/img/ibank/test"        ]              "certificateName": "外观专利证书或授权书证书",   }] |
| `promotionUrl` | String | yes | CPS commission rebate link | 直接跳转1688商品详情链接 |
| `descriptionTrans` | String | yes | Detailed description - after AIGC processing and translation | https://xxx |
| `productImageTrans` | [message:product.search.queryProductDetail.model.ProductImage](#m-product-search-queryproductdetail-model-productimage) | yes | Product image - after AIGC processing and translation | {"whiteImage":"xxx","images":[]} |
| `companyName` | String | yes | Supplier's name (shop name) | 杭州宜路发发电子商务有限公司 |
| `invoiceInfo` | [message:product.search.queryProductDetail.model.InvoiceInfo](#m-product-search-queryproductdetail-model-invoiceinfo) | yes | Invoice information | {     "supportOnlineInvoice": true,     "supportFastInvoice": false,     "invoiceTypes": [         "普票"     ],     "taxpayerType": null } |

<a id="m-product-search-queryproductdetail-model-productimage"></a>
#### product.search.queryProductDetail.model.ProductImage

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `images` | java.lang.String[] | yes | Image | 图片 |
| `whiteImage` | String | yes | White-background image | https://cbu01.alicdn.com/img/ibank/O1CN01pEsVS41Bs2uny9oka_!!0-0-cib.jpg |

<a id="m-product-search-queryproductdetail-model-productattribute[]"></a>
#### product.search.queryProductDetail.model.ProductAttribute[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `attributeId` | java.lang.String | yes | Attribute id | 1 |
| `attributeName` | java.lang.String | yes | Attribute name | 1 |
| `value` | java.lang.String | yes | Attribute value | 1 |
| `attributeNameTrans` | java.lang.String | yes | Attribute name translation | 1 |
| `valueTrans` | java.lang.String | yes | Attribute value translation | 1 |

<a id="m-product-search-queryproductdetail-model-skuinfo[]"></a>
#### product.search.queryProductDetail.model.SkuInfo[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `amountOnSale` | Integer | yes | Stock | 库存 |
| `price` | java.lang.String | yes | Price | 价格 |
| `jxhyPrice` | java.lang.String | yes | Deprecated | 精选货源价格 |
| `skuId` | Long | yes | sku | sku |
| `specId` | java.lang.String | yes | specid | specid |
| `skuAttributes` | [message:product.search.queryProductDetail.model.SkuAttribute[]](#m-product-search-queryproductdetail-model-skuattribute[]) | yes | Attribute | 属性 |
| `pfJxhyPrice` | String | yes | Deprecated | 1 |
| `consignPrice` | String | yes | Deprecated | 1 |
| `cargoNumber` | String | yes | SKU level | 1 |
| `promotionPrice` | String | yes | Marketing price | 1 |
| `fenxiaoPriceInfo` | [message:com.alibaba.cbu.offer.model.FenxiaoPriceInfo](#m-com-alibaba-cbu-offer-model-fenxiaopriceinfo) | yes | Distribution price | 1 |
| `foreignCurrencyPrice` | BigDecimal | yes | Foreign currency price | 3.50 |
| `foreignCurrencyPromotionPrice` | BigDecimal | yes | Foreign currency marketing price | 3.50 |
| `retailPrice` | String | yes | Product retail price | 1.1 |
| `foreignCurrencyRetailPrice` | BigDecimal | yes | Retail price in foreign currency | 1.1 |

<a id="m-product-search-queryproductdetail-model-skuattribute[]"></a>
#### product.search.queryProductDetail.model.SkuAttribute[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `attributeId` | java.lang.Long | yes | Attribute id | 1 |
| `attributeName` | java.lang.String | yes | Attribute name | 1 |
| `attributeNameTrans` | java.lang.String | yes | Attribute name translation | 1 |
| `value` | java.lang.String | yes | Value | 1 |
| `valueTrans` | java.lang.String | yes | Value translation | 1 |
| `skuImageUrl` | java.lang.String | yes | SKU image | 1 |
| `skuImageUrlTrans` | String | yes | SKU image - after AIGC processing and translation | 1 |

<a id="m-com-alibaba-cbu-offer-model-fenxiaopriceinfo"></a>
#### com.alibaba.cbu.offer.model.FenxiaoPriceInfo

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `onePiecePrice` | String | yes | Dropshipping free-shipping price | 15 |
| `offerPrice` | String | yes | Distribution price | 12 |

<a id="m-product-search-queryproductdetail-model-productsaleinfo"></a>
#### product.search.queryProductDetail.model.ProductSaleInfo

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `amountOnSale` | Integer | yes | Product stock | 99999 |
| `priceRangeList` | [message:product.search.queryProductDetail.model.PriceRangeV2[]](#m-product-search-queryproductdetail-model-pricerangev2[]) | yes | Price range | 数组 |
| `foreignCurrencyPrice` | BigInteger | yes | Foreign currency amount | 3.50 |
| `quoteType` | Integer | yes | 0 - no SKU, quote by product quantity; 1 - quote by SKU spec; 2 - has SKU, quote by product quantity | 1 |
| `consignPrice` | String | yes | Deprecated | 1 |
| `jxhyPrice` | String | yes | Deprecated | 1 |
| `unitInfo` | [message:com.alibaba.cbu.offer.model.UnitInfo](#m-com-alibaba-cbu-offer-model-unitinfo) | yes | Unit info | unit |
| `fenxiaoSaleInfo` | [message:com.alibaba.cbu.offer.model.FenxiaoSaleInfo](#m-com-alibaba-cbu-offer-model-fenxiaosaleinfo) | yes | Distribution sales information | 1 |
| `retailPrice` | String | yes | Product retail price | 1.1 |
| `foreignCurrencyRetailPrice` | BigDecimal | yes | Retail price in foreign currency | 1.1 |

<a id="m-product-search-queryproductdetail-model-pricerangev2[]"></a>
#### product.search.queryProductDetail.model.PriceRangeV2[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `startQuantity` | Integer | yes | Minimum order quantity | 10 |
| `price` | String | yes | Wholesale price | 10 |
| `promotionPrice` | String | yes | Marketing price | 9 |
| `foreignCurrencyPrice` | BigDecimal | yes | Foreign currency price | 3.50 |
| `foreignCurrencyPromotionPrice` | BigDecimal | yes | Foreign currency marketing price | 3.50 |

<a id="m-com-alibaba-cbu-offer-model-unitinfo"></a>
#### com.alibaba.cbu.offer.model.UnitInfo

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `unit` | String | yes | Unit in Chinese | 盒 |
| `transUnit` | String | yes | Translation unit | Box |

<a id="m-com-alibaba-cbu-offer-model-fenxiaosaleinfo"></a>
#### com.alibaba.cbu.offer.model.FenxiaoSaleInfo

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `onePieceFreePostage` | Boolean | yes | Whether dropshipping includes free shipping | true |
| `startQuantity` | Integer | yes | Distribution minimum order quantity | 2 |
| `onePiecePrice` | String | yes | Dropshipping free-shipping price | 15 |
| `offerPrice` | String | yes | Distribution price | 12 |

<a id="m-product-search-queryproductdetail-model-productshippinginfo"></a>
#### product.search.queryProductDetail.model.ProductShippingInfo

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `sendGoodsAddressText` | java.lang.String | yes |  |  |
| `weight` | Double | yes | Weight, in kg | 200 |
| `width` | Double | yes | Width, in cm | 1 |
| `height` | Double | yes | Height, in cm | 3 |
| `length` | Double | yes | Length, in cm | 2 |
| `skuShippingInfoList` | [message:com.alibaba.cbu.offer.model.SkuShippingInfo[]](#m-com-alibaba-cbu-offer-model-skushippinginfo[]) | yes | SKU logistics specification information | sku信息 |
| `shippingTimeGuarantee` | String | yes | Shipping guarantee | shipIn24Hours-24小时发货 shipIn48Hours-48小时发货 |
| `skuShippingDetails` | [message:com.alibaba.cbu.offer.model.SkuShippingDetail[]](#m-com-alibaba-cbu-offer-model-skushippingdetail[]) | yes | Set of SKU weight and dimensions | 1 |
| `pkgSizeSource` | String | yes | Source of weight/dimension data | 预测-同款放大，商家自填，官方测量 |
| `officialLength` | Double | yes | Product's officially measured length, in cm | 12 |
| `officialWidth` | Double | yes | Officially measured product width, in cm | 5 |
| `officialHeight` | Double | yes | Officially measured product height, in cm | 14 |
| `officialWeight` | Double | yes | Product official measured weight, unit: kg | 0.001 |

<a id="m-com-alibaba-cbu-offer-model-skushippinginfo[]"></a>
#### com.alibaba.cbu.offer.model.SkuShippingInfo[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `specId` | String | yes | SKU id | 2b36920d5139fd431a2030090d1e2599 |
| `skuId` | Long | yes | skuId | 5104790281451 |
| `width` | Double | yes | Width, in cm | 1 |
| `length` | Double | yes | Length, unit cm | 2 |
| `height` | Double | yes | Height, in cm | 3 |
| `weight` | Long | yes | Weight, in g | 4 |

<a id="m-com-alibaba-cbu-offer-model-skushippingdetail[]"></a>
#### com.alibaba.cbu.offer.model.SkuShippingDetail[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `skuId` | String | yes | skuId | 12123313 |
| `width` | Double | yes | Width, cm | 10 |
| `length` | Double | yes | Length, cm | 10 |
| `height` | Double | yes | Height, cm | 10 |
| `weight` | Double | yes | Weight, kg | 1.2 |
| `pkgSizeSource` | String | yes | Source of piece weight and dimensions | 预测-同款放大，商家自填，官方测量 |
| `officialLength` | Double | yes | Officially measured length, in cm | 12 |
| `officialWidth` | Double | yes | Officially measured width, in cm | 5 |
| `officialHeight` | Double | yes | Officially measured height, in cm | 14 |
| `officialWeight` | Double | yes | Officially measured weight, in kg | 0.001 |
| `aiWeight` | Double | yes | AI-predicted weight, in kg | 0.001 |
| `aiWeightAccuracy` | String | yes | The accuracy of the AI-predicted weight within this product's leaf category | 80% |

<a id="m-com-alibaba-cbu-offer-model-out-producttaginfo[]"></a>
#### com.alibaba.cbu.offer.model.out.ProductTagInfo[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `key` | String | yes | Service name. isOnePsale - dropshipping, select - cross-border select assortment | isOnePsale |
| `value` | Boolean | yes | Whether it is enabled | true |

<a id="m-product-search-queryproductdetail-model-sellermixsetting"></a>
#### product.search.queryProductDetail.model.SellerMixSetting

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `generalHunpi` | Boolean | yes | Whether it is a regular mixed batch | true |
| `mixAmount` | Integer | yes | Mixed batch amount | 100 |
| `mixNumber` | Integer | yes | Mixed batch quantity | 2 |

<a id="m-product-search-queryproductdetail-model-sellerdatainfo"></a>
#### product.search.queryProductDetail.model.SellerDataInfo

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `tradeMedalLevel` | String | yes | Seller trade badge | 5 |
| `compositeServiceScore` | String | yes | Comprehensive service score | 3.5 |
| `logisticsExperienceScore` | String | yes | Logistics experience score | 4.5 |
| `disputeComplaintScore` | String | yes | Dispute resolution score | 3.0 |
| `offerExperienceScore` | String | yes | Product experience score | 4.0 |
| `consultingExperienceScore` | String | yes | Inquiry experience score | 5.0 |
| `repeatPurchasePercent` | String | yes | Seller repeat purchase rate | 0.4666 |
| `afterSalesExperienceScore` | String | yes | Return/exchange experience score | 3.0 |
| `collect30DayWithin48HPercent` | String | yes | 48-hour pickup rate over the last 30 days | 1 |
| `qualityRefundWithin30Day` | String | yes | Quality-related refund rate over the last 30 days | 2 |

<a id="m-product-search-queryproductdetail-model-channelprice"></a>
#### product.search.queryProductDetail.model.ChannelPrice

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `channelSkuPriceList` | [message:product.search.queryProductDetail.model.ChannelSkuPrice[]](#m-product-search-queryproductdetail-model-channelskuprice[]) | yes | Channel SKU price list | 如下 |

<a id="m-product-search-queryproductdetail-model-channelskuprice[]"></a>
#### product.search.queryProductDetail.model.ChannelSkuPrice[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `skuId` | Long | yes | sku id | 435234325 |
| `currentPrice` | String | yes | Channel price | 23.41 |

<a id="m-com-alibaba-cbu-offer-model-promotionmodel"></a>
#### com.alibaba.cbu.offer.model.PromotionModel

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `hasPromotion` | Boolean | yes | Whether there is marketing/promotion | true |
| `promotionType` | String | yes | Marketing type | plus-plus会员 |

<a id="m-com-alibaba-cbu-offer-model-offercertificatemodel[]"></a>
#### com.alibaba.cbu.offer.model.OfferCertificateModel[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `certificateName` | String | yes | Certificate name | 外观专利证书或授权书证书 |
| `certificateCode` | String | yes | Certificate number | ZL 2018 3 0658542.7 |
| `certificatePhotoList` | String[] | yes | Certificate image |   "certificatePhotoList": [                     "https://cbu01.alicdn.com/img/ibank/test.jpg"    ] |

<a id="m-product-search-queryproductdetail-model-invoiceinfo"></a>
#### product.search.queryProductDetail.model.InvoiceInfo

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `supportOnlineInvoice` | Boolean | yes | Whether online invoicing is supported | true |
| `supportFastInvoice` | Boolean | yes | Whether instant invoicing is supported | false |
| `invoiceTypes` | String[] | yes | List of supported invoicing types | ["普票","专票"] |
| `taxpayerType` | String | yes | Taxpayer type | 一般纳税人 \|  小规模纳税人 |

## Samples

**Input parameter example**

```
{
  "offerDetailParam": {
    "offerId": 1,
    "country": "ja-日语 en-英语",
    "outMemberId": "1"
  }
}
```

**Output parameter example**

```
{
  "result": {
    "success": true,
    "code": "200",
    "message": "成功",
    "result": {
      "offerId": 34513534353434,
      "categoryId": 32,
      "categoryName": "裙子",
      "subject": "猫砂盆超大号防外溅猫厕所加高巨无霸开放式宠物用品厂家直销批发",
      "subjectTrans": "Cat litter box extra large splash-proof cat toilet heightened giant open pet supplies factory direct sales wholesale",
      "description": "1",
      "mainVideo": "1",
      "detailVideo": "1",
      "productImage": {
        "images": [
          "图片"
        ],
        "whiteImage": "https://cbu01.alicdn.com/img/ibank/O1CN01pEsVS41Bs2uny9oka_!!0-0-cib.jpg"
      },
      "productAttribute": [
        {
          "attributeId": "1",
          "attributeName": "1",
          "value": "1",
          "attributeNameTrans": "1",
          "valueTrans": "1"
        }
      ],
      "productSkuInfos": [
        {
          "amountOnSale": 123,
          "price": "价格",
          "jxhyPrice": "精选货源价格",
          "skuId": 123,
          "specId": "specid",
          "skuAttributes": [
            {
              "attributeId": 1,
              "attributeName": "1",
              "attributeNameTrans": "1",
              "value": "1",
              "valueTrans": "1",
              "skuImageUrl": "1",
              "skuImageUrlTrans": "1"
            }
          ],
          "pfJxhyPrice": "1",
          "consignPrice": "1",
          "cargoNumber": "1",
          "promotionPrice": "1",
          "fenxiaoPriceInfo": {
            "onePiecePrice": "15",
            "offerPrice": "12"
          }
        }
      ],
      "productSaleInfo": {
        "amountOnSale": 99999,
        "priceRangeList": [
          {
            "startQuantity": 10,
            "price": "10",
            "promotionPrice": "9"
          }
        ],
        "quoteType": 1,
        "consignPrice": "1",
        "jxhyPrice": "1",
        "unitInfo": {
          "unit": "盒",
          "transUnit": "Box"
        },
        "fenxiaoSaleInfo": {
          "onePieceFreePostage": true,
          "startQuantity": 2,
          "onePiecePrice": "15",
          "offerPrice": "12"
        }
      },
      "productShippingInfo": {
        "sendGoodsAddressText": "",
        "weight": 200.0,
        "width": 1.0,
        "height": 3.0,
        "length": 2.0,
        "skuShippingInfoList": [
          {
            "specId": "2b36920d5139fd431a2030090d1e2599",
            "skuId": 5104790281451,
            "width": 1.0,
            "length": 2.0,
            "height": 3.0,
            "weight": 4
          }
        ],
        "shippingTimeGuarantee": "shipIn24Hours-24小时发货 shipIn48Hours-48小时发货",
        "skuShippingDetails": [
          {
            "skuId": "12123313",
            "width": 10.0,
            "length": 10.0,
            "height": 10.0,
            "weight": 1.2,
            "pkgSizeSource": "预测-同款放大，商家自填，官方测量",
            "officialLength": 12.0,
            "officialWidth": 5.0,
            "officialHeight": 14.0,
            "officialWeight": 0.001,
            "aiWeight": 0.001,
            "aiWeightAccuracy": "80%"
          }
        ],
        "pkgSizeSource": "预测-同款放大，商家自填，官方测量",
        "officialLength": 12.0,
        "officialWidth": 5.0,
        "officialHeight": 14.0,
        "officialWeight": 0.001
      },
      "isJxhy": true,
      "sellerOpenId": "23tsdvcdsjngp3oj4j3i5",
      "minOrderQuantity": 1,
      "batchNumber": 200,
      "status": "published",
      "tagInfoList": [
        {
          "key": "isOnePsale",
          "value": true
        }
      ],
      "traceInfo": "object_id@620201390233^object_type@offer",
      "sellerMixSetting": {
        "generalHunpi": true,
        "mixAmount": 100,
        "mixNumber": 2
      },
      "productCargoNumber": "1",
      "sellerDataInfo": {
        "tradeMedalLevel": "5",
        "compositeServiceScore": "3.5",
        "logisticsExperienceScore": "4.5",
        "disputeComplaintScore": "3.0",
        "offerExperienceScore": "4.0",
        "consultingExperienceScore": "5.0",
        "repeatPurchasePercent": "0.4666",
        "afterSalesExperienceScore": "3.0",
        "collect30DayWithin48HPercent": "1",
        "qualityRefundWithin30Day": "2"
      },
      "soldOut": "12342",
      "channelPrice": {
        "channelSkuPriceList": [
          {
            "skuId": 435234325,
            "currentPrice": "23.41"
          }
        ]
      },
      "promotionModel": {
        "hasPromotion": true,
        "promotionType": "plus-plus会员"
      },
      "tradeScore": "5.0",
      "topCategoryId": 1,
      "secondCategoryId": 2,
      "thirdCategoryId": 3,
      "sellingPoint": [
        "卖点:卖点描述"
      ],
      "offerIdentities": [
        "超级工厂,实力商家，诚信通"
      ],
      "createDate": "2024-05-24 00:00:00",
      "isSelect": "true",
      "certificateList": [
        {
          "certificateName": "外观专利证书或授权书证书",
          "certificateCode": "ZL 2018 3 0658542.7",
          "certificatePhotoList": [
            "  \"certificatePhotoList\": [                     \"https://cbu01.alicdn.com/img/ibank/test.jpg\"    ]"
          ]
        }
      ],
      "promotionUrl": "直接跳转1688商品详情链接",
      "descriptionTrans": "https://xxx",
      "productImageTrans": {
        "images": [
          "图片"
        ],
        "whiteImage": "https://cbu01.alicdn.com/img/ibank/O1CN01pEsVS41Bs2uny9oka_!!0-0-cib.jpg"
      },
      "companyName": "杭州宜路发发电子商务有限公司",
      "invoiceInfo": {
        "supportOnlineInvoice": true,
        "supportFastInvoice": false,
        "invoiceTypes": ["普票","专票"]        ,
        "taxpayerType": "一般纳税人"
      }
    }
  }
}
```
