# Logistics company list (all companies)

Original name: 物流公司列表-所有的物流公司  
API: `com.alibaba.logistics:alibaba.logistics.OpQueryLogisticCompanyList:1` · Category: Logistics  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.logistics:alibaba.logistics.OpQueryLogisticCompanyList-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.logistics/alibaba.logistics.OpQueryLogisticCompanyList/{appKey}`  
Requires user authorization (access_token) · Requires signature

Get the names of all logistics companies.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.logistics.OpLogisticsCompanyModel[]](#m-alibaba-logistics-oplogisticscompanymodel[]) | yes | List of logistics companies | [] |
| `success` | Boolean | yes | Whether successful | true |
| `errorCode` | String | yes | Error code | 500 |
| `errorMessage` | String | yes | Error code description | 错误码描述 |
| `extErrorMessage` | String | yes | Extended error code description |   |

<a id="m-alibaba-logistics-oplogisticscompanymodel[]"></a>
#### alibaba.logistics.OpLogisticsCompanyModel[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `id` | Long | yes |  |   |
| `companyName` | String | yes | Logistics company name | xxx |
| `companyNo` | String | yes | Logistics company number | xxx |
| `companyPhone` | String | yes | Logistics company service phone number | xxx |
| `supportPrint` | Boolean | yes | Whether printing is supported | xxx |
| `spelling` | String | yes | Full pinyin | xxx |

## Samples

**Output parameter example**
