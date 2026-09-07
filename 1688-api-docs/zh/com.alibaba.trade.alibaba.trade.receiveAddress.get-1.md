# 买家获取保存的收货地址信息列表

API: `com.alibaba.trade:alibaba.trade.receiveAddress.get:1` · Category: 物流  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.trade.receiveAddress.get-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.trade.receiveAddress.get/{appKey}`  
需要授权 (access_token) · 需要签名

买家获取保存的收货地址信息列表

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:alibaba.trade.ReceiveAddressResult](#m-alibaba-trade-receiveaddressresult) | 是 | 返回结果 | {} |
| `success` | Boolean | 是 | 是否成功 | true |
| `code` | String | 是 | 错误码 | 400 |
| `message` | String | 是 | 错误信息 | api need authorized |

<a id="m-alibaba-trade-receiveaddressresult"></a>
#### alibaba.trade.ReceiveAddressResult

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `receiveAddressItems` | [message:alibaba.trade.ReceiveAddressItem[]](#m-alibaba-trade-receiveaddressitem[]) | 是 | 收货地址列表 | [] |

<a id="m-alibaba-trade-receiveaddressitem[]"></a>
#### alibaba.trade.ReceiveAddressItem[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `id` | Long | 是 | addressId | 322683081 |
| `fullName` | String | 是 | 收货人姓名 | 张三 |
| `address` | String | 是 | 街道地址，不包括省市编码 | 网商路699 |
| `post` | String | 是 | 邮编 | 340000 |
| `phone` | String | 是 | 电话 | 0517-8888888 |
| `mobilePhone` | String | 是 | 手机号 | 18012345678 |
| `addressCode` | String | 是 | 地址区域编码 | 330108 |
| `addressCodeText` | String | 是 | 地址区域编码对应的文本（包括国家，省，城市） | 浙江省 杭州市 滨江区 |
| `isDefault` | Boolean | 是 | 是否为默认 | false |
| `townCode` | String | 是 | 镇编码 | 123 |
| `townName` | String | 是 | 镇地址 | 长河镇 |

## 示例

**返回结果JSON示例**

```
{"result":{"receiveAddressItems":[{"id":560954849,"fullName":"b2b-测试账号006","address":"AAAAb2b-测试账号00613071801119","post":"810600","mobilePhone":"13071801119","addressCode":"632122","addressCodeText":"青海省 海东市 民和回族土族自治县","townCode":"630222100","townName":"川口镇"},{"id":547009002,"fullName":"2323","address":"所得税负担","post":"000000","mobilePhone":"15678906543","addressCode":"420203","addressCodeText":"湖北省 黄石市 西塞山区","townCode":"420203402","townName":"西塞山工业园区管委会"},{"id":586953001,"fullName":"放大法","address":"发的发","post":"210046","mobilePhone":"18058423787","addressCode":"320102","addressCodeText":"江苏省 南京市 玄武区","townCode":"320102003","townName":"新街口街道"},{"id":561960001,"fullName":"eeee ","address":"ccccc ","post":"145896","phone":"","mobilePhone":"13514782569","addressCode":"110102","addressCodeText":"北京 北京市 西城区","townCode":"110102009","townName":"展览路街道"},{"id":548576123,"fullName":"sdfs","address":"sdfsdf","post":"122311","mobilePhone":"15890986543","addressCode":"540121","addressCodeText":"西藏自治区 拉萨市 林周县","townCode":"540121200","townName":"春堆乡"},{"id":584414269,"fullName":"张三","address":"办事处","post":"000000","mobilePhone":"15251556677","addressCode":"441900","addressCodeText":"广东省 东莞市","townCode":"441901004","townName":"南城街道"}]},"success":true}
```
