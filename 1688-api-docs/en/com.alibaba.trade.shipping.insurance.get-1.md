# Query shipping-insurance info

Original name: 运费险信息查询  
API: `com.alibaba.trade:shipping.insurance.get:1` · Category: Logistics  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:shipping.insurance.get-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/shipping.insurance.get/{appKey}`  
Requires user authorization (access_token) · Requires signature

Query shipping-insurance information.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `orderId` | Long | yes | Order number | 订单号 |
| `type` | String | yes | Shipping insurance type | givenByPlatform平台赠送，givenByMerchant商家赠送 |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.ocean.openplatform.common.ResultModel](#m-alibaba-ocean-openplatform-common-resultmodel) | yes |  |  |

<a id="m-alibaba-ocean-openplatform-common-resultmodel"></a>
#### alibaba.ocean.openplatform.common.ResultModel

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `success` | java.lang.Boolean | yes | Whether successful | 是否成功 |
| `code` | java.lang.String | yes | Response code | 响应码 |
| `message` | java.lang.String | yes | Response information | 响应信息 |
| `result` | [message:alibaba.ocean.openplatform.biz.trade.result.TradeFreightPolicyResult](#m-alibaba-ocean-openplatform-biz-trade-result-tradefreightpolicyresult) | yes | Return result | 返回结果 |

<a id="m-alibaba-ocean-openplatform-biz-trade-result-tradefreightpolicyresult"></a>
#### alibaba.ocean.openplatform.biz.trade.result.TradeFreightPolicyResult

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `insuranceId` | java.lang.Long | yes | Policy ID | 保单id |
| `orderId` | java.lang.Long | yes | Order id | 订单id |
| `tradeClaimList` | [message:alibaba.ocean.openplatform.biz.trade.result.TradeClaimResult[]](#m-alibaba-ocean-openplatform-biz-trade-result-tradeclaimresult[]) | yes | Claim order information | 理赔单信息 |

<a id="m-alibaba-ocean-openplatform-biz-trade-result-tradeclaimresult[]"></a>
#### alibaba.ocean.openplatform.biz.trade.result.TradeClaimResult[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `applicationTime` | java.util.Date | yes | Application time | 申请时间	 |
| `claimAmount` | java.lang.Long | yes | Claim amount | 理赔金额 |
| `claimId` | java.lang.String | yes | Claim order id | 理赔单id |
| `payTime` | java.util.Date | yes | Payment time | 打款时间 |
| `status` | java.lang.String | yes | Claim status | 理赔状态 |
| `tradeNO` | java.lang.String | yes | Alipay transaction serial number | 支付宝交易流水号 |
