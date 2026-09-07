# Buyer adds an order message

Original name: 买家补充订单留言接口  
API: `com.alibaba.trade:alibaba.trade.addFeedback:1` · Category: Orders  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.trade.addFeedback-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.trade.addFeedback/{appKey}`  
Requires user authorization (access_token) · Requires signature

Buyer adds a supplementary message to an order. The total message length must not exceed 500 characters.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `tradeFeedbackParam` | [message:alibaba.ocean.openplatform.biz.trade.param.TradeFeedbackParam](#m-alibaba-ocean-openplatform-biz-trade-param-tradefeedbackparam) | yes | Request parameters | {"feedback":"test","orderId":"123123213"} |

<a id="m-alibaba-ocean-openplatform-biz-trade-param-tradefeedbackparam"></a>
#### alibaba.ocean.openplatform.biz.trade.param.TradeFeedbackParam

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `feedback` | java.lang.String | yes | Message/comment | 留言 |
| `orderId` | java.lang.String | yes | Order ID | 12344444555545 |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.ocean.openplatform.biz.trade.result.TradeFeedbackResult](#m-alibaba-ocean-openplatform-biz-trade-result-tradefeedbackresult) | yes | Return result | {} |
| `code` | String | yes | Error code | 500_2 |
| `message` | String | yes | Error description | remote service error |
| `success` | Boolean | yes | Whether successful | false |

<a id="m-alibaba-ocean-openplatform-biz-trade-result-tradefeedbackresult"></a>
#### alibaba.ocean.openplatform.biz.trade.result.TradeFeedbackResult

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `errorInfo` | java.lang.String | yes | Error description | 错误描述 |
| `errorCode` | java.lang.String | yes | Error code | 400_1 |
| `success` | Boolean | yes | Whether successful | true |

## Error codes

| Error | Symptom | How to fix |
|---|---|---|
| 400_1 | parameter invalid | Check whether the parameter was passed and whether the parameter type meets the requirements |
| 400_2 | need authorization | Authorized login required |
| 500_1 | invoke remote service error | Exception calling the remote service. Please consult technical support in the ISV community group or submit a ticket to the help center. |
| 500_2 | remote service error | Service exception, please try again later, or ask for support in the ISV community group or submit a ticket at the help center |
| 500_2 | invalid parameter error | Invalid request parameters; please consult technical support in the ISV exchange group or submit a ticket in the help center |
| 500_2 | user order not exist error | The order corresponding to the order number is not an order of the currently authorized user; no permission to operate |
| 500_2 | order not exist error | The order number does not correspond to any order, please check whether the order number is correct |

## Samples

**Request example**

```
{"feedback":"test","orderId":"159095856057498520"}
```

**Return example**

```
{"result":{"success":true},"success":true}
```
