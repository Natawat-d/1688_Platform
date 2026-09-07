# 查询是否开通免密支付

API: `com.alibaba.trade:alibaba.trade.pay.protocolPay.isopen:1` · Category: 支付  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.trade.pay.protocolPay.isopen-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.trade.pay.protocolPay.isopen/{appKey}`  
需要授权 (access_token) · 需要签名

查询是否开通代扣协议

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:alibaba.ocean.openplatform.biz.trade.result.TradeWithholdStatusResultModel](#m-alibaba-ocean-openplatform-biz-trade-result-tradewithholdstatusresultmodel) | 是 | 签约情况返回值，支付宝和诚E赊的自动代扣情况.payChannel=SHEGOU且signedStatus=true表示诚E赊自动代扣已经签约，payChannel=ALIPAY且signedStatus=true表示支付宝代扣已签约。任何一个signedStatus=true都可以调用代扣接口，完成代扣。 | {   "result": {     "result": {       "paymentAgreements": [         {           "bindingStatus": "false",           "payChannel": "SHEGOU",           "signedStatus": "false"         },         {           "bindingStatus": "true",           "payChannel": "ALIPAY",           "signUrl": "https://mapi.alipay.com/gateway.do?sales_product_code=TAOBAO_TRAVEL_HOTEL¬ify_url=http%3A%2F%2Fpre-center.finnet.alibaba.com%2Fcallback%2Fv1%2Ffm-PG.MID.CBU.NewRetail-alipay.tbapi.escrow-sign%2Fip_33_7_69_219-0%2F589540333&product_code=PERSONAL_TRAVEL_HOTEL&scene=INDUSTRY%7CCBU_AOTUPAY&partner=2088006300088887&service=alipay.dut.customer.agreement.page.sign&external_sign_no=119901210704561230970333&sign_type=DSA&sign=jm4udyi_d_a7_d46_yxsb_m_w3_xdb3_w_j9_wx_sg_u_w_g3d_d_rv_epkws_h_hid_e_h6i_cg%253D%253D",           "signedStatus": "false"         }       ]     },     "success": true   } } |

<a id="m-alibaba-ocean-openplatform-biz-trade-result-tradewithholdstatusresultmodel"></a>
#### alibaba.ocean.openplatform.biz.trade.result.TradeWithholdStatusResultModel

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `success` | Boolean | 是 | 是否成功 | true |
| `code` | String | 是 | 错误码 | null |
| `message` | String | 是 | 错误消息 | null |
| `result` | [message:alibaba.ocean.openplatform.biz.trade.result.TradeWithholdStatusResult](#m-alibaba-ocean-openplatform-biz-trade-result-tradewithholdstatusresult) | 是 | 签约状态 | {} |

<a id="m-alibaba-ocean-openplatform-biz-trade-result-tradewithholdstatusresult"></a>
#### alibaba.ocean.openplatform.biz.trade.result.TradeWithholdStatusResult

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `paymentAgreements` | [message:alibaba.ocean.openplatform.biz.trade.result.TradePaymentAgreement[]](#m-alibaba-ocean-openplatform-biz-trade-result-tradepaymentagreement[]) | 是 |  |  |

<a id="m-alibaba-ocean-openplatform-biz-trade-result-tradepaymentagreement[]"></a>
#### alibaba.ocean.openplatform.biz.trade.result.TradePaymentAgreement[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `payChannel` | java.lang.String | 是 | 支付成功的渠道，支付不成功渠道为null | ALIPAY |
| `bindingStatus` | java.lang.String | 是 | 支付宝或者诚E赊是否已设置绑定，signedStatus和bindingStatus均为true才能发起代扣 | true |
| `signedStatus` | java.lang.String | 是 | 支付宝或者诚E赊是否已签约代扣，signedStatus和bindingStatus均为true才能发起代扣 | true |
| `signUrl` | java.lang.String | 是 | 签约URl | http:// |
| `agreementNo` | java.lang.String | 是 | 签约单号 | 55622148344 |
