# 保存账号所属业务线

API: `com.alibaba.fenxiao.crossborder:account.business.save:1` · Category: 回传数据  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:account.business.save-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/account.business.save/{appKey}`  
需要授权 (access_token) · 需要签名

保存账号所属业务线

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `accountPerformance` | [message:account.business.save.AccountPerformance](#m-account-business-save-accountperformance) | 是 | 入参 | 如下 |

<a id="m-account-business-save-accountperformance"></a>
#### account.business.save.AccountPerformance

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `account` | java.lang.String | 是 | 账号名称 | b测试账号009 或 b测试账号009:lctest |
| `business` | java.lang.String | 是 | 所属业务线 | 东南亚sea，南亚sa，日韩jk，港澳台hmt，中东me，北美na，拉美la，西欧we，泛俄ru，非洲af |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:account.business.save.AccountPerformanceResult](#m-account-business-save-accountperformanceresult) | 是 | 结果 | 如下 |

<a id="m-account-business-save-accountperformanceresult"></a>
#### account.business.save.AccountPerformanceResult

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `saveResult` | Boolean | 是 | 保存结果，true成功，false失败 | true |
