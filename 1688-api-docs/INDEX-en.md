# 1688 API docs (English)

Solution 1703167397970 · pulled from open.1688.com on 2026-09-03

## Membership & Accounts

- [Batch-add sub-account authorizations](en/system.oauth2.subaccount.auth.add-1.md) `system.oauth2:subaccount.auth.add:1` — 批量添加子账号授权
- [Batch-cancel sub-account authorizations](en/system.oauth2.subaccount.auth.cancel-1.md) `system.oauth2:subaccount.auth.cancel:1` — 批量取消子账号授权
- [Batch-query sub-account authorizations](en/system.oauth2.subaccount.auth.list-1.md) `system.oauth2:subaccount.auth.list:1` — 批量查询子账号授权
- [1688 member registration](en/com.alibaba.fenxiao.crossborder.account.user.register-1.md) `com.alibaba.fenxiao.crossborder:account.user.register:1` — 1688会员注册
- [Get basic info of a non-authorized user (cross-border)](en/com.alibaba.account.alibaba.account.agent.crossBasic-1.md) `com.alibaba.account:alibaba.account.agent.crossBasic:1` — 跨境场景获取非授权用户的基本信息
- [Get basic info of the authorized user](en/com.alibaba.account.alibaba.account.basic-1.md) `com.alibaba.account:alibaba.account.basic:1` — 获取授权用户的基本信息
- [Query sub-account info](en/cn.alibaba.open.querySubAccount-1.md) `cn.alibaba.open:querySubAccount:1` — 查询子账号信息

## Tools

- [Pull products from a product pool](en/com.alibaba.fenxiao.crossborder.pool.product.pull-1.md) `com.alibaba.fenxiao.crossborder:pool.product.pull:1` — 拉取商品池中商品数据
- [Query product count in a product pool](en/com.alibaba.fenxiao.crossborder.pool.product.total-1.md) `com.alibaba.fenxiao.crossborder:pool.product.total:1` — 查询商品池中商品总数
- [Encrypt a user loginId into an OpenUID](en/com.alibaba.account.loginid.openuid.encrypt-1.md) `com.alibaba.account:loginid.openuid.encrypt:1` — 用户loginId加密转换为Openuid接口
- [Get a link that opens a WangWang chat](en/com.alibaba.account.account.wangwangUrl.get-1.md) `com.alibaba.account:account.wangwangUrl.get:1` — 获取唤起旺旺聊天的链接
- [Decrypt an OpenUID into a WangWang nickname (only for launching WangWang)](en/com.alibaba.account.wangwangnick.openuid.decrypt-1.md) `com.alibaba.account:wangwangnick.openuid.decrypt:1` — Openuid转换解密为旺旺昵称接口（仅可使用于用户唤起旺旺）
- [Add buyer-seller distribution relationship](en/com.alibaba.fenxiao.alibaba.fenxiao.relationadd-1.md) `com.alibaba.fenxiao:alibaba.fenxiao.relationadd:1` — 买卖家分销关系添加

## Market Insights

- [Trending product search keywords](en/com.alibaba.fenxiao.crossborder.product.search.topKeyword-1.md) `com.alibaba.fenxiao.crossborder:product.search.topKeyword:1` — 商品热搜词
- [Query ranking lists](en/com.alibaba.fenxiao.crossborder.product.topList.query-1.md) `com.alibaba.fenxiao.crossborder:product.topList.query:1` — 查询榜单列表
- [Get daily sales-quantity trend for a product (new)](en/com.alibaba.fenxiao.crossborder.product.analyze.getPerdaySellQuantityTrendNew-1.md) `com.alibaba.fenxiao.crossborder:product.analyze.getPerdaySellQuantityTrendNew:1` — 获取商品每日销售数量趋势（新）
- [Get 30-day median-price trend for a product (new)](en/com.alibaba.fenxiao.crossborder.product.analyze.getThirtyDayMedianPriceTrendNew-1.md) `com.alibaba.fenxiao.crossborder:product.analyze.getThirtyDayMedianPriceTrendNew:1` — 获取商品30天价格中位数指标趋势（新）
- [Get 90-day repurchase-rate trend for a product (new)](en/com.alibaba.fenxiao.crossborder.product.analyze.getRepurchaseRateTrendNew-1.md) `com.alibaba.fenxiao.crossborder:product.analyze.getRepurchaseRateTrendNew:1` — 获取商品90天复购率指标趋势(新） 

## Categories

- [Query multilingual category by category ID](en/com.alibaba.fenxiao.crossborder.category.translation.getById-1.md) `com.alibaba.fenxiao.crossborder:category.translation.getById:1` — 根据类目ID查询多语言类目
- [Query multilingual category by category name](en/com.alibaba.fenxiao.crossborder.category.translation.getByKeyword-1.md) `com.alibaba.fenxiao.crossborder:category.translation.getByKeyword:1` — 根据类目名称查询多语言类目
- [Get leaf-category attributes](en/com.alibaba.product.alibaba.category.attribute.get-1.md) `com.alibaba.product:alibaba.category.attribute.get:1` — 获取叶子类目属性

## Products

- [Multilingual keyword search](en/com.alibaba.fenxiao.crossborder.product.search.keywordQuery-1.md) `com.alibaba.fenxiao.crossborder:product.search.keywordQuery:1` — 多语言关键词搜索
- [Multilingual image search](en/com.alibaba.fenxiao.crossborder.product.search.imageQuery-1.md) `com.alibaba.fenxiao.crossborder:product.search.imageQuery:1` — 多语言图搜
- [Multilingual in-store product search](en/com.alibaba.fenxiao.crossborder.product.search.querySellerOfferList-1.md) `com.alibaba.fenxiao.crossborder:product.search.querySellerOfferList:1` — 多语言商品店搜
- [Multilingual product detail](en/com.alibaba.fenxiao.crossborder.product.search.queryProductDetail-1.md) `com.alibaba.fenxiao.crossborder:product.search.queryProductDetail:1` — 多语言商详
- [Product recommendations](en/com.alibaba.fenxiao.crossborder.product.search.offerRecommend-1.md) `com.alibaba.fenxiao.crossborder:product.search.offerRecommend:1` — 商品推荐
- [Multilingual search navigation](en/com.alibaba.fenxiao.crossborder.product.search.keywordSNQuery-1.md) `com.alibaba.fenxiao.crossborder:product.search.keywordSNQuery:1` — 多语言搜索导航
- [Related product recommendations](en/com.alibaba.fenxiao.crossborder.product.related.recommend-1.md) `com.alibaba.fenxiao.crossborder:product.related.recommend:1` — 相关性商品推荐
- [Upload an image to get an imageId](en/com.alibaba.fenxiao.crossborder.product.image.upload-1.md) `com.alibaba.fenxiao.crossborder:product.image.upload:1` — 上传图片获取imageId
- [Claim the optimal coupon for a product](en/com.alibaba.marketing.coupon.optimal.claim-1.md) `com.alibaba.marketing:coupon.optimal.claim:1` — 通过商品领取最优化的优惠券
- [Recommend products by keyword (cross-border)](en/com.alibaba.product.alibaba.product.suggest.crossBorder-1.md) `com.alibaba.product:alibaba.product.suggest.crossBorder:1` — 跨境场景根据关键字推荐商品
- [Get simple product info from a previously purchased supplier](en/com.alibaba.product.alibaba.product.simple.get-1.md) `com.alibaba.product:alibaba.product.simple.get:1` — 获取已购买过商家的商品简单信息
- [Follow a product (cross-border)](en/com.alibaba.fenxiao.crossborder.product.kjdistribute.addRelation-1.md) `com.alibaba.fenxiao.crossborder:product.kjdistribute.addRelation:1` — 增加跨境关注商品
- [Unfollow a product (cross-border)](en/com.alibaba.fenxiao.crossborder.product.kjdistribute.removeRelation-1.md) `com.alibaba.fenxiao.crossborder:product.kjdistribute.removeRelation:1` — 取消跨境关注商品
- [Get product selling points for listing](en/com.alibaba.fenxiao.crossborder.product.distribute.getDistributeInfo-1.md) `com.alibaba.fenxiao.crossborder:product.distribute.getDistributeInfo:1` — 获取商品铺货卖点
- [Query leaf-category attribute and value mappings](en/com.alibaba.fenxiao.crossborder.product.category.getAttrById-1.md) `com.alibaba.fenxiao.crossborder:product.category.getAttrById:1` — 叶子类目属性属性值映射查询

## Product Selection & Listing

- [Get product details (cross-border)](en/com.alibaba.product.alibaba.cross.productInfo-1.md) `com.alibaba.product:alibaba.cross.productInfo:1` — 跨境场景获取商品详情
- [Add products to the listing list (cross-border)](en/com.alibaba.product.push.alibaba.cross.syncProductListPushed-1.md) `com.alibaba.product.push:alibaba.cross.syncProductListPushed:1` — 跨境场景下将商品加入铺货列表
- [Sync listing results](en/com.alibaba.product.push.alibaba.product.push.syncPushProductResult-1.md) `com.alibaba.product.push:alibaba.product.push.syncPushProductResult:1` — 同步铺货结果
- [Get product list (cross-border)](en/com.alibaba.product.alibaba.cross.productList-1.md) `com.alibaba.product:alibaba.cross.productList:1` — 跨境场景获取商品列表
- [Unfollow a product](en/com.alibaba.product.alibaba.product.unfollow.crossborder-1.md) `com.alibaba.product:alibaba.product.unfollow.crossborder:1` — 解除关注商品
- [Query category by category ID](en/com.alibaba.product.alibaba.category.get-1.md) `com.alibaba.product:alibaba.category.get:1` — 根据类目Id查询类目
- [Follow a product](en/com.alibaba.product.alibaba.product.follow.crossborder-1.md) `com.alibaba.product:alibaba.product.follow.crossborder:1` — 关注商品
- [[Relationship] Distributor: query supplier list](en/cn.alibaba.open.alibaba.relation.querySuppliers-1.md) `cn.alibaba.open:alibaba.relation.querySuppliers:1` — 【关系】分销商-查询供应商列表

## Orders

- [Create cross-border order](en/com.alibaba.trade.alibaba.trade.createCrossOrder-1.md) `com.alibaba.trade:alibaba.trade.createCrossOrder:1` — 跨境订单创建
- [View order details (buyer view)](en/com.alibaba.trade.alibaba.trade.get.buyerView-1.md) `com.alibaba.trade:alibaba.trade.get.buyerView:1` — 订单详情查看(买家视角)
- [Get order logistics info (buyer view)](en/com.alibaba.logistics.alibaba.trade.getLogisticsInfos.buyerView-1.md) `com.alibaba.logistics:alibaba.trade.getLogisticsInfos.buyerView:1` — 获取交易订单的物流信息(买家视角)
- [Get order logistics tracking info (buyer view)](en/com.alibaba.logistics.alibaba.trade.getLogisticsTraceInfo.buyerView-1.md) `com.alibaba.logistics:alibaba.trade.getLogisticsTraceInfo.buyerView:1` — 获取交易订单的物流跟踪信息(买家视角)
- [Cancel transaction](en/com.alibaba.trade.alibaba.trade.cancel-1.md) `com.alibaba.trade:alibaba.trade.cancel:1` — 取消交易
- [View order list (buyer view)](en/com.alibaba.trade.alibaba.trade.getBuyerOrderList-1.md) `com.alibaba.trade:alibaba.trade.getBuyerOrderList:1` — 订单列表查看(买家视角)
- [Preview data before creating an order](en/com.alibaba.trade.alibaba.createOrder.preview-1.md) `com.alibaba.trade:alibaba.createOrder.preview:1` — 创建订单前预览数据接口
- [Update order memo](en/com.alibaba.trade.alibaba.order.memoAdd-1.md) `com.alibaba.trade:alibaba.order.memoAdd:1` — 修改订单备忘
- [Get sub-account list](en/com.alibaba.account.alibaba.subAccount.list-1.md) `com.alibaba.account:alibaba.subAccount.list:1` — 获取子账号列表
- [Buyer adds an order message](en/com.alibaba.trade.alibaba.trade.addFeedback-1.md) `com.alibaba.trade:alibaba.trade.addFeedback:1` — 买家补充订单留言接口
- [Buyer confirms receipt](en/com.alibaba.trade.trade.receivegoods.confirm-1.md) `com.alibaba.trade:trade.receivegoods.confirm:1` — 买家确认收货
- [Buyer deletes a closed order](en/com.alibaba.trade.trade.order.buyerdelete-1.md) `com.alibaba.trade:trade.order.buyerdelete:1` — 买家删除已关闭的订单
- [Buyer requests a shipping-address change](en/com.alibaba.trade.order.receiveAddress.buyerUpdate-1.md) `com.alibaba.trade:order.receiveAddress.buyerUpdate:1` — 买家申请修改收货地址

## Payment

- [Batch-get payment links for orders](en/com.alibaba.trade.alibaba.alipay.url.get-1.md) `com.alibaba.trade:alibaba.alipay.url.get:1` — 批量获取订单的支付链接
- [Get payment link for Cross-Border Pay (Kuajingbao)](en/com.alibaba.trade.alibaba.crossBorderPay.url.get-1.md) `com.alibaba.trade:alibaba.crossBorderPay.url.get:1` — 获取使用跨境宝支付的支付链接
- [Get payment link for Cheng-e-She credit pay](en/com.alibaba.trade.alibaba.creditPay.url.get-1.md) `com.alibaba.trade:alibaba.creditPay.url.get:1` — 获取使用诚e赊支付的支付链接
- [Buyer views all granted credit terms](en/com.alibaba.trade.alibaba.accountPeriod.list.buyerView-1.md) `com.alibaba.trade:alibaba.accountPeriod.list.buyerView:1` — 买家查看获得的所有账期授信
- [Query payment channels supported by an order](en/com.alibaba.trade.alibaba.trade.payWay.query-1.md) `com.alibaba.trade:alibaba.trade.payWay.query:1` — 查询订单可以支持的支付渠道
- [Check whether password-free payment is enabled](en/com.alibaba.trade.alibaba.trade.pay.protocolPay.isopen-1.md) `com.alibaba.trade:alibaba.trade.pay.protocolPay.isopen:1` — 查询是否开通免密支付
- [Initiate password-free payment](en/com.alibaba.trade.alibaba.trade.pay.protocolPay.preparePay-1.md) `com.alibaba.trade:alibaba.trade.pay.protocolPay.preparePay:1` — 发起免密支付
- [Order payment consultation](en/com.alibaba.trade.trade.orderpay.analysis-1.md) `com.alibaba.trade:trade.orderpay.analysis:1` — 交易订单支付咨询
- [Get combined-cashier URL](en/com.alibaba.trade.alibaba.trade.grouppay.url.get-1.md) `com.alibaba.trade:alibaba.trade.grouppay.url.get:1` — 组合收银台url获取

## Logistics

- [Estimate domestic (China) shipping fee for a product](en/com.alibaba.fenxiao.crossborder.product.freight.estimate-1.md) `com.alibaba.fenxiao.crossborder:product.freight.estimate:1` — 商品中国国内运费预估
- [Query external order ID by waybill number or unclaimed-parcel code](en/com.alibaba.fenxiao.crossborder.logistics.order.getOutOrderId-1.md) `com.alibaba.fenxiao.crossborder:logistics.order.getOutOrderId:1` — 根据运单号或无主件码查询外部订单ID
- [Query shipping-insurance info](en/com.alibaba.trade.shipping.insurance.get-1.md) `com.alibaba.trade:shipping.insurance.get:1` — 运费险信息查询
- [Get shipping-template details](en/com.alibaba.logistics.alibaba.logistics.myFreightTemplate.list.get-1.md) `com.alibaba.logistics:alibaba.logistics.myFreightTemplate.list.get:1` — 获取物流模板详情
- [Buyer gets saved shipping addresses](en/com.alibaba.trade.alibaba.trade.receiveAddress.get-1.md) `com.alibaba.trade:alibaba.trade.receiveAddress.get:1` — 买家获取保存的收货地址信息列表
- [Parse an address into area codes](en/com.alibaba.trade.alibaba.trade.addresscode.parse-1.md) `com.alibaba.trade:alibaba.trade.addresscode.parse:1` — 根据地址解析地区码
- [Query seller mixed-batch settings](en/com.alibaba.trade.alibaba.trade.OpQueryMarketingMixConfig-1.md) `com.alibaba.trade:alibaba.trade.OpQueryMarketingMixConfig:1` — 查询卖家混批设置
- [Logistics company list (all companies)](en/com.alibaba.logistics.alibaba.logistics.OpQueryLogisticCompanyList-1.md) `com.alibaba.logistics:alibaba.logistics.OpQueryLogisticCompanyList:1` — 物流公司列表-所有的物流公司
- [Parse overseas address](en/com.alibaba.trade.trade.address.parseText-1.md) `com.alibaba.trade:trade.address.parseText:1` — 海外地址解析
- [Urge seller to ship](en/com.alibaba.logistics.logistics.delivery.urge-1.md) `com.alibaba.logistics:logistics.delivery.urge:1` — 催卖家发货
- [Receive external print-label (UDF) URL](en/com.alibaba.fenxiao.crossborder.label.url.receive-1.md) `com.alibaba.fenxiao.crossborder:label.url.receive:1` — 接收外部打印UDF链接
- [Get waybill-number set](en/com.alibaba.fenxiao.crossborder.bigcustomer.mailNo.query-1.md) `com.alibaba.fenxiao.crossborder:bigcustomer.mailNo.query:1` — 获取运单号集合

## Official Return Pickup

- [Create official return-pickup logistics order](en/com.alibaba.logistics.refundofficialdelivery.order.create-1.md) `com.alibaba.logistics:refundofficialdelivery.order.create:1` — 创建退货官方物流上门揽订单
- [Modify official return-pickup logistics order](en/com.alibaba.logistics.refundofficialdelivery.order.modify-1.md) `com.alibaba.logistics:refundofficialdelivery.order.modify:1` — 退货官方物流上门揽订单修改
- [Get official return-pickup logistics order details](en/com.alibaba.logistics.refundofficialdelivery.order.get-1.md) `com.alibaba.logistics:refundofficialdelivery.order.get:1` — 退货官方物流上门揽订单详情获取
- [Cancel official return-pickup logistics order](en/com.alibaba.logistics.refundofficialdelivery.order.cancel-1.md) `com.alibaba.logistics:refundofficialdelivery.order.cancel:1` — 取消退货官方物流上门揽订单
- [Query official return-pickup plans](en/com.alibaba.logistics.refundofficialdelivery.solution.get-1.md) `com.alibaba.logistics:refundofficialdelivery.solution.get:1` — 官方退货上门揽方案查询

## Data Write-back

- [Write back mapping between end-customer orders and 1688 orders](en/com.alibaba.fenxiao.crossborder.order.relation.write-1.md) `com.alibaba.fenxiao.crossborder:order.relation.write:1` — 回传机构真实用户订单和1688订单的映射关系
- [Write back country-site logistics orders](en/com.alibaba.fenxiao.crossborder.trade.cross.logisticsOrderSync-1.md) `com.alibaba.fenxiao.crossborder:trade.cross.logisticsOrderSync:1` — 国家站物流单回传
- [Sync downstream sales orders](en/com.alibaba.fenxiao.crossborder.trade.cross.orderSync-1.md) `com.alibaba.fenxiao.crossborder:trade.cross.orderSync:1` — 下游销售订单同步
- [Save the business line an account belongs to](en/com.alibaba.fenxiao.crossborder.account.business.save-1.md) `com.alibaba.fenxiao.crossborder:account.business.save:1` — 保存账号所属业务线

## Returns & Refunds

- [Query refund details by order ID (buyer view)](en/com.alibaba.trade.alibaba.trade.refund.OpQueryBatchRefundByOrderIdAndStatus-1.md) `com.alibaba.trade:alibaba.trade.refund.OpQueryBatchRefundByOrderIdAndStatus:1` — 查询退款单详情-根据订单ID（买家视角）
- [Refund operation history (buyer view)](en/com.alibaba.trade.alibaba.trade.refund.OpQueryOrderRefundOperationList-1.md) `com.alibaba.trade:alibaba.trade.refund.OpQueryOrderRefundOperationList:1` — 退款单操作记录列表（买家视角）
- [Query refund details by refund ID (buyer view)](en/com.alibaba.trade.alibaba.trade.refund.OpQueryOrderRefund-1.md) `com.alibaba.trade:alibaba.trade.refund.OpQueryOrderRefund:1` — 查询退款单详情-根据退款单ID（买家视角）
- [Query refund list (buyer view)](en/com.alibaba.trade.alibaba.trade.refund.buyer.queryOrderRefundList-1.md) `com.alibaba.trade:alibaba.trade.refund.buyer.queryOrderRefundList:1` — 查询退款单列表(买家视角)
- 根据理赔单或保险单号查询对应的订单号 `com.alibaba.trade:alibaba.trade.queryOrderByInsure:1` — detail page unavailable on open.1688.com
- [Create refund/return request](en/com.alibaba.trade.alibaba.trade.createRefund-1.md) `com.alibaba.trade:alibaba.trade.createRefund:1` — 创建退款退货申请
- [Query refund/return reasons (for creating a request)](en/com.alibaba.trade.alibaba.trade.getRefundReasonList-1.md) `com.alibaba.trade:alibaba.trade.getRefundReasonList:1` — 查询退款退货原因（用于创建退款退货）
- [Upload refund/return evidence](en/com.alibaba.trade.alibaba.trade.uploadRefundVoucher-1.md) `com.alibaba.trade:alibaba.trade.uploadRefundVoucher:1` — 上传退款退货凭证
- [Buyer submits return-shipment info](en/com.alibaba.trade.alibaba.trade.refund.returnGoods-1.md) `com.alibaba.trade:alibaba.trade.refund.returnGoods:1` — 买家提交退款货信息
- [Cancel refund/return request](en/com.alibaba.trade.alibaba.trade.cancelRefund-1.md) `com.alibaba.trade:alibaba.trade.cancelRefund:1` — 取消退款退货申请 
- [Query maximum refundable amount when applying](en/com.alibaba.trade.alibaba.trade.getMaxRefundFee-1.md) `com.alibaba.trade:alibaba.trade.getMaxRefundFee:1` — 申请退款时查询最大可退费用
- [Apply for trade arbitration](en/com.alibaba.trade.trade.arbitration.apply-1.md) `com.alibaba.trade:trade.arbitration.apply:1` — 申请交易仲裁
- [Query the merchant's return address](en/com.alibaba.fenxiao.refund.address.get-1.md) `com.alibaba.fenxiao:refund.address.get:1` — 查询商家退货地址

## Messaging

- [Batch-confirm failed messages](en/cn.alibaba.open.push.message.confirm-1.md) `cn.alibaba.open:push.message.confirm:1` — 失败消息批量确认
- [Fetch failed messages (query style)](en/cn.alibaba.open.push.query.messageList-1.md) `cn.alibaba.open:push.query.messageList:1` — 查询式获取失败的消息列表
- [Fetch failed messages (cursor style)](en/cn.alibaba.open.push.cursor.messageList-1.md) `cn.alibaba.open:push.cursor.messageList:1` — 游标式获取失败的消息列表
- [View order list (seller view)](en/com.alibaba.trade.alibaba.trade.getSellerOrderList-1.md) `com.alibaba.trade:alibaba.trade.getSellerOrderList:1` — 订单列表查看(卖家视角)
- [Ship: seller arranges own logistics](en/com.alibaba.logistics.alibaba.logistics.OpDeliverySendOrder.offline-1.md) `com.alibaba.logistics:alibaba.logistics.OpDeliverySendOrder.offline:1` — 物流发货-自己联系物流发货
- [Ship: no logistics needed](en/com.alibaba.logistics.alibaba.logistics.OpDeliverySendOrder.dummy-1.md) `com.alibaba.logistics:alibaba.logistics.OpDeliverySendOrder.dummy:1` — 物流发货-无需物流

## Light Customization

- [Get artwork info for a customization order](en/com.alibaba.fenxiao.crossborder.product.podOrderDesign.get-1.md) `com.alibaba.fenxiao.crossborder:product.podOrderDesign.get:1` — 获取加工定制订单稿件信息

## Suppliers

- [Run AI supplier-search task (async)](en/com.alibaba.fenxiao.crossborder.account.search.getInfo.async-1.md) `com.alibaba.fenxiao.crossborder:account.search.getInfo.async:1` — 异步执行AI找商任务
- [Query AI supplier-search task result](en/com.alibaba.fenxiao.crossborder.account.search.getInfoResult-1.md) `com.alibaba.fenxiao.crossborder:account.search.getInfoResult:1` — 查询AI找商任务执行结果

## Inquiries (Newton Cloud)

- [Newton Cloud: create long-running task](en/com.alibaba.agent.newtoncloud.task.create-1.md) `com.alibaba.agent:newtoncloud.task.create:1` — 牛顿云-创建长程任务
- [Newton Cloud: query task](en/com.alibaba.agent.newtoncloud.task.get-1.md) `com.alibaba.agent:newtoncloud.task.get:1` — 牛顿云-查询任务服务
- [Newton Cloud: list tasks](en/com.alibaba.agent.newtoncloud.task.list-1.md) `com.alibaba.agent:newtoncloud.task.list:1` — 牛顿云-查询任务列表
- [Newton Cloud: terminate task](en/com.alibaba.agent.newtoncloud.task.kill-1.md) `com.alibaba.agent:newtoncloud.task.kill:1` — 牛顿云-终止任务
- [Newton Cloud: query task table](en/com.alibaba.agent.newtoncloud.task.fetch-1.md) `com.alibaba.agent:newtoncloud.task.fetch:1` — 牛顿云-查询任务表格服务
- [Newton Cloud: resume task](en/com.alibaba.agent.newtoncloud.task.resume-1.md) `com.alibaba.agent:newtoncloud.task.resume:1` — 牛顿云-恢复任务
- [Newton Cloud: query batch-inquiry results](en/com.alibaba.agent.newtoncloud.batchInquiry.getResult-1.md) `com.alibaba.agent:newtoncloud.batchInquiry.getResult:1` — 牛顿云-查询批量询盘结果
- [Newton Cloud: upload file](en/com.alibaba.agent.newtoncloud.file.upload-1.md) `com.alibaba.agent:newtoncloud.file.upload:1` — 牛顿云-上传文件
- [Newton Cloud: list available model tiers](en/com.alibaba.agent.newtoncloud.model.list-1.md) `com.alibaba.agent:newtoncloud.model.list:1` — 牛顿云-查询可用模型档位列表
- [Newton Cloud: query points details](en/com.alibaba.agent.newtoncloud.points.query-1.md) `com.alibaba.agent:newtoncloud.points.query:1` — 牛顿云-查询积分详情

## 88 ShengYiTong (Business Link)

- [88 ShengYiTong: buyer drafts purchase order](en/com.alibaba.syt.syt.buyer.draftPurchaseOrder-2.md) `com.alibaba.syt:syt.buyer.draftPurchaseOrder:2` — 88生意通买家起草采购单
- [88 ShengYiTong: confirm transaction complete](en/com.alibaba.syt.syt.contract.confirm-1.md) `com.alibaba.syt:syt.contract.confirm:1` — 88生意通确认交易完成
- [88 ShengYiTong: void contract](en/com.alibaba.syt.syt.contract.invalid-1.md) `com.alibaba.syt:syt.contract.invalid:1` — 88生意通合同作废
- [88 ShengYiTong: contract refund request](en/com.alibaba.syt.syt.contract.refund-1.md) `com.alibaba.syt:syt.contract.refund:1` — 88生意通合同退款申请
- [88 ShengYiTong: buyer payment, get cashier URL](en/com.alibaba.syt.syt.contract.pay-1.md) `com.alibaba.syt:syt.contract.pay:1` —  88生意通 买家支付获取收银台 URL
- [88 ShengYiTong: WorldFirst transfer payment (WorldFirst B2C accounts only)](en/com.alibaba.syt.syt.contract.payTransfer-1.md) `com.alibaba.syt:syt.contract.payTransfer:1` — 88生意通万里汇转账支付-仅支持万里汇B2C账号
- [88 ShengYiTong: buyer confirms purchase order](en/com.alibaba.syt.syt.buyer.confirmPurchaseOrder-1.md) `com.alibaba.syt:syt.buyer.confirmPurchaseOrder:1` — 88生意通买家确认采购单
- [88 ShengYiTong: buyer queries purchase-order details](en/com.alibaba.syt.syt.buyer.queryContractDetail-1.md) `com.alibaba.syt:syt.buyer.queryContractDetail:1` — 88生意通买家查询采购单详情
- [88 ShengYiTong: check whether customer is verified](en/com.alibaba.syt.syt.customer.queryUserAuthStatus-1.md) `com.alibaba.syt:syt.customer.queryUserAuthStatus:1` — 88生意通查询客户是否已认证

## Invoicing

- [Query buyer invoice titles (paginated)](en/com.alibaba.trade.trade.invoiceTitle.getPageList-1.md) `com.alibaba.trade:trade.invoiceTitle.getPageList:1` — 分页查询买家抬头列表
- [Query invoiceable amount of orders](en/com.alibaba.trade.trade.invoiceAmount.getList-1.md) `com.alibaba.trade:trade.invoiceAmount.getList:1` — 查询订单可开票金额
- [Query invoice applications (buyer view, paginated)](en/com.alibaba.trade.trade.invoiceApply.getPageListBuyerView-1.md) `com.alibaba.trade:trade.invoiceApply.getPageListBuyerView:1` — 分页查询发票申请列表（买家视角）
- [Query all issued invoices for an order (buyer view)](en/com.alibaba.trade.trade.invoice.getListBuyerView-1.md) `com.alibaba.trade:trade.invoice.getListBuyerView:1` — 查询交易单下关联的所有开具的发票信息（买家视角）
- [Buyer requests invoice](en/com.alibaba.trade.trade.invoice.apply-1.md) `com.alibaba.trade:trade.invoice.apply:1` — 买家申请开票
- [Merged-invoice consultation](en/com.alibaba.trade.trade.invoice.consult-1.md) `com.alibaba.trade:trade.invoice.consult:1` — 合并开票咨询
- [Submit merged-invoice request](en/com.alibaba.trade.trade.invoice.mergeapply-1.md) `com.alibaba.trade:trade.invoice.mergeapply:1` — 提交合并开票
- [Query merged-invoice relationships by order or invoice application](en/com.alibaba.trade.trade.invoice.sellerqueryrelatedorders-1.md) `com.alibaba.trade:trade.invoice.sellerqueryrelatedorders:1` — 基于交易单或发票申请单查询合单关联关系

## Fully Managed (Consignment)

- [Warehouse receipt / shelving of goods](en/com.alibaba.fenxiao.crossborder.consignment.co.status.sync-1.md) `com.alibaba.fenxiao.crossborder:consignment.co.status.sync:1` — 仓库签收/上架商品
- [Warehouse creates discrepancy tally sheet](en/com.alibaba.fenxiao.crossborder.consignment.tally.create-1.md) `com.alibaba.fenxiao.crossborder:consignment.tally.create:1` — 仓库创建差异理货单
- [Update warehouse product inventory](en/com.alibaba.fenxiao.crossborder.consignment.co.inventory-1.md) `com.alibaba.fenxiao.crossborder:consignment.co.inventory:1` — 仓库商品库存更新

## Listing (Publishing)

- [Publish result callback](en/com.alibaba.fenxiao.crossborder.publish.result.callback-1.md) `com.alibaba.fenxiao.crossborder:publish.result.callback:1` — 发布结果回调
- [Listable product list](en/com.alibaba.fenxiao.crossborder.publish.product.list-1.md) `com.alibaba.fenxiao.crossborder:publish.product.list:1` — 铺货商品列表
- [Listing product card](en/com.alibaba.fenxiao.crossborder.publish.card.get-1.md) `com.alibaba.fenxiao.crossborder:publish.card.get:1` — 铺货商品卡片

## Repurchase

- [Query repurchase contract](en/com.alibaba.trade.repurchase.contract.get-1.md) `com.alibaba.trade:repurchase.contract.get:1` — 复购合约查询
