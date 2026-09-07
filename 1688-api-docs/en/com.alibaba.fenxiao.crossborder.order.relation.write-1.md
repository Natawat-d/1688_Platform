# Write back mapping between end-customer orders and 1688 orders

Original name: 回传机构真实用户订单和1688订单的映射关系  
API: `com.alibaba.fenxiao.crossborder:order.relation.write:1` · Category: Data Write-back  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:order.relation.write-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/order.relation.write/{appKey}`  
Requires user authorization (access_token) · Requires signature

Write back to 1688 the mapping between the organisation's real end-customer orders and the 1688 orders generated under the organisation's account.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `orderRelationParam` | [message:alibaba.cbu.order.param.OrderRelationParam](#m-alibaba-cbu-order-param-orderrelationparam) | yes | Order relationship parameter | 如下 |

<a id="m-alibaba-cbu-order-param-orderrelationparam"></a>
#### alibaba.cbu.order.param.OrderRelationParam

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `orderId` | java.lang.String | yes | External institution sub-order ID | 4590347523948375 |
| `parentOrderId` | java.lang.String | yes | External institution's main order ID | 4590347523948370 |
| `purchaseOrderId` | java.lang.Long | yes | 1688 purchase sub-order ID | 3209572465452734 |
| `purchaseParentOrderId` | java.lang.Long | yes | 1688 purchase sub-order ID | 3209572465452730 |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.openapi.shared.common.ResultModel](#m-alibaba-openapi-shared-common-resultmodel) | yes | Return result | 如下 |

<a id="m-alibaba-openapi-shared-common-resultmodel"></a>
#### alibaba.openapi.shared.common.ResultModel

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `success` | boolean | yes | Whether successful | true |
| `code` | java.lang.String | yes | Error code | 0 |
| `message` | java.lang.String | yes | Error description | 无 |

## Samples

**Input parameter**

```
{
    "orderId": "234325234234542324",
    "parentOrderId": "234325234234542320",
    "purchaseOrderId": "2352424230582342934",
    "purchaseParentOrderId": "2352424230582342930"
}
```

**Result**

```
{
    "success": true,
    "code": "0",
    "message": ""
}
```
