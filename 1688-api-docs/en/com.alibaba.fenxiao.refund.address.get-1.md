# Query the merchant's return address

Original name: 查询商家退货地址  
API: `com.alibaba.fenxiao:refund.address.get:1` · Category: Returns & Refunds  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao:refund.address.get-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao/refund.address.get/{appKey}`  
Requires user authorization (access_token) · Requires signature

Query the return address.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `mainOrderId` | String | yes | Main order id | 123 |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:com.alibaba.distributionhub.common.model.RefundAddressResult](#m-com-alibaba-distributionhub-common-model-refundaddressresult) | yes | Result | 1 |

<a id="m-com-alibaba-distributionhub-common-model-refundaddressresult"></a>
#### com.alibaba.distributionhub.common.model.RefundAddressResult

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `data` | [message:com.alibaba.distributionhub.refund.model.RefundAddressInfoModel[]](#m-com-alibaba-distributionhub-refund-model-refundaddressinfomodel[]) | yes | Return address list | 1 |
| `errorCode` | String | yes | Error code | 101 |
| `errorInfo` | String | yes | Error message | 系统异常 |
| `success` | Boolean | yes | API call result | true |

<a id="m-com-alibaba-distributionhub-refund-model-refundaddressinfomodel[]"></a>
#### com.alibaba.distributionhub.refund.model.RefundAddressInfoModel[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `cityText` | java.lang.String | yes | City text | 1 |
| `contactPhone` | java.lang.String | yes | Contact phone number (mobile or landline) | 1 |
| `districtText` | java.lang.String | yes | District text | 1 |
| `mobile` | java.lang.String | yes | Mobile phone | 1 |
| `offerId` | java.lang.Long | yes | Product ID | 1 |
| `phone` | java.lang.String | yes | Landline | 1 |
| `provinceText` | java.lang.String | yes | Province text | 1 |
| `receiverName` | java.lang.String | yes | Recipient name | 1 |
| `streetAddress` | java.lang.String | yes | Street address | 1 |
| `streetAddressCode` | java.lang.String | yes | Street address code | 1 |
| `townCode` | java.lang.String | yes | Town address code | 1 |
| `townText` | java.lang.String | yes | Town text | 1 |
| `zipCode` | java.lang.String | yes | Postal code | 1 |

## Samples

**Input parameter example**

```
{
  "mainOrderId": "123"
}
```

**Output parameter example**

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
