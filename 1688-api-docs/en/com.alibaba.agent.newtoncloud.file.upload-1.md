# Newton Cloud: upload file

Original name: 牛顿云-上传文件  
API: `com.alibaba.agent:newtoncloud.file.upload:1` · Category: Inquiries (Newton Cloud)  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.agent:newtoncloud.file.upload-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.agent/newtoncloud.file.upload/{appKey}`  
Requires user authorization (access_token) · Requires signature

Upload a local file to Newton Cloud OSS. Returns a temporary public download link and a stable download address, which can be used directly as the fileUrls parameter of newtoncloud.task.create. File content is passed as Base64; a single file of 3 MB or less is recommended.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `contentBase64` | java.lang.String | yes | Base64-encoded string of the file content. Supports standard and URL-safe encoding, and a data URI prefix; the original file is recommended not to exceed 3MB. | iVBORw0KGgoAAAANSUhEUgAA... |
| `filename` | String | yes | File name, must include the extension | 报价单.xlsx |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.cbu.selleragent.api.newton.dto.NewtonCloudUploadResult](#m-alibaba-cbu-selleragent-api-newton-dto-newtonclouduploadresult) | yes | Newton Cloud file upload result | {} |

<a id="m-alibaba-cbu-selleragent-api-newton-dto-newtonclouduploadresult"></a>
#### alibaba.cbu.selleragent.api.newton.dto.NewtonCloudUploadResult

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `downloadUrl` | java.lang.String | yes | Stable download URL, valid long-term; requires logging into the corresponding 1688 account to access | https://selleragent.1688.com/api/newton/cloud/files/download?relativePath=uploads%2F%E6%8A%A5%E4%BB%B7%E.xlsx |
| `eagleTraceId` | java.lang.String | yes | Trace ID | 213e1a2f17220000000001 |
| `fileSize` | long | yes | File size, unit: byte | 102400 |
| `relativePath` | java.lang.String | yes | The relative path of the file in the Newton Cloud workspace, which can be used for subsequent operations such as requesting a download link | uploads/报价单.xlsx |
| `agentDownloadUrl` | String | yes | Login-free public network download address, which can be used directly for the fileUrls of task.create. | https://selleragent.1688.com/api/seller/knowledge/file/download/xxx |
