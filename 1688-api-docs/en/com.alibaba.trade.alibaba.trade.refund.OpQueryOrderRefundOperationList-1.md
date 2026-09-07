# Refund operation history (buyer view)

Original name: 退款单操作记录列表（买家视角）  
API: `com.alibaba.trade:alibaba.trade.refund.OpQueryOrderRefundOperationList:1` · Category: Returns & Refunds  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.trade.refund.OpQueryOrderRefundOperationList-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.trade.refund.OpQueryOrderRefundOperationList/{appKey}`  
Requires user authorization (access_token) · Requires signature

For buyers; sellers should use alibaba.trade.refund.OpQueryOrderRefund.sellerView. Returns the buyer-side refund operation records.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `refundId` | String | yes | Refund order ID | TQ1043162**46961198 |
| `pageNo` | String | yes | Current page number | 1 |
| `pageSize` | String | yes | Page size | 100 |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.trade.refund.OpQueryOrderRefundOperationListResult](#m-alibaba-trade-refund-opqueryorderrefundoperationlistresult) | yes | Return result | {} |
| `errorMessage` | String | yes | Error message | 退款单错误 |
| `extErrorMessage` | String | yes | Additional error information | 退款单错误 |
| `errorCode` | String | yes | Error code | 500 |

<a id="m-alibaba-trade-refund-opqueryorderrefundoperationlistresult"></a>
#### alibaba.trade.refund.OpQueryOrderRefundOperationListResult

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `opOrderRefundOperationModels` | [message:alibaba.trade.refund.OpOrderRefundOperationModel[]](#m-alibaba-trade-refund-oporderrefundoperationmodel[]) | yes | Refund info | [] |

<a id="m-alibaba-trade-refund-oporderrefundoperationmodel[]"></a>
#### alibaba.trade.refund.OpOrderRefundOperationModel[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `afterOperateStatus` | java.lang.String | yes | Refund status after the operation | waitselleragree |
| `beforeOperateStatus` | java.lang.String | yes | Refund status before the operation | waitselleragree |
| `closeRefundStepId` | long | yes | Stage ID when a staged order's forward operation closes the refund | 0 |
| `crmModifyRefund` | Boolean | yes | Whether the refund order has been modified by customer service (Xiaoer) |   |
| `discription` | java.lang.String | yes | Description, notes | 卖家同意仅退款协议，仅退款成功 |
| `freightBill` | java.lang.String | yes | Waybill number |   |
| `gmtCreate` | java.util.Date | yes | Creation time | 20180611180901000+0800 |
| `gmtModified` | java.util.Date | yes | Modification time | 20180611180901000+0800 |
| `id` | java.lang.Long | yes | Primary key, refund operation record serial number | 60556968242 |
| `messageStatus` | int | yes | Voucher status. 1: normal; 2: blocked by backend staff | 3 |
| `msgType` | int | yes | Message type. 3: staff message to buyer and seller; 4: message to buyer; 5: message to seller; 7: regular CBU message, equivalent to Taobao's 1 | 7 |
| `operateRemark` | java.lang.String | yes | Operation remarks | 退款成功 |
| `operateTypeInt` | int | yes | Operation type, replaces operateType | 19  |
| `operatorId` | java.lang.String | yes | Operator - memberID |   |
| `operatorLoginId` | java.lang.String | yes | Operator - loginID | alitestforisv02 |
| `operatorRoleId` | java.lang.Integer | yes | Operator role name: buyer, seller, system | 2 |
| `refundAddress` | java.lang.String | yes | Return address |   |
| `refundId` | java.lang.String | yes | Refund record ID | TQ10143867704492085 |
| `rejectReason` | java.lang.String | yes | Reason for seller's refund rejection |   |
| `vouchers` | java.util.List | yes | Voucher image URL |   |
| `logisticsCompany` | [message:alibaba.trade.refund.OpLogisticsCompanyModel](#m-alibaba-trade-refund-oplogisticscompanymodel) | yes | Logistics company details | 暂时不返回 |

<a id="m-alibaba-trade-refund-oplogisticscompanymodel"></a>
#### alibaba.trade.refund.OpLogisticsCompanyModel

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `companyName` | java.lang.String | yes | Courier company name |   |
| `companyNo` | java.lang.String | yes | Logistics company number |   |
| `companyPhone` | java.lang.String | yes | Logistics company service phone number |   |
| `gmtCreate` | java.util.Date | yes | Creation time |   |
| `gmtModified` | java.util.Date | yes | Modification time |   |
| `id` | java.lang.Long | yes | ID |   |
| `spelling` | java.lang.String | yes | Full pinyin |   |
| `supportPrint` | Boolean | yes | Whether printing is supported |   |

## Error codes

| Error | Symptom | How to fix |
|---|---|---|
| 500 | {\&quot;errorCode\&quot;:\&quot;003\&quot;,\&quot;errorMessage\&quot;:\&quot;SERVICE: failed to query data, please check the input parameters and try again.\&quot;,\&quot;cause\&quot;:\&quot;\&quot;} | Refund order of an unauthorized user |
| 500 | {\&quot;errorCode\&quot;:\&quot;003002\&quot;,\&quot;errorMessage\&quot;:\&quot;SERVICE:INVOKE_FAIL:Refund data error, please check the refund order number; the refund order number format is usually TD+id\&quot;,\&quot;cause\&quot;:\&quot;errorCode:REFUND_DATA_ERROR,errorMsg:refundId : 11043002311780591,cause:null\&quot;} | Refund order is incorrect |
| 500 | {\&quot;errorCode\&quot;:\&quot;003002\&quot;,\&quot;errorMessage\&quot;:\&quot;SERVICE:INVOKE_FAIL:OrderRefundService#queryOrderRefundOperationList\&quot;,\&quot;cause\&quot;:\&quot;errorCode:INVALID_PARAM,errorMsg:null,cause:null\&quot;} | Refund order is incorrect |

## Samples

**Request parameter example**

```
{
    "refundId":"TQ10143867704492085",
    "pageNo":1,
    "pageSize":10
}
```

**Example of return parameters**

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
