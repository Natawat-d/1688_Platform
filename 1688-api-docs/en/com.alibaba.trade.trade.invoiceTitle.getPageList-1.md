# Query buyer invoice titles (paginated)

Original name: 分页查询买家抬头列表  
API: `com.alibaba.trade:trade.invoiceTitle.getPageList:1` · Category: Invoicing  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:trade.invoiceTitle.getPageList-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/trade.invoiceTitle.getPageList/{appKey}`  
Requires user authorization (access_token) · Requires signature

Paginated query of the buyer's invoice titles.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `input` | [message:alibaba.ocean.openplatform.biz.trade.param.OpInvoiceTitleQueryParam](#m-alibaba-ocean-openplatform-biz-trade-param-opinvoicetitlequeryparam) | yes |  |   |

<a id="m-alibaba-ocean-openplatform-biz-trade-param-opinvoicetitlequeryparam"></a>
#### alibaba.ocean.openplatform.biz.trade.param.OpInvoiceTitleQueryParam

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `page` | int | no | Current page, cannot be less than 1 | 1 |
| `pageSize` | int | no | Page size; cannot be less than 1 or greater than 100 | 50 |
| `titleType` | java.lang.String | no | Invoice title type. Both individual and social organization map to PERSONAL: PERSONAL(0,&quot;individual&quot;),COMPANY(1,&quot;company&quot;),SOCIAL_ORGANIZATION(3,&quot;social organization (government agency, public institution, etc.)&quot;),PARENT_VIRTUAL(-1,&quot;parent invoice virtual type&quot;); | PERSONAL |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.ocean.openplatform.common.OpInvoiceTitlePageSingleResult](#m-alibaba-ocean-openplatform-common-opinvoicetitlepagesingleresult) | yes |  |   |

<a id="m-alibaba-ocean-openplatform-common-opinvoicetitlepagesingleresult"></a>
#### alibaba.ocean.openplatform.common.OpInvoiceTitlePageSingleResult

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `currentPage` | java.lang.Integer | yes | Current page number | 1 |
| `errorCode` | java.lang.String | yes | Error code | 404 |
| `errorInfo` | java.lang.String | yes | Error message | 错误信息 |
| `pageSize` | java.lang.Integer | yes | Page size | 10 |
| `result` | [message:alibaba.ocean.openplatform.biz.trade.result.OpInvoiceTitleQueryResult](#m-alibaba-ocean-openplatform-biz-trade-result-opinvoicetitlequeryresult) | yes | Return result | {} |
| `success` | boolean | yes | Whether successful | true |
| `totalNum` | java.lang.Long | yes | Total record count | 100 |

<a id="m-alibaba-ocean-openplatform-biz-trade-result-opinvoicetitlequeryresult"></a>
#### alibaba.ocean.openplatform.biz.trade.result.OpInvoiceTitleQueryResult

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `companyInvoiceTitles` | [message:alibaba.ocean.openplatform.biz.trade.result.OpInvoiceTitleModel[]](#m-alibaba-ocean-openplatform-biz-trade-result-opinvoicetitlemodel[]) | yes | Enterprise invoice title | [] |
| `personalInvoiceTitles` | [message:alibaba.ocean.openplatform.biz.trade.result.OpInvoiceTitleModel[]](#m-alibaba-ocean-openplatform-biz-trade-result-opinvoicetitlemodel[]) | yes | Personal invoice title | [] |

<a id="m-alibaba-ocean-openplatform-biz-trade-result-opinvoicetitlemodel[]"></a>
#### alibaba.ocean.openplatform.biz.trade.result.OpInvoiceTitleModel[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `bankAccountId` | java.lang.String | yes | Bank account number | 123 |
| `bankName` | java.lang.String | yes | Bank name | 开户行 |
| `email` | java.lang.String | yes | Email | 邮箱 |
| `gmtCreate` | java.util.Date | yes | Creation time |   |
| `gmtModified` | java.util.Date | yes | Modification time |   |
| `invoiceTitleId` | java.lang.Long | yes | Invoice title primary key id | 123 |
| `isDefault` | java.lang.String | yes | Whether it is default | true |
| `name` | java.lang.String | yes | Enterprise contact person | 企业联系人 |
| `receiverPhone` | java.lang.String | yes | Recipient phone number | 123 |
| `registerAddress` | java.lang.String | yes | Enterprise registered address | 企业注册地址 |
| `registerPhone` | java.lang.String | yes | Enterprise phone | 123 |
| `taxpayerIdentify` | java.lang.String | yes | Taxpayer identification number | 纳税人识别号 |
| `title` | java.lang.String | yes | Title | 抬头 |
| `titleType` | java.lang.String | yes | Invoice title type. Both individual and social organization map to PERSONAL: PERSONAL(0,&quot;individual&quot;),COMPANY(1,&quot;company&quot;),SOCIAL_ORGANIZATION(3,&quot;social organization (government agency, public institution, etc.)&quot;),PARENT_VIRTUAL(-1,&quot;parent invoice virtual type&quot;); | PERSONAL |
| `openUid` | String | yes | openUid | OU |
