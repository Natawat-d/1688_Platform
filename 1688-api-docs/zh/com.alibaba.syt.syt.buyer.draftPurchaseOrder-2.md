# 88生意通买家起草采购单

API: `com.alibaba.syt:syt.buyer.draftPurchaseOrder:2` · Category: 88生意通  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.syt:syt.buyer.draftPurchaseOrder-2  
Request URL: `https://gw.open.1688.com/openapi/param2/2/com.alibaba.syt/syt.buyer.draftPurchaseOrder/{appKey}`  
需要授权 (access_token) · 需要签名

88生意通解决方案
88生意通买家起草采购单

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `request` | [message:alibaba.china.share.creditpay.gateway.api.request.ContractDraftAndSignApplyApiRequest](#m-alibaba-china-share-creditpay-gateway-api-request-contractdraftandsignapplyapirequest) | 是 | 请求参数 | {     "origin": "ERP",     "requestNo": "REQ2026051100001",     "contractType": "PURCHASE_ORDER",     "draftNo": null,     "payTerm": {         "amount": 10000,         "payMethod": "SINGLE",         "payTimeType": "IMMEDIATELY",         "payTimeContent": null     },     "contentTerm": {         "contentType": "PURCHASE_ORDER",         "purchaseItems": [             {                 "key": "SKU001",                 "productName": "商品A",                 "productSpec": "红色/L码",                 "quantity": 10,                 "unitPrice": 1000,                 "subtotal": 10000,                 "productImage": {                     "fileName": "product_a.jpg",                     "fileUrl": "https://example.com/images/product_a.jpg"                 }             }         ],         "attachments": [          ]     },     "contractRoleInfoTerm": {         "counterpartyOrigin": "LOGIN_ID_1688_MATCH",         "counterpartyLoginId": "seller1688",         "counterpartyName": null,         "counterpartyLicenseNo": null,         "counterpartyLicenseType": null     },     "postFeeTerm": {         "postFee": 0     },     "confirmTerm": {         "contractConfirmType": "BUYER_MANUAL_CONFIRM",         "autoConfirmCondition": null     },     "drafterNeedSignConfirmAgain": false } |

<a id="m-alibaba-china-share-creditpay-gateway-api-request-contractdraftandsignapplyapirequest"></a>
#### alibaba.china.share.creditpay.gateway.api.request.ContractDraftAndSignApplyApiRequest

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `origin` | java.lang.String | 是 | 调用来源 | ERP，固定值 |
| `requestNo` | java.lang.String | 是 | 请求幂等号 | 本次创建采购单唯一请求号 |
| `draftNo` | java.lang.String | 是 | 采购单编号-创建时不填 | 空，不传值 |
| `payTerm` | [message:alibaba.china.share.creditpay.gateway.api.dto.ApiContractPayTerm](#m-alibaba-china-share-creditpay-gateway-api-dto-apicontractpayterm) | 是 | 支付条款 | {     "amount": 10000.00,     "payMethod": "SINGLE",     "payTimeType": "IMMEDIATELY",     "payTimeContent": null   } |
| `contentTerm` | [message:alibaba.china.share.creditpay.gateway.api.dto.ApiContractContentTerm](#m-alibaba-china-share-creditpay-gateway-api-dto-apicontractcontentterm) | 是 | 内容条款 | {     "contentType": "PURCHASE_ORDER",     "purchaseItems": [       {         "key": "SKU001",         "productName": "商品A",         "productSpec": "红色/L码",         "quantity": 10,         "unitPrice": 1000.00,         "subtotal": 10000.00,         "productImage": {           "fileName": "product_a.jpg",           "fileUrl": "https://example.com/images/product_a.jpg"         }       }     ],     "attachments": []   } |
| `contractRoleInfoTerm` | [message:alibaba.china.share.creditpay.gateway.api.dto.ApiContractRoleInfoTerm](#m-alibaba-china-share-creditpay-gateway-api-dto-apicontractroleinfoterm) | 是 | 对手方约定 | {     "counterpartyOrigin": "LOGIN_ID_1688_MATCH",     "counterpartyLoginId": "seller1688",     "counterpartyName": null,     "counterpartyLicenseNo": null,     "counterpartyLicenseType": null   } |
| `drafterNeedSignConfirmAgain` | Boolean | 是 | 起草方是否需要二次确认 false不需要 | false |
| `postFeeTerm` | [message:com.alibaba.china.share.creditpay.gateway.api.dto.ApiContractPostFeeTerm](#m-com-alibaba-china-share-creditpay-gateway-api-dto-apicontractpostfeeterm) | 是 | 邮费条款 | {   "postFee": 1.00 } |
| `confirmTerm` | [message:com.alibaba.china.share.creditpay.gateway.api.dto.ApiContractConfirmTerm](#m-com-alibaba-china-share-creditpay-gateway-api-dto-apicontractconfirmterm) | 是 | 确认收货条款 | {   "contractConfirmType": "AUTO_CONFIRM",   "autoConfirmCondition": {     "conditionEnum": "AFTER_DELIVERY",     "xDay": 7   } } |

<a id="m-alibaba-china-share-creditpay-gateway-api-dto-apicontractpayterm"></a>
#### alibaba.china.share.creditpay.gateway.api.dto.ApiContractPayTerm

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `amount` | BigDecimal | 是 | 支付金额，单位元，币种人民币 CNY | 1 |
| `payMethod` | java.lang.String | 是 | 支付方法-固定值 | SINGLE |
| `payTimeType` | java.lang.String | 是 | 支付时间类型-固定值 | IMMEDIATELY |
| `payTimeContent` | java.lang.String | 否 | 支付时间内容 | 不传 |

<a id="m-alibaba-china-share-creditpay-gateway-api-dto-apicontractcontentterm"></a>
#### alibaba.china.share.creditpay.gateway.api.dto.ApiContractContentTerm

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `contentType` | java.lang.String | 是 | 内容类型-固定值 | PURCHASE_ORDER |
| `purchaseItems` | [message:alibaba.china.share.creditpay.gateway.api.dto.ApiPurchaseItem[]](#m-alibaba-china-share-creditpay-gateway-api-dto-apipurchaseitem[]) | 是 | 采购单详情 | [       {         "key": "SKU001",         "productName": "商品A",         "productSpec": "红色/L码",         "quantity": 10,         "unitPrice": 1000.00,         "subtotal": 10000.00       } ] |

<a id="m-alibaba-china-share-creditpay-gateway-api-dto-apipurchaseitem[]"></a>
#### alibaba.china.share.creditpay.gateway.api.dto.ApiPurchaseItem[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `key` | java.lang.String | 是 | 采购单序号 | sku001 |
| `productName` | java.lang.String | 是 | 产品名称 | 水杯 |
| `productSpec` | java.lang.String | 是 | 产品 sku | 1L、白色、大肚款 |
| `quantity` | java.lang.Integer | 是 | 购买数量 | 1 |
| `unitPrice` | BigDecimal | 是 | 单价，单位元 | 1 |
| `subtotal` | BigDecimal | 是 | 总价，单位元 | 1 |

<a id="m-alibaba-china-share-creditpay-gateway-api-dto-apicontractroleinfoterm"></a>
#### alibaba.china.share.creditpay.gateway.api.dto.ApiContractRoleInfoTerm

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `counterpartyOrigin` | java.lang.String | 是 | LOGIN_ID_1688_MATCH（通过1688loginId指定对手方）、DRAFTER_INPUT（通过对方主体信息指定）、COUNTERPARTY_INPUT（不指定交给对方自己认领） | LOGIN_ID_1688_MATCH |
| `counterpartyLoginId` | java.lang.String | 否 | 1688loginId 当 counterpartyOrigin = LOGIN_ID_1688_MATCH 必填 | ces测试002 |
| `counterpartyName` | java.lang.String | 否 | 对手方名称 当 counterpartyOrigin = DRAFTER_INPUT 必填 | 张三 |
| `counterpartyLicenseNo` | java.lang.String | 否 | 对手方身份证号或者营业执照号 当 counterpartyOrigin = DRAFTER_INPUT 必填 | 330xx |
| `counterpartyLicenseType` | java.lang.String | 否 | 证件类型 当 counterpartyOrigin = DRAFTER_INPUT 必填 | 统一社会信用代码UNITY |

<a id="m-com-alibaba-china-share-creditpay-gateway-api-dto-apicontractpostfeeterm"></a>
#### com.alibaba.china.share.creditpay.gateway.api.dto.ApiContractPostFeeTerm

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `postFee` | BigDecimal | 是 | postFee | 0.01 |

<a id="m-com-alibaba-china-share-creditpay-gateway-api-dto-apicontractconfirmterm"></a>
#### com.alibaba.china.share.creditpay.gateway.api.dto.ApiContractConfirmTerm

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `contractConfirmType` | String | 是 | 买家手动确认收货： BUYER_MANUAL_CONFIRM 自动确认收货 ：AUTO_CONFIRM ，传值为这个时：autoConfirmCondition 不可以为空； | BUYER_MANUAL_CONFIRM |
| `autoConfirmCondition` | [message:com.alibaba.china.share.creditpay.gateway.api.dto.ApiContractAutoConfirmCondition](#m-com-alibaba-china-share-creditpay-gateway-api-dto-apicontractautoconfirmcondition) | 是 | 自动确认收货条件，当contractConfirmType 取值 AUTO_CONFIRM 不可以为空 | {   "contractConfirmType": "AUTO_CONFIRM",   "autoConfirmCondition": {     "conditionEnum": "AFTER_DELIVERY",     "xDay": 7   } } |

<a id="m-com-alibaba-china-share-creditpay-gateway-api-dto-apicontractautoconfirmcondition"></a>
#### com.alibaba.china.share.creditpay.gateway.api.dto.ApiContractAutoConfirmCondition

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `conditionEnum` | String | 是 | 条件枚举 | PAY_SUCCESS_TIME_PLUS_X_DAYS |
| `xDay` | Integer | 是 | 天数 | 1 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:alibaba.china.share.creditpay.gateway.api.response.ContractDraftAndSignApplyApiResponse](#m-alibaba-china-share-creditpay-gateway-api-response-contractdraftandsignapplyapiresponse) | 是 | 影响参数 | {   "success": true,   "code": "200",   "message": "合同起草并签署申请成功",   "draftNo": "DRAFT202601051853001",   "contractStatus": "DRAFTED",   "signUrl": "https://sign.example.com/contract/DRAFT202601051853001" } |

<a id="m-alibaba-china-share-creditpay-gateway-api-response-contractdraftandsignapplyapiresponse"></a>
#### alibaba.china.share.creditpay.gateway.api.response.ContractDraftAndSignApplyApiResponse

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `draftNo` | java.lang.String | 是 | 采购单号 | 88SYT1233333 |
| `contractCurrentStatus` | java.lang.String | 是 | 当前状态 | PAYING |
| `success` | java.lang.Boolean | 是 | 是否成功 | true |
| `errorCode` | java.lang.String | 是 | 错误码 | SUCCESS |
| `errorMsg` | java.lang.String | 是 | 错误信息 | 调用成功 |
| `traceId` | String | 是 | 调用trace | 123322 |
| `counterpartyConfirmUrl` | String | 是 | 需要对手方认领时不为空，返回认领 URL | https://syt.1688.com/page/SYT/seller-contract-simple?draftNo={draftNo}&tracelog=api&isInvite=true |

## 示例

**入参示例**

```
{
  "request": {
    "drafterUserId": "xxx",
    "drafterRole": "PART_A",
    "origin": "ERP",
    "requestNo": "xx",
    "contractType": "PURCHASE_ORDER",
    "draftNo": "88SYT2015121000013",
    "payTerm": {
      "amount": 1,
      "payMethod": "SINGLE",
      "payTimeType": "IMMEDIATELY",
      "payTimeContent": "无"
    },
    "contentTerm": {
      "contentType": "PURCHASE_ORDER_DETAIL",
      "purchaseItems": [
        "[{\"productName\":\"产品名称\",\"productSpec\":\"1\",\"quantity\":1}]"
      ]
    },
    "contractRoleInfoTerm": {
      "counterpartyOrigin": "LOGIN_ID_1688_MATCH",
      "counterpartyLoginId": "ces测试xx",
      "counterpartyName": "张三",
      "counterpartyLicenseNo": "310222201512109985",
      "counterpartyLicenseType": "IDENTITY_CARD"
    }
  }
}
```

**出参示例**

```
{
  "result": {
    "draftNo": "88SYT202512301454",
    "contractCurrentStatus": "PAYING"
  }
}
```
