# 跨境采购助手一键铺货消息

Topic: `CROSSBOARD_LP_DISTRIBUTION` · Group: CROSSBOARD (跨境消息)  
Doc page: https://open.1688.com/doc/topicDetail.htm?id=CROSSBOARD_LP_DISTRIBUTION

跨境采购助手页面将1688的商品一键铺货到接入的ERP系统中

## 消息字段

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `customerId` | String | 是 | 接入选品页客户id | testIsv1 |
| `userKey` | String | 是 | isv系统内用户的加密id，不变 | qwerty |
| `appkey` | String | 是 | 开放平台appkey | 123 |
| `userToken` | String | 是 | isv系统内用户token信息（isv能够解析出自己的用户登陆态信息） | asdfggjkl |
| `offerList` | Object[] | 是 | 铺货商品列表 | "offerList":[{  "offerId":"1244567" ，   "pic":"商品图片"}] |
| &nbsp;&nbsp;↳ `offerId` | String | 否 | 商品id | 1244567 |
| &nbsp;&nbsp;↳ `pic` | String | 否 | 商品主图url | https://cbu01.alicdn.com/xxx.jpg |
| `sign` | String | 是 | 签名信息 | ABCDEFG |
| `timestamp` | String | 是 | 当前铺货操作的时间戳（毫秒） | 123454654645 |

## 消息示例

```json
{
  "customerId": "testIsv1",
  "userKey": "qwerty",
  "appkey": "123",
  "userToken": "asdfggjkl",
  "offerList": [
    {
      "offerId": "1244567",
      "pic": "https://cbu01.alicdn.com/xxx.jpg"
    }
  ],
  "sign": "ABCDEFG",
  "timestamp": "123454654645"
}
```
