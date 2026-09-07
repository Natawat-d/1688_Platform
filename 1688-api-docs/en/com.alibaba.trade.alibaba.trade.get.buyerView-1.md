# View order details (buyer view)

Original name: 订单详情查看(买家视角)  
API: `com.alibaba.trade:alibaba.trade.get.buyerView:1` · Category: Orders  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.trade.get.buyerView-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.trade.get.buyerView/{appKey}`  
Requires user authorization (access_token) · Requires signature

Get the details of a single transaction; buyer calls only. Permission must be requested from the Alibaba Open Platform to use this API.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `webSite` | String | yes | Site information, specifies whether the called API belongs to the international site (alibaba) or the 1688 website (1688). | 1688 |
| `orderId` | Long | yes | Order ID of the transaction | 123456 |
| `includeFields` | String | no | Fields included in the query result: GuaranteesTerms - guarantee terms, NativeLogistics - logistics information, RateDetail - rating details, OrderInvoice - invoice information. By default, GuaranteesTerms, NativeLogistics, and OrderInvoice are returned. InvoicingSetting - invoicing settings | GuaranteesTerms,NativeLogistics,RateDetail,OrderInvoice |
| `attributeKeys` | String[] | no | attributeKeys in the vertical table | [] |
| `outOrderId` | String | no | External order ID, for idempotency control | 1556246 |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.openplatform.trade.model.TradeInfo](#m-alibaba-openplatform-trade-model-tradeinfo) | yes | Order detail info | {} |
| `errorCode` | String | yes | Error code |   |
| `errorMessage` | String | yes | Error description |   |
| `success` | String | yes | Whether successful | true |

<a id="m-alibaba-openplatform-trade-model-tradeinfo"></a>
#### alibaba.openplatform.trade.model.TradeInfo

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `baseInfo` | [message:alibaba.openplatform.trade.model.OrderBaseInfo](#m-alibaba-openplatform-trade-model-orderbaseinfo) | yes | Basic order information | {} |
| `orderBizInfo` | [message:alibaba.order.bizInfo](#m-alibaba-order-bizinfo) | yes | Order business info | {} |
| `tradeTerms` | [message:alibaba.openplatform.trade.model.TradeTermsInfo[]](#m-alibaba-openplatform-trade-model-tradetermsinfo[]) | yes | Trade terms | {} |
| `productItems` | [message:alibaba.openplatform.trade.model.ProductItemInfo[]](#m-alibaba-openplatform-trade-model-productiteminfo[]) | yes | Product entry information | {} |
| `nativeLogistics` | [message:alibaba.openplatform.trade.model.NativeLogisticsInfo](#m-alibaba-openplatform-trade-model-nativelogisticsinfo) | yes | Domestic logistics | {} |
| `orderInvoiceInfo` | [message:alibaba.invoice.OrderInvoiceModel](#m-alibaba-invoice-orderinvoicemodel) | yes | Invoice information | {} |
| `guaranteesTerms` | [message:alibaba.openplatform.trade.model.GuaranteeTermsInfo](#m-alibaba-openplatform-trade-model-guaranteetermsinfo) | yes | Protection terms | {} |
| `orderRateInfo` | [message:alibaba.trade.OrderRateInfo](#m-alibaba-trade-orderrateinfo) | yes | Order review information | {} |
| `overseasExtraAddress` | [message:alibaba.trade.OverseasExtraAddress](#m-alibaba-trade-overseasextraaddress) | yes | Cross-border address extended information | {} |
| `customs` | [message:alibaba.trade.Customs](#m-alibaba-trade-customs) | yes | Cross-border customs declaration information | {} |
| `quoteList` | [message:alibaba.orderDetail.caigouQuoteInfo[]](#m-alibaba-orderdetail-caigouquoteinfo[]) | yes | Purchase order detail list; a field exclusive to large-enterprise procurement orders. | {} |
| `extAttributes` | [message:alibaba.openplatform.trade.KeyValuePair[]](#m-alibaba-openplatform-trade-keyvaluepair[]) | yes | Order extended attributes | {} |
| `fromEncryptOrder` | Boolean | yes | Whether the order was created with downstream desensitized information | true |
| `encryptOutOrderInfo` | [message:alibaba.trade.get.sellerView.tradeinfo.EncryptOutOrderInfo](#m-alibaba-trade-get-sellerview-tradeinfo-encryptoutorderinfo) | yes | External order information | {} |
| `overseaLogisticsInfo` | [message:com.alibaba.ocean.openplatform.biz.trade.common.model.OverseaLogisticsInfo](#m-com-alibaba-ocean-openplatform-biz-trade-common-model-oversealogisticsinfo) | yes | Overseas logistics information | {} |
| `invoicingSettingModel` | [message:com.alibaba.ocean.openplatform.biz.trade.common.model.OpSellerInvoiceTradeSettingModel](#m-com-alibaba-ocean-openplatform-biz-trade-common-model-opsellerinvoicetradesettingmodel) | yes | Invoicing settings | {} |

<a id="m-alibaba-openplatform-trade-model-orderbaseinfo"></a>
#### alibaba.openplatform.trade.model.OrderBaseInfo

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `allDeliveredTime` | java.util.Date | yes | Complete shipment time | 20180614101942000+0800 |
| `sellerCreditLevel` | java.lang.String | yes | Seller credibility rating | L1 |
| `payTime` | java.util.Date | yes | Payment time; if there were multiple payments, this returns the time of the first payment | 20180614101942000+0800 |
| `discount` | Long | yes | Discount information, in cents (fen) | 11 |
| `alipayTradeId` | java.lang.String | yes | External payment transaction ID | 123123121111 |
| `sumProductPayment` | java.math.BigDecimal | yes | Total product amount (the sum of the product amounts in the order's product detail table), unit: yuan | 1212 |
| `buyerFeedback` | String | yes | Buyer's message, no more than 500 characters | 留言 |
| `flowTemplateCode` | String | yes | 4.0 transaction process template code | flow |
| `sellerOrder` | java.lang.Boolean | yes | Whether it is a self-initiated order (invitation order) | false |
| `buyerLoginId` | java.lang.String | yes | Buyer's loginId, WangWang ID | alitestforusv01 |
| `modifyTime` | java.util.Date | yes | Modification time | 20180614101942000+0800 |
| `subBuyerLoginId` | String | yes | Buyer sub-account | alitestforusv02:temp |
| `id` | java.lang.Long | yes | Transaction ID | 1231231231111 |
| `closeReason` | java.lang.String | yes | Close reason. buyerCancel: buyer cancelled the order, sellerGoodsLack: seller out of stock, other: other | buyerCancel |
| `buyerContact` | [message:alibaba.trade.tradeContact](#m-alibaba-trade-tradecontact) | yes | Buyer contact person | {} |
| `sellerAlipayId` | java.lang.String | yes | Seller's Alipay ID | 12312311111 |
| `completeTime` | java.util.Date | yes | Completion time | 20180614101942000+0800 |
| `sellerLoginId` | java.lang.String | yes | Seller oginId, WangWang ID | alitestforusv02 |
| `buyerID` | java.lang.String | yes | Buyer's main account id | 1234531 |
| `closeOperateType` | String | yes | Order closing operation type. CLOSE_TRADE_BY_SELLER: order closed by seller, CLOSE_TRADE_BY_BOPS: order closed by BOPS backend, CLOSE_TRADE_BY_SYSTEM: order closed by system (timeout), CLOSE_TRADE_BY_BUYER: order closed by buyer, CLOSE_TRADE_BY_CREADIT: closed due to trust & safety guarantee complaint | CLOSE_TRADE_BY_SELLER |
| `totalAmount` | java.math.BigDecimal | yes | Total payable amount, totalAmount = ∑itemAmount + shippingFee, in yuan | 1000 |
| `sellerID` | java.lang.String | yes | Seller's main account id | 123123123123 |
| `shippingFee` | java.math.BigDecimal | yes | Shipping fee, in yuan | 1 |
| `buyerUserId` | java.lang.Long | yes | Buyer numeric ID | 12314144 |
| `buyerMemo` | java.lang.String | yes | Buyer memo info | 备忘 |
| `refund` | java.math.BigDecimal | yes | Refund amount, in yuan | 1 |
| `status` | java.lang.String | yes | Transaction status: waitbuyerpay: waiting for buyer payment; waitsellersend: waiting for seller to ship; waitbuyerreceive: waiting for buyer to receive goods; confirm_goods: goods received; success: transaction successful; cancel: transaction canceled; terminated: transaction terminated; not enumerated: other status | waitbuyerpay |
| `refundPayment` | Long | yes | Refund amount | 1 |
| `sellerContact` | [message:alibaba.trade.tradeSellerContact](#m-alibaba-trade-tradesellercontact) | yes | Seller contact information | {} |
| `couponFee` | java.math.BigDecimal | yes | Red packet amount; the paid amount (totalAmount) has already accounted for the red packet amount | 7.5 |
| `buyerRemarkIcon` | String | yes | Buyer memo flag | 1 |
| `receiverInfo` | [message:alibaba.trade.orderReceiverInfo](#m-alibaba-trade-orderreceiverinfo) | yes | Recipient info | {} |
| `refundStatus` | String | yes | The order's in-sale refund status: waiting for seller to agree: waitselleragree, pending buyer modification: waitbuyermodify, waiting for buyer to return goods: waitbuyersend, waiting for seller to confirm receipt: waitsellerreceive, refund successful: refundsuccess, refund failed: refundclose | refundclose |
| `remark` | java.lang.String | yes | Remark; on 1688 this refers to the remark entered when placing the order | 备注 |
| `preOrderId` | java.lang.Long | yes | Pre-order ID | 123123 |
| `confirmedTime` | java.util.Date | yes | Confirmation time | 20180614101942000+0800 |
| `closeRemark` | String | yes | Order closing remark | 备注 |
| `tradeType` | String | yes | 1: Escrow (guaranteed) transaction<br>2: Pre-deposit transaction<br>3: ETC overseas acquiring transaction<br>4: Instant payment transaction<br>5: Security deposit protected transaction<br>6: Unified transaction process<br>7: Staged payment<br>8. Cash on delivery transaction<br>9. Credit voucher payment transaction<br>10. Account period payment transaction, 50060 Trade 4.0 | 50060 |
| `receivingTime` | java.util.Date | yes | Receipt time; the time returned here is the complete receipt time | 20180614101942000+0800 |
| `stepAgreementPath` | java.lang.String | yes | Staged legal agreement address |   |
| `idOfStr` | String | yes | Transaction id (string format) | 123121212123 |
| `refundStatusForAs` | String | yes | Order's after-sales refund status |   |
| `stepPayAll` | java.lang.Boolean | yes | Whether it is a one-time payment | false |
| `sellerUserId` | java.lang.Long | yes | Seller numeric ID | 12312422 |
| `stepOrderList` | [message:alibaba.trade.StepOrderModel[]](#m-alibaba-trade-stepordermodel[]) | yes | [Trade 3.0] Staged transaction, staged order list |   |
| `newStepOrderList` | [message:alibaba.trade.BizNewStepOrderModel[]](#m-alibaba-trade-biznewstepordermodel[]) | yes | [Transaction 4.0] Staged transaction, list of staged orders |   |
| `buyerAlipayId` | java.lang.String | yes | Buyer's Alipay id | 12312311233 |
| `createTime` | java.util.Date | yes | Creation time | 20180614101942000+0800 |
| `businessType` | java.lang.String | yes | Business type. International site: ta (Trade Assurance), wholesale (online wholesale).<br>China site: regular order type = &quot;cn&quot;;<br>large-value wholesale order type = &quot;ws&quot;;<br>regular sample order type = &quot;yp&quot;;<br>one-cent sample order type = &quot;yf&quot;;<br>reverse wholesale (limited-time discount) order type = &quot;fs&quot;;<br>processing/customization order type = &quot;cz&quot;;<br>agreement procurement order type = &quot;ag&quot;;<br>group-buy (huopin) order type = &quot;hp&quot;;<br>supply-distribution order type = &quot;supply&quot;;<br>Taobao Factory order = &quot;factory&quot;;<br>quick order placement = &quot;quick&quot;;<br>Xiangpin order = &quot;xiangpin&quot;;<br>face-to-face payment = &quot;f2f&quot;;<br>sample storage service = &quot;cyfw&quot;;<br>consignment order = &quot;sp&quot;;<br>Weigong order = &quot;wg&quot;; Lingshoutong (Retail Link) = &quot;lst&quot;; cross-border = 'cb'; distribution = 'distribution'; Cai Yuan Bao = 'cab'; processing/customization = &quot;manufact&quot; | cn |
| `overSeaOrder` | java.lang.Boolean | yes | Whether it is an overseas dropshipping order; yes: true | true |
| `refundId` | java.lang.String | yes | Refund order ID | TQ4562212313 |
| `tradeTypeDesc` | String | yes | Transaction method specified when placing the order | 担保交易 |
| `payChannelList` | String[] | yes | List of payment channel names. An order may involve multiple payment channels. Enum values: Alipay, MYbank Trust Pay, Cheng-e-She, corporate bank transfer, credit sale (Shexiaobao), account period payment, combined payment channels, payment platform, declared payment, MYbank electronic bank acceptance draft, bank transfer, Kuajingbao, red packet, other | ["支付宝","跨境宝","银行转账"] |
| `tradeTypeCode` | String | yes | The tradeType (transaction method) specified when placing the order | assureTrade |
| `payTimeout` | Long | yes | Payment timeout duration; when fixed-length, the unit is seconds - currently always fixed-length | 43200 |
| `payTimeoutType` | Integer | yes | Payment timeout TYPE, 0: fixed length, 1: fixed time | 0 |
| `payChannelCodeList` | String[] | yes | Payment channel code; see payChannelList for the Chinese description of payChannelCodeList | ["alipay"] |
| `inventoryMode` | String | yes | Supply stock mode, jit (JIT mode) or cang (warehouse fulfillment mode) | jit |
| `outOrderId` | String | yes | External order number | 1928919827731 |
| `officialSolutionOrderId` | String | yes | Official logistics pickup order ID | 123 |
| `officialSolutionCost` | Long | yes | Official logistics pickup order fee | 770 |

<a id="m-alibaba-trade-tradecontact"></a>
#### alibaba.trade.tradeContact

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `phone` | String | yes | Contact phone |  |
| `fax` | String | yes | Fax |  |
| `email` | String | yes | Email |  |
| `imInPlatform` | String | yes | The contact's IM account on the platform |  |
| `name` | String | yes | Contact name |  |
| `mobile` | String | yes | Contact person's mobile number |  |
| `companyName` | java.lang.String | yes | Company name |  |

<a id="m-alibaba-trade-tradesellercontact"></a>
#### alibaba.trade.tradeSellerContact

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `phone` | String | yes | Contact phone | 189982787712 |
| `fax` | String | yes | Fax | 189982787712 |
| `email` | String | yes | Email | xxx@mail.com |
| `imInPlatform` | String | yes | The contact's IM account on the platform | wangwang |
| `name` | String | yes | Contact name | 张三 |
| `mobile` | String | yes | Contact person's mobile number | 189982787712 |
| `companyName` | java.lang.String | yes | Company name | 三生网媒邮箱公司 |
| `wgSenderName` | java.lang.String | yes | Sender name, set by the distributor in distribution scenarios such as WeiGong (micro-supply) | 张** |
| `wgSenderPhone` | java.lang.String | yes | Sender's phone number; set by the distributor in distribution scenarios such as Weigong | 13800000000 |
| `shopName` | String | yes | Shop name |  三生网媒邮箱公司 |

<a id="m-alibaba-trade-orderreceiverinfo"></a>
#### alibaba.trade.orderReceiverInfo

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `toFullName` | java.lang.String | yes | Recipient |  |
| `toDivisionCode` | java.lang.String | yes | Recipient address region code |  |
| `toMobile` | java.lang.String | yes | Recipient's mobile phone number |  |
| `toPhone` | java.lang.String | yes | Recipient phone number |  |
| `toPost` | java.lang.String | yes | Postal code |  |
| `toTownCode` | java.lang.String | yes | Recipient's street or town region code, may be empty |  |
| `toArea` | java.lang.String | yes | Shipping address |  |

<a id="m-alibaba-trade-stepordermodel[]"></a>
#### alibaba.trade.StepOrderModel[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `stepOrderId` | java.lang.Long | yes | Stage id |  |
| `stepOrderStatus` | java.lang.String | yes | waitactivate  Not started (pending activation)<br>waitsellerpush Waiting for seller to proceed<br>success This stage completed<br>settlebill Settlement (split payment)<br>cancel This stage terminated<br>inactiveandcancel This stage terminated before starting<br>waitbuyerpay Waiting for buyer payment<br>waitsellersend Waiting for seller to ship<br>waitbuyerreceive Waiting for buyer to confirm receipt<br>waitselleract Waiting for seller to perform XX action<br>waitbuyerconfirmaction Waiting for buyer to confirm XX action |  |
| `stepPayStatus` | java.lang.Integer | yes | 1 Not frozen/not paid<br>2 Frozen/paid<br>4 Refunded<br>6 Transferred transaction<br>8 Transaction closed due to non-payment |  |
| `stepNo` | java.lang.Integer | yes | Stage sequence: 1, 2, 3... |  |
| `lastStep` | java.lang.Boolean | yes | Whether it is the last stage |  |
| `hasDisbursed` | java.lang.Boolean | yes | Whether payment has been made to the seller |  |
| `payFee` | java.math.BigDecimal | yes | The amount payable at creation, excluding shipping fee |  |
| `actualPayFee` | java.math.BigDecimal | yes | Amount payable (including shipping fee) = unit price x quantity - item discount - shop discount + shipping fee + adjusted amount (except for shipping fee, all amounts refer to the allocated amount) |  |
| `discountFee` | java.math.BigDecimal | yes | Store discount allocated for this stage |  |
| `itemDiscountFee` | java.math.BigDecimal | yes | Item discount allocated to this stage |  |
| `price` | java.math.BigDecimal | yes | Unit price allocated to this stage |  |
| `amount` | java.lang.Long | yes | Purchase quantity |  |
| `postFee` | java.math.BigDecimal | yes | Shipping fee |  |
| `adjustFee` | java.math.BigDecimal | yes | Amount changed by the price modification |  |
| `gmtCreate` | java.util.Date | yes | Creation time |  |
| `gmtModified` | java.util.Date | yes | Modification time |  |
| `enterTime` | java.util.Date | yes | Start time |  |
| `payTime` | java.util.Date | yes | Payment time |  |
| `sellerActionTime` | java.util.Date | yes | Seller's operation time |  |
| `endTime` | java.util.Date | yes | End time of this stage |  |
| `messagePath` | java.lang.String | yes | Path of the seller's operation message |  |
| `picturePath` | java.lang.String | yes | Path of the image evidence uploaded by the seller |  |
| `message` | java.lang.String | yes | Seller operation message |  |
| `templateId` | java.lang.Long | yes | Template ID used |  |
| `stepName` | java.lang.String | yes | Name of the current step |  |
| `sellerActionName` | java.lang.String | yes | Seller operation name |  |
| `buyerPayTimeout` | java.lang.Long | yes | Timeout for buyer non-payment (seconds) |  |
| `buyerConfirmTimeout` | java.lang.Long | yes | Timeout for buyer non-confirmation |  |
| `needLogistics` | java.lang.Boolean | yes | Whether logistics is required |  |
| `needSellerAction` | java.lang.Boolean | yes | Whether seller action and buyer confirmation are required |  |
| `transferAfterConfirm` | java.lang.Boolean | yes | Whether payment is disbursed when the stage ends |  |
| `needSellerCallNext` | java.lang.Boolean | yes | Whether seller action is required to proceed |  |
| `instantPay` | java.lang.Boolean | yes | Whether instant crediting is allowed |  |

<a id="m-alibaba-trade-biznewstepordermodel[]"></a>
#### alibaba.trade.BizNewStepOrderModel[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `gmtStart` | java.util.Date | yes | Stage start time | 20180604092517000+0800 |
| `gmtPay` | java.util.Date | yes | Payment time | 20180604093243000+0800 |
| `gmtEnd` | java.util.Date | yes | Stage end time | 20180604093243000+0800 |
| `stepNo` | Integer | yes | Stage sequence number | 1 |
| `lastStep` | Boolean | yes | Whether it is the last stage | true |
| `stepName` | String | yes | Stage name | 全款交易 |
| `activeStatus` | Integer | yes | Activation status. 0 indicates not activated, 1 indicates activated. | 1 |
| `payStatus` | Integer | yes | Stage payment status. 1 unpaid, 2 paid, 8 cancelled before payment, 12 overage/shortage supplementary payment | 2 |
| `logisticsStatus` | Integer | yes | Logistics stage status: 1 not shipped, 2 shipped, 3 received, 4 fully returned, 7 canceled before shipment | 2 |
| `payFee` | java.math.BigDecimal | yes | Amount payable for the stage (including shipping fee), in yuan | 0.03 |
| `paidFee` | java.math.BigDecimal | yes | Amount paid for this stage (including shipping), in yuan | 0.03 |
| `goodsFee` | java.math.BigDecimal | yes | Stage product price allocation, in yuan | 0 |
| `adjustFee` | java.math.BigDecimal | yes | Stage adjusted price, in yuan | -3175.97 |
| `discountFee` | java.math.BigDecimal | yes | Discounted price for this stage, in yuan | 0 |
| `postFee` | java.math.BigDecimal | yes | Shipping fee payable for this stage, in yuan | 0 |
| `paidPostFee` | java.math.BigDecimal | yes | Shipping fee already paid for the stage, in yuan | 0 |

<a id="m-alibaba-order-bizinfo"></a>
#### alibaba.order.bizInfo

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `odsCyd` | java.lang.Boolean | yes | Whether it is a Caiyuanbao order | true |
| `accountPeriodTime` | java.lang.String | yes | Payment arrival time for account period transaction orders | yyyy-MM-dd HH:mm:ss |
| `creditOrder` | java.lang.Boolean | yes | If true, indicates that the Cheng-e-She transaction method was selected when placing the order. Note this is not the same as "Cheng-e-She payment" — the actual payment may be made via Alipay; check tradeTerms.payWay for the specific payment method | false |
| `creditOrderDetail` | [message:alibaba.creditOrder.forDetail](#m-alibaba-creditorder-fordetail) | yes | Cheng-e-She payment details; returned only when paying via Cheng-e-She |   |
| `preOrderInfo` | [message:alibaba.order.preOrder.forRead](#m-alibaba-order-preorder-forread) | yes | Pre-order information | {} |
| `lstOrderInfo` | [message:alibaba.lst.tradeInfo](#m-alibaba-lst-tradeinfo) | yes | Lingshoutong (Retail Link) order information | {} |
| `erpBuyerUserId` | String | yes | ERP user ID | U001012121 |
| `erpOrderId` | String | yes | ERP order number | O123331 |
| `erpBuyerOrgId` | String | yes | ERP organization ID | OG4331113 |
| `isCz` | Boolean | yes | Whether it is a processing/customization order | false |
| `isDz` | Boolean | yes | Whether it is a customized order | false |
| `dz` | Boolean | yes | Whether it is a customized order | false |
| `dropshipping` | Boolean | yes | Whether it is a dropshipping order; this type of order does not allow combined shipping | true |
| `shippingInsurance` | String | yes | givenByPlatform: shipping insurance provided by the platform; givenByMerchant: shipping insurance provided by the merchant; empty means the order has no shipping insurance | givenByPlatform |
| `hyperLinkCangFaOrder` | Boolean | yes | Large-store warehouse-shipped order | true |
| `hyperLinkOrder` | Boolean | yes | Chaolian (super-link) first-stage order | flase |
| `hyperLinkSecondStepOrder` | Boolean | yes | The second-stage order of a Chaolian (super-link) flagship store two-stage order | true |
| `hyperLinkShipType` | String | yes | Shipping mode for Chaolian stage-one orders: 0 warehouse shipment, 1 merchant shipment | 0 |
| `lightningWarehouse` | Boolean | yes | Flash warehouse order | true |
| `aeDoorPickUp` | Boolean | yes | AE door-to-door pickup order | true |
| `fz` | Boolean | yes | Split-payment order | true |
| `tgOfficialPickUp` | Boolean | yes | Managed official door-to-door pickup order | true |
| `officialPickUp` | Boolean | yes | Official logistics pickup | true |
| `abnormalPriceChange` | Boolean | yes | Whether it is an abnormal price change | true |

<a id="m-alibaba-creditorder-fordetail"></a>
#### alibaba.creditOrder.forDetail

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `payAmount` | java.lang.Long | yes | Order amount | 10 |
| `createTime` | java.lang.String | yes | Payment time | 2018-01-01 00:00:00 |
| `status` | java.lang.String | yes | Status | END |
| `gracePeriodEndTime` | java.lang.String | yes | No longer recommended for use | 2018-01-01 00:00:00 |
| `statusStr` | java.lang.String | yes | Status description | 已完结 |
| `restRepayAmount` | java.lang.Long | yes | Amount due for repayment | 11 |
| `lastRepayTime` | String | yes | Latest repayment time | 2023-01-01 00:00:00 |
| `repaySource` | String | yes | Repayment source: KJPAY - Kuajingbao repayment, OWN_FUNDS - repayment from own funds, INSTALLMENT_REPAY - installment/loan repayment | KJPAY |

<a id="m-alibaba-order-preorder-forread"></a>
#### alibaba.order.preOrder.forRead

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `marketName` | String | yes | The market name passed in when creating the pre-order | dxc |
| `createPreOrderApp` | Boolean | yes | Whether the pre-order was created by the ERP used for the current query | false |

<a id="m-alibaba-lst-tradeinfo"></a>
#### alibaba.lst.tradeInfo

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `lstWarehouseType` | String | yes | Retail Tong (Lingshoutong) warehouse type. customer: virtual warehouse; cainiao: physical warehouse | cainiao |

<a id="m-alibaba-openplatform-trade-model-tradetermsinfo[]"></a>
#### alibaba.openplatform.trade.model.TradeTermsInfo[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `payStatus` | java.lang.String | yes | Payment status. International site: WAIT_PAY (unpaid), PAYER_PAID (payment completed), PART_SUCCESS (partial payment succeeded), PAY_SUCCESS (payment succeeded), CLOSED (closed due to risk control), CANCELLED (payment cancelled), SUCCESS (success), FAIL (failed).<br>1688: 1 (unpaid); 2 (paid); 4 (full refund); 6 (seller has received payment, payment collection completed); 7 (external payment order not created); 8 (cancelled before payment); 9 (payment in progress); 12 (account period payment, pending arrival) |   |
| `payTime` | java.util.Date | yes | Time when the stage payment was completed |   |
| `payWay` | java.lang.String | yes | Payment method.<br>International site: ECL (financing payment), CC (credit card), TT (offline TT), ACH (echecking payment).<br>1688: 1-Alipay, 2-MYbank Trusted Payment, 3-Cheng-e-She, 4-bank transfer, 5-Shexiaobao, 6-electronic acceptance bill, 7-account period payment, 8-combined payment channel, 9-no payment, 10-Lingshoutong credit purchase, 13-payment platform, 12-declared payment |   |
| `phasAmount` | java.math.BigDecimal | yes | Payment amount |   |
| `phase` | java.lang.Long | yes | Stage order id |   |
| `phaseCondition` | java.lang.String | yes | Stage condition; not applicable on 1688. |   |
| `phaseDate` | java.lang.String | yes | Stage time; not applicable to 1688 |   |
| `cardPay` | java.lang.Boolean | yes | Whether bank card payment is used |   |
| `expressPay` | java.lang.Boolean | yes | Whether quick payment is used |   |
| `payWayDesc` | String | yes | Payment method | 支付宝 |

<a id="m-alibaba-openplatform-trade-model-productiteminfo[]"></a>
#### alibaba.openplatform.trade.model.ProductItemInfo[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `cargoNumber` | java.lang.String | yes | Specifies the SKU (single item) item number; not applicable for the international site. This field does not always have a value — it is only recorded when the order is placed (if the seller has set an item number for the SKU). For other order types, the item number can only be obtained via the product API. Note: the item number obtained via the product API and the item number at the time of order placement may differ, because the seller may modify the product information after the order is completed, changing the item number. |   |
| `description` | java.lang.String | yes | Description; not applicable to 1688 |   |
| `itemAmount` | java.math.BigDecimal | yes | Actual amount paid, in yuan |   |
| `name` | java.lang.String | yes | Product name |   |
| `price` | java.math.BigDecimal | yes | Original unit price, in yuan |   |
| `productID` | java.lang.Long | yes | Product ID (empty for non-online products) |   |
| `productImgUrl` | String[] | yes | Product image URL |   |
| `productSnapshotUrl` | java.lang.String | yes | Product snapshot URL. When a trade order is created, a snapshot of the product at that time is automatically recorded for reference in the event of a subsequent dispute. |   |
| `quantity` | java.math.BigDecimal | yes | Quantity in units of unit, e.g. how many pieces, items, boxes, or tons |   |
| `refund` | java.math.BigDecimal | yes | Refund amount, in yuan |   |
| `skuID` | java.lang.Long | yes | skuID |   |
| `sort` | java.lang.Integer | yes | Sort field. The product list is sorted by this field, starting from 0; not provided by 1688. |   |
| `status` | java.lang.String | yes | Sub-order status |   |
| `subItemID` | java.lang.Long | yes | Sub-order number, or product detail line item ID |   |
| `type` | java.lang.String | yes | Type, used on the international site, for sellers to mark the type the product belongs to |   |
| `unit` | java.lang.String | yes | Selling unit	E.g.: piece, item, box, ton |   |
| `weight` | java.lang.String | yes | Weight	Weight calculated in weight units, e.g.: 100 |   |
| `weightUnit` | java.lang.String | yes | Weight unit	e.g. g, kg, t |   |
| `guaranteesTerms` | [message:alibaba.openplatform.trade.model.GuaranteeTermsInfo[]](#m-alibaba-openplatform-trade-model-guaranteetermsinfo[]) | yes | Guarantee terms; this field applies only to 1688 |   |
| `productCargoNumber` | java.lang.String | yes | Specifies the product item number. This field does not always have a value — it is only recorded when the order is placed. For other order types, the item number can only be obtained via the product API. Note: the item number obtained via the product API and the item number at the time of order placement may differ, because the seller may modify the product information after the order is completed, changing the item number. The difference between this field and cargoNUmber is: this field is the item number defined at the product level, while cargoNUmber is the item number defined at the SKU (single item) level. |   |
| `skuInfos` | [message:alibaba.trade.SkuItemDesc[]](#m-alibaba-trade-skuitemdesc[]) | yes |  |    |
| `entryDiscount` | Long | yes | Amount of price increase or decrease in the order line item |   |
| `specId` | java.lang.String | yes | Order sales attribute ID |   |
| `quantityFactor` | java.math.BigDecimal | yes | The quantity precision factor in units of unit; the value is a power of 10. For example: quantityFactor=1000, unit=ton, then the minimum precision of quantity is 0.001 ton. |   |
| `statusStr` | java.lang.String | yes | Sub-order status description |   |
| `refundStatus` | java.lang.String | yes | WAIT_SELLER_AGREE Waiting for seller to agree<br>REFUND_SUCCESS Refund successful<br>REFUND_CLOSED Refund closed<br>WAIT_BUYER_MODIFY Pending buyer modification<br>WAIT_BUYER_SEND Waiting for buyer to return goods<br>WAIT_SELLER_RECEIVE Waiting for seller to confirm receipt |   |
| `closeReason` | java.lang.String | yes | Closure reason |   |
| `logisticsStatus` | java.lang.Integer | yes | 1 Not shipped<br>2 Shipped<br>3 Received<br>4 Returned<br>5 Partially shipped<br>8 Logistics order not yet created |   |
| `gmtCreate` | java.util.Date | yes | Creation time |   |
| `gmtModified` | java.util.Date | yes | Modification time |   |
| `gmtCompleted` | java.util.Date | yes | Line item completion time |   |
| `gmtPayExpireTime` | String | yes | Stock timeout time, in the format "yyyy-MM-dd HH:mm:ss" |   |
| `refundId` | String | yes | In-sale refund order number |   |
| `subItemIDString` | String | yes | Sub-order number, or product detail line item ID (string type; since a Long-type ID may cause processing issues in JS and PHP, it can be handled as a string type) |   |
| `refundIdForAs` | String | yes | After-sales refund order number |   |
| `sharePostage` | BigDecimal | yes | Shared shipping fee, in yuan |   |

<a id="m-alibaba-openplatform-trade-model-guaranteetermsinfo[]"></a>
#### alibaba.openplatform.trade.model.GuaranteeTermsInfo[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `assuranceInfo` | java.lang.String | yes | Protection terms | 自愿选择向买家提供“交期保障”服务 |
| `assuranceType` | java.lang.String | yes | Guarantee method. International site: TA (Trade Assurance) | jqbz |
| `qualityAssuranceType` | java.lang.String | yes | Quality assurance type. International site: pre_shipment (before shipment), post_delivery (after delivery) | 交期保障 |
| `value` | String | yes | Guarantee term value; for example, in a delivery time guarantee, 6 represents 6 days | 6 |

<a id="m-alibaba-trade-skuitemdesc[]"></a>
#### alibaba.trade.SkuItemDesc[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `name` | String | yes | Attribute name |  |
| `value` | String | yes | Attribute value |  |

<a id="m-alibaba-openplatform-trade-model-nativelogisticsinfo"></a>
#### alibaba.openplatform.trade.model.NativeLogisticsInfo

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `address` | java.lang.String | yes | Detailed address |  |
| `area` | java.lang.String | yes | County, district |  |
| `areaCode` | java.lang.String | yes | Province/city/district code |  |
| `city` | java.lang.String | yes | City |  |
| `contactPerson` | java.lang.String | yes | Contact person name |  |
| `fax` | java.lang.String | yes | Fax |  |
| `mobile` | java.lang.String | yes | Mobile phone |  |
| `province` | java.lang.String | yes | Province |  |
| `telephone` | java.lang.String | yes | Phone number |  |
| `zip` | java.lang.String | yes | Postal code |  |
| `logisticsItems` | [message:alibaba.openplatform.trade.model.NativeLogisticsItemsInfo[]](#m-alibaba-openplatform-trade-model-nativelogisticsitemsinfo[]) | yes | Waybill details |  |
| `townCode` | java.lang.String | yes | Town/street address code |  |
| `town` | java.lang.String | yes | Town, street |  |

<a id="m-alibaba-openplatform-trade-model-nativelogisticsitemsinfo[]"></a>
#### alibaba.openplatform.trade.model.NativeLogisticsItemsInfo[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `deliveredTime` | java.util.Date | yes | Shipping time |   |
| `logisticsCode` | java.lang.String | yes | Logistics number |   |
| `type` | java.lang.String | yes | SELF_SEND_GOODS (&quot;0&quot;) self-delivery, online delivery ONLINE_SEND_GOODS (&quot;1&quot;), delivery without logistics NO_LOGISTICS_SEND_GOODS (&quot;2&quot;) |   |
| `id` | java.lang.Long | yes | Primary key id |   |
| `status` | java.lang.String | yes | Status |   |
| `gmtModified` | java.util.Date | yes | Modification time |   |
| `gmtCreate` | java.util.Date | yes | Creation time |   |
| `carriage` | java.math.BigDecimal | yes | Shipping fee (in yuan) |   |
| `fromProvince` | java.lang.String | yes | Shipping province |   |
| `fromCity` | java.lang.String | yes | Shipping city |   |
| `fromArea` | java.lang.String | yes | Shipping district |   |
| `fromAddress` | java.lang.String | yes | Shipping street address |   |
| `fromPhone` | java.lang.String | yes | Shipping contact phone number |   |
| `fromMobile` | java.lang.String | yes | Shipping contact mobile number |   |
| `fromPost` | java.lang.String | yes | Shipping address postal code |   |
| `logisticsCompanyId` | java.lang.Long | yes | Logistics company Id |   |
| `logisticsCompanyNo` | java.lang.String | yes | Logistics company number |   |
| `logisticsCompanyName` | java.lang.String | yes | Logistics company name |   |
| `logisticsBillNo` | java.lang.String | yes | Logistics company waybill number |   |
| `subItemIds` | java.lang.String | yes | Product line item ID; if multiple, separated by commas |   |
| `toProvince` | java.lang.String | yes | Receiving province |   |
| `toCity` | java.lang.String | yes | Receiving city |   |
| `toArea` | java.lang.String | yes | Receiving district |   |
| `toAddress` | java.lang.String | yes | Receiving street address |   |
| `toPhone` | java.lang.String | yes | Recipient contact phone number |   |
| `toMobile` | java.lang.String | yes | Receiving contact mobile number |   |
| `toPost` | java.lang.String | yes | Shipping address postal code |   |
| `noLogisticsName` | String | yes | Logistics contact name |   |
| `noLogisticsTel` | String | yes | Contact information |   |
| `noLogisticsBillNo` | String | yes | Business order number for no-logistics-required |   |
| `noLogisticsCondition` | String | yes | No-logistics-required category. noLogisticsCondition=1 means other third-party logistics, small logistics providers, fleets, etc.; noLogisticsCondition=2 means shipping fee/price difference supplement; noLogisticsCondition=3 means seller delivery; noLogisticsCondition=4 means buyer self-pickup; noLogisticsCondition=5 means other reasons |   |

<a id="m-alibaba-invoice-orderinvoicemodel"></a>
#### alibaba.invoice.OrderInvoiceModel

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `invoiceCompanyName` | java.lang.String | yes | Invoice company name (i.e. the invoice title) |   |
| `invoiceType` | java.lang.Integer | yes | Invoice type. 0: general invoice, 1: VAT invoice, 9: unknown type |   |
| `localInvoiceId` | java.lang.Long | yes | Local invoice number |   |
| `orderId` | java.lang.Long | yes | Order ID |   |
| `receiveCode` | java.lang.String | yes | (Recipient) address region code |   |
| `receiveCodeText` | java.lang.String | yes | (Recipient) text corresponding to the province/city/district code (VAT invoice information) |   |
| `receiveMobile` | java.lang.String | yes | (Recipient) invoice recipient's mobile phone number |   |
| `receiveName` | java.lang.String | yes | (Recipient) invoice recipient |   |
| `receivePhone` | java.lang.String | yes | (Recipient) invoice recipient's phone number |   |
| `receivePost` | java.lang.String | yes | (Recipient) invoice shipping address postal code |   |
| `receiveStreet` | java.lang.String | yes | (Recipient) street address (VAT invoice information) |   |
| `registerAccountId` | java.lang.String | yes | (Company) bank account number |   |
| `registerBank` | java.lang.String | yes | (Company) bank of deposit |   |
| `registerCode` | java.lang.String | yes | (Registration) province/city/district code |   |
| `registerCodeText` | java.lang.String | yes | (Registration) province/city/district text |   |
| `registerPhone` | java.lang.String | yes | (Company) registered phone number |   |
| `registerStreet` | java.lang.String | yes | (Registered) street address |   |
| `taxpayerIdentify` | java.lang.String | yes | Taxpayer identification number |   |
| `amount` | Long | yes | Invoice total amount including tax (cents/fen) | 1 |
| `status` | String | yes | Invoice status: {&quot;ISSUED&quot;:&quot;Issued&quot;,&quot;RED_ISSUING&quot;:&quot;Red-flush in progress&quot;,&quot;RED_ALL_ISSUED&quot;:&quot;Fully red-flushed&quot;,&quot;CLOSED&quot;:&quot;Closed&quot;,&quot;RED_PART_ISSUED&quot;:&quot;Partially red-flushed&quot;,&quot;VERIFYING&quot;:&quot;Verifying&quot;,&quot;DEPRECATED&quot;:&quot;Voided&quot;,&quot;RETURNING&quot;:&quot;Return in progress&quot;,&quot;INVALIDING&quot;:&quot;Voiding&quot;,&quot;INIT&quot;:&quot;Initialized&quot;,&quot;VERIFY_FAILED&quot;:&quot;Verification failed&quot;,&quot;ISSUING&quot;:&quot;Issuing&quot;,&quot;FAILED&quot;:&quot;Issuance failed&quot;,&quot;RETURNED&quot;:&quot;Returned&quot;} | issued |

<a id="m-alibaba-openplatform-trade-model-guaranteetermsinfo"></a>
#### alibaba.openplatform.trade.model.GuaranteeTermsInfo

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `assuranceInfo` | java.lang.String | yes | Protection terms | 自愿选择向买家提供“交期保障”服务 |
| `assuranceType` | java.lang.String | yes | Guarantee method. International site: TA (Trade Assurance) | jqbz |
| `qualityAssuranceType` | java.lang.String | yes | Quality assurance type. International site: pre_shipment (before shipment), post_delivery (after delivery) | 交期保障 |
| `value` | String | yes | Guarantee term value; for example, in a delivery time guarantee, 6 represents 6 days | 6 |

<a id="m-alibaba-trade-orderrateinfo"></a>
#### alibaba.trade.OrderRateInfo

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `buyerRateStatus` | Integer | yes | Buyer review status (4: reviewed, 5: not reviewed, 6: review not required) |  |
| `sellerRateStatus` | Integer | yes | Seller review status (4: reviewed, 5: not reviewed, 6: review not required) |  |
| `buyerRateList` | [message:alibaba.order.rateDetail[]](#m-alibaba-order-ratedetail[]) | yes | Seller's rating of the buyer |  |
| `sellerRateList` | [message:alibaba.order.rateDetail[]](#m-alibaba-order-ratedetail[]) | yes | Buyer's rating of the seller |  |

<a id="m-alibaba-order-ratedetail[]"></a>
#### alibaba.order.rateDetail[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `starLevel` | java.lang.Integer | yes | Review star rating |  |
| `content` | java.lang.String | yes | Review details |  |
| `receiverNick` | java.lang.String | yes | Nickname of the user who received the review |  |
| `posterNick` | java.lang.String | yes | Nickname of the user who submitted the review |  |
| `publishTime` | java.util.Date | yes | Review publish time |  |

<a id="m-alibaba-trade-overseasextraaddress"></a>
#### alibaba.trade.OverseasExtraAddress

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `channelName` | String | yes | Route name | 欧洲小包 |
| `channelId` | String | yes | Route id | 1 |
| `shippingCompanyId` | String | yes | Freight forwarder company ID | 222 |
| `shippingCompanyName` | String | yes | Freight forwarder company name | 货代公司1 |
| `countryCode` | String | yes | Country code | UK |
| `country` | String | yes | Country | 英国 |
| `email` | String | yes | Buyer email | aaa@gmail.com |

<a id="m-alibaba-trade-customs"></a>
#### alibaba.trade.Customs

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `id` | Long | yes | id | 1 |
| `gmtCreate` | java.util.Date | yes | Creation time | 20170806114526000+0800 |
| `gmtModified` | java.util.Date | yes | Modification time | 20170806114526000+0800 |
| `buyerId` | Long | yes | Buyer ID | 123456 |
| `orderId` | String | yes | Main order id | 12312312312312 |
| `type` | Integer | yes | Business data type, default 1: customs declaration | 1 |
| `attributes` | [message:alibaba.trade.CustomsAttributesInfo[]](#m-alibaba-trade-customsattributesinfo[]) | yes | List of customs declaration information |  |

<a id="m-alibaba-trade-customsattributesinfo[]"></a>
#### alibaba.trade.CustomsAttributesInfo[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `sku` | String | yes | SKU identifier | 1234 |
| `cName` | String | yes | Chinese name | 测试 |
| `enName` | String | yes | English name | test |
| `amount` | Double | yes | Declared value | 3000.0 |
| `quantity` | Double | yes | Quantity | 1.0 |
| `weight` | Double | yes | Weight (kg) | 0.5 |
| `currency` | String | yes | Customs declaration currency | CNY |

<a id="m-alibaba-orderdetail-caigouquoteinfo[]"></a>
#### alibaba.orderDetail.caigouQuoteInfo[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `productQuoteName` | String | yes | Name of the supply line item | 物料01 |
| `price` | java.math.BigDecimal | yes | Price, in yuan | 100 |
| `count` | Double | yes | Purchase quantity | 10 |

<a id="m-alibaba-openplatform-trade-keyvaluepair[]"></a>
#### alibaba.openplatform.trade.KeyValuePair[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `key` | String | yes | Key |  |
| `value` | String | yes | Value |  |
| `description` | String | yes | Description |  |

<a id="m-alibaba-trade-get-sellerview-tradeinfo-encryptoutorderinfo"></a>
#### alibaba.trade.get.sellerView.tradeinfo.EncryptOutOrderInfo

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `outPlatformOrderNo` | String | yes | External order number | 25662254541 |
| `outPlatformCode` | String | yes | Downstream platform: Taobao-thyny, Tmall-tm, Taote-taote, Alibaba C2M-c2m, JD-jingdong, Pinduoduo-pinduoduo, WeChat-weixin, Cross-border-kuajing, Kuaishou-kuaishou, Youzan-youzan, Douyin-douyin, Siku-siku, Meituan Tuanhaohuo-meituan, Xiaohongshu-xiaohongshu, Dangdang-dangdang, Suning-suning, DaVdian-davdian, Xingyun-xingyun, Miya-miya, Boluopai Mall-boluo, Other-other | tm |
| `outPlatformAppkey` | String | yes | The appkey of the downstream platform used to obtain downstream order information | 65345 |
| `oaid` | String | yes | Taobao oaid | xxx-xxxx-xxx |
| `encryptReceiverName` | String | yes | Downstream encrypted recipient name | *** |
| `encryptReceiverMobile` | String | yes | Downstream encrypted recipient phone number | *** |
| `encryptReceiverAddress` | String | yes | Downstream encrypted recipient address | *** |
| `outPatformExtraInfo` | String | yes | Other extended information | {} |
| `outShopId` | String | yes | Downstream platform shopId | 21343123 |
| `outOriginAddress` | [message:com.alibaba.ocean.openplatform.biz.trade.param.OutAddress](#m-com-alibaba-ocean-openplatform-biz-trade-param-outaddress) | yes | External original address information | {} |
| `outShopName` | String | yes | Downstream platform shop name | 三生智能 |
| `outPlatformSubCode` | String | yes | Downstream channel sub-business code, e.g. Douyin sub-channel 101, used for supply chain order identification and for printing encrypted orders. | 101 |
| `outPlatformSupplyOrderNo` | String | yes | Downstream platform supply chain purchase order number, e.g. Douyin supply chain, etc. | CT7403712218870825259 |
| `outSupplierId` | String | yes | Downstream platform supply chain provider id | 1231123 |

<a id="m-com-alibaba-ocean-openplatform-biz-trade-param-outaddress"></a>
#### com.alibaba.ocean.openplatform.biz.trade.param.OutAddress

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `province` | [message:com.alibaba.ocean.openplatform.biz.trade.param.Place](#m-com-alibaba-ocean-openplatform-biz-trade-param-place) | yes | Province | {"name":"四川省","code":"51000"} |
| `city` | [message:com.alibaba.ocean.openplatform.biz.trade.param.Place](#m-com-alibaba-ocean-openplatform-biz-trade-param-place) | yes | City | {} |
| `area` | [message:com.alibaba.ocean.openplatform.biz.trade.param.Place](#m-com-alibaba-ocean-openplatform-biz-trade-param-place) | yes | District | {} |
| `town` | [message:com.alibaba.ocean.openplatform.biz.trade.param.Place](#m-com-alibaba-ocean-openplatform-biz-trade-param-place) | no | Town/street | {} |
| `address` | String | no | Detailed address | 网商路699号 |
| `postCode` | String | no | Postal code | 511304 |

<a id="m-com-alibaba-ocean-openplatform-biz-trade-param-place"></a>
#### com.alibaba.ocean.openplatform.biz.trade.param.Place

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `code` | String | yes | Address code | 511300 |
| `name` | String | yes | Address name | 南充 |

<a id="m-com-alibaba-ocean-openplatform-biz-trade-common-model-oversealogisticsinfo"></a>
#### com.alibaba.ocean.openplatform.biz.trade.common.model.OverseaLogisticsInfo

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `overseasTransportUserAddrInfo` | [message:com.alibaba.ocean.openplatform.biz.cross.model.OverseasTransportUserAddrInfo](#m-com-alibaba-ocean-openplatform-biz-cross-model-overseastransportuseraddrinfo) | yes | Overseas shipping address | "addressDetail":"test", 						"cityName":"Г.Курчатов", 						"areaName":" ", 						"warehouseContactName":"Zeus", 						"mobile":"123456789", 						"fixedPhone":"123456789", 						"postCode":"1111", 						"provinceName":"Абайская область", 						"countryName":"Казахстан", |
| `overseasLogisticsIds` | String[] | yes | Overseas logistics tracking number | ["TN0003077L" ] |

<a id="m-com-alibaba-ocean-openplatform-biz-cross-model-overseastransportuseraddrinfo"></a>
#### com.alibaba.ocean.openplatform.biz.cross.model.OverseasTransportUserAddrInfo

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `postCode` | String | yes | Logistics code | 123 |
| `countryName` | String | yes | Country | 越南 |
| `provinceName` | String | yes | Province | xx |
| `cityName` | String | yes | City | x x |
| `areaName` | String | yes | District | xx |
| `addressDetail` | String | yes | Address | xx |
| `mobile` | String | yes | Mobile phone number | 123 |
| `fixedPhone` | String | yes | Mobile phone number | 123 |
| `warehouseContactName` | String | yes | Warehouse contact person | aa |
| `warehouseName` | String | yes | Warehouse name | aa |

<a id="m-com-alibaba-ocean-openplatform-biz-trade-common-model-opsellerinvoicetradesettingmodel"></a>
#### com.alibaba.ocean.openplatform.biz.trade.common.model.OpSellerInvoiceTradeSettingModel

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `tradeInvoiceStatus` | String | yes | Transaction invoicing status: -1: invoicing not allowed; 0: invoicing allowed; 1: applied; 2: invoiced | 1 |
| `sellerInvoiceType` | String | yes | Invoicing type: 0 - general VAT invoice; 1 - special VAT invoice; 2 - general/special VAT invoice | 0 |

## Samples

**Order detail output example**

```
{"result":{"baseInfo":{"allDeliveredTime":"20170913231916000-0700","businessType":"cn","buyerID":"b2b-2248544159","createTime":"20170913231708000-0700","id":58218860983545944,"modifyTime":"20170913234725000-0700","payTime":"20170913231727000-0700","refund":0,"sellerID":"b2b-2248564064","shippingFee":6,"status":"waitbuyerreceive","totalAmount":6.15,"discount":0,"buyerContact":{"phone":"86-0571-81895955","imInPlatform":"b测试账号002","name":"乔的石","mobile":"18668184036","companyName":"阿里巴巴网络科技有限公司"},"sellerContact":{"phone":"86-0571-88881888","email":"caigouwfw_1688@163.com","imInPlatform":"b测试账号110","name":"孟舒","mobile":"13312919596","companyName":"阿里巴巴网络有限公司"},"tradeType":"50060","refundStatus":"waitselleragree","refundPayment":0,"idOfStr":"58218860983545941","alipayTradeId":"2017091421001008480237437679","receiverInfo":{"toFullName":"童恩杰","toDivisionCode":"330108","toPost":"312000","toTownCode":"330108002","toArea":"浙江省 杭州市 滨江区 长河街道"},"buyerLoginId":"b测试账号002","sellerLoginId":"b测试账号110","buyerUserId":2248544159,"sellerUserId":2248564064,"buyerAlipayId":"2088611489970483","sellerAlipayId":"2088611383470360","sumProductPayment":0.3,"stepPayAll":false},"nativeLogistics":{"address":"杭州市滨江区网商路699号","area":"滨江区","areaCode":"330108","city":"杭州市","contactPerson":"童恩杰","mobile":"13666836263","province":"浙江省","zip":"312000","logisticsItems":[{"deliveredTime":"20170913231917000-0700","logisticsCode":"BX107450035961376","type":"2","id":107450035961376,"status":"alreadysend","gmtModified":"20170913231916000-0700","gmtCreate":"20170913231916000-0700","fromPhone":"86-0571-88881888","fromMobile":"13312919596","logisticsBillNo":"不需要物流","subItemIds":"58218860984545941,58218860985545941,58218860986545941"}],"townCode":"330108002","town":"长河街道"},"productItems":[{"itemAmount":0.3,"name":"测试扫码购富光勿拍2L*6件","price":0.3,"productID":547486647009,"productImgUrl":["http://cbu01.alicdn.com/img/ibank/2017/757/125/3950521757.80x80.jpg","http://cbu01.alicdn.com/img/ibank/2017/757/125/3950521757.jpg"],"productSnapshotUrl":"https://trade.1688.com/order/offer_snapshot.htm?order_entry_id=128403042259997715","quantity":1,"refund":0,"skuID":3315536521048,"status":"waitsellersend","subItemID":128403042259997710,"type":"common","unit":"箱","guaranteesTerms":[],"skuInfos":[{"name":"颜色","value":"白色"},{"name":"容量","value":"1.1L-2L"}],"entryDiscount":0,"specId":"28b952ec96c8b5c3ab0affc1b74923f0","quantityFactor":1,"statusStr":"等待卖家发货","refundStatus":"WAIT_SELLER_AGREE","logisticsStatus":1,"gmtCreate":"20180206152758000+0800","gmtModified":"20180206153325000+0800","gmtPayExpireTime":"2018-02-07 15:27:58","refundId":"TQ7257793022991577","subItemIDString":"128403042259997715"},{"itemAmount":0.05,"name":"这个一个很好看好看的鞋子用于服务测试（大家不要动）","price":0.1,"productID":558700975520,"productImgUrl":[],"productSnapshotUrl":"https://trade.1688.com/order/offer_snapshot.htm?order_entry_id=58218860985545941","quantity":1,"refund":0,"skuID":3638916762844,"status":"waitbuyerreceive","subItemID":58218860985545944,"type":"common","unit":"双","guaranteesTerms":[],"skuInfos":[{"name":"颜色","value":"白色"},{"name":"尺码","value":"38"}],"entryDiscount":0,"specId":"6ff6071792c8ab520b9b867c61b990bd","quantityFactor":1,"statusStr":"等待买家收货","refundStatus":"WAIT_SELLER_AGREE","logisticsStatus":2},{"itemAmount":0.05,"name":"这个一个很好看好看的鞋子用于服务测试（大家不要动）","price":0.1,"productID":558700975520,"productImgUrl":[],"productSnapshotUrl":"https://trade.1688.com/order/offer_snapshot.htm?order_entry_id=58218860986545941","quantity":1,"refund":0,"skuID":3638916762843,"status":"waitbuyerreceive","subItemID":58218860986545944,"type":"common","unit":"双","guaranteesTerms":[],"skuInfos":[{"name":"颜色","value":"白色"},{"name":"尺码","value":"34"}],"entryDiscount":0,"specId":"8ecd5cb401df85d3f2485eb120f670df","quantityFactor":1,"statusStr":"等待买家收货","refundStatus":"WAIT_SELLER_AGREE","logisticsStatus":2}],"tradeTerms":[{"payStatus":"2","payTime":"20170913231727000-0700","payWay":"1","phasAmount":6.15,"phase":655062941545941}],"extAttributes":[],"orderRateInfo":{"buyerRateStatus":5,"sellerRateStatus":5}}}
```
