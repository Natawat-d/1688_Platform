# Query refund list (buyer view)

Original name: 查询退款单列表(买家视角)  
API: `com.alibaba.trade:alibaba.trade.refund.buyer.queryOrderRefundList:1` · Category: Returns & Refunds  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.trade.refund.buyer.queryOrderRefundList-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.trade.refund.buyer.queryOrderRefundList/{appKey}`  
Requires user authorization (access_token) · Requires signature

Buyer views the refund list. This interface does not support sub-account queries; authorize with the main account before querying.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `orderId` | Long | no | Order ID | 179087886005498520 |
| `applyStartTime` | java.util.Date | no | Refund application time (start) | 20170926114526000+0800 |
| `applyEndTime` | java.util.Date | no | Refund application time (end/deadline) | 20220926114526000+0800 |
| `refundStatusSet` | String[] | no | Refund status list | 等待卖家同意 waitselleragree;退款成功 refundsuccess;退款关闭 refundclose;待买家修改 waitbuyermodify;等待买家退货 waitbuyersend;等待卖家确认收货 waitsellerreceive |
| `sellerMemberId` | String | no | Seller memberId | b2b-1623492085 |
| `currentPageNum` | Integer | no | Current page number | 0 |
| `pageSize` | Integer | no | Number of records per page | 20 |
| `logisticsNo` | String | no | Return logistics waybill number (when passing this field for a query, sellerMemberId must also be passed) | 3101***159271 |
| `modifyStartTime` | java.util.Date | no | Refund modification time (start) | 20170926114526000+0800 |
| `modifyEndTime` | java.util.Date | no | Refund modification time (end) | 20220926114526000+0800 |
| `dipsuteType` | Integer | no | 1: in-sale refund, 2: after-sales refund; 0: all refund orders | 1 |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.trade.refund.OpQueryOrderRefundListResult](#m-alibaba-trade-refund-opqueryorderrefundlistresult) | yes | Query result | {} |
| `errorCode` | String | yes | Error code |   |
| `errorMsg` | String | yes | Error message |   |
| `success` | Boolean | yes | Whether successful | true |

<a id="m-alibaba-trade-refund-opqueryorderrefundlistresult"></a>
#### alibaba.trade.refund.OpQueryOrderRefundListResult

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `opOrderRefundModels` | [message:alibaba.trade.refund.OpOrderRefundModel[]](#m-alibaba-trade-refund-oporderrefundmodel[]) | yes | List of refund orders |  |
| `totalCount` | int | yes | Total number of records that meet the criteria |  |
| `currentPageNum` | int | yes | Current page number of the query |  |

<a id="m-alibaba-trade-refund-oporderrefundmodel[]"></a>
#### alibaba.trade.refund.OpOrderRefundModel[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `aftersaleAgreeTimeout` | Boolean | yes | After-sales timeout flag |   |
| `aftersaleAutoDisburse` | Boolean | yes | After-sales automatic payment |   |
| `alipayPaymentId` | java.lang.String | yes | Alipay transaction number |   |
| `applyCarriage` | java.lang.Long | yes | Shipping fee refund amount requested, in cents |   |
| `applyExpect` | java.lang.Long | yes | Refund amount originally entered by the buyer (can be empty) |   |
| `applyPayment` | java.lang.Long | yes | Refund amount requested by the buyer, in cents |   |
| `applyReason` | java.lang.String | yes | Application reason |   |
| `applyReasonId` | int | yes | Application reason ID |   |
| `applySubReason` | java.lang.String | yes | Secondary refund reason |   |
| `applySubReasonId` | int | yes | Secondary refund reason ID |   |
| `asynErrCode` | java.lang.String | yes |  |   |
| `asynSubErrCode` | java.lang.String | yes |  |   |
| `buyerAlipayId` | java.lang.String | yes | Buyer's Alipay ID |   |
| `buyerLogisticsName` | java.lang.String | yes | Buyer's return logistics company name |   |
| `buyerMemberId` | java.lang.String | yes | Buyer member ID |   |
| `buyerSendGoods` | Boolean | yes | Whether the buyer has already shipped (if there is a return process) |   |
| `buyerUserId` | java.lang.Long | yes | Buyer's Alibaba account ID (including Taobao account ID) |   |
| `canRefundPayment` | java.lang.Long | yes | Maximum refundable amount, in cents |   |
| `crmModifyRefund` | Boolean | yes | Whether the refund order has been modified by customer service (Xiaoer) |   |
| `disburseChannel` | java.lang.String | yes | Instant-arrival payment channel |   |
| `disputeRequest` | int | yes | After-sales refund requirement |   |
| `disputeType` | int | yes | Dispute type: in-sale refund, after-sales refund; defaults to in-sale refund |   |
| `extInfo` | java.util.Map | yes | Extended information |   |
| `freightBill` | java.lang.String | yes | Waybill number |   |
| `frozenFund` | long | yes | Actual frozen account amount, unit: cent |   |
| `gmtApply` | java.util.Date | yes | Refund application time |   |
| `gmtCompleted` | java.util.Date | yes | Completion time |   |
| `gmtCreate` | java.util.Date | yes | Creation time |   |
| `gmtFreezed` | java.util.Date | yes | The timeout freeze start time for this refund order |   |
| `gmtModified` | java.util.Date | yes | Modification time |   |
| `gmtTimeOut` | java.util.Date | yes | The time limit by which this refund order must be completed before timing out |   |
| `goodsReceived` | Boolean | yes | Whether the buyer has received the goods |   |
| `goodsStatus` | int | yes | 1: buyer has not received the goods<br>2: buyer has received the goods<br>3: buyer has returned the goods |   |
| `id` | java.lang.Long | yes | Refund order number |   |
| `instantRefundType` | java.lang.String | yes | Instant refund type |   |
| `insufficientAccount` | Boolean | yes | Transaction 4.0 refund balance insufficient |   |
| `insufficientBail` | Boolean | yes | Insufficient deposit for instant refund |   |
| `newRefundReturn` | Boolean | yes | Whether the refund/return was created via the new process |   |
| `onlyRefund` | Boolean | yes | Whether it is refund only |   |
| `orderEntryCountMap` | java.util.Map | yes | Sub-order return quantity |   |
| `orderEntryIdList` | java.util.List | yes | Order details included in the refund order, listed in reverse chronological order |   |
| `orderId` | java.lang.Long | yes | Order number corresponding to the refund order |   |
| `prepaidBalance` | java.lang.Long | yes | Express refund advance funding amount. When this value is not empty, it only means this refund order is eligible for the advance funding process, not that the advance funding will necessarily succeed |   |
| `productName` | java.lang.String | yes | Product name (the item name from the order detail associated with the refund order) |   |
| `refundCarriage` | java.lang.Long | yes | Actual refund amount for shipping fee, in cents |   |
| `refundGoods` | Boolean | yes | Whether return is required |   |
| `refundId` | java.lang.String | yes | Refund order logical primary key |   |
| `refundPayment` | java.lang.Long | yes | Actual refund amount, in cents |   |
| `rejectReason` | java.lang.String | yes | Seller rejection reason |   |
| `rejectReasonId` | int | yes | Seller rejection reason Id |   |
| `rejectTimes` | int | yes | Number of times the refund order was rejected |   |
| `sellerAlipayId` | java.lang.String | yes | Seller's Alipay ID |   |
| `sellerDelayDisburse` | Boolean | yes | Whether the seller delays payment, i.e., safe refund |   |
| `sellerMemberId` | java.lang.String | yes | Seller member ID |   |
| `sellerMobile` | java.lang.String | yes | Recipient's mobile phone |   |
| `sellerRealName` | java.lang.String | yes | Recipient name |   |
| `sellerReceiveAddress` | java.lang.String | yes | Seller's receiving address for buyer returns |   |
| `sellerTel` | java.lang.String | yes | Consignee phone number |   |
| `sellerUserId` | java.lang.Long | yes | Seller's Alibaba account ID (including Taobao account ID) |   |
| `status` | java.lang.String | yes | Refund status |   |
| `supportNewSteppay` | Boolean | yes | Whether Transaction 4.0 is supported |   |
| `taskStatus` | java.lang.String | yes | Ticket sub-status; empty when the ticket has not yet flowed to CRM for creation |   |
| `timeOutFreeze` | Boolean | yes | Whether frozen by the system due to timeout. true means frozen, false means not frozen. |   |
| `timeOutOperateType` | java.lang.String | yes | Action executed after timeout |   |
| `tradeTypeStr` | java.lang.String | yes | Trade type, used to replace the enum type tradeType |   |
| `refundOperationList` | [message:alibaba.trade.refund.OpOrderRefundOperationModel[]](#m-alibaba-trade-refund-oporderrefundoperationmodel[]) | yes | Operation record list |   |
| `buyerLoginId` | java.lang.String | yes | Buyer LoginId |   |
| `sellerLoginId` | java.lang.String | yes | Seller LoginId |   |
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
| APPLY_TIME_ERROR | Only refund orders from the last 730 days can be queried. Please check the refund application time! | Please check the refund application time! |

## Samples

**Description of the extInfo parameter**

```
payMode     资金分流（alipay:支付宝，fundFreeze:资金平台账户冻结） 
7d          7天无理由订单标记（1表示7天无理由订单）     
```

**Request parameter example**

```
{"orderId":"177681528398969811",
"dipsuteType":0}
```

**Example of return parameters**

```
{"result":{"opOrderRefundModels":[{"alipayPaymentId":"2018061421001008760569923894","applyCarriage":0,"applyPayment":1,"applyReason":"不想买了，已与卖家协商一致","applyReasonId":20028,"applySubReasonId":-1,"buyerMemberId":"b2b-1624961198","buyerUserId":1624961198,"canRefundPayment":1,"disputeRequest":3,"disputeType":1,"extInfo":{"EXmrf":"1","ttid":"2","reason":"20028","apply_text_id":"null","newRefund":"rp2","bizCode":"cbu.general.refund","lastOrder":"0","seller_agreed_refund_fee":"1","old_reason_id":"20028","b2b_seller_mId":"b2b-1623492085","seller_batch":"true","itemBuyAmount":"0","b2b_buyer_mId":"b2b-1624961198","ability":"1","seller_audit":"0","ee_trace_id":"0ab2dbd215300798520175927d07d3","ol_tf":"0","opRole":"timeout","apply_init_refund_fee":"1","isVirtual":"0","logisticsCompanyId":"-1","itemPrice":"0","interceptStatus":"0","refundFrom":"2","restartForXiaoer":"1","appName":"tosp-aftersales","abnormal_dispute_status":"0","payMode":"alipay","workflowName":"cbu_return_and_refund","sgr":"1","enfunddetail":"1","bgmtc":"2018-06-14 14:20:21"},"frozenFund":-1,"gmtApply":"20180622141034000+0800","gmtCompleted":"20180702141113000+0800","gmtCreate":"20180622141034000+0800","gmtModified":"20180702141113000+0800","goodsStatus":2,"id":10426553019961198,"orderId":177681528398969811,"productName":"短袖撞色领保罗衫定做纯棉广告衫定制印字logo刺绣男现货批发 白色/橙领 S等3种","refundCarriage":0,"refundId":"TQ10426553019961198","refundPayment":1,"rejectReasonId":0,"rejectTimes":0,"sellerMemberId":"b2b-1623492085","sellerUserId":1623492085,"status":"refundclose","tradeTypeStr":"50060","refundOperationList":[]}],"totalCount":1,"currentPageNum":0},"success":true}
```
