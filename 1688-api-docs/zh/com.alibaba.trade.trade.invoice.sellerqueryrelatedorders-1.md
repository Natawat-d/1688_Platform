# 基于交易单或发票申请单查询合单关联关系

API: `com.alibaba.trade:trade.invoice.sellerqueryrelatedorders:1` · Category: 发票  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:trade.invoice.sellerqueryrelatedorders-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/trade.invoice.sellerqueryrelatedorders/{appKey}`  
需要授权 (access_token) · 需要签名

卖家基于交易单 orderId或发票申请单outbizid查询合单关联关系；
如果同时传入，以outbizid为准；
如果查询出不是合并开票，那么mergeInvoice会返回false

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `reqDTO` | [message:alibaba.china.app.invoice.common.model.req.QueryMergeInvoiceRelatedOrdersOpenReqDTO](#m-alibaba-china-app-invoice-common-model-req-querymergeinvoicerelatedordersopenreqdto) | 是 | 入参 | 入参对象 |

<a id="m-alibaba-china-app-invoice-common-model-req-querymergeinvoicerelatedordersopenreqdto"></a>
#### alibaba.china.app.invoice.common.model.req.QueryMergeInvoiceRelatedOrdersOpenReqDTO

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `orderId` | java.lang.String | 是 | 交易单 | xxx |
| `outBizId` | java.lang.String | 是 | 发票申请单唯一键outbizid | xxx |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:alibaba.china.app.invoice.common.model.SingleResultDTO](#m-alibaba-china-app-invoice-common-model-singleresultdto) | 是 | 出参 | 出参对象 |

<a id="m-alibaba-china-app-invoice-common-model-singleresultdto"></a>
#### alibaba.china.app.invoice.common.model.SingleResultDTO

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `errorCode` | java.lang.String | 是 | 错误码 | 大写字母下划线分割 |
| `errorDesc` | java.lang.String | 是 | 错误原因 | 中文描述 |
| `isRetry` | boolean | 是 | 是否支持重试 | false/true |
| `result` | [message:alibaba.china.app.invoice.common.model.res.QueryMergeInvoiceRelatedOrdersOpenResDTO](#m-alibaba-china-app-invoice-common-model-res-querymergeinvoicerelatedordersopenresdto) | 是 | 出参 | 出参 |
| `success` | boolean | 是 | 服务是否成功 | false/true |

<a id="m-alibaba-china-app-invoice-common-model-res-querymergeinvoicerelatedordersopenresdto"></a>
#### alibaba.china.app.invoice.common.model.res.QueryMergeInvoiceRelatedOrdersOpenResDTO

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `mergeInvoice` | java.lang.Boolean | 是 | 是否是发生了合并开票 | false/true |
| `orderIds` | java.lang.String[] | 是 | 合并开票关联的所有交易单列表 | ["xxx","xxx"] |
| `outBizId` | java.lang.String | 是 | 发票申请单唯一键outbizid | xxx |

## 示例

**入参示例**

```
{
  "reqDTO": {
    "orderId": "xxx",
    "outBizId": "xxx"
  }
}
```

**出参示例**

```
{
  "result": {
    "errorCode": "大写字母下划线分割",
    "errorDesc": "中文描述",
    "isRetry": true,
    "result": {
      "mergeInvoice": true,
      "orderIds": [
        "[\"xxx\",\"xxx\"]"
      ],
      "outBizId": "xxx"
    },
    "success": true
  }
}
```
