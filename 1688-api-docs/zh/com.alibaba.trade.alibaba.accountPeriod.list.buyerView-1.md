# 买家查看获得的所有账期授信

API: `com.alibaba.trade:alibaba.accountPeriod.list.buyerView:1` · Category: 支付  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.accountPeriod.list.buyerView-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.accountPeriod.list.buyerView/{appKey}`  
需要授权 (access_token) · 需要签名

买家维度查看所有获得的账期授信。可翻页查询，每次返回不超过10条

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `pageIndex` | Long | 否 | 页码 | 1 |
| `sellerLoginId` | String | 否 | 卖家ID，不填则查询全部 | alitestforisv01 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `success` | Boolean | 是 | 是否成功 | true |
| `errorCode` | String | 是 | 错误码 | 500_1 |
| `errorMsg` | String | 是 | 错误信息 |   |
| `resultList` | [message:accountPeriod.list.buyerView.result](#m-accountperiod-list-buyerview-result) | 是 | 返回数据结果 | {} |

<a id="m-accountperiod-list-buyerview-result"></a>
#### accountPeriod.list.buyerView.result

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `totalCount` | String | 是 | 总数据条数 | 100 |
| `accountPeriodList` | [message:alibaba.ocean.openplatform.biz.trade.common.model.AccountPeriodInfo[]](#m-alibaba-ocean-openplatform-biz-trade-common-model-accountperiodinfo[]) | 是 | 授信列表 | [] |

<a id="m-alibaba-ocean-openplatform-biz-trade-common-model-accountperiodinfo[]"></a>
#### alibaba.ocean.openplatform.biz.trade.common.model.AccountPeriodInfo[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `sellerLoginId` | java.lang.String | 是 | 卖家loginId | alitestforisv02 |
| `sellerCompanyName` | String | 是 | 卖家公司名 | 公司名 |
| `gmtQuota` | java.util.Date | 是 | 授信日期 | 20170913231727000+0800 |
| `quota` | java.lang.Long | 是 | 授信额度值,单位为分 | 10000 |
| `surplusQuota` | java.lang.Long | 是 | 可用授信额度值,单位为分 | 10000 |
| `statusStr` | java.lang.String | 是 | 状态描述 | 有效 |
| `tapDateStr` | java.lang.String | 是 | 账期日期描述 | 两个月一结，20号 |
| `tapOverdue` | java.lang.Integer | 是 | 逾期次数 | 0 |

## 示例

**出参示例**

```
{
  "resultList": {
    "accountPeriodList": [
      {
        "gmtQuota": "20190110114140000+0800",
        "quota": 90000000,
        "sellerCompanyName": "AOP对外测试账号01",
        "sellerLoginId": "alitestforisv01",
        "statusStr": "有效",
        "surplusQuota": 90000000,
        "tapDateStr": "120天",
        "tapOverdue": 0,
      }
    ],
    "totalCount": "1"
  },
  "success": true
}
```
