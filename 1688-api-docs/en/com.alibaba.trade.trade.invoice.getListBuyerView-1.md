# Query all issued invoices for an order (buyer view)

Original name: 查询交易单下关联的所有开具的发票信息（买家视角）  
API: `com.alibaba.trade:trade.invoice.getListBuyerView:1` · Category: Invoicing  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:trade.invoice.getListBuyerView-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/trade.invoice.getListBuyerView/{appKey}`  
Requires user authorization (access_token) · Requires signature

Query all issued invoices associated with an order (buyer view).

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `input` | [message:alibaba.ocean.openplatform.biz.trade.param.OpInvoiceByOrderIdQueryParam](#m-alibaba-ocean-openplatform-biz-trade-param-opinvoicebyorderidqueryparam) | yes |  |   |

<a id="m-alibaba-ocean-openplatform-biz-trade-param-opinvoicebyorderidqueryparam"></a>
#### alibaba.ocean.openplatform.biz.trade.param.OpInvoiceByOrderIdQueryParam

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `orderId` | Long | yes | Order ID | 123 |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.ocean.openplatform.common.OpInvoiceModelListResultModel](#m-alibaba-ocean-openplatform-common-opinvoicemodellistresultmodel) | yes |  |   |

<a id="m-alibaba-ocean-openplatform-common-opinvoicemodellistresultmodel"></a>
#### alibaba.ocean.openplatform.common.OpInvoiceModelListResultModel

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `code` | String | yes |  |   |
| `message` | String | yes |  |   |
| `result` | [message:alibaba.ocean.openplatform.biz.trade.common.model.OpInvoiceModel[]](#m-alibaba-ocean-openplatform-biz-trade-common-model-opinvoicemodel[]) | yes |  |   |
| `retCodes` | String[] | yes |  |   |
| `subCode` | String | yes |  |   |
| `subMessage` | String | yes |  |   |
| `success` | Boolean | yes |  |   |

<a id="m-alibaba-ocean-openplatform-biz-trade-common-model-opinvoicemodel[]"></a>
#### alibaba.ocean.openplatform.biz.trade.common.model.OpInvoiceModel[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `amount` | java.lang.Long | yes | Total invoice amount including tax, in cents | 2 |
| `bizStatus` | java.lang.String | yes | Invoice status: ISSUING(1,&quot;Issuing&quot;),ISSUED(2,&quot;Issued&quot;),VERIFYING(5,&quot;Verifying&quot;),VERIFY_FAILED(6,&quot;Verification failed&quot;),RETURNING(10,&quot;Returning&quot;),RETURNED(11,&quot;Returned&quot;),INVALIDING(20,&quot;Invalidating&quot;),DEPRECATED(21,&quot;Invalidated&quot;),RED_ISSUING(30,&quot;Red-flush in progress&quot;),RED_PART_ISSUED(31,&quot;Partially red-flushed&quot;),RED_ALL_ISSUED(32,&quot;Fully red-flushed&quot;),FAILED(40,&quot;Issuance failed&quot;),CLOSED(50,&quot;Closed&quot;),; | ISSUED |
| `gmtCreate` | java.util.Date | yes | Creation time |   |
| `gmtModified` | java.util.Date | yes | Modification time |   |
| `id` | java.lang.Long | yes | Primary key | 123 |
| `invoiceDate` | java.lang.Long | yes | The actual invoicing time of the invoice, as distinct from the submission time, invoice application time, etc. |   |
| `invoiceDeadline` | java.lang.Long | yes | Invoicing deadline, millisecond timestamp | 123 |
| `invoiceDownloadUrl` | java.lang.String | yes | Invoice download link | 发票下载链接 |
| `invoiceErrMsg` | java.lang.String | yes | Reason for invoicing failure | 开票失败原因 |
| `invoiceFileType` | java.lang.String | yes | Invoice file type: PDF(&quot;PDF&quot;,&quot;PDF format&quot;),OFD(&quot;OFD&quot;,&quot;OFD format&quot;),GIF(&quot;GIF&quot;,&quot;GIF format&quot;),TIF(&quot;TIF&quot;,&quot;TIF format&quot;),BMP(&quot;BMP&quot;,&quot;BMP format&quot;),JPG(&quot;JPG&quot;,&quot;JPG format&quot;),XML(&quot;XML&quot;,&quot;XML format&quot;),PNG(&quot;PNG&quot;,&quot;PNG format&quot;); | PDF |
| `invoiceMaterial` | java.lang.String | yes | Invoice material: ELECTRON(1, &quot;electronic&quot;); | ELECTRON |
| `invoiceNo` | java.lang.String | yes | Invoice number | 发票号 |
| `invoiceType` | java.lang.String | yes | Invoice type (including special invoice types): VATAX_COMM(1, &quot;VAT general invoice&quot;), VATAX_SPEC(2, &quot;VAT special invoice&quot;); | VATAX_COMM |
| `isRedInvoice` | boolean | yes | Whether it is a red invoice | false |
| `ossInvoiceFileName` | java.lang.String | yes | File name | 文件名称 |
| `outBizId` | java.lang.String | yes | Business order number: used for interaction with upstream systems, must be unique; because a trade order may have multiple invoices, a unique outBizId must be passed in when creating an invoice. | 123 |
| `purchaserOpenUid` | String | yes | Buyer's openUid | POU |
| `purchaserInvoiceTitleModel` | [message:alibaba.ocean.openplatform.biz.trade.result.OpInvoiceTitleModel](#m-alibaba-ocean-openplatform-biz-trade-result-opinvoicetitlemodel) | yes | Buyer information | {} |
| `requestNo` | java.lang.String | yes | The request number used for interaction with downstream systems; pass in the trade order number when applying based on the trade order. |   |
| `sellerInputInvoiceTime` | java.lang.Long | yes | Time when the merchant entered the invoice | 123 |
| `sellerInvoiceTitleModel` | [message:alibaba.ocean.openplatform.biz.trade.result.OpInvoiceTitleModel](#m-alibaba-ocean-openplatform-biz-trade-result-opinvoicetitlemodel) | yes | Seller information | {} |
| `sellerMemberId` | java.lang.String | yes | Seller's member id | 销售方会员id |
| `sellerOpenUid` | String | yes | Seller's UID | SOU |
| `source` | java.lang.String | yes | Invoice source |   |

<a id="m-alibaba-ocean-openplatform-biz-trade-result-opinvoicetitlemodel"></a>
#### alibaba.ocean.openplatform.biz.trade.result.OpInvoiceTitleModel

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `titleType` | java.lang.String | yes | Invoice title type. Both individuals and social organizations use PERSONAL: PERSONAL(0,&quot;Individual and social organization&quot;),COMPANY(1,&quot;Enterprise&quot;); | PERSONAL |
| `title` | java.lang.String | yes | Title | 抬头 |
| `taxpayerIdentify` | java.lang.String | yes | Taxpayer identification number (required for enterprise invoicing) | 123 |
| `bankName` | java.lang.String | yes | Bank name (required for enterprise special invoice) | 开户行 |
| `bankAccountId` | java.lang.String | yes | Bank account number (required for enterprises issuing special VAT invoices) | 123 |
| `registerAddress` | java.lang.String | yes | Enterprise registered address (required for enterprises issuing special VAT invoices) | 企业注册地址 |
| `registerPhone` | java.lang.String | yes | Company phone number (required for enterprise special invoice) | 123 |
| `OpenUid` | String | yes | openUid |  OUD |
