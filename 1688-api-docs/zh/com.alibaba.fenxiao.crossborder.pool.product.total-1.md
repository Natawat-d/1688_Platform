# 查询商品池中商品总数

API: `com.alibaba.fenxiao.crossborder:pool.product.total:1` · Category: 工具  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:pool.product.total-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/pool.product.total/{appKey}`  
无需授权 · 需要签名

查询商品池中商品总数

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 否 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `appKey` | java.lang.String | 是 | appKey | 123 |
| `palletId` | java.lang.Long | 是 | palletId | 123 |
| `categoryId` | java.lang.String | 否 | categoryId | 1 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:alibaba.national.pool.product.total.new.ResultModel](#m-alibaba-national-pool-product-total-new-resultmodel) | 是 | 结果 | {} |

<a id="m-alibaba-national-pool-product-total-new-resultmodel"></a>
#### alibaba.national.pool.product.total.new.ResultModel

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `msg` | String | 是 | msg | 1 |
| `code` | String | 是 | code | 001 |
| `traceId` | String | 是 | traceId | 12121212333443212sada |
| `success` | String | 是 | 请求是否成功标志 | true |
| `model` | Long | 是 | 总数 | 100 |
