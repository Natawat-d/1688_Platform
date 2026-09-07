# 创建订单前预览数据接口

API: `com.alibaba.trade:alibaba.createOrder.preview:1` · Category: 订单  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.createOrder.preview-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.createOrder.preview/{appKey}`  
需要授权 (access_token) · 需要签名

订单创建只允许购买同一个供应商的商品。本接口返回创建订单相关的优惠等信息。
1、校验商品数据是否允许订购。
2、校验代销关系
3、校验库存、起批量、是否满足混批条件

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `addressParam` | [message:alibaba.trade.fast.address](#m-alibaba-trade-fast-address) | 是 | 收货地址信息 | {"address":"网商路699号","phone": "0517-88990077","mobile": "15251667788","fullName": "张三","postCode": "000000","areaText": "滨江区","townText": "","cityText": "杭州市","provinceText": "浙江省"} |
| `cargoParamList` | [message:alibaba.trade.fast.cargo[]](#m-alibaba-trade-fast-cargo[]) | 是 | 商品信息 | [{"specId": "b266e0726506185beaf205cbae88530d","quantity": 5,"offerId": 554456348334},{"specId": "2ba3d63866a71fbae83909d9b4814f01","quantity": 6,"offerId": 554456348334}] |
| `invoiceParam` | [message:alibaba.trade.fast.invoice](#m-alibaba-trade-fast-invoice) | 否 | 发票信息 | {"invoiceType":0,"cityText": "杭州市","provinceText": "浙江省","address": "网商路699号","phone": "0517-88990077","mobile": "15251667788","fullName": "张五","postCode": "000000","areaText": "滨江区","companyName": "测试公司","taxpayerIdentifier": "123455"} |
| `flow` | String | 否 | general（创建大市场订单），fenxiao（创建分销订单）,paired(天天特卖),repurchase(复购合约下单),saleproxy流程将校验分销关系,boutiquefenxiao(精选货源分销价下单，采购量1个使用包邮)， boutiquepifa(精选货源批发价下单，采购量大于2使用), flow如果为空的情况，会比价择优预览，并返回最优下单方式flow。 复购合约下单不在比价范围内。所有的下单flow渠道，可能导致价格和买保服务差异。 | general |
| `instanceId` | String | 否 | 批发团instanceId,从alibaba.pifatuan.product.list获取 | 4063139_1662080400000 |
| `encryptOutOrderInfo` | [message:alibaba.trade.fastCreateOrder.EncryptOutOrderInfo](#m-alibaba-trade-fastcreateorder-encryptoutorderinfo) | 否 | 下游加密订单信息，用于下游打单使用 | {} |
| `proxySettleRecordId` | String | 否 | 分账普通下单采购单id，交易flow为“proxy” | 4051300002 |
| `inventoryMode` | String | 否 | 库存模式，jit（jit模式）或 cang（仓发模式）,目前只提供给AE使用 | jit |
| `outOrderId` | String | 否 | 外部订单号 | 988129883123 |
| `pickupService` | String | 否 | 上门揽收,目前AE供货可用，其他场景暂不开通 | y或n,默认为n |
| `crossBorderLogisticsSolutionId` | String | 否 | 用户选择的官方物流跨境解决方案sourceId | GLOBAL_CAINIAO_VN_TRANSIT_LAND  |
| `useBorderLogisticsSolution` | Boolean | 否 | 是否使用跨境物流解决方案预览，默认为false，当海外地址时，参数不生效。当为港澳台地址时，参数生效 | true |
| `isvBizType` | String | 否 | 开放平台业务码 | cross |

<a id="m-alibaba-trade-fast-address"></a>
#### alibaba.trade.fast.address

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `addressId` | Long | 是 | 收货地址id | 1234 |
| `fullName` | String | 是 | 收货人姓名 | 张三 |
| `mobile` | String | 是 | 手机 | 15251667788 |
| `phone` | String | 是 | 电话 | 0517-88990077 |
| `postCode` | String | 是 | 邮编 | 000000 |
| `cityText` | String | 是 | 市文本 | 杭州市 |
| `provinceText` | String | 是 | 省份文本 | 浙江省 |
| `areaText` | String | 是 | 区文本 | 滨江区 |
| `townText` | String | 是 | 镇文本 | 长河镇 |
| `address` | String | 是 | 街道地址 | 网商路699号 |
| `districtCode` | String | 是 | 地址编码 | 310107 |

<a id="m-alibaba-trade-fast-cargo[]"></a>
#### alibaba.trade.fast.cargo[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `offerId` | Long | 是 | 商品对应的offer id | 554456348334 |
| `specId` | String | 是 | 商品规格id | b266e0726506185beaf205cbae88530d |
| `quantity` | Double | 是 | 商品数量(计算金额用) | 5 |
| `openOfferId` | String | 否 | 加密offerId，当搜索返回只有openOfferId时，需采用openOfferId替代offerId下单 | Wcv2w970KoL1BCpJGQp3vwKvcBThOPlpHnUtkK3T0n4= |
| `outMemberId` | String | 否 | 外部下游会员ID | 98928912-23 |
| `bizFenXiaoOutSubOrderInfo` | [message:com.alibaba.ocean.openplatform.biz.trade.common.model.BizFenXiaoOutSubOrderInfo[]](#m-com-alibaba-ocean-openplatform-biz-trade-common-model-bizfenxiaooutsuborderinfo[]) | 否 | 分销下游子单数据 | {} |

<a id="m-com-alibaba-ocean-openplatform-biz-trade-common-model-bizfenxiaooutsuborderinfo[]"></a>
#### com.alibaba.ocean.openplatform.biz.trade.common.model.BizFenXiaoOutSubOrderInfo[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `outSubOrderId` | String | 否 | 外部订单id | 123 |
| `outSkuId` | String | 否 | 外部sku | 123 |
| `outSubGmv` | Double | 否 | gmv | 123 |
| `outSubQty` | Double | 否 | 订单量 | 123 |
| `outSubOrderPayTime` | String | 否 | 下游子单下单时间 | 2025-03-18 09:20:21 |
| `outSubOrderLatestDeliveryTime` | String | 否 | 下游主单最晚发货时间 | 2025-03-18 09:20:21 |
| `outSubGmvDiscount` | Double | 否 | 下游子订单优惠金额 | 10 |
| `outSubGmvPost` | Double | 否 | 下游子订单邮费 | 5 |
| `outSubGmvReceive` | Double | 否 | 下游子订单实收金额 | 118 |

<a id="m-alibaba-trade-fast-invoice"></a>
#### alibaba.trade.fast.invoice

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `invoiceType` | Integer | 是 | 发票类型<br>0：普通发票，1:增值税发票 | 0 |
| `provinceText` | String | 是 | 省份文本 | 浙江省 |
| `cityText` | String | 是 | 城市文本 | 杭州市 |
| `areaText` | String | 是 | 地区文本 | 滨江区 |
| `townText` | String | 是 | 镇文本 | 长河镇 |
| `postCode` | String | 是 | 邮编 | 333333 |
| `address` | String | 是 | 街道 | 网商路699号 |
| `fullName` | String | 是 | 收票人姓名 | 张三 |
| `phone` | String | 是 | 电话 | 0517-88990077 |
| `mobile` | String | 是 | 手机 | 15251667788 |
| `companyName` | String | 是 | 购货公司名（发票抬头） | 测试公司 |
| `taxpayerIdentifier` | String | 是 | 纳税识别码 | 12345 |
| `bankAndAccount` | String | 是 | 开户行及帐号 | 网商银行 |
| `localInvoiceId` | String | 是 | 增值税本地发票号 | 123123123 |

<a id="m-alibaba-trade-fastcreateorder-encryptoutorderinfo"></a>
#### alibaba.trade.fastCreateOrder.EncryptOutOrderInfo

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `encryptOrder` | Boolean | 是 | 是否加密订单 | true |
| `outPlatformOrderNo` | String | 是 | 下游平台订单号 | 12365452354551 |
| `outPlatformSupplyOrderNo` | String | 否 | 下游平台供应链采购订单 | CT7403712218870825259 |
| `outPlatformCode` | String | 是 | 淘宝-thyny，天猫-tm，淘特-taote，阿里巴巴C2M-c2m，京东-jingdong，拼多多-pinduoduo，微信小店-weixin，跨境-kuajing，快手-kuaishou，有赞-youzan，抖音-douyin，寺库-siku，美团团好货-meituan，小红书-xiaohongshu，当当-dangdang，苏宁-suning，大V店-davdian，行云-xingyun，蜜芽-miya，菠萝派商城-boluo，其他-other | taote |
| `outPlatformAppkey` | String | 是 | 下游平台获取订单的appkey | 32154 |
| `outShopId` | String | 否 | 下游平台店铺Id | 1879283 |
| `outShopName` | String | 否 | 下游平台店铺名称 | 三生科技 |
| `outOriginAddress` | [message:com.alibaba.ocean.openplatform.biz.trade.param.OutAddress](#m-com-alibaba-ocean-openplatform-biz-trade-param-outaddress) | 否 | 外部原始地址信息 | {} |
| `oaid` | String | 否 | 淘宝oaid | 265646-52342354-2354Akf-w3654SF |
| `outPatformExtraInfo` | String | 否 | 下游平台其他扩展信息 | {} |
| `encryptReceiverName` | String | 否 | 下游加密收货人姓名 | *** |
| `encryptReceiverMobile` | String | 否 | 下游加密收货人电话 | *** |
| `encryptReceiverAddress` | String | 否 | 下游加密收货人地址 | *** |
| `outPlatformSubCode` | String | 否 | 下游渠道子业务编码，比如抖音101子渠道，用于供应链订单识别 | 101 |

<a id="m-com-alibaba-ocean-openplatform-biz-trade-param-outaddress"></a>
#### com.alibaba.ocean.openplatform.biz.trade.param.OutAddress

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `province` | [message:com.alibaba.ocean.openplatform.biz.trade.param.Place](#m-com-alibaba-ocean-openplatform-biz-trade-param-place) | 是 | 省 | {"name":"四川省","code":"51000"} |
| `city` | [message:com.alibaba.ocean.openplatform.biz.trade.param.Place](#m-com-alibaba-ocean-openplatform-biz-trade-param-place) | 是 | 市 | {} |
| `area` | [message:com.alibaba.ocean.openplatform.biz.trade.param.Place](#m-com-alibaba-ocean-openplatform-biz-trade-param-place) | 是 | 区 | {} |
| `town` | [message:com.alibaba.ocean.openplatform.biz.trade.param.Place](#m-com-alibaba-ocean-openplatform-biz-trade-param-place) | 否 | 镇/街道 | {} |
| `address` | String | 否 | 详细地址 | 网商路699号 |
| `postCode` | String | 否 | 邮编 | 511304 |

<a id="m-com-alibaba-ocean-openplatform-biz-trade-param-place"></a>
#### com.alibaba.ocean.openplatform.biz.trade.param.Place

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `code` | String | 是 | 地址code | 511300 |
| `name` | String | 是 | 地址name | 南充 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `orderPreviewResuslt` | [message:alibaba.createOrder.preview.result.model[]](#m-alibaba-createorder-preview-result-model[]) | 是 | 订单预览结果，过自动拆单会返回多个记录 | [] |
| `success` | Boolean | 是 | 是否成功 | true |
| `errorCode` | String | 是 | 错误码 | 500_1 |
| `errorMsg` | String | 是 | 错误信息 | 错误 |
| `postFeeByDescOfferList` | Long[] | 是 | 运费说明的商品列表 | [12324324234,12312422] |
| `consignOfferList` | Long[] | 是 | 代销商品列表 | [12324324234,12312422] |
| `unsupportedCrossBorderPayOfferList` | Long[] | 是 | 不支持跨境宝支付的商品列表 | [12324324234,12312422] |
| `extPairList` | [message:com.alibaba.ocean.openplatform.biz.trade.common.model.ExtrPair[]](#m-com-alibaba-ocean-openplatform-biz-trade-common-model-extrpair[]) | 是 | 扩展对象列表 | [] |

<a id="m-alibaba-createorder-preview-result-model[]"></a>
#### alibaba.createOrder.preview.result.model[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `discountFee` | java.lang.Long | 是 | 计算完货品金额后再次进行的减免金额. 单位: 分 | ''  |
| `tradeModeNameList` | String[] | 是 | 当前交易在使用下单接口时可以支持的交易方式列表，其中的元素可以直接用于下单接口的tradeType入参。列表为空，当前交易不可通过接口下单，需要在1688页面下单。 |  ''  |
| `status` | boolean | 是 | 状态 |  ''  |
| `taoSampleSinglePromotion` | boolean | 是 | 是否有淘货源单品优惠  false:有单品优惠   true：没有单品优惠 |  ''  |
| `sumPayment` | long | 是 | 订单总费用, 单位为分. |  ''  |
| `message` | java.lang.String | 是 | 返回信息 |  ''  |
| `sumCarriage` | long | 是 | 总运费信息, 单位为分. |  ''  |
| `resultCode` | java.lang.String | 是 | 返回码 |  ''  |
| `sumPaymentNoCarriage` | long | 是 | 不包含运费的货品总费用, 单位为分. |  ''  |
| `additionalFee` | java.lang.Long | 是 | 附加费,单位，分 | ''   |
| `flowFlag` | java.lang.String | 是 | 订单下单流程 |  ''  |
| `cargoList` | [message:alibaba.createOrder.preview.resultCargo.model[]](#m-alibaba-createorder-preview-resultcargo-model[]) | 是 | 规格信息 |  ''  |
| `shopPromotionList` | [message:alibaba.trade.promotion.model[]](#m-alibaba-trade-promotion-model[]) | 是 | 可用店铺级别优惠列表 |  ''  |
| `tradeModelList` | [message:tradeModelExtensionList[]](#m-trademodelextensionlist[]) | 是 | 当前交易可以支持的交易方式列表。结果可以参照1688下单预览页面的交易方式。 |  ''  |
| `payChannelInfos` | [message:payChaneelList[]](#m-paychaneellist[]) | 是 | 当前交易支持的支付渠道信息 | [] |
| `tradeServiceList` | [message:com.alibaba.ocean.openplatform.biz.trade.common.model.TradeServiceGroup[]](#m-com-alibaba-ocean-openplatform-biz-trade-common-model-tradeservicegroup[]) | 是 | 提供的服务列表，比如跨境物流服务等 | {} |
| `extPairList` | [message:com.alibaba.ocean.openplatform.biz.trade.common.model.ExtrPair[]](#m-com-alibaba-ocean-openplatform-biz-trade-common-model-extrpair[]) | 是 | 扩展对象列表 | [] |
| `orderGroup` | String | 是 | 订单分组group，指定官方物流提货方案时，需要按该字段分组指定 | 123 |
| `canUseOfficialSolution` | Boolean | 是 | 是否可用官方物流提货 | true |
| `officialSolutionModelList` | [message:com.alibaba.ocean.openplatform.biz.trade.common.model.OfficialSolutionModel[]](#m-com-alibaba-ocean-openplatform-biz-trade-common-model-officialsolutionmodel[]) | 是 | 官方物流提货服务列表 | [{"solutionCode":"471","solutionName":"特惠当日上门","totalCost":770}] |
| `totalFundUsageAmount` | Long | 是 | 红包使用总金额，分 | 10 |
| `totalPostFundUsageAmount` | Long | 是 | 运费券使用总金额，分 | 100 |

<a id="m-alibaba-createorder-preview-resultcargo-model[]"></a>
#### alibaba.createOrder.preview.resultCargo.model[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `amount` | java.lang.Double | 是 | 产品总金额 | 10 |
| `message` | java.lang.String | 是 | 返回信息 | null |
| `finalUnitPrice` | java.lang.Double | 是 | 最终单价 | 100 |
| `specId` | java.lang.String | 是 | 规格ID，offer内唯一 | 98SDKTKSLFasbFiAlf |
| `skuId` | java.lang.Long | 是 | 规格ID，全局唯一 | 245252612367 |
| `resultCode` | java.lang.String | 是 | 返回码 | 200 |
| `offerId` | java.lang.Long | 是 | 商品ID | 909872988917 |
| `openOfferId` | String | 是 | 加密商品ID | 9JlsKsfcKLAKDFS |
| `cargoPromotionList` | [message:alibaba.trade.promotion.model[]](#m-alibaba-trade-promotion-model[]) | 是 | 商品优惠列表 | {} |
| `extPairList` | [message:com.alibaba.ocean.openplatform.biz.trade.common.model.ExtrPair[]](#m-com-alibaba-ocean-openplatform-biz-trade-common-model-extrpair[]) | 是 | 扩展对象列表 | [] |
| `totalFundUsageAmount` | Long | 是 | 红包使用总金额 分 | 10 |
| `totalPostFundUsageAmount` | Long | 是 | 运费券使用总金额 分 | 10 |
| `tradeServiceList` | [message:com.alibaba.ocean.openplatform.biz.trade.common.model.TradeServiceGroup[]](#m-com-alibaba-ocean-openplatform-biz-trade-common-model-tradeservicegroup[]) | 是 | 服务-跟商品相关的服务 | {} |

<a id="m-alibaba-trade-promotion-model[]"></a>
#### alibaba.trade.promotion.model[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `promotionId` | java.lang.String | 是 | 优惠券ID |   |
| `selected` | boolean | 是 | 是否默认选中 |   |
| `text` | java.lang.String | 是 | 优惠券名称 |   |
| `desc` | java.lang.String | 是 | 优惠券描述 |   |
| `freePostage` | boolean | 是 | 是否免邮 |   |
| `discountFee` | java.lang.Long | 是 | 减去金额，单位为分 |   |

<a id="m-trademodelextensionlist[]"></a>
#### tradeModelExtensionList[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `name` | String | 是 | 交易方式名称，1688下单预览页面展示的名称 |   |
| `description` | String | 是 | 交易描述 |   |
| `tradeType` | String | 是 | 做为入参传入下单接口的tradeType字段 |   |
| `opSupport` | Boolean | 是 | 开放平台下单是否支持此种交易模式。如果为true,该交易方式可做为下单接口tradeType参数的入参；如果为false,则不可做为下单接口的入参。 |   |

<a id="m-paychaneellist[]"></a>
#### payChaneelList[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `name` | String | 是 | 支付渠道 | "alipay" |
| `amountLimit` | Long | 是 | 可用额度金额，单位为分 | 100 |

<a id="m-com-alibaba-ocean-openplatform-biz-trade-common-model-tradeservicegroup[]"></a>
#### com.alibaba.ocean.openplatform.biz.trade.common.model.TradeServiceGroup[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `outId` | String | 是 | id | '' |
| `groupTitle` | String | 是 | 组名称 | '' |
| `groupCode` | String | 是 | 组编码 | '' |
| `groupDescLink` | String | 是 | 组描述链接 | '' |
| `tradeServiceList` | [message:com.alibaba.ocean.openplatform.biz.trade.common.model.TradeService[]](#m-com-alibaba-ocean-openplatform-biz-trade-common-model-tradeservice[]) | 是 | 服务信息 | '' |
| `extraMapStr` | String | 是 | 扩展信息 | '' |
| `combinedPrice` | Long | 是 | 合并价格 | '' |
| `showCombinedPrice` | String | 是 | 展示合并价格 | '' |

<a id="m-com-alibaba-ocean-openplatform-biz-trade-common-model-tradeservice[]"></a>
#### com.alibaba.ocean.openplatform.biz.trade.common.model.TradeService[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `sourceType` | String | 是 | 源类型 | '' |
| `serviceType` | String | 是 | 服务类型 | '' |
| `sourceId` | String | 是 | 源id | '' |
| `title` | String | 是 | 名称 | '' |
| `unitPrice` | Long | 是 | 单位价格 | '' |
| `quantity` | Double | 是 | 数量 | '' |
| `serviceTips` | String | 是 | 服务tip | '' |
| `serviceUrl` | String | 是 | 服务url | '' |
| `groupTitle` | String | 是 | 组名称 | '' |
| `groupCode` | String | 是 | 组code | '' |
| `sponsorDesc` | String | 是 | 描述 | '' |
| `code` | String | 是 | code | '' |
| `extMapStr` | String | 是 | 扩展信息 | '' |

<a id="m-com-alibaba-ocean-openplatform-biz-trade-common-model-officialsolutionmodel[]"></a>
#### com.alibaba.ocean.openplatform.biz.trade.common.model.OfficialSolutionModel[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `solutionCode` | String | 是 | 物流解决方案编码 | 471 |
| `solutionName` | String | 是 | 物流解决方案名称 | 特惠当日上门 |
| `totalCost` | Long | 是 | 总费用(分) | 770 |

<a id="m-com-alibaba-ocean-openplatform-biz-trade-common-model-extrpair[]"></a>
#### com.alibaba.ocean.openplatform.biz.trade.common.model.ExtrPair[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `key` | String | 是 | 键 | doorPickup |
| `value` | String | 是 | 值 | {       "serviceCode": "mkd",       "price": 1200, // 人民币，分     } |
| `desc` | String | 是 | 描述 | 是否使用上门揽 |

## 错误码

| 错误码 | 现象 | 解决方案 |
|---|---|---|
| 500_001 | 商品[offerId]不支持在线交易，无法下单。 | 商品不支持在线交易，目前不能购买 |
| 500_002 | 商品[offerId]不属于同一卖家或者没有指定specId。 | 存在多个卖家的商品或者商品没有指定specId |
| 500_003 | 商品[offerId ]不属于同一卖家或者规格[specId] 不属于商品[offerId] | 存在多个卖家的商品或者商品不存在specId的规格 |
| 500_004 | 商品[offerId_specId]库存不足，请核实库存后订购。 | 商品的某个规格库存不足 |
| 500_005 | 商品[offerId]的购买数量不满足起批量限制。 | 商品的购买数量小于起批量 |
| 500_006 | 商品[offerId]的购买数量或者价格不满足混批限制。 | 商品的购买数量或者总金额均不满足混批条件 |
| 500_007 | 与供应商的代销关系不存在,不能使用saleproxy通道下单。 | flow不能使用slproxy |
| 500_009 | 商品[offerId]的购买数量不满足批售起批量限制。 | 检查批售商品的购买数量 |
| 500_008 | 商品规格[offerId_specId]的价格为0，不可以下单，请检查后重新提交。 | 检查下商品规格的价格 |

## 示例

**出参示例**

```
{
    "orderPreviewResuslt":[
        {
            "tradeModeNameList":[
                "nzassure"
            ],
            "status":true,
            "taoSampleSinglePromotion":false,
            "sumPayment":4400,
            "sumCarriage":800,
            "sumPaymentNoCarriage":3600,
            "flowFlag":"general",
            "cargoList":[
                {
                    "amount":36,
                    "finalUnitPrice":18,
                    "specId":"eb81c61de14f4adb405ffcc2c8a4a3fb",
                    "skuId":3332085412530,
                    "offerId":548818919805,
                    "cargoPromotionList":[

                    ]
                }
            ],
            "shopPromotionList":[

            ],
            "tradeModelList":[
                {
                    "tradeType":"nzassure",
                    "name":"担保交易",
                    "description":"买家下单5天内全款支付。自卖家发货起，买家需在10天内确认收货，确认收货后打款给卖家。特别提醒：如卖家已开通极速到账服务，将优先适用极速到账交易，买家支付款项将直接打入卖家账户，详见《极速到账交易及争议处理规则》。",
                    "opSupport":true
                }
            ]
        }
    ],
    "success":true
}
```
