# 获取子账号列表

API: `com.alibaba.account:alibaba.subAccount.list:1` · Category: 订单  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.account:alibaba.subAccount.list-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.account/alibaba.subAccount.list/{appKey}`  
需要授权 (access_token) · 需要签名

获取用户的主账号及子账号信息。
如果调用API时的授权为子账号，则只返回支持子账号对应的主账号；
如果调用API时的授权为主账号，则只返回所有子账号列表；

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
| `mainUserId` | String | 是 | 主账号Userid | 12346 |
| `mainLoginId` | String | 是 | 主账号loginId | alitestforisv01 |
| `mainMemberId` | String | 是 | 主账号MemberId | b2b-123565 |
| `subAccountList` | [message:alibaba.account.simpleAccountInfo[]](#m-alibaba-account-simpleaccountinfo[]) | 是 | 子账号列表 | [] |
| `errorMsg` | String | 是 | 错误描述 | 获取帐号信息失败 |
| `errorCode` | String | 是 | 错误码 | 500_001 |

<a id="m-alibaba-account-simpleaccountinfo[]"></a>
#### alibaba.account.simpleAccountInfo[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `loginId` | java.lang.String | 是 | 子账号登录名 | alitestforisv01:01 |
| `userId` | java.lang.Long | 是 | 子账号用户ID | 12124134 |
| `memberId` | java.lang.String | 是 | 子账号memberId | b2b-12315435 |
| `status` | String | 是 | 子账号状态 | enabled |

## 示例

**出参示例**

```
{
  "mainLoginId": "b测试账号003",
  "mainMemberId": "b2b-22**4086",
  "mainUserId": "22**86",
  "subAccountList": [
    {
      "loginId": "b测试账号003:yms",
      "userId": 400137**76,
      "memberId": "b2b-22**4086"
    },
    {
      "loginId": "b测试账号003:yansen",
      "userId": 4000**0833,
      "memberId": "b2b-22**86"
    }
  ]
}

```
