# 多语言搜索导航

API: `com.alibaba.fenxiao.crossborder:product.search.keywordSNQuery:1` · Category: 商品  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:product.search.keywordSNQuery-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/product.search.keywordSNQuery/{appKey}`  
需要授权 (access_token) · 需要签名

获取多语言关键词搜索导航列表

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `snParams` | [message:product.search.keywordSNQuery.KeywordSNQueryParams](#m-product-search-keywordsnquery-keywordsnqueryparams) | 是 |  |  |

<a id="m-product-search-keywordsnquery-keywordsnqueryparams"></a>
#### product.search.keywordSNQuery.KeywordSNQueryParams

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `keyword` | java.lang.String | 是 | 搜索关键词 | skirt |
| `language` | java.lang.String | 是 | 语言，如英语en_US，参考【开发人员参考】 | en_US |
| `region` | java.lang.String | 是 | 地区，如美国US，参考【开发人员参考】 | US |
| `currency` | java.lang.String | 是 | 币种，如美元USD，参考【开发人员参考】 | USD |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:product.search.keywordSNQuery.CommonResult](#m-product-search-keywordsnquery-commonresult) | 是 | 结果 | 如下 |

<a id="m-product-search-keywordsnquery-commonresult"></a>
#### product.search.keywordSNQuery.CommonResult

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `success` | java.lang.Boolean | 是 | 调用结果 | true |
| `retCode` | java.lang.String | 是 | 错误码 | S0000 |
| `retMsg` | java.lang.String | 是 | 错误描述 | 成功 |
| `result` | [message:product.search.keywordSNQuery.KeywordSNModel[]](#m-product-search-keywordsnquery-keywordsnmodel[]) | 是 | 实际结果 | 如下 |

<a id="m-product-search-keywordsnquery-keywordsnmodel[]"></a>
#### product.search.keywordSNQuery.KeywordSNModel[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `id` | java.lang.String | 是 | 搜索主导航ID | 973 |
| `name` | java.lang.String | 是 | 搜索主导航中文名称 | 风格 |
| `translateName` | java.lang.String | 是 | 搜索主导航中文名称 | Style |
| `children` | [message:product.search.keywordSNQuery.KeywordSubSNModel[]](#m-product-search-keywordsnquery-keywordsubsnmodel[]) | 是 | 搜索子导航列表 | 如下 |

<a id="m-product-search-keywordsnquery-keywordsubsnmodel[]"></a>
#### product.search.keywordSNQuery.KeywordSubSNModel[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `id` | java.lang.String | 是 | 搜索子导航ID | 973:28105 |
| `name` | java.lang.String | 是 | 搜索子导航中文名称 | 韩版 |
| `translateName` | java.lang.String | 是 | 搜索子导航译文名称 | Korean version |
