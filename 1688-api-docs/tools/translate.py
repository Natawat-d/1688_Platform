# -*- coding: utf-8 -*-
"""Build English translation of the 1688 solution page (key 1703167397970)."""
import json, html, datetime, sys

SOL = json.load(open('solution.json', encoding='utf-8'))['result']
LST = json.load(open('apilist.json', encoding='utf-8'))

# ---------------------------------------------------------------- translations
CATS = {
    '会员': 'Membership & Accounts',
    '工具': 'Tools',
    '商机': 'Market Insights',
    '类目': 'Categories',
    '商品': 'Products',
    '选品铺货': 'Product Selection & Listing',
    '订单': 'Orders',
    '支付': 'Payment',
    '物流': 'Logistics',
    '官方退上门取件': 'Official Return Pickup',
    '回传数据': 'Data Write-back',
    '退货退款': 'Returns & Refunds',
    '消息': 'Messaging',
    '轻定制': 'Light Customization',
    '商家': 'Suppliers',
    '询盘': 'Inquiries (Newton Cloud)',
    '88生意通': '88 ShengYiTong (Business Link)',
    '发票': 'Invoicing',
    '全托管': 'Fully Managed (Consignment)',
    '铺货': 'Listing (Publishing)',
    '复购': 'Repurchase',
    # message groups
    '交易消息': 'Trade Messages',
    '物流消息': 'Logistics Messages',
    '商品消息': 'Product Messages',
}

# fullName -> (display name EN, description EN)
API = {
 # Membership
 'system.oauth2:subaccount.auth.add:1': ('Batch-add sub-account authorizations',
   'Batch-add authorization for the sub-accounts under a main account. The main account must already be authorized.'),
 'system.oauth2:subaccount.auth.cancel:1': ('Batch-cancel sub-account authorizations',
   'Cancel the authorization of sub-accounts in batch.'),
 'system.oauth2:subaccount.auth.list:1': ('Batch-query sub-account authorizations',
   'Batch-query the authorization status of the sub-accounts under a main account.'),
 'com.alibaba.fenxiao.crossborder:account.user.register:1': ('1688 member registration',
   'Register a 1688 member.'),
 'com.alibaba.account:alibaba.account.agent.crossBasic:1': ('Get basic info of a non-authorized user (cross-border)',
   'View another user\'s basic information. For use in cross-border scenarios.'),
 'com.alibaba.account:alibaba.account.basic:1': ('Get basic info of the authorized user',
   'Get the basic information of the authorized user.'),
 'cn.alibaba.open:querySubAccount:1': ('Query sub-account info',
   'Query sub-account information.'),
 # Tools
 'com.alibaba.fenxiao.crossborder:pool.product.pull:1': ('Pull products from a product pool',
   'Batch-pull product data directly from a product pool by pool ID.'),
 'com.alibaba.fenxiao.crossborder:pool.product.total:1': ('Query product count in a product pool',
   'Query the total number of products in a product pool.'),
 'com.alibaba.account:loginid.openuid.encrypt:1': ('Encrypt a user loginId into an OpenUID',
   'Encrypts a user\'s loginId into an OpenUID. This interface is risk-controlled: batch operations are not allowed, and it may only be triggered manually by the merchant, for example when searching a user\'s orders or configuring rules.'),
 'com.alibaba.account:account.wangwangUrl.get:1': ('Get a link that opens a WangWang chat',
   'Get a link that launches an AliWangWang chat.'),
 'com.alibaba.account:wangwangnick.openuid.decrypt:1': ('Decrypt an OpenUID into a WangWang nickname (only for launching WangWang)',
   'Decrypts an OpenUID into a WangWang nickname. This interface is risk-controlled: it may only be used when a user needs to launch WangWang. Automated batch operations are not allowed, and it must not be used as a decryption interface to show plaintext to users. WangWang supports recall after encryption; do not use it for any other scenario.'),
 'com.alibaba.fenxiao:alibaba.fenxiao.relationadd:1': ('Add buyer-seller distribution relationship',
   'Add a buyer-seller distribution relationship by product ID.'),
 # Market insights
 'com.alibaba.fenxiao.crossborder:product.search.topKeyword:1': ('Trending product search keywords',
   'Trending product search keywords.'),
 'com.alibaba.fenxiao.crossborder:product.topList.query:1': ('Query ranking lists',
   'Query ranking (top) lists.'),
 'com.alibaba.fenxiao.crossborder:product.analyze.getPerdaySellQuantityTrendNew:1': ('Get daily sales-quantity trend for a product (new)',
   'Get a product\'s daily sales-quantity trend over 90 days (at most 90 days of data can be queried). New interface.'),
 'com.alibaba.fenxiao.crossborder:product.analyze.getThirtyDayMedianPriceTrendNew:1': ('Get 30-day median-price trend for a product (new)',
   'Get a product\'s 30-day median-price metric trend. New interface.'),
 'com.alibaba.fenxiao.crossborder:product.analyze.getRepurchaseRateTrendNew:1': ('Get 90-day repurchase-rate trend for a product (new)',
   'Get a product\'s 90-day repurchase-rate metric trend (at most 30 days can be queried). New interface.'),
 # Categories
 'com.alibaba.fenxiao.crossborder:category.translation.getById:1': ('Query multilingual category by category ID',
   'Multilingual category query. Returns the category details in the requested language for the given category ID, including the list of its child categories.'),
 'com.alibaba.fenxiao.crossborder:category.translation.getByKeyword:1': ('Query multilingual category by category name',
   'Multilingual category query. Returns the list of matching category details in the requested language for the given category name. Child-category data is not included.'),
 'com.alibaba.product:alibaba.category.attribute.get:1': ('Get leaf-category attributes',
   'Get category attributes by leaf-category ID.'),
 # Products
 'com.alibaba.fenxiao.crossborder:product.search.keywordQuery:1': ('Multilingual keyword search', 'Multilingual keyword search.'),
 'com.alibaba.fenxiao.crossborder:product.search.imageQuery:1': ('Multilingual image search', 'Multilingual image search.'),
 'com.alibaba.fenxiao.crossborder:product.search.querySellerOfferList:1': ('Multilingual in-store product search',
   'Multilingual search of the products in a seller\'s store.'),
 'com.alibaba.fenxiao.crossborder:product.search.queryProductDetail:1': ('Multilingual product detail', 'Multilingual product detail.'),
 'com.alibaba.fenxiao.crossborder:product.search.offerRecommend:1': ('Product recommendations', 'Personalized product recommendations.'),
 'com.alibaba.fenxiao.crossborder:product.search.keywordSNQuery:1': ('Multilingual search navigation',
   'Get the multilingual keyword-search navigation list.'),
 'com.alibaba.fenxiao.crossborder:product.related.recommend:1': ('Related product recommendations', 'Related product recommendations.'),
 'com.alibaba.fenxiao.crossborder:product.image.upload:1': ('Upload an image to get an imageId', 'Upload an image and receive an imageId.'),
 'com.alibaba.marketing:coupon.optimal.claim:1': ('Claim the optimal coupon for a product',
   'Claim the best-value coupon for a product. Usually called before placing an order to complete the optimal coupon-claiming strategy.'),
 'com.alibaba.product:alibaba.product.suggest.crossBorder:1': ('Recommend products by keyword (cross-border)',
   'Recommend products by keyword and category in cross-border scenarios, sorted by sales volume. Note: this API is rate-limited and is only suitable for manual-association scenarios.'),
 'com.alibaba.product:alibaba.product.simple.get:1': ('Get simple product info from a previously purchased supplier',
   'Get product details by product ID. This interface returns simple information for products of suppliers you have already purchased from. Access to this interface is paid. It returns only basic information and is mainly intended for data association in ERP systems.'),
 'com.alibaba.fenxiao.crossborder:product.kjdistribute.addRelation:1': ('Follow a product (cross-border)', 'Add a cross-border followed product.'),
 'com.alibaba.fenxiao.crossborder:product.kjdistribute.removeRelation:1': ('Unfollow a product (cross-border)', 'Remove a cross-border followed product.'),
 'com.alibaba.fenxiao.crossborder:product.distribute.getDistributeInfo:1': ('Get product selling points for listing',
   'For listing scenarios: get the selling-point information needed to list a product.'),
 'com.alibaba.fenxiao.crossborder:product.category.getAttrById:1': ('Query leaf-category attribute and value mappings',
   'Category mapping query.'),
 # Product selection & listing
 'com.alibaba.product:alibaba.cross.productInfo:1': ('Get product details (cross-border)',
   'Get product details in cross-border scenarios. A cross-border listing relationship must be established before details can be retrieved.'),
 'com.alibaba.product.push:alibaba.cross.syncProductListPushed:1': ('Add products to the listing list (cross-border)',
   'Cross-border only. Adds products to the listing list (that is, creates listing relationships); at most 20 items per call. Only after a product is added can its details be queried through the product-detail interface. Contact the cross-border operations staff to configure permissions manually before calling.'),
 'com.alibaba.product.push:alibaba.product.push.syncPushProductResult:1': ('Sync listing results',
   'Sync listing results. When an ISV lists products from the source platform (1688) onto a target platform (for example TAOBAO), the ISV must return the listing result. The listing-status descriptions must match those defined by the source platform (1688). This interface also supports operations such as delisting, all of which are expressed through the listing status.'),
 'com.alibaba.product:alibaba.cross.productList:1': ('Get product list (cross-border)',
   'This interface validates the cross-border listing relationship and is for cross-border business only. Product model V2.'),
 'com.alibaba.product:alibaba.product.unfollow.crossborder:1': ('Unfollow a product', 'Unfollow a product.'),
 'com.alibaba.product:alibaba.category.get:1': ('Query category by category ID',
   'Category query. To retrieve all 1688 categories, traverse the whole category tree starting from the root: first pass 0 to get all level-1 category IDs, then iterate the level-1 IDs to get all level-2 categories, and finally iterate the level-2 IDs to get the level-3 categories. Note: 1688 categories have only three levels; level-3 categories are the leaf categories required for publishing products.'),
 'com.alibaba.product:alibaba.product.follow.crossborder:1': ('Follow a product', 'Follow a product.'),
 'cn.alibaba.open:alibaba.relation.querySuppliers:1': ('[Relationship] Distributor: query supplier list',
   'Get the supplier list for a distributor by userID.'),
 # Orders
 'com.alibaba.trade:alibaba.trade.createCrossOrder:1': ('Create cross-border order',
   'Cross-border-only order creation. An order may contain at most 50 SKUs, all from the same supplier. For multiple suppliers or more than 50 SKUs, split the order yourself before submitting. In some special cases several orders are created at once and several order numbers are returned. Supports both the open-marketplace and distribution scenarios. Orders are placed under the main account or a sub-account depending on the currently authorized user.'),
 'com.alibaba.trade:alibaba.trade.get.buyerView:1': ('View order details (buyer view)',
   'Get the details of a single transaction; buyer calls only. Permission must be requested from the Alibaba Open Platform to use this API.'),
 'com.alibaba.logistics:alibaba.trade.getLogisticsInfos.buyerView:1': ('Get order logistics info (buyer view)',
   'Requires the order buyer\'s authorization and returns the logistics details of the buyer\'s order. In procurement or distribution scenarios, buyers also need logistics details. This interface looks up logistics details by order number, including the sender, the recipient and the details of the goods shipped. Because of when logistics records are entered, API tracking queries may be delayed. Permission must be requested from the Open Platform to access this API.'),
 'com.alibaba.logistics:alibaba.trade.getLogisticsTraceInfo.buyerView:1': ('Get order logistics tracking info (buyer view)',
   'Requires the order buyer\'s authorization and returns the logistics tracking information of the buyer\'s order. In procurement or distribution scenarios, buyers also need logistics details. This interface looks up tracking information by logistics (waybill) number. Because of when logistics records are entered, API tracking queries may be delayed. Permission must be requested from the Open Platform to access this API.'),
 'com.alibaba.trade:alibaba.trade.cancel:1': ('Cancel transaction',
   'Buyer or seller cancels a transaction. Only transactions in specific statuses can be cancelled; on 1688 this is used to cancel unpaid orders. If an order is closed less than 10 seconds after it was created, the error CLOSE_ORDER_TOO_FAST is returned.'),
 'com.alibaba.trade:alibaba.trade.getBuyerOrderList:1': ('View order list (buyer view)',
   'Get the buyer\'s order list; the user\'s memberId must equal the buyer memberId on the orders. This interface returns only basic order information, not logistics or invoice information. For logistics information call the order-detail interface; for invoice information call the invoice-information API.'),
 'com.alibaba.trade:alibaba.createOrder.preview:1': ('Preview data before creating an order',
   'Orders may only contain products from a single supplier. This interface returns discount and related information for order creation. It 1. validates whether the products may be ordered; 2. validates the consignment (distribution) relationship; 3. validates stock, minimum order quantity and whether mixed-batch conditions are met.'),
 'com.alibaba.trade:alibaba.order.memoAdd:1': ('Update order memo',
   'If the authorized user is the seller, updates the seller memo; if the buyer, updates the buyer memo. Note: this interface can be called repeatedly, and the memo overwrites the content of the previous call.'),
 'com.alibaba.account:alibaba.subAccount.list:1': ('Get sub-account list',
   'Get the user\'s main-account and sub-account information. If the API is authorized as a sub-account, only the main account that the sub-account belongs to is returned. If authorized as a main account, the list of all sub-accounts is returned.'),
 'com.alibaba.trade:alibaba.trade.addFeedback:1': ('Buyer adds an order message',
   'Buyer adds a supplementary message to an order. The total message length must not exceed 500 characters.'),
 'com.alibaba.trade:trade.receivegoods.confirm:1': ('Buyer confirms receipt', 'Buyer confirms receipt of goods.'),
 'com.alibaba.trade:trade.order.buyerdelete:1': ('Buyer deletes a closed order', 'Buyer deletes an order that has been closed.'),
 'com.alibaba.trade:order.receiveAddress.buyerUpdate:1': ('Buyer requests a shipping-address change',
   'Buyer changes the shipping address. If the new address is in a remote area or the shipping fee must be recalculated, the seller must confirm the change and may reject it.'),
 # Payment
 'com.alibaba.trade:alibaba.alipay.url.get:1': ('Batch-get payment links for orders',
   'When paying through an ERP, use this API to get a cashier link for batch payment. A single order returns the 1688 cashier URL; multiple orders return the Alipay cashier URL. The ERP can redirect the user to the cashier link to complete payment. The user\'s 1688 login status is verified before payment.'),
 'com.alibaba.trade:alibaba.crossBorderPay.url.get:1': ('Get payment link for Cross-Border Pay (Kuajingbao)',
   'Get a payment link for paying with Cross-Border Pay (Kuajingbao).'),
 'com.alibaba.trade:alibaba.creditPay.url.get:1': ('Get payment link for Cheng-e-She credit pay',
   'Get a payment link for paying with Cheng-e-She (buy-now-pay-later credit).'),
 'com.alibaba.trade:alibaba.accountPeriod.list.buyerView:1': ('Buyer views all granted credit terms',
   'View, from the buyer\'s side, all account-period (credit-term) lines the buyer has been granted. Paginated; at most 10 records per call.'),
 'com.alibaba.trade:alibaba.trade.payWay.query:1': ('Query payment channels supported by an order',
   'Query the payment methods or channels available for an unpaid order.'),
 'com.alibaba.trade:alibaba.trade.pay.protocolPay.isopen:1': ('Check whether password-free payment is enabled',
   'Check whether an auto-debit (withholding) agreement is enabled.'),
 'com.alibaba.trade:alibaba.trade.pay.protocolPay.preparePay:1': ('Initiate password-free payment',
   'Initiates a password-free payment. Automatically detects whether Alipay or Cheng-e-She password-free payment is enabled and initiates the debit. Cheng-e-She auto-debit is tried first; if it fails, Alipay auto-debit is attempted. The error codes returned by this interface are currently not detailed; after a failed debit, retry up to 3 times rather than indefinitely.'),
 'com.alibaba.trade:trade.orderpay.analysis:1': ('Order payment consultation',
   'Order-payment consultation interface, used to analyse which payment method an order uses, and so on.'),
 'com.alibaba.trade:alibaba.trade.grouppay.url.get:1': ('Get combined-cashier URL', 'Get the combined-cashier URL.'),
 # Logistics
 'com.alibaba.fenxiao.crossborder:product.freight.estimate:1': ('Estimate domestic (China) shipping fee for a product',
   'Estimate a product\'s shipping fee from the product ID and the province/city/district codes of a delivery address within mainland China.'),
 'com.alibaba.fenxiao.crossborder:logistics.order.getOutOrderId:1': ('Query external order ID by waybill number or unclaimed-parcel code',
   'Query the external order ID by waybill number or unclaimed-parcel code.'),
 'com.alibaba.trade:shipping.insurance.get:1': ('Query shipping-insurance info', 'Query shipping-insurance information.'),
 'com.alibaba.logistics:alibaba.logistics.myFreightTemplate.list.get:1': ('Get shipping-template details',
   'Get the seller\'s shipping template by template ID. Template ID 0 means "shipping fee to be explained"; 1 means the seller bears the shipping fee.'),
 'com.alibaba.trade:alibaba.trade.receiveAddress.get:1': ('Buyer gets saved shipping addresses',
   'Get the buyer\'s list of saved shipping addresses.'),
 'com.alibaba.trade:alibaba.trade.addresscode.parse:1': ('Parse an address into area codes',
   'Parse area codes from address information.'),
 'com.alibaba.trade:alibaba.trade.OpQueryMarketingMixConfig:1': ('Query seller mixed-batch settings',
   'Query the seller\'s mixed-batch (mixed wholesale) settings.'),
 'com.alibaba.logistics:alibaba.logistics.OpQueryLogisticCompanyList:1': ('Logistics company list (all companies)',
   'Get the names of all logistics companies.'),
 'com.alibaba.trade:trade.address.parseText:1': ('Parse overseas address', 'Parse an overseas address.'),
 'com.alibaba.logistics:logistics.delivery.urge:1': ('Urge seller to ship',
   'Urge the seller to ship. The order must still be in a not-yet-shipped status. Limited to once per 24 hours.'),
 'com.alibaba.fenxiao.crossborder:label.url.receive:1': ('Receive external print-label (UDF) URL',
   'Receive an external print-label (UDF) URL.'),
 'com.alibaba.fenxiao.crossborder:bigcustomer.mailNo.query:1': ('Get waybill-number set', 'Get the set of waybill numbers.'),
 # Official return pickup
 'com.alibaba.logistics:refundofficialdelivery.order.create:1': ('Create official return-pickup logistics order',
   'Create an official-logistics door-to-door pickup order for a return.'),
 'com.alibaba.logistics:refundofficialdelivery.order.modify:1': ('Modify official return-pickup logistics order',
   'Modify an official-logistics door-to-door pickup order for a return.'),
 'com.alibaba.logistics:refundofficialdelivery.order.get:1': ('Get official return-pickup logistics order details',
   'Get the details of an official-logistics door-to-door pickup order for a return.'),
 'com.alibaba.logistics:refundofficialdelivery.order.cancel:1': ('Cancel official return-pickup logistics order',
   'Cancel an official-logistics door-to-door pickup order for a return.'),
 'com.alibaba.logistics:refundofficialdelivery.solution.get:1': ('Query official return-pickup plans',
   'Query the available official door-to-door return-pickup plans.'),
 # Data write-back
 'com.alibaba.fenxiao.crossborder:order.relation.write:1': ('Write back mapping between end-customer orders and 1688 orders',
   'Write back to 1688 the mapping between the organisation\'s real end-customer orders and the 1688 orders generated under the organisation\'s account.'),
 'com.alibaba.fenxiao.crossborder:trade.cross.logisticsOrderSync:1': ('Write back country-site logistics orders',
   'Write back logistics orders from the country site.'),
 'com.alibaba.fenxiao.crossborder:trade.cross.orderSync:1': ('Sync downstream sales orders', 'Sync downstream sales orders.'),
 'com.alibaba.fenxiao.crossborder:account.business.save:1': ('Save the business line an account belongs to',
   'Save the business line that an account belongs to.'),
 # Returns & refunds
 'com.alibaba.trade:alibaba.trade.refund.OpQueryBatchRefundByOrderIdAndStatus:1': ('Query refund details by order ID (buyer view)',
   'For buyers; sellers should use alibaba.trade.refund.OpQueryBatchRefundByOrderIdAndStatus.sellerView. Queries the refund list for an order in real time. Currently only in-sale refunds (before the transaction completes) can be queried.'),
 'com.alibaba.trade:alibaba.trade.refund.OpQueryOrderRefundOperationList:1': ('Refund operation history (buyer view)',
   'For buyers; sellers should use alibaba.trade.refund.OpQueryOrderRefund.sellerView. Returns the buyer-side refund operation records.'),
 'com.alibaba.trade:alibaba.trade.refund.OpQueryOrderRefund:1': ('Query refund details by refund ID (buyer view)',
   'For buyers; sellers should use alibaba.trade.refund.OpQueryOrderRefund.sellerView. Queries refund details, including the list of refund operations. Permission must be requested from Alibaba to access this API.'),
 'com.alibaba.trade:alibaba.trade.refund.buyer.queryOrderRefundList:1': ('Query refund list (buyer view)',
   'Buyer views the refund list. This interface does not support sub-account queries; authorize with the main account before querying.'),
 'com.alibaba.trade:alibaba.trade.queryOrderByInsure:1': ('Query order number by claim or insurance-policy number',
   'Query the order number that corresponds to a claim or insurance-policy number. Pass type=lp for claims and type=bx for shipping insurance.'),
 'com.alibaba.trade:alibaba.trade.createRefund:1': ('Create refund/return request', 'Create a refund or return request.'),
 'com.alibaba.trade:alibaba.trade.getRefundReasonList:1': ('Query refund/return reasons (for creating a request)',
   'Query the refund/return reasons (used when creating a refund or return request).'),
 'com.alibaba.trade:alibaba.trade.uploadRefundVoucher:1': ('Upload refund/return evidence',
   'Upload evidence for a refund/return request. To convert a file stream to a byte array, org.apache.commons.io.IOUtils#toByteArray(java.io.InputStream) is recommended.'),
 'com.alibaba.trade:alibaba.trade.refund.returnGoods:1': ('Buyer submits return-shipment info',
   'Used after the seller approves the buyer\'s return/refund request, for the buyer to submit the return-shipment information. First call alibaba.logistics.OpQueryLogisticCompanyList.offline to look up logistics companies, and use the logistics-company code it returns.'),
 'com.alibaba.trade:alibaba.trade.cancelRefund:1': ('Cancel refund/return request', 'Cancel a refund or return request.'),
 'com.alibaba.trade:alibaba.trade.getMaxRefundFee:1': ('Query maximum refundable amount when applying',
   'Query the maximum refundable amount when applying for a refund.'),
 'com.alibaba.trade:trade.arbitration.apply:1': ('Apply for trade arbitration', 'Trade arbitration application.'),
 'com.alibaba.fenxiao:refund.address.get:1': ('Query the merchant\'s return address', 'Query the return address.'),
 # Messaging
 'cn.alibaba.open:push.message.confirm:1': ('Batch-confirm failed messages',
   'Manually call the confirmation API to confirm that messages have been consumed successfully. Only needed when using the query-style API to fetch failed messages.'),
 'cn.alibaba.open:push.query.messageList:1': ('Fetch failed messages (query style)',
   'Query-style retrieval of sent messages. Retrieved messages are not confirmed automatically; the caller must call the confirmation API to confirm the consumption status. Note that confirming affects the data returned by pagination.'),
 'cn.alibaba.open:push.cursor.messageList:1': ('Fetch failed messages (cursor style)',
   'Cursor-style retrieval of failed messages. Retrieved messages are automatically confirmed as consumed, so the next call with the same conditions returns the remaining data, until the result is empty.'),
 'com.alibaba.trade:alibaba.trade.getSellerOrderList:1': ('View order list (seller view)',
   'Get the seller\'s order list; the user\'s memberId must equal the sellerMemberId on the orders. This interface returns only basic order information, not logistics or invoice information. For logistics information call the order-detail interface; for invoice information call the invoice-information API.'),
 'com.alibaba.logistics:alibaba.logistics.OpDeliverySendOrder.offline:1': ('Ship: seller arranges own logistics',
   'For 1688 open-marketplace orders where the seller arranges logistics themselves. Supports combined shipping (several orders in one shipment) and shipping at sub-order (orderEntry) level. Shipping by quantity is not supported.'),
 'com.alibaba.logistics:alibaba.logistics.OpDeliverySendOrder.dummy:1': ('Ship: no logistics needed',
   'For 1688 open-marketplace orders that need no logistics. Supports combined shipping (several orders in one shipment) and shipping at sub-order (orderEntry) level. Shipping by quantity is not supported.'),
 # Light customization
 'com.alibaba.fenxiao.crossborder:product.podOrderDesign.get:1': ('Get artwork info for a customization order',
   'Get the design/artwork information of a processing-and-customization (print-on-demand) order.'),
 # Suppliers
 'com.alibaba.fenxiao.crossborder:account.search.getInfo.async:1': ('Run AI supplier-search task (async)',
   'Execute an AI supplier-search task asynchronously.'),
 'com.alibaba.fenxiao.crossborder:account.search.getInfoResult:1': ('Query AI supplier-search task result',
   'Query the execution result of an AI supplier-search task.'),
 # Inquiries / Newton Cloud
 'com.alibaba.agent:newtoncloud.task.create:1': ('Newton Cloud: create long-running task',
   'Create a Newton Cloud long-running task. It executes asynchronously and returns taskId/sessionId/status.'),
 'com.alibaba.agent:newtoncloud.task.get:1': ('Newton Cloud: query task', 'Newton Cloud task query service.'),
 'com.alibaba.agent:newtoncloud.task.list:1': ('Newton Cloud: list tasks',
   'Query the current user\'s task list. Returns each task\'s taskId/sessionId/status/taskType and its creation and completion times.'),
 'com.alibaba.agent:newtoncloud.task.kill:1': ('Newton Cloud: terminate task',
   'Terminate the task with the given taskId, mark it with the KILL status, and return killed.'),
 'com.alibaba.agent:newtoncloud.task.fetch:1': ('Newton Cloud: query task table', 'Newton Cloud task-table query service.'),
 'com.alibaba.agent:newtoncloud.task.resume:1': ('Newton Cloud: resume task', 'Resume a Newton Cloud task.'),
 'com.alibaba.agent:newtoncloud.batchInquiry.getResult:1': ('Newton Cloud: query batch-inquiry results',
   'Query inquiry results by batch-inquiry task ID (wwTaskId). After the ISV receives an AGENT_NEWTON_CLOUD_TASK_NOTIFY notification (stage=BATCH_INQUIRY, stageDetail.phase=COMPLETE), call this interface with the stageDetail.wwTaskId from the notification. It returns structured inquiry results synchronously, including supplier reply summaries, quotation details and recommendation reasons.'),
 'com.alibaba.agent:newtoncloud.file.upload:1': ('Newton Cloud: upload file',
   'Upload a local file to Newton Cloud OSS. Returns a temporary public download link and a stable download address, which can be used directly as the fileUrls parameter of newtoncloud.task.create. File content is passed as Base64; a single file of 3 MB or less is recommended.'),
 'com.alibaba.agent:newtoncloud.model.list:1': ('Newton Cloud: list available model tiers',
   'Query the model tiers currently available on Newton Cloud. Returns each tier\'s code (id), display name (displayName) and other information. The tier code can be used as the model parameter of newtoncloud.task.create. No business parameters; only access_token and signature are required.'),
 'com.alibaba.agent:newtoncloud.points.query:1': ('Newton Cloud: query points details',
   'Query the Newton Cloud points information of the currently authorized account, including total points, used points, available points, points sources, expiry information and paginated usage details. The user is identified automatically from the access_token; no userId is needed.'),
 # 88 ShengYiTong
 'com.alibaba.syt:syt.buyer.draftPurchaseOrder:2': ('88 ShengYiTong: buyer drafts purchase order',
   '88 ShengYiTong solution. The buyer drafts a purchase order.'),
 'com.alibaba.syt:syt.contract.confirm:1': ('88 ShengYiTong: confirm transaction complete',
   '88 ShengYiTong: confirm that the transaction is complete.'),
 'com.alibaba.syt:syt.contract.invalid:1': ('88 ShengYiTong: void contract', '88 ShengYiTong: void (invalidate) a contract.'),
 'com.alibaba.syt:syt.contract.refund:1': ('88 ShengYiTong: contract refund request', '88 ShengYiTong: apply for a contract refund.'),
 'com.alibaba.syt:syt.contract.pay:1': ('88 ShengYiTong: buyer payment, get cashier URL', '88 ShengYiTong solution.'),
 'com.alibaba.syt:syt.contract.payTransfer:1': ('88 ShengYiTong: WorldFirst transfer payment (WorldFirst B2C accounts only)',
   '88 ShengYiTong payment-transfer interface.'),
 'com.alibaba.syt:syt.buyer.confirmPurchaseOrder:1': ('88 ShengYiTong: buyer confirms purchase order',
   '88 ShengYiTong: the buyer confirms a purchase order.'),
 'com.alibaba.syt:syt.buyer.queryContractDetail:1': ('88 ShengYiTong: buyer queries purchase-order details',
   '88 ShengYiTong: the buyer queries the details of a purchase order.'),
 'com.alibaba.syt:syt.customer.queryUserAuthStatus:1': ('88 ShengYiTong: check whether customer is verified',
   '88 ShengYiTong: check whether a customer has completed verification.'),
 # Invoicing
 'com.alibaba.trade:trade.invoiceTitle.getPageList:1': ('Query buyer invoice titles (paginated)',
   'Paginated query of the buyer\'s invoice titles.'),
 'com.alibaba.trade:trade.invoiceAmount.getList:1': ('Query invoiceable amount of orders',
   'Query the invoiceable amount of orders. To find invoiceable orders first: 1. alibaba.trade.getBuyerOrderList-1: pass needInvoicingSetting=true and check invoicingSettingModel.tradeInvoiceStatus in the response to see whether an order can be invoiced. 2. alibaba.trade.get.buyerView-1: include InvoicingSetting in the includeFields parameter.'),
 'com.alibaba.trade:trade.invoiceApply.getPageListBuyerView:1': ('Query invoice applications (buyer view, paginated)',
   'Paginated query of invoice applications (buyer view).'),
 'com.alibaba.trade:trade.invoice.getListBuyerView:1': ('Query all issued invoices for an order (buyer view)',
   'Query all issued invoices associated with an order (buyer view).'),
 'com.alibaba.trade:trade.invoice.apply:1': ('Buyer requests invoice',
   'Buyer requests an invoice; batch requests are supported. Prerequisites: get the buyer\'s invoice titles; get the invoiceable orders and the invoice types they support; get the orders\' invoiceable amounts.'),
 'com.alibaba.trade:trade.invoice.consult:1': ('Merged-invoice consultation',
   'Submit orders to check whether they can be invoiced together and how they are grouped.'),
 'com.alibaba.trade:trade.invoice.mergeapply:1': ('Submit merged-invoice request',
   'First call the merged-invoice consultation trade.invoice.consult to get the grouping result, then call merged invoicing.'),
 'com.alibaba.trade:trade.invoice.sellerqueryrelatedorders:1': ('Query merged-invoice relationships by order or invoice application',
   'The seller queries merged-invoice relationships by order ID (orderId) or invoice-application ID (outbizid). If both are passed, outbizid takes precedence. If the result is not a merged invoice, mergeInvoice returns false.'),
 # Fully managed
 'com.alibaba.fenxiao.crossborder:consignment.co.status.sync:1': ('Warehouse receipt / shelving of goods',
   'Open to Buffalo. Syncs the Buffalo warehouse\'s receipt or shelving of goods to the 1688 shipment-order status.'),
 'com.alibaba.fenxiao.crossborder:consignment.tally.create:1': ('Warehouse creates discrepancy tally sheet',
   'The Buffalo warehouse creates a discrepancy tally sheet.'),
 'com.alibaba.fenxiao.crossborder:consignment.co.inventory:1': ('Update warehouse product inventory',
   'Update the inventory of products in the warehouse.'),
 # Listing (publishing)
 'com.alibaba.fenxiao.crossborder:publish.result.callback:1': ('Publish result callback',
   'WB-1688 integration Method 3: publish-result callback interface.'),
 'com.alibaba.fenxiao.crossborder:publish.product.list:1': ('Listable product list',
   'Query the list of products available for listing. Supports paginated queries of product-ID sets by assortment (goods-pool) ID.'),
 'com.alibaba.fenxiao.crossborder:publish.card.get:1': ('Listing product card',
   'Get the card data for a single 1688 product, converted to WB format.'),
 # Repurchase
 'com.alibaba.trade:repurchase.contract.get:1': ('Query repurchase contract',
   'Query whether a product supports a repurchase contract. A repurchase contract is a repeat-purchase discount agreement signed between buyer and seller, under which the seller gives repurchasing users special prices and guarantees. Orders under a repurchase contract must be placed through a specific trade flow.'),
}

# typeName -> description EN
MSG = {
 'ORDER_BUYER_VIEW_BUYER_MAKE': 'Order created (buyer view)',
 'ORDER_BUYER_VIEW_ORDER_PRICE_MODIFY': 'Order price modified (buyer view)',
 'ORDER_BUYER_VIEW_ORDER_PAY': 'Order paid (buyer view)',
 'ORDER_BUYER_VIEW_ANNOUNCE_SENDGOODS': 'Order shipped (buyer view)',
 'ORDER_BUYER_VIEW_PART_PART_SENDGOODS': 'Order partially shipped (buyer view)',
 'ORDER_BUYER_VIEW_ORDER_COMFIRM_RECEIVEGOODS': 'Receipt of goods confirmed (buyer view)',
 'ORDER_BUYER_VIEW_ORDER_SUCCESS': 'Transaction completed successfully (buyer view)',
 'ORDER_BUYER_VIEW_ORDER_BUYER_CLOSE': 'Order closed by buyer (buyer view)',
 'ORDER_BUYER_VIEW_ORDER_BOPS_CLOSE': 'Order closed by operations back-office (buyer view)',
 'ORDER_BUYER_VIEW_ORDER_BUYER_REFUND_IN_SALES': 'In-sale refund on order (buyer view)',
 'ORDER_BUYER_VIEW_ORDER_REFUND_AFTER_SALES': 'After-sales refund on order (buyer view)',
 'ORDER_BUYER_VIEW_ORDER_STEP_PAY': 'Order stage payment (buyer view)',
 'ORDER_BATCH_PAY': 'Order batch-payment status sync',
 'ORDER_BUYER_VIEW_ORDER_SELLER_CLOSE': 'Order closed by seller (buyer view)',
 'ORDER_ORDER_PRICE_MODIFY': 'Order price modified (seller view)',
 'LOGISTICS_BUYER_VIEW_TRACE': 'Logistics order status changed (buyer view)',
 'LOGISTICS_MAIL_NO_CHANGE': 'Waybill number changed',
 'LOGISTICS_GLOBAL_1688_PACKAGE_CHANGE': '1688 cross-border logistics package update',
 'PRODUCT_RELATION_VIEW_PRODUCT_EXPIRE': 'Product delisted (related-user view)',
 'PRODUCT_RELATION_VIEW_PRODUCT_NEW_OR_MODIFY': 'Product added or modified (related-user view)',
 'PRODUCT_RELATION_VIEW_PRODUCT_DELETE': 'Product deleted (related-user view)',
 'PRODUCT_RELATION_VIEW_PRODUCT_REPOST': 'Product relisted (related-user view)',
 'PRODUCT_PRODUCT_INVENTORY_CHANGE': 'Product inventory changed (related-user view)',
 'PRODUCT_RELATION_VIEW_PRODUCT_AUDIT': 'Product audit result (related-user view)',
 'PRODUCT_PFT_OFFER_QUIT': 'Curated-supply product delisted',
 'PRODUCT_PFT_OFFER_PRICE_MODIFY': 'Curated-supply product price changed',
 'PRODUCT_PRODUCT_CROSSBOARD_INFORM': 'One-click listing notification',
 'PRODUCT_RELATION_VIEW_PRODUCT_CHANGE': 'Product changed (related-user view; covers every product change action)',
 'FENXIAO_PRICE_CHANGE': 'Distribution price changed',
 'CROSSBOARD_CROSSBOARD_ADD_SUPPLY': 'Product set as cross-border supply source',
 'CROSSBOARD_XYT_PLATFORM_OFFER_CHANGE': 'XunYuanTong workbench product change',
 'CROSSBOARD_LP_DISTRIBUTION': 'Cross-border purchasing assistant one-click listing',
 'SYT_CONTRACT_REJECT': 'Purchase order / contract rejected',
 'SYT_CONTRACT_PAY_SUCCESS': 'Contract / purchase order paid successfully by both parties',
 'SYT_CONTRACT_WAIT_SIGN': 'Contract or purchase order awaiting signature',
 'SYT_CONTRACT_SIGN_SUCCESS': 'Contract / purchase order signed successfully by both parties',
 'SYT_ADD_CONTRACT_CONTENT': 'New transaction evidence record added',
 'SYT_REFUND_FINISH_NOTICE': 'Contract or purchase order refund result',
 'AGENT_NEWTON_CLOUD_TASK_NOTIFY': 'Newton Cloud task execution notification queue',
}

# change-log keys (ns:name-ver) that are no longer in the current list
EXTRA = {
 'com.alibaba.fenxiao.crossborder:inquiry.task.batchOrder-1': '1688 cross-border: order inquiry',
 'com.alibaba.fenxiao.crossborder:product.analyze.getRepurchaseRateTrend-1': 'Get 90-day repurchase-rate trend for a product',
 'com.alibaba.fenxiao.crossborder:product.analyze.getPerdaySellQuantityTrend-1': 'Get daily sales-quantity trend for a product',
 'com.alibaba.fenxiao.crossborder:product.analyze.getThirtyDayMedianPriceTrend-1': 'Get 30-day median-price trend for a product',
 'com.alibaba.fenxiao.crossborder:image.elements.remove-1': 'Smart image object removal',
 'com.alibaba.fenxiao.crossborder:image.elements.enlarge-1': 'HD image upscaling',
 'com.alibaba.fenxiao.crossborder:text.generate.description-1': 'Product description generation',
 'com.alibaba.fenxiao.crossborder:image.elements.translate-1': 'Image translation',
 'com.alibaba.fenxiao.crossborder:text.generate.title-1': 'Product title generation',
 'com.alibaba.fenxiao.crossborder:image.elements.matting-1': 'Smart image cutout (matting)',
 'com.alibaba.fenxiao.crossborder:image.elements.recognition-1': 'Image element recognition',
 'com.alibaba.fenxiao.crossborder:product.text.translate-1': 'Product text translation',
 'com.alibaba.image:image.elements.cut-1': 'Image cropping',
}

SOLUTION_EN = 'XunYuanTong Sourcing Solution for Cross-Border Key Accounts'
ROLE_EN = 'Cross-border key account'
IDENTITY_EN = 'Procurement service provider'
NOTE_DOC_URL = 'https://alidocs.dingtalk.com/i/nodes/ZX6GRezwJlzeYoPLFbr4Na6DWdqbropQ'
NOTE_PERM = 'Permission requests: contact the business team and join the DingTalk group to obtain access to the documentation.'

SDK_TEXT = [
 'Welcome to the Alibaba Open Platform SDK. The SDK already encapsulates request signing and signature verification, so you can call the APIs directly through it.',
 'When calling an API, first work out the class name that corresponds to the interface. In the SDK, capitalise the first letter of each word, remove the "." separators, and append Param (request class) or Result (response class). For example, the interface alibaba.product.add corresponds to AlibabaProductAddParam (request class) and AlibabaProductAddResult (response class).',
 'Then complete the API call with the following code:',
]
SDK_CODE = '''ApiExecutor apiExecutor = new ApiExecutor("{your app key}", "{your app secret}");
AlibabaProductAddParam param = new AlibabaProductAddParam();
//TODO set the param fields
apiExecutor.execute(param, "{your access token}");'''
SDK_VERSIONS = ['Java 1.6 or later', 'PHP 5.6 or later', '.NET requires .NET Framework 4.0 or later']

# ---------------------------------------------------------------- coverage check
missing = []
for fam in LST['apis']:
    if fam['categoryFamilyName'] not in CATS: missing.append('CAT ' + fam['categoryFamilyName'])
    for m in fam['modules']:
        if m['fullName'] not in API: missing.append('API ' + m['fullName'])
for fam in LST['messages']:
    if fam['categoryFamilyName'] not in CATS: missing.append('MCAT ' + fam['categoryFamilyName'])
    for m in fam['modules']:
        if m['typeName'] not in MSG: missing.append('MSG ' + m['typeName'])
api_by_key = {}
for fam in LST['apis']:
    for m in fam['modules']:
        api_by_key[f"{m['namespace']}:{m['name']}-{m['version']}"] = m
for ch in SOL['changeLog']:
    for k in ('addedApis', 'removedApis'):
        for x in ch[k]:
            if x['key'] not in api_by_key and x['key'] not in EXTRA: missing.append('CHG ' + x['key'])
    for k in ('addedMessages', 'removedMessages'):
        for x in ch[k]:
            if x['key'] not in MSG: missing.append('CHGMSG ' + x['key'])
if missing:
    print('MISSING TRANSLATIONS:'); print('\n'.join(missing)); sys.exit(1)

# ---------------------------------------------------------------- derived data
n_api = sum(len(f['modules']) for f in LST['apis'])
n_msg = sum(len(f['modules']) for f in LST['messages'])
n_cat = len(LST['apis'])
n_chg = len(SOL['changeLog'])
updated = datetime.datetime.utcfromtimestamp(int(SOL['gmtModifyed']) / 1000).strftime('%Y-%m-%d')
sdk = json.loads([e for e in SOL['solutionDescList'] if e['subCategory'] == 'sdk'][0]['content'])
PAGE_URL = 'https://open.1688.com/solution/detail?key=1703167397970'

def api_url(m): return f"https://open.1688.com/api/apidocdetail.htm?id={m['namespace']}:{m['name']}-{m['version']}"
def msg_url(t): return f"https://open.1688.com/doc/topicDetail.htm?id={t}"
def slug(s): return ''.join(ch if ch.isalnum() else '-' for ch in s.lower()).strip('-')
def chg_name(key):
    if key in api_by_key: return API[api_by_key[key]['fullName']][0]
    return EXTRA[key]

# ---------------------------------------------------------------- markdown
md = []
md.append(f'# {SOLUTION_EN}\n')
md.append(f'*English translation of the 1688 Open Platform solution page* [{PAGE_URL}]({PAGE_URL})  ')
md.append(f'Original name: {SOL["solutionName"]} · Solution key {SOL["bizKey"]} · Version {SOL["version"]} · Role: {ROLE_EN} ({SOL["roleName"]}) · Identity: {IDENTITY_EN} ({SOL["identityName"]}) · Last updated {updated}\n')
md.append('## Overview\n')
md.append(f'- Documentation: {NOTE_DOC_URL}')
md.append(f'- {NOTE_PERM}\n')
md.append(f'**Contents:** {n_api} APIs in {n_cat} categories · {n_msg} message topics · {n_chg} change-log entries · SDK downloads\n')
md.append('## API list\n')
md.append('API identifiers, parameter names and error codes are kept exactly as in the original. Each API name links to its detail page (request/response parameters) on open.1688.com.\n')
for fam in LST['apis']:
    md.append(f"### {CATS[fam['categoryFamilyName']]} ({fam['categoryFamilyName']})\n")
    md.append('| API | Original name | Description |')
    md.append('|---|---|---|')
    for m in fam['modules']:
        en, desc = API[m['fullName']]
        md.append(f"| [{en}]({api_url(m)})<br>`{m['namespace']}:{m['name']}:{m['version']}` | {m['displayName'].strip()} | {desc} |")
    md.append('')
md.append('## Message topics\n')
for fam in LST['messages']:
    md.append(f"### {CATS[fam['categoryFamilyName']]} ({fam['categoryFamilyName']})\n")
    md.append('| Topic | Description | Original |')
    md.append('|---|---|---|')
    for m in fam['modules']:
        md.append(f"| [`{m['typeName']}`]({msg_url(m['typeName'])}) | {MSG[m['typeName']]} | {m['description']} |")
    md.append('')
md.append('## Change log\n')
for ch in SOL['changeLog']:
    md.append(f"### {ch['title']}\n")
    for x in ch['addedApis']: md.append(f"- Added API: {chg_name(x['key'])} (`{x['key']}`)")
    for x in ch['addedMessages']: md.append(f"- Added message: {MSG[x['key']]} (`{x['key']}`)")
    for x in ch['removedApis']: md.append(f"- Removed API: {chg_name(x['key'])} (`{x['key']}`)")
    for x in ch['removedMessages']: md.append(f"- Removed message: {MSG[x['key']]} (`{x['key']}`)")
    md.append('')
md.append('## SDK\n')
for p in SDK_TEXT: md.append(p + '\n')
md.append('```java\n' + SDK_CODE + '\n```\n')
md.append('SDK language versions:\n')
for v in SDK_VERSIONS: md.append(f'- {v}')
md.append('\nDownloads (SDK guide: http://open.1688.com/doc/apiSdk.htm):\n')
for f in sdk['sdkFiles']: md.append(f"- {f['platform']}: {f['downloadUrl']}")
md.append('')
open('/Users/natawatd/Desktop/temp/1688-xunyuantong-solution-api-docs-en.md', 'w', encoding='utf-8').write('\n'.join(md))

# ---------------------------------------------------------------- html
e = html.escape
H = []
H.append(f'''<title>XunYuanTong Solution API Reference</title>
<link rel="preconnect" href="https://fonts.googleapis.com">
<link rel="stylesheet" href="https://fonts.googleapis.com/css2?family=IBM+Plex+Sans:wght@400;500;600&family=IBM+Plex+Sans+Condensed:wght@500;600&family=IBM+Plex+Mono:wght@400;500&display=swap">
<style>
:root{{
  --bg:#F5F6F8; --surface:#FFFFFF; --ink:#1B1F26; --muted:#5F6875; --line:#DCE0E6; --line-strong:#C3C9D2;
  --accent:#D85A00; --accent-ink:#FFFFFF; --accent-soft:#FFF1E6; --code-bg:#EEF0F3;
  --added:#1E7A4A; --added-soft:#E4F4EA; --removed:#B4261E; --removed-soft:#FBE7E5;
  --focus:#D85A00;
}}
@media (prefers-color-scheme: dark){{
  :root:not([data-theme="light"]){{
    --bg:#121519; --surface:#1A1E24; --ink:#E7E9ED; --muted:#98A2AF; --line:#2A303A; --line-strong:#3A424E;
    --accent:#FF8B45; --accent-ink:#1A1208; --accent-soft:#2A1E14; --code-bg:#22272F;
    --added:#63C48F; --added-soft:#17302A; --removed:#F08C84; --removed-soft:#3A1F1E;
    --focus:#FF8B45;
  }}
}}
:root[data-theme="dark"]{{
  --bg:#121519; --surface:#1A1E24; --ink:#E7E9ED; --muted:#98A2AF; --line:#2A303A; --line-strong:#3A424E;
  --accent:#FF8B45; --accent-ink:#1A1208; --accent-soft:#2A1E14; --code-bg:#22272F;
  --added:#63C48F; --added-soft:#17302A; --removed:#F08C84; --removed-soft:#3A1F1E;
  --focus:#FF8B45;
}}
*{{box-sizing:border-box}}
html{{scroll-behavior:smooth}}
@media (prefers-reduced-motion: reduce){{ html{{scroll-behavior:auto}} }}
body{{margin:0;background:var(--bg);color:var(--ink);font-family:"IBM Plex Sans","Helvetica Neue",Arial,"PingFang SC","Microsoft YaHei",sans-serif;font-size:15px;line-height:1.55;-webkit-font-smoothing:antialiased}}
a{{color:var(--accent);text-decoration:none}}
a:hover{{text-decoration:underline}}
a:focus-visible,input:focus-visible,button:focus-visible{{outline:2px solid var(--focus);outline-offset:2px}}
code,.mono{{font-family:"IBM Plex Mono",ui-monospace,SFMono-Regular,Menlo,Consolas,monospace;font-size:.86em}}
code{{background:var(--code-bg);padding:.1em .4em;border-radius:3px;word-break:break-all}}
h1,h2,h3{{font-family:"IBM Plex Sans Condensed","IBM Plex Sans","Helvetica Neue",Arial,"PingFang SC",sans-serif;font-weight:600;line-height:1.15;margin:0;text-wrap:balance}}
h1{{font-size:2.2rem;letter-spacing:-.01em}}
h2{{font-size:1.5rem;padding-top:8px}}
h3{{font-size:1.15rem}}
.cn{{color:var(--muted);font-weight:400}}
.eyebrow{{font-size:.74rem;letter-spacing:.09em;text-transform:uppercase;color:var(--muted);font-weight:500}}

.shell{{display:grid;grid-template-columns:236px minmax(0,1fr);gap:44px;max-width:1260px;margin:0 auto;padding:36px 28px 80px}}
nav.side{{position:sticky;top:20px;align-self:start;max-height:calc(100vh - 40px);overflow:auto;font-size:.88rem;padding-right:8px}}
nav.side .grp{{margin:0 0 18px}}
nav.side .grp .eyebrow{{display:block;margin-bottom:6px}}
nav.side a{{display:flex;justify-content:space-between;gap:10px;color:var(--ink);padding:4px 8px;border-left:2px solid transparent;border-radius:0 3px 3px 0}}
nav.side a:hover{{background:var(--accent-soft);text-decoration:none}}
nav.side a.on{{border-left-color:var(--accent);color:var(--accent)}}
nav.side a .n{{color:var(--muted);font-variant-numeric:tabular-nums}}

header.top{{display:flex;flex-direction:column;gap:14px;padding-bottom:26px;border-bottom:1px solid var(--line-strong);margin-bottom:28px}}
.meta{{display:flex;flex-wrap:wrap;gap:8px 22px;font-size:.88rem;color:var(--muted)}}
.meta b{{color:var(--ink);font-weight:500}}
.stats{{display:flex;flex-wrap:wrap;gap:6px 26px;font-size:.92rem;font-variant-numeric:tabular-nums}}
.stats span b{{font-size:1.25rem;font-family:"IBM Plex Sans Condensed",sans-serif;font-weight:600;margin-right:4px}}
.note{{background:var(--accent-soft);border-left:3px solid var(--accent);padding:12px 16px;font-size:.92rem;border-radius:0 4px 4px 0}}
.note p{{margin:0}} .note p+p{{margin-top:6px}}

.search{{display:flex;align-items:center;gap:12px;position:sticky;top:0;z-index:2;background:var(--bg);padding:10px 0 12px;margin-bottom:6px}}
.search input{{flex:1;max-width:520px;padding:9px 12px;border:1px solid var(--line-strong);border-radius:4px;background:var(--surface);color:var(--ink);font:inherit;font-size:.95rem}}
.search .count{{font-size:.86rem;color:var(--muted);font-variant-numeric:tabular-nums}}

section.cat{{margin:0 0 34px}}
section.cat h2,section.msg h2{{display:flex;align-items:baseline;gap:12px;flex-wrap:wrap;margin-bottom:12px}}
section.cat h2 .n,section.msg h2 .n{{font-size:.8rem;color:var(--muted);font-family:"IBM Plex Sans",sans-serif;font-weight:500}}
.tbl{{overflow-x:auto;background:var(--surface);border:1px solid var(--line);border-radius:5px}}
table{{border-collapse:collapse;width:100%;font-size:.92rem}}
th{{text-align:left;font-size:.72rem;letter-spacing:.08em;text-transform:uppercase;color:var(--muted);font-weight:500;padding:9px 14px;border-bottom:1px solid var(--line);background:var(--surface)}}
td{{padding:12px 14px;border-bottom:1px solid var(--line);vertical-align:top}}
tr:last-child td{{border-bottom:0}}
td.api{{width:38%;min-width:260px}}
td.api .name{{font-weight:500;display:block}}
td.api .orig{{display:block;color:var(--muted);font-size:.86rem;margin:2px 0 6px}}
td.api code{{display:inline-block;font-size:.78rem;color:var(--ink);background:var(--code-bg)}}
td.desc{{color:var(--ink)}}
td.desc .orig{{display:block;color:var(--muted);font-size:.84rem;margin-top:6px}}
tr.hide{{display:none}}
.ns{{display:inline-block;font-size:.7rem;letter-spacing:.04em;color:var(--muted);border:1px solid var(--line);border-radius:3px;padding:0 6px;margin-left:6px;vertical-align:middle;font-family:"IBM Plex Mono",monospace}}
.empty{{padding:14px;color:var(--muted);font-size:.9rem;display:none}}
section.cat.none .tbl{{display:none}} section.cat.none .empty{{display:block}}

.divider{{border:0;border-top:1px solid var(--line-strong);margin:44px 0 30px}}
.log{{display:grid;grid-template-columns:150px minmax(0,1fr);gap:0 22px}}
.log .when{{font-family:"IBM Plex Mono",monospace;font-size:.82rem;color:var(--muted);padding:14px 0;border-top:1px solid var(--line);font-variant-numeric:tabular-nums}}
.log .what{{padding:14px 0;border-top:1px solid var(--line)}}
.log ul{{margin:0;padding:0;list-style:none;display:flex;flex-direction:column;gap:5px}}
.log li{{display:flex;gap:10px;align-items:baseline;font-size:.9rem}}
.tag{{flex:none;font-size:.68rem;letter-spacing:.06em;text-transform:uppercase;font-weight:500;padding:1px 7px;border-radius:3px;min-width:88px;text-align:center}}
.tag.add{{color:var(--added);background:var(--added-soft)}}
.tag.rm{{color:var(--removed);background:var(--removed-soft)}}
.log li code{{font-size:.76rem;color:var(--muted);background:transparent;padding:0}}
.sdk p{{max-width:68ch;margin:0 0 12px}}
pre{{background:var(--code-bg);padding:14px 16px;border-radius:5px;overflow-x:auto;font-size:.84rem;line-height:1.5;margin:0 0 16px}}
pre code{{background:transparent;padding:0;font-size:inherit}}
.sdk ul{{margin:0 0 14px 20px;padding:0}}
.dl{{display:flex;flex-wrap:wrap;gap:8px}}
.dl a{{border:1px solid var(--line-strong);border-radius:4px;padding:6px 12px;font-size:.88rem;color:var(--ink);background:var(--surface)}}
.dl a:hover{{border-color:var(--accent);color:var(--accent);text-decoration:none}}
footer{{margin-top:50px;padding-top:18px;border-top:1px solid var(--line);font-size:.84rem;color:var(--muted);max-width:70ch}}
@media (max-width:900px){{
  .shell{{grid-template-columns:1fr;gap:20px;padding:22px 16px 60px}}
  nav.side{{position:static;max-height:none;display:flex;gap:4px 10px;flex-wrap:wrap;border-bottom:1px solid var(--line);padding-bottom:12px}}
  nav.side .grp{{display:contents}} nav.side .grp .eyebrow{{display:none}}
  nav.side a{{border:1px solid var(--line);border-radius:999px;padding:3px 10px}}
  nav.side a.on{{border-color:var(--accent)}}
  h1{{font-size:1.7rem}}
  .log{{grid-template-columns:1fr}} .log .when{{border-top:0;padding-bottom:0}} .log .what{{padding-top:6px}}
  td.api{{min-width:220px}}
}}
</style>
<div class="shell">
<nav class="side" aria-label="Sections">
  <div class="grp"><span class="eyebrow">APIs</span>''')
for fam in LST['apis']:
    H.append(f'    <a href="#cat-{slug(CATS[fam["categoryFamilyName"]])}">{e(CATS[fam["categoryFamilyName"]])}<span class="n">{len(fam["modules"])}</span></a>')
H.append('  </div><div class="grp"><span class="eyebrow">Messages</span>')
for fam in LST['messages']:
    H.append(f'    <a href="#msg-{slug(CATS[fam["categoryFamilyName"]])}">{e(CATS[fam["categoryFamilyName"]])}<span class="n">{len(fam["modules"])}</span></a>')
H.append('  </div><div class="grp"><span class="eyebrow">More</span><a href="#changelog">Change log</a><a href="#sdk">SDK</a></div>')
H.append('</nav>\n<main>')
H.append(f'''<header class="top">
  <span class="eyebrow">1688 Open Platform · Solution {e(SOL["bizKey"])} · English translation</span>
  <h1>{e(SOLUTION_EN)}</h1>
  <div class="cn">{e(SOL["solutionName"])} · <a href="{PAGE_URL}" target="_blank" rel="noopener">original page</a></div>
  <div class="meta"><span>Role <b>{ROLE_EN}</b> <span class="cn">{e(SOL["roleName"])}</span></span><span>Identity <b>{IDENTITY_EN}</b> <span class="cn">{e(SOL["identityName"])}</span></span><span>Version <b>{e(SOL["version"])}</b></span><span>Last updated <b>{updated}</b></span></div>
  <div class="stats"><span><b>{n_api}</b> APIs</span><span><b>{n_cat}</b> categories</span><span><b>{n_msg}</b> message topics</span><span><b>{n_chg}</b> change-log entries</span></div>
  <div class="note"><p><b>Documentation:</b> <a href="{NOTE_DOC_URL}" target="_blank" rel="noopener">{NOTE_DOC_URL}</a></p><p>{e(NOTE_PERM)}</p></div>
</header>
<div class="search"><input id="q" type="search" placeholder="Filter APIs and messages by name, ID or description (English or 中文)" aria-label="Filter"><span class="count" id="count">{n_api} APIs · {n_msg} messages</span></div>
<p class="cn" style="margin:0 0 22px;font-size:.9rem;max-width:72ch">API identifiers, parameter names and error codes are kept exactly as in the original. Each API name links to its detail page on open.1688.com, where the request and response parameters are documented.</p>''')

for fam in LST['apis']:
    cn = fam['categoryFamilyName']; en = CATS[cn]
    H.append(f'<section class="cat" id="cat-{slug(en)}"><h2>{e(en)} <span class="cn">{e(cn)}</span> <span class="n">{len(fam["modules"])} APIs</span></h2><div class="tbl"><table><thead><tr><th>API</th><th>Description</th></tr></thead><tbody>')
    for m in fam['modules']:
        en_name, desc = API[m['fullName']]
        H.append(f'<tr class="row"><td class="api"><a class="name" href="{api_url(m)}" target="_blank" rel="noopener">{e(en_name)}</a><span class="orig">{e(m["displayName"].strip())}</span><code>{e(m["name"])}:{e(m["version"])}</code><span class="ns">{e(m["namespace"])}</span></td><td class="desc">{e(desc)}<span class="orig">{e(" ".join(m["description"].split()))}</span></td></tr>')
    H.append('</tbody></table></div><div class="empty">No APIs in this category match the filter.</div></section>')

H.append('<hr class="divider"><h2 id="messages">Message topics <span class="cn">消息</span></h2><p class="cn" style="font-size:.9rem;margin:8px 0 20px">Push-message topics the solution subscribes to. Topic names link to the topic detail page.</p>')
for fam in LST['messages']:
    cn = fam['categoryFamilyName']; en = CATS[cn]
    H.append(f'<section class="cat msg" id="msg-{slug(en)}"><h2 style="font-size:1.2rem">{e(en)} <span class="cn">{e(cn)}</span> <span class="n">{len(fam["modules"])} topics</span></h2><div class="tbl"><table><thead><tr><th>Topic</th><th>Description</th></tr></thead><tbody>')
    for m in fam['modules']:
        H.append(f'<tr class="row"><td class="api"><a class="name mono" href="{msg_url(m["typeName"])}" target="_blank" rel="noopener">{e(m["typeName"])}</a></td><td class="desc">{e(MSG[m["typeName"]])}<span class="orig">{e(m["description"])}</span></td></tr>')
    H.append('</tbody></table></div><div class="empty">No topics in this group match the filter.</div></section>')

H.append('<hr class="divider"><h2 id="changelog">Change log <span class="cn">变更记录</span></h2><p class="cn" style="font-size:.9rem;margin:8px 0 18px">APIs and message topics added to or removed from this solution, newest first.</p><div class="log">')
for ch in SOL['changeLog']:
    items = []
    for x in ch['addedApis']: items.append(f'<li><span class="tag add">Added API</span><span>{e(chg_name(x["key"]))} <code>{e(x["key"])}</code></span></li>')
    for x in ch['addedMessages']: items.append(f'<li><span class="tag add">Added msg</span><span>{e(MSG[x["key"]])} <code>{e(x["key"])}</code></span></li>')
    for x in ch['removedApis']: items.append(f'<li><span class="tag rm">Removed API</span><span>{e(chg_name(x["key"]))} <code>{e(x["key"])}</code></span></li>')
    for x in ch['removedMessages']: items.append(f'<li><span class="tag rm">Removed msg</span><span>{e(MSG[x["key"]])} <code>{e(x["key"])}</code></span></li>')
    H.append(f'<div class="when">{e(ch["title"])}</div><div class="what"><ul>{"".join(items)}</ul></div>')
H.append('</div>')

H.append('<hr class="divider"><section class="sdk" id="sdk"><h2>SDK <span class="cn">SDK下载</span></h2><div style="height:14px"></div>')
for p in SDK_TEXT: H.append(f'<p>{e(p)}</p>')
H.append(f'<pre><code>{e(SDK_CODE)}</code></pre><p><b>SDK language versions</b></p><ul>')
for v in SDK_VERSIONS: H.append(f'<li>{e(v)}</li>')
H.append('</ul><p><b>Downloads</b> <span class="cn">(SDK guide: <a href="http://open.1688.com/doc/apiSdk.htm" target="_blank" rel="noopener">open.1688.com/doc/apiSdk.htm</a>)</span></p><div class="dl">')
for f in sdk['sdkFiles']:
    H.append(f'<a href="{e(f["downloadUrl"])}" target="_blank" rel="noopener">{e(f["platform"].upper() if f["platform"]!="net" else ".NET")} SDK · {e(f["fileName"])}</a>')
H.append('</div></section>')
H.append(f'<footer>Translated from the Chinese original at open.1688.com (solution {e(SOL["bizKey"])}, data as of {updated}). Product names such as XunYuanTong (寻源通), ShengYiTong (生意通), Cheng-e-She (诚e赊), Kuajingbao (跨境宝) and Newton Cloud (牛顿云) are transliterated; the original Chinese is shown under each entry.</footer>')
H.append('''</main></div>
<script>
(function(){
  var q=document.getElementById('q'),count=document.getElementById('count');
  var rows=[].slice.call(document.querySelectorAll('tr.row'));
  var secs=[].slice.call(document.querySelectorAll('section.cat'));
  var nApi=rows.filter(function(r){return !r.closest('section.msg')}).length, nMsg=rows.length-nApi;
  function apply(){
    var t=q.value.trim().toLowerCase(), a=0, m=0;
    rows.forEach(function(r){
      var hit=!t||r.textContent.toLowerCase().indexOf(t)>-1;
      r.classList.toggle('hide',!hit);
      if(hit){ if(r.closest('section.msg')) m++; else a++; }
    });
    secs.forEach(function(s){ s.classList.toggle('none',!s.querySelector('tr.row:not(.hide)')); });
    count.textContent=t?(a+' of '+nApi+' APIs · '+m+' of '+nMsg+' messages'):(nApi+' APIs · '+nMsg+' messages');
  }
  q.addEventListener('input',apply);
  var links=[].slice.call(document.querySelectorAll('nav.side a[href^="#"]'));
  var targets=links.map(function(l){return document.getElementById(l.getAttribute('href').slice(1))}).filter(Boolean);
  if('IntersectionObserver' in window){
    var io=new IntersectionObserver(function(es){
      es.forEach(function(en){ if(en.isIntersecting){ links.forEach(function(l){l.classList.toggle('on',l.getAttribute('href')==='#'+en.target.id)}); } });
    },{rootMargin:'-10% 0px -80% 0px'});
    targets.forEach(function(t){io.observe(t)});
  }
})();
</script>''')
open('xunyuantong-api-reference.html', 'w', encoding='utf-8').write('\n'.join(H))
print('OK', n_api, 'APIs', n_msg, 'messages', n_cat, 'categories', n_chg, 'changelog; updated', updated)
