# 提交合并开票

API: `com.alibaba.trade:trade.invoice.mergeapply:1` · Category: 发票  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:trade.invoice.mergeapply-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/trade.invoice.mergeapply/{appKey}`  
需要授权 (access_token) · 需要签名

先调用合并开票咨询trade.invoice.consult，获取分组结果后，调用合并开票

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `reqDTO` | [message:alibaba.china.app.invoice.common.model.req.MergeInvoiceApplyOpenReqDTO](#m-alibaba-china-app-invoice-common-model-req-mergeinvoiceapplyopenreqdto) | 是 | 入参 | 入参对象 |

<a id="m-alibaba-china-app-invoice-common-model-req-mergeinvoiceapplyopenreqdto"></a>
#### alibaba.china.app.invoice.common.model.req.MergeInvoiceApplyOpenReqDTO

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `invoiceType` | java.lang.String | 是 | 发票类型 | 1:普票，2：专票 |
| `orderIdGroupsJsonList` | java.lang.String | 是 | 合并开票分组列表（每个分组对应一张票),每个分组逗号分隔的字符串 | ["123,456,789","101,202"] |
| `purchaserInvoiceTitle` | [message:alibaba.china.app.invoice.common.model.invoiceTitle.InvoiceTitleModel](#m-alibaba-china-app-invoice-common-model-invoicetitle-invoicetitlemodel) | 是 | 购方抬头对象 | 对象 |

<a id="m-alibaba-china-app-invoice-common-model-invoicetitle-invoicetitlemodel"></a>
#### alibaba.china.app.invoice.common.model.invoiceTitle.InvoiceTitleModel

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `bankAccountId` | java.lang.String | 否 | 银行账户 | xxx |
| `bankName` | java.lang.String | 否 | 开户银行名称 | xx银行 |
| `email` | java.lang.String | 否 | 邮箱 | xxx |
| `isDefault` | java.lang.String | 否 | 是否默认抬头 | Y：是/N：否 |
| `name` | java.lang.String | 否 | 企业联系人 | xxx |
| `receiverPhone` | java.lang.String | 否 | 收件人电话 | xx |
| `registerAddress` | java.lang.String | 否 | 企业注册地址 | xxx |
| `registerPhone` | java.lang.String | 否 | 企业电话 | xxx |
| `taxpayerIdentify` | java.lang.String | 否 | 税号 | 专票必填 |
| `title` | java.lang.String | 是 | 抬头 | xxx |
| `titleType` | java.lang.Integer | 是 | 发票类型 | 0：个人和社会组织，1：企业 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:alibaba.china.app.invoice.common.model.SingleResultDTO](#m-alibaba-china-app-invoice-common-model-singleresultdto) | 是 | 出参 | 出参对象 |

<a id="m-alibaba-china-app-invoice-common-model-singleresultdto"></a>
#### alibaba.china.app.invoice.common.model.SingleResultDTO

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `errorCode` | java.lang.String | 是 | 错误码 | 英文字母下划线分割 |
| `errorDesc` | java.lang.String | 是 | 错误描述 | 中文描述 |
| `isIdempotent` | java.lang.Boolean | 是 | 幂等 | false/true |
| `isRetry` | boolean | 是 | 是否支持重试 | false/true |
| `result` | [message:alibaba.china.app.invoice.common.model.res.MergeInvoiceApplyOpenResDTO](#m-alibaba-china-app-invoice-common-model-res-mergeinvoiceapplyopenresdto) | 是 | 对象 | 对象 |
| `success` | boolean | 是 | 服务是否成功 | false/true |

<a id="m-alibaba-china-app-invoice-common-model-res-mergeinvoiceapplyopenresdto"></a>
#### alibaba.china.app.invoice.common.model.res.MergeInvoiceApplyOpenResDTO

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `failReason` | java.lang.String | 是 | 业务失败原因 | 中文描述 |
| `idempotent` | java.lang.Boolean | 是 | 是否幂等 | false/true |
| `result` | boolean | 是 | 业务调用是否成功 | false/true |

## 示例

**入参示例**

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

**出参示例**

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
