# 订单列表查看(买家视角)

API: `com.alibaba.trade:alibaba.trade.getBuyerOrderList:1` · Category: 订单  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.trade.getBuyerOrderList-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.trade.getBuyerOrderList/{appKey}`  
需要授权 (access_token) · 需要签名

获取买家的订单列表，也就是用户的memberId必须等于订单里的买家memberId。该接口仅仅返回订单基本信息，不会返回订单的物流信息和发票信息；如果需要获取物流信息，请调用获取订单详情接口；如果需要获取发票信息，请调用获取发票信息的API

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `bizTypes` | String[] | 否 | 业务类型，支持： &quot;cn&quot;(普通订单类型), &quot;ws&quot;(大额批发订单类型), &quot;yp&quot;(普通拿样订单类型), &quot;yf&quot;(一分钱拿样订单类型), &quot;fs&quot;(倒批(限时折扣)订单类型), &quot;cz&quot;(加工定制订单类型), &quot;ag&quot;(协议采购订单类型), &quot;hp&quot;(伙拼订单类型), &quot;gc&quot;(国采订单类型), &quot;supply&quot;(供销订单类型), &quot;nyg&quot;(nyg订单类型), &quot;factory&quot;(淘工厂订单类型), &quot;quick&quot;(快订下单), &quot;xiangpin&quot;(享拼订单), &quot;nest&quot;(采购商城-鸟巢), &quot;f2f&quot;(当面付), &quot;cyfw&quot;(存样服务), &quot;sp&quot;(代销订单标记), &quot;wg&quot;(微供订单), &quot;factorysamp&quot;(淘工厂打样订单), &quot;factorybig&quot;(淘工厂大货订单) | ["cn","ws"] |
| `createEndTime` | java.util.Date | 否 | 下单结束时间 | 20180802211113000+0800 |
| `createStartTime` | java.util.Date | 否 | 下单开始时间 | 20180102211113000+0800 |
| `isHis` | boolean | 否 | 是否查询历史订单表,默认查询当前表，即默认值为false | false |
| `modifyEndTime` | java.util.Date | 否 | 查询修改时间结束 | 20180802211113000+0800 |
| `modifyStartTime` | java.util.Date | 否 | 查询修改时间开始 | 20180102211113000+0800 |
| `orderStatus` | java.lang.String | 否 | 订单状态，值有 success, cancel(交易取消，违约金等交割完毕), waitbuyerpay(等待卖家付款)， waitsellersend(等待卖家发货), waitbuyerreceive(等待买家收货 ) | success |
| `page` | int | 否 | 查询分页页码，从1开始 | 1 |
| `pageSize` | int | 否 | 查询的每页的数量 | 20 |
| `refundStatus` | java.lang.String | 否 | 退款状态，支持： &quot;waitselleragree&quot;(等待卖家同意), &quot;refundsuccess&quot;(退款成功), &quot;refundclose&quot;(退款关闭), &quot;waitbuyermodify&quot;(待买家修改), &quot;waitbuyersend&quot;(等待买家退货), &quot;waitsellerreceive&quot;(等待卖家确认收货) | refundsuccess |
| `sellerMemberId` | java.lang.String | 否 | 卖家memberId | b2b-1624961198 |
| `sellerLoginId` | String | 否 | 卖家loginId | alitestforisv02 |
| `sellerRateStatus` | java.lang.Integer | 否 | 卖家评价状态 (4:已评价,5:未评价,6;不需要评价) | 6 |
| `tradeType` | java.lang.String | 否 | 交易类型:<br>担保交易(1),<br>预存款交易(2),<br>ETC境外收单交易(3),<br>即时到帐交易(4),<br>保障金安全交易(5),<br>统一交易流程(6),<br>分阶段交易(7),<br>货到付款交易(8),<br>信用凭证支付交易(9),<br>账期支付交易(10),<br>1688交易4.0，新分阶段交易(50060),<br>当面付的交易流程(50070),<br>服务类的交易流程(50080) | 50060 |
| `productName` | java.lang.String | 否 | 商品名称 | 测试商品 |
| `needBuyerAddressAndPhone` | java.lang.Boolean | 否 | 是否需要查询买家的详细地址信息和电话 | false |
| `needMemoInfo` | java.lang.Boolean | 否 | 是否需要查询备注信息 | false |
| `outOrderId` | String | 否 | 外部订单号，可用于控制幂等 | 90187872898371 |
| `needInvoicingSetting` | Boolean | 否 | 需要查询开票设置 | true |
| `orderIds` | Long[] | 否 | 查询多个订单id | [90187872898371,90187872898372] |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:alibaba.openplatform.trade.model.TradeInfo[]](#m-alibaba-openplatform-trade-model-tradeinfo[]) | 是 | 查询返回列表 | [] |
| `errorCode` | String | 是 | 错误码 |   |
| `errorMessage` | String | 是 | 错误信息 |   |
| `totalRecord` | Long | 是 | 总记录数 | 528 |

<a id="m-alibaba-openplatform-trade-model-tradeinfo[]"></a>
#### alibaba.openplatform.trade.model.TradeInfo[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `quoteList` | [message:alibaba.orderDetail.caigouQuoteInfo[]](#m-alibaba-orderdetail-caigouquoteinfo[]) | 是 | 采购单详情列表，为大企业采购订单独有域。 | {} |
| `extAttributes` | [message:alibaba.openplatform.trade.KeyValuePair[]](#m-alibaba-openplatform-trade-keyvaluepair[]) | 是 | 订单扩展属性 | [] |
| `orderRateInfo` | [message:alibaba.trade.OrderRateInfo](#m-alibaba-trade-orderrateinfo) | 是 | 订单评价信息 | {} |
| `orderInvoiceInfo` | [message:alibaba.invoice.OrderInvoiceModel](#m-alibaba-invoice-orderinvoicemodel) | 是 | 发票信息 | {} |
| `tradeTerms` | [message:alibaba.openplatform.trade.model.TradeTermsInfo[]](#m-alibaba-openplatform-trade-model-tradetermsinfo[]) | 是 | 交易条款 | {} |
| `nativeLogistics` | [message:alibaba.openplatform.trade.model.NativeLogisticsInfo](#m-alibaba-openplatform-trade-model-nativelogisticsinfo) | 是 | 国内物流 | {} |
| `guaranteesTerms` | [message:alibaba.openplatform.trade.model.GuaranteeTermsInfo](#m-alibaba-openplatform-trade-model-guaranteetermsinfo) | 是 | 保障条款 | {} |
| `baseInfo` | [message:alibaba.openplatform.trade.model.OrderBaseInfo](#m-alibaba-openplatform-trade-model-orderbaseinfo) | 是 | 订单基础信息 | {} |
| `orderBizInfo` | [message:alibaba.order.bizInfo](#m-alibaba-order-bizinfo) | 是 | 订单业务信息 | {} |
| `productItems` | [message:alibaba.openplatform.trade.model.ProductItemInfo[]](#m-alibaba-openplatform-trade-model-productiteminfo[]) | 是 | 商品条目信息 | {} |
| `overseasExtraAddress` | [message:alibaba.trade.OverseasExtraAddress](#m-alibaba-trade-overseasextraaddress) | 是 | 跨境地址扩展信息 | {} |
| `customs` | [message:alibaba.trade.Customs](#m-alibaba-trade-customs) | 是 | 跨境报关信息 | {} |
| `overseaLogisticsInfo` | [message:com.alibaba.ocean.openplatform.biz.trade.common.model.OverseaLogisticsInfo](#m-com-alibaba-ocean-openplatform-biz-trade-common-model-oversealogisticsinfo) | 是 | 海外收货信息 | { 					"overseasTransportUserAddrInfo":{ 						"addressDetail":"test", 						"cityName":"Г.Курчатов", 						"areaName":" ", 						"warehouseContactName":"Zeus", 						"mobile":"123456789", 						"fixedPhone":"123456789", 						"postCode":"1111", 						"provinceName":"Абайская область", 						"countryName":"Казахстан", 						"@type":"com.alibaba.ocean.openplatform.biz.cross.model.OverseasTransportUserAddrInfo", 						"warehouseName":" " 					}, 					"overseasLogisticsIds":[ 						"TN0003077L" 					],} |
| `invoicingSettingModel` | [message:com.alibaba.ocean.openplatform.biz.trade.common.model.OpSellerInvoiceTradeSettingModel](#m-com-alibaba-ocean-openplatform-biz-trade-common-model-opsellerinvoicetradesettingmodel) | 是 | 开票设置 | {} |

<a id="m-alibaba-orderdetail-caigouquoteinfo[]"></a>
#### alibaba.orderDetail.caigouQuoteInfo[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `productQuoteName` | String | 是 | 供应单项的名称 | 物料01 |
| `price` | Long | 是 | 价格，单位：分 | 100 |
| `count` | Double | 是 | 购买数量 | 10 |

<a id="m-alibaba-openplatform-trade-keyvaluepair[]"></a>
#### alibaba.openplatform.trade.KeyValuePair[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `key` | String | 是 | 键 |   |
| `value` | String | 是 | 值 |   |
| `description` | String | 是 | 描述 |   |

<a id="m-alibaba-trade-orderrateinfo"></a>
#### alibaba.trade.OrderRateInfo

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `buyerRateStatus` | Integer | 是 | 买家评价状态(4:已评论,5:未评论,6;不需要评论) | 5 |
| `sellerRateStatus` | Integer | 是 | 卖家评价状态(4:已评论,5:未评论,6;不需要评论) | 5 |

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
| `amount` | Long | 是 | 发票含税总额（分） | 1 |
| `status` | String | 是 | 发票状态：{&quot;ISSUED&quot;:&quot;已开票&quot;,&quot;RED_ISSUING&quot;:&quot;冲红中&quot;,&quot;RED_ALL_ISSUED&quot;:&quot;全部冲红&quot;,&quot;CLOSED&quot;:&quot;关闭&quot;,&quot;RED_PART_ISSUED&quot;:&quot;部分冲红&quot;,&quot;VERIFYING&quot;:&quot;验票中&quot;,&quot;DEPRECATED&quot;:&quot;已作废&quot;,&quot;RETURNING&quot;:&quot;退票中&quot;,&quot;INVALIDING&quot;:&quot;作废中&quot;,&quot;INIT&quot;:&quot;初始化&quot;,&quot;VERIFY_FAILED&quot;:&quot;验票失败&quot;,&quot;ISSUING&quot;:&quot;开票中&quot;,&quot;FAILED&quot;:&quot;开票失败&quot;,&quot;RETURNED&quot;:&quot;已退票&quot;} | issued |

<a id="m-alibaba-openplatform-trade-model-tradetermsinfo[]"></a>
#### alibaba.openplatform.trade.model.TradeTermsInfo[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `payStatus` | java.lang.String | 是 | 支付状态。国际站：WAIT_PAY(未支付),PAYER_PAID(已完成支付),PART_SUCCESS(部分支付成功),PAY_SUCCESS(支付成功),CLOSED(风控关闭),CANCELLED(支付撤销),SUCCESS(成功),FAIL(失败)。<br>1688:未付款，1;已付款，2;全额退款，4;卖家有收到钱，回款完成，6 ;未创建外部支付单，7; 付款前取消 ，8 ; 正在支付中 ，9 | 6 |
| `payTime` | java.util.Date | 是 | 完成阶段支付时间 | 20180807153006000+0800 |
| `payWay` | java.lang.String | 是 | 支付方式。国际站：ECL(融资支付),CC(信用卡),TT(线下TT),ACH(echecking支付)。<br>1688:支付宝1;网商银行金票2;赊购(诚e赊)3;网商银行大额转账4;电子承兑汇票 6;账期支付 7;虚拟支付8;不进行任何支付动作 9;零售通赊购支付 10 | 1 |
| `phasAmount` | java.math.BigDecimal | 是 | 阶段金额 | 1158 |
| `phase` | java.lang.Long | 是 | 阶段 | 2988886630908522 |
| `phaseCondition` | java.lang.String | 是 | 阶段条件，1688无此内容 |   |
| `phaseDate` | java.lang.String | 是 | 阶段时间，1688无此内容 |   |
| `payWayDesc` | String | 是 | 支付方式 | 支付宝 |

<a id="m-alibaba-openplatform-trade-model-nativelogisticsinfo"></a>
#### alibaba.openplatform.trade.model.NativeLogisticsInfo

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `address` | java.lang.String | 是 | 详细地址 | 浙江省杭州市滨江区 |
| `area` | java.lang.String | 是 | 县，区 | 滨江区 |
| `areaCode` | java.lang.String | 是 | 省市区编码 | 330108 |
| `city` | java.lang.String | 是 | 城市 | 杭州市 |
| `contactPerson` | java.lang.String | 是 | 联系人姓名 | 张三 |
| `fax` | java.lang.String | 是 | 传真 |   |
| `mobile` | java.lang.String | 是 | 手机 | 15000001111 |
| `province` | java.lang.String | 是 | 省份 | 浙江省 |
| `telephone` | java.lang.String | 是 | 电话 | 7314065 |
| `zip` | java.lang.String | 是 | 邮编 | 888888 |
| `townCode` | java.lang.String | 是 | 镇，街道地址码 | 4403011 |
| `town` | java.lang.String | 是 | 镇，街道 | 测试街道 |

<a id="m-alibaba-openplatform-trade-model-guaranteetermsinfo"></a>
#### alibaba.openplatform.trade.model.GuaranteeTermsInfo

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `assuranceInfo` | java.lang.String | 是 | 保障条款 |  自愿选择向买家提供“交期保障”服务 |
| `assuranceType` | java.lang.String | 是 | 保障方式。国际站：TA(信保) |  jqbz |
| `qualityAssuranceType` | java.lang.String | 是 | 质量保证类型。国际站：pre_shipment(发货前),post_delivery(发货后) |  交期保障 |
| `value` | String | 是 | 保障条款值，比如交期保障里，6表示6天 | 6 |

<a id="m-alibaba-openplatform-trade-model-orderbaseinfo"></a>
#### alibaba.openplatform.trade.model.OrderBaseInfo

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `allDeliveredTime` | java.util.Date | 是 | 完全发货时间 |   |
| `businessType` | java.lang.String | 是 | 业务类型。国际站：ta(信保),wholesale(在线批发)。<br>中文站：普通订单类型 = &quot;cn&quot;;<br>大额批发订单类型 = &quot;ws&quot;;<br>普通拿样订单类型 = &quot;yp&quot;;<br>一分钱拿样订单类型 = &quot;yf&quot;;<br>倒批(限时折扣)订单类型 = &quot;fs&quot;;<br>加工定制订单类型 = &quot;cz&quot;;<br>协议采购订单类型 = &quot;ag&quot;;<br>伙拼订单类型 = &quot;hp&quot;;<br>供销订单类型 = &quot;supply&quot;;<br>淘工厂订单 = &quot;factory&quot;;<br>快订下单  = &quot;quick&quot;;<br>享拼订单  = &quot;xiangpin&quot;;<br>当面付 = &quot;f2f&quot;;<br>存样服务 = &quot;cyfw&quot;;<br>代销订单 = &quot;sp&quot;;<br>微供订单 = &quot;wg&quot;;零售通 = &quot;lst&quot;; | cb |
| `buyerID` | java.lang.String | 是 | 买家主账号id | b2b-397390228 |
| `buyerMemo` | java.lang.String | 是 | 买家备忘信息 |   |
| `buyerSubID` | java.lang.Long | 是 | 买家子账号id，1688无此内容 |   |
| `completeTime` | java.util.Date | 是 | 完成时间 |   |
| `createTime` | java.util.Date | 是 | 创建时间 | 20180807152315000+0800 |
| `currency` | java.lang.String | 是 | 币种，币种，整个交易单使用同一个币种。值范围：USD,RMB,HKD,GBP,CAD,AUD,JPY,KRW,EUR |   |
| `id` | java.lang.Long | 是 | 交易id | 1997017368139085 |
| `modifyTime` | java.util.Date | 是 | 修改时间 | 20180807152533000+0800 |
| `payTime` | java.util.Date | 是 | 付款时间，如果有多次付款，这里返回的是首次付款时间 | 20180807153006000+0800 |
| `receivingTime` | java.util.Date | 是 | 收货时间，这里返回的是完全收货时间 |   |
| `refund` | java.math.BigDecimal | 是 | 退款金额，单位为元 | 0 |
| `remark` | java.lang.String | 是 | 备注，1688指下单时的备注 | 测试备注 |
| `sellerID` | java.lang.String | 是 | 卖家主账号id | b2b-23187037 |
| `sellerSubID` | java.lang.Long | 是 | 卖家子账号id，1688无此内容 |   |
| `shippingFee` | java.math.BigDecimal | 是 | 运费，单位为元 | 10 |
| `status` | java.lang.String | 是 | 交易状态，waitbuyerpay:等待买家付款;waitsellersend:等待卖家发货;waitlogisticstakein:等待物流公司揽件;waitbuyerreceive:等待买家收货;waitbuyersign:等待买家签收;signinsuccess:买家已签收;confirm_goods:已收货;success:交易成功;cancel:交易取消;terminated:交易终止;未枚举:其他状态 | waitbuyerpay |
| `totalAmount` | java.math.BigDecimal | 是 | 应付款总金额，totalAmount = ∑itemAmount + shippingFee，单位为元 | 528 |
| `buyerRemarkIcon` | String | 是 | 买家备忘标志 |   |
| `discount` | Long | 是 | 折扣信息，单位分 | -226 |
| `buyerContact` | [message:alibaba.trade.tradeContact](#m-alibaba-trade-tradecontact) | 是 | 买家联系人 | {} |
| `sellerContact` | [message:alibaba.trade.tradeSellerContact](#m-alibaba-trade-tradesellercontact) | 是 | 卖家联系人 | {} |
| `tradeType` | String | 是 | 1:担保交易<br>2:预存款交易<br>3:ETC境外收单交易<br>4:即时到帐交易<br>5:保障金安全交易<br>6:统一交易流程<br>7:分阶段付款<br>8.货到付款交易<br>9.信用凭证支付交易<br>10.账期支付交易 | 50060 |
| `refundStatus` | String | 是 | 订单的售中退款状态 |   |
| `refundStatusForAs` | String | 是 | 订单的售后退款状态 |   |
| `refundPayment` | Long | 是 | 退款金额 | 0 |
| `idOfStr` | String | 是 | 交易id(字符串格式) | 1997017368139085 |
| `alipayTradeId` | java.lang.String | 是 | 外部支付交易Id | UNCREAT |
| `receiverInfo` | [message:alibaba.trade.orderReceiverInfo](#m-alibaba-trade-orderreceiverinfo) | 是 | 收件人信息 | {} |
| `buyerLoginId` | java.lang.String | 是 | 买家loginId，旺旺Id | alitestforisv01 |
| `sellerLoginId` | java.lang.String | 是 | 卖家oginId，旺旺Id | alitestforisv02 |
| `buyerUserId` | java.lang.Long | 是 | 买家数字id | 39739022 |
| `sellerUserId` | java.lang.Long | 是 | 卖家数字id | 23187037 |
| `buyerAlipayId` | java.lang.String | 是 | 买家支付宝id | 20881314957453 |
| `sellerAlipayId` | java.lang.String | 是 | 卖家支付宝id | 20887127689595 |
| `confirmedTime` | java.util.Date | 是 | 确认时间 |   |
| `closeReason` | java.lang.String | 是 | 关闭原因 |   |
| `sumProductPayment` | java.math.BigDecimal | 是 | 产品总金额(该订单产品明细表中的产品金额的和)，单位元 | 258 |
| `stepOrderList` | [message:alibaba.trade.StepOrderModel[]](#m-alibaba-trade-stepordermodel[]) | 是 | [交易3.0]分阶段交易，分阶段订单list |   |
| `stepAgreementPath` | java.lang.String | 是 | 分阶段法务协议地址 |   |
| `stepPayAll` | java.lang.Boolean | 是 | 是否一次性付款 | false |
| `buyerFeedback` | java.lang.String | 是 | 买家留言 | 测试买家留言 |
| `relatedCode` | java.lang.String | 是 | 订单关联码(如属于哪个活动等) |   |
| `overSeaOrder` | java.lang.Boolean | 是 | 是否海外代发订单，是：true | true |
| `subBuyerLoginId` | String | 是 | 买家子账号 |   |
| `sellerOrder` | java.lang.Boolean | 是 | 是否自主订单（邀约订单） |   |
| `preOrderId` | java.lang.Long | 是 | 预订单ID |   |
| `refundId` | java.lang.String | 是 | 退款单ID |   |
| `flowTemplateCode` | String | 是 | 4.0交易流程模板code |   |
| `officialSolutionOrderId` | String | 是 | 官方物流提货订单ID | 123 |
| `officialSolutionCost` | Long | 是 | 官方物流提货订单费用（分） | 770 |

<a id="m-alibaba-trade-tradecontact"></a>
#### alibaba.trade.tradeContact

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `phone` | String | 是 | 联系电话 | 0912-2525263 |
| `fax` | String | 是 | 传真 |   |
| `email` | String | 是 | 邮箱 |   |
| `imInPlatform` | String | 是 | 联系人在平台的IM账号 | alitestforisv01 |
| `name` | String | 是 | 联系人名称 | 张三 |
| `mobile` | String | 是 | 联系人手机号 | 15123543625 |
| `companyName` | java.lang.String | 是 | 公司名称 | 阿里巴巴 |

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

<a id="m-alibaba-trade-orderreceiverinfo"></a>
#### alibaba.trade.orderReceiverInfo

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `toFullName` | java.lang.String | 是 | 收件人 | 张三 |
| `toDivisionCode` | java.lang.String | 是 | 收货人地址区域编码 | 440311 |
| `toMobile` | java.lang.String | 是 | 收件人移动电话 | 15236457596 |
| `toPhone` | java.lang.String | 是 | 收件人电话 | 0912-44548585 |
| `toPost` | java.lang.String | 是 | 邮编 | 000000 |
| `toTownCode` | java.lang.String | 是 | 收货人街道或镇区域编码，可能为空 | 440306011 |
| `toArea` | java.lang.String | 是 | 收货地址 | 广东省 深圳市 龙华区 大浪街道 |

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
| `abnormalPriceChange` | Boolean | 是 | 是否已经改价 | true |

<a id="m-alibaba-creditorder-fordetail"></a>
#### alibaba.creditOrder.forDetail

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `payAmount` | java.lang.Long | 是 | 订单金额 | 10 |
| `createTime` | java.lang.String | 是 | 支付时间 | 2018-01-01 00:00:00 |
| `status` | java.lang.String | 是 | 状态 | END |
| `gracePeriodEndTime` | java.lang.String | 是 | 最晚还款时间 | 2018-01-01 00:00:00 |
| `statusStr` | java.lang.String | 是 | 状态描述 | 已完结 |
| `restRepayAmount` | java.lang.Long | 是 | 应还金额 | 11 |

<a id="m-alibaba-order-preorder-forread"></a>
#### alibaba.order.preOrder.forRead

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `appkey` | String | 是 | 创建预订单的appkey | 12345 |
| `marketName` | String | 是 | 创建预订单时传入的市场名 | dxc |
| `createPreOrderApp` | Boolean | 是 | 预订单是否为当前查询的通过当前查询的ERP创建 | false |

<a id="m-alibaba-lst-tradeinfo"></a>
#### alibaba.lst.tradeInfo

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `lstWarehouseType` | String | 是 | 零售通仓库类型。customer：虚仓；cainiao：实仓 | cainiao |

<a id="m-alibaba-openplatform-trade-model-productiteminfo[]"></a>
#### alibaba.openplatform.trade.model.ProductItemInfo[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `cargoNumber` | java.lang.String | 是 | 指定单品货号，国际站无需关注。该字段不一定有值，仅仅在下单时才会把货号记录(如果卖家设置了单品货号的话)。别的订单类型的货号只能通过商品接口去获取。请注意：通过商品接口获取时的货号和下单时的货号可能不一致，因为下单完成后卖家可能修改商品信息，改变了货号。 | E0003 |
| `description` | java.lang.String | 是 | 描述,1688无此信息 |   |
| `itemAmount` | java.math.BigDecimal | 是 | 实付金额，单位为元 | 279 |
| `name` | java.lang.String | 是 | 商品名称 | 测试商品 |
| `price` | java.math.BigDecimal | 是 | 原始单价，以元为单位 | 3 |
| `productID` | java.lang.Long | 是 | 产品ID（非在线产品为空） | 129527213581 |
| `productImgUrl` | String[] | 是 | 商品图片url | http://cbu01.alicdn.com/img/order/trading/025/894/036055839091/8536232543_347415001.80x80.jpg |
| `productSnapshotUrl` | java.lang.String | 是 | 产品快照url，交易订单产生时会自动记录下当时的商品快照，供后续纠纷时参考 | https://trade.1688.com/order/offer_snapshot.htm?order_entry_id=190938550630498520 |
| `quantity` | java.math.BigDecimal | 是 | 以unit为单位的数量，例如多少个、多少件、多少箱、多少吨 | 20 |
| `refund` | java.math.BigDecimal | 是 | 退款金额，单位为元 | 0 |
| `skuID` | java.lang.Long | 是 | skuID | 352367821371 |
| `sort` | java.lang.Integer | 是 | 排序字段，商品列表按此字段进行排序，从0开始，1688不提供 |   |
| `status` | java.lang.String | 是 | 子订单状态 | waitbuyerpay |
| `subItemID` | java.lang.Long | 是 | 商品明细条目ID | 20015919587908522 |
| `type` | java.lang.String | 是 | 类型，国际站使用，供卖家标注商品所属类型 | common |
| `unit` | java.lang.String | 是 | 售卖单位	例如：个、件、箱、吨 | 件 |
| `weight` | java.lang.String | 是 | 重量	按重量单位计算的重量，例如：100 |   |
| `weightUnit` | java.lang.String | 是 | 重量单位	例如：g，kg，t |   |
| `guaranteesTerms` | [message:alibaba.openplatform.trade.model.GuaranteeTermsInfo[]](#m-alibaba-openplatform-trade-model-guaranteetermsinfo[]) | 是 | 保障条款，此字段仅针对1688 |   |
| `productCargoNumber` | java.lang.String | 是 | 指定商品货号，该字段不一定有值，在下单时才会把货号记录。别的订单类型的货号只能通过商品接口去获取。请注意：通过商品接口获取时的货号和下单时的货号可能不一致，因为下单完成后卖家可能修改商品信息，改变了货号。该字段和cargoNUmber的区别是：该字段是定义在商品级别上的货号，cargoNUmber是定义在单品级别的货号 | C017 |
| `skuInfos` | [message:alibaba.trade.SkuItemDesc[]](#m-alibaba-trade-skuitemdesc[]) | 是 | SKU属性描述 | [] |
| `entryDiscount` | Long | 是 | 订单明细涨价或降价的金额 | 0 |
| `specId` | java.lang.String | 是 | 订单销售属性ID | 2b3878b01d251c057668066c085d75 |
| `quantityFactor` | java.math.BigDecimal | 是 | 以unit为单位的quantity精度系数，值为10的幂次，例如:quantityFactor=1000,unit=吨，那么quantity的最小精度为0.001吨 | 1 |
| `statusStr` | java.lang.String | 是 | 子订单状态描述 | 等待买家付款 |
| `refundStatus` | java.lang.String | 是 | WAIT_SELLER_AGREE 等待卖家同意<br>REFUND_SUCCESS 退款成功<br>REFUND_CLOSED 退款关闭<br>WAIT_BUYER_MODIFY 待买家修改<br>WAIT_BUYER_SEND 等待买家退货<br>WAIT_SELLER_RECEIVE 等待卖家确认收货 |   |
| `closeReason` | java.lang.String | 是 | 关闭原因 |   |
| `logisticsStatus` | java.lang.Integer | 是 | 1 未发货<br>2 已发货<br>3 已收货<br>4 已经退货<br>5 部分发货<br>8 还未创建物流订单 | 1 |
| `refundId` | String | 是 | 售中退款单号 | TQ123123 |
| `refundIdForAs` | String | 是 | 售后退款单号 | TQ123123 |
| `relatedCode` | String | 是 | 子订单关联码 | 1365465 |
| `sharePostage` | BigDecimal | 是 | 分担邮费，单位为元 |   |

<a id="m-alibaba-openplatform-trade-model-guaranteetermsinfo[]"></a>
#### alibaba.openplatform.trade.model.GuaranteeTermsInfo[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `assuranceInfo` | java.lang.String | 是 | 保障条款 |  自愿选择向买家提供“交期保障”服务 |
| `assuranceType` | java.lang.String | 是 | 保障方式。国际站：TA(信保) |  jqbz |
| `qualityAssuranceType` | java.lang.String | 是 | 质量保证类型。国际站：pre_shipment(发货前),post_delivery(发货后) |  交期保障 |
| `value` | String | 是 | 保障条款值，比如交期保障里，6表示6天 | 6 |

<a id="m-alibaba-trade-skuitemdesc[]"></a>
#### alibaba.trade.SkuItemDesc[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `name` | String | 是 | 属性名 | 颜色 |
| `value` | String | 是 | 属性值 | 黑色 |

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
| `attributes` | [message:alibaba.trade.CustomsAttributesInfo[]](#m-alibaba-trade-customsattributesinfo[]) | 是 | 报关信息列表 |   |

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

**返回参数示例**

```
{
    "result":[
        {
            "baseInfo":{
                "businessType":"cn",
                "buyerID":"b2b-1623492085",
                "completeTime":"20180806104425000+0800",
                "createTime":"20180801104333000+0800",
                "id":196965465451498520,
                "modifyTime":"20180806104425000+0800",
                "refund":0,
                "sellerID":"b2b-1624747073",
                "shippingFee":0,
                "status":"cancel",
                "totalAmount":0.01,
                "discount":-3989,
                "buyerContact":{
                    "phone":"86-0591-13805040134",
                    "name":"测试账号一暮星",
                    "imInPlatform":"alitestforisv01",
                    "companyName":"AOP对外测试账号01"
                },
                "sellerContact":{
                    "phone":"86-",
                    "name":"暮星",
                    "imInPlatform":"alitestforisv04",
                    "companyName":"AOP对外测试账号04"
                },
                "tradeType":"50060",
                "refundPayment":0,
                "idOfStr":196965465451498520,
                "alipayTradeId":"2018080121001008530224243852",
                "receiverInfo":{
                    "toFullName":"洪帮",
                    "toDivisionCode":"330782",
                    "toPost":"322000",
                    "toArea":"浙江省 金华市 义乌市"
                },
                "buyerLoginId":"alitestforisv01",
                "sellerLoginId":"alitestforisv04",
                "buyerUserId":1623492085,
                "sellerUserId":1624747073,
                "buyerAlipayId":"2088611492691533",
                "sellerAlipayId":"2088521180335604",
                "closeReason":"BUYER_NO_PAY",
                "sumProductPayment":39.9,
                "stepPayAll":false
            },
            "nativeLogistics":{
                "area":"义乌市",
                "areaCode":"330782",
                "city":"金华市",
                "contactPerson":"洪帮",
                "province":"浙江省",
                "zip":"322000"
            },
            "productItems":[
                {
                    "itemAmount":0.01,
                    "name":"夏季亚麻九分裤韩版潮流棉麻薄款修身小脚裤时尚休闲哈伦裤潮男裤",
                    "price":39.9,
                    "productID":574273466269,
                    "productImgUrl":[
                        "http://cbu01.alicdn.com/img/order/trading/025/894/154564569691/9134583276_1665066988.80x80.jpg",
                        "http://cbu01.alicdn.com/img/order/trading/025/894/154564569691/9134583276_1665066988.jpg"
                    ],
                    "productSnapshotUrl":"https://trade.1688.com/order/offer_snapshot.htm?order_entry_id=196965465451498520",
                    "quantity":1,
                    "refund":0,
                    "skuID":3753684468637,
                    "status":"cancel",
                    "subItemID":196965465451498520,
                    "type":"common",
                    "unit":"件",
                    "guaranteesTerms":[

                    ],
                    "productCargoNumber":"KY",
                    "skuInfos":[
                        {
                            "name":"颜色",
                            "value":"深灰色"
                        },
                        {
                            "name":"尺码",
                            "value":"3XL"
                        }
                    ],
                    "entryDiscount":-3989,
                    "specId":"7a2087a6594c7d6030d8ff5c42c8a2e8",
                    "quantityFactor":1,
                    "statusStr":"交易取消",
                    "closeReason":"BUYER_NO_PAY",
                    "logisticsStatus":1
                }
            ],
            "tradeTerms":[
                {
                    "payStatus":"8",
                    "payTime":"20180801104336000+0800",
                    "payWay":"1",
                    "phasAmount":0.01,
                    "phase":2960551254498520
                }
            ],
            "orderRateInfo":{
                "buyerRateStatus":5,
                "sellerRateStatus":5
            }
        }
    ],
    "totalRecord":537
}
```
