# XunYuanTong Sourcing Solution for Cross-Border Key Accounts

*English translation of the 1688 Open Platform solution page* [https://open.1688.com/solution/detail?key=1703167397970](https://open.1688.com/solution/detail?key=1703167397970)  
Original name: 跨境大客户寻源通解决方案 · Solution key 1703167397970 · Version 7.0 · Role: Cross-border key account (跨境大客户) · Identity: Procurement service provider (采购服务商) · Last updated 2026-08-25

## Overview

- Documentation: https://alidocs.dingtalk.com/i/nodes/ZX6GRezwJlzeYoPLFbr4Na6DWdqbropQ
- Permission requests: contact the business team and join the DingTalk group to obtain access to the documentation.

**Contents:** 143 APIs in 21 categories · 39 message topics · 18 change-log entries · SDK downloads

## API list

API identifiers, parameter names and error codes are kept exactly as in the original. Each API name links to its detail page (request/response parameters) on open.1688.com.

### Membership & Accounts (会员)

| API | Original name | Description |
|---|---|---|
| [Batch-add sub-account authorizations](https://open.1688.com/api/apidocdetail.htm?id=system.oauth2:subaccount.auth.add-1)<br>`system.oauth2:subaccount.auth.add:1` | 批量添加子账号授权 | Batch-add authorization for the sub-accounts under a main account. The main account must already be authorized. |
| [Batch-cancel sub-account authorizations](https://open.1688.com/api/apidocdetail.htm?id=system.oauth2:subaccount.auth.cancel-1)<br>`system.oauth2:subaccount.auth.cancel:1` | 批量取消子账号授权 | Cancel the authorization of sub-accounts in batch. |
| [Batch-query sub-account authorizations](https://open.1688.com/api/apidocdetail.htm?id=system.oauth2:subaccount.auth.list-1)<br>`system.oauth2:subaccount.auth.list:1` | 批量查询子账号授权 | Batch-query the authorization status of the sub-accounts under a main account. |
| [1688 member registration](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:account.user.register-1)<br>`com.alibaba.fenxiao.crossborder:account.user.register:1` | 1688会员注册 | Register a 1688 member. |
| [Get basic info of a non-authorized user (cross-border)](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.account:alibaba.account.agent.crossBasic-1)<br>`com.alibaba.account:alibaba.account.agent.crossBasic:1` | 跨境场景获取非授权用户的基本信息 | View another user's basic information. For use in cross-border scenarios. |
| [Get basic info of the authorized user](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.account:alibaba.account.basic-1)<br>`com.alibaba.account:alibaba.account.basic:1` | 获取授权用户的基本信息 | Get the basic information of the authorized user. |
| [Query sub-account info](https://open.1688.com/api/apidocdetail.htm?id=cn.alibaba.open:querySubAccount-1)<br>`cn.alibaba.open:querySubAccount:1` | 查询子账号信息 | Query sub-account information. |

### Tools (工具)

| API | Original name | Description |
|---|---|---|
| [Pull products from a product pool](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:pool.product.pull-1)<br>`com.alibaba.fenxiao.crossborder:pool.product.pull:1` | 拉取商品池中商品数据 | Batch-pull product data directly from a product pool by pool ID. |
| [Query product count in a product pool](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:pool.product.total-1)<br>`com.alibaba.fenxiao.crossborder:pool.product.total:1` | 查询商品池中商品总数 | Query the total number of products in a product pool. |
| [Encrypt a user loginId into an OpenUID](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.account:loginid.openuid.encrypt-1)<br>`com.alibaba.account:loginid.openuid.encrypt:1` | 用户loginId加密转换为Openuid接口 | Encrypts a user's loginId into an OpenUID. This interface is risk-controlled: batch operations are not allowed, and it may only be triggered manually by the merchant, for example when searching a user's orders or configuring rules. |
| [Get a link that opens a WangWang chat](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.account:account.wangwangUrl.get-1)<br>`com.alibaba.account:account.wangwangUrl.get:1` | 获取唤起旺旺聊天的链接 | Get a link that launches an AliWangWang chat. |
| [Decrypt an OpenUID into a WangWang nickname (only for launching WangWang)](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.account:wangwangnick.openuid.decrypt-1)<br>`com.alibaba.account:wangwangnick.openuid.decrypt:1` | Openuid转换解密为旺旺昵称接口（仅可使用于用户唤起旺旺） | Decrypts an OpenUID into a WangWang nickname. This interface is risk-controlled: it may only be used when a user needs to launch WangWang. Automated batch operations are not allowed, and it must not be used as a decryption interface to show plaintext to users. WangWang supports recall after encryption; do not use it for any other scenario. |
| [Add buyer-seller distribution relationship](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao:alibaba.fenxiao.relationadd-1)<br>`com.alibaba.fenxiao:alibaba.fenxiao.relationadd:1` | 买卖家分销关系添加 | Add a buyer-seller distribution relationship by product ID. |

### Market Insights (商机)

| API | Original name | Description |
|---|---|---|
| [Trending product search keywords](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:product.search.topKeyword-1)<br>`com.alibaba.fenxiao.crossborder:product.search.topKeyword:1` | 商品热搜词 | Trending product search keywords. |
| [Query ranking lists](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:product.topList.query-1)<br>`com.alibaba.fenxiao.crossborder:product.topList.query:1` | 查询榜单列表 | Query ranking (top) lists. |
| [Get daily sales-quantity trend for a product (new)](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:product.analyze.getPerdaySellQuantityTrendNew-1)<br>`com.alibaba.fenxiao.crossborder:product.analyze.getPerdaySellQuantityTrendNew:1` | 获取商品每日销售数量趋势（新） | Get a product's daily sales-quantity trend over 90 days (at most 90 days of data can be queried). New interface. |
| [Get 30-day median-price trend for a product (new)](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:product.analyze.getThirtyDayMedianPriceTrendNew-1)<br>`com.alibaba.fenxiao.crossborder:product.analyze.getThirtyDayMedianPriceTrendNew:1` | 获取商品30天价格中位数指标趋势（新） | Get a product's 30-day median-price metric trend. New interface. |
| [Get 90-day repurchase-rate trend for a product (new)](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:product.analyze.getRepurchaseRateTrendNew-1)<br>`com.alibaba.fenxiao.crossborder:product.analyze.getRepurchaseRateTrendNew:1` | 获取商品90天复购率指标趋势(新） | Get a product's 90-day repurchase-rate metric trend (at most 30 days can be queried). New interface. |

### Categories (类目)

| API | Original name | Description |
|---|---|---|
| [Query multilingual category by category ID](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:category.translation.getById-1)<br>`com.alibaba.fenxiao.crossborder:category.translation.getById:1` | 根据类目ID查询多语言类目 | Multilingual category query. Returns the category details in the requested language for the given category ID, including the list of its child categories. |
| [Query multilingual category by category name](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:category.translation.getByKeyword-1)<br>`com.alibaba.fenxiao.crossborder:category.translation.getByKeyword:1` | 根据类目名称查询多语言类目 | Multilingual category query. Returns the list of matching category details in the requested language for the given category name. Child-category data is not included. |
| [Get leaf-category attributes](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.product:alibaba.category.attribute.get-1)<br>`com.alibaba.product:alibaba.category.attribute.get:1` | 获取叶子类目属性 | Get category attributes by leaf-category ID. |

### Products (商品)

| API | Original name | Description |
|---|---|---|
| [Multilingual keyword search](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:product.search.keywordQuery-1)<br>`com.alibaba.fenxiao.crossborder:product.search.keywordQuery:1` | 多语言关键词搜索 | Multilingual keyword search. |
| [Multilingual image search](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:product.search.imageQuery-1)<br>`com.alibaba.fenxiao.crossborder:product.search.imageQuery:1` | 多语言图搜 | Multilingual image search. |
| [Multilingual in-store product search](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:product.search.querySellerOfferList-1)<br>`com.alibaba.fenxiao.crossborder:product.search.querySellerOfferList:1` | 多语言商品店搜 | Multilingual search of the products in a seller's store. |
| [Multilingual product detail](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:product.search.queryProductDetail-1)<br>`com.alibaba.fenxiao.crossborder:product.search.queryProductDetail:1` | 多语言商详 | Multilingual product detail. |
| [Product recommendations](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:product.search.offerRecommend-1)<br>`com.alibaba.fenxiao.crossborder:product.search.offerRecommend:1` | 商品推荐 | Personalized product recommendations. |
| [Multilingual search navigation](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:product.search.keywordSNQuery-1)<br>`com.alibaba.fenxiao.crossborder:product.search.keywordSNQuery:1` | 多语言搜索导航 | Get the multilingual keyword-search navigation list. |
| [Related product recommendations](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:product.related.recommend-1)<br>`com.alibaba.fenxiao.crossborder:product.related.recommend:1` | 相关性商品推荐 | Related product recommendations. |
| [Upload an image to get an imageId](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:product.image.upload-1)<br>`com.alibaba.fenxiao.crossborder:product.image.upload:1` | 上传图片获取imageId | Upload an image and receive an imageId. |
| [Claim the optimal coupon for a product](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.marketing:coupon.optimal.claim-1)<br>`com.alibaba.marketing:coupon.optimal.claim:1` | 通过商品领取最优化的优惠券 | Claim the best-value coupon for a product. Usually called before placing an order to complete the optimal coupon-claiming strategy. |
| [Recommend products by keyword (cross-border)](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.product:alibaba.product.suggest.crossBorder-1)<br>`com.alibaba.product:alibaba.product.suggest.crossBorder:1` | 跨境场景根据关键字推荐商品 | Recommend products by keyword and category in cross-border scenarios, sorted by sales volume. Note: this API is rate-limited and is only suitable for manual-association scenarios. |
| [Get simple product info from a previously purchased supplier](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.product:alibaba.product.simple.get-1)<br>`com.alibaba.product:alibaba.product.simple.get:1` | 获取已购买过商家的商品简单信息 | Get product details by product ID. This interface returns simple information for products of suppliers you have already purchased from. Access to this interface is paid. It returns only basic information and is mainly intended for data association in ERP systems. |
| [Follow a product (cross-border)](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:product.kjdistribute.addRelation-1)<br>`com.alibaba.fenxiao.crossborder:product.kjdistribute.addRelation:1` | 增加跨境关注商品 | Add a cross-border followed product. |
| [Unfollow a product (cross-border)](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:product.kjdistribute.removeRelation-1)<br>`com.alibaba.fenxiao.crossborder:product.kjdistribute.removeRelation:1` | 取消跨境关注商品 | Remove a cross-border followed product. |
| [Get product selling points for listing](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:product.distribute.getDistributeInfo-1)<br>`com.alibaba.fenxiao.crossborder:product.distribute.getDistributeInfo:1` | 获取商品铺货卖点 | For listing scenarios: get the selling-point information needed to list a product. |
| [Query leaf-category attribute and value mappings](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:product.category.getAttrById-1)<br>`com.alibaba.fenxiao.crossborder:product.category.getAttrById:1` | 叶子类目属性属性值映射查询 | Category mapping query. |

### Product Selection & Listing (选品铺货)

| API | Original name | Description |
|---|---|---|
| [Get product details (cross-border)](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.product:alibaba.cross.productInfo-1)<br>`com.alibaba.product:alibaba.cross.productInfo:1` | 跨境场景获取商品详情 | Get product details in cross-border scenarios. A cross-border listing relationship must be established before details can be retrieved. |
| [Add products to the listing list (cross-border)](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.product.push:alibaba.cross.syncProductListPushed-1)<br>`com.alibaba.product.push:alibaba.cross.syncProductListPushed:1` | 跨境场景下将商品加入铺货列表 | Cross-border only. Adds products to the listing list (that is, creates listing relationships); at most 20 items per call. Only after a product is added can its details be queried through the product-detail interface. Contact the cross-border operations staff to configure permissions manually before calling. |
| [Sync listing results](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.product.push:alibaba.product.push.syncPushProductResult-1)<br>`com.alibaba.product.push:alibaba.product.push.syncPushProductResult:1` | 同步铺货结果 | Sync listing results. When an ISV lists products from the source platform (1688) onto a target platform (for example TAOBAO), the ISV must return the listing result. The listing-status descriptions must match those defined by the source platform (1688). This interface also supports operations such as delisting, all of which are expressed through the listing status. |
| [Get product list (cross-border)](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.product:alibaba.cross.productList-1)<br>`com.alibaba.product:alibaba.cross.productList:1` | 跨境场景获取商品列表 | This interface validates the cross-border listing relationship and is for cross-border business only. Product model V2. |
| [Unfollow a product](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.product:alibaba.product.unfollow.crossborder-1)<br>`com.alibaba.product:alibaba.product.unfollow.crossborder:1` | 解除关注商品 | Unfollow a product. |
| [Query category by category ID](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.product:alibaba.category.get-1)<br>`com.alibaba.product:alibaba.category.get:1` | 根据类目Id查询类目 | Category query. To retrieve all 1688 categories, traverse the whole category tree starting from the root: first pass 0 to get all level-1 category IDs, then iterate the level-1 IDs to get all level-2 categories, and finally iterate the level-2 IDs to get the level-3 categories. Note: 1688 categories have only three levels; level-3 categories are the leaf categories required for publishing products. |
| [Follow a product](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.product:alibaba.product.follow.crossborder-1)<br>`com.alibaba.product:alibaba.product.follow.crossborder:1` | 关注商品 | Follow a product. |
| [[Relationship] Distributor: query supplier list](https://open.1688.com/api/apidocdetail.htm?id=cn.alibaba.open:alibaba.relation.querySuppliers-1)<br>`cn.alibaba.open:alibaba.relation.querySuppliers:1` | 【关系】分销商-查询供应商列表 | Get the supplier list for a distributor by userID. |

### Orders (订单)

| API | Original name | Description |
|---|---|---|
| [Create cross-border order](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.trade.createCrossOrder-1)<br>`com.alibaba.trade:alibaba.trade.createCrossOrder:1` | 跨境订单创建 | Cross-border-only order creation. An order may contain at most 50 SKUs, all from the same supplier. For multiple suppliers or more than 50 SKUs, split the order yourself before submitting. In some special cases several orders are created at once and several order numbers are returned. Supports both the open-marketplace and distribution scenarios. Orders are placed under the main account or a sub-account depending on the currently authorized user. |
| [View order details (buyer view)](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.trade.get.buyerView-1)<br>`com.alibaba.trade:alibaba.trade.get.buyerView:1` | 订单详情查看(买家视角) | Get the details of a single transaction; buyer calls only. Permission must be requested from the Alibaba Open Platform to use this API. |
| [Get order logistics info (buyer view)](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.logistics:alibaba.trade.getLogisticsInfos.buyerView-1)<br>`com.alibaba.logistics:alibaba.trade.getLogisticsInfos.buyerView:1` | 获取交易订单的物流信息(买家视角) | Requires the order buyer's authorization and returns the logistics details of the buyer's order. In procurement or distribution scenarios, buyers also need logistics details. This interface looks up logistics details by order number, including the sender, the recipient and the details of the goods shipped. Because of when logistics records are entered, API tracking queries may be delayed. Permission must be requested from the Open Platform to access this API. |
| [Get order logistics tracking info (buyer view)](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.logistics:alibaba.trade.getLogisticsTraceInfo.buyerView-1)<br>`com.alibaba.logistics:alibaba.trade.getLogisticsTraceInfo.buyerView:1` | 获取交易订单的物流跟踪信息(买家视角) | Requires the order buyer's authorization and returns the logistics tracking information of the buyer's order. In procurement or distribution scenarios, buyers also need logistics details. This interface looks up tracking information by logistics (waybill) number. Because of when logistics records are entered, API tracking queries may be delayed. Permission must be requested from the Open Platform to access this API. |
| [Cancel transaction](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.trade.cancel-1)<br>`com.alibaba.trade:alibaba.trade.cancel:1` | 取消交易 | Buyer or seller cancels a transaction. Only transactions in specific statuses can be cancelled; on 1688 this is used to cancel unpaid orders. If an order is closed less than 10 seconds after it was created, the error CLOSE_ORDER_TOO_FAST is returned. |
| [View order list (buyer view)](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.trade.getBuyerOrderList-1)<br>`com.alibaba.trade:alibaba.trade.getBuyerOrderList:1` | 订单列表查看(买家视角) | Get the buyer's order list; the user's memberId must equal the buyer memberId on the orders. This interface returns only basic order information, not logistics or invoice information. For logistics information call the order-detail interface; for invoice information call the invoice-information API. |
| [Preview data before creating an order](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.createOrder.preview-1)<br>`com.alibaba.trade:alibaba.createOrder.preview:1` | 创建订单前预览数据接口 | Orders may only contain products from a single supplier. This interface returns discount and related information for order creation. It 1. validates whether the products may be ordered; 2. validates the consignment (distribution) relationship; 3. validates stock, minimum order quantity and whether mixed-batch conditions are met. |
| [Update order memo](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.order.memoAdd-1)<br>`com.alibaba.trade:alibaba.order.memoAdd:1` | 修改订单备忘 | If the authorized user is the seller, updates the seller memo; if the buyer, updates the buyer memo. Note: this interface can be called repeatedly, and the memo overwrites the content of the previous call. |
| [Get sub-account list](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.account:alibaba.subAccount.list-1)<br>`com.alibaba.account:alibaba.subAccount.list:1` | 获取子账号列表 | Get the user's main-account and sub-account information. If the API is authorized as a sub-account, only the main account that the sub-account belongs to is returned. If authorized as a main account, the list of all sub-accounts is returned. |
| [Buyer adds an order message](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.trade.addFeedback-1)<br>`com.alibaba.trade:alibaba.trade.addFeedback:1` | 买家补充订单留言接口 | Buyer adds a supplementary message to an order. The total message length must not exceed 500 characters. |
| [Buyer confirms receipt](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:trade.receivegoods.confirm-1)<br>`com.alibaba.trade:trade.receivegoods.confirm:1` | 买家确认收货 | Buyer confirms receipt of goods. |
| [Buyer deletes a closed order](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:trade.order.buyerdelete-1)<br>`com.alibaba.trade:trade.order.buyerdelete:1` | 买家删除已关闭的订单 | Buyer deletes an order that has been closed. |
| [Buyer requests a shipping-address change](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:order.receiveAddress.buyerUpdate-1)<br>`com.alibaba.trade:order.receiveAddress.buyerUpdate:1` | 买家申请修改收货地址 | Buyer changes the shipping address. If the new address is in a remote area or the shipping fee must be recalculated, the seller must confirm the change and may reject it. |

### Payment (支付)

| API | Original name | Description |
|---|---|---|
| [Batch-get payment links for orders](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.alipay.url.get-1)<br>`com.alibaba.trade:alibaba.alipay.url.get:1` | 批量获取订单的支付链接 | When paying through an ERP, use this API to get a cashier link for batch payment. A single order returns the 1688 cashier URL; multiple orders return the Alipay cashier URL. The ERP can redirect the user to the cashier link to complete payment. The user's 1688 login status is verified before payment. |
| [Get payment link for Cross-Border Pay (Kuajingbao)](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.crossBorderPay.url.get-1)<br>`com.alibaba.trade:alibaba.crossBorderPay.url.get:1` | 获取使用跨境宝支付的支付链接 | Get a payment link for paying with Cross-Border Pay (Kuajingbao). |
| [Get payment link for Cheng-e-She credit pay](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.creditPay.url.get-1)<br>`com.alibaba.trade:alibaba.creditPay.url.get:1` | 获取使用诚e赊支付的支付链接 | Get a payment link for paying with Cheng-e-She (buy-now-pay-later credit). |
| [Buyer views all granted credit terms](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.accountPeriod.list.buyerView-1)<br>`com.alibaba.trade:alibaba.accountPeriod.list.buyerView:1` | 买家查看获得的所有账期授信 | View, from the buyer's side, all account-period (credit-term) lines the buyer has been granted. Paginated; at most 10 records per call. |
| [Query payment channels supported by an order](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.trade.payWay.query-1)<br>`com.alibaba.trade:alibaba.trade.payWay.query:1` | 查询订单可以支持的支付渠道 | Query the payment methods or channels available for an unpaid order. |
| [Check whether password-free payment is enabled](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.trade.pay.protocolPay.isopen-1)<br>`com.alibaba.trade:alibaba.trade.pay.protocolPay.isopen:1` | 查询是否开通免密支付 | Check whether an auto-debit (withholding) agreement is enabled. |
| [Initiate password-free payment](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.trade.pay.protocolPay.preparePay-1)<br>`com.alibaba.trade:alibaba.trade.pay.protocolPay.preparePay:1` | 发起免密支付 | Initiates a password-free payment. Automatically detects whether Alipay or Cheng-e-She password-free payment is enabled and initiates the debit. Cheng-e-She auto-debit is tried first; if it fails, Alipay auto-debit is attempted. The error codes returned by this interface are currently not detailed; after a failed debit, retry up to 3 times rather than indefinitely. |
| [Order payment consultation](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:trade.orderpay.analysis-1)<br>`com.alibaba.trade:trade.orderpay.analysis:1` | 交易订单支付咨询 | Order-payment consultation interface, used to analyse which payment method an order uses, and so on. |
| [Get combined-cashier URL](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.trade.grouppay.url.get-1)<br>`com.alibaba.trade:alibaba.trade.grouppay.url.get:1` | 组合收银台url获取 | Get the combined-cashier URL. |

### Logistics (物流)

| API | Original name | Description |
|---|---|---|
| [Estimate domestic (China) shipping fee for a product](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:product.freight.estimate-1)<br>`com.alibaba.fenxiao.crossborder:product.freight.estimate:1` | 商品中国国内运费预估 | Estimate a product's shipping fee from the product ID and the province/city/district codes of a delivery address within mainland China. |
| [Query external order ID by waybill number or unclaimed-parcel code](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:logistics.order.getOutOrderId-1)<br>`com.alibaba.fenxiao.crossborder:logistics.order.getOutOrderId:1` | 根据运单号或无主件码查询外部订单ID | Query the external order ID by waybill number or unclaimed-parcel code. |
| [Query shipping-insurance info](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:shipping.insurance.get-1)<br>`com.alibaba.trade:shipping.insurance.get:1` | 运费险信息查询 | Query shipping-insurance information. |
| [Get shipping-template details](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.logistics:alibaba.logistics.myFreightTemplate.list.get-1)<br>`com.alibaba.logistics:alibaba.logistics.myFreightTemplate.list.get:1` | 获取物流模板详情 | Get the seller's shipping template by template ID. Template ID 0 means "shipping fee to be explained"; 1 means the seller bears the shipping fee. |
| [Buyer gets saved shipping addresses](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.trade.receiveAddress.get-1)<br>`com.alibaba.trade:alibaba.trade.receiveAddress.get:1` | 买家获取保存的收货地址信息列表 | Get the buyer's list of saved shipping addresses. |
| [Parse an address into area codes](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.trade.addresscode.parse-1)<br>`com.alibaba.trade:alibaba.trade.addresscode.parse:1` | 根据地址解析地区码 | Parse area codes from address information. |
| [Query seller mixed-batch settings](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.trade.OpQueryMarketingMixConfig-1)<br>`com.alibaba.trade:alibaba.trade.OpQueryMarketingMixConfig:1` | 查询卖家混批设置 | Query the seller's mixed-batch (mixed wholesale) settings. |
| [Logistics company list (all companies)](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.logistics:alibaba.logistics.OpQueryLogisticCompanyList-1)<br>`com.alibaba.logistics:alibaba.logistics.OpQueryLogisticCompanyList:1` | 物流公司列表-所有的物流公司 | Get the names of all logistics companies. |
| [Parse overseas address](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:trade.address.parseText-1)<br>`com.alibaba.trade:trade.address.parseText:1` | 海外地址解析 | Parse an overseas address. |
| [Urge seller to ship](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.logistics:logistics.delivery.urge-1)<br>`com.alibaba.logistics:logistics.delivery.urge:1` | 催卖家发货 | Urge the seller to ship. The order must still be in a not-yet-shipped status. Limited to once per 24 hours. |
| [Receive external print-label (UDF) URL](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:label.url.receive-1)<br>`com.alibaba.fenxiao.crossborder:label.url.receive:1` | 接收外部打印UDF链接 | Receive an external print-label (UDF) URL. |
| [Get waybill-number set](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:bigcustomer.mailNo.query-1)<br>`com.alibaba.fenxiao.crossborder:bigcustomer.mailNo.query:1` | 获取运单号集合 | Get the set of waybill numbers. |

### Official Return Pickup (官方退上门取件)

| API | Original name | Description |
|---|---|---|
| [Create official return-pickup logistics order](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.logistics:refundofficialdelivery.order.create-1)<br>`com.alibaba.logistics:refundofficialdelivery.order.create:1` | 创建退货官方物流上门揽订单 | Create an official-logistics door-to-door pickup order for a return. |
| [Modify official return-pickup logistics order](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.logistics:refundofficialdelivery.order.modify-1)<br>`com.alibaba.logistics:refundofficialdelivery.order.modify:1` | 退货官方物流上门揽订单修改 | Modify an official-logistics door-to-door pickup order for a return. |
| [Get official return-pickup logistics order details](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.logistics:refundofficialdelivery.order.get-1)<br>`com.alibaba.logistics:refundofficialdelivery.order.get:1` | 退货官方物流上门揽订单详情获取 | Get the details of an official-logistics door-to-door pickup order for a return. |
| [Cancel official return-pickup logistics order](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.logistics:refundofficialdelivery.order.cancel-1)<br>`com.alibaba.logistics:refundofficialdelivery.order.cancel:1` | 取消退货官方物流上门揽订单 | Cancel an official-logistics door-to-door pickup order for a return. |
| [Query official return-pickup plans](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.logistics:refundofficialdelivery.solution.get-1)<br>`com.alibaba.logistics:refundofficialdelivery.solution.get:1` | 官方退货上门揽方案查询 | Query the available official door-to-door return-pickup plans. |

### Data Write-back (回传数据)

| API | Original name | Description |
|---|---|---|
| [Write back mapping between end-customer orders and 1688 orders](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:order.relation.write-1)<br>`com.alibaba.fenxiao.crossborder:order.relation.write:1` | 回传机构真实用户订单和1688订单的映射关系 | Write back to 1688 the mapping between the organisation's real end-customer orders and the 1688 orders generated under the organisation's account. |
| [Write back country-site logistics orders](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:trade.cross.logisticsOrderSync-1)<br>`com.alibaba.fenxiao.crossborder:trade.cross.logisticsOrderSync:1` | 国家站物流单回传 | Write back logistics orders from the country site. |
| [Sync downstream sales orders](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:trade.cross.orderSync-1)<br>`com.alibaba.fenxiao.crossborder:trade.cross.orderSync:1` | 下游销售订单同步 | Sync downstream sales orders. |
| [Save the business line an account belongs to](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:account.business.save-1)<br>`com.alibaba.fenxiao.crossborder:account.business.save:1` | 保存账号所属业务线 | Save the business line that an account belongs to. |

### Returns & Refunds (退货退款)

| API | Original name | Description |
|---|---|---|
| [Query refund details by order ID (buyer view)](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.trade.refund.OpQueryBatchRefundByOrderIdAndStatus-1)<br>`com.alibaba.trade:alibaba.trade.refund.OpQueryBatchRefundByOrderIdAndStatus:1` | 查询退款单详情-根据订单ID（买家视角） | For buyers; sellers should use alibaba.trade.refund.OpQueryBatchRefundByOrderIdAndStatus.sellerView. Queries the refund list for an order in real time. Currently only in-sale refunds (before the transaction completes) can be queried. |
| [Refund operation history (buyer view)](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.trade.refund.OpQueryOrderRefundOperationList-1)<br>`com.alibaba.trade:alibaba.trade.refund.OpQueryOrderRefundOperationList:1` | 退款单操作记录列表（买家视角） | For buyers; sellers should use alibaba.trade.refund.OpQueryOrderRefund.sellerView. Returns the buyer-side refund operation records. |
| [Query refund details by refund ID (buyer view)](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.trade.refund.OpQueryOrderRefund-1)<br>`com.alibaba.trade:alibaba.trade.refund.OpQueryOrderRefund:1` | 查询退款单详情-根据退款单ID（买家视角） | For buyers; sellers should use alibaba.trade.refund.OpQueryOrderRefund.sellerView. Queries refund details, including the list of refund operations. Permission must be requested from Alibaba to access this API. |
| [Query refund list (buyer view)](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.trade.refund.buyer.queryOrderRefundList-1)<br>`com.alibaba.trade:alibaba.trade.refund.buyer.queryOrderRefundList:1` | 查询退款单列表(买家视角) | Buyer views the refund list. This interface does not support sub-account queries; authorize with the main account before querying. |
| [Query order number by claim or insurance-policy number](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.trade.queryOrderByInsure-1)<br>`com.alibaba.trade:alibaba.trade.queryOrderByInsure:1` | 根据理赔单或保险单号查询对应的订单号 | Query the order number that corresponds to a claim or insurance-policy number. Pass type=lp for claims and type=bx for shipping insurance. |
| [Create refund/return request](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.trade.createRefund-1)<br>`com.alibaba.trade:alibaba.trade.createRefund:1` | 创建退款退货申请 | Create a refund or return request. |
| [Query refund/return reasons (for creating a request)](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.trade.getRefundReasonList-1)<br>`com.alibaba.trade:alibaba.trade.getRefundReasonList:1` | 查询退款退货原因（用于创建退款退货） | Query the refund/return reasons (used when creating a refund or return request). |
| [Upload refund/return evidence](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.trade.uploadRefundVoucher-1)<br>`com.alibaba.trade:alibaba.trade.uploadRefundVoucher:1` | 上传退款退货凭证 | Upload evidence for a refund/return request. To convert a file stream to a byte array, org.apache.commons.io.IOUtils#toByteArray(java.io.InputStream) is recommended. |
| [Buyer submits return-shipment info](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.trade.refund.returnGoods-1)<br>`com.alibaba.trade:alibaba.trade.refund.returnGoods:1` | 买家提交退款货信息 | Used after the seller approves the buyer's return/refund request, for the buyer to submit the return-shipment information. First call alibaba.logistics.OpQueryLogisticCompanyList.offline to look up logistics companies, and use the logistics-company code it returns. |
| [Cancel refund/return request](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.trade.cancelRefund-1)<br>`com.alibaba.trade:alibaba.trade.cancelRefund:1` | 取消退款退货申请 | Cancel a refund or return request. |
| [Query maximum refundable amount when applying](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.trade.getMaxRefundFee-1)<br>`com.alibaba.trade:alibaba.trade.getMaxRefundFee:1` | 申请退款时查询最大可退费用 | Query the maximum refundable amount when applying for a refund. |
| [Apply for trade arbitration](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:trade.arbitration.apply-1)<br>`com.alibaba.trade:trade.arbitration.apply:1` | 申请交易仲裁 | Trade arbitration application. |
| [Query the merchant's return address](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao:refund.address.get-1)<br>`com.alibaba.fenxiao:refund.address.get:1` | 查询商家退货地址 | Query the return address. |

### Messaging (消息)

| API | Original name | Description |
|---|---|---|
| [Batch-confirm failed messages](https://open.1688.com/api/apidocdetail.htm?id=cn.alibaba.open:push.message.confirm-1)<br>`cn.alibaba.open:push.message.confirm:1` | 失败消息批量确认 | Manually call the confirmation API to confirm that messages have been consumed successfully. Only needed when using the query-style API to fetch failed messages. |
| [Fetch failed messages (query style)](https://open.1688.com/api/apidocdetail.htm?id=cn.alibaba.open:push.query.messageList-1)<br>`cn.alibaba.open:push.query.messageList:1` | 查询式获取失败的消息列表 | Query-style retrieval of sent messages. Retrieved messages are not confirmed automatically; the caller must call the confirmation API to confirm the consumption status. Note that confirming affects the data returned by pagination. |
| [Fetch failed messages (cursor style)](https://open.1688.com/api/apidocdetail.htm?id=cn.alibaba.open:push.cursor.messageList-1)<br>`cn.alibaba.open:push.cursor.messageList:1` | 游标式获取失败的消息列表 | Cursor-style retrieval of failed messages. Retrieved messages are automatically confirmed as consumed, so the next call with the same conditions returns the remaining data, until the result is empty. |
| [View order list (seller view)](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.trade.getSellerOrderList-1)<br>`com.alibaba.trade:alibaba.trade.getSellerOrderList:1` | 订单列表查看(卖家视角) | Get the seller's order list; the user's memberId must equal the sellerMemberId on the orders. This interface returns only basic order information, not logistics or invoice information. For logistics information call the order-detail interface; for invoice information call the invoice-information API. |
| [Ship: seller arranges own logistics](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.logistics:alibaba.logistics.OpDeliverySendOrder.offline-1)<br>`com.alibaba.logistics:alibaba.logistics.OpDeliverySendOrder.offline:1` | 物流发货-自己联系物流发货 | For 1688 open-marketplace orders where the seller arranges logistics themselves. Supports combined shipping (several orders in one shipment) and shipping at sub-order (orderEntry) level. Shipping by quantity is not supported. |
| [Ship: no logistics needed](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.logistics:alibaba.logistics.OpDeliverySendOrder.dummy-1)<br>`com.alibaba.logistics:alibaba.logistics.OpDeliverySendOrder.dummy:1` | 物流发货-无需物流 | For 1688 open-marketplace orders that need no logistics. Supports combined shipping (several orders in one shipment) and shipping at sub-order (orderEntry) level. Shipping by quantity is not supported. |

### Light Customization (轻定制)

| API | Original name | Description |
|---|---|---|
| [Get artwork info for a customization order](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:product.podOrderDesign.get-1)<br>`com.alibaba.fenxiao.crossborder:product.podOrderDesign.get:1` | 获取加工定制订单稿件信息 | Get the design/artwork information of a processing-and-customization (print-on-demand) order. |

### Suppliers (商家)

| API | Original name | Description |
|---|---|---|
| [Run AI supplier-search task (async)](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:account.search.getInfo.async-1)<br>`com.alibaba.fenxiao.crossborder:account.search.getInfo.async:1` | 异步执行AI找商任务 | Execute an AI supplier-search task asynchronously. |
| [Query AI supplier-search task result](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:account.search.getInfoResult-1)<br>`com.alibaba.fenxiao.crossborder:account.search.getInfoResult:1` | 查询AI找商任务执行结果 | Query the execution result of an AI supplier-search task. |

### Inquiries (Newton Cloud) (询盘)

| API | Original name | Description |
|---|---|---|
| [Newton Cloud: create long-running task](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.agent:newtoncloud.task.create-1)<br>`com.alibaba.agent:newtoncloud.task.create:1` | 牛顿云-创建长程任务 | Create a Newton Cloud long-running task. It executes asynchronously and returns taskId/sessionId/status. |
| [Newton Cloud: query task](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.agent:newtoncloud.task.get-1)<br>`com.alibaba.agent:newtoncloud.task.get:1` | 牛顿云-查询任务服务 | Newton Cloud task query service. |
| [Newton Cloud: list tasks](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.agent:newtoncloud.task.list-1)<br>`com.alibaba.agent:newtoncloud.task.list:1` | 牛顿云-查询任务列表 | Query the current user's task list. Returns each task's taskId/sessionId/status/taskType and its creation and completion times. |
| [Newton Cloud: terminate task](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.agent:newtoncloud.task.kill-1)<br>`com.alibaba.agent:newtoncloud.task.kill:1` | 牛顿云-终止任务 | Terminate the task with the given taskId, mark it with the KILL status, and return killed. |
| [Newton Cloud: query task table](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.agent:newtoncloud.task.fetch-1)<br>`com.alibaba.agent:newtoncloud.task.fetch:1` | 牛顿云-查询任务表格服务 | Newton Cloud task-table query service. |
| [Newton Cloud: resume task](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.agent:newtoncloud.task.resume-1)<br>`com.alibaba.agent:newtoncloud.task.resume:1` | 牛顿云-恢复任务 | Resume a Newton Cloud task. |
| [Newton Cloud: query batch-inquiry results](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.agent:newtoncloud.batchInquiry.getResult-1)<br>`com.alibaba.agent:newtoncloud.batchInquiry.getResult:1` | 牛顿云-查询批量询盘结果 | Query inquiry results by batch-inquiry task ID (wwTaskId). After the ISV receives an AGENT_NEWTON_CLOUD_TASK_NOTIFY notification (stage=BATCH_INQUIRY, stageDetail.phase=COMPLETE), call this interface with the stageDetail.wwTaskId from the notification. It returns structured inquiry results synchronously, including supplier reply summaries, quotation details and recommendation reasons. |
| [Newton Cloud: upload file](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.agent:newtoncloud.file.upload-1)<br>`com.alibaba.agent:newtoncloud.file.upload:1` | 牛顿云-上传文件 | Upload a local file to Newton Cloud OSS. Returns a temporary public download link and a stable download address, which can be used directly as the fileUrls parameter of newtoncloud.task.create. File content is passed as Base64; a single file of 3 MB or less is recommended. |
| [Newton Cloud: list available model tiers](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.agent:newtoncloud.model.list-1)<br>`com.alibaba.agent:newtoncloud.model.list:1` | 牛顿云-查询可用模型档位列表 | Query the model tiers currently available on Newton Cloud. Returns each tier's code (id), display name (displayName) and other information. The tier code can be used as the model parameter of newtoncloud.task.create. No business parameters; only access_token and signature are required. |
| [Newton Cloud: query points details](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.agent:newtoncloud.points.query-1)<br>`com.alibaba.agent:newtoncloud.points.query:1` | 牛顿云-查询积分详情 | Query the Newton Cloud points information of the currently authorized account, including total points, used points, available points, points sources, expiry information and paginated usage details. The user is identified automatically from the access_token; no userId is needed. |

### 88 ShengYiTong (Business Link) (88生意通)

| API | Original name | Description |
|---|---|---|
| [88 ShengYiTong: buyer drafts purchase order](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.syt:syt.buyer.draftPurchaseOrder-2)<br>`com.alibaba.syt:syt.buyer.draftPurchaseOrder:2` | 88生意通买家起草采购单 | 88 ShengYiTong solution. The buyer drafts a purchase order. |
| [88 ShengYiTong: confirm transaction complete](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.syt:syt.contract.confirm-1)<br>`com.alibaba.syt:syt.contract.confirm:1` | 88生意通确认交易完成 | 88 ShengYiTong: confirm that the transaction is complete. |
| [88 ShengYiTong: void contract](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.syt:syt.contract.invalid-1)<br>`com.alibaba.syt:syt.contract.invalid:1` | 88生意通合同作废 | 88 ShengYiTong: void (invalidate) a contract. |
| [88 ShengYiTong: contract refund request](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.syt:syt.contract.refund-1)<br>`com.alibaba.syt:syt.contract.refund:1` | 88生意通合同退款申请 | 88 ShengYiTong: apply for a contract refund. |
| [88 ShengYiTong: buyer payment, get cashier URL](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.syt:syt.contract.pay-1)<br>`com.alibaba.syt:syt.contract.pay:1` | 88生意通 买家支付获取收银台 URL | 88 ShengYiTong solution. |
| [88 ShengYiTong: WorldFirst transfer payment (WorldFirst B2C accounts only)](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.syt:syt.contract.payTransfer-1)<br>`com.alibaba.syt:syt.contract.payTransfer:1` | 88生意通万里汇转账支付-仅支持万里汇B2C账号 | 88 ShengYiTong payment-transfer interface. |
| [88 ShengYiTong: buyer confirms purchase order](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.syt:syt.buyer.confirmPurchaseOrder-1)<br>`com.alibaba.syt:syt.buyer.confirmPurchaseOrder:1` | 88生意通买家确认采购单 | 88 ShengYiTong: the buyer confirms a purchase order. |
| [88 ShengYiTong: buyer queries purchase-order details](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.syt:syt.buyer.queryContractDetail-1)<br>`com.alibaba.syt:syt.buyer.queryContractDetail:1` | 88生意通买家查询采购单详情 | 88 ShengYiTong: the buyer queries the details of a purchase order. |
| [88 ShengYiTong: check whether customer is verified](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.syt:syt.customer.queryUserAuthStatus-1)<br>`com.alibaba.syt:syt.customer.queryUserAuthStatus:1` | 88生意通查询客户是否已认证 | 88 ShengYiTong: check whether a customer has completed verification. |

### Invoicing (发票)

| API | Original name | Description |
|---|---|---|
| [Query buyer invoice titles (paginated)](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:trade.invoiceTitle.getPageList-1)<br>`com.alibaba.trade:trade.invoiceTitle.getPageList:1` | 分页查询买家抬头列表 | Paginated query of the buyer's invoice titles. |
| [Query invoiceable amount of orders](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:trade.invoiceAmount.getList-1)<br>`com.alibaba.trade:trade.invoiceAmount.getList:1` | 查询订单可开票金额 | Query the invoiceable amount of orders. To find invoiceable orders first: 1. alibaba.trade.getBuyerOrderList-1: pass needInvoicingSetting=true and check invoicingSettingModel.tradeInvoiceStatus in the response to see whether an order can be invoiced. 2. alibaba.trade.get.buyerView-1: include InvoicingSetting in the includeFields parameter. |
| [Query invoice applications (buyer view, paginated)](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:trade.invoiceApply.getPageListBuyerView-1)<br>`com.alibaba.trade:trade.invoiceApply.getPageListBuyerView:1` | 分页查询发票申请列表（买家视角） | Paginated query of invoice applications (buyer view). |
| [Query all issued invoices for an order (buyer view)](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:trade.invoice.getListBuyerView-1)<br>`com.alibaba.trade:trade.invoice.getListBuyerView:1` | 查询交易单下关联的所有开具的发票信息（买家视角） | Query all issued invoices associated with an order (buyer view). |
| [Buyer requests invoice](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:trade.invoice.apply-1)<br>`com.alibaba.trade:trade.invoice.apply:1` | 买家申请开票 | Buyer requests an invoice; batch requests are supported. Prerequisites: get the buyer's invoice titles; get the invoiceable orders and the invoice types they support; get the orders' invoiceable amounts. |
| [Merged-invoice consultation](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:trade.invoice.consult-1)<br>`com.alibaba.trade:trade.invoice.consult:1` | 合并开票咨询 | Submit orders to check whether they can be invoiced together and how they are grouped. |
| [Submit merged-invoice request](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:trade.invoice.mergeapply-1)<br>`com.alibaba.trade:trade.invoice.mergeapply:1` | 提交合并开票 | First call the merged-invoice consultation trade.invoice.consult to get the grouping result, then call merged invoicing. |
| [Query merged-invoice relationships by order or invoice application](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:trade.invoice.sellerqueryrelatedorders-1)<br>`com.alibaba.trade:trade.invoice.sellerqueryrelatedorders:1` | 基于交易单或发票申请单查询合单关联关系 | The seller queries merged-invoice relationships by order ID (orderId) or invoice-application ID (outbizid). If both are passed, outbizid takes precedence. If the result is not a merged invoice, mergeInvoice returns false. |

### Fully Managed (Consignment) (全托管)

| API | Original name | Description |
|---|---|---|
| [Warehouse receipt / shelving of goods](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:consignment.co.status.sync-1)<br>`com.alibaba.fenxiao.crossborder:consignment.co.status.sync:1` | 仓库签收/上架商品 | Open to Buffalo. Syncs the Buffalo warehouse's receipt or shelving of goods to the 1688 shipment-order status. |
| [Warehouse creates discrepancy tally sheet](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:consignment.tally.create-1)<br>`com.alibaba.fenxiao.crossborder:consignment.tally.create:1` | 仓库创建差异理货单 | The Buffalo warehouse creates a discrepancy tally sheet. |
| [Update warehouse product inventory](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:consignment.co.inventory-1)<br>`com.alibaba.fenxiao.crossborder:consignment.co.inventory:1` | 仓库商品库存更新 | Update the inventory of products in the warehouse. |

### Listing (Publishing) (铺货)

| API | Original name | Description |
|---|---|---|
| [Publish result callback](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:publish.result.callback-1)<br>`com.alibaba.fenxiao.crossborder:publish.result.callback:1` | 发布结果回调 | WB-1688 integration Method 3: publish-result callback interface. |
| [Listable product list](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:publish.product.list-1)<br>`com.alibaba.fenxiao.crossborder:publish.product.list:1` | 铺货商品列表 | Query the list of products available for listing. Supports paginated queries of product-ID sets by assortment (goods-pool) ID. |
| [Listing product card](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:publish.card.get-1)<br>`com.alibaba.fenxiao.crossborder:publish.card.get:1` | 铺货商品卡片 | Get the card data for a single 1688 product, converted to WB format. |

### Repurchase (复购)

| API | Original name | Description |
|---|---|---|
| [Query repurchase contract](https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:repurchase.contract.get-1)<br>`com.alibaba.trade:repurchase.contract.get:1` | 复购合约查询 | Query whether a product supports a repurchase contract. A repurchase contract is a repeat-purchase discount agreement signed between buyer and seller, under which the seller gives repurchasing users special prices and guarantees. Orders under a repurchase contract must be placed through a specific trade flow. |

## Message topics

### Trade Messages (交易消息)

| Topic | Description | Original |
|---|---|---|
| [`ORDER_BUYER_VIEW_BUYER_MAKE`](https://open.1688.com/doc/topicDetail.htm?id=ORDER_BUYER_VIEW_BUYER_MAKE) | Order created (buyer view) | 1688创建订单（买家视角）/order created (buyer view) |
| [`ORDER_BUYER_VIEW_ORDER_PRICE_MODIFY`](https://open.1688.com/doc/topicDetail.htm?id=ORDER_BUYER_VIEW_ORDER_PRICE_MODIFY) | Order price modified (buyer view) | 1688修改订单价格（买家视角）/order price modification (buyer view) |
| [`ORDER_BUYER_VIEW_ORDER_PAY`](https://open.1688.com/doc/topicDetail.htm?id=ORDER_BUYER_VIEW_ORDER_PAY) | Order paid (buyer view) | 1688交易付款（买家视角）/1688 transaction payment (buyer view) |
| [`ORDER_BUYER_VIEW_ANNOUNCE_SENDGOODS`](https://open.1688.com/doc/topicDetail.htm?id=ORDER_BUYER_VIEW_ANNOUNCE_SENDGOODS) | Order shipped (buyer view) | 1688订单发货（买家视角）/1688 order delivery (buyer view) |
| [`ORDER_BUYER_VIEW_PART_PART_SENDGOODS`](https://open.1688.com/doc/topicDetail.htm?id=ORDER_BUYER_VIEW_PART_PART_SENDGOODS) | Order partially shipped (buyer view) | 1688订单部分发货（买家视角）/Partial delivery of 1688 order (buyer view) |
| [`ORDER_BUYER_VIEW_ORDER_COMFIRM_RECEIVEGOODS`](https://open.1688.com/doc/topicDetail.htm?id=ORDER_BUYER_VIEW_ORDER_COMFIRM_RECEIVEGOODS) | Receipt of goods confirmed (buyer view) | 1688订单确认收货（买家视角）/order receipt confirmation (buyer view) |
| [`ORDER_BUYER_VIEW_ORDER_SUCCESS`](https://open.1688.com/doc/topicDetail.htm?id=ORDER_BUYER_VIEW_ORDER_SUCCESS) | Transaction completed successfully (buyer view) | 1688交易成功（买家视角） |
| [`ORDER_BUYER_VIEW_ORDER_BUYER_CLOSE`](https://open.1688.com/doc/topicDetail.htm?id=ORDER_BUYER_VIEW_ORDER_BUYER_CLOSE) | Order closed by buyer (buyer view) | 1688买家关闭订单（买家视角）/buyer closing order (buyer view) |
| [`ORDER_BUYER_VIEW_ORDER_BOPS_CLOSE`](https://open.1688.com/doc/topicDetail.htm?id=ORDER_BUYER_VIEW_ORDER_BOPS_CLOSE) | Order closed by operations back-office (buyer view) | 1688运营后台关闭订单（买家视角） |
| [`ORDER_BUYER_VIEW_ORDER_BUYER_REFUND_IN_SALES`](https://open.1688.com/doc/topicDetail.htm?id=ORDER_BUYER_VIEW_ORDER_BUYER_REFUND_IN_SALES) | In-sale refund on order (buyer view) | 1688订单售中退款（买家视角） |
| [`ORDER_BUYER_VIEW_ORDER_REFUND_AFTER_SALES`](https://open.1688.com/doc/topicDetail.htm?id=ORDER_BUYER_VIEW_ORDER_REFUND_AFTER_SALES) | After-sales refund on order (buyer view) | 1688订单售后退款（买家视角） |
| [`ORDER_BUYER_VIEW_ORDER_STEP_PAY`](https://open.1688.com/doc/topicDetail.htm?id=ORDER_BUYER_VIEW_ORDER_STEP_PAY) | Order stage payment (buyer view) | 1688订单阶段付款（买家视角） |
| [`ORDER_BATCH_PAY`](https://open.1688.com/doc/topicDetail.htm?id=ORDER_BATCH_PAY) | Order batch-payment status sync | 1688订单批量支付状态同步消息 |
| [`ORDER_BUYER_VIEW_ORDER_SELLER_CLOSE`](https://open.1688.com/doc/topicDetail.htm?id=ORDER_BUYER_VIEW_ORDER_SELLER_CLOSE) | Order closed by seller (buyer view) | 1688卖家关闭订单（买家视角）/seller closing order (buyer view) |
| [`ORDER_ORDER_PRICE_MODIFY`](https://open.1688.com/doc/topicDetail.htm?id=ORDER_ORDER_PRICE_MODIFY) | Order price modified (seller view) | 1688修改订单价格（卖家视角） |

### Logistics Messages (物流消息)

| Topic | Description | Original |
|---|---|---|
| [`LOGISTICS_BUYER_VIEW_TRACE`](https://open.1688.com/doc/topicDetail.htm?id=LOGISTICS_BUYER_VIEW_TRACE) | Logistics order status changed (buyer view) | 物流单状态变更（买家视角） |
| [`LOGISTICS_MAIL_NO_CHANGE`](https://open.1688.com/doc/topicDetail.htm?id=LOGISTICS_MAIL_NO_CHANGE) | Waybill number changed | 物流单号修改消息 |
| [`LOGISTICS_GLOBAL_1688_PACKAGE_CHANGE`](https://open.1688.com/doc/topicDetail.htm?id=LOGISTICS_GLOBAL_1688_PACKAGE_CHANGE) | 1688 cross-border logistics package update | 1688跨境物流包裹消息 |

### Product Messages (商品消息)

| Topic | Description | Original |
|---|---|---|
| [`PRODUCT_RELATION_VIEW_PRODUCT_EXPIRE`](https://open.1688.com/doc/topicDetail.htm?id=PRODUCT_RELATION_VIEW_PRODUCT_EXPIRE) | Product delisted (related-user view) | 1688产品下架（关系用户视角） |
| [`PRODUCT_RELATION_VIEW_PRODUCT_NEW_OR_MODIFY`](https://open.1688.com/doc/topicDetail.htm?id=PRODUCT_RELATION_VIEW_PRODUCT_NEW_OR_MODIFY) | Product added or modified (related-user view) | 1688产品新增或修改（关系用户视角） |
| [`PRODUCT_RELATION_VIEW_PRODUCT_DELETE`](https://open.1688.com/doc/topicDetail.htm?id=PRODUCT_RELATION_VIEW_PRODUCT_DELETE) | Product deleted (related-user view) | 1688产品删除（关系用户视角） |
| [`PRODUCT_RELATION_VIEW_PRODUCT_REPOST`](https://open.1688.com/doc/topicDetail.htm?id=PRODUCT_RELATION_VIEW_PRODUCT_REPOST) | Product relisted (related-user view) | 1688产品上架（关系用户视角） |
| [`PRODUCT_PRODUCT_INVENTORY_CHANGE`](https://open.1688.com/doc/topicDetail.htm?id=PRODUCT_PRODUCT_INVENTORY_CHANGE) | Product inventory changed (related-user view) | 1688商品库存变更消息（关系用户视角） |
| [`PRODUCT_RELATION_VIEW_PRODUCT_AUDIT`](https://open.1688.com/doc/topicDetail.htm?id=PRODUCT_RELATION_VIEW_PRODUCT_AUDIT) | Product audit result (related-user view) | 1688产品审核（关系用户视角） |
| [`PRODUCT_PFT_OFFER_QUIT`](https://open.1688.com/doc/topicDetail.htm?id=PRODUCT_PFT_OFFER_QUIT) | Curated-supply product delisted | 精选货源商品下架消息 |
| [`PRODUCT_PFT_OFFER_PRICE_MODIFY`](https://open.1688.com/doc/topicDetail.htm?id=PRODUCT_PFT_OFFER_PRICE_MODIFY) | Curated-supply product price changed | 精选货源商品价格变动消息 |
| [`PRODUCT_PRODUCT_CROSSBOARD_INFORM`](https://open.1688.com/doc/topicDetail.htm?id=PRODUCT_PRODUCT_CROSSBOARD_INFORM) | One-click listing notification | 一键铺货消息 |
| [`PRODUCT_RELATION_VIEW_PRODUCT_CHANGE`](https://open.1688.com/doc/topicDetail.htm?id=PRODUCT_RELATION_VIEW_PRODUCT_CHANGE) | Product changed (related-user view; covers every product change action) | 商品变更消息(关系用户视角、包含所有商品变更动作) |
| [`FENXIAO_PRICE_CHANGE`](https://open.1688.com/doc/topicDetail.htm?id=FENXIAO_PRICE_CHANGE) | Distribution price changed | 分销价格变更 |
| [`CROSSBOARD_CROSSBOARD_ADD_SUPPLY`](https://open.1688.com/doc/topicDetail.htm?id=CROSSBOARD_CROSSBOARD_ADD_SUPPLY) | Product set as cross-border supply source | 跨境设为货源 |
| [`CROSSBOARD_XYT_PLATFORM_OFFER_CHANGE`](https://open.1688.com/doc/topicDetail.htm?id=CROSSBOARD_XYT_PLATFORM_OFFER_CHANGE) | XunYuanTong workbench product change | 寻源通工作台商品变更信息 |
| [`CROSSBOARD_LP_DISTRIBUTION`](https://open.1688.com/doc/topicDetail.htm?id=CROSSBOARD_LP_DISTRIBUTION) | Cross-border purchasing assistant one-click listing | 跨境采购助手一键铺货消息 |

### 88 ShengYiTong (Business Link) (88生意通)

| Topic | Description | Original |
|---|---|---|
| [`SYT_CONTRACT_REJECT`](https://open.1688.com/doc/topicDetail.htm?id=SYT_CONTRACT_REJECT) | Purchase order / contract rejected | 采购单/合同拒绝消息通知 |
| [`SYT_CONTRACT_PAY_SUCCESS`](https://open.1688.com/doc/topicDetail.htm?id=SYT_CONTRACT_PAY_SUCCESS) | Contract / purchase order paid successfully by both parties | 合同/采购单双方支付成功消息通知 |
| [`SYT_CONTRACT_WAIT_SIGN`](https://open.1688.com/doc/topicDetail.htm?id=SYT_CONTRACT_WAIT_SIGN) | Contract or purchase order awaiting signature | 合同或采购单待签署消息 |
| [`SYT_CONTRACT_SIGN_SUCCESS`](https://open.1688.com/doc/topicDetail.htm?id=SYT_CONTRACT_SIGN_SUCCESS) | Contract / purchase order signed successfully by both parties | 合同/采购单双方签署成功通知 |
| [`SYT_ADD_CONTRACT_CONTENT`](https://open.1688.com/doc/topicDetail.htm?id=SYT_ADD_CONTRACT_CONTENT) | New transaction evidence record added | 新增交易存证通知消息 |
| [`SYT_REFUND_FINISH_NOTICE`](https://open.1688.com/doc/topicDetail.htm?id=SYT_REFUND_FINISH_NOTICE) | Contract or purchase order refund result | 合同或采购单退款结果消息 |

### Inquiries (Newton Cloud) (询盘)

| Topic | Description | Original |
|---|---|---|
| [`AGENT_NEWTON_CLOUD_TASK_NOTIFY`](https://open.1688.com/doc/topicDetail.htm?id=AGENT_NEWTON_CLOUD_TASK_NOTIFY) | Newton Cloud task execution notification queue | 牛顿云任务执行通知队列 |

## Change log

### 2026-08-25 11:51:28

- Added API: Query repurchase contract (`com.alibaba.trade:repurchase.contract.get-1`)

### 2026-08-19 11:54:59

- Added API: Submit merged-invoice request (`com.alibaba.trade:trade.invoice.mergeapply-1`)
- Added API: Merged-invoice consultation (`com.alibaba.trade:trade.invoice.consult-1`)
- Added API: Query merged-invoice relationships by order or invoice application (`com.alibaba.trade:trade.invoice.sellerqueryrelatedorders-1`)

### 2026-08-17 10:22:37

- Removed API: 1688 cross-border: order inquiry (`com.alibaba.fenxiao.crossborder:inquiry.task.batchOrder-1`)

### 2026-08-13 14:45:30

- Added API: Newton Cloud: query points details (`com.alibaba.agent:newtoncloud.points.query-1`)

### 2026-08-07 11:55:44

- Added API: Newton Cloud: list available model tiers (`com.alibaba.agent:newtoncloud.model.list-1`)
- Added API: Newton Cloud: query task table (`com.alibaba.agent:newtoncloud.task.fetch-1`)
- Added API: Newton Cloud: list tasks (`com.alibaba.agent:newtoncloud.task.list-1`)
- Added API: Newton Cloud: query batch-inquiry results (`com.alibaba.agent:newtoncloud.batchInquiry.getResult-1`)
- Added API: Newton Cloud: create long-running task (`com.alibaba.agent:newtoncloud.task.create-1`)
- Added API: Newton Cloud: resume task (`com.alibaba.agent:newtoncloud.task.resume-1`)
- Added API: Newton Cloud: terminate task (`com.alibaba.agent:newtoncloud.task.kill-1`)
- Added API: Newton Cloud: upload file (`com.alibaba.agent:newtoncloud.file.upload-1`)
- Added API: Newton Cloud: query task (`com.alibaba.agent:newtoncloud.task.get-1`)
- Added message: Newton Cloud task execution notification queue (`AGENT_NEWTON_CLOUD_TASK_NOTIFY`)

### 2026-07-30 21:14:07

- Added API: Publish result callback (`com.alibaba.fenxiao.crossborder:publish.result.callback-1`)
- Added API: Listable product list (`com.alibaba.fenxiao.crossborder:publish.product.list-1`)
- Added API: Listing product card (`com.alibaba.fenxiao.crossborder:publish.card.get-1`)

### 2026-07-08 16:53:30

- Added API: Query the merchant's return address (`com.alibaba.fenxiao:refund.address.get-1`)

### 2026-06-17 17:42:15

- Added API: Get a link that opens a WangWang chat (`com.alibaba.account:account.wangwangUrl.get-1`)
- Added API: 88 ShengYiTong: contract refund request (`com.alibaba.syt:syt.contract.refund-1`)
- Added API: Cancel official return-pickup logistics order (`com.alibaba.logistics:refundofficialdelivery.order.cancel-1`)
- Added API: Query official return-pickup plans (`com.alibaba.logistics:refundofficialdelivery.solution.get-1`)
- Added API: Claim the optimal coupon for a product (`com.alibaba.marketing:coupon.optimal.claim-1`)
- Added API: Follow a product (`com.alibaba.product:alibaba.product.follow.crossborder-1`)
- Added API: Get 90-day repurchase-rate trend for a product (new) (`com.alibaba.fenxiao.crossborder:product.analyze.getRepurchaseRateTrendNew-1`)
- Added API: Get artwork info for a customization order (`com.alibaba.fenxiao.crossborder:product.podOrderDesign.get-1`)
- Added API: Batch-confirm failed messages (`cn.alibaba.open:push.message.confirm-1`)
- Added API: 88 ShengYiTong: buyer confirms purchase order (`com.alibaba.syt:syt.buyer.confirmPurchaseOrder-1`)
- Added API: Get product details (cross-border) (`com.alibaba.product:alibaba.cross.productInfo-1`)
- Added API: Query category by category ID (`com.alibaba.product:alibaba.category.get-1`)
- Added API: View order list (seller view) (`com.alibaba.trade:alibaba.trade.getSellerOrderList-1`)
- Added API: 88 ShengYiTong: buyer payment, get cashier URL (`com.alibaba.syt:syt.contract.pay-1`)
- Added API: 1688 cross-border: order inquiry (`com.alibaba.fenxiao.crossborder:inquiry.task.batchOrder-1`)
- Added API: Batch-query sub-account authorizations (`system.oauth2:subaccount.auth.list-1`)
- Added API: Add products to the listing list (cross-border) (`com.alibaba.product.push:alibaba.cross.syncProductListPushed-1`)
- Added API: 88 ShengYiTong: check whether customer is verified (`com.alibaba.syt:syt.customer.queryUserAuthStatus-1`)
- Added API: 88 ShengYiTong: WorldFirst transfer payment (WorldFirst B2C accounts only) (`com.alibaba.syt:syt.contract.payTransfer-1`)
- Added API: Save the business line an account belongs to (`com.alibaba.fenxiao.crossborder:account.business.save-1`)
- Added API: Ship: no logistics needed (`com.alibaba.logistics:alibaba.logistics.OpDeliverySendOrder.dummy-1`)
- Added API: Decrypt an OpenUID into a WangWang nickname (only for launching WangWang) (`com.alibaba.account:wangwangnick.openuid.decrypt-1`)
- Added API: Get simple product info from a previously purchased supplier (`com.alibaba.product:alibaba.product.simple.get-1`)
- Added API: Write back country-site logistics orders (`com.alibaba.fenxiao.crossborder:trade.cross.logisticsOrderSync-1`)
- Added API: Query ranking lists (`com.alibaba.fenxiao.crossborder:product.topList.query-1`)
- Added API: Sync listing results (`com.alibaba.product.push:alibaba.product.push.syncPushProductResult-1`)
- Added API: Urge seller to ship (`com.alibaba.logistics:logistics.delivery.urge-1`)
- Added API: [Relationship] Distributor: query supplier list (`cn.alibaba.open:alibaba.relation.querySuppliers-1`)
- Added API: Get 30-day median-price trend for a product (new) (`com.alibaba.fenxiao.crossborder:product.analyze.getThirtyDayMedianPriceTrendNew-1`)
- Added API: Buyer deletes a closed order (`com.alibaba.trade:trade.order.buyerdelete-1`)
- Added API: Cancel refund/return request (`com.alibaba.trade:alibaba.trade.cancelRefund-1`)
- Added API: 88 ShengYiTong: confirm transaction complete (`com.alibaba.syt:syt.contract.confirm-1`)
- Added API: Query seller mixed-batch settings (`com.alibaba.trade:alibaba.trade.OpQueryMarketingMixConfig-1`)
- Added API: 88 ShengYiTong: void contract (`com.alibaba.syt:syt.contract.invalid-1`)
- Added API: Get daily sales-quantity trend for a product (new) (`com.alibaba.fenxiao.crossborder:product.analyze.getPerdaySellQuantityTrendNew-1`)
- Added API: Query shipping-insurance info (`com.alibaba.trade:shipping.insurance.get-1`)
- Added API: Run AI supplier-search task (async) (`com.alibaba.fenxiao.crossborder:account.search.getInfo.async-1`)
- Added API: Modify official return-pickup logistics order (`com.alibaba.logistics:refundofficialdelivery.order.modify-1`)
- Added API: Ship: seller arranges own logistics (`com.alibaba.logistics:alibaba.logistics.OpDeliverySendOrder.offline-1`)
- Added API: Create official return-pickup logistics order (`com.alibaba.logistics:refundofficialdelivery.order.create-1`)
- Added API: Parse overseas address (`com.alibaba.trade:trade.address.parseText-1`)
- Added API: Get official return-pickup logistics order details (`com.alibaba.logistics:refundofficialdelivery.order.get-1`)
- Added API: Fetch failed messages (cursor style) (`cn.alibaba.open:push.cursor.messageList-1`)
- Added API: Add buyer-seller distribution relationship (`com.alibaba.fenxiao:alibaba.fenxiao.relationadd-1`)
- Added API: Unfollow a product (`com.alibaba.product:alibaba.product.unfollow.crossborder-1`)
- Added API: 88 ShengYiTong: buyer drafts purchase order (`com.alibaba.syt:syt.buyer.draftPurchaseOrder-2`)
- Added API: 88 ShengYiTong: buyer queries purchase-order details (`com.alibaba.syt:syt.buyer.queryContractDetail-1`)
- Added API: Query AI supplier-search task result (`com.alibaba.fenxiao.crossborder:account.search.getInfoResult-1`)
- Added API: Write back mapping between end-customer orders and 1688 orders (`com.alibaba.fenxiao.crossborder:order.relation.write-1`)
- Added API: Get product list (cross-border) (`com.alibaba.product:alibaba.cross.productList-1`)
- Added API: Recommend products by keyword (cross-border) (`com.alibaba.product:alibaba.product.suggest.crossBorder-1`)
- Added API: Encrypt a user loginId into an OpenUID (`com.alibaba.account:loginid.openuid.encrypt-1`)
- Added API: Query refund details by refund ID (buyer view) (`com.alibaba.trade:alibaba.trade.refund.OpQueryOrderRefund-1`)
- Added message: Product added or modified (related-user view) (`PRODUCT_RELATION_VIEW_PRODUCT_NEW_OR_MODIFY`)
- Added message: Product set as cross-border supply source (`CROSSBOARD_CROSSBOARD_ADD_SUPPLY`)
- Added message: 1688 cross-border logistics package update (`LOGISTICS_GLOBAL_1688_PACKAGE_CHANGE`)
- Added message: Order closed by operations back-office (buyer view) (`ORDER_BUYER_VIEW_ORDER_BOPS_CLOSE`)
- Added message: Purchase order / contract rejected (`SYT_CONTRACT_REJECT`)
- Added message: Contract / purchase order signed successfully by both parties (`SYT_CONTRACT_SIGN_SUCCESS`)
- Added message: New transaction evidence record added (`SYT_ADD_CONTRACT_CONTENT`)
- Added message: Cross-border purchasing assistant one-click listing (`CROSSBOARD_LP_DISTRIBUTION`)
- Added message: Curated-supply product price changed (`PRODUCT_PFT_OFFER_PRICE_MODIFY`)
- Added message: One-click listing notification (`PRODUCT_PRODUCT_CROSSBOARD_INFORM`)
- Added message: Order stage payment (buyer view) (`ORDER_BUYER_VIEW_ORDER_STEP_PAY`)
- Added message: Contract or purchase order refund result (`SYT_REFUND_FINISH_NOTICE`)
- Added message: Curated-supply product delisted (`PRODUCT_PFT_OFFER_QUIT`)
- Added message: Contract or purchase order awaiting signature (`SYT_CONTRACT_WAIT_SIGN`)
- Added message: Contract / purchase order paid successfully by both parties (`SYT_CONTRACT_PAY_SUCCESS`)
- Added message: Order partially shipped (buyer view) (`ORDER_BUYER_VIEW_PART_PART_SENDGOODS`)
- Removed API: Get 90-day repurchase-rate trend for a product (`com.alibaba.fenxiao.crossborder:product.analyze.getRepurchaseRateTrend-1`)
- Removed API: Get daily sales-quantity trend for a product (`com.alibaba.fenxiao.crossborder:product.analyze.getPerdaySellQuantityTrend-1`)
- Removed API: Get 30-day median-price trend for a product (`com.alibaba.fenxiao.crossborder:product.analyze.getThirtyDayMedianPriceTrend-1`)

### 2026-04-20 10:07:02

- Added API: Query maximum refundable amount when applying (`com.alibaba.trade:alibaba.trade.getMaxRefundFee-1`)

### 2026-04-14 19:17:10

- Added API: 1688 member registration (`com.alibaba.fenxiao.crossborder:account.user.register-1`)
- Added API: Sync downstream sales orders (`com.alibaba.fenxiao.crossborder:trade.cross.orderSync-1`)
- Added API: Query multilingual category by category name (`com.alibaba.fenxiao.crossborder:category.translation.getByKeyword-1`)

### 2026-03-03 16:52:59

- Added API: Query all issued invoices for an order (buyer view) (`com.alibaba.trade:trade.invoice.getListBuyerView-1`)
- Added API: Query invoiceable amount of orders (`com.alibaba.trade:trade.invoiceAmount.getList-1`)
- Added API: Query buyer invoice titles (paginated) (`com.alibaba.trade:trade.invoiceTitle.getPageList-1`)
- Added API: Query invoice applications (buyer view, paginated) (`com.alibaba.trade:trade.invoiceApply.getPageListBuyerView-1`)
- Added API: Buyer requests invoice (`com.alibaba.trade:trade.invoice.apply-1`)
- Added API: Buyer requests a shipping-address change (`com.alibaba.trade:order.receiveAddress.buyerUpdate-1`)

### 2026-02-27 10:23:15

- Added API: Get combined-cashier URL (`com.alibaba.trade:alibaba.trade.grouppay.url.get-1`)

### 2025-12-29 15:56:11

- Removed API: Smart image object removal (`com.alibaba.fenxiao.crossborder:image.elements.remove-1`)
- Removed API: HD image upscaling (`com.alibaba.fenxiao.crossborder:image.elements.enlarge-1`)
- Removed API: Product description generation (`com.alibaba.fenxiao.crossborder:text.generate.description-1`)
- Removed API: Image translation (`com.alibaba.fenxiao.crossborder:image.elements.translate-1`)
- Removed API: Product title generation (`com.alibaba.fenxiao.crossborder:text.generate.title-1`)
- Removed API: Smart image cutout (matting) (`com.alibaba.fenxiao.crossborder:image.elements.matting-1`)
- Removed API: Image element recognition (`com.alibaba.fenxiao.crossborder:image.elements.recognition-1`)
- Removed API: Product text translation (`com.alibaba.fenxiao.crossborder:product.text.translate-1`)
- Removed API: Image cropping (`com.alibaba.image:image.elements.cut-1`)

### 2025-12-04 20:53:15

- Added API: Batch-cancel sub-account authorizations (`system.oauth2:subaccount.auth.cancel-1`)
- Added API: Batch-add sub-account authorizations (`system.oauth2:subaccount.auth.add-1`)

### 2025-09-10 13:56:41

- Added API: Unfollow a product (cross-border) (`com.alibaba.fenxiao.crossborder:product.kjdistribute.removeRelation-1`)

### 2025-05-28 11:51:44

- Added message: Order closed by seller (buyer view) (`ORDER_BUYER_VIEW_ORDER_SELLER_CLOSE`)

### 2025-04-16 11:20:55

- Added message: Order price modified (buyer view) (`ORDER_BUYER_VIEW_ORDER_PRICE_MODIFY`)

### 2025-03-10 16:58:58

- Added API: Smart image object removal (`com.alibaba.fenxiao.crossborder:image.elements.remove-1`)
- Added API: HD image upscaling (`com.alibaba.fenxiao.crossborder:image.elements.enlarge-1`)
- Added API: Product description generation (`com.alibaba.fenxiao.crossborder:text.generate.description-1`)
- Added API: Multilingual search navigation (`com.alibaba.fenxiao.crossborder:product.search.keywordSNQuery-1`)
- Added API: Image translation (`com.alibaba.fenxiao.crossborder:image.elements.translate-1`)
- Added API: Product title generation (`com.alibaba.fenxiao.crossborder:text.generate.title-1`)
- Added API: Trending product search keywords (`com.alibaba.fenxiao.crossborder:product.search.topKeyword-1`)
- Added API: Smart image cutout (matting) (`com.alibaba.fenxiao.crossborder:image.elements.matting-1`)
- Added API: Image element recognition (`com.alibaba.fenxiao.crossborder:image.elements.recognition-1`)
- Added API: Product text translation (`com.alibaba.fenxiao.crossborder:product.text.translate-1`)
- Added API: Image cropping (`com.alibaba.image:image.elements.cut-1`)
- Added API: Related product recommendations (`com.alibaba.fenxiao.crossborder:product.related.recommend-1`)

## SDK

Welcome to the Alibaba Open Platform SDK. The SDK already encapsulates request signing and signature verification, so you can call the APIs directly through it.

When calling an API, first work out the class name that corresponds to the interface. In the SDK, capitalise the first letter of each word, remove the "." separators, and append Param (request class) or Result (response class). For example, the interface alibaba.product.add corresponds to AlibabaProductAddParam (request class) and AlibabaProductAddResult (response class).

Then complete the API call with the following code:

```java
ApiExecutor apiExecutor = new ApiExecutor("{your app key}", "{your app secret}");
AlibabaProductAddParam param = new AlibabaProductAddParam();
//TODO set the param fields
apiExecutor.execute(param, "{your access token}");
```

SDK language versions:

- Java 1.6 or later
- PHP 5.6 or later
- .NET requires .NET Framework 4.0 or later

Downloads (SDK guide: http://open.1688.com/doc/apiSdk.htm):

- python: https://ocean-sdk.oss-cn-hangzhou.aliyuncs.com/1703167397970-1-python.zip
- java: https://ocean-sdk.oss-cn-hangzhou.aliyuncs.com/1703167397970-1-java.zip
- php: https://ocean-sdk.oss-cn-hangzhou.aliyuncs.com/1703167397970-1-php.zip
- net: https://ocean-sdk.oss-cn-hangzhou.aliyuncs.com/1703167397970-1-net.zip
