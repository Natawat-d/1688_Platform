# 1688会员注册

API: `com.alibaba.fenxiao.crossborder:account.user.register:1` · Category: 会员  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:account.user.register-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/account.user.register/{appKey}`  
需要授权 (access_token) · 需要签名

注册1688会员

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `countryAccount` | [message:account.user.register.param.CountryAccount](#m-account-user-register-param-countryaccount) | 是 | 账号注册入参 | {     "country": "japan",     "site": "sniff",     "outLoginId": "18899993333",     "outMemberId": "c5b9e8a658554771852063f3a44d4e3d",     "mobile": "18899993333",     "mobileArea": "JP",     "ip": "11.11.11.11",     "email": "123@163.com" } |

<a id="m-account-user-register-param-countryaccount"></a>
#### account.user.register.param.CountryAccount

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `country` | java.lang.String | 是 | 国家，支持的参数见country枚举。 | japan |
| `site` | java.lang.String | 是 | 站点，请按照公司名或电商产品名来传入。要求无空格并小写。不超过16字符。 | sniff |
| `outLoginId` | java.lang.String | 是 | 用户在机构的登陆名，不超过60位字符。 | 1234test |
| `outMemberId` | java.lang.String | 是 | 用户在机构的唯一标识，不超过60位。 | 1232fdsf |
| `email` | java.lang.String | 是 | 邮箱，会校验格式，不超过60位。 | 123@163.com |
| `mobile` | java.lang.String | 是 | 手机号，会校验格式，不超过30位。 | 1234567890 |
| `mobileArea` | java.lang.String | 是 | 手机号所属地区，支持的参数见mobileArea枚举。 | JP |
| `ip` | java.lang.String | 是 | IP地址，会校验格式。 | 11.11.11.11 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:account.user.register.result.ResultModel](#m-account-user-register-result-resultmodel) | 是 |  |  |

<a id="m-account-user-register-result-resultmodel"></a>
#### account.user.register.result.ResultModel

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `success` | boolean | 是 | 是否调用成功，true成功，false失败。 | true |
| `code` | java.lang.String | 是 | 错误码，如S0000代表成功。 | S0000 |
| `message` | java.lang.String | 是 | 错误信息，如成功。 | 成功 |
| `result` | java.lang.Boolean | 是 | 返回结果。 | true |

## 示例

**入参示例**

```
{
    "country": "japan",
    "outLoginId": "1234test",
    "outUserId": "3333333333",
    "site": "test",
    "mobileArea": "JP",
    "ip": "11.11.11.11",
    "email": "122323@234.com",
    "mobile": "1234"
}
```

**出参示例**

```
{
    "success": true,
    "code": "S0000",
    "message": "success",
    "result": true
}
```
