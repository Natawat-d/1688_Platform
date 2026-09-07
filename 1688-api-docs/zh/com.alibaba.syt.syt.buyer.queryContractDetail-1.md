# 88生意通买家查询采购单详情

API: `com.alibaba.syt:syt.buyer.queryContractDetail:1` · Category: 88生意通  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.syt:syt.buyer.queryContractDetail-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.syt/syt.buyer.queryContractDetail/{appKey}`  
需要授权 (access_token) · 需要签名

88生意通买家查询采购单详情

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `request` | [message:alibaba.china.share.creditpay.gateway.api.request.ContractQueryDetailApiRequest](#m-alibaba-china-share-creditpay-gateway-api-request-contractquerydetailapirequest) | 是 | 请求 | {     "draftNo": "CT2026051100001"     } |

<a id="m-alibaba-china-share-creditpay-gateway-api-request-contractquerydetailapirequest"></a>
#### alibaba.china.share.creditpay.gateway.api.request.ContractQueryDetailApiRequest

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `draftNo` | java.lang.String | 是 | 采购单编号 | CT2026051100001 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:alibaba.china.share.creditpay.gateway.api.response.ContractQueryDetailApiResponse](#m-alibaba-china-share-creditpay-gateway-api-response-contractquerydetailapiresponse) | 是 | 响应 | {   "contractDTO": {     "amount": 0.01,     "contentTerm": {       "contentType": "PURCHASE_ORDER",       "purchaseItems": [         {           "key": "fd8ede4d-58bd-4a08-a4ae-cde6790330b5",           "productName": "杯子",           "productSpec": "黑色",           "quantity": 1,           "subtotal": 0.01,           "unitPrice": 0.01         }       ]     },     "draftNo": "88SYT20260415598003",     "partAName": "杭州XXXX公司",     "partBName": "XXX",     "status": "SIGN_SUCCESS"   },   "errorCode": "SUCCESS",   "traceId": "213e01f717791723747816123ea56f" } |

<a id="m-alibaba-china-share-creditpay-gateway-api-response-contractquerydetailapiresponse"></a>
#### alibaba.china.share.creditpay.gateway.api.response.ContractQueryDetailApiResponse

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `contractDTO` | [message:alibaba.china.share.creditpay.gateway.api.dto.ContractApiDTO](#m-alibaba-china-share-creditpay-gateway-api-dto-contractapidto) | 是 | 采购单详情 | {   "amount": 0.01,   "contentTerm": {     "contentType": "PURCHASE_ORDER",     "purchaseItems": [       {         "key": "fd8ede4d-58bd-4a08-a4ae-cde6790330b5",         "productName": "杯子",         "productSpec": "黑色",         "quantity": 1,         "subtotal": 0.01,         "unitPrice": 0.01       }     ]   },   "draftNo": "88SYT20260415598003",   "partAName": "杭州XXX限公司",   "partBName": "XX",   "status": "SIGN_SUCCESS" } |
| `errorCode` | java.lang.String | 是 | 错误 code | 错误 code |
| `errorMsg` | java.lang.String | 是 | 调用结果描述 | 调用成功 |
| `isSuccess` | java.lang.Boolean | 是 | 是否成功 | true |
| `traceId` | String | 是 | 调用 traceId | traceId |
| `fundPayOrderList` | [message:com.alibaba.china.share.creditpay.gateway.api.dto.FundPayOrderApiDTO[]](#m-com-alibaba-china-share-creditpay-gateway-api-dto-fundpayorderapidto[]) | 是 | 支付单列表 | [     {       "payNo": "8896010FP09006126043000000618013",       "payTime": 1777543672000,       "payChannel": "BANK_TRANSFER",       "class": "com.alibaba.china.share.creditpay.gateway.api.dto.FundPayOrderApiDTO"     }   ] |

<a id="m-alibaba-china-share-creditpay-gateway-api-dto-contractapidto"></a>
#### alibaba.china.share.creditpay.gateway.api.dto.ContractApiDTO

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `amount` | java.math.BigDecimal | 是 | 金额 | 10000.00 |
| `contentTerm` | [message:alibaba.china.share.creditpay.gateway.api.dto.ApiContractContentTerm](#m-alibaba-china-share-creditpay-gateway-api-dto-apicontractcontentterm) | 是 | 内容明细 | {   "contentType": "PURCHASE_ORDER",   "purchaseItems": [     {       "key": "SKU001",       "productName": "商品A",       "quantity": 10,       "unitPrice": 100.00,       "subtotal": 1000.00     }   ],   "attachments": [] } |
| `draftNo` | java.lang.String | 是 | 采购单编号 | 88SYT2026051100001 |
| `partAName` | java.lang.String | 是 | 买家名称 | 张三 |
| `partBName` | java.lang.String | 是 | 卖家名称 | 李四公司 |
| `status` | java.lang.String | 是 | 状态- DRAFT（草稿状态，可编辑）、DATA_SUPPLYING（补充资料中，也是草稿态一种）、SIGNING（签署中）、SIGN_REJECT（拒绝签署）、  SIGN_SUCCESS（签署完成，待付款）、PAYING（付款中）、PAID（支付完成）、CONFIRMED（买家已确认交易完成）、  REFUNDING（退款中）、FINISHED（交易完成）、CLOSED（采购单关闭-已退款）、INVALID（合同/采购单失效）、DELETED（采购单已删除） | CONFIRMED |
| `supplyVoucherList` | [message:alibaba.china.share.creditpay.gateway.api.dto.ContractSupplyVoucherApiDTO[]](#m-alibaba-china-share-creditpay-gateway-api-dto-contractsupplyvoucherapidto[]) | 是 | 交易凭证列表 | [       {         "voucherNo": "VOUCHER001",         "content": "货物已收到",         "createTime": "2026-05-11 10:00:00"       }     ] |
| `buyerContactName` | String | 是 | 买家联系人姓名 | 买家联系人姓名 |
| `buyerContactPhone` | String | 是 | 买家联系电话 | 15611112222 |
| `sellerContactName` | String | 是 | 卖家联系人姓名 | 卖家联系人姓名 |
| `sellerContactPhone` | String | 是 | 卖家联系电话 | 15611112222 |
| `confirmReceiveRule` | String | 是 | 确认收货规则 | AUTO |
| `autoConfirmTime` | Date | 是 | 自动确认收货时间，仅自动确认规则下有值 | 123456789 |

<a id="m-alibaba-china-share-creditpay-gateway-api-dto-apicontractcontentterm"></a>
#### alibaba.china.share.creditpay.gateway.api.dto.ApiContractContentTerm

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `contentType` | java.lang.String | 是 | 内容类型 | PURCHASE_ORDER |
| `purchaseItems` | [message:alibaba.china.share.creditpay.gateway.api.dto.ApiPurchaseItem[]](#m-alibaba-china-share-creditpay-gateway-api-dto-apipurchaseitem[]) | 是 | 采购单明细 | {       "contentType": "PURCHASE_ORDER",       "purchaseItems": [         {           "key": "SKU002",           "productName": "商品B",           "productSpec": "蓝色/M码",           "quantity": 5,           "unitPrice": 200.00,           "subtotal": 1000.00,           "productImage": null         }       ],       "attachments": [         {           "fileName": "contract.pdf",           "fileUrl": "https://example.com/files/contract.pdf"         }       ]     } |

<a id="m-alibaba-china-share-creditpay-gateway-api-dto-apipurchaseitem[]"></a>
#### alibaba.china.share.creditpay.gateway.api.dto.ApiPurchaseItem[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `key` | java.lang.String | 是 | 编号 | SKU002 |
| `productName` | java.lang.String | 是 | 产品名称 | 商品B |
| `productSpec` | java.lang.String | 是 | 产品规则 | 蓝色/M码 |
| `quantity` | java.lang.Integer | 是 | 数量 | 1 |
| `subtotal` | java.math.BigDecimal | 是 | 总价，单位人民币，元 | 1 |
| `unitPrice` | java.math.BigDecimal | 是 | 单价，单位人民币，元 | 1 |

<a id="m-alibaba-china-share-creditpay-gateway-api-dto-contractsupplyvoucherapidto[]"></a>
#### alibaba.china.share.creditpay.gateway.api.dto.ContractSupplyVoucherApiDTO[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `submitterName` | java.lang.String | 是 | 提交人名字 | 张三 |
| `submitterRole` | java.lang.String | 是 | 补充角色，PART_A 代表买家，PART_B打标卖家 | PART_A |
| `supplyTime` | java.util.Date | 是 | 补充时间 | 2026-05-11 10:00:00 |
| `voucherContent` | java.lang.String | 是 | 凭证内容 | 货物已收到，质量符合要求 |
| `voucherFiles` | [message:alibaba.china.share.creditpay.gateway.api.dto.ApiFileItem[]](#m-alibaba-china-share-creditpay-gateway-api-dto-apifileitem[]) | 是 | 图片 | [           {             "fileName": "delivery_proof.jpg",             "fileUrl": "https://example.com/files/delivery_proof.jpg"           }         ]       } |
| `voucherNo` | java.lang.String | 是 | 凭证记录唯一标识 | VOUCHER001 |

<a id="m-alibaba-china-share-creditpay-gateway-api-dto-apifileitem[]"></a>
#### alibaba.china.share.creditpay.gateway.api.dto.ApiFileItem[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `fileData` | byte[] | 是 | 二进制数据，和 fileUrl 二选一 | 12 |
| `fileName` | java.lang.String | 是 | 文件名，上传带有后缀的文件名 | contract.pdf |
| `fileUrl` | java.lang.String | 是 | 文件 URL，和 fileData 二选一就可以 | https://example.com/files/contract.pdf |

<a id="m-com-alibaba-china-share-creditpay-gateway-api-dto-fundpayorderapidto[]"></a>
#### com.alibaba.china.share.creditpay.gateway.api.dto.FundPayOrderApiDTO[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `payNo` | String | 是 | 支付单号 | 12345777777 |
| `payTime` | Date | 是 | 支付时间 | 12345777777 |
| `payChannel` | String | 是 | 支付渠道 | BANK_TRANSFER |
