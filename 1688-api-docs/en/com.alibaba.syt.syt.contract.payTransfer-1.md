# 88 ShengYiTong: WorldFirst transfer payment (WorldFirst B2C accounts only)

Original name: 88生意通万里汇转账支付-仅支持万里汇B2C账号  
API: `com.alibaba.syt:syt.contract.payTransfer:1` · Category: 88 ShengYiTong (Business Link)  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.syt:syt.contract.payTransfer-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.syt/syt.contract.payTransfer/{appKey}`  
Requires user authorization (access_token) · Requires signature

88 ShengYiTong payment-transfer interface.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `request` | [message:alibaba.china.share.creditpay.gateway.api.request.ContractPayTransferApiRequest](#m-alibaba-china-share-creditpay-gateway-api-request-contractpaytransferapirequest) | yes | Request input parameter | {   "draftNo": "DRAFT20251230001",   "userId": "user123456",   "payOrderNo": "PAY20251230001",   "escrowAccount": {     "accountNumber": "6222021234567890123",     "bankName": "中国工商银行",     "bankBranchCode": "102290000123"   },   "transferToAmount": 100000,   "transferFromCurrency": "USD",   "transferFundChannel": "WORLDFIRST",   "transferFundChannelConfig": {     "channelCode": "WF_USD",     "accountType": "BUSINESS"   } } |

<a id="m-alibaba-china-share-creditpay-gateway-api-request-contractpaytransferapirequest"></a>
#### alibaba.china.share.creditpay.gateway.api.request.ContractPayTransferApiRequest

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `draftNo` | java.lang.String | yes | Contract drafting number | 88SYT20251204069031 |
| `payOrderNo` | java.lang.String | yes | Payment order number | 88999xxx |
| `escrowAccount` | [message:alibaba.cbu.institution.platform.api.common.dto.account.BankAccount](#m-alibaba-cbu-institution-platform-api-common-dto-account-bankaccount) | yes | Escrow account the funds are transferred into | {   "accountName": "张三",   "accountNo": "6228765030160001109",   "bankName": "江苏银行股份有限公司" } |
| `transferToAmount` | BigDecimal | yes | Transferred fund amount, equal to the contract amount, unit: yuan, currency: RMB CNY | 1 |
| `transferFromCurrency` | java.lang.String | yes | Source currency, US Dollar USD | USD |
| `transferFundChannel` | java.lang.String | yes | Funding channel, specified value | WORLD_FIRST |
| `transferFundChannelConfig` | [message:com.alibaba.china.share.creditpay.gateway.api.dto.ApiWfTransferFundChannelConfig](#m-com-alibaba-china-share-creditpay-gateway-api-dto-apiwftransferfundchannelconfig) | yes | WorldFirst customer configuration | {"clientId":"xxx","customerId":"万里汇客户 ID"} |

<a id="m-alibaba-cbu-institution-platform-api-common-dto-account-bankaccount"></a>
#### alibaba.cbu.institution.platform.api.common.dto.account.BankAccount

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `cardHolderType` | java.lang.String | yes | Payee type, enterprise COMPANY or individual PERSON | COMPANY |
| `accountNumber` | java.lang.String | yes | Payee account | 6228765030160001109 |
| `holder` | java.lang.String | yes | Payee name | 张三 |

<a id="m-com-alibaba-china-share-creditpay-gateway-api-dto-apiwftransferfundchannelconfig"></a>
#### com.alibaba.china.share.creditpay.gateway.api.dto.ApiWfTransferFundChannelConfig

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `clientId` | java.lang.String | yes | Wanlihui (WorldFirst) customer ID | 11 |
| `customerId` | java.lang.String | yes | Wanlihui (WorldFirst) customer ID | 11 |
| `privateKeyCiphertext` | java.lang.String | yes | WorldFirst customer key ciphertext | 11 |
| `privateKeyVersion` | java.lang.String | yes | WorldFirst customer key version | 1 |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.china.share.creditpay.gateway.api.response.ContractPayTransferApiResponse](#m-alibaba-china-share-creditpay-gateway-api-response-contractpaytransferapiresponse) | yes | Response result | {   "success": true,   "code": "200",   "message": "转账成功",   "escrowAccount": {     "accountNumber": "6222021234567890123",     "bankName": "中国工商银行",     "bankBranchCode": "102290000123"   },   "transferToCurrency": "CNY",   "transferToAmount": 100000,   "transferFromAmount": 13850,   "transferFromCurrency": "USD",   "quotePrice": 7.22,   "feeAmount": 50,   "feeCurrency": "USD" } |

<a id="m-alibaba-china-share-creditpay-gateway-api-response-contractpaytransferapiresponse"></a>
#### alibaba.china.share.creditpay.gateway.api.response.ContractPayTransferApiResponse

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `escrowAccount` | [message:alibaba.cbu.institution.platform.api.common.dto.account.BankAccount](#m-alibaba-cbu-institution-platform-api-common-dto-account-bankaccount) | yes | Fund transfer-in account information | {   "accountName": "张三",   "accountNo": "6228765030160001109",   "bankName": "江苏银行股份有限公司" } |
| `transferToCurrency` | java.lang.String | yes | Target currency | CNY |
| `transferToAmount` | java.lang.Long | yes | Target amount, contract amount, unit: cent | 1 |
| `transferFromAmount` | java.lang.Long | yes | Transfer-out amount, unit: US cents, USD | 1 |
| `transferFromCurrency` | java.lang.String | yes | Payment currency | USD |
| `quotePrice` | java.math.BigDecimal | yes | Exchange rate | 6.89 |
| `feeAmount` | java.lang.Long | yes | Fee amount, in US cents | 1 |
| `feeCurrency` | java.lang.String | yes | Fee currency | USD |
| `success` | String | yes | Whether the call succeeded, true for success, false for failure | true |
| `errorCode` | String | yes | Call code |  SUCCESS |
| `errorMsg` | String | yes | Call result description | 调用成功 |
| `traceId` | String | yes | traceId, for troubleshooting | 1122 |

<a id="m-alibaba-cbu-institution-platform-api-common-dto-account-bankaccount"></a>
#### alibaba.cbu.institution.platform.api.common.dto.account.BankAccount

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `cardHolderType` | java.lang.String | yes | Payee type, enterprise COMPANY or individual PERSON | COMPANY |
| `accountNumber` | java.lang.String | yes | Payee account | 6228765030160001109 |
| `holder` | java.lang.String | yes | Payee name | 张三 |
