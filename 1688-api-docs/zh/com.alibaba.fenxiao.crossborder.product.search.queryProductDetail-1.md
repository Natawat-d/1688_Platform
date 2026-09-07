# 多语言商详

API: `com.alibaba.fenxiao.crossborder:product.search.queryProductDetail:1` · Category: 商品  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:product.search.queryProductDetail-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/product.search.queryProductDetail/{appKey}`  
需要授权 (access_token) · 需要签名

多语言商详

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `offerDetailParam` | [message:product.search.queryProductDetail.param.OfferDetailParam](#m-product-search-queryproductdetail-param-offerdetailparam) | 是 | 参数 | {"offerId":1,"country":"en"} |

<a id="m-product-search-queryproductdetail-param-offerdetailparam"></a>
#### product.search.queryProductDetail.param.OfferDetailParam

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `offerId` | java.lang.Long | 是 | 商品id | 1 |
| `country` | java.lang.String | 是 | 语言 | ja-日语 en-英语 |
| `outMemberId` | java.lang.String | 否 | 外部用户id | 1 |
| `currency` | String | 否 | 币种编码 | HKD |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:product.search.queryProductDetail.resut.ResultModel](#m-product-search-queryproductdetail-resut-resultmodel) | 是 | 结果 | {} |

<a id="m-product-search-queryproductdetail-resut-resultmodel"></a>
#### product.search.queryProductDetail.resut.ResultModel

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `success` | boolean | 是 | 是否成功 | true |
| `code` | java.lang.String | 是 | 返回码 | 200 |
| `message` | java.lang.String | 是 | 提示 | 成功 |
| `result` | [message:product.search.queryProductDetail.model.ProductDetailModel](#m-product-search-queryproductdetail-model-productdetailmodel) | 是 | 结果 | 1 |

<a id="m-product-search-queryproductdetail-model-productdetailmodel"></a>
#### product.search.queryProductDetail.model.ProductDetailModel

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `offerId` | java.lang.Long | 是 | 商品id | 34513534353434 |
| `categoryId` | java.lang.Long | 是 | 类目id | 32 |
| `categoryName` | java.lang.String | 是 | 类目名称 | 裙子 |
| `subject` | java.lang.String | 是 | 中文标题 | 猫砂盆超大号防外溅猫厕所加高巨无霸开放式宠物用品厂家直销批发 |
| `subjectTrans` | java.lang.String | 是 | 译文标题 | Cat litter box extra large splash-proof cat toilet heightened giant open pet supplies factory direct sales wholesale |
| `description` | java.lang.String | 是 | 详情描述 | 1 |
| `mainVideo` | java.lang.String | 是 | 主视频 | 1 |
| `detailVideo` | java.lang.String | 是 | 详情视频 | 1 |
| `productImage` | [message:product.search.queryProductDetail.model.ProductImage](#m-product-search-queryproductdetail-model-productimage) | 是 | 图片模型 | 1 |
| `productAttribute` | [message:product.search.queryProductDetail.model.ProductAttribute[]](#m-product-search-queryproductdetail-model-productattribute[]) | 是 | 商品CPV属性 | 1 |
| `productSkuInfos` | [message:product.search.queryProductDetail.model.SkuInfo[]](#m-product-search-queryproductdetail-model-skuinfo[]) | 是 | 商品SKU | 1 |
| `productSaleInfo` | [message:product.search.queryProductDetail.model.ProductSaleInfo](#m-product-search-queryproductdetail-model-productsaleinfo) | 是 | 商品销售信息 | 1 |
| `productShippingInfo` | [message:product.search.queryProductDetail.model.ProductShippingInfo](#m-product-search-queryproductdetail-model-productshippinginfo) | 是 | 商品包裹配送相关数据 | 1 |
| `isJxhy` | boolean | 是 | 是否精选货源 | true |
| `sellerOpenId` | java.lang.String | 是 | 商家加密ID | 23tsdvcdsjngp3oj4j3i5 |
| `minOrderQuantity` | java.lang.Integer | 是 | 最小起批量模型 | 1 |
| `batchNumber` | Integer | 是 | 一手数量 | 200 |
| `status` | String | 是 | 商品状态。published:上网状态;member expired:会员撤销;auto expired:自然过期;expired:过期(包含手动过期与自动过期);member deleted:会员删除;modified:修改;new:新发;deleted:删除;TBD:to be delete;approved:审批通过;auditing:审核中;untread:审核不通过; | published |
| `tagInfoList` | [message:com.alibaba.cbu.offer.model.out.ProductTagInfo[]](#m-com-alibaba-cbu-offer-model-out-producttaginfo[]) | 是 | 商品服务标签 | 1 |
| `traceInfo` | String | 是 | 打点信息，用于向1688上报打点数据 | object_id@620201390233^object_type@offer |
| `sellerMixSetting` | [message:product.search.queryProductDetail.model.SellerMixSetting](#m-product-search-queryproductdetail-model-sellermixsetting) | 是 | 卖家混批配置 | 1 |
| `productCargoNumber` | String | 是 | 商品货号 | 1 |
| `sellerDataInfo` | [message:product.search.queryProductDetail.model.SellerDataInfo](#m-product-search-queryproductdetail-model-sellerdatainfo) | 是 | 商家属性数据 | {} |
| `soldOut` | String | 是 | 商品销量 | 12342 |
| `channelPrice` | [message:product.search.queryProductDetail.model.ChannelPrice](#m-product-search-queryproductdetail-model-channelprice) | 是 | 渠道价格数据 | 如下 |
| `promotionModel` | [message:com.alibaba.cbu.offer.model.PromotionModel](#m-com-alibaba-cbu-offer-model-promotionmodel) | 是 | 营销 | 营销 |
| `tradeScore` | String | 是 | 商品交易评分 | 5.0 |
| `topCategoryId` | Long | 是 | 一级类目 | 1 |
| `secondCategoryId` | Long | 是 | 二级类目 | 2 |
| `thirdCategoryId` | Long | 是 | 三级类目 | 3 |
| `sellingPoint` | String[] | 是 | 多语言卖点 | 卖点:卖点描述 |
| `offerIdentities` | String[] | 是 | 商家身份 | 超级工厂,实力商家，诚信通 |
| `createDate` | String | 是 | 创建时间 | 2024-05-24 00:00:00 |
| `isSelect` | String | 是 | 跨境select货盘 | true |
| `certificateList` | [message:com.alibaba.cbu.offer.model.OfferCertificateModel[]](#m-com-alibaba-cbu-offer-model-offercertificatemodel[]) | 是 | 商品证书列表 | [             {                 "certificateCode": "ZL 2018 3 0658542.7",       "certificatePhotoList": [      "https://cbu01.alicdn.com/img/ibank/test"        ]              "certificateName": "外观专利证书或授权书证书",   }] |
| `promotionUrl` | String | 是 | 返佣cps链接 | 直接跳转1688商品详情链接 |
| `descriptionTrans` | String | 是 | 详情描述-aigc处理和翻译后 | https://xxx |
| `productImageTrans` | [message:product.search.queryProductDetail.model.ProductImage](#m-product-search-queryproductdetail-model-productimage) | 是 | 商品图片-aigc处理和翻译后 | {"whiteImage":"xxx","images":[]} |
| `companyName` | String | 是 | 供应商的名称（店铺名字） | 杭州宜路发发电子商务有限公司 |
| `invoiceInfo` | [message:product.search.queryProductDetail.model.InvoiceInfo](#m-product-search-queryproductdetail-model-invoiceinfo) | 是 | 发票信息 | {     "supportOnlineInvoice": true,     "supportFastInvoice": false,     "invoiceTypes": [         "普票"     ],     "taxpayerType": null } |

<a id="m-product-search-queryproductdetail-model-productimage"></a>
#### product.search.queryProductDetail.model.ProductImage

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `images` | java.lang.String[] | 是 | 图片 | 图片 |
| `whiteImage` | String | 是 | 白底图 | https://cbu01.alicdn.com/img/ibank/O1CN01pEsVS41Bs2uny9oka_!!0-0-cib.jpg |

<a id="m-product-search-queryproductdetail-model-productattribute[]"></a>
#### product.search.queryProductDetail.model.ProductAttribute[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `attributeId` | java.lang.String | 是 | 属性id | 1 |
| `attributeName` | java.lang.String | 是 | 属性名称 | 1 |
| `value` | java.lang.String | 是 | 属性值 | 1 |
| `attributeNameTrans` | java.lang.String | 是 | 属性名称翻译 | 1 |
| `valueTrans` | java.lang.String | 是 | 属性值翻译 | 1 |

<a id="m-product-search-queryproductdetail-model-skuinfo[]"></a>
#### product.search.queryProductDetail.model.SkuInfo[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `amountOnSale` | Integer | 是 | 库存 | 库存 |
| `price` | java.lang.String | 是 | 价格 | 价格 |
| `jxhyPrice` | java.lang.String | 是 | 废弃 | 精选货源价格 |
| `skuId` | Long | 是 | sku | sku |
| `specId` | java.lang.String | 是 | specid | specid |
| `skuAttributes` | [message:product.search.queryProductDetail.model.SkuAttribute[]](#m-product-search-queryproductdetail-model-skuattribute[]) | 是 | 属性 | 属性 |
| `pfJxhyPrice` | String | 是 | 废弃 | 1 |
| `consignPrice` | String | 是 | 废弃 | 1 |
| `cargoNumber` | String | 是 | sku级别 | 1 |
| `promotionPrice` | String | 是 | 营销价 | 1 |
| `fenxiaoPriceInfo` | [message:com.alibaba.cbu.offer.model.FenxiaoPriceInfo](#m-com-alibaba-cbu-offer-model-fenxiaopriceinfo) | 是 | 分销价格 | 1 |
| `foreignCurrencyPrice` | BigDecimal | 是 | 外币价格 | 3.50 |
| `foreignCurrencyPromotionPrice` | BigDecimal | 是 | 外币营销价格 | 3.50 |
| `retailPrice` | String | 是 | 商品零售价 | 1.1 |
| `foreignCurrencyRetailPrice` | BigDecimal | 是 | 零售价外币金额 | 1.1 |

<a id="m-product-search-queryproductdetail-model-skuattribute[]"></a>
#### product.search.queryProductDetail.model.SkuAttribute[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `attributeId` | java.lang.Long | 是 | 属性id | 1 |
| `attributeName` | java.lang.String | 是 | 属性名 | 1 |
| `attributeNameTrans` | java.lang.String | 是 | 属性名翻译 | 1 |
| `value` | java.lang.String | 是 | 值 | 1 |
| `valueTrans` | java.lang.String | 是 | 值翻译 | 1 |
| `skuImageUrl` | java.lang.String | 是 | sku图片 | 1 |
| `skuImageUrlTrans` | String | 是 | sku图片-aigc处理翻译后 | 1 |

<a id="m-com-alibaba-cbu-offer-model-fenxiaopriceinfo"></a>
#### com.alibaba.cbu.offer.model.FenxiaoPriceInfo

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `onePiecePrice` | String | 是 | 一件代发包邮价格 | 15 |
| `offerPrice` | String | 是 | 分销价 | 12 |

<a id="m-product-search-queryproductdetail-model-productsaleinfo"></a>
#### product.search.queryProductDetail.model.ProductSaleInfo

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `amountOnSale` | Integer | 是 | 商品库存 | 99999 |
| `priceRangeList` | [message:product.search.queryProductDetail.model.PriceRangeV2[]](#m-product-search-queryproductdetail-model-pricerangev2[]) | 是 | 价格区间 | 数组 |
| `foreignCurrencyPrice` | BigInteger | 是 | 外币金额 | 3.50 |
| `quoteType` | Integer | 是 | 0-无sku按商品数量报价，1-按sku规格报价 2-有sku按商品数量报价 | 1 |
| `consignPrice` | String | 是 | 废弃 | 1 |
| `jxhyPrice` | String | 是 | 废弃 | 1 |
| `unitInfo` | [message:com.alibaba.cbu.offer.model.UnitInfo](#m-com-alibaba-cbu-offer-model-unitinfo) | 是 | 单位信息 | unit |
| `fenxiaoSaleInfo` | [message:com.alibaba.cbu.offer.model.FenxiaoSaleInfo](#m-com-alibaba-cbu-offer-model-fenxiaosaleinfo) | 是 | 分销销售信息 | 1 |
| `retailPrice` | String | 是 | 商品零售价 | 1.1 |
| `foreignCurrencyRetailPrice` | BigDecimal | 是 | 零售价外币金额 | 1.1 |

<a id="m-product-search-queryproductdetail-model-pricerangev2[]"></a>
#### product.search.queryProductDetail.model.PriceRangeV2[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `startQuantity` | Integer | 是 | 起批量 | 10 |
| `price` | String | 是 | 批发价 | 10 |
| `promotionPrice` | String | 是 | 营销价 | 9 |
| `foreignCurrencyPrice` | BigDecimal | 是 | 外币价格 | 3.50 |
| `foreignCurrencyPromotionPrice` | BigDecimal | 是 | 外币营销价格 | 3.50 |

<a id="m-com-alibaba-cbu-offer-model-unitinfo"></a>
#### com.alibaba.cbu.offer.model.UnitInfo

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `unit` | String | 是 | 中文单位 | 盒 |
| `transUnit` | String | 是 | 译文单位 | Box |

<a id="m-com-alibaba-cbu-offer-model-fenxiaosaleinfo"></a>
#### com.alibaba.cbu.offer.model.FenxiaoSaleInfo

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `onePieceFreePostage` | Boolean | 是 | 是否一件代发包邮 | true |
| `startQuantity` | Integer | 是 | 分销起批量 | 2 |
| `onePiecePrice` | String | 是 | 一件代发包邮价 | 15 |
| `offerPrice` | String | 是 | 分销价 | 12 |

<a id="m-product-search-queryproductdetail-model-productshippinginfo"></a>
#### product.search.queryProductDetail.model.ProductShippingInfo

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `sendGoodsAddressText` | java.lang.String | 是 |  |  |
| `weight` | Double | 是 | 重，单位kg | 200 |
| `width` | Double | 是 | 宽，单位cm | 1 |
| `height` | Double | 是 | 高，单位cm | 3 |
| `length` | Double | 是 | 长，单位cm | 2 |
| `skuShippingInfoList` | [message:com.alibaba.cbu.offer.model.SkuShippingInfo[]](#m-com-alibaba-cbu-offer-model-skushippinginfo[]) | 是 | sku物流规格信息 | sku信息 |
| `shippingTimeGuarantee` | String | 是 | 发货保障 | shipIn24Hours-24小时发货 shipIn48Hours-48小时发货 |
| `skuShippingDetails` | [message:com.alibaba.cbu.offer.model.SkuShippingDetail[]](#m-com-alibaba-cbu-offer-model-skushippingdetail[]) | 是 | sku件重尺集合 | 1 |
| `pkgSizeSource` | String | 是 | 件重尺数据来源 | 预测-同款放大，商家自填，官方测量 |
| `officialLength` | Double | 是 | 商品官方测量长度，单位cm | 12 |
| `officialWidth` | Double | 是 | 商品官方测量宽度，单位cm | 5 |
| `officialHeight` | Double | 是 | 商品官方测量高度，单位cm | 14 |
| `officialWeight` | Double | 是 | 商品官方测量重量,单位kg | 0.001 |

<a id="m-com-alibaba-cbu-offer-model-skushippinginfo[]"></a>
#### com.alibaba.cbu.offer.model.SkuShippingInfo[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `specId` | String | 是 | 规格id | 2b36920d5139fd431a2030090d1e2599 |
| `skuId` | Long | 是 | skuId | 5104790281451 |
| `width` | Double | 是 | 宽,单位cm | 1 |
| `length` | Double | 是 | 长,单位cm | 2 |
| `height` | Double | 是 | 高,单位cm | 3 |
| `weight` | Long | 是 | 重,单位g | 4 |

<a id="m-com-alibaba-cbu-offer-model-skushippingdetail[]"></a>
#### com.alibaba.cbu.offer.model.SkuShippingDetail[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `skuId` | String | 是 | skuId | 12123313 |
| `width` | Double | 是 | 宽，cm | 10 |
| `length` | Double | 是 | 长，cm | 10 |
| `height` | Double | 是 | 高，cm | 10 |
| `weight` | Double | 是 | 重，kg | 1.2 |
| `pkgSizeSource` | String | 是 | 件重尺来源 | 预测-同款放大，商家自填，官方测量 |
| `officialLength` | Double | 是 | 官方测量长度，单位cm | 12 |
| `officialWidth` | Double | 是 | 官方测量宽度，单位cm | 5 |
| `officialHeight` | Double | 是 | 官方测量高度，单位cm | 14 |
| `officialWeight` | Double | 是 | 官方测量重量，单位kg | 0.001 |
| `aiWeight` | Double | 是 | AI 预测重量，单位kg | 0.001 |
| `aiWeightAccuracy` | String | 是 | AI 预测重量在该商品叶子类目准确率 | 80% |

<a id="m-com-alibaba-cbu-offer-model-out-producttaginfo[]"></a>
#### com.alibaba.cbu.offer.model.out.ProductTagInfo[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `key` | String | 是 | 服务名，isOnePsale-一件代发，select-跨境select货盘 | isOnePsale |
| `value` | Boolean | 是 | 是否开通 | true |

<a id="m-product-search-queryproductdetail-model-sellermixsetting"></a>
#### product.search.queryProductDetail.model.SellerMixSetting

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `generalHunpi` | Boolean | 是 | 是否普通混批 | true |
| `mixAmount` | Integer | 是 | 混批金额 | 100 |
| `mixNumber` | Integer | 是 | 混批数量 | 2 |

<a id="m-product-search-queryproductdetail-model-sellerdatainfo"></a>
#### product.search.queryProductDetail.model.SellerDataInfo

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `tradeMedalLevel` | String | 是 | 卖家交易勋章 | 5 |
| `compositeServiceScore` | String | 是 | 综合服务分 | 3.5 |
| `logisticsExperienceScore` | String | 是 | 物流体验分 | 4.5 |
| `disputeComplaintScore` | String | 是 | 纠纷解决分 | 3.0 |
| `offerExperienceScore` | String | 是 | 商品体验分 | 4.0 |
| `consultingExperienceScore` | String | 是 | 咨询体验分 | 5.0 |
| `repeatPurchasePercent` | String | 是 | 卖家回头率 | 0.4666 |
| `afterSalesExperienceScore` | String | 是 | 退换体验分 | 3.0 |
| `collect30DayWithin48HPercent` | String | 是 | 最近30天48H揽收率 | 1 |
| `qualityRefundWithin30Day` | String | 是 | 最近30天品质退款率 | 2 |

<a id="m-product-search-queryproductdetail-model-channelprice"></a>
#### product.search.queryProductDetail.model.ChannelPrice

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `channelSkuPriceList` | [message:product.search.queryProductDetail.model.ChannelSkuPrice[]](#m-product-search-queryproductdetail-model-channelskuprice[]) | 是 | 渠道sku价格列表 | 如下 |

<a id="m-product-search-queryproductdetail-model-channelskuprice[]"></a>
#### product.search.queryProductDetail.model.ChannelSkuPrice[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `skuId` | Long | 是 | sku id | 435234325 |
| `currentPrice` | String | 是 | 渠道价格 | 23.41 |

<a id="m-com-alibaba-cbu-offer-model-promotionmodel"></a>
#### com.alibaba.cbu.offer.model.PromotionModel

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `hasPromotion` | Boolean | 是 | 是否有营销 | true |
| `promotionType` | String | 是 | 营销类型 | plus-plus会员 |

<a id="m-com-alibaba-cbu-offer-model-offercertificatemodel[]"></a>
#### com.alibaba.cbu.offer.model.OfferCertificateModel[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `certificateName` | String | 是 | 证书名字 | 外观专利证书或授权书证书 |
| `certificateCode` | String | 是 | 证书编号 | ZL 2018 3 0658542.7 |
| `certificatePhotoList` | String[] | 是 | 证书图片 |   "certificatePhotoList": [                     "https://cbu01.alicdn.com/img/ibank/test.jpg"    ] |

<a id="m-product-search-queryproductdetail-model-invoiceinfo"></a>
#### product.search.queryProductDetail.model.InvoiceInfo

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `supportOnlineInvoice` | Boolean | 是 | 是否支持在线开票 | true |
| `supportFastInvoice` | Boolean | 是 | 是否支持极速开票 | false |
| `invoiceTypes` | String[] | 是 | 支持的开票类型列表 | ["普票","专票"] |
| `taxpayerType` | String | 是 | 纳税人类型 | 一般纳税人 \|  小规模纳税人 |

## 示例

**入参示例**

```
{
  "offerDetailParam": {
    "offerId": 1,
    "country": "ja-日语 en-英语",
    "outMemberId": "1"
  }
}
```

**出参示例**

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
