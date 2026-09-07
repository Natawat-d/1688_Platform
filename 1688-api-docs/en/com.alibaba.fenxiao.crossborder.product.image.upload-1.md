# Upload an image to get an imageId

Original name: 上传图片获取imageId  
API: `com.alibaba.fenxiao.crossborder:product.image.upload:1` · Category: Products  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:product.image.upload-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/product.image.upload/{appKey}`  
Requires user authorization (access_token) · Requires signature

Upload an image and receive an imageId.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `uploadImageParam` | [message:product.image.upload.param.UploadImageParam](#m-product-image-upload-param-uploadimageparam) | yes |  |  |

<a id="m-product-image-upload-param-uploadimageparam"></a>
#### product.image.upload.param.UploadImageParam

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `imageBase64` | java.lang.String | yes | Image base64 | 12 |
| `outMemberId` | java.lang.String | no | External user ID | 2 |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:product.image.upload.result.ResultModel](#m-product-image-upload-result-resultmodel) | yes |  |  |

<a id="m-product-image-upload-result-resultmodel"></a>
#### product.image.upload.result.ResultModel

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `success` | java.lang.String | yes |  |  |
| `code` | java.lang.String | yes |  |  |
| `message` | java.lang.String | yes |  |  |
| `result` | java.lang.String | yes |  |  |
