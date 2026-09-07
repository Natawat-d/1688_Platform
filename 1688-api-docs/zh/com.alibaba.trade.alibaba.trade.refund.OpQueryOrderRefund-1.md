# 查询退款单详情-根据退款单ID（买家视角）

API: `com.alibaba.trade:alibaba.trade.refund.OpQueryOrderRefund:1` · Category: 退货退款  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.trade.refund.OpQueryOrderRefund-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.trade.refund.OpQueryOrderRefund/{appKey}`  
需要授权 (access_token) · 需要签名

该API为买家使用，卖家查询请使用alibaba.trade.refund.OpQueryOrderRefund.sellerView，查询退款单详情，同时可以查询到退款操作列表。 该API需要向阿里巴巴申请权限才能访问。

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `refundId` | String | 是 | 退款单业务主键 TQ+ID | TQ11173622***991577 |
| `needTimeOutInfo` | boolean | 否 | 需要退款单的超时信息 | true |
| `needOrderRefundOperation` | boolean | 否 | 需要退款单伴随的所有退款操作信息 | true |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:alibaba.trade.refund.result.OpQueryOrderRefund](#m-alibaba-trade-refund-result-opqueryorderrefund) | 是 | 查询结果 | {} |
| `errorCode` | String | 是 | 错误码 | 500 |
| `errorMessage` | String | 是 | 错误描述信息 | 退款单错误 |
| `extErrorMessage` | String | 是 | 补充错误描述信息 | 退款单错误 |

<a id="m-alibaba-trade-refund-result-opqueryorderrefund"></a>
#### alibaba.trade.refund.result.OpQueryOrderRefund

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `opOrderRefundModelDetail` | [message:alibaba.trade.refund.OpOrderRefundModel](#m-alibaba-trade-refund-oporderrefundmodel) | 是 | 返回值 | {} |

<a id="m-alibaba-trade-refund-oporderrefundmodel"></a>
#### alibaba.trade.refund.OpOrderRefundModel

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `aftersaleAgreeTimeout` | Boolean | 是 | 售后超时标记 |   |
| `aftersaleAutoDisburse` | Boolean | 是 | 售后自动打款 |   |
| `alipayPaymentId` | java.lang.String | 是 | 支付宝交易号 |   |
| `applyCarriage` | java.lang.Long | 是 | 运费的申请退款金额，单位：分 |   |
| `applyExpect` | java.lang.Long | 是 | 买家原始输入的退款金额(可以为空) |   |
| `applyPayment` | java.lang.Long | 是 | 买家申请退款金额，单位：分 |   |
| `applyReason` | java.lang.String | 是 | 申请原因 |   |
| `applyReasonId` | int | 是 | 申请原因ID |   |
| `applySubReason` | java.lang.String | 是 | 二级退款原因 |   |
| `applySubReasonId` | int | 是 | 二级退款原因Id |   |
| `asynErrCode` | java.lang.String | 是 |  |    |
| `asynSubErrCode` | java.lang.String | 是 |  |     |
| `buyerAlipayId` | java.lang.String | 是 | 买家支付宝ID |    |
| `buyerLogisticsName` | java.lang.String | 是 | 买家退货物流公司名 |    |
| `buyerMemberId` | java.lang.String | 是 | 买家会员ID |    |
| `buyerSendGoods` | Boolean | 是 | 买家是否已经发货（如果有退货的流程） |    |
| `buyerUserId` | java.lang.Long | 是 | 买家阿里帐号ID(包括淘宝帐号Id) |    |
| `canRefundPayment` | java.lang.Long | 是 | 最大能够退款金额，单位：分 |    |
| `crmModifyRefund` | Boolean | 是 | 是否小二修改过退款单 |    |
| `disburseChannel` | java.lang.String | 是 | 极速到账打款渠道 |    |
| `disputeRequest` | int | 是 | 售后退款要求 |    |
| `disputeType` | int | 是 | 纠纷类型：售中退款 售后退款，默认为售中退款 |    |
| `extInfo` | java.util.Map | 是 | 扩展信息 |    |
| `freightBill` | java.lang.String | 是 | 运单号 |    |
| `frozenFund` | long | 是 | 实际冻结账户金额,单位：分 |    |
| `gmtApply` | java.util.Date | 是 | 申请退款时间 |    |
| `gmtCompleted` | java.util.Date | 是 | 完成时间 |    |
| `gmtCreate` | java.util.Date | 是 | 创建时间 |    |
| `gmtFreezed` | java.util.Date | 是 | 该退款单超时冻结开始时间 |    |
| `gmtModified` | java.util.Date | 是 | 修改时间 |    |
| `gmtTimeOut` | java.util.Date | 是 | 该退款单超时完成的时间期限 |    |
| `goodsReceived` | Boolean | 是 | 买家是否已收到货 |    |
| `goodsStatus` | int | 是 | 1：买家未收到货<br>2：买家已收到货<br>3：买家已退货 |    |
| `id` | java.lang.Long | 是 | 退款单编号 |    |
| `instantRefundType` | java.lang.String | 是 | 极速到账退款类型 |    |
| `insufficientAccount` | Boolean | 是 | 交易4.0退款余额不足 |    |
| `insufficientBail` | Boolean | 是 | 极速到账退款保证金不足 |    |
| `newRefundReturn` | Boolean | 是 | 是否新流程创建的退款退货 |    |
| `onlyRefund` | Boolean | 是 | 是否仅退款 |    |
| `orderEntryCountMap` | java.util.Map | 是 | 子订单退货数量 |    |
| `orderEntryIdList` | java.util.List | 是 | 退款单包含的订单明细，时间逆序排列 |    |
| `orderId` | java.lang.Long | 是 | 退款单对应的订单编号 |    |
| `prepaidBalance` | java.lang.Long | 是 | 极速退款垫资金额,该值不为空时,只代表该退款单可以走垫资流程,但不代表一定垫资成功 |    |
| `productName` | java.lang.String | 是 | 产品名称(退款单关联订单明细的货品名称) |    |
| `refundCarriage` | java.lang.Long | 是 | 运费的实际退款金额，单位：分 |    |
| `refundGoods` | Boolean | 是 | 是否要求退货 |    |
| `refundId` | java.lang.String | 是 | 退款单逻辑主键 |    |
| `refundPayment` | java.lang.Long | 是 | 实际退款金额，单位：分 |    |
| `rejectReason` | java.lang.String | 是 | 卖家拒绝原因 |    |
| `rejectReasonId` | int | 是 | 卖家拒绝原因Id |    |
| `rejectTimes` | int | 是 | 退款单被拒绝的次数 |    |
| `sellerAlipayId` | java.lang.String | 是 | 卖家支付宝ID |    |
| `sellerDelayDisburse` | Boolean | 是 | 是否卖家延迟打款，即安全退款 |    |
| `sellerMemberId` | java.lang.String | 是 | 卖家会员ID |    |
| `sellerMobile` | java.lang.String | 是 | 收货人手机 |    |
| `sellerRealName` | java.lang.String | 是 | 收货人姓名 |    |
| `sellerReceiveAddress` | java.lang.String | 是 | 买家退货时卖家收货地址 |     |
| `sellerTel` | java.lang.String | 是 | 收货人电话 |     |
| `sellerUserId` | java.lang.Long | 是 | 卖家阿里帐号ID(包括淘宝帐号Id) |     |
| `status` | java.lang.String | 是 | 退款状态 |     |
| `supportNewSteppay` | Boolean | 是 | 是否支持交易4.0 |    |
| `taskStatus` | java.lang.String | 是 | 工单子状态，没有流到CRM创建工单时为空 |    |
| `timeOutFreeze` | Boolean | 是 | 是否超时系统冻结，true代表冻结，false代表不冻结 |    |
| `timeOutOperateType` | java.lang.String | 是 | 超时后执行的动作 |   |
| `tradeTypeStr` | java.lang.String | 是 | 交易类型，用来替换枚举类型的tradeType |   |
| `refundOperationList` | [message:alibaba.trade.refund.OpOrderRefundOperationModel[]](#m-alibaba-trade-refund-oporderrefundoperationmodel[]) | 是 | 操作记录列表 |   |
| `buyerLoginId` | java.lang.String | 是 | 买家LoginId |   |
| `sellerLoginId` | java.lang.String | 是 | 卖家LoginId |   |
| `refundOfficialSolutionCost` | Long | 是 | 退官方物流提货订单费用（分） | 770 |

<a id="m-alibaba-trade-refund-oporderrefundoperationmodel[]"></a>
#### alibaba.trade.refund.OpOrderRefundOperationModel[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `afterOperateStatus` | java.lang.String | 是 | 操作后的退款状态 | waitsellerreceive |
| `beforeOperateStatus` | java.lang.String | 是 | 操作前的退款状态 | waitsellerreceive |
| `closeRefundStepId` | long | 是 | 分阶段订单正向操作关闭退款时的阶段ID | 0 |
| `crmModifyRefund` | Boolean | 是 | 是否小二修改过退款单 | false |
| `discription` | java.lang.String | 是 | 描述、说明 |  退款 |
| `email` | java.lang.String | 是 | 联系人EMAIL |   |
| `freightBill` | java.lang.String | 是 | 运单号 | 7060***4101371 |
| `gmtCreate` | java.util.Date | 是 | 创建时间 | 20180521141339000+0800 |
| `gmtModified` | java.util.Date | 是 | 修改时间 | 20180521141339000+0800 |
| `id` | java.lang.Long | 是 | 主键，退款操作记录流水号 | 27266762834 |
| `messageStatus` | int | 是 | 凭证状态，1:正常 2:后台小二屏蔽 | 1 |
| `mobile` | java.lang.String | 是 | 联系人手机 |   |
| `msgType` | int | 是 | 留言类型 3:小二留言给买家和卖家 4:给买家的留言 5:给卖家的留言 7:cbu的普通留言等同于淘宝的1 | 7 |
| `operateRemark` | java.lang.String | 是 | 操作备注 |  卖家拒绝协议，等待买家修改 |
| `operateTypeInt` | int | 是 | 操作类型 取代operateType | 15 |
| `operatorId` | java.lang.String | 是 | 操作者-memberID |   |
| `operatorLoginId` | java.lang.String | 是 | 操作者-loginID | alitestforisv02 |
| `operatorRoleId` | java.lang.Integer | 是 | 操作者角色名称 买家 卖家 系统 | 4 |
| `operatorUserId` | java.lang.Long | 是 | 操作者-userID | 1623***085 |
| `phone` | java.lang.String | 是 | 联系人电话 |   |
| `refundAddress` | java.lang.String | 是 | 退货地址 | 网商路66号 |
| `refundId` | java.lang.String | 是 | 退款记录ID | TQ83489**0492085 |
| `rejectReason` | java.lang.String | 是 | 卖家拒绝退款原因 | 无方案 |
| `vouchers` | java.util.List | 是 | 凭证图片地址 | null |
| `logisticsCompany` | [message:alibaba.trade.refund.OpLogisticsCompanyModel](#m-alibaba-trade-refund-oplogisticscompanymodel) | 是 | 物流公司详情 | {} |

<a id="m-alibaba-trade-refund-oplogisticscompanymodel"></a>
#### alibaba.trade.refund.OpLogisticsCompanyModel

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `companyName` | java.lang.String | 是 | 快递公司名 | 百世快递 |
| `companyNo` | java.lang.String | 是 | 物流公司编号 | HTKY |
| `companyPhone` | java.lang.String | 是 | 物流公司服务电话 |   |
| `gmtCreate` | java.util.Date | 是 | 创建时间 |   |
| `gmtModified` | java.util.Date | 是 | 修改时间 | 20160926142611000+0800 |
| `id` | java.lang.Long | 是 | ID | 352 |
| `spelling` | java.lang.String | 是 | 全拼 |   |
| `supportPrint` | Boolean | 是 | 是否支持打印 |   |

## 错误码

| 错误码 | 现象 | 解决方案 |
|---|---|---|
| 500 | {\&quot;errorCode\&quot;:\&quot;003\&quot;,\&quot;errorMessage\&quot;:\&quot;SERVICE:查询数据失败，请检查输入参数后重试。\&quot;,\&quot;cause\&quot;:\&quot;\&quot;} | 非授权用户的退款单 |
| 500 | {\&quot;errorCode\&quot;:\&quot;003002\&quot;,\&quot;errorMessage\&quot;:\&quot;SERVICE:INVOKE_FAIL:退款数据错误，请检查退款单号，退款单号的格式一般为TD+id\&quot;,\&quot;cause\&quot;:\&quot;errorCode:REFUND_DATA_ERROR,errorMsg:refundId : 11043002311780591,cause:null\&quot;} | 退款单不正确 |
| 500 | {\&quot;errorCode\&quot;:\&quot;003002\&quot;,\&quot;errorMessage\&quot;:\&quot;SERVICE:INVOKE_FAIL:OrderRefundService#queryOrderRefundOperationList\&quot;,\&quot;cause\&quot;:\&quot;errorCode:INVALID_PARAM,errorMsg:null,cause:null\&quot;} | 退款单不正确 |

## 示例

**extInfo参数说明**

```
payMode     资金分流（alipay:支付宝，fundFreeze:资金平台账户冻结） 
 7d          7天无理由订单标记（1表示7天无理由订单） 
```

**请求参数示例**

```
{"refundId":"TQ865149372**61198",
"needTimeOutInfo":true,
"needOrderRefundOperation":true}
```

**返回参数示例**

```
{"result":{"opOrderRefundModelDetail":{"alipayPaymentId":"2018042321001008760555342267","applyCarriage":0,"applyPayment":1,"applyReason":"颜色/图案/款式不符","applyReasonId":20021,"applySubReasonId":0,"buyerLogisticsName":"其他","buyerMemberId":"b2b-1624961198","buyerSendGoods":true,"buyerUserId":1624961198,"canRefundPayment":1,"crmModifyRefund":false,"disputeRequest":3,"disputeType":1,"extInfo":{"EXmrf":"1","ttid":"2","prepaidFailure":"QUERY_BUYER_CREDIT_LEVEL_FAIL","pay_lock":"seller","reason":"20021","apply_text_id":"null","newRefund":"rp2","bizCode":"cbu.general.refund","disputeTradeStatus":"4","lastOrder":"0","seller_agreed_refund_fee":"1","old_reason_id":"20021","sync":"0","b2b_seller_mId":"b2b-1623492085","seller_batch":"true","itemBuyAmount":"0","b2b_buyer_mId":"b2b-1624961198","ability":"1","refundPostFee":"0","seller_audit":"0","ee_trace_id":"0b802ca915244768332374685efa19","ol_tf":"1","opRole":"user","apply_init_refund_fee":"1","isVirtual":"0","logisticsCompanyId":"-1","itemPrice":"0","interceptStatus":"0","refundFrom":"2","restartForXiaoer":"1","appName":"tosp-aftersales","abnormal_dispute_status":"0","payMode":"alipay","workflowName":"cbu_return_and_refund","sgr":"1","enfunddetail":"1","sellerDoRefundNick":"alitestforisv01","bgmtc":"2018-04-23 09:48:12"},"freightBill":"12345134135435234","frozenFund":-1,"gmtApply":"20180423171116000+0800","gmtCompleted":"20180423174713000+0800","gmtCreate":"20180423171116000+0800","gmtModified":"20180423174714000+0800","goodsReceived":true,"goodsStatus":3,"id":8651493722961198,"newRefundReturn":true,"onlyRefund":false,"orderEntryCountMap":{"151267031009969811":2},"orderEntryIdList":[151267031009969811],"orderId":151267031008969811,"productName":"欧洲站2018新款宽松显瘦烫金印花字母老鹰短袖T恤打底裤套装女潮","refundCarriage":0,"refundGoods":true,"refundId":"TQ8651493722961198","refundPayment":1,"rejectReasonId":0,"rejectTimes":0,"sellerDelayDisburse":false,"sellerMemberId":"b2b-1623492085","sellerMobile":"19926555555","sellerRealName":"琳琳","sellerReceiveAddress":"山东 聊城 解决了交流交流链接连接","sellerUserId":1623492085,"status":"refundsuccess","supportNewSteppay":true,"timeOutFreeze":false,"tradeTypeStr":"50060","refundOperationList":[{"afterOperateStatus":"waitsellerreceive","beforeOperateStatus":"waitsellerreceive","closeRefundStepId":0,"crmModifyRefund":false,"discription":"卖家同意退货退款协议，退货退款成功","gmtCreate":"20180423174714000+0800","gmtModified":"20180423174714000+0800","id":50027795018,"messageStatus":3,"msgType":7,"operateRemark":"退款成功","operateTypeInt":8,"operatorLoginId":"alitestforisv01","operatorRoleId":2,"operatorUserId":1623492085,"refundId":"TQ8651493722961198"},{"afterOperateStatus":"waitbuyersend","beforeOperateStatus":"waitbuyersend","closeRefundStepId":0,"crmModifyRefund":false,"discription":"买家声明退货，等待卖家确认。|物流公司：其他|物流单号：12345134135435234|说明：asdada|","gmtCreate":"20180423174651000+0800","gmtModified":"20180423174651000+0800","id":50027571423,"messageStatus":3,"msgType":7,"operateRemark":"买家退货","operateTypeInt":10,"operatorLoginId":"alitestforisv02","operatorRoleId":1,"operatorUserId":1624961198,"refundId":"TQ8651493722961198"},{"afterOperateStatus":"waitselleragree","beforeOperateStatus":"waitselleragree","closeRefundStepId":0,"crmModifyRefund":false,"discription":"卖家确认收货地址：琳琳,19926555555,山东 聊城 解决了交流交流链接连接","gmtCreate":"20180423174544000+0800","gmtModified":"20180423174544000+0800","id":49967844797,"messageStatus":3,"msgType":7,"operateRemark":"确认收货地址","operateTypeInt":19,"operatorLoginId":"alitestforisv01","operatorRoleId":2,"operatorUserId":1623492085,"refundId":"TQ8651493722961198"},{"closeRefundStepId":0,"crmModifyRefund":false,"discription":"退款诉求：退货退款|申请原因：颜色/图案/款式不符|货品情况：已收到货|退款货品金额：0.01|退款运费金额：0.00|退款说明：asdad","gmtCreate":"20180423171116000+0800","gmtModified":"20180423171116000+0800","id":49962356896,"messageStatus":3,"msgType":7,"operateRemark":"买家申请退款协议，等待卖家确认","operateTypeInt":1,"operatorLoginId":"alitestforisv02","operatorRoleId":1,"operatorUserId":1624961198,"refundId":"TQ8651493722961198"}],"buyerLoginId":"alitestforisv02","sellerLoginId":"alitestforisv01"}}}
```
