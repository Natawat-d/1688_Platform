# 物流公司列表-所有的物流公司

API: `com.alibaba.logistics:alibaba.logistics.OpQueryLogisticCompanyList:1` · Category: 物流  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.logistics:alibaba.logistics.OpQueryLogisticCompanyList-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.logistics/alibaba.logistics.OpQueryLogisticCompanyList/{appKey}`  
需要授权 (access_token) · 需要签名

获取所有的物流公司名称

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
| `result` | [message:alibaba.logistics.OpLogisticsCompanyModel[]](#m-alibaba-logistics-oplogisticscompanymodel[]) | 是 | 物流公司列表 | [] |
| `success` | Boolean | 是 | 是否成功 | true |
| `errorCode` | String | 是 | 错误码 | 500 |
| `errorMessage` | String | 是 | 错误码描述 | 错误码描述 |
| `extErrorMessage` | String | 是 | 扩展错误码描述 |   |

<a id="m-alibaba-logistics-oplogisticscompanymodel[]"></a>
#### alibaba.logistics.OpLogisticsCompanyModel[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `id` | Long | 是 |  |   |
| `companyName` | String | 是 | 物流公司名称 | xxx |
| `companyNo` | String | 是 | 物流公司编号 | xxx |
| `companyPhone` | String | 是 | 物流公司服务电话 | xxx |
| `supportPrint` | Boolean | 是 | 是否支持打印 | xxx |
| `spelling` | String | 是 | 全拼 | xxx |

## 示例

**出参示例**
