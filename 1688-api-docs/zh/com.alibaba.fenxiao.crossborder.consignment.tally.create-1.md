# 仓库创建差异理货单

API: `com.alibaba.fenxiao.crossborder:consignment.tally.create:1` · Category: 全托管  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:consignment.tally.create-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/consignment.tally.create/{appKey}`  
无需授权 · 需要签名

Buffalo仓库创建差异理货单

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 否 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `bizInboundOrderId` | java.lang.String | 是 | 入库单id |  |
| `status` | java.lang.String | 是 | tallysheet |  |
| `processTime` | java.lang.Long | 是 | 时间戳 |  |
| `tallySheetId` | java.lang.String | 是 | 仓库差异单id |  |
| `tallySheetReason` | java.lang.String | 是 | 理货差异原因 |  |
| `inboundCount` | java.lang.String | 是 | 入库数量 |  |
| `varianceCount` | java.lang.String | 是 | 差异数量 |  |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:alibaba.national.common.result.ResultModel](#m-alibaba-national-common-result-resultmodel) | 是 |  |  |

<a id="m-alibaba-national-common-result-resultmodel"></a>
#### alibaba.national.common.result.ResultModel

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `msg` | java.lang.String | 是 | 错误信息 | 参数为空 |
| `traceId` | java.lang.String | 是 | traceId | 111111 |
| `success` | java.lang.String | 是 | true/false | true |
| `model` | java.lang.Object | 是 | null | null |
