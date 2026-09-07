# Get 90-day repurchase-rate trend for a product (new)

Original name: 获取商品90天复购率指标趋势(新）   
API: `com.alibaba.fenxiao.crossborder:product.analyze.getRepurchaseRateTrendNew:1` · Category: Market Insights  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:product.analyze.getRepurchaseRateTrendNew-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/product.analyze.getRepurchaseRateTrendNew/{appKey}`  
Requires user authorization (access_token) · Requires signature

Get a product's 90-day repurchase-rate metric trend (at most 30 days can be queried). New interface.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `trendQueryParam` | [message:com.alibaba.cbu.offer.param.OfferSellTrendQueryParam](#m-com-alibaba-cbu-offer-param-offerselltrendqueryparam) | yes |  |  |

<a id="m-com-alibaba-cbu-offer-param-offerselltrendqueryparam"></a>
#### com.alibaba.cbu.offer.param.OfferSellTrendQueryParam

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `offerId` | java.lang.Long | yes | Product ID | 123 |
| `startDate` | java.lang.String | yes | Query start time | 20240701(起始时间和截止时间不要超过一个月） |
| `endDate` | java.lang.String | yes | Query end time | 20240710 |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:com.alibaba.cbu.common.result.CommonResult](#m-com-alibaba-cbu-common-result-commonresult) | yes |  |  |

<a id="m-com-alibaba-cbu-common-result-commonresult"></a>
#### com.alibaba.cbu.common.result.CommonResult

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `success` | java.lang.Boolean | yes | Whether successful | true |
| `retCode` | java.lang.String | yes | Error code | S000 |
| `retMsg` | java.lang.String | yes | Return message | 成功 |
| `result` | [message:com.alibaba.cbu.offer.model.OfferSellTrendDataModel[]](#m-com-alibaba-cbu-offer-model-offerselltrenddatamodel[]) | yes | Result | [ { "date": "20240701", "value": "20" }, { "date": "20240702", "value": "20" }, { "date": "20240703", "value": "20" }, { "date": "20240704", "value": "20" }, { "date": "20240705", "value": "" } ] |

<a id="m-com-alibaba-cbu-offer-model-offerselltrenddatamodel[]"></a>
#### com.alibaba.cbu.offer.model.OfferSellTrendDataModel[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `date` | java.lang.String | yes | The date corresponding to the queried metric data | 20240701 |
| `value` | java.lang.String | yes | Metric data | 10% |

## Samples

**Input parameter example**

```
{
  "trendQueryParam": {
    "offerId": 123,
    "startDate": "20240701(起始时间和截止时间不要超过一个月）",
    "endDate": "20240710"
  }
}
```

**Output parameter example**

```
{
  "result": {
    "success": true,
    "retCode": "S000",
    "retMsg": "成功",
    "result": [
      {
        "date": "20240701",
        "value": "10%"
      }
    ]
  }
}
```
