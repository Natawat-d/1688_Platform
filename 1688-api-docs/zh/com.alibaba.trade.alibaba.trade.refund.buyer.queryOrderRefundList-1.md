# 查询退款单列表(买家视角)

API: `com.alibaba.trade:alibaba.trade.refund.buyer.queryOrderRefundList:1` · Category: 退货退款  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.trade.refund.buyer.queryOrderRefundList-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.trade.refund.buyer.queryOrderRefundList/{appKey}`  
需要授权 (access_token) · 需要签名

买家查看退款单列表，该接口不支持子账号查询，请使用主账号授权后查询

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `orderId` | Long | 否 | 订单Id | 179087886005498520 |
| `applyStartTime` | java.util.Date | 否 | 退款申请时间（起始） | 20170926114526000+0800 |
| `applyEndTime` | java.util.Date | 否 | 退款申请时间（截止） | 20220926114526000+0800 |
| `refundStatusSet` | String[] | 否 | 退款状态列表 | 等待卖家同意 waitselleragree;退款成功 refundsuccess;退款关闭 refundclose;待买家修改 waitbuyermodify;等待买家退货 waitbuyersend;等待卖家确认收货 waitsellerreceive |
| `sellerMemberId` | String | 否 | 卖家memberId | b2b-1623492085 |
| `currentPageNum` | Integer | 否 | 当前页码 | 0 |
| `pageSize` | Integer | 否 | 每页条数 | 20 |
| `logisticsNo` | String | 否 | 退货物流单号（传此字段查询时，需同时传入sellerMemberId） | 3101***159271 |
| `modifyStartTime` | java.util.Date | 否 | 退款修改时间(起始) | 20170926114526000+0800 |
| `modifyEndTime` | java.util.Date | 否 | 退款修改时间(截止) | 20220926114526000+0800 |
| `dipsuteType` | Integer | 否 | 1:售中退款，2:售后退款；0:所有退款单 | 1 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:alibaba.trade.refund.OpQueryOrderRefundListResult](#m-alibaba-trade-refund-opqueryorderrefundlistresult) | 是 | 查询结果 | {} |
| `errorCode` | String | 是 | 错误码 |   |
| `errorMsg` | String | 是 | 错误信息 |   |
| `success` | Boolean | 是 | 是否成功 | true |

<a id="m-alibaba-trade-refund-opqueryorderrefundlistresult"></a>
#### alibaba.trade.refund.OpQueryOrderRefundListResult

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `opOrderRefundModels` | [message:alibaba.trade.refund.OpOrderRefundModel[]](#m-alibaba-trade-refund-oporderrefundmodel[]) | 是 | 退款单列表 |  |
| `totalCount` | int | 是 | 符合条件总的记录条数 |  |
| `currentPageNum` | int | 是 | 查询的当前页码 |  |

<a id="m-alibaba-trade-refund-oporderrefundmodel[]"></a>
#### alibaba.trade.refund.OpOrderRefundModel[]

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
| `asynErrCode` | java.lang.String | 是 |  |   |
| `asynSubErrCode` | java.lang.String | 是 |  |   |
| `buyerAlipayId` | java.lang.String | 是 | 买家支付宝ID |   |
| `buyerLogisticsName` | java.lang.String | 是 | 买家退货物流公司名 |   |
| `buyerMemberId` | java.lang.String | 是 | 买家会员ID |   |
| `buyerSendGoods` | Boolean | 是 | 买家是否已经发货（如果有退货的流程） |   |
| `buyerUserId` | java.lang.Long | 是 | 买家阿里帐号ID(包括淘宝帐号Id) |   |
| `canRefundPayment` | java.lang.Long | 是 | 最大能够退款金额，单位：分 |   |
| `crmModifyRefund` | Boolean | 是 | 是否小二修改过退款单 |   |
| `disburseChannel` | java.lang.String | 是 | 极速到账打款渠道 |   |
| `disputeRequest` | int | 是 | 售后退款要求 |   |
| `disputeType` | int | 是 | 纠纷类型：售中退款 售后退款，默认为售中退款 |   |
| `extInfo` | java.util.Map | 是 | 扩展信息 |   |
| `freightBill` | java.lang.String | 是 | 运单号 |   |
| `frozenFund` | long | 是 | 实际冻结账户金额,单位：分 |   |
| `gmtApply` | java.util.Date | 是 | 申请退款时间 |   |
| `gmtCompleted` | java.util.Date | 是 | 完成时间 |   |
| `gmtCreate` | java.util.Date | 是 | 创建时间 |   |
| `gmtFreezed` | java.util.Date | 是 | 该退款单超时冻结开始时间 |   |
| `gmtModified` | java.util.Date | 是 | 修改时间 |   |
| `gmtTimeOut` | java.util.Date | 是 | 该退款单超时完成的时间期限 |   |
| `goodsReceived` | Boolean | 是 | 买家是否已收到货 |   |
| `goodsStatus` | int | 是 | 1：买家未收到货<br>2：买家已收到货<br>3：买家已退货 |   |
| `id` | java.lang.Long | 是 | 退款单编号 |   |
| `instantRefundType` | java.lang.String | 是 | 极速到账退款类型 |   |
| `insufficientAccount` | Boolean | 是 | 交易4.0退款余额不足 |   |
| `insufficientBail` | Boolean | 是 | 极速到账退款保证金不足 |   |
| `newRefundReturn` | Boolean | 是 | 是否新流程创建的退款退货 |   |
| `onlyRefund` | Boolean | 是 | 是否仅退款 |   |
| `orderEntryCountMap` | java.util.Map | 是 | 子订单退货数量 |   |
| `orderEntryIdList` | java.util.List | 是 | 退款单包含的订单明细，时间逆序排列 |   |
| `orderId` | java.lang.Long | 是 | 退款单对应的订单编号 |   |
| `prepaidBalance` | java.lang.Long | 是 | 极速退款垫资金额,该值不为空时,只代表该退款单可以走垫资流程,但不代表一定垫资成功 |   |
| `productName` | java.lang.String | 是 | 产品名称(退款单关联订单明细的货品名称) |   |
| `refundCarriage` | java.lang.Long | 是 | 运费的实际退款金额，单位：分 |   |
| `refundGoods` | Boolean | 是 | 是否要求退货 |   |
| `refundId` | java.lang.String | 是 | 退款单逻辑主键 |   |
| `refundPayment` | java.lang.Long | 是 | 实际退款金额，单位：分 |   |
| `rejectReason` | java.lang.String | 是 | 卖家拒绝原因 |   |
| `rejectReasonId` | int | 是 | 卖家拒绝原因Id |   |
| `rejectTimes` | int | 是 | 退款单被拒绝的次数 |   |
| `sellerAlipayId` | java.lang.String | 是 | 卖家支付宝ID |   |
| `sellerDelayDisburse` | Boolean | 是 | 是否卖家延迟打款，即安全退款 |   |
| `sellerMemberId` | java.lang.String | 是 | 卖家会员ID |   |
| `sellerMobile` | java.lang.String | 是 | 收货人手机 |   |
| `sellerRealName` | java.lang.String | 是 | 收货人姓名 |   |
| `sellerReceiveAddress` | java.lang.String | 是 | 买家退货时卖家收货地址 |   |
| `sellerTel` | java.lang.String | 是 | 收货人电话 |   |
| `sellerUserId` | java.lang.Long | 是 | 卖家阿里帐号ID(包括淘宝帐号Id) |   |
| `status` | java.lang.String | 是 | 退款状态 |   |
| `supportNewSteppay` | Boolean | 是 | 是否支持交易4.0 |   |
| `taskStatus` | java.lang.String | 是 | 工单子状态，没有流到CRM创建工单时为空 |   |
| `timeOutFreeze` | Boolean | 是 | 是否超时系统冻结，true代表冻结，false代表不冻结 |   |
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
| APPLY_TIME_ERROR | 只能查询近730天的退款单，请检查退款申请时间! | 请检查退款申请时间! |

## 示例

**extInfo参数说明**

```
payMode     资金分流（alipay:支付宝，fundFreeze:资金平台账户冻结） 
7d          7天无理由订单标记（1表示7天无理由订单）     
```

**请求参数示例**

```
{"orderId":"177681528398969811",
"dipsuteType":0}
```

**返回参数示例**

```
{"result":{"opOrderRefundModels":[{"alipayPaymentId":"2018061421001008760569923894","applyCarriage":0,"applyPayment":1,"applyReason":"不想买了，已与卖家协商一致","applyReasonId":20028,"applySubReasonId":-1,"buyerMemberId":"b2b-1624961198","buyerUserId":1624961198,"canRefundPayment":1,"disputeRequest":3,"disputeType":1,"extInfo":{"EXmrf":"1","ttid":"2","reason":"20028","apply_text_id":"null","newRefund":"rp2","bizCode":"cbu.general.refund","lastOrder":"0","seller_agreed_refund_fee":"1","old_reason_id":"20028","b2b_seller_mId":"b2b-1623492085","seller_batch":"true","itemBuyAmount":"0","b2b_buyer_mId":"b2b-1624961198","ability":"1","seller_audit":"0","ee_trace_id":"0ab2dbd215300798520175927d07d3","ol_tf":"0","opRole":"timeout","apply_init_refund_fee":"1","isVirtual":"0","logisticsCompanyId":"-1","itemPrice":"0","interceptStatus":"0","refundFrom":"2","restartForXiaoer":"1","appName":"tosp-aftersales","abnormal_dispute_status":"0","payMode":"alipay","workflowName":"cbu_return_and_refund","sgr":"1","enfunddetail":"1","bgmtc":"2018-06-14 14:20:21"},"frozenFund":-1,"gmtApply":"20180622141034000+0800","gmtCompleted":"20180702141113000+0800","gmtCreate":"20180622141034000+0800","gmtModified":"20180702141113000+0800","goodsStatus":2,"id":10426553019961198,"orderId":177681528398969811,"productName":"短袖撞色领保罗衫定做纯棉广告衫定制印字logo刺绣男现货批发 白色/橙领 S等3种","refundCarriage":0,"refundId":"TQ10426553019961198","refundPayment":1,"rejectReasonId":0,"rejectTimes":0,"sellerMemberId":"b2b-1623492085","sellerUserId":1623492085,"status":"refundclose","tradeTypeStr":"50060","refundOperationList":[]}],"totalCount":1,"currentPageNum":0},"success":true}
```
