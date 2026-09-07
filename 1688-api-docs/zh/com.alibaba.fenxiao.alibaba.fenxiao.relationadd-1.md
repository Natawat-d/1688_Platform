# 买卖家分销关系添加

API: `com.alibaba.fenxiao:alibaba.fenxiao.relationadd:1` · Category: 工具  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao:alibaba.fenxiao.relationadd-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao/alibaba.fenxiao.relationadd/{appKey}`  
需要授权 (access_token) · 需要签名

通过商品id，添加买卖家分销关系

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `offerId` | java.lang.Long | 是 | 商品id | 98129931 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:alibaba.ocean.openplatform.common.addrelation.ResultModel](#m-alibaba-ocean-openplatform-common-addrelation-resultmodel) | 是 |  |  |

<a id="m-alibaba-ocean-openplatform-common-addrelation-resultmodel"></a>
#### alibaba.ocean.openplatform.common.addrelation.ResultModel

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `success` | boolean | 是 |  |  |
| `code` | java.lang.String | 是 |  |  |
| `message` | java.lang.String | 是 |  |  |
| `result` | java.lang.Boolean | 是 |  |  |
