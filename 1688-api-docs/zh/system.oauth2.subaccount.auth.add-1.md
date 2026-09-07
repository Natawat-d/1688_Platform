# 批量添加子账号授权

API: `system.oauth2:subaccount.auth.add:1` · Category: 会员  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=system.oauth2:subaccount.auth.add-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/system.oauth2/subaccount.auth.add/{appKey}`  
需要授权 (access_token) · 需要签名

批量对某个主账号下面的子账号添加授权，前提是主账号已经授权

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `subUserIdentityList` | java.lang.String[] | 是 | 子账号id列表 | abb:test1,abb:test2 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:alibaba.ocean.auth.dto.AuthResultModelMap](#m-alibaba-ocean-auth-dto-authresultmodelmap) | 是 | 返回结果对象 | {} |

<a id="m-alibaba-ocean-auth-dto-authresultmodelmap"></a>
#### alibaba.ocean.auth.dto.AuthResultModelMap

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `errorCode` | java.lang.String | 是 | 错误码 | param_error |
| `errorMessage` | java.lang.String | 是 | 错误信息 | invalid param |
| `returnValue` | java.util.Map | 是 | 返回结果 | map |
| `success` | boolean | 是 | 是否成功 | true |

## 示例

**请求示例**

```
https://gw.open.1688.com/openapi/param2/1/system.oauth2/subaccount.auth.add/YOUR_APPKEY?access_token=xx&subUserIdentityList=a,b
```
