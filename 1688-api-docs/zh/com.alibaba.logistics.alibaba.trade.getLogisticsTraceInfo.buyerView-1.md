# 获取交易订单的物流跟踪信息(买家视角)

API: `com.alibaba.logistics:alibaba.trade.getLogisticsTraceInfo.buyerView:1` · Category: 订单  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.logistics:alibaba.trade.getLogisticsTraceInfo.buyerView-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.logistics/alibaba.trade.getLogisticsTraceInfo.buyerView/{appKey}`  
需要授权 (access_token) · 需要签名

该接口需要获取订单买家的授权，获取买家的订单的物流跟踪信息，在采购或者分销场景中，作为买家也有获取物流详情的需求。该接口能查能根据物流单号查看物流单跟踪信息。由于物流单录入的原因，可能跟踪信息的API查询会有延迟。该API需要向开放平台申请权限才能访问。In the procurement or distribution scenario, buyers can obtain information on logistics tracking. The interface can view the logistics tracking information according to the logistics tacking number. Depending on the logistics information entry time, there may be a delay in API queries regarding the information tracking.

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `logisticsId` | String | 否 | 该订单下的物流编号 | AL8234243 |
| `orderId` | Long | 是 | 订单号 | 13342343 |
| `webSite` | String | 是 | 是1688业务还是icbu业务 | 1688或者alibaba |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `logisticsTrace` | [message:alibaba.logistics.OpenPlatformLogisticsTrace[]](#m-alibaba-logistics-openplatformlogisticstrace[]) | 是 | 跟踪单详情 |   |
| `errorCode` | String | 是 | 错误码 |   |
| `errorMessage` | String | 是 | 错误描述 |   |
| `success` | Boolean | 是 | 是否成功 |   |
| `crossPackageFulfillmentDTO` | [message:com.alibaba.ocean.openplatform.biz.cross.model.CrossPackageFulfillmentDTO[]](#m-com-alibaba-ocean-openplatform-biz-cross-model-crosspackagefulfillmentdto[]) | 是 | 跨境海外运踪 | {} |

<a id="m-alibaba-logistics-openplatformlogisticstrace[]"></a>
#### alibaba.logistics.OpenPlatformLogisticsTrace[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `logisticsId` | String | 是 | 物流编号，如BX110096003841234 |   |
| `orderId` | Long | 是 | 订单编号 |   |
| `logisticsBillNo` | String | 是 | 物流单编号，如480330616596 |   |
| `logisticsSteps` | [message:alibaba.logistics.OpenPlatformLogisticsStep[]](#m-alibaba-logistics-openplatformlogisticsstep[]) | 是 | 物流跟踪步骤 |   |

<a id="m-alibaba-logistics-openplatformlogisticsstep[]"></a>
#### alibaba.logistics.OpenPlatformLogisticsStep[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `acceptTime` | String | 是 | 物流跟踪单该步骤的时间 |  |
| `remark` | String | 是 | 备注，如：“在浙江浦江县公司进行下级地点扫描，即将发往：广东深圳公司” |  |

<a id="m-com-alibaba-ocean-openplatform-biz-cross-model-crosspackagefulfillmentdto[]"></a>
#### com.alibaba.ocean.openplatform.biz.cross.model.CrossPackageFulfillmentDTO[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `crossMailNo` | String | 是 | 单号 | 123 |
| `orderId` | String | 是 | 订单id | 123 |
| `crossPackageTraceDTOList` | [message:com.alibaba.ocean.openplatform.biz.cross.model.CrossPackageTraceDTO[]](#m-com-alibaba-ocean-openplatform-biz-cross-model-crosspackagetracedto[]) | 是 | 运踪极点list | {} |

<a id="m-com-alibaba-ocean-openplatform-biz-cross-model-crosspackagetracedto[]"></a>
#### com.alibaba.ocean.openplatform.biz.cross.model.CrossPackageTraceDTO[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `buyerUserId` | String | 是 | 买家id | 123 |
| `carrierPartnerCode` | String | 是 | 物流公司CODE | 1 |
| `outTraceNodeCode` | String | 是 | 外部运踪Code | 23 |
| `opCode` | String | 是 | 操作code，异常节点的opCode | 123 |
| `nodeActionTime` | String | 是 | nodeActionTime | 12 |
| `nodeDetail` | String | 是 | nodeDetail | 23 |
| `stageType` | String | 是 | stageType | 1 |
| `traceNode` | [message:com.alibaba.ocean.openplatform.biz.cross.model.CrossPackageTraceDTO.TraceNode](#m-com-alibaba-ocean-openplatform-biz-cross-model-crosspackagetracedto-tracenode) | 是 | 运踪节点 | {} |
| `traceException` | [message:com.alibaba.ocean.openplatform.biz.cross.model.CrossPackageTraceDTO.ErrorTraceInfo](#m-com-alibaba-ocean-openplatform-biz-cross-model-crosspackagetracedto-errortraceinfo) | 是 | 异常节点 | {} |

<a id="m-com-alibaba-ocean-openplatform-biz-cross-model-crosspackagetracedto-tracenode"></a>
#### com.alibaba.ocean.openplatform.biz.cross.model.CrossPackageTraceDTO.TraceNode

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `traceNodeName` | String | 是 | 节点名称 | 仓库 |
| `traceNodeCode` | String | 是 | 节点 | 123 |

<a id="m-com-alibaba-ocean-openplatform-biz-cross-model-crosspackagetracedto-errortraceinfo"></a>
#### com.alibaba.ocean.openplatform.biz.cross.model.CrossPackageTraceDTO.ErrorTraceInfo

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `exceptionCode` | String | 是 | 异常code | 123 |
| `traceNodeCode` | String | 是 | 异常节点 | 1 |
| `exceptionDesc` | String | 是 | 异常 | 12 |

## 错误码

| 错误码 | 现象 | 解决方案 |
|---|---|---|
| 404 | 无法找到相对应的物流单跟踪信息。 | 无法找到相对应的物流单跟踪信息。 |
| order.nopermission.buyer | 你没有权限获取该订单详情(买家端) | 不是该授权用户的订单 |
| order.createtime.history | 不支持查询一年前的物流轨迹信息 | 不支持查询一年前的物流轨迹信息 |

## 示例

**出参示例**

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
