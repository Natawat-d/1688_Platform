# Preview data before creating an order

Original name: 创建订单前预览数据接口  
API: `com.alibaba.trade:alibaba.createOrder.preview:1` · Category: Orders  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.createOrder.preview-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.createOrder.preview/{appKey}`  
Requires user authorization (access_token) · Requires signature

Orders may only contain products from a single supplier. This interface returns discount and related information for order creation. It 1. validates whether the products may be ordered; 2. validates the consignment (distribution) relationship; 3. validates stock, minimum order quantity and whether mixed-batch conditions are met.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `addressParam` | [message:alibaba.trade.fast.address](#m-alibaba-trade-fast-address) | yes | Shipping address info | {"address":"网商路699号","phone": "0517-88990077","mobile": "15251667788","fullName": "张三","postCode": "000000","areaText": "滨江区","townText": "","cityText": "杭州市","provinceText": "浙江省"} |
| `cargoParamList` | [message:alibaba.trade.fast.cargo[]](#m-alibaba-trade-fast-cargo[]) | yes | Product information | [{"specId": "b266e0726506185beaf205cbae88530d","quantity": 5,"offerId": 554456348334},{"specId": "2ba3d63866a71fbae83909d9b4814f01","quantity": 6,"offerId": 554456348334}] |
| `invoiceParam` | [message:alibaba.trade.fast.invoice](#m-alibaba-trade-fast-invoice) | no | Invoice information | {"invoiceType":0,"cityText": "杭州市","provinceText": "浙江省","address": "网商路699号","phone": "0517-88990077","mobile": "15251667788","fullName": "张五","postCode": "000000","areaText": "滨江区","companyName": "测试公司","taxpayerIdentifier": "123455"} |
| `flow` | String | no | general (create open marketplace order), fenxiao (create distribution order), paired (Tiantian Temai / Daily Deals), repurchase (repurchase contract order), saleproxy process will verify the distribution relationship, boutiquefenxiao (curated supply distribution-price order, free shipping when purchase quantity is 1), boutiquepifa (curated supply wholesale-price order, used when purchase quantity is greater than 2). If flow is empty, the system compares prices and previews the best option, returning the optimal order flow. Repurchase contract orders are not included in the price comparison. All order flow channels may cause differences in price and buyer protection services. | general |
| `instanceId` | String | no | Wholesale group-buy instanceId, obtained from alibaba.pifatuan.product.list | 4063139_1662080400000 |
| `encryptOutOrderInfo` | [message:alibaba.trade.fastCreateOrder.EncryptOutOrderInfo](#m-alibaba-trade-fastcreateorder-encryptoutorderinfo) | no | Downstream encrypted order information, used for downstream label printing | {} |
| `proxySettleRecordId` | String | no | Purchase order ID for split-payment regular order placement, where the transaction flow is "proxy" | 4051300002 |
| `inventoryMode` | String | no | Stock mode, jit (JIT mode) or cang (warehouse shipping mode); currently only available for AE | jit |
| `outOrderId` | String | no | External order number | 988129883123 |
| `pickupService` | String | no | Door-to-door pickup; currently available for AE supply, not yet enabled for other scenarios | y或n,默认为n |
| `crossBorderLogisticsSolutionId` | String | no | The sourceId of the official cross-border logistics solution selected by the user | GLOBAL_CAINIAO_VN_TRANSIT_LAND  |
| `useBorderLogisticsSolution` | Boolean | no | Whether to use the cross-border logistics solution preview. Defaults to false. This parameter does not take effect for overseas addresses. It takes effect for Hong Kong, Macao, and Taiwan addresses. | true |
| `isvBizType` | String | no | Open platform business code | cross |

<a id="m-alibaba-trade-fast-address"></a>
#### alibaba.trade.fast.address

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `addressId` | Long | yes | Shipping address ID | 1234 |
| `fullName` | String | yes | Recipient name | 张三 |
| `mobile` | String | yes | Mobile phone | 15251667788 |
| `phone` | String | yes | Phone number | 0517-88990077 |
| `postCode` | String | yes | Postal code | 000000 |
| `cityText` | String | yes | City text | 杭州市 |
| `provinceText` | String | yes | Province text | 浙江省 |
| `areaText` | String | yes | District text | 滨江区 |
| `townText` | String | yes | Town text | 长河镇 |
| `address` | String | yes | Street address | 网商路699号 |
| `districtCode` | String | yes | Address code | 310107 |

<a id="m-alibaba-trade-fast-cargo[]"></a>
#### alibaba.trade.fast.cargo[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `offerId` | Long | yes | offer id corresponding to the product | 554456348334 |
| `specId` | String | yes | Product SKU ID | b266e0726506185beaf205cbae88530d |
| `quantity` | Double | yes | Product quantity (used for amount calculation) | 5 |
| `openOfferId` | String | no | Encrypted offerId. When the search result returns only openOfferId, use openOfferId in place of offerId to place the order. | Wcv2w970KoL1BCpJGQp3vwKvcBThOPlpHnUtkK3T0n4= |
| `outMemberId` | String | no | External downstream member ID | 98928912-23 |
| `bizFenXiaoOutSubOrderInfo` | [message:com.alibaba.ocean.openplatform.biz.trade.common.model.BizFenXiaoOutSubOrderInfo[]](#m-com-alibaba-ocean-openplatform-biz-trade-common-model-bizfenxiaooutsuborderinfo[]) | no | Downstream sub-order data for distribution | {} |

<a id="m-com-alibaba-ocean-openplatform-biz-trade-common-model-bizfenxiaooutsuborderinfo[]"></a>
#### com.alibaba.ocean.openplatform.biz.trade.common.model.BizFenXiaoOutSubOrderInfo[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `outSubOrderId` | String | no | External order id | 123 |
| `outSkuId` | String | no | External sku | 123 |
| `outSubGmv` | Double | no | gmv | 123 |
| `outSubQty` | Double | no | Order volume | 123 |
| `outSubOrderPayTime` | String | no | Order placement time of the downstream sub-order | 2025-03-18 09:20:21 |
| `outSubOrderLatestDeliveryTime` | String | no | Latest shipping time for the downstream main order | 2025-03-18 09:20:21 |
| `outSubGmvDiscount` | Double | no | Discount amount of the downstream sub-order | 10 |
| `outSubGmvPost` | Double | no | Shipping fee for the downstream sub-order | 5 |
| `outSubGmvReceive` | Double | no | Actual amount received for the downstream sub-order | 118 |

<a id="m-alibaba-trade-fast-invoice"></a>
#### alibaba.trade.fast.invoice

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `invoiceType` | Integer | yes | Invoice type<br>0: regular invoice, 1: VAT invoice | 0 |
| `provinceText` | String | yes | Province text | 浙江省 |
| `cityText` | String | yes | City text | 杭州市 |
| `areaText` | String | yes | Region text | 滨江区 |
| `townText` | String | yes | Town text | 长河镇 |
| `postCode` | String | yes | Postal code | 333333 |
| `address` | String | yes | Street | 网商路699号 |
| `fullName` | String | yes | Invoice recipient's name | 张三 |
| `phone` | String | yes | Phone number | 0517-88990077 |
| `mobile` | String | yes | Mobile phone | 15251667788 |
| `companyName` | String | yes | Purchasing company name (invoice title) | 测试公司 |
| `taxpayerIdentifier` | String | yes | Tax identification number | 12345 |
| `bankAndAccount` | String | yes | Bank name and account number | 网商银行 |
| `localInvoiceId` | String | yes | VAT local invoice number | 123123123 |

<a id="m-alibaba-trade-fastcreateorder-encryptoutorderinfo"></a>
#### alibaba.trade.fastCreateOrder.EncryptOutOrderInfo

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `encryptOrder` | Boolean | yes | Whether the order is encrypted | true |
| `outPlatformOrderNo` | String | yes | Downstream platform order number | 12365452354551 |
| `outPlatformSupplyOrderNo` | String | no | Downstream platform supply chain purchase order | CT7403712218870825259 |
| `outPlatformCode` | String | yes | Taobao-thyny, Tmall-tm, Taote-taote, Alibaba C2M-c2m, JD.com-jingdong, Pinduoduo-pinduoduo, WeChat Store-weixin, Cross-border-kuajing, Kuaishou-kuaishou, Youzan-youzan, Douyin-douyin, Secoo-siku, Meituan Tuanhaohuo-meituan, Xiaohongshu-xiaohongshu, Dangdang-dangdang, Suning-suning, DaVdian-davdian, Xingyun-xingyun, Miya-miya, Boluopai Mall-boluo, Other-other | taote |
| `outPlatformAppkey` | String | yes | The appkey used by the downstream platform to retrieve orders | 32154 |
| `outShopId` | String | no | Downstream platform shop Id | 1879283 |
| `outShopName` | String | no | Downstream platform shop name | 三生科技 |
| `outOriginAddress` | [message:com.alibaba.ocean.openplatform.biz.trade.param.OutAddress](#m-com-alibaba-ocean-openplatform-biz-trade-param-outaddress) | no | External original address information | {} |
| `oaid` | String | no | Taobao oaid | 265646-52342354-2354Akf-w3654SF |
| `outPatformExtraInfo` | String | no | Other extended information from the downstream platform | {} |
| `encryptReceiverName` | String | no | Downstream encrypted recipient name | *** |
| `encryptReceiverMobile` | String | no | Downstream encrypted recipient phone number | *** |
| `encryptReceiverAddress` | String | no | Downstream encrypted recipient address | *** |
| `outPlatformSubCode` | String | no | Downstream channel sub-business code, e.g. Douyin sub-channel 101, used to identify supply chain orders | 101 |

<a id="m-com-alibaba-ocean-openplatform-biz-trade-param-outaddress"></a>
#### com.alibaba.ocean.openplatform.biz.trade.param.OutAddress

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `province` | [message:com.alibaba.ocean.openplatform.biz.trade.param.Place](#m-com-alibaba-ocean-openplatform-biz-trade-param-place) | yes | Province | {"name":"四川省","code":"51000"} |
| `city` | [message:com.alibaba.ocean.openplatform.biz.trade.param.Place](#m-com-alibaba-ocean-openplatform-biz-trade-param-place) | yes | City | {} |
| `area` | [message:com.alibaba.ocean.openplatform.biz.trade.param.Place](#m-com-alibaba-ocean-openplatform-biz-trade-param-place) | yes | District | {} |
| `town` | [message:com.alibaba.ocean.openplatform.biz.trade.param.Place](#m-com-alibaba-ocean-openplatform-biz-trade-param-place) | no | Town/street | {} |
| `address` | String | no | Detailed address | 网商路699号 |
| `postCode` | String | no | Postal code | 511304 |

<a id="m-com-alibaba-ocean-openplatform-biz-trade-param-place"></a>
#### com.alibaba.ocean.openplatform.biz.trade.param.Place

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `code` | String | yes | Address code | 511300 |
| `name` | String | yes | Address name | 南充 |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `orderPreviewResuslt` | [message:alibaba.createOrder.preview.result.model[]](#m-alibaba-createorder-preview-result-model[]) | yes | Order preview result; if automatic order splitting occurs, multiple records will be returned | [] |
| `success` | Boolean | yes | Whether successful | true |
| `errorCode` | String | yes | Error code | 500_1 |
| `errorMsg` | String | yes | Error message | 错误 |
| `postFeeByDescOfferList` | Long[] | yes | Product list for the freight description | [12324324234,12312422] |
| `consignOfferList` | Long[] | yes | Consignment product list | [12324324234,12312422] |
| `unsupportedCrossBorderPayOfferList` | Long[] | yes | List of products that do not support Kuajingbao payment | [12324324234,12312422] |
| `extPairList` | [message:com.alibaba.ocean.openplatform.biz.trade.common.model.ExtrPair[]](#m-com-alibaba-ocean-openplatform-biz-trade-common-model-extrpair[]) | yes | List of extended objects | [] |

<a id="m-alibaba-createorder-preview-result-model[]"></a>
#### alibaba.createOrder.preview.result.model[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `discountFee` | java.lang.Long | yes | The additional reduction amount applied after calculating the product amount. Unit: cents (fen). | ''  |
| `tradeModeNameList` | String[] | yes | List of transaction methods supported by the current transaction when using the place-order API; elements in this list can be used directly as the tradeType input parameter of the place-order API. If the list is empty, the current transaction cannot be placed via the API and must be placed on the 1688 page. |  ''  |
| `status` | boolean | yes | Status |  ''  |
| `taoSampleSinglePromotion` | boolean | yes | Whether there is a Taobao supply-source single-item discount. false: has single-item discount; true: no single-item discount |  ''  |
| `sumPayment` | long | yes | Total order fee, in cents (fen). |  ''  |
| `message` | java.lang.String | yes | Return message |  ''  |
| `sumCarriage` | long | yes | Total shipping fee information, in cents. |  ''  |
| `resultCode` | java.lang.String | yes | Return code |  ''  |
| `sumPaymentNoCarriage` | long | yes | Total product cost excluding shipping fee, in cents. |  ''  |
| `additionalFee` | java.lang.Long | yes | Additional fee, in cents (fen) | ''   |
| `flowFlag` | java.lang.String | yes | Order placement process |  ''  |
| `cargoList` | [message:alibaba.createOrder.preview.resultCargo.model[]](#m-alibaba-createorder-preview-resultcargo-model[]) | yes | SKU information |  ''  |
| `shopPromotionList` | [message:alibaba.trade.promotion.model[]](#m-alibaba-trade-promotion-model[]) | yes | List of available shop-level discounts |  ''  |
| `tradeModelList` | [message:tradeModelExtensionList[]](#m-trademodelextensionlist[]) | yes | List of transaction methods supported by the current transaction. Refer to the transaction methods shown on the 1688 order preview page for the results. |  ''  |
| `payChannelInfos` | [message:payChaneelList[]](#m-paychaneellist[]) | yes | Payment channel information supported by the current transaction | [] |
| `tradeServiceList` | [message:com.alibaba.ocean.openplatform.biz.trade.common.model.TradeServiceGroup[]](#m-com-alibaba-ocean-openplatform-biz-trade-common-model-tradeservicegroup[]) | yes | List of services provided, such as cross-border logistics services, etc. | {} |
| `extPairList` | [message:com.alibaba.ocean.openplatform.biz.trade.common.model.ExtrPair[]](#m-com-alibaba-ocean-openplatform-biz-trade-common-model-extrpair[]) | yes | List of extended objects | [] |
| `orderGroup` | String | yes | Order group. When specifying an official logistics pickup plan, grouping must be specified using this field. | 123 |
| `canUseOfficialSolution` | Boolean | yes | Whether official logistics pickup is available | true |
| `officialSolutionModelList` | [message:com.alibaba.ocean.openplatform.biz.trade.common.model.OfficialSolutionModel[]](#m-com-alibaba-ocean-openplatform-biz-trade-common-model-officialsolutionmodel[]) | yes | List of official logistics pickup services | [{"solutionCode":"471","solutionName":"特惠当日上门","totalCost":770}] |
| `totalFundUsageAmount` | Long | yes | Total amount of red envelope used, in cents (fen) | 10 |
| `totalPostFundUsageAmount` | Long | yes | Total amount of shipping coupon used, in cents | 100 |

<a id="m-alibaba-createorder-preview-resultcargo-model[]"></a>
#### alibaba.createOrder.preview.resultCargo.model[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `amount` | java.lang.Double | yes | Total product amount | 10 |
| `message` | java.lang.String | yes | Return message | null |
| `finalUnitPrice` | java.lang.Double | yes | Final unit price | 100 |
| `specId` | java.lang.String | yes | SKU ID, unique within the offer | 98SDKTKSLFasbFiAlf |
| `skuId` | java.lang.Long | yes | SKU ID, globally unique | 245252612367 |
| `resultCode` | java.lang.String | yes | Return code | 200 |
| `offerId` | java.lang.Long | yes | Product ID | 909872988917 |
| `openOfferId` | String | yes | Encrypted product ID | 9JlsKsfcKLAKDFS |
| `cargoPromotionList` | [message:alibaba.trade.promotion.model[]](#m-alibaba-trade-promotion-model[]) | yes | List of product discounts | {} |
| `extPairList` | [message:com.alibaba.ocean.openplatform.biz.trade.common.model.ExtrPair[]](#m-com-alibaba-ocean-openplatform-biz-trade-common-model-extrpair[]) | yes | List of extended objects | [] |
| `totalFundUsageAmount` | Long | yes | Total amount of red packet used, in cents | 10 |
| `totalPostFundUsageAmount` | Long | yes | Total amount of shipping coupon used, in cents (fen) | 10 |
| `tradeServiceList` | [message:com.alibaba.ocean.openplatform.biz.trade.common.model.TradeServiceGroup[]](#m-com-alibaba-ocean-openplatform-biz-trade-common-model-tradeservicegroup[]) | yes | Service - services related to the product | {} |

<a id="m-alibaba-trade-promotion-model[]"></a>
#### alibaba.trade.promotion.model[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `promotionId` | java.lang.String | yes | Coupon ID |   |
| `selected` | boolean | yes | Whether it is selected by default |   |
| `text` | java.lang.String | yes | Coupon name |   |
| `desc` | java.lang.String | yes | Coupon description |   |
| `freePostage` | boolean | yes | Whether shipping is free |   |
| `discountFee` | java.lang.Long | yes | Amount deducted, in cents (fen) |   |

<a id="m-trademodelextensionlist[]"></a>
#### tradeModelExtensionList[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `name` | String | yes | Transaction method name, the name shown on the 1688 order preview page |   |
| `description` | String | yes | Transaction description |   |
| `tradeType` | String | yes | The tradeType field passed as an input parameter to the order placement API |   |
| `opSupport` | Boolean | yes | Whether the Open Platform order placement supports this transaction mode. If true, this transaction method can be used as the input value for the tradeType parameter of the order placement API; if false, it cannot be used as an input value for the order placement API. |   |

<a id="m-paychaneellist[]"></a>
#### payChaneelList[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `name` | String | yes | Payment channel | "alipay" |
| `amountLimit` | Long | yes | Available credit amount, in cents | 100 |

<a id="m-com-alibaba-ocean-openplatform-biz-trade-common-model-tradeservicegroup[]"></a>
#### com.alibaba.ocean.openplatform.biz.trade.common.model.TradeServiceGroup[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `outId` | String | yes | id | '' |
| `groupTitle` | String | yes | Group name | '' |
| `groupCode` | String | yes | Group code | '' |
| `groupDescLink` | String | yes | Group description link | '' |
| `tradeServiceList` | [message:com.alibaba.ocean.openplatform.biz.trade.common.model.TradeService[]](#m-com-alibaba-ocean-openplatform-biz-trade-common-model-tradeservice[]) | yes | Service information | '' |
| `extraMapStr` | String | yes | Extended information | '' |
| `combinedPrice` | Long | yes | Combined price | '' |
| `showCombinedPrice` | String | yes | Displayed combined price | '' |

<a id="m-com-alibaba-ocean-openplatform-biz-trade-common-model-tradeservice[]"></a>
#### com.alibaba.ocean.openplatform.biz.trade.common.model.TradeService[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `sourceType` | String | yes | Source type | '' |
| `serviceType` | String | yes | Service type | '' |
| `sourceId` | String | yes | Source id | '' |
| `title` | String | yes | Name | '' |
| `unitPrice` | Long | yes | Unit price | '' |
| `quantity` | Double | yes | Quantity | '' |
| `serviceTips` | String | yes | Service tip | '' |
| `serviceUrl` | String | yes | Service URL | '' |
| `groupTitle` | String | yes | Group name | '' |
| `groupCode` | String | yes | Group code | '' |
| `sponsorDesc` | String | yes | Description | '' |
| `code` | String | yes | code | '' |
| `extMapStr` | String | yes | Extended information | '' |

<a id="m-com-alibaba-ocean-openplatform-biz-trade-common-model-officialsolutionmodel[]"></a>
#### com.alibaba.ocean.openplatform.biz.trade.common.model.OfficialSolutionModel[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `solutionCode` | String | yes | Logistics solution code | 471 |
| `solutionName` | String | yes | Logistics solution name | 特惠当日上门 |
| `totalCost` | Long | yes | Total cost (cents) | 770 |

<a id="m-com-alibaba-ocean-openplatform-biz-trade-common-model-extrpair[]"></a>
#### com.alibaba.ocean.openplatform.biz.trade.common.model.ExtrPair[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `key` | String | yes | Key | doorPickup |
| `value` | String | yes | Value | {       "serviceCode": "mkd",       "price": 1200, // 人民币，分     } |
| `desc` | String | yes | Description | 是否使用上门揽 |

## Error codes

| Error | Symptom | How to fix |
|---|---|---|
| 500_001 | Product [offerId] does not support online transactions; an order cannot be placed. | The product does not support online transactions and currently cannot be purchased |
| 500_002 | Product [offerId] does not belong to the same seller, or specId was not specified. | There are products from multiple sellers, or a product does not have a specId specified |
| 500_003 | Product [offerId] does not belong to the same seller, or spec [specId] does not belong to product [offerId] | Product belongs to multiple sellers, or the product does not have a SKU with this specId |
| 500_004 | Product [offerId_specId] has insufficient stock; please verify the stock before ordering. | Insufficient stock for a certain SKU of the product |
| 500_005 | The purchase quantity for product [offerId] does not meet the minimum order quantity limit. | The purchase quantity of the product is less than the minimum order quantity |
| 500_006 | The purchase quantity or price of product [offerId] does not meet the mixed batch limit. | Neither the purchase quantity nor the total amount of the product meets the mixed batch condition |
| 500_007 | The consignment relationship with the supplier does not exist; orders cannot be placed via the saleproxy channel. | flow cannot use slproxy |
| 500_009 | The purchase quantity of product [offerId] does not meet the wholesale minimum order quantity (MOQ) limit. | Check the purchase quantity of the wholesale product |
| 500_008 | The price of product SKU [offerId_specId] is 0; an order cannot be placed. Please check and resubmit. | Check the price of the product SKU |

## Samples

**Output parameter example**

```
{
    "orderPreviewResuslt":[
        {
            "tradeModeNameList":[
                "nzassure"
            ],
            "status":true,
            "taoSampleSinglePromotion":false,
            "sumPayment":4400,
            "sumCarriage":800,
            "sumPaymentNoCarriage":3600,
            "flowFlag":"general",
            "cargoList":[
                {
                    "amount":36,
                    "finalUnitPrice":18,
                    "specId":"eb81c61de14f4adb405ffcc2c8a4a3fb",
                    "skuId":3332085412530,
                    "offerId":548818919805,
                    "cargoPromotionList":[

                    ]
                }
            ],
            "shopPromotionList":[

            ],
            "tradeModelList":[
                {
                    "tradeType":"nzassure",
                    "name":"担保交易",
                    "description":"买家下单5天内全款支付。自卖家发货起，买家需在10天内确认收货，确认收货后打款给卖家。特别提醒：如卖家已开通极速到账服务，将优先适用极速到账交易，买家支付款项将直接打入卖家账户，详见《极速到账交易及争议处理规则》。",
                    "opSupport":true
                }
            ]
        }
    ],
    "success":true
}
```
