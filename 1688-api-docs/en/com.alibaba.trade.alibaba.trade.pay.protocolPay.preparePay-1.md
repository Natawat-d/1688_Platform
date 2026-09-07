# Initiate password-free payment

Original name: 发起免密支付  
API: `com.alibaba.trade:alibaba.trade.pay.protocolPay.preparePay:1` · Category: Payment  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.trade.pay.protocolPay.preparePay-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.trade.pay.protocolPay.preparePay/{appKey}`  
Requires user authorization (access_token) · Requires signature

Initiates a password-free payment. Automatically detects whether Alipay or Cheng-e-She password-free payment is enabled and initiates the debit. Cheng-e-She auto-debit is tried first; if it fails, Alipay auto-debit is attempted. The error codes returned by this interface are currently not detailed; after a failed debit, retry up to 3 times rather than indefinitely.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `tradeWithholdPreparePayParam` | [message:alibaba.ocean.openplatform.biz.trade.param.TradeWithholdPreparePayParam](#m-alibaba-ocean-openplatform-biz-trade-param-tradewithholdpreparepayparam) | yes | Initiate password-free payment |   |

<a id="m-alibaba-ocean-openplatform-biz-trade-param-tradewithholdpreparepayparam"></a>
#### alibaba.ocean.openplatform.biz.trade.param.TradeWithholdPreparePayParam

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `orderId` | java.lang.Long | yes | Order ID | 1938489823 |
| `payChannel` | String | no | Pass shegou for Buy-Now-Pay-Later payment, and alipay for Alipay payment. If no value is passed, the deduction defaults to using the priority described in the API documentation. | alipay |
| `payAmount` | Long | no | Total payment amount, in cents | 123 |
| `opRequestId` | String | no | requestid | 134134134 |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:com.alibaba.ocean.openplatform.biz.trade.result.TradeWithholdPreparePayResult.rresult](#m-com-alibaba-ocean-openplatform-biz-trade-result-tradewithholdpreparepayresult-rresult) | yes | Password-free payment result | {} |

<a id="m-com-alibaba-ocean-openplatform-biz-trade-result-tradewithholdpreparepayresult-rresult"></a>
#### com.alibaba.ocean.openplatform.biz.trade.result.TradeWithholdPreparePayResult.rresult

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `success` | Boolean | yes | Whether successful | true |
| `code` | String | yes | Error code | null |
| `message` | String | yes | Error message | null |
| `result` | [message:com.alibaba.ocean.openplatform.biz.trade.result.TradeWithholdPreparePayResult.mresult](#m-com-alibaba-ocean-openplatform-biz-trade-result-tradewithholdpreparepayresult-mresult) | yes | Deduction return value | {} |

<a id="m-com-alibaba-ocean-openplatform-biz-trade-result-tradewithholdpreparepayresult-mresult"></a>
#### com.alibaba.ocean.openplatform.biz.trade.result.TradeWithholdPreparePayResult.mresult

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `payChannel` | String | yes | Successful payment channel | Alipay |
| `paySuccess` | Boolean | yes | Whether the payment was successful. In case of a timeout, false may be returned even though the deduction actually succeeded; you need to query the order's actual payment status | true |
