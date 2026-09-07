# 跨境场景获取商品详情

API: `com.alibaba.product:alibaba.cross.productInfo:1` · Category: 选品铺货  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.product:alibaba.cross.productInfo-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.product/alibaba.cross.productInfo/{appKey}`  
需要授权 (access_token) · 需要签名

跨境场景获取商品详情，需要建立跨境铺货关系之后才能获取详情

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `productId` | java.lang.Long | 是 | 1688商品ID | 573741401425 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `productInfo` | [message:alibaba.product.ProductInfo](#m-alibaba-product-productinfo) | 是 | 商品详情 |  {} |
| `bizGroupInfos` | [message:alibaba.product.ProductBizGroupInfo[]](#m-alibaba-product-productbizgroupinfo[]) | 是 | 产品业务的支持信息,support为false说明不支持. | {} |
| `success` | Boolean | 是 | 是否成功 | true |
| `message` | String | 是 | 调用信息 |   |

<a id="m-alibaba-product-productinfo"></a>
#### alibaba.product.ProductInfo

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `productID` | Long | 是 | 商品ID | 584051070147 |
| `productType` | String | 是 | 商品类型，在线批发商品(wholesale)或者询盘商品(sourcing)，1688网站缺省为wholesale | wholesale |
| `categoryID` | Long | 是 | 类目ID，标识商品所属类目 | 1048182 |
| `attributes` | [message:alibaba.product.ProductAttribute[]](#m-alibaba-product-productattribute[]) | 是 | 商品属性和属性值 | [] |
| `groupID` | Long[] | 是 | 分组ID，确定商品所属分组。1688可传入多个分组ID，国际站同一个商品只能属于一个分组，因此默认只取第一个 | [107331682] |
| `status` | String | 是 | 商品状态。published:上网状态;member expired:会员撤销;auto expired:自然过期;expired:过期(包含手动过期与自动过期);member deleted:会员删除;modified:修改;new:新发;deleted:删除;TBD:to be delete;approved:审批通过;auditing:审核中;untread:审核不通过; | published |
| `subject` | String | 是 | 商品标题，最多128个字符 | 高端气质OL韩版雪纺女装套头半高领长袖修身型蕾丝衫 |
| `description` | String | 是 | 商品详情描述，可包含图片中心的图片URL | 高端气质OL韩版雪纺女装套头半高领长袖修身型蕾丝衫 |
| `language` | String | 是 | 语种，参见FAQ 语种枚举值，1688网站默认传入CHINESE | ENGLISH |
| `periodOfValidity` | Integer | 是 | 信息有效期，按天计算，国际站无此信息 | 3650 |
| `bizType` | Integer | 是 | 业务类型。1：商品，2：加工，3：代理，4：合作，5：商务服务。国际站按默认商品。 | 1 |
| `pictureAuth` | Boolean | 是 | 是否图片私密信息，国际站此字段无效 | false |
| `image` | [message:alibaba.product.ProductImageInfo](#m-alibaba-product-productimageinfo) | 是 | 商品主图 | {} |
| `skuInfos` | [message:alibaba.product.ProductSKUInfo[]](#m-alibaba-product-productskuinfo[]) | 是 | sku信息 | [] |
| `saleInfo` | [message:alibaba.product.ProductSaleInfo](#m-alibaba-product-productsaleinfo) | 是 | 商品销售信息 | {} |
| `shippingInfo` | [message:alibaba.product.ProductShippingInfo](#m-alibaba-product-productshippinginfo) | 是 | 商品物流信息 | {} |
| `extendInfos` | [message:alibaba.product.ProductExtendInfo[]](#m-alibaba-product-productextendinfo[]) | 是 | 商品扩展信息 | [] |
| `supplierUserId` | java.lang.String | 是 | 供应商用户ID | 1234 |
| `qualityLevel` | Integer | 是 | 质量星级(0-5) | 5 |
| `supplierLoginId` | java.lang.String | 是 | 供应商loginId | alitestforisv01 |
| `categoryName` | java.lang.String | 是 | 类目名 | 连衣裙 |
| `mainVedio` | java.lang.String | 是 | 主图视频播放地址 | https://cloud.video.taobao.com/play/u/1685/p/1/e/6/t/1/5224**.mp4 |
| `productCargoNumber` | java.lang.String | 是 | 商品货号，产品属性中的货号 | 666 |
| `crossBorderOffer` | Boolean | 是 | 是否海外代发 | true |
| `referencePrice` | java.lang.String | 是 | 参考价格，返回价格区间，可能为空 | 500 |
| `createTime` | java.util.Date | 是 | 创建时间 | 20181213201638000+0800 |
| `lastUpdateTime` | java.util.Date | 是 | 最后操作时间 | 20181219175505000+0800 |
| `expireTime` | java.util.Date | 是 | 过期时间 | 20281216175505000+0800 |
| `modifyTime` | java.util.Date | 是 | 修改时间 | 20281216175505000+0800 |
| `approvedTime` | java.util.Date | 是 | 审核时间 | 20181219175505000+0800 |
| `lastRepostTime` | java.util.Date | 是 | 最后重发时间 | 20181217090842000+0800 |
| `bookedCount` | String | 是 | 成交量 | 1999 |
| `productLine` | String | 是 | 产品线 | 默认 |
| `detailVedio` | String | 是 | 详情视频 | https://cloud.video.taobao.com/play/u/1685/p/1/e/6/t/1/5224**.mp4 |
| `internationalTradeInfo` | [message:alibaba.product.ProductInternationalTradeInfo](#m-alibaba-product-productinternationaltradeinfo) | 是 | 商品国际贸易信息，1688无需处理此字段 |   |
| `bizGroupInfos` | [message:alibaba.product.ProductBizGroupInfo[]](#m-alibaba-product-productbizgroupinfo[]) | 是 | 产品业务的支持信息,support为false说明不支持. | [] |
| `sellerLoginId` | String | 是 | 卖家旺旺ID | 卖家旺旺ID |
| `intelligentInfo` | [message:com.alibaba.ocean.openplatform.biz.product.common.model.ProductIntelligentInfo](#m-com-alibaba-ocean-openplatform-biz-product-common-model-productintelligentinfo) | 是 | 商品算法智能改写信息，包含算法优化后的商品标题和图片信息，未改写的则直接返回原标题和原图片 | [] |
| `sellStartTime` | Date | 是 | 预售开售时间 | 20181219175505000+0800 |
| `sevenDaysRefunds` | Boolean | 是 | 是否支持七天无理由退货 | true |

<a id="m-alibaba-product-productattribute[]"></a>
#### alibaba.product.ProductAttribute[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `attributeID` | Long | 是 | 属性ID | 123456 |
| `attributeName` | String | 是 | 属性名称 | color |
| `valueID` | Long | 是 | 属性值ID | 123456 |
| `value` | String | 是 | 属性值 | grey |
| `isCustom` | Boolean | 是 | 是否为自定义属性，国际站无需关注 | true |

<a id="m-alibaba-product-productimageinfo"></a>
#### alibaba.product.ProductImageInfo

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `images` | String[] | 是 | 主图列表，使用相对路径，需要增加域名：https://cbu01.alicdn.com/ | ["img/ibank/2014/766/624/1652426667_642119312.jpg","img/ibank/2014/656/624/1652426656_642119312.jpg","img/ibank/2014/236/624/1652426632_642119312.jpg"] |
| `isWatermark` | Boolean | 是 | 是否打水印，是(true)或否(false)，1688无需关注此字段，1688的水印信息在上传图片时处理 |  |
| `isWatermarkFrame` | Boolean | 是 | 水印是否有边框，有边框(true)或者无边框(false)，1688无需关注此字段，1688的水印信息在上传图片时处理 |  |
| `watermarkPosition` | String | 是 | 水印位置，在中间(center)或者在底部(bottom)，1688无需关注此字段，1688的水印信息在上传图片时处理 |  |

<a id="m-alibaba-product-productskuinfo[]"></a>
#### alibaba.product.ProductSKUInfo[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `attributes` | [message:alibaba.product.SKUAttrInfo[]](#m-alibaba-product-skuattrinfo[]) | 是 | SKU属性值，可填多组信息 | [] |
| `cargoNumber` | String | 是 | 指定规格的货号，国际站无需关注 | cy小黄人 |
| `amountOnSale` | Integer | 是 | 可销售数量，国际站无需关注 | 1200 |
| `retailPrice` | double | 是 | 建议零售价，国际站无需关注 | 40 |
| `price` | Double | 是 | 报价时该规格的单价，国际站注意要点：含有SKU属性的在线批发产品设定具体价格时使用此值，若设置阶梯价格则使用priceRange | 40 |
| `priceRange` | [message:alibaba.product.ProductPriceRange[]](#m-alibaba-product-productpricerange[]) | 是 | 阶梯报价，1688无需关注 | [] |
| `skuCode` | String | 是 | 商品编码，1688无需关注 | 3746778972330 |
| `skuId` | Long | 是 | skuId, 国际站无需关注 | 3746778972330 |
| `specId` | String | 是 | specId, 国际站无需关注 | fb82997a5b64af3c729cea89bef44409 |
| `consignPrice` | Double | 是 | 分销基准价 |  0.2 |
| `pftPrice` | Double | 是 | 批发团价格 | 0.1 |
| `jxhyPrice` | Double | 是 | 精选货源价一件包邮价 | 5.9 |
| `jxhyPfPrice` | Double | 是 | 精选货源价批发价 | 5.5 |

<a id="m-alibaba-product-skuattrinfo[]"></a>
#### alibaba.product.SKUAttrInfo[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `attributeID` | Long | 是 | sku属性ID | 3216 |
| `attValueID` | Long | 是 | sku值ID，1688不用关注 |   |
| `attributeValue` | String | 是 | sku值内容，国际站不用关注 | 黑色 |
| `customValueName` | String | 是 | 自定义属性值名称，1688无需关注 |   |
| `skuImageUrl` | String | 是 | sku图片 | img/ibank/2018/221/909/9143909122_1606139362.jpg |
| `attributeDisplayName` | String | 是 | sku属性ID所对应的显示名，比如颜色，尺码 |   |

<a id="m-alibaba-product-productpricerange[]"></a>
#### alibaba.product.ProductPriceRange[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `startQuantity` | Integer | 是 | 起批量 | 1 |
| `price` | Double | 是 | 商品价格 | 40 |

<a id="m-alibaba-product-productsaleinfo"></a>
#### alibaba.product.ProductSaleInfo

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `supportOnlineTrade` | Boolean | 是 | 是否支持网上交易。true：支持 false：不支持，国际站不需关注此字段 | TRUE |
| `mixWholeSale` | Boolean | 是 | 是否支持混批，国际站无需关注此字段 | FALSE |
| `saleType` | String | 是 | 销售方式，按件卖(normal)或者按批卖(batch)，1688站点无需关注此字段 | batch |
| `priceAuth` | Boolean | 是 | 是否价格私密信息，国际站无需关注此字段 | FALSE |
| `priceRanges` | [message:alibaba.product.ProductPriceRange[]](#m-alibaba-product-productpricerange[]) | 是 | 区间价格。按数量范围设定的区间价格 | [] |
| `amountOnSale` | Double | 是 | 可售数量，国际站无需关注此字段 | 24000 |
| `unit` | String | 是 | 计量单位 | 件 |
| `minOrderQuantity` | Integer | 是 | 最小起订量，范围是1-99999。 | 1 |
| `batchNumber` | Integer | 是 | 每批数量，默认为空或者非零值，该属性不为空时sellunit为必填 | 1 |
| `retailprice` | Double | 是 | 建议零售价，国际站无需关注 | 40 |
| `tax` | String | 是 | 税率相关信息，内容由用户自定，国际站无需关注 |   |
| `sellunit` | String | 是 | 售卖单位，如果为批量售卖，代表售卖的单位，该属性不为空时batchNumber为必填，例如1&quot;手&quot;=12“件&quot;的&quot;手&quot;，国际站无需关注 | 1手等于1件 |
| `quoteType` | Integer | 是 | 普通报价-FIXED_PRICE(&quot;0&quot;),SKU规格报价-SKU_PRICE(&quot;1&quot;),SKU区间报价（商品维度）-SKU_PRICE_RANGE_FOR_OFFER(&quot;2&quot;),SKU区间报价（SKU维度）-SKU_PRICE_RANGE(&quot;3&quot;)，国际站无需关注 | 1 |
| `consignPrice` | Double | 是 | 分销基准价 |  0.5 |
| `pftPrice` | Double | 是 | 批发团商品价 | 0.1 |
| `jxhyPrice` | Double | 是 | 精选货源一件包邮价 | 5.0 |
| `jxhyPfPrice` | Double | 是 | 精选货源价批发价 | 5.8 |
| `beginAmount` | Integer | 是 | 批发起批量 | 2 |
| `retailBeginAmount` | Long | 是 | 零售起售量 | 1 |

<a id="m-alibaba-product-productshippinginfo"></a>
#### alibaba.product.ProductShippingInfo

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `unitWeight` | Double | 是 | 重量/毛重，单位千克/件 | 121 |
| `volume` | Integer | 是 | 体积，单位是立方厘米，范围是1-9999999，1688无需关注此字段 | 500 |
| `handlingTime` | Integer | 是 | 备货期，单位是天，范围是1-60。1688无需处理此字段 | 12 |
| `freightTemplateID` | Long | 是 | 运费模板ID，0表示运费说明，1表示卖家承担运费，其他值表示使用运费模版。此参数可调用运费模板相关API获取 | 121133 |
| `suttleWeight` | Double | 是 | 净重，单位千克/件 | 1001 |
| `sendGoodsAddressText` | java.lang.String | 是 | 发货地描述 | asda |
| `width` | Double | 是 | 宽度，单位厘米 | 30 |
| `height` | Double | 是 | 高度，单位厘米 | 20 |
| `length` | Double | 是 | 长度，单位厘米 | 10 |
| `packageSize` | String | 是 | 尺寸，单位是厘米，长宽高范围是1-9999999。1688无需关注此字段 | 10x20x50 |
| `sendGoodsAddressId` | Long | 是 | 发货地址ID，国际站无需处理此字段 | 124431 |

<a id="m-alibaba-product-productextendinfo[]"></a>
#### alibaba.product.ProductExtendInfo[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `key` | String | 是 | 扩展结构的key | 代销价格,consignPrice;<br>买家保障,buyerProtection; |
| `value` | String | 是 | 扩展结构的value | 代销价格,key为skuId，value为用户设置的代销价，<br>示例：31151771910:2088.0;31151771909:2088.0;31151771908:2088.0;31152339121:2088.0;<br>买家保障,string数组，value为买保全拼，<br>示例：["psbj","swtwlybt","swtbh","ssbxsfh"] |

<a id="m-alibaba-product-productinternationaltradeinfo"></a>
#### alibaba.product.ProductInternationalTradeInfo

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `fobCurrency` | String | 是 | FOB价格货币，参见FAQ 货币枚举值 |  |
| `fobMinPrice` | String | 是 | FOB最小价格 |  |
| `fobMaxPrice` | String | 是 | FOB最大价格 |  |
| `fobUnitType` | String | 是 | FOB计量单位，参见FAQ 计量单位枚举值 |  |
| `paymentMethods` | String[] | 是 | 付款方式，参见FAQ 付款方式枚举值 |  |
| `minOrderQuantity` | Integer | 是 | 最小起订量 |  |
| `minOrderUnitType` | String | 是 | 最小起订量计量单位，参见FAQ 计量单位枚举值 |  |
| `supplyQuantity` | Integer | 是 | supplyQuantity |  |
| `supplyUnitType` | String | 是 | 供货能力计量单位，参见FAQ 计量单位枚举值 |  |
| `supplyPeriodType` | String | 是 | 供货能力周期，参见FAQ 时间周期枚举值 |  |
| `deliveryPort` | String | 是 | 发货港口 |  |
| `deliveryTime` | String | 是 | 发货期限 |  |
| `consignmentDate` | Integer | 是 | 新发货期限 |  |
| `packagingDesc` | String | 是 | 常规包装 |  |

<a id="m-com-alibaba-ocean-openplatform-biz-product-common-model-productintelligentinfo"></a>
#### com.alibaba.ocean.openplatform.biz.product.common.model.ProductIntelligentInfo

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `title` | String | 是 | 算法优化后的商品标题 | 算法优化后的商品标题 |
| `images` | String[] | 是 | 算法优化后的商品图片 | ["https://cbu01.alicdn.com/img/ibank/2020/932/210/13529012239_321095253.jpg"] |
| `skuImages` | [message:com.alibaba.ocean.openplatform.biz.product.common.model.SkuIntelligentInfo[]](#m-com-alibaba-ocean-openplatform-biz-product-common-model-skuintelligentinfo[]) | 是 | 算法优化后的规格图片 | [] |
| `descriptionImages` | String[] | 是 | 算法优化后的详情图片 | ["https://cbu01.alicdn.com/img/ibank/2020/932/210/13529012239_321095253.jpg"] |

<a id="m-com-alibaba-ocean-openplatform-biz-product-common-model-skuintelligentinfo[]"></a>
#### com.alibaba.ocean.openplatform.biz.product.common.model.SkuIntelligentInfo[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `skuId` | Long | 是 | skuId | 123433333 |
| `imageUrl` | String | 是 | 算法处理后的图片地址，未处理则返回原图片地址 | https://cbu01.alicdn.com/img/ibank/2020/932/210/13529012239_321095253.jpg |

<a id="m-alibaba-product-productbizgroupinfo[]"></a>
#### alibaba.product.ProductBizGroupInfo[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `support` | java.lang.Boolean | 是 | 是否支持 |  |
| `description` | java.lang.String | 是 | 垂直市场名字，如微供市场、货品市场 |  |
| `code` | java.lang.String | 是 | 垂直市场标记 |  |

## 错误码

| 错误码 | 现象 | 解决方案 |
|---|---|---|
| 没有该商品的查询权限 | 没有该商品的查询权限 | 没有建立铺货关系，在商品页面点击一键铺货或者调用铺货同步接口 |
| 商品[57374140142]不存在 | 商品[57374140142]不存在 | 商品ID错误，检查商品ID |

## 示例

**入参示例**

```
{
  productID:573741401425
}
```
