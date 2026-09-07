# Order payment consultation

Original name: 交易订单支付咨询  
API: `com.alibaba.trade:trade.orderpay.analysis:1` · Category: Payment  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:trade.orderpay.analysis-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/trade.orderpay.analysis/{appKey}`  
Requires user authorization (access_token) · Requires signature

Order-payment consultation interface, used to analyse which payment method an order uses, and so on.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `orderIds` | java.lang.Long[] | yes | Order list | [3535260697417660107] |
| `payChannel` | java.lang.String | yes | Payment channel  alipay (Alipay), shegou (Cheng-e-She), kjpayV2 (Kuajingbao) | kjpayV2 |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.ocean.openplatform.common.tradeorderpayanalysis.ResultModel](#m-alibaba-ocean-openplatform-common-tradeorderpayanalysis-resultmodel) | yes |  |  |

<a id="m-alibaba-ocean-openplatform-common-tradeorderpayanalysis-resultmodel"></a>
#### alibaba.ocean.openplatform.common.tradeorderpayanalysis.ResultModel

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `success` | boolean | yes |  |  |
| `code` | java.lang.String | yes |  |  |
| `message` | java.lang.String | yes |  |  |
| `result` | [message:alibaba.ocean.openplatform.biz.trade.result.PayAnalysisResult](#m-alibaba-ocean-openplatform-biz-trade-result-payanalysisresult) | yes |  |  |

<a id="m-alibaba-ocean-openplatform-biz-trade-result-payanalysisresult"></a>
#### alibaba.ocean.openplatform.biz.trade.result.PayAnalysisResult

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `orderTotalCount` | java.lang.Long | yes | Total order quantity | 2 |
| `totalPayFee` | java.lang.Long | yes | Total payment amount | 150000 |
| `orderIds` | java.lang.Long[] | yes | Order list | [265354896466,854635455] |
| `payChannel` | [message:com.alibaba.ocean.openplatform.biz.trade.result.PayChannel](#m-com-alibaba-ocean-openplatform-biz-trade-result-paychannel) | yes | Payment channel | {"channel":"alipay","avaliable":true} |

<a id="m-com-alibaba-ocean-openplatform-biz-trade-result-paychannel"></a>
#### com.alibaba.ocean.openplatform.biz.trade.result.PayChannel

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `channel` | String | yes | Payment channel | alipay |
| `avaliable` | Boolean | yes | Whether supported | true |
