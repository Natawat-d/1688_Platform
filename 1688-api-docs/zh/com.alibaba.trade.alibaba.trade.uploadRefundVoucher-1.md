# 上传退款退货凭证

API: `com.alibaba.trade:alibaba.trade.uploadRefundVoucher:1` · Category: 退货退款  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.trade.uploadRefundVoucher-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.trade.uploadRefundVoucher/{appKey}`  
需要授权 (access_token) · 需要签名

上传退款退货凭证，用于退款退货申请，文件流转byte数组推荐使用org.apache.commons.io.IOUtils#toByteArray(java.io.InputStream)

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `imageData` | byte[] | 是 | 凭证图片数据。小于1M，jpg格式。 |   |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:alibaba.ocean.openplatform.common.OrderRefundUploadVoucherResult](#m-alibaba-ocean-openplatform-common-orderrefunduploadvoucherresult) | 是 | 返回结果 |   |

<a id="m-alibaba-ocean-openplatform-common-orderrefunduploadvoucherresult"></a>
#### alibaba.ocean.openplatform.common.OrderRefundUploadVoucherResult

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `code` | String | 是 | 错误码 |   |
| `message` | String | 是 | 错误信息 |   |
| `result` | [message:alibaba.ocean.openplatform.biz.trade.result.OrderRefundUploadVoucherResult](#m-alibaba-ocean-openplatform-biz-trade-result-orderrefunduploadvoucherresult) | 是 | 成功结果 |   |
| `success` | Boolean | 是 | 是否成功 |   |

<a id="m-alibaba-ocean-openplatform-biz-trade-result-orderrefunduploadvoucherresult"></a>
#### alibaba.ocean.openplatform.biz.trade.result.OrderRefundUploadVoucherResult

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `imageDomain` | java.lang.String | 是 | 图片域名 |   |
| `imageRelativeUrl` | java.lang.String | 是 | 图片相对路径 |   |

## 错误码

| 错误码 | 现象 | 解决方案 |
|---|---|---|
| 4000 |  | 调用者错误：未正确传入imageData参数；<br>服务方错误：网络延迟、内部服务异常 |

## 示例

**通过表单提交示例**

```
 <!--参数：appkeyValue, signatureValue,tokenValue-->
 <!--imageData部分不参与签名-->

<form name="test" action="https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.trade.uploadRefundVoucher/appkeyValue?access_token=tokenValue&_aop_signature=signatureValue"  method="post" enctype ="multipart/form-data">
    <input type="file" name="imageData">
    <input type="submit" value="提交">
</form>
```
