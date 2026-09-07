# 查询交易单下关联的所有开具的发票信息（买家视角）

API: `com.alibaba.trade:trade.invoice.getListBuyerView:1` · Category: 发票  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:trade.invoice.getListBuyerView-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/trade.invoice.getListBuyerView/{appKey}`  
需要授权 (access_token) · 需要签名

查询交易单下关联的所有开具的发票信息（买家视角）

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `input` | [message:alibaba.ocean.openplatform.biz.trade.param.OpInvoiceByOrderIdQueryParam](#m-alibaba-ocean-openplatform-biz-trade-param-opinvoicebyorderidqueryparam) | 是 |  |   |

<a id="m-alibaba-ocean-openplatform-biz-trade-param-opinvoicebyorderidqueryparam"></a>
#### alibaba.ocean.openplatform.biz.trade.param.OpInvoiceByOrderIdQueryParam

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `orderId` | Long | 是 | 订单ID | 123 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:alibaba.ocean.openplatform.common.OpInvoiceModelListResultModel](#m-alibaba-ocean-openplatform-common-opinvoicemodellistresultmodel) | 是 |  |   |

<a id="m-alibaba-ocean-openplatform-common-opinvoicemodellistresultmodel"></a>
#### alibaba.ocean.openplatform.common.OpInvoiceModelListResultModel

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `code` | String | 是 |  |   |
| `message` | String | 是 |  |   |
| `result` | [message:alibaba.ocean.openplatform.biz.trade.common.model.OpInvoiceModel[]](#m-alibaba-ocean-openplatform-biz-trade-common-model-opinvoicemodel[]) | 是 |  |   |
| `retCodes` | String[] | 是 |  |   |
| `subCode` | String | 是 |  |   |
| `subMessage` | String | 是 |  |   |
| `success` | Boolean | 是 |  |   |

<a id="m-alibaba-ocean-openplatform-biz-trade-common-model-opinvoicemodel[]"></a>
#### alibaba.ocean.openplatform.biz.trade.common.model.OpInvoiceModel[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `amount` | java.lang.Long | 是 | 发票合计含税金额，单位分 | 2 |
| `bizStatus` | java.lang.String | 是 | 发票状态：ISSUING(1,&quot;开票中&quot;),ISSUED(2,&quot;已开票&quot;),VERIFYING(5,&quot;验票中&quot;),VERIFY_FAILED(6,&quot;验票失败&quot;),RETURNING(10,&quot;退票中&quot;),RETURNED(11,&quot;已退票&quot;),INVALIDING(20,&quot;作废中&quot;),DEPRECATED(21,&quot;已作废&quot;),RED_ISSUING(30,&quot;冲红中&quot;),RED_PART_ISSUED(31,&quot;部分冲红&quot;),RED_ALL_ISSUED(32,&quot;全部冲红&quot;),FAILED(40,&quot;开票失败&quot;),CLOSED(50,&quot;关闭&quot;),; | ISSUED |
| `gmtCreate` | java.util.Date | 是 | 创建时间 |   |
| `gmtModified` | java.util.Date | 是 | 修改时间 |   |
| `id` | java.lang.Long | 是 | 主键 | 123 |
| `invoiceDate` | java.lang.Long | 是 | 发票实际开票时间；区别于提交时间；发票申请时间等 |   |
| `invoiceDeadline` | java.lang.Long | 是 | 开票截止时间，毫秒时间戳 | 123 |
| `invoiceDownloadUrl` | java.lang.String | 是 | 发票下载链接 | 发票下载链接 |
| `invoiceErrMsg` | java.lang.String | 是 | 开票失败原因 | 开票失败原因 |
| `invoiceFileType` | java.lang.String | 是 | 发票文件类型：PDF(&quot;PDF&quot;,&quot;PDF格式&quot;),OFD(&quot;OFD&quot;,&quot;OFD格式&quot;),GIF(&quot;GIF&quot;,&quot;GIF格式&quot;),TIF(&quot;TIF&quot;,&quot;TIF格式&quot;),BMP(&quot;BMP&quot;,&quot;BMP格式&quot;),JPG(&quot;JPG&quot;,&quot;JPG格式&quot;),XML(&quot;XML&quot;,&quot;XML格式&quot;),PNG(&quot;PNG&quot;,&quot;PNG格式&quot;); | PDF |
| `invoiceMaterial` | java.lang.String | 是 | 发票材质：ELECTRON(1,&quot;电子&quot;); | ELECTRON |
| `invoiceNo` | java.lang.String | 是 | 发票号 | 发票号 |
| `invoiceType` | java.lang.String | 是 | 发票类型（包含特殊发票类型）：VATAX_COMM(1,&quot;增值税普通发票&quot;),VATAX_SPEC(2,&quot;增值税专用发票&quot;); | VATAX_COMM |
| `isRedInvoice` | boolean | 是 | 是否红票 | false |
| `ossInvoiceFileName` | java.lang.String | 是 | 文件名称 | 文件名称 |
| `outBizId` | java.lang.String | 是 | 业务单号：和上游交互使用，要保证唯一；因为交易单可能存在多张发票；创建发票的时候，一定要传入唯一的outBizId | 123 |
| `purchaserOpenUid` | String | 是 | 购买方openUid | POU |
| `purchaserInvoiceTitleModel` | [message:alibaba.ocean.openplatform.biz.trade.result.OpInvoiceTitleModel](#m-alibaba-ocean-openplatform-biz-trade-result-opinvoicetitlemodel) | 是 | 购方信息 | {} |
| `requestNo` | java.lang.String | 是 | 和下游交互的请求no，基于交易单申请的时候传入交易单 |   |
| `sellerInputInvoiceTime` | java.lang.Long | 是 | 商家录入发票时间 | 123 |
| `sellerInvoiceTitleModel` | [message:alibaba.ocean.openplatform.biz.trade.result.OpInvoiceTitleModel](#m-alibaba-ocean-openplatform-biz-trade-result-opinvoicetitlemodel) | 是 | 销售方信息 | {} |
| `sellerMemberId` | java.lang.String | 是 | 销售方会员id | 销售方会员id |
| `sellerOpenUid` | String | 是 | 销售方UID | SOU |
| `source` | java.lang.String | 是 | 发票来源 |   |

<a id="m-alibaba-ocean-openplatform-biz-trade-result-opinvoicetitlemodel"></a>
#### alibaba.ocean.openplatform.biz.trade.result.OpInvoiceTitleModel

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `titleType` | java.lang.String | 是 | 抬头类型，个人和社会组织都是PERSONAL：PERSONAL(0,&quot;个人和社会组织&quot;),COMPANY(1,&quot;企业&quot;); | PERSONAL |
| `title` | java.lang.String | 是 | 抬头 | 抬头 |
| `taxpayerIdentify` | java.lang.String | 是 | 纳税人识别号(企业开票必填) | 123 |
| `bankName` | java.lang.String | 是 | 开户行(企业开专票必填) | 开户行 |
| `bankAccountId` | java.lang.String | 是 | 银行账号(企业开专票必填) | 123 |
| `registerAddress` | java.lang.String | 是 | 企业注册地址(企业开专票必填) | 企业注册地址 |
| `registerPhone` | java.lang.String | 是 | 企业电话(企业开专票必填) | 123 |
| `OpenUid` | String | 是 | openUid |  OUD |
