# Parse overseas address

Original name: 海外地址解析  
API: `com.alibaba.trade:trade.address.parseText:1` · Category: Logistics  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:trade.address.parseText-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/trade.address.parseText/{appKey}`  
Requires user authorization (access_token) · Requires signature

Parse an overseas address.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `param` | [message:alibaba.global.trade.api.open.order.param.OrderAddressParseParam](#m-alibaba-global-trade-api-open-order-param-orderaddressparseparam) | yes | Input parameter | {} |

<a id="m-alibaba-global-trade-api-open-order-param-orderaddressparseparam"></a>
#### alibaba.global.trade.api.open.order.param.OrderAddressParseParam

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `area` | java.lang.String | yes | District | Кокшетау |
| `city` | java.lang.String | yes | City | Г.Кокшетау |
| `country` | java.lang.String | yes | Country | Казахстан |
| `province` | java.lang.String | yes | Province | Акмолинская область |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.global.trade.api.common.model.ResultModel](#m-alibaba-global-trade-api-common-model-resultmodel) | yes | Result | {} |

<a id="m-alibaba-global-trade-api-common-model-resultmodel"></a>
#### alibaba.global.trade.api.common.model.ResultModel

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `success` | java.lang.Boolean | yes | Whether successful | true |
| `code` | java.lang.String | yes | Return code | PARAM_CAN_NOT_BE_NULL |
| `message` | java.lang.String | yes | Attached information | PARAM_CAN_NOT_BE_NULL |
| `data` | [message:alibaba.global.trade.api.open.order.model.OrderAddressParseModel](#m-alibaba-global-trade-api-open-order-model-orderaddressparsemodel) | yes | Data | {} |

<a id="m-alibaba-global-trade-api-open-order-model-orderaddressparsemodel"></a>
#### alibaba.global.trade.api.open.order.model.OrderAddressParseModel

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `addressCode` | java.lang.String | yes | Lowest-level address code | 910800030036000000 |
| `area` | java.lang.String | yes | District | Кокшетау |
| `city` | java.lang.String | yes | City | Г.Кокшетау |
| `country` | java.lang.String | yes | Country | Казахстан |
| `province` | java.lang.String | yes | Province | Акмолинская область |

## Samples

**Input parameter example**

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

**Output parameter example**

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
