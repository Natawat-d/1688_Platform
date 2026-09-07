# 查询订单可以支持的支付渠道

API: `com.alibaba.trade:alibaba.trade.payWay.query:1` · Category: 支付  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.trade.payWay.query-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.trade.payWay.query/{appKey}`  
需要授权 (access_token) · 需要签名

查询未支付订单可以使用的支付方式或者支付渠道

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `orderId` | String | 是 | 订单号 | 123123 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `success` | String | 是 | 是否成功 | true |
| `errorCode` | String | 是 | 错误码 | 500_1 |
| `errorMsg` | String | 是 | 错误信息 |   |
| `resultList` | [message:alibaba.ocean.openplatform.biz.trade.result.TradePayTypeResult](#m-alibaba-ocean-openplatform-biz-trade-result-tradepaytyperesult) | 是 | 返回结果 | {} |

<a id="m-alibaba-ocean-openplatform-biz-trade-result-tradepaytyperesult"></a>
#### alibaba.ocean.openplatform.biz.trade.result.TradePayTypeResult

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `channels` | [message:alibaba.ocean.openplatform.biz.trade.result.PayTypeInfo[]](#m-alibaba-ocean-openplatform-biz-trade-result-paytypeinfo[]) | 是 | 可用支付渠道列表 | [] |
| `orderId` | java.lang.String | 是 | 订单号 | 1231231211 |
| `payFee` | java.lang.Long | 是 | 支付金额，单位分 | 100 |
| `timeout` | java.lang.String | 是 | 最晚支付时间 | 2018-10-01 00:00:00 |

<a id="m-alibaba-ocean-openplatform-biz-trade-result-paytypeinfo[]"></a>
#### alibaba.ocean.openplatform.biz.trade.result.PayTypeInfo[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `code` | java.lang.Long | 是 | 支付渠道编码，1:支付宝 2:网商银行信任付 3:诚e赊 4:对公转账 5:赊销宝 6:电子承兑票据 7:账期支付 8:合并支付渠道 9:无打款 10:零售通赊购 12:声明付款 13:支付平台 14:网商电子银行承兑汇票 15:银行转账 16:跨境宝 17:红包 20:跨境宝 35:网商银行跨境直采 | 1 |
| `name` | java.lang.String | 是 | 支付渠道名称，1:支付宝 2:网商银行信任付 3:诚e赊 4:对公转账 5:赊销宝 6:电子承兑票据 7:账期支付 8:合并支付渠道 9:无打款 10:零售通赊购 12:声明付款 13:支付平台 14:网商电子银行承兑汇票 15:银行转账 16:跨境宝 17:红包 20:跨境宝 35:网商银行跨境直采 | 支付宝 |

## 错误码

| 错误码 | 现象 | 解决方案 |
|---|---|---|
| 500_2 | 没有权限获取该订单可支付方式。 | 检查授权用户，授权账号为买家，且必须为买家主账号。 |

## 示例

**返回结果示例**

```
{
  "resultList": {
    "channels": [
      {
        "code": 1,
        "name": "支付宝"
      }
    ],
    "orderId": "239695213738498520",
    "payFee": 32120,
    "timeout": "2018-11-04 14:00:45"
  },
  "success": true
}
```
