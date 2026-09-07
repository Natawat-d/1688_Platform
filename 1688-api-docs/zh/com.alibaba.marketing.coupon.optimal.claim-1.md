# 通过商品领取最优化的优惠券

API: `com.alibaba.marketing:coupon.optimal.claim:1` · Category: 商品  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.marketing:coupon.optimal.claim-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.marketing/coupon.optimal.claim/{appKey}`  
需要授权 (access_token) · 需要签名

通过商品领取最优化的优惠券，一般在在下单前调用该接口，完成最优化的领取优惠券策略。

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `offerIds` | Long[] | 是 | 商品id列表 | [24910983123,2799731973] |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:alibaba.openapi.shared.common.coupon.optimal.claim.ResultModel](#m-alibaba-openapi-shared-common-coupon-optimal-claim-resultmodel) | 是 |  |  |

<a id="m-alibaba-openapi-shared-common-coupon-optimal-claim-resultmodel"></a>
#### alibaba.openapi.shared.common.coupon.optimal.claim.ResultModel

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `success` | boolean | 是 |  |  |
| `code` | java.lang.String | 是 |  |  |
| `message` | java.lang.String | 是 |  |  |
| `result` | [message:alibaba.ocean.openplatform.biz.market.result.BestBizCouponGetResult](#m-alibaba-ocean-openplatform-biz-market-result-bestbizcoupongetresult) | 是 |  |  |

<a id="m-alibaba-ocean-openplatform-biz-market-result-bestbizcoupongetresult"></a>
#### alibaba.ocean.openplatform.biz.market.result.BestBizCouponGetResult

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `couponIds` | java.lang.String[] | 是 |  |  |
