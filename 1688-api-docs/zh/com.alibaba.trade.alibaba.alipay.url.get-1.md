# 批量获取订单的支付链接

API: `com.alibaba.trade:alibaba.alipay.url.get:1` · Category: 支付  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.alipay.url.get-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.alipay.url.get/{appKey}`  
需要授权 (access_token) · 需要签名

通过ERP付款时，可以通过本API获取批量支付的收银台的链接。
单个订单返回1688收银台地址，多个订单返回支付宝收银台地址。
ERP可以引导用户跳转到收银台链接完成支付动作，支付前会校验用户在1688的登陆状态。

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `orderIdList` | Long[] | 是 | 订单Id列表,最多批量100个订单，跨境宝批量最大只支持30笔。 | [74321349391498520] |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `erroMsg` | String | 是 | 错误信息 |   |
| `payUrl` | String | 是 | 支付链接 | https://a.b.com |
| `success` | Boolean | 是 | 是否成功，可能部分成功，需要结合payFailureOrderList查看 | true |
| `errorCode` | String | 是 | 错误码 |   |
| `payFailureOrderList` | Long[] | 是 | 部分失败订单id | [1299871823,19798172783] |

## 错误码

| 错误码 | 现象 | 解决方案 |
|---|---|---|
| Batch pay : not surport MANUAL-TRADE! | Batch pay : not surport MANUAL-TRADE! | 未补充买家收件信息的邀约订单不能合并付款，是否邀约订单可以通过订单详情里面的baseInfo.sellerOrder字段判断 |
| 操作库存失败:PRODUCT_TRADE_STAT_ERROR | inventoryErrorIds:[16397675**722128**] | 存在付款减库存订单且扣减库存失败。 |

## 示例

**出参：多个订单的支付链接**

```
{
  "payUrl": "https://trade.1688.com/order/cashier.htm?orderId=154051432607498520;151923545459498520",
  "success": true
}
```

**出参：单个订单支付链接**

```
{
  "payUrl": "https://trade.1688.com/order/cashier.htm?orderId=151923545459498520",
  "success": true
}
```
