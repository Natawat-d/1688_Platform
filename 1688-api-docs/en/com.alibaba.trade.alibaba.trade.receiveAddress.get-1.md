# Buyer gets saved shipping addresses

Original name: 买家获取保存的收货地址信息列表  
API: `com.alibaba.trade:alibaba.trade.receiveAddress.get:1` · Category: Logistics  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.trade.receiveAddress.get-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.trade.receiveAddress.get/{appKey}`  
Requires user authorization (access_token) · Requires signature

Get the buyer's list of saved shipping addresses.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.trade.ReceiveAddressResult](#m-alibaba-trade-receiveaddressresult) | yes | Return result | {} |
| `success` | Boolean | yes | Whether successful | true |
| `code` | String | yes | Error code | 400 |
| `message` | String | yes | Error message | api need authorized |

<a id="m-alibaba-trade-receiveaddressresult"></a>
#### alibaba.trade.ReceiveAddressResult

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `receiveAddressItems` | [message:alibaba.trade.ReceiveAddressItem[]](#m-alibaba-trade-receiveaddressitem[]) | yes | List of shipping addresses | [] |

<a id="m-alibaba-trade-receiveaddressitem[]"></a>
#### alibaba.trade.ReceiveAddressItem[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `id` | Long | yes | addressId | 322683081 |
| `fullName` | String | yes | Recipient name | 张三 |
| `address` | String | yes | Street address, excluding province/city codes | 网商路699 |
| `post` | String | yes | Postal code | 340000 |
| `phone` | String | yes | Phone number | 0517-8888888 |
| `mobilePhone` | String | yes | Mobile phone number | 18012345678 |
| `addressCode` | String | yes | Address region code | 330108 |
| `addressCodeText` | String | yes | The text corresponding to the address region code (including country, province, city) | 浙江省 杭州市 滨江区 |
| `isDefault` | Boolean | yes | Whether it is the default | false |
| `townCode` | String | yes | Town code | 123 |
| `townName` | String | yes | Town address | 长河镇 |

## Samples

**Example JSON of the return result**

```
{"result":{"receiveAddressItems":[{"id":560954849,"fullName":"b2b-测试账号006","address":"AAAAb2b-测试账号00613071801119","post":"810600","mobilePhone":"13071801119","addressCode":"632122","addressCodeText":"青海省 海东市 民和回族土族自治县","townCode":"630222100","townName":"川口镇"},{"id":547009002,"fullName":"2323","address":"所得税负担","post":"000000","mobilePhone":"15678906543","addressCode":"420203","addressCodeText":"湖北省 黄石市 西塞山区","townCode":"420203402","townName":"西塞山工业园区管委会"},{"id":586953001,"fullName":"放大法","address":"发的发","post":"210046","mobilePhone":"18058423787","addressCode":"320102","addressCodeText":"江苏省 南京市 玄武区","townCode":"320102003","townName":"新街口街道"},{"id":561960001,"fullName":"eeee ","address":"ccccc ","post":"145896","phone":"","mobilePhone":"13514782569","addressCode":"110102","addressCodeText":"北京 北京市 西城区","townCode":"110102009","townName":"展览路街道"},{"id":548576123,"fullName":"sdfs","address":"sdfsdf","post":"122311","mobilePhone":"15890986543","addressCode":"540121","addressCodeText":"西藏自治区 拉萨市 林周县","townCode":"540121200","townName":"春堆乡"},{"id":584414269,"fullName":"张三","address":"办事处","post":"000000","mobilePhone":"15251556677","addressCode":"441900","addressCodeText":"广东省 东莞市","townCode":"441901004","townName":"南城街道"}]},"success":true}
```
