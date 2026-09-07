# Apply for trade arbitration

Original name: 申请交易仲裁  
API: `com.alibaba.trade:trade.arbitration.apply:1` · Category: Returns & Refunds  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:trade.arbitration.apply-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/trade.arbitration.apply/{appKey}`  
Requires user authorization (access_token) · Requires signature

Trade arbitration application.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `param` | [message:alibaba.ocean.openplatform.biz.trade.param.TradeArbitrateApplyParam](#m-alibaba-ocean-openplatform-biz-trade-param-tradearbitrateapplyparam) | yes | Appeal parameter model | {} |

<a id="m-alibaba-ocean-openplatform-biz-trade-param-tradearbitrateapplyparam"></a>
#### alibaba.ocean.openplatform.biz.trade.param.TradeArbitrateApplyParam

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `type` | String | yes | Complaint type: RefundComplaint (refund complaint), SafeRefundComplaint (secure refund complaint), AfterSalesComplaint (after-sales case complaint), TradeComplaint (after-sales order complaint) | RefundComplaint |
| `refundId` | String | no | Return/refund order ID; refundId is required for refund complaints | TQ129988192325 |
| `orderId` | Long | yes | Order ID | 3356489569898 |
| `reasonText` | java.lang.String | yes | Application reason | 货没收到 |
| `reasonId` | java.lang.Integer | yes | Application reason ID | 12 |
| `picUrlEvidence` | java.lang.String | no | Evidence image | https://cnd01.1688.com/evidence.png |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.openapi.shared.common.tradearbitrationapplyResultModel](#m-alibaba-openapi-shared-common-tradearbitrationapplyresultmodel) | yes | Appeal return model | {} |

<a id="m-alibaba-openapi-shared-common-tradearbitrationapplyresultmodel"></a>
#### alibaba.openapi.shared.common.tradearbitrationapplyResultModel

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `success` | boolean | yes | Whether successful | true |
| `code` | java.lang.String | yes | Error code | null |
| `message` | java.lang.String | yes | Error message | null |
| `result` | java.lang.Long | yes | Appeal order ID | 3655462545 |
