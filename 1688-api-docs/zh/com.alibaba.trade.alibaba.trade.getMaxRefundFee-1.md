# 申请退款时查询最大可退费用

API: `com.alibaba.trade:alibaba.trade.getMaxRefundFee:1` · Category: 退货退款  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.trade.getMaxRefundFee-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.trade.getMaxRefundFee/{appKey}`  
需要授权 (access_token) · 需要签名

申请退款时查询最大可退费用

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `input` | [message:alibaba.ocean.openplatform.biz.trade.param.OpMaxRefundFeeGetParam](#m-alibaba-ocean-openplatform-biz-trade-param-opmaxrefundfeegetparam) | 是 |  |   |

<a id="m-alibaba-ocean-openplatform-biz-trade-param-opmaxrefundfeegetparam"></a>
#### alibaba.ocean.openplatform.biz.trade.param.OpMaxRefundFeeGetParam

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `goodsStatus` | java.lang.String | 是 | 货物状态 |  售中等待卖家发货:"refundWaitSellerSend"; 售中等待买家收货:"refundWaitBuyerReceive"; 售中已收货（未确认完成交易）:"refundBuyerReceived" 售后未收货:"aftersaleBuyerNotReceived"; 售后已收到货:"aftersaleBuyerReceived" |
| `orderId` | java.lang.Long | 是 | 订单ID | 123 |
| `refundId` | java.lang.String | 否 | 退款单必须处于退款中，可不传 | TQ123 |
| `refundGoodsCountList` | [message:alibaba.ocean.openplatform.biz.trade.common.model.OrderEntryCountModel[]](#m-alibaba-ocean-openplatform-biz-trade-common-model-orderentrycountmodel[]) | 是 | 退货数量 | [{1: 1}] |

<a id="m-alibaba-ocean-openplatform-biz-trade-common-model-orderentrycountmodel[]"></a>
#### alibaba.ocean.openplatform.biz.trade.common.model.OrderEntryCountModel[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `id` | Long | 是 | 子订单id | 1 |
| `count` | Integer | 是 | 子订单购买商品数量 | 1 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:alibaba.ocean.openplatform.common.OpMaxRefundFeeResultModel](#m-alibaba-ocean-openplatform-common-opmaxrefundfeeresultmodel) | 是 |  |   |

<a id="m-alibaba-ocean-openplatform-common-opmaxrefundfeeresultmodel"></a>
#### alibaba.ocean.openplatform.common.OpMaxRefundFeeResultModel

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `code` | java.lang.String | 是 |  |   |
| `message` | java.lang.String | 是 |  |   |
| `result` | [message:alibaba.ocean.openplatform.biz.trade.common.model.OpMaxRefundFeeModel](#m-alibaba-ocean-openplatform-biz-trade-common-model-opmaxrefundfeemodel) | 是 |  |   |
| `retCodes` | java.lang.String[] | 是 |  |   |
| `subCode` | java.lang.String | 是 |  |   |
| `subMessage` | java.lang.String | 是 |  |   |
| `success` | boolean | 是 |  |   |

<a id="m-alibaba-ocean-openplatform-biz-trade-common-model-opmaxrefundfeemodel"></a>
#### alibaba.ocean.openplatform.biz.trade.common.model.OpMaxRefundFeeModel

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `maxGoodsFee` | java.lang.Long | 是 | 最大可退款货品金额，单位：分 | 7 |
| `maxPostFee` | java.lang.Long | 是 | 最大可退运费，单位：分 | 2 |
| `maxRefundFee` | java.lang.Long | 是 | 最大可退款金额，即运费+货品金额的上限值，单位：分 | 9 |
| `maxSubstituteFetchFee` | java.lang.Long | 是 | 最大可退官方代提服务费，单位：分 | 0 |
