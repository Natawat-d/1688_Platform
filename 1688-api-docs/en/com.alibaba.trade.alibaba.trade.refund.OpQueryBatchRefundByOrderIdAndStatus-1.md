# Query refund details by order ID (buyer view)

Original name: 查询退款单详情-根据订单ID（买家视角）  
API: `com.alibaba.trade:alibaba.trade.refund.OpQueryBatchRefundByOrderIdAndStatus:1` · Category: Returns & Refunds  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.trade.refund.OpQueryBatchRefundByOrderIdAndStatus-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.trade.refund.OpQueryBatchRefundByOrderIdAndStatus/{appKey}`  
Requires user authorization (access_token) · Requires signature

For buyers; sellers should use alibaba.trade.refund.OpQueryBatchRefundByOrderIdAndStatus.sellerView. Queries the refund list for an order in real time. Currently only in-sale refunds (before the transaction completes) can be queried.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `orderId` | String | yes | Order id | 151267031**8969811 |
| `queryType` | String | yes | 1: active; 3: refund successful (only refund-in-progress and refund-successful are supported) | 3 |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.trade.refund.OpQueryBatchRefundByOrderIdAndStatusResultModel](#m-alibaba-trade-refund-opquerybatchrefundbyorderidandstatusresultmodel) | yes | Query result | {} |
| `errorCode` | String | yes | Error code | 500 |
| `errorMessage` | String | yes | Error message | "{\"errorCode\":\"003002\",\"errorMessage\":\"SERVICE:INVOKE_FAIL:OrderRefundService.queryBatchRefundByOrderIdAndStatus\",\"cause\":\"errorCode:ORDER_NOT_EXIST,errorMsg:null,cause:null\"}" |
| `extErrorMessage` | String | yes | Additional information | 订单号有误 |

<a id="m-alibaba-trade-refund-opquerybatchrefundbyorderidandstatusresultmodel"></a>
#### alibaba.trade.refund.OpQueryBatchRefundByOrderIdAndStatusResultModel

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `opOrderRefundModels` | [message:alibaba.trade.refund.OpOrderRefundModel[]](#m-alibaba-trade-refund-oporderrefundmodel[]) | yes | Refund order information | {} |

<a id="m-alibaba-trade-refund-oporderrefundmodel[]"></a>
#### alibaba.trade.refund.OpOrderRefundModel[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `alipayPaymentId` | java.lang.String | yes | Alipay transaction number | 2018042321****08760555342267 |
| `applyCarriage` | java.lang.Long | yes | Shipping fee refund amount requested, in cents | 0 |
| `applyExpect` | java.lang.Long | yes | Refund amount originally entered by the buyer (can be empty) |   |
| `applyPayment` | java.lang.Long | yes | Refund amount requested by the buyer, in cents | 1 |
| `applyReason` | java.lang.String | yes | Application reason | 颜色/图案/款式不符 |
| `applyReasonId` | int | yes | Application reason ID | 20021 |
| `applySubReason` | java.lang.String | yes | Secondary refund reason |   |
| `asynErrCode` | java.lang.String | yes |  |   |
| `applySubReasonId` | int | yes | Secondary refund reason ID | -1 |
| `asynSubErrCode` | java.lang.String | yes |  |   |
| `buyerAlipayId` | java.lang.String | yes | Buyer's Alipay ID |   |
| `buyerLogisticsName` | java.lang.String | yes | Buyer's return logistics company name | 其他 |
| `buyerMemberId` | java.lang.String | yes | Buyer member ID | b2b-1624961198 |
| `buyerUserId` | java.lang.Long | yes | Buyer's Alibaba account ID (including Taobao account ID) | 1624961198 |
| `canRefundPayment` | java.lang.Long | yes | Maximum refundable amount, in cents | 1 |
| `disburseChannel` | java.lang.String | yes | Instant-arrival payment channel |   |
| `disputeRequest` | int | yes | After-sales refund requirement | 3 |
| `disputeType` | int | yes | Dispute type: in-sale refund, after-sales refund; defaults to in-sale refund | 1 |
| `extInfo` | java.util.Map | yes | Extended information | {} |
| `freightBill` | java.lang.String | yes | Waybill number | 12345134135435234 |
| `frozenFund` | long | yes | Actual frozen account amount, unit: cent | -1 |
| `gmtApply` | java.util.Date | yes | Refund application time | 20180423171116000+0800 |
| `gmtCompleted` | java.util.Date | yes | Completion time | 20180423174713000+0800 |
| `gmtCreate` | java.util.Date | yes | Creation time | 20180423171116000+0800 |
| `gmtFreezed` | java.util.Date | yes | The timeout freeze start time for this refund order |   |
| `gmtModified` | java.util.Date | yes | Modification time | 20180423174714000+0800 |
| `gmtTimeOut` | java.util.Date | yes | The time limit by which this refund order must be completed before timing out |   |
| `goodsStatus` | int | yes | 1: buyer has not received the goods<br>2: buyer has received the goods<br>3: buyer has returned the goods | 3 |
| `id` | java.lang.Long | yes | Refund order number | 8651493722961198 |
| `instantRefundType` | java.lang.String | yes | Instant refund type |   |
| `orderEntryCountMap` | java.util.Map | yes | Sub-order return quantity |   |
| `orderEntryIdList` | java.util.List | yes | Order details included in the refund order, listed in reverse chronological order |   |
| `orderId` | java.lang.Long | yes | Order number corresponding to the refund order | 151267031008969811 |
| `prepaidBalance` | java.lang.Long | yes | Express refund advance funding amount. When this value is not empty, it only means this refund order is eligible for the advance funding process, not that the advance funding will necessarily succeed |   |
| `productName` | java.lang.String | yes | Product name (the item name from the order detail associated with the refund order) | 欧洲站2018新款宽松显瘦烫金印花字母老鹰短袖T恤打底裤套装女潮 |
| `refundCarriage` | java.lang.Long | yes | Actual refund amount for shipping fee, in cents | 0 |
| `refundId` | java.lang.String | yes | Refund order logical primary key | TQ8651493722961198 |
| `refundPayment` | java.lang.Long | yes | Actual refund amount, in cents | 1 |
| `rejectReason` | java.lang.String | yes | Seller rejection reason |   |
| `rejectReasonId` | int | yes | Seller rejection reason Id | 0 |
| `rejectTimes` | int | yes | Number of times the refund order was rejected | 0 |
| `sellerAlipayId` | java.lang.String | yes | Seller's Alipay ID |   |
| `sellerMemberId` | java.lang.String | yes | Seller member ID | b2b-1623492085 |
| `sellerMobile` | java.lang.String | yes | Recipient's mobile phone |   |
| `sellerRealName` | java.lang.String | yes | Recipient name |   |
| `sellerReceiveAddress` | java.lang.String | yes | Seller's receiving address for buyer returns |   |
| `sellerTel` | java.lang.String | yes | Consignee phone number |   |
| `sellerUserId` | java.lang.Long | yes | Seller's Alibaba account ID (including Taobao account ID) | 1623492085 |
| `status` | java.lang.String | yes | Refund status | refundsuccess |
| `taskStatus` | java.lang.String | yes | Ticket sub-status; empty when the ticket has not yet flowed to CRM for creation |   |
| `timeOutOperateType` | java.lang.String | yes | Action executed after timeout |   |
| `tradeTypeStr` | java.lang.String | yes | Trade type, used to replace the enum type tradeType | 50060 |
| `success` | Boolean | yes | Whether successful | true |
| `refundOperationList` | [message:alibaba.trade.refund.OpOrderRefundOperationModel[]](#m-alibaba-trade-refund-oporderrefundoperationmodel[]) | yes | Operation record list | 暂不返回 |
| `buyerLoginId` | String | yes | Buyer member ID | alitestforisv02 |
| `sellerLoginId` | String | yes | Seller member ID | alitestforisv01 |
| `isCrmModifyRefund` | Boolean | yes | Whether the refund order has been modified by customer service (Xiaoer) | false |
| `isTimeOutFreeze` | Boolean | yes | Whether frozen by the system due to timeout. true means frozen, false means not frozen. | false |
| `isInsufficientAccount` | Boolean | yes | Transaction 4.0 refund balance insufficient | false |
| `isGoodsReceived` | Boolean | yes | Whether the buyer has received the goods | true |
| `isOnlyRefund` | Boolean | yes | Whether it is refund only | false |
| `isRefundGoods` | Boolean | yes | Whether return is required | true |
| `isSellerDelayDisburse` | Boolean | yes | Whether the seller delays payment, i.e., safe refund | false |
| `isAftersaleAutoDisburse` | Boolean | yes | After-sales automatic payment | false |
| `isSupportNewSteppay` | Boolean | yes | Whether Transaction 4.0 is supported | true |
| `isNewRefundReturn` | Boolean | yes | Whether the refund/return was created via the new process | true |
| `isBuyerSendGoods` | Boolean | yes | Whether the buyer has already shipped (if there is a return process) | true |
| `isAftersaleAgreeTimeout` | Boolean | yes | After-sales timeout flag | false |
| `isInsufficientBail` | Boolean | yes | Insufficient deposit for instant refund | false |
| `refundOfficialSolutionCost` | Long | yes | Fee for returning an official logistics pickup order (cents/fen) | 770 |

<a id="m-alibaba-trade-refund-oporderrefundoperationmodel[]"></a>
#### alibaba.trade.refund.OpOrderRefundOperationModel[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `afterOperateStatus` | java.lang.String | yes | Refund status after the operation |  |
| `beforeOperateStatus` | java.lang.String | yes | Refund status before the operation |  |
| `closeRefundStepId` | long | yes | Stage ID when a staged order's forward operation closes the refund |  |
| `crmModifyRefund` | Boolean | yes | Whether the refund order has been modified by customer service (Xiaoer) |  |
| `discription` | java.lang.String | yes | Description, notes |  |
| `email` | java.lang.String | yes | Contact EMAIL |  |
| `freightBill` | java.lang.String | yes | Waybill number |  |
| `gmtCreate` | java.util.Date | yes | Creation time |  |
| `gmtModified` | java.util.Date | yes | Modification time |  |
| `id` | java.lang.Long | yes | Primary key, refund operation record serial number |  |
| `messageStatus` | int | yes | Voucher status. 1: normal; 2: blocked by backend staff |  |
| `mobile` | java.lang.String | yes | Contact person's mobile number |  |
| `msgType` | int | yes | Message type. 3: staff message to buyer and seller; 4: message to buyer; 5: message to seller; 7: regular CBU message, equivalent to Taobao's 1 |  |
| `operateRemark` | java.lang.String | yes | Operation remarks |  |
| `operateTypeInt` | int | yes | Operation type, replaces operateType |  |
| `operatorId` | java.lang.String | yes | Operator - memberID |  |
| `operatorLoginId` | java.lang.String | yes | Operator - loginID |  |
| `operatorRoleId` | java.lang.Integer | yes | Operator role name: buyer, seller, system |  |
| `operatorUserId` | java.lang.Long | yes | Operator - userID |  |
| `phone` | java.lang.String | yes | Contact phone number |  |
| `refundAddress` | java.lang.String | yes | Return address |  |
| `refundId` | java.lang.String | yes | Refund record ID |  |
| `rejectReason` | java.lang.String | yes | Reason for seller's refund rejection |  |
| `vouchers` | java.util.List | yes | Voucher image URL |  |
| `logisticsCompany` | [message:alibaba.trade.refund.OpLogisticsCompanyModel](#m-alibaba-trade-refund-oplogisticscompanymodel) | yes | Logistics company details |  |

<a id="m-alibaba-trade-refund-oplogisticscompanymodel"></a>
#### alibaba.trade.refund.OpLogisticsCompanyModel

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `companyName` | java.lang.String | yes | Courier company name |  |
| `companyNo` | java.lang.String | yes | Logistics company number |  |
| `companyPhone` | java.lang.String | yes | Logistics company service phone number |  |
| `gmtCreate` | java.util.Date | yes | Creation time |  |
| `gmtModified` | java.util.Date | yes | Modification time |  |
| `id` | java.lang.Long | yes | ID |  |
| `spelling` | java.lang.String | yes | Full pinyin |  |
| `supportPrint` | Boolean | yes | Whether printing is supported |  |

## Error codes

| Error | Symptom | How to fix |
|---|---|---|
| 500 | &quot;{\&quot;errorCode\&quot;:\&quot;003002\&quot;,\&quot;errorMessage\&quot;:\&quot;SERVICE:INVOKE_FAIL:OrderRefundService.queryBatchRefundByOrderIdAndStatus\&quot;,\&quot;cause\&quot;:\&quot;errorCode:ORDER_NOT_EXIST,errorMsg:null,cause:null\&quot;}&quot; | Incorrect order number |

## Samples

**Description of the extInfo parameter**

```
payMode     资金分流（alipay:支付宝，fundFreeze:资金平台账户冻结） 
 7d          7天无理由订单标记（1表示7天无理由订单）  
```

**Request parameter example**

```
{"orderId":"151267031008969811",
"queryType" :"3",
"dipsuteType":1}
```

**Example of return parameters**

```
{"result":{"opOrderRefundModels":[{"applyReasonId":20021,"applySubReasonId":-1,"buyerUserId":1624961198,"alipayPaymentId":"2018042321001008760555342267","isCrmModifyRefund":false,"rejectTimes":0,"isTimeOutFreeze":false,"applyReason":"颜色/图案/款式不符","isInsufficientAccount":false,"rejectReasonId":0,"buyerLoginId":"alitestforisv02","refundCarriage":0,"isGoodsReceived":true,"sellerUserId":1623492085,"isOnlyRefund":false,"status":"refundsuccess","isRefundGoods":true,"refundOperationList":[],"isSellerDelayDisburse":false,"sellerMemberId":"b2b-1623492085","sellerLoginId":"alitestforisv01","gmtModified":"20180423174714000+0800","isAftersaleAutoDisburse":false,"disputeRequest":3,"disputeType":1,"id":8651493722961198,"isSupportNewSteppay":true,"canRefundPayment":1,"tradeTypeStr":"50060","freightBill":"12345134135435234","gmtApply":"20180423171116000+0800","goodsStatus":3,"gmtCompleted":"20180423174713000+0800","applyPayment":1,"buyerLogisticsName":"其他","orderId":151267031008969811,"isNewRefundReturn":true,"applyCarriage":0,"gmtCreate":"20180423171116000+0800","isBuyerSendGoods":true,"buyerMemberId":"b2b-1624961198","frozenFund":-1,"refundId":"TQ8651493722961198","extInfo":{"EXmrf":"1","ttid":"2","prepaidFailure":"QUERY_BUYER_CREDIT_LEVEL_FAIL","pay_lock":"seller","reason":"20021","apply_text_id":"null","newRefund":"rp2","bizCode":"cbu.general.refund","disputeTradeStatus":"4","lastOrder":"0","seller_agreed_refund_fee":"1","old_reason_id":"20021","sync":"0","b2b_seller_mId":"b2b-1623492085","seller_batch":"true","itemBuyAmount":"0","b2b_buyer_mId":"b2b-1624961198","ability":"1","refundPostFee":"0","seller_audit":"0","ee_trace_id":"0b802ca915244768332374685efa19","ol_tf":"1","opRole":"user","apply_init_refund_fee":"1","isVirtual":"0","logisticsCompanyId":"-1","itemPrice":"0","interceptStatus":"0","refundFrom":"2","restartForXiaoer":"1","appName":"tosp-aftersales","abnormal_dispute_status":"0","payMode":"alipay","workflowName":"cbu_return_and_refund","sgr":"1","enfunddetail":"1","sellerDoRefundNick":"alitestforisv01","bgmtc":"2018-04-23 09:48:12"},"refundPayment":1,"productName":"欧洲站2018新款宽松显瘦烫金印花字母老鹰短袖T恤打底裤套装女潮 黑 XL","success":true,"isAftersaleAgreeTimeout":false,"isInsufficientBail":false}]}}
```
