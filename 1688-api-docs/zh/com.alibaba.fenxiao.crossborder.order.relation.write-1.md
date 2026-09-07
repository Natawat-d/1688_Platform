# 回传机构真实用户订单和1688订单的映射关系

API: `com.alibaba.fenxiao.crossborder:order.relation.write:1` · Category: 回传数据  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:order.relation.write-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/order.relation.write/{appKey}`  
需要授权 (access_token) · 需要签名

将机构真实用户的订单和机构账号产生的1688订单的映射关系回传1688。

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `orderRelationParam` | [message:alibaba.cbu.order.param.OrderRelationParam](#m-alibaba-cbu-order-param-orderrelationparam) | 是 | 订单关系参数 | 如下 |

<a id="m-alibaba-cbu-order-param-orderrelationparam"></a>
#### alibaba.cbu.order.param.OrderRelationParam

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `orderId` | java.lang.String | 是 | 外部机构子订单ID | 4590347523948375 |
| `parentOrderId` | java.lang.String | 是 | 外部机构主订单ID | 4590347523948370 |
| `purchaseOrderId` | java.lang.Long | 是 | 1688采购子订单ID | 3209572465452734 |
| `purchaseParentOrderId` | java.lang.Long | 是 | 1688采购子订单ID | 3209572465452730 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:alibaba.openapi.shared.common.ResultModel](#m-alibaba-openapi-shared-common-resultmodel) | 是 | 返回结果 | 如下 |

<a id="m-alibaba-openapi-shared-common-resultmodel"></a>
#### alibaba.openapi.shared.common.ResultModel

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `success` | boolean | 是 | 是否成功 | true |
| `code` | java.lang.String | 是 | 错误码 | 0 |
| `message` | java.lang.String | 是 | 错误描述 | 无 |

## 示例

**入参**

```
{
    "orderId": "234325234234542324",
    "parentOrderId": "234325234234542320",
    "purchaseOrderId": "2352424230582342934",
    "purchaseParentOrderId": "2352424230582342930"
}
```

**结果**

```
{
    "success": true,
    "code": "0",
    "message": ""
}
```
