# 异步执行AI找商任务

API: `com.alibaba.fenxiao.crossborder:account.search.getInfo.async:1` · Category: 商家  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:account.search.getInfo.async-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/account.search.getInfo.async/{appKey}`  
需要授权 (access_token) · 需要签名

异步执行AI找商任务

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `request` | [message:com.alibaba.global1688.common.request.SpApiExecuteRequest](#m-com-alibaba-global1688-common-request-spapiexecuterequest) | 是 | 请求参数 | {} |

<a id="m-com-alibaba-global1688-common-request-spapiexecuterequest"></a>
#### com.alibaba.global1688.common.request.SpApiExecuteRequest

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `query` | String | 是 | 查询语句 | 毛巾 |
| `spFilterCondition` | [message:com.alibaba.global1688.common.request.SpFilterCondition](#m-com-alibaba-global1688-common-request-spfiltercondition) | 否 | 过滤条件 | 过滤条件 |

<a id="m-com-alibaba-global1688-common-request-spfiltercondition"></a>
#### com.alibaba.global1688.common.request.SpFilterCondition

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `businessModel` | String | 否 | 业务模式, 仅支持传入 factory (标识源头工厂) | factory |
| `establishedYearsMin` | Integer | 否 | 最低成立年限 (筛选成立时间 ≥ 此年限的商家) | 3 |
| `repeatPurchaseRateMin` | Integer | 否 | 30 天回头率最低值, 范围 [0, 100] | 50 |
| `supportCustomization` | Boolean | 否 | 是否支持加工定制, 仅支持传入 true | true |
| `afterSalesScoreMin` | Double | 否 | 最低售后体验分, 范围 [0, 5] | 4.0 |
| `productQualityScoreMin` | Double | 否 | 最低商品体验分, 范围 [0, 5] | 4.0 |
| `consultationResponseScoreMin` | Double | 否 | 最低咨询体验分, 范围 [0, 5] | 4.0 |
| `logisticsTimelinessScoreMin` | Double | 否 | 最低物流体验分, 范围 [0, 5] | 4.0 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:com.alibaba.global1688.common.response.ApiResultExecute](#m-com-alibaba-global1688-common-response-apiresultexecute) | 是 | 返回信息 | {} |

<a id="m-com-alibaba-global1688-common-response-apiresultexecute"></a>
#### com.alibaba.global1688.common.response.ApiResultExecute

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `success` | Boolean | 是 | AI找商任务是否发起成功 | true |
| `result` | String | 是 | 任务唯一标识，查询任务状态时需要该值 | xxxxxx |
| `code` | String | 是 | 状态码 | SUCCESS |
| `message` | String | 是 | 异常信息 | error_msg_xxxxx |

## 示例

**入参示例**

```
{
    "spFilterCondition": {
        "repeatPurchaseRateMin": 70,
        "consultationResponseScoreMin": 3.5,
        "establishedYearsMin": 6,
        "productQualityScoreMin": 3.5,
        "businessModel": "factory",
        "afterSalesScoreMin": 3.5,
        "logisticsTimelinessScoreMin": 3.5
    },
    "query": "连衣裙"
}
```

**出参示例**

```
{
  "result": "taskIdxxxxxxxxxx",
  "code": "SUCCESS",
  "success": true
}
```
