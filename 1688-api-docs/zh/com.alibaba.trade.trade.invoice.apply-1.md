# 买家申请开票

API: `com.alibaba.trade:trade.invoice.apply:1` · Category: 发票  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:trade.invoice.apply-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/trade.invoice.apply/{appKey}`  
需要授权 (access_token) · 需要签名

买家申请开票，支持批量。
前置操作：
获取买家抬头；
获取可开票订单及支持的开票类型；
获取订单的可开票金额；

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `input` | [message:alibaba.ocean.openplatform.biz.trade.param.OpApplyInvoiceParam](#m-alibaba-ocean-openplatform-biz-trade-param-opapplyinvoiceparam) | 是 | 入参 | {} |

<a id="m-alibaba-ocean-openplatform-biz-trade-param-opapplyinvoiceparam"></a>
#### alibaba.ocean.openplatform.biz.trade.param.OpApplyInvoiceParam

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `invoiceApplyModelList` | [message:alibaba.ocean.openplatform.biz.trade.common.model.OpInvoiceApplyModel[]](#m-alibaba-ocean-openplatform-biz-trade-common-model-opinvoiceapplymodel[]) | 是 | 发票申请模型 | {} |

<a id="m-alibaba-ocean-openplatform-biz-trade-common-model-opinvoiceapplymodel[]"></a>
#### alibaba.ocean.openplatform.biz.trade.common.model.OpInvoiceApplyModel[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `orderId` | java.lang.Long | 是 | 主订单号 | 123 |
| `amount` | java.lang.Long | 是 | 发票合计含税金额，通过接口trade.invoiceAmount.getList获取 | 2 |
| `invoiceType` | java.lang.String | 是 | 发票类型：VATAX_SPEC(2, &quot;增值税专用发票&quot;)；VATAX_COMM(1, &quot;增值税普通发票&quot;)。 | VATAX_COMM |
| `purchaserInvoiceTitleModel` | [message:alibaba.ocean.openplatform.biz.trade.result.OpInvoiceTitleModel](#m-alibaba-ocean-openplatform-biz-trade-result-opinvoicetitlemodel) | 是 | 买家发票抬头，通过接口trade.invoiceTitle.getPageList获取，请将查到的抬头原样传入 | {} |

<a id="m-alibaba-ocean-openplatform-biz-trade-result-opinvoicetitlemodel"></a>
#### alibaba.ocean.openplatform.biz.trade.result.OpInvoiceTitleModel

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `titleType` | java.lang.String | 是 | 抬头类型，个人和社会组织都是PERSONAL：PERSONAL(0,&quot;个人和社会组织&quot;),COMPANY(1,&quot;企业&quot;); | PERSONAL |
| `title` | java.lang.String | 是 | 抬头 | 抬头 |
| `taxpayerIdentify` | java.lang.String | 否 | 纳税人识别号(企业开票必填) | 123 |
| `bankName` | java.lang.String | 否 | 开户行(企业开专票必填) | 开户行 |
| `bankAccountId` | java.lang.String | 否 | 银行账号(企业开专票必填) | 123 |
| `registerAddress` | java.lang.String | 否 | 企业注册地址(企业开专票必填) | 企业注册地址 |
| `registerPhone` | java.lang.String | 否 | 企业电话(企业开专票必填) | 123 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:ApplyInvoiceResultModel](#m-applyinvoiceresultmodel) | 是 | 请求结果 | {} |

<a id="m-applyinvoiceresultmodel"></a>
#### ApplyInvoiceResultModel

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `code` | String | 是 | code |   |
| `message` | String | 是 | message |   |
| `result` | [message:alibaba.ocean.openplatform.biz.trade.result.OpApplyInvoiceResult](#m-alibaba-ocean-openplatform-biz-trade-result-opapplyinvoiceresult) | 是 | 开票结果 | {} |
| `retCodes` | String[] | 是 | retCodes | [] |
| `subCode` | String | 是 | subCode |   |
| `subMessage` | String | 是 | subMessage |   |
| `success` | Boolean | 是 | 是否成功 | true |

<a id="m-alibaba-ocean-openplatform-biz-trade-result-opapplyinvoiceresult"></a>
#### alibaba.ocean.openplatform.biz.trade.result.OpApplyInvoiceResult

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `failedList` | [message:alibaba.ocean.openplatform.biz.trade.common.model.OpInvoicingResultModel[]](#m-alibaba-ocean-openplatform-biz-trade-common-model-opinvoicingresultmodel[]) | 否 | 开票失败列表 | [] |
| `successList` | [message:alibaba.ocean.openplatform.biz.trade.common.model.OpInvoicingResultModel[]](#m-alibaba-ocean-openplatform-biz-trade-common-model-opinvoicingresultmodel[]) | 否 | 开票成功列表 | [] |

<a id="m-alibaba-ocean-openplatform-biz-trade-common-model-opinvoicingresultmodel[]"></a>
#### alibaba.ocean.openplatform.biz.trade.common.model.OpInvoicingResultModel[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `outBizId` | java.lang.String | 是 | 主订单id | 123 |
| `result` | boolean | 是 | 开票结果 | true |
| `orderId` | Long | 是 | 主订单id | 123 |
| `tradeOrderCompleted` | Boolean | 是 | 是否交易完结 | true |
| `errorCode` | String | 是 | 错误码（失败时填写） | XXX |
| `errorDesc` | String | 是 | 错误描述（失败时填写） | XXX |

## 错误码

| 错误码 | 现象 | 解决方案 |
|---|---|---|
| KINGSUNS_BASIC_PARAM_CHECK_ERROR | 基础参数检查异常 入参为null，或入参购方id为空 |  |
| PARAM_INVALID | 参数异常;抬头参数异常 |  |
| ORDER_CANCELED_ERROR | 订单已取消，不允许开票 |  |
| INVOICE_ALREADY_APPLIED | 曾经从别的渠道历史已申请过发票 |  |
| INVOICE_NOT_SUPPORT_SPEC | 该订单的商品付款时不支持开具专票，请您联系商家咨询，开启\&quot;专票\&quot;开票设置&quot; |  |
| INVOICE_NOT_SUPPORT_COMMON | 该订单的商品付款时不支持开具普票，请您联系商家咨询，开启\&quot;全店商品\&quot;开票设置&quot; |  |
| APPLY_CREATE_INVOICE_EXCEPTION | 申请创建发票异常 兜底异常 |  |
