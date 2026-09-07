# Buyer views all granted credit terms

Original name: 买家查看获得的所有账期授信  
API: `com.alibaba.trade:alibaba.accountPeriod.list.buyerView:1` · Category: Payment  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.accountPeriod.list.buyerView-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.accountPeriod.list.buyerView/{appKey}`  
Requires user authorization (access_token) · Requires signature

View, from the buyer's side, all account-period (credit-term) lines the buyer has been granted. Paginated; at most 10 records per call.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `pageIndex` | Long | no | Page number | 1 |
| `sellerLoginId` | String | no | Seller ID; if not filled in, all are queried | alitestforisv01 |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `success` | Boolean | yes | Whether successful | true |
| `errorCode` | String | yes | Error code | 500_1 |
| `errorMsg` | String | yes | Error message |   |
| `resultList` | [message:accountPeriod.list.buyerView.result](#m-accountperiod-list-buyerview-result) | yes | Returned data result | {} |

<a id="m-accountperiod-list-buyerview-result"></a>
#### accountPeriod.list.buyerView.result

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `totalCount` | String | yes | Total number of records | 100 |
| `accountPeriodList` | [message:alibaba.ocean.openplatform.biz.trade.common.model.AccountPeriodInfo[]](#m-alibaba-ocean-openplatform-biz-trade-common-model-accountperiodinfo[]) | yes | Credit list | [] |

<a id="m-alibaba-ocean-openplatform-biz-trade-common-model-accountperiodinfo[]"></a>
#### alibaba.ocean.openplatform.biz.trade.common.model.AccountPeriodInfo[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `sellerLoginId` | java.lang.String | yes | Seller loginId | alitestforisv02 |
| `sellerCompanyName` | String | yes | Seller's company name | 公司名 |
| `gmtQuota` | java.util.Date | yes | Credit date | 20170913231727000+0800 |
| `quota` | java.lang.Long | yes | Credit line amount, in cents (fen) | 10000 |
| `surplusQuota` | java.lang.Long | yes | Available credit line amount, in cents (fen) | 10000 |
| `statusStr` | java.lang.String | yes | Status description | 有效 |
| `tapDateStr` | java.lang.String | yes | Account period date description | 两个月一结，20号 |
| `tapOverdue` | java.lang.Integer | yes | Number of overdue occurrences | 0 |

## Samples

**Output parameter example**

```
{
  "resultList": {
    "accountPeriodList": [
      {
        "gmtQuota": "20190110114140000+0800",
        "quota": 90000000,
        "sellerCompanyName": "AOP对外测试账号01",
        "sellerLoginId": "alitestforisv01",
        "statusStr": "有效",
        "surplusQuota": 90000000,
        "tapDateStr": "120天",
        "tapOverdue": 0,
      }
    ],
    "totalCount": "1"
  },
  "success": true
}
```
