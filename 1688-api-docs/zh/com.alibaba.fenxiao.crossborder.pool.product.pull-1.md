# 拉取商品池中商品数据

API: `com.alibaba.fenxiao.crossborder:pool.product.pull:1` · Category: 工具  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:pool.product.pull-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/pool.product.pull/{appKey}`  
需要授权 (access_token) · 需要签名

通过商品池ID直接批量拉取池中商品数据

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `offerPoolQueryParam` | [message:pool.product.pull.OfferPoolQueryParam](#m-pool-product-pull-offerpoolqueryparam) | 是 |  |  |

<a id="m-pool-product-pull-offerpoolqueryparam"></a>
#### pool.product.pull.OfferPoolQueryParam

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `offerPoolId` | java.lang.Long | 是 | 品池id（业务定制且有权限控制，从对接的业务获取，随便传会报错，寻源通代采建议走词搜接口） | 111 |
| `cateId` | java.lang.Long | 否 | 类目ID | 11 |
| `taskId` | java.lang.String | 是 | 查询任务ID，假如货盘有10000商品，每页1000个查询10次将10000商品查走，这10次都需要传同一个taskId，机构需要在分页查询的时候固定一个taskId留存下来，然后每次分页都传同一个taskId来查该接口 | 1 |
| `language` | java.lang.String | 否 | 语言 | en |
| `pageNo` | java.lang.Integer | 是 | 页码 | 1 |
| `pageSize` | java.lang.Integer | 是 | 每页数量 | 10 |
| `sortField` | String | 否 | 排序字段 | order1m/buyer1m（order1m：最近1个月销售额排序；buyer1m：最近1个月买家数） |
| `sortType` | String | 否 | 排序规则 | ASC/DESC |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:pool.product.pull.ResultModel](#m-pool-product-pull-resultmodel) | 是 |  |  |

<a id="m-pool-product-pull-resultmodel"></a>
#### pool.product.pull.ResultModel

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `success` | java.lang.String | 是 | 是否成功 | true |
| `code` | java.lang.String | 是 | 错误码 | S0000 |
| `message` | java.lang.String | 是 | 错误描述 | 成功 |
| `result` | [message:pool.product.pull.ProductPoolModel[]](#m-pool-product-pull-productpoolmodel[]) | 是 | 结果 | 结果 |

<a id="m-pool-product-pull-productpoolmodel[]"></a>
#### pool.product.pull.ProductPoolModel[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `offerId` | java.lang.Long | 是 | 商品ID | 111111 |
| `bizCategoryId` | java.lang.String | 是 | 机构的类目ID | 111111 |
| `offerPoolTotal` | Integer | 是 | 商品池总数(每个offer都返回) | 122211 |
