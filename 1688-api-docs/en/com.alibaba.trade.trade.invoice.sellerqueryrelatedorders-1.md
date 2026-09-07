# Query merged-invoice relationships by order or invoice application

Original name: 基于交易单或发票申请单查询合单关联关系  
API: `com.alibaba.trade:trade.invoice.sellerqueryrelatedorders:1` · Category: Invoicing  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:trade.invoice.sellerqueryrelatedorders-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/trade.invoice.sellerqueryrelatedorders/{appKey}`  
Requires user authorization (access_token) · Requires signature

The seller queries merged-invoice relationships by order ID (orderId) or invoice-application ID (outbizid). If both are passed, outbizid takes precedence. If the result is not a merged invoice, mergeInvoice returns false.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `reqDTO` | [message:alibaba.china.app.invoice.common.model.req.QueryMergeInvoiceRelatedOrdersOpenReqDTO](#m-alibaba-china-app-invoice-common-model-req-querymergeinvoicerelatedordersopenreqdto) | yes | Input parameter | 入参对象 |

<a id="m-alibaba-china-app-invoice-common-model-req-querymergeinvoicerelatedordersopenreqdto"></a>
#### alibaba.china.app.invoice.common.model.req.QueryMergeInvoiceRelatedOrdersOpenReqDTO

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `orderId` | java.lang.String | yes | Transaction order | xxx |
| `outBizId` | java.lang.String | yes | Unique key outbizid of the invoice application order | xxx |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.china.app.invoice.common.model.SingleResultDTO](#m-alibaba-china-app-invoice-common-model-singleresultdto) | yes | Output parameters | 出参对象 |

<a id="m-alibaba-china-app-invoice-common-model-singleresultdto"></a>
#### alibaba.china.app.invoice.common.model.SingleResultDTO

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `errorCode` | java.lang.String | yes | Error code | 大写字母下划线分割 |
| `errorDesc` | java.lang.String | yes | Error reason | 中文描述 |
| `isRetry` | boolean | yes | Whether retry is supported | false/true |
| `result` | [message:alibaba.china.app.invoice.common.model.res.QueryMergeInvoiceRelatedOrdersOpenResDTO](#m-alibaba-china-app-invoice-common-model-res-querymergeinvoicerelatedordersopenresdto) | yes | Output parameters | 出参 |
| `success` | boolean | yes | Whether the service was successful | false/true |

<a id="m-alibaba-china-app-invoice-common-model-res-querymergeinvoicerelatedordersopenresdto"></a>
#### alibaba.china.app.invoice.common.model.res.QueryMergeInvoiceRelatedOrdersOpenResDTO

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `mergeInvoice` | java.lang.Boolean | yes | Whether combined invoicing occurred | false/true |
| `orderIds` | java.lang.String[] | yes | List of all transaction orders associated with the merged invoicing | ["xxx","xxx"] |
| `outBizId` | java.lang.String | yes | Unique key outbizid of the invoice application order | xxx |

## Samples

**Input parameter example**

```
{
  "reqDTO": {
    "orderId": "xxx",
    "outBizId": "xxx"
  }
}
```

**Output parameter example**

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
