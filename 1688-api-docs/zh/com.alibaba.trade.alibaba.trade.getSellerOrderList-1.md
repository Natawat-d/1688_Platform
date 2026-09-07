# 订单列表查看(卖家视角)

API: `com.alibaba.trade:alibaba.trade.getSellerOrderList:1` · Category: 消息  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.trade.getSellerOrderList-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.trade.getSellerOrderList/{appKey}`  
需要授权 (access_token) · 需要签名

获取卖家订单列表，也就是用户的memberId必须等于订单的sellerMemberId。该接口仅仅返回订单基本信息，不会返回订单的物流信息和发票信息；如果需要获取物流信息，请调用获取订单详情接口；如果需要获取发票信息，请调用获取发票信息的API

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `createStartTime` | java.util.Date | 否 | 下单开始时间 | 20180721172608000+0800 |
| `createEndTime` | java.util.Date | 否 | 下单结束时间 | 20180721172608000+0800 |
| `modifyStartTime` | java.util.Date | 否 | 查询修改时间开始 | 20180721172608000+0800 |
| `modifyEndTime` | java.util.Date | 否 | 查询修改时间结束 | 20180721172608000+0800 |
| `page` | int | 否 | 查询分页页码，从1开始 | 1 |
| `pageSize` | int | 否 | 查询的每页的数量(最高20) | 10 |
| `orderStatus` | java.lang.String | 否 | 订单状态，值有success,cancel(交易取消，违约金等交割完毕),waitbuyerpay(等待买家付款)，waitsellersend(等待卖家发货),waitbuyerreceive(等待买家收货) | waitbuyerpay |
| `refundStatus` | java.lang.String | 否 | 退款状态，支持：<br>&quot;waitselleragree&quot;(等待卖家同意),<br>&quot;refundsuccess&quot;(退款成功),<br>&quot;refundclose&quot;(退款关闭),<br>&quot;waitbuyermodify&quot;(待买家修改),<br>&quot;waitbuyersend&quot;(等待买家退货),<br>&quot;waitsellerreceive&quot;(等待卖家确认收货) | waitselleragree |
| `buyerMemberId` | java.lang.String | 否 | 买家memberId或者buyerOpenUid（买家加密ID） | b2b-1234325 |
| `buyerLoginId` | String | 否 | 买家LoginId或者buyerOpenUid（买家加密ID） | alitestforisv02 |
| `tradeType` | java.lang.String | 否 | 交易类型:<br>担保交易(1),<br>预存款交易(2),<br>ETC境外收单交易(3),<br>即时到帐交易(4),<br>保障金安全交易(5),<br>统一交易流程(6),<br>分阶段交易(7),<br>货到付款交易(8),<br>信用凭证支付交易(9),<br>账期支付交易(10),<br>1688交易4.0，新分阶段交易(50060),<br>当面付的交易流程(50070),<br>服务类的交易流程(50080) | 5 |
| `bizTypes` | String[] | 否 | 业务类型，支持： &quot;cn&quot;(普通订单类型), &quot;ws&quot;(大额批发订单类型), &quot;yp&quot;(普通拿样订单类型), &quot;yf&quot;(一分钱拿样订单类型), &quot;fs&quot;(倒批(限时折扣)订单类型), &quot;cz&quot;(加工定制订单类型), &quot;ag&quot;(协议采购订单类型), &quot;hp&quot;(伙拼订单类型), &quot;gc&quot;(国采订单类型), &quot;supply&quot;(供销订单类型), &quot;nyg&quot;(nyg订单类型), &quot;factory&quot;(淘工厂订单类型), &quot;quick&quot;(快订下单), &quot;xiangpin&quot;(享拼订单), &quot;nest&quot;(采购商城-鸟巢), &quot;f2f&quot;(当面付), &quot;cyfw&quot;(存样服务), &quot;sp&quot;(代销订单标记), &quot;wg&quot;(微供订单), &quot;factorysamp&quot;(淘工厂打样订单), &quot;factorybig&quot;(淘工厂大货订单) | ["cn","ws"] |
| `isHis` | boolean | 否 | 是否查询历史订单表,默认查询当前表 | false |
| `productName` | java.lang.String | 否 | 商品名称 | 测试商品名 |
| `needBuyerAddressAndPhone` | java.lang.Boolean | 否 | 是否需要查询买家的详细地址信息和电话 | false |
| `needMemoInfo` | java.lang.Boolean | 否 | 是否需要查询备注信息 | false |
| `tousuStatus` | boolean | 否 | 是否查找投诉中的地拟改单 | false |
| `buyerRateStatus` | java.lang.Integer | 否 | 买家评价状态 (4:已评价,5:未评价,6;不需要评价) | 5 |
| `sellerRateStatus` | java.lang.Integer | 否 | 卖家评价状态 (4:已评价,5:未评价,6;不需要评价) | 5 |
| `needCheckSend` | Boolean | 否 | 是否需要返回发货校验信息，对于打单发货场景必须要填，非常重要 | true |
| `needSendGoodsOverdueRisk` | Boolean | 否 | 是否需要订单发货超期风险标 | false |
| `needDeliverGoodsOverdueRisk` | Boolean | 否 | 是否需要揽收超期风险标 | false |
| `needOfficialLogisticOrder` | Boolean | 否 | 是否需要筛选官方直送保障服务的订单 | false |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:alibaba.openplatform.trade.model.TradeInfo[]](#m-alibaba-openplatform-trade-model-tradeinfo[]) | 是 | 查询返回结果 | {} |
| `errorCode` | String | 是 | 错误码 |   |
| `errorMessage` | String | 是 | 错误信息 |   |
| `totalRecord` | Long | 是 | 总记录数 | 101 |
| `success` | Boolean | 是 | 是否调用成功 | true |
| `retCodes` | String[] | 是 | 脱敏信息码 | ["xxx","xxx"] |

<a id="m-alibaba-openplatform-trade-model-tradeinfo[]"></a>
#### alibaba.openplatform.trade.model.TradeInfo[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `baseInfo` | [message:alibaba.openplatform.trade.model.OrderBaseInfo](#m-alibaba-openplatform-trade-model-orderbaseinfo) | 是 | 订单基础信息 | [] |
| `nativeLogistics` | [message:alibaba.openplatform.trade.model.NativeLogisticsInfo](#m-alibaba-openplatform-trade-model-nativelogisticsinfo) | 是 | 国内物流 | {} |
| `overseasExtraAddress` | [message:alibaba.trade.OverseasExtraAddress](#m-alibaba-trade-overseasextraaddress) | 是 | 跨境地址扩展信息 | {} |
| `productItems` | [message:alibaba.openplatform.trade.model.ProductItemInfo[]](#m-alibaba-openplatform-trade-model-productiteminfo[]) | 是 | 商品条目信息 | [] |
| `customs` | [message:alibaba.trade.Customs](#m-alibaba-trade-customs) | 是 | 跨境报关信息 | {} |
| `tradeTerms` | [message:alibaba.openplatform.trade.model.TradeTermsInfo[]](#m-alibaba-openplatform-trade-model-tradetermsinfo[]) | 是 | 交易条款 | [] |
| `orderRateInfo` | [message:alibaba.trade.OrderRateInfo](#m-alibaba-trade-orderrateinfo) | 是 | 订单评价信息 | {} |
| `orderInvoiceInfo` | [message:alibaba.invoice.OrderInvoiceModel](#m-alibaba-invoice-orderinvoicemodel) | 是 | 发票信息 | {} |
| `guaranteesTerms` | [message:alibaba.openplatform.trade.model.GuaranteeTermsInfo](#m-alibaba-openplatform-trade-model-guaranteetermsinfo) | 是 | 保障条款 | {} |
| `extAttributes` | [message:alibaba.openplatform.trade.KeyValuePair[]](#m-alibaba-openplatform-trade-keyvaluepair[]) | 是 | 订单扩展属性 | [] |
| `fromEncryptOrder` | Boolean | 是 | 是否下游脱敏信息创建的订单 | true |
| `sendGoodsOverdueRisk` | String | 是 | 订单发货超期风险标，PTMO:潜在超时；TMOT:已超期 | TMOT |
| `officialLogisticOrder` | Boolean | 是 | 为true时，表示官方直送保障服务订单，为空或者false时，非官方直送保障服务订单 | null |

<a id="m-alibaba-openplatform-trade-model-orderbaseinfo"></a>
#### alibaba.openplatform.trade.model.OrderBaseInfo

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `allDeliveredTime` | java.util.Date | 是 | 完全发货时间 | 20180614101942000+0800 |
| `payTime` | java.util.Date | 是 | 付款时间，如果有多次付款，这里返回的是首次付款时间 | 20180614101942000+0800 |
| `buyerRemarkIcon` | String | 是 | 买家备忘标志 | 1 |
| `receiverInfo` | [message:alibaba.trade.orderReceiverInfo](#m-alibaba-trade-orderreceiverinfo) | 是 | 收件人信息 | {} |
| `discount` | Long | 是 | 折扣信息，单位分 | 11 |
| `refundStatus` | String | 是 | 订单的售中退款状态，等待卖家同意：waitselleragree ，待买家修改：waitbuyermodify，等待买家退货：waitbuyersend，等待卖家确认收货：waitsellerreceive，退款成功：refundsuccess，退款失败：refundclose | refundclose |
| `alipayTradeId` | java.lang.String | 是 | 外部支付交易Id | 123123121111 |
| `remark` | java.lang.String | 是 | 备注，1688指下单时的备注 | 备注 |
| `sumProductPayment` | java.math.BigDecimal | 是 | 产品总金额(该订单产品明细表中的产品金额的和)，单位元 | 1212 |
| `buyerFeedback` | java.lang.String | 是 | 买家留言，不超过500字 | 留言 |
| `flowTemplateCode` | String | 是 | 4.0交易流程模板code | flow |
| `sellerOrder` | java.lang.Boolean | 是 | 是否自主订单（邀约订单） | false |
| `preOrderId` | java.lang.Long | 是 | 预订单ID | 123123 |
| `buyerLoginId` | java.lang.String | 是 | 买家loginId，旺旺Id | alitestforusv01 |
| `modifyTime` | java.util.Date | 是 | 修改时间 | 20180614101942000+0800 |
| `subBuyerLoginId` | String | 是 | 买家子账号 | alitestforusv02:temp |
| `confirmedTime` | java.util.Date | 是 | 确认时间 | 20180614101942000+0800 |
| `currency` | java.lang.String | 是 | 币种，币种，整个交易单使用同一个币种。值范围：USD,RMB,HKD,GBP,CAD,AUD,JPY,KRW,EUR | EUR |
| `id` | java.lang.Long | 是 | 交易id | 1231231231111 |
| `closeReason` | java.lang.String | 是 | 关闭原因。buyerCancel:买家取消订单，sellerGoodsLack:卖家库存不足，other:其它 | buyerCancel |
| `tradeType` | String | 是 | 1:担保交易<br>2:预存款交易<br>3:ETC境外收单交易<br>4:即时到帐交易<br>5:保障金安全交易<br>6:统一交易流程<br>7:分阶段付款<br>8.货到付款交易<br>9.信用凭证支付交易<br>10.账期支付交易，50060 交易4.0 | 50060 |
| `ccid` | String | 是 | 联系人信息解密ID，用于电商平台联系人信息加密场景使用，非订单加密场景请勿使用。 | xxx |
| `buyerContact` | [message:alibaba.trade.tradeContact](#m-alibaba-trade-tradecontact) | 是 | 买家联系人 | {} |
| `receivingTime` | java.util.Date | 是 | 收货时间，这里返回的是完全收货时间 | 20180614101942000+0800 |
| `stepAgreementPath` | java.lang.String | 是 | 分阶段法务协议地址 |   |
| `idOfStr` | String | 是 | 交易id(字符串格式) | 123121212123 |
| `refundStatusForAs` | String | 是 | 订单的售后退款状态 |   |
| `stepPayAll` | java.lang.Boolean | 是 | 是否一次性付款 | false |
| `completeTime` | java.util.Date | 是 | 完成时间 | 20180614101942000+0800 |
| `sellerLoginId` | java.lang.String | 是 | 卖家oginId，旺旺Id | alitestforusv02 |
| `buyerID` | java.lang.String | 是 | 买家主账号id | 1234531 |
| `stepOrderList` | [message:alibaba.trade.StepOrderModel[]](#m-alibaba-trade-stepordermodel[]) | 是 | [交易3.0]分阶段交易，分阶段订单list |   |
| `totalAmount` | java.math.BigDecimal | 是 | 应付款总金额，totalAmount = ∑itemAmount + shippingFee，单位为元 | 1000 |
| `sellerID` | java.lang.String | 是 | 卖家主账号id | 123123123123 |
| `shippingFee` | java.math.BigDecimal | 是 | 运费，单位为元 | 1 |
| `createTime` | java.util.Date | 是 | 创建时间 | 20180614101942000+0800 |
| `sellerRemarkIcon` | String | 是 | 卖家备忘标志 | 1 |
| `sellerMemo` | java.lang.String | 是 | 卖家备忘信息 | 备忘 |
| `businessType` | java.lang.String | 是 | 业务类型。国际站：ta(信保),wholesale(在线批发)。<br>中文站：普通订单类型 = &quot;cn&quot;;<br>大额批发订单类型 = &quot;ws&quot;;<br>普通拿样订单类型 = &quot;yp&quot;;<br>一分钱拿样订单类型 = &quot;yf&quot;;<br>倒批(限时折扣)订单类型 = &quot;fs&quot;;<br>加工定制订单类型 = &quot;cz&quot;;<br>协议采购订单类型 = &quot;ag&quot;;<br>伙拼订单类型 = &quot;hp&quot;;<br>供销订单类型 = &quot;supply&quot;;<br>淘工厂订单 = &quot;factory&quot;;<br>快订下单  = &quot;quick&quot;;<br>享拼订单  = &quot;xiangpin&quot;;<br>当面付 = &quot;f2f&quot;;<br>存样服务 = &quot;cyfw&quot;;<br>代销订单 = &quot;sp&quot;;<br>微供订单 = &quot;wg&quot;;零售通 = &quot;lst&quot;;跨境=&#39;cb&#39;;分销=&#39;distribution&#39;;采源宝=&#39;cab&#39;;加工定制=&quot;manufact&quot; | cn |
| `overSeaOrder` | java.lang.Boolean | 是 | 是否海外代发订单，是：true | true |
| `refundId` | java.lang.String | 是 | 退款单ID | TQ4562212313 |
| `refund` | java.math.BigDecimal | 是 | 退款金额，单位为元 | 1 |
| `status` | java.lang.String | 是 | 交易状态，waitbuyerpay:等待买家付款;waitsellersend:等待卖家发货;waitlogisticstakein:等待物流公司揽件;waitbuyerreceive:等待买家收货;waitbuyersign:等待买家签收;signinsuccess:买家已签收;confirm_goods:已收货;success:交易成功;cancel:交易取消;terminated:交易终止;未枚举:其他状态 | waitbuyerpay |
| `refundPayment` | Long | 是 | 退款金额 | 1 |
| `sellerContact` | [message:alibaba.trade.tradeSellerContact](#m-alibaba-trade-tradesellercontact) | 是 | 卖家联系人信息 | {} |
| `relatedCode` | String | 是 | 关联code | 229195140003841187 |
| `buyerOpenUid` | String | 是 | 买家加密ID，可通过接口解密，该ID唯一。注意：该ID在不同appkey内返回不一致 | xxx |
| `buyerSubOpenUid` | String | 是 | 买家子账号加密ID，可通过接口解密，该ID唯一。注意：该ID在不同appkey内返回不一致 | xxx |
| `buyerUserId` | Long | 是 | 买家Userid | 22919514 |
| `sellerUserId` | Long | 是 | 卖家UserId | 22919514 |
| `buyerAlipayId` | String | 是 | 买家AlipayId | 441951400038411872 |
| `sellerAlipayId` | String | 是 | 卖家AlipayId | 221400038411871123 |

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

<a id="m-alibaba-trade-tradesellercontact"></a>
#### alibaba.trade.tradeSellerContact

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `phone` | String | 是 | 联系电话 |  |
| `fax` | String | 是 | 传真 |  |
| `email` | String | 是 | 邮箱 |  |
| `imInPlatform` | String | 是 | 联系人在平台的IM账号 |  |
| `name` | String | 是 | 联系人名称 |  |
| `mobile` | String | 是 | 联系人手机号 |  |
| `companyName` | java.lang.String | 是 | 公司名称 |  |
| `wgSenderName` | java.lang.String | 是 | 发件人名称，在微供等分销场景下由分销商设置 | 张** |
| `wgSenderPhone` | java.lang.String | 是 | 发件人电话，在微供等分销场景下由分销商设置 | 13800000000 |
| `shopName` | String | 是 | 旺铺名称 |  三生网媒邮箱公司 |

<a id="m-alibaba-openplatform-trade-model-nativelogisticsinfo"></a>
#### alibaba.openplatform.trade.model.NativeLogisticsInfo

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `caid` | String | 是 | 解密地址ID，用于电商平台收货人信息加密场景使用，非订单加密场景请勿使用。 |   |
| `address` | java.lang.String | 是 | 详细地址 |   |
| `area` | java.lang.String | 是 | 县，区 |   |
| `areaCode` | java.lang.String | 是 | 省市区编码 |   |
| `city` | java.lang.String | 是 | 城市 |   |
| `contactPerson` | java.lang.String | 是 | 联系人姓名 |   |
| `fax` | java.lang.String | 是 | 传真 |   |
| `mobile` | java.lang.String | 是 | 手机 |   |
| `province` | java.lang.String | 是 | 省份 |   |
| `telephone` | java.lang.String | 是 | 电话 |   |
| `zip` | java.lang.String | 是 | 邮编 |   |
| `logisticsItems` | [message:alibaba.openplatform.trade.model.NativeLogisticsItemsInfo[]](#m-alibaba-openplatform-trade-model-nativelogisticsitemsinfo[]) | 是 | 运单明细 |   |
| `townCode` | java.lang.String | 是 | 镇，街道地址码 |   |
| `town` | java.lang.String | 是 | 镇，街道 |   |

<a id="m-alibaba-openplatform-trade-model-nativelogisticsitemsinfo[]"></a>
#### alibaba.openplatform.trade.model.NativeLogisticsItemsInfo[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `deliveredTime` | java.util.Date | 是 | 发货时间 |  |
| `logisticsCode` | java.lang.String | 是 | 物流编号 |  |
| `type` | java.lang.String | 是 | SELF_SEND_GOODS(&quot;0&quot;)自行发货，在线发货ONLINE_SEND_GOODS(&quot;1&quot;，不需要物流的发货 NO_LOGISTICS_SEND_GOODS(&quot;2&quot;) |  |
| `id` | java.lang.Long | 是 | 主键id |  |
| `status` | java.lang.String | 是 | 状态 |  |
| `gmtModified` | java.util.Date | 是 | 修改时间 |  |
| `gmtCreate` | java.util.Date | 是 | 创建时间 |  |
| `carriage` | java.math.BigDecimal | 是 | 运费(单位为元) |  |
| `fromProvince` | java.lang.String | 是 | 发货省 |  |
| `fromCity` | java.lang.String | 是 | 发货市 |  |
| `fromArea` | java.lang.String | 是 | 发货区 |  |
| `fromAddress` | java.lang.String | 是 | 发货街道地址 |  |
| `fromPhone` | java.lang.String | 是 | 发货联系电话 |  |
| `fromMobile` | java.lang.String | 是 | 发货联系手机 |  |
| `fromPost` | java.lang.String | 是 | 发货地址邮编 |  |
| `logisticsCompanyId` | java.lang.Long | 是 | 物流公司Id |  |
| `logisticsCompanyNo` | java.lang.String | 是 | 物流公司编号 |  |
| `logisticsCompanyName` | java.lang.String | 是 | 物流公司名称 |  |
| `logisticsBillNo` | java.lang.String | 是 | 物流公司运单号 |  |
| `subItemIds` | java.lang.String | 是 | 商品明细条目id，如有多个以,分隔 |  |
| `toProvince` | java.lang.String | 是 | 收货省 |  |
| `toCity` | java.lang.String | 是 | 收货市 |  |
| `toArea` | java.lang.String | 是 | 收货区 |  |
| `toAddress` | java.lang.String | 是 | 收货街道地址 |  |
| `toPhone` | java.lang.String | 是 | 收货联系电话 |  |
| `toMobile` | java.lang.String | 是 | 收货联系手机 |  |
| `toPost` | java.lang.String | 是 | 收货地址邮编 |  |

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

<a id="m-alibaba-openplatform-trade-model-productiteminfo[]"></a>
#### alibaba.openplatform.trade.model.ProductItemInfo[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `cargoNumber` | java.lang.String | 是 | 指定单品货号，国际站无需关注。该字段不一定有值，仅仅在下单时才会把货号记录(如果卖家设置了单品货号的话)。别的订单类型的货号只能通过商品接口去获取。请注意：通过商品接口获取时的货号和下单时的货号可能不一致，因为下单完成后卖家可能修改商品信息，改变了货号。 | 123 |
| `description` | java.lang.String | 是 | 描述,1688无此信息 | 描述 |
| `itemAmount` | java.math.BigDecimal | 是 | 实付金额，单位为元 | 12 |
| `name` | java.lang.String | 是 | 商品名称 | 商品名称 |
| `price` | java.math.BigDecimal | 是 | 原始单价，以元为单位 | 12.5 |
| `productID` | java.lang.Long | 是 | 产品ID（非在线产品为空） | 12345666 |
| `productImgUrl` | String[] | 是 | 商品图片url | ["http://cbu01.alicdn.com/img/ibank/2019/700/221/12771122007.80x80.jpg"] |
| `productSnapshotUrl` | java.lang.String | 是 | 产品快照url，交易订单产生时会自动记录下当时的商品快照，供后续纠纷时参考 | https://trade.1688.com/order/offer_snapshot.htm?order_entry_id=731312321827747370 |
| `quantity` | java.math.BigDecimal | 是 | 以unit为单位的数量，例如多少个、多少件、多少箱、多少吨 | 12 |
| `refund` | java.math.BigDecimal | 是 | 退款金额，单位为元 | 12 |
| `skuID` | java.lang.Long | 是 | skuID | 12 |
| `sort` | java.lang.Integer | 是 | 排序字段，商品列表按此字段进行排序，从0开始，1688不提供 | 0 |
| `status` | java.lang.String | 是 | 子订单状态，如果订单状态为待发货状态（waitsellersend），还需要根据canSendGoods来判断当前子订单能否发货，如果为false则当前子订单不能发货，为true则说明当前子订单可以正常发货 | waitsellersend |
| `subItemID` | java.lang.Long | 是 | 子订单号，或商品明细条目ID | 731312321827747370 |
| `type` | java.lang.String | 是 | 类型，国际站使用，供卖家标注商品所属类型 | common |
| `unit` | java.lang.String | 是 | 售卖单位	例如：个、件、箱、吨 | 个 |
| `weight` | java.lang.String | 是 | 重量	按重量单位计算的重量，例如：100 | 100 |
| `weightUnit` | java.lang.String | 是 | 重量单位	例如：g，kg，t | g |
| `guaranteesTerms` | [message:alibaba.openplatform.trade.model.GuaranteeTermsInfo[]](#m-alibaba-openplatform-trade-model-guaranteetermsinfo[]) | 是 | 保障条款，此字段仅针对1688 | {} |
| `productCargoNumber` | java.lang.String | 是 | 指定商品货号，该字段不一定有值，在下单时才会把货号记录。别的订单类型的货号只能通过商品接口去获取。请注意：通过商品接口获取时的货号和下单时的货号可能不一致，因为下单完成后卖家可能修改商品信息，改变了货号。该字段和cargoNUmber的区别是：该字段是定义在商品级别上的货号，cargoNUmber是定义在单品级别的货号 | 123 |
| `skuInfos` | [message:alibaba.trade.SkuItemDesc[]](#m-alibaba-trade-skuitemdesc[]) | 是 | 规格信息 | {} |
| `entryDiscount` | Long | 是 | 订单明细涨价或降价的金额 | 0 |
| `specId` | java.lang.String | 是 | 订单销售属性ID | 213123123123213ecfw12331 |
| `quantityFactor` | java.math.BigDecimal | 是 | 以unit为单位的quantity精度系数，值为10的幂次，例如:quantityFactor=1000,unit=吨，那么quantity的最小精度为0.001吨 | 1000 |
| `statusStr` | java.lang.String | 是 | 子订单状态描述 | 等待卖家同意 |
| `refundStatus` | java.lang.String | 是 | WAIT_SELLER_AGREE 等待卖家同意<br>REFUND_SUCCESS 退款成功<br>REFUND_CLOSED 退款关闭<br>WAIT_BUYER_MODIFY 待买家修改<br>WAIT_BUYER_SEND 等待买家退货<br>WAIT_SELLER_RECEIVE 等待卖家确认收货 | WAIT_SELLER_AGREE |
| `closeReason` | java.lang.String | 是 | 关闭原因 | 测试 |
| `logisticsStatus` | java.lang.Integer | 是 | 1 未发货<br>2 已发货<br>3 已收货<br>4 已经退货<br>5 部分发货<br>8 还未创建物流订单 | 1 |
| `gmtCreate` | java.util.Date | 是 | 创建时间 | 20190801154220368+0800 |
| `gmtModified` | java.util.Date | 是 | 修改时间 | 20190801154220368+0800 |
| `gmtCompleted` | java.util.Date | 是 | 明细完成时间 | 20190801154220368+0800 |
| `gmtPayExpireTime` | String | 是 | 库存超时时间，格式为“yyyy-MM-dd HH:mm:ss” | 2020-03-01 08:00:00 |
| `refundId` | String | 是 | 退款单号 | TQ12345t5 |
| `subItemIDString` | String | 是 | 子订单号，或商品明细条目ID(字符串类型，由于Long类型的ID可能在JS和PHP中处理有问题，所以可以用字符串类型来处理) | 731312321827747370 |
| `canSendGoods` | Boolean | 是 | 当前子订单是否可发货，如果当前子订单状态为待发货状态（waitsellersend），只要canSendGoods为false，不区分业务，该子订单则不可以发货，为true则可以正常发货 | true |
| `cantSendReason` | String | 是 | 不可发货的原因，拼团、C2M业务都有 | 未成团，不可发货/该订单已由系统自动安排菜鸟仓发货，请勿重复发货 |
| `permitLogisticsCpCode` | String | 是 | 允许声明发货的物流公司cpCode，多个cpCode以半角逗号间隔 |   |
| `deliverGoodsOverdueRisk` | String | 是 | 即将揽收超时标签，overdue：即将揽收超时 | overdue |

<a id="m-alibaba-openplatform-trade-model-guaranteetermsinfo[]"></a>
#### alibaba.openplatform.trade.model.GuaranteeTermsInfo[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `assuranceInfo` | java.lang.String | 是 | 保障条款 |  |
| `assuranceType` | java.lang.String | 是 | 保障方式。国际站：TA(信保) |  |
| `qualityAssuranceType` | java.lang.String | 是 | 质量保证类型。国际站：pre_shipment(发货前),post_delivery(发货后) |  |

<a id="m-alibaba-trade-skuitemdesc[]"></a>
#### alibaba.trade.SkuItemDesc[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `name` | String | 是 | 属性名 |  |
| `value` | String | 是 | 属性值 |  |

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

<a id="m-alibaba-invoice-orderinvoicemodel"></a>
#### alibaba.invoice.OrderInvoiceModel

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `invoiceCompanyName` | java.lang.String | 是 | 发票公司名称(即发票抬头-title) |  |
| `invoiceType` | java.lang.Integer | 是 | 发票类型. 0：普通发票，1:增值税发票，9未知类型 |  |
| `localInvoiceId` | java.lang.Long | 是 | 本地发票号 |  |
| `orderId` | java.lang.Long | 是 | 订单Id |  |
| `receiveCode` | java.lang.String | 是 | (收件人)址区域编码 |  |
| `receiveCodeText` | java.lang.String | 是 | (收件人) 省市区编码对应的文案(增值税发票信息) |  |
| `receiveMobile` | java.lang.String | 是 | （收件者）发票收货人手机 |  |
| `receiveName` | java.lang.String | 是 | （收件者）发票收货人 |  |
| `receivePhone` | java.lang.String | 是 | （收件者）发票收货人电话 |  |
| `receivePost` | java.lang.String | 是 | （收件者）发票收货地址邮编 |  |
| `receiveStreet` | java.lang.String | 是 | (收件人) 街道地址(增值税发票信息) |  |
| `registerAccountId` | java.lang.String | 是 | (公司)银行账号 |  |
| `registerBank` | java.lang.String | 是 | (公司)开户银行 |  |
| `registerCode` | java.lang.String | 是 | (注册)省市区编码 |  |
| `registerCodeText` | java.lang.String | 是 | (注册)省市区文本 |  |
| `registerPhone` | java.lang.String | 是 | （公司）注册电话 |  |
| `registerStreet` | java.lang.String | 是 | (注册)街道地址 |  |
| `taxpayerIdentify` | java.lang.String | 是 | 纳税人识别号 |  |

<a id="m-alibaba-openplatform-trade-model-guaranteetermsinfo"></a>
#### alibaba.openplatform.trade.model.GuaranteeTermsInfo

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `assuranceInfo` | java.lang.String | 是 | 保障条款 |  |
| `assuranceType` | java.lang.String | 是 | 保障方式。国际站：TA(信保) |  |
| `qualityAssuranceType` | java.lang.String | 是 | 质量保证类型。国际站：pre_shipment(发货前),post_delivery(发货后) |  |

<a id="m-alibaba-openplatform-trade-keyvaluepair[]"></a>
#### alibaba.openplatform.trade.KeyValuePair[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `key` | String | 是 | 键 |  |
| `value` | String | 是 | 值 |  |
| `description` | String | 是 | 描述 |  |

## 示例

**非常重要：关于订单是否能发货的说明**

```
只有当子订单状态为待发货状态（waitsellersend）时，且当前子订单下canSendGoods为true时，此子订单对应的商品才可以发货，如果子订单状态为待发货状态（waitsellersend），且当前子订单canSendGoods为false，则该子订单不可以发货，否则商家会造成资损，目前有拼团和C2M等业务都会有管种情况的订单，拼团的订单不能发货的原因（cantSendReason）是订单未成团，C2M业务不能发货的原因是该订单已经由系统自动安排签约的菜鸟仓发货。注意：对于发货场景，请求参数includeFields必须要传CanSendCheck才可以获取此信息。
```

**出参示例**

```
{
  "result": [
	 {
      "nativeLogistics": {
        "area": "东城区",
        "areaCode": "110101",
        "city": "北京市",
        "contactPerson": "黎明",
        "province": "北京",
        "zip": "721000",
        "townCode": "110101001",
        "town": "东华门街道"
      },
      "productItems": [
        {
          "itemAmount": 69.3,
          "name": "散装太空豆柱形橄榄形小小号大中小100个1盒",
          "price": 9,
          "productID": 570959966604,
          "productImgUrl": [
            "http://cbu01.alicdn.com/img/order/trading/118/969/342303144391/9163043580_1606139362.310x310.80x80.jpg",
            "http://cbu01.alicdn.com/img/order/trading/118/969/342303144391/9163043580_1606139362.310x310.jpg"
          ],
          "productSnapshotUrl": "https://trade.1688.com/order/offer_snapshot.htm?order_entry_id=193441303243969811",
          "quantity": 11,
          "refund": 0,
          "status": "waitbuyerpay",
          "subItemID": 193441303243969820,
          "type": "common",
          "guaranteesTerms": [],
          "entryDiscount": 0,
          "quantityFactor": 1,
          "statusStr": "等待买家付款",
          "logisticsStatus": 1,
          "gmtCreate": "20180721172607000+0800",
          "gmtModified": "20180721172608000+0800",
          "subItemIDString": "193441303243969811"
        }
      ],
      "tradeTerms": [],
      "orderRateInfo": {
        "buyerRateStatus": 5,
        "sellerRateStatus": 5
      },
      "baseInfo": {
        "businessType": "cn",
        "buyerID": "b2b-1624961198",
        "createTime": "20180721172607000+0800",
        "id": 193441303243969820,
        "modifyTime": "20180721172608000+0800",
        "refund": 0,
        "sellerID": "b2b-1623492085",
        "shippingFee": 1,
        "status": "waitbuyerpay",
        "totalAmount": 70.3,
        "discount": 0,
        "buyerContact": {
          "phone": "86-0591-400-800-533",
          "imInPlatform": "alitestforisv02",
          "name": "测试账号暮星",
          "companyName": "AOP对外测试账号02"
        },
        "sellerContact": {
          "phone": "86-0591-13805040134",
          "imInPlatform": "alitestforisv01",
          "name": "测试账号一暮星",
          "companyName": "AOP对外测试账号01"
        },
        "tradeType": "50060",
        "refundPayment": 0,
        "idOfStr": "193441303243969811",
        "alipayTradeId": "UNCREATED",
        "receiverInfo": {
          "toFullName": "黎明",
          "toDivisionCode": "110101",
          "toPost": "721000",
          "toTownCode": "110101001",
          "toArea": "北京 北京市 东城区 东华门街道"
        },
        "buyerLoginId": "alitestforisv02",
        "sellerLoginId": "alitestforisv01",
        "buyerUserId": 16249****8,
        "sellerUserId": 162349****5,
        "buyerAlipayId": "208861149****764",
        "sellerAlipayId": "208861149****533",
        "sumProductPayment": 99,
        "stepPayAll": false,
        "overSeaOrder": false,
        "sellerOrder": false,
        "flowTemplateCode": "assureTradeAlipay"
      }
    }
  ],
  "totalRecord": 615
}
```
