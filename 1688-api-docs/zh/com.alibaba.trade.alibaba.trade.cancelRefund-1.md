# 取消退款退货申请 

API: `com.alibaba.trade:alibaba.trade.cancelRefund:1` · Category: 退货退款  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.trade.cancelRefund-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.trade.cancelRefund/{appKey}`  
无需授权 · 需要签名

取消退款退货申请

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 否 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `refundId` | String | 是 | 退款单id | TQ267395256051660259 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:alibaba.ocean.openplatform.biz.trade.result.RefundCancelResult](#m-alibaba-ocean-openplatform-biz-trade-result-refundcancelresult) | 是 |  |  |

<a id="m-alibaba-ocean-openplatform-biz-trade-result-refundcancelresult"></a>
#### alibaba.ocean.openplatform.biz.trade.result.RefundCancelResult

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `success` | boolean | 是 |  |  |
| `errorInfo` | java.lang.String | 是 |  |  |
| `errorCode` | java.lang.String | 是 |  |  |
