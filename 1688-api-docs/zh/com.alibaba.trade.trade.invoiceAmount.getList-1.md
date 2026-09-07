# 查询订单可开票金额

API: `com.alibaba.trade:trade.invoiceAmount.getList:1` · Category: 发票  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:trade.invoiceAmount.getList-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/trade.invoiceAmount.getList/{appKey}`  
需要授权 (access_token) · 需要签名

查询订单可开票金额。
前置获取可开票订单的方法：
1. alibaba.trade.getBuyerOrderList-1：入参needInvoicingSetting传true；通过出参invoicingSettingModel.tradeInvoiceStatus判断是否可开票。
2. alibaba.trade.get.buyerView-1：入参includeFields包含InvoicingSetting

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `input` | [message:alibaba.ocean.openplatform.biz.trade.param.OpOrderInvoiceAmountQueryParam](#m-alibaba-ocean-openplatform-biz-trade-param-oporderinvoiceamountqueryparam) | 是 |  |   |

<a id="m-alibaba-ocean-openplatform-biz-trade-param-oporderinvoiceamountqueryparam"></a>
#### alibaba.ocean.openplatform.biz.trade.param.OpOrderInvoiceAmountQueryParam

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `orderIds` | java.lang.Long[] | 是 | 订单ID列表 | [123, 456] |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:alibaba.ocean.openplatform.common.OpOrderInvoiceAmountResultModel](#m-alibaba-ocean-openplatform-common-oporderinvoiceamountresultmodel) | 是 |  |   |

<a id="m-alibaba-ocean-openplatform-common-oporderinvoiceamountresultmodel"></a>
#### alibaba.ocean.openplatform.common.OpOrderInvoiceAmountResultModel

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `code` | java.lang.String | 是 |  |   |
| `message` | java.lang.String | 是 |  |   |
| `result` | [message:alibaba.ocean.openplatform.biz.trade.result.OpOrderInvoiceAmountQueryResult](#m-alibaba-ocean-openplatform-biz-trade-result-oporderinvoiceamountqueryresult) | 是 |  |   |
| `retCodes` | java.lang.String[] | 是 |  |   |
| `subCode` | java.lang.String | 是 |  |   |
| `subMessage` | java.lang.String | 是 |  |   |
| `success` | boolean | 是 |  |   |

<a id="m-alibaba-ocean-openplatform-biz-trade-result-oporderinvoiceamountqueryresult"></a>
#### alibaba.ocean.openplatform.biz.trade.result.OpOrderInvoiceAmountQueryResult

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `orderInvoiceAmountModelList` | [message:alibaba.ocean.openplatform.biz.trade.common.model.OrderInvoiceAmountModel[]](#m-alibaba-ocean-openplatform-biz-trade-common-model-orderinvoiceamountmodel[]) | 是 | 订单可开票金额模型列表 | [] |

<a id="m-alibaba-ocean-openplatform-biz-trade-common-model-orderinvoiceamountmodel[]"></a>
#### alibaba.ocean.openplatform.biz.trade.common.model.OrderInvoiceAmountModel[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `amount` | java.lang.Long | 是 | 可开票金额，单位：分 | 2 |
| `orderId` | java.lang.Long | 是 | 订单ID | 123 |
