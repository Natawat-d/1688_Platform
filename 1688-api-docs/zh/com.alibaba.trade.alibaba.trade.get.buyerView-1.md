# 订单详情查看(买家视角)

API: `com.alibaba.trade:alibaba.trade.get.buyerView:1` · Category: 订单  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.trade.get.buyerView-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.trade.get.buyerView/{appKey}`  
需要授权 (access_token) · 需要签名

获取单个交易明细信息，仅限买家调用。该API需要向阿里巴巴开放平台申请权限才能使用。Get a single transaction detail, only for users to call.

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `webSite` | String | 是 | 站点信息，指定调用的API是属于国际站（alibaba）还是1688网站（1688） | 1688 |
| `orderId` | Long | 是 | 交易的订单id | 123456 |
| `includeFields` | String | 否 | 查询结果中包含的域，GuaranteesTerms：保障条款，NativeLogistics：物流信息，RateDetail：评价详情，OrderInvoice：发票信息。默认返回GuaranteesTerms、NativeLogistics、OrderInvoice。InvoicingSetting-开票设置 | GuaranteesTerms,NativeLogistics,RateDetail,OrderInvoice |
| `attributeKeys` | String[] | 否 | 垂直表中的attributeKeys | [] |
| `outOrderId` | String | 否 | 外部订单id，控制幂等 | 1556246 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:alibaba.openplatform.trade.model.TradeInfo](#m-alibaba-openplatform-trade-model-tradeinfo) | 是 | 订单详情信息 | {} |
| `errorCode` | String | 是 | 错误代码 |   |
| `errorMessage` | String | 是 | 错误描述 |   |
| `success` | String | 是 | 是否成功 | true |

<a id="m-alibaba-openplatform-trade-model-tradeinfo"></a>
#### alibaba.openplatform.trade.model.TradeInfo

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `baseInfo` | [message:alibaba.openplatform.trade.model.OrderBaseInfo](#m-alibaba-openplatform-trade-model-orderbaseinfo) | 是 | 订单基础信息 | {} |
| `orderBizInfo` | [message:alibaba.order.bizInfo](#m-alibaba-order-bizinfo) | 是 | 订单业务信息 | {} |
| `tradeTerms` | [message:alibaba.openplatform.trade.model.TradeTermsInfo[]](#m-alibaba-openplatform-trade-model-tradetermsinfo[]) | 是 | 交易条款 | {} |
| `productItems` | [message:alibaba.openplatform.trade.model.ProductItemInfo[]](#m-alibaba-openplatform-trade-model-productiteminfo[]) | 是 | 商品条目信息 | {} |
| `nativeLogistics` | [message:alibaba.openplatform.trade.model.NativeLogisticsInfo](#m-alibaba-openplatform-trade-model-nativelogisticsinfo) | 是 | 国内物流 | {} |
| `orderInvoiceInfo` | [message:alibaba.invoice.OrderInvoiceModel](#m-alibaba-invoice-orderinvoicemodel) | 是 | 发票信息 | {} |
| `guaranteesTerms` | [message:alibaba.openplatform.trade.model.GuaranteeTermsInfo](#m-alibaba-openplatform-trade-model-guaranteetermsinfo) | 是 | 保障条款 | {} |
| `orderRateInfo` | [message:alibaba.trade.OrderRateInfo](#m-alibaba-trade-orderrateinfo) | 是 | 订单评价信息 | {} |
| `overseasExtraAddress` | [message:alibaba.trade.OverseasExtraAddress](#m-alibaba-trade-overseasextraaddress) | 是 | 跨境地址扩展信息 | {} |
| `customs` | [message:alibaba.trade.Customs](#m-alibaba-trade-customs) | 是 | 跨境报关信息 | {} |
| `quoteList` | [message:alibaba.orderDetail.caigouQuoteInfo[]](#m-alibaba-orderdetail-caigouquoteinfo[]) | 是 | 采购单详情列表，为大企业采购订单独有域。 | {} |
| `extAttributes` | [message:alibaba.openplatform.trade.KeyValuePair[]](#m-alibaba-openplatform-trade-keyvaluepair[]) | 是 | 订单扩展属性 | {} |
| `fromEncryptOrder` | Boolean | 是 | 是否下游脱敏信息创建的订单 | true |
| `encryptOutOrderInfo` | [message:alibaba.trade.get.sellerView.tradeinfo.EncryptOutOrderInfo](#m-alibaba-trade-get-sellerview-tradeinfo-encryptoutorderinfo) | 是 | 外部订单信息 | {} |
| `overseaLogisticsInfo` | [message:com.alibaba.ocean.openplatform.biz.trade.common.model.OverseaLogisticsInfo](#m-com-alibaba-ocean-openplatform-biz-trade-common-model-oversealogisticsinfo) | 是 | 海外物流信息 | {} |
| `invoicingSettingModel` | [message:com.alibaba.ocean.openplatform.biz.trade.common.model.OpSellerInvoiceTradeSettingModel](#m-com-alibaba-ocean-openplatform-biz-trade-common-model-opsellerinvoicetradesettingmodel) | 是 | 开票设置 | {} |

<a id="m-alibaba-openplatform-trade-model-orderbaseinfo"></a>
#### alibaba.openplatform.trade.model.OrderBaseInfo

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `allDeliveredTime` | java.util.Date | 是 | 完全发货时间 | 20180614101942000+0800 |
| `sellerCreditLevel` | java.lang.String | 是 | 卖家诚信等级 | L1 |
| `payTime` | java.util.Date | 是 | 付款时间，如果有多次付款，这里返回的是首次付款时间 | 20180614101942000+0800 |
| `discount` | Long | 是 | 折扣信息，单位分 | 11 |
| `alipayTradeId` | java.lang.String | 是 | 外部支付交易Id | 123123121111 |
| `sumProductPayment` | java.math.BigDecimal | 是 | 产品总金额(该订单产品明细表中的产品金额的和)，单位元 | 1212 |
| `buyerFeedback` | String | 是 | 买家留言，不超过500字 | 留言 |
| `flowTemplateCode` | String | 是 | 4.0交易流程模板code | flow |
| `sellerOrder` | java.lang.Boolean | 是 | 是否自主订单（邀约订单） | false |
| `buyerLoginId` | java.lang.String | 是 | 买家loginId，旺旺Id | alitestforusv01 |
| `modifyTime` | java.util.Date | 是 | 修改时间 | 20180614101942000+0800 |
| `subBuyerLoginId` | String | 是 | 买家子账号 | alitestforusv02:temp |
| `id` | java.lang.Long | 是 | 交易id | 1231231231111 |
| `closeReason` | java.lang.String | 是 | 关闭原因。buyerCancel:买家取消订单，sellerGoodsLack:卖家库存不足，other:其它 | buyerCancel |
| `buyerContact` | [message:alibaba.trade.tradeContact](#m-alibaba-trade-tradecontact) | 是 | 买家联系人 | {} |
| `sellerAlipayId` | java.lang.String | 是 | 卖家支付宝id | 12312311111 |
| `completeTime` | java.util.Date | 是 | 完成时间 | 20180614101942000+0800 |
| `sellerLoginId` | java.lang.String | 是 | 卖家oginId，旺旺Id | alitestforusv02 |
| `buyerID` | java.lang.String | 是 | 买家主账号id | 1234531 |
| `closeOperateType` | String | 是 | 关闭订单操作类型。CLOSE_TRADE_BY_SELLER:卖家关闭交易,CLOSE_TRADE_BY_BOPS:BOPS后台关闭交易,CLOSE_TRADE_BY_SYSTEM:系统（超时）关闭交易,CLOSE_TRADE_BY_BUYER:买家关闭交易,CLOSE_TRADE_BY_CREADIT:诚信保障投诉关闭 | CLOSE_TRADE_BY_SELLER |
| `totalAmount` | java.math.BigDecimal | 是 | 应付款总金额，totalAmount = ∑itemAmount + shippingFee，单位为元 | 1000 |
| `sellerID` | java.lang.String | 是 | 卖家主账号id | 123123123123 |
| `shippingFee` | java.math.BigDecimal | 是 | 运费，单位为元 | 1 |
| `buyerUserId` | java.lang.Long | 是 | 买家数字id | 12314144 |
| `buyerMemo` | java.lang.String | 是 | 买家备忘信息 | 备忘 |
| `refund` | java.math.BigDecimal | 是 | 退款金额，单位为元 | 1 |
| `status` | java.lang.String | 是 | 交易状态，waitbuyerpay:等待买家付款;waitsellersend:等待卖家发货;waitbuyerreceive:等待买家收货;confirm_goods:已收货;success:交易成功;cancel:交易取消;terminated:交易终止;未枚举:其他状态 | waitbuyerpay |
| `refundPayment` | Long | 是 | 退款金额 | 1 |
| `sellerContact` | [message:alibaba.trade.tradeSellerContact](#m-alibaba-trade-tradesellercontact) | 是 | 卖家联系人信息 | {} |
| `couponFee` | java.math.BigDecimal | 是 | 红包金额，实付金额（totalAmount）已经计算过红包金额 | 7.5 |
| `buyerRemarkIcon` | String | 是 | 买家备忘标志 | 1 |
| `receiverInfo` | [message:alibaba.trade.orderReceiverInfo](#m-alibaba-trade-orderreceiverinfo) | 是 | 收件人信息 | {} |
| `refundStatus` | String | 是 | 订单的售中退款状态，等待卖家同意：waitselleragree ，待买家修改：waitbuyermodify，等待买家退货：waitbuyersend，等待卖家确认收货：waitsellerreceive，退款成功：refundsuccess，退款失败：refundclose | refundclose |
| `remark` | java.lang.String | 是 | 备注，1688指下单时的备注 | 备注 |
| `preOrderId` | java.lang.Long | 是 | 预订单ID | 123123 |
| `confirmedTime` | java.util.Date | 是 | 确认时间 | 20180614101942000+0800 |
| `closeRemark` | String | 是 | 关闭订单备注 | 备注 |
| `tradeType` | String | 是 | 1:担保交易<br>2:预存款交易<br>3:ETC境外收单交易<br>4:即时到帐交易<br>5:保障金安全交易<br>6:统一交易流程<br>7:分阶段付款<br>8.货到付款交易<br>9.信用凭证支付交易<br>10.账期支付交易，50060 交易4.0 | 50060 |
| `receivingTime` | java.util.Date | 是 | 收货时间，这里返回的是完全收货时间 | 20180614101942000+0800 |
| `stepAgreementPath` | java.lang.String | 是 | 分阶段法务协议地址 |   |
| `idOfStr` | String | 是 | 交易id(字符串格式) | 123121212123 |
| `refundStatusForAs` | String | 是 | 订单的售后退款状态 |   |
| `stepPayAll` | java.lang.Boolean | 是 | 是否一次性付款 | false |
| `sellerUserId` | java.lang.Long | 是 | 卖家数字id | 12312422 |
| `stepOrderList` | [message:alibaba.trade.StepOrderModel[]](#m-alibaba-trade-stepordermodel[]) | 是 | [交易3.0]分阶段交易，分阶段订单list |   |
| `newStepOrderList` | [message:alibaba.trade.BizNewStepOrderModel[]](#m-alibaba-trade-biznewstepordermodel[]) | 是 | [交易4.0]分阶段交易，分阶段订单list |   |
| `buyerAlipayId` | java.lang.String | 是 | 买家支付宝id | 12312311233 |
| `createTime` | java.util.Date | 是 | 创建时间 | 20180614101942000+0800 |
| `businessType` | java.lang.String | 是 | 业务类型。国际站：ta(信保),wholesale(在线批发)。<br>中文站：普通订单类型 = &quot;cn&quot;;<br>大额批发订单类型 = &quot;ws&quot;;<br>普通拿样订单类型 = &quot;yp&quot;;<br>一分钱拿样订单类型 = &quot;yf&quot;;<br>倒批(限时折扣)订单类型 = &quot;fs&quot;;<br>加工定制订单类型 = &quot;cz&quot;;<br>协议采购订单类型 = &quot;ag&quot;;<br>伙拼订单类型 = &quot;hp&quot;;<br>供销订单类型 = &quot;supply&quot;;<br>淘工厂订单 = &quot;factory&quot;;<br>快订下单  = &quot;quick&quot;;<br>享拼订单  = &quot;xiangpin&quot;;<br>当面付 = &quot;f2f&quot;;<br>存样服务 = &quot;cyfw&quot;;<br>代销订单 = &quot;sp&quot;;<br>微供订单 = &quot;wg&quot;;零售通 = &quot;lst&quot;;跨境=&#39;cb&#39;;分销=&#39;distribution&#39;;采源宝=&#39;cab&#39;;加工定制=&quot;manufact&quot; | cn |
| `overSeaOrder` | java.lang.Boolean | 是 | 是否海外代发订单，是：true | true |
| `refundId` | java.lang.String | 是 | 退款单ID | TQ4562212313 |
| `tradeTypeDesc` | String | 是 | 下单时指定的交易方式 | 担保交易 |
| `payChannelList` | String[] | 是 | 支付渠道名称列表。一笔订单可能存在多种支付渠道。枚举值：支付宝,网商银行信任付,诚e赊,对公转账,赊销宝,账期支付,合并支付渠道,支付平台,声明付款,网商电子银行承兑汇票,银行转账,跨境宝,红包,其它 | ["支付宝","跨境宝","银行转账"] |
| `tradeTypeCode` | String | 是 | 下单时指定的交易方式tradeType | assureTrade |
| `payTimeout` | Long | 是 | 支付超时时间，定长情况时单位：秒，目前都是定长 | 43200 |
| `payTimeoutType` | Integer | 是 | 支付超时TYPE，0：定长，1：固定时间 | 0 |
| `payChannelCodeList` | String[] | 是 | 支付渠道code，payChannelCodeList的中文示意参见payChannelList | ["alipay"] |
| `inventoryMode` | String | 是 | 供货库存模式，jit（jit模式）或cang（仓发模式） | jit |
| `outOrderId` | String | 是 | 外部订单号 | 1928919827731 |
| `officialSolutionOrderId` | String | 是 | 官方物流提货订单ID | 123 |
| `officialSolutionCost` | Long | 是 | 官方物流提货订单费用 | 770 |

<a id="m-alibaba-trade-tradecontact"></a>
#### alibaba.trade.tradeContact

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `phone` | String | 是 | 联系电话 |  |
| `fax` | String | 是 | 传真 |  |
| `email` | String | 是 | 邮箱 |  |
| `imInPlatform` | String | 是 | 联系人在平台的IM账号 |  |
| `name` | String | 是 | 联系人名称 |  |
| `mobile` | String | 是 | 联系人手机号 |  |
| `companyName` | java.lang.String | 是 | 公司名称 |  |

<a id="m-alibaba-trade-tradesellercontact"></a>
#### alibaba.trade.tradeSellerContact

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `phone` | String | 是 | 联系电话 | 189982787712 |
| `fax` | String | 是 | 传真 | 189982787712 |
| `email` | String | 是 | 邮箱 | xxx@mail.com |
| `imInPlatform` | String | 是 | 联系人在平台的IM账号 | wangwang |
| `name` | String | 是 | 联系人名称 | 张三 |
| `mobile` | String | 是 | 联系人手机号 | 189982787712 |
| `companyName` | java.lang.String | 是 | 公司名称 | 三生网媒邮箱公司 |
| `wgSenderName` | java.lang.String | 是 | 发件人名称，在微供等分销场景下由分销商设置 | 张** |
| `wgSenderPhone` | java.lang.String | 是 | 发件人电话，在微供等分销场景下由分销商设置 | 13800000000 |
| `shopName` | String | 是 | 旺铺名称 |  三生网媒邮箱公司 |

<a id="m-alibaba-trade-orderreceiverinfo"></a>
#### alibaba.trade.orderReceiverInfo

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `toFullName` | java.lang.String | 是 | 收件人 |  |
| `toDivisionCode` | java.lang.String | 是 | 收货人地址区域编码 |  |
| `toMobile` | java.lang.String | 是 | 收件人移动电话 |  |
| `toPhone` | java.lang.String | 是 | 收件人电话 |  |
| `toPost` | java.lang.String | 是 | 邮编 |  |
| `toTownCode` | java.lang.String | 是 | 收货人街道或镇区域编码，可能为空 |  |
| `toArea` | java.lang.String | 是 | 收货地址 |  |

<a id="m-alibaba-trade-stepordermodel[]"></a>
#### alibaba.trade.StepOrderModel[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `stepOrderId` | java.lang.Long | 是 | 阶段id |  |
| `stepOrderStatus` | java.lang.String | 是 | waitactivate  未开始（待激活）<br>waitsellerpush 等待卖家推进<br>success 本阶段完成<br>settlebill 分账<br>cancel 本阶段终止<br>inactiveandcancel 本阶段未开始便终止<br>waitbuyerpay 等待买家付款<br>waitsellersend 等待卖家发货<br>waitbuyerreceive 等待买家确认收货<br>waitselleract 等待卖家XX操作<br>waitbuyerconfirmaction 等待买家确认XX操作 |  |
| `stepPayStatus` | java.lang.Integer | 是 | 1 未冻结/未付款<br>2 已冻结/已付款<br>4 已退款<br>6 已转交易<br>8 交易未付款被关闭 |  |
| `stepNo` | java.lang.Integer | 是 | 阶段序列：1、2、3... |  |
| `lastStep` | java.lang.Boolean | 是 | 是否最后一个阶段 |  |
| `hasDisbursed` | java.lang.Boolean | 是 | 是否已打款给卖家 |  |
| `payFee` | java.math.BigDecimal | 是 | 创建时需要付款的金额，不含运费 |  |
| `actualPayFee` | java.math.BigDecimal | 是 | 应付款（含运费）= 单价×数量-单品优惠-店铺优惠+运费+修改的金额（除运费外，均指分摊后的金额） |  |
| `discountFee` | java.math.BigDecimal | 是 | 本阶段分摊的店铺优惠 |  |
| `itemDiscountFee` | java.math.BigDecimal | 是 | 本阶段分摊的单品优惠 |  |
| `price` | java.math.BigDecimal | 是 | 本阶段分摊的单价 |  |
| `amount` | java.lang.Long | 是 | 购买数量 |  |
| `postFee` | java.math.BigDecimal | 是 | 运费 |  |
| `adjustFee` | java.math.BigDecimal | 是 | 修改价格修改的金额 |  |
| `gmtCreate` | java.util.Date | 是 | 创建时间 |  |
| `gmtModified` | java.util.Date | 是 | 修改时间 |  |
| `enterTime` | java.util.Date | 是 | 开始时间 |  |
| `payTime` | java.util.Date | 是 | 付款时间 |  |
| `sellerActionTime` | java.util.Date | 是 | 卖家操作时间 |  |
| `endTime` | java.util.Date | 是 | 本阶段结束时间 |  |
| `messagePath` | java.lang.String | 是 | 卖家操作留言路径 |  |
| `picturePath` | java.lang.String | 是 | 卖家上传图片凭据路径 |  |
| `message` | java.lang.String | 是 | 卖家操作留言 |  |
| `templateId` | java.lang.Long | 是 | 使用的模板id |  |
| `stepName` | java.lang.String | 是 | 当前步骤的名称 |  |
| `sellerActionName` | java.lang.String | 是 | 卖家操作名称 |  |
| `buyerPayTimeout` | java.lang.Long | 是 | 买家不付款的超时时间(秒) |  |
| `buyerConfirmTimeout` | java.lang.Long | 是 | 买家不确认的超时时间 |  |
| `needLogistics` | java.lang.Boolean | 是 | 是否需要物流 |  |
| `needSellerAction` | java.lang.Boolean | 是 | 是否需要卖家操作和买家确认 |  |
| `transferAfterConfirm` | java.lang.Boolean | 是 | 阶段结束是否打款 |  |
| `needSellerCallNext` | java.lang.Boolean | 是 | 是否需要卖家推进 |  |
| `instantPay` | java.lang.Boolean | 是 | 是否允许即时到帐 |  |

<a id="m-alibaba-trade-biznewstepordermodel[]"></a>
#### alibaba.trade.BizNewStepOrderModel[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `gmtStart` | java.util.Date | 是 | 阶段开始时间 | 20180604092517000+0800 |
| `gmtPay` | java.util.Date | 是 | 付款时间 | 20180604093243000+0800 |
| `gmtEnd` | java.util.Date | 是 | 阶段结束时间 | 20180604093243000+0800 |
| `stepNo` | Integer | 是 | 阶段顺序编号 | 1 |
| `lastStep` | Boolean | 是 | 是否最后一个阶段 | true |
| `stepName` | String | 是 | 阶段名称 | 全款交易 |
| `activeStatus` | Integer | 是 | 激活状态。0表示未激活，1表示已激活 | 1 |
| `payStatus` | Integer | 是 | 阶段付款状态。1未付款、2已付款、8付款前取消、12溢短补付款 | 2 |
| `logisticsStatus` | Integer | 是 | 物流环节状态：1未发货、2已发货、3已收货、4已全部退货、7发货前取消 | 2 |
| `payFee` | java.math.BigDecimal | 是 | 阶段应付款（包含运费），单位为元 | 0.03 |
| `paidFee` | java.math.BigDecimal | 是 | 阶段已付款（包含运费），单位为元 | 0.03 |
| `goodsFee` | java.math.BigDecimal | 是 | 阶段商品价格分摊 ，单位为元 | 0 |
| `adjustFee` | java.math.BigDecimal | 是 | 阶段调整价格 ，单位为元 | -3175.97 |
| `discountFee` | java.math.BigDecimal | 是 | 阶段优惠价格，单位为元 | 0 |
| `postFee` | java.math.BigDecimal | 是 | 阶段的应付邮费，单位为元 | 0 |
| `paidPostFee` | java.math.BigDecimal | 是 | 阶段已付的邮费，单位为元 | 0 |

<a id="m-alibaba-order-bizinfo"></a>
#### alibaba.order.bizInfo

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `odsCyd` | java.lang.Boolean | 是 | 是否采源宝订单 | true |
| `accountPeriodTime` | java.lang.String | 是 | 账期交易订单的到账时间 | yyyy-MM-dd HH:mm:ss |
| `creditOrder` | java.lang.Boolean | 是 | 为true，表示下单时选择了诚e赊交易方式。注意不等同于“诚e赊支付”，支付时有可能是支付宝付款，具体支付方式查询tradeTerms.payWay | false |
| `creditOrderDetail` | [message:alibaba.creditOrder.forDetail](#m-alibaba-creditorder-fordetail) | 是 | 诚e赊支付详情，只有使用诚e赊付款时返回 |   |
| `preOrderInfo` | [message:alibaba.order.preOrder.forRead](#m-alibaba-order-preorder-forread) | 是 | 预订单信息 | {} |
| `lstOrderInfo` | [message:alibaba.lst.tradeInfo](#m-alibaba-lst-tradeinfo) | 是 | 零售通订单信息 | {} |
| `erpBuyerUserId` | String | 是 | ERP的用户ID | U001012121 |
| `erpOrderId` | String | 是 | ERP的订单编号 | O123331 |
| `erpBuyerOrgId` | String | 是 | ERP组织ID | OG4331113 |
| `isCz` | Boolean | 是 | 是否加工定制订单 | false |
| `isDz` | Boolean | 是 | 是否定制订单 | false |
| `dz` | Boolean | 是 | 是否定制订单 | false |
| `dropshipping` | Boolean | 是 | 是否dropshipping订单，该类型订单不允许合并发货 | true |
| `shippingInsurance` | String | 是 | givenByPlatform:平台赠送运费险 givenByMerchant:商家赠送运费险，为空表示订单无运费险 | givenByPlatform |
| `hyperLinkCangFaOrder` | Boolean | 是 | 大店仓发订单 | true |
| `hyperLinkOrder` | Boolean | 是 | 超链一阶段订单 | flase |
| `hyperLinkSecondStepOrder` | Boolean | 是 | 超链大店二阶段订单的第二阶段订单 | true |
| `hyperLinkShipType` | String | 是 | 超链一阶段订单的发货模式 0 仓发 1 商发 | 0 |
| `lightningWarehouse` | Boolean | 是 | 闪电仓订单 | true |
| `aeDoorPickUp` | Boolean | 是 | ae上门揽订单 | true |
| `fz` | Boolean | 是 | 分账订单 | true |
| `tgOfficialPickUp` | Boolean | 是 | 托管官方上门揽订单 | true |
| `officialPickUp` | Boolean | 是 | 官方物流提货 | true |
| `abnormalPriceChange` | Boolean | 是 | 是否异常改价 | true |

<a id="m-alibaba-creditorder-fordetail"></a>
#### alibaba.creditOrder.forDetail

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `payAmount` | java.lang.Long | 是 | 订单金额 | 10 |
| `createTime` | java.lang.String | 是 | 支付时间 | 2018-01-01 00:00:00 |
| `status` | java.lang.String | 是 | 状态 | END |
| `gracePeriodEndTime` | java.lang.String | 是 | 不再建议使用 | 2018-01-01 00:00:00 |
| `statusStr` | java.lang.String | 是 | 状态描述 | 已完结 |
| `restRepayAmount` | java.lang.Long | 是 | 应还金额 | 11 |
| `lastRepayTime` | String | 是 | 最晚还款时间 | 2023-01-01 00:00:00 |
| `repaySource` | String | 是 | 还款来源 KJPAY-跨境宝还款 OWN_FUNDS-自有资金还款 INSTALLMENT_REPAY-分期、贷款付还款 | KJPAY |

<a id="m-alibaba-order-preorder-forread"></a>
#### alibaba.order.preOrder.forRead

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `marketName` | String | 是 | 创建预订单时传入的市场名 | dxc |
| `createPreOrderApp` | Boolean | 是 | 预订单是否为当前查询的通过当前查询的ERP创建 | false |

<a id="m-alibaba-lst-tradeinfo"></a>
#### alibaba.lst.tradeInfo

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `lstWarehouseType` | String | 是 | 零售通仓库类型。customer：虚仓；cainiao：实仓 | cainiao |

<a id="m-alibaba-openplatform-trade-model-tradetermsinfo[]"></a>
#### alibaba.openplatform.trade.model.TradeTermsInfo[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `payStatus` | java.lang.String | 是 | 支付状态。国际站：WAIT_PAY(未支付),PAYER_PAID(已完成支付),PART_SUCCESS(部分支付成功),PAY_SUCCESS(支付成功),CLOSED(风控关闭),CANCELLED(支付撤销),SUCCESS(成功),FAIL(失败)。<br>1688:1(未付款);2(已付款);4(全额退款);6(卖家有收到钱，回款完成) ;7(未创建外部支付单);8 (付款前取消) ; 9(正在支付中);12(账期支付,待到账) |   |
| `payTime` | java.util.Date | 是 | 完成阶段支付时间 |   |
| `payWay` | java.lang.String | 是 | 支付方式。<br>国际站：ECL(融资支付),CC(信用卡),TT(线下TT),ACH(echecking支付)。<br>1688:1-支付宝,2-网商银行信任付,3-诚e赊,4-银行转账,5-赊销宝,6-电子承兑票据,7-账期支付,8-合并支付渠道,9-无打款,10-零售通赊购,13-支付平台,12-声明付款 |   |
| `phasAmount` | java.math.BigDecimal | 是 | 付款额 |   |
| `phase` | java.lang.Long | 是 | 阶段单id |   |
| `phaseCondition` | java.lang.String | 是 | 阶段条件，1688无此内容 |   |
| `phaseDate` | java.lang.String | 是 | 阶段时间，1688无此内容 |   |
| `cardPay` | java.lang.Boolean | 是 | 是否银行卡支付 |   |
| `expressPay` | java.lang.Boolean | 是 | 是否快捷支付 |   |
| `payWayDesc` | String | 是 | 支付方式 | 支付宝 |

<a id="m-alibaba-openplatform-trade-model-productiteminfo[]"></a>
#### alibaba.openplatform.trade.model.ProductItemInfo[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `cargoNumber` | java.lang.String | 是 | 指定单品货号，国际站无需关注。该字段不一定有值，仅仅在下单时才会把货号记录(如果卖家设置了单品货号的话)。别的订单类型的货号只能通过商品接口去获取。请注意：通过商品接口获取时的货号和下单时的货号可能不一致，因为下单完成后卖家可能修改商品信息，改变了货号。 |   |
| `description` | java.lang.String | 是 | 描述,1688无此信息 |   |
| `itemAmount` | java.math.BigDecimal | 是 | 实付金额，单位为元 |   |
| `name` | java.lang.String | 是 | 商品名称 |   |
| `price` | java.math.BigDecimal | 是 | 原始单价，以元为单位 |   |
| `productID` | java.lang.Long | 是 | 产品ID（非在线产品为空） |   |
| `productImgUrl` | String[] | 是 | 商品图片url |   |
| `productSnapshotUrl` | java.lang.String | 是 | 产品快照url，交易订单产生时会自动记录下当时的商品快照，供后续纠纷时参考 |   |
| `quantity` | java.math.BigDecimal | 是 | 以unit为单位的数量，例如多少个、多少件、多少箱、多少吨 |   |
| `refund` | java.math.BigDecimal | 是 | 退款金额，单位为元 |   |
| `skuID` | java.lang.Long | 是 | skuID |   |
| `sort` | java.lang.Integer | 是 | 排序字段，商品列表按此字段进行排序，从0开始，1688不提供 |   |
| `status` | java.lang.String | 是 | 子订单状态 |   |
| `subItemID` | java.lang.Long | 是 | 子订单号，或商品明细条目ID |   |
| `type` | java.lang.String | 是 | 类型，国际站使用，供卖家标注商品所属类型 |   |
| `unit` | java.lang.String | 是 | 售卖单位	例如：个、件、箱、吨 |   |
| `weight` | java.lang.String | 是 | 重量	按重量单位计算的重量，例如：100 |   |
| `weightUnit` | java.lang.String | 是 | 重量单位	例如：g，kg，t |   |
| `guaranteesTerms` | [message:alibaba.openplatform.trade.model.GuaranteeTermsInfo[]](#m-alibaba-openplatform-trade-model-guaranteetermsinfo[]) | 是 | 保障条款，此字段仅针对1688 |   |
| `productCargoNumber` | java.lang.String | 是 | 指定商品货号，该字段不一定有值，在下单时才会把货号记录。别的订单类型的货号只能通过商品接口去获取。请注意：通过商品接口获取时的货号和下单时的货号可能不一致，因为下单完成后卖家可能修改商品信息，改变了货号。该字段和cargoNUmber的区别是：该字段是定义在商品级别上的货号，cargoNUmber是定义在单品级别的货号 |   |
| `skuInfos` | [message:alibaba.trade.SkuItemDesc[]](#m-alibaba-trade-skuitemdesc[]) | 是 |  |    |
| `entryDiscount` | Long | 是 | 订单明细涨价或降价的金额 |   |
| `specId` | java.lang.String | 是 | 订单销售属性ID |   |
| `quantityFactor` | java.math.BigDecimal | 是 | 以unit为单位的quantity精度系数，值为10的幂次，例如:quantityFactor=1000,unit=吨，那么quantity的最小精度为0.001吨 |   |
| `statusStr` | java.lang.String | 是 | 子订单状态描述 |   |
| `refundStatus` | java.lang.String | 是 | WAIT_SELLER_AGREE 等待卖家同意<br>REFUND_SUCCESS 退款成功<br>REFUND_CLOSED 退款关闭<br>WAIT_BUYER_MODIFY 待买家修改<br>WAIT_BUYER_SEND 等待买家退货<br>WAIT_SELLER_RECEIVE 等待卖家确认收货 |   |
| `closeReason` | java.lang.String | 是 | 关闭原因 |   |
| `logisticsStatus` | java.lang.Integer | 是 | 1 未发货<br>2 已发货<br>3 已收货<br>4 已经退货<br>5 部分发货<br>8 还未创建物流订单 |   |
| `gmtCreate` | java.util.Date | 是 | 创建时间 |   |
| `gmtModified` | java.util.Date | 是 | 修改时间 |   |
| `gmtCompleted` | java.util.Date | 是 | 明细完成时间 |   |
| `gmtPayExpireTime` | String | 是 | 库存超时时间，格式为“yyyy-MM-dd HH:mm:ss” |   |
| `refundId` | String | 是 | 售中退款单号 |   |
| `subItemIDString` | String | 是 | 子订单号，或商品明细条目ID(字符串类型，由于Long类型的ID可能在JS和PHP中处理有问题，所以可以用字符串类型来处理) |   |
| `refundIdForAs` | String | 是 | 售后退款单号 |   |
| `sharePostage` | BigDecimal | 是 | 分担邮费，单位为元 |   |

<a id="m-alibaba-openplatform-trade-model-guaranteetermsinfo[]"></a>
#### alibaba.openplatform.trade.model.GuaranteeTermsInfo[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `assuranceInfo` | java.lang.String | 是 | 保障条款 | 自愿选择向买家提供“交期保障”服务 |
| `assuranceType` | java.lang.String | 是 | 保障方式。国际站：TA(信保) | jqbz |
| `qualityAssuranceType` | java.lang.String | 是 | 质量保证类型。国际站：pre_shipment(发货前),post_delivery(发货后) | 交期保障 |
| `value` | String | 是 | 保障条款值，比如交期保障里，6表示6天 | 6 |

<a id="m-alibaba-trade-skuitemdesc[]"></a>
#### alibaba.trade.SkuItemDesc[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `name` | String | 是 | 属性名 |  |
| `value` | String | 是 | 属性值 |  |

<a id="m-alibaba-openplatform-trade-model-nativelogisticsinfo"></a>
#### alibaba.openplatform.trade.model.NativeLogisticsInfo

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `address` | java.lang.String | 是 | 详细地址 |  |
| `area` | java.lang.String | 是 | 县，区 |  |
| `areaCode` | java.lang.String | 是 | 省市区编码 |  |
| `city` | java.lang.String | 是 | 城市 |  |
| `contactPerson` | java.lang.String | 是 | 联系人姓名 |  |
| `fax` | java.lang.String | 是 | 传真 |  |
| `mobile` | java.lang.String | 是 | 手机 |  |
| `province` | java.lang.String | 是 | 省份 |  |
| `telephone` | java.lang.String | 是 | 电话 |  |
| `zip` | java.lang.String | 是 | 邮编 |  |
| `logisticsItems` | [message:alibaba.openplatform.trade.model.NativeLogisticsItemsInfo[]](#m-alibaba-openplatform-trade-model-nativelogisticsitemsinfo[]) | 是 | 运单明细 |  |
| `townCode` | java.lang.String | 是 | 镇，街道地址码 |  |
| `town` | java.lang.String | 是 | 镇，街道 |  |

<a id="m-alibaba-openplatform-trade-model-nativelogisticsitemsinfo[]"></a>
#### alibaba.openplatform.trade.model.NativeLogisticsItemsInfo[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `deliveredTime` | java.util.Date | 是 | 发货时间 |   |
| `logisticsCode` | java.lang.String | 是 | 物流编号 |   |
| `type` | java.lang.String | 是 | SELF_SEND_GOODS(&quot;0&quot;)自行发货，在线发货ONLINE_SEND_GOODS(&quot;1&quot;，不需要物流的发货 NO_LOGISTICS_SEND_GOODS(&quot;2&quot;) |   |
| `id` | java.lang.Long | 是 | 主键id |   |
| `status` | java.lang.String | 是 | 状态 |   |
| `gmtModified` | java.util.Date | 是 | 修改时间 |   |
| `gmtCreate` | java.util.Date | 是 | 创建时间 |   |
| `carriage` | java.math.BigDecimal | 是 | 运费(单位为元) |   |
| `fromProvince` | java.lang.String | 是 | 发货省 |   |
| `fromCity` | java.lang.String | 是 | 发货市 |   |
| `fromArea` | java.lang.String | 是 | 发货区 |   |
| `fromAddress` | java.lang.String | 是 | 发货街道地址 |   |
| `fromPhone` | java.lang.String | 是 | 发货联系电话 |   |
| `fromMobile` | java.lang.String | 是 | 发货联系手机 |   |
| `fromPost` | java.lang.String | 是 | 发货地址邮编 |   |
| `logisticsCompanyId` | java.lang.Long | 是 | 物流公司Id |   |
| `logisticsCompanyNo` | java.lang.String | 是 | 物流公司编号 |   |
| `logisticsCompanyName` | java.lang.String | 是 | 物流公司名称 |   |
| `logisticsBillNo` | java.lang.String | 是 | 物流公司运单号 |   |
| `subItemIds` | java.lang.String | 是 | 商品明细条目id，如有多个以,分隔 |   |
| `toProvince` | java.lang.String | 是 | 收货省 |   |
| `toCity` | java.lang.String | 是 | 收货市 |   |
| `toArea` | java.lang.String | 是 | 收货区 |   |
| `toAddress` | java.lang.String | 是 | 收货街道地址 |   |
| `toPhone` | java.lang.String | 是 | 收货联系电话 |   |
| `toMobile` | java.lang.String | 是 | 收货联系手机 |   |
| `toPost` | java.lang.String | 是 | 收货地址邮编 |   |
| `noLogisticsName` | String | 是 | 物流姓名 |   |
| `noLogisticsTel` | String | 是 | 联系方式 |   |
| `noLogisticsBillNo` | String | 是 | 无需物流业务单号 |   |
| `noLogisticsCondition` | String | 是 | 无需物流类别,noLogisticsCondition=1， 表示其他第三方物流、小型物充商、车队等, noLogisticsCondition=2 表示补运费、差价, noLogisticsCondition=3 表示卖家配送, noLogisticsCondition=4 表示买家自提 noLogisticsCondition=5 表示其他原因 |   |

<a id="m-alibaba-invoice-orderinvoicemodel"></a>
#### alibaba.invoice.OrderInvoiceModel

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `invoiceCompanyName` | java.lang.String | 是 | 发票公司名称(即发票抬头-title) |   |
| `invoiceType` | java.lang.Integer | 是 | 发票类型. 0：普通发票，1:增值税发票，9未知类型 |   |
| `localInvoiceId` | java.lang.Long | 是 | 本地发票号 |   |
| `orderId` | java.lang.Long | 是 | 订单Id |   |
| `receiveCode` | java.lang.String | 是 | (收件人)址区域编码 |   |
| `receiveCodeText` | java.lang.String | 是 | (收件人) 省市区编码对应的文案(增值税发票信息) |   |
| `receiveMobile` | java.lang.String | 是 | （收件者）发票收货人手机 |   |
| `receiveName` | java.lang.String | 是 | （收件者）发票收货人 |   |
| `receivePhone` | java.lang.String | 是 | （收件者）发票收货人电话 |   |
| `receivePost` | java.lang.String | 是 | （收件者）发票收货地址邮编 |   |
| `receiveStreet` | java.lang.String | 是 | (收件人) 街道地址(增值税发票信息) |   |
| `registerAccountId` | java.lang.String | 是 | (公司)银行账号 |   |
| `registerBank` | java.lang.String | 是 | (公司)开户银行 |   |
| `registerCode` | java.lang.String | 是 | (注册)省市区编码 |   |
| `registerCodeText` | java.lang.String | 是 | (注册)省市区文本 |   |
| `registerPhone` | java.lang.String | 是 | （公司）注册电话 |   |
| `registerStreet` | java.lang.String | 是 | (注册)街道地址 |   |
| `taxpayerIdentify` | java.lang.String | 是 | 纳税人识别号 |   |
| `amount` | Long | 是 | 发票含税总额（分） | 1 |
| `status` | String | 是 | 发票状态：{&quot;ISSUED&quot;:&quot;已开票&quot;,&quot;RED_ISSUING&quot;:&quot;冲红中&quot;,&quot;RED_ALL_ISSUED&quot;:&quot;全部冲红&quot;,&quot;CLOSED&quot;:&quot;关闭&quot;,&quot;RED_PART_ISSUED&quot;:&quot;部分冲红&quot;,&quot;VERIFYING&quot;:&quot;验票中&quot;,&quot;DEPRECATED&quot;:&quot;已作废&quot;,&quot;RETURNING&quot;:&quot;退票中&quot;,&quot;INVALIDING&quot;:&quot;作废中&quot;,&quot;INIT&quot;:&quot;初始化&quot;,&quot;VERIFY_FAILED&quot;:&quot;验票失败&quot;,&quot;ISSUING&quot;:&quot;开票中&quot;,&quot;FAILED&quot;:&quot;开票失败&quot;,&quot;RETURNED&quot;:&quot;已退票&quot;} | issued |

<a id="m-alibaba-openplatform-trade-model-guaranteetermsinfo"></a>
#### alibaba.openplatform.trade.model.GuaranteeTermsInfo

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `assuranceInfo` | java.lang.String | 是 | 保障条款 | 自愿选择向买家提供“交期保障”服务 |
| `assuranceType` | java.lang.String | 是 | 保障方式。国际站：TA(信保) | jqbz |
| `qualityAssuranceType` | java.lang.String | 是 | 质量保证类型。国际站：pre_shipment(发货前),post_delivery(发货后) | 交期保障 |
| `value` | String | 是 | 保障条款值，比如交期保障里，6表示6天 | 6 |

<a id="m-alibaba-trade-orderrateinfo"></a>
#### alibaba.trade.OrderRateInfo

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `buyerRateStatus` | Integer | 是 | 买家评价状态(4:已评论,5:未评论,6;不需要评论) |  |
| `sellerRateStatus` | Integer | 是 | 卖家评价状态(4:已评论,5:未评论,6;不需要评论) |  |
| `buyerRateList` | [message:alibaba.order.rateDetail[]](#m-alibaba-order-ratedetail[]) | 是 | 卖家給买家的评价 |  |
| `sellerRateList` | [message:alibaba.order.rateDetail[]](#m-alibaba-order-ratedetail[]) | 是 | 买家給卖家的评价 |  |

<a id="m-alibaba-order-ratedetail[]"></a>
#### alibaba.order.rateDetail[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `starLevel` | java.lang.Integer | 是 | 评价星级 |  |
| `content` | java.lang.String | 是 | 评价详情 |  |
| `receiverNick` | java.lang.String | 是 | 收到评价的用户昵称 |  |
| `posterNick` | java.lang.String | 是 | 发送评价的用户昵称 |  |
| `publishTime` | java.util.Date | 是 | 评价上线时间 |  |

<a id="m-alibaba-trade-overseasextraaddress"></a>
#### alibaba.trade.OverseasExtraAddress

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `channelName` | String | 是 | 路线名称 | 欧洲小包 |
| `channelId` | String | 是 | 路线id | 1 |
| `shippingCompanyId` | String | 是 | 货代公司id | 222 |
| `shippingCompanyName` | String | 是 | 货代公司名称 | 货代公司1 |
| `countryCode` | String | 是 | 国家code | UK |
| `country` | String | 是 | 国家 | 英国 |
| `email` | String | 是 | 买家邮箱 | aaa@gmail.com |

<a id="m-alibaba-trade-customs"></a>
#### alibaba.trade.Customs

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `id` | Long | 是 | id | 1 |
| `gmtCreate` | java.util.Date | 是 | 创建时间 | 20170806114526000+0800 |
| `gmtModified` | java.util.Date | 是 | 修改时间 | 20170806114526000+0800 |
| `buyerId` | Long | 是 | 买家id | 123456 |
| `orderId` | String | 是 | 主订单id | 12312312312312 |
| `type` | Integer | 是 | 业务数据类型,默认1：报关单 | 1 |
| `attributes` | [message:alibaba.trade.CustomsAttributesInfo[]](#m-alibaba-trade-customsattributesinfo[]) | 是 | 报关信息列表 |  |

<a id="m-alibaba-trade-customsattributesinfo[]"></a>
#### alibaba.trade.CustomsAttributesInfo[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `sku` | String | 是 | sku标识 | 1234 |
| `cName` | String | 是 | 中文名称 | 测试 |
| `enName` | String | 是 | 英文名称 | test |
| `amount` | Double | 是 | 申报价值 | 3000.0 |
| `quantity` | Double | 是 | 数量 | 1.0 |
| `weight` | Double | 是 | 重量（kg） | 0.5 |
| `currency` | String | 是 | 报关币种 | CNY |

<a id="m-alibaba-orderdetail-caigouquoteinfo[]"></a>
#### alibaba.orderDetail.caigouQuoteInfo[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `productQuoteName` | String | 是 | 供应单项的名称 | 物料01 |
| `price` | java.math.BigDecimal | 是 | 价格，单位：元 | 100 |
| `count` | Double | 是 | 购买数量 | 10 |

<a id="m-alibaba-openplatform-trade-keyvaluepair[]"></a>
#### alibaba.openplatform.trade.KeyValuePair[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `key` | String | 是 | 键 |  |
| `value` | String | 是 | 值 |  |
| `description` | String | 是 | 描述 |  |

<a id="m-alibaba-trade-get-sellerview-tradeinfo-encryptoutorderinfo"></a>
#### alibaba.trade.get.sellerView.tradeinfo.EncryptOutOrderInfo

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `outPlatformOrderNo` | String | 是 | 外部订单号 | 25662254541 |
| `outPlatformCode` | String | 是 | 下游平台,淘宝-thyny，天猫-tm，淘特-taote，阿里巴巴C2M-c2m，京东-jingdong，拼多多-pinduoduo，微信-weixin，跨境-kuajing，快手-kuaishou，有赞-youzan，抖音-douyin，寺库-siku，美团团好货-meituan，小红书-xiaohongshu，当当-dangdang，苏宁-suning，大V店-davdian，行云-xingyun，蜜芽-miya，菠萝派商城-boluo，其他-other | tm |
| `outPlatformAppkey` | String | 是 | 获取下游订单信息的下游平台的appkey | 65345 |
| `oaid` | String | 是 | 淘宝oaid | xxx-xxxx-xxx |
| `encryptReceiverName` | String | 是 | 下游加密收货人姓名 | *** |
| `encryptReceiverMobile` | String | 是 | 下游加密收货人电话 | *** |
| `encryptReceiverAddress` | String | 是 | 下游加密收货人地址 | *** |
| `outPatformExtraInfo` | String | 是 | 其他扩展信息 | {} |
| `outShopId` | String | 是 | 下游平台shopId | 21343123 |
| `outOriginAddress` | [message:com.alibaba.ocean.openplatform.biz.trade.param.OutAddress](#m-com-alibaba-ocean-openplatform-biz-trade-param-outaddress) | 是 | 外部原始地址信息 | {} |
| `outShopName` | String | 是 | 下游平台店铺名称 | 三生智能 |
| `outPlatformSubCode` | String | 是 | 下游渠道子业务编码，比如抖音101子渠道，用于供应链订单识别,并用于加密订单打单 | 101 |
| `outPlatformSupplyOrderNo` | String | 是 | 下游平台供应链采购订单号，如抖音供应链等 | CT7403712218870825259 |
| `outSupplierId` | String | 是 | 下游平台供应链提供商id | 1231123 |

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

<a id="m-com-alibaba-ocean-openplatform-biz-trade-common-model-oversealogisticsinfo"></a>
#### com.alibaba.ocean.openplatform.biz.trade.common.model.OverseaLogisticsInfo

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `overseasTransportUserAddrInfo` | [message:com.alibaba.ocean.openplatform.biz.cross.model.OverseasTransportUserAddrInfo](#m-com-alibaba-ocean-openplatform-biz-cross-model-overseastransportuseraddrinfo) | 是 | 海外收货地址 | "addressDetail":"test", 						"cityName":"Г.Курчатов", 						"areaName":" ", 						"warehouseContactName":"Zeus", 						"mobile":"123456789", 						"fixedPhone":"123456789", 						"postCode":"1111", 						"provinceName":"Абайская область", 						"countryName":"Казахстан", |
| `overseasLogisticsIds` | String[] | 是 | 海外物流单号 | ["TN0003077L" ] |

<a id="m-com-alibaba-ocean-openplatform-biz-cross-model-overseastransportuseraddrinfo"></a>
#### com.alibaba.ocean.openplatform.biz.cross.model.OverseasTransportUserAddrInfo

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `postCode` | String | 是 | 物流码 | 123 |
| `countryName` | String | 是 | 国家 | 越南 |
| `provinceName` | String | 是 | 省份 | xx |
| `cityName` | String | 是 | 城市 | x x |
| `areaName` | String | 是 | 区 | xx |
| `addressDetail` | String | 是 | 地址 | xx |
| `mobile` | String | 是 | 手机号 | 123 |
| `fixedPhone` | String | 是 | 手机号 | 123 |
| `warehouseContactName` | String | 是 | 仓库联系人 | aa |
| `warehouseName` | String | 是 | 仓库名称 | aa |

<a id="m-com-alibaba-ocean-openplatform-biz-trade-common-model-opsellerinvoicetradesettingmodel"></a>
#### com.alibaba.ocean.openplatform.biz.trade.common.model.OpSellerInvoiceTradeSettingModel

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `tradeInvoiceStatus` | String | 是 | 交易开票状态，-1:不可开票；0-允许开票；1-已申请；2-已开票 | 1 |
| `sellerInvoiceType` | String | 是 | 开票类型，0-普票；1专票；2-普票/专票 | 0 |

## 示例

**订单详情出参示例**

```
{"result":{"baseInfo":{"allDeliveredTime":"20170913231916000-0700","businessType":"cn","buyerID":"b2b-2248544159","createTime":"20170913231708000-0700","id":58218860983545944,"modifyTime":"20170913234725000-0700","payTime":"20170913231727000-0700","refund":0,"sellerID":"b2b-2248564064","shippingFee":6,"status":"waitbuyerreceive","totalAmount":6.15,"discount":0,"buyerContact":{"phone":"86-0571-81895955","imInPlatform":"b测试账号002","name":"乔的石","mobile":"18668184036","companyName":"阿里巴巴网络科技有限公司"},"sellerContact":{"phone":"86-0571-88881888","email":"caigouwfw_1688@163.com","imInPlatform":"b测试账号110","name":"孟舒","mobile":"13312919596","companyName":"阿里巴巴网络有限公司"},"tradeType":"50060","refundStatus":"waitselleragree","refundPayment":0,"idOfStr":"58218860983545941","alipayTradeId":"2017091421001008480237437679","receiverInfo":{"toFullName":"童恩杰","toDivisionCode":"330108","toPost":"312000","toTownCode":"330108002","toArea":"浙江省 杭州市 滨江区 长河街道"},"buyerLoginId":"b测试账号002","sellerLoginId":"b测试账号110","buyerUserId":2248544159,"sellerUserId":2248564064,"buyerAlipayId":"2088611489970483","sellerAlipayId":"2088611383470360","sumProductPayment":0.3,"stepPayAll":false},"nativeLogistics":{"address":"杭州市滨江区网商路699号","area":"滨江区","areaCode":"330108","city":"杭州市","contactPerson":"童恩杰","mobile":"13666836263","province":"浙江省","zip":"312000","logisticsItems":[{"deliveredTime":"20170913231917000-0700","logisticsCode":"BX107450035961376","type":"2","id":107450035961376,"status":"alreadysend","gmtModified":"20170913231916000-0700","gmtCreate":"20170913231916000-0700","fromPhone":"86-0571-88881888","fromMobile":"13312919596","logisticsBillNo":"不需要物流","subItemIds":"58218860984545941,58218860985545941,58218860986545941"}],"townCode":"330108002","town":"长河街道"},"productItems":[{"itemAmount":0.3,"name":"测试扫码购富光勿拍2L*6件","price":0.3,"productID":547486647009,"productImgUrl":["http://cbu01.alicdn.com/img/ibank/2017/757/125/3950521757.80x80.jpg","http://cbu01.alicdn.com/img/ibank/2017/757/125/3950521757.jpg"],"productSnapshotUrl":"https://trade.1688.com/order/offer_snapshot.htm?order_entry_id=128403042259997715","quantity":1,"refund":0,"skuID":3315536521048,"status":"waitsellersend","subItemID":128403042259997710,"type":"common","unit":"箱","guaranteesTerms":[],"skuInfos":[{"name":"颜色","value":"白色"},{"name":"容量","value":"1.1L-2L"}],"entryDiscount":0,"specId":"28b952ec96c8b5c3ab0affc1b74923f0","quantityFactor":1,"statusStr":"等待卖家发货","refundStatus":"WAIT_SELLER_AGREE","logisticsStatus":1,"gmtCreate":"20180206152758000+0800","gmtModified":"20180206153325000+0800","gmtPayExpireTime":"2018-02-07 15:27:58","refundId":"TQ7257793022991577","subItemIDString":"128403042259997715"},{"itemAmount":0.05,"name":"这个一个很好看好看的鞋子用于服务测试（大家不要动）","price":0.1,"productID":558700975520,"productImgUrl":[],"productSnapshotUrl":"https://trade.1688.com/order/offer_snapshot.htm?order_entry_id=58218860985545941","quantity":1,"refund":0,"skuID":3638916762844,"status":"waitbuyerreceive","subItemID":58218860985545944,"type":"common","unit":"双","guaranteesTerms":[],"skuInfos":[{"name":"颜色","value":"白色"},{"name":"尺码","value":"38"}],"entryDiscount":0,"specId":"6ff6071792c8ab520b9b867c61b990bd","quantityFactor":1,"statusStr":"等待买家收货","refundStatus":"WAIT_SELLER_AGREE","logisticsStatus":2},{"itemAmount":0.05,"name":"这个一个很好看好看的鞋子用于服务测试（大家不要动）","price":0.1,"productID":558700975520,"productImgUrl":[],"productSnapshotUrl":"https://trade.1688.com/order/offer_snapshot.htm?order_entry_id=58218860986545941","quantity":1,"refund":0,"skuID":3638916762843,"status":"waitbuyerreceive","subItemID":58218860986545944,"type":"common","unit":"双","guaranteesTerms":[],"skuInfos":[{"name":"颜色","value":"白色"},{"name":"尺码","value":"34"}],"entryDiscount":0,"specId":"8ecd5cb401df85d3f2485eb120f670df","quantityFactor":1,"statusStr":"等待买家收货","refundStatus":"WAIT_SELLER_AGREE","logisticsStatus":2}],"tradeTerms":[{"payStatus":"2","payTime":"20170913231727000-0700","payWay":"1","phasAmount":6.15,"phase":655062941545941}],"extAttributes":[],"orderRateInfo":{"buyerRateStatus":5,"sellerRateStatus":5}}}
```
