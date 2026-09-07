# 海外地址解析

API: `com.alibaba.trade:trade.address.parseText:1` · Category: 物流  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:trade.address.parseText-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/trade.address.parseText/{appKey}`  
需要授权 (access_token) · 需要签名

海外地址解析

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `param` | [message:alibaba.global.trade.api.open.order.param.OrderAddressParseParam](#m-alibaba-global-trade-api-open-order-param-orderaddressparseparam) | 是 | 入参 | {} |

<a id="m-alibaba-global-trade-api-open-order-param-orderaddressparseparam"></a>
#### alibaba.global.trade.api.open.order.param.OrderAddressParseParam

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `area` | java.lang.String | 是 | 区 | Кокшетау |
| `city` | java.lang.String | 是 | 市 | Г.Кокшетау |
| `country` | java.lang.String | 是 | 国家 | Казахстан |
| `province` | java.lang.String | 是 | 省 | Акмолинская область |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:alibaba.global.trade.api.common.model.ResultModel](#m-alibaba-global-trade-api-common-model-resultmodel) | 是 | 结果 | {} |

<a id="m-alibaba-global-trade-api-common-model-resultmodel"></a>
#### alibaba.global.trade.api.common.model.ResultModel

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `success` | java.lang.Boolean | 是 | 是否成功 | true |
| `code` | java.lang.String | 是 | 返回码 | PARAM_CAN_NOT_BE_NULL |
| `message` | java.lang.String | 是 | 附带信息 | PARAM_CAN_NOT_BE_NULL |
| `data` | [message:alibaba.global.trade.api.open.order.model.OrderAddressParseModel](#m-alibaba-global-trade-api-open-order-model-orderaddressparsemodel) | 是 | 数据 | {} |

<a id="m-alibaba-global-trade-api-open-order-model-orderaddressparsemodel"></a>
#### alibaba.global.trade.api.open.order.model.OrderAddressParseModel

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `addressCode` | java.lang.String | 是 | 末级地址code | 910800030036000000 |
| `area` | java.lang.String | 是 | 区 | Кокшетау |
| `city` | java.lang.String | 是 | 市 | Г.Кокшетау |
| `country` | java.lang.String | 是 | 国家 | Казахстан |
| `province` | java.lang.String | 是 | 省 | Акмолинская область |

## 示例

**入参示例**

```
{
  "param": {
    "area": "Кокшетау",
    "city": "Г.Кокшетау",
    "country": "Казахстан",
    "province": "Акмолинская область"
  }
}
```

**出参示例**

```
{
  "result": {
    "success": true,
    "code": "PARAM_CAN_NOT_BE_NULL",
    "message": "PARAM_CAN_NOT_BE_NULL",
    "data": {
      "addressCode": "910800030036000000",
      "area": "Кокшетау",
      "city": "Г.Кокшетау",
      "country": "Казахстан",
      "province": "Акмолинская область"
    }
  }
}
```
