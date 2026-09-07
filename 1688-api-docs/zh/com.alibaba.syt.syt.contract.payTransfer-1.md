# 88生意通万里汇转账支付-仅支持万里汇B2C账号

API: `com.alibaba.syt:syt.contract.payTransfer:1` · Category: 88生意通  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.syt:syt.contract.payTransfer-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.syt/syt.contract.payTransfer/{appKey}`  
需要授权 (access_token) · 需要签名

88生意通支付转账接口

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `request` | [message:alibaba.china.share.creditpay.gateway.api.request.ContractPayTransferApiRequest](#m-alibaba-china-share-creditpay-gateway-api-request-contractpaytransferapirequest) | 是 | 请求入参 | {   "draftNo": "DRAFT20251230001",   "userId": "user123456",   "payOrderNo": "PAY20251230001",   "escrowAccount": {     "accountNumber": "6222021234567890123",     "bankName": "中国工商银行",     "bankBranchCode": "102290000123"   },   "transferToAmount": 100000,   "transferFromCurrency": "USD",   "transferFundChannel": "WORLDFIRST",   "transferFundChannelConfig": {     "channelCode": "WF_USD",     "accountType": "BUSINESS"   } } |

<a id="m-alibaba-china-share-creditpay-gateway-api-request-contractpaytransferapirequest"></a>
#### alibaba.china.share.creditpay.gateway.api.request.ContractPayTransferApiRequest

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `draftNo` | java.lang.String | 是 | 合同起草号 | 88SYT20251204069031 |
| `payOrderNo` | java.lang.String | 是 | 支付单号 | 88999xxx |
| `escrowAccount` | [message:alibaba.cbu.institution.platform.api.common.dto.account.BankAccount](#m-alibaba-cbu-institution-platform-api-common-dto-account-bankaccount) | 是 | 转入的资金托管账户 | {   "accountName": "张三",   "accountNo": "6228765030160001109",   "bankName": "江苏银行股份有限公司" } |
| `transferToAmount` | BigDecimal | 是 | 转入资金金额，等于合同金额，单位元，币种人民币 CNY | 1 |
| `transferFromCurrency` | java.lang.String | 是 | 转出币种，美元 USD | USD |
| `transferFundChannel` | java.lang.String | 是 | 资金渠道，规定值 | WORLD_FIRST |
| `transferFundChannelConfig` | [message:com.alibaba.china.share.creditpay.gateway.api.dto.ApiWfTransferFundChannelConfig](#m-com-alibaba-china-share-creditpay-gateway-api-dto-apiwftransferfundchannelconfig) | 是 | 万里汇客户配置 | {"clientId":"xxx","customerId":"万里汇客户 ID"} |

<a id="m-alibaba-cbu-institution-platform-api-common-dto-account-bankaccount"></a>
#### alibaba.cbu.institution.platform.api.common.dto.account.BankAccount

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `cardHolderType` | java.lang.String | 是 | 收款人类型，企业 COMPANY 还是个人 PERSON | COMPANY |
| `accountNumber` | java.lang.String | 是 | 收款人账号 | 6228765030160001109 |
| `holder` | java.lang.String | 是 | 收款人姓名 | 张三 |

<a id="m-com-alibaba-china-share-creditpay-gateway-api-dto-apiwftransferfundchannelconfig"></a>
#### com.alibaba.china.share.creditpay.gateway.api.dto.ApiWfTransferFundChannelConfig

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `clientId` | java.lang.String | 是 | 万里汇客户 ID | 11 |
| `customerId` | java.lang.String | 是 | 万里汇客户 ID | 11 |
| `privateKeyCiphertext` | java.lang.String | 是 | 万里汇客户 密钥密文 | 11 |
| `privateKeyVersion` | java.lang.String | 是 | 万里汇客户 密钥版本 | 1 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:alibaba.china.share.creditpay.gateway.api.response.ContractPayTransferApiResponse](#m-alibaba-china-share-creditpay-gateway-api-response-contractpaytransferapiresponse) | 是 | 响应结果 | {   "success": true,   "code": "200",   "message": "转账成功",   "escrowAccount": {     "accountNumber": "6222021234567890123",     "bankName": "中国工商银行",     "bankBranchCode": "102290000123"   },   "transferToCurrency": "CNY",   "transferToAmount": 100000,   "transferFromAmount": 13850,   "transferFromCurrency": "USD",   "quotePrice": 7.22,   "feeAmount": 50,   "feeCurrency": "USD" } |

<a id="m-alibaba-china-share-creditpay-gateway-api-response-contractpaytransferapiresponse"></a>
#### alibaba.china.share.creditpay.gateway.api.response.ContractPayTransferApiResponse

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `escrowAccount` | [message:alibaba.cbu.institution.platform.api.common.dto.account.BankAccount](#m-alibaba-cbu-institution-platform-api-common-dto-account-bankaccount) | 是 | 资金转入账户信息 | {   "accountName": "张三",   "accountNo": "6228765030160001109",   "bankName": "江苏银行股份有限公司" } |
| `transferToCurrency` | java.lang.String | 是 | 目标币种 | CNY |
| `transferToAmount` | java.lang.Long | 是 | 目标金额，合同金额，单位分 | 1 |
| `transferFromAmount` | java.lang.Long | 是 | 转出金额，单位美分，USD | 1 |
| `transferFromCurrency` | java.lang.String | 是 | 支付币种 | USD |
| `quotePrice` | java.math.BigDecimal | 是 | 汇率 | 6.89 |
| `feeAmount` | java.lang.Long | 是 | 费用金额，单位美分 | 1 |
| `feeCurrency` | java.lang.String | 是 | 费用币种 | USD |
| `success` | String | 是 | 调用是否成功，true 成功，false 失败 | true |
| `errorCode` | String | 是 | 调用 code |  SUCCESS |
| `errorMsg` | String | 是 | 调用结果描述 | 调用成功 |
| `traceId` | String | 是 | traceId，排查问题 | 1122 |

<a id="m-alibaba-cbu-institution-platform-api-common-dto-account-bankaccount"></a>
#### alibaba.cbu.institution.platform.api.common.dto.account.BankAccount

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `cardHolderType` | java.lang.String | 是 | 收款人类型，企业 COMPANY 还是个人 PERSON | COMPANY |
| `accountNumber` | java.lang.String | 是 | 收款人账号 | 6228765030160001109 |
| `holder` | java.lang.String | 是 | 收款人姓名 | 张三 |
