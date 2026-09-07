# 买家确认收货

API: `com.alibaba.trade:trade.receivegoods.confirm:1` · Category: 订单  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:trade.receivegoods.confirm-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/trade.receivegoods.confirm/{appKey}`  
需要授权 (access_token) · 需要签名

买家确认收货

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `orderId` | java.lang.Long | 是 | 订单ID | 56623232655125698 |
| `orderEntryIds` | java.lang.Long[] | 是 | 子订单ID | 562356635566365512 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:alibaba.ocean.openplatform.biz.trade.result.TradeConfirmReceiptResult](#m-alibaba-ocean-openplatform-biz-trade-result-tradeconfirmreceiptresult) | 是 |  |  |

<a id="m-alibaba-ocean-openplatform-biz-trade-result-tradeconfirmreceiptresult"></a>
#### alibaba.ocean.openplatform.biz.trade.result.TradeConfirmReceiptResult

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `success` | boolean | 是 |  |  |
| `errorInfo` | java.lang.String | 是 |  |  |
| `errorCode` | java.lang.String | 是 |  |  |
