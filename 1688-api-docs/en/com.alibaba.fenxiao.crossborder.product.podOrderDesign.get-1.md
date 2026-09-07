# Get artwork info for a customization order

Original name: 获取加工定制订单稿件信息  
API: `com.alibaba.fenxiao.crossborder:product.podOrderDesign.get:1` · Category: Light Customization  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:product.podOrderDesign.get-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/product.podOrderDesign.get/{appKey}`  
Requires user authorization (access_token) · Requires signature

Get the design/artwork information of a processing-and-customization (print-on-demand) order.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `orderDesignParam` | [message:com.alibaba.cbu.pod.design.param.OrderDesignOpenApiParam](#m-com-alibaba-cbu-pod-design-param-orderdesignopenapiparam) | yes | Order design query parameters | {"orderId":123} |

<a id="m-com-alibaba-cbu-pod-design-param-orderdesignopenapiparam"></a>
#### com.alibaba.cbu.pod.design.param.OrderDesignOpenApiParam

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `orderId` | java.lang.Long | yes | 1688 order id | 123 |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:com.alibaba.openapi.shared.common.ResultModel](#m-com-alibaba-openapi-shared-common-resultmodel) | yes | Return result | {} |

<a id="m-com-alibaba-openapi-shared-common-resultmodel"></a>
#### com.alibaba.openapi.shared.common.ResultModel

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `success` | boolean | yes | Whether successful | true |
| `code` | java.lang.String | yes | Response code | S0000 |
| `message` | java.lang.String | yes | Error message | 成功 |
| `result` | [message:com.alibaba.cbu.pod.design.model.OrderDesignOpenResultModel](#m-com-alibaba-cbu-pod-design-model-orderdesignopenresultmodel) | yes | Return result | {} |

<a id="m-com-alibaba-cbu-pod-design-model-orderdesignopenresultmodel"></a>
#### com.alibaba.cbu.pod.design.model.OrderDesignOpenResultModel

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `orderId` | java.lang.Long | yes | Order id | 123 |
| `designId` | java.lang.String | yes | Design draft id | designId1 |
| `regionInfo` | [message:com.alibaba.cbu.pod.design.model.DesignImageModel[]](#m-com-alibaba-cbu-pod-design-model-designimagemodel[]) | yes | Design draft image information | [{"region":"正面","maskImageUrl":"https://"}] |

<a id="m-com-alibaba-cbu-pod-design-model-designimagemodel[]"></a>
#### com.alibaba.cbu.pod.design.model.DesignImageModel[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `regionName` | java.lang.String | yes | Design draft image area name (Chinese) | 正面 |
| `maskImageUrl` | java.lang.String | yes | Design draft image - if maskImageUrl and previewImageUrl are both non-empty, customization was performed for this area. | https:/xxx |
| `previewImageUrl` | java.lang.String | yes | Preview image - if both maskImageUrl and previewImageUrl are not empty, customization has been done for this area. | https:/xxx |

## Samples

**Input parameter example**

```
{
  "orderDesignParam": {
    "orderId": 123
  }
}
```

**Output parameter example**

```
{
  "result": {
    "success": true,
    "code": "S0000",
    "message": "成功",
    "result": {
      "orderId": 123,
      "designId": "designId1",
      "regionInfo": [
        {
          "code": "front",
          "regionName": "正面",
          "maskImageUrl": "https:/xxx",
          "previewImageUrl": "https:/xxx"
        }
      ]
    }
  }
}
```
