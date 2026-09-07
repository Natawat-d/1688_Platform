# 获取已购买过商家的商品简单信息

API: `com.alibaba.product:alibaba.product.simple.get:1` · Category: 商品  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.product:alibaba.product.simple.get-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.product/alibaba.product.simple.get/{appKey}`  
需要授权 (access_token) · 需要签名

由商品ID获取商品详细信息，接口能获取已购买过商家的商品简单信息。该接口需要付费才能访问。该接口只返回简单信息，主要是用来在ERP系统中做数据关联

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `productID` | Long | 是 | 商品ID | 565507182121 |
| `webSite` | String | 是 | 站点信息，指定调用的API是属于国际站（alibaba）还是1688网站（1688） | 1688 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `productInfo` | [message:alibaba.product.ProductInfo](#m-alibaba-product-productinfo) | 是 | 商品详细信息 | {} |
| `bizGroupInfos` | [message:alibaba.product.ProductBizGroupInfo[]](#m-alibaba-product-productbizgroupinfo[]) | 是 | 产品业务的支持信息,support为false说明不支持. | [] |
| `errMsg` | String | 是 | 返回错误信息 | 商品[57053081292]不存在 |

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
| `qualityLevel` | Integer | 是 | 质量星级(1-7) | 5 |
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
| `internationalTradeInfo` | [message:alibaba.product.ProductInternationalTradeInfo](#m-alibaba-product-productinternationaltradeinfo) | 是 | 商品国际贸易信息，1688无需处理此字段 |  |
| `bizGroupInfos` | [message:alibaba.product.ProductBizGroupInfo[]](#m-alibaba-product-productbizgroupinfo[]) | 是 | 产品业务的支持信息,support为false说明不支持. | [] |
| `sellerLoginId` | String | 是 | 卖家旺旺ID | 卖家旺旺ID |
| `intelligentInfo` | [message:com.alibaba.ocean.openplatform.biz.product.common.model.ProductIntelligentInfo](#m-com-alibaba-ocean-openplatform-biz-product-common-model-productintelligentinfo) | 是 | 商品算法智能改写信息，包含算法优化后的商品标题和图片信息，未改写的则直接返回原标题和原图片 | [] |
| `processing` | [message:alibaba.product.ProductProcessing](#m-alibaba-product-productprocessing) | 否 | 加工定制信息，传入此参数表示是加工定制产品 | 默认 |
| `productProcessing` | [message:alibaba.product.ProductProcessing](#m-alibaba-product-productprocessing) | 是 | 加工定制信息 | 默认 |
| `attributeChanges` | String[] | 是 | 商品修改的具体属性 | [image] |
| `productLineId` | Long | 否 | 产品线 | -1 |
| `processingOfferId` | Long | 是 | 现货商品关联的定制商品id | 6388 |
| `sevenDaysRefunds` | Boolean | 是 | 是否支持七天无理由退货 | true |
| `privateChannelInfo` | [message:com.alibaba.forGetcommon.model.PrivateChannelInfo](#m-com-alibaba-forgetcommon-model-privatechannelinfo) | 是 | 商品私密价格信息 |   |
| `reserveInfo` | [message:alibaba.product.ReserveInfo](#m-alibaba-product-reserveinfo) | 是 | 订货数据 | {} |
| `sellStartTime` | Date | 是 | 预售开售时间 | 20181219175505000+0800 |
| `productOfficialLogisticsModel` | [message:com.alibaba.ocean.openplatform.biz.product.common.model.ProductOfficialLogisticsModel](#m-com-alibaba-ocean-openplatform-biz-product-common-model-productofficiallogisticsmodel) | 是 | 官方物流件重尺模型 | {} |

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
| `consignPrice` | Double | 是 | 分销基准价 |   |
| `takeSamplePrice` | Double | 是 | SKU 维度拿样价, 国际站无需关注 | 1.00 |
| `pftPrice` | Double | 是 | 批发团价格 | 0.1 |
| `jxhyPrice` | Double | 是 | 精选货源价 | 0.1 |
| `jxhyPfPrice` | Double | 是 | 精选货源价批发价 | 5.5 |
| `dailySpecialPrice` | Double | 是 | 天天特卖价 | 5.4 |
| `eleXdChannelPrice` | Double | 是 | 饿了么小店专属价 | 5.4 |

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
| `attrType` | String | 是 | 属性类型 | 1 规格属性 2规格扩展属性 ，默认规格属性 |
| `attributeName` | String | 是 | sku属性ID所对应的显示名，比如颜色，尺码 |    |

<a id="m-alibaba-product-productpricerange[]"></a>
#### alibaba.product.ProductPriceRange[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `startQuantity` | Integer | 是 | 起批量 | 3 |
| `price` | Double | 是 | 价格 | 445 |

<a id="m-alibaba-product-productsaleinfo"></a>
#### alibaba.product.ProductSaleInfo

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `supportOnlineTrade` | Boolean | 是 | 是否支持网上交易。true：支持 false：不支持，国际站不需关注此字段 | TRUE |
| `mixWholeSale` | Boolean | 是 | 是否支持混批，国际站无需关注此字段 | TRUE |
| `saleType` | String | 是 | 销售方式，按件卖(normal)或者按批卖(batch)，1688站点无需关注此字段 | normal |
| `priceAuth` | Boolean | 是 | 是否价格私密信息，国际站无需关注此字段 | TRUE |
| `priceRanges` | [message:alibaba.product.ProductPriceRange[]](#m-alibaba-product-productpricerange[]) | 是 | 区间价格。按数量范围设定的区间价格 | [] |
| `amountOnSale` | Double | 是 | 可售数量，国际站无需关注此字段 | 108 |
| `unit` | String | 是 | 计量单位 | 件 |
| `minOrderQuantity` | Integer | 是 | 最小起订量，范围是1-99999。 | 3 |
| `batchNumber` | Integer | 是 | 每批数量，默认为空或者非零值，该属性不为空时sellunit为必填 | 1 |
| `retailprice` | Double | 是 | 建议零售价，国际站无需关注 | 20600 |
| `tax` | String | 是 | 税率相关信息，内容由用户自定，国际站无需关注 |   |
| `sellunit` | String | 是 | 售卖单位，如果为批量售卖，代表售卖的单位，该属性不为空时batchNumber为必填，例如1&quot;手&quot;=12“件&quot;的&quot;手&quot;，国际站无需关注 | 1手等于1件 |
| `quoteType` | Integer | 是 | 普通报价-FIXED_PRICE(&quot;0&quot;),SKU规格报价-SKU_PRICE(&quot;1&quot;),SKU区间报价（商品维度）-SKU_PRICE_RANGE_FOR_OFFER(&quot;2&quot;),SKU区间报价（SKU维度）-SKU_PRICE_RANGE(&quot;3&quot;)，国际站无需关注 | 1 |
| `consignPrice` | Double | 是 | 分销基准价。代销场景均使用该价格。有SKU商品查看skuInfo中的consignPrice | 20600 |
| `deliveryLimit` | Integer | 是 | 发货时间限制（非买保发货周期），按开计算 | 5 |
| `invReduceType` | String | 是 | 库存扣减方式，1是下单减库存，2是付款减库存 | 1 |
| `pftPrice` | Double | 是 | 批发团商品价 | 0.1 |
| `jxhyPrice` | Double | 是 | 精选货源一件包邮价 | 5.0 |
| `jxhyPfPrice` | Double | 是 | 精选货源价批发价 | 5.8 |
| `dailySpecialPrice` | Double | 是 | 天天特卖价 | 5.5 |
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
| `offerSuttleWeight` | Double | 是 | OFFER净重（kg） | 2 |
| `offerWidth` | Double | 是 | OFFER包装宽度（cm） | 30 |
| `offerHeight` | Double | 是 | OFFER包装高度（cm） | 30 |
| `offerLength` | Double | 是 | OFFER包装长度（cm） | 30 |
| `freightTemplate` | [message:alibaba.product.FreightTemplate[]](#m-alibaba-product-freighttemplate[]) | 是 | 商品运费费率 | [] |
| `channelPriceFreePostage` | Boolean | 是 | 厂货通渠道专享价是否包邮，要结合非包邮地址，如果收货地址在非包邮地区则商品为不包邮 | true |
| `channelPriceExcludeAreaCodes` | [message:com.alibaba.ocean.openplatform.biz.trade.result.ProductAddressCode[]](#m-com-alibaba-ocean-openplatform-biz-trade-result-productaddresscode[]) | 是 | 厂货通渠道专享价非包邮地区（地址信息列表，省份信息） | [{}] |

<a id="m-alibaba-product-freighttemplate[]"></a>
#### alibaba.product.FreightTemplate[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `addressCodeText` | String | 是 | 地址区域编码对应的文本（包括省市区，用空格隔开） | 福建省 福州市 鼓楼区 |
| `fromAreaCode` | String | 是 | 发货地址地区码 | 350102 |
| `id` | Long | 是 | 地址ID | 1234 |
| `name` | String | 是 | 模板名称 | 2019 |
| `remark` | String | 是 | 备注 | 2019 |
| `status` | Integer | 是 | 状态：1表示有效，-1表示失效 | 1 |
| `expressSubTemplate` | [message:alibaba.product.DeliverySubTemplateDetailDTO](#m-alibaba-product-deliverysubtemplatedetaildto) | 是 | 快递子模版 | {} |
| `logisticsSubTemplate` | [message:alibaba.product.DeliverySubTemplateDetailDTO](#m-alibaba-product-deliverysubtemplatedetaildto) | 是 | 货运子模版 | {} |
| `codSubTemplate` | [message:alibaba.product.DeliverySubTemplateDetailDTO](#m-alibaba-product-deliverysubtemplatedetaildto) | 是 | 货到付款子模版 | {} |

<a id="m-alibaba-product-deliverysubtemplatedetaildto"></a>
#### alibaba.product.DeliverySubTemplateDetailDTO

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `subTemplateDTO` | [message:alibaba.product.DeliverySubTemplateDTO](#m-alibaba-product-deliverysubtemplatedto) | 是 | 子模板 | {} |
| `rateList` | [message:alibaba.product.DeliveryRateDetailDTO[]](#m-alibaba-product-deliveryratedetaildto[]) | 是 | 费率 | [] |

<a id="m-alibaba-product-deliverysubtemplatedto"></a>
#### alibaba.product.DeliverySubTemplateDTO

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `chargeType` | Integer | 是 | 计件类型。0:重量 1:件数 2:体积 | 0 |
| `isSysTemplate` | Boolean | 是 | 是否系统模板 | false |
| `serviceChargeType` | Integer | 是 | 运费承担类型 卖家承担：0；买家承担：1。 | 1 |
| `serviceType` | Integer | 是 | 服务类型。0:快递 1:货运 2:货到付款 | 0 |
| `type` | Integer | 是 | 子模板类型 0基准 1增值。默认0。 | 0 |

<a id="m-alibaba-product-deliveryratedetaildto[]"></a>
#### alibaba.product.DeliveryRateDetailDTO[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `isSysRate` | Boolean | 是 | 是否系统模板 | false |
| `toAreaCodeText` | String | 是 | 地址编码文本，用顿号隔开。例如：上海、福建省、广东省 | 上海、福建省、广东省 |
| `rateDTO` | [message:alibaba.product.DeliveryRateDTO](#m-alibaba-product-deliveryratedto) | 是 | 普通子模板费率 | {} |
| `sysRateDTO` | [message:alibaba.product.DeliverySysRateDTO](#m-alibaba-product-deliverysysratedto) | 是 | 系统子模板费率 | {} |

<a id="m-alibaba-product-deliveryratedto"></a>
#### alibaba.product.DeliveryRateDTO

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `firstUnit` | Long | 是 | 首重（单位：克）或首件（单位：件） | 1 |
| `firstUnitFee` | Long | 是 | 首重或首件的价格（单位：分） | 600 |
| `leastExpenses` | Long | 是 | 最低一票 | 1 |
| `nextUnit` | Long | 是 | 续重件单位 | 2 |
| `nextUnitFee` | Long | 是 | 续重件价格（单位：分） | 100 |

<a id="m-alibaba-product-deliverysysratedto"></a>
#### alibaba.product.DeliverySysRateDTO

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `firstUnit` | Long | 是 | 首重（单位：克）或首件（单位：件） | 1 |
| `firstUnitFee` | Long | 是 | 首重或首件的价格（单位：分） | 600 |
| `leastExpenses` | Long | 是 | 最低一票 | 1 |
| `nextUnit` | Long | 是 | 续重（单位：克）或续件（单位：件）单位 | 1 |
| `nextUnitFee` | Long | 是 | 续重件价格（单位：分） | 100 |

<a id="m-com-alibaba-ocean-openplatform-biz-trade-result-productaddresscode[]"></a>
#### com.alibaba.ocean.openplatform.biz.trade.result.ProductAddressCode[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `code` | String | 是 | 地址编码 | 620000 |
| `name` | String | 是 | 地址名称 | 甘肃省 |

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

<a id="m-alibaba-product-productprocessing"></a>
#### alibaba.product.ProductProcessing

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `images` | String[] | 是 | 定制商品样图 | ["https://1.jpg","https://1.jpg"] |
| `customType` | String[] | 是 | 定制方式，可选值: processing、sku | ["processing"] |
| `minPrice` | Double | 是 | 价格区间的最低价 | 10.00 |
| `maxPrice` | Double | 是 | 价格区间的最高价 | 100.00 |
| `minOrderQuantity` | Long | 是 | 最小起订量 | 50 |
| `stockOfferIds` | Long[] | 是 | 关联现货商品的id列表 | [6386,6387] |
| `reserveRanges` | [message:alibaba.product.ReserveRange[]](#m-alibaba-product-reserverange[]) | 是 | 出货周期 | [{"quantity":100,"period":10}] |
| `sampleInfo` | [message:alibaba.product.SampleInfo](#m-alibaba-product-sampleinfo) | 是 | 打样信息 | 	{"isSupportSample":true,"price":10,"period":2} |

<a id="m-alibaba-product-reserverange[]"></a>
#### alibaba.product.ReserveRange[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `beginAmount` | Long | 是 | 订货数量 | 100 |
| `date` | Integer | 是 | 预计出货时间（天） | 10 |
| `endAmount` | Long | 是 | 订货截止数量 | 100 |
| `quantity` | Long | 是 | 订货起订数量 | 1000 |
| `period` | Integer | 是 | 预计出货时间（天） | 10 |

<a id="m-alibaba-product-sampleinfo"></a>
#### alibaba.product.SampleInfo

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `price` | Double | 是 | 打样价格（单位：元） | 10.00 |
| `period` | Integer | 是 | 打样周期（单位：天） | 2 |
| `isSupportSample` | Boolean | 是 | 是否支持打样 | true |

<a id="m-com-alibaba-forgetcommon-model-privatechannelinfo"></a>
#### com.alibaba.forGetcommon.model.PrivateChannelInfo

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `offerPrivatePriceInfo` | [message:com.alibaba.forGet.common.model.OfferPrivatePriceInfo[]](#m-com-alibaba-forget-common-model-offerprivatepriceinfo[]) | 是 | 私密商品私密价格信息 | [] |

<a id="m-com-alibaba-forget-common-model-offerprivatepriceinfo[]"></a>
#### com.alibaba.forGet.common.model.OfferPrivatePriceInfo[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `skuId` | Long | 是 | skuid，非规格报价为0 | 0 |
| `privatePriceInfo` | String | 是 | 具体会员等级价格（普通，高级，vip，至尊vip，初级）注意会员顺序 | [] |

<a id="m-alibaba-product-reserveinfo"></a>
#### alibaba.product.ReserveInfo

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `supportReserve` | Boolean | 是 | 是否支持订货 | false |
| `minQuantity` | Long | 是 | 最小订货量 | 10 |
| `maxQuantity` | Long | 是 | 最大订货量 | 1000 |
| `reserveRangeInfos` | [message:alibaba.product.ReserveRangeInfo[]](#m-alibaba-product-reserverangeinfo[]) | 是 | 订货区间信息 | [] |

<a id="m-alibaba-product-reserverangeinfo[]"></a>
#### alibaba.product.ReserveRangeInfo[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `quantity` | Long | 是 | 订货数量 | 1000 |
| `period` | Integer | 是 | 出货时间(天) | 10 |

<a id="m-com-alibaba-ocean-openplatform-biz-product-common-model-productofficiallogisticsmodel"></a>
#### com.alibaba.ocean.openplatform.biz.product.common.model.ProductOfficialLogisticsModel

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `officialLogisticsInfoModel` | [message:com.alibaba.ocean.openplatform.biz.product.common.model.ProductOfficialLogisticsInfoModel](#m-com-alibaba-ocean-openplatform-biz-product-common-model-productofficiallogisticsinfomodel) | 是 | 无sku时，官方商品物流模型，长宽高等 | {} |
| `officialLogisticsSkuInfoModels` | [message:com.alibaba.ocean.openplatform.biz.product.common.model.ProductOfficialLogisticsSkuInfoModel[]](#m-com-alibaba-ocean-openplatform-biz-product-common-model-productofficiallogisticsskuinfomodel[]) | 是 | 有sku时，官方商品的sku的物流模型，长宽高等 | {} |

<a id="m-com-alibaba-ocean-openplatform-biz-product-common-model-productofficiallogisticsinfomodel"></a>
#### com.alibaba.ocean.openplatform.biz.product.common.model.ProductOfficialLogisticsInfoModel

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `length` | Double | 是 | 长 | 1 |
| `width` | Double | 是 | 宽 | 1 |
| `height` | Double | 是 | 高 | 1 |
| `weight` | Long | 是 | 重量 | 100 |
| `volume` | Double | 是 | 体积 | 1 |

<a id="m-com-alibaba-ocean-openplatform-biz-product-common-model-productofficiallogisticsskuinfomodel[]"></a>
#### com.alibaba.ocean.openplatform.biz.product.common.model.ProductOfficialLogisticsSkuInfoModel[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `skuId` | Long | 是 | skuid | 121315446134 |
| `specId` | String | 是 | specid | 1o9KdfOwelir12934kDKFuikA87LIk |
| `length` | Double | 是 | 长 | 1 |
| `width` | Double | 是 | 宽 | 1 |
| `height` | Double | 是 | 高 | 1 |
| `weight` | Long | 是 | 重量 | 100 |
| `volume` | Double | 是 | 体积 | 1 |

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
| 商品[56550718212]不存在 | 商品[56550718212]不存在 | 检查商品ID是否正确 |

## 示例

**出参示例**

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
