# Multilingual keyword search

Original name: 多语言关键词搜索  
API: `com.alibaba.fenxiao.crossborder:product.search.keywordQuery:1` · Category: Products  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:product.search.keywordQuery-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/product.search.keywordQuery/{appKey}`  
Requires user authorization (access_token) · Requires signature

Multilingual keyword search.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `offerQueryParam` | [message:product.search.keywordQuery.param.OfferQueryParam](#m-product-search-keywordquery-param-offerqueryparam) | yes | Query parameter | {} |

<a id="m-product-search-keywordquery-param-offerqueryparam"></a>
#### product.search.keywordQuery.param.OfferQueryParam

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `keyword` | java.lang.String | yes | Keyword | 饼干 |
| `beginPage` | java.lang.Integer | yes | Pagination | 1 |
| `pageSize` | java.lang.Integer | yes | Pagination | 1 |
| `filter` | java.lang.String | no | Filter parameters, multiple values separated by commas; see the solution introduction for the enumeration | shipInToday,ksCiphertext |
| `sort` | java.lang.String | no | Sorting parameter; for enum values see the solution introduction | {"price":"asc"} |
| `outMemberId` | java.lang.String | no | External user ID | 123 |
| `priceStart` | java.lang.String | no | Wholesale price start | 1 |
| `priceEnd` | java.lang.String | no | Wholesale price range end | 10 |
| `categoryId` | java.lang.Long | no | Category ID | 1 |
| `categoryIdList` | String | no | List of category IDs, separated by English commas; supports the union of multiple categories | 2,45 |
| `country` | java.lang.String | yes | Language | 如en-英语，详细枚举参考开发人员参考菜单 |
| `regionOpp` | String | no | Business opportunity | 枚举值见开发人员参考菜单 |
| `productCollectionId` | String | no | Xunyuantong workbench assortment id | 174316138 |
| `snId` | String | no | Search navigation ID, e.g. 978 or 978:1352 | 978:1352 |
| `keywordTranslate` | Boolean | no | Whether the keyword has already been translated; default is not translated. If true, keyword translation is skipped | false |
| `saleFilterList` | [message:com.alibaba.cbu.offer.param.SaleFilterParam[]](#m-com-alibaba-cbu-offer-param-salefilterparam[]) | no | Sales volume filter parameter | [{"saleType":"sales7","saleStart":"10","saleEnd":"100"}] |

<a id="m-com-alibaba-cbu-offer-param-salefilterparam[]"></a>
#### com.alibaba.cbu.offer.param.SaleFilterParam[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `saleType` | String | no | Sales volume type | sales7:近7天销量,sales14:近14天销量,sales30:近30天销量,totalSales:总销量，传入多个取交集 |
| `saleStart` | String | no | Minimum sales volume | 10 |
| `saleEnd` | String | no | Maximum sales volume | 100 |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:product.search.keywordQuery.result.ResultModelV3](#m-product-search-keywordquery-result-resultmodelv3) | yes | Return message | 返回信息 |

<a id="m-product-search-keywordquery-result-resultmodelv3"></a>
#### product.search.keywordQuery.result.ResultModelV3

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `success` | Boolean | yes | Whether normal | 正否正常 |
| `code` | String | yes | Status code | 状态码 |
| `message` | String | yes | Prompt | 提示 |
| `result` | [message:product.search.keywordQuery.model.PageInfoV3](#m-product-search-keywordquery-model-pageinfov3) | yes | Content | 内容 |

<a id="m-product-search-keywordquery-model-pageinfov3"></a>
#### product.search.keywordQuery.model.PageInfoV3

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `totalRecords` | Integer | yes | Total count | 分页 |
| `totalPage` | Integer | yes | Total page count | 分页 |
| `pageSize` | Integer | yes | Pagination | 分页 |
| `currentPage` | Integer | yes | Pagination | 分页 |
| `data` | [message:product.search.keywordQuery.model.ProductInfoModelV2[]](#m-product-search-keywordquery-model-productinfomodelv2[]) | yes | Data | 数据 |

<a id="m-product-search-keywordquery-model-productinfomodelv2[]"></a>
#### product.search.keywordQuery.model.ProductInfoModelV2[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `imageUrl` | String | yes | Image URL | 图片地址 |
| `aigcImageUrl` | String | yes | Image URL - after AIGC processing and translation | https:// |
| `subject` | String | yes | Chinese title | 中文标题 |
| `subjectTrans` | String | yes | Foreign-language title | 外文标题 |
| `offerId` | Long | yes | Product ID | 2 |
| `isJxhy` | Boolean | yes | Whether it is curated supply | true |
| `priceInfo` | [message:product.search.keywordQuery.model.PriceInfoV2](#m-product-search-keywordquery-model-priceinfov2) | yes | Price | 1 |
| `repurchaseRate` | String | yes | Repurchase rate | 10% |
| `monthSold` | Integer | yes | 30-day sales volume | 1213 |
| `traceInfo` | String | yes | Report tracking data to 1688 | object_id@620201390233^object_type@offer |
| `isOnePsale` | Boolean | yes | Whether it is dropshipping | true |
| `sellerIdentities` | String[] | yes | Merchant identity | super_factory-超级工厂 powerful_merchants-实力商家 tp_member-诚信通会员 |
| `offerIdentities` | String[] | yes | Product tag | yx-严选，select-跨境select |
| `tradeScore` | String | yes | Product transaction rating | 5.0 |
| `whiteImage` | String | yes | Product white-background image | 商品白底图 |
| `promotionModel` | [message:product.search.keywordQuery.model.PromotionModelV2](#m-product-search-keywordquery-model-promotionmodelv2) | yes | Whether there is marketing information | 目前只透plus |
| `topCategoryId` | Long | yes | Top-level category | 1 |
| `secondCategoryId` | Long | yes | Secondary category | 2 |
| `thirdCategoryId` | Long | yes | Third-level category | 3 |
| `isPatentProduct` | Boolean | yes | Whether it is a patented product | true |
| `createDate` | String | yes | Product listing time | 2024-04-20 08:00:00 |
| `modifyDate` | String | yes | Product modification time | 2024-04-20 08:00:00 |
| `isSelect` | Boolean | yes | Cross-border select assortment | true |
| `minOrderQuantity` | Integer | yes | Minimum order quantity | 1 |
| `sellerDataInfo` | [message:product.search.keywordQuery.model.SellerDataInfoV1](#m-product-search-keywordquery-model-sellerdatainfov1) | yes | Product data | 1 |
| `productSimpleShippingInfo` | [message:product.search.keywordQuery.model.ProductSimpleShippingInfo](#m-product-search-keywordquery-model-productsimpleshippinginfo) | yes | Brief shipping information | 1 |
| `token` | String | yes | Plugin rebate token | abc |
| `promotionURL` | String | yes | Link to the 1688 product detail page with the [AI Cross-Border Operations Assistant] module | 商品详情页链接 |
| `sales7d` | String | yes | Sales volume in the last 7 days | 1 |
| `productTradeInfo` | [message:com.alibaba.cbu.offer.model.ProductTradeInfo](#m-com-alibaba-cbu-offer-model-producttradeinfo) | yes | Product transaction data | 1 |
| `invoiceInfo` | [message:product.search.queryProductDetail.model.InvoiceInfo](#m-product-search-queryproductdetail-model-invoiceinfo) | yes | Invoice information | "invoiceInfo": {         "supportOnlineInvoice": false,         "supportFastInvoice": false,         "invoiceTypes": [           "普票"         ],         "taxpayerType": "一般纳税人"       } |

<a id="m-product-search-keywordquery-model-priceinfov2"></a>
#### product.search.keywordQuery.model.PriceInfoV2

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `price` | String | yes | Wholesale price | 10 |
| `jxhyPrice` | String | yes | Curated supply price for dropshipping | 10 |
| `pfJxhyPrice` | String | yes | Curated supply wholesale price | 10 |
| `consignPrice` | String | yes | Dropshipping price | 10 |
| `promotionPrice` | String | yes | Marketing price | 10 |

<a id="m-product-search-keywordquery-model-promotionmodelv2"></a>
#### product.search.keywordQuery.model.PromotionModelV2

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `hasPromotion` | Boolean | yes | Whether there is marketing/promotion | true |
| `promotionType` | String | yes | Marketing type | plus |

<a id="m-product-search-keywordquery-model-sellerdatainfov1"></a>
#### product.search.keywordQuery.model.SellerDataInfoV1

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `tradeMedalLevel` | String | yes | Seller's trade medal level | 1 |
| `compositeServiceScore` | String | yes | Overall service experience score | 1 |
| `logisticsExperienceScore` | String | yes | Logistics experience score | 1 |
| `disputeComplaintScore` | String | yes | Dispute/complaint handling score | 1 |
| `offerExperienceScore` | String | yes | Product experience score | 1 |
| `afterSalesExperienceScore` | String | yes | After-sales experience score | 1 |
| `consultingExperienceScore` | String | yes | Inquiry experience score | 1 |
| `repeatPurchasePercent` | String | yes | Repeat purchase rate | 11 |
| `tpYear` | Integer | yes | Chengxintong membership years | 1 |

<a id="m-product-search-keywordquery-model-productsimpleshippinginfo"></a>
#### product.search.keywordQuery.model.ProductSimpleShippingInfo

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `shippingTimeGuarantee` | String | yes | shipIn24Hours - 24-hour shipping shipIn48Hours - 48-hour shipping | shipIn24Hours |
| `perfectFulfillmentRate30d` | String | yes | Perfect fulfillment rate over the last 30 days; empty value means 0 | 100.00% |
| `pickupWithin24hRate30d` | String | yes | 24-hour pickup rate over the last 30 days; a null value represents 0. | 99.12% |
| `qualityReturnRate30d` | String | yes | Quality return rate over the last 30 days; empty value means 0 | 0.12% |
| `perfectFulfillmentRate7d` | String | yes | Perfect fulfillment rate over the last 7 days; a null value represents 0. | 0.00% |
| `pickupWithin24hRate7d` | String | yes | 24-hour pickup rate over the last 7 days; a null value represents 0. | 95.12% |
| `delayedShippingRate7d` | String | yes | Delayed shipment rate over the last 7 days; empty value means 0 | 12.34% |

<a id="m-com-alibaba-cbu-offer-model-producttradeinfo"></a>
#### com.alibaba.cbu.offer.model.ProductTradeInfo

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `addCartCount7d` | String | yes | Number of times added to cart in the last 7 days | 1 |
| `payBuyerCount7d` | String | yes | Number of paying buyers in the last 7 days | 2 |
| `addCartCount30d` | String | yes | Number of times added to cart in the last 30 days | 3 |
| `payBuyerCount30d` | String | yes | Number of paying buyers in the last 30 days | 4 |

<a id="m-product-search-queryproductdetail-model-invoiceinfo"></a>
#### product.search.queryProductDetail.model.InvoiceInfo

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `supportOnlineInvoice` | Boolean | yes | Whether online invoicing is supported | true |
| `supportFastInvoice` | Boolean | yes | Whether instant invoicing is supported | false |
| `invoiceTypes` | String[] | yes | List of supported invoicing types | ["普票","专票"] |
| `taxpayerType` | String | yes | Taxpayer type | 一般纳税人 \|  小规模纳税人 |

## Samples

**Input parameter example**

```
{
  "offerQueryParam": {
    "keyword": "饼干",
    "beginPage": 1,
    "pageSize": 1,
    "filter": "shipInToday,ksCiphertext",
    "sort": "{\"price\":\"asc\"}",
    "outMemberId": "123",
    "priceStart": "1",
    "priceEnd": "10",
    "categoryId": 1,
    "categoryIdList": "2,45",
    "country": "如en-英语，详细枚举参考开发人员参考菜单",
    "regionOpp": "枚举值见开发人员参考菜单",
    "productCollectionId": "174316138",
    "snId": "978:1352",
    "keywordTranslate": false,
    "saleFilterList": [
      {
        "saleType": "sales7:近7天销量,sales14:近14天销量,sales30:近30天销量,totalSales:总销量，传入多个取交集",
        "saleStart": "10",
        "saleEnd": "100"
      }
    ]
  }
}
```

**Output parameter example**

```
{
    "result": {
        "success": true,
        "code": "状态码",
        "message": "提示",
        "result": {
            "totalRecords": 123,
            "totalPage": 123,
            "pageSize": 123,
            "currentPage": 123,
            "data": [
                {
                    "imageUrl": "图片地址",
                    "aigcImageUrl": "https://",
                    "subject": "中文标题",
                    "subjectTrans": "外文标题",
                    "offerId": 2,
                    "isJxhy": true,
                    "priceInfo": {
                        "price": "10",
                        "jxhyPrice": "10",
                        "pfJxhyPrice": "10",
                        "consignPrice": "10",
                        "promotionPrice": "10"
                    },
                    "repurchaseRate": "10%",
                    "monthSold": 1213,
                    "traceInfo": "object_id@620201390233^object_type@offer",
                    "isOnePsale": true,
                    "sellerIdentities": [
                        "super_factory-超级工厂 powerful_merchants-实力商家 tp_member-诚信通会员"
                    ],
                    "offerIdentities": [
                        "yx-严选，select-跨境select"
                    ],
                    "tradeScore": "5.0",
                    "whiteImage": "商品白底图",
                    "promotionModel": {
                        "hasPromotion": true,
                        "promotionType": "plus"
                    },
                    "topCategoryId": 1,
                    "secondCategoryId": 2,
                    "thirdCategoryId": 3,
                    "isPatentProduct": true,
                    "createDate": "2024-04-20 08:00:00",
                    "modifyDate": "2024-04-20 08:00:00",
                    "isSelect": true,
                    "minOrderQuantity": 1,
                    "sellerDataInfo": {
                        "tradeMedalLevel": "1",
                        "compositeServiceScore": "1",
                        "logisticsExperienceScore": "1",
                        "disputeComplaintScore": "1",
                        "offerExperienceScore": "1",
                        "afterSalesExperienceScore": "1",
                        "consultingExperienceScore": "1",
                        "repeatPurchasePercent": "11",
                        "tpYear": 1
                    },
                    "productSimpleShippingInfo": {
                        "shippingTimeGuarantee": "shipIn24Hours",
                        "perfectFulfillmentRate30d": "100.00%",
                        "pickupWithin24hRate30d": "99.12%",
                        "qualityReturnRate30d": "0.12%",
                        "perfectFulfillmentRate7d": "0.00%",
                        "pickupWithin24hRate7d": "95.12%",
                        "delayedShippingRate7d": "12.34%"
                    },
                    "token": "abc",
                    "promotionURL": "商品详情页链接",
                    "sales7d": "1",
                    "productTradeInfo": {
                        "addCartCount7d": "1",
                        "payBuyerCount7d": "2",
                        "addCartCount30d": "3",
                        "payBuyerCount30d": "4"
                    },
                    "invoiceInfo": {
                        "supportOnlineInvoice": true,
                        "supportFastInvoice": false,
                        "invoiceTypes": [
                            "普票",
                            "专票"
                        ],
                        "taxpayerType": "一般纳税人"
                    }
                }
            ]
        }
    }
}
```
