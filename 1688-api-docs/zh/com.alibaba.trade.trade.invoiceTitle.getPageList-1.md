# 分页查询买家抬头列表

API: `com.alibaba.trade:trade.invoiceTitle.getPageList:1` · Category: 发票  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:trade.invoiceTitle.getPageList-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/trade.invoiceTitle.getPageList/{appKey}`  
需要授权 (access_token) · 需要签名

分页查询买家抬头列表

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `input` | [message:alibaba.ocean.openplatform.biz.trade.param.OpInvoiceTitleQueryParam](#m-alibaba-ocean-openplatform-biz-trade-param-opinvoicetitlequeryparam) | 是 |  |   |

<a id="m-alibaba-ocean-openplatform-biz-trade-param-opinvoicetitlequeryparam"></a>
#### alibaba.ocean.openplatform.biz.trade.param.OpInvoiceTitleQueryParam

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `page` | int | 否 | 当前页，不能小于1 | 1 |
| `pageSize` | int | 否 | 分页大小，不能小于1，不能大于100 | 50 |
| `titleType` | java.lang.String | 否 | 发票抬头类型，个人和社会组织都是PERSONAL：PERSONAL(0,&quot;个人&quot;),COMPANY(1,&quot;企业&quot;),SOCIAL_ORGANIZATION(3,&quot;社会组织（机关事业单位等）&quot;),PARENT_VIRTUAL(-1,&quot;父票虚拟类型&quot;); | PERSONAL |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:alibaba.ocean.openplatform.common.OpInvoiceTitlePageSingleResult](#m-alibaba-ocean-openplatform-common-opinvoicetitlepagesingleresult) | 是 |  |   |

<a id="m-alibaba-ocean-openplatform-common-opinvoicetitlepagesingleresult"></a>
#### alibaba.ocean.openplatform.common.OpInvoiceTitlePageSingleResult

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `currentPage` | java.lang.Integer | 是 | 当前页码 | 1 |
| `errorCode` | java.lang.String | 是 | 错误编码 | 404 |
| `errorInfo` | java.lang.String | 是 | 错误信息 | 错误信息 |
| `pageSize` | java.lang.Integer | 是 | 分页大小 | 10 |
| `result` | [message:alibaba.ocean.openplatform.biz.trade.result.OpInvoiceTitleQueryResult](#m-alibaba-ocean-openplatform-biz-trade-result-opinvoicetitlequeryresult) | 是 | 返回结果 | {} |
| `success` | boolean | 是 | 是否成功 | true |
| `totalNum` | java.lang.Long | 是 | 总记录数 | 100 |

<a id="m-alibaba-ocean-openplatform-biz-trade-result-opinvoicetitlequeryresult"></a>
#### alibaba.ocean.openplatform.biz.trade.result.OpInvoiceTitleQueryResult

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `companyInvoiceTitles` | [message:alibaba.ocean.openplatform.biz.trade.result.OpInvoiceTitleModel[]](#m-alibaba-ocean-openplatform-biz-trade-result-opinvoicetitlemodel[]) | 是 | 企业抬头 | [] |
| `personalInvoiceTitles` | [message:alibaba.ocean.openplatform.biz.trade.result.OpInvoiceTitleModel[]](#m-alibaba-ocean-openplatform-biz-trade-result-opinvoicetitlemodel[]) | 是 | 个人抬头 | [] |

<a id="m-alibaba-ocean-openplatform-biz-trade-result-opinvoicetitlemodel[]"></a>
#### alibaba.ocean.openplatform.biz.trade.result.OpInvoiceTitleModel[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `bankAccountId` | java.lang.String | 是 | 银行账号 | 123 |
| `bankName` | java.lang.String | 是 | 开户行 | 开户行 |
| `email` | java.lang.String | 是 | 邮箱 | 邮箱 |
| `gmtCreate` | java.util.Date | 是 | 创建时间 |   |
| `gmtModified` | java.util.Date | 是 | 修改时间 |   |
| `invoiceTitleId` | java.lang.Long | 是 | 抬头主键 id | 123 |
| `isDefault` | java.lang.String | 是 | 是否默认 | true |
| `name` | java.lang.String | 是 | 企业联系人 | 企业联系人 |
| `receiverPhone` | java.lang.String | 是 | 收件人电话 | 123 |
| `registerAddress` | java.lang.String | 是 | 企业注册地址 | 企业注册地址 |
| `registerPhone` | java.lang.String | 是 | 企业电话 | 123 |
| `taxpayerIdentify` | java.lang.String | 是 | 纳税人识别号 | 纳税人识别号 |
| `title` | java.lang.String | 是 | 抬头 | 抬头 |
| `titleType` | java.lang.String | 是 | 发票抬头类型，个人和社会组织都是PERSONAL：PERSONAL(0,&quot;个人&quot;),COMPANY(1,&quot;企业&quot;),SOCIAL_ORGANIZATION(3,&quot;社会组织（机关事业单位等）&quot;),PARENT_VIRTUAL(-1,&quot;父票虚拟类型&quot;); | PERSONAL |
| `openUid` | String | 是 | openUid | OU |
