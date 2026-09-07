# Multilingual search navigation

Original name: 多语言搜索导航  
API: `com.alibaba.fenxiao.crossborder:product.search.keywordSNQuery:1` · Category: Products  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:product.search.keywordSNQuery-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/product.search.keywordSNQuery/{appKey}`  
Requires user authorization (access_token) · Requires signature

Get the multilingual keyword-search navigation list.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `snParams` | [message:product.search.keywordSNQuery.KeywordSNQueryParams](#m-product-search-keywordsnquery-keywordsnqueryparams) | yes |  |  |

<a id="m-product-search-keywordsnquery-keywordsnqueryparams"></a>
#### product.search.keywordSNQuery.KeywordSNQueryParams

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `keyword` | java.lang.String | yes | Search keyword | skirt |
| `language` | java.lang.String | yes | Language, e.g. English en_US; see [Developer Reference] | en_US |
| `region` | java.lang.String | yes | Region, e.g. United States US, refer to [Developer Reference] | US |
| `currency` | java.lang.String | yes | Currency, e.g. US Dollar USD, refer to [Developer Reference] | USD |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:product.search.keywordSNQuery.CommonResult](#m-product-search-keywordsnquery-commonresult) | yes | Result | 如下 |

<a id="m-product-search-keywordsnquery-commonresult"></a>
#### product.search.keywordSNQuery.CommonResult

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `success` | java.lang.Boolean | yes | Call result | true |
| `retCode` | java.lang.String | yes | Error code | S0000 |
| `retMsg` | java.lang.String | yes | Error description | 成功 |
| `result` | [message:product.search.keywordSNQuery.KeywordSNModel[]](#m-product-search-keywordsnquery-keywordsnmodel[]) | yes | Actual result | 如下 |

<a id="m-product-search-keywordsnquery-keywordsnmodel[]"></a>
#### product.search.keywordSNQuery.KeywordSNModel[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `id` | java.lang.String | yes | Search main navigation ID | 973 |
| `name` | java.lang.String | yes | Chinese name of the search main navigation | 风格 |
| `translateName` | java.lang.String | yes | Chinese name of the search main navigation | Style |
| `children` | [message:product.search.keywordSNQuery.KeywordSubSNModel[]](#m-product-search-keywordsnquery-keywordsubsnmodel[]) | yes | List of search sub-navigations | 如下 |

<a id="m-product-search-keywordsnquery-keywordsubsnmodel[]"></a>
#### product.search.keywordSNQuery.KeywordSubSNModel[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `id` | java.lang.String | yes | Search sub-navigation ID | 973:28105 |
| `name` | java.lang.String | yes | Chinese name of the search sub-navigation | 韩版 |
| `translateName` | java.lang.String | yes | Translated name of the search sub-navigation | Korean version |
