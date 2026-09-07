# 物流发货-自己联系物流发货

API: `com.alibaba.logistics:alibaba.logistics.OpDeliverySendOrder.offline:1` · Category: 消息  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.logistics:alibaba.logistics.OpDeliverySendOrder.offline-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.logistics/alibaba.logistics.OpDeliverySendOrder.offline/{appKey}`  
需要授权 (access_token) · 需要签名

1688大市场订单，卖家自己联系物流发货，支持合并发货，即：多个订单一次发货；支持子订单(orderEntry)级别的发货，不支持按数量发货。

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `multiPackage` | Boolean | 否 | 是否使用多包裹发货 | true：使用多包裹发货，false或不传：使用单包裹发货 |
| `sendGoods` | [message:alibaba.logistics.OpSendGood[]](#m-alibaba-logistics-opsendgood[]) | 是 | 发货对象列表 |   |
| `remarks` | String | 否 | 备注 |   |
| `gmtSend` | java.util.Date | 否 | 发货时间 | 标准时间格式：yyyyMMddHHmmssSSSZ，例如：20120801154220368+0800 |
| `extBody` | String | 否 | 单包裹发货时传的JSON字符串，cpCode 为物流公司code，对应物流公司信息获取接口的companyNo（非数字Id）; logisticsCpName:物流公司名称，对应物流公司信息获取接口的companyName； mailNo：运单号； 可以通过获取 &quot;物流公司列表-自联物流&quot; 接口查询到数据。 | {"cpCode":"SF","logisticsCpName":"顺丰","mailNo":"a123"} |
| `extParam` | String | 否 | JSON 字符串 |   |
| `receiverInfo` | [message:alibaba.logistics.OpReceiveContacter](#m-alibaba-logistics-opreceivecontacter) | 否 | 收货地址,优先级大于订单收货地址，为空时，使用订单收货地址 |   |
| `isEncryptOrderSend` | String | 否 | 是否下游加密订单取号发货，当卖家是通过下游平台加密取号时传入Y | Y,N |

<a id="m-alibaba-logistics-opsendgood[]"></a>
#### alibaba.logistics.OpSendGood[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `sourceId` | String | 是 | 发货对象id,一般是订单id |  |
| `sendGoodEntries` | [message:alibaba.logistics.OpSendGoodEntry[]](#m-alibaba-logistics-opsendgoodentry[]) | 是 | 发货对象明细列表 |  |

<a id="m-alibaba-logistics-opsendgoodentry[]"></a>
#### alibaba.logistics.OpSendGoodEntry[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `sourceEntryId` | String | 是 | 发货对象明细id，对应子订单Id |   |
| `amount` | Long | 是 | 发货对象实发数量 |   |
| `weight` | Double | 是 | 发货对象实发重量，重量单位默认为千克 |   |
| `extBody` | String | 否 | 多包裹发货时传的JSON数组字符串，cpCode 为物流公司code，对应物流公司信息获取接口的companyNo（非数字Id）; logisticsCpName:物流公司名称，对应物流公司信息获取接口的companyName； mailNo：运单号； 可以通过获取 &quot;物流公司列表-自联物流&quot; 接口查询到数据。quantity：包裹内发货对象数量，必传。 | [{"cpCode":"SF","logisticsCpName":"顺丰","mailNo":"a1231","quantity":1},{"cpCode":"SF","logisticsCpName":"顺丰","mailNo":"a1232","quantity":1}] |

<a id="m-alibaba-logistics-opreceivecontacter"></a>
#### alibaba.logistics.OpReceiveContacter

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `provinceCode` | String | 是 | 省编码 |  |
| `cityCode` | String | 是 | 市编码 |  |
| `areaCode` | String | 是 | 地区编码 |  |
| `townCode` | String | 是 | 镇或街道编码 |  |
| `province` | String | 是 | 省名称，如果传了code，则可以不传 |  |
| `city` | String | 是 | 市名称，如果传了code，则可以不传 |  |
| `area` | String | 是 | 区名称，如果传了code，则可以不传 |  |
| `town` | String | 是 | 镇或街道名称 |  |
| `address` | String | 是 | 详细地址 |  |
| `fullName` | String | 是 | 姓名 |  |
| `corpName` | String | 是 | 公司名称 |  |
| `post` | String | 是 | 邮编 |  |
| `phone` | String | 是 | 固定电话 |  |
| `mobile` | String | 是 | 移动电话 |  |
| `warehouse` | String | 是 | 仓库 |  |
| `codeType` | String | 是 | 地址编码类型,默认菜鸟标准编码 |  |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:alibaba.logistic.result.OpSendOrderModelResult](#m-alibaba-logistic-result-opsendordermodelresult) | 是 | 发货明细 |   |
| `success` | Boolean | 是 | 是否成功 |   |
| `errorCode` | String | 是 | 错误码 |   |
| `errorMessage` | String | 是 | 错误描述 |   |
| `extErrorMessage` | String | 是 | 扩展错误描述 |    |

<a id="m-alibaba-logistic-result-opsendordermodelresult"></a>
#### alibaba.logistic.result.OpSendOrderModelResult

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `logisticsId` | String | 是 | 物流编号 |   |
| `sendGoods` | [message:alibaba.logistics.OpSendGood[]](#m-alibaba-logistics-opsendgood[]) | 是 | 发货明细 |   |
| `sendSuccessList` | [message:com.alibaba.ocean.openplatform.biz.logistics.result.OpSendOrderResult[]](#m-com-alibaba-ocean-openplatform-biz-logistics-result-opsendorderresult[]) | 是 | 多包裹发货成功的包裹列表 | [{"sourceId":1,"sourceEntryId":11,"extBody":"{\"cpCode\":\"SF\",\"logisticsCpName\":\"顺丰\",\"mailNo\":\"a1231\",\"quantity\":1}","success":true,"logisticsId":"123"}] |
| `sendFailList` | [message:com.alibaba.ocean.openplatform.biz.logistics.result.OpSendOrderResult[]](#m-com-alibaba-ocean-openplatform-biz-logistics-result-opsendorderresult[]) | 是 | 多包裹发货失败的包裹列表 | [{"sourceId":1,"sourceEntryId":11,"extBody":"{\"cpCode\":\"SF\",\"logisticsCpName\":\"顺丰\",\"mailNo\":\"a1231\",\"quantity\":1}","success":false,"errorMessage":"错误原因"}] |

<a id="m-alibaba-logistics-opsendgood[]"></a>
#### alibaba.logistics.OpSendGood[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `sourceId` | String | 是 | 发货对象id,一般是订单id |   |
| `sendGoodEntries` | [message:alibaba.logistics.OpSendGoodEntry[]](#m-alibaba-logistics-opsendgoodentry[]) | 是 | 发货对象明细列表 |   |

<a id="m-alibaba-logistics-opsendgoodentry[]"></a>
#### alibaba.logistics.OpSendGoodEntry[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `sourceEntryId` | String | 是 | 发货对象明细id，对应子订单Id |   |
| `amount` | Long | 是 | 发货对象实发数量 |   |
| `weight` | Double | 是 | 发货对象实发重量，重量单位默认为千克 |   |
| `extBody` | String | 是 | 多包裹发货时传的JSON数组字符串，cpCode 为物流公司code，对应物流公司信息获取接口的companyNo（非数字Id）; logisticsCpName:物流公司名称，对应物流公司信息获取接口的companyName； mailNo：运单号； 可以通过获取 &quot;物流公司列表-自联物流&quot; 接口查询到数据。quantity：包裹内发货对象数量，必传。 | [{"cpCode":"SF","logisticsCpName":"顺丰","mailNo":"a1231","quantity":1},{"cpCode":"SF","logisticsCpName":"顺丰","mailNo":"a1232","quantity":1}] |

<a id="m-com-alibaba-ocean-openplatform-biz-logistics-result-opsendorderresult[]"></a>
#### com.alibaba.ocean.openplatform.biz.logistics.result.OpSendOrderResult[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `sourceId` | String | 是 | 主单id | 1 |
| `sourceEntryId` | String | 是 | 子单id | 11 |
| `extBody` | String | 是 | JSON字符串（对应入参多包裹中的一个） | {"cpCode":"SF","logisticsCpName":"顺丰","mailNo":"a1231","quantity":1} |
| `success` | Boolean | 是 | 是否成功 | true |
| `logisticsId` | String | 是 | 物流单号 | 123 |
| `errorMessage` | String | 是 | 错误信息 | aaa |

## 错误码

| 错误码 | 现象 | 解决方案 |
|---|---|---|
| 系统繁忙，请稍后再试Case1 | 调用接口返回“系统繁忙，请稍后再试” | 请检查订单号是否正确，另外请确认物流公司CpCode是否正确。cpcode为物流公司信息的companyNo而非Id，例如顺丰的coCode为“SF”而非106422 |
| 系统繁忙，请稍后再试Case2 | 使用测试工具时返回“系统繁忙，请稍后再试” | 不要传递receiverInfo字段 |
| HSF Server unexpected exception | HSF Server unexpected exception | 传入参数的字段类型不正确，或某些必传字段没有传递 |
| INVALID_PARAM | 运单号不符合规则或已被使用 | 请检查运单号是否正确，另外请注意一个运单号最多使用15次 |
| 5001 | 下游销售订单退款中，暂不支持发货，建议与您的客户沟通处理退款，再确定是否继续发货 | 下游销售订单退款中，暂不支持发货，建议与您的客户沟通处理退款，再确定是否继续发货，继续发货需要到1688work工作台上发货 |
| 5002 | 下游销售订单已退款成功/交易关闭，不支持发货 | 下游销售订单已退款成功/交易关闭，不支持发货 |

## 示例

**出参示例**

```
{
    "result":{
        "logisticsId":"ZX113988430670174",
        "sendGoods":[
            {
                "sourceId":206026099675498520,
                "sendGoodEntries":[
                    {
                        "sourceEntryId":206026099675498520,
                        "amount":80,
                        "weight":10
                    }
                ]
            }
        ]
    },
    "success":true
}
```
