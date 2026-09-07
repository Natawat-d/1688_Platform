# Query AI supplier-search task result

Original name: 查询AI找商任务执行结果  
API: `com.alibaba.fenxiao.crossborder:account.search.getInfoResult:1` · Category: Suppliers  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:account.search.getInfoResult-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/account.search.getInfoResult/{appKey}`  
Requires user authorization (access_token) · Requires signature

Query the execution result of an AI supplier-search task.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `request` | [message:alibaba.global1688.common.request.SpApiExecuteStatusRequest](#m-alibaba-global1688-common-request-spapiexecutestatusrequest) | yes | Request parameters | {} |

<a id="m-alibaba-global1688-common-request-spapiexecutestatusrequest"></a>
#### alibaba.global1688.common.request.SpApiExecuteStatusRequest

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `taskId` | java.lang.String | yes | Task ID, obtained from the return value of the task execution API | taskIdxxxxxxxxxx |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.global1688.common.response.ApiSpExecuteTaskResult](#m-alibaba-global1688-common-response-apispexecutetaskresult) | yes | Return value | {} |

<a id="m-alibaba-global1688-common-response-apispexecutetaskresult"></a>
#### alibaba.global1688.common.response.ApiSpExecuteTaskResult

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `success` | java.lang.Boolean | yes | Whether execution succeeded | true |
| `result` | [message:alibaba.global1688.common.response.SpExecuteTask](#m-alibaba-global1688-common-response-spexecutetask) | yes | AI supplier-sourcing task running status | {} |
| `code` | java.lang.String | yes | Status code | SUCCESS |
| `message` | java.lang.String | yes | Error message | error_msg_xxxxx |

<a id="m-alibaba-global1688-common-response-spexecutetask"></a>
#### alibaba.global1688.common.response.SpExecuteTask

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `status` | java.lang.String | yes | Task execution status. There are two possible return values: RUNNING - in progress, and FINISHED - completed. | RUNNING |
| `providerList` | [message:alibaba.global1688.common.response.ProviderApiResult[]](#m-alibaba-global1688-common-response-providerapiresult[]) | yes | AI supplier-finding result | [] |

<a id="m-alibaba-global1688-common-response-providerapiresult[]"></a>
#### alibaba.global1688.common.response.ProviderApiResult[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `companyName` | java.lang.String | yes | Company name | xxxxx工厂 |
| `factoryUrl` | java.lang.String | yes | Factory 1688 link | https://xxxx |
| `businessModel` | java.lang.String | yes | Business model. Returned values include factory - source factory and trading - trading company | factory |
| `establishedYears` | java.lang.Integer | yes | Number of years the company has been established | 8 |
| `repeatPurchaseRate` | java.lang.Integer | yes | Repeat purchase rate over the last 90 days, returned in [0, 100] | 90 |
| `supportCustomization` | java.lang.Boolean | yes | Whether processing/customization is supported | true |
| `afterSalesScore` | java.lang.Double | yes | After-sales experience score | 4.0 |
| `productQualityScore` | java.lang.Double | yes | Product experience score | 4.0 |
| `consultationResponseScore` | java.lang.Double | yes | Inquiry experience score | 4.0 |
| `logisticsTimelinessScore` | java.lang.Double | yes | Logistics experience score | 4.0 |
| `itemList` | [message:alibaba.global1688.common.response.ItemApiInfo[]](#m-alibaba-global1688-common-response-itemapiinfo[]) | yes | Recommended product | [] |

<a id="m-alibaba-global1688-common-response-itemapiinfo[]"></a>
#### alibaba.global1688.common.response.ItemApiInfo[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `title` | java.lang.String | yes | Product title | 毛巾 |
| `itemId` | java.lang.Long | yes | Product ID | 11111111 |
| `itemPrice` | java.lang.Double | yes | Product price | 10.5 |
| `imageUrl` | java.lang.String | yes | Product image link | https://xxxxxx |
| `totalOnlineSales1y` | java.lang.Long | yes | Platform-wide sales volume over the past year | 600 |
| `itemSales1y` | java.lang.Long | yes | Sales volume of this product over the past year | 50 |

## Samples

**Input parameter example**

```
{
    "taskId": "52aa3fb35b494e15a93086f981968907"
}
```

**Output parameter example**

```
{
    "result": {
        "status": "FINISHED",
        "providerList": [
            {
                "factoryUrl": "https://xxxxxxx",
                "businessModel": "factory",
                "productQualityScore": 5,
                "companyName": "xxxxxxx服饰有限公司",
                "repeatPurchaseRate": 100,
                "itemList": [
                    {
                        "itemSales1y": 100,
                        "itemId": 100000,
                        "totalOnlineSales1y": 999,
                        "imageUrl": "https://xxxxxxx",
                        "title": "xxxxxxx连衣裙"
                    },
                    {
                        "itemSales1y": 10,
                        "itemId": 100000,
                        "totalOnlineSales1y": 999,
                        "imageUrl": "https://xxxxxxx",
                        "title": "xxxxxx连衣裙"
                    },
                    {
                        "itemSales1y": 9,
                        "itemId": 100000,
                        "totalOnlineSales1y": 9999,
                        "imageUrl": "https://xxxxxxx",
                        "title": "xxxxxxx连体裤"
                    }
                ],
                "establishedYears": 17,
                "supportCustomization": true,
                "consultationResponseScore": 4,
                "afterSalesScore": 5,
                "logisticsTimelinessScore": 5
            },
            {
                "factoryUrl": "https://xxxxxxx",
                "businessModel": "trading",
                "productQualityScore": 4.33,
                "companyName": "xxxxxxx经贸有限公司",
                "repeatPurchaseRate": 100,
                "itemList": [
                    {
                        "itemSales1y": 888,
                        "itemId": 100000,
                        "totalOnlineSales1y": 9999,
                        "imageUrl": "https://xxxxxxx",
                        "title": "xxxxx连衣裙"
                    },
                    {
                        "itemSales1y": 888,
                        "itemId": 100000,
                        "totalOnlineSales1y": 9999,
                        "imageUrl": "https://xxxxxxx",
                        "title": "xxxxxxx连衣裙"
                    },
                    {
                        "itemSales1y": 888,
                        "itemId": 100000,
                        "totalOnlineSales1y": 9999,
                        "imageUrl": "https://xxxxxxx",
                        "title": "xxxxxxx连衣裙"
                    }
                ],
                "establishedYears": 15,
                "supportCustomization": false,
                "consultationResponseScore": 4,
                "afterSalesScore": 4.67,
                "logisticsTimelinessScore": 3.86
            }
        ]
    },
    "code": "SUCCESS",
    "success": true
}
```
