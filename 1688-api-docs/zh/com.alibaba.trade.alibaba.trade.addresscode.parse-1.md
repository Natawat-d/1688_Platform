# 根据地址解析地区码

API: `com.alibaba.trade:alibaba.trade.addresscode.parse:1` · Category: 物流  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.trade.addresscode.parse-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.trade.addresscode.parse/{appKey}`  
需要授权 (access_token) · 需要签名

根据地址信息，解析地区码

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `addressInfo` | String | 是 | 地址信息 | 浙江省 杭州市 滨江区网商路699号 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:alibaba.trade.ReceiveAddress](#m-alibaba-trade-receiveaddress) | 是 | 解析后的收获地址 | {} |
| `errorCode` | String | 是 | 错误码 |   |
| `errorMessage` | String | 是 | 错误信息 |   |

<a id="m-alibaba-trade-receiveaddress"></a>
#### alibaba.trade.ReceiveAddress

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `address` | java.lang.String | 是 | 街道地址，不包括省市编码 | 网商路699号 |
| `addressCode` | java.lang.String | 是 | 地址区域编码 | 330108 |
| `addressCodeText` | java.lang.String | 是 | 地址区域编码对应的文本（包括国家，省，城市） | 浙江省 杭州市 滨江区 |
| `addressId` | java.lang.Long | 是 | addressId | 322683081 |
| `bizType` | java.lang.String | 是 | 记录收货地址的业务类型 | 无须关注 |
| `isDefault` | boolean | 是 | 是否为默认 | false |
| `fullName` | java.lang.String | 是 | 收货人姓名 | 张三 |
| `latest` | boolean | 是 | 是否是最后选择的收货地址 | false |
| `mobile` | java.lang.String | 是 | 手机号 | 18012345678 |
| `phone` | java.lang.String | 是 | 电话 | 0517-8888888 |
| `postCode` | java.lang.String | 是 | 邮编 | 310051 |

## 示例

**请求参数示例**

```
{"addressInfo":"浙江省 杭州市 滨江区网商路699号"}
```

**返回参数示例**

```
{"result": {"address": "网商路699号","addressCode": "330108","isDefault": false,"latest": false,"postCode": "310051"}}
```
