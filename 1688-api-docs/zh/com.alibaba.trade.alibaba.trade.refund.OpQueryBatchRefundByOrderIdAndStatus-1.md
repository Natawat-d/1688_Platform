# 查询退款单详情-根据订单ID（买家视角）

API: `com.alibaba.trade:alibaba.trade.refund.OpQueryBatchRefundByOrderIdAndStatus:1` · Category: 退货退款  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.trade.refund.OpQueryBatchRefundByOrderIdAndStatus-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.trade.refund.OpQueryBatchRefundByOrderIdAndStatus/{appKey}`  
需要授权 (access_token) · 需要签名

该API为买家使用，卖家查询请使用alibaba.trade.refund.OpQueryBatchRefundByOrderIdAndStatus.sellerView，根据订单号实时查询退款单列表，目前只能查询到售中的退款单

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `orderId` | String | 是 | 订单id | 151267031**8969811 |
| `queryType` | String | 是 | 1：活动；3:退款成功（只支持退款中和退款成功） | 3 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:alibaba.trade.refund.OpQueryBatchRefundByOrderIdAndStatusResultModel](#m-alibaba-trade-refund-opquerybatchrefundbyorderidandstatusresultmodel) | 是 | 查询结果 | {} |
| `errorCode` | String | 是 | 错误码 | 500 |
| `errorMessage` | String | 是 | 错误信息 | "{\"errorCode\":\"003002\",\"errorMessage\":\"SERVICE:INVOKE_FAIL:OrderRefundService.queryBatchRefundByOrderIdAndStatus\",\"cause\":\"errorCode:ORDER_NOT_EXIST,errorMsg:null,cause:null\"}" |
| `extErrorMessage` | String | 是 | 附加信息 | 订单号有误 |

<a id="m-alibaba-trade-refund-opquerybatchrefundbyorderidandstatusresultmodel"></a>
#### alibaba.trade.refund.OpQueryBatchRefundByOrderIdAndStatusResultModel

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `opOrderRefundModels` | [message:alibaba.trade.refund.OpOrderRefundModel[]](#m-alibaba-trade-refund-oporderrefundmodel[]) | 是 | 退款单信息 | {} |

<a id="m-alibaba-trade-refund-oporderrefundmodel[]"></a>
#### alibaba.trade.refund.OpOrderRefundModel[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `alipayPaymentId` | java.lang.String | 是 | 支付宝交易号 | 2018042321****08760555342267 |
| `applyCarriage` | java.lang.Long | 是 | 运费的申请退款金额，单位：分 | 0 |
| `applyExpect` | java.lang.Long | 是 | 买家原始输入的退款金额(可以为空) |   |
| `applyPayment` | java.lang.Long | 是 | 买家申请退款金额，单位：分 | 1 |
| `applyReason` | java.lang.String | 是 | 申请原因 | 颜色/图案/款式不符 |
| `applyReasonId` | int | 是 | 申请原因ID | 20021 |
| `applySubReason` | java.lang.String | 是 | 二级退款原因 |   |
| `asynErrCode` | java.lang.String | 是 |  |   |
| `applySubReasonId` | int | 是 | 二级退款原因Id | -1 |
| `asynSubErrCode` | java.lang.String | 是 |  |   |
| `buyerAlipayId` | java.lang.String | 是 | 买家支付宝ID |   |
| `buyerLogisticsName` | java.lang.String | 是 | 买家退货物流公司名 | 其他 |
| `buyerMemberId` | java.lang.String | 是 | 买家会员ID | b2b-1624961198 |
| `buyerUserId` | java.lang.Long | 是 | 买家阿里帐号ID(包括淘宝帐号Id) | 1624961198 |
| `canRefundPayment` | java.lang.Long | 是 | 最大能够退款金额，单位：分 | 1 |
| `disburseChannel` | java.lang.String | 是 | 极速到账打款渠道 |   |
| `disputeRequest` | int | 是 | 售后退款要求 | 3 |
| `disputeType` | int | 是 | 纠纷类型：售中退款 售后退款，默认为售中退款 | 1 |
| `extInfo` | java.util.Map | 是 | 扩展信息 | {} |
| `freightBill` | java.lang.String | 是 | 运单号 | 12345134135435234 |
| `frozenFund` | long | 是 | 实际冻结账户金额,单位：分 | -1 |
| `gmtApply` | java.util.Date | 是 | 申请退款时间 | 20180423171116000+0800 |
| `gmtCompleted` | java.util.Date | 是 | 完成时间 | 20180423174713000+0800 |
| `gmtCreate` | java.util.Date | 是 | 创建时间 | 20180423171116000+0800 |
| `gmtFreezed` | java.util.Date | 是 | 该退款单超时冻结开始时间 |   |
| `gmtModified` | java.util.Date | 是 | 修改时间 | 20180423174714000+0800 |
| `gmtTimeOut` | java.util.Date | 是 | 该退款单超时完成的时间期限 |   |
| `goodsStatus` | int | 是 | 1：买家未收到货<br>2：买家已收到货<br>3：买家已退货 | 3 |
| `id` | java.lang.Long | 是 | 退款单编号 | 8651493722961198 |
| `instantRefundType` | java.lang.String | 是 | 极速到账退款类型 |   |
| `orderEntryCountMap` | java.util.Map | 是 | 子订单退货数量 |   |
| `orderEntryIdList` | java.util.List | 是 | 退款单包含的订单明细，时间逆序排列 |   |
| `orderId` | java.lang.Long | 是 | 退款单对应的订单编号 | 151267031008969811 |
| `prepaidBalance` | java.lang.Long | 是 | 极速退款垫资金额,该值不为空时,只代表该退款单可以走垫资流程,但不代表一定垫资成功 |   |
| `productName` | java.lang.String | 是 | 产品名称(退款单关联订单明细的货品名称) | 欧洲站2018新款宽松显瘦烫金印花字母老鹰短袖T恤打底裤套装女潮 |
| `refundCarriage` | java.lang.Long | 是 | 运费的实际退款金额，单位：分 | 0 |
| `refundId` | java.lang.String | 是 | 退款单逻辑主键 | TQ8651493722961198 |
| `refundPayment` | java.lang.Long | 是 | 实际退款金额，单位：分 | 1 |
| `rejectReason` | java.lang.String | 是 | 卖家拒绝原因 |   |
| `rejectReasonId` | int | 是 | 卖家拒绝原因Id | 0 |
| `rejectTimes` | int | 是 | 退款单被拒绝的次数 | 0 |
| `sellerAlipayId` | java.lang.String | 是 | 卖家支付宝ID |   |
| `sellerMemberId` | java.lang.String | 是 | 卖家会员ID | b2b-1623492085 |
| `sellerMobile` | java.lang.String | 是 | 收货人手机 |   |
| `sellerRealName` | java.lang.String | 是 | 收货人姓名 |   |
| `sellerReceiveAddress` | java.lang.String | 是 | 买家退货时卖家收货地址 |   |
| `sellerTel` | java.lang.String | 是 | 收货人电话 |   |
| `sellerUserId` | java.lang.Long | 是 | 卖家阿里帐号ID(包括淘宝帐号Id) | 1623492085 |
| `status` | java.lang.String | 是 | 退款状态 | refundsuccess |
| `taskStatus` | java.lang.String | 是 | 工单子状态，没有流到CRM创建工单时为空 |   |
| `timeOutOperateType` | java.lang.String | 是 | 超时后执行的动作 |   |
| `tradeTypeStr` | java.lang.String | 是 | 交易类型，用来替换枚举类型的tradeType | 50060 |
| `success` | Boolean | 是 | 是否成功 | true |
| `refundOperationList` | [message:alibaba.trade.refund.OpOrderRefundOperationModel[]](#m-alibaba-trade-refund-oporderrefundoperationmodel[]) | 是 | 操作记录列表 | 暂不返回 |
| `buyerLoginId` | String | 是 | 买家会员ID | alitestforisv02 |
| `sellerLoginId` | String | 是 | 卖家会员ID | alitestforisv01 |
| `isCrmModifyRefund` | Boolean | 是 | 是否小二修改过退款单 | false |
| `isTimeOutFreeze` | Boolean | 是 | 是否超时系统冻结，true代表冻结，false代表不冻结 | false |
| `isInsufficientAccount` | Boolean | 是 | 交易4.0退款余额不足 | false |
| `isGoodsReceived` | Boolean | 是 | 买家是否已收到货 | true |
| `isOnlyRefund` | Boolean | 是 | 是否仅退款 | false |
| `isRefundGoods` | Boolean | 是 | 是否要求退货 | true |
| `isSellerDelayDisburse` | Boolean | 是 | 是否卖家延迟打款，即安全退款 | false |
| `isAftersaleAutoDisburse` | Boolean | 是 | 售后自动打款 | false |
| `isSupportNewSteppay` | Boolean | 是 | 是否支持交易4.0 | true |
| `isNewRefundReturn` | Boolean | 是 | 是否新流程创建的退款退货 | true |
| `isBuyerSendGoods` | Boolean | 是 | 买家是否已经发货（如果有退货的流程） | true |
| `isAftersaleAgreeTimeout` | Boolean | 是 | 售后超时标记 | false |
| `isInsufficientBail` | Boolean | 是 | 极速到账退款保证金不足 | false |
| `refundOfficialSolutionCost` | Long | 是 | 退官方物流提货订单费用（分） | 770 |

<a id="m-alibaba-trade-refund-oporderrefundoperationmodel[]"></a>
#### alibaba.trade.refund.OpOrderRefundOperationModel[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `afterOperateStatus` | java.lang.String | 是 | 操作后的退款状态 |  |
| `beforeOperateStatus` | java.lang.String | 是 | 操作前的退款状态 |  |
| `closeRefundStepId` | long | 是 | 分阶段订单正向操作关闭退款时的阶段ID |  |
| `crmModifyRefund` | Boolean | 是 | 是否小二修改过退款单 |  |
| `discription` | java.lang.String | 是 | 描述、说明 |  |
| `email` | java.lang.String | 是 | 联系人EMAIL |  |
| `freightBill` | java.lang.String | 是 | 运单号 |  |
| `gmtCreate` | java.util.Date | 是 | 创建时间 |  |
| `gmtModified` | java.util.Date | 是 | 修改时间 |  |
| `id` | java.lang.Long | 是 | 主键，退款操作记录流水号 |  |
| `messageStatus` | int | 是 | 凭证状态，1:正常 2:后台小二屏蔽 |  |
| `mobile` | java.lang.String | 是 | 联系人手机 |  |
| `msgType` | int | 是 | 留言类型 3:小二留言给买家和卖家 4:给买家的留言 5:给卖家的留言 7:cbu的普通留言等同于淘宝的1 |  |
| `operateRemark` | java.lang.String | 是 | 操作备注 |  |
| `operateTypeInt` | int | 是 | 操作类型 取代operateType |  |
| `operatorId` | java.lang.String | 是 | 操作者-memberID |  |
| `operatorLoginId` | java.lang.String | 是 | 操作者-loginID |  |
| `operatorRoleId` | java.lang.Integer | 是 | 操作者角色名称 买家 卖家 系统 |  |
| `operatorUserId` | java.lang.Long | 是 | 操作者-userID |  |
| `phone` | java.lang.String | 是 | 联系人电话 |  |
| `refundAddress` | java.lang.String | 是 | 退货地址 |  |
| `refundId` | java.lang.String | 是 | 退款记录ID |  |
| `rejectReason` | java.lang.String | 是 | 卖家拒绝退款原因 |  |
| `vouchers` | java.util.List | 是 | 凭证图片地址 |  |
| `logisticsCompany` | [message:alibaba.trade.refund.OpLogisticsCompanyModel](#m-alibaba-trade-refund-oplogisticscompanymodel) | 是 | 物流公司详情 |  |

<a id="m-alibaba-trade-refund-oplogisticscompanymodel"></a>
#### alibaba.trade.refund.OpLogisticsCompanyModel

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `companyName` | java.lang.String | 是 | 快递公司名 |  |
| `companyNo` | java.lang.String | 是 | 物流公司编号 |  |
| `companyPhone` | java.lang.String | 是 | 物流公司服务电话 |  |
| `gmtCreate` | java.util.Date | 是 | 创建时间 |  |
| `gmtModified` | java.util.Date | 是 | 修改时间 |  |
| `id` | java.lang.Long | 是 | ID |  |
| `spelling` | java.lang.String | 是 | 全拼 |  |
| `supportPrint` | Boolean | 是 | 是否支持打印 |  |

## 错误码

| 错误码 | 现象 | 解决方案 |
|---|---|---|
| 500 | &quot;{\&quot;errorCode\&quot;:\&quot;003002\&quot;,\&quot;errorMessage\&quot;:\&quot;SERVICE:INVOKE_FAIL:OrderRefundService.queryBatchRefundByOrderIdAndStatus\&quot;,\&quot;cause\&quot;:\&quot;errorCode:ORDER_NOT_EXIST,errorMsg:null,cause:null\&quot;}&quot; | 订单号有误 |

## 示例

**extInfo参数说明**

```
payMode     资金分流（alipay:支付宝，fundFreeze:资金平台账户冻结） 
 7d          7天无理由订单标记（1表示7天无理由订单）  
```

**请求参数示例**

```
{"orderId":"151267031008969811",
"queryType" :"3",
"dipsuteType":1}
```

**返回参数示例**

```
{"result":{"opOrderRefundModels":[{"applyReasonId":20021,"applySubReasonId":-1,"buyerUserId":1624961198,"alipayPaymentId":"2018042321001008760555342267","isCrmModifyRefund":false,"rejectTimes":0,"isTimeOutFreeze":false,"applyReason":"颜色/图案/款式不符","isInsufficientAccount":false,"rejectReasonId":0,"buyerLoginId":"alitestforisv02","refundCarriage":0,"isGoodsReceived":true,"sellerUserId":1623492085,"isOnlyRefund":false,"status":"refundsuccess","isRefundGoods":true,"refundOperationList":[],"isSellerDelayDisburse":false,"sellerMemberId":"b2b-1623492085","sellerLoginId":"alitestforisv01","gmtModified":"20180423174714000+0800","isAftersaleAutoDisburse":false,"disputeRequest":3,"disputeType":1,"id":8651493722961198,"isSupportNewSteppay":true,"canRefundPayment":1,"tradeTypeStr":"50060","freightBill":"12345134135435234","gmtApply":"20180423171116000+0800","goodsStatus":3,"gmtCompleted":"20180423174713000+0800","applyPayment":1,"buyerLogisticsName":"其他","orderId":151267031008969811,"isNewRefundReturn":true,"applyCarriage":0,"gmtCreate":"20180423171116000+0800","isBuyerSendGoods":true,"buyerMemberId":"b2b-1624961198","frozenFund":-1,"refundId":"TQ8651493722961198","extInfo":{"EXmrf":"1","ttid":"2","prepaidFailure":"QUERY_BUYER_CREDIT_LEVEL_FAIL","pay_lock":"seller","reason":"20021","apply_text_id":"null","newRefund":"rp2","bizCode":"cbu.general.refund","disputeTradeStatus":"4","lastOrder":"0","seller_agreed_refund_fee":"1","old_reason_id":"20021","sync":"0","b2b_seller_mId":"b2b-1623492085","seller_batch":"true","itemBuyAmount":"0","b2b_buyer_mId":"b2b-1624961198","ability":"1","refundPostFee":"0","seller_audit":"0","ee_trace_id":"0b802ca915244768332374685efa19","ol_tf":"1","opRole":"user","apply_init_refund_fee":"1","isVirtual":"0","logisticsCompanyId":"-1","itemPrice":"0","interceptStatus":"0","refundFrom":"2","restartForXiaoer":"1","appName":"tosp-aftersales","abnormal_dispute_status":"0","payMode":"alipay","workflowName":"cbu_return_and_refund","sgr":"1","enfunddetail":"1","sellerDoRefundNick":"alitestforisv01","bgmtc":"2018-04-23 09:48:12"},"refundPayment":1,"productName":"欧洲站2018新款宽松显瘦烫金印花字母老鹰短袖T恤打底裤套装女潮 黑 XL","success":true,"isAftersaleAgreeTimeout":false,"isInsufficientBail":false}]}}
```
