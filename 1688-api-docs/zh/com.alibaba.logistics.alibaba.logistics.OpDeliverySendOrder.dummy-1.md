# 物流发货-无需物流

API: `com.alibaba.logistics:alibaba.logistics.OpDeliverySendOrder.dummy:1` · Category: 消息  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.logistics:alibaba.logistics.OpDeliverySendOrder.dummy-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.logistics/alibaba.logistics.OpDeliverySendOrder.dummy/{appKey}`  
需要授权 (access_token) · 需要签名

1688大市场订单，无需物流，支持合并发货，即：多个订单一次发货；支持子订单(orderEntry)级别的发货，不支持按数量发货。

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `sendGoods` | [message:alibaba.logistics.OpSendGood[]](#m-alibaba-logistics-opsendgood[]) | 是 | 发货对象列表 |   |
| `remarks` | String | 否 | 备注 |   |
| `gmtSend` | java.util.Date | 否 | 发货时间 | 标准时间格式：yyyyMMddHHmmssSSSZ，例如：20120801154220368+0800 |
| `extBody` | String | 是 | JSON字符串,extBodyJson中的noLogisticsCondition必填，取值字符串1到5：“1”：其他第三方物流、小型物充商、车队等（noLogisticsName、noLogisticsTel必填）；“2”：补运费、差价（noLogisticsBillNo必填）；“3”：卖家配送（noLogisticsName、noLogisticsTel必填）；“4”：买家自提；“5”：其他原因（remarks必填）。其他字段根据noLogisticsCondition的值不同，必填要求不同。字段说明：无需物流原因:noLogisticsCondition;无需物流名称:noLogisticsName;无需物流电话:noLogisticsTel;无需物流单号，不同的无需物流原因，该字段解释不同:noLogisticsBillNo;发货凭证列表:noLogisticsFiles。 | {"noLogisticsBillNo":"111111111111111111","noLogisticsCondition":"3","noLogisticsName":"张三","noLogisticsTel":"13999999999","noLogisticsFiles":["https://cbu01.alicdn.com/xxx.jpg"]} |
| `extParam` | String | 否 | {} | JSON 字符串 |
| `receiverInfo` | [message:alibaba.logistics.OpReceiveContacter](#m-alibaba-logistics-opreceivecontacter) | 否 | 收货地址 | 优先级大于订单收货地址，为空时，使用订单收货地址 |

<a id="m-alibaba-logistics-opsendgood[]"></a>
#### alibaba.logistics.OpSendGood[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `sourceId` | String | 是 | 发货对象id,一般是订单id |   |
| `sendGoodEntries` | [message:alibaba.logistics.OpSendGoodEntry[]](#m-alibaba-logistics-opsendgoodentry[]) | 是 | 发货对象明细列表 |    |

<a id="m-alibaba-logistics-opsendgoodentry[]"></a>
#### alibaba.logistics.OpSendGoodEntry[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `sourceEntryId` | String | 是 | 发货对象明细id，对应子订单Id |  |
| `amount` | Long | 是 | 发货对象实发数量 |  |
| `weight` | Double | 是 | 发货对象实发重量，重量单位默认为千克 |  |
| `extBody` | String | 否 | 多包裹发货时传的JSON数组字符串，cpCode 为物流公司code，对应物流公司信息获取接口的companyNo（非数字Id）; logisticsCpName:物流公司名称，对应物流公司信息获取接口的companyName； mailNo：运单号； 可以通过获取 &quot;物流公司列表-自联物流&quot; 接口查询到数据。quantity：包裹内发货对象数量，必传。 | [{"cpCode":"SF","logisticsCpName":"顺丰","mailNo":"a1231","quantity":1},{"cpCode":"SF","logisticsCpName":"顺丰","mailNo":"a1232","quantity":1}] |

<a id="m-alibaba-logistics-opreceivecontacter"></a>
#### alibaba.logistics.OpReceiveContacter

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `provinceCode` | String | 是 | 省编码 |   |
| `cityCode` | String | 是 | 市编码 |   |
| `areaCode` | String | 是 | 地区编码 |   |
| `townCode` | String | 是 | 镇或街道编码 |   |
| `province` | String | 是 | 省名称，如果传了code，则可以不传 |   |
| `city` | String | 是 | 市名称，如果传了code，则可以不传 |   |
| `area` | String | 是 | 区名称，如果传了code，则可以不传 |   |
| `town` | String | 是 | 镇或街道名称 |   |
| `address` | String | 是 | 详细地址 |   |
| `fullName` | String | 是 | 姓名 |   |
| `corpName` | String | 是 | 公司名称 |   |
| `post` | String | 是 | 邮编 |   |
| `phone` | String | 是 | 固定电话 |   |
| `mobile` | String | 是 | 移动电话 |   |
| `warehouse` | String | 是 | 仓库 |   |
| `codeType` | String | 是 | 地址编码类型,默认菜鸟标准编码 |   |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:alibaba.logistic.result.OpSendOrderModelResult](#m-alibaba-logistic-result-opsendordermodelresult) | 是 | 发货明细 |   |
| `success` | Boolean | 是 | 是否成功 |   |
| `errorCode` | String | 是 | 错误码 |   |
| `errorMessage` | String | 是 | 错误描述 |   |
| `extErrorMessage` | String | 是 | 扩展错误描述 |   |

<a id="m-alibaba-logistic-result-opsendordermodelresult"></a>
#### alibaba.logistic.result.OpSendOrderModelResult

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `logisticsId` | String | 是 | 物流编号 |    |
| `sendGoods` | [message:alibaba.logistics.OpSendGood[]](#m-alibaba-logistics-opsendgood[]) | 是 | 发货明细 |   |

<a id="m-alibaba-logistics-opsendgood[]"></a>
#### alibaba.logistics.OpSendGood[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `sourceId` | Long | 是 | 发货对象id,一般是订单id |  |
| `sendGoodEntries` | [message:alibaba.logistics.OpSendGoodEntry[]](#m-alibaba-logistics-opsendgoodentry[]) | 是 | 发货对象明细列表 |  |

<a id="m-alibaba-logistics-opsendgoodentry[]"></a>
#### alibaba.logistics.OpSendGoodEntry[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `sourceEntryId` | Long | 是 | 发货对象明细id，对应子订单Id |  |
| `amount` | Long | 是 | 发货对象实发数量 |  |
| `weight` | Double | 是 | 发货对象实发重量，重量单位默认为千克 |  |

## 示例

**extBodyJson说明**

```
注意：JSON字符串,extBodyJson中的noLogisticsCondition必填，取值字符串0到5：“0”：历史无需物流的订单；“1”：其他第三方物流、小型物充商、车队等（noLogisticsName、noLogisticsTel必填）；“2”：补运费、差价（noLogisticsBillNo必填）；“3”：卖家配送（noLogisticsName、noLogisticsTel必填）；“4”：买家自提；“5”：其他原因（remarks必填）。其他字段根据noLogisticsCondition的值不同，必填要求不同。字段说明：无需物流原因:noLogisticsCondition;无需物流名称:noLogisticsName;无需物流电话:noLogisticsTel;无需物流单号，不同的无需物流原因，该字段解释不同:noLogisticsBillNo;发货凭证列表:noLogisticsFiles
```
