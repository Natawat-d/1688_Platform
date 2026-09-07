# 运费险信息查询

API: `com.alibaba.trade:shipping.insurance.get:1` · Category: 物流  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:shipping.insurance.get-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/shipping.insurance.get/{appKey}`  
需要授权 (access_token) · 需要签名

运费险信息查询

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `orderId` | Long | 是 | 订单号 | 订单号 |
| `type` | String | 是 | 运费险类型 | givenByPlatform平台赠送，givenByMerchant商家赠送 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:alibaba.ocean.openplatform.common.ResultModel](#m-alibaba-ocean-openplatform-common-resultmodel) | 是 |  |  |

<a id="m-alibaba-ocean-openplatform-common-resultmodel"></a>
#### alibaba.ocean.openplatform.common.ResultModel

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `success` | java.lang.Boolean | 是 | 是否成功 | 是否成功 |
| `code` | java.lang.String | 是 | 响应码 | 响应码 |
| `message` | java.lang.String | 是 | 响应信息 | 响应信息 |
| `result` | [message:alibaba.ocean.openplatform.biz.trade.result.TradeFreightPolicyResult](#m-alibaba-ocean-openplatform-biz-trade-result-tradefreightpolicyresult) | 是 | 返回结果 | 返回结果 |

<a id="m-alibaba-ocean-openplatform-biz-trade-result-tradefreightpolicyresult"></a>
#### alibaba.ocean.openplatform.biz.trade.result.TradeFreightPolicyResult

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `insuranceId` | java.lang.Long | 是 | 保单id | 保单id |
| `orderId` | java.lang.Long | 是 | 订单id | 订单id |
| `tradeClaimList` | [message:alibaba.ocean.openplatform.biz.trade.result.TradeClaimResult[]](#m-alibaba-ocean-openplatform-biz-trade-result-tradeclaimresult[]) | 是 | 理赔单信息 | 理赔单信息 |

<a id="m-alibaba-ocean-openplatform-biz-trade-result-tradeclaimresult[]"></a>
#### alibaba.ocean.openplatform.biz.trade.result.TradeClaimResult[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `applicationTime` | java.util.Date | 是 | 申请时间 | 申请时间	 |
| `claimAmount` | java.lang.Long | 是 | 理赔金额 | 理赔金额 |
| `claimId` | java.lang.String | 是 | 理赔单id | 理赔单id |
| `payTime` | java.util.Date | 是 | 打款时间 | 打款时间 |
| `status` | java.lang.String | 是 | 理赔状态 | 理赔状态 |
| `tradeNO` | java.lang.String | 是 | 支付宝交易流水号 | 支付宝交易流水号 |
