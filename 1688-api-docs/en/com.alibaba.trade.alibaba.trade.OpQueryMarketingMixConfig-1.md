# Query seller mixed-batch settings

Original name: 查询卖家混批设置  
API: `com.alibaba.trade:alibaba.trade.OpQueryMarketingMixConfig:1` · Category: Logistics  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.trade.OpQueryMarketingMixConfig-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.trade.OpQueryMarketingMixConfig/{appKey}`  
Requires user authorization (access_token) · Requires signature

Query the seller's mixed-batch (mixed wholesale) settings.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `sellerMemberId` | String | no | Seller memberId | b2b-1623492085 |
| `sellerLoginId` | String | no | Seller LoginId. When sellerMemberId is empty, loginId takes precedence. | alitestforisv01 |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.openplatform.trade.result.OpMarketingMixConfigModel](#m-alibaba-openplatform-trade-result-opmarketingmixconfigmodel) | yes | Return result | {} |
| `errorCode` | String | yes | Error code | 错误码 |
| `errorMessage` | String | yes | Error message |  错误信息 |
| `extErrorMessage` | String | yes | Error info extension |  错误信息 |
| `success` | Boolean | yes | Whether successful | true |

<a id="m-alibaba-openplatform-trade-result-opmarketingmixconfigmodel"></a>
#### alibaba.openplatform.trade.result.OpMarketingMixConfigModel

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `generalHunpi` | boolean | yes | Whether it is a regular mixed batch | true |
| `gmtCreate` | java.util.Date | yes | Creation time | 20130522193706000+0800 |
| `gmtModified` | java.util.Date | yes | Modification time | 20180710083636000+0800 |
| `memberId` | java.lang.String | yes | Seller memberID | b2b-1623492085 |
| `mixAmount` | java.lang.Integer | yes | Mixed batch amount | 99 |
| `mixNumber` | java.lang.Integer | yes | Mixed batch quantity | 1 |

## Error codes

| Error | Symptom | How to fix |
|---|---|---|
| 500_1 | Seller information must not be empty. | Check sllerMemberId and sellerLoginId; these two cannot both be empty |

## Samples

**Example of return parameters**

```
{
  "result": {
    "generalHunpi": true,
    "gmtCreate": "20171127164631000+0800",
    "gmtModified": "20171127164631000+0800",
    "memberId": "b2b-1624786331",
    "mixAmount": 10,
    "mixNumber": 1
  },
  "success": true
}
```
