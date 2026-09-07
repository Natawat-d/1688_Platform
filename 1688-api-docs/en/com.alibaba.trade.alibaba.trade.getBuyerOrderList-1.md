# View order list (buyer view)

Original name: 订单列表查看(买家视角)  
API: `com.alibaba.trade:alibaba.trade.getBuyerOrderList:1` · Category: Orders  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.trade.getBuyerOrderList-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.trade.getBuyerOrderList/{appKey}`  
Requires user authorization (access_token) · Requires signature

Get the buyer's order list; the user's memberId must equal the buyer memberId on the orders. This interface returns only basic order information, not logistics or invoice information. For logistics information call the order-detail interface; for invoice information call the invoice-information API.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `bizTypes` | String[] | no | Business type, supports: &quot;cn&quot; (regular order type), &quot;ws&quot; (large-value wholesale order type), &quot;yp&quot; (regular sample order type), &quot;yf&quot; (one-cent sample order type), &quot;fs&quot; (flash sale (limited-time discount) order type), &quot;cz&quot; (processing/customization order type), &quot;ag&quot; (agreement procurement order type), &quot;hp&quot; (group-buy order type), &quot;gc&quot; (national procurement order type), &quot;supply&quot; (supply-and-marketing order type), &quot;nyg&quot; (nyg order type), &quot;factory&quot; (Taobao Factory order type), &quot;quick&quot; (quick order placement), &quot;xiangpin&quot; (Xiangpin order), &quot;nest&quot; (Procurement Mall - Nest), &quot;f2f&quot; (face-to-face payment), &quot;cyfw&quot; (sample storage service), &quot;sp&quot; (consignment order flag), &quot;wg&quot; (WeiGong order), &quot;factorysamp&quot; (Taobao Factory sampling order), &quot;factorybig&quot; (Taobao Factory bulk order) | ["cn","ws"] |
| `createEndTime` | java.util.Date | no | Order placement end time | 20180802211113000+0800 |
| `createStartTime` | java.util.Date | no | Order start time | 20180102211113000+0800 |
| `isHis` | boolean | no | Whether to query the historical order table; default queries the current table, i.e. the default value is false | false |
| `modifyEndTime` | java.util.Date | no | End of the modification time query range | 20180802211113000+0800 |
| `modifyStartTime` | java.util.Date | no | Query modification time start | 20180102211113000+0800 |
| `orderStatus` | java.lang.String | no | Order status, values are success, cancel (transaction cancelled, penalty and other settlements completed), waitbuyerpay (waiting for seller to pay), waitsellersend (waiting for seller to ship), waitbuyerreceive (waiting for buyer to receive goods) | success |
| `page` | int | no | Query page number, starting from 1 | 1 |
| `pageSize` | int | no | Number of items per page for the query | 20 |
| `refundStatus` | java.lang.String | no | Refund status, supported values: &quot;waitselleragree&quot; (waiting for seller to agree), &quot;refundsuccess&quot; (refund successful), &quot;refundclose&quot; (refund closed), &quot;waitbuyermodify&quot; (waiting for buyer to modify), &quot;waitbuyersend&quot; (waiting for buyer to return the goods), &quot;waitsellerreceive&quot; (waiting for seller to confirm receipt) | refundsuccess |
| `sellerMemberId` | java.lang.String | no | Seller memberId | b2b-1624961198 |
| `sellerLoginId` | String | no | Seller loginId | alitestforisv02 |
| `sellerRateStatus` | java.lang.Integer | no | Seller rating status (4: rated, 5: not rated, 6: rating not required) | 6 |
| `tradeType` | java.lang.String | no | Transaction type:<br>Escrow transaction (1),<br>Prepaid deposit transaction (2),<br>ETC overseas acquiring transaction (3),<br>Instant payment transaction (4),<br>Guarantee fund security transaction (5),<br>Unified transaction process (6),<br>Staged transaction (7),<br>Cash on delivery transaction (8),<br>Credit voucher payment transaction (9),<br>Account period payment transaction (10),<br>1688 Transaction 4.0, new staged transaction (50060),<br>Face-to-face payment transaction process (50070),<br>Service-type transaction process (50080) | 50060 |
| `productName` | java.lang.String | no | Product name | 测试商品 |
| `needBuyerAddressAndPhone` | java.lang.Boolean | no | Whether the buyer's detailed address information and phone number need to be queried | false |
| `needMemoInfo` | java.lang.Boolean | no | Whether remark information needs to be queried | false |
| `outOrderId` | String | no | External order number, can be used for idempotency control | 90187872898371 |
| `needInvoicingSetting` | Boolean | no | Invoicing settings need to be queried | true |
| `orderIds` | Long[] | no | Query multiple order IDs | [90187872898371,90187872898372] |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.openplatform.trade.model.TradeInfo[]](#m-alibaba-openplatform-trade-model-tradeinfo[]) | yes | Query return list | [] |
| `errorCode` | String | yes | Error code |   |
| `errorMessage` | String | yes | Error message |   |
| `totalRecord` | Long | yes | Total record count | 528 |

<a id="m-alibaba-openplatform-trade-model-tradeinfo[]"></a>
#### alibaba.openplatform.trade.model.TradeInfo[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `quoteList` | [message:alibaba.orderDetail.caigouQuoteInfo[]](#m-alibaba-orderdetail-caigouquoteinfo[]) | yes | Purchase order detail list; a field exclusive to large-enterprise procurement orders. | {} |
| `extAttributes` | [message:alibaba.openplatform.trade.KeyValuePair[]](#m-alibaba-openplatform-trade-keyvaluepair[]) | yes | Order extended attributes | [] |
| `orderRateInfo` | [message:alibaba.trade.OrderRateInfo](#m-alibaba-trade-orderrateinfo) | yes | Order review information | {} |
| `orderInvoiceInfo` | [message:alibaba.invoice.OrderInvoiceModel](#m-alibaba-invoice-orderinvoicemodel) | yes | Invoice information | {} |
| `tradeTerms` | [message:alibaba.openplatform.trade.model.TradeTermsInfo[]](#m-alibaba-openplatform-trade-model-tradetermsinfo[]) | yes | Trade terms | {} |
| `nativeLogistics` | [message:alibaba.openplatform.trade.model.NativeLogisticsInfo](#m-alibaba-openplatform-trade-model-nativelogisticsinfo) | yes | Domestic logistics | {} |
| `guaranteesTerms` | [message:alibaba.openplatform.trade.model.GuaranteeTermsInfo](#m-alibaba-openplatform-trade-model-guaranteetermsinfo) | yes | Protection terms | {} |
| `baseInfo` | [message:alibaba.openplatform.trade.model.OrderBaseInfo](#m-alibaba-openplatform-trade-model-orderbaseinfo) | yes | Basic order information | {} |
| `orderBizInfo` | [message:alibaba.order.bizInfo](#m-alibaba-order-bizinfo) | yes | Order business info | {} |
| `productItems` | [message:alibaba.openplatform.trade.model.ProductItemInfo[]](#m-alibaba-openplatform-trade-model-productiteminfo[]) | yes | Product entry information | {} |
| `overseasExtraAddress` | [message:alibaba.trade.OverseasExtraAddress](#m-alibaba-trade-overseasextraaddress) | yes | Cross-border address extended information | {} |
| `customs` | [message:alibaba.trade.Customs](#m-alibaba-trade-customs) | yes | Cross-border customs declaration information | {} |
| `overseaLogisticsInfo` | [message:com.alibaba.ocean.openplatform.biz.trade.common.model.OverseaLogisticsInfo](#m-com-alibaba-ocean-openplatform-biz-trade-common-model-oversealogisticsinfo) | yes | Overseas receiving information | { 					"overseasTransportUserAddrInfo":{ 						"addressDetail":"test", 						"cityName":"Г.Курчатов", 						"areaName":" ", 						"warehouseContactName":"Zeus", 						"mobile":"123456789", 						"fixedPhone":"123456789", 						"postCode":"1111", 						"provinceName":"Абайская область", 						"countryName":"Казахстан", 						"@type":"com.alibaba.ocean.openplatform.biz.cross.model.OverseasTransportUserAddrInfo", 						"warehouseName":" " 					}, 					"overseasLogisticsIds":[ 						"TN0003077L" 					],} |
| `invoicingSettingModel` | [message:com.alibaba.ocean.openplatform.biz.trade.common.model.OpSellerInvoiceTradeSettingModel](#m-com-alibaba-ocean-openplatform-biz-trade-common-model-opsellerinvoicetradesettingmodel) | yes | Invoicing settings | {} |

<a id="m-alibaba-orderdetail-caigouquoteinfo[]"></a>
#### alibaba.orderDetail.caigouQuoteInfo[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `productQuoteName` | String | yes | Name of the supply line item | 物料01 |
| `price` | Long | yes | Price, unit: cents (fen) | 100 |
| `count` | Double | yes | Purchase quantity | 10 |

<a id="m-alibaba-openplatform-trade-keyvaluepair[]"></a>
#### alibaba.openplatform.trade.KeyValuePair[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `key` | String | yes | Key |   |
| `value` | String | yes | Value |   |
| `description` | String | yes | Description |   |

<a id="m-alibaba-trade-orderrateinfo"></a>
#### alibaba.trade.OrderRateInfo

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `buyerRateStatus` | Integer | yes | Buyer review status (4: reviewed, 5: not reviewed, 6: review not required) | 5 |
| `sellerRateStatus` | Integer | yes | Seller review status (4: reviewed, 5: not reviewed, 6: review not required) | 5 |

<a id="m-alibaba-invoice-orderinvoicemodel"></a>
#### alibaba.invoice.OrderInvoiceModel

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `invoiceCompanyName` | java.lang.String | yes | Invoice company name (i.e. the invoice title) |  |
| `invoiceType` | java.lang.Integer | yes | Invoice type. 0: general invoice, 1: VAT invoice, 9: unknown type |  |
| `localInvoiceId` | java.lang.Long | yes | Local invoice number |  |
| `orderId` | java.lang.Long | yes | Order ID |  |
| `receiveCode` | java.lang.String | yes | (Recipient) address region code |  |
| `receiveCodeText` | java.lang.String | yes | (Recipient) text corresponding to the province/city/district code (VAT invoice information) |  |
| `receiveMobile` | java.lang.String | yes | (Recipient) invoice recipient's mobile phone number |  |
| `receiveName` | java.lang.String | yes | (Recipient) invoice recipient |  |
| `receivePhone` | java.lang.String | yes | (Recipient) invoice recipient's phone number |  |
| `receivePost` | java.lang.String | yes | (Recipient) invoice shipping address postal code |  |
| `receiveStreet` | java.lang.String | yes | (Recipient) street address (VAT invoice information) |  |
| `registerAccountId` | java.lang.String | yes | (Company) bank account number |  |
| `registerBank` | java.lang.String | yes | (Company) bank of deposit |  |
| `registerCode` | java.lang.String | yes | (Registration) province/city/district code |  |
| `registerCodeText` | java.lang.String | yes | (Registration) province/city/district text |  |
| `registerPhone` | java.lang.String | yes | (Company) registered phone number |  |
| `registerStreet` | java.lang.String | yes | (Registered) street address |  |
| `taxpayerIdentify` | java.lang.String | yes | Taxpayer identification number |  |
| `amount` | Long | yes | Invoice total amount including tax (cents/fen) | 1 |
| `status` | String | yes | Invoice status: {&quot;ISSUED&quot;:&quot;Issued&quot;,&quot;RED_ISSUING&quot;:&quot;Red-flush in progress&quot;,&quot;RED_ALL_ISSUED&quot;:&quot;Fully red-flushed&quot;,&quot;CLOSED&quot;:&quot;Closed&quot;,&quot;RED_PART_ISSUED&quot;:&quot;Partially red-flushed&quot;,&quot;VERIFYING&quot;:&quot;Verifying&quot;,&quot;DEPRECATED&quot;:&quot;Voided&quot;,&quot;RETURNING&quot;:&quot;Return in progress&quot;,&quot;INVALIDING&quot;:&quot;Voiding&quot;,&quot;INIT&quot;:&quot;Initialized&quot;,&quot;VERIFY_FAILED&quot;:&quot;Verification failed&quot;,&quot;ISSUING&quot;:&quot;Issuing&quot;,&quot;FAILED&quot;:&quot;Issuance failed&quot;,&quot;RETURNED&quot;:&quot;Returned&quot;} | issued |

<a id="m-alibaba-openplatform-trade-model-tradetermsinfo[]"></a>
#### alibaba.openplatform.trade.model.TradeTermsInfo[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `payStatus` | java.lang.String | yes | Payment status. International site: WAIT_PAY (unpaid), PAYER_PAID (payment completed), PART_SUCCESS (partial payment succeeded), PAY_SUCCESS (payment succeeded), CLOSED (closed due to risk control), CANCELLED (payment cancelled), SUCCESS (success), FAIL (failed).<br>1688: unpaid, 1; paid, 2; full refund, 4; seller has received payment, payment collection completed, 6; external payment order not created, 7; cancelled before payment, 8; payment in progress, 9 | 6 |
| `payTime` | java.util.Date | yes | Time when the stage payment was completed | 20180807153006000+0800 |
| `payWay` | java.lang.String | yes | Payment method. International site: ECL (financing payment), CC (credit card), TT (offline TT), ACH (echecking payment).<br>1688: Alipay 1; MYbank Gold Bill 2; credit purchase (Cheng-e-She) 3; MYbank large-amount transfer 4; electronic acceptance bill 6; account period payment 7; virtual payment 8; no payment action 9; Lingshoutong credit purchase payment 10 | 1 |
| `phasAmount` | java.math.BigDecimal | yes | Stage amount | 1158 |
| `phase` | java.lang.Long | yes | Stage | 2988886630908522 |
| `phaseCondition` | java.lang.String | yes | Stage condition; not applicable on 1688. |   |
| `phaseDate` | java.lang.String | yes | Stage time; not applicable to 1688 |   |
| `payWayDesc` | String | yes | Payment method | 支付宝 |

<a id="m-alibaba-openplatform-trade-model-nativelogisticsinfo"></a>
#### alibaba.openplatform.trade.model.NativeLogisticsInfo

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `address` | java.lang.String | yes | Detailed address | 浙江省杭州市滨江区 |
| `area` | java.lang.String | yes | County, district | 滨江区 |
| `areaCode` | java.lang.String | yes | Province/city/district code | 330108 |
| `city` | java.lang.String | yes | City | 杭州市 |
| `contactPerson` | java.lang.String | yes | Contact person name | 张三 |
| `fax` | java.lang.String | yes | Fax |   |
| `mobile` | java.lang.String | yes | Mobile phone | 15000001111 |
| `province` | java.lang.String | yes | Province | 浙江省 |
| `telephone` | java.lang.String | yes | Phone number | 7314065 |
| `zip` | java.lang.String | yes | Postal code | 888888 |
| `townCode` | java.lang.String | yes | Town/street address code | 4403011 |
| `town` | java.lang.String | yes | Town, street | 测试街道 |

<a id="m-alibaba-openplatform-trade-model-guaranteetermsinfo"></a>
#### alibaba.openplatform.trade.model.GuaranteeTermsInfo

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `assuranceInfo` | java.lang.String | yes | Protection terms |  自愿选择向买家提供“交期保障”服务 |
| `assuranceType` | java.lang.String | yes | Guarantee method. International site: TA (Trade Assurance) |  jqbz |
| `qualityAssuranceType` | java.lang.String | yes | Quality assurance type. International site: pre_shipment (before shipment), post_delivery (after delivery) |  交期保障 |
| `value` | String | yes | Guarantee term value; for example, in a delivery time guarantee, 6 represents 6 days | 6 |

<a id="m-alibaba-openplatform-trade-model-orderbaseinfo"></a>
#### alibaba.openplatform.trade.model.OrderBaseInfo

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `allDeliveredTime` | java.util.Date | yes | Complete shipment time |   |
| `businessType` | java.lang.String | yes | Business type. International site: ta (Trade Assurance), wholesale (online wholesale).<br>Chinese site: regular order type = &quot;cn&quot;;<br>large wholesale order type = &quot;ws&quot;;<br>regular sample order type = &quot;yp&quot;;<br>one-cent sample order type = &quot;yf&quot;;<br>reverse-batch (limited-time discount) order type = &quot;fs&quot;;<br>processing/customization order type = &quot;cz&quot;;<br>agreement procurement order type = &quot;ag&quot;;<br>group-buy order type = &quot;hp&quot;;<br>supply-distribution order type = &quot;supply&quot;;<br>Taogongchang (Taobao Factory) order = &quot;factory&quot;;<br>quick order placement = &quot;quick&quot;;<br>Xiangpin (group buy) order = &quot;xiangpin&quot;;<br>face-to-face payment = &quot;f2f&quot;;<br>sample storage service = &quot;cyfw&quot;;<br>consignment order = &quot;sp&quot;;<br>Weigong (micro-supply) order = &quot;wg&quot;; Lingshoutong (Retail Link) = &quot;lst&quot;; | cb |
| `buyerID` | java.lang.String | yes | Buyer's main account id | b2b-397390228 |
| `buyerMemo` | java.lang.String | yes | Buyer memo info |   |
| `buyerSubID` | java.lang.Long | yes | Buyer sub-account ID; not applicable to 1688 |   |
| `completeTime` | java.util.Date | yes | Completion time |   |
| `createTime` | java.util.Date | yes | Creation time | 20180807152315000+0800 |
| `currency` | java.lang.String | yes | Currency. The entire trade order uses the same currency. Value range: USD, RMB, HKD, GBP, CAD, AUD, JPY, KRW, EUR |   |
| `id` | java.lang.Long | yes | Transaction ID | 1997017368139085 |
| `modifyTime` | java.util.Date | yes | Modification time | 20180807152533000+0800 |
| `payTime` | java.util.Date | yes | Payment time; if there were multiple payments, this returns the time of the first payment | 20180807153006000+0800 |
| `receivingTime` | java.util.Date | yes | Receipt time; the time returned here is the complete receipt time |   |
| `refund` | java.math.BigDecimal | yes | Refund amount, in yuan | 0 |
| `remark` | java.lang.String | yes | Remark; on 1688 this refers to the remark entered when placing the order | 测试备注 |
| `sellerID` | java.lang.String | yes | Seller's main account id | b2b-23187037 |
| `sellerSubID` | java.lang.Long | yes | Seller sub-account ID; not applicable on 1688. |   |
| `shippingFee` | java.math.BigDecimal | yes | Shipping fee, in yuan | 10 |
| `status` | java.lang.String | yes | Trade status. waitbuyerpay: waiting for buyer payment; waitsellersend: waiting for seller to ship; waitlogisticstakein: waiting for logistics company pickup; waitbuyerreceive: waiting for buyer to receive; waitbuyersign: waiting for buyer to sign for; signinsuccess: buyer has signed for; confirm_goods: goods received; success: transaction successful; cancel: transaction canceled; terminated: transaction terminated; unenumerated: other status | waitbuyerpay |
| `totalAmount` | java.math.BigDecimal | yes | Total payable amount, totalAmount = ∑itemAmount + shippingFee, in yuan | 528 |
| `buyerRemarkIcon` | String | yes | Buyer memo flag |   |
| `discount` | Long | yes | Discount information, in cents (fen) | -226 |
| `buyerContact` | [message:alibaba.trade.tradeContact](#m-alibaba-trade-tradecontact) | yes | Buyer contact person | {} |
| `sellerContact` | [message:alibaba.trade.tradeSellerContact](#m-alibaba-trade-tradesellercontact) | yes | Seller contact | {} |
| `tradeType` | String | yes | 1: Escrow (guaranteed) transaction<br>2: Pre-deposit transaction<br>3: ETC overseas acquiring transaction<br>4: Instant payment transaction<br>5: Security deposit protected transaction<br>6: Unified transaction process<br>7: Staged payment<br>8. Cash on delivery transaction<br>9. Credit voucher payment transaction<br>10. Account period payment transaction | 50060 |
| `refundStatus` | String | yes | In-sale refund status of the order |   |
| `refundStatusForAs` | String | yes | Order's after-sales refund status |   |
| `refundPayment` | Long | yes | Refund amount | 0 |
| `idOfStr` | String | yes | Transaction id (string format) | 1997017368139085 |
| `alipayTradeId` | java.lang.String | yes | External payment transaction ID | UNCREAT |
| `receiverInfo` | [message:alibaba.trade.orderReceiverInfo](#m-alibaba-trade-orderreceiverinfo) | yes | Recipient info | {} |
| `buyerLoginId` | java.lang.String | yes | Buyer's loginId, WangWang ID | alitestforisv01 |
| `sellerLoginId` | java.lang.String | yes | Seller oginId, WangWang ID | alitestforisv02 |
| `buyerUserId` | java.lang.Long | yes | Buyer numeric ID | 39739022 |
| `sellerUserId` | java.lang.Long | yes | Seller numeric ID | 23187037 |
| `buyerAlipayId` | java.lang.String | yes | Buyer's Alipay id | 20881314957453 |
| `sellerAlipayId` | java.lang.String | yes | Seller's Alipay ID | 20887127689595 |
| `confirmedTime` | java.util.Date | yes | Confirmation time |   |
| `closeReason` | java.lang.String | yes | Closure reason |   |
| `sumProductPayment` | java.math.BigDecimal | yes | Total product amount (the sum of the product amounts in the order's product detail table), unit: yuan | 258 |
| `stepOrderList` | [message:alibaba.trade.StepOrderModel[]](#m-alibaba-trade-stepordermodel[]) | yes | [Trade 3.0] Staged transaction, staged order list |   |
| `stepAgreementPath` | java.lang.String | yes | Staged legal agreement address |   |
| `stepPayAll` | java.lang.Boolean | yes | Whether it is a one-time payment | false |
| `buyerFeedback` | java.lang.String | yes | Buyer message | 测试买家留言 |
| `relatedCode` | java.lang.String | yes | Order association code (e.g. which promotional activity it belongs to, etc.) |   |
| `overSeaOrder` | java.lang.Boolean | yes | Whether it is an overseas dropshipping order; yes: true | true |
| `subBuyerLoginId` | String | yes | Buyer sub-account |   |
| `sellerOrder` | java.lang.Boolean | yes | Whether it is a self-initiated order (invitation order) |   |
| `preOrderId` | java.lang.Long | yes | Pre-order ID |   |
| `refundId` | java.lang.String | yes | Refund order ID |   |
| `flowTemplateCode` | String | yes | 4.0 transaction process template code |   |
| `officialSolutionOrderId` | String | yes | Official logistics pickup order ID | 123 |
| `officialSolutionCost` | Long | yes | Official logistics pickup order fee (in cents/fen) | 770 |

<a id="m-alibaba-trade-tradecontact"></a>
#### alibaba.trade.tradeContact

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `phone` | String | yes | Contact phone | 0912-2525263 |
| `fax` | String | yes | Fax |   |
| `email` | String | yes | Email |   |
| `imInPlatform` | String | yes | The contact's IM account on the platform | alitestforisv01 |
| `name` | String | yes | Contact name | 张三 |
| `mobile` | String | yes | Contact person's mobile number | 15123543625 |
| `companyName` | java.lang.String | yes | Company name | 阿里巴巴 |

<a id="m-alibaba-trade-tradesellercontact"></a>
#### alibaba.trade.tradeSellerContact

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `phone` | String | yes | Contact phone |  |
| `fax` | String | yes | Fax |  |
| `email` | String | yes | Email |  |
| `imInPlatform` | String | yes | The contact's IM account on the platform |  |
| `name` | String | yes | Contact name |  |
| `mobile` | String | yes | Contact person's mobile number |  |
| `companyName` | java.lang.String | yes | Company name |  |
| `wgSenderName` | java.lang.String | yes | Sender name, set by the distributor in distribution scenarios such as WeiGong (micro-supply) | 张** |
| `wgSenderPhone` | java.lang.String | yes | Sender's phone number; set by the distributor in distribution scenarios such as Weigong | 13800000000 |
| `shopName` | String | yes | Shop name |  三生网媒邮箱公司 |

<a id="m-alibaba-trade-orderreceiverinfo"></a>
#### alibaba.trade.orderReceiverInfo

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `toFullName` | java.lang.String | yes | Recipient | 张三 |
| `toDivisionCode` | java.lang.String | yes | Recipient address region code | 440311 |
| `toMobile` | java.lang.String | yes | Recipient's mobile phone number | 15236457596 |
| `toPhone` | java.lang.String | yes | Recipient phone number | 0912-44548585 |
| `toPost` | java.lang.String | yes | Postal code | 000000 |
| `toTownCode` | java.lang.String | yes | Recipient's street or town region code, may be empty | 440306011 |
| `toArea` | java.lang.String | yes | Shipping address | 广东省 深圳市 龙华区 大浪街道 |

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
| `abnormalPriceChange` | Boolean | yes | Whether the price has been changed | true |

<a id="m-alibaba-creditorder-fordetail"></a>
#### alibaba.creditOrder.forDetail

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `payAmount` | java.lang.Long | yes | Order amount | 10 |
| `createTime` | java.lang.String | yes | Payment time | 2018-01-01 00:00:00 |
| `status` | java.lang.String | yes | Status | END |
| `gracePeriodEndTime` | java.lang.String | yes | Latest repayment time | 2018-01-01 00:00:00 |
| `statusStr` | java.lang.String | yes | Status description | 已完结 |
| `restRepayAmount` | java.lang.Long | yes | Amount due for repayment | 11 |

<a id="m-alibaba-order-preorder-forread"></a>
#### alibaba.order.preOrder.forRead

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `appkey` | String | yes | The appkey that created the pre-order | 12345 |
| `marketName` | String | yes | The market name passed in when creating the pre-order | dxc |
| `createPreOrderApp` | Boolean | yes | Whether the pre-order was created by the ERP used for the current query | false |

<a id="m-alibaba-lst-tradeinfo"></a>
#### alibaba.lst.tradeInfo

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `lstWarehouseType` | String | yes | Retail Tong (Lingshoutong) warehouse type. customer: virtual warehouse; cainiao: physical warehouse | cainiao |

<a id="m-alibaba-openplatform-trade-model-productiteminfo[]"></a>
#### alibaba.openplatform.trade.model.ProductItemInfo[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `cargoNumber` | java.lang.String | yes | Specifies the SKU (single item) item number; not applicable for the international site. This field does not always have a value — it is only recorded when the order is placed (if the seller has set an item number for the SKU). For other order types, the item number can only be obtained via the product API. Note: the item number obtained via the product API and the item number at the time of order placement may differ, because the seller may modify the product information after the order is completed, changing the item number. | E0003 |
| `description` | java.lang.String | yes | Description; not applicable to 1688 |   |
| `itemAmount` | java.math.BigDecimal | yes | Actual amount paid, in yuan | 279 |
| `name` | java.lang.String | yes | Product name | 测试商品 |
| `price` | java.math.BigDecimal | yes | Original unit price, in yuan | 3 |
| `productID` | java.lang.Long | yes | Product ID (empty for non-online products) | 129527213581 |
| `productImgUrl` | String[] | yes | Product image URL | http://cbu01.alicdn.com/img/order/trading/025/894/036055839091/8536232543_347415001.80x80.jpg |
| `productSnapshotUrl` | java.lang.String | yes | Product snapshot URL. When a trade order is created, a snapshot of the product at that time is automatically recorded for reference in the event of a subsequent dispute. | https://trade.1688.com/order/offer_snapshot.htm?order_entry_id=190938550630498520 |
| `quantity` | java.math.BigDecimal | yes | Quantity in units of unit, e.g. how many pieces, items, boxes, or tons | 20 |
| `refund` | java.math.BigDecimal | yes | Refund amount, in yuan | 0 |
| `skuID` | java.lang.Long | yes | skuID | 352367821371 |
| `sort` | java.lang.Integer | yes | Sort field. The product list is sorted by this field, starting from 0; not provided by 1688. |   |
| `status` | java.lang.String | yes | Sub-order status | waitbuyerpay |
| `subItemID` | java.lang.Long | yes | Product line item ID | 20015919587908522 |
| `type` | java.lang.String | yes | Type, used on the international site, for sellers to mark the type the product belongs to | common |
| `unit` | java.lang.String | yes | Selling unit	E.g.: piece, item, box, ton | 件 |
| `weight` | java.lang.String | yes | Weight	Weight calculated in weight units, e.g.: 100 |   |
| `weightUnit` | java.lang.String | yes | Weight unit	e.g. g, kg, t |   |
| `guaranteesTerms` | [message:alibaba.openplatform.trade.model.GuaranteeTermsInfo[]](#m-alibaba-openplatform-trade-model-guaranteetermsinfo[]) | yes | Guarantee terms; this field applies only to 1688 |   |
| `productCargoNumber` | java.lang.String | yes | Specifies the product item number. This field does not always have a value — it is only recorded when the order is placed. For other order types, the item number can only be obtained via the product API. Note: the item number obtained via the product API and the item number at the time of order placement may differ, because the seller may modify the product information after the order is completed, changing the item number. The difference between this field and cargoNUmber is: this field is the item number defined at the product level, while cargoNUmber is the item number defined at the SKU (single item) level. | C017 |
| `skuInfos` | [message:alibaba.trade.SkuItemDesc[]](#m-alibaba-trade-skuitemdesc[]) | yes | SKU attribute description | [] |
| `entryDiscount` | Long | yes | Amount of price increase or decrease in the order line item | 0 |
| `specId` | java.lang.String | yes | Order sales attribute ID | 2b3878b01d251c057668066c085d75 |
| `quantityFactor` | java.math.BigDecimal | yes | The quantity precision factor in units of unit; the value is a power of 10. For example: quantityFactor=1000, unit=ton, then the minimum precision of quantity is 0.001 ton. | 1 |
| `statusStr` | java.lang.String | yes | Sub-order status description | 等待买家付款 |
| `refundStatus` | java.lang.String | yes | WAIT_SELLER_AGREE Waiting for seller to agree<br>REFUND_SUCCESS Refund successful<br>REFUND_CLOSED Refund closed<br>WAIT_BUYER_MODIFY Pending buyer modification<br>WAIT_BUYER_SEND Waiting for buyer to return goods<br>WAIT_SELLER_RECEIVE Waiting for seller to confirm receipt |   |
| `closeReason` | java.lang.String | yes | Closure reason |   |
| `logisticsStatus` | java.lang.Integer | yes | 1 Not shipped<br>2 Shipped<br>3 Received<br>4 Returned<br>5 Partially shipped<br>8 Logistics order not yet created | 1 |
| `refundId` | String | yes | In-sale refund order number | TQ123123 |
| `refundIdForAs` | String | yes | After-sales refund order number | TQ123123 |
| `relatedCode` | String | yes | Sub-order association code | 1365465 |
| `sharePostage` | BigDecimal | yes | Shared shipping fee, in yuan |   |

<a id="m-alibaba-openplatform-trade-model-guaranteetermsinfo[]"></a>
#### alibaba.openplatform.trade.model.GuaranteeTermsInfo[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `assuranceInfo` | java.lang.String | yes | Protection terms |  自愿选择向买家提供“交期保障”服务 |
| `assuranceType` | java.lang.String | yes | Guarantee method. International site: TA (Trade Assurance) |  jqbz |
| `qualityAssuranceType` | java.lang.String | yes | Quality assurance type. International site: pre_shipment (before shipment), post_delivery (after delivery) |  交期保障 |
| `value` | String | yes | Guarantee term value; for example, in a delivery time guarantee, 6 represents 6 days | 6 |

<a id="m-alibaba-trade-skuitemdesc[]"></a>
#### alibaba.trade.SkuItemDesc[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `name` | String | yes | Attribute name | 颜色 |
| `value` | String | yes | Attribute value | 黑色 |

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
| `attributes` | [message:alibaba.trade.CustomsAttributesInfo[]](#m-alibaba-trade-customsattributesinfo[]) | yes | List of customs declaration information |   |

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

**Example of return parameters**

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
