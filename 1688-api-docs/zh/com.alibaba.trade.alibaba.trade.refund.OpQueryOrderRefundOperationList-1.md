# 退款单操作记录列表（买家视角）

API: `com.alibaba.trade:alibaba.trade.refund.OpQueryOrderRefundOperationList:1` · Category: 退货退款  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.trade.refund.OpQueryOrderRefundOperationList-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.trade.refund.OpQueryOrderRefundOperationList/{appKey}`  
需要授权 (access_token) · 需要签名

该API为买家使用，卖家查询请使用alibaba.trade.refund.OpQueryOrderRefund.sellerView，买方退款操作记录

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `refundId` | String | 是 | 退款单Id | TQ1043162**46961198 |
| `pageNo` | String | 是 | 当前页号 | 1 |
| `pageSize` | String | 是 | 页大小 | 100 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:alibaba.trade.refund.OpQueryOrderRefundOperationListResult](#m-alibaba-trade-refund-opqueryorderrefundoperationlistresult) | 是 | 返回结果 | {} |
| `errorMessage` | String | 是 | 错误信息 | 退款单错误 |
| `extErrorMessage` | String | 是 | 附加错误信息 | 退款单错误 |
| `errorCode` | String | 是 | 错误码 | 500 |

<a id="m-alibaba-trade-refund-opqueryorderrefundoperationlistresult"></a>
#### alibaba.trade.refund.OpQueryOrderRefundOperationListResult

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `opOrderRefundOperationModels` | [message:alibaba.trade.refund.OpOrderRefundOperationModel[]](#m-alibaba-trade-refund-oporderrefundoperationmodel[]) | 是 | 退款信息 | [] |

<a id="m-alibaba-trade-refund-oporderrefundoperationmodel[]"></a>
#### alibaba.trade.refund.OpOrderRefundOperationModel[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `afterOperateStatus` | java.lang.String | 是 | 操作后的退款状态 | waitselleragree |
| `beforeOperateStatus` | java.lang.String | 是 | 操作前的退款状态 | waitselleragree |
| `closeRefundStepId` | long | 是 | 分阶段订单正向操作关闭退款时的阶段ID | 0 |
| `crmModifyRefund` | Boolean | 是 | 是否小二修改过退款单 |   |
| `discription` | java.lang.String | 是 | 描述、说明 | 卖家同意仅退款协议，仅退款成功 |
| `freightBill` | java.lang.String | 是 | 运单号 |   |
| `gmtCreate` | java.util.Date | 是 | 创建时间 | 20180611180901000+0800 |
| `gmtModified` | java.util.Date | 是 | 修改时间 | 20180611180901000+0800 |
| `id` | java.lang.Long | 是 | 主键，退款操作记录流水号 | 60556968242 |
| `messageStatus` | int | 是 | 凭证状态，1:正常 2:后台小二屏蔽 | 3 |
| `msgType` | int | 是 | 留言类型 3:小二留言给买家和卖家 4:给买家的留言 5:给卖家的留言 7:cbu的普通留言等同于淘宝的1 | 7 |
| `operateRemark` | java.lang.String | 是 | 操作备注 | 退款成功 |
| `operateTypeInt` | int | 是 | 操作类型 取代operateType | 19  |
| `operatorId` | java.lang.String | 是 | 操作者-memberID |   |
| `operatorLoginId` | java.lang.String | 是 | 操作者-loginID | alitestforisv02 |
| `operatorRoleId` | java.lang.Integer | 是 | 操作者角色名称 买家 卖家 系统 | 2 |
| `refundAddress` | java.lang.String | 是 | 退货地址 |   |
| `refundId` | java.lang.String | 是 | 退款记录ID | TQ10143867704492085 |
| `rejectReason` | java.lang.String | 是 | 卖家拒绝退款原因 |   |
| `vouchers` | java.util.List | 是 | 凭证图片地址 |   |
| `logisticsCompany` | [message:alibaba.trade.refund.OpLogisticsCompanyModel](#m-alibaba-trade-refund-oplogisticscompanymodel) | 是 | 物流公司详情 | 暂时不返回 |

<a id="m-alibaba-trade-refund-oplogisticscompanymodel"></a>
#### alibaba.trade.refund.OpLogisticsCompanyModel

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `companyName` | java.lang.String | 是 | 快递公司名 |   |
| `companyNo` | java.lang.String | 是 | 物流公司编号 |   |
| `companyPhone` | java.lang.String | 是 | 物流公司服务电话 |   |
| `gmtCreate` | java.util.Date | 是 | 创建时间 |   |
| `gmtModified` | java.util.Date | 是 | 修改时间 |   |
| `id` | java.lang.Long | 是 | ID |   |
| `spelling` | java.lang.String | 是 | 全拼 |   |
| `supportPrint` | Boolean | 是 | 是否支持打印 |   |

## 错误码

| 错误码 | 现象 | 解决方案 |
|---|---|---|
| 500 | {\&quot;errorCode\&quot;:\&quot;003\&quot;,\&quot;errorMessage\&quot;:\&quot;SERVICE:查询数据失败，请检查输入参数后重试。\&quot;,\&quot;cause\&quot;:\&quot;\&quot;} | 非授权用户的退款单 |
| 500 | {\&quot;errorCode\&quot;:\&quot;003002\&quot;,\&quot;errorMessage\&quot;:\&quot;SERVICE:INVOKE_FAIL:退款数据错误，请检查退款单号，退款单号的格式一般为TD+id\&quot;,\&quot;cause\&quot;:\&quot;errorCode:REFUND_DATA_ERROR,errorMsg:refundId : 11043002311780591,cause:null\&quot;} | 退款单不正确 |
| 500 | {\&quot;errorCode\&quot;:\&quot;003002\&quot;,\&quot;errorMessage\&quot;:\&quot;SERVICE:INVOKE_FAIL:OrderRefundService#queryOrderRefundOperationList\&quot;,\&quot;cause\&quot;:\&quot;errorCode:INVALID_PARAM,errorMsg:null,cause:null\&quot;} | 退款单不正确 |

## 示例

**请求参数示例**

```
{
    "refundId":"TQ10143867704492085",
    "pageNo":1,
    "pageSize":10
}
```

**返回参数示例**

```
{
    "result":{
        "opOrderRefundOperationModels":[
            {
                "afterOperateStatus":"waitselleragree",
                "beforeOperateStatus":"waitselleragree",
                "closeRefundStepId":0,
                "discription":"卖家同意仅退款协议，仅退款成功",
                "gmtCreate":"20180611180901000+0800",
                "gmtModified":"20180611180901000+0800",
                "id":60556968242,
                "messageStatus":3,
                "msgType":7,
                "operateRemark":"退款成功",
                "operateTypeInt":19,
                "operatorLoginId":"alitestforisv02",
                "operatorRoleId":2,
                "operatorUserId":1624961198,
                "refundId":"TQ10143867704492085"
            },
            {
                "closeRefundStepId":0,
                "discription":"退款诉求：仅退款|申请原因：不想买了/等不及/拍错|货品情况：未收到货|退款货品金额：0.01|退款运费金额：0.01|退款说明：Ian testing for buyer refund",
                "gmtCreate":"20180611180414000+0800",
                "gmtModified":"20180611180414000+0800",
                "id":60660515733,
                "messageStatus":3,
                "msgType":7,
                "operateRemark":"买家申请退款协议，等待卖家确认",
                "operateTypeInt":1,
                "operatorLoginId":"alitestforisv01",
                "operatorRoleId":1,
                "operatorUserId":1623492085,
                "refundId":"TQ10143867704492085"
            }
        ]
    }
}
```
