# Parse an address into area codes

Original name: 根据地址解析地区码  
API: `com.alibaba.trade:alibaba.trade.addresscode.parse:1` · Category: Logistics  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.trade.addresscode.parse-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.trade.addresscode.parse/{appKey}`  
Requires user authorization (access_token) · Requires signature

Parse area codes from address information.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `addressInfo` | String | yes | Address information | 浙江省 杭州市 滨江区网商路699号 |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.trade.ReceiveAddress](#m-alibaba-trade-receiveaddress) | yes | Parsed shipping address | {} |
| `errorCode` | String | yes | Error code |   |
| `errorMessage` | String | yes | Error message |   |

<a id="m-alibaba-trade-receiveaddress"></a>
#### alibaba.trade.ReceiveAddress

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `address` | java.lang.String | yes | Street address, excluding province/city codes | 网商路699号 |
| `addressCode` | java.lang.String | yes | Address region code | 330108 |
| `addressCodeText` | java.lang.String | yes | The text corresponding to the address region code (including country, province, city) | 浙江省 杭州市 滨江区 |
| `addressId` | java.lang.Long | yes | addressId | 322683081 |
| `bizType` | java.lang.String | yes | Business type recorded for the shipping address | 无须关注 |
| `isDefault` | boolean | yes | Whether it is the default | false |
| `fullName` | java.lang.String | yes | Recipient name | 张三 |
| `latest` | boolean | yes | Whether it is the last-selected shipping address | false |
| `mobile` | java.lang.String | yes | Mobile phone number | 18012345678 |
| `phone` | java.lang.String | yes | Phone number | 0517-8888888 |
| `postCode` | java.lang.String | yes | Postal code | 310051 |

## Samples

**Request parameter example**

```
{"addressInfo":"浙江省 杭州市 滨江区网商路699号"}
```

**Example of return parameters**

```
{"result": {"address": "网商路699号","addressCode": "330108","isDefault": false,"latest": false,"postCode": "310051"}}
```
