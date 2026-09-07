# Claim the optimal coupon for a product

Original name: 通过商品领取最优化的优惠券  
API: `com.alibaba.marketing:coupon.optimal.claim:1` · Category: Products  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.marketing:coupon.optimal.claim-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.marketing/coupon.optimal.claim/{appKey}`  
Requires user authorization (access_token) · Requires signature

Claim the best-value coupon for a product. Usually called before placing an order to complete the optimal coupon-claiming strategy.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `offerIds` | Long[] | yes | List of product IDs | [24910983123,2799731973] |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.openapi.shared.common.coupon.optimal.claim.ResultModel](#m-alibaba-openapi-shared-common-coupon-optimal-claim-resultmodel) | yes |  |  |

<a id="m-alibaba-openapi-shared-common-coupon-optimal-claim-resultmodel"></a>
#### alibaba.openapi.shared.common.coupon.optimal.claim.ResultModel

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `success` | boolean | yes |  |  |
| `code` | java.lang.String | yes |  |  |
| `message` | java.lang.String | yes |  |  |
| `result` | [message:alibaba.ocean.openplatform.biz.market.result.BestBizCouponGetResult](#m-alibaba-ocean-openplatform-biz-market-result-bestbizcoupongetresult) | yes |  |  |

<a id="m-alibaba-ocean-openplatform-biz-market-result-bestbizcoupongetresult"></a>
#### alibaba.ocean.openplatform.biz.market.result.BestBizCouponGetResult

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `couponIds` | java.lang.String[] | yes |  |  |
