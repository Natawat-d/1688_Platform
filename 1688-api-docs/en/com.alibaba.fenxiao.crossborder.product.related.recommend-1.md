# Related product recommendations

Original name: 相关性商品推荐  
API: `com.alibaba.fenxiao.crossborder:product.related.recommend:1` · Category: Products  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:product.related.recommend-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/product.related.recommend/{appKey}`  
Requires user authorization (access_token) · Requires signature

Related product recommendations.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `relatedQueryParams` | [message:product.related.recommend.RelatedQueryParams](#m-product-related-recommend-relatedqueryparams) | yes | Input parameter | 如下 |

<a id="m-product-related-recommend-relatedqueryparams"></a>
#### product.related.recommend.RelatedQueryParams

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `offerId` | java.lang.Long | yes | Product ID | 689337673960 |
| `pageNo` | java.lang.Integer | yes | Page number | 1 |
| `pageSize` | java.lang.Integer | yes | Number per page, maximum 10 | 10 |
| `language` | java.lang.String | yes | Language | en |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:product.related.recommend.ResultModel](#m-product-related-recommend-resultmodel) | yes | Result | 如下 |

<a id="m-product-related-recommend-resultmodel"></a>
#### product.related.recommend.ResultModel

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `success` | boolean | yes | Whether successful | true |
| `code` | java.lang.String | yes | Error code | S0000 |
| `message` | java.lang.String | yes | Error description | 成功 |
| `result` | [message:product.related.recommend.ProductInfoModel[]](#m-product-related-recommend-productinfomodel[]) | yes | Result | 如下 |

<a id="m-product-related-recommend-productinfomodel[]"></a>
#### product.related.recommend.ProductInfoModel[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `imageUrl` | java.lang.String | yes | Image link | https://cbu01.alicdn.com/img/ibank/O1CN01IEn8mi1i9xzXcd2GU_!!2212446184371-0-cib.jpg |
| `subject` | java.lang.String | yes | Chinese title | 跨境自制果蔬全自动面膜机diy英文版带语音智能美容仪面膜机批发 |
| `subjectTrans` | java.lang.String | yes | Translated title | Cross-border homemade fruit and vegetable automatic mask machine diy English version with voice intelligent beauty instrument mask machine wholesale |
| `priceInfo` | [message:product.related.recommend.PriceInfo](#m-product-related-recommend-priceinfo) | yes | Price object | 如下 |
| `offerId` | java.lang.Long | yes | Product ID | 688766724221 |
| `monthSold` | java.lang.Integer | yes | Number of items sold in the last 90 days | 98374 |
| `sellerIdentities` | java.lang.String[] | yes | Merchant mark | super_factory |
| `offerIdentities` | java.lang.String[] | yes | Product tag | yx |
| `topCategoryId` | java.lang.Long | yes | Top-level (first-level) category ID | 1 |
| `secondCategoryId` | java.lang.Long | yes | Second-level category ID | 1 |
| `thirdCategoryId` | java.lang.Long | yes | Level-3 category ID | 1 |

<a id="m-product-related-recommend-priceinfo"></a>
#### product.related.recommend.PriceInfo

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `price` | java.lang.String | yes | Lowest wholesale price of the product | 178.0 |
