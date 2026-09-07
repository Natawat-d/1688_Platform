# Buyer requests invoice

Original name: 买家申请开票  
API: `com.alibaba.trade:trade.invoice.apply:1` · Category: Invoicing  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:trade.invoice.apply-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/trade.invoice.apply/{appKey}`  
Requires user authorization (access_token) · Requires signature

Buyer requests an invoice; batch requests are supported. Prerequisites: get the buyer's invoice titles; get the invoiceable orders and the invoice types they support; get the orders' invoiceable amounts.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `input` | [message:alibaba.ocean.openplatform.biz.trade.param.OpApplyInvoiceParam](#m-alibaba-ocean-openplatform-biz-trade-param-opapplyinvoiceparam) | yes | Input parameter | {} |

<a id="m-alibaba-ocean-openplatform-biz-trade-param-opapplyinvoiceparam"></a>
#### alibaba.ocean.openplatform.biz.trade.param.OpApplyInvoiceParam

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `invoiceApplyModelList` | [message:alibaba.ocean.openplatform.biz.trade.common.model.OpInvoiceApplyModel[]](#m-alibaba-ocean-openplatform-biz-trade-common-model-opinvoiceapplymodel[]) | yes | Invoice application model | {} |

<a id="m-alibaba-ocean-openplatform-biz-trade-common-model-opinvoiceapplymodel[]"></a>
#### alibaba.ocean.openplatform.biz.trade.common.model.OpInvoiceApplyModel[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `orderId` | java.lang.Long | yes | Main order number | 123 |
| `amount` | java.lang.Long | yes | Total invoice amount including tax, obtained via the trade.invoiceAmount.getList API | 2 |
| `invoiceType` | java.lang.String | yes | Invoice type: VATAX_SPEC(2, &quot;VAT special invoice&quot;); VATAX_COMM(1, &quot;VAT general invoice&quot;). | VATAX_COMM |
| `purchaserInvoiceTitleModel` | [message:alibaba.ocean.openplatform.biz.trade.result.OpInvoiceTitleModel](#m-alibaba-ocean-openplatform-biz-trade-result-opinvoicetitlemodel) | yes | Buyer's invoice title, obtained via the trade.invoiceTitle.getPageList API; pass the retrieved title through as-is. | {} |

<a id="m-alibaba-ocean-openplatform-biz-trade-result-opinvoicetitlemodel"></a>
#### alibaba.ocean.openplatform.biz.trade.result.OpInvoiceTitleModel

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `titleType` | java.lang.String | yes | Invoice title type. Both individuals and social organizations use PERSONAL: PERSONAL(0,&quot;Individual and social organization&quot;),COMPANY(1,&quot;Enterprise&quot;); | PERSONAL |
| `title` | java.lang.String | yes | Title | 抬头 |
| `taxpayerIdentify` | java.lang.String | no | Taxpayer identification number (required for enterprise invoicing) | 123 |
| `bankName` | java.lang.String | no | Bank name (required for enterprise special invoice) | 开户行 |
| `bankAccountId` | java.lang.String | no | Bank account number (required for enterprises issuing special VAT invoices) | 123 |
| `registerAddress` | java.lang.String | no | Enterprise registered address (required for enterprises issuing special VAT invoices) | 企业注册地址 |
| `registerPhone` | java.lang.String | no | Company phone number (required for enterprise special invoice) | 123 |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:ApplyInvoiceResultModel](#m-applyinvoiceresultmodel) | yes | Request result | {} |

<a id="m-applyinvoiceresultmodel"></a>
#### ApplyInvoiceResultModel

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `code` | String | yes | code |   |
| `message` | String | yes | message |   |
| `result` | [message:alibaba.ocean.openplatform.biz.trade.result.OpApplyInvoiceResult](#m-alibaba-ocean-openplatform-biz-trade-result-opapplyinvoiceresult) | yes | Invoicing result | {} |
| `retCodes` | String[] | yes | retCodes | [] |
| `subCode` | String | yes | subCode |   |
| `subMessage` | String | yes | subMessage |   |
| `success` | Boolean | yes | Whether successful | true |

<a id="m-alibaba-ocean-openplatform-biz-trade-result-opapplyinvoiceresult"></a>
#### alibaba.ocean.openplatform.biz.trade.result.OpApplyInvoiceResult

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `failedList` | [message:alibaba.ocean.openplatform.biz.trade.common.model.OpInvoicingResultModel[]](#m-alibaba-ocean-openplatform-biz-trade-common-model-opinvoicingresultmodel[]) | no | List of invoicing failures | [] |
| `successList` | [message:alibaba.ocean.openplatform.biz.trade.common.model.OpInvoicingResultModel[]](#m-alibaba-ocean-openplatform-biz-trade-common-model-opinvoicingresultmodel[]) | no | Successful invoicing list | [] |

<a id="m-alibaba-ocean-openplatform-biz-trade-common-model-opinvoicingresultmodel[]"></a>
#### alibaba.ocean.openplatform.biz.trade.common.model.OpInvoicingResultModel[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `outBizId` | java.lang.String | yes | Main order id | 123 |
| `result` | boolean | yes | Invoicing result | true |
| `orderId` | Long | yes | Main order id | 123 |
| `tradeOrderCompleted` | Boolean | yes | Whether the transaction is completed | true |
| `errorCode` | String | yes | Error code (filled in on failure) | XXX |
| `errorDesc` | String | yes | Error description (filled in on failure) | XXX |

## Error codes

| Error | Symptom | How to fix |
|---|---|---|
| KINGSUNS_BASIC_PARAM_CHECK_ERROR | Basic parameter check exception: the input parameter is null, or the input buyer id is empty. |  |
| PARAM_INVALID | Parameter exception; invoice title parameter exception |  |
| ORDER_CANCELED_ERROR | The order has been cancelled; invoicing is not allowed |  |
| INVOICE_ALREADY_APPLIED | An invoice has already been previously applied for via another channel |  |
| INVOICE_NOT_SUPPORT_SPEC | This order's product does not support issuing a special VAT invoice at payment time; please contact the merchant to enable the \&quot;special VAT invoice\&quot; invoicing setting&quot; |  |
| INVOICE_NOT_SUPPORT_COMMON | This order's product does not support issuing a general VAT invoice at payment time; please contact the merchant to enable the \&quot;store-wide products\&quot; invoicing setting&quot; |  |
| APPLY_CREATE_INVOICE_EXCEPTION | Exception applying to create invoice; fallback exception |  |
