# Query payment channels supported by an order

Original name: 查询订单可以支持的支付渠道  
API: `com.alibaba.trade:alibaba.trade.payWay.query:1` · Category: Payment  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.trade.payWay.query-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.trade.payWay.query/{appKey}`  
Requires user authorization (access_token) · Requires signature

Query the payment methods or channels available for an unpaid order.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `orderId` | String | yes | Order number | 123123 |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `success` | String | yes | Whether successful | true |
| `errorCode` | String | yes | Error code | 500_1 |
| `errorMsg` | String | yes | Error message |   |
| `resultList` | [message:alibaba.ocean.openplatform.biz.trade.result.TradePayTypeResult](#m-alibaba-ocean-openplatform-biz-trade-result-tradepaytyperesult) | yes | Return result | {} |

<a id="m-alibaba-ocean-openplatform-biz-trade-result-tradepaytyperesult"></a>
#### alibaba.ocean.openplatform.biz.trade.result.TradePayTypeResult

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `channels` | [message:alibaba.ocean.openplatform.biz.trade.result.PayTypeInfo[]](#m-alibaba-ocean-openplatform-biz-trade-result-paytypeinfo[]) | yes | List of available payment channels | [] |
| `orderId` | java.lang.String | yes | Order number | 1231231211 |
| `payFee` | java.lang.Long | yes | Payment amount, in cents | 100 |
| `timeout` | java.lang.String | yes | Latest payment time | 2018-10-01 00:00:00 |

<a id="m-alibaba-ocean-openplatform-biz-trade-result-paytypeinfo[]"></a>
#### alibaba.ocean.openplatform.biz.trade.result.PayTypeInfo[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `code` | java.lang.Long | yes | Payment channel code. 1: Alipay 2: MYbank Trust Pay 3: Cheng-e-She 4: Corporate transfer 5: Shexiaobao (credit sale) 6: Electronic acceptance draft 7: Account period payment 8: Combined payment channel 9: No payment 10: Lingshoutong credit purchase 12: Declared payment 13: Payment platform 14: MYbank electronic bank acceptance draft 15: Bank transfer 16: Kuajingbao (Cross-Border Pay) 17: Red envelope 20: Kuajingbao (Cross-Border Pay) 35: MYbank cross-border direct purchase | 1 |
| `name` | java.lang.String | yes | Payment channel name. 1: Alipay; 2: MYbank Trusted Payment; 3: Cheng-e-She; 4: corporate bank transfer; 5: Shexiaobao (credit sales); 6: electronic acceptance bill; 7: account period payment; 8: combined payment channel; 9: no payment; 10: Lingshoutong credit purchase; 12: declared payment; 13: payment platform; 14: MYbank electronic bank acceptance bill; 15: bank transfer; 16: Kuajingbao; 17: red packet; 20: Kuajingbao; 35: MYbank cross-border direct purchase | 支付宝 |

## Error codes

| Error | Symptom | How to fix |
|---|---|---|
| 500_2 | No permission to obtain the available payment methods for this order. | Check the authorized user; the authorized account must be the buyer, and must be the buyer's main account. |

## Samples

**Example of returned result**

```
{
  "resultList": {
    "channels": [
      {
        "code": 1,
        "name": "支付宝"
      }
    ],
    "orderId": "239695213738498520",
    "payFee": 32120,
    "timeout": "2018-11-04 14:00:45"
  },
  "success": true
}
```
