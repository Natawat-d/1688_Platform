# View order list (seller view)

Original name: 订单列表查看(卖家视角)  
API: `com.alibaba.trade:alibaba.trade.getSellerOrderList:1` · Category: Messaging  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.trade.getSellerOrderList-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.trade.getSellerOrderList/{appKey}`  
Requires user authorization (access_token) · Requires signature

Get the seller's order list; the user's memberId must equal the sellerMemberId on the orders. This interface returns only basic order information, not logistics or invoice information. For logistics information call the order-detail interface; for invoice information call the invoice-information API.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `createStartTime` | java.util.Date | no | Order start time | 20180721172608000+0800 |
| `createEndTime` | java.util.Date | no | Order placement end time | 20180721172608000+0800 |
| `modifyStartTime` | java.util.Date | no | Query modification time start | 20180721172608000+0800 |
| `modifyEndTime` | java.util.Date | no | End of the modification time query range | 20180721172608000+0800 |
| `page` | int | no | Query page number, starting from 1 | 1 |
| `pageSize` | int | no | Number of records per page for the query (maximum 20) | 10 |
| `orderStatus` | java.lang.String | no | Order status, possible values: success, cancel (transaction cancelled; liquidated damages, etc. have been settled), waitbuyerpay (waiting for buyer to pay), waitsellersend (waiting for seller to ship), waitbuyerreceive (waiting for buyer to receive the goods) | waitbuyerpay |
| `refundStatus` | java.lang.String | no | Refund status, supports:<br>&quot;waitselleragree&quot; (waiting for seller to agree),<br>&quot;refundsuccess&quot; (refund successful),<br>&quot;refundclose&quot; (refund closed),<br>&quot;waitbuyermodify&quot; (pending buyer modification),<br>&quot;waitbuyersend&quot; (waiting for buyer to return goods),<br>&quot;waitsellerreceive&quot; (waiting for seller to confirm receipt) | waitselleragree |
| `buyerMemberId` | java.lang.String | no | Buyer memberId or buyerOpenUid (buyer's encrypted ID) | b2b-1234325 |
| `buyerLoginId` | String | no | Buyer LoginId or buyerOpenUid (buyer's encrypted ID) | alitestforisv02 |
| `tradeType` | java.lang.String | no | Transaction type:<br>Escrow transaction (1),<br>Prepaid deposit transaction (2),<br>ETC overseas acquiring transaction (3),<br>Instant payment transaction (4),<br>Guarantee fund security transaction (5),<br>Unified transaction process (6),<br>Staged transaction (7),<br>Cash on delivery transaction (8),<br>Credit voucher payment transaction (9),<br>Account period payment transaction (10),<br>1688 Transaction 4.0, new staged transaction (50060),<br>Face-to-face payment transaction process (50070),<br>Service-type transaction process (50080) | 5 |
| `bizTypes` | String[] | no | Business type, supports: &quot;cn&quot; (regular order type), &quot;ws&quot; (large-value wholesale order type), &quot;yp&quot; (regular sample order type), &quot;yf&quot; (one-cent sample order type), &quot;fs&quot; (flash sale (limited-time discount) order type), &quot;cz&quot; (processing/customization order type), &quot;ag&quot; (agreement procurement order type), &quot;hp&quot; (group-buy order type), &quot;gc&quot; (national procurement order type), &quot;supply&quot; (supply-and-marketing order type), &quot;nyg&quot; (nyg order type), &quot;factory&quot; (Taobao Factory order type), &quot;quick&quot; (quick order placement), &quot;xiangpin&quot; (Xiangpin order), &quot;nest&quot; (Procurement Mall - Nest), &quot;f2f&quot; (face-to-face payment), &quot;cyfw&quot; (sample storage service), &quot;sp&quot; (consignment order flag), &quot;wg&quot; (WeiGong order), &quot;factorysamp&quot; (Taobao Factory sampling order), &quot;factorybig&quot; (Taobao Factory bulk order) | ["cn","ws"] |
| `isHis` | boolean | no | Whether to query the historical order table; default queries the current table | false |
| `productName` | java.lang.String | no | Product name | 测试商品名 |
| `needBuyerAddressAndPhone` | java.lang.Boolean | no | Whether the buyer's detailed address information and phone number need to be queried | false |
| `needMemoInfo` | java.lang.Boolean | no | Whether remark information needs to be queried | false |
| `tousuStatus` | boolean | no | Whether to search for the pending order-change request under complaint | false |
| `buyerRateStatus` | java.lang.Integer | no | Buyer review status (4: reviewed, 5: not reviewed, 6: review not required) | 5 |
| `sellerRateStatus` | java.lang.Integer | no | Seller rating status (4: rated, 5: not rated, 6: rating not required) | 5 |
| `needCheckSend` | Boolean | no | Whether shipment verification information needs to be returned; required and very important for label-printing and shipping scenarios | true |
| `needSendGoodsOverdueRisk` | Boolean | no | Whether the order shipment overdue risk flag is needed | false |
| `needDeliverGoodsOverdueRisk` | Boolean | no | Whether a pickup overdue risk flag is needed | false |
| `needOfficialLogisticOrder` | Boolean | no | Whether to filter orders with the official direct-delivery guarantee service | false |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.openplatform.trade.model.TradeInfo[]](#m-alibaba-openplatform-trade-model-tradeinfo[]) | yes | Query return result | {} |
| `errorCode` | String | yes | Error code |   |
| `errorMessage` | String | yes | Error message |   |
| `totalRecord` | Long | yes | Total record count | 101 |
| `success` | Boolean | yes | Whether the call was successful | true |
| `retCodes` | String[] | yes | Masked information code | ["xxx","xxx"] |

<a id="m-alibaba-openplatform-trade-model-tradeinfo[]"></a>
#### alibaba.openplatform.trade.model.TradeInfo[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `baseInfo` | [message:alibaba.openplatform.trade.model.OrderBaseInfo](#m-alibaba-openplatform-trade-model-orderbaseinfo) | yes | Basic order information | [] |
| `nativeLogistics` | [message:alibaba.openplatform.trade.model.NativeLogisticsInfo](#m-alibaba-openplatform-trade-model-nativelogisticsinfo) | yes | Domestic logistics | {} |
| `overseasExtraAddress` | [message:alibaba.trade.OverseasExtraAddress](#m-alibaba-trade-overseasextraaddress) | yes | Cross-border address extended information | {} |
| `productItems` | [message:alibaba.openplatform.trade.model.ProductItemInfo[]](#m-alibaba-openplatform-trade-model-productiteminfo[]) | yes | Product entry information | [] |
| `customs` | [message:alibaba.trade.Customs](#m-alibaba-trade-customs) | yes | Cross-border customs declaration information | {} |
| `tradeTerms` | [message:alibaba.openplatform.trade.model.TradeTermsInfo[]](#m-alibaba-openplatform-trade-model-tradetermsinfo[]) | yes | Trade terms | [] |
| `orderRateInfo` | [message:alibaba.trade.OrderRateInfo](#m-alibaba-trade-orderrateinfo) | yes | Order review information | {} |
| `orderInvoiceInfo` | [message:alibaba.invoice.OrderInvoiceModel](#m-alibaba-invoice-orderinvoicemodel) | yes | Invoice information | {} |
| `guaranteesTerms` | [message:alibaba.openplatform.trade.model.GuaranteeTermsInfo](#m-alibaba-openplatform-trade-model-guaranteetermsinfo) | yes | Protection terms | {} |
| `extAttributes` | [message:alibaba.openplatform.trade.KeyValuePair[]](#m-alibaba-openplatform-trade-keyvaluepair[]) | yes | Order extended attributes | [] |
| `fromEncryptOrder` | Boolean | yes | Whether the order was created with downstream desensitized information | true |
| `sendGoodsOverdueRisk` | String | yes | Order shipment overdue risk flag. PTMO: potentially overdue; TMOT: overdue | TMOT |
| `officialLogisticOrder` | Boolean | yes | When true, indicates an official direct-delivery guarantee service order; when empty or false, it is not an official direct-delivery guarantee service order. | null |

<a id="m-alibaba-openplatform-trade-model-orderbaseinfo"></a>
#### alibaba.openplatform.trade.model.OrderBaseInfo

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `allDeliveredTime` | java.util.Date | yes | Complete shipment time | 20180614101942000+0800 |
| `payTime` | java.util.Date | yes | Payment time; if there were multiple payments, this returns the time of the first payment | 20180614101942000+0800 |
| `buyerRemarkIcon` | String | yes | Buyer memo flag | 1 |
| `receiverInfo` | [message:alibaba.trade.orderReceiverInfo](#m-alibaba-trade-orderreceiverinfo) | yes | Recipient info | {} |
| `discount` | Long | yes | Discount information, in cents (fen) | 11 |
| `refundStatus` | String | yes | The order's in-sale refund status: waiting for seller to agree: waitselleragree, pending buyer modification: waitbuyermodify, waiting for buyer to return goods: waitbuyersend, waiting for seller to confirm receipt: waitsellerreceive, refund successful: refundsuccess, refund failed: refundclose | refundclose |
| `alipayTradeId` | java.lang.String | yes | External payment transaction ID | 123123121111 |
| `remark` | java.lang.String | yes | Remark; on 1688 this refers to the remark entered when placing the order | 备注 |
| `sumProductPayment` | java.math.BigDecimal | yes | Total product amount (the sum of the product amounts in the order's product detail table), unit: yuan | 1212 |
| `buyerFeedback` | java.lang.String | yes | Buyer's message, no more than 500 characters | 留言 |
| `flowTemplateCode` | String | yes | 4.0 transaction process template code | flow |
| `sellerOrder` | java.lang.Boolean | yes | Whether it is a self-initiated order (invitation order) | false |
| `preOrderId` | java.lang.Long | yes | Pre-order ID | 123123 |
| `buyerLoginId` | java.lang.String | yes | Buyer's loginId, WangWang ID | alitestforusv01 |
| `modifyTime` | java.util.Date | yes | Modification time | 20180614101942000+0800 |
| `subBuyerLoginId` | String | yes | Buyer sub-account | alitestforusv02:temp |
| `confirmedTime` | java.util.Date | yes | Confirmation time | 20180614101942000+0800 |
| `currency` | java.lang.String | yes | Currency. The entire trade order uses the same currency. Value range: USD, RMB, HKD, GBP, CAD, AUD, JPY, KRW, EUR | EUR |
| `id` | java.lang.Long | yes | Transaction ID | 1231231231111 |
| `closeReason` | java.lang.String | yes | Close reason. buyerCancel: buyer cancelled the order, sellerGoodsLack: seller out of stock, other: other | buyerCancel |
| `tradeType` | String | yes | 1: Escrow (guaranteed) transaction<br>2: Pre-deposit transaction<br>3: ETC overseas acquiring transaction<br>4: Instant payment transaction<br>5: Security deposit protected transaction<br>6: Unified transaction process<br>7: Staged payment<br>8. Cash on delivery transaction<br>9. Credit voucher payment transaction<br>10. Account period payment transaction, 50060 Trade 4.0 | 50060 |
| `ccid` | String | yes | Contact information decryption ID, used for e-commerce platform contact information encryption scenarios. Do not use for non-order encryption scenarios. | xxx |
| `buyerContact` | [message:alibaba.trade.tradeContact](#m-alibaba-trade-tradecontact) | yes | Buyer contact person | {} |
| `receivingTime` | java.util.Date | yes | Receipt time; the time returned here is the complete receipt time | 20180614101942000+0800 |
| `stepAgreementPath` | java.lang.String | yes | Staged legal agreement address |   |
| `idOfStr` | String | yes | Transaction id (string format) | 123121212123 |
| `refundStatusForAs` | String | yes | Order's after-sales refund status |   |
| `stepPayAll` | java.lang.Boolean | yes | Whether it is a one-time payment | false |
| `completeTime` | java.util.Date | yes | Completion time | 20180614101942000+0800 |
| `sellerLoginId` | java.lang.String | yes | Seller oginId, WangWang ID | alitestforusv02 |
| `buyerID` | java.lang.String | yes | Buyer's main account id | 1234531 |
| `stepOrderList` | [message:alibaba.trade.StepOrderModel[]](#m-alibaba-trade-stepordermodel[]) | yes | [Trade 3.0] Staged transaction, staged order list |   |
| `totalAmount` | java.math.BigDecimal | yes | Total payable amount, totalAmount = ∑itemAmount + shippingFee, in yuan | 1000 |
| `sellerID` | java.lang.String | yes | Seller's main account id | 123123123123 |
| `shippingFee` | java.math.BigDecimal | yes | Shipping fee, in yuan | 1 |
| `createTime` | java.util.Date | yes | Creation time | 20180614101942000+0800 |
| `sellerRemarkIcon` | String | yes | Seller's memo flag | 1 |
| `sellerMemo` | java.lang.String | yes | Seller's memo information | 备忘 |
| `businessType` | java.lang.String | yes | Business type. International site: ta (Trade Assurance), wholesale (online wholesale).<br>China site: regular order type = &quot;cn&quot;;<br>large-value wholesale order type = &quot;ws&quot;;<br>regular sample order type = &quot;yp&quot;;<br>one-cent sample order type = &quot;yf&quot;;<br>reverse wholesale (limited-time discount) order type = &quot;fs&quot;;<br>processing/customization order type = &quot;cz&quot;;<br>agreement procurement order type = &quot;ag&quot;;<br>group-buy (huopin) order type = &quot;hp&quot;;<br>supply-distribution order type = &quot;supply&quot;;<br>Taobao Factory order = &quot;factory&quot;;<br>quick order placement = &quot;quick&quot;;<br>Xiangpin order = &quot;xiangpin&quot;;<br>face-to-face payment = &quot;f2f&quot;;<br>sample storage service = &quot;cyfw&quot;;<br>consignment order = &quot;sp&quot;;<br>Weigong order = &quot;wg&quot;; Lingshoutong (Retail Link) = &quot;lst&quot;; cross-border = 'cb'; distribution = 'distribution'; Cai Yuan Bao = 'cab'; processing/customization = &quot;manufact&quot; | cn |
| `overSeaOrder` | java.lang.Boolean | yes | Whether it is an overseas dropshipping order; yes: true | true |
| `refundId` | java.lang.String | yes | Refund order ID | TQ4562212313 |
| `refund` | java.math.BigDecimal | yes | Refund amount, in yuan | 1 |
| `status` | java.lang.String | yes | Trade status. waitbuyerpay: waiting for buyer payment; waitsellersend: waiting for seller to ship; waitlogisticstakein: waiting for logistics company pickup; waitbuyerreceive: waiting for buyer to receive; waitbuyersign: waiting for buyer to sign for; signinsuccess: buyer has signed for; confirm_goods: goods received; success: transaction successful; cancel: transaction canceled; terminated: transaction terminated; unenumerated: other status | waitbuyerpay |
| `refundPayment` | Long | yes | Refund amount | 1 |
| `sellerContact` | [message:alibaba.trade.tradeSellerContact](#m-alibaba-trade-tradesellercontact) | yes | Seller contact information | {} |
| `relatedCode` | String | yes | Associated code | 229195140003841187 |
| `buyerOpenUid` | String | yes | Buyer's encrypted ID, which can be decrypted via the API; this ID is unique. Note: this ID is returned inconsistently across different appkeys. | xxx |
| `buyerSubOpenUid` | String | yes | Buyer's sub-account encrypted ID, which can be decrypted via an API; this ID is unique. Note: this ID is returned differently across different appkeys. | xxx |
| `buyerUserId` | Long | yes | Buyer's Userid | 22919514 |
| `sellerUserId` | Long | yes | Seller's UserId | 22919514 |
| `buyerAlipayId` | String | yes | Buyer's AlipayId | 441951400038411872 |
| `sellerAlipayId` | String | yes | Seller's AlipayId | 221400038411871123 |

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

<a id="m-alibaba-openplatform-trade-model-nativelogisticsinfo"></a>
#### alibaba.openplatform.trade.model.NativeLogisticsInfo

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `caid` | String | yes | Decrypted address ID, used for e-commerce platform recipient information encryption scenarios; do not use outside order encryption scenarios. |   |
| `address` | java.lang.String | yes | Detailed address |   |
| `area` | java.lang.String | yes | County, district |   |
| `areaCode` | java.lang.String | yes | Province/city/district code |   |
| `city` | java.lang.String | yes | City |   |
| `contactPerson` | java.lang.String | yes | Contact person name |   |
| `fax` | java.lang.String | yes | Fax |   |
| `mobile` | java.lang.String | yes | Mobile phone |   |
| `province` | java.lang.String | yes | Province |   |
| `telephone` | java.lang.String | yes | Phone number |   |
| `zip` | java.lang.String | yes | Postal code |   |
| `logisticsItems` | [message:alibaba.openplatform.trade.model.NativeLogisticsItemsInfo[]](#m-alibaba-openplatform-trade-model-nativelogisticsitemsinfo[]) | yes | Waybill details |   |
| `townCode` | java.lang.String | yes | Town/street address code |   |
| `town` | java.lang.String | yes | Town, street |   |

<a id="m-alibaba-openplatform-trade-model-nativelogisticsitemsinfo[]"></a>
#### alibaba.openplatform.trade.model.NativeLogisticsItemsInfo[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `deliveredTime` | java.util.Date | yes | Shipping time |  |
| `logisticsCode` | java.lang.String | yes | Logistics number |  |
| `type` | java.lang.String | yes | SELF_SEND_GOODS (&quot;0&quot;) self-delivery, online delivery ONLINE_SEND_GOODS (&quot;1&quot;), delivery without logistics NO_LOGISTICS_SEND_GOODS (&quot;2&quot;) |  |
| `id` | java.lang.Long | yes | Primary key id |  |
| `status` | java.lang.String | yes | Status |  |
| `gmtModified` | java.util.Date | yes | Modification time |  |
| `gmtCreate` | java.util.Date | yes | Creation time |  |
| `carriage` | java.math.BigDecimal | yes | Shipping fee (in yuan) |  |
| `fromProvince` | java.lang.String | yes | Shipping province |  |
| `fromCity` | java.lang.String | yes | Shipping city |  |
| `fromArea` | java.lang.String | yes | Shipping district |  |
| `fromAddress` | java.lang.String | yes | Shipping street address |  |
| `fromPhone` | java.lang.String | yes | Shipping contact phone number |  |
| `fromMobile` | java.lang.String | yes | Shipping contact mobile number |  |
| `fromPost` | java.lang.String | yes | Shipping address postal code |  |
| `logisticsCompanyId` | java.lang.Long | yes | Logistics company Id |  |
| `logisticsCompanyNo` | java.lang.String | yes | Logistics company number |  |
| `logisticsCompanyName` | java.lang.String | yes | Logistics company name |  |
| `logisticsBillNo` | java.lang.String | yes | Logistics company waybill number |  |
| `subItemIds` | java.lang.String | yes | Product line item ID; if multiple, separated by commas |  |
| `toProvince` | java.lang.String | yes | Receiving province |  |
| `toCity` | java.lang.String | yes | Receiving city |  |
| `toArea` | java.lang.String | yes | Receiving district |  |
| `toAddress` | java.lang.String | yes | Receiving street address |  |
| `toPhone` | java.lang.String | yes | Recipient contact phone number |  |
| `toMobile` | java.lang.String | yes | Receiving contact mobile number |  |
| `toPost` | java.lang.String | yes | Shipping address postal code |  |

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

<a id="m-alibaba-openplatform-trade-model-productiteminfo[]"></a>
#### alibaba.openplatform.trade.model.ProductItemInfo[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `cargoNumber` | java.lang.String | yes | Specifies the SKU (single item) item number; not applicable for the international site. This field does not always have a value — it is only recorded when the order is placed (if the seller has set an item number for the SKU). For other order types, the item number can only be obtained via the product API. Note: the item number obtained via the product API and the item number at the time of order placement may differ, because the seller may modify the product information after the order is completed, changing the item number. | 123 |
| `description` | java.lang.String | yes | Description; not applicable to 1688 | 描述 |
| `itemAmount` | java.math.BigDecimal | yes | Actual amount paid, in yuan | 12 |
| `name` | java.lang.String | yes | Product name | 商品名称 |
| `price` | java.math.BigDecimal | yes | Original unit price, in yuan | 12.5 |
| `productID` | java.lang.Long | yes | Product ID (empty for non-online products) | 12345666 |
| `productImgUrl` | String[] | yes | Product image URL | ["http://cbu01.alicdn.com/img/ibank/2019/700/221/12771122007.80x80.jpg"] |
| `productSnapshotUrl` | java.lang.String | yes | Product snapshot URL. When a trade order is created, a snapshot of the product at that time is automatically recorded for reference in the event of a subsequent dispute. | https://trade.1688.com/order/offer_snapshot.htm?order_entry_id=731312321827747370 |
| `quantity` | java.math.BigDecimal | yes | Quantity in units of unit, e.g. how many pieces, items, boxes, or tons | 12 |
| `refund` | java.math.BigDecimal | yes | Refund amount, in yuan | 12 |
| `skuID` | java.lang.Long | yes | skuID | 12 |
| `sort` | java.lang.Integer | yes | Sort field. The product list is sorted by this field, starting from 0; not provided by 1688. | 0 |
| `status` | java.lang.String | yes | Sub-order status. If the order status is pending shipment (waitsellersend), you also need to use canSendGoods to determine whether the current sub-order can be shipped: if false, the current sub-order cannot be shipped; if true, the current sub-order can be shipped normally. | waitsellersend |
| `subItemID` | java.lang.Long | yes | Sub-order number, or product detail line item ID | 731312321827747370 |
| `type` | java.lang.String | yes | Type, used on the international site, for sellers to mark the type the product belongs to | common |
| `unit` | java.lang.String | yes | Selling unit	E.g.: piece, item, box, ton | 个 |
| `weight` | java.lang.String | yes | Weight	Weight calculated in weight units, e.g.: 100 | 100 |
| `weightUnit` | java.lang.String | yes | Weight unit	e.g. g, kg, t | g |
| `guaranteesTerms` | [message:alibaba.openplatform.trade.model.GuaranteeTermsInfo[]](#m-alibaba-openplatform-trade-model-guaranteetermsinfo[]) | yes | Guarantee terms; this field applies only to 1688 | {} |
| `productCargoNumber` | java.lang.String | yes | Specifies the product item number. This field does not always have a value — it is only recorded when the order is placed. For other order types, the item number can only be obtained via the product API. Note: the item number obtained via the product API and the item number at the time of order placement may differ, because the seller may modify the product information after the order is completed, changing the item number. The difference between this field and cargoNUmber is: this field is the item number defined at the product level, while cargoNUmber is the item number defined at the SKU (single item) level. | 123 |
| `skuInfos` | [message:alibaba.trade.SkuItemDesc[]](#m-alibaba-trade-skuitemdesc[]) | yes | SKU information | {} |
| `entryDiscount` | Long | yes | Amount of price increase or decrease in the order line item | 0 |
| `specId` | java.lang.String | yes | Order sales attribute ID | 213123123123213ecfw12331 |
| `quantityFactor` | java.math.BigDecimal | yes | The quantity precision factor in units of unit; the value is a power of 10. For example: quantityFactor=1000, unit=ton, then the minimum precision of quantity is 0.001 ton. | 1000 |
| `statusStr` | java.lang.String | yes | Sub-order status description | 等待卖家同意 |
| `refundStatus` | java.lang.String | yes | WAIT_SELLER_AGREE Waiting for seller to agree<br>REFUND_SUCCESS Refund successful<br>REFUND_CLOSED Refund closed<br>WAIT_BUYER_MODIFY Pending buyer modification<br>WAIT_BUYER_SEND Waiting for buyer to return goods<br>WAIT_SELLER_RECEIVE Waiting for seller to confirm receipt | WAIT_SELLER_AGREE |
| `closeReason` | java.lang.String | yes | Closure reason | 测试 |
| `logisticsStatus` | java.lang.Integer | yes | 1 Not shipped<br>2 Shipped<br>3 Received<br>4 Returned<br>5 Partially shipped<br>8 Logistics order not yet created | 1 |
| `gmtCreate` | java.util.Date | yes | Creation time | 20190801154220368+0800 |
| `gmtModified` | java.util.Date | yes | Modification time | 20190801154220368+0800 |
| `gmtCompleted` | java.util.Date | yes | Line item completion time | 20190801154220368+0800 |
| `gmtPayExpireTime` | String | yes | Stock timeout time, in the format "yyyy-MM-dd HH:mm:ss" | 2020-03-01 08:00:00 |
| `refundId` | String | yes | Refund order number | TQ12345t5 |
| `subItemIDString` | String | yes | Sub-order number, or product detail line item ID (string type; since a Long-type ID may cause processing issues in JS and PHP, it can be handled as a string type) | 731312321827747370 |
| `canSendGoods` | Boolean | yes | Whether the current sub-order can be shipped. If the current sub-order status is waiting-to-ship (waitsellersend), as long as canSendGoods is false, the sub-order cannot be shipped regardless of business type; if true, it can be shipped normally. | true |
| `cantSendReason` | String | yes | The reason shipment is not possible; applies to both group-buying and C2M business. | 未成团，不可发货/该订单已由系统自动安排菜鸟仓发货，请勿重复发货 |
| `permitLogisticsCpCode` | String | yes | The logistics company cpCode allowed for declared shipment; multiple cpCodes are separated by half-width commas. |   |
| `deliverGoodsOverdueRisk` | String | yes | Upcoming pickup-overdue tag. overdue: pickup is about to be overdue | overdue |

<a id="m-alibaba-openplatform-trade-model-guaranteetermsinfo[]"></a>
#### alibaba.openplatform.trade.model.GuaranteeTermsInfo[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `assuranceInfo` | java.lang.String | yes | Protection terms |  |
| `assuranceType` | java.lang.String | yes | Guarantee method. International site: TA (Trade Assurance) |  |
| `qualityAssuranceType` | java.lang.String | yes | Quality assurance type. International site: pre_shipment (before shipment), post_delivery (after delivery) |  |

<a id="m-alibaba-trade-skuitemdesc[]"></a>
#### alibaba.trade.SkuItemDesc[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `name` | String | yes | Attribute name |  |
| `value` | String | yes | Attribute value |  |

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

<a id="m-alibaba-openplatform-trade-model-guaranteetermsinfo"></a>
#### alibaba.openplatform.trade.model.GuaranteeTermsInfo

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `assuranceInfo` | java.lang.String | yes | Protection terms |  |
| `assuranceType` | java.lang.String | yes | Guarantee method. International site: TA (Trade Assurance) |  |
| `qualityAssuranceType` | java.lang.String | yes | Quality assurance type. International site: pre_shipment (before shipment), post_delivery (after delivery) |  |

<a id="m-alibaba-openplatform-trade-keyvaluepair[]"></a>
#### alibaba.openplatform.trade.KeyValuePair[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `key` | String | yes | Key |  |
| `value` | String | yes | Value |  |
| `description` | String | yes | Description |  |

## Samples

**Very important: description of whether the order can be shipped**

```
只有当子订单状态为待发货状态（waitsellersend）时，且当前子订单下canSendGoods为true时，此子订单对应的商品才可以发货，如果子订单状态为待发货状态（waitsellersend），且当前子订单canSendGoods为false，则该子订单不可以发货，否则商家会造成资损，目前有拼团和C2M等业务都会有管种情况的订单，拼团的订单不能发货的原因（cantSendReason）是订单未成团，C2M业务不能发货的原因是该订单已经由系统自动安排签约的菜鸟仓发货。注意：对于发货场景，请求参数includeFields必须要传CanSendCheck才可以获取此信息。
```

**Output parameter example**

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
