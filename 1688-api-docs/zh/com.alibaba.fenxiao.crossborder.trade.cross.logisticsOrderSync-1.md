# 国家站物流单回传

API: `com.alibaba.fenxiao.crossborder:trade.cross.logisticsOrderSync:1` · Category: 回传数据  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:trade.cross.logisticsOrderSync-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/trade.cross.logisticsOrderSync/{appKey}`  
需要授权 (access_token) · 需要签名

国家站物流单回传

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `orderLogisticsParam` | [message:alibaba.cbu.order.out.param.OrderLogisticsParam](#m-alibaba-cbu-order-out-param-orderlogisticsparam) | 是 |  |  |

<a id="m-alibaba-cbu-order-out-param-orderlogisticsparam"></a>
#### alibaba.cbu.order.out.param.OrderLogisticsParam

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `orderId` | java.lang.String | 是 | 国家站主单id | 1213 |
| `logisticsId` | java.lang.String | 是 | 下游物流单id | 1·21213 |
| `logisticsStatus` | java.lang.String | 是 | 物流状态send-已发货 sign-签收 partSign-部分收货 refund-退货 | sign |
| `createTime` | java.util.Date | 是 | 创建时间 | 2023-10-10 10:12:12 |
| `signTime` | java.util.Date | 是 | 签收时间 | 2023-10-10 10:12:12 |
| `province` | java.lang.String | 是 | 省 | 江西省 |
| `city` | java.lang.String | 是 | 市 | 南昌市 |
| `area` | java.lang.String | 是 | 区 | 白云区 |
| `address` | java.lang.String | 是 | 详细地址 | 江西省南昌市白云区xx号 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:com.alibaba.openapi.shared.common.ResultModel](#m-com-alibaba-openapi-shared-common-resultmodel) | 是 | 返回信息 |  |

<a id="m-com-alibaba-openapi-shared-common-resultmodel"></a>
#### com.alibaba.openapi.shared.common.ResultModel

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `success` | Boolean | 是 | 是否成功 | true |
| `code` | String | 是 | code | 200 |
| `message` | String | 是 | 返回信息 | 成功 |
