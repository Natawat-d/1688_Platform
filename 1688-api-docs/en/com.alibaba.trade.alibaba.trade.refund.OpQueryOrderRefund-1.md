# Query refund details by refund ID (buyer view)

Original name: 查询退款单详情-根据退款单ID（买家视角）  
API: `com.alibaba.trade:alibaba.trade.refund.OpQueryOrderRefund:1` · Category: Returns & Refunds  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.trade.refund.OpQueryOrderRefund-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.trade.refund.OpQueryOrderRefund/{appKey}`  
Requires user authorization (access_token) · Requires signature

For buyers; sellers should use alibaba.trade.refund.OpQueryOrderRefund.sellerView. Queries refund details, including the list of refund operations. Permission must be requested from Alibaba to access this API.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `refundId` | String | yes | Business primary key of the refund order: TQ+ID | TQ11173622***991577 |
| `needTimeOutInfo` | boolean | no | Timeout information needed for the refund order | true |
| `needOrderRefundOperation` | boolean | no | All refund operation information that must accompany the refund order | true |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.trade.refund.result.OpQueryOrderRefund](#m-alibaba-trade-refund-result-opqueryorderrefund) | yes | Query result | {} |
| `errorCode` | String | yes | Error code | 500 |
| `errorMessage` | String | yes | Error description | 退款单错误 |
| `extErrorMessage` | String | yes | Additional error description information | 退款单错误 |

<a id="m-alibaba-trade-refund-result-opqueryorderrefund"></a>
#### alibaba.trade.refund.result.OpQueryOrderRefund

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `opOrderRefundModelDetail` | [message:alibaba.trade.refund.OpOrderRefundModel](#m-alibaba-trade-refund-oporderrefundmodel) | yes | Return value | {} |

<a id="m-alibaba-trade-refund-oporderrefundmodel"></a>
#### alibaba.trade.refund.OpOrderRefundModel

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
| `asynErrCode` | java.lang.String | yes |  |    |
| `asynSubErrCode` | java.lang.String | yes |  |     |
| `buyerAlipayId` | java.lang.String | yes | Buyer's Alipay ID |    |
| `buyerLogisticsName` | java.lang.String | yes | Buyer's return logistics company name |    |
| `buyerMemberId` | java.lang.String | yes | Buyer member ID |    |
| `buyerSendGoods` | Boolean | yes | Whether the buyer has already shipped (if there is a return process) |    |
| `buyerUserId` | java.lang.Long | yes | Buyer's Alibaba account ID (including Taobao account ID) |    |
| `canRefundPayment` | java.lang.Long | yes | Maximum refundable amount, in cents |    |
| `crmModifyRefund` | Boolean | yes | Whether the refund order has been modified by customer service (Xiaoer) |    |
| `disburseChannel` | java.lang.String | yes | Instant-arrival payment channel |    |
| `disputeRequest` | int | yes | After-sales refund requirement |    |
| `disputeType` | int | yes | Dispute type: in-sale refund, after-sales refund; defaults to in-sale refund |    |
| `extInfo` | java.util.Map | yes | Extended information |    |
| `freightBill` | java.lang.String | yes | Waybill number |    |
| `frozenFund` | long | yes | Actual frozen account amount, unit: cent |    |
| `gmtApply` | java.util.Date | yes | Refund application time |    |
| `gmtCompleted` | java.util.Date | yes | Completion time |    |
| `gmtCreate` | java.util.Date | yes | Creation time |    |
| `gmtFreezed` | java.util.Date | yes | The timeout freeze start time for this refund order |    |
| `gmtModified` | java.util.Date | yes | Modification time |    |
| `gmtTimeOut` | java.util.Date | yes | The time limit by which this refund order must be completed before timing out |    |
| `goodsReceived` | Boolean | yes | Whether the buyer has received the goods |    |
| `goodsStatus` | int | yes | 1: buyer has not received the goods<br>2: buyer has received the goods<br>3: buyer has returned the goods |    |
| `id` | java.lang.Long | yes | Refund order number |    |
| `instantRefundType` | java.lang.String | yes | Instant refund type |    |
| `insufficientAccount` | Boolean | yes | Transaction 4.0 refund balance insufficient |    |
| `insufficientBail` | Boolean | yes | Insufficient deposit for instant refund |    |
| `newRefundReturn` | Boolean | yes | Whether the refund/return was created via the new process |    |
| `onlyRefund` | Boolean | yes | Whether it is refund only |    |
| `orderEntryCountMap` | java.util.Map | yes | Sub-order return quantity |    |
| `orderEntryIdList` | java.util.List | yes | Order details included in the refund order, listed in reverse chronological order |    |
| `orderId` | java.lang.Long | yes | Order number corresponding to the refund order |    |
| `prepaidBalance` | java.lang.Long | yes | Express refund advance funding amount. When this value is not empty, it only means this refund order is eligible for the advance funding process, not that the advance funding will necessarily succeed |    |
| `productName` | java.lang.String | yes | Product name (the item name from the order detail associated with the refund order) |    |
| `refundCarriage` | java.lang.Long | yes | Actual refund amount for shipping fee, in cents |    |
| `refundGoods` | Boolean | yes | Whether return is required |    |
| `refundId` | java.lang.String | yes | Refund order logical primary key |    |
| `refundPayment` | java.lang.Long | yes | Actual refund amount, in cents |    |
| `rejectReason` | java.lang.String | yes | Seller rejection reason |    |
| `rejectReasonId` | int | yes | Seller rejection reason Id |    |
| `rejectTimes` | int | yes | Number of times the refund order was rejected |    |
| `sellerAlipayId` | java.lang.String | yes | Seller's Alipay ID |    |
| `sellerDelayDisburse` | Boolean | yes | Whether the seller delays payment, i.e., safe refund |    |
| `sellerMemberId` | java.lang.String | yes | Seller member ID |    |
| `sellerMobile` | java.lang.String | yes | Recipient's mobile phone |    |
| `sellerRealName` | java.lang.String | yes | Recipient name |    |
| `sellerReceiveAddress` | java.lang.String | yes | Seller's receiving address for buyer returns |     |
| `sellerTel` | java.lang.String | yes | Consignee phone number |     |
| `sellerUserId` | java.lang.Long | yes | Seller's Alibaba account ID (including Taobao account ID) |     |
| `status` | java.lang.String | yes | Refund status |     |
| `supportNewSteppay` | Boolean | yes | Whether Transaction 4.0 is supported |    |
| `taskStatus` | java.lang.String | yes | Ticket sub-status; empty when the ticket has not yet flowed to CRM for creation |    |
| `timeOutFreeze` | Boolean | yes | Whether frozen by the system due to timeout. true means frozen, false means not frozen. |    |
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
| `afterOperateStatus` | java.lang.String | yes | Refund status after the operation | waitsellerreceive |
| `beforeOperateStatus` | java.lang.String | yes | Refund status before the operation | waitsellerreceive |
| `closeRefundStepId` | long | yes | Stage ID when a staged order's forward operation closes the refund | 0 |
| `crmModifyRefund` | Boolean | yes | Whether the refund order has been modified by customer service (Xiaoer) | false |
| `discription` | java.lang.String | yes | Description, notes |  退款 |
| `email` | java.lang.String | yes | Contact EMAIL |   |
| `freightBill` | java.lang.String | yes | Waybill number | 7060***4101371 |
| `gmtCreate` | java.util.Date | yes | Creation time | 20180521141339000+0800 |
| `gmtModified` | java.util.Date | yes | Modification time | 20180521141339000+0800 |
| `id` | java.lang.Long | yes | Primary key, refund operation record serial number | 27266762834 |
| `messageStatus` | int | yes | Voucher status. 1: normal; 2: blocked by backend staff | 1 |
| `mobile` | java.lang.String | yes | Contact person's mobile number |   |
| `msgType` | int | yes | Message type. 3: staff message to buyer and seller; 4: message to buyer; 5: message to seller; 7: regular CBU message, equivalent to Taobao's 1 | 7 |
| `operateRemark` | java.lang.String | yes | Operation remarks |  卖家拒绝协议，等待买家修改 |
| `operateTypeInt` | int | yes | Operation type, replaces operateType | 15 |
| `operatorId` | java.lang.String | yes | Operator - memberID |   |
| `operatorLoginId` | java.lang.String | yes | Operator - loginID | alitestforisv02 |
| `operatorRoleId` | java.lang.Integer | yes | Operator role name: buyer, seller, system | 4 |
| `operatorUserId` | java.lang.Long | yes | Operator - userID | 1623***085 |
| `phone` | java.lang.String | yes | Contact phone number |   |
| `refundAddress` | java.lang.String | yes | Return address | 网商路66号 |
| `refundId` | java.lang.String | yes | Refund record ID | TQ83489**0492085 |
| `rejectReason` | java.lang.String | yes | Reason for seller's refund rejection | 无方案 |
| `vouchers` | java.util.List | yes | Voucher image URL | null |
| `logisticsCompany` | [message:alibaba.trade.refund.OpLogisticsCompanyModel](#m-alibaba-trade-refund-oplogisticscompanymodel) | yes | Logistics company details | {} |

<a id="m-alibaba-trade-refund-oplogisticscompanymodel"></a>
#### alibaba.trade.refund.OpLogisticsCompanyModel

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `companyName` | java.lang.String | yes | Courier company name | 百世快递 |
| `companyNo` | java.lang.String | yes | Logistics company number | HTKY |
| `companyPhone` | java.lang.String | yes | Logistics company service phone number |   |
| `gmtCreate` | java.util.Date | yes | Creation time |   |
| `gmtModified` | java.util.Date | yes | Modification time | 20160926142611000+0800 |
| `id` | java.lang.Long | yes | ID | 352 |
| `spelling` | java.lang.String | yes | Full pinyin |   |
| `supportPrint` | Boolean | yes | Whether printing is supported |   |

## Error codes

| Error | Symptom | How to fix |
|---|---|---|
| 500 | {\&quot;errorCode\&quot;:\&quot;003\&quot;,\&quot;errorMessage\&quot;:\&quot;SERVICE: failed to query data, please check the input parameters and try again.\&quot;,\&quot;cause\&quot;:\&quot;\&quot;} | Refund order of an unauthorized user |
| 500 | {\&quot;errorCode\&quot;:\&quot;003002\&quot;,\&quot;errorMessage\&quot;:\&quot;SERVICE:INVOKE_FAIL:Refund data error, please check the refund order number; the refund order number format is usually TD+id\&quot;,\&quot;cause\&quot;:\&quot;errorCode:REFUND_DATA_ERROR,errorMsg:refundId : 11043002311780591,cause:null\&quot;} | Refund order is incorrect |
| 500 | {\&quot;errorCode\&quot;:\&quot;003002\&quot;,\&quot;errorMessage\&quot;:\&quot;SERVICE:INVOKE_FAIL:OrderRefundService#queryOrderRefundOperationList\&quot;,\&quot;cause\&quot;:\&quot;errorCode:INVALID_PARAM,errorMsg:null,cause:null\&quot;} | Refund order is incorrect |

## Samples

**Description of the extInfo parameter**

```
payMode     资金分流（alipay:支付宝，fundFreeze:资金平台账户冻结） 
 7d          7天无理由订单标记（1表示7天无理由订单） 
```

**Request parameter example**

```
{"refundId":"TQ865149372**61198",
"needTimeOutInfo":true,
"needOrderRefundOperation":true}
```

**Example of return parameters**

```
{"result":{"opOrderRefundModelDetail":{"alipayPaymentId":"2018042321001008760555342267","applyCarriage":0,"applyPayment":1,"applyReason":"颜色/图案/款式不符","applyReasonId":20021,"applySubReasonId":0,"buyerLogisticsName":"其他","buyerMemberId":"b2b-1624961198","buyerSendGoods":true,"buyerUserId":1624961198,"canRefundPayment":1,"crmModifyRefund":false,"disputeRequest":3,"disputeType":1,"extInfo":{"EXmrf":"1","ttid":"2","prepaidFailure":"QUERY_BUYER_CREDIT_LEVEL_FAIL","pay_lock":"seller","reason":"20021","apply_text_id":"null","newRefund":"rp2","bizCode":"cbu.general.refund","disputeTradeStatus":"4","lastOrder":"0","seller_agreed_refund_fee":"1","old_reason_id":"20021","sync":"0","b2b_seller_mId":"b2b-1623492085","seller_batch":"true","itemBuyAmount":"0","b2b_buyer_mId":"b2b-1624961198","ability":"1","refundPostFee":"0","seller_audit":"0","ee_trace_id":"0b802ca915244768332374685efa19","ol_tf":"1","opRole":"user","apply_init_refund_fee":"1","isVirtual":"0","logisticsCompanyId":"-1","itemPrice":"0","interceptStatus":"0","refundFrom":"2","restartForXiaoer":"1","appName":"tosp-aftersales","abnormal_dispute_status":"0","payMode":"alipay","workflowName":"cbu_return_and_refund","sgr":"1","enfunddetail":"1","sellerDoRefundNick":"alitestforisv01","bgmtc":"2018-04-23 09:48:12"},"freightBill":"12345134135435234","frozenFund":-1,"gmtApply":"20180423171116000+0800","gmtCompleted":"20180423174713000+0800","gmtCreate":"20180423171116000+0800","gmtModified":"20180423174714000+0800","goodsReceived":true,"goodsStatus":3,"id":8651493722961198,"newRefundReturn":true,"onlyRefund":false,"orderEntryCountMap":{"151267031009969811":2},"orderEntryIdList":[151267031009969811],"orderId":151267031008969811,"productName":"欧洲站2018新款宽松显瘦烫金印花字母老鹰短袖T恤打底裤套装女潮","refundCarriage":0,"refundGoods":true,"refundId":"TQ8651493722961198","refundPayment":1,"rejectReasonId":0,"rejectTimes":0,"sellerDelayDisburse":false,"sellerMemberId":"b2b-1623492085","sellerMobile":"19926555555","sellerRealName":"琳琳","sellerReceiveAddress":"山东 聊城 解决了交流交流链接连接","sellerUserId":1623492085,"status":"refundsuccess","supportNewSteppay":true,"timeOutFreeze":false,"tradeTypeStr":"50060","refundOperationList":[{"afterOperateStatus":"waitsellerreceive","beforeOperateStatus":"waitsellerreceive","closeRefundStepId":0,"crmModifyRefund":false,"discription":"卖家同意退货退款协议，退货退款成功","gmtCreate":"20180423174714000+0800","gmtModified":"20180423174714000+0800","id":50027795018,"messageStatus":3,"msgType":7,"operateRemark":"退款成功","operateTypeInt":8,"operatorLoginId":"alitestforisv01","operatorRoleId":2,"operatorUserId":1623492085,"refundId":"TQ8651493722961198"},{"afterOperateStatus":"waitbuyersend","beforeOperateStatus":"waitbuyersend","closeRefundStepId":0,"crmModifyRefund":false,"discription":"买家声明退货，等待卖家确认。|物流公司：其他|物流单号：12345134135435234|说明：asdada|","gmtCreate":"20180423174651000+0800","gmtModified":"20180423174651000+0800","id":50027571423,"messageStatus":3,"msgType":7,"operateRemark":"买家退货","operateTypeInt":10,"operatorLoginId":"alitestforisv02","operatorRoleId":1,"operatorUserId":1624961198,"refundId":"TQ8651493722961198"},{"afterOperateStatus":"waitselleragree","beforeOperateStatus":"waitselleragree","closeRefundStepId":0,"crmModifyRefund":false,"discription":"卖家确认收货地址：琳琳,19926555555,山东 聊城 解决了交流交流链接连接","gmtCreate":"20180423174544000+0800","gmtModified":"20180423174544000+0800","id":49967844797,"messageStatus":3,"msgType":7,"operateRemark":"确认收货地址","operateTypeInt":19,"operatorLoginId":"alitestforisv01","operatorRoleId":2,"operatorUserId":1623492085,"refundId":"TQ8651493722961198"},{"closeRefundStepId":0,"crmModifyRefund":false,"discription":"退款诉求：退货退款|申请原因：颜色/图案/款式不符|货品情况：已收到货|退款货品金额：0.01|退款运费金额：0.00|退款说明：asdad","gmtCreate":"20180423171116000+0800","gmtModified":"20180423171116000+0800","id":49962356896,"messageStatus":3,"msgType":7,"operateRemark":"买家申请退款协议，等待卖家确认","operateTypeInt":1,"operatorLoginId":"alitestforisv02","operatorRoleId":1,"operatorUserId":1624961198,"refundId":"TQ8651493722961198"}],"buyerLoginId":"alitestforisv02","sellerLoginId":"alitestforisv01"}}}
```
