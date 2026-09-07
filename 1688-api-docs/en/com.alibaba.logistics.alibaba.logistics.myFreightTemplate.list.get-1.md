# Get shipping-template details

Original name: 获取物流模板详情  
API: `com.alibaba.logistics:alibaba.logistics.myFreightTemplate.list.get:1` · Category: Logistics  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.logistics:alibaba.logistics.myFreightTemplate.list.get-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.logistics/alibaba.logistics.myFreightTemplate.list.get/{appKey}`  
Requires user authorization (access_token) · Requires signature

Get the seller's shipping template by template ID. Template ID 0 means "shipping fee to be explained"; 1 means the seller bears the shipping fee.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `templateId` | java.lang.Long | no | Template id, used for single-record query scenarios | xxx |
| `querySubTemplate` | java.lang.Boolean | no | Whether to query the sub-template | false |
| `queryRate` | java.lang.Boolean | no | Whether to query the sub-template rate | false |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.logistics.FreightTemplate[]](#m-alibaba-logistics-freighttemplate[]) | yes | Return result | [] |
| `errorCode` | java.lang.String | yes | Error code | 错误码 |
| `errorMsg` | java.lang.String | yes | Error description | 错误描述 |

<a id="m-alibaba-logistics-freighttemplate[]"></a>
#### alibaba.logistics.FreightTemplate[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `addressCodeText` | java.lang.String | yes | Address information | xxx |
| `fromAreaCode` | java.lang.String | yes | Shipping address region code | xxx |
| `id` | java.lang.Long | yes | Address ID | xxx |
| `memberId` | java.lang.String | yes | Member ID | xxx |
| `name` | java.lang.String | yes | Name | xxx |
| `remark` | java.lang.String | yes | Remark | xxx |
| `status` | java.lang.Integer | yes | Status | xxx |
| `expressSubTemplate` | [message:alibaba.openplatform.logistics.DeliverySubTemplateDetailDTO](#m-alibaba-openplatform-logistics-deliverysubtemplatedetaildto) | yes | Express delivery sub-template | xxx |
| `logisticsSubTemplate` | [message:alibaba.openplatform.logistics.DeliverySubTemplateDetailDTO](#m-alibaba-openplatform-logistics-deliverysubtemplatedetaildto) | yes | Freight sub-template | xxx |
| `codSubTemplate` | [message:alibaba.openplatform.logistics.DeliverySubTemplateDetailDTO](#m-alibaba-openplatform-logistics-deliverysubtemplatedetaildto) | yes | Cash-on-delivery sub-template | xxx |
| `type` | String | yes | Type: 3 - official logistics template; 2 or no value - user logistics template | 3 |

<a id="m-alibaba-openplatform-logistics-deliverysubtemplatedetaildto"></a>
#### alibaba.openplatform.logistics.DeliverySubTemplateDetailDTO

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `subTemplateDTO` | [message:alibaba.openplatform.logistics.DeliverySubTemplateDTO](#m-alibaba-openplatform-logistics-deliverysubtemplatedto) | yes | Sub-template |  |
| `rateList` | [message:alibaba.openplatform.logistics.DeliveryRateDetailDTO[]](#m-alibaba-openplatform-logistics-deliveryratedetaildto[]) | yes | Rate |  |

<a id="m-alibaba-openplatform-logistics-deliverysubtemplatedto"></a>
#### alibaba.openplatform.logistics.DeliverySubTemplateDTO

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `chargeType` | java.lang.Integer | yes | Counting type. 0: weight 1: piece count 2: volume |  |
| `isSysTemplate` | java.lang.Boolean | yes | Whether it is a system template |  |
| `serviceChargeType` | java.lang.Integer | yes | Shipping fee bearer type. Seller bears it: 0; buyer bears it: 1. |  |
| `serviceType` | java.lang.Integer | yes | Service type. 0: express delivery 1: freight 2: cash on delivery |  |
| `type` | java.lang.Integer | yes | Sub-template type. 0: baseline; 1: value-added. Default 0. |  |

<a id="m-alibaba-openplatform-logistics-deliveryratedetaildto[]"></a>
#### alibaba.openplatform.logistics.DeliveryRateDetailDTO[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `isSysRate` | boolean | yes | Whether it is a system template |  |
| `toAreaCodeText` | java.lang.String | yes | Address code text, separated by 、 (Chinese enumeration comma). Example: Shanghai、Fujian Province、Guangdong Province |  |
| `rateDTO` | [message:alibaba.openplatform.logistics.DeliveryRateDTO](#m-alibaba-openplatform-logistics-deliveryratedto) | yes | Regular sub-template rate |  |
| `sysRateDTO` | [message:alibaba.openplatform.logistics.DeliverySysRateDTO](#m-alibaba-openplatform-logistics-deliverysysratedto) | yes | System sub-template rate |  |

<a id="m-alibaba-openplatform-logistics-deliveryratedto"></a>
#### alibaba.openplatform.logistics.DeliveryRateDTO

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `firstUnit` | java.lang.Long | yes | First weight unit (in grams) or first item unit (in pieces) |  |
| `firstUnitFee` | java.lang.Long | yes | Price for the first weight unit or first item |  |
| `leastExpenses` | java.lang.Long | yes | Minimum per shipment |  |

<a id="m-alibaba-openplatform-logistics-deliverysysratedto"></a>
#### alibaba.openplatform.logistics.DeliverySysRateDTO

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `firstUnit` | java.lang.Long | yes | First weight unit (in grams) or first item unit (in pieces) |  |
| `firstUnitFee` | java.lang.Long | yes | Price for the first weight unit or first item |  |
| `leastExpenses` | java.lang.Long | yes | Minimum per shipment |  |
| `nextUnit` | java.lang.Long | yes | Additional weight (in grams) or additional item unit (in pieces) |  |

## Samples

**Output parameter example**

```
      
{
    "result":[
        {
            "addressCodeText":"湖北省 荆门市 东宝区",
            "fromAreaCode":"420802",
            "id":11864709,
            "memberId":"b2b-1624961198",
            "name":"test168",
            "remark":"",
            "status":1
        }
    ]
}
```
