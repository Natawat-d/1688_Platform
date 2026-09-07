# Merged-invoice consultation

Original name: 合并开票咨询  
API: `com.alibaba.trade:trade.invoice.consult:1` · Category: Invoicing  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:trade.invoice.consult-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/trade.invoice.consult/{appKey}`  
Requires user authorization (access_token) · Requires signature

Submit orders to check whether they can be invoiced together and how they are grouped.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `reqDTO` | [message:alibaba.china.app.invoice.common.model.req.MergeInvoiceConsultOpenReqDTO](#m-alibaba-china-app-invoice-common-model-req-mergeinvoiceconsultopenreqdto) | yes | Input parameter | 入参 |

<a id="m-alibaba-china-app-invoice-common-model-req-mergeinvoiceconsultopenreqdto"></a>
#### alibaba.china.app.invoice.common.model.req.MergeInvoiceConsultOpenReqDTO

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `orderIdsJsonList` | java.lang.String | yes | List of orders to be consulted (JSON string) | ["123","456","789"] |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.china.app.invoice.common.model.SingleResultDTO](#m-alibaba-china-app-invoice-common-model-singleresultdto) | yes | Output parameters | 出参 |

<a id="m-alibaba-china-app-invoice-common-model-singleresultdto"></a>
#### alibaba.china.app.invoice.common.model.SingleResultDTO

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `errorCode` | java.lang.String | yes | Error code | 大写字母下划线分割 |
| `errorDesc` | java.lang.String | yes | Error description | 中文 |
| `isRetry` | boolean | yes | Whether to retry | false/true |
| `result` | [message:alibaba.china.app.invoice.common.model.res.MergeInvoiceConsultOpenResDTO](#m-alibaba-china-app-invoice-common-model-res-mergeinvoiceconsultopenresdto) | yes | Result object | 对象 |
| `success` | boolean | yes | Whether the service was successful | false/true |

<a id="m-alibaba-china-app-invoice-common-model-res-mergeinvoiceconsultopenresdto"></a>
#### alibaba.china.app.invoice.common.model.res.MergeInvoiceConsultOpenResDTO

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `groups` | [message:alibaba.china.app.invoice.common.model.normal.mtop.MergeInvoiceGroupDTO[]](#m-alibaba-china-app-invoice-common-model-normal-mtop-mergeinvoicegroupdto[]) | yes | Grouping object for merged invoicing inquiries | 对象 |
| `supportMerge` | java.lang.Boolean | yes | Whether combined invoicing is supported | false/true |

<a id="m-alibaba-china-app-invoice-common-model-normal-mtop-mergeinvoicegroupdto[]"></a>
#### alibaba.china.app.invoice.common.model.normal.mtop.MergeInvoiceGroupDTO[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `orderIds` | java.lang.String[] | yes | List of order IDs under this group: one invoice per group | ["xxx","xxx"] |
| `sellerLoginId` | java.lang.String | yes | Seller name | 测试卖家001 |
| `totalAmount` | java.lang.Long | yes | Invoiceable amount for the group | 100L |

## Samples

**Input parameter example**

```
{
  "reqDTO": {
    "orderIdsJsonList": "[\"123\",\"456\",\"789\"]"
  }
}
```

**Output parameter example**

```
{
  "result": {
    "errorCode": "大写字母下划线分割",
    "errorDesc": "中文",
    "isRetry": true,
    "result": {
      "groups": [
        {
          "orderIds": [
            "[\"xxx\",\"xxx\"]"
          ],
          "sellerLoginId": "测试卖家001",
          "totalAmount": 123
        }
      ],
      "supportMerge": true
    },
    "success": true
  }
}
```
