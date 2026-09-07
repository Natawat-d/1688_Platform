# 交易订单支付咨询

API: `com.alibaba.trade:trade.orderpay.analysis:1` · Category: 支付  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:trade.orderpay.analysis-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/trade.orderpay.analysis/{appKey}`  
需要授权 (access_token) · 需要签名

交易订单支付咨询接口，用于分析订单使用什么支付方式等。

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `orderIds` | java.lang.Long[] | 是 | 订单列表 | [3535260697417660107] |
| `payChannel` | java.lang.String | 是 | 支付渠道  alipay(支付宝),shegou(诚e赊),kjpayV2(跨境宝) | kjpayV2 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:alibaba.ocean.openplatform.common.tradeorderpayanalysis.ResultModel](#m-alibaba-ocean-openplatform-common-tradeorderpayanalysis-resultmodel) | 是 |  |  |

<a id="m-alibaba-ocean-openplatform-common-tradeorderpayanalysis-resultmodel"></a>
#### alibaba.ocean.openplatform.common.tradeorderpayanalysis.ResultModel

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `success` | boolean | 是 |  |  |
| `code` | java.lang.String | 是 |  |  |
| `message` | java.lang.String | 是 |  |  |
| `result` | [message:alibaba.ocean.openplatform.biz.trade.result.PayAnalysisResult](#m-alibaba-ocean-openplatform-biz-trade-result-payanalysisresult) | 是 |  |  |

<a id="m-alibaba-ocean-openplatform-biz-trade-result-payanalysisresult"></a>
#### alibaba.ocean.openplatform.biz.trade.result.PayAnalysisResult

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `orderTotalCount` | java.lang.Long | 是 | 总订单数量 | 2 |
| `totalPayFee` | java.lang.Long | 是 | 总付款金额 | 150000 |
| `orderIds` | java.lang.Long[] | 是 | 订单列表 | [265354896466,854635455] |
| `payChannel` | [message:com.alibaba.ocean.openplatform.biz.trade.result.PayChannel](#m-com-alibaba-ocean-openplatform-biz-trade-result-paychannel) | 是 | 支付渠道 | {"channel":"alipay","avaliable":true} |

<a id="m-com-alibaba-ocean-openplatform-biz-trade-result-paychannel"></a>
#### com.alibaba.ocean.openplatform.biz.trade.result.PayChannel

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `channel` | String | 是 | 支付渠道 | alipay |
| `avaliable` | Boolean | 是 | 是否支持 | true |
