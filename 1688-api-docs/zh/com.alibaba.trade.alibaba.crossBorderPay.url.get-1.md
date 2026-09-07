# 获取使用跨境宝支付的支付链接

API: `com.alibaba.trade:alibaba.crossBorderPay.url.get:1` · Category: 支付  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.crossBorderPay.url.get-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.crossBorderPay.url.get/{appKey}`  
需要授权 (access_token) · 需要签名

获取使用跨境宝支付的支付链接

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `orderIdList` | Long[] | 是 | 订单Id列表,最多批量30个订单，订单过多会导致超时，建议一次10个订单 | [111111,22222333] |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `success` | String | 是 | 是否成功 | true |
| `errorCode` | String | 是 | 错误码 | 400_1 |
| `errorMsg` | String | 是 | 错误描述 |   |
| `payUrl` | String | 是 | 收银台支付链接 | https://trade.1688.com/order/cashier.htm?orderId=15405143260 |
| `cantPayOrderList` | Long[] | 是 | 由于额度及风控原因不能批量支付的订单列表 | [123123,23123123] |

## 错误码

| 错误码 | 现象 | 解决方案 |
|---|---|---|
| 400_4 | 无可使用支付渠道[跨境宝]付款的订单 | 订单列表里面没有可以使用跨境宝支付的订单 |

## 示例

**返回示例**

```
{
  "payUrl": "https://trade.1688.com/order/cashier.htm?orderId=151923545498520",
  "cantPayOrderList":[12123123,12312222222]
  "success": true
}
```
