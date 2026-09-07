# 申请交易仲裁

API: `com.alibaba.trade:trade.arbitration.apply:1` · Category: 退货退款  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:trade.arbitration.apply-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/trade.arbitration.apply/{appKey}`  
需要授权 (access_token) · 需要签名

交易仲裁申请

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `param` | [message:alibaba.ocean.openplatform.biz.trade.param.TradeArbitrateApplyParam](#m-alibaba-ocean-openplatform-biz-trade-param-tradearbitrateapplyparam) | 是 | 申诉参数模型 | {} |

<a id="m-alibaba-ocean-openplatform-biz-trade-param-tradearbitrateapplyparam"></a>
#### alibaba.ocean.openplatform.biz.trade.param.TradeArbitrateApplyParam

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `type` | String | 是 | 投诉类型 RefundComplaint(退款投诉),SafeRefundComplaint（安全退款投诉）,AfterSalesComplaint（售后单投诉）,TradeComplaint(售后订单投诉) | RefundComplaint |
| `refundId` | String | 否 | 退货退款单ID，退款投诉时，refundId必填 | TQ129988192325 |
| `orderId` | Long | 是 | 订单ID | 3356489569898 |
| `reasonText` | java.lang.String | 是 | 申请原因 | 货没收到 |
| `reasonId` | java.lang.Integer | 是 | 申请原因id | 12 |
| `picUrlEvidence` | java.lang.String | 否 | 证据图 | https://cnd01.1688.com/evidence.png |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:alibaba.openapi.shared.common.tradearbitrationapplyResultModel](#m-alibaba-openapi-shared-common-tradearbitrationapplyresultmodel) | 是 | 申诉返回模型 | {} |

<a id="m-alibaba-openapi-shared-common-tradearbitrationapplyresultmodel"></a>
#### alibaba.openapi.shared.common.tradearbitrationapplyResultModel

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `success` | boolean | 是 | 是否成功 | true |
| `code` | java.lang.String | 是 | 错误码 | null |
| `message` | java.lang.String | 是 | 错误信息 | null |
| `result` | java.lang.Long | 是 | 申诉单ID | 3655462545 |
