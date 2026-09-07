# 仓库商品库存更新

API: `com.alibaba.fenxiao.crossborder:consignment.co.inventory:1` · Category: 全托管  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:consignment.co.inventory-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/consignment.co.inventory/{appKey}`  
无需授权 · 需要签名

仓库商品库存更新

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 否 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `bizSkuId` | java.lang.String | 是 | 外部skuId | 12212122 |
| `sourceItemId` | java.lang.String | 是 | 仓库货品id | 12312323123 |
| `outboundQuantity` | java.lang.String | 是 | 出库数量 | 12 |
| `bizId` | java.lang.String | 是 | 业务id，用来保证库存扣减幂等 | 12321312 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:alibaba.national.common.result.ResultModel](#m-alibaba-national-common-result-resultmodel) | 是 |  |  |

<a id="m-alibaba-national-common-result-resultmodel"></a>
#### alibaba.national.common.result.ResultModel

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `msg` | java.lang.String | 是 | 信息 | 成功 |
| `traceId` | java.lang.String | 是 | 链路id | 12312312313 |
| `success` | java.lang.Boolean | 是 | 成功/失败 | true |
| `model` | java.lang.Object | 是 | 返回结果 | null |
