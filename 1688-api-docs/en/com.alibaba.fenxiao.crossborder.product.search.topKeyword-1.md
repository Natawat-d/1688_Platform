# Trending product search keywords

Original name: 商品热搜词  
API: `com.alibaba.fenxiao.crossborder:product.search.topKeyword:1` · Category: Market Insights  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:product.search.topKeyword-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/product.search.topKeyword/{appKey}`  
No user authorization · Requires signature

Trending product search keywords.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | no | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `topSeKeywordParam` | [message:alibaba.cbu.offer.param.TopSeKeywordParam](#m-alibaba-cbu-offer-param-topsekeywordparam) | yes |  |  |

<a id="m-alibaba-cbu-offer-param-topsekeywordparam"></a>
#### alibaba.cbu.offer.param.TopSeKeywordParam

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `country` | java.lang.String | yes | Language; refer to the developer reference enum | en-英语 |
| `sourceId` | java.lang.String | yes | Query ID, e.g. category ID | 1-类目id |
| `hotKeywordType` | java.lang.String | yes | Hot search type; currently only the category dimension is provided, pass the fixed value cate | cate-类目 |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.openapi.shared.common.ResultModel](#m-alibaba-openapi-shared-common-resultmodel) | yes |  |  |

<a id="m-alibaba-openapi-shared-common-resultmodel"></a>
#### alibaba.openapi.shared.common.ResultModel

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `success` | java.lang.String | yes |  |  |
| `code` | java.lang.String | yes |  |  |
| `message` | java.lang.String | yes |  |  |
| `result` | [message:alibaba.cbu.offer.model.TopSeKeywordModel[]](#m-alibaba-cbu-offer-model-topsekeywordmodel[]) | yes |  |  |

<a id="m-alibaba-cbu-offer-model-topsekeywordmodel[]"></a>
#### alibaba.cbu.offer.model.TopSeKeywordModel[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `seKeyword` | java.lang.String | yes | Chinese hot keyword; use this term for keyword search. | 大型宠物犬 |
| `seKeywordTranslation` | java.lang.String | yes | Translated hot keyword, for display use | Large Pet Dog |
