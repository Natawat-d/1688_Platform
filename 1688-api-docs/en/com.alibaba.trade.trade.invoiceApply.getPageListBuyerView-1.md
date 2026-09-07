# Query invoice applications (buyer view, paginated)

Original name: 分页查询发票申请列表（买家视角）  
API: `com.alibaba.trade:trade.invoiceApply.getPageListBuyerView:1` · Category: Invoicing  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:trade.invoiceApply.getPageListBuyerView-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/trade.invoiceApply.getPageListBuyerView/{appKey}`  
Requires user authorization (access_token) · Requires signature

Paginated query of invoice applications (buyer view).

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `input` | [message:alibaba.ocean.openplatform.biz.trade.param.OpInvoiceListByPageParam](#m-alibaba-ocean-openplatform-biz-trade-param-opinvoicelistbypageparam) | yes |  |   |

<a id="m-alibaba-ocean-openplatform-biz-trade-param-opinvoicelistbypageparam"></a>
#### alibaba.ocean.openplatform.biz.trade.param.OpInvoiceListByPageParam

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `bizStatusList` | java.lang.String[] | no | Invoice status: ISSUING(1,&quot;Issuing&quot;),ISSUED(2,&quot;Issued&quot;),VERIFYING(5,&quot;Verifying&quot;),VERIFY_FAILED(6,&quot;Verification failed&quot;),RETURNING(10,&quot;Returning&quot;),RETURNED(11,&quot;Returned&quot;),INVALIDING(20,&quot;Invalidating&quot;),DEPRECATED(21,&quot;Invalidated&quot;),RED_ISSUING(30,&quot;Red-flush in progress&quot;),RED_PART_ISSUED(31,&quot;Partially red-flushed&quot;),RED_ALL_ISSUED(32,&quot;Fully red-flushed&quot;),FAILED(40,&quot;Issuance failed&quot;),CLOSED(50,&quot;Closed&quot;),; | ISSUING |
| `fuzzyInvoiceTitle` | java.lang.String | no | Fuzzy-matched invoice title | 模糊发票抬头 |
| `isRedInvoice` | Boolean | no | Whether it is a red invoice | false |
| `orderId` | java.lang.String | no | Transaction order ID | 123 |
| `outBizId` | java.lang.String | no | Unique ID of the invoice record (used when querying application records); usually the transaction order ID, but changes for red-letter invoices | 123 |
| `page` | int | yes | Current page, can be less than 1 | 1 |
| `pageSize` | int | yes | Page size; must not be greater than 100 or less than 1 | 10 |
| `createMillTimeStart` | Long | no | Creation start time in milliseconds | 1773849600000 |
| `createMillTimeEnd` | Long | no | Creation end time in milliseconds | 1773936000000 |
| `modifyMillTimeStart` | Long | no | Modification start time in milliseconds | 1773849600000 |
| `modifyMillTimeEnd` | Long | no | Modification end time in milliseconds | 1773936000000 |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.ocean.openplatform.common.OpInvoiceModelPageResult](#m-alibaba-ocean-openplatform-common-opinvoicemodelpageresult) | yes |  |   |

<a id="m-alibaba-ocean-openplatform-common-opinvoicemodelpageresult"></a>
#### alibaba.ocean.openplatform.common.OpInvoiceModelPageResult

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `errorCode` | java.lang.String | yes |  |   |
| `errorInfo` | java.lang.String | yes |  |   |
| `pageIndex` | int | yes |  |   |
| `resultList` | [message:alibaba.ocean.openplatform.biz.trade.common.model.OpInvoiceModel[]](#m-alibaba-ocean-openplatform-biz-trade-common-model-opinvoicemodel[]) | yes |  |   |
| `sizePerPage` | int | yes |  |   |
| `success` | boolean | yes |  |   |
| `totalRecords` | int | yes |  |   |

<a id="m-alibaba-ocean-openplatform-biz-trade-common-model-opinvoicemodel[]"></a>
#### alibaba.ocean.openplatform.biz.trade.common.model.OpInvoiceModel[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `amount` | java.lang.Long | yes | Total invoice amount including tax, in cents | 2 |
| `bizStatus` | java.lang.String | yes | Invoice status: ISSUING(1,&quot;Issuing&quot;),ISSUED(2,&quot;Issued&quot;),VERIFYING(5,&quot;Verifying&quot;),VERIFY_FAILED(6,&quot;Verification failed&quot;),RETURNING(10,&quot;Returning&quot;),RETURNED(11,&quot;Returned&quot;),INVALIDING(20,&quot;Invalidating&quot;),DEPRECATED(21,&quot;Invalidated&quot;),RED_ISSUING(30,&quot;Red-flush in progress&quot;),RED_PART_ISSUED(31,&quot;Partially red-flushed&quot;),RED_ALL_ISSUED(32,&quot;Fully red-flushed&quot;),FAILED(40,&quot;Issuance failed&quot;),CLOSED(50,&quot;Closed&quot;),; | ISSUED |
| `gmtCreate` | java.util.Date | yes | Creation time |   |
| `gmtModified` | java.util.Date | yes | Modification time |   |
| `id` | java.lang.Long | yes | Primary key | 123 |
| `invoiceDeadline` | java.lang.Long | yes | Invoicing deadline, millisecond timestamp | 123 |
| `invoiceErrMsg` | java.lang.String | yes | Reason for invoicing failure | 开票失败原因 |
| `invoiceMaterial` | java.lang.String | yes | Invoice material: ELECTRON(1, &quot;electronic&quot;); | ELECTRON |
| `invoiceType` | java.lang.String | yes | Invoice type (including special invoice types): VATAX_COMM(1, &quot;VAT general invoice&quot;), VATAX_SPEC(2, &quot;VAT special invoice&quot;); | VATAX_COMM |
| `isRedInvoice` | boolean | yes | Whether it is a red invoice | false |
| `outBizId` | java.lang.String | yes | Business order number: used for interaction with upstream systems, must be unique; because a trade order may have multiple invoices, a unique outBizId must be passed in when creating an invoice. | 123 |
| `purchaserOpenUid` | String | yes | Buyer's openUid | POU |
| `purchaserInvoiceTitleModel` | [message:alibaba.ocean.openplatform.biz.trade.result.OpInvoiceTitleModel](#m-alibaba-ocean-openplatform-biz-trade-result-opinvoicetitlemodel) | yes | Buyer information | {} |
| `requestNo` | java.lang.String | yes | The request number used for interaction with downstream systems; pass in the trade order number when applying based on the trade order. |   |
| `sellerInputInvoiceTime` | java.lang.Long | yes | Time when the merchant entered the invoice | 123 |
| `sellerInvoiceTitleModel` | [message:alibaba.ocean.openplatform.biz.trade.result.OpInvoiceTitleModel](#m-alibaba-ocean-openplatform-biz-trade-result-opinvoicetitlemodel) | yes | Seller information | {} |
| `sellerMemberId` | java.lang.String | yes | Seller's member id | 销售方会员id |
| `sellerOpenUid` | String | yes | Seller's UID | SOU |
| `source` | java.lang.String | yes | Invoice source |   |
| `invoiceDetailDTOList` | [message:alibaba.ocean.openplatform.biz.trade.common.model.OpInvoiceDetailModel[]](#m-alibaba-ocean-openplatform-biz-trade-common-model-opinvoicedetailmodel[]) | yes | List of invoiced product details | [] |

<a id="m-alibaba-ocean-openplatform-biz-trade-result-opinvoicetitlemodel"></a>
#### alibaba.ocean.openplatform.biz.trade.result.OpInvoiceTitleModel

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `bankAccountId` | java.lang.String | yes | Bank account number |  银行账号 |
| `bankName` | java.lang.String | yes | Bank name |  开户行 |
| `registerAddress` | java.lang.String | yes | Enterprise registered address |  企业注册地址 |
| `registerPhone` | java.lang.String | yes | Enterprise phone |  企业电话 |
| `taxpayerIdentify` | java.lang.String | yes | Taxpayer identification number |  纳税人识别号 |
| `title` | java.lang.String | yes | Title |  抬头 |
| `titleType` | java.lang.String | yes | Invoice title type. Both individuals and social organizations are PERSONAL: PERSONAL(0, &quot;individual and social organization&quot;), COMPANY(1, &quot;enterprise&quot;); |  PERSONAL |
| `OpenUid` | String | yes | openUid |  OUD |

<a id="m-alibaba-ocean-openplatform-biz-trade-common-model-opinvoicedetailmodel[]"></a>
#### alibaba.ocean.openplatform.biz.trade.common.model.OpInvoiceDetailModel[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `subOutBizId` | java.lang.String | yes | Business sub-order number | 业务子单号 |
| `cargoName` | java.lang.String | yes | Goods/service name | 货物/服务名称 |
| `skuId` | String | yes | skuId | skuId |
| `specifications` | java.lang.String | yes | Specification/model | 规格型号 |
| `specId` | String | yes | SKU specId | 规格specId |
| `invoiceAmountWithTax` | java.lang.Long | yes | Amount including tax | 123 |
| `quantity` | java.lang.Long | yes | Quantity | 1 |
| `unit` | java.lang.String | yes | Unit | 单位 |
| `currency` | java.lang.String | yes | Currency | 币种 |
| `taxRate` | java.lang.Integer | yes | Tax rate (0-100), integer | 1 |
| `purchaserOpenUid` | String | yes | Buyer's openUid | POU |
| `sellerOpenUid` | String | yes | Seller's openUid | SOU |
| `relatedInvoiceId` | java.lang.Long | yes | Associated main invoice | 123 |
