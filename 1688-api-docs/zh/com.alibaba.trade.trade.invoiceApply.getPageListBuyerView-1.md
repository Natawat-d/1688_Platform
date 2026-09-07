# 分页查询发票申请列表（买家视角）

API: `com.alibaba.trade:trade.invoiceApply.getPageListBuyerView:1` · Category: 发票  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:trade.invoiceApply.getPageListBuyerView-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/trade.invoiceApply.getPageListBuyerView/{appKey}`  
需要授权 (access_token) · 需要签名

分页查询发票申请列表（买家视角）

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `input` | [message:alibaba.ocean.openplatform.biz.trade.param.OpInvoiceListByPageParam](#m-alibaba-ocean-openplatform-biz-trade-param-opinvoicelistbypageparam) | 是 |  |   |

<a id="m-alibaba-ocean-openplatform-biz-trade-param-opinvoicelistbypageparam"></a>
#### alibaba.ocean.openplatform.biz.trade.param.OpInvoiceListByPageParam

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `bizStatusList` | java.lang.String[] | 否 | 发票状态：ISSUING(1,&quot;开票中&quot;),ISSUED(2,&quot;已开票&quot;),VERIFYING(5,&quot;验票中&quot;),VERIFY_FAILED(6,&quot;验票失败&quot;),RETURNING(10,&quot;退票中&quot;),RETURNED(11,&quot;已退票&quot;),INVALIDING(20,&quot;作废中&quot;),DEPRECATED(21,&quot;已作废&quot;),RED_ISSUING(30,&quot;冲红中&quot;),RED_PART_ISSUED(31,&quot;部分冲红&quot;),RED_ALL_ISSUED(32,&quot;全部冲红&quot;),FAILED(40,&quot;开票失败&quot;),CLOSED(50,&quot;关闭&quot;),; | ISSUING |
| `fuzzyInvoiceTitle` | java.lang.String | 否 | 模糊发票抬头 | 模糊发票抬头 |
| `isRedInvoice` | Boolean | 否 | 是否红票 | false |
| `orderId` | java.lang.String | 否 | 交易单id | 123 |
| `outBizId` | java.lang.String | 否 | 发票记录的唯一Id(查询申请记录时使用)；通常是交易单id，但是红票会变 | 123 |
| `page` | int | 是 | 当前页，能小于1 | 1 |
| `pageSize` | int | 是 | 分页大小，不能大于100，不能小于1 | 10 |
| `createMillTimeStart` | Long | 否 | 创建开始毫秒时间 | 1773849600000 |
| `createMillTimeEnd` | Long | 否 | 创建结束毫秒时间 | 1773936000000 |
| `modifyMillTimeStart` | Long | 否 | 修改开始毫秒时间 | 1773849600000 |
| `modifyMillTimeEnd` | Long | 否 | 修改结束毫秒时间 | 1773936000000 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:alibaba.ocean.openplatform.common.OpInvoiceModelPageResult](#m-alibaba-ocean-openplatform-common-opinvoicemodelpageresult) | 是 |  |   |

<a id="m-alibaba-ocean-openplatform-common-opinvoicemodelpageresult"></a>
#### alibaba.ocean.openplatform.common.OpInvoiceModelPageResult

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `errorCode` | java.lang.String | 是 |  |   |
| `errorInfo` | java.lang.String | 是 |  |   |
| `pageIndex` | int | 是 |  |   |
| `resultList` | [message:alibaba.ocean.openplatform.biz.trade.common.model.OpInvoiceModel[]](#m-alibaba-ocean-openplatform-biz-trade-common-model-opinvoicemodel[]) | 是 |  |   |
| `sizePerPage` | int | 是 |  |   |
| `success` | boolean | 是 |  |   |
| `totalRecords` | int | 是 |  |   |

<a id="m-alibaba-ocean-openplatform-biz-trade-common-model-opinvoicemodel[]"></a>
#### alibaba.ocean.openplatform.biz.trade.common.model.OpInvoiceModel[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `amount` | java.lang.Long | 是 | 发票合计含税金额，单位分 | 2 |
| `bizStatus` | java.lang.String | 是 | 发票状态：ISSUING(1,&quot;开票中&quot;),ISSUED(2,&quot;已开票&quot;),VERIFYING(5,&quot;验票中&quot;),VERIFY_FAILED(6,&quot;验票失败&quot;),RETURNING(10,&quot;退票中&quot;),RETURNED(11,&quot;已退票&quot;),INVALIDING(20,&quot;作废中&quot;),DEPRECATED(21,&quot;已作废&quot;),RED_ISSUING(30,&quot;冲红中&quot;),RED_PART_ISSUED(31,&quot;部分冲红&quot;),RED_ALL_ISSUED(32,&quot;全部冲红&quot;),FAILED(40,&quot;开票失败&quot;),CLOSED(50,&quot;关闭&quot;),; | ISSUED |
| `gmtCreate` | java.util.Date | 是 | 创建时间 |   |
| `gmtModified` | java.util.Date | 是 | 修改时间 |   |
| `id` | java.lang.Long | 是 | 主键 | 123 |
| `invoiceDeadline` | java.lang.Long | 是 | 开票截止时间，毫秒时间戳 | 123 |
| `invoiceErrMsg` | java.lang.String | 是 | 开票失败原因 | 开票失败原因 |
| `invoiceMaterial` | java.lang.String | 是 | 发票材质：ELECTRON(1,&quot;电子&quot;); | ELECTRON |
| `invoiceType` | java.lang.String | 是 | 发票类型（包含特殊发票类型）：VATAX_COMM(1,&quot;增值税普通发票&quot;),VATAX_SPEC(2,&quot;增值税专用发票&quot;); | VATAX_COMM |
| `isRedInvoice` | boolean | 是 | 是否红票 | false |
| `outBizId` | java.lang.String | 是 | 业务单号：和上游交互使用，要保证唯一；因为交易单可能存在多张发票；创建发票的时候，一定要传入唯一的outBizId | 123 |
| `purchaserOpenUid` | String | 是 | 购买方openUid | POU |
| `purchaserInvoiceTitleModel` | [message:alibaba.ocean.openplatform.biz.trade.result.OpInvoiceTitleModel](#m-alibaba-ocean-openplatform-biz-trade-result-opinvoicetitlemodel) | 是 | 购方信息 | {} |
| `requestNo` | java.lang.String | 是 | 和下游交互的请求no，基于交易单申请的时候传入交易单 |   |
| `sellerInputInvoiceTime` | java.lang.Long | 是 | 商家录入发票时间 | 123 |
| `sellerInvoiceTitleModel` | [message:alibaba.ocean.openplatform.biz.trade.result.OpInvoiceTitleModel](#m-alibaba-ocean-openplatform-biz-trade-result-opinvoicetitlemodel) | 是 | 销售方信息 | {} |
| `sellerMemberId` | java.lang.String | 是 | 销售方会员id | 销售方会员id |
| `sellerOpenUid` | String | 是 | 销售方UID | SOU |
| `source` | java.lang.String | 是 | 发票来源 |   |
| `invoiceDetailDTOList` | [message:alibaba.ocean.openplatform.biz.trade.common.model.OpInvoiceDetailModel[]](#m-alibaba-ocean-openplatform-biz-trade-common-model-opinvoicedetailmodel[]) | 是 | 开票商品详情列表 | [] |

<a id="m-alibaba-ocean-openplatform-biz-trade-result-opinvoicetitlemodel"></a>
#### alibaba.ocean.openplatform.biz.trade.result.OpInvoiceTitleModel

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `bankAccountId` | java.lang.String | 是 | 银行账号 |  银行账号 |
| `bankName` | java.lang.String | 是 | 开户行 |  开户行 |
| `registerAddress` | java.lang.String | 是 | 企业注册地址 |  企业注册地址 |
| `registerPhone` | java.lang.String | 是 | 企业电话 |  企业电话 |
| `taxpayerIdentify` | java.lang.String | 是 | 纳税人识别号 |  纳税人识别号 |
| `title` | java.lang.String | 是 | 抬头 |  抬头 |
| `titleType` | java.lang.String | 是 | 发票抬头类型，个人和社会组织都是PERSONAL：PERSONAL(0,&quot;个人和社会组织&quot;),COMPANY(1,&quot;企业&quot;); |  PERSONAL |
| `OpenUid` | String | 是 | openUid |  OUD |

<a id="m-alibaba-ocean-openplatform-biz-trade-common-model-opinvoicedetailmodel[]"></a>
#### alibaba.ocean.openplatform.biz.trade.common.model.OpInvoiceDetailModel[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `subOutBizId` | java.lang.String | 是 | 业务子单号 | 业务子单号 |
| `cargoName` | java.lang.String | 是 | 货物/服务名称 | 货物/服务名称 |
| `skuId` | String | 是 | skuId | skuId |
| `specifications` | java.lang.String | 是 | 规格型号 | 规格型号 |
| `specId` | String | 是 | 规格specId | 规格specId |
| `invoiceAmountWithTax` | java.lang.Long | 是 | 含税金额 | 123 |
| `quantity` | java.lang.Long | 是 | 数量 | 1 |
| `unit` | java.lang.String | 是 | 单位 | 单位 |
| `currency` | java.lang.String | 是 | 币种 | 币种 |
| `taxRate` | java.lang.Integer | 是 | 税率(0-100)整数 | 1 |
| `purchaserOpenUid` | String | 是 | 买方openUid | POU |
| `sellerOpenUid` | String | 是 | 卖方openUid | SOU |
| `relatedInvoiceId` | java.lang.Long | 是 | 关联的主发票单 | 123 |
