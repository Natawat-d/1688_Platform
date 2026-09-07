# Write back country-site logistics orders

Original name: 国家站物流单回传  
API: `com.alibaba.fenxiao.crossborder:trade.cross.logisticsOrderSync:1` · Category: Data Write-back  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:trade.cross.logisticsOrderSync-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/trade.cross.logisticsOrderSync/{appKey}`  
Requires user authorization (access_token) · Requires signature

Write back logistics orders from the country site.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `orderLogisticsParam` | [message:alibaba.cbu.order.out.param.OrderLogisticsParam](#m-alibaba-cbu-order-out-param-orderlogisticsparam) | yes |  |  |

<a id="m-alibaba-cbu-order-out-param-orderlogisticsparam"></a>
#### alibaba.cbu.order.out.param.OrderLogisticsParam

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `orderId` | java.lang.String | yes | Country site main order ID | 1213 |
| `logisticsId` | java.lang.String | yes | Downstream logistics order ID | 1·21213 |
| `logisticsStatus` | java.lang.String | yes | Logistics status: send - shipped, sign - signed for, partSign - partially received, refund - returned | sign |
| `createTime` | java.util.Date | yes | Creation time | 2023-10-10 10:12:12 |
| `signTime` | java.util.Date | yes | Signed receipt time | 2023-10-10 10:12:12 |
| `province` | java.lang.String | yes | Province | 江西省 |
| `city` | java.lang.String | yes | City | 南昌市 |
| `area` | java.lang.String | yes | District | 白云区 |
| `address` | java.lang.String | yes | Detailed address | 江西省南昌市白云区xx号 |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:com.alibaba.openapi.shared.common.ResultModel](#m-com-alibaba-openapi-shared-common-resultmodel) | yes | Return message |  |

<a id="m-com-alibaba-openapi-shared-common-resultmodel"></a>
#### com.alibaba.openapi.shared.common.ResultModel

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `success` | Boolean | yes | Whether successful | true |
| `code` | String | yes | code | 200 |
| `message` | String | yes | Return message | 成功 |
