# 查询卖家混批设置

API: `com.alibaba.trade:alibaba.trade.OpQueryMarketingMixConfig:1` · Category: 物流  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.trade.OpQueryMarketingMixConfig-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.trade.OpQueryMarketingMixConfig/{appKey}`  
需要授权 (access_token) · 需要签名

查询卖家混批设置。Query seller settings for mixed batch.

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `sellerMemberId` | String | 否 | 卖家memberId | b2b-1623492085 |
| `sellerLoginId` | String | 否 | 卖家LoginId，sellerMemberId为空时，以loginId为准 | alitestforisv01 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:alibaba.openplatform.trade.result.OpMarketingMixConfigModel](#m-alibaba-openplatform-trade-result-opmarketingmixconfigmodel) | 是 | 返回结果 | {} |
| `errorCode` | String | 是 | 错误码 | 错误码 |
| `errorMessage` | String | 是 | 错误信息 |  错误信息 |
| `extErrorMessage` | String | 是 | 错误信息扩展 |  错误信息 |
| `success` | Boolean | 是 | 是否成功 | true |

<a id="m-alibaba-openplatform-trade-result-opmarketingmixconfigmodel"></a>
#### alibaba.openplatform.trade.result.OpMarketingMixConfigModel

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `generalHunpi` | boolean | 是 | 是否普通混批 | true |
| `gmtCreate` | java.util.Date | 是 | 创建时间 | 20130522193706000+0800 |
| `gmtModified` | java.util.Date | 是 | 修改时间 | 20180710083636000+0800 |
| `memberId` | java.lang.String | 是 | 卖家memberID | b2b-1623492085 |
| `mixAmount` | java.lang.Integer | 是 | 混批金额 | 99 |
| `mixNumber` | java.lang.Integer | 是 | 混批数量 | 1 |

## 错误码

| 错误码 | 现象 | 解决方案 |
|---|---|---|
| 500_1 | 卖家信息不得为空。 | 检查sllerMemberId和sellerLoginId，这两个不能同时为空 |

## 示例

**返回参数示例**

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
