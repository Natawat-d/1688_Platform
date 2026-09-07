# 相关性商品推荐

API: `com.alibaba.fenxiao.crossborder:product.related.recommend:1` · Category: 商品  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:product.related.recommend-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/product.related.recommend/{appKey}`  
需要授权 (access_token) · 需要签名

相关性商品推荐

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `relatedQueryParams` | [message:product.related.recommend.RelatedQueryParams](#m-product-related-recommend-relatedqueryparams) | 是 | 入参 | 如下 |

<a id="m-product-related-recommend-relatedqueryparams"></a>
#### product.related.recommend.RelatedQueryParams

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `offerId` | java.lang.Long | 是 | 商品ID | 689337673960 |
| `pageNo` | java.lang.Integer | 是 | 页号 | 1 |
| `pageSize` | java.lang.Integer | 是 | 每页数量，最大10 | 10 |
| `language` | java.lang.String | 是 | 语言 | en |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:product.related.recommend.ResultModel](#m-product-related-recommend-resultmodel) | 是 | 结果 | 如下 |

<a id="m-product-related-recommend-resultmodel"></a>
#### product.related.recommend.ResultModel

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `success` | boolean | 是 | 是否成功 | true |
| `code` | java.lang.String | 是 | 错误码 | S0000 |
| `message` | java.lang.String | 是 | 错误描述 | 成功 |
| `result` | [message:product.related.recommend.ProductInfoModel[]](#m-product-related-recommend-productinfomodel[]) | 是 | 结果 | 如下 |

<a id="m-product-related-recommend-productinfomodel[]"></a>
#### product.related.recommend.ProductInfoModel[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `imageUrl` | java.lang.String | 是 | 图片链接 | https://cbu01.alicdn.com/img/ibank/O1CN01IEn8mi1i9xzXcd2GU_!!2212446184371-0-cib.jpg |
| `subject` | java.lang.String | 是 | 中文标题 | 跨境自制果蔬全自动面膜机diy英文版带语音智能美容仪面膜机批发 |
| `subjectTrans` | java.lang.String | 是 | 译文标题 | Cross-border homemade fruit and vegetable automatic mask machine diy English version with voice intelligent beauty instrument mask machine wholesale |
| `priceInfo` | [message:product.related.recommend.PriceInfo](#m-product-related-recommend-priceinfo) | 是 | 价格对象 | 如下 |
| `offerId` | java.lang.Long | 是 | 商品ID | 688766724221 |
| `monthSold` | java.lang.Integer | 是 | 最近90天销售件数 | 98374 |
| `sellerIdentities` | java.lang.String[] | 是 | 商家标 | super_factory |
| `offerIdentities` | java.lang.String[] | 是 | 商品标 | yx |
| `topCategoryId` | java.lang.Long | 是 | 一级类目ID | 1 |
| `secondCategoryId` | java.lang.Long | 是 | 二级类目ID | 1 |
| `thirdCategoryId` | java.lang.Long | 是 | 三级类目ID | 1 |

<a id="m-product-related-recommend-priceinfo"></a>
#### product.related.recommend.PriceInfo

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `price` | java.lang.String | 是 | 商品批发价格最低价 | 178.0 |
