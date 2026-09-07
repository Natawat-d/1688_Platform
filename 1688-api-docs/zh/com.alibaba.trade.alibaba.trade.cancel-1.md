# 取消交易

API: `com.alibaba.trade:alibaba.trade.cancel:1` · Category: 订单  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.trade.cancel-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.trade.cancel/{appKey}`  
需要授权 (access_token) · 需要签名

买家或者卖家取消交易，注意只有特定状态的交易才能取消，1688可用于取消未付款的订单。
当订单从创建到关闭时间小于10s的时候，会报“CLOSE_ORDER_TOO_FAST”错误。

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `webSite` | String | 是 | 站点信息，指定调用的API是属于国际站（alibaba）还是1688网站（1688） | 1688 |
| `tradeID` | Long | 是 | 交易id，订单号 | 123456 |
| `cancelReason` | String | 是 | 原因描述；buyerCancel:买家取消订单;sellerGoodsLack:卖家库存不足;other:其它 | other |
| `remark` | String | 否 | 备注 | 备注 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `success` | Boolean | 是 | 是否处理成功：true为成功，false为失败，失败原因见error | true |
| `errorCode` | String | 是 | 错误码 | ORDER_STATUS_ERROR |
| `errorMessage` | String | 是 | 错误信息 | 订单状态错误 |

## 错误码

| 错误码 | 现象 | 解决方案 |
|---|---|---|
| ORDER_STATUS_ERROR | 订单状态错误 | 只能取消待支付的订单，需要先确认订单状态 |
| 400_3 | 没有权限取消该订单 | 只有订单的买卖双方才能取消订单，确认授权用户是否该订单的买卖双方 |
| ORDER_NOT_EXIST | 订单不存在 | 确认订单号是否正确 |

## 示例

**请求参数示例**

```
{"webSite":"1688","tradeID":"202711458975969812","cancelReason":"other","remark":"取消订单"}
```

**返回参数示例**

```
{"success":true}
```
