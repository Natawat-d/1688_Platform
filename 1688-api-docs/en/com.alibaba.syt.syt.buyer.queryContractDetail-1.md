# 88 ShengYiTong: buyer queries purchase-order details

Original name: 88生意通买家查询采购单详情  
API: `com.alibaba.syt:syt.buyer.queryContractDetail:1` · Category: 88 ShengYiTong (Business Link)  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.syt:syt.buyer.queryContractDetail-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.syt/syt.buyer.queryContractDetail/{appKey}`  
Requires user authorization (access_token) · Requires signature

88 ShengYiTong: the buyer queries the details of a purchase order.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `request` | [message:alibaba.china.share.creditpay.gateway.api.request.ContractQueryDetailApiRequest](#m-alibaba-china-share-creditpay-gateway-api-request-contractquerydetailapirequest) | yes | Request | {     "draftNo": "CT2026051100001"     } |

<a id="m-alibaba-china-share-creditpay-gateway-api-request-contractquerydetailapirequest"></a>
#### alibaba.china.share.creditpay.gateway.api.request.ContractQueryDetailApiRequest

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `draftNo` | java.lang.String | yes | Purchase order number | CT2026051100001 |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.china.share.creditpay.gateway.api.response.ContractQueryDetailApiResponse](#m-alibaba-china-share-creditpay-gateway-api-response-contractquerydetailapiresponse) | yes | Response | {   "contractDTO": {     "amount": 0.01,     "contentTerm": {       "contentType": "PURCHASE_ORDER",       "purchaseItems": [         {           "key": "fd8ede4d-58bd-4a08-a4ae-cde6790330b5",           "productName": "杯子",           "productSpec": "黑色",           "quantity": 1,           "subtotal": 0.01,           "unitPrice": 0.01         }       ]     },     "draftNo": "88SYT20260415598003",     "partAName": "杭州XXXX公司",     "partBName": "XXX",     "status": "SIGN_SUCCESS"   },   "errorCode": "SUCCESS",   "traceId": "213e01f717791723747816123ea56f" } |

<a id="m-alibaba-china-share-creditpay-gateway-api-response-contractquerydetailapiresponse"></a>
#### alibaba.china.share.creditpay.gateway.api.response.ContractQueryDetailApiResponse

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `contractDTO` | [message:alibaba.china.share.creditpay.gateway.api.dto.ContractApiDTO](#m-alibaba-china-share-creditpay-gateway-api-dto-contractapidto) | yes | Purchase order details | {   "amount": 0.01,   "contentTerm": {     "contentType": "PURCHASE_ORDER",     "purchaseItems": [       {         "key": "fd8ede4d-58bd-4a08-a4ae-cde6790330b5",         "productName": "杯子",         "productSpec": "黑色",         "quantity": 1,         "subtotal": 0.01,         "unitPrice": 0.01       }     ]   },   "draftNo": "88SYT20260415598003",   "partAName": "杭州XXX限公司",   "partBName": "XX",   "status": "SIGN_SUCCESS" } |
| `errorCode` | java.lang.String | yes | Error code | 错误 code |
| `errorMsg` | java.lang.String | yes | Call result description | 调用成功 |
| `isSuccess` | java.lang.Boolean | yes | Whether successful | true |
| `traceId` | String | yes | Call traceId | traceId |
| `fundPayOrderList` | [message:com.alibaba.china.share.creditpay.gateway.api.dto.FundPayOrderApiDTO[]](#m-com-alibaba-china-share-creditpay-gateway-api-dto-fundpayorderapidto[]) | yes | List of payment orders | [     {       "payNo": "8896010FP09006126043000000618013",       "payTime": 1777543672000,       "payChannel": "BANK_TRANSFER",       "class": "com.alibaba.china.share.creditpay.gateway.api.dto.FundPayOrderApiDTO"     }   ] |

<a id="m-alibaba-china-share-creditpay-gateway-api-dto-contractapidto"></a>
#### alibaba.china.share.creditpay.gateway.api.dto.ContractApiDTO

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `amount` | java.math.BigDecimal | yes | Amount | 10000.00 |
| `contentTerm` | [message:alibaba.china.share.creditpay.gateway.api.dto.ApiContractContentTerm](#m-alibaba-china-share-creditpay-gateway-api-dto-apicontractcontentterm) | yes | Content details | {   "contentType": "PURCHASE_ORDER",   "purchaseItems": [     {       "key": "SKU001",       "productName": "商品A",       "quantity": 10,       "unitPrice": 100.00,       "subtotal": 1000.00     }   ],   "attachments": [] } |
| `draftNo` | java.lang.String | yes | Purchase order number | 88SYT2026051100001 |
| `partAName` | java.lang.String | yes | Buyer name | 张三 |
| `partBName` | java.lang.String | yes | Seller name | 李四公司 |
| `status` | java.lang.String | yes | Status - DRAFT (draft status, editable), DATA_SUPPLYING (supplementing data, also a type of draft status), SIGNING (signing in progress), SIGN_REJECT (signing rejected),  SIGN_SUCCESS (signing completed, pending payment), PAYING (payment in progress), PAID (payment completed), CONFIRMED (buyer has confirmed transaction completion),  REFUNDING (refund in progress), FINISHED (transaction completed), CLOSED (purchase order closed - refunded), INVALID (contract/purchase order invalid), DELETED (purchase order deleted) | CONFIRMED |
| `supplyVoucherList` | [message:alibaba.china.share.creditpay.gateway.api.dto.ContractSupplyVoucherApiDTO[]](#m-alibaba-china-share-creditpay-gateway-api-dto-contractsupplyvoucherapidto[]) | yes | List of transaction vouchers | [       {         "voucherNo": "VOUCHER001",         "content": "货物已收到",         "createTime": "2026-05-11 10:00:00"       }     ] |
| `buyerContactName` | String | yes | Buyer contact name | 买家联系人姓名 |
| `buyerContactPhone` | String | yes | Buyer contact phone number | 15611112222 |
| `sellerContactName` | String | yes | Seller contact person's name | 卖家联系人姓名 |
| `sellerContactPhone` | String | yes | Seller contact phone | 15611112222 |
| `confirmReceiveRule` | String | yes | Receipt confirmation rule | AUTO |
| `autoConfirmTime` | Date | yes | Automatic receipt confirmation time; only has a value under the automatic confirmation rule. | 123456789 |

<a id="m-alibaba-china-share-creditpay-gateway-api-dto-apicontractcontentterm"></a>
#### alibaba.china.share.creditpay.gateway.api.dto.ApiContractContentTerm

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `contentType` | java.lang.String | yes | Content type | PURCHASE_ORDER |
| `purchaseItems` | [message:alibaba.china.share.creditpay.gateway.api.dto.ApiPurchaseItem[]](#m-alibaba-china-share-creditpay-gateway-api-dto-apipurchaseitem[]) | yes | Purchase order details | {       "contentType": "PURCHASE_ORDER",       "purchaseItems": [         {           "key": "SKU002",           "productName": "商品B",           "productSpec": "蓝色/M码",           "quantity": 5,           "unitPrice": 200.00,           "subtotal": 1000.00,           "productImage": null         }       ],       "attachments": [         {           "fileName": "contract.pdf",           "fileUrl": "https://example.com/files/contract.pdf"         }       ]     } |

<a id="m-alibaba-china-share-creditpay-gateway-api-dto-apipurchaseitem[]"></a>
#### alibaba.china.share.creditpay.gateway.api.dto.ApiPurchaseItem[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `key` | java.lang.String | yes | Number | SKU002 |
| `productName` | java.lang.String | yes | Product name | 商品B |
| `productSpec` | java.lang.String | yes | Product rule | 蓝色/M码 |
| `quantity` | java.lang.Integer | yes | Quantity | 1 |
| `subtotal` | java.math.BigDecimal | yes | Total price, in RMB yuan | 1 |
| `unitPrice` | java.math.BigDecimal | yes | Unit price, in RMB, yuan | 1 |

<a id="m-alibaba-china-share-creditpay-gateway-api-dto-contractsupplyvoucherapidto[]"></a>
#### alibaba.china.share.creditpay.gateway.api.dto.ContractSupplyVoucherApiDTO[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `submitterName` | java.lang.String | yes | Submitter's name | 张三 |
| `submitterRole` | java.lang.String | yes | Supplementary role. PART_A represents the buyer, PART_B is tagged as the seller. | PART_A |
| `supplyTime` | java.util.Date | yes | Supplement time | 2026-05-11 10:00:00 |
| `voucherContent` | java.lang.String | yes | Voucher content | 货物已收到，质量符合要求 |
| `voucherFiles` | [message:alibaba.china.share.creditpay.gateway.api.dto.ApiFileItem[]](#m-alibaba-china-share-creditpay-gateway-api-dto-apifileitem[]) | yes | Image | [           {             "fileName": "delivery_proof.jpg",             "fileUrl": "https://example.com/files/delivery_proof.jpg"           }         ]       } |
| `voucherNo` | java.lang.String | yes | Unique identifier of the voucher record | VOUCHER001 |

<a id="m-alibaba-china-share-creditpay-gateway-api-dto-apifileitem[]"></a>
#### alibaba.china.share.creditpay.gateway.api.dto.ApiFileItem[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `fileData` | byte[] | yes | Binary data; choose either this or fileUrl. | 12 |
| `fileName` | java.lang.String | yes | File name; upload a file name with an extension. | contract.pdf |
| `fileUrl` | java.lang.String | yes | File URL; choose either this or fileData. | https://example.com/files/contract.pdf |

<a id="m-com-alibaba-china-share-creditpay-gateway-api-dto-fundpayorderapidto[]"></a>
#### com.alibaba.china.share.creditpay.gateway.api.dto.FundPayOrderApiDTO[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `payNo` | String | yes | Payment order number | 12345777777 |
| `payTime` | Date | yes | Payment time | 12345777777 |
| `payChannel` | String | yes | Payment channel | BANK_TRANSFER |
