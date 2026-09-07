# 发起免密支付

API: `com.alibaba.trade:alibaba.trade.pay.protocolPay.preparePay:1` · Category: 支付  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.trade.pay.protocolPay.preparePay-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.trade.pay.protocolPay.preparePay/{appKey}`  
需要授权 (access_token) · 需要签名

发起免密支付，会自动判断是否开通了支付宝或者诚E赊的免密支付，并发起扣款。优先发起诚E赊自动扣款，如果失败，则尝试支付宝自动扣款。该接口目前返回错误码不详，在发起扣款失败后，建议重试3次，不要无限制重试。

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `tradeWithholdPreparePayParam` | [message:alibaba.ocean.openplatform.biz.trade.param.TradeWithholdPreparePayParam](#m-alibaba-ocean-openplatform-biz-trade-param-tradewithholdpreparepayparam) | 是 | 发起免密支付 |   |

<a id="m-alibaba-ocean-openplatform-biz-trade-param-tradewithholdpreparepayparam"></a>
#### alibaba.ocean.openplatform.biz.trade.param.TradeWithholdPreparePayParam

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `orderId` | java.lang.Long | 是 | 订单ID | 1938489823 |
| `payChannel` | String | 否 | 先采后付支付传入shegou，支付宝支付传入alipay。不传值默认使用API描述优先级进行代扣 | alipay |
| `payAmount` | Long | 否 | 付款总金额,单位分 | 123 |
| `opRequestId` | String | 否 | requestid | 134134134 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:com.alibaba.ocean.openplatform.biz.trade.result.TradeWithholdPreparePayResult.rresult](#m-com-alibaba-ocean-openplatform-biz-trade-result-tradewithholdpreparepayresult-rresult) | 是 | 免密支付结果 | {} |

<a id="m-com-alibaba-ocean-openplatform-biz-trade-result-tradewithholdpreparepayresult-rresult"></a>
#### com.alibaba.ocean.openplatform.biz.trade.result.TradeWithholdPreparePayResult.rresult

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `success` | Boolean | 是 | 是否成功 | true |
| `code` | String | 是 | 错误码 | null |
| `message` | String | 是 | 错误消息 | null |
| `result` | [message:com.alibaba.ocean.openplatform.biz.trade.result.TradeWithholdPreparePayResult.mresult](#m-com-alibaba-ocean-openplatform-biz-trade-result-tradewithholdpreparepayresult-mresult) | 是 | 扣款返回值 | {} |

<a id="m-com-alibaba-ocean-openplatform-biz-trade-result-tradewithholdpreparepayresult-mresult"></a>
#### com.alibaba.ocean.openplatform.biz.trade.result.TradeWithholdPreparePayResult.mresult

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `payChannel` | String | 是 | 支付成功渠道 | Alipay |
| `paySuccess` | Boolean | 是 | 支付是否成功,在超时的情况下，可能返回false但实际扣款成功的情况，需要查询订单实际支付状态 | true |
