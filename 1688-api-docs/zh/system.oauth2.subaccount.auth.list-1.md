# 批量查询子账号授权

API: `system.oauth2:subaccount.auth.list:1` · Category: 会员  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=system.oauth2:subaccount.auth.list-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/system.oauth2/subaccount.auth.list/{appKey}`  
需要授权 (access_token) · 需要签名

批量查询主账号下面子账号的授权情况

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `subUserIdentityList` | java.lang.String[] | 是 | 子账号id列表 | "a:b","a:c" |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:alibaba.ocean.auth.dto.AuthResultModelRelation](#m-alibaba-ocean-auth-dto-authresultmodelrelation) | 是 | 查询结果 | xx |

<a id="m-alibaba-ocean-auth-dto-authresultmodelrelation"></a>
#### alibaba.ocean.auth.dto.AuthResultModelRelation

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `errorCode` | java.lang.String | 是 | 错误码 | param_error |
| `errorMessage` | java.lang.String | 是 | 错误描述 | lack of param |
| `returnValue` | [message:alibaba.ocean.auth.dto.AuthRelationDTO[]](#m-alibaba-ocean-auth-dto-authrelationdto[]) | 是 | 返回结果 | 结构体 |
| `success` | boolean | 是 | 是否成功 | true |

<a id="m-alibaba-ocean-auth-dto-authrelationdto[]"></a>
#### alibaba.ocean.auth.dto.AuthRelationDTO[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `accessToken` | java.lang.String | 是 | 授权凭证 | xx |
| `adminOwnerId` | java.lang.String | 是 | 主账号loginId | xx |
| `adminUserId` | java.lang.Long | 是 | 主账号userId | xx |
| `clientId` | java.lang.String | 是 | appKey | xx |
| `clientName` | java.lang.String | 是 | appName | xx |
| `gmtExpired` | java.util.Date | 是 | 授权过期时间 | xx |
| `memberId` | java.lang.String | 是 | 授权用户memberId | xx |
| `ownerId` | java.lang.String | 是 | 授权用户loginId | xx |
| `resourceScopes` | java.lang.String | 是 | 资源域 | xx |
| `site` | java.lang.String | 是 | 授权站点 | xx |
| `status` | java.lang.String | 是 | 授权状态 | xx |
| `subAuth` | java.lang.Boolean | 是 | 是否子账号授权 | xx |
| `subOwnerId` | java.lang.String | 是 | 子账号loginId | xx |
| `subUserId` | java.lang.Long | 是 | 子账号userId | xx |
| `userId` | java.lang.Long | 是 | 授权用户userId | xx |

## 示例

**请求示例**

```
https://gw.open.1688.com/openapi/param2/1/system.oauth2/subaccount.auth.list/YOUR_APPKEY?access_token=xx&subUserIdentityList=a,b
```
