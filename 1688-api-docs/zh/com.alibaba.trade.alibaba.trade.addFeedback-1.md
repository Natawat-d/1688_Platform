# 买家补充订单留言接口

API: `com.alibaba.trade:alibaba.trade.addFeedback:1` · Category: 订单  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.trade.addFeedback-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.trade.addFeedback/{appKey}`  
需要授权 (access_token) · 需要签名

买家补充订单留言接口，注意留言总长不超500字符

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `tradeFeedbackParam` | [message:alibaba.ocean.openplatform.biz.trade.param.TradeFeedbackParam](#m-alibaba-ocean-openplatform-biz-trade-param-tradefeedbackparam) | 是 | 请求参数 | {"feedback":"test","orderId":"123123213"} |

<a id="m-alibaba-ocean-openplatform-biz-trade-param-tradefeedbackparam"></a>
#### alibaba.ocean.openplatform.biz.trade.param.TradeFeedbackParam

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `feedback` | java.lang.String | 是 | 留言 | 留言 |
| `orderId` | java.lang.String | 是 | 订单ID | 12344444555545 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:alibaba.ocean.openplatform.biz.trade.result.TradeFeedbackResult](#m-alibaba-ocean-openplatform-biz-trade-result-tradefeedbackresult) | 是 | 返回结果 | {} |
| `code` | String | 是 | 错误码 | 500_2 |
| `message` | String | 是 | 错误描述 | remote service error |
| `success` | Boolean | 是 | 是否成功 | false |

<a id="m-alibaba-ocean-openplatform-biz-trade-result-tradefeedbackresult"></a>
#### alibaba.ocean.openplatform.biz.trade.result.TradeFeedbackResult

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `errorInfo` | java.lang.String | 是 | 错误描述 | 错误描述 |
| `errorCode` | java.lang.String | 是 | 错误码 | 400_1 |
| `success` | Boolean | 是 | 是否成功 | true |

## 错误码

| 错误码 | 现象 | 解决方案 |
|---|---|---|
| 400_1 | parameter invalid | 检查参数是否有传及参数类型是否符合要求 |
| 400_2 | need authorization | 需要授权登录 |
| 500_1 | invoke remote service error | 调用远程服务异常，请在ISV交流群里咨询技术支持或帮助中心提单 |
| 500_2 | remote service error | 服务异常，请稍后再试，或在ISV交流群里咨询技术支持或帮助中心提单 |
| 500_2 | invalid parameter error | 请求参数无效，请在ISV交流群里咨询技术支持或帮助中心提单 |
| 500_2 | user order not exist error | 订单号对应的订单不是当前授权用户的订单，无权操作 |
| 500_2 | order not exist error | 订单号不存在对应订单，请检查订单号是否正确 |

## 示例

**请求示例**

```
{"feedback":"test","orderId":"159095856057498520"}
```

**返回示例**

```
{"result":{"success":true},"success":true}
```
