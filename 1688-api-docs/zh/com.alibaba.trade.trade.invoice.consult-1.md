# 合并开票咨询

API: `com.alibaba.trade:trade.invoice.consult:1` · Category: 发票  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:trade.invoice.consult-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/trade.invoice.consult/{appKey}`  
需要授权 (access_token) · 需要签名

提交交易单，咨询能否合并开票，以及开票分组

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `reqDTO` | [message:alibaba.china.app.invoice.common.model.req.MergeInvoiceConsultOpenReqDTO](#m-alibaba-china-app-invoice-common-model-req-mergeinvoiceconsultopenreqdto) | 是 | 入参 | 入参 |

<a id="m-alibaba-china-app-invoice-common-model-req-mergeinvoiceconsultopenreqdto"></a>
#### alibaba.china.app.invoice.common.model.req.MergeInvoiceConsultOpenReqDTO

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `orderIdsJsonList` | java.lang.String | 是 | 需要咨询的订单列表（JSON字符串） | ["123","456","789"] |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:alibaba.china.app.invoice.common.model.SingleResultDTO](#m-alibaba-china-app-invoice-common-model-singleresultdto) | 是 | 出参 | 出参 |

<a id="m-alibaba-china-app-invoice-common-model-singleresultdto"></a>
#### alibaba.china.app.invoice.common.model.SingleResultDTO

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `errorCode` | java.lang.String | 是 | 错误码 | 大写字母下划线分割 |
| `errorDesc` | java.lang.String | 是 | 错误描述 | 中文 |
| `isRetry` | boolean | 是 | 是否重试 | false/true |
| `result` | [message:alibaba.china.app.invoice.common.model.res.MergeInvoiceConsultOpenResDTO](#m-alibaba-china-app-invoice-common-model-res-mergeinvoiceconsultopenresdto) | 是 | 结果对象 | 对象 |
| `success` | boolean | 是 | 服务是否成功 | false/true |

<a id="m-alibaba-china-app-invoice-common-model-res-mergeinvoiceconsultopenresdto"></a>
#### alibaba.china.app.invoice.common.model.res.MergeInvoiceConsultOpenResDTO

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `groups` | [message:alibaba.china.app.invoice.common.model.normal.mtop.MergeInvoiceGroupDTO[]](#m-alibaba-china-app-invoice-common-model-normal-mtop-mergeinvoicegroupdto[]) | 是 | 合并开票咨询的分组对象 | 对象 |
| `supportMerge` | java.lang.Boolean | 是 | 是否支持合并开票 | false/true |

<a id="m-alibaba-china-app-invoice-common-model-normal-mtop-mergeinvoicegroupdto[]"></a>
#### alibaba.china.app.invoice.common.model.normal.mtop.MergeInvoiceGroupDTO[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `orderIds` | java.lang.String[] | 是 | 该分组下的订单ID列表：每个分组一张发票 | ["xxx","xxx"] |
| `sellerLoginId` | java.lang.String | 是 | 卖家名称 | 测试卖家001 |
| `totalAmount` | java.lang.Long | 是 | 分组可开票金额 | 100L |

## 示例

**入参示例**

```
{
  "reqDTO": {
    "orderIdsJsonList": "[\"123\",\"456\",\"789\"]"
  }
}
```

**出参示例**

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
