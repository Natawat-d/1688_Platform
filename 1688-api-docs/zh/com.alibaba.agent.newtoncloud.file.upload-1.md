# 牛顿云-上传文件

API: `com.alibaba.agent:newtoncloud.file.upload:1` · Category: 询盘  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.agent:newtoncloud.file.upload-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.agent/newtoncloud.file.upload/{appKey}`  
需要授权 (access_token) · 需要签名

上传本地文件到牛顿云 OSS，返回临时公网下载链接与稳定下载地址，可直接作为newtoncloud.task.create 的 fileUrls 参数使用。文件内容以 Base64 传入，建议单文件 ≤3MB

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `contentBase64` | java.lang.String | 是 | 文件内容的 Base64 编码字符串，支持标准与 URL-safe 编码、data URI 前缀；建议原文件不超过 3MB | iVBORw0KGgoAAAANSUhEUgAA... |
| `filename` | String | 是 | 文件名，需包含扩展名 | 报价单.xlsx |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:alibaba.cbu.selleragent.api.newton.dto.NewtonCloudUploadResult](#m-alibaba-cbu-selleragent-api-newton-dto-newtonclouduploadresult) | 是 | 牛顿云文件上传结果 | {} |

<a id="m-alibaba-cbu-selleragent-api-newton-dto-newtonclouduploadresult"></a>
#### alibaba.cbu.selleragent.api.newton.dto.NewtonCloudUploadResult

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `downloadUrl` | java.lang.String | 是 | 稳定下载地址，长期有效，需登录对应的 1688 账号后访问 | https://selleragent.1688.com/api/newton/cloud/files/download?relativePath=uploads%2F%E6%8A%A5%E4%BB%B7%E.xlsx |
| `eagleTraceId` | java.lang.String | 是 | 链路追踪 ID | 213e1a2f17220000000001 |
| `fileSize` | long | 是 | 文件大小，单位字节 | 102400 |
| `relativePath` | java.lang.String | 是 | 文件在牛顿云工作区的相对路径，可用于后续申请下载链接等操作 | uploads/报价单.xlsx |
| `agentDownloadUrl` | String | 是 | 免登录公网下载地址，可直接用于 task.create 的 fileUrls。 | https://selleragent.1688.com/api/seller/knowledge/file/download/xxx |
