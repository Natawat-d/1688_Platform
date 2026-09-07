# Submit merged-invoice request

Original name: 提交合并开票  
API: `com.alibaba.trade:trade.invoice.mergeapply:1` · Category: Invoicing  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:trade.invoice.mergeapply-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/trade.invoice.mergeapply/{appKey}`  
Requires user authorization (access_token) · Requires signature

First call the merged-invoice consultation trade.invoice.consult to get the grouping result, then call merged invoicing.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `reqDTO` | [message:alibaba.china.app.invoice.common.model.req.MergeInvoiceApplyOpenReqDTO](#m-alibaba-china-app-invoice-common-model-req-mergeinvoiceapplyopenreqdto) | yes | Input parameter | 入参对象 |

<a id="m-alibaba-china-app-invoice-common-model-req-mergeinvoiceapplyopenreqdto"></a>
#### alibaba.china.app.invoice.common.model.req.MergeInvoiceApplyOpenReqDTO

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `invoiceType` | java.lang.String | yes | Invoice type | 1:普票，2：专票 |
| `orderIdGroupsJsonList` | java.lang.String | yes | Merged invoicing group list (each group corresponds to one invoice); each group is a comma-separated string | ["123,456,789","101,202"] |
| `purchaserInvoiceTitle` | [message:alibaba.china.app.invoice.common.model.invoiceTitle.InvoiceTitleModel](#m-alibaba-china-app-invoice-common-model-invoicetitle-invoicetitlemodel) | yes | Buyer invoice title object | 对象 |

<a id="m-alibaba-china-app-invoice-common-model-invoicetitle-invoicetitlemodel"></a>
#### alibaba.china.app.invoice.common.model.invoiceTitle.InvoiceTitleModel

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `bankAccountId` | java.lang.String | no | Bank account | xxx |
| `bankName` | java.lang.String | no | Name of the account-holding bank | xx银行 |
| `email` | java.lang.String | no | Email | xxx |
| `isDefault` | java.lang.String | no | Whether it is the default invoice title | Y：是/N：否 |
| `name` | java.lang.String | no | Enterprise contact person | xxx |
| `receiverPhone` | java.lang.String | no | Recipient phone number | xx |
| `registerAddress` | java.lang.String | no | Enterprise registered address | xxx |
| `registerPhone` | java.lang.String | no | Enterprise phone | xxx |
| `taxpayerIdentify` | java.lang.String | no | Tax number | 专票必填 |
| `title` | java.lang.String | yes | Title | xxx |
| `titleType` | java.lang.Integer | yes | Invoice type | 0：个人和社会组织，1：企业 |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.china.app.invoice.common.model.SingleResultDTO](#m-alibaba-china-app-invoice-common-model-singleresultdto) | yes | Output parameters | 出参对象 |

<a id="m-alibaba-china-app-invoice-common-model-singleresultdto"></a>
#### alibaba.china.app.invoice.common.model.SingleResultDTO

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `errorCode` | java.lang.String | yes | Error code | 英文字母下划线分割 |
| `errorDesc` | java.lang.String | yes | Error description | 中文描述 |
| `isIdempotent` | java.lang.Boolean | yes | Idempotent | false/true |
| `isRetry` | boolean | yes | Whether retry is supported | false/true |
| `result` | [message:alibaba.china.app.invoice.common.model.res.MergeInvoiceApplyOpenResDTO](#m-alibaba-china-app-invoice-common-model-res-mergeinvoiceapplyopenresdto) | yes | Object | 对象 |
| `success` | boolean | yes | Whether the service was successful | false/true |

<a id="m-alibaba-china-app-invoice-common-model-res-mergeinvoiceapplyopenresdto"></a>
#### alibaba.china.app.invoice.common.model.res.MergeInvoiceApplyOpenResDTO

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `failReason` | java.lang.String | yes | Business failure reason | 中文描述 |
| `idempotent` | java.lang.Boolean | yes | Whether idempotent | false/true |
| `result` | boolean | yes | Whether the business call was successful | false/true |

## Samples

**Input parameter example**

```
{
  "reqDTO": {
    "invoiceType": "1:普票，2：专票",
    "loginUserId": 123,
    "orderIdGroupsJsonList": "[\"123,456,789\",\"101,202\"]",
    "purchaserInvoiceTitle": {
      "bankAccountId": "xxx",
      "bankName": "xx银行",
      "email": "xxx",
      "isDefault": "Y：是/N：否",
      "name": "xxx",
      "receiverPhone": "xx",
      "registerAddress": "xxx",
      "registerPhone": "xxx",
      "taxpayerIdentify": "xx",
      "title": "xxx",
      "titleType": 123
    }
  }
}
```

**Output parameter example**

```
{
  "result": {
    "errorCode": "英文字母下划线分割",
    "errorDesc": "中文描述",
    "isIdempotent": true,
    "isRetry": true,
    "result": {
      "failReason": "中文描述",
      "idempotent": true,
      "result": true
    },
    "success": true
  }
}
```
