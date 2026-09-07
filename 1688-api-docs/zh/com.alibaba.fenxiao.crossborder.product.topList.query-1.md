# 查询榜单列表

API: `com.alibaba.fenxiao.crossborder:product.topList.query:1` · Category: 商机  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:product.topList.query-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/product.topList.query/{appKey}`  
需要授权 (access_token) · 需要签名

查询榜单列表

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `rankQueryParams` | [message:product.topList.query.RankQueryParams](#m-product-toplist-query-rankqueryparams) | 是 |  |  |

<a id="m-product-toplist-query-rankqueryparams"></a>
#### product.topList.query.RankQueryParams

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `rankId` | java.lang.String | 是 | 榜单ID，可传入类目ID，目前支持类目榜单 | 1111 |
| `rankType` | java.lang.String | 是 | 榜单类型，complex综合榜，hot热卖榜，goodPrice好价榜 | complex |
| `limit` | java.lang.Integer | 是 | 榜单商品个数，最多20 | 10 |
| `language` | java.lang.String | 是 | 榜单商品语言 | en |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:product.topList.query.ResultModel](#m-product-toplist-query-resultmodel) | 是 |  |  |

<a id="m-product-toplist-query-resultmodel"></a>
#### product.topList.query.ResultModel

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `success` | java.lang.Boolean | 是 | 是否成功 | true |
| `code` | java.lang.String | 是 | 错误码 | S0000 |
| `message` | java.lang.String | 是 | 错误描述 | 成功 |
| `result` | [message:product.topList.query.RankModel](#m-product-toplist-query-rankmodel) | 是 | 结果 | 结果 |

<a id="m-product-toplist-query-rankmodel"></a>
#### product.topList.query.RankModel

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `rankId` | java.lang.String | 是 | 榜单ID | 111 |
| `rankName` | java.lang.String | 是 | 榜单名称 | Comprehensive List |
| `rankType` | java.lang.String | 是 | 榜单类型 | complex |
| `rankProductModels` | [message:product.topList.query.RankProductModel[]](#m-product-toplist-query-rankproductmodel[]) | 是 | 榜单结果 | 如下 |

<a id="m-product-toplist-query-rankproductmodel[]"></a>
#### product.topList.query.RankProductModel[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `itemId` | java.lang.Long | 是 | 商品ID | 699490252651 |
| `title` | String | 是 | 商品中文标题 | 2023厚底女士凉鞋软底拖鞋夏季现货踩屎感注塑鞋出口家居防滑凉鞋 |
| `translateTitle` | String | 是 | 商品译文标题 | 2023 thick-soled ladies sandals soft-soled slippers summer spot poop injection shoes export home non-slip sandals |
| `imgUrl` | java.lang.String | 是 | 商品图片 | http://img.china.alibaba.com/img/ibank/O1CN01p4SIPo1D6Wx4c0xs8_!!2201053890167-0-cib.search.jpg |
| `sort` | java.lang.Integer | 是 | 商品排行 | 1 |
| `serviceList` | java.lang.String[] | 是 | 商品包含的服务，24小时发货sendGoods24H，48小时发货sendGoods48H | ["sendGoods48H"] |
| `buyerNum` | java.lang.Integer | 是 | 最近30天买家数 | 1334 |
| `soldOut` | java.lang.Integer | 是 | 最近30天商品售卖件数 | 433454 |
| `goodsScore` | String | 是 | 商品交易评分 | 5 |
