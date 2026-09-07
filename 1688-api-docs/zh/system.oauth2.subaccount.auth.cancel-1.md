# 批量取消子账号授权

API: `system.oauth2:subaccount.auth.cancel:1` · Category: 会员  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=system.oauth2:subaccount.auth.cancel-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/system.oauth2/subaccount.auth.cancel/{appKey}`  
无需授权 · 需要签名

对子账号的授权批量取消

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 否 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `subUserIdentityList` | java.lang.String[] | 是 | 子账号id列表 | a:b,a:c |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:alibaba.ocean.auth.dto.AuthResultModelBoolean](#m-alibaba-ocean-auth-dto-authresultmodelboolean) | 是 | 返回结果 | xx |

<a id="m-alibaba-ocean-auth-dto-authresultmodelboolean"></a>
#### alibaba.ocean.auth.dto.AuthResultModelBoolean

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `errorCode` | java.lang.String | 是 | 错误码 | param_error |
| `errorMessage` | java.lang.String | 是 | 错误描述 | invalid param |
| `returnValue` | java.lang.Boolean | 是 | 返回结果 | true |
| `success` | boolean | 是 | 是否成功 | true |

## 示例

**请求示例**

```
https://gw.open.1688.com/openapi/param2/1/system.oauth2/subaccount.auth.cancel/YOUR_APPKEY?access_token=xx&subUserIdentityList=a,b
```
