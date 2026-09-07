# Check whether password-free payment is enabled

Original name: 查询是否开通免密支付  
API: `com.alibaba.trade:alibaba.trade.pay.protocolPay.isopen:1` · Category: Payment  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.trade.pay.protocolPay.isopen-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.trade.pay.protocolPay.isopen/{appKey}`  
Requires user authorization (access_token) · Requires signature

Check whether an auto-debit (withholding) agreement is enabled.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.ocean.openplatform.biz.trade.result.TradeWithholdStatusResultModel](#m-alibaba-ocean-openplatform-biz-trade-result-tradewithholdstatusresultmodel) | yes | Signing status return value, indicating the auto-debit signing status for Alipay and Cheng-e-She. payChannel=SHEGOU and signedStatus=true means Cheng-e-She auto-debit has been signed; payChannel=ALIPAY and signedStatus=true means Alipay auto-debit has been signed. If either signedStatus=true, the auto-debit interface can be called to complete the auto-debit. | {   "result": {     "result": {       "paymentAgreements": [         {           "bindingStatus": "false",           "payChannel": "SHEGOU",           "signedStatus": "false"         },         {           "bindingStatus": "true",           "payChannel": "ALIPAY",           "signUrl": "https://mapi.alipay.com/gateway.do?sales_product_code=TAOBAO_TRAVEL_HOTEL¬ify_url=http%3A%2F%2Fpre-center.finnet.alibaba.com%2Fcallback%2Fv1%2Ffm-PG.MID.CBU.NewRetail-alipay.tbapi.escrow-sign%2Fip_33_7_69_219-0%2F589540333&product_code=PERSONAL_TRAVEL_HOTEL&scene=INDUSTRY%7CCBU_AOTUPAY&partner=2088006300088887&service=alipay.dut.customer.agreement.page.sign&external_sign_no=119901210704561230970333&sign_type=DSA&sign=jm4udyi_d_a7_d46_yxsb_m_w3_xdb3_w_j9_wx_sg_u_w_g3d_d_rv_epkws_h_hid_e_h6i_cg%253D%253D",           "signedStatus": "false"         }       ]     },     "success": true   } } |

<a id="m-alibaba-ocean-openplatform-biz-trade-result-tradewithholdstatusresultmodel"></a>
#### alibaba.ocean.openplatform.biz.trade.result.TradeWithholdStatusResultModel

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `success` | Boolean | yes | Whether successful | true |
| `code` | String | yes | Error code | null |
| `message` | String | yes | Error message | null |
| `result` | [message:alibaba.ocean.openplatform.biz.trade.result.TradeWithholdStatusResult](#m-alibaba-ocean-openplatform-biz-trade-result-tradewithholdstatusresult) | yes | Contract signing status | {} |

<a id="m-alibaba-ocean-openplatform-biz-trade-result-tradewithholdstatusresult"></a>
#### alibaba.ocean.openplatform.biz.trade.result.TradeWithholdStatusResult

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `paymentAgreements` | [message:alibaba.ocean.openplatform.biz.trade.result.TradePaymentAgreement[]](#m-alibaba-ocean-openplatform-biz-trade-result-tradepaymentagreement[]) | yes |  |  |

<a id="m-alibaba-ocean-openplatform-biz-trade-result-tradepaymentagreement[]"></a>
#### alibaba.ocean.openplatform.biz.trade.result.TradePaymentAgreement[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `payChannel` | java.lang.String | yes | Channel for a successful payment; null if payment was not successful | ALIPAY |
| `bindingStatus` | java.lang.String | yes | Whether Alipay or Cheng-e-She has been bound/set up. Both signedStatus and bindingStatus must be true to initiate automatic debit. | true |
| `signedStatus` | java.lang.String | yes | Whether Alipay or Cheng-e-She has signed up for automatic debit. Both signedStatus and bindingStatus must be true to initiate automatic debit. | true |
| `signUrl` | java.lang.String | yes | Contract signing URL | http:// |
| `agreementNo` | java.lang.String | yes | Contract order number | 55622148344 |
