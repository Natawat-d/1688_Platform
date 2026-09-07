# 上传图片获取imageId

API: `com.alibaba.fenxiao.crossborder:product.image.upload:1` · Category: 商品  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:product.image.upload-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/product.image.upload/{appKey}`  
需要授权 (access_token) · 需要签名

上传图片获取imageId

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `uploadImageParam` | [message:product.image.upload.param.UploadImageParam](#m-product-image-upload-param-uploadimageparam) | 是 |  |  |

<a id="m-product-image-upload-param-uploadimageparam"></a>
#### product.image.upload.param.UploadImageParam

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `imageBase64` | java.lang.String | 是 | 图片base64 | 12 |
| `outMemberId` | java.lang.String | 否 | 外部用户id | 2 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:product.image.upload.result.ResultModel](#m-product-image-upload-result-resultmodel) | 是 |  |  |

<a id="m-product-image-upload-result-resultmodel"></a>
#### product.image.upload.result.ResultModel

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `success` | java.lang.String | 是 |  |  |
| `code` | java.lang.String | 是 |  |  |
| `message` | java.lang.String | 是 |  |  |
| `result` | java.lang.String | 是 |  |  |
