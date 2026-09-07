# Query invoiceable amount of orders

Original name: 查询订单可开票金额  
API: `com.alibaba.trade:trade.invoiceAmount.getList:1` · Category: Invoicing  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:trade.invoiceAmount.getList-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/trade.invoiceAmount.getList/{appKey}`  
Requires user authorization (access_token) · Requires signature

Query the invoiceable amount of orders. To find invoiceable orders first: 1. alibaba.trade.getBuyerOrderList-1: pass needInvoicingSetting=true and check invoicingSettingModel.tradeInvoiceStatus in the response to see whether an order can be invoiced. 2. alibaba.trade.get.buyerView-1: include InvoicingSetting in the includeFields parameter.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `input` | [message:alibaba.ocean.openplatform.biz.trade.param.OpOrderInvoiceAmountQueryParam](#m-alibaba-ocean-openplatform-biz-trade-param-oporderinvoiceamountqueryparam) | yes |  |   |

<a id="m-alibaba-ocean-openplatform-biz-trade-param-oporderinvoiceamountqueryparam"></a>
#### alibaba.ocean.openplatform.biz.trade.param.OpOrderInvoiceAmountQueryParam

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `orderIds` | java.lang.Long[] | yes | List of order IDs | [123, 456] |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.ocean.openplatform.common.OpOrderInvoiceAmountResultModel](#m-alibaba-ocean-openplatform-common-oporderinvoiceamountresultmodel) | yes |  |   |

<a id="m-alibaba-ocean-openplatform-common-oporderinvoiceamountresultmodel"></a>
#### alibaba.ocean.openplatform.common.OpOrderInvoiceAmountResultModel

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `code` | java.lang.String | yes |  |   |
| `message` | java.lang.String | yes |  |   |
| `result` | [message:alibaba.ocean.openplatform.biz.trade.result.OpOrderInvoiceAmountQueryResult](#m-alibaba-ocean-openplatform-biz-trade-result-oporderinvoiceamountqueryresult) | yes |  |   |
| `retCodes` | java.lang.String[] | yes |  |   |
| `subCode` | java.lang.String | yes |  |   |
| `subMessage` | java.lang.String | yes |  |   |
| `success` | boolean | yes |  |   |

<a id="m-alibaba-ocean-openplatform-biz-trade-result-oporderinvoiceamountqueryresult"></a>
#### alibaba.ocean.openplatform.biz.trade.result.OpOrderInvoiceAmountQueryResult

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `orderInvoiceAmountModelList` | [message:alibaba.ocean.openplatform.biz.trade.common.model.OrderInvoiceAmountModel[]](#m-alibaba-ocean-openplatform-biz-trade-common-model-orderinvoiceamountmodel[]) | yes | List of order invoiceable amount models | [] |

<a id="m-alibaba-ocean-openplatform-biz-trade-common-model-orderinvoiceamountmodel[]"></a>
#### alibaba.ocean.openplatform.biz.trade.common.model.OrderInvoiceAmountModel[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `amount` | java.lang.Long | yes | Invoiceable amount, in cents | 2 |
| `orderId` | java.lang.Long | yes | Order ID | 123 |
