# 商品热搜词

API: `com.alibaba.fenxiao.crossborder:product.search.topKeyword:1` · Category: 商机  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:product.search.topKeyword-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/product.search.topKeyword/{appKey}`  
无需授权 · 需要签名

商品热搜词

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 否 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `topSeKeywordParam` | [message:alibaba.cbu.offer.param.TopSeKeywordParam](#m-alibaba-cbu-offer-param-topsekeywordparam) | 是 |  |  |

<a id="m-alibaba-cbu-offer-param-topsekeywordparam"></a>
#### alibaba.cbu.offer.param.TopSeKeywordParam

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `country` | java.lang.String | 是 | 语言，参考开发参考枚举 | en-英语 |
| `sourceId` | java.lang.String | 是 | 查询id，如类目id | 1-类目id |
| `hotKeywordType` | java.lang.String | 是 | 热搜类型，目前只提供类目纬度，此固定传cate | cate-类目 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:alibaba.openapi.shared.common.ResultModel](#m-alibaba-openapi-shared-common-resultmodel) | 是 |  |  |

<a id="m-alibaba-openapi-shared-common-resultmodel"></a>
#### alibaba.openapi.shared.common.ResultModel

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `success` | java.lang.String | 是 |  |  |
| `code` | java.lang.String | 是 |  |  |
| `message` | java.lang.String | 是 |  |  |
| `result` | [message:alibaba.cbu.offer.model.TopSeKeywordModel[]](#m-alibaba-cbu-offer-model-topsekeywordmodel[]) | 是 |  |  |

<a id="m-alibaba-cbu-offer-model-topsekeywordmodel[]"></a>
#### alibaba.cbu.offer.model.TopSeKeywordModel[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `seKeyword` | java.lang.String | 是 | 中文热词，请用词进行词搜 | 大型宠物犬 |
| `seKeywordTranslation` | java.lang.String | 是 | 译文热词，展示使用 | Large Pet Dog |
