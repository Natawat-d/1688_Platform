# 跨境场景获取非授权用户的基本信息

API: `com.alibaba.account:alibaba.account.agent.crossBasic:1` · Category: 会员  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.account:alibaba.account.agent.crossBasic-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.account/alibaba.account.agent.crossBasic/{appKey}`  
需要授权 (access_token) · 需要签名

可以查看他人的用户信息，使用在跨境场景

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `loginId` | String | 否 | 用户的loginId，入参不可同时为空 | alitestforisv01 |
| `domain` | String | 否 | 旺铺域名，入参不可同时为空 | trgm66666.1688.com |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:alibaba.account.simpleAccountInfo](#m-alibaba-account-simpleaccountinfo) | 是 | 会员信息 | {} |
| `errorCode` | String | 是 | 错误编码 | 001 |
| `errorMessage` | String | 是 | 错误信息 |   |
| `success` | Boolean | 是 | 是否成功 | true |

<a id="m-alibaba-account-simpleaccountinfo"></a>
#### alibaba.account.simpleAccountInfo

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `loginId` | java.lang.String | 是 | 登录名 | alitestforis01 |
| `categoryName` | java.lang.String | 是 | 主营行业的类目名称 | 连衣裙 |
| `companyName` | java.lang.String | 是 | 公司名 | 测试公司1 |
| `shopUrl` | java.lang.String | 是 | 旺铺首页地址 | https://alitest.1688.com	 |
| `supplierName` | java.lang.String | 是 | 供应商名称 | alitestforis01 |
| `kuaJingBao` | Boolean | 是 | 是否跨境宝 | true |

## 示例

**输出结果示例**

```
{"result":{"loginId":"alitestforisv01","categoryName":"办公、文教","companyName":"AOP对外测试账号01","shopUrl":"https://trgm66666.1688.com"}}
```
