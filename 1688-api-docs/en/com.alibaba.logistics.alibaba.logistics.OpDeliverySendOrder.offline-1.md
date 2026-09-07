# Ship: seller arranges own logistics

Original name: 物流发货-自己联系物流发货  
API: `com.alibaba.logistics:alibaba.logistics.OpDeliverySendOrder.offline:1` · Category: Messaging  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.logistics:alibaba.logistics.OpDeliverySendOrder.offline-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.logistics/alibaba.logistics.OpDeliverySendOrder.offline/{appKey}`  
Requires user authorization (access_token) · Requires signature

For 1688 open-marketplace orders where the seller arranges logistics themselves. Supports combined shipping (several orders in one shipment) and shipping at sub-order (orderEntry) level. Shipping by quantity is not supported.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `multiPackage` | Boolean | no | Whether to use multi-package shipment | true：使用多包裹发货，false或不传：使用单包裹发货 |
| `sendGoods` | [message:alibaba.logistics.OpSendGood[]](#m-alibaba-logistics-opsendgood[]) | yes | Shipment object list |   |
| `remarks` | String | no | Remark |   |
| `gmtSend` | java.util.Date | no | Shipping time | 标准时间格式：yyyyMMddHHmmssSSSZ，例如：20120801154220368+0800 |
| `extBody` | String | no | JSON string passed for single-package shipment. cpCode is the logistics company code, corresponding to companyNo (not the numeric Id) from the logistics company info API; logisticsCpName: logistics company name, corresponding to companyName from the logistics company info API; mailNo: waybill number. This data can be queried via the &quot;Logistics Company List - Self-Connected Logistics&quot; API. | {"cpCode":"SF","logisticsCpName":"顺丰","mailNo":"a123"} |
| `extParam` | String | no | JSON string |   |
| `receiverInfo` | [message:alibaba.logistics.OpReceiveContacter](#m-alibaba-logistics-opreceivecontacter) | no | Shipping address; takes priority over the order's shipping address. If empty, the order's shipping address is used |   |
| `isEncryptOrderSend` | String | no | Whether to ship using a downstream encrypted order number; pass Y when the seller obtains the number via downstream platform encryption. | Y,N |

<a id="m-alibaba-logistics-opsendgood[]"></a>
#### alibaba.logistics.OpSendGood[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `sourceId` | String | yes | Shipment object ID, generally the order ID |  |
| `sendGoodEntries` | [message:alibaba.logistics.OpSendGoodEntry[]](#m-alibaba-logistics-opsendgoodentry[]) | yes | List of shipment object details |  |

<a id="m-alibaba-logistics-opsendgoodentry[]"></a>
#### alibaba.logistics.OpSendGoodEntry[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `sourceEntryId` | String | yes | Shipment target line item ID, corresponding to the sub-order ID |   |
| `amount` | Long | yes | Actual shipped quantity for the shipment target |   |
| `weight` | Double | yes | The actual shipped weight of the shipment object; the weight unit defaults to kilograms. |   |
| `extBody` | String | no | JSON array string passed when shipping multiple packages. cpCode is the logistics company code, corresponding to companyNo (not the numeric Id) from the logistics company information API; logisticsCpName: logistics company name, corresponding to companyName from the logistics company information API; mailNo: waybill number; this data can be queried via the &quot;Logistics Company List - Self-linked Logistics&quot; API. quantity: number of shipment objects in the package, required. | [{"cpCode":"SF","logisticsCpName":"顺丰","mailNo":"a1231","quantity":1},{"cpCode":"SF","logisticsCpName":"顺丰","mailNo":"a1232","quantity":1}] |

<a id="m-alibaba-logistics-opreceivecontacter"></a>
#### alibaba.logistics.OpReceiveContacter

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `provinceCode` | String | yes | Province code |  |
| `cityCode` | String | yes | City code |  |
| `areaCode` | String | yes | Region code |  |
| `townCode` | String | yes | Town or street code |  |
| `province` | String | yes | Province name; can be omitted if code is passed |  |
| `city` | String | yes | City name; can be omitted if code is passed |  |
| `area` | String | yes | District name; if code is passed, this can be omitted |  |
| `town` | String | yes | Town or street name |  |
| `address` | String | yes | Detailed address |  |
| `fullName` | String | yes | Name |  |
| `corpName` | String | yes | Company name |  |
| `post` | String | yes | Postal code |  |
| `phone` | String | yes | Landline phone number |  |
| `mobile` | String | yes | Mobile phone number |  |
| `warehouse` | String | yes | Warehouse |  |
| `codeType` | String | yes | Address code type; defaults to the Cainiao standard code |  |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.logistic.result.OpSendOrderModelResult](#m-alibaba-logistic-result-opsendordermodelresult) | yes | Shipping details |   |
| `success` | Boolean | yes | Whether successful |   |
| `errorCode` | String | yes | Error code |   |
| `errorMessage` | String | yes | Error description |   |
| `extErrorMessage` | String | yes | Extended error description |    |

<a id="m-alibaba-logistic-result-opsendordermodelresult"></a>
#### alibaba.logistic.result.OpSendOrderModelResult

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `logisticsId` | String | yes | Logistics number |   |
| `sendGoods` | [message:alibaba.logistics.OpSendGood[]](#m-alibaba-logistics-opsendgood[]) | yes | Shipping details |   |
| `sendSuccessList` | [message:com.alibaba.ocean.openplatform.biz.logistics.result.OpSendOrderResult[]](#m-com-alibaba-ocean-openplatform-biz-logistics-result-opsendorderresult[]) | yes | List of packages successfully shipped in multi-package shipment | [{"sourceId":1,"sourceEntryId":11,"extBody":"{\"cpCode\":\"SF\",\"logisticsCpName\":\"顺丰\",\"mailNo\":\"a1231\",\"quantity\":1}","success":true,"logisticsId":"123"}] |
| `sendFailList` | [message:com.alibaba.ocean.openplatform.biz.logistics.result.OpSendOrderResult[]](#m-com-alibaba-ocean-openplatform-biz-logistics-result-opsendorderresult[]) | yes | List of packages that failed to ship in a multi-package shipment | [{"sourceId":1,"sourceEntryId":11,"extBody":"{\"cpCode\":\"SF\",\"logisticsCpName\":\"顺丰\",\"mailNo\":\"a1231\",\"quantity\":1}","success":false,"errorMessage":"错误原因"}] |

<a id="m-alibaba-logistics-opsendgood[]"></a>
#### alibaba.logistics.OpSendGood[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `sourceId` | String | yes | Shipment object ID, generally the order ID |   |
| `sendGoodEntries` | [message:alibaba.logistics.OpSendGoodEntry[]](#m-alibaba-logistics-opsendgoodentry[]) | yes | List of shipment object details |   |

<a id="m-alibaba-logistics-opsendgoodentry[]"></a>
#### alibaba.logistics.OpSendGoodEntry[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `sourceEntryId` | String | yes | Shipment target line item ID, corresponding to the sub-order ID |   |
| `amount` | Long | yes | Actual shipped quantity for the shipment target |   |
| `weight` | Double | yes | The actual shipped weight of the shipment object; the weight unit defaults to kilograms. |   |
| `extBody` | String | yes | JSON array string passed when shipping multiple packages. cpCode is the logistics company code, corresponding to companyNo (not the numeric Id) from the logistics company information API; logisticsCpName: logistics company name, corresponding to companyName from the logistics company information API; mailNo: waybill number; this data can be queried via the &quot;Logistics Company List - Self-linked Logistics&quot; API. quantity: number of shipment objects in the package, required. | [{"cpCode":"SF","logisticsCpName":"顺丰","mailNo":"a1231","quantity":1},{"cpCode":"SF","logisticsCpName":"顺丰","mailNo":"a1232","quantity":1}] |

<a id="m-com-alibaba-ocean-openplatform-biz-logistics-result-opsendorderresult[]"></a>
#### com.alibaba.ocean.openplatform.biz.logistics.result.OpSendOrderResult[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `sourceId` | String | yes | Main order id | 1 |
| `sourceEntryId` | String | yes | Sub-order ID | 11 |
| `extBody` | String | yes | JSON string (corresponding to one item in the multi-package input parameter) | {"cpCode":"SF","logisticsCpName":"顺丰","mailNo":"a1231","quantity":1} |
| `success` | Boolean | yes | Whether successful | true |
| `logisticsId` | String | yes | Logistics tracking number | 123 |
| `errorMessage` | String | yes | Error message | aaa |

## Error codes

| Error | Symptom | How to fix |
|---|---|---|
| System busy, please try again later Case1 | The API call returned “System busy, please try again later” | Please check whether the order number is correct, and also confirm whether the logistics company CpCode is correct. cpcode is the companyNo of the logistics company information, not the Id — for example, SF Express's coCode is "SF", not 106422. |
| System busy, please try again later Case2 | Returns "System busy, please try again later" when using the testing tool | Do not pass the receiverInfo field |
| HSF Server unexpected exception | HSF Server unexpected exception | The field type of the input parameter is incorrect, or some required fields were not passed |
| INVALID_PARAM | Waybill number does not conform to the rules or has already been used | Please check whether the waybill number is correct; also note that a waybill number can be used at most 15 times |
| 5001 | The downstream sales order is being refunded; shipping is not currently supported. It is recommended to communicate with your customer to handle the refund before deciding whether to continue shipping | The downstream sales order is under refund; shipping is not currently supported. It is recommended to communicate with your customer to process the refund first, then decide whether to continue shipping; to continue shipping, please ship via the 1688work workbench |
| 5002 | The downstream sales order has already been successfully refunded / the transaction is closed; shipping is not supported | The downstream sales order has already been successfully refunded / the transaction is closed; shipping is not supported |

## Samples

**Output parameter example**

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
