# 查询AI找商任务执行结果

API: `com.alibaba.fenxiao.crossborder:account.search.getInfoResult:1` · Category: 商家  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:account.search.getInfoResult-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/account.search.getInfoResult/{appKey}`  
需要授权 (access_token) · 需要签名

查询AI找商任务执行结果

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `request` | [message:alibaba.global1688.common.request.SpApiExecuteStatusRequest](#m-alibaba-global1688-common-request-spapiexecutestatusrequest) | 是 | 请求参数 | {} |

<a id="m-alibaba-global1688-common-request-spapiexecutestatusrequest"></a>
#### alibaba.global1688.common.request.SpApiExecuteStatusRequest

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `taskId` | java.lang.String | 是 | 任务ID, 通过任务执行 API 返回值获取 | taskIdxxxxxxxxxx |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:alibaba.global1688.common.response.ApiSpExecuteTaskResult](#m-alibaba-global1688-common-response-apispexecutetaskresult) | 是 | 返回值 | {} |

<a id="m-alibaba-global1688-common-response-apispexecutetaskresult"></a>
#### alibaba.global1688.common.response.ApiSpExecuteTaskResult

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `success` | java.lang.Boolean | 是 | 是否执行成功 | true |
| `result` | [message:alibaba.global1688.common.response.SpExecuteTask](#m-alibaba-global1688-common-response-spexecutetask) | 是 | AI找商任务运行情况 | {} |
| `code` | java.lang.String | 是 | 状态码 | SUCCESS |
| `message` | java.lang.String | 是 | 错误信息 | error_msg_xxxxx |

<a id="m-alibaba-global1688-common-response-spexecutetask"></a>
#### alibaba.global1688.common.response.SpExecuteTask

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `status` | java.lang.String | 是 | 任务执行状态, 返回值有两种状态, 分别是 RUNNING - 运行中 和 FINISHED - 已完成 | RUNNING |
| `providerList` | [message:alibaba.global1688.common.response.ProviderApiResult[]](#m-alibaba-global1688-common-response-providerapiresult[]) | 是 | AI找商结果 | [] |

<a id="m-alibaba-global1688-common-response-providerapiresult[]"></a>
#### alibaba.global1688.common.response.ProviderApiResult[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `companyName` | java.lang.String | 是 | 公司名称 | xxxxx工厂 |
| `factoryUrl` | java.lang.String | 是 | 工厂 1688 链接 | https://xxxx |
| `businessModel` | java.lang.String | 是 | 经营模式, 返回有 factory - 源头工厂 和 trading - 贸易商 | factory |
| `establishedYears` | java.lang.Integer | 是 | 公司成立年限 | 8 |
| `repeatPurchaseRate` | java.lang.Integer | 是 | 近 90 天回头率, 返回 [0, 100] | 90 |
| `supportCustomization` | java.lang.Boolean | 是 | 是否支持加工定制 | true |
| `afterSalesScore` | java.lang.Double | 是 | 售后体验分 | 4.0 |
| `productQualityScore` | java.lang.Double | 是 | 商品体验分 | 4.0 |
| `consultationResponseScore` | java.lang.Double | 是 | 咨询体验分 | 4.0 |
| `logisticsTimelinessScore` | java.lang.Double | 是 | 物流体验分 | 4.0 |
| `itemList` | [message:alibaba.global1688.common.response.ItemApiInfo[]](#m-alibaba-global1688-common-response-itemapiinfo[]) | 是 | 推荐商品 | [] |

<a id="m-alibaba-global1688-common-response-itemapiinfo[]"></a>
#### alibaba.global1688.common.response.ItemApiInfo[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `title` | java.lang.String | 是 | 商品标题 | 毛巾 |
| `itemId` | java.lang.Long | 是 | 商品ID | 11111111 |
| `itemPrice` | java.lang.Double | 是 | 商品价格 | 10.5 |
| `imageUrl` | java.lang.String | 是 | 商品图片链接 | https://xxxxxx |
| `totalOnlineSales1y` | java.lang.Long | 是 | 近一年全网销量 | 600 |
| `itemSales1y` | java.lang.Long | 是 | 当前商品近一年销量 | 50 |

## 示例

**入参示例**

```
{
    "taskId": "52aa3fb35b494e15a93086f981968907"
}
```

**出参示例**

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
