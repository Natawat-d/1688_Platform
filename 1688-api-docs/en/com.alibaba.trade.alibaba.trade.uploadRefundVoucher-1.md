# Upload refund/return evidence

Original name: 上传退款退货凭证  
API: `com.alibaba.trade:alibaba.trade.uploadRefundVoucher:1` · Category: Returns & Refunds  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.trade.uploadRefundVoucher-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.trade.uploadRefundVoucher/{appKey}`  
Requires user authorization (access_token) · Requires signature

Upload evidence for a refund/return request. To convert a file stream to a byte array, org.apache.commons.io.IOUtils#toByteArray(java.io.InputStream) is recommended.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `imageData` | byte[] | yes | Voucher image data. Less than 1MB, jpg format. |   |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.ocean.openplatform.common.OrderRefundUploadVoucherResult](#m-alibaba-ocean-openplatform-common-orderrefunduploadvoucherresult) | yes | Return result |   |

<a id="m-alibaba-ocean-openplatform-common-orderrefunduploadvoucherresult"></a>
#### alibaba.ocean.openplatform.common.OrderRefundUploadVoucherResult

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `code` | String | yes | Error code |   |
| `message` | String | yes | Error message |   |
| `result` | [message:alibaba.ocean.openplatform.biz.trade.result.OrderRefundUploadVoucherResult](#m-alibaba-ocean-openplatform-biz-trade-result-orderrefunduploadvoucherresult) | yes | Success result |   |
| `success` | Boolean | yes | Whether successful |   |

<a id="m-alibaba-ocean-openplatform-biz-trade-result-orderrefunduploadvoucherresult"></a>
#### alibaba.ocean.openplatform.biz.trade.result.OrderRefundUploadVoucherResult

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `imageDomain` | java.lang.String | yes | Image domain |   |
| `imageRelativeUrl` | java.lang.String | yes | Image relative path |   |

## Error codes

| Error | Symptom | How to fix |
|---|---|---|
| 4000 |  | Caller error: the imageData parameter was not passed correctly;<br>Service-side error: network latency, internal service exception |

## Samples

**Example of submission via form**

```
 <!--参数：appkeyValue, signatureValue,tokenValue-->
 <!--imageData部分不参与签名-->

<form name="test" action="https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.trade.uploadRefundVoucher/appkeyValue?access_token=tokenValue&_aop_signature=signatureValue"  method="post" enctype ="multipart/form-data">
    <input type="file" name="imageData">
    <input type="submit" value="提交">
</form>
```
