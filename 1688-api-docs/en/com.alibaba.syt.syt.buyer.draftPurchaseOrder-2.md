# 88 ShengYiTong: buyer drafts purchase order

Original name: 88生意通买家起草采购单  
API: `com.alibaba.syt:syt.buyer.draftPurchaseOrder:2` · Category: 88 ShengYiTong (Business Link)  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.syt:syt.buyer.draftPurchaseOrder-2  
Request URL: `https://gw.open.1688.com/openapi/param2/2/com.alibaba.syt/syt.buyer.draftPurchaseOrder/{appKey}`  
Requires user authorization (access_token) · Requires signature

88 ShengYiTong solution. The buyer drafts a purchase order.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `request` | [message:alibaba.china.share.creditpay.gateway.api.request.ContractDraftAndSignApplyApiRequest](#m-alibaba-china-share-creditpay-gateway-api-request-contractdraftandsignapplyapirequest) | yes | Request parameters | {     "origin": "ERP",     "requestNo": "REQ2026051100001",     "contractType": "PURCHASE_ORDER",     "draftNo": null,     "payTerm": {         "amount": 10000,         "payMethod": "SINGLE",         "payTimeType": "IMMEDIATELY",         "payTimeContent": null     },     "contentTerm": {         "contentType": "PURCHASE_ORDER",         "purchaseItems": [             {                 "key": "SKU001",                 "productName": "商品A",                 "productSpec": "红色/L码",                 "quantity": 10,                 "unitPrice": 1000,                 "subtotal": 10000,                 "productImage": {                     "fileName": "product_a.jpg",                     "fileUrl": "https://example.com/images/product_a.jpg"                 }             }         ],         "attachments": [          ]     },     "contractRoleInfoTerm": {         "counterpartyOrigin": "LOGIN_ID_1688_MATCH",         "counterpartyLoginId": "seller1688",         "counterpartyName": null,         "counterpartyLicenseNo": null,         "counterpartyLicenseType": null     },     "postFeeTerm": {         "postFee": 0     },     "confirmTerm": {         "contractConfirmType": "BUYER_MANUAL_CONFIRM",         "autoConfirmCondition": null     },     "drafterNeedSignConfirmAgain": false } |

<a id="m-alibaba-china-share-creditpay-gateway-api-request-contractdraftandsignapplyapirequest"></a>
#### alibaba.china.share.creditpay.gateway.api.request.ContractDraftAndSignApplyApiRequest

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `origin` | java.lang.String | yes | Call source | ERP，固定值 |
| `requestNo` | java.lang.String | yes | Request idempotency key | 本次创建采购单唯一请求号 |
| `draftNo` | java.lang.String | yes | Purchase order number - leave empty when creating | 空，不传值 |
| `payTerm` | [message:alibaba.china.share.creditpay.gateway.api.dto.ApiContractPayTerm](#m-alibaba-china-share-creditpay-gateway-api-dto-apicontractpayterm) | yes | Payment terms | {     "amount": 10000.00,     "payMethod": "SINGLE",     "payTimeType": "IMMEDIATELY",     "payTimeContent": null   } |
| `contentTerm` | [message:alibaba.china.share.creditpay.gateway.api.dto.ApiContractContentTerm](#m-alibaba-china-share-creditpay-gateway-api-dto-apicontractcontentterm) | yes | Content terms | {     "contentType": "PURCHASE_ORDER",     "purchaseItems": [       {         "key": "SKU001",         "productName": "商品A",         "productSpec": "红色/L码",         "quantity": 10,         "unitPrice": 1000.00,         "subtotal": 10000.00,         "productImage": {           "fileName": "product_a.jpg",           "fileUrl": "https://example.com/images/product_a.jpg"         }       }     ],     "attachments": []   } |
| `contractRoleInfoTerm` | [message:alibaba.china.share.creditpay.gateway.api.dto.ApiContractRoleInfoTerm](#m-alibaba-china-share-creditpay-gateway-api-dto-apicontractroleinfoterm) | yes | Counterparty agreement | {     "counterpartyOrigin": "LOGIN_ID_1688_MATCH",     "counterpartyLoginId": "seller1688",     "counterpartyName": null,     "counterpartyLicenseNo": null,     "counterpartyLicenseType": null   } |
| `drafterNeedSignConfirmAgain` | Boolean | yes | Whether the drafting party needs secondary confirmation; false means not needed | false |
| `postFeeTerm` | [message:com.alibaba.china.share.creditpay.gateway.api.dto.ApiContractPostFeeTerm](#m-com-alibaba-china-share-creditpay-gateway-api-dto-apicontractpostfeeterm) | yes | Postage terms | {   "postFee": 1.00 } |
| `confirmTerm` | [message:com.alibaba.china.share.creditpay.gateway.api.dto.ApiContractConfirmTerm](#m-com-alibaba-china-share-creditpay-gateway-api-dto-apicontractconfirmterm) | yes | Confirm receipt terms | {   "contractConfirmType": "AUTO_CONFIRM",   "autoConfirmCondition": {     "conditionEnum": "AFTER_DELIVERY",     "xDay": 7   } } |

<a id="m-alibaba-china-share-creditpay-gateway-api-dto-apicontractpayterm"></a>
#### alibaba.china.share.creditpay.gateway.api.dto.ApiContractPayTerm

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `amount` | BigDecimal | yes | Payment amount, unit: yuan, currency: RMB (CNY) | 1 |
| `payMethod` | java.lang.String | yes | Payment method - fixed value | SINGLE |
| `payTimeType` | java.lang.String | yes | Payment time type - fixed value | IMMEDIATELY |
| `payTimeContent` | java.lang.String | no | Payment time content | 不传 |

<a id="m-alibaba-china-share-creditpay-gateway-api-dto-apicontractcontentterm"></a>
#### alibaba.china.share.creditpay.gateway.api.dto.ApiContractContentTerm

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `contentType` | java.lang.String | yes | Content type - fixed value | PURCHASE_ORDER |
| `purchaseItems` | [message:alibaba.china.share.creditpay.gateway.api.dto.ApiPurchaseItem[]](#m-alibaba-china-share-creditpay-gateway-api-dto-apipurchaseitem[]) | yes | Purchase order details | [       {         "key": "SKU001",         "productName": "商品A",         "productSpec": "红色/L码",         "quantity": 10,         "unitPrice": 1000.00,         "subtotal": 10000.00       } ] |

<a id="m-alibaba-china-share-creditpay-gateway-api-dto-apipurchaseitem[]"></a>
#### alibaba.china.share.creditpay.gateway.api.dto.ApiPurchaseItem[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `key` | java.lang.String | yes | Purchase order sequence number | sku001 |
| `productName` | java.lang.String | yes | Product name | 水杯 |
| `productSpec` | java.lang.String | yes | Product sku | 1L、白色、大肚款 |
| `quantity` | java.lang.Integer | yes | Purchase quantity | 1 |
| `unitPrice` | BigDecimal | yes | Unit price, in yuan | 1 |
| `subtotal` | BigDecimal | yes | Total price, in yuan | 1 |

<a id="m-alibaba-china-share-creditpay-gateway-api-dto-apicontractroleinfoterm"></a>
#### alibaba.china.share.creditpay.gateway.api.dto.ApiContractRoleInfoTerm

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `counterpartyOrigin` | java.lang.String | yes | LOGIN_ID_1688_MATCH (specify the counterparty via 1688 loginId), DRAFTER_INPUT (specify via the counterparty's entity information), COUNTERPARTY_INPUT (not specified, left for the counterparty to claim themselves) | LOGIN_ID_1688_MATCH |
| `counterpartyLoginId` | java.lang.String | no | 1688loginId, required when counterpartyOrigin = LOGIN_ID_1688_MATCH | ces测试002 |
| `counterpartyName` | java.lang.String | no | Counterparty name. Required when counterpartyOrigin = DRAFTER_INPUT | 张三 |
| `counterpartyLicenseNo` | java.lang.String | no | Counterparty's ID card number or business license number. Required when counterpartyOrigin = DRAFTER_INPUT | 330xx |
| `counterpartyLicenseType` | java.lang.String | no | ID document type; required when counterpartyOrigin = DRAFTER_INPUT | 统一社会信用代码UNITY |

<a id="m-com-alibaba-china-share-creditpay-gateway-api-dto-apicontractpostfeeterm"></a>
#### com.alibaba.china.share.creditpay.gateway.api.dto.ApiContractPostFeeTerm

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `postFee` | BigDecimal | yes | postFee | 0.01 |

<a id="m-com-alibaba-china-share-creditpay-gateway-api-dto-apicontractconfirmterm"></a>
#### com.alibaba.china.share.creditpay.gateway.api.dto.ApiContractConfirmTerm

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `contractConfirmType` | String | yes | Buyer manually confirms receipt: BUYER_MANUAL_CONFIRM; automatically confirms receipt: AUTO_CONFIRM. When this value is passed, autoConfirmCondition cannot be empty; | BUYER_MANUAL_CONFIRM |
| `autoConfirmCondition` | [message:com.alibaba.china.share.creditpay.gateway.api.dto.ApiContractAutoConfirmCondition](#m-com-alibaba-china-share-creditpay-gateway-api-dto-apicontractautoconfirmcondition) | yes | Auto-confirm receipt condition; cannot be empty when contractConfirmType is AUTO_CONFIRM | {   "contractConfirmType": "AUTO_CONFIRM",   "autoConfirmCondition": {     "conditionEnum": "AFTER_DELIVERY",     "xDay": 7   } } |

<a id="m-com-alibaba-china-share-creditpay-gateway-api-dto-apicontractautoconfirmcondition"></a>
#### com.alibaba.china.share.creditpay.gateway.api.dto.ApiContractAutoConfirmCondition

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `conditionEnum` | String | yes | Condition enum | PAY_SUCCESS_TIME_PLUS_X_DAYS |
| `xDay` | Integer | yes | Number of days | 1 |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.china.share.creditpay.gateway.api.response.ContractDraftAndSignApplyApiResponse](#m-alibaba-china-share-creditpay-gateway-api-response-contractdraftandsignapplyapiresponse) | yes | Affected parameters | {   "success": true,   "code": "200",   "message": "合同起草并签署申请成功",   "draftNo": "DRAFT202601051853001",   "contractStatus": "DRAFTED",   "signUrl": "https://sign.example.com/contract/DRAFT202601051853001" } |

<a id="m-alibaba-china-share-creditpay-gateway-api-response-contractdraftandsignapplyapiresponse"></a>
#### alibaba.china.share.creditpay.gateway.api.response.ContractDraftAndSignApplyApiResponse

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `draftNo` | java.lang.String | yes | Purchase order number | 88SYT1233333 |
| `contractCurrentStatus` | java.lang.String | yes | Current status | PAYING |
| `success` | java.lang.Boolean | yes | Whether successful | true |
| `errorCode` | java.lang.String | yes | Error code | SUCCESS |
| `errorMsg` | java.lang.String | yes | Error message | 调用成功 |
| `traceId` | String | yes | Call trace | 123322 |
| `counterpartyConfirmUrl` | String | yes | Not empty when the counterparty needs to claim it; returns the claim URL | https://syt.1688.com/page/SYT/seller-contract-simple?draftNo={draftNo}&tracelog=api&isInvite=true |

## Samples

**Input parameter example**

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

**Output parameter example**

```
{
  "result": {
    "draftNo": "88SYT202512301454",
    "contractCurrentStatus": "PAYING"
  }
}
```
