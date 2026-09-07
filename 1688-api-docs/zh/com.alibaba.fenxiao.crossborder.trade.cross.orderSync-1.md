# 下游销售订单同步

API: `com.alibaba.fenxiao.crossborder:trade.cross.orderSync:1` · Category: 回传数据  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:trade.cross.orderSync-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/trade.cross.orderSync/{appKey}`  
需要授权 (access_token) · 需要签名

下游销售订单同步

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `orderParam` | [message:com.alibaba.cbu.order.param.OrderParamV](#m-com-alibaba-cbu-order-param-orderparamv) | 是 | 订单详情 |  |

<a id="m-com-alibaba-cbu-order-param-orderparamv"></a>
#### com.alibaba.cbu.order.param.OrderParamV

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `orderId` | String | 是 | 订单id | 1234 |
| `productId` | String | 否 | 商品id，子单或主子和一订单必填 | 12 |
| `productName` | String | 否 | 商品名称，子单或主子和一订单必填 | 手套 |
| `skuId` | String | 否 | skuId，子单或主子和一订单必填 | adsds1213 |
| `skuName` | String | 否 | sku名称，子单或主子和一订单必填 | 绿色 |
| `buyAmount` | Long | 否 | 购买数量 | 2 |
| `createTime` | String | 是 | 创建时间 | 2023-10-01 10:10:10 |
| `outMemberId` | String | 是 | 下游用户id | 1212113 |
| `payTime` | String | 是 | 支付时间 | 2023-10-01 10:10:10 |
| `endTime` | String | 否 | 完成时间 | 2023-10-01 10:10:10 |
| `paidFee` | Long | 是 | 实付金额 分 | 200 |
| `refundStatus` | String | 否 | 退款状态 0-未退款（退款关闭，未申请） 1-退款 | 0 |
| `status` | String | 是 | payed-已支付  success-交易成功 close-交易关闭 | success |
| `subOrderParamList` | [message:com.alibaba.cbu.order.param.SubOrderParam[]](#m-com-alibaba-cbu-order-param-suborderparam[]) | 否 | 一主多子需传子单信息 | 子单信息 |

<a id="m-com-alibaba-cbu-order-param-suborderparam[]"></a>
#### com.alibaba.cbu.order.param.SubOrderParam[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `orderId` | String | 是 | 子单id | 2 |
| `productId` | String | 是 | 1688商品id需明文 | 1 |
| `productName` | String | 是 | 1688商品名称需明文 | 袜子 |
| `skuId` | String | 是 | 1688skuId需明文 | 1 |
| `skuName` | String | 是 | 1688sku名称需明文 | 绿色 |
| `buyAmount` | Long | 是 | 子单购买数量 | 1 |
| `paidFee` | Long | 是 | 实付金额 单位分 | 100 |
| `refundStatus` | String | 是 | 退款状态 0-未退款或退款关闭 1-退款 | 0 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:alibaba.openapi.shared.out.common.ResultModel](#m-alibaba-openapi-shared-out-common-resultmodel) | 是 |  |  |

<a id="m-alibaba-openapi-shared-out-common-resultmodel"></a>
#### alibaba.openapi.shared.out.common.ResultModel

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `success` | boolean | 是 |  |  |
| `code` | java.lang.String | 是 |  |  |
| `message` | java.lang.String | 是 |  |  |
