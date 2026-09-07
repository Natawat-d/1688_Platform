# 获取商品每日销售数量趋势（新）

API: `com.alibaba.fenxiao.crossborder:product.analyze.getPerdaySellQuantityTrendNew:1` · Category: 商机  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:product.analyze.getPerdaySellQuantityTrendNew-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/product.analyze.getPerdaySellQuantityTrendNew/{appKey}`  
需要授权 (access_token) · 需要签名

获取商品90天每日销售数量趋势（最多查询90天数据），新接口

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `tendQueryParam` | [message:com.alibaba.cbu.offer.param.OfferSellTrendQueryParam](#m-com-alibaba-cbu-offer-param-offerselltrendqueryparam) | 是 |  |  |

<a id="m-com-alibaba-cbu-offer-param-offerselltrendqueryparam"></a>
#### com.alibaba.cbu.offer.param.OfferSellTrendQueryParam

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `offerId` | java.lang.Long | 是 | 商品id | 123 |
| `startDate` | java.lang.String | 是 | 查询起始时间 | 20240701(起始时间和截止时间不要超过1个月） |
| `endDate` | java.lang.String | 是 | 查询截止时间 | 20240704 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:com.alibaba.cbu.common.result.CommonResult](#m-com-alibaba-cbu-common-result-commonresult) | 是 |  |  |

<a id="m-com-alibaba-cbu-common-result-commonresult"></a>
#### com.alibaba.cbu.common.result.CommonResult

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `success` | java.lang.Boolean | 是 | 是否成功 | true |
| `retCode` | java.lang.String | 是 | 错误码 | S000 |
| `retMsg` | java.lang.String | 是 | 返回信息 | 成功 |
| `result` | [message:com.alibaba.cbu.offer.model.OfferSellTrendDataModel[]](#m-com-alibaba-cbu-offer-model-offerselltrenddatamodel[]) | 是 | 结果 | [     {       "date": "20240701",       "value": "20"     },     {       "date": "20240702",       "value": "20"     },     {       "date": "20240703",       "value": "20"     },     {       "date": "20240704",       "value": "20"     },     {       "date": "20240705",       "value": ""     } |

<a id="m-com-alibaba-cbu-offer-model-offerselltrenddatamodel[]"></a>
#### com.alibaba.cbu.offer.model.OfferSellTrendDataModel[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `date` | java.lang.String | 是 | 数据指标对应的日期 | 20240620 |
| `value` | java.lang.String | 是 | 指标数据 | 100 |

## 示例

**入参示例**

```
{
  "tendQueryParam": {
    "offerId": 123,
    "startDate": "20240701(起始时间和截止时间不要超过1个月）",
    "endDate": "20240704"
  }
}
```

**出参示例**

```
{
  "result": {
    "success": true,
    "retCode": "S000",
    "retMsg": "成功",
    "result": [
      {
        "date": "20240620",
        "value": "100"
      }
    ]
  }
}
```
