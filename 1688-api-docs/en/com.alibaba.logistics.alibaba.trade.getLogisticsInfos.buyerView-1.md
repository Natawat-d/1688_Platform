# Get order logistics info (buyer view)

Original name: 获取交易订单的物流信息(买家视角)  
API: `com.alibaba.logistics:alibaba.trade.getLogisticsInfos.buyerView:1` · Category: Orders  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.logistics:alibaba.trade.getLogisticsInfos.buyerView-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.logistics/alibaba.trade.getLogisticsInfos.buyerView/{appKey}`  
Requires user authorization (access_token) · Requires signature

Requires the order buyer's authorization and returns the logistics details of the buyer's order. In procurement or distribution scenarios, buyers also need logistics details. This interface looks up logistics details by order number, including the sender, the recipient and the details of the goods shipped. Because of when logistics records are entered, API tracking queries may be delayed. Permission must be requested from the Open Platform to access this API.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `orderId` | Long | yes | Order number | 1221434 |
| `fields` | String | no | Fields to be returned. Currently available: company.name, sender, receiver, sendgood. The returned fields must be separated by English commas. | company,name,sender,receiver,sendgood |
| `webSite` | String | yes | Whether it is a 1688 business or an icbu business | 1688或者alibaba |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.logistics.OpenPlatformLogisticsOrder[]](#m-alibaba-logistics-openplatformlogisticsorder[]) | yes | Return result | {} |
| `errorCode` | String | yes | Error code |   |
| `errorMessage` | String | yes | Error message |   |
| `success` | Boolean | yes | Whether successful | true |

<a id="m-alibaba-logistics-openplatformlogisticsorder[]"></a>
#### alibaba.logistics.OpenPlatformLogisticsOrder[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `logisticsId` | String | yes | Logistics information ID | 物流信息ID |
| `logisticsBillNo` | String | yes | Logistics tracking number, waybill number | 物流单号，运单号 |
| `orderEntryIds` | String | yes | List of order numbers. Equal to the main order number if there are no sub-orders; otherwise, the corresponding list of sub-order numbers. | 129232515787615400,129232515788615400,129232515789615400,129232515790615400 |
| `status` | String | yes | Logistics status. WAITACCEPT: not accepted; CANCEL: canceled; ACCEPT: accepted; TRANSPORT: in transit; NOGET: pickup failed; SIGN: signed for; UNSIGN: sign-for exception | WAITACCEPT |
| `logisticsCompanyId` | String | yes | Logistics company ID | 物流公司ID |
| `logisticsCompanyName` | String | yes | Logistics company code | 物流公司编码 |
| `logisticsCompanyNo` | String | yes | Logistics company number | 物流公司编号 |
| `remarks` | String | yes | Remark | 备注 |
| `serviceFeature` | String | yes | serviceFeature | serviceFeature |
| `gmtSystemSend` | String | yes | gmtSystemSend | gmtSystemSend	 |
| `sendGoods` | [message:alibaba.logistics.OpenPlatformLogisticsSendGood[]](#m-alibaba-logistics-openplatformlogisticssendgood[]) | yes | Product information | 商品信息 |
| `receiver` | [message:alibaba.logistics.OpenPlatformLogisticsReceiver](#m-alibaba-logistics-openplatformlogisticsreceiver) | yes | Recipient info | 收件人信息 |
| `sender` | [message:alibaba.logistics.OpenPlatformLogisticsSender](#m-alibaba-logistics-openplatformlogisticssender) | yes | Sender info | 发件人信息 |
| `logisticsOrderGoods` | [message:com.alibaba.ocean.openplatform.biz.logistics.common.model.OpenPlatformLogisticsOrderSendGood[]](#m-com-alibaba-ocean-openplatform-biz-logistics-common-model-openplatformlogisticsordersendgood[]) | yes | Model of the relationship between logistics order and product | {} |
| `logisticsOrderSendGood` | [message:com.alibaba.ocean.openplatform.biz.logistics.common.model.seller.OpenPlatformLogisticsOrderSendGood[]](#m-com-alibaba-ocean-openplatform-biz-logistics-common-model-seller-openplatformlogisticsordersendgood[]) | yes | Model of the association between logistics and order | {} |

<a id="m-alibaba-logistics-openplatformlogisticssendgood[]"></a>
#### alibaba.logistics.OpenPlatformLogisticsSendGood[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `goodName` | String | yes | Product name |  |
| `quantity` | String | yes | Product quantity |  |
| `unit` | String | yes | Product unit |  |

<a id="m-alibaba-logistics-openplatformlogisticsreceiver"></a>
#### alibaba.logistics.OpenPlatformLogisticsReceiver

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `receiverName` | String | yes | Recipient's name |  |
| `receiverPhone` | String | yes | Recipient phone number |  |
| `receiverMobile` | String | yes | Recipient phone number |  |
| `encrypt` | String | yes |  |  |
| `receiverProvinceCode` | String | yes | Province code |  |
| `receiverCityCode` | String | yes | City code |  |
| `receiverCountyCode` | String | yes | Country code |  |
| `receiverAddress` | String | yes | Address |  |
| `receiverProvince` | String | yes | Province |  |
| `receiverCity` | String | yes | City |  |
| `receiverCounty` | String | yes | Country |  |

<a id="m-alibaba-logistics-openplatformlogisticssender"></a>
#### alibaba.logistics.OpenPlatformLogisticsSender

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `senderName` | String | yes | Sender's name |  |
| `senderPhone` | String | yes | Sender's phone number |  |
| `senderMobile` | String | yes | Sender's phone number |  |
| `encrypt` | String | yes |  |  |
| `senderProvinceCode` | String | yes | Province code |  |
| `senderCityCode` | String | yes | City code |  |
| `senderCountyCode` | String | yes | Country code |  |
| `senderAddress` | String | yes | Shipper's address |  |
| `senderProvince` | String | yes | Province |  |
| `senderCity` | String | yes | City |  |
| `senderCounty` | String | yes | Country |  |

<a id="m-com-alibaba-ocean-openplatform-biz-logistics-common-model-openplatformlogisticsordersendgood[]"></a>
#### com.alibaba.ocean.openplatform.biz.logistics.common.model.OpenPlatformLogisticsOrderSendGood[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `logisticsId` | String | yes | Logistics number | LP00616288919385 |
| `tradeOrderId` | Long | yes | Main trade order number | 3657499272179232333 |
| `tradeOrderItemId` | Long | yes | Transaction sub-order number | 3657499272179232333 |
| `description` | String | yes | SKU description | 颜色: 黑色; 尺码: M; |
| `quantity` | Double | yes | Quantity | 1 |
| `unit` | String | yes | Unit | 件 |
| `productName` | String | yes | Product name | 品名 |

<a id="m-com-alibaba-ocean-openplatform-biz-logistics-common-model-seller-openplatformlogisticsordersendgood[]"></a>
#### com.alibaba.ocean.openplatform.biz.logistics.common.model.seller.OpenPlatformLogisticsOrderSendGood[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `logisticsId` | String | yes | Logistics number | LP00616288919385 |
| `tradeOrderId` | Long | yes | Main trade order number | 3657499272179232333 |
| `tradeOrderItemId` | Long | yes | Transaction sub-order number | 3657499272179232333 |
| `description` | String | yes | SKU description | 颜色: 黑色; 尺码: M; |
| `quantity` | Double | yes | Quantity | 1 |
| `unit` | String | yes | Unit | 件 |
| `productName` | String | yes | Product name | 品名 |

## Error codes

| Error | Symptom | How to fix |
|---|---|---|
| 500_2 | The order has not been shipped yet; no logistics details are available. Please try again later. | The order has not shipped yet. It is recommended to listen for the order shipment message and query the order logistics status after receiving it. Reference: https://open.1688.com/doc/topicDetail.htm?spm=a260s.11630592.0.0.555655ed4QUvy5&topicGroup=ORDER&id=ORDER_BUYER_VIEW_ANNOUNCE_SENDGOODS |

## Samples

**Return result**

```
{
  "result": [
    {
      "logisticsId": "BX111841674232006",
      "orderEntryIds": "149989279191615400",
      "status": "SIGN",
      "logisticsCompanyId": "8",
      "remarks": "asdasdasdasdsad啊啊啊",
      "sendGoods": [
        {
          "unit": "个",
          "quantity": 2,
          "goodName": "【茂茂回归】12.9公开-sku区间价-淘货源"
        }
      ],
      "receiver": {
        "receiverCountyCode": "310105",
        "receiverCity": "上海市",
        "receiverProvinceCode": "310000",
        "receiverCityCode": "310100",
        "encrypt": "CN",
        "receiverCounty": "长宁区",
        "receiverMobile": "13800138000",
        "receiverProvince": "上海",
        "receiverName": "XXX",
        "receiverAddress": "XXX街道 XX路999",
        "receiverPhone": ""
      },
      "sender": {
        "senderCityCode": "330100",
        "senderAddress": "网商路699号",
        "senderName": "张三",
        "senderProvince": "浙江省",
        "encrypt": "CN",
        "senderCity": "杭州市",
        "senderCountyCode": "330108",
        "senderProvinceCode": "330000",
        "senderMobile": "12345678900",
        "senderPhone": "",
        "senderCounty": "滨江区"
      }
    }
  ]
}
```
