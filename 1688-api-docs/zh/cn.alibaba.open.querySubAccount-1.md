# 查询子账号信息

API: `cn.alibaba.open:querySubAccount:1` · Category: 会员  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=cn.alibaba.open:querySubAccount-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/cn.alibaba.open/querySubAccount/{appKey}`  
无需授权 · 需要签名

查询子账号信息

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 否 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `loginId` | String | 是 | loginid | bonlientest:zhagnshan |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:ResultModel](#m-resultmodel) | 是 | 返回模型 | {"isSuccess":true,"data":{"userId":"3453","name":"张三"}} |

<a id="m-resultmodel"></a>
#### ResultModel

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `errorMsg` | java.lang.String | 否 | 错误信息 | null |
| `resultCode` | java.lang.String | 否 | 结果码 | null |
| `isSuccess` | boolean | 是 | 成功标识 | true |
| `data` | [message:CoopSubAccountModel](#m-coopsubaccountmodel) | 是 | 数据 | {"departmentId":"352","userId":"3454234","name":"zhangsan"} |

<a id="m-coopsubaccountmodel"></a>
#### CoopSubAccountModel

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `userId` | java.lang.Long | 是 | userId | 345345 |
| `loginId` | java.lang.String | 是 | 登录ID | bonlinetest:zhangsan |
| `employeeId` | java.lang.String | 否 | employeeid | 23453465 |
| `name` | java.lang.String | 否 | 姓名 | 张三 |
| `sex` | java.lang.String | 否 | 性别 | 男 |
| `personalPhone` | java.lang.String | 否 | 电话 | 13234567854 |
| `mail` | java.lang.String | 否 | 邮件 | 234324343@qq.com |
| `departmentId` | java.lang.Long | 否 | 部门ID | 36456 |
| `departmentName` | java.lang.String | 否 | 不部门名称 | 设计部 |

## 示例

****

```
入参
loginID：bonlinetest:zhangsan

返回
{"isSuccess":true,"data":{"userId":"3453","loginId":"bonlinetest:zhangsan","employeeId":"3453456","name":"张三"}}
```
