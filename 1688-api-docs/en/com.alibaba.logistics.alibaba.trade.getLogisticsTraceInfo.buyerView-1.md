# Get order logistics tracking info (buyer view)

Original name: 获取交易订单的物流跟踪信息(买家视角)  
API: `com.alibaba.logistics:alibaba.trade.getLogisticsTraceInfo.buyerView:1` · Category: Orders  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.logistics:alibaba.trade.getLogisticsTraceInfo.buyerView-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.logistics/alibaba.trade.getLogisticsTraceInfo.buyerView/{appKey}`  
Requires user authorization (access_token) · Requires signature

Requires the order buyer's authorization and returns the logistics tracking information of the buyer's order. In procurement or distribution scenarios, buyers also need logistics details. This interface looks up tracking information by logistics (waybill) number. Because of when logistics records are entered, API tracking queries may be delayed. Permission must be requested from the Open Platform to access this API.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `logisticsId` | String | no | The logistics number under this order | AL8234243 |
| `orderId` | Long | yes | Order number | 13342343 |
| `webSite` | String | yes | Whether it is a 1688 business or an icbu business | 1688或者alibaba |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `logisticsTrace` | [message:alibaba.logistics.OpenPlatformLogisticsTrace[]](#m-alibaba-logistics-openplatformlogisticstrace[]) | yes | Tracking order details |   |
| `errorCode` | String | yes | Error code |   |
| `errorMessage` | String | yes | Error description |   |
| `success` | Boolean | yes | Whether successful |   |
| `crossPackageFulfillmentDTO` | [message:com.alibaba.ocean.openplatform.biz.cross.model.CrossPackageFulfillmentDTO[]](#m-com-alibaba-ocean-openplatform-biz-cross-model-crosspackagefulfillmentdto[]) | yes | Cross-border overseas shipment tracking | {} |

<a id="m-alibaba-logistics-openplatformlogisticstrace[]"></a>
#### alibaba.logistics.OpenPlatformLogisticsTrace[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `logisticsId` | String | yes | Logistics number, e.g. BX110096003841234 |   |
| `orderId` | Long | yes | Order number |   |
| `logisticsBillNo` | String | yes | Logistics order number, e.g. 480330616596 |   |
| `logisticsSteps` | [message:alibaba.logistics.OpenPlatformLogisticsStep[]](#m-alibaba-logistics-openplatformlogisticsstep[]) | yes | Logistics tracking step |   |

<a id="m-alibaba-logistics-openplatformlogisticsstep[]"></a>
#### alibaba.logistics.OpenPlatformLogisticsStep[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `acceptTime` | String | yes | The time of this step on the logistics tracking record |  |
| `remark` | String | yes | Remarks, e.g.: "Scanned at a sub-location by the company in Pujiang County, Zhejiang, about to be sent to: the company in Shenzhen, Guangdong" |  |

<a id="m-com-alibaba-ocean-openplatform-biz-cross-model-crosspackagefulfillmentdto[]"></a>
#### com.alibaba.ocean.openplatform.biz.cross.model.CrossPackageFulfillmentDTO[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `crossMailNo` | String | yes | Order number | 123 |
| `orderId` | String | yes | Order id | 123 |
| `crossPackageTraceDTOList` | [message:com.alibaba.ocean.openplatform.biz.cross.model.CrossPackageTraceDTO[]](#m-com-alibaba-ocean-openplatform-biz-cross-model-crosspackagetracedto[]) | yes | Shipment tracking node list | {} |

<a id="m-com-alibaba-ocean-openplatform-biz-cross-model-crosspackagetracedto[]"></a>
#### com.alibaba.ocean.openplatform.biz.cross.model.CrossPackageTraceDTO[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `buyerUserId` | String | yes | Buyer ID | 123 |
| `carrierPartnerCode` | String | yes | Logistics company CODE | 1 |
| `outTraceNodeCode` | String | yes | External shipment tracking code | 23 |
| `opCode` | String | yes | Operation code, the opCode of the exception node | 123 |
| `nodeActionTime` | String | yes | nodeActionTime | 12 |
| `nodeDetail` | String | yes | nodeDetail | 23 |
| `stageType` | String | yes | stageType | 1 |
| `traceNode` | [message:com.alibaba.ocean.openplatform.biz.cross.model.CrossPackageTraceDTO.TraceNode](#m-com-alibaba-ocean-openplatform-biz-cross-model-crosspackagetracedto-tracenode) | yes | Shipment tracking node | {} |
| `traceException` | [message:com.alibaba.ocean.openplatform.biz.cross.model.CrossPackageTraceDTO.ErrorTraceInfo](#m-com-alibaba-ocean-openplatform-biz-cross-model-crosspackagetracedto-errortraceinfo) | yes | Exception node | {} |

<a id="m-com-alibaba-ocean-openplatform-biz-cross-model-crosspackagetracedto-tracenode"></a>
#### com.alibaba.ocean.openplatform.biz.cross.model.CrossPackageTraceDTO.TraceNode

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `traceNodeName` | String | yes | Node name | 仓库 |
| `traceNodeCode` | String | yes | Node | 123 |

<a id="m-com-alibaba-ocean-openplatform-biz-cross-model-crosspackagetracedto-errortraceinfo"></a>
#### com.alibaba.ocean.openplatform.biz.cross.model.CrossPackageTraceDTO.ErrorTraceInfo

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `exceptionCode` | String | yes | Exception code | 123 |
| `traceNodeCode` | String | yes | Exception node | 1 |
| `exceptionDesc` | String | yes | Exception | 12 |

## Error codes

| Error | Symptom | How to fix |
|---|---|---|
| 404 | Unable to find the corresponding logistics tracking information. | Unable to find the corresponding logistics tracking information. |
| order.nopermission.buyer | You do not have permission to obtain this order's details (buyer side) | Not an order belonging to this authorized user |
| order.createtime.history | Does not support querying logistics tracking info from more than a year ago | Does not support querying logistics tracking info from more than a year ago |

## Samples

**Output parameter example**

```
{
    "logisticsTrace":[
        {
            "logisticsId":"LP00106397027178",
            "logisticsBillNo":"3832890717253",
            "orderId":188983797838441800,
            "logisticsSteps":[
                {
                    "acceptTime":"2018-07-24 21:55:33",
                    "remark":"在广东广州天河区天平架一公司进行揽件扫描"
                },
                {
                    "acceptTime":"2018-07-24 22:10:50",
                    "remark":"在广东广州天河区天平架一公司进行下级地点扫描，即将发往：浙江宁波分拨中心"
                },
                {
                    "acceptTime":"2018-07-25 01:45:05",
                    "remark":"在分拨中心广东广州分拨中心进行称重扫描"
                },
                {
                    "acceptTime":"2018-07-25 01:47:42",
                    "remark":"在广东广州分拨中心进行装车扫描，即将发往：浙江宁波分拨中心"
                },
                {
                    "acceptTime":"2018-07-26 03:01:41",
                    "remark":"在分拨中心浙江宁波分拨中心进行卸车扫描"
                },
                {
                    "acceptTime":"2018-07-26 03:21:34",
                    "remark":"从浙江宁波分拨中心发出，本次转运目的地：浙江宁波鄞州区邱隘公司"
                },
                {
                    "acceptTime":"2018-07-26 07:06:21",
                    "remark":"到达目的地网点浙江宁波鄞州区邱隘公司，快件将很快进行派送"
                },
                {
                    "acceptTime":"2018-07-26 08:54:15",
                    "remark":"在浙江宁波鄞州区邱隘公司进行派件扫描；派送业务员：徐洲；联系电话：xxxxx"
                },
                {
                    "acceptTime":"2018-07-26 13:33:11",
                    "remark":"快件已被 已签收 签收"
                }
            ]
        }
    ]
}
```
