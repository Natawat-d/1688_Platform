# 1688 XunYuanTong solution: all APIs, tokens and request bodies

Solution 1703167397970 · pulled from open.1688.com 2026-09-03 · 142 APIs (one listed API has no detail doc on 1688).

Every call is an HTTP POST (form-encoded) to `https://gw.open.1688.com/openapi/param2/{version}/{namespace}/{name}/{appKey}`. System parameters go alongside the body: `access_token` (user authorization token, from the OAuth flow), `_aop_signature` (HMAC-SHA1 request signature), `_aop_timestamp` (optional). Fields whose type is an object are passed as JSON strings; nested fields are shown as dotted paths (`parent.child`, `list[].child`).

## Contents

- **Membership & Accounts** (会员): [Batch-add sub-account authorizations](#subaccountauthadd-1), [Batch-cancel sub-account authorizations](#subaccountauthcancel-1), [Batch-query sub-account authorizations](#subaccountauthlist-1), [1688 member registration](#accountuserregister-1), [Get basic info of a non-authorized user (cross-border)](#alibabaaccountagentcrossbasic-1), [Get basic info of the authorized user](#alibabaaccountbasic-1), [Query sub-account info](#querysubaccount-1)
- **Tools** (工具): [Pull products from a product pool](#poolproductpull-1), [Query product count in a product pool](#poolproducttotal-1), [Encrypt a user loginId into an OpenUID](#loginidopenuidencrypt-1), [Get a link that opens a WangWang chat](#accountwangwangurlget-1), [Decrypt an OpenUID into a WangWang nickname (only for launching WangWang)](#wangwangnickopenuiddecrypt-1), [Add buyer-seller distribution relationship](#alibabafenxiaorelationadd-1)
- **Market Insights** (商机): [Trending product search keywords](#productsearchtopkeyword-1), [Query ranking lists](#producttoplistquery-1), [Get daily sales-quantity trend for a product (new)](#productanalyzegetperdaysellquantitytrendnew-1), [Get 30-day median-price trend for a product (new)](#productanalyzegetthirtydaymedianpricetrendnew-1), [Get 90-day repurchase-rate trend for a product (new)](#productanalyzegetrepurchaseratetrendnew-1)
- **Categories** (类目): [Query multilingual category by category ID](#categorytranslationgetbyid-1), [Query multilingual category by category name](#categorytranslationgetbykeyword-1), [Get leaf-category attributes](#alibabacategoryattributeget-1)
- **Products** (商品): [Multilingual keyword search](#productsearchkeywordquery-1), [Multilingual image search](#productsearchimagequery-1), [Multilingual in-store product search](#productsearchquerysellerofferlist-1), [Multilingual product detail](#productsearchqueryproductdetail-1), [Product recommendations](#productsearchofferrecommend-1), [Multilingual search navigation](#productsearchkeywordsnquery-1), [Related product recommendations](#productrelatedrecommend-1), [Upload an image to get an imageId](#productimageupload-1), [Claim the optimal coupon for a product](#couponoptimalclaim-1), [Recommend products by keyword (cross-border)](#alibabaproductsuggestcrossborder-1), [Get simple product info from a previously purchased supplier](#alibabaproductsimpleget-1), [Follow a product (cross-border)](#productkjdistributeaddrelation-1), [Unfollow a product (cross-border)](#productkjdistributeremoverelation-1), [Get product selling points for listing](#productdistributegetdistributeinfo-1), [Query leaf-category attribute and value mappings](#productcategorygetattrbyid-1)
- **Product Selection & Listing** (选品铺货): [Get product details (cross-border)](#alibabacrossproductinfo-1), [Add products to the listing list (cross-border)](#alibabacrosssyncproductlistpushed-1), [Sync listing results](#alibabaproductpushsyncpushproductresult-1), [Get product list (cross-border)](#alibabacrossproductlist-1), [Unfollow a product](#alibabaproductunfollowcrossborder-1), [Query category by category ID](#alibabacategoryget-1), [Follow a product](#alibabaproductfollowcrossborder-1), [[Relationship] Distributor: query supplier list](#alibabarelationquerysuppliers-1)
- **Orders** (订单): [Create cross-border order](#alibabatradecreatecrossorder-1), [View order details (buyer view)](#alibabatradegetbuyerview-1), [Get order logistics info (buyer view)](#alibabatradegetlogisticsinfosbuyerview-1), [Get order logistics tracking info (buyer view)](#alibabatradegetlogisticstraceinfobuyerview-1), [Cancel transaction](#alibabatradecancel-1), [View order list (buyer view)](#alibabatradegetbuyerorderlist-1), [Preview data before creating an order](#alibabacreateorderpreview-1), [Update order memo](#alibabaordermemoadd-1), [Get sub-account list](#alibabasubaccountlist-1), [Buyer adds an order message](#alibabatradeaddfeedback-1), [Buyer confirms receipt](#tradereceivegoodsconfirm-1), [Buyer deletes a closed order](#tradeorderbuyerdelete-1), [Buyer requests a shipping-address change](#orderreceiveaddressbuyerupdate-1)
- **Payment** (支付): [Batch-get payment links for orders](#alibabaalipayurlget-1), [Get payment link for Cross-Border Pay (Kuajingbao)](#alibabacrossborderpayurlget-1), [Get payment link for Cheng-e-She credit pay](#alibabacreditpayurlget-1), [Buyer views all granted credit terms](#alibabaaccountperiodlistbuyerview-1), [Query payment channels supported by an order](#alibabatradepaywayquery-1), [Check whether password-free payment is enabled](#alibabatradepayprotocolpayisopen-1), [Initiate password-free payment](#alibabatradepayprotocolpaypreparepay-1), [Order payment consultation](#tradeorderpayanalysis-1), [Get combined-cashier URL](#alibabatradegrouppayurlget-1)
- **Logistics** (物流): [Estimate domestic (China) shipping fee for a product](#productfreightestimate-1), [Query external order ID by waybill number or unclaimed-parcel code](#logisticsordergetoutorderid-1), [Query shipping-insurance info](#shippinginsuranceget-1), [Get shipping-template details](#alibabalogisticsmyfreighttemplatelistget-1), [Buyer gets saved shipping addresses](#alibabatradereceiveaddressget-1), [Parse an address into area codes](#alibabatradeaddresscodeparse-1), [Query seller mixed-batch settings](#alibabatradeopquerymarketingmixconfig-1), [Logistics company list (all companies)](#alibabalogisticsopquerylogisticcompanylist-1), [Parse overseas address](#tradeaddressparsetext-1), [Urge seller to ship](#logisticsdeliveryurge-1), [Receive external print-label (UDF) URL](#labelurlreceive-1), [Get waybill-number set](#bigcustomermailnoquery-1)
- **Official Return Pickup** (官方退上门取件): [Create official return-pickup logistics order](#refundofficialdeliveryordercreate-1), [Modify official return-pickup logistics order](#refundofficialdeliveryordermodify-1), [Get official return-pickup logistics order details](#refundofficialdeliveryorderget-1), [Cancel official return-pickup logistics order](#refundofficialdeliveryordercancel-1), [Query official return-pickup plans](#refundofficialdeliverysolutionget-1)
- **Data Write-back** (回传数据): [Write back mapping between end-customer orders and 1688 orders](#orderrelationwrite-1), [Write back country-site logistics orders](#tradecrosslogisticsordersync-1), [Sync downstream sales orders](#tradecrossordersync-1), [Save the business line an account belongs to](#accountbusinesssave-1)
- **Returns & Refunds** (退货退款): [Query refund details by order ID (buyer view)](#alibabatraderefundopquerybatchrefundbyorderidandstatus-1), [Refund operation history (buyer view)](#alibabatraderefundopqueryorderrefundoperationlist-1), [Query refund details by refund ID (buyer view)](#alibabatraderefundopqueryorderrefund-1), [Query refund list (buyer view)](#alibabatraderefundbuyerqueryorderrefundlist-1), [Query order number by claim or insurance-policy number](#alibabatradequeryorderbyinsure-1), [Create refund/return request](#alibabatradecreaterefund-1), [Query refund/return reasons (for creating a request)](#alibabatradegetrefundreasonlist-1), [Upload refund/return evidence](#alibabatradeuploadrefundvoucher-1), [Buyer submits return-shipment info](#alibabatraderefundreturngoods-1), [Cancel refund/return request](#alibabatradecancelrefund-1), [Query maximum refundable amount when applying](#alibabatradegetmaxrefundfee-1), [Apply for trade arbitration](#tradearbitrationapply-1), [Query the merchant's return address](#refundaddressget-1)
- **Messaging** (消息): [Batch-confirm failed messages](#pushmessageconfirm-1), [Fetch failed messages (query style)](#pushquerymessagelist-1), [Fetch failed messages (cursor style)](#pushcursormessagelist-1), [View order list (seller view)](#alibabatradegetsellerorderlist-1), [Ship: seller arranges own logistics](#alibabalogisticsopdeliverysendorderoffline-1), [Ship: no logistics needed](#alibabalogisticsopdeliverysendorderdummy-1)
- **Light Customization** (轻定制): [Get artwork info for a customization order](#productpodorderdesignget-1)
- **Suppliers** (商家): [Run AI supplier-search task (async)](#accountsearchgetinfoasync-1), [Query AI supplier-search task result](#accountsearchgetinforesult-1)
- **Inquiries (Newton Cloud)** (询盘): [Newton Cloud: create long-running task](#newtoncloudtaskcreate-1), [Newton Cloud: query task](#newtoncloudtaskget-1), [Newton Cloud: list tasks](#newtoncloudtasklist-1), [Newton Cloud: terminate task](#newtoncloudtaskkill-1), [Newton Cloud: query task table](#newtoncloudtaskfetch-1), [Newton Cloud: resume task](#newtoncloudtaskresume-1), [Newton Cloud: query batch-inquiry results](#newtoncloudbatchinquirygetresult-1), [Newton Cloud: upload file](#newtoncloudfileupload-1), [Newton Cloud: list available model tiers](#newtoncloudmodellist-1), [Newton Cloud: query points details](#newtoncloudpointsquery-1)
- **88 ShengYiTong (Business Link)** (88生意通): [88 ShengYiTong: buyer drafts purchase order](#sytbuyerdraftpurchaseorder-2), [88 ShengYiTong: confirm transaction complete](#sytcontractconfirm-1), [88 ShengYiTong: void contract](#sytcontractinvalid-1), [88 ShengYiTong: contract refund request](#sytcontractrefund-1), [88 ShengYiTong: buyer payment, get cashier URL](#sytcontractpay-1), [88 ShengYiTong: WorldFirst transfer payment (WorldFirst B2C accounts only)](#sytcontractpaytransfer-1), [88 ShengYiTong: buyer confirms purchase order](#sytbuyerconfirmpurchaseorder-1), [88 ShengYiTong: buyer queries purchase-order details](#sytbuyerquerycontractdetail-1), [88 ShengYiTong: check whether customer is verified](#sytcustomerqueryuserauthstatus-1)
- **Invoicing** (发票): [Query buyer invoice titles (paginated)](#tradeinvoicetitlegetpagelist-1), [Query invoiceable amount of orders](#tradeinvoiceamountgetlist-1), [Query invoice applications (buyer view, paginated)](#tradeinvoiceapplygetpagelistbuyerview-1), [Query all issued invoices for an order (buyer view)](#tradeinvoicegetlistbuyerview-1), [Buyer requests invoice](#tradeinvoiceapply-1), [Merged-invoice consultation](#tradeinvoiceconsult-1), [Submit merged-invoice request](#tradeinvoicemergeapply-1), [Query merged-invoice relationships by order or invoice application](#tradeinvoicesellerqueryrelatedorders-1)
- **Fully Managed (Consignment)** (全托管): [Warehouse receipt / shelving of goods](#consignmentcostatussync-1), [Warehouse creates discrepancy tally sheet](#consignmenttallycreate-1), [Update warehouse product inventory](#consignmentcoinventory-1)
- **Listing (Publishing)** (铺货): [Publish result callback](#publishresultcallback-1), [Listable product list](#publishproductlist-1), [Listing product card](#publishcardget-1)
- **Repurchase** (复购): [Query repurchase contract](#repurchasecontractget-1)

<a id="subaccountauthadd-1"></a>
## 1. Batch-add sub-account authorizations

`system.oauth2:subaccount.auth.add:1` · Membership & Accounts · 批量添加子账号授权  
POST `https://gw.open.1688.com/openapi/param2/1/system.oauth2/subaccount.auth.add/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Batch-add authorization for the sub-accounts under a main account. The main account must already be authorized.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `subUserIdentityList` | java.lang.String[] | yes | List of sub-account ids | abb:test1,abb:test2 |

**Response (top level)**: `result` AuthResultModelMap – Return result object; `result.errorCode` java.lang.String – Error code; `result.errorMessage` java.lang.String – Error message; `result.returnValue` java.util.Map – Return result; `result.success` boolean – Whether successful

<a id="subaccountauthcancel-1"></a>
## 2. Batch-cancel sub-account authorizations

`system.oauth2:subaccount.auth.cancel:1` · Membership & Accounts · 批量取消子账号授权  
POST `https://gw.open.1688.com/openapi/param2/1/system.oauth2/subaccount.auth.cancel/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (optional) · user authorization not required · signature required  
Cancel the authorization of sub-accounts in batch.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `subUserIdentityList` | java.lang.String[] | yes | List of sub-account ids | a:b,a:c |

**Response (top level)**: `result` AuthResultModelBoolean – Return result; `result.errorCode` java.lang.String – Error code; `result.errorMessage` java.lang.String – Error description; `result.returnValue` java.lang.Boolean – Return result; `result.success` boolean – Whether successful

<a id="subaccountauthlist-1"></a>
## 3. Batch-query sub-account authorizations

`system.oauth2:subaccount.auth.list:1` · Membership & Accounts · 批量查询子账号授权  
POST `https://gw.open.1688.com/openapi/param2/1/system.oauth2/subaccount.auth.list/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Batch-query the authorization status of the sub-accounts under a main account.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `subUserIdentityList` | java.lang.String[] | yes | List of sub-account ids | "a:b","a:c" |

**Response (top level)**: `result` AuthResultModelRelation – Query result; `result.errorCode` java.lang.String – Error code; `result.errorMessage` java.lang.String – Error description; `result.returnValue` AuthRelationDTO[] – Return result; `result.success` boolean – Whether successful; …

<a id="accountuserregister-1"></a>
## 4. 1688 member registration

`com.alibaba.fenxiao.crossborder:account.user.register:1` · Membership & Accounts · 1688会员注册  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/account.user.register/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Register a 1688 member.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `countryAccount` | CountryAccount | yes | Account registration input parameters | {     "country": "japan",     "site": "sniff",     "outLoginId": "18899993333",     "outMe… |
| `countryAccount.country` | java.lang.String | yes | Country; for supported parameters see the country enum. | japan |
| `countryAccount.site` | java.lang.String | yes | Site; please pass in the company name or e-commerce product name. Must contain no spaces and be lowercase. No more than 16 characters. | sniff |
| `countryAccount.outLoginId` | java.lang.String | yes | The user's login name within the organization, no more than 60 characters. | 1234test |
| `countryAccount.outMemberId` | java.lang.String | yes | The user's unique identifier within the organization, no more than 60 characters. | 1232fdsf |
| `countryAccount.email` | java.lang.String | yes | Email; format will be validated, no more than 60 characters. | 123@163.com |
| `countryAccount.mobile` | java.lang.String | yes | Mobile number; format will be validated, no more than 30 characters. | 1234567890 |
| `countryAccount.mobileArea` | java.lang.String | yes | The region the mobile number belongs to; see the mobileArea enum for supported parameters. | JP |
| `countryAccount.ip` | java.lang.String | yes | IP address; format will be validated. | 11.11.11.11 |

**Response (top level)**: `result` ResultModel – ; `result.success` boolean – Whether the call succeeded, true for success, false for failure.; `result.code` java.lang.String – Error code, e.g. S0000 indicates success.; `result.message` java.lang.String – Error message, e.g. success.; `result.result` java.lang.Boolean – Return result.

<a id="alibabaaccountagentcrossbasic-1"></a>
## 5. Get basic info of a non-authorized user (cross-border)

`com.alibaba.account:alibaba.account.agent.crossBasic:1` · Membership & Accounts · 跨境场景获取非授权用户的基本信息  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.account/alibaba.account.agent.crossBasic/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
View another user's basic information. For use in cross-border scenarios.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `loginId` | String | no | The user's loginId; the input parameters cannot both be empty | alitestforisv01 |
| `domain` | String | no | WangPu (storefront) domain name; input parameters cannot all be empty at the same time | trgm66666.1688.com |

**Response (top level)**: `result` simpleAccountInfo – Member information; `result.loginId` java.lang.String – Login name; `result.categoryName` java.lang.String – Category name of the main industry; `result.companyName` java.lang.String – Company name; `result.shopUrl` java.lang.String – WangPu (shop) homepage URL; `result.supplierName` java.lang.String – Supplier name; `result.kuaJingBao` Boolean – Whether it is Kuajingbao (Cross-Border Pay); `errorCode` String – Error code; `errorMessage` String – Error message; `success` Boolean – Whether successful

<a id="alibabaaccountbasic-1"></a>
## 6. Get basic info of the authorized user

`com.alibaba.account:alibaba.account.basic:1` · Membership & Accounts · 获取授权用户的基本信息  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.account/alibaba.account.basic/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Get the basic information of the authorized user.

**Request body**: no application parameters.

**Response (top level)**: `result` simpleAccountInfo – Member information; `result.loginId` java.lang.String – Login name; `result.saleRate` java.lang.Long – Seller level; `result.maturity` int – Member maturity level; `result.memo` java.lang.String – Remark; `result.modifyDate` java.util.Date – Modification date; `result.categoryName` java.lang.String – Category name of the main industry; `result.trustScore` int – Chengxintong (TrustPass) index; `result.userId` java.lang.Long – User ID; `result.enterpriseAccount` boolean – Whether it is a company member; `result.createDate` java.util.Date – Creation time; `result.communityLevel` java.lang.String – TradeManager (Maoyitong) user identifier; `result.joinFrom` java.lang.String – Registration source; `result.rateNum` java.lang.Long – Buyer's credit score (Taobao); …

<a id="querysubaccount-1"></a>
## 7. Query sub-account info

`cn.alibaba.open:querySubAccount:1` · Membership & Accounts · 查询子账号信息  
POST `https://gw.open.1688.com/openapi/param2/1/cn.alibaba.open/querySubAccount/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (optional) · user authorization not required · signature required  
Query sub-account information.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `loginId` | String | yes | loginid | bonlientest:zhagnshan |

**Response (top level)**: `result` ResultModel – Return model; `result.errorMsg` java.lang.String – Error message; `result.resultCode` java.lang.String – Result code; `result.isSuccess` boolean – Success flag; `result.data` CoopSubAccountModel – Data; …

<a id="poolproductpull-1"></a>
## 8. Pull products from a product pool

`com.alibaba.fenxiao.crossborder:pool.product.pull:1` · Tools · 拉取商品池中商品数据  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/pool.product.pull/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Batch-pull product data directly from a product pool by pool ID.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `offerPoolQueryParam` | OfferPoolQueryParam | yes |  |  |
| `offerPoolQueryParam.offerPoolId` | java.lang.Long | yes | Product pool ID (business-customized and permission-controlled; obtain it from the integrated business - passing an arbitrary value will cause an error. For Xunyuantong purchasing-on-behalf, using the keyword search interface is recommended) | 111 |
| `offerPoolQueryParam.cateId` | java.lang.Long | no | Category ID | 11 |
| `offerPoolQueryParam.taskId` | java.lang.String | yes | Query task ID. For example, if an assortment has 10,000 products, querying 1,000 per page for 10 queries retrieves all 10,000 products; the same taskId must be passed for all 10 queries. The organization needs to fix and retain one taskId during paginated queries, then pass the same taskId each time when paging through this API. | 1 |
| `offerPoolQueryParam.language` | java.lang.String | no | Language | en |
| `offerPoolQueryParam.pageNo` | java.lang.Integer | yes | Page number | 1 |
| `offerPoolQueryParam.pageSize` | java.lang.Integer | yes | Number per page | 10 |
| `offerPoolQueryParam.sortField` | String | no | Sort field | order1m/buyer1m（order1m：最近1个月销售额排序；buyer1m：最近1个月买家数） |
| `offerPoolQueryParam.sortType` | String | no | Sorting rule | ASC/DESC |

**Response (top level)**: `result` ResultModel – ; `result.success` java.lang.String – Whether successful; `result.code` java.lang.String – Error code; `result.message` java.lang.String – Error description; `result.result` ProductPoolModel[] – Result; …

<a id="poolproducttotal-1"></a>
## 9. Query product count in a product pool

`com.alibaba.fenxiao.crossborder:pool.product.total:1` · Tools · 查询商品池中商品总数  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/pool.product.total/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (optional) · user authorization not required · signature required  
Query the total number of products in a product pool.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `appKey` | java.lang.String | yes | appKey | 123 |
| `palletId` | java.lang.Long | yes | palletId | 123 |
| `categoryId` | java.lang.String | no | categoryId | 1 |

**Response (top level)**: `result` ResultModel – Result; `result.msg` String – msg; `result.code` String – code; `result.traceId` String – traceId; `result.success` String – Flag indicating whether the request was successful; `result.model` Long – Total count

<a id="loginidopenuidencrypt-1"></a>
## 10. Encrypt a user loginId into an OpenUID

`com.alibaba.account:loginid.openuid.encrypt:1` · Tools · 用户loginId加密转换为Openuid接口  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.account/loginid.openuid.encrypt/{appKey}`  
Token/system params: `_aop_timestamp` (required), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Encrypts a user's loginId into an OpenUID. This interface is risk-controlled: batch operations are not allowed, and it may only be triggered manually by the merchant, for example when searching a user's orders or configuring rules.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `loginId` | String | yes | User login name |  |

**Response (top level)**: `openUid` String – openUid

<a id="accountwangwangurlget-1"></a>
## 11. Get a link that opens a WangWang chat

`com.alibaba.account:account.wangwangUrl.get:1` · Tools · 获取唤起旺旺聊天的链接  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.account/account.wangwangUrl.get/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Get a link that launches an AliWangWang chat.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `toOpenUid` | String | yes | Chat counterpart's openUid |  |

**Response (top level)**: `result` ResultModelGTlfvn7n – ; `result.code` java.lang.String – ; `result.message` java.lang.String – ; `result.result` java.lang.String – WangWang chat link; `result.retCodes` java.lang.String[] – ; `result.subCode` java.lang.String – ; `result.subMessage` java.lang.String – ; `result.success` boolean – 

<a id="wangwangnickopenuiddecrypt-1"></a>
## 12. Decrypt an OpenUID into a WangWang nickname (only for launching WangWang)

`com.alibaba.account:wangwangnick.openuid.decrypt:1` · Tools · Openuid转换解密为旺旺昵称接口（仅可使用于用户唤起旺旺）  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.account/wangwangnick.openuid.decrypt/{appKey}`  
Token/system params: `_aop_timestamp` (required), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Decrypts an OpenUID into a WangWang nickname. This interface is risk-controlled: it may only be used when a user needs to launch WangWang. Automated batch operations are not allowed, and it must not be used as a decryption interface to show plaintext to users. WangWang supports recall after encryption; do not use it for any other scenario.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `openUid` | java.lang.String | yes | openUid to be decrypted |  |

**Response (top level)**: `wangwangNick` java.lang.String – WangWang nickname used to launch WangWang

<a id="alibabafenxiaorelationadd-1"></a>
## 13. Add buyer-seller distribution relationship

`com.alibaba.fenxiao:alibaba.fenxiao.relationadd:1` · Tools · 买卖家分销关系添加  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao/alibaba.fenxiao.relationadd/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Add a buyer-seller distribution relationship by product ID.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `offerId` | java.lang.Long | yes | Product ID | 98129931 |

**Response (top level)**: `result` ResultModel – ; `result.success` boolean – ; `result.code` java.lang.String – ; `result.message` java.lang.String – ; `result.result` java.lang.Boolean – 

<a id="productsearchtopkeyword-1"></a>
## 14. Trending product search keywords

`com.alibaba.fenxiao.crossborder:product.search.topKeyword:1` · Market Insights · 商品热搜词  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/product.search.topKeyword/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (optional) · user authorization not required · signature required  
Trending product search keywords.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `topSeKeywordParam` | TopSeKeywordParam | yes |  |  |
| `topSeKeywordParam.country` | java.lang.String | yes | Language; refer to the developer reference enum | en-英语 |
| `topSeKeywordParam.sourceId` | java.lang.String | yes | Query ID, e.g. category ID | 1-类目id |
| `topSeKeywordParam.hotKeywordType` | java.lang.String | yes | Hot search type; currently only the category dimension is provided, pass the fixed value cate | cate-类目 |

**Response (top level)**: `result` ResultModel – ; `result.success` java.lang.String – ; `result.code` java.lang.String – ; `result.message` java.lang.String – ; `result.result` TopSeKeywordModel[] – ; …

<a id="producttoplistquery-1"></a>
## 15. Query ranking lists

`com.alibaba.fenxiao.crossborder:product.topList.query:1` · Market Insights · 查询榜单列表  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/product.topList.query/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Query ranking (top) lists.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `rankQueryParams` | RankQueryParams | yes |  |  |
| `rankQueryParams.rankId` | java.lang.String | yes | Ranking list ID; a category ID can be passed. Category rankings are currently supported | 1111 |
| `rankQueryParams.rankType` | java.lang.String | yes | Ranking list type: complex (comprehensive ranking), hot (best-seller ranking), goodPrice (best-price ranking) | complex |
| `rankQueryParams.limit` | java.lang.Integer | yes | Number of products on the ranking list, maximum 20 | 10 |
| `rankQueryParams.language` | java.lang.String | yes | Ranking list product language | en |

**Response (top level)**: `result` ResultModel – ; `result.success` java.lang.Boolean – Whether successful; `result.code` java.lang.String – Error code; `result.message` java.lang.String – Error description; `result.result` RankModel – Result; …

<a id="productanalyzegetperdaysellquantitytrendnew-1"></a>
## 16. Get daily sales-quantity trend for a product (new)

`com.alibaba.fenxiao.crossborder:product.analyze.getPerdaySellQuantityTrendNew:1` · Market Insights · 获取商品每日销售数量趋势（新）  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/product.analyze.getPerdaySellQuantityTrendNew/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Get a product's daily sales-quantity trend over 90 days (at most 90 days of data can be queried). New interface.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `tendQueryParam` | OfferSellTrendQueryParam | yes |  |  |
| `tendQueryParam.offerId` | java.lang.Long | yes | Product ID | 123 |
| `tendQueryParam.startDate` | java.lang.String | yes | Query start time | 20240701(起始时间和截止时间不要超过1个月） |
| `tendQueryParam.endDate` | java.lang.String | yes | Query end time | 20240704 |

**Response (top level)**: `result` CommonResult – ; `result.success` java.lang.Boolean – Whether successful; `result.retCode` java.lang.String – Error code; `result.retMsg` java.lang.String – Return message; `result.result` OfferSellTrendDataModel[] – Result; …

<a id="productanalyzegetthirtydaymedianpricetrendnew-1"></a>
## 17. Get 30-day median-price trend for a product (new)

`com.alibaba.fenxiao.crossborder:product.analyze.getThirtyDayMedianPriceTrendNew:1` · Market Insights · 获取商品30天价格中位数指标趋势（新）  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/product.analyze.getThirtyDayMedianPriceTrendNew/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Get a product's 30-day median-price metric trend. New interface.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `trendQueryParam` | OfferSellTrendQueryParam | yes |  |  |
| `trendQueryParam.offerId` | java.lang.Long | yes | Product ID | 123 |
| `trendQueryParam.startDate` | java.lang.String | yes | Query start time | 20240701(开始时间到截止时间不要超过1个月) |
| `trendQueryParam.endDate` | java.lang.String | yes | Query end time | 20240704 |

**Response (top level)**: `result` CommonResult – ; `result.success` java.lang.Boolean – Whether successful; `result.retCode` java.lang.String – Error code; `result.retMsg` java.lang.String – Return message; `result.result` OfferSellTrendDataModel[] – Result; …

<a id="productanalyzegetrepurchaseratetrendnew-1"></a>
## 18. Get 90-day repurchase-rate trend for a product (new)

`com.alibaba.fenxiao.crossborder:product.analyze.getRepurchaseRateTrendNew:1` · Market Insights · 获取商品90天复购率指标趋势(新）  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/product.analyze.getRepurchaseRateTrendNew/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Get a product's 90-day repurchase-rate metric trend (at most 30 days can be queried). New interface.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `trendQueryParam` | OfferSellTrendQueryParam | yes |  |  |
| `trendQueryParam.offerId` | java.lang.Long | yes | Product ID | 123 |
| `trendQueryParam.startDate` | java.lang.String | yes | Query start time | 20240701(起始时间和截止时间不要超过一个月） |
| `trendQueryParam.endDate` | java.lang.String | yes | Query end time | 20240710 |

**Response (top level)**: `result` CommonResult – ; `result.success` java.lang.Boolean – Whether successful; `result.retCode` java.lang.String – Error code; `result.retMsg` java.lang.String – Return message; `result.result` OfferSellTrendDataModel[] – Result; …

<a id="categorytranslationgetbyid-1"></a>
## 19. Query multilingual category by category ID

`com.alibaba.fenxiao.crossborder:category.translation.getById:1` · Categories · 根据类目ID查询多语言类目  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/category.translation.getById/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Multilingual category query. Returns the category details in the requested language for the given category ID, including the list of its child categories.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `outMemberId` | String | no | User's unique ID within the organization. No more than 64 characters, consisting of digits and letters. | 23423532fwef |
| `language` | java.lang.String | yes | Language. See the enum in the FAQ. | ja |
| `categoryId` | java.lang.Long | yes | Category ID | 0 |
| `parentCateId` | Long | no | Parent ID | 0 |

**Response (top level)**: `result` ResultModel – Return result; `result.success` boolean – Whether successful; `result.code` java.lang.String – Error code; `result.message` java.lang.String – Error description; `result.result` Category – Actual result; …

<a id="categorytranslationgetbykeyword-1"></a>
## 20. Query multilingual category by category name

`com.alibaba.fenxiao.crossborder:category.translation.getByKeyword:1` · Categories · 根据类目名称查询多语言类目  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/category.translation.getByKeyword/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Multilingual category query. Returns the list of matching category details in the requested language for the given category name. Child-category data is not included.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `outMemberId` | java.lang.String | no | User's unique ID within the organization. No more than 64 characters, consisting of digits and letters. | 2423523f13tr12412f |
| `language` | java.lang.String | yes | Language. See the enum in the FAQ. | ja |
| `cateName` | java.lang.String | yes | Category name. Supports fuzzy search. | 裙子 |

**Response (top level)**: `result` ResultModel – Return result; `result.success` boolean – Whether successful; `result.code` java.lang.String – Error code; `result.message` java.lang.String – Error description; `result.result` Category[] – Actual result; …

<a id="alibabacategoryattributeget-1"></a>
## 21. Get leaf-category attributes

`com.alibaba.product:alibaba.category.attribute.get:1` · Categories · 获取叶子类目属性  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.product/alibaba.category.attribute.get/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Get category attributes by leaf-category ID.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `categoryID` | Long | yes | Category ID |  |
| `webSite` | String | yes | Site information, specifies whether the called API belongs to the international site (alibaba) or the 1688 website (1688). |  |
| `scene` | String | no | Scene value; optional values are empty and processing, default is empty |  |

**Response (top level)**: `attributes` AttributeInfo[] – Category attribute info; `levelAttrRelList` PostLevelAttrRel[] – (Deprecated) Category attribute cascading relationship; this field is only returned for 1688 business; `attributeLevelMapStr` java.util.Map – Cascading information string, can be cast to a map; `errorMsg` String – Error description; `errorCode` String – Error code; `success` Boolean – Whether successful; …

<a id="productsearchkeywordquery-1"></a>
## 22. Multilingual keyword search

`com.alibaba.fenxiao.crossborder:product.search.keywordQuery:1` · Products · 多语言关键词搜索  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/product.search.keywordQuery/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Multilingual keyword search.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `offerQueryParam` | OfferQueryParam | yes | Query parameter | {} |
| `offerQueryParam.keyword` | java.lang.String | yes | Keyword | 饼干 |
| `offerQueryParam.beginPage` | java.lang.Integer | yes | Pagination | 1 |
| `offerQueryParam.pageSize` | java.lang.Integer | yes | Pagination | 1 |
| `offerQueryParam.filter` | java.lang.String | no | Filter parameters, multiple values separated by commas; see the solution introduction for the enumeration | shipInToday,ksCiphertext |
| `offerQueryParam.sort` | java.lang.String | no | Sorting parameter; for enum values see the solution introduction | {"price":"asc"} |
| `offerQueryParam.outMemberId` | java.lang.String | no | External user ID | 123 |
| `offerQueryParam.priceStart` | java.lang.String | no | Wholesale price start | 1 |
| `offerQueryParam.priceEnd` | java.lang.String | no | Wholesale price range end | 10 |
| `offerQueryParam.categoryId` | java.lang.Long | no | Category ID | 1 |
| `offerQueryParam.categoryIdList` | String | no | List of category IDs, separated by English commas; supports the union of multiple categories | 2,45 |
| `offerQueryParam.country` | java.lang.String | yes | Language | 如en-英语，详细枚举参考开发人员参考菜单 |
| `offerQueryParam.regionOpp` | String | no | Business opportunity | 枚举值见开发人员参考菜单 |
| `offerQueryParam.productCollectionId` | String | no | Xunyuantong workbench assortment id | 174316138 |
| `offerQueryParam.snId` | String | no | Search navigation ID, e.g. 978 or 978:1352 | 978:1352 |
| `offerQueryParam.keywordTranslate` | Boolean | no | Whether the keyword has already been translated; default is not translated. If true, keyword translation is skipped | false |
| `offerQueryParam.saleFilterList` | SaleFilterParam[] | no | Sales volume filter parameter | [{"saleType":"sales7","saleStart":"10","saleEnd":"100"}] |
| `offerQueryParam.saleFilterList[].saleType` | String | no | Sales volume type | sales7:近7天销量,sales14:近14天销量,sales30:近30天销量,totalSales:总销量，传入多个取交集 |
| `offerQueryParam.saleFilterList[].saleStart` | String | no | Minimum sales volume | 10 |
| `offerQueryParam.saleFilterList[].saleEnd` | String | no | Maximum sales volume | 100 |

**Response (top level)**: `result` ResultModelV3 – Return message; `result.success` Boolean – Whether normal; `result.code` String – Status code; `result.message` String – Prompt; `result.result` PageInfoV3 – Content; …

<a id="productsearchimagequery-1"></a>
## 23. Multilingual image search

`com.alibaba.fenxiao.crossborder:product.search.imageQuery:1` · Products · 多语言图搜  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/product.search.imageQuery/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Multilingual image search.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `offerQueryParam` | OfferQueryParam | yes | {&quot;beginPage&quot;:1,&quot;country&quot;:&quot;en&quot;,&quot;imageAddress&quot;:&quot;https://cbu01.alicdn.com/img/ibank/O1CN01q5lIoD1ZPh3gN3www_!!2928623187-0-cib.jpg&quot;,&quot;pageSize&quot;:1,&quot;userId&quot;:0} | {"beginPage":1,"country":"en","imageAddress":"https://cbu01.alicdn.com/img/ibank/O1CN01q5l… |
| `offerQueryParam.imageId` | java.lang.String | yes | Image ID, required | 图片id |
| `offerQueryParam.beginPage` | java.lang.Integer | yes | Pagination | 分页 |
| `offerQueryParam.pageSize` | java.lang.Integer | yes | Pagination, maximum 50; 20 is recommended for best results | 分页 |
| `offerQueryParam.region` | java.lang.String | no | Entity selection | 266,799,48,581 |
| `offerQueryParam.filter` | java.lang.String | no | Filter parameters, multiple values separated by commas; see the solution introduction for the enumeration | shipInToday,ksCiphertext |
| `offerQueryParam.sort` | java.lang.String | no | Sorting parameter; for enum values see the solution introduction | {"price":"asc"} |
| `offerQueryParam.outMemberId` | java.lang.String | no | External user uid | 外部用户uid |
| `offerQueryParam.priceStart` | java.lang.String | no | Wholesale price start | 10 |
| `offerQueryParam.priceEnd` | java.lang.String | no | Wholesale price range end | 20 |
| `offerQueryParam.categoryId` | java.lang.Long | no | Category ID | 类目id |
| `offerQueryParam.imageAddress` | java.lang.String | no | Image URL; only used in the scenario of querying with a 1688 image link, other cases are not guaranteed to return data | 图片地址 |
| `offerQueryParam.country` | java.lang.String | yes | Language | 如en-英语，详细枚举请参考开发人员参考菜单 |
| `offerQueryParam.keyword` | String | no | Search within results | 书本 |
| `offerQueryParam.auxiliaryText` | String | no | Multimodal image-search copy text | 热卖的 |
| `offerQueryParam.productCollectionId` | String | no | Xunyuantong workbench assortment ID | 21432232 |
| `offerQueryParam.keywordTranslate` | Boolean | no | Whether the search term has already been translated; if true, search directly without translating the keyword | false |
| `offerQueryParam.itemTitle` | String | no | Product title | 商品标题 |

**Response (top level)**: `result` ResultModelV5 – Return value; `result.success` String – Whether successful; `result.code` String – code; `result.message` String – message; `result.result` PageInfoV4 – Result; …

<a id="productsearchquerysellerofferlist-1"></a>
## 24. Multilingual in-store product search

`com.alibaba.fenxiao.crossborder:product.search.querySellerOfferList:1` · Products · 多语言商品店搜  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/product.search.querySellerOfferList/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Multilingual search of the products in a seller's store.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `offerQueryParam` | OfferQueryParam | yes | Request parameters | {} |
| `offerQueryParam.keyword` | String | no | Keyword | 饼干 |
| `offerQueryParam.beginPage` | Integer | yes | Pagination | 1 |
| `offerQueryParam.pageSize` | Integer | yes | Pagination | 1 |
| `offerQueryParam.filter` | String | no | Filter parameters, multiple values separated by commas; see the solution introduction for the enumeration | shipInToday,ksCiphertext |
| `offerQueryParam.sort` | String | no | Sorting parameter; for enum values see the solution introduction | {"price":"asc"} |
| `offerQueryParam.outMemberId` | String | no | External user ID | 123 |
| `offerQueryParam.priceStart` | String | no | Wholesale price start | 1 |
| `offerQueryParam.priceEnd` | String | no | Wholesale price range end | 1 |
| `offerQueryParam.categoryId` | Long | no | Category ID | 1 |
| `offerQueryParam.country` | String | yes | City | japan |
| `offerQueryParam.sellerOpenId` | String | yes | Masked merchant store ID | 123 |

**Response (top level)**: `result` ResultModelV3 – Return message; `result.success` Boolean – Whether normal; `result.code` String – Status code; `result.message` String – Prompt; `result.result` PageInfoV3 – Content; …

<a id="productsearchqueryproductdetail-1"></a>
## 25. Multilingual product detail

`com.alibaba.fenxiao.crossborder:product.search.queryProductDetail:1` · Products · 多语言商详  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/product.search.queryProductDetail/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Multilingual product detail.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `offerDetailParam` | OfferDetailParam | yes | Parameter | {"offerId":1,"country":"en"} |
| `offerDetailParam.offerId` | java.lang.Long | yes | Product ID | 1 |
| `offerDetailParam.country` | java.lang.String | yes | Language | ja-日语 en-英语 |
| `offerDetailParam.outMemberId` | java.lang.String | no | External user ID | 1 |
| `offerDetailParam.currency` | String | no | Currency code | HKD |

**Response (top level)**: `result` ResultModel – Result; `result.success` boolean – Whether successful; `result.code` java.lang.String – Return code; `result.message` java.lang.String – Prompt; `result.result` ProductDetailModel – Result; …

<a id="productsearchofferrecommend-1"></a>
## 26. Product recommendations

`com.alibaba.fenxiao.crossborder:product.search.offerRecommend:1` · Products · 商品推荐  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/product.search.offerRecommend/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Personalized product recommendations.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `recommendOfferParam` | RecommendOfferParam | yes | Recommended parameters | 1 |
| `recommendOfferParam.beginPage` | java.lang.Integer | yes | Starting page number | 1 |
| `recommendOfferParam.pageSize` | java.lang.Integer | yes | Page size | 20 |
| `recommendOfferParam.country` | java.lang.String | yes | Language | en-英语，ja-日语，具体请参考开放参考菜单 |
| `recommendOfferParam.outMemberId` | java.lang.String | yes | External organization user id; this id is the service provider platform id | dferg0001 |

**Response (top level)**: `result` ResultModel – Return message; `result.success` Boolean – Whether successful; `result.code` String – Message code; `result.message` String – Message content; `result.result` ProductInfoModel[] – Result; …

<a id="productsearchkeywordsnquery-1"></a>
## 27. Multilingual search navigation

`com.alibaba.fenxiao.crossborder:product.search.keywordSNQuery:1` · Products · 多语言搜索导航  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/product.search.keywordSNQuery/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Get the multilingual keyword-search navigation list.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `snParams` | KeywordSNQueryParams | yes |  |  |
| `snParams.keyword` | java.lang.String | yes | Search keyword | skirt |
| `snParams.language` | java.lang.String | yes | Language, e.g. English en_US; see [Developer Reference] | en_US |
| `snParams.region` | java.lang.String | yes | Region, e.g. United States US, refer to [Developer Reference] | US |
| `snParams.currency` | java.lang.String | yes | Currency, e.g. US Dollar USD, refer to [Developer Reference] | USD |

**Response (top level)**: `result` CommonResult – Result; `result.success` java.lang.Boolean – Call result; `result.retCode` java.lang.String – Error code; `result.retMsg` java.lang.String – Error description; `result.result` KeywordSNModel[] – Actual result; …

<a id="productrelatedrecommend-1"></a>
## 28. Related product recommendations

`com.alibaba.fenxiao.crossborder:product.related.recommend:1` · Products · 相关性商品推荐  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/product.related.recommend/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Related product recommendations.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `relatedQueryParams` | RelatedQueryParams | yes | Input parameter | 如下 |
| `relatedQueryParams.offerId` | java.lang.Long | yes | Product ID | 689337673960 |
| `relatedQueryParams.pageNo` | java.lang.Integer | yes | Page number | 1 |
| `relatedQueryParams.pageSize` | java.lang.Integer | yes | Number per page, maximum 10 | 10 |
| `relatedQueryParams.language` | java.lang.String | yes | Language | en |

**Response (top level)**: `result` ResultModel – Result; `result.success` boolean – Whether successful; `result.code` java.lang.String – Error code; `result.message` java.lang.String – Error description; `result.result` ProductInfoModel[] – Result; …

<a id="productimageupload-1"></a>
## 29. Upload an image to get an imageId

`com.alibaba.fenxiao.crossborder:product.image.upload:1` · Products · 上传图片获取imageId  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/product.image.upload/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Upload an image and receive an imageId.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `uploadImageParam` | UploadImageParam | yes |  |  |
| `uploadImageParam.imageBase64` | java.lang.String | yes | Image base64 | 12 |
| `uploadImageParam.outMemberId` | java.lang.String | no | External user ID | 2 |

**Response (top level)**: `result` ResultModel – ; `result.success` java.lang.String – ; `result.code` java.lang.String – ; `result.message` java.lang.String – ; `result.result` java.lang.String – 

<a id="couponoptimalclaim-1"></a>
## 30. Claim the optimal coupon for a product

`com.alibaba.marketing:coupon.optimal.claim:1` · Products · 通过商品领取最优化的优惠券  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.marketing/coupon.optimal.claim/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Claim the best-value coupon for a product. Usually called before placing an order to complete the optimal coupon-claiming strategy.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `offerIds` | Long[] | yes | List of product IDs | [24910983123,2799731973] |

**Response (top level)**: `result` ResultModel – ; `result.success` boolean – ; `result.code` java.lang.String – ; `result.message` java.lang.String – ; `result.result` BestBizCouponGetResult – ; …

<a id="alibabaproductsuggestcrossborder-1"></a>
## 31. Recommend products by keyword (cross-border)

`com.alibaba.product:alibaba.product.suggest.crossBorder:1` · Products · 跨境场景根据关键字推荐商品  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.product/alibaba.product.suggest.crossBorder/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Recommend products by keyword and category in cross-border scenarios, sorted by sales volume. Note: this API is rate-limited and is only suitable for manual-association scenarios.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `keyWord` | String | yes | Search keyword for the product, usually the product title | 商品标题 |
| `loginId` | String | no | Seller loginId | alitestforisv01 |

**Response (top level)**: `resultList` ProductSearchResultInfo[] – Search return results; `success` String – Whether successful; `errorMsg` String – Error description; `errorCode` String – Error code; …

<a id="alibabaproductsimpleget-1"></a>
## 32. Get simple product info from a previously purchased supplier

`com.alibaba.product:alibaba.product.simple.get:1` · Products · 获取已购买过商家的商品简单信息  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.product/alibaba.product.simple.get/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Get product details by product ID. This interface returns simple information for products of suppliers you have already purchased from. Access to this interface is paid. It returns only basic information and is mainly intended for data association in ERP systems.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `productID` | Long | yes | Product ID | 565507182121 |
| `webSite` | String | yes | Site information, specifies whether the called API belongs to the international site (alibaba) or the 1688 website (1688). | 1688 |

**Response (top level)**: `productInfo` ProductInfo – Product detailed information; `productInfo.productID` Long – Product ID; `productInfo.productType` String – Product type: online wholesale product (wholesale) or inquiry product (sourcing). Defaults to wholesale on the 1688 website.; `productInfo.categoryID` Long – Category ID, identifies the category the product belongs to; `productInfo.attributes` ProductAttribute[] – Product attributes and attribute values; `productInfo.groupID` Long[] – Group ID, determines the group the product belongs to. On 1688, multiple group IDs can be passed in; on the international site, a product can belong to only one group, so by default only the first one is taken.; `productInfo.status` String – Product status. published: online status; member expired: revoked by member; auto expired: naturally expired; expired: expired (includes both manually and automatically expired); member deleted: deleted by member; modified: modified; new: newly published; deleted: deleted; TBD: to be delete; approved: approved; auditing: under review; untread: review not passed;; `productInfo.subject` String – Product title, up to 128 characters; `productInfo.description` String – Product detail description, may include image URLs from the image center; `productInfo.language` String – Language; see the FAQ for language enum values. The 1688 website passes CHINESE by default; `productInfo.periodOfValidity` Integer – Information validity period, calculated in days; not applicable for the international site; `productInfo.bizType` Integer – Business type. 1: Product, 2: Processing, 3: Agency, 4: Cooperation, 5: Business service. The international site defaults to Product.; `productInfo.pictureAuth` Boolean – Whether the image is private information; this field is invalid for the international site; `productInfo.image` ProductImageInfo – Product main image; …

<a id="productkjdistributeaddrelation-1"></a>
## 33. Follow a product (cross-border)

`com.alibaba.fenxiao.crossborder:product.kjdistribute.addRelation:1` · Products · 增加跨境关注商品  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/product.kjdistribute.addRelation/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (optional) · user authorization not required · signature required  
Add a cross-border followed product.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `offerId` | java.lang.Long | yes | Product ID |  |

**Response (top level)**: `result` ResultModel – ; `result.msg` java.lang.String – ; `result.success` java.lang.Boolean – ; `result.traceId` java.lang.String – 

<a id="productkjdistributeremoverelation-1"></a>
## 34. Unfollow a product (cross-border)

`com.alibaba.fenxiao.crossborder:product.kjdistribute.removeRelation:1` · Products · 取消跨境关注商品  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/product.kjdistribute.removeRelation/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Remove a cross-border followed product.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `offerId` | Long | yes | Product ID | 232324312 |

**Response (top level)**: `result` ResultModel – Return result; `result.model` java.lang.String[] – ; `result.msg` java.lang.String – ; `result.traceId` java.lang.String – ; `result.success` java.lang.Boolean – 

<a id="productdistributegetdistributeinfo-1"></a>
## 35. Get product selling points for listing

`com.alibaba.fenxiao.crossborder:product.distribute.getDistributeInfo:1` · Products · 获取商品铺货卖点  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/product.distribute.getDistributeInfo/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (optional) · user authorization not required · signature required  
For listing scenarios: get the selling-point information needed to list a product.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `offerId` | java.lang.String | yes | offerId | 1 |

**Response (top level)**: `result` Result – Return result; `result.success` java.lang.Boolean – Whether the request was successful; `result.msg` java.lang.String – Error message; `result.data` ItemAIGCVO – Specific business data; …

<a id="productcategorygetattrbyid-1"></a>
## 36. Query leaf-category attribute and value mappings

`com.alibaba.fenxiao.crossborder:product.category.getAttrById:1` · Products · 叶子类目属性属性值映射查询  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/product.category.getAttrById/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Category mapping query.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `categoryId` | java.lang.String | yes |  |  |

**Response (top level)**: `success` Boolean – ; `msg` String – ; `data` PVMappingVO – ; `data.categoryId` java.lang.String – Category ID; `data.categoryName` java.lang.String – Category name; `data.externalCategoryId` java.lang.String – External category ID; `data.externalCategoryName` java.lang.String – External category name; `data.attributes` PVMappingAttribute[] – Attribute; `data.productShippingInfo` ProductShippingInfo – Shipping information; …

<a id="alibabacrossproductinfo-1"></a>
## 37. Get product details (cross-border)

`com.alibaba.product:alibaba.cross.productInfo:1` · Product Selection & Listing · 跨境场景获取商品详情  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.product/alibaba.cross.productInfo/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Get product details in cross-border scenarios. A cross-border listing relationship must be established before details can be retrieved.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `productId` | java.lang.Long | yes | 1688 product ID | 573741401425 |

**Response (top level)**: `productInfo` ProductInfo – Product details; `productInfo.productID` Long – Product ID; `productInfo.productType` String – Product type: online wholesale product (wholesale) or inquiry product (sourcing). Defaults to wholesale on the 1688 website.; `productInfo.categoryID` Long – Category ID, identifies the category the product belongs to; `productInfo.attributes` ProductAttribute[] – Product attributes and attribute values; `productInfo.groupID` Long[] – Group ID, determines the group the product belongs to. On 1688, multiple group IDs can be passed in; on the international site, a product can belong to only one group, so by default only the first one is taken.; `productInfo.status` String – Product status. published: online status; member expired: revoked by member; auto expired: naturally expired; expired: expired (includes both manually and automatically expired); member deleted: deleted by member; modified: modified; new: newly published; deleted: deleted; TBD: to be delete; approved: approved; auditing: under review; untread: review not passed;; `productInfo.subject` String – Product title, up to 128 characters; `productInfo.description` String – Product detail description, may include image URLs from the image center; `productInfo.language` String – Language; see the FAQ for language enum values. The 1688 website passes CHINESE by default; `productInfo.periodOfValidity` Integer – Information validity period, calculated in days; not applicable for the international site; `productInfo.bizType` Integer – Business type. 1: Product, 2: Processing, 3: Agency, 4: Cooperation, 5: Business service. The international site defaults to Product.; `productInfo.pictureAuth` Boolean – Whether the image is private information; this field is invalid for the international site; `productInfo.image` ProductImageInfo – Product main image; …

<a id="alibabacrosssyncproductlistpushed-1"></a>
## 38. Add products to the listing list (cross-border)

`com.alibaba.product.push:alibaba.cross.syncProductListPushed:1` · Product Selection & Listing · 跨境场景下将商品加入铺货列表  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.product.push/alibaba.cross.syncProductListPushed/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Cross-border only. Adds products to the listing list (that is, creates listing relationships); at most 20 items per call. Only after a product is added can its details be queried through the product-detail interface. Contact the cross-border operations staff to configure permissions manually before calling.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `productIdList` | Long[] | yes | List of 1688 product IDs; the list length must not exceed 20. | [123456] |

**Response (top level)**: `result` commonResult – Sync result; `result.errorCode` java.lang.String – Error code; `result.errorMsg` java.lang.String – Error message; `result.success` boolean – Whether successful

<a id="alibabaproductpushsyncpushproductresult-1"></a>
## 39. Sync listing results

`com.alibaba.product.push:alibaba.product.push.syncPushProductResult:1` · Product Selection & Listing · 同步铺货结果  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.product.push/alibaba.product.push.syncPushProductResult/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Sync listing results. When an ISV lists products from the source platform (1688) onto a target platform (for example TAOBAO), the ISV must return the listing result. The listing-status descriptions must match those defined by the source platform (1688). This interface also supports operations such as delisting, all of which are expressed through the listing status.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `platformDefinition` | PlatformDefinition | yes | Definition of the target platform | {} |
| `platformDefinition.definitionId` | java.lang.String | yes | Platform ID, defined by Alibaba. For example, Taobao is www.taobao.com. Enum values: Amazon: AMAZON, AliExpress: AE, Wish: WISH, eBay: EBAY, Lazada: LAZADA, Taobao: TAOBAO |  |
| `pushRecordIdentity` | Identity | no | During bulk listing, the source platform may generate a batch for each listing operation and pass it to the ISV. The ISV can return this field in the sync notification. This field is passed from the platform to the ISV and is not required. | {} |
| `pushRecordIdentity.content` | java.lang.String | yes |  |  |
| `pushProductResults` | PushProductResult | yes | Product-level listing (distribution) result | {} |
| `pushProductResults.productIdInTargetPlatform` | java.lang.String | yes | Product ID on the target platform, a numeric-only string |  |
| `pushProductResults.productIdInPartner` | java.lang.String | yes | Product ID on the third-party platform |  |
| `pushProductResults.productIdInSource` | java.lang.String | yes | Product ID on the source platform, a pure numeric string |  |
| `pushProductResults.productPushStatus` | java.lang.String | yes | Listing status. 0: not successful, 1: successful |  |
| `pushProductResults.productInfoInTargetPlatform` | SimpleItemDesc | yes | Product-level listing (distribution) result |  |
| `pushProductResults.productInfoInTargetPlatform.id` | java.lang.String | yes | Unique identifier |  |
| `pushProductResults.productInfoInTargetPlatform.price` | java.lang.Double | yes | Price |  |
| `pushProductResults.productInfoInTargetPlatform.subject` | java.lang.String | yes | Product name |  |
| `pushProductResults.productInfoInTargetPlatform.description` | java.lang.String | yes | Description |  |
| `pushProductResults.productInfoInTargetPlatform.url` | java.lang.String | yes | Product URL |  |
| `pushProductResults.productInfoInTargetPlatform.priceRanges` | priceRanges[] | yes | Price range |  |
| `pushProductResults.productInfoInTargetPlatform.priceRanges[].price` | double | yes |  |  |
| `pushProductResults.productInfoInTargetPlatform.priceRanges[].startQuantity` | Integer | yes |  |  |
| `pushProductResults.skus` | PushProductSKUResult[] | yes | SKU array |  |
| `pushProductResults.skus[].skuIdInSource` | java.lang.String | yes | SKU identifier on the source platform | 312312312312 |
| `pushProductResults.skus[].skuIdInPartner` | java.lang.String | no | SKU identifier on the third-party platform | 101123123 |
| `pushProductResults.skus[].skuPushStatus` | java.lang.String | yes | Listing status | success |
| `pushProductResults.skus[].skuIdInTargetPlatform` | java.lang.String | yes | SKU identifier on the target platform | 10111 |
| `pushProductResults.skus[].skuInfoInTargetPlatform` | SimpleItemDesc | no | SKU listing result | {} |
| `pushProductResults.skus[].skuInfoInTargetPlatform.id` | java.lang.String | yes | Unique identifier |  |
| `pushProductResults.skus[].skuInfoInTargetPlatform.price` | java.lang.Double | yes | Price |  |
| `pushProductResults.skus[].skuInfoInTargetPlatform.subject` | java.lang.String | yes | Product name |  |
| `pushProductResults.skus[].skuInfoInTargetPlatform.description` | java.lang.String | yes | Description |  |
| `pushProductResults.skus[].skuInfoInTargetPlatform.url` | java.lang.String | yes | Product URL |  |
| `pushProductResults.skus[].skuInfoInTargetPlatform.priceRanges` | priceRanges[] | yes | Price range |  |
| `pushProductResults.skus[].skuInfoInTargetPlatform.priceRanges[].price` | double | yes |  |  |
| `pushProductResults.skus[].skuInfoInTargetPlatform.priceRanges[].startQuantity` | Integer | yes |  |  |
| `pushProductResults.skus[].mappingInfo` | attributeRelationMapping[] | no | Attribute mapping relationship | {} |
| `pushProductResults.skus[].mappingInfo[].propertyIdInSource` | java.lang.String | yes | The attribute ID on the source platform (1688) | 123 |
| `pushProductResults.skus[].mappingInfo[].propertyTextInSource` | java.lang.String | yes | The attribute ID text on the source platform (1688) | 颜色 |
| `pushProductResults.skus[].mappingInfo[].valueIdInSource` | java.lang.String | yes | Attribute value ID on the source platform (1688) | 1233 |
| `pushProductResults.skus[].mappingInfo[].valueTextInSource` | java.lang.String | yes | The attribute value text on the source platform (1688) | 红色 |
| `pushProductResults.skus[].mappingInfo[].propertyIdInTarget` | java.lang.String | yes | The attribute ID on the target platform (e.g. TAOBAO) | 234 |
| `pushProductResults.skus[].mappingInfo[].propertyTextInTarget` | java.lang.String | yes | Attribute ID text on the target platform (e.g. TAOBAO) | 颜色 |
| `pushProductResults.skus[].mappingInfo[].valueIdInTarget` | java.lang.String | yes | The attribute value ID on the target platform (e.g. TAOBAO) | 2341 |
| `pushProductResults.skus[].mappingInfo[].valueTextInTarget` | java.lang.String | yes | Attribute value text on the target platform (e.g. TAOBAO) | 橙红色 |
| `pushProductResults.userShopIdInTargetPlatform` | java.lang.String | yes | The user's store identifier on the target platform, a pure numeric string, corresponding to targetUserId on the backend. For example, if the target platform is Taobao, pass the customer's user ID on Taobao. |  |

**Response (top level)**: `errorCode` String – Error code; `errorMessage` String – Error description; `success` boolean – Whether successful

<a id="alibabacrossproductlist-1"></a>
## 40. Get product list (cross-border)

`com.alibaba.product:alibaba.cross.productList:1` · Product Selection & Listing · 跨境场景获取商品列表  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.product/alibaba.cross.productList/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
This interface validates the cross-border listing relationship and is for cross-border business only. Product model V2.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `productIdList` | Long[] | yes | Product Id list | [574325651942,570027659932] |

**Response (top level)**: `productList` ProductInfo[] – Product list; `success` Boolean – Whether successful; `message` String – Return message; …

<a id="alibabaproductunfollowcrossborder-1"></a>
## 41. Unfollow a product

`com.alibaba.product:alibaba.product.unfollow.crossborder:1` · Product Selection & Listing · 解除关注商品  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.product/alibaba.product.unfollow.crossborder/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Unfollow a product.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `productId` | Long | yes | Product ID | 36143645361 |

**Response (top level)**: `code` int – code, 0 means success; `message` String – Description of the result

<a id="alibabacategoryget-1"></a>
## 42. Query category by category ID

`com.alibaba.product:alibaba.category.get:1` · Product Selection & Listing · 根据类目Id查询类目  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.product/alibaba.category.get/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (optional) · user authorization not required · signature required  
Category query. To retrieve all 1688 categories, traverse the whole category tree starting from the root: first pass 0 to get all level-1 category IDs, then iterate the level-1 IDs to get all level-2 categories, and finally iterate the level-2 IDs to get the level-3 categories. Note: 1688 categories have only three levels; level-3 categories are the leaf categories required for publishing products.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `categoryID` | Long | yes | Category ID, must be greater than or equal to 0. If 0, all first-level categories are queried | 1031910 |

**Response (top level)**: `succes` String – Whether successful; `categoryInfo` CategoryInfo[] – List of categories; `errorMsg` String – Error message; `errorCode` String – Error code; …

<a id="alibabaproductfollowcrossborder-1"></a>
## 43. Follow a product

`com.alibaba.product:alibaba.product.follow.crossborder:1` · Product Selection & Listing · 关注商品  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.product/alibaba.product.follow.crossborder/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Follow a product.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `productId` | Long | yes | Product ID | 52312121144 |

**Response (top level)**: `code` int – code, 0 means success; `message` String – Description of the result

<a id="alibabarelationquerysuppliers-1"></a>
## 44. [Relationship] Distributor: query supplier list

`cn.alibaba.open:alibaba.relation.querySuppliers:1` · Product Selection & Listing · 【关系】分销商-查询供应商列表  
POST `https://gw.open.1688.com/openapi/param2/1/cn.alibaba.open/alibaba.relation.querySuppliers/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Get the supplier list for a distributor by userID.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `supplierLoginId` | String | no | Supplier login ID; specifying this parameter allows querying the distribution relationship between the authorized user and the specified supplier | 李定国 |
| `currentPage` | Integer | no | Current page number | 1 |
| `pageSize` | Integer | no | Number of items per page | 10 |

**Response (top level)**: `result` suppliers-result – ; `result.count` Integer – Total number of items; `result.currentPage` Integer – Current page number; `result.pageSize` Integer – Number of items per page; `result.relationModels` supplierModel[] – Result set; …

<a id="alibabatradecreatecrossorder-1"></a>
## 45. Create cross-border order

`com.alibaba.trade:alibaba.trade.createCrossOrder:1` · Orders · 跨境订单创建  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.trade.createCrossOrder/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Cross-border-only order creation. An order may contain at most 50 SKUs, all from the same supplier. For multiple suppliers or more than 50 SKUs, split the order yourself before submitting. In some special cases several orders are created at once and several order numbers are returned. Supports both the open-marketplace and distribution scenarios. Orders are placed under the main account or a sub-account depending on the currently authorized user.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `flow` | String | yes | general (create wholesale order), fenxiao (create distribution order), the saleproxy process will verify the distribution relationship, paired (Tiantian Texmai / Daily Deals), repurchase (repurchase contract order), boutiquefenxiao (curated supply distribution-price order; used when purchase quantity is 1, free shipping applies), boutiquepifa (curated supply wholesale-price order; used when purchase quantity is greater than 2). Different order-placing flow channels may cause differences in price and purchase protection services. | general |
| `message` | String | no | Buyer message | 留言 |
| `isvBizType` | String | no | Open platform business code, default is cross, representing the 1688 business side being integrated with. cross (cross-border business), cross_daigou (cross-border purchasing-agent business), cross_distribution (overseas distribution business); you can contact your business manager (Xiaoer) to confirm — this field does not affect order placement | cross |
| `addressParam` | address | yes | Shipping address info | {"address":"网商路699号","phone": "0517-88990077","mobile": "15251667788","fullName": "张三","po… |
| `addressParam.addressId` | Long | yes | Shipping address ID | 1234 |
| `addressParam.fullName` | String | yes | Recipient name | 张三 |
| `addressParam.mobile` | String | yes | Mobile phone | 15251667788 |
| `addressParam.phone` | String | yes | Phone number | 0517-88990077  注意: 只能填电话号码,不能加其他中文 |
| `addressParam.postCode` | String | yes | Postal code | 000000 |
| `addressParam.cityText` | String | yes | City text | 杭州市 |
| `addressParam.provinceText` | String | yes | Province text | 浙江省 |
| `addressParam.areaText` | String | yes | District text | 滨江区 |
| `addressParam.townText` | String | yes | Town text | 长河镇 |
| `addressParam.address` | String | yes | Street address | 网商路699号 |
| `addressParam.districtCode` | String | yes | Address code | 310107 |
| `addressParam.addressCode` | String | yes | Address code | 地址码 |
| `cargoParamList` | cargo[] | yes | Product information | [{"specId": "b266e0726506185beaf205cbae88530d","quantity": 5,"offerId": 554456348334},{"sp… |
| `cargoParamList[].offerId` | Long | yes | offer id corresponding to the product | 554456348334 |
| `cargoParamList[].specId` | String | yes | Product SKU ID | b266e0726506185beaf205cbae88530d |
| `cargoParamList[].quantity` | Double | yes | Product quantity (used for amount calculation) | 5 |
| `cargoParamList[].openOfferId` | String | no | Encrypted offerId. When the search result returns only openOfferId, use openOfferId in place of offerId to place the order. | Wcv2w970KoL1BCpJGQp3vwKvcBThOPlpHnUtkK3T0n4= |
| `cargoParamList[].outMemberId` | String | no | External downstream member ID | 98928912-23 |
| `invoiceParam` | invoice | no | Invoice information | {"invoiceType":0,"cityText": "杭州市","provinceText": "浙江省","address": "网商路699号","phone": "05… |
| `invoiceParam.invoiceType` | Integer | yes | Invoice type 0: regular invoice, 1: VAT invoice | 0 |
| `invoiceParam.provinceText` | String | yes | Province text | 浙江省 |
| `invoiceParam.cityText` | String | yes | City text | 杭州市 |
| `invoiceParam.areaText` | String | yes | Region text | 滨江区 |
| `invoiceParam.townText` | String | yes | Town text | 长河镇 |
| `invoiceParam.postCode` | String | yes | Postal code | 333333 |
| `invoiceParam.address` | String | yes | Street | 网商路699号 |
| `invoiceParam.fullName` | String | yes | Invoice recipient's name | 张三 |
| `invoiceParam.phone` | String | yes | Phone number | 0517-88990077 |
| `invoiceParam.mobile` | String | yes | Mobile phone | 15251667788 |
| `invoiceParam.companyName` | String | yes | Purchasing company name (invoice title) | 测试公司 |
| `invoiceParam.taxpayerIdentifier` | String | yes | Tax identification number | 12345 |
| `invoiceParam.bankAndAccount` | String | yes | Bank name and account number | 网商银行 |
| `invoiceParam.localInvoiceId` | String | yes | VAT local invoice number | 123123123 |
| `tradeType` | String | no | Since different products support different trade modes and no single trade mode applies universally, the trade mode usable for the current order must be obtained from tradeModeNameList in the order preview API. Trade mode type descriptions: assureTrade (Trade 4.0 universal escrow transaction), alipay (Alipay escrow transaction used generally in the open marketplace; currently being phased out and will be removed later), period (regular account period transaction), assure (the escrow transaction process required when a large buyer enterprise places an order via procurement inquiry/quote), creditBuy (Cheng-e-She), bank (bank transfer), 631staged (631 staged payment), 37staged (37 staged payment). If this field is not passed, the system defaults to selecting an available trade mode for the order; if Cheng-e-She is enabled the default is creditBuy (Cheng-e-She), otherwise the default is Alipay escrow transaction. | assureTrade |
| `shopPromotionId` | String | no | Store discount ID, obtained via the "preview data before order creation" API. If empty, the default discount is used | itemCoupon-5600812521_31032085284-398517001570 |
| `anonymousBuyer` | Boolean | no | Whether the order is anonymous |  |
| `fenxiaoChannel` | String | no | Downstream platform for reflow orders: Taobao-thyny, Tmall-tm, Taote-taote, Alibaba C2M-c2m, JD-jingdong, Pinduoduo-pinduoduo, WeChat-weixin, Cross-border-kuajing, Kuaishou-kuaishou, Youzan-youzan, Douyin-douyin, Siku-siku, Meituan Tuanhaohuo-meituan, Xiaohongshu-xiaohongshu, Dangdang-dangdang, Suning-suning, DaVdian-davdian, Xingyun-xingyun, Miya-miya, Boluopai Mall-boluo, Kuaituantuan-kuaituantuan, Other-other | douyin |
| `inventoryMode` | String | no | Stock mode, JIT (JIT mode) or NORMAL (warehouse fulfillment mode). Currently only available for AE. | JIT |
| `outOrderId` | String | no | External order number | 988129883123 |
| `pickupService` | String | no | Door-to-door pickup, currently available for AE supply, not yet enabled for other scenarios. y or n, default is n | n |
| `warehouseCode` | String | no | Door-to-door pickup warehouse code | any |
| `preSelectPayChannel` | String | no | Pre-selected payment channel, used for financial order routing. Returned by the order information query API as result.exAttributes.preSelectPayChannel; this value is the pre-selected payment channel flag passed in when creating the order via the create-order API. | alipay |
| `smallProcurement` | String | no | Whether it is a small-amount purchase. Currently available for AE supply. Value: y/n, default is n. | y |
| `useRedEnvelope` | String | no | Whether to use a red packet (coupon): y - use, n - do not use. Uses red packet by default | n |
| `dropshipping` | String | no | Whether it is a transshipment order. Value: y/n, default is n. | y |
| `addedService` | String | no | Value-added service, toB / toC | toB |
| `crossBorderLogisticsSolutionId` | String | no | The sourceId of the official cross-border logistics solution selected by the user | GLOBAL_CAINIAO_VN_TRANSIT_LAND |
| `useCrossBorderLogisticsSolution` | Boolean | no | Whether to use the cross-border logistics solution preview. Defaults to false. This parameter does not take effect for overseas addresses. It takes effect for Hong Kong, Macao, and Taiwan addresses. | true |
| `extendParam` | ExtendParam | no | Order placement extension information | {} |
| `extendParam.crossBorderLogisticsSelfPickupAddress` | CrossBorderLogisticsSelfPickupAddress | yes | Cross-border consolidation warehouse pickup point address and contact person | {} |
| `extendParam.crossBorderLogisticsSelfPickupAddress.addressText` | String | yes | Consolidation warehouse self-pickup address; province/city/district do not need to be filled in — they follow the default solution's consolidation warehouse | 网商路699号 |
| `extendParam.crossBorderLogisticsSelfPickupAddress.mobile` | String | yes | Consolidation warehouse self-pickup mobile number | 18792983782 |
| `extendParam.crossBorderLogisticsSelfPickupAddress.contractName` | String | yes | Consolidation warehouse self-pickup contact person | 无名 |
| `extendParam.extPairList` | ExtrPair[] | no | List of extended objects | [] |
| `extendParam.extPairList[].key` | String | yes | Key | doorPickup |
| `extendParam.extPairList[].value` | String | yes | Value | {       "serviceCode": "mkd",       "price": 1200, // 人民币，分     } |
| `extendParam.extPairList[].desc` | String | yes | Description | 是否使用上门揽 |
| `extendParam.offerIsvCargoFromList` | KeyValuePair[] | no | Lightweight app closed-loop order placement, offer source info list | [] |
| `extendParam.offerIsvCargoFromList[].key` | java.lang.String | yes | offerId |  |
| `extendParam.offerIsvCargoFromList[].value` | java.lang.String | yes | Offer source information |  |
| `extendParam.selectedTradeServiceList` | TradeService[] | no | Trade service parameters | [] |
| `extendParam.selectedTradeServiceList[].code` | String | yes | Trade service code | vas |
| `extendParam.selectedTradeServiceList[].sourceId` | String | yes | Trade service sourceId | 3C |
| `useOfficialSolution` | Boolean | no | Delivery method uses official logistics pickup | false |
| `useOfficialSolutionModelList` | UseOfficialSolutionModel[] | no | Official logistics pickup solution (required when useOfficialSolution is true) | [] |
| `useOfficialSolutionModelList[].orderGroup` | String | yes | Order group | 123 |
| `useOfficialSolutionModelList[].useOfficialSolutionCode` | String | yes | Official logistics pickup solution code | 123 |
| `fromAgent` | String | no | External agent | qoderwork |

**Response (top level)**: `result` result – Order creation result; `result.totalSuccessAmount` Long – Total order amount (in cents). This field is empty when multiple orders are created at once.; `result.orderId` String – Order ID; this field is empty when multiple orders are created at once; `result.success` Boolean – Whether successful; `result.code` String – Error code; `result.message` String – Error message; `result.accountPeriod` period – Account period information; returns empty for orders not paid via account period; `result.failedOfferList` offer[] – Failed product information; `result.postFee` Long – Freight, unit: cents (fen); when creating multiple orders at once, this field is empty; `result.orderList` BizSimpleOrder[] – Create multiple orders at once; `success` Boolean – Whether successful; `code` String – Error code; `message` String – Error information; …

<a id="alibabatradegetbuyerview-1"></a>
## 46. View order details (buyer view)

`com.alibaba.trade:alibaba.trade.get.buyerView:1` · Orders · 订单详情查看(买家视角)  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.trade.get.buyerView/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Get the details of a single transaction; buyer calls only. Permission must be requested from the Alibaba Open Platform to use this API.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `webSite` | String | yes | Site information, specifies whether the called API belongs to the international site (alibaba) or the 1688 website (1688). | 1688 |
| `orderId` | Long | yes | Order ID of the transaction | 123456 |
| `includeFields` | String | no | Fields included in the query result: GuaranteesTerms - guarantee terms, NativeLogistics - logistics information, RateDetail - rating details, OrderInvoice - invoice information. By default, GuaranteesTerms, NativeLogistics, and OrderInvoice are returned. InvoicingSetting - invoicing settings | GuaranteesTerms,NativeLogistics,RateDetail,OrderInvoice |
| `attributeKeys` | String[] | no | attributeKeys in the vertical table | [] |
| `outOrderId` | String | no | External order ID, for idempotency control | 1556246 |

**Response (top level)**: `result` TradeInfo – Order detail info; `result.baseInfo` OrderBaseInfo – Basic order information; `result.orderBizInfo` bizInfo – Order business info; `result.tradeTerms` TradeTermsInfo[] – Trade terms; `result.productItems` ProductItemInfo[] – Product entry information; `result.nativeLogistics` NativeLogisticsInfo – Domestic logistics; `result.orderInvoiceInfo` OrderInvoiceModel – Invoice information; `result.guaranteesTerms` GuaranteeTermsInfo – Protection terms; `result.orderRateInfo` OrderRateInfo – Order review information; `result.overseasExtraAddress` OverseasExtraAddress – Cross-border address extended information; `result.customs` Customs – Cross-border customs declaration information; `result.quoteList` caigouQuoteInfo[] – Purchase order detail list; a field exclusive to large-enterprise procurement orders.; `result.extAttributes` KeyValuePair[] – Order extended attributes; `result.fromEncryptOrder` Boolean – Whether the order was created with downstream desensitized information; …

<a id="alibabatradegetlogisticsinfosbuyerview-1"></a>
## 47. Get order logistics info (buyer view)

`com.alibaba.logistics:alibaba.trade.getLogisticsInfos.buyerView:1` · Orders · 获取交易订单的物流信息(买家视角)  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.logistics/alibaba.trade.getLogisticsInfos.buyerView/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Requires the order buyer's authorization and returns the logistics details of the buyer's order. In procurement or distribution scenarios, buyers also need logistics details. This interface looks up logistics details by order number, including the sender, the recipient and the details of the goods shipped. Because of when logistics records are entered, API tracking queries may be delayed. Permission must be requested from the Open Platform to access this API.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `orderId` | Long | yes | Order number | 1221434 |
| `fields` | String | no | Fields to be returned. Currently available: company.name, sender, receiver, sendgood. The returned fields must be separated by English commas. | company,name,sender,receiver,sendgood |
| `webSite` | String | yes | Whether it is a 1688 business or an icbu business | 1688或者alibaba |

**Response (top level)**: `result` OpenPlatformLogisticsOrder[] – Return result; `errorCode` String – Error code; `errorMessage` String – Error message; `success` Boolean – Whether successful; …

<a id="alibabatradegetlogisticstraceinfobuyerview-1"></a>
## 48. Get order logistics tracking info (buyer view)

`com.alibaba.logistics:alibaba.trade.getLogisticsTraceInfo.buyerView:1` · Orders · 获取交易订单的物流跟踪信息(买家视角)  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.logistics/alibaba.trade.getLogisticsTraceInfo.buyerView/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Requires the order buyer's authorization and returns the logistics tracking information of the buyer's order. In procurement or distribution scenarios, buyers also need logistics details. This interface looks up tracking information by logistics (waybill) number. Because of when logistics records are entered, API tracking queries may be delayed. Permission must be requested from the Open Platform to access this API.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `logisticsId` | String | no | The logistics number under this order | AL8234243 |
| `orderId` | Long | yes | Order number | 13342343 |
| `webSite` | String | yes | Whether it is a 1688 business or an icbu business | 1688或者alibaba |

**Response (top level)**: `logisticsTrace` OpenPlatformLogisticsTrace[] – Tracking order details; `errorCode` String – Error code; `errorMessage` String – Error description; `success` Boolean – Whether successful; `crossPackageFulfillmentDTO` CrossPackageFulfillmentDTO[] – Cross-border overseas shipment tracking; …

<a id="alibabatradecancel-1"></a>
## 49. Cancel transaction

`com.alibaba.trade:alibaba.trade.cancel:1` · Orders · 取消交易  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.trade.cancel/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Buyer or seller cancels a transaction. Only transactions in specific statuses can be cancelled; on 1688 this is used to cancel unpaid orders. If an order is closed less than 10 seconds after it was created, the error CLOSE_ORDER_TOO_FAST is returned.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `webSite` | String | yes | Site information, specifies whether the called API belongs to the international site (alibaba) or the 1688 website (1688). | 1688 |
| `tradeID` | Long | yes | Transaction ID, order number | 123456 |
| `cancelReason` | String | yes | Reason description; buyerCancel: buyer cancelled the order; sellerGoodsLack: seller out of stock; other: other | other |
| `remark` | String | no | Remark | 备注 |

**Response (top level)**: `success` Boolean – Whether processed successfully: true for success, false for failure; see error for the failure reason; `errorCode` String – Error code; `errorMessage` String – Error message

<a id="alibabatradegetbuyerorderlist-1"></a>
## 50. View order list (buyer view)

`com.alibaba.trade:alibaba.trade.getBuyerOrderList:1` · Orders · 订单列表查看(买家视角)  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.trade.getBuyerOrderList/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Get the buyer's order list; the user's memberId must equal the buyer memberId on the orders. This interface returns only basic order information, not logistics or invoice information. For logistics information call the order-detail interface; for invoice information call the invoice-information API.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `bizTypes` | String[] | no | Business type, supports: &quot;cn&quot; (regular order type), &quot;ws&quot; (large-value wholesale order type), &quot;yp&quot; (regular sample order type), &quot;yf&quot; (one-cent sample order type), &quot;fs&quot; (flash sale (limited-time discount) order type), &quot;cz&quot; (processing/customization order type), &quot;ag&quot; (agreement procurement order type), &quot;hp&quot; (group-buy order type), &quot;gc&quot; (national procurement order type), &quot;supply&quot; (supply-and-marketing order type), &quot;nyg&quot; (nyg order type), &quot;factory&quot; (Taobao Factory order type), &quot;quick&quot; (quick order placement), &quot;xiangpin&quot; (Xiangpin order), &quot;nest&quot; (Procurement Mall - Nest), &quot;f2f&quot; (face-to-face payment), &quot;cyfw&quot; (sample storage service), &quot;sp&quot; (consignment order flag), &quot;wg&quot; (WeiGong order), &quot;factorysamp&quot; (Taobao Factory sampling order), &quot;factorybig&quot; (Taobao Factory bulk order) | ["cn","ws"] |
| `createEndTime` | java.util.Date | no | Order placement end time | 20180802211113000+0800 |
| `createStartTime` | java.util.Date | no | Order start time | 20180102211113000+0800 |
| `isHis` | boolean | no | Whether to query the historical order table; default queries the current table, i.e. the default value is false | false |
| `modifyEndTime` | java.util.Date | no | End of the modification time query range | 20180802211113000+0800 |
| `modifyStartTime` | java.util.Date | no | Query modification time start | 20180102211113000+0800 |
| `orderStatus` | java.lang.String | no | Order status, values are success, cancel (transaction cancelled, penalty and other settlements completed), waitbuyerpay (waiting for seller to pay), waitsellersend (waiting for seller to ship), waitbuyerreceive (waiting for buyer to receive goods) | success |
| `page` | int | no | Query page number, starting from 1 | 1 |
| `pageSize` | int | no | Number of items per page for the query | 20 |
| `refundStatus` | java.lang.String | no | Refund status, supported values: &quot;waitselleragree&quot; (waiting for seller to agree), &quot;refundsuccess&quot; (refund successful), &quot;refundclose&quot; (refund closed), &quot;waitbuyermodify&quot; (waiting for buyer to modify), &quot;waitbuyersend&quot; (waiting for buyer to return the goods), &quot;waitsellerreceive&quot; (waiting for seller to confirm receipt) | refundsuccess |
| `sellerMemberId` | java.lang.String | no | Seller memberId | b2b-1624961198 |
| `sellerLoginId` | String | no | Seller loginId | alitestforisv02 |
| `sellerRateStatus` | java.lang.Integer | no | Seller rating status (4: rated, 5: not rated, 6: rating not required) | 6 |
| `tradeType` | java.lang.String | no | Transaction type: Escrow transaction (1), Prepaid deposit transaction (2), ETC overseas acquiring transaction (3), Instant payment transaction (4), Guarantee fund security transaction (5), Unified transaction process (6), Staged transaction (7), Cash on delivery transaction (8), Credit voucher payment transaction (9), Account period payment transaction (10), 1688 Transaction 4.0, new staged transaction (50060), Face-to-face payment transaction process (50070), Service-type transaction process (50080) | 50060 |
| `productName` | java.lang.String | no | Product name | 测试商品 |
| `needBuyerAddressAndPhone` | java.lang.Boolean | no | Whether the buyer's detailed address information and phone number need to be queried | false |
| `needMemoInfo` | java.lang.Boolean | no | Whether remark information needs to be queried | false |
| `outOrderId` | String | no | External order number, can be used for idempotency control | 90187872898371 |
| `needInvoicingSetting` | Boolean | no | Invoicing settings need to be queried | true |
| `orderIds` | Long[] | no | Query multiple order IDs | [90187872898371,90187872898372] |

**Response (top level)**: `result` TradeInfo[] – Query return list; `errorCode` String – Error code; `errorMessage` String – Error message; `totalRecord` Long – Total record count; …

<a id="alibabacreateorderpreview-1"></a>
## 51. Preview data before creating an order

`com.alibaba.trade:alibaba.createOrder.preview:1` · Orders · 创建订单前预览数据接口  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.createOrder.preview/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Orders may only contain products from a single supplier. This interface returns discount and related information for order creation. It 1. validates whether the products may be ordered; 2. validates the consignment (distribution) relationship; 3. validates stock, minimum order quantity and whether mixed-batch conditions are met.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `addressParam` | address | yes | Shipping address info | {"address":"网商路699号","phone": "0517-88990077","mobile": "15251667788","fullName": "张三","po… |
| `addressParam.addressId` | Long | yes | Shipping address ID | 1234 |
| `addressParam.fullName` | String | yes | Recipient name | 张三 |
| `addressParam.mobile` | String | yes | Mobile phone | 15251667788 |
| `addressParam.phone` | String | yes | Phone number | 0517-88990077 |
| `addressParam.postCode` | String | yes | Postal code | 000000 |
| `addressParam.cityText` | String | yes | City text | 杭州市 |
| `addressParam.provinceText` | String | yes | Province text | 浙江省 |
| `addressParam.areaText` | String | yes | District text | 滨江区 |
| `addressParam.townText` | String | yes | Town text | 长河镇 |
| `addressParam.address` | String | yes | Street address | 网商路699号 |
| `addressParam.districtCode` | String | yes | Address code | 310107 |
| `cargoParamList` | cargo[] | yes | Product information | [{"specId": "b266e0726506185beaf205cbae88530d","quantity": 5,"offerId": 554456348334},{"sp… |
| `cargoParamList[].offerId` | Long | yes | offer id corresponding to the product | 554456348334 |
| `cargoParamList[].specId` | String | yes | Product SKU ID | b266e0726506185beaf205cbae88530d |
| `cargoParamList[].quantity` | Double | yes | Product quantity (used for amount calculation) | 5 |
| `cargoParamList[].openOfferId` | String | no | Encrypted offerId. When the search result returns only openOfferId, use openOfferId in place of offerId to place the order. | Wcv2w970KoL1BCpJGQp3vwKvcBThOPlpHnUtkK3T0n4= |
| `cargoParamList[].outMemberId` | String | no | External downstream member ID | 98928912-23 |
| `cargoParamList[].bizFenXiaoOutSubOrderInfo` | BizFenXiaoOutSubOrderInfo[] | no | Downstream sub-order data for distribution | {} |
| `cargoParamList[].bizFenXiaoOutSubOrderInfo[].outSubOrderId` | String | no | External order id | 123 |
| `cargoParamList[].bizFenXiaoOutSubOrderInfo[].outSkuId` | String | no | External sku | 123 |
| `cargoParamList[].bizFenXiaoOutSubOrderInfo[].outSubGmv` | Double | no | gmv | 123 |
| `cargoParamList[].bizFenXiaoOutSubOrderInfo[].outSubQty` | Double | no | Order volume | 123 |
| `cargoParamList[].bizFenXiaoOutSubOrderInfo[].outSubOrderPayTime` | String | no | Order placement time of the downstream sub-order | 2025-03-18 09:20:21 |
| `cargoParamList[].bizFenXiaoOutSubOrderInfo[].outSubOrderLatestDeliveryTime` | String | no | Latest shipping time for the downstream main order | 2025-03-18 09:20:21 |
| `cargoParamList[].bizFenXiaoOutSubOrderInfo[].outSubGmvDiscount` | Double | no | Discount amount of the downstream sub-order | 10 |
| `cargoParamList[].bizFenXiaoOutSubOrderInfo[].outSubGmvPost` | Double | no | Shipping fee for the downstream sub-order | 5 |
| `cargoParamList[].bizFenXiaoOutSubOrderInfo[].outSubGmvReceive` | Double | no | Actual amount received for the downstream sub-order | 118 |
| `invoiceParam` | invoice | no | Invoice information | {"invoiceType":0,"cityText": "杭州市","provinceText": "浙江省","address": "网商路699号","phone": "05… |
| `invoiceParam.invoiceType` | Integer | yes | Invoice type 0: regular invoice, 1: VAT invoice | 0 |
| `invoiceParam.provinceText` | String | yes | Province text | 浙江省 |
| `invoiceParam.cityText` | String | yes | City text | 杭州市 |
| `invoiceParam.areaText` | String | yes | Region text | 滨江区 |
| `invoiceParam.townText` | String | yes | Town text | 长河镇 |
| `invoiceParam.postCode` | String | yes | Postal code | 333333 |
| `invoiceParam.address` | String | yes | Street | 网商路699号 |
| `invoiceParam.fullName` | String | yes | Invoice recipient's name | 张三 |
| `invoiceParam.phone` | String | yes | Phone number | 0517-88990077 |
| `invoiceParam.mobile` | String | yes | Mobile phone | 15251667788 |
| `invoiceParam.companyName` | String | yes | Purchasing company name (invoice title) | 测试公司 |
| `invoiceParam.taxpayerIdentifier` | String | yes | Tax identification number | 12345 |
| `invoiceParam.bankAndAccount` | String | yes | Bank name and account number | 网商银行 |
| `invoiceParam.localInvoiceId` | String | yes | VAT local invoice number | 123123123 |
| `flow` | String | no | general (create open marketplace order), fenxiao (create distribution order), paired (Tiantian Temai / Daily Deals), repurchase (repurchase contract order), saleproxy process will verify the distribution relationship, boutiquefenxiao (curated supply distribution-price order, free shipping when purchase quantity is 1), boutiquepifa (curated supply wholesale-price order, used when purchase quantity is greater than 2). If flow is empty, the system compares prices and previews the best option, returning the optimal order flow. Repurchase contract orders are not included in the price comparison. All order flow channels may cause differences in price and buyer protection services. | general |
| `instanceId` | String | no | Wholesale group-buy instanceId, obtained from alibaba.pifatuan.product.list | 4063139_1662080400000 |
| `encryptOutOrderInfo` | EncryptOutOrderInfo | no | Downstream encrypted order information, used for downstream label printing | {} |
| `encryptOutOrderInfo.encryptOrder` | Boolean | yes | Whether the order is encrypted | true |
| `encryptOutOrderInfo.outPlatformOrderNo` | String | yes | Downstream platform order number | 12365452354551 |
| `encryptOutOrderInfo.outPlatformSupplyOrderNo` | String | no | Downstream platform supply chain purchase order | CT7403712218870825259 |
| `encryptOutOrderInfo.outPlatformCode` | String | yes | Taobao-thyny, Tmall-tm, Taote-taote, Alibaba C2M-c2m, JD.com-jingdong, Pinduoduo-pinduoduo, WeChat Store-weixin, Cross-border-kuajing, Kuaishou-kuaishou, Youzan-youzan, Douyin-douyin, Secoo-siku, Meituan Tuanhaohuo-meituan, Xiaohongshu-xiaohongshu, Dangdang-dangdang, Suning-suning, DaVdian-davdian, Xingyun-xingyun, Miya-miya, Boluopai Mall-boluo, Other-other | taote |
| `encryptOutOrderInfo.outPlatformAppkey` | String | yes | The appkey used by the downstream platform to retrieve orders | 32154 |
| `encryptOutOrderInfo.outShopId` | String | no | Downstream platform shop Id | 1879283 |
| `encryptOutOrderInfo.outShopName` | String | no | Downstream platform shop name | 三生科技 |
| `encryptOutOrderInfo.outOriginAddress` | OutAddress | no | External original address information | {} |
| `encryptOutOrderInfo.outOriginAddress.province` | Place | yes | Province | {"name":"四川省","code":"51000"} |
| `encryptOutOrderInfo.outOriginAddress.province.code` | String | yes | Address code | 511300 |
| `encryptOutOrderInfo.outOriginAddress.province.name` | String | yes | Address name | 南充 |
| `encryptOutOrderInfo.outOriginAddress.city` | Place | yes | City | {} |
| `encryptOutOrderInfo.outOriginAddress.city.code` | String | yes | Address code | 511300 |
| `encryptOutOrderInfo.outOriginAddress.city.name` | String | yes | Address name | 南充 |
| `encryptOutOrderInfo.outOriginAddress.area` | Place | yes | District | {} |
| `encryptOutOrderInfo.outOriginAddress.area.code` | String | yes | Address code | 511300 |
| `encryptOutOrderInfo.outOriginAddress.area.name` | String | yes | Address name | 南充 |
| `encryptOutOrderInfo.outOriginAddress.town` | Place | no | Town/street | {} |
| `encryptOutOrderInfo.outOriginAddress.town.code` | String | yes | Address code | 511300 |
| `encryptOutOrderInfo.outOriginAddress.town.name` | String | yes | Address name | 南充 |
| `encryptOutOrderInfo.outOriginAddress.address` | String | no | Detailed address | 网商路699号 |
| `encryptOutOrderInfo.outOriginAddress.postCode` | String | no | Postal code | 511304 |
| `encryptOutOrderInfo.oaid` | String | no | Taobao oaid | 265646-52342354-2354Akf-w3654SF |
| `encryptOutOrderInfo.outPatformExtraInfo` | String | no | Other extended information from the downstream platform | {} |
| `encryptOutOrderInfo.encryptReceiverName` | String | no | Downstream encrypted recipient name | *** |
| `encryptOutOrderInfo.encryptReceiverMobile` | String | no | Downstream encrypted recipient phone number | *** |
| `encryptOutOrderInfo.encryptReceiverAddress` | String | no | Downstream encrypted recipient address | *** |
| `encryptOutOrderInfo.outPlatformSubCode` | String | no | Downstream channel sub-business code, e.g. Douyin sub-channel 101, used to identify supply chain orders | 101 |
| `proxySettleRecordId` | String | no | Purchase order ID for split-payment regular order placement, where the transaction flow is "proxy" | 4051300002 |
| `inventoryMode` | String | no | Stock mode, jit (JIT mode) or cang (warehouse shipping mode); currently only available for AE | jit |
| `outOrderId` | String | no | External order number | 988129883123 |
| `pickupService` | String | no | Door-to-door pickup; currently available for AE supply, not yet enabled for other scenarios | y或n,默认为n |
| `crossBorderLogisticsSolutionId` | String | no | The sourceId of the official cross-border logistics solution selected by the user | GLOBAL_CAINIAO_VN_TRANSIT_LAND |
| `useBorderLogisticsSolution` | Boolean | no | Whether to use the cross-border logistics solution preview. Defaults to false. This parameter does not take effect for overseas addresses. It takes effect for Hong Kong, Macao, and Taiwan addresses. | true |
| `isvBizType` | String | no | Open platform business code | cross |

**Response (top level)**: `orderPreviewResuslt` model[] – Order preview result; if automatic order splitting occurs, multiple records will be returned; `success` Boolean – Whether successful; `errorCode` String – Error code; `errorMsg` String – Error message; `postFeeByDescOfferList` Long[] – Product list for the freight description; `consignOfferList` Long[] – Consignment product list; `unsupportedCrossBorderPayOfferList` Long[] – List of products that do not support Kuajingbao payment; `extPairList` ExtrPair[] – List of extended objects; …

<a id="alibabaordermemoadd-1"></a>
## 52. Update order memo

`com.alibaba.trade:alibaba.order.memoAdd:1` · Orders · 修改订单备忘  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.order.memoAdd/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
If the authorized user is the seller, updates the seller memo; if the buyer, updates the buyer memo. Note: this interface can be called repeatedly, and the memo overwrites the content of the previous call.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `orderId` | java.lang.Long | yes | Order ID | 1234567 |
| `memo` | java.lang.String | yes | Memo info | 订单备忘详情 |
| `remarkIcon` | String | yes | Memo icon, currently only supports numbers. 1 is red icon, 2 is blue icon, 3 is green icon, 4 is yellow icon | 2 |

**Response (top level)**: `success` java.lang.Boolean – Whether successful; `errorCode` java.lang.String – Error code; `errorMsg` java.lang.String – Error message

<a id="alibabasubaccountlist-1"></a>
## 53. Get sub-account list

`com.alibaba.account:alibaba.subAccount.list:1` · Orders · 获取子账号列表  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.account/alibaba.subAccount.list/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Get the user's main-account and sub-account information. If the API is authorized as a sub-account, only the main account that the sub-account belongs to is returned. If authorized as a main account, the list of all sub-accounts is returned.

**Request body**: no application parameters.

**Response (top level)**: `mainUserId` String – Main account Userid; `mainLoginId` String – Main account loginId; `mainMemberId` String – Main account MemberId; `subAccountList` simpleAccountInfo[] – Sub-account list; `errorMsg` String – Error description; `errorCode` String – Error code; …

<a id="alibabatradeaddfeedback-1"></a>
## 54. Buyer adds an order message

`com.alibaba.trade:alibaba.trade.addFeedback:1` · Orders · 买家补充订单留言接口  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.trade.addFeedback/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Buyer adds a supplementary message to an order. The total message length must not exceed 500 characters.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `tradeFeedbackParam` | TradeFeedbackParam | yes | Request parameters | {"feedback":"test","orderId":"123123213"} |
| `tradeFeedbackParam.feedback` | java.lang.String | yes | Message/comment | 留言 |
| `tradeFeedbackParam.orderId` | java.lang.String | yes | Order ID | 12344444555545 |

**Response (top level)**: `result` TradeFeedbackResult – Return result; `result.errorInfo` java.lang.String – Error description; `result.errorCode` java.lang.String – Error code; `result.success` Boolean – Whether successful; `code` String – Error code; `message` String – Error description; `success` Boolean – Whether successful

<a id="tradereceivegoodsconfirm-1"></a>
## 55. Buyer confirms receipt

`com.alibaba.trade:trade.receivegoods.confirm:1` · Orders · 买家确认收货  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/trade.receivegoods.confirm/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Buyer confirms receipt of goods.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `orderId` | java.lang.Long | yes | Order ID | 56623232655125698 |
| `orderEntryIds` | java.lang.Long[] | yes | Sub-order ID | 562356635566365512 |

**Response (top level)**: `result` TradeConfirmReceiptResult – ; `result.success` boolean – ; `result.errorInfo` java.lang.String – ; `result.errorCode` java.lang.String – 

<a id="tradeorderbuyerdelete-1"></a>
## 56. Buyer deletes a closed order

`com.alibaba.trade:trade.order.buyerdelete:1` · Orders · 买家删除已关闭的订单  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/trade.order.buyerdelete/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Buyer deletes an order that has been closed.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `orderId` | java.lang.Long | yes | Order id | 12213412341 |

**Response (top level)**: `result` ResultModel – Return value; `result.success` java.lang.Boolean – Whether successful; `result.code` java.lang.String – Error code; `result.message` java.lang.String – Error message; `result.result` java.lang.Boolean – Whether the business operation was successful

<a id="orderreceiveaddressbuyerupdate-1"></a>
## 57. Buyer requests a shipping-address change

`com.alibaba.trade:order.receiveAddress.buyerUpdate:1` · Orders · 买家申请修改收货地址  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/order.receiveAddress.buyerUpdate/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Buyer changes the shipping address. If the new address is in a remote area or the shipping fee must be recalculated, the seller must confirm the change and may reject it.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `param` | BuyerDeliveryAddressModifyParam | yes | Modify address info object | {} |
| `param.orderId` | java.lang.Long | yes | Order id | 1231231123123 |
| `param.receiveAddress` | AddressParam | yes | Modify address | {} |
| `param.receiveAddress.addressId` | java.lang.Long | yes |  |  |
| `param.receiveAddress.fullName` | java.lang.String | yes |  |  |
| `param.receiveAddress.mobile` | java.lang.String | yes |  |  |
| `param.receiveAddress.phone` | java.lang.String | yes |  |  |
| `param.receiveAddress.postCode` | java.lang.String | yes |  |  |
| `param.receiveAddress.cityText` | java.lang.String | yes |  |  |
| `param.receiveAddress.cityCode` | java.lang.String | yes |  |  |
| `param.receiveAddress.provinceText` | java.lang.String | yes |  |  |
| `param.receiveAddress.provinceCode` | java.lang.String | yes |  |  |
| `param.receiveAddress.areaText` | java.lang.String | yes |  |  |
| `param.receiveAddress.areaCode` | java.lang.String | yes |  |  |
| `param.receiveAddress.townText` | java.lang.String | yes |  |  |
| `param.receiveAddress.townCode` | java.lang.String | yes |  |  |
| `param.receiveAddress.address` | java.lang.String | yes |  |  |
| `param.receiveAddress.districtCode` | java.lang.String | yes |  |  |

**Response (top level)**: `result` ResultModel – Return value; `result.success` boolean – Whether the gateway interface succeeded; `result.code` java.lang.String – Error code; `result.message` java.lang.String – Error message; `result.result` Boolean – Whether the address modification request was successful

<a id="alibabaalipayurlget-1"></a>
## 58. Batch-get payment links for orders

`com.alibaba.trade:alibaba.alipay.url.get:1` · Payment · 批量获取订单的支付链接  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.alipay.url.get/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
When paying through an ERP, use this API to get a cashier link for batch payment. A single order returns the 1688 cashier URL; multiple orders return the Alipay cashier URL. The ERP can redirect the user to the cashier link to complete payment. The user's 1688 login status is verified before payment.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `orderIdList` | Long[] | yes | Order ID list. Up to 100 orders per batch; for Kuajingbao, batches support a maximum of 30 orders. | [74321349391498520] |

**Response (top level)**: `erroMsg` String – Error message; `payUrl` String – Payment link; `success` Boolean – Whether successful; may be partially successful, needs to be checked together with payFailureOrderList; `errorCode` String – Error code; `payFailureOrderList` Long[] – Partially failed order id

<a id="alibabacrossborderpayurlget-1"></a>
## 59. Get payment link for Cross-Border Pay (Kuajingbao)

`com.alibaba.trade:alibaba.crossBorderPay.url.get:1` · Payment · 获取使用跨境宝支付的支付链接  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.crossBorderPay.url.get/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Get a payment link for paying with Cross-Border Pay (Kuajingbao).

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `orderIdList` | Long[] | yes | Order Id list, maximum 30 orders per batch; too many orders will cause timeout, 10 orders at a time is recommended | [111111,22222333] |

**Response (top level)**: `success` String – Whether successful; `errorCode` String – Error code; `errorMsg` String – Error description; `payUrl` String – Cashier payment link; `cantPayOrderList` Long[] – List of orders that cannot be paid in batch due to credit limit or risk control reasons

<a id="alibabacreditpayurlget-1"></a>
## 60. Get payment link for Cheng-e-She credit pay

`com.alibaba.trade:alibaba.creditPay.url.get:1` · Payment · 获取使用诚e赊支付的支付链接  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.creditPay.url.get/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Get a payment link for paying with Cheng-e-She (buy-now-pay-later credit).

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `orderIdList` | Long[] | yes | Order Id list, maximum 30 orders per batch; too many orders will cause timeout, 10 orders at a time is recommended | [111111,22222333] |

**Response (top level)**: `success` String – Whether successful; `errorCode` String – Error code; `errorMsg` String – Error description; `payUrl` String – Cashier payment link; `cantPayOrderList` Long[] – List of orders that cannot be paid in batch due to credit limit or risk control reasons

<a id="alibabaaccountperiodlistbuyerview-1"></a>
## 61. Buyer views all granted credit terms

`com.alibaba.trade:alibaba.accountPeriod.list.buyerView:1` · Payment · 买家查看获得的所有账期授信  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.accountPeriod.list.buyerView/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
View, from the buyer's side, all account-period (credit-term) lines the buyer has been granted. Paginated; at most 10 records per call.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `pageIndex` | Long | no | Page number | 1 |
| `sellerLoginId` | String | no | Seller ID; if not filled in, all are queried | alitestforisv01 |

**Response (top level)**: `success` Boolean – Whether successful; `errorCode` String – Error code; `errorMsg` String – Error message; `resultList` result – Returned data result; `resultList.totalCount` String – Total number of records; `resultList.accountPeriodList` AccountPeriodInfo[] – Credit list; …

<a id="alibabatradepaywayquery-1"></a>
## 62. Query payment channels supported by an order

`com.alibaba.trade:alibaba.trade.payWay.query:1` · Payment · 查询订单可以支持的支付渠道  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.trade.payWay.query/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Query the payment methods or channels available for an unpaid order.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `orderId` | String | yes | Order number | 123123 |

**Response (top level)**: `success` String – Whether successful; `errorCode` String – Error code; `errorMsg` String – Error message; `resultList` TradePayTypeResult – Return result; `resultList.channels` PayTypeInfo[] – List of available payment channels; `resultList.orderId` java.lang.String – Order number; `resultList.payFee` java.lang.Long – Payment amount, in cents; `resultList.timeout` java.lang.String – Latest payment time; …

<a id="alibabatradepayprotocolpayisopen-1"></a>
## 63. Check whether password-free payment is enabled

`com.alibaba.trade:alibaba.trade.pay.protocolPay.isopen:1` · Payment · 查询是否开通免密支付  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.trade.pay.protocolPay.isopen/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Check whether an auto-debit (withholding) agreement is enabled.

**Request body**: no application parameters.

**Response (top level)**: `result` TradeWithholdStatusResultModel – Signing status return value, indicating the auto-debit signing status for Alipay and Cheng-e-She. payChannel=SHEGOU and signedStatus=true means Cheng-e-She auto-debit has been signed; payChannel=ALIPAY and signedStatus=true means Alipay auto-debit has been signed. If either signedStatus=true, the auto-debit interface can be called to complete the auto-debit.; `result.success` Boolean – Whether successful; `result.code` String – Error code; `result.message` String – Error message; `result.result` TradeWithholdStatusResult – Contract signing status; …

<a id="alibabatradepayprotocolpaypreparepay-1"></a>
## 64. Initiate password-free payment

`com.alibaba.trade:alibaba.trade.pay.protocolPay.preparePay:1` · Payment · 发起免密支付  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.trade.pay.protocolPay.preparePay/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Initiates a password-free payment. Automatically detects whether Alipay or Cheng-e-She password-free payment is enabled and initiates the debit. Cheng-e-She auto-debit is tried first; if it fails, Alipay auto-debit is attempted. The error codes returned by this interface are currently not detailed; after a failed debit, retry up to 3 times rather than indefinitely.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `tradeWithholdPreparePayParam` | TradeWithholdPreparePayParam | yes | Initiate password-free payment |  |
| `tradeWithholdPreparePayParam.orderId` | java.lang.Long | yes | Order ID | 1938489823 |
| `tradeWithholdPreparePayParam.payChannel` | String | no | Pass shegou for Buy-Now-Pay-Later payment, and alipay for Alipay payment. If no value is passed, the deduction defaults to using the priority described in the API documentation. | alipay |
| `tradeWithholdPreparePayParam.payAmount` | Long | no | Total payment amount, in cents | 123 |
| `tradeWithholdPreparePayParam.opRequestId` | String | no | requestid | 134134134 |

**Response (top level)**: `result` rresult – Password-free payment result; `result.success` Boolean – Whether successful; `result.code` String – Error code; `result.message` String – Error message; `result.result` mresult – Deduction return value; …

<a id="tradeorderpayanalysis-1"></a>
## 65. Order payment consultation

`com.alibaba.trade:trade.orderpay.analysis:1` · Payment · 交易订单支付咨询  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/trade.orderpay.analysis/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Order-payment consultation interface, used to analyse which payment method an order uses, and so on.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `orderIds` | java.lang.Long[] | yes | Order list | [3535260697417660107] |
| `payChannel` | java.lang.String | yes | Payment channel  alipay (Alipay), shegou (Cheng-e-She), kjpayV2 (Kuajingbao) | kjpayV2 |

**Response (top level)**: `result` ResultModel – ; `result.success` boolean – ; `result.code` java.lang.String – ; `result.message` java.lang.String – ; `result.result` PayAnalysisResult – ; …

<a id="alibabatradegrouppayurlget-1"></a>
## 66. Get combined-cashier URL

`com.alibaba.trade:alibaba.trade.grouppay.url.get:1` · Payment · 组合收银台url获取  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.trade.grouppay.url.get/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Get the combined-cashier URL.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `orderIds` | Long[] | yes | Order list | [123123413,1223234] |
| `payPlatformType` | String | no | PC or WIRELESS | PC |

**Response (top level)**: `results` TradeCreateGroupPayUrlResult – ; `results.payUrl` java.lang.String – Return payurl; `results.success` boolean – Whether successful; `results.errorInfo` java.lang.String – Error message; `results.errorCode` java.lang.String – Error code

<a id="productfreightestimate-1"></a>
## 67. Estimate domestic (China) shipping fee for a product

`com.alibaba.fenxiao.crossborder:product.freight.estimate:1` · Logistics · 商品中国国内运费预估  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/product.freight.estimate/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Estimate a product's shipping fee from the product ID and the province/city/district codes of a delivery address within mainland China.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `productFreightQueryParamsNew` | ProductFreightQueryParamsNew | yes | Input parameter | 如下 |
| `productFreightQueryParamsNew.offerId` | Long | yes | Product ID | 111111111 |
| `productFreightQueryParamsNew.toProvinceCode` | String | yes | China province code | 如浙江省330000 |
| `productFreightQueryParamsNew.toCityCode` | String | yes | China city code | 如杭州市330100 |
| `productFreightQueryParamsNew.toCountryCode` | String | yes | China region code | 如滨江区330108 |
| `productFreightQueryParamsNew.totalNum` | Long | yes | Purchase quantity | 3 |
| `productFreightQueryParamsNew.logisticsSkuNumModels` | LogisticsSkuNumModel[] | yes | SKU quantity | 1 |
| `productFreightQueryParamsNew.logisticsSkuNumModels[].skuId` | String | yes | skuId | 12345 |
| `productFreightQueryParamsNew.logisticsSkuNumModels[].number` | Long | yes | Quantity | 1 |

**Response (top level)**: `result` ResultModel – ; `result.success` boolean – Result; `result.code` java.lang.String – Error code; `result.message` java.lang.String – Error description; `result.result` ProductFreightModel – Internal result; …

<a id="logisticsordergetoutorderid-1"></a>
## 68. Query external order ID by waybill number or unclaimed-parcel code

`com.alibaba.fenxiao.crossborder:logistics.order.getOutOrderId:1` · Logistics · 根据运单号或无主件码查询外部订单ID  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/logistics.order.getOutOrderId/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Query the external order ID by waybill number or unclaimed-parcel code.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `shipmentId` | java.lang.String | no | Waybill number | 1 |
| `noMainPartCode` | java.lang.String | no | Unclaimed item code | 1 |

**Response (top level)**: `result` ResultModel – Result; `result.success` java.lang.Boolean – Success; `result.msg` java.lang.String – Info; `result.model` ScOutOrderIdQueryModel – Return value; …

<a id="shippinginsuranceget-1"></a>
## 69. Query shipping-insurance info

`com.alibaba.trade:shipping.insurance.get:1` · Logistics · 运费险信息查询  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/shipping.insurance.get/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Query shipping-insurance information.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `orderId` | Long | yes | Order number | 订单号 |
| `type` | String | yes | Shipping insurance type | givenByPlatform平台赠送，givenByMerchant商家赠送 |

**Response (top level)**: `result` ResultModel – ; `result.success` java.lang.Boolean – Whether successful; `result.code` java.lang.String – Response code; `result.message` java.lang.String – Response information; `result.result` TradeFreightPolicyResult – Return result; …

<a id="alibabalogisticsmyfreighttemplatelistget-1"></a>
## 70. Get shipping-template details

`com.alibaba.logistics:alibaba.logistics.myFreightTemplate.list.get:1` · Logistics · 获取物流模板详情  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.logistics/alibaba.logistics.myFreightTemplate.list.get/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Get the seller's shipping template by template ID. Template ID 0 means "shipping fee to be explained"; 1 means the seller bears the shipping fee.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `templateId` | java.lang.Long | no | Template id, used for single-record query scenarios | xxx |
| `querySubTemplate` | java.lang.Boolean | no | Whether to query the sub-template | false |
| `queryRate` | java.lang.Boolean | no | Whether to query the sub-template rate | false |

**Response (top level)**: `result` FreightTemplate[] – Return result; `errorCode` java.lang.String – Error code; `errorMsg` java.lang.String – Error description; …

<a id="alibabatradereceiveaddressget-1"></a>
## 71. Buyer gets saved shipping addresses

`com.alibaba.trade:alibaba.trade.receiveAddress.get:1` · Logistics · 买家获取保存的收货地址信息列表  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.trade.receiveAddress.get/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Get the buyer's list of saved shipping addresses.

**Request body**: no application parameters.

**Response (top level)**: `result` ReceiveAddressResult – Return result; `result.receiveAddressItems` ReceiveAddressItem[] – List of shipping addresses; `success` Boolean – Whether successful; `code` String – Error code; `message` String – Error message; …

<a id="alibabatradeaddresscodeparse-1"></a>
## 72. Parse an address into area codes

`com.alibaba.trade:alibaba.trade.addresscode.parse:1` · Logistics · 根据地址解析地区码  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.trade.addresscode.parse/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Parse area codes from address information.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `addressInfo` | String | yes | Address information | 浙江省 杭州市 滨江区网商路699号 |

**Response (top level)**: `result` ReceiveAddress – Parsed shipping address; `result.address` java.lang.String – Street address, excluding province/city codes; `result.addressCode` java.lang.String – Address region code; `result.addressCodeText` java.lang.String – The text corresponding to the address region code (including country, province, city); `result.addressId` java.lang.Long – addressId; `result.bizType` java.lang.String – Business type recorded for the shipping address; `result.isDefault` boolean – Whether it is the default; `result.fullName` java.lang.String – Recipient name; `result.latest` boolean – Whether it is the last-selected shipping address; `result.mobile` java.lang.String – Mobile phone number; `result.phone` java.lang.String – Phone number; `result.postCode` java.lang.String – Postal code; `errorCode` String – Error code; `errorMessage` String – Error message

<a id="alibabatradeopquerymarketingmixconfig-1"></a>
## 73. Query seller mixed-batch settings

`com.alibaba.trade:alibaba.trade.OpQueryMarketingMixConfig:1` · Logistics · 查询卖家混批设置  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.trade.OpQueryMarketingMixConfig/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Query the seller's mixed-batch (mixed wholesale) settings.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `sellerMemberId` | String | no | Seller memberId | b2b-1623492085 |
| `sellerLoginId` | String | no | Seller LoginId. When sellerMemberId is empty, loginId takes precedence. | alitestforisv01 |

**Response (top level)**: `result` OpMarketingMixConfigModel – Return result; `result.generalHunpi` boolean – Whether it is a regular mixed batch; `result.gmtCreate` java.util.Date – Creation time; `result.gmtModified` java.util.Date – Modification time; `result.memberId` java.lang.String – Seller memberID; `result.mixAmount` java.lang.Integer – Mixed batch amount; `result.mixNumber` java.lang.Integer – Mixed batch quantity; `errorCode` String – Error code; `errorMessage` String – Error message; `extErrorMessage` String – Error info extension; `success` Boolean – Whether successful

<a id="alibabalogisticsopquerylogisticcompanylist-1"></a>
## 74. Logistics company list (all companies)

`com.alibaba.logistics:alibaba.logistics.OpQueryLogisticCompanyList:1` · Logistics · 物流公司列表-所有的物流公司  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.logistics/alibaba.logistics.OpQueryLogisticCompanyList/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Get the names of all logistics companies.

**Request body**: no application parameters.

**Response (top level)**: `result` OpLogisticsCompanyModel[] – List of logistics companies; `success` Boolean – Whether successful; `errorCode` String – Error code; `errorMessage` String – Error code description; `extErrorMessage` String – Extended error code description; …

<a id="tradeaddressparsetext-1"></a>
## 75. Parse overseas address

`com.alibaba.trade:trade.address.parseText:1` · Logistics · 海外地址解析  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/trade.address.parseText/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Parse an overseas address.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `param` | OrderAddressParseParam | yes | Input parameter | {} |
| `param.area` | java.lang.String | yes | District | Кокшетау |
| `param.city` | java.lang.String | yes | City | Г.Кокшетау |
| `param.country` | java.lang.String | yes | Country | Казахстан |
| `param.province` | java.lang.String | yes | Province | Акмолинская область |

**Response (top level)**: `result` ResultModel – Result; `result.success` java.lang.Boolean – Whether successful; `result.code` java.lang.String – Return code; `result.message` java.lang.String – Attached information; `result.data` OrderAddressParseModel – Data; …

<a id="logisticsdeliveryurge-1"></a>
## 76. Urge seller to ship

`com.alibaba.logistics:logistics.delivery.urge:1` · Logistics · 催卖家发货  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.logistics/logistics.delivery.urge/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Urge the seller to ship. The order must still be in a not-yet-shipped status. Limited to once per 24 hours.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `orderId` | String | yes | Order id | 12898772891323 |

**Response (top level)**: `result` ResultModel – ; `result.success` boolean – ; `result.code` java.lang.String – ; `result.message` java.lang.String – ; `result.result` java.lang.Boolean – 

<a id="labelurlreceive-1"></a>
## 77. Receive external print-label (UDF) URL

`com.alibaba.fenxiao.crossborder:label.url.receive:1` · Logistics · 接收外部打印UDF链接  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/label.url.receive/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (optional) · user authorization not required · signature required  
Receive an external print-label (UDF) URL.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `orderId` | java.lang.String | yes | Order ID |  |
| `outOrderId` | java.lang.String | no | External order ID |  |
| `printUrls` | java.lang.String | yes | Shipping label print link |  |

**Response (top level)**: `result` Result – ; `result.success` java.lang.Boolean – ; `result.data` java.lang.Boolean – ; `result.msg` java.lang.String – 

<a id="bigcustomermailnoquery-1"></a>
## 78. Get waybill-number set

`com.alibaba.fenxiao.crossborder:bigcustomer.mailNo.query:1` · Logistics · 获取运单号集合  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/bigcustomer.mailNo.query/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Get the set of waybill numbers.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `outOrderId` | java.lang.String | yes | Order id | 1234 |

**Response (top level)**: `result` ResultModel – Return result; `result.model` java.util.Map – The specific data returned; `result.msg` java.lang.String – msg; `result.traceId` java.lang.String – traceId; `result.success` java.lang.Boolean – Whether successful

<a id="refundofficialdeliveryordercreate-1"></a>
## 79. Create official return-pickup logistics order

`com.alibaba.logistics:refundofficialdelivery.order.create:1` · Official Return Pickup · 创建退货官方物流上门揽订单  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.logistics/refundofficialdelivery.order.create/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Create an official-logistics door-to-door pickup order for a return.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `param` | OfficialDeliveryOrderCreateParam | yes |  |  |
| `param.couponId` | java.lang.String | no | Coupon id | 1234123 |
| `param.dateStr` | java.lang.String | yes | Door-to-door pickup time, format yyyy-MM-dd | 2024-07-11 |
| `param.gmtAcceptStart` | java.lang.String | yes | Door-to-door pickup start time, HH:mm | 9:00 |
| `param.gmtAcceptEnd` | java.lang.String | yes | Door-to-door pickup end time, HH:mm | 21:00 |
| `param.goodsType` | java.lang.String | yes | Product type (&quot;DAILY_NECESSITIES&quot;, &quot;Daily necessities&quot;); (&quot;FOOD&quot;, &quot;Food&quot;); (&quot;FURNITURE&quot;, &quot;Furniture&quot;); (&quot;METALS&quot;, &quot;Hardware&quot;); (&quot;COSMETICS&quot;, &quot;Cosmetics&quot;); (&quot;DIGITAL&quot;, &quot;Digital&quot;); (&quot;GIFT&quot;, &quot;Gift&quot;); (&quot;OTHER&quot;, &quot;Other&quot;); | DAILY_NECESSITIES |
| `param.packageCount` | java.lang.String | yes | Number of packages | 1 |
| `param.receiverInfo` | AddressParam | no | Recipient info | {} |
| `param.receiverInfo.address` | java.lang.String | yes |  |  |
| `param.receiverInfo.addressId` | java.lang.Long | yes |  |  |
| `param.receiverInfo.areaCode` | java.lang.String | yes |  |  |
| `param.receiverInfo.areaText` | java.lang.String | yes |  |  |
| `param.receiverInfo.cityCode` | java.lang.String | yes |  |  |
| `param.receiverInfo.cityText` | java.lang.String | yes |  |  |
| `param.receiverInfo.districtCode` | java.lang.String | yes |  |  |
| `param.receiverInfo.fullName` | java.lang.String | yes |  |  |
| `param.receiverInfo.mobile` | java.lang.String | yes |  |  |
| `param.receiverInfo.phone` | java.lang.String | yes |  |  |
| `param.receiverInfo.postCode` | java.lang.String | yes |  |  |
| `param.receiverInfo.provinceCode` | java.lang.String | yes |  |  |
| `param.receiverInfo.provinceText` | java.lang.String | yes |  |  |
| `param.receiverInfo.townCode` | java.lang.String | yes |  |  |
| `param.receiverInfo.townText` | java.lang.String | yes |  |  |
| `param.receiverInfoStr` | java.lang.String | no | Recipient info text | "" |
| `param.refundId` | java.lang.String | yes | Refund order ID | TQ284160529017092048 |
| `param.senderInfo` | AddressParam | no | Sender info | {} |
| `param.senderInfo.address` | java.lang.String | yes |  |  |
| `param.senderInfo.addressId` | java.lang.Long | yes |  |  |
| `param.senderInfo.areaCode` | java.lang.String | yes |  |  |
| `param.senderInfo.areaText` | java.lang.String | yes |  |  |
| `param.senderInfo.cityCode` | java.lang.String | yes |  |  |
| `param.senderInfo.cityText` | java.lang.String | yes |  |  |
| `param.senderInfo.districtCode` | java.lang.String | yes |  |  |
| `param.senderInfo.fullName` | java.lang.String | yes |  |  |
| `param.senderInfo.mobile` | java.lang.String | yes |  |  |
| `param.senderInfo.phone` | java.lang.String | yes |  |  |
| `param.senderInfo.postCode` | java.lang.String | yes |  |  |
| `param.senderInfo.provinceCode` | java.lang.String | yes |  |  |
| `param.senderInfo.provinceText` | java.lang.String | yes |  |  |
| `param.senderInfo.townCode` | java.lang.String | yes |  |  |
| `param.senderInfo.townText` | java.lang.String | yes |  |  |
| `param.senderInfoStr` | java.lang.String | no | Sender text | "" |
| `param.solutionCode` | java.lang.String | yes | Door-to-door pickup scheme code | 434 |
| `param.totalVolume` | java.lang.String | no | Total volume, unit: cm^3 | 10 |
| `param.totalWeight` | java.lang.String | no | Total weight, in g | 1000 |

**Response (top level)**: `result` ResultModel – ; `result.success` boolean – Whether successful; `result.code` java.lang.String – Error code; `result.message` java.lang.String – Error message; `result.result` Long – Returned order ID

<a id="refundofficialdeliveryordermodify-1"></a>
## 80. Modify official return-pickup logistics order

`com.alibaba.logistics:refundofficialdelivery.order.modify:1` · Official Return Pickup · 退货官方物流上门揽订单修改  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.logistics/refundofficialdelivery.order.modify/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Modify an official-logistics door-to-door pickup order for a return.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `param` | OfficialDeliveryOrderModifyParam | yes |  |  |
| `param.officialDeliveryOrderId` | java.lang.Long | yes | Door-to-door pickup order ID | 7210001 |
| `param.couponId` | java.lang.String | no | Coupon id | 1234123 |
| `param.dateStr` | java.lang.String | yes | Door-to-door pickup date, yyyy-MM-dd | 2024-07-11 |
| `param.gmtAcceptEnd` | java.lang.String | yes | Door-to-door pickup start time, HH:mm | 9:00 |
| `param.gmtAcceptStart` | java.lang.String | yes | Door-to-door pickup end time, HH:mm | 21:00 |
| `param.goodsType` | java.lang.String | yes | Product type (&quot;DAILY_NECESSITIES&quot;, &quot;Daily necessities&quot;); (&quot;FOOD&quot;, &quot;Food&quot;); (&quot;FURNITURE&quot;, &quot;Furniture&quot;); (&quot;METALS&quot;, &quot;Hardware&quot;); (&quot;COSMETICS&quot;, &quot;Cosmetics&quot;); (&quot;DIGITAL&quot;, &quot;Digital&quot;); (&quot;GIFT&quot;, &quot;Gift&quot;); (&quot;OTHER&quot;, &quot;Other&quot;); | DAILY_NECESSITIES |
| `param.packageCount` | java.lang.String | yes | Number of packages | 1 |
| `param.receiverInfo` | AddressParam | no | Recipient info | {} |
| `param.receiverInfo.address` | java.lang.String | yes |  |  |
| `param.receiverInfo.addressId` | java.lang.Long | yes |  |  |
| `param.receiverInfo.areaCode` | java.lang.String | yes |  |  |
| `param.receiverInfo.areaText` | java.lang.String | yes |  |  |
| `param.receiverInfo.cityCode` | java.lang.String | yes |  |  |
| `param.receiverInfo.cityText` | java.lang.String | yes |  |  |
| `param.receiverInfo.districtCode` | java.lang.String | yes |  |  |
| `param.receiverInfo.fullName` | java.lang.String | yes |  |  |
| `param.receiverInfo.mobile` | java.lang.String | yes |  |  |
| `param.receiverInfo.phone` | java.lang.String | yes |  |  |
| `param.receiverInfo.postCode` | java.lang.String | yes |  |  |
| `param.receiverInfo.provinceCode` | java.lang.String | yes |  |  |
| `param.receiverInfo.provinceText` | java.lang.String | yes |  |  |
| `param.receiverInfo.townCode` | java.lang.String | yes |  |  |
| `param.receiverInfo.townText` | java.lang.String | yes |  |  |
| `param.receiverInfoStr` | java.lang.String | no | Recipient info text | "" |
| `param.refundId` | java.lang.String | yes | Refund order ID | TQ284160529017092048 |
| `param.senderInfo` | AddressParam | no | Sender info | {} |
| `param.senderInfo.address` | java.lang.String | yes |  |  |
| `param.senderInfo.addressId` | java.lang.Long | yes |  |  |
| `param.senderInfo.areaCode` | java.lang.String | yes |  |  |
| `param.senderInfo.areaText` | java.lang.String | yes |  |  |
| `param.senderInfo.cityCode` | java.lang.String | yes |  |  |
| `param.senderInfo.cityText` | java.lang.String | yes |  |  |
| `param.senderInfo.districtCode` | java.lang.String | yes |  |  |
| `param.senderInfo.fullName` | java.lang.String | yes |  |  |
| `param.senderInfo.mobile` | java.lang.String | yes |  |  |
| `param.senderInfo.phone` | java.lang.String | yes |  |  |
| `param.senderInfo.postCode` | java.lang.String | yes |  |  |
| `param.senderInfo.provinceCode` | java.lang.String | yes |  |  |
| `param.senderInfo.provinceText` | java.lang.String | yes |  |  |
| `param.senderInfo.townCode` | java.lang.String | yes |  |  |
| `param.senderInfo.townText` | java.lang.String | yes |  |  |
| `param.senderInfoStr` | java.lang.String | no | Sender text | "" |
| `param.solutionCode` | java.lang.String | yes | Door-to-door pickup scheme code | 434 |
| `param.totalVolume` | java.lang.String | no | Total volume, unit: cm^3 | 10 |
| `param.totalWeight` | java.lang.String | no | Total weight, in g | 1000 |

**Response (top level)**: `result` ResultModel – ; `result.success` boolean – Whether successful; `result.code` java.lang.String – Error code; `result.message` java.lang.String – Error message; `result.result` java.lang.Boolean – Whether the modification was successful

<a id="refundofficialdeliveryorderget-1"></a>
## 81. Get official return-pickup logistics order details

`com.alibaba.logistics:refundofficialdelivery.order.get:1` · Official Return Pickup · 退货官方物流上门揽订单详情获取  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.logistics/refundofficialdelivery.order.get/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Get the details of an official-logistics door-to-door pickup order for a return.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `officialDeliveryOrderId` | String | yes | Official logistics door-to-door pickup order ID | 1988772 |

**Response (top level)**: `result` ResultModel – Return value; `result.success` boolean – Whether successful; `result.code` java.lang.String – Error code; `result.message` java.lang.String – Error message; `result.result` OfficialDeliveryOrderInfo – Order details; …

<a id="refundofficialdeliveryordercancel-1"></a>
## 82. Cancel official return-pickup logistics order

`com.alibaba.logistics:refundofficialdelivery.order.cancel:1` · Official Return Pickup · 取消退货官方物流上门揽订单  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.logistics/refundofficialdelivery.order.cancel/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Cancel an official-logistics door-to-door pickup order for a return.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `param` | OfficialDeliveryOrderCancelParam | yes |  |  |
| `param.cancelText` | java.lang.String | yes | Cancellation reason | 计划有变，暂时不需要寄了 |
| `param.cancelType` | java.lang.String | yes | Cancellation reason type: INFORMATION_INCORRECT(&quot;Information entered incorrectly (time/address etc. needs to be changed)&quot;), WANT_TO_SEND_MYSELF(&quot;Want to go to a nearby service point to send it myself&quot;), PLAN_CHANGED(&quot;Plans changed, no need to ship for now&quot;), WANT_TO_CHANGE_PICKUP_TIME(&quot;Want to change the door-to-door pickup time&quot;), PRICE_TOO_HIGH(&quot;I think the price is a bit high&quot;), 	ATE(&quot;Courier did not arrive for pickup on time&quot;), COURIER_NOT_COMING(&quot;Courier does not come for pickup&quot;), COURIER_BAD_ATTITUDE(&quot;Courier's service attitude is poor&quot;), ITEM_CANNOT_BE_SHIPPED(&quot;Item type cannot be shipped&quot;), COURIER_TOO_BUSY(&quot;Courier reported unable to pick up due to tight capacity&quot;), OTHER(&quot;Other&quot;), | PLAN_CHANGED |
| `param.officialDeliveryOrderId` | java.lang.Long | yes | Door-to-door pickup order ID | 7210001 |

**Response (top level)**: `result` ResultModel – ; `result.success` boolean – Whether successful; `result.code` java.lang.String – Error code; `result.message` java.lang.String – Error message; `result.result` java.lang.Boolean – Whether the order cancellation was successful

<a id="refundofficialdeliverysolutionget-1"></a>
## 83. Query official return-pickup plans

`com.alibaba.logistics:refundofficialdelivery.solution.get:1` · Official Return Pickup · 官方退货上门揽方案查询  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.logistics/refundofficialdelivery.solution.get/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Query the available official door-to-door return-pickup plans.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `param` | OfficialDeliverySolutionQueryParam | yes |  |  |
| `param.receiverInfo` | AddressParam | no | Recipient, i.e. the seller's shipping address information | {} |
| `param.receiverInfo.address` | java.lang.String | yes |  |  |
| `param.receiverInfo.addressId` | java.lang.Long | yes |  |  |
| `param.receiverInfo.areaCode` | java.lang.String | yes |  |  |
| `param.receiverInfo.areaText` | java.lang.String | yes |  |  |
| `param.receiverInfo.cityCode` | java.lang.String | yes |  |  |
| `param.receiverInfo.cityText` | java.lang.String | yes |  |  |
| `param.receiverInfo.districtCode` | java.lang.String | yes |  |  |
| `param.receiverInfo.fullName` | java.lang.String | yes |  |  |
| `param.receiverInfo.mobile` | java.lang.String | yes |  |  |
| `param.receiverInfo.phone` | java.lang.String | yes |  |  |
| `param.receiverInfo.postCode` | java.lang.String | yes |  |  |
| `param.receiverInfo.provinceCode` | java.lang.String | yes |  |  |
| `param.receiverInfo.provinceText` | java.lang.String | yes |  |  |
| `param.receiverInfo.townCode` | java.lang.String | yes |  |  |
| `param.receiverInfo.townText` | java.lang.String | yes |  |  |
| `param.senderInfo` | AddressParam | no | Sender, i.e., the buyer's door-to-door pickup address | {} |
| `param.senderInfo.address` | java.lang.String | yes |  |  |
| `param.senderInfo.addressId` | java.lang.Long | yes |  |  |
| `param.senderInfo.areaCode` | java.lang.String | yes |  |  |
| `param.senderInfo.areaText` | java.lang.String | yes |  |  |
| `param.senderInfo.cityCode` | java.lang.String | yes |  |  |
| `param.senderInfo.cityText` | java.lang.String | yes |  |  |
| `param.senderInfo.districtCode` | java.lang.String | yes |  |  |
| `param.senderInfo.fullName` | java.lang.String | yes |  |  |
| `param.senderInfo.mobile` | java.lang.String | yes |  |  |
| `param.senderInfo.phone` | java.lang.String | yes |  |  |
| `param.senderInfo.postCode` | java.lang.String | yes |  |  |
| `param.senderInfo.provinceCode` | java.lang.String | yes |  |  |
| `param.senderInfo.provinceText` | java.lang.String | yes |  |  |
| `param.senderInfo.townCode` | java.lang.String | yes |  |  |
| `param.senderInfo.townText` | java.lang.String | yes |  |  |
| `param.receiverInfoStr` | java.lang.String | no | Recipient, i.e., the seller's receiving address; text | {\"townName\":\"东华门街道\",\"townCode\":\"110101001\",\"cityCode\":\"110100\",\"contactorName… |
| `param.senderInfoStr` | java.lang.String | no | Sender, i.e., the buyer's door-to-door pickup address; text | {\"townName\":\"东华门街道\",\"townCode\":\"110101001\",\"cityCode\":\"110100\",\"contactorName… |
| `param.refundId` | java.lang.String | yes | Refund order ID | TQ284019625557092048 |
| `param.packageCount` | java.lang.Integer | yes | Number of packages | 1 |
| `param.totalVolume` | java.lang.String | no | Total volume, cm^3 | 1 |
| `param.totalWeight` | java.lang.String | no | Total weight, g | 1 |
| `param.couponId` | String | no | Coupon id |  |

**Response (top level)**: `result` ResultModel – ; `result.success` boolean – Whether successful; `result.code` java.lang.String – Error code; `result.message` java.lang.String – Error message; `result.result` OfficialDeliverySolutionResult – Return value; …

<a id="orderrelationwrite-1"></a>
## 84. Write back mapping between end-customer orders and 1688 orders

`com.alibaba.fenxiao.crossborder:order.relation.write:1` · Data Write-back · 回传机构真实用户订单和1688订单的映射关系  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/order.relation.write/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Write back to 1688 the mapping between the organisation's real end-customer orders and the 1688 orders generated under the organisation's account.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `orderRelationParam` | OrderRelationParam | yes | Order relationship parameter | 如下 |
| `orderRelationParam.orderId` | java.lang.String | yes | External institution sub-order ID | 4590347523948375 |
| `orderRelationParam.parentOrderId` | java.lang.String | yes | External institution's main order ID | 4590347523948370 |
| `orderRelationParam.purchaseOrderId` | java.lang.Long | yes | 1688 purchase sub-order ID | 3209572465452734 |
| `orderRelationParam.purchaseParentOrderId` | java.lang.Long | yes | 1688 purchase sub-order ID | 3209572465452730 |

**Response (top level)**: `result` ResultModel – Return result; `result.success` boolean – Whether successful; `result.code` java.lang.String – Error code; `result.message` java.lang.String – Error description

<a id="tradecrosslogisticsordersync-1"></a>
## 85. Write back country-site logistics orders

`com.alibaba.fenxiao.crossborder:trade.cross.logisticsOrderSync:1` · Data Write-back · 国家站物流单回传  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/trade.cross.logisticsOrderSync/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Write back logistics orders from the country site.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `orderLogisticsParam` | OrderLogisticsParam | yes |  |  |
| `orderLogisticsParam.orderId` | java.lang.String | yes | Country site main order ID | 1213 |
| `orderLogisticsParam.logisticsId` | java.lang.String | yes | Downstream logistics order ID | 1·21213 |
| `orderLogisticsParam.logisticsStatus` | java.lang.String | yes | Logistics status: send - shipped, sign - signed for, partSign - partially received, refund - returned | sign |
| `orderLogisticsParam.createTime` | java.util.Date | yes | Creation time | 2023-10-10 10:12:12 |
| `orderLogisticsParam.signTime` | java.util.Date | yes | Signed receipt time | 2023-10-10 10:12:12 |
| `orderLogisticsParam.province` | java.lang.String | yes | Province | 江西省 |
| `orderLogisticsParam.city` | java.lang.String | yes | City | 南昌市 |
| `orderLogisticsParam.area` | java.lang.String | yes | District | 白云区 |
| `orderLogisticsParam.address` | java.lang.String | yes | Detailed address | 江西省南昌市白云区xx号 |

**Response (top level)**: `result` ResultModel – Return message; `result.success` Boolean – Whether successful; `result.code` String – code; `result.message` String – Return message

<a id="tradecrossordersync-1"></a>
## 86. Sync downstream sales orders

`com.alibaba.fenxiao.crossborder:trade.cross.orderSync:1` · Data Write-back · 下游销售订单同步  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/trade.cross.orderSync/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Sync downstream sales orders.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `orderParam` | OrderParamV | yes | Order details |  |
| `orderParam.orderId` | String | yes | Order id | 1234 |
| `orderParam.productId` | String | no | Product ID; required for sub-orders or combined main-sub orders | 12 |
| `orderParam.productName` | String | no | Product name; required for sub-orders or merged main/sub orders | 手套 |
| `orderParam.skuId` | String | no | skuId; required for sub-orders or merged main/sub orders | adsds1213 |
| `orderParam.skuName` | String | no | SKU name; required for sub-orders or combined main-sub orders | 绿色 |
| `orderParam.buyAmount` | Long | no | Purchase quantity | 2 |
| `orderParam.createTime` | String | yes | Creation time | 2023-10-01 10:10:10 |
| `orderParam.outMemberId` | String | yes | Downstream user ID | 1212113 |
| `orderParam.payTime` | String | yes | Payment time | 2023-10-01 10:10:10 |
| `orderParam.endTime` | String | no | Completion time | 2023-10-01 10:10:10 |
| `orderParam.paidFee` | Long | yes | Actual payment amount, in cents | 200 |
| `orderParam.refundStatus` | String | no | Refund status 0 - not refunded (refund closed, not applied) 1 - refunded | 0 |
| `orderParam.status` | String | yes | payed - paid, success - transaction successful, close - transaction closed | success |
| `orderParam.subOrderParamList` | SubOrderParam[] | no | For one main order with multiple sub-orders, sub-order info must be passed | 子单信息 |
| `orderParam.subOrderParamList[].orderId` | String | yes | Sub-order ID | 2 |
| `orderParam.subOrderParamList[].productId` | String | yes | 1688 product ID, must be in plaintext | 1 |
| `orderParam.subOrderParamList[].productName` | String | yes | 1688 product name, must be in plaintext | 袜子 |
| `orderParam.subOrderParamList[].skuId` | String | yes | 1688 skuId must be plaintext | 1 |
| `orderParam.subOrderParamList[].skuName` | String | yes | 1688 SKU name, must be in plaintext | 绿色 |
| `orderParam.subOrderParamList[].buyAmount` | Long | yes | Purchase quantity of the sub-order | 1 |
| `orderParam.subOrderParamList[].paidFee` | Long | yes | Actual amount paid, in cents (fen) | 100 |
| `orderParam.subOrderParamList[].refundStatus` | String | yes | Refund status: 0 - not refunded or refund closed, 1 - refunded | 0 |

**Response (top level)**: `result` ResultModel – ; `result.success` boolean – ; `result.code` java.lang.String – ; `result.message` java.lang.String – 

<a id="accountbusinesssave-1"></a>
## 87. Save the business line an account belongs to

`com.alibaba.fenxiao.crossborder:account.business.save:1` · Data Write-back · 保存账号所属业务线  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/account.business.save/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Save the business line that an account belongs to.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `accountPerformance` | AccountPerformance | yes | Input parameter | 如下 |
| `accountPerformance.account` | java.lang.String | yes | Account name | b测试账号009 或 b测试账号009:lctest |
| `accountPerformance.business` | java.lang.String | yes | Business line it belongs to | 东南亚sea，南亚sa，日韩jk，港澳台hmt，中东me，北美na，拉美la，西欧we，泛俄ru，非洲af |

**Response (top level)**: `result` AccountPerformanceResult – Result; `result.saveResult` Boolean – Save result, true for success, false for failure

<a id="alibabatraderefundopquerybatchrefundbyorderidandstatus-1"></a>
## 88. Query refund details by order ID (buyer view)

`com.alibaba.trade:alibaba.trade.refund.OpQueryBatchRefundByOrderIdAndStatus:1` · Returns & Refunds · 查询退款单详情-根据订单ID（买家视角）  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.trade.refund.OpQueryBatchRefundByOrderIdAndStatus/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
For buyers; sellers should use alibaba.trade.refund.OpQueryBatchRefundByOrderIdAndStatus.sellerView. Queries the refund list for an order in real time. Currently only in-sale refunds (before the transaction completes) can be queried.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `orderId` | String | yes | Order id | 151267031**8969811 |
| `queryType` | String | yes | 1: active; 3: refund successful (only refund-in-progress and refund-successful are supported) | 3 |

**Response (top level)**: `result` OpQueryBatchRefundByOrderIdAndStatusResultModel – Query result; `result.opOrderRefundModels` OpOrderRefundModel[] – Refund order information; `errorCode` String – Error code; `errorMessage` String – Error message; `extErrorMessage` String – Additional information; …

<a id="alibabatraderefundopqueryorderrefundoperationlist-1"></a>
## 89. Refund operation history (buyer view)

`com.alibaba.trade:alibaba.trade.refund.OpQueryOrderRefundOperationList:1` · Returns & Refunds · 退款单操作记录列表（买家视角）  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.trade.refund.OpQueryOrderRefundOperationList/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
For buyers; sellers should use alibaba.trade.refund.OpQueryOrderRefund.sellerView. Returns the buyer-side refund operation records.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `refundId` | String | yes | Refund order ID | TQ1043162**46961198 |
| `pageNo` | String | yes | Current page number | 1 |
| `pageSize` | String | yes | Page size | 100 |

**Response (top level)**: `result` OpQueryOrderRefundOperationListResult – Return result; `result.opOrderRefundOperationModels` OpOrderRefundOperationModel[] – Refund info; `errorMessage` String – Error message; `extErrorMessage` String – Additional error information; `errorCode` String – Error code; …

<a id="alibabatraderefundopqueryorderrefund-1"></a>
## 90. Query refund details by refund ID (buyer view)

`com.alibaba.trade:alibaba.trade.refund.OpQueryOrderRefund:1` · Returns & Refunds · 查询退款单详情-根据退款单ID（买家视角）  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.trade.refund.OpQueryOrderRefund/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
For buyers; sellers should use alibaba.trade.refund.OpQueryOrderRefund.sellerView. Queries refund details, including the list of refund operations. Permission must be requested from Alibaba to access this API.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `refundId` | String | yes | Business primary key of the refund order: TQ+ID | TQ11173622***991577 |
| `needTimeOutInfo` | boolean | no | Timeout information needed for the refund order | true |
| `needOrderRefundOperation` | boolean | no | All refund operation information that must accompany the refund order | true |

**Response (top level)**: `result` OpQueryOrderRefund – Query result; `result.opOrderRefundModelDetail` OpOrderRefundModel – Return value; `errorCode` String – Error code; `errorMessage` String – Error description; `extErrorMessage` String – Additional error description information; …

<a id="alibabatraderefundbuyerqueryorderrefundlist-1"></a>
## 91. Query refund list (buyer view)

`com.alibaba.trade:alibaba.trade.refund.buyer.queryOrderRefundList:1` · Returns & Refunds · 查询退款单列表(买家视角)  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.trade.refund.buyer.queryOrderRefundList/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Buyer views the refund list. This interface does not support sub-account queries; authorize with the main account before querying.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `orderId` | Long | no | Order ID | 179087886005498520 |
| `applyStartTime` | java.util.Date | no | Refund application time (start) | 20170926114526000+0800 |
| `applyEndTime` | java.util.Date | no | Refund application time (end/deadline) | 20220926114526000+0800 |
| `refundStatusSet` | String[] | no | Refund status list | 等待卖家同意 waitselleragree;退款成功 refundsuccess;退款关闭 refundclose;待买家修改 waitbuyermodify;等待买家退货 wa… |
| `sellerMemberId` | String | no | Seller memberId | b2b-1623492085 |
| `currentPageNum` | Integer | no | Current page number | 0 |
| `pageSize` | Integer | no | Number of records per page | 20 |
| `logisticsNo` | String | no | Return logistics waybill number (when passing this field for a query, sellerMemberId must also be passed) | 3101***159271 |
| `modifyStartTime` | java.util.Date | no | Refund modification time (start) | 20170926114526000+0800 |
| `modifyEndTime` | java.util.Date | no | Refund modification time (end) | 20220926114526000+0800 |
| `dipsuteType` | Integer | no | 1: in-sale refund, 2: after-sales refund; 0: all refund orders | 1 |

**Response (top level)**: `result` OpQueryOrderRefundListResult – Query result; `result.opOrderRefundModels` OpOrderRefundModel[] – List of refund orders; `result.totalCount` int – Total number of records that meet the criteria; `result.currentPageNum` int – Current page number of the query; `errorCode` String – Error code; `errorMsg` String – Error message; `success` Boolean – Whether successful; …

<a id="alibabatradequeryorderbyinsure-1"></a>
## 92. Query order number by claim or insurance-policy number

`com.alibaba.trade:alibaba.trade.queryOrderByInsure:1` · Returns & Refunds · 根据理赔单或保险单号查询对应的订单号  
Detail doc unavailable on open.1688.com (server error "error finding API").

<a id="alibabatradecreaterefund-1"></a>
## 93. Create refund/return request

`com.alibaba.trade:alibaba.trade.createRefund:1` · Returns & Refunds · 创建退款退货申请  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.trade.createRefund/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Create a refund or return request.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `orderId` | Long | yes | Main order |  |
| `orderEntryIds` | Long[] | yes | Sub-order |  |
| `disputeRequest` | String | yes | Refund / refund and return. Refund and return can only be selected once the goods have been received. | 退款:"refund"; 退款退货:"returnRefund" |
| `applyPayment` | Long | yes | Refund amount (unit: cent). Must not exceed the actual payment amount; when waiting for the seller to ship, it must equal the product's actual payment amount. |  |
| `applyCarriage` | Long | yes | Refunded freight amount (unit: cents/fen). |  |
| `applyReasonId` | Long | yes | Refund reason ID (obtained from the API getRefundReasonList) |  |
| `description` | String | yes | Reason for refund application, 2-150 characters |  |
| `goodsStatus` | String | yes | Cargo status | 售中等待卖家发货:"refundWaitSellerSend"; 售中等待买家收货:"refundWaitBuyerReceive"; 售中已收货（未确认完成交易）:"refund… |
| `vouchers` | String[] | no | Voucher image URLs. 1-5 images, must use the “image domain/relative path” returned by the uploadRefundVoucher API | [https://cbu01.alicdn.com/img/ibank/2019/901/930/11848039109.jpg] |
| `orderEntryCountList` | OrderEntryCountModel[] | no | Sub-order refund quantity. The return quantity can be specified only when the buyer has already received the goods during in-sale (refund and return); by default, all goods are returned. | [{"id":586683458996743215,"count":1}] |
| `orderEntryCountList[].id` | Long | yes | Sub-order ID | 1 |
| `orderEntryCountList[].count` | Integer | yes | Quantity of products purchased in the sub-order | 1 |
| `customRefund` | Boolean | no | Whether the return shipping fee is customized | true：订单退款运费按照isv回传为准，不做优化；false（默认）：订单退款金额按照平台规则为准，需要做优化； |
| `refundRemark` | String | no | Refund description, used to receive the actual description of the downstream refund order | 克重不对，偷工减料了 |

**Response (top level)**: `result` OrderRefundCreateResult – Return result; `result.code` String – Error code; `result.message` String – Error message; `result.result` OrderRefundCreateResult – Result; `result.success` boolean – Whether successful; …

<a id="alibabatradegetrefundreasonlist-1"></a>
## 94. Query refund/return reasons (for creating a request)

`com.alibaba.trade:alibaba.trade.getRefundReasonList:1` · Returns & Refunds · 查询退款退货原因（用于创建退款退货）  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.trade.getRefundReasonList/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Query the refund/return reasons (used when creating a refund or return request).

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `orderId` | Long | yes | Main order id |  |
| `orderEntryIds` | Long[] | yes | Sub-order ID |  |
| `goodsStatus` | String | yes | Cargo status | 售中等待买家发货:”refundWaitSellerSend"; 售中等待买家收货:"refundWaitBuyerReceive"; 售中已收货（未确认完成交易）:"refund… |

**Response (top level)**: `result` OrderRefundReasonListResult – Return result; `result.code` String – Error code; `result.message` String – Error message; `result.result` OrderRefundReasonListResult – Result; `result.success` Boolean – Whether successful; …

<a id="alibabatradeuploadrefundvoucher-1"></a>
## 95. Upload refund/return evidence

`com.alibaba.trade:alibaba.trade.uploadRefundVoucher:1` · Returns & Refunds · 上传退款退货凭证  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.trade.uploadRefundVoucher/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Upload evidence for a refund/return request. To convert a file stream to a byte array, org.apache.commons.io.IOUtils#toByteArray(java.io.InputStream) is recommended.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `imageData` | byte[] | yes | Voucher image data. Less than 1MB, jpg format. |  |

**Response (top level)**: `result` OrderRefundUploadVoucherResult – Return result; `result.code` String – Error code; `result.message` String – Error message; `result.result` OrderRefundUploadVoucherResult – Success result; `result.success` Boolean – Whether successful; …

<a id="alibabatraderefundreturngoods-1"></a>
## 96. Buyer submits return-shipment info

`com.alibaba.trade:alibaba.trade.refund.returnGoods:1` · Returns & Refunds · 买家提交退款货信息  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.trade.refund.returnGoods/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Used after the seller approves the buyer's return/refund request, for the buyer to submit the return-shipment information. First call alibaba.logistics.OpQueryLogisticCompanyList.offline to look up logistics companies, and use the logistics-company code it returns.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `refundId` | String | yes | Refund order number, starts with TQ | TQ36706338027991577 |
| `logisticsCompanyNo` | String | yes | Logistics company code, query via the alibaba.logistics.OpQueryLogisticCompanyList.offline API | ZTO |
| `freightBill` | String | yes | Logistics company waybill number. Please fill it in accurately, otherwise the seller has the right to refuse the refund. | 3110044550034338 |
| `description` | String | no | Shipping note; content must be between 2-200 characters | 发货说明 |
| `vouchers` | String[] | no | Voucher image URLs. Must use the “image domain/relative path” returned by the API alibaba.trade.uploadRefundVoucher. Up to 10 images can be uploaded; each image must not exceed 1MB; jpg, gif, jpeg, png, and bmp formats are supported. Please upload vouchers as they are required for a subsequent claim (a claim cannot be filed without them). | [https://cbu01.alicdn.com/img/ibank/2019/901/930/11848039109.jpg] |

**Response (top level)**: `result` RefundReturnGoodsResult – Return result; `result.errorCode` java.lang.String – Error code; `result.errorInfo` java.lang.String – Error description; `result.success` boolean – Whether the submission succeeded

<a id="alibabatradecancelrefund-1"></a>
## 97. Cancel refund/return request

`com.alibaba.trade:alibaba.trade.cancelRefund:1` · Returns & Refunds · 取消退款退货申请  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.trade.cancelRefund/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (optional) · user authorization not required · signature required  
Cancel a refund or return request.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `refundId` | String | yes | Refund order ID | TQ267395256051660259 |

**Response (top level)**: `result` RefundCancelResult – ; `result.success` boolean – ; `result.errorInfo` java.lang.String – ; `result.errorCode` java.lang.String – 

<a id="alibabatradegetmaxrefundfee-1"></a>
## 98. Query maximum refundable amount when applying

`com.alibaba.trade:alibaba.trade.getMaxRefundFee:1` · Returns & Refunds · 申请退款时查询最大可退费用  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.trade.getMaxRefundFee/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Query the maximum refundable amount when applying for a refund.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `input` | OpMaxRefundFeeGetParam | yes |  |  |
| `input.goodsStatus` | java.lang.String | yes | Cargo status | 售中等待卖家发货:"refundWaitSellerSend"; 售中等待买家收货:"refundWaitBuyerReceive"; 售中已收货（未确认完成交易）:"refund… |
| `input.orderId` | java.lang.Long | yes | Order ID | 123 |
| `input.refundId` | java.lang.String | no | The refund order must be in refunding status; this can be omitted | TQ123 |
| `input.refundGoodsCountList` | OrderEntryCountModel[] | yes | Return quantity | [{1: 1}] |
| `input.refundGoodsCountList[].id` | Long | yes | Sub-order ID | 1 |
| `input.refundGoodsCountList[].count` | Integer | yes | Quantity of products purchased in the sub-order | 1 |

**Response (top level)**: `result` OpMaxRefundFeeResultModel – ; `result.code` java.lang.String – ; `result.message` java.lang.String – ; `result.result` OpMaxRefundFeeModel – ; `result.retCodes` java.lang.String[] – ; `result.subCode` java.lang.String – ; `result.subMessage` java.lang.String – ; `result.success` boolean – ; …

<a id="tradearbitrationapply-1"></a>
## 99. Apply for trade arbitration

`com.alibaba.trade:trade.arbitration.apply:1` · Returns & Refunds · 申请交易仲裁  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/trade.arbitration.apply/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Trade arbitration application.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `param` | TradeArbitrateApplyParam | yes | Appeal parameter model | {} |
| `param.type` | String | yes | Complaint type: RefundComplaint (refund complaint), SafeRefundComplaint (secure refund complaint), AfterSalesComplaint (after-sales case complaint), TradeComplaint (after-sales order complaint) | RefundComplaint |
| `param.refundId` | String | no | Return/refund order ID; refundId is required for refund complaints | TQ129988192325 |
| `param.orderId` | Long | yes | Order ID | 3356489569898 |
| `param.reasonText` | java.lang.String | yes | Application reason | 货没收到 |
| `param.reasonId` | java.lang.Integer | yes | Application reason ID | 12 |
| `param.picUrlEvidence` | java.lang.String | no | Evidence image | https://cnd01.1688.com/evidence.png |

**Response (top level)**: `result` tradearbitrationapplyResultModel – Appeal return model; `result.success` boolean – Whether successful; `result.code` java.lang.String – Error code; `result.message` java.lang.String – Error message; `result.result` java.lang.Long – Appeal order ID

<a id="refundaddressget-1"></a>
## 100. Query the merchant's return address

`com.alibaba.fenxiao:refund.address.get:1` · Returns & Refunds · 查询商家退货地址  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao/refund.address.get/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Query the return address.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `mainOrderId` | String | yes | Main order id | 123 |

**Response (top level)**: `result` RefundAddressResult – Result; `result.data` RefundAddressInfoModel[] – Return address list; `result.errorCode` String – Error code; `result.errorInfo` String – Error message; `result.success` Boolean – API call result; …

<a id="pushmessageconfirm-1"></a>
## 101. Batch-confirm failed messages

`cn.alibaba.open:push.message.confirm:1` · Messaging · 失败消息批量确认  
POST `https://gw.open.1688.com/openapi/param2/1/cn.alibaba.open/push.message.confirm/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (optional) · user authorization not required · signature required  
Manually call the confirmation API to confirm that messages have been consumed successfully. Only needed when using the query-style API to fetch failed messages.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `msgIdList` | java.util.List | no | List of message IDs pending confirmation | [123,456] |

**Response (top level)**: `isSuccess` boolean – Whether the operation succeeded

<a id="pushquerymessagelist-1"></a>
## 102. Fetch failed messages (query style)

`cn.alibaba.open:push.query.messageList:1` · Messaging · 查询式获取失败的消息列表  
POST `https://gw.open.1688.com/openapi/param2/1/cn.alibaba.open/push.query.messageList/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (optional) · user authorization not required · signature required  
Query-style retrieval of sent messages. Retrieved messages are not confirmed automatically; the caller must call the confirmation API to confirm the consumption status. Note that confirming affects the data returned by pagination.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `createStartTime` | java.util.Date | no | Start of the message creation time search range | 20130417000000000+0800 |
| `createEndTime` | java.util.Date | no | End of the message creation time search range | 20130417000000000+0800 |
| `page` | int | no | Current data page, default is 1 | 1 |
| `pageSize` | int | no | Number of records retrieved per page, range 20-200, default 20 | 20 |
| `type` | String | no | Message type | ORDER_BUYER_MAKER |
| `userInfo` | String | no | User Id | b2b-4137495171f2513 |

**Response (top level)**: `pushMessagePage` PushMessagePage – Paginated data; `pushMessagePage.datas` PushMessage[] – Paginated list of message data; `pushMessagePage.totalCount` int – Total number of messages; …

<a id="pushcursormessagelist-1"></a>
## 103. Fetch failed messages (cursor style)

`cn.alibaba.open:push.cursor.messageList:1` · Messaging · 游标式获取失败的消息列表  
POST `https://gw.open.1688.com/openapi/param2/1/cn.alibaba.open/push.cursor.messageList/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (optional) · user authorization not required · signature required  
Cursor-style retrieval of failed messages. Retrieved messages are automatically confirmed as consumed, so the next call with the same conditions returns the remaining data, until the result is empty.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `createStartTime` | java.util.Date | no | Start of the message creation time search range | 20130417000000000+0800 |
| `createEndTime` | java.util.Date | no | End of the message creation time search range | 20130417000000000+0800 |
| `quantity` | int | no | Number of records fetched per request, range 20-200, default 20 | 20 |
| `type` | String | no | Message type | ORDER_BUYER_MAKER |
| `userInfo` | String | no | User Id | b2b-4137495171f2513 |

**Response (top level)**: `pushMessageList` PushMessage[] – List of push messages; …

<a id="alibabatradegetsellerorderlist-1"></a>
## 104. View order list (seller view)

`com.alibaba.trade:alibaba.trade.getSellerOrderList:1` · Messaging · 订单列表查看(卖家视角)  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.trade.getSellerOrderList/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Get the seller's order list; the user's memberId must equal the sellerMemberId on the orders. This interface returns only basic order information, not logistics or invoice information. For logistics information call the order-detail interface; for invoice information call the invoice-information API.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `createStartTime` | java.util.Date | no | Order start time | 20180721172608000+0800 |
| `createEndTime` | java.util.Date | no | Order placement end time | 20180721172608000+0800 |
| `modifyStartTime` | java.util.Date | no | Query modification time start | 20180721172608000+0800 |
| `modifyEndTime` | java.util.Date | no | End of the modification time query range | 20180721172608000+0800 |
| `page` | int | no | Query page number, starting from 1 | 1 |
| `pageSize` | int | no | Number of records per page for the query (maximum 20) | 10 |
| `orderStatus` | java.lang.String | no | Order status, possible values: success, cancel (transaction cancelled; liquidated damages, etc. have been settled), waitbuyerpay (waiting for buyer to pay), waitsellersend (waiting for seller to ship), waitbuyerreceive (waiting for buyer to receive the goods) | waitbuyerpay |
| `refundStatus` | java.lang.String | no | Refund status, supports: &quot;waitselleragree&quot; (waiting for seller to agree), &quot;refundsuccess&quot; (refund successful), &quot;refundclose&quot; (refund closed), &quot;waitbuyermodify&quot; (pending buyer modification), &quot;waitbuyersend&quot; (waiting for buyer to return goods), &quot;waitsellerreceive&quot; (waiting for seller to confirm receipt) | waitselleragree |
| `buyerMemberId` | java.lang.String | no | Buyer memberId or buyerOpenUid (buyer's encrypted ID) | b2b-1234325 |
| `buyerLoginId` | String | no | Buyer LoginId or buyerOpenUid (buyer's encrypted ID) | alitestforisv02 |
| `tradeType` | java.lang.String | no | Transaction type: Escrow transaction (1), Prepaid deposit transaction (2), ETC overseas acquiring transaction (3), Instant payment transaction (4), Guarantee fund security transaction (5), Unified transaction process (6), Staged transaction (7), Cash on delivery transaction (8), Credit voucher payment transaction (9), Account period payment transaction (10), 1688 Transaction 4.0, new staged transaction (50060), Face-to-face payment transaction process (50070), Service-type transaction process (50080) | 5 |
| `bizTypes` | String[] | no | Business type, supports: &quot;cn&quot; (regular order type), &quot;ws&quot; (large-value wholesale order type), &quot;yp&quot; (regular sample order type), &quot;yf&quot; (one-cent sample order type), &quot;fs&quot; (flash sale (limited-time discount) order type), &quot;cz&quot; (processing/customization order type), &quot;ag&quot; (agreement procurement order type), &quot;hp&quot; (group-buy order type), &quot;gc&quot; (national procurement order type), &quot;supply&quot; (supply-and-marketing order type), &quot;nyg&quot; (nyg order type), &quot;factory&quot; (Taobao Factory order type), &quot;quick&quot; (quick order placement), &quot;xiangpin&quot; (Xiangpin order), &quot;nest&quot; (Procurement Mall - Nest), &quot;f2f&quot; (face-to-face payment), &quot;cyfw&quot; (sample storage service), &quot;sp&quot; (consignment order flag), &quot;wg&quot; (WeiGong order), &quot;factorysamp&quot; (Taobao Factory sampling order), &quot;factorybig&quot; (Taobao Factory bulk order) | ["cn","ws"] |
| `isHis` | boolean | no | Whether to query the historical order table; default queries the current table | false |
| `productName` | java.lang.String | no | Product name | 测试商品名 |
| `needBuyerAddressAndPhone` | java.lang.Boolean | no | Whether the buyer's detailed address information and phone number need to be queried | false |
| `needMemoInfo` | java.lang.Boolean | no | Whether remark information needs to be queried | false |
| `tousuStatus` | boolean | no | Whether to search for the pending order-change request under complaint | false |
| `buyerRateStatus` | java.lang.Integer | no | Buyer review status (4: reviewed, 5: not reviewed, 6: review not required) | 5 |
| `sellerRateStatus` | java.lang.Integer | no | Seller rating status (4: rated, 5: not rated, 6: rating not required) | 5 |
| `needCheckSend` | Boolean | no | Whether shipment verification information needs to be returned; required and very important for label-printing and shipping scenarios | true |
| `needSendGoodsOverdueRisk` | Boolean | no | Whether the order shipment overdue risk flag is needed | false |
| `needDeliverGoodsOverdueRisk` | Boolean | no | Whether a pickup overdue risk flag is needed | false |
| `needOfficialLogisticOrder` | Boolean | no | Whether to filter orders with the official direct-delivery guarantee service | false |

**Response (top level)**: `result` TradeInfo[] – Query return result; `errorCode` String – Error code; `errorMessage` String – Error message; `totalRecord` Long – Total record count; `success` Boolean – Whether the call was successful; `retCodes` String[] – Masked information code; …

<a id="alibabalogisticsopdeliverysendorderoffline-1"></a>
## 105. Ship: seller arranges own logistics

`com.alibaba.logistics:alibaba.logistics.OpDeliverySendOrder.offline:1` · Messaging · 物流发货-自己联系物流发货  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.logistics/alibaba.logistics.OpDeliverySendOrder.offline/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
For 1688 open-marketplace orders where the seller arranges logistics themselves. Supports combined shipping (several orders in one shipment) and shipping at sub-order (orderEntry) level. Shipping by quantity is not supported.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `multiPackage` | Boolean | no | Whether to use multi-package shipment | true：使用多包裹发货，false或不传：使用单包裹发货 |
| `sendGoods` | OpSendGood[] | yes | Shipment object list |  |
| `sendGoods[].sourceId` | String | yes | Shipment object ID, generally the order ID |  |
| `sendGoods[].sendGoodEntries` | OpSendGoodEntry[] | yes | List of shipment object details |  |
| `sendGoods[].sendGoodEntries[].sourceEntryId` | String | yes | Shipment target line item ID, corresponding to the sub-order ID |  |
| `sendGoods[].sendGoodEntries[].amount` | Long | yes | Actual shipped quantity for the shipment target |  |
| `sendGoods[].sendGoodEntries[].weight` | Double | yes | The actual shipped weight of the shipment object; the weight unit defaults to kilograms. |  |
| `sendGoods[].sendGoodEntries[].extBody` | String | no | JSON array string passed when shipping multiple packages. cpCode is the logistics company code, corresponding to companyNo (not the numeric Id) from the logistics company information API; logisticsCpName: logistics company name, corresponding to companyName from the logistics company information API; mailNo: waybill number; this data can be queried via the &quot;Logistics Company List - Self-linked Logistics&quot; API. quantity: number of shipment objects in the package, required. | [{"cpCode":"SF","logisticsCpName":"顺丰","mailNo":"a1231","quantity":1},{"cpCode":"SF","logi… |
| `remarks` | String | no | Remark |  |
| `gmtSend` | java.util.Date | no | Shipping time | 标准时间格式：yyyyMMddHHmmssSSSZ，例如：20120801154220368+0800 |
| `extBody` | String | no | JSON string passed for single-package shipment. cpCode is the logistics company code, corresponding to companyNo (not the numeric Id) from the logistics company info API; logisticsCpName: logistics company name, corresponding to companyName from the logistics company info API; mailNo: waybill number. This data can be queried via the &quot;Logistics Company List - Self-Connected Logistics&quot; API. | {"cpCode":"SF","logisticsCpName":"顺丰","mailNo":"a123"} |
| `extParam` | String | no | JSON string |  |
| `receiverInfo` | OpReceiveContacter | no | Shipping address; takes priority over the order's shipping address. If empty, the order's shipping address is used |  |
| `receiverInfo.provinceCode` | String | yes | Province code |  |
| `receiverInfo.cityCode` | String | yes | City code |  |
| `receiverInfo.areaCode` | String | yes | Region code |  |
| `receiverInfo.townCode` | String | yes | Town or street code |  |
| `receiverInfo.province` | String | yes | Province name; can be omitted if code is passed |  |
| `receiverInfo.city` | String | yes | City name; can be omitted if code is passed |  |
| `receiverInfo.area` | String | yes | District name; if code is passed, this can be omitted |  |
| `receiverInfo.town` | String | yes | Town or street name |  |
| `receiverInfo.address` | String | yes | Detailed address |  |
| `receiverInfo.fullName` | String | yes | Name |  |
| `receiverInfo.corpName` | String | yes | Company name |  |
| `receiverInfo.post` | String | yes | Postal code |  |
| `receiverInfo.phone` | String | yes | Landline phone number |  |
| `receiverInfo.mobile` | String | yes | Mobile phone number |  |
| `receiverInfo.warehouse` | String | yes | Warehouse |  |
| `receiverInfo.codeType` | String | yes | Address code type; defaults to the Cainiao standard code |  |
| `isEncryptOrderSend` | String | no | Whether to ship using a downstream encrypted order number; pass Y when the seller obtains the number via downstream platform encryption. | Y,N |

**Response (top level)**: `result` OpSendOrderModelResult – Shipping details; `result.logisticsId` String – Logistics number; `result.sendGoods` OpSendGood[] – Shipping details; `result.sendSuccessList` OpSendOrderResult[] – List of packages successfully shipped in multi-package shipment; `result.sendFailList` OpSendOrderResult[] – List of packages that failed to ship in a multi-package shipment; `success` Boolean – Whether successful; `errorCode` String – Error code; `errorMessage` String – Error description; `extErrorMessage` String – Extended error description; …

<a id="alibabalogisticsopdeliverysendorderdummy-1"></a>
## 106. Ship: no logistics needed

`com.alibaba.logistics:alibaba.logistics.OpDeliverySendOrder.dummy:1` · Messaging · 物流发货-无需物流  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.logistics/alibaba.logistics.OpDeliverySendOrder.dummy/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
For 1688 open-marketplace orders that need no logistics. Supports combined shipping (several orders in one shipment) and shipping at sub-order (orderEntry) level. Shipping by quantity is not supported.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `sendGoods` | OpSendGood[] | yes | Shipment object list |  |
| `sendGoods[].sourceId` | String | yes | Shipment object ID, generally the order ID |  |
| `sendGoods[].sendGoodEntries` | OpSendGoodEntry[] | yes | List of shipment object details |  |
| `sendGoods[].sendGoodEntries[].sourceEntryId` | String | yes | Shipment target line item ID, corresponding to the sub-order ID |  |
| `sendGoods[].sendGoodEntries[].amount` | Long | yes | Actual shipped quantity for the shipment target |  |
| `sendGoods[].sendGoodEntries[].weight` | Double | yes | The actual shipped weight of the shipment object; the weight unit defaults to kilograms. |  |
| `sendGoods[].sendGoodEntries[].extBody` | String | no | JSON array string passed when shipping multiple packages. cpCode is the logistics company code, corresponding to companyNo (not the numeric Id) from the logistics company information API; logisticsCpName: logistics company name, corresponding to companyName from the logistics company information API; mailNo: waybill number; this data can be queried via the &quot;Logistics Company List - Self-linked Logistics&quot; API. quantity: number of shipment objects in the package, required. | [{"cpCode":"SF","logisticsCpName":"顺丰","mailNo":"a1231","quantity":1},{"cpCode":"SF","logi… |
| `remarks` | String | no | Remark |  |
| `gmtSend` | java.util.Date | no | Shipping time | 标准时间格式：yyyyMMddHHmmssSSSZ，例如：20120801154220368+0800 |
| `extBody` | String | yes | JSON string. noLogisticsCondition in extBodyJson is required, value is a string from 1 to 5: "1": other third-party logistics, small logistics provider, fleet, etc. (noLogisticsName and noLogisticsTel required); "2": supplementary freight, price difference (noLogisticsBillNo required); "3": seller delivery (noLogisticsName and noLogisticsTel required); "4": buyer self-pickup; "5": other reasons (remarks required). Other fields have different required-ness depending on the value of noLogisticsCondition. Field descriptions: reason no logistics needed: noLogisticsCondition; name when no logistics needed: noLogisticsName; phone when no logistics needed: noLogisticsTel; waybill number when no logistics needed, meaning varies by the reason: noLogisticsBillNo; shipping proof file list: noLogisticsFiles. | {"noLogisticsBillNo":"111111111111111111","noLogisticsCondition":"3","noLogisticsName":"张三… |
| `extParam` | String | no | {} | JSON 字符串 |
| `receiverInfo` | OpReceiveContacter | no | Shipping address | 优先级大于订单收货地址，为空时，使用订单收货地址 |
| `receiverInfo.provinceCode` | String | yes | Province code |  |
| `receiverInfo.cityCode` | String | yes | City code |  |
| `receiverInfo.areaCode` | String | yes | Region code |  |
| `receiverInfo.townCode` | String | yes | Town or street code |  |
| `receiverInfo.province` | String | yes | Province name; can be omitted if code is passed |  |
| `receiverInfo.city` | String | yes | City name; can be omitted if code is passed |  |
| `receiverInfo.area` | String | yes | District name; if code is passed, this can be omitted |  |
| `receiverInfo.town` | String | yes | Town or street name |  |
| `receiverInfo.address` | String | yes | Detailed address |  |
| `receiverInfo.fullName` | String | yes | Name |  |
| `receiverInfo.corpName` | String | yes | Company name |  |
| `receiverInfo.post` | String | yes | Postal code |  |
| `receiverInfo.phone` | String | yes | Landline phone number |  |
| `receiverInfo.mobile` | String | yes | Mobile phone number |  |
| `receiverInfo.warehouse` | String | yes | Warehouse |  |
| `receiverInfo.codeType` | String | yes | Address code type; defaults to the Cainiao standard code |  |

**Response (top level)**: `result` OpSendOrderModelResult – Shipping details; `result.logisticsId` String – Logistics number; `result.sendGoods` OpSendGood[] – Shipping details; `success` Boolean – Whether successful; `errorCode` String – Error code; `errorMessage` String – Error description; `extErrorMessage` String – Extended error description; …

<a id="productpodorderdesignget-1"></a>
## 107. Get artwork info for a customization order

`com.alibaba.fenxiao.crossborder:product.podOrderDesign.get:1` · Light Customization · 获取加工定制订单稿件信息  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/product.podOrderDesign.get/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Get the design/artwork information of a processing-and-customization (print-on-demand) order.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `orderDesignParam` | OrderDesignOpenApiParam | yes | Order design query parameters | {"orderId":123} |
| `orderDesignParam.orderId` | java.lang.Long | yes | 1688 order id | 123 |

**Response (top level)**: `result` ResultModel – Return result; `result.success` boolean – Whether successful; `result.code` java.lang.String – Response code; `result.message` java.lang.String – Error message; `result.result` OrderDesignOpenResultModel – Return result; …

<a id="accountsearchgetinfoasync-1"></a>
## 108. Run AI supplier-search task (async)

`com.alibaba.fenxiao.crossborder:account.search.getInfo.async:1` · Suppliers · 异步执行AI找商任务  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/account.search.getInfo.async/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Execute an AI supplier-search task asynchronously.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `request` | SpApiExecuteRequest | yes | Request parameters | {} |
| `request.query` | String | yes | Query statement | 毛巾 |
| `request.spFilterCondition` | SpFilterCondition | no | Filter condition | 过滤条件 |
| `request.spFilterCondition.businessModel` | String | no | Business mode; only supports passing factory (indicates a source factory) | factory |
| `request.spFilterCondition.establishedYearsMin` | Integer | no | Minimum years established (filters for merchants established for ≥ this many years) | 3 |
| `request.spFilterCondition.repeatPurchaseRateMin` | Integer | no | Minimum 30-day repeat purchase rate, range [0, 100] | 50 |
| `request.spFilterCondition.supportCustomization` | Boolean | no | Whether processing/customization is supported; only true can be passed | true |
| `request.spFilterCondition.afterSalesScoreMin` | Double | no | Minimum after-sales experience score, range [0, 5] | 4.0 |
| `request.spFilterCondition.productQualityScoreMin` | Double | no | Minimum product experience score, range [0, 5] | 4.0 |
| `request.spFilterCondition.consultationResponseScoreMin` | Double | no | Minimum inquiry experience score, range [0, 5] | 4.0 |
| `request.spFilterCondition.logisticsTimelinessScoreMin` | Double | no | Minimum logistics experience score, range [0, 5] | 4.0 |

**Response (top level)**: `result` ApiResultExecute – Return message; `result.success` Boolean – Whether the AI product-sourcing task was successfully initiated; `result.result` String – Unique task identifier; required when querying task status; `result.code` String – Status code; `result.message` String – Exception info

<a id="accountsearchgetinforesult-1"></a>
## 109. Query AI supplier-search task result

`com.alibaba.fenxiao.crossborder:account.search.getInfoResult:1` · Suppliers · 查询AI找商任务执行结果  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/account.search.getInfoResult/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Query the execution result of an AI supplier-search task.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `request` | SpApiExecuteStatusRequest | yes | Request parameters | {} |
| `request.taskId` | java.lang.String | yes | Task ID, obtained from the return value of the task execution API | taskIdxxxxxxxxxx |

**Response (top level)**: `result` ApiSpExecuteTaskResult – Return value; `result.success` java.lang.Boolean – Whether execution succeeded; `result.result` SpExecuteTask – AI supplier-sourcing task running status; `result.code` java.lang.String – Status code; `result.message` java.lang.String – Error message; …

<a id="newtoncloudtaskcreate-1"></a>
## 110. Newton Cloud: create long-running task

`com.alibaba.agent:newtoncloud.task.create:1` · Inquiries (Newton Cloud) · 牛顿云-创建长程任务  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.agent/newtoncloud.task.create/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Create a Newton Cloud long-running task. It executes asynchronously and returns taskId/sessionId/status.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `message` | String | yes | The question initiated by the user when starting the task | 你有哪些skill供我调用？ |
| `sessionId` | String | no | 1688- and user-generated conversation-level ID; one sessionId can be linked to multiple taskIds | internal_a1b2c3d4-e5f6-7890-abcd-ef1234567890 |
| `taskId` | String | no | Task-dimension ID generated by 1688 or the user; a taskId always belongs to one conversation message | fdd93e89-2f43-4d58-b5b4-356da6cda5e4 |
| `auto` | Boolean | no | Whether automatic mode is enabled | true/false |
| `fileUrls` | String[] | no | List of URLs for files uploaded by the user | https://selleragent.1688.com/api/seller/knowledge/file/download/xxx |
| `fileNames` | String[] | no | List of file names, corresponding one-to-one with fileUrls | 11111.jpg |
| `model` | String | no | Selectable model | flagship |
| `workflowName` | String | no | Specifies the name of the Workflow to execute directly. When passed, the main Agent's intent recognition and planning process is skipped, and the corresponding Workflow is executed directly. If the name does not exist, the task fails. | 1688-supplychain-procurement-search |

**Response (top level)**: `success` Boolean – Whether the request was successful; `taskId` String – Task-dimension ID generated by 1688 or the user; a taskId always belongs to one conversation message; `sessionId` String – 1688- and user-generated conversation-level ID; one sessionId can be linked to multiple taskIds; `status` String – Task status. There are 6 possible task statuses: INIT (created, the task has entered the queue awaiting execution), RUNNING (currently executing; the Agent is performing reasoning and tool calls), WAIT_SKILL (the Agent has invoked an asynchronous skill; the task is paused awaiting the skill execution result callback), WAIT_USER (the task requires user confirmation or additional input; it is paused awaiting the user to resume it via the resume interface), END (the task completed execution normally and has produced a final result), KILL (the task was terminated; this may be a forced termination caused by the user actively cancelling it, execution timing out, or a system exception). END and KILL are terminal states; once entered, they cannot be changed again.; `error` String – Task failure description; `eagleTraceId` String – Unique request ID; eagleTraceId is used for trace troubleshooting; `errorCode` String – Failure error code, returned only when success=false. Values: RATE_LIMIT_CONCURRENT (concurrency limit reached), RATE_LIMIT_PENDING_TOTAL (total queue size exceeded), RATE_LIMIT_QUEUE_FULL (queue full), RATE_LIMIT_HOURLY (hourly quota exhausted), RATE_LIMIT_DAILY (daily quota exhausted)

<a id="newtoncloudtaskget-1"></a>
## 111. Newton Cloud: query task

`com.alibaba.agent:newtoncloud.task.get:1` · Inquiries (Newton Cloud) · 牛顿云-查询任务服务  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.agent/newtoncloud.task.get/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Newton Cloud task query service.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `taskId` | String | yes | Task-level ID generated by the business side; one task ID always belongs to one conversation ID. | 009e1847-d652-4fa1-be5e-346b9b09ac38 |
| `fromIndex` | Long | no | Starting index for incremental output. Pass 0 on the first call, and the nextIndex returned by the previous call on subsequent calls. | 0 |
| `includeBlocks` | Boolean | no | Whether to return rich text/block structure; recommended to pass true when retrieving chunks | true |

**Response (top level)**: `success` Boolean – Whether the API request succeeded; true for success, false for failure; `taskId` String – Task-level ID generated by the business side; one task ID always belongs to one conversation ID.; `sessionId` String – Conversation-dimension ID generated by the business side; one conversation ID can be linked to multiple task IDs; `taskType` String – Task type. LONG_RUNNING indicates a long-running task; NORMAL indicates a regular conversation.; `status` String – Task status. There are 6 possible task statuses: INIT (created, the task has entered the queue awaiting execution), RUNNING (currently executing; the Agent is performing reasoning and tool calls), WAIT_SKILL (the Agent has invoked an asynchronous skill; the task is paused awaiting the skill execution result callback), WAIT_USER (the task requires user confirmation or additional input; it is paused awaiting the user to resume it via the resume interface), END (the task completed execution normally and has produced a final result), KILL (the task was terminated; this may be a forced termination caused by the user actively cancelling it, execution timing out, or a system exception). END and KILL are terminal states; once entered, they cannot be changed again.; `createdAt` String – Task creation time; `startedAt` String – Task start time; `content` String – Output of the current streaming conversation; only returned when the task is not in a terminal state; `messages` message[] – List of all AI response messages for the current task; only returned when the task reaches a final state; `errorMessage` String – Task failure description; `error` String – Description of this request's failure; `eagleTraceId` String – Unique request ID; eagleTraceId is used for trace troubleshooting; `nextIndex` Long – The fromIndex to pass on the next poll, used to incrementally fetch subsequent output chunks and avoid re-fetching.; `outputStatus` String – Current task output stream status. producing means still generating, done means output complete, error means output error, wait_user means waiting for user interaction.; …

<a id="newtoncloudtasklist-1"></a>
## 112. Newton Cloud: list tasks

`com.alibaba.agent:newtoncloud.task.list:1` · Inquiries (Newton Cloud) · 牛顿云-查询任务列表  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.agent/newtoncloud.task.list/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Query the current user's task list. Returns each task's taskId/sessionId/status/taskType and its creation and completion times.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `pageNo` | Integer | no | Page number; if this parameter is not passed, a full query is performed | 1 |
| `pageSize` | Integer | no | Number of items per page | 20 |

**Response (top level)**: `success` Boolean – Whether the interface call succeeded; `data` taskItem[] – Task list; each item is the detailed information of a task. See Query Task Service for details; `error` String – Task failure description; `eagleTraceId` String – Unique request ID; eagleTraceId is used for trace troubleshooting; `total` Integer – Total number of tasks; …

<a id="newtoncloudtaskkill-1"></a>
## 113. Newton Cloud: terminate task

`com.alibaba.agent:newtoncloud.task.kill:1` · Inquiries (Newton Cloud) · 牛顿云-终止任务  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.agent/newtoncloud.task.kill/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Terminate the task with the given taskId, mark it with the KILL status, and return killed.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `taskId` | String | yes | Task-level ID generated by the business side; one task ID always belongs to one conversation ID. | 009e1847-d652-4fa1-be5e-346b9b09ac3 |
| `reason` | String | yes | Reason for task termination, for record-keeping purposes only | 测试收集样例，主动终止 |

**Response (top level)**: `success` Boolean – Whether the interface call succeeded; `taskId` String – Task-level ID generated by the business side; one task ID always belongs to one conversation ID.; `killed` Boolean – Whether termination was successful; `error` String – Task failure description; `eagleTraceId` String – Unique request ID; eagleTraceId is used for trace troubleshooting; `message` String – Termination result description. Returns "the queued task has been canceled" when a queued task is canceled, "the task has been terminated" when a running task is interrupted, and the reason when it fails.

<a id="newtoncloudtaskfetch-1"></a>
## 114. Newton Cloud: query task table

`com.alibaba.agent:newtoncloud.task.fetch:1` · Inquiries (Newton Cloud) · 牛顿云-查询任务表格服务  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.agent/newtoncloud.task.fetch/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Newton Cloud task-table query service.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `taskId` | String | yes | Original task ID | 80c230b3-0cd1-4094-a223-27cc257bc8e0 |
| `id` | String | yes | Table instance ID | 69ec9bdcc878431c873f68e953c30e31_2631391132_20260714211145 |
| `scene` | String | yes | Use the value from complex_table.content; in procurement scenarios this is usually newton | newton |
| `stage` | String | no | Stage, e.g. recall / inquiry | recall |
| `subScene` | String | yes | Use the value from complex_table.content; the procurement scenario is usually purchase | purchase |
| `pageNo` | Integer | no | Requested page number, defaults to the first page | 1 |
| `pageSize` | Integer | no | Requested page size, default 10 | 10 |

**Response (top level)**: `result` result – Service-side response body, including whether the request succeeded, table data, error information, etc.; `result.success` java.lang.Boolean – Whether the request succeeded; true for success, false for failure; `result.error` java.lang.String – Error code when the request fails. INVALID_REQUEST=missing parameter, TASK_NOT_FOUND=task does not exist or no permission, UNSUPPORTED_IN_OPENAPI=sub-scenario not yet supported, CREDENTIAL_UNAVAILABLE=failed to obtain credential, UPSTREAM_TIMEOUT=upstream timeout, UPSTREAM_ERROR=upstream exception, INTERNAL_ERROR=internal error. Detailed error description when the request fails; `result.message` java.lang.String – Error detail description when the request fails; `result.data` java.lang.String – Table data JSON string, needs to be used after JSON.parse. Contains result (product array), total (total count), fieldSpec (field name and type definitions, can be used to dynamically render table columns); `result.eagleTraceId` java.lang.String – Unique request ID, used for tracing

<a id="newtoncloudtaskresume-1"></a>
## 115. Newton Cloud: resume task

`com.alibaba.agent:newtoncloud.task.resume:1` · Inquiries (Newton Cloud) · 牛顿云-恢复任务  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.agent/newtoncloud.task.resume/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Resume a Newton Cloud task.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `sessionId` | String | yes | sessionId | chat_03575ef4 |
| `scene` | String | yes | scene | open_api |
| `taskId` | String | yes | taskId | 68c83ae1-334a-46e2-9d26-e2db37cd1ba4 |
| `eventField` | resumeEventObject | no | User-selected content; choose either this or selectedData — selectedData is recommended | {"type":"EXTERNAL_EXECUTION_RESULT","executionResults":[{"type":"tool_result","id":"call_3… |
| `eventField.type` | String | yes | type | EXTERNAL_EXECUTION_RESULT |
| `eventField.executionResults` | executionResults[] | yes | User selection content | [{"type":"tool_result","id":"call_fab53346e8fe4103b726da70","name":"show_interaction","out… |
| `eventField.executionResults[].type` | String | yes | tool_result | tool_result |
| `eventField.executionResults[].id` | String | yes | Result id | call_9f0a309daaa045b2844ee563 |
| `eventField.executionResults[].name` | String | yes | show_interaction | show_interaction |
| `eventField.executionResults[].state` | String | yes | Status | interrupted |
| `eventField.executionResults[].output` | output[] | yes | Selected content | [{"type":"text","text":"{\"skipped\":true}","id":"ir_1784020010765"}] |
| `eventField.executionResults[].output[].type` | String | yes | Type | text |
| `eventField.executionResults[].output[].text` | String | yes | Content | {\"skipped\":true} |
| `eventField.executionResults[].output[].id` | String | yes | Option ID | ir_1784020010765 |
| `selectedData` | questionAndSelected[] | no | User selection content | [{"question":"请问您需要采购多少件西装？","selected":"10件"},{"question":"请问您的预算范围是多少？","selected":"50-1… |
| `selectedData[].question` | String | yes | Question | 请问您需要采购多少件西装？ |
| `selectedData[].selected` | String | yes | Selected content | 10件 |
| `skipped` | Boolean | no | Whether to skip; true means skip | true |
| `userInput` | String | no | User input requirement | 我想要便宜一些的 |

**Response (top level)**: `result` result – Return result; `result.success` java.lang.Boolean – Whether successful; `result.taskId` java.lang.String – Task ID; `result.status` java.lang.String – Status after resuming; normally RUNNING; `result.message` java.lang.String – Result description; `result.error` java.lang.String – Error description on failure; `result.eagleTraceId` java.lang.String – Trace ID; `result.errorCode` String – Failure error code. If this API is called while the task is still queued (status=QUEUED), it returns INVALID_REQUEST. In this case, poll task.get first and wait for the task to enter RUNNING before calling resume.

<a id="newtoncloudbatchinquirygetresult-1"></a>
## 116. Newton Cloud: query batch-inquiry results

`com.alibaba.agent:newtoncloud.batchInquiry.getResult:1` · Inquiries (Newton Cloud) · 牛顿云-查询批量询盘结果  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.agent/newtoncloud.batchInquiry.getResult/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Query inquiry results by batch-inquiry task ID (wwTaskId). After the ISV receives an AGENT_NEWTON_CLOUD_TASK_NOTIFY notification (stage=BATCH_INQUIRY, stageDetail.phase=COMPLETE), call this interface with the stageDetail.wwTaskId from the notification. It returns structured inquiry results synchronously, including supplier reply summaries, quotation details and recommendation reasons.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `wwTaskId` | String | no | Batch inquiry task ID | 1784790418611 |
| `taskId` | String | no | Newton Cloud task ID; pass this in when there is no wwTaskId, and the API automatically queries the inquiry result based on the Newton Cloud task ID | e0e5e0c0-6911-4b5a-9082-140474666exxx |

**Response (top level)**: `success` Boolean – Whether the request was successful; `data` String – Inquiry result JSON string; `eagleTraceId` String – Unique request ID; eagleTraceId is used for trace troubleshooting; `error` String – Query failure description; `inquiryStatus` String – Batch inquiry task status

<a id="newtoncloudfileupload-1"></a>
## 117. Newton Cloud: upload file

`com.alibaba.agent:newtoncloud.file.upload:1` · Inquiries (Newton Cloud) · 牛顿云-上传文件  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.agent/newtoncloud.file.upload/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Upload a local file to Newton Cloud OSS. Returns a temporary public download link and a stable download address, which can be used directly as the fileUrls parameter of newtoncloud.task.create. File content is passed as Base64; a single file of 3 MB or less is recommended.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `contentBase64` | java.lang.String | yes | Base64-encoded string of the file content. Supports standard and URL-safe encoding, and a data URI prefix; the original file is recommended not to exceed 3MB. | iVBORw0KGgoAAAANSUhEUgAA... |
| `filename` | String | yes | File name, must include the extension | 报价单.xlsx |

**Response (top level)**: `result` NewtonCloudUploadResult – Newton Cloud file upload result; `result.downloadUrl` java.lang.String – Stable download URL, valid long-term; requires logging into the corresponding 1688 account to access; `result.eagleTraceId` java.lang.String – Trace ID; `result.fileSize` long – File size, unit: byte; `result.relativePath` java.lang.String – The relative path of the file in the Newton Cloud workspace, which can be used for subsequent operations such as requesting a download link; `result.agentDownloadUrl` String – Login-free public network download address, which can be used directly for the fileUrls of task.create.

<a id="newtoncloudmodellist-1"></a>
## 118. Newton Cloud: list available model tiers

`com.alibaba.agent:newtoncloud.model.list:1` · Inquiries (Newton Cloud) · 牛顿云-查询可用模型档位列表  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.agent/newtoncloud.model.list/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Query the model tiers currently available on Newton Cloud. Returns each tier's code (id), display name (displayName) and other information. The tier code can be used as the model parameter of newtoncloud.task.create. No business parameters; only access_token and signature are required.

**Request body**: no application parameters.

**Response (top level)**: `success` Boolean – Whether the request was successful; `error` String – Error description on failure; `models` model[] – List of available model tiers. The returned content is driven by platform configuration and only includes tiers open to the public; …

<a id="newtoncloudpointsquery-1"></a>
## 119. Newton Cloud: query points details

`com.alibaba.agent:newtoncloud.points.query:1` · Inquiries (Newton Cloud) · 牛顿云-查询积分详情  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.agent/newtoncloud.points.query/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Query the Newton Cloud points information of the currently authorized account, including total points, used points, available points, points sources, expiry information and paginated usage details. The user is identified automatically from the access_token; no userId is needed.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `pageNo` | Integer | no | Page number for points consumption details, starting from 1 | 1 |
| `pageSize` | Integer | no | Number of records per page, maximum 100 | 20 |
| `startTime` | String | no | Query start time, format yyyy-MM-dd HH:mm:ss | 2026-08-01 00:00:00 |
| `endTime` | String | no | Query end time, format yyyy-MM-dd HH:mm:ss | 2026-08-12 23:59:59 |

**Response (top level)**: `result` pointsQueryResponse – Newton points query result; `result.success` Boolean – Whether the query was successful; `result.availableCredits` BigDecimal – Remaining points in the current account; `result.pageNo` Integer – Current page number; `result.pageSize` Integer – Records per page; `result.total` Integer – Total number of conversations; `result.records` pointsSessionUsage[] – Conversation points consumption record; `result.msgInfo` String – Error description when the query fails; `result.eagleTraceId` String – Request trace ID; …

<a id="sytbuyerdraftpurchaseorder-2"></a>
## 120. 88 ShengYiTong: buyer drafts purchase order

`com.alibaba.syt:syt.buyer.draftPurchaseOrder:2` · 88 ShengYiTong (Business Link) · 88生意通买家起草采购单  
POST `https://gw.open.1688.com/openapi/param2/2/com.alibaba.syt/syt.buyer.draftPurchaseOrder/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
88 ShengYiTong solution. The buyer drafts a purchase order.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `request` | ContractDraftAndSignApplyApiRequest | yes | Request parameters | {     "origin": "ERP",     "requestNo": "REQ2026051100001",     "contractType": "PURCHASE_… |
| `request.origin` | java.lang.String | yes | Call source | ERP，固定值 |
| `request.requestNo` | java.lang.String | yes | Request idempotency key | 本次创建采购单唯一请求号 |
| `request.draftNo` | java.lang.String | yes | Purchase order number - leave empty when creating | 空，不传值 |
| `request.payTerm` | ApiContractPayTerm | yes | Payment terms | {     "amount": 10000.00,     "payMethod": "SINGLE",     "payTimeType": "IMMEDIATELY",    … |
| `request.payTerm.amount` | BigDecimal | yes | Payment amount, unit: yuan, currency: RMB (CNY) | 1 |
| `request.payTerm.payMethod` | java.lang.String | yes | Payment method - fixed value | SINGLE |
| `request.payTerm.payTimeType` | java.lang.String | yes | Payment time type - fixed value | IMMEDIATELY |
| `request.payTerm.payTimeContent` | java.lang.String | no | Payment time content | 不传 |
| `request.contentTerm` | ApiContractContentTerm | yes | Content terms | {     "contentType": "PURCHASE_ORDER",     "purchaseItems": [       {         "key": "SKU0… |
| `request.contentTerm.contentType` | java.lang.String | yes | Content type - fixed value | PURCHASE_ORDER |
| `request.contentTerm.purchaseItems` | ApiPurchaseItem[] | yes | Purchase order details | [       {         "key": "SKU001",         "productName": "商品A",         "productSpec": "红… |
| `request.contentTerm.purchaseItems[].key` | java.lang.String | yes | Purchase order sequence number | sku001 |
| `request.contentTerm.purchaseItems[].productName` | java.lang.String | yes | Product name | 水杯 |
| `request.contentTerm.purchaseItems[].productSpec` | java.lang.String | yes | Product sku | 1L、白色、大肚款 |
| `request.contentTerm.purchaseItems[].quantity` | java.lang.Integer | yes | Purchase quantity | 1 |
| `request.contentTerm.purchaseItems[].unitPrice` | BigDecimal | yes | Unit price, in yuan | 1 |
| `request.contentTerm.purchaseItems[].subtotal` | BigDecimal | yes | Total price, in yuan | 1 |
| `request.contractRoleInfoTerm` | ApiContractRoleInfoTerm | yes | Counterparty agreement | {     "counterpartyOrigin": "LOGIN_ID_1688_MATCH",     "counterpartyLoginId": "seller1688"… |
| `request.contractRoleInfoTerm.counterpartyOrigin` | java.lang.String | yes | LOGIN_ID_1688_MATCH (specify the counterparty via 1688 loginId), DRAFTER_INPUT (specify via the counterparty's entity information), COUNTERPARTY_INPUT (not specified, left for the counterparty to claim themselves) | LOGIN_ID_1688_MATCH |
| `request.contractRoleInfoTerm.counterpartyLoginId` | java.lang.String | no | 1688loginId, required when counterpartyOrigin = LOGIN_ID_1688_MATCH | ces测试002 |
| `request.contractRoleInfoTerm.counterpartyName` | java.lang.String | no | Counterparty name. Required when counterpartyOrigin = DRAFTER_INPUT | 张三 |
| `request.contractRoleInfoTerm.counterpartyLicenseNo` | java.lang.String | no | Counterparty's ID card number or business license number. Required when counterpartyOrigin = DRAFTER_INPUT | 330xx |
| `request.contractRoleInfoTerm.counterpartyLicenseType` | java.lang.String | no | ID document type; required when counterpartyOrigin = DRAFTER_INPUT | 统一社会信用代码UNITY |
| `request.drafterNeedSignConfirmAgain` | Boolean | yes | Whether the drafting party needs secondary confirmation; false means not needed | false |
| `request.postFeeTerm` | ApiContractPostFeeTerm | yes | Postage terms | {   "postFee": 1.00 } |
| `request.postFeeTerm.postFee` | BigDecimal | yes | postFee | 0.01 |
| `request.confirmTerm` | ApiContractConfirmTerm | yes | Confirm receipt terms | {   "contractConfirmType": "AUTO_CONFIRM",   "autoConfirmCondition": {     "conditionEnum"… |
| `request.confirmTerm.contractConfirmType` | String | yes | Buyer manually confirms receipt: BUYER_MANUAL_CONFIRM; automatically confirms receipt: AUTO_CONFIRM. When this value is passed, autoConfirmCondition cannot be empty; | BUYER_MANUAL_CONFIRM |
| `request.confirmTerm.autoConfirmCondition` | ApiContractAutoConfirmCondition | yes | Auto-confirm receipt condition; cannot be empty when contractConfirmType is AUTO_CONFIRM | {   "contractConfirmType": "AUTO_CONFIRM",   "autoConfirmCondition": {     "conditionEnum"… |
| `request.confirmTerm.autoConfirmCondition.conditionEnum` | String | yes | Condition enum | PAY_SUCCESS_TIME_PLUS_X_DAYS |
| `request.confirmTerm.autoConfirmCondition.xDay` | Integer | yes | Number of days | 1 |

**Response (top level)**: `result` ContractDraftAndSignApplyApiResponse – Affected parameters; `result.draftNo` java.lang.String – Purchase order number; `result.contractCurrentStatus` java.lang.String – Current status; `result.success` java.lang.Boolean – Whether successful; `result.errorCode` java.lang.String – Error code; `result.errorMsg` java.lang.String – Error message; `result.traceId` String – Call trace; `result.counterpartyConfirmUrl` String – Not empty when the counterparty needs to claim it; returns the claim URL

<a id="sytcontractconfirm-1"></a>
## 121. 88 ShengYiTong: confirm transaction complete

`com.alibaba.syt:syt.contract.confirm:1` · 88 ShengYiTong (Business Link) · 88生意通确认交易完成  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.syt/syt.contract.confirm/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
88 ShengYiTong: confirm that the transaction is complete.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `request` | ContractConfirmApiRequest | yes | Request parameters | {"draftNo":"88SYT20251204069031"} |
| `request.draftNo` | java.lang.String | yes | Contract number | 88SYT20251204069031 |

**Response (top level)**: `result` ContractConfirmApiResponse – Return response; `result.draftNo` java.lang.String – Contract number; `result.contractCurrentStatus` java.lang.String – Current latest contract status; `result.success` Boolean – Call result; `result.errorCode` String – Error code; `result.errorMsg` String – Error message; `result.traceId` String – Call traceId

<a id="sytcontractinvalid-1"></a>
## 122. 88 ShengYiTong: void contract

`com.alibaba.syt:syt.contract.invalid:1` · 88 ShengYiTong (Business Link) · 88生意通合同作废  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.syt/syt.contract.invalid/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
88 ShengYiTong: void (invalidate) a contract.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `request` | ContractInvalidApiRequest | yes | Request parameters | {   "draftNo": "DRAFT20251230001" } |
| `request.draftNo` | java.lang.String | yes | Contract number | 88SYT20251201169011 |

**Response (top level)**: `result` ContractInvalidApiResponse – Response parameters; `result.draftNo` java.lang.String – Contract number; `result.contractCurrentStatus` java.lang.String – Current contract status; `result.success` Boolean – Call result; `result.errorCode` String – Error message; `result.errorMsg` String – Error description; `result.traceId` String – traceId. Must be recorded and provided to the platform for troubleshooting.

<a id="sytcontractrefund-1"></a>
## 123. 88 ShengYiTong: contract refund request

`com.alibaba.syt:syt.contract.refund:1` · 88 ShengYiTong (Business Link) · 88生意通合同退款申请  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.syt/syt.contract.refund/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
88 ShengYiTong: apply for a contract refund.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `request` | ContractRefundApplyApiRequest | yes | Refund application parameters | {   "draftNo": "DRAFT20251230001",   "buyerId": "buyer123456",   "amount": 100000,   "requ… |
| `request.draftNo` | java.lang.String | yes | Contract number | 88SYT20251201169011 |
| `request.amount` | BigDecimal | yes | Refund amount, in yuan, currency RMB (CNY) | 1 |
| `request.requestNo` | java.lang.String | yes | Refund request idempotency key, identifies a unique request | requestNo123abc |
| `request.payOrderNo` | java.lang.String | yes | Payment order number, from the payment interface | 8899x |
| `request.refundReason` | java.lang.String | yes | Refund application reason | 已经线下沟通退款10元 |

**Response (top level)**: `result` ContractRefundApplyApiResponse – Response; `result.refundOrderNo` java.lang.String – Refund order number; `result.success` Boolean – Whether the call succeeded, true for success, false for failure; `result.errorMsg` String – Call description; `result.errorCode` String – Call error code; `result.traceId` String – Call trace ID, for troubleshooting

<a id="sytcontractpay-1"></a>
## 124. 88 ShengYiTong: buyer payment, get cashier URL

`com.alibaba.syt:syt.contract.pay:1` · 88 ShengYiTong (Business Link) · 88生意通 买家支付获取收银台 URL  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.syt/syt.contract.pay/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
88 ShengYiTong solution.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `request` | ContractPayApiRequest | yes | Payment request parameters | {   "draftNo": "DRAFT20251230001",   "userId": "user123456",   "payChannel": "ALIPAY",   "… |
| `request.draftNo` | java.lang.String | yes | Contract draft | 88SYT20251204069031 |
| `request.amount` | BigDecimal | yes | Payment amount, unit: yuan, currency: RMB (CNY) | 1 |

**Response (top level)**: `result` ContractPayApiResponse – Response; `result.success` java.lang.Boolean – Whether successful; `result.errorCode` java.lang.String – Error code; `result.errorMsg` java.lang.String – Error description; `result.cashierUrl` String – Cashier URL

<a id="sytcontractpaytransfer-1"></a>
## 125. 88 ShengYiTong: WorldFirst transfer payment (WorldFirst B2C accounts only)

`com.alibaba.syt:syt.contract.payTransfer:1` · 88 ShengYiTong (Business Link) · 88生意通万里汇转账支付-仅支持万里汇B2C账号  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.syt/syt.contract.payTransfer/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
88 ShengYiTong payment-transfer interface.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `request` | ContractPayTransferApiRequest | yes | Request input parameter | {   "draftNo": "DRAFT20251230001",   "userId": "user123456",   "payOrderNo": "PAY202512300… |
| `request.draftNo` | java.lang.String | yes | Contract drafting number | 88SYT20251204069031 |
| `request.payOrderNo` | java.lang.String | yes | Payment order number | 88999xxx |
| `request.escrowAccount` | BankAccount | yes | Escrow account the funds are transferred into | {   "accountName": "张三",   "accountNo": "6228765030160001109",   "bankName": "江苏银行股份有限公司" … |
| `request.escrowAccount.cardHolderType` | java.lang.String | yes | Payee type, enterprise COMPANY or individual PERSON | COMPANY |
| `request.escrowAccount.accountNumber` | java.lang.String | yes | Payee account | 6228765030160001109 |
| `request.escrowAccount.holder` | java.lang.String | yes | Payee name | 张三 |
| `request.transferToAmount` | BigDecimal | yes | Transferred fund amount, equal to the contract amount, unit: yuan, currency: RMB CNY | 1 |
| `request.transferFromCurrency` | java.lang.String | yes | Source currency, US Dollar USD | USD |
| `request.transferFundChannel` | java.lang.String | yes | Funding channel, specified value | WORLD_FIRST |
| `request.transferFundChannelConfig` | ApiWfTransferFundChannelConfig | yes | WorldFirst customer configuration | {"clientId":"xxx","customerId":"万里汇客户 ID"} |
| `request.transferFundChannelConfig.clientId` | java.lang.String | yes | Wanlihui (WorldFirst) customer ID | 11 |
| `request.transferFundChannelConfig.customerId` | java.lang.String | yes | Wanlihui (WorldFirst) customer ID | 11 |
| `request.transferFundChannelConfig.privateKeyCiphertext` | java.lang.String | yes | WorldFirst customer key ciphertext | 11 |
| `request.transferFundChannelConfig.privateKeyVersion` | java.lang.String | yes | WorldFirst customer key version | 1 |

**Response (top level)**: `result` ContractPayTransferApiResponse – Response result; `result.escrowAccount` BankAccount – Fund transfer-in account information; `result.transferToCurrency` java.lang.String – Target currency; `result.transferToAmount` java.lang.Long – Target amount, contract amount, unit: cent; `result.transferFromAmount` java.lang.Long – Transfer-out amount, unit: US cents, USD; `result.transferFromCurrency` java.lang.String – Payment currency; `result.quotePrice` java.math.BigDecimal – Exchange rate; `result.feeAmount` java.lang.Long – Fee amount, in US cents; `result.feeCurrency` java.lang.String – Fee currency; `result.success` String – Whether the call succeeded, true for success, false for failure; `result.errorCode` String – Call code; `result.errorMsg` String – Call result description; `result.traceId` String – traceId, for troubleshooting; …

<a id="sytbuyerconfirmpurchaseorder-1"></a>
## 126. 88 ShengYiTong: buyer confirms purchase order

`com.alibaba.syt:syt.buyer.confirmPurchaseOrder:1` · 88 ShengYiTong (Business Link) · 88生意通买家确认采购单  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.syt/syt.buyer.confirmPurchaseOrder/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
88 ShengYiTong: the buyer confirms a purchase order.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `request` | SignConfirmApiRequest | yes | Request parameters | {   "draftNo": "CT2026051100001" } |
| `request.draftNo` | java.lang.String | yes | Purchase order number | 88SYT123456 |

**Response (top level)**: `result` SignConfirmApiResponse – Response result; `result.draftNo` java.lang.String – Purchase order number; `result.errorCode` java.lang.String – Call result code; relevant when isSuccess=true; `result.errorMsg` java.lang.String – Call result description; relevant when isSuccess=true; `result.traceId` String – traceId; `result.isSuccess` java.lang.Boolean – true

<a id="sytbuyerquerycontractdetail-1"></a>
## 127. 88 ShengYiTong: buyer queries purchase-order details

`com.alibaba.syt:syt.buyer.queryContractDetail:1` · 88 ShengYiTong (Business Link) · 88生意通买家查询采购单详情  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.syt/syt.buyer.queryContractDetail/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
88 ShengYiTong: the buyer queries the details of a purchase order.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `request` | ContractQueryDetailApiRequest | yes | Request | {     "draftNo": "CT2026051100001"     } |
| `request.draftNo` | java.lang.String | yes | Purchase order number | CT2026051100001 |

**Response (top level)**: `result` ContractQueryDetailApiResponse – Response; `result.contractDTO` ContractApiDTO – Purchase order details; `result.errorCode` java.lang.String – Error code; `result.errorMsg` java.lang.String – Call result description; `result.isSuccess` java.lang.Boolean – Whether successful; `result.traceId` String – Call traceId; `result.fundPayOrderList` FundPayOrderApiDTO[] – List of payment orders; …

<a id="sytcustomerqueryuserauthstatus-1"></a>
## 128. 88 ShengYiTong: check whether customer is verified

`com.alibaba.syt:syt.customer.queryUserAuthStatus:1` · 88 ShengYiTong (Business Link) · 88生意通查询客户是否已认证  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.syt/syt.customer.queryUserAuthStatus/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
88 ShengYiTong: check whether a customer has completed verification.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `request` | UserAuthStatusQueryApiRequest | yes | Input parameters | {   "loginId": "user1688" } |
| `request.loginId` | java.lang.String | yes | 1688 loginId | ces测试002 |

**Response (top level)**: `result` UserAuthStatusQueryApiResponse – Output result; `result.authenticated` java.lang.Boolean – Whether real-name verification has been completed; `result.errorCode` java.lang.String – Error code, to be noted when isSuccess=true; `result.errorMsg` java.lang.String – Error message description; relevant when isSuccess=true; `result.isSuccess` java.lang.Boolean – Whether the call succeeded; `result.traceId` String – traceId

<a id="tradeinvoicetitlegetpagelist-1"></a>
## 129. Query buyer invoice titles (paginated)

`com.alibaba.trade:trade.invoiceTitle.getPageList:1` · Invoicing · 分页查询买家抬头列表  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/trade.invoiceTitle.getPageList/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Paginated query of the buyer's invoice titles.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `input` | OpInvoiceTitleQueryParam | yes |  |  |
| `input.page` | int | no | Current page, cannot be less than 1 | 1 |
| `input.pageSize` | int | no | Page size; cannot be less than 1 or greater than 100 | 50 |
| `input.titleType` | java.lang.String | no | Invoice title type. Both individual and social organization map to PERSONAL: PERSONAL(0,&quot;individual&quot;),COMPANY(1,&quot;company&quot;),SOCIAL_ORGANIZATION(3,&quot;social organization (government agency, public institution, etc.)&quot;),PARENT_VIRTUAL(-1,&quot;parent invoice virtual type&quot;); | PERSONAL |

**Response (top level)**: `result` OpInvoiceTitlePageSingleResult – ; `result.currentPage` java.lang.Integer – Current page number; `result.errorCode` java.lang.String – Error code; `result.errorInfo` java.lang.String – Error message; `result.pageSize` java.lang.Integer – Page size; `result.result` OpInvoiceTitleQueryResult – Return result; `result.success` boolean – Whether successful; `result.totalNum` java.lang.Long – Total record count; …

<a id="tradeinvoiceamountgetlist-1"></a>
## 130. Query invoiceable amount of orders

`com.alibaba.trade:trade.invoiceAmount.getList:1` · Invoicing · 查询订单可开票金额  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/trade.invoiceAmount.getList/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Query the invoiceable amount of orders. To find invoiceable orders first: 1. alibaba.trade.getBuyerOrderList-1: pass needInvoicingSetting=true and check invoicingSettingModel.tradeInvoiceStatus in the response to see whether an order can be invoiced. 2. alibaba.trade.get.buyerView-1: include InvoicingSetting in the includeFields parameter.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `input` | OpOrderInvoiceAmountQueryParam | yes |  |  |
| `input.orderIds` | java.lang.Long[] | yes | List of order IDs | [123, 456] |

**Response (top level)**: `result` OpOrderInvoiceAmountResultModel – ; `result.code` java.lang.String – ; `result.message` java.lang.String – ; `result.result` OpOrderInvoiceAmountQueryResult – ; `result.retCodes` java.lang.String[] – ; `result.subCode` java.lang.String – ; `result.subMessage` java.lang.String – ; `result.success` boolean – ; …

<a id="tradeinvoiceapplygetpagelistbuyerview-1"></a>
## 131. Query invoice applications (buyer view, paginated)

`com.alibaba.trade:trade.invoiceApply.getPageListBuyerView:1` · Invoicing · 分页查询发票申请列表（买家视角）  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/trade.invoiceApply.getPageListBuyerView/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Paginated query of invoice applications (buyer view).

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `input` | OpInvoiceListByPageParam | yes |  |  |
| `input.bizStatusList` | java.lang.String[] | no | Invoice status: ISSUING(1,&quot;Issuing&quot;),ISSUED(2,&quot;Issued&quot;),VERIFYING(5,&quot;Verifying&quot;),VERIFY_FAILED(6,&quot;Verification failed&quot;),RETURNING(10,&quot;Returning&quot;),RETURNED(11,&quot;Returned&quot;),INVALIDING(20,&quot;Invalidating&quot;),DEPRECATED(21,&quot;Invalidated&quot;),RED_ISSUING(30,&quot;Red-flush in progress&quot;),RED_PART_ISSUED(31,&quot;Partially red-flushed&quot;),RED_ALL_ISSUED(32,&quot;Fully red-flushed&quot;),FAILED(40,&quot;Issuance failed&quot;),CLOSED(50,&quot;Closed&quot;),; | ISSUING |
| `input.fuzzyInvoiceTitle` | java.lang.String | no | Fuzzy-matched invoice title | 模糊发票抬头 |
| `input.isRedInvoice` | Boolean | no | Whether it is a red invoice | false |
| `input.orderId` | java.lang.String | no | Transaction order ID | 123 |
| `input.outBizId` | java.lang.String | no | Unique ID of the invoice record (used when querying application records); usually the transaction order ID, but changes for red-letter invoices | 123 |
| `input.page` | int | yes | Current page, can be less than 1 | 1 |
| `input.pageSize` | int | yes | Page size; must not be greater than 100 or less than 1 | 10 |
| `input.createMillTimeStart` | Long | no | Creation start time in milliseconds | 1773849600000 |
| `input.createMillTimeEnd` | Long | no | Creation end time in milliseconds | 1773936000000 |
| `input.modifyMillTimeStart` | Long | no | Modification start time in milliseconds | 1773849600000 |
| `input.modifyMillTimeEnd` | Long | no | Modification end time in milliseconds | 1773936000000 |

**Response (top level)**: `result` OpInvoiceModelPageResult – ; `result.errorCode` java.lang.String – ; `result.errorInfo` java.lang.String – ; `result.pageIndex` int – ; `result.resultList` OpInvoiceModel[] – ; `result.sizePerPage` int – ; `result.success` boolean – ; `result.totalRecords` int – ; …

<a id="tradeinvoicegetlistbuyerview-1"></a>
## 132. Query all issued invoices for an order (buyer view)

`com.alibaba.trade:trade.invoice.getListBuyerView:1` · Invoicing · 查询交易单下关联的所有开具的发票信息（买家视角）  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/trade.invoice.getListBuyerView/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Query all issued invoices associated with an order (buyer view).

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `input` | OpInvoiceByOrderIdQueryParam | yes |  |  |
| `input.orderId` | Long | yes | Order ID | 123 |

**Response (top level)**: `result` OpInvoiceModelListResultModel – ; `result.code` String – ; `result.message` String – ; `result.result` OpInvoiceModel[] – ; `result.retCodes` String[] – ; `result.subCode` String – ; `result.subMessage` String – ; `result.success` Boolean – ; …

<a id="tradeinvoiceapply-1"></a>
## 133. Buyer requests invoice

`com.alibaba.trade:trade.invoice.apply:1` · Invoicing · 买家申请开票  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/trade.invoice.apply/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Buyer requests an invoice; batch requests are supported. Prerequisites: get the buyer's invoice titles; get the invoiceable orders and the invoice types they support; get the orders' invoiceable amounts.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `input` | OpApplyInvoiceParam | yes | Input parameter | {} |
| `input.invoiceApplyModelList` | OpInvoiceApplyModel[] | yes | Invoice application model | {} |
| `input.invoiceApplyModelList[].orderId` | java.lang.Long | yes | Main order number | 123 |
| `input.invoiceApplyModelList[].amount` | java.lang.Long | yes | Total invoice amount including tax, obtained via the trade.invoiceAmount.getList API | 2 |
| `input.invoiceApplyModelList[].invoiceType` | java.lang.String | yes | Invoice type: VATAX_SPEC(2, &quot;VAT special invoice&quot;); VATAX_COMM(1, &quot;VAT general invoice&quot;). | VATAX_COMM |
| `input.invoiceApplyModelList[].purchaserInvoiceTitleModel` | OpInvoiceTitleModel | yes | Buyer's invoice title, obtained via the trade.invoiceTitle.getPageList API; pass the retrieved title through as-is. | {} |
| `input.invoiceApplyModelList[].purchaserInvoiceTitleModel.titleType` | java.lang.String | yes | Invoice title type. Both individuals and social organizations use PERSONAL: PERSONAL(0,&quot;Individual and social organization&quot;),COMPANY(1,&quot;Enterprise&quot;); | PERSONAL |
| `input.invoiceApplyModelList[].purchaserInvoiceTitleModel.title` | java.lang.String | yes | Title | 抬头 |
| `input.invoiceApplyModelList[].purchaserInvoiceTitleModel.taxpayerIdentify` | java.lang.String | no | Taxpayer identification number (required for enterprise invoicing) | 123 |
| `input.invoiceApplyModelList[].purchaserInvoiceTitleModel.bankName` | java.lang.String | no | Bank name (required for enterprise special invoice) | 开户行 |
| `input.invoiceApplyModelList[].purchaserInvoiceTitleModel.bankAccountId` | java.lang.String | no | Bank account number (required for enterprises issuing special VAT invoices) | 123 |
| `input.invoiceApplyModelList[].purchaserInvoiceTitleModel.registerAddress` | java.lang.String | no | Enterprise registered address (required for enterprises issuing special VAT invoices) | 企业注册地址 |
| `input.invoiceApplyModelList[].purchaserInvoiceTitleModel.registerPhone` | java.lang.String | no | Company phone number (required for enterprise special invoice) | 123 |

**Response (top level)**: `result` ApplyInvoiceResultModel – Request result; `result.code` String – code; `result.message` String – message; `result.result` OpApplyInvoiceResult – Invoicing result; `result.retCodes` String[] – retCodes; `result.subCode` String – subCode; `result.subMessage` String – subMessage; `result.success` Boolean – Whether successful; …

<a id="tradeinvoiceconsult-1"></a>
## 134. Merged-invoice consultation

`com.alibaba.trade:trade.invoice.consult:1` · Invoicing · 合并开票咨询  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/trade.invoice.consult/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Submit orders to check whether they can be invoiced together and how they are grouped.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `reqDTO` | MergeInvoiceConsultOpenReqDTO | yes | Input parameter | 入参 |
| `reqDTO.orderIdsJsonList` | java.lang.String | yes | List of orders to be consulted (JSON string) | ["123","456","789"] |

**Response (top level)**: `result` SingleResultDTO – Output parameters; `result.errorCode` java.lang.String – Error code; `result.errorDesc` java.lang.String – Error description; `result.isRetry` boolean – Whether to retry; `result.result` MergeInvoiceConsultOpenResDTO – Result object; `result.success` boolean – Whether the service was successful; …

<a id="tradeinvoicemergeapply-1"></a>
## 135. Submit merged-invoice request

`com.alibaba.trade:trade.invoice.mergeapply:1` · Invoicing · 提交合并开票  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/trade.invoice.mergeapply/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
First call the merged-invoice consultation trade.invoice.consult to get the grouping result, then call merged invoicing.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `reqDTO` | MergeInvoiceApplyOpenReqDTO | yes | Input parameter | 入参对象 |
| `reqDTO.invoiceType` | java.lang.String | yes | Invoice type | 1:普票，2：专票 |
| `reqDTO.orderIdGroupsJsonList` | java.lang.String | yes | Merged invoicing group list (each group corresponds to one invoice); each group is a comma-separated string | ["123,456,789","101,202"] |
| `reqDTO.purchaserInvoiceTitle` | InvoiceTitleModel | yes | Buyer invoice title object | 对象 |
| `reqDTO.purchaserInvoiceTitle.bankAccountId` | java.lang.String | no | Bank account | xxx |
| `reqDTO.purchaserInvoiceTitle.bankName` | java.lang.String | no | Name of the account-holding bank | xx银行 |
| `reqDTO.purchaserInvoiceTitle.email` | java.lang.String | no | Email | xxx |
| `reqDTO.purchaserInvoiceTitle.isDefault` | java.lang.String | no | Whether it is the default invoice title | Y：是/N：否 |
| `reqDTO.purchaserInvoiceTitle.name` | java.lang.String | no | Enterprise contact person | xxx |
| `reqDTO.purchaserInvoiceTitle.receiverPhone` | java.lang.String | no | Recipient phone number | xx |
| `reqDTO.purchaserInvoiceTitle.registerAddress` | java.lang.String | no | Enterprise registered address | xxx |
| `reqDTO.purchaserInvoiceTitle.registerPhone` | java.lang.String | no | Enterprise phone | xxx |
| `reqDTO.purchaserInvoiceTitle.taxpayerIdentify` | java.lang.String | no | Tax number | 专票必填 |
| `reqDTO.purchaserInvoiceTitle.title` | java.lang.String | yes | Title | xxx |
| `reqDTO.purchaserInvoiceTitle.titleType` | java.lang.Integer | yes | Invoice type | 0：个人和社会组织，1：企业 |

**Response (top level)**: `result` SingleResultDTO – Output parameters; `result.errorCode` java.lang.String – Error code; `result.errorDesc` java.lang.String – Error description; `result.isIdempotent` java.lang.Boolean – Idempotent; `result.isRetry` boolean – Whether retry is supported; `result.result` MergeInvoiceApplyOpenResDTO – Object; `result.success` boolean – Whether the service was successful; …

<a id="tradeinvoicesellerqueryrelatedorders-1"></a>
## 136. Query merged-invoice relationships by order or invoice application

`com.alibaba.trade:trade.invoice.sellerqueryrelatedorders:1` · Invoicing · 基于交易单或发票申请单查询合单关联关系  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/trade.invoice.sellerqueryrelatedorders/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
The seller queries merged-invoice relationships by order ID (orderId) or invoice-application ID (outbizid). If both are passed, outbizid takes precedence. If the result is not a merged invoice, mergeInvoice returns false.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `reqDTO` | QueryMergeInvoiceRelatedOrdersOpenReqDTO | yes | Input parameter | 入参对象 |
| `reqDTO.orderId` | java.lang.String | yes | Transaction order | xxx |
| `reqDTO.outBizId` | java.lang.String | yes | Unique key outbizid of the invoice application order | xxx |

**Response (top level)**: `result` SingleResultDTO – Output parameters; `result.errorCode` java.lang.String – Error code; `result.errorDesc` java.lang.String – Error reason; `result.isRetry` boolean – Whether retry is supported; `result.result` QueryMergeInvoiceRelatedOrdersOpenResDTO – Output parameters; `result.success` boolean – Whether the service was successful; …

<a id="consignmentcostatussync-1"></a>
## 137. Warehouse receipt / shelving of goods

`com.alibaba.fenxiao.crossborder:consignment.co.status.sync:1` · Fully Managed (Consignment) · 仓库签收/上架商品  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/consignment.co.status.sync/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (optional) · user authorization not required · signature required  
Open to Buffalo. Syncs the Buffalo warehouse's receipt or shelving of goods to the 1688 shipment-order status.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `bizInboundOrderId` | java.lang.String | yes | Inbound (warehouse receipt) order id |  |
| `status` | java.lang.String | yes | signed/onShelves |  |
| `processTime` | java.lang.Long | yes | Timestamp |  |
| `num` | String | yes | Quantity |  |

**Response (top level)**: `result` ResultModel – ; `result.msg` java.lang.String – Error message; `result.traceId` java.lang.String – traceId; `result.success` java.lang.Boolean – true/false; `result.model` java.lang.Object – null

<a id="consignmenttallycreate-1"></a>
## 138. Warehouse creates discrepancy tally sheet

`com.alibaba.fenxiao.crossborder:consignment.tally.create:1` · Fully Managed (Consignment) · 仓库创建差异理货单  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/consignment.tally.create/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (optional) · user authorization not required · signature required  
The Buffalo warehouse creates a discrepancy tally sheet.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `bizInboundOrderId` | java.lang.String | yes | Inbound (warehouse receipt) order id |  |
| `status` | java.lang.String | yes | tallysheet |  |
| `processTime` | java.lang.Long | yes | Timestamp |  |
| `tallySheetId` | java.lang.String | yes | Warehouse discrepancy order ID |  |
| `tallySheetReason` | java.lang.String | yes | Reason for tally discrepancy |  |
| `inboundCount` | java.lang.String | yes | Inbound quantity |  |
| `varianceCount` | java.lang.String | yes | Discrepancy quantity |  |

**Response (top level)**: `result` ResultModel – ; `result.msg` java.lang.String – Error message; `result.traceId` java.lang.String – traceId; `result.success` java.lang.String – true/false; `result.model` java.lang.Object – null

<a id="consignmentcoinventory-1"></a>
## 139. Update warehouse product inventory

`com.alibaba.fenxiao.crossborder:consignment.co.inventory:1` · Fully Managed (Consignment) · 仓库商品库存更新  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/consignment.co.inventory/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (optional) · user authorization not required · signature required  
Update the inventory of products in the warehouse.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `bizSkuId` | java.lang.String | yes | External skuId | 12212122 |
| `sourceItemId` | java.lang.String | yes | Warehouse product ID | 12312323123 |
| `outboundQuantity` | java.lang.String | yes | Outbound quantity | 12 |
| `bizId` | java.lang.String | yes | Business ID, used to ensure idempotency of stock deduction | 12321312 |

**Response (top level)**: `result` ResultModel – ; `result.msg` java.lang.String – Info; `result.traceId` java.lang.String – Trace id; `result.success` java.lang.Boolean – Success/Failure; `result.model` java.lang.Object – Return result

<a id="publishresultcallback-1"></a>
## 140. Publish result callback

`com.alibaba.fenxiao.crossborder:publish.result.callback:1` · Listing (Publishing) · 发布结果回调  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/publish.result.callback/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
WB-1688 integration Method 3: publish-result callback interface.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `request` | WbPublishCallbackRequestDTO | yes | Callback request parameters | {"offerId":123456789,"imtId":"imt123","status":"SUCCESS","requestId":"req123","cards":[]} |
| `request.offerId` | java.lang.Long | yes | 1688 product ID | 123456789 |
| `request.imtId` | java.lang.String | yes | WB product IMT ID | imt123 |
| `request.status` | java.lang.String | yes | Publish status | SUCCESS |
| `request.requestId` | java.lang.String | yes | Request ID | req123 |
| `request.cards` | WbCardResult[] | yes | Card result list | [] |

**Response (top level)**: `result` ResultGT52f0rf – Callback result; `result.success` java.lang.Boolean – Whether successful; `result.result` WbPublishCallbackResponseDTO – Return result; `result.code` java.lang.String – Error code; `result.permissionName` java.lang.String – Permission name; `result.message` java.lang.String – Error message; …

<a id="publishproductlist-1"></a>
## 141. Listable product list

`com.alibaba.fenxiao.crossborder:publish.product.list:1` · Listing (Publishing) · 铺货商品列表  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/publish.product.list/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Query the list of products available for listing. Supports paginated queries of product-ID sets by assortment (goods-pool) ID.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `param` | WbPublishableOfferQueryParam | yes | Request parameters for querying listable products | {"palletId":"example","pageNo":1,"pageSize":10,"subjects":[578,571]} |
| `param.pageNo` | java.lang.Integer | yes | pageNo field | 1 |
| `param.pageSize` | java.lang.Integer | yes | pageSize field | 1 |
| `param.palletId` | java.lang.String | yes | palletId field | example |
| `param.subjects` | Integer[] | no | List of WB category subjectIDs, optional; when provided, only products in the corresponding categories are returned. | [578,571] |

**Response (top level)**: `result` ResultGTuskmls – Query result; `result.code` java.lang.String – code field; `result.message` java.lang.String – message field; `result.permissionName` java.lang.String – permissionName field; `result.result` WbPublishableOfferPageDTO – result field; `result.success` java.lang.Boolean – success field; …

<a id="publishcardget-1"></a>
## 142. Listing product card

`com.alibaba.fenxiao.crossborder:publish.card.get:1` · Listing (Publishing) · 铺货商品卡片  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/publish.card.get/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Get the card data for a single 1688 product, converted to WB format.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `param` | WbProductCardQueryParam | yes | Parameters for querying the product card | {"offerId":1234567890,"country":"ru"} |
| `param.country` | java.lang.String | yes | country field | country_example |
| `param.offerId` | java.lang.Long | yes | offerId field | 1 |

**Response (top level)**: `result` WbProductCardResponseDTO – Product card data return result; `result.offerId` java.lang.Long – 1688 product ID; `result.subjectID` java.lang.Integer – WB product category ID; `result.subject` String – WB category externalCategoryId, in the format parentID#subjectID; top-level or unknown categories degrade to a plain subjectID; `result.status` java.lang.String – Product card status; `result.productRating` java.lang.Float – Product rating (0-5); `result.images` java.util.List – List of product images; `result.videos` java.util.List – List of product videos; `result.cards` WbProductCardDTO[] – List of product cards (split by color variant); …

<a id="repurchasecontractget-1"></a>
## 143. Query repurchase contract

`com.alibaba.trade:repurchase.contract.get:1` · Repurchase · 复购合约查询  
POST `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/repurchase.contract.get/{appKey}`  
Token/system params: `_aop_timestamp` (optional), `_aop_signature` (required), `access_token` (required) · user authorization required · signature required  
Query whether a product supports a repurchase contract. A repurchase contract is a repeat-purchase discount agreement signed between buyer and seller, under which the seller gives repurchasing users special prices and guarantees. Orders under a repurchase contract must be placed through a specific trade flow.

**Request body**

| Field | Type | Req | Description | Example |
|---|---|---|---|---|
| `param` | TradeRepurchaseQueryParam | yes | Query parameter | {"offerIds": [ 867442475017]} |
| `param.offerIds` | java.lang.Long[] | yes | List of product IDs | [867442475017] |

**Response (top level)**: `result` ResultModel – Return value; `result.success` boolean – Whether successful; `result.code` java.lang.String – Error code; `result.message` java.lang.String – Error message; `result.result` TradeRepurchaseQueryResult – Repurchase contract model; …
