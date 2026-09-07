# 组合收银台url获取

API: `com.alibaba.trade:alibaba.trade.grouppay.url.get:1` · Category: 支付  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.trade.grouppay.url.get-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.trade.grouppay.url.get/{appKey}`  
需要授权 (access_token) · 需要签名

组合收银台url获取

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `orderIds` | Long[] | 是 | 订单列表 | [123123413,1223234] |
| `payPlatformType` | String | 否 | PC或WIRELESS | PC |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `results` | [message:alibaba.ocean.openplatform.biz.trade.result.TradeCreateGroupPayUrlResult](#m-alibaba-ocean-openplatform-biz-trade-result-tradecreategrouppayurlresult) | 是 |  |  |

<a id="m-alibaba-ocean-openplatform-biz-trade-result-tradecreategrouppayurlresult"></a>
#### alibaba.ocean.openplatform.biz.trade.result.TradeCreateGroupPayUrlResult

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `payUrl` | java.lang.String | 是 | 返回payurl | https://payment2.m.1688.com/page/cashier.html?orderId=154612245 |
| `success` | boolean | 是 | 是否成功 | true |
| `errorInfo` | java.lang.String | 是 | 错误信息 | null |
| `errorCode` | java.lang.String | 是 | 错误码 | null |
