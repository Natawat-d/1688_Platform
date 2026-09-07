# 查询商家退货地址

API: `com.alibaba.fenxiao:refund.address.get:1` · Category: 退货退款  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao:refund.address.get-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao/refund.address.get/{appKey}`  
需要授权 (access_token) · 需要签名

查询退货地址

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `mainOrderId` | String | 是 | 主订单id | 123 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:com.alibaba.distributionhub.common.model.RefundAddressResult](#m-com-alibaba-distributionhub-common-model-refundaddressresult) | 是 | 结果 | 1 |

<a id="m-com-alibaba-distributionhub-common-model-refundaddressresult"></a>
#### com.alibaba.distributionhub.common.model.RefundAddressResult

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `data` | [message:com.alibaba.distributionhub.refund.model.RefundAddressInfoModel[]](#m-com-alibaba-distributionhub-refund-model-refundaddressinfomodel[]) | 是 | 退货地址列表 | 1 |
| `errorCode` | String | 是 | 错误编码 | 101 |
| `errorInfo` | String | 是 | 错误信息 | 系统异常 |
| `success` | Boolean | 是 | 接口调用结果 | true |

<a id="m-com-alibaba-distributionhub-refund-model-refundaddressinfomodel[]"></a>
#### com.alibaba.distributionhub.refund.model.RefundAddressInfoModel[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `cityText` | java.lang.String | 是 | 市文本 | 1 |
| `contactPhone` | java.lang.String | 是 | 联系电话（手机或固话） | 1 |
| `districtText` | java.lang.String | 是 | 区文本 | 1 |
| `mobile` | java.lang.String | 是 | 手机 | 1 |
| `offerId` | java.lang.Long | 是 | 商品id | 1 |
| `phone` | java.lang.String | 是 | 固话 | 1 |
| `provinceText` | java.lang.String | 是 | 省份文本 | 1 |
| `receiverName` | java.lang.String | 是 | 收货人姓名 | 1 |
| `streetAddress` | java.lang.String | 是 | 街道地址 | 1 |
| `streetAddressCode` | java.lang.String | 是 | 街道地址编码 | 1 |
| `townCode` | java.lang.String | 是 | 镇地址编码 | 1 |
| `townText` | java.lang.String | 是 | 镇文本 | 1 |
| `zipCode` | java.lang.String | 是 | 邮编 | 1 |

## 示例

**入参示例**

```
{
  "mainOrderId": "123"
}
```

**出参示例**

```
{
  "result": {
    "data": [
      {
        "cityText": "1",
        "contactPhone": "1",
        "districtText": "1",
        "mobile": "1",
        "offerId": 1,
        "phone": "1",
        "provinceText": "1",
        "receiverName": "1",
        "streetAddress": "1",
        "streetAddressCode": "1",
        "townCode": "1",
        "townText": "1",
        "zipCode": "1"
      }
    ],
    "errorCode": "101",
    "errorInfo": "系统异常",
    "success": true
  }
}
```
