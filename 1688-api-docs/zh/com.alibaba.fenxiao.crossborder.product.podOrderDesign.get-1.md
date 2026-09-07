# 获取加工定制订单稿件信息

API: `com.alibaba.fenxiao.crossborder:product.podOrderDesign.get:1` · Category: 轻定制  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:product.podOrderDesign.get-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/product.podOrderDesign.get/{appKey}`  
需要授权 (access_token) · 需要签名

获取加工定制订单稿件信息

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `orderDesignParam` | [message:com.alibaba.cbu.pod.design.param.OrderDesignOpenApiParam](#m-com-alibaba-cbu-pod-design-param-orderdesignopenapiparam) | 是 | 订单设计查询参数 | {"orderId":123} |

<a id="m-com-alibaba-cbu-pod-design-param-orderdesignopenapiparam"></a>
#### com.alibaba.cbu.pod.design.param.OrderDesignOpenApiParam

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `orderId` | java.lang.Long | 是 | 1688订单id | 123 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:com.alibaba.openapi.shared.common.ResultModel](#m-com-alibaba-openapi-shared-common-resultmodel) | 是 | 返回结果 | {} |

<a id="m-com-alibaba-openapi-shared-common-resultmodel"></a>
#### com.alibaba.openapi.shared.common.ResultModel

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `success` | boolean | 是 | 是否成功 | true |
| `code` | java.lang.String | 是 | 响应码 | S0000 |
| `message` | java.lang.String | 是 | 错误信息 | 成功 |
| `result` | [message:com.alibaba.cbu.pod.design.model.OrderDesignOpenResultModel](#m-com-alibaba-cbu-pod-design-model-orderdesignopenresultmodel) | 是 | 返回结果 | {} |

<a id="m-com-alibaba-cbu-pod-design-model-orderdesignopenresultmodel"></a>
#### com.alibaba.cbu.pod.design.model.OrderDesignOpenResultModel

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `orderId` | java.lang.Long | 是 | 订单id | 123 |
| `designId` | java.lang.String | 是 | 设计稿id | designId1 |
| `regionInfo` | [message:com.alibaba.cbu.pod.design.model.DesignImageModel[]](#m-com-alibaba-cbu-pod-design-model-designimagemodel[]) | 是 | 设计稿图片信息 | [{"region":"正面","maskImageUrl":"https://"}] |

<a id="m-com-alibaba-cbu-pod-design-model-designimagemodel[]"></a>
#### com.alibaba.cbu.pod.design.model.DesignImageModel[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `regionName` | java.lang.String | 是 | 设计稿图片区域名字(中文) | 正面 |
| `maskImageUrl` | java.lang.String | 是 | 设计稿图片-maskImageUrl和previewImageUrl同时不为空，则进行了该区域的定制 | https:/xxx |
| `previewImageUrl` | java.lang.String | 是 | 预览图-maskImageUrl和previewImageUrl同时不为空，则进行了该区域的定制 | https:/xxx |

## 示例

**入参示例**

```
{
  "orderDesignParam": {
    "orderId": 123
  }
}
```

**出参示例**

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
