# Ship: no logistics needed

Original name: 物流发货-无需物流  
API: `com.alibaba.logistics:alibaba.logistics.OpDeliverySendOrder.dummy:1` · Category: Messaging  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.logistics:alibaba.logistics.OpDeliverySendOrder.dummy-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.logistics/alibaba.logistics.OpDeliverySendOrder.dummy/{appKey}`  
Requires user authorization (access_token) · Requires signature

For 1688 open-marketplace orders that need no logistics. Supports combined shipping (several orders in one shipment) and shipping at sub-order (orderEntry) level. Shipping by quantity is not supported.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `sendGoods` | [message:alibaba.logistics.OpSendGood[]](#m-alibaba-logistics-opsendgood[]) | yes | Shipment object list |   |
| `remarks` | String | no | Remark |   |
| `gmtSend` | java.util.Date | no | Shipping time | 标准时间格式：yyyyMMddHHmmssSSSZ，例如：20120801154220368+0800 |
| `extBody` | String | yes | JSON string. noLogisticsCondition in extBodyJson is required, value is a string from 1 to 5: "1": other third-party logistics, small logistics provider, fleet, etc. (noLogisticsName and noLogisticsTel required); "2": supplementary freight, price difference (noLogisticsBillNo required); "3": seller delivery (noLogisticsName and noLogisticsTel required); "4": buyer self-pickup; "5": other reasons (remarks required). Other fields have different required-ness depending on the value of noLogisticsCondition. Field descriptions: reason no logistics needed: noLogisticsCondition; name when no logistics needed: noLogisticsName; phone when no logistics needed: noLogisticsTel; waybill number when no logistics needed, meaning varies by the reason: noLogisticsBillNo; shipping proof file list: noLogisticsFiles. | {"noLogisticsBillNo":"111111111111111111","noLogisticsCondition":"3","noLogisticsName":"张三","noLogisticsTel":"13999999999","noLogisticsFiles":["https://cbu01.alicdn.com/xxx.jpg"]} |
| `extParam` | String | no | {} | JSON 字符串 |
| `receiverInfo` | [message:alibaba.logistics.OpReceiveContacter](#m-alibaba-logistics-opreceivecontacter) | no | Shipping address | 优先级大于订单收货地址，为空时，使用订单收货地址 |

<a id="m-alibaba-logistics-opsendgood[]"></a>
#### alibaba.logistics.OpSendGood[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `sourceId` | String | yes | Shipment object ID, generally the order ID |   |
| `sendGoodEntries` | [message:alibaba.logistics.OpSendGoodEntry[]](#m-alibaba-logistics-opsendgoodentry[]) | yes | List of shipment object details |    |

<a id="m-alibaba-logistics-opsendgoodentry[]"></a>
#### alibaba.logistics.OpSendGoodEntry[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `sourceEntryId` | String | yes | Shipment target line item ID, corresponding to the sub-order ID |  |
| `amount` | Long | yes | Actual shipped quantity for the shipment target |  |
| `weight` | Double | yes | The actual shipped weight of the shipment object; the weight unit defaults to kilograms. |  |
| `extBody` | String | no | JSON array string passed when shipping multiple packages. cpCode is the logistics company code, corresponding to companyNo (not the numeric Id) from the logistics company information API; logisticsCpName: logistics company name, corresponding to companyName from the logistics company information API; mailNo: waybill number; this data can be queried via the &quot;Logistics Company List - Self-linked Logistics&quot; API. quantity: number of shipment objects in the package, required. | [{"cpCode":"SF","logisticsCpName":"顺丰","mailNo":"a1231","quantity":1},{"cpCode":"SF","logisticsCpName":"顺丰","mailNo":"a1232","quantity":1}] |

<a id="m-alibaba-logistics-opreceivecontacter"></a>
#### alibaba.logistics.OpReceiveContacter

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `provinceCode` | String | yes | Province code |   |
| `cityCode` | String | yes | City code |   |
| `areaCode` | String | yes | Region code |   |
| `townCode` | String | yes | Town or street code |   |
| `province` | String | yes | Province name; can be omitted if code is passed |   |
| `city` | String | yes | City name; can be omitted if code is passed |   |
| `area` | String | yes | District name; if code is passed, this can be omitted |   |
| `town` | String | yes | Town or street name |   |
| `address` | String | yes | Detailed address |   |
| `fullName` | String | yes | Name |   |
| `corpName` | String | yes | Company name |   |
| `post` | String | yes | Postal code |   |
| `phone` | String | yes | Landline phone number |   |
| `mobile` | String | yes | Mobile phone number |   |
| `warehouse` | String | yes | Warehouse |   |
| `codeType` | String | yes | Address code type; defaults to the Cainiao standard code |   |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.logistic.result.OpSendOrderModelResult](#m-alibaba-logistic-result-opsendordermodelresult) | yes | Shipping details |   |
| `success` | Boolean | yes | Whether successful |   |
| `errorCode` | String | yes | Error code |   |
| `errorMessage` | String | yes | Error description |   |
| `extErrorMessage` | String | yes | Extended error description |   |

<a id="m-alibaba-logistic-result-opsendordermodelresult"></a>
#### alibaba.logistic.result.OpSendOrderModelResult

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `logisticsId` | String | yes | Logistics number |    |
| `sendGoods` | [message:alibaba.logistics.OpSendGood[]](#m-alibaba-logistics-opsendgood[]) | yes | Shipping details |   |

<a id="m-alibaba-logistics-opsendgood[]"></a>
#### alibaba.logistics.OpSendGood[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `sourceId` | Long | yes | Shipment object ID, generally the order ID |  |
| `sendGoodEntries` | [message:alibaba.logistics.OpSendGoodEntry[]](#m-alibaba-logistics-opsendgoodentry[]) | yes | List of shipment object details |  |

<a id="m-alibaba-logistics-opsendgoodentry[]"></a>
#### alibaba.logistics.OpSendGoodEntry[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `sourceEntryId` | Long | yes | Shipment target line item ID, corresponding to the sub-order ID |  |
| `amount` | Long | yes | Actual shipped quantity for the shipment target |  |
| `weight` | Double | yes | The actual shipped weight of the shipment object; the weight unit defaults to kilograms. |  |

## Samples

**Description of extBodyJson**

```
注意：JSON字符串,extBodyJson中的noLogisticsCondition必填，取值字符串0到5：“0”：历史无需物流的订单；“1”：其他第三方物流、小型物充商、车队等（noLogisticsName、noLogisticsTel必填）；“2”：补运费、差价（noLogisticsBillNo必填）；“3”：卖家配送（noLogisticsName、noLogisticsTel必填）；“4”：买家自提；“5”：其他原因（remarks必填）。其他字段根据noLogisticsCondition的值不同，必填要求不同。字段说明：无需物流原因:noLogisticsCondition;无需物流名称:noLogisticsName;无需物流电话:noLogisticsTel;无需物流单号，不同的无需物流原因，该字段解释不同:noLogisticsBillNo;发货凭证列表:noLogisticsFiles
```
