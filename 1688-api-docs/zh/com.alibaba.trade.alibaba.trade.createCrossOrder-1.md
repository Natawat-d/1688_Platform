# 跨境订单创建

API: `com.alibaba.trade:alibaba.trade.createCrossOrder:1` · Category: 订单  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.trade.createCrossOrder-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.trade.createCrossOrder/{appKey}`  
需要授权 (access_token) · 需要签名

跨境专用订单创建。创建订单最多允许50个SKU，且必须为同一个供应商的商品。多个供应商或多于50个SKU的情况，请自行拆单后提交。一些特殊情况会一次创建多个订单并返回多个订单号。
支持大市场及分销两个场景。根据当前授权用户,区分主子账号下单

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `flow` | String | 是 | general（创建批发订单），fenxiao（创建分销订单）,saleproxy流程将校验分销关系,paired(天天特卖),repurchase(复购合约下单),boutiquefenxiao(精选货源分销价下单，采购量1个使用包邮)， boutiquepifa(精选货源批发价下单，采购量大于2使用).不同的下单flow渠道，可能导致价格和买保服务差异。 | general |
| `message` | String | 否 | 买家留言 | 留言 |
| `isvBizType` | String | 否 | 开放平台业务码,默认为cross，代表对接的1688业务方。cross(跨境业务),cross_daigou（跨境代购业务）cross_distribution（海外分销业务）；可以联系你的业务小二确定，本字段不影响下单 | cross |
| `addressParam` | [message:alibaba.trade.fast.address](#m-alibaba-trade-fast-address) | 是 | 收货地址信息 | {"address":"网商路699号","phone": "0517-88990077","mobile": "15251667788","fullName": "张三","postCode": "000000","areaText": "滨江区","townText": "","cityText": "杭州市","provinceText": "浙江省"} |
| `cargoParamList` | [message:alibaba.trade.fast.cargo[]](#m-alibaba-trade-fast-cargo[]) | 是 | 商品信息 | [{"specId": "b266e0726506185beaf205cbae88530d","quantity": 5,"offerId": 554456348334},{"specId": "2ba3d63866a71fbae83909d9b4814f01","quantity": 6,"offerId": 554456348334}] |
| `invoiceParam` | [message:alibaba.trade.fast.invoice](#m-alibaba-trade-fast-invoice) | 否 | 发票信息 | {"invoiceType":0,"cityText": "杭州市","provinceText": "浙江省","address": "网商路699号","phone": "0517-88990077","mobile": "15251667788","fullName": "张五","postCode": "000000","areaText": "滨江区","companyName": "测试公司","taxpayerIdentifier": "123455"} |
| `tradeType` | String | 否 | 由于不同的商品支持的交易方式不同，没有一种交易方式是全局通用的，所以当前下单可使用的交易方式必须通过下单预览接口的tradeModeNameList获取。交易方式类型说明：assureTrade（交易4.0通用担保交易），alipay（大市场通用的支付宝担保交易（目前在做切流，后续会下掉）），period（普通账期交易）, assure（大买家企业采购询报价下单时需要使用的担保交易流程）, creditBuy（诚E赊），bank（银行转账），631staged（631分阶段付款），37staged（37分阶段）；此字段不传则系统默认会选取一个可用的交易方式下单，如果开通了诚E赊默认是creditBuy（诚E赊），未开通诚E赊默认使用的方式是支付宝担宝交易。 | assureTrade |
| `shopPromotionId` | String | 否 | 店铺优惠ID，通过“创建订单前预览数据接口”获得。为空默认使用默认优惠 | itemCoupon-5600812521_31032085284-398517001570 |
| `anonymousBuyer` | Boolean | 否 | 是否匿名下单 |   |
| `fenxiaoChannel` | String | 否 | 回流订单下游平台 淘宝-thyny，天猫-tm，淘特-taote，阿里巴巴C2M-c2m，京东-jingdong，拼多多-pinduoduo，微信-weixin，跨境-kuajing，快手-kuaishou，有赞-youzan，抖音-douyin，寺库-siku，美团团好货-meituan，小红书-xiaohongshu，当当-dangdang，苏宁-suning，大V店-davdian，行云-xingyun，蜜芽-miya，菠萝派商城-boluo，快团团-kuaituantuan，其他-other | douyin |
| `inventoryMode` | String | 否 | 库存模式，JIT（jit模式）或 NORMAL（仓发模式）,目前只提供给AE使用 | JIT |
| `outOrderId` | String | 否 | 外部订单号 | 988129883123 |
| `pickupService` | String | 否 | 上门揽收,目前AE供货可用，其他场景暂不开通.y或n,默认为n | n |
| `warehouseCode` | String | 否 | 上门揽仓库code | any |
| `preSelectPayChannel` | String | 否 | 预选的支付渠道，用作财务订单分流。订单信息查询接口返回：result.exAttributes.preSelectPayChannel ，该值是创建订单接口时传入的预选的支付渠道标记。 | alipay |
| `smallProcurement` | String | 否 | 是否小额采购，目前AE供货可用，取值y/n，默认为n | y |
| `useRedEnvelope` | String | 否 | 使用红包：y使用，n不使用。默认使用红包 | n |
| `dropshipping` | String | 否 | 是否转运订单，取值y/n，默认为n | y |
| `addedService` | String | 否 | 增值服务 toB、toC | toB |
| `crossBorderLogisticsSolutionId` | String | 否 | 用户选择的官方物流跨境解决方案sourceId | GLOBAL_CAINIAO_VN_TRANSIT_LAND |
| `useCrossBorderLogisticsSolution` | Boolean | 否 | 是否使用跨境物流解决方案预览，默认为false，当海外地址时，参数不生效。当为港澳台地址时，参数生效 | true |
| `extendParam` | [message:com.alibaba.ocean.openplatform.biz.trade.param.ExtendParam](#m-com-alibaba-ocean-openplatform-biz-trade-param-extendparam) | 否 | 下单扩展信息 | {} |
| `useOfficialSolution` | Boolean | 否 | 配送方式使用官方物流提货 | false |
| `useOfficialSolutionModelList` | [message:com.alibaba.ocean.openplatform.biz.trade.common.model.UseOfficialSolutionModel[]](#m-com-alibaba-ocean-openplatform-biz-trade-common-model-useofficialsolutionmodel[]) | 否 | 官方物流提货物流解决方案（useOfficialSolution为true时必填） | [] |
| `fromAgent` | String | 否 | 外部agent | qoderwork |

<a id="m-alibaba-trade-fast-address"></a>
#### alibaba.trade.fast.address

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `addressId` | Long | 是 | 收货地址id | 1234 |
| `fullName` | String | 是 | 收货人姓名 | 张三 |
| `mobile` | String | 是 | 手机 | 15251667788 |
| `phone` | String | 是 | 电话 | 0517-88990077  注意: 只能填电话号码,不能加其他中文 |
| `postCode` | String | 是 | 邮编 | 000000 |
| `cityText` | String | 是 | 市文本 | 杭州市 |
| `provinceText` | String | 是 | 省份文本 | 浙江省 |
| `areaText` | String | 是 | 区文本 | 滨江区 |
| `townText` | String | 是 | 镇文本 | 长河镇 |
| `address` | String | 是 | 街道地址 | 网商路699号 |
| `districtCode` | String | 是 | 地址编码 | 310107 |
| `addressCode` | String | 是 | 地址编码 | 地址码 |

<a id="m-alibaba-trade-fast-cargo[]"></a>
#### alibaba.trade.fast.cargo[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `offerId` | Long | 是 | 商品对应的offer id | 554456348334 |
| `specId` | String | 是 | 商品规格id | b266e0726506185beaf205cbae88530d |
| `quantity` | Double | 是 | 商品数量(计算金额用) | 5 |
| `openOfferId` | String | 否 | 加密offerId，当搜索返回只有openOfferId时，需采用openOfferId替代offerId下单 | Wcv2w970KoL1BCpJGQp3vwKvcBThOPlpHnUtkK3T0n4= |
| `outMemberId` | String | 否 | 外部下游会员ID | 98928912-23 |

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

<a id="m-com-alibaba-ocean-openplatform-biz-trade-param-extendparam"></a>
#### com.alibaba.ocean.openplatform.biz.trade.param.ExtendParam

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `crossBorderLogisticsSelfPickupAddress` | [message:com.alibaba.ocean.openplatform.biz.trade.param.CrossBorderLogisticsSelfPickupAddress](#m-com-alibaba-ocean-openplatform-biz-trade-param-crossborderlogisticsselfpickupaddress) | 是 | 跨境集运仓自提点地址和联系人 | {} |
| `extPairList` | [message:com.alibaba.ocean.openplatform.biz.trade.common.model.ExtrPair[]](#m-com-alibaba-ocean-openplatform-biz-trade-common-model-extrpair[]) | 否 | 扩展对象列表 | [] |
| `offerIsvCargoFromList` | [message:alibaba.ocean.openplatform.common.KeyValuePair[]](#m-alibaba-ocean-openplatform-common-keyvaluepair[]) | 否 | 轻应用闭环下单，offer来源信息列表 | [] |
| `selectedTradeServiceList` | [message:com.alibaba.ocean.openplatform.biz.trade.param.TradeService[]](#m-com-alibaba-ocean-openplatform-biz-trade-param-tradeservice[]) | 否 | 交易服务参数 | [] |

<a id="m-com-alibaba-ocean-openplatform-biz-trade-param-crossborderlogisticsselfpickupaddress"></a>
#### com.alibaba.ocean.openplatform.biz.trade.param.CrossBorderLogisticsSelfPickupAddress

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `addressText` | String | 是 | 集运仓自提地址，不用填写省市区，省市区跟默认解决方案集运仓一致 | 网商路699号 |
| `mobile` | String | 是 | 集运仓自提手机号 | 18792983782 |
| `contractName` | String | 是 | 集运仓自提联系人 | 无名 |

<a id="m-com-alibaba-ocean-openplatform-biz-trade-common-model-extrpair[]"></a>
#### com.alibaba.ocean.openplatform.biz.trade.common.model.ExtrPair[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `key` | String | 是 | 键 | doorPickup |
| `value` | String | 是 | 值 | {       "serviceCode": "mkd",       "price": 1200, // 人民币，分     } |
| `desc` | String | 是 | 描述 | 是否使用上门揽 |

<a id="m-alibaba-ocean-openplatform-common-keyvaluepair[]"></a>
#### alibaba.ocean.openplatform.common.KeyValuePair[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `key` | java.lang.String | 是 | offerId |    |
| `value` | java.lang.String | 是 | offer来源信息 |    |

<a id="m-com-alibaba-ocean-openplatform-biz-trade-param-tradeservice[]"></a>
#### com.alibaba.ocean.openplatform.biz.trade.param.TradeService[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `code` | String | 是 | 交易服务code | vas |
| `sourceId` | String | 是 | 交易服务sourceId | 3C |

<a id="m-com-alibaba-ocean-openplatform-biz-trade-common-model-useofficialsolutionmodel[]"></a>
#### com.alibaba.ocean.openplatform.biz.trade.common.model.UseOfficialSolutionModel[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `orderGroup` | String | 是 | 订单分组 | 123 |
| `useOfficialSolutionCode` | String | 是 | 官方物流提货物流解决方案编码 | 123 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:alibaba.trade.cross.result](#m-alibaba-trade-cross-result) | 是 | 创建订单结果 | {} |
| `success` | Boolean | 是 | 是否成功 | true |
| `code` | String | 是 | 错误码 | 400 |
| `message` | String | 是 | 错误信息 | AddressId invalid:AddressId invalid |

<a id="m-alibaba-trade-cross-result"></a>
#### alibaba.trade.cross.result

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `totalSuccessAmount` | Long | 是 | 订单总金额（单位分），一次创建多个订单时，该字段为空 | 100 |
| `orderId` | String | 是 | 订单ID，一次创建多个订单时，该字段为空 | 111111111 |
| `success` | Boolean | 是 | 是否成功 | true |
| `code` | String | 是 | 错误码 |  |
| `message` | String | 是 | 错误信息 |  |
| `accountPeriod` | [message:alibaba.trade.cross.period](#m-alibaba-trade-cross-period) | 是 | 账期信息，非账期支付订单返回空 |  |
| `failedOfferList` | [message:alibaba.trade.fast.offer[]](#m-alibaba-trade-fast-offer[]) | 是 | 失败商品信息 |  |
| `postFee` | Long | 是 | 运费，单位：分，一次创建多个订单时，该字段为空 |  |
| `orderList` | [message:alibaba.tradeResult.BizSimpleOrder[]](#m-alibaba-traderesult-bizsimpleorder[]) | 是 | 一次创建多个订单 |  |

<a id="m-alibaba-trade-cross-period"></a>
#### alibaba.trade.cross.period

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `tapType` | Integer | 是 | 账期的类型,1：一个月指定日期结算一次，3：两个月指定日期结算一次，6：三个月指定日期结算一次，5：按收货时间和账期日期结算 | 1 |
| `tapDate` | Integer | 是 | 根据账期类型不同而不同，按月结算类型此值代表具体某日，按收货时间结算时此值代表结算时间周期 | 12 |
| `tapOverdue` | Integer | 是 | 逾期次数 | 0 |

<a id="m-alibaba-trade-fast-offer[]"></a>
#### alibaba.trade.fast.offer[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `offerId` | String | 是 | 下单失败的商品 | 554456348334 |
| `specId` | String | 是 | 下单失败商品的规格ID | b266e0726506185beaf205cbae88530d |
| `errorCode` | String | 是 | 下单失败的错误编码 |  |
| `errorMessage` | String | 是 | 下单失败的错误描述 |  |

<a id="m-alibaba-traderesult-bizsimpleorder[]"></a>
#### alibaba.tradeResult.BizSimpleOrder[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `postFee` | java.lang.Long | 是 | 运费 | 10 |
| `orderAmmount` | java.lang.Long | 是 | 订单实付款金额，单位为分 | 100 |
| `message` | java.lang.String | 是 | 描述信息 | null |
| `resultCode` | java.lang.String | 是 | 返回码 | 200 |
| `success` | java.lang.Boolean | 是 | 是否成功 | true |
| `orderId` | java.lang.String | 是 | 订单号 | 123124124 |
| `payChannel` | String | 是 | 支付渠道 | shegou |

## 错误码

| 错误码 | 现象 | 解决方案 |
|---|---|---|
| not support tradeType:【XXXX】 | 不支持的交易类型 | 当前交易可支持的交易类型，需要通过预览接口的tradeModeNameList字段获取 |

## 示例

**支付宝方式下单返回示例**

```
{"totalSuccessAmount":156800,"orderId":"87407346014789305","success":false}
```

**账期支付下单返回示例**

```
{"totalSuccessAmount":156800,"orderId":"87407346015789305","success":false,"accountPeriod":{"tapType":5,"tapDate":360,"tapOverdue":1}}
```

**收货地址说明**

```
1、使用保存的收货地址传参示例：{"addressId":593861699}，其中addressId是调用“买家获取保存的收货地址信息列表”接口获取；
2、使用地址编码传参示例：{"address":"桃浦镇 金达路888号贸易8楼","phone": "0517-88990077","mobile": "15251667788","fullName": "张三","postCode": "000000","districtCode": "310107"}，其中districtCode需要调用“根据地址解析地区码”接口获取；
3、直接使用文本地址示例：{"address":"网商路699号","phone": "0517-88990077","mobile": "15251667788","fullName": "张三","postCode": "000000","areaText": "滨江区","townText": "","cityText": "杭州市","provinceText": "浙江省"}，省市区要传文本；
4、优先级说明，如果同时传入以上参数，系统按从1至3的优先级获取地址，满足1的条件下不会使用上示2中的参数，满足2的条件下不会使用上示3中的参数；
```

**同时创建多个订单的返回值**

```
{
    "result":{
        "success":true,
        "orderList":[
            {
                "postFee":48000,
                "orderAmmount":4848000,
                "discount":0,
                "sumPaymentNoCarriageFromClient":4800000,
                "mergePay":false,
                "orderId":"105581756010628640",
                "chooseFreeFreight":false
            },
            {
                "postFee":50000,
                "orderAmmount":5050000,
                "discount":0,
                "sumPaymentNoCarriageFromClient":5000000,
                "mergePay":false,
                "orderId":"105637283010628640",
                "chooseFreeFreight":false
            }
        ]
    },
    "success":true
}
```
