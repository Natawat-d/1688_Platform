# Multilingual image search

Original name: 多语言图搜  
API: `com.alibaba.fenxiao.crossborder:product.search.imageQuery:1` · Category: Products  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:product.search.imageQuery-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/product.search.imageQuery/{appKey}`  
Requires user authorization (access_token) · Requires signature

Multilingual image search.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `offerQueryParam` | [message:product.search.imageQuery.param.OfferQueryParam](#m-product-search-imagequery-param-offerqueryparam) | yes | {&quot;beginPage&quot;:1,&quot;country&quot;:&quot;en&quot;,&quot;imageAddress&quot;:&quot;https://cbu01.alicdn.com/img/ibank/O1CN01q5lIoD1ZPh3gN3www_!!2928623187-0-cib.jpg&quot;,&quot;pageSize&quot;:1,&quot;userId&quot;:0} | {"beginPage":1,"country":"en","imageAddress":"https://cbu01.alicdn.com/img/ibank/O1CN01q5lIoD1ZPh3gN3www_!!2928623187-0-cib.jpg","pageSize":1,"userId":0} |

<a id="m-product-search-imagequery-param-offerqueryparam"></a>
#### product.search.imageQuery.param.OfferQueryParam

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `imageId` | java.lang.String | yes | Image ID, required | 图片id |
| `beginPage` | java.lang.Integer | yes | Pagination | 分页 |
| `pageSize` | java.lang.Integer | yes | Pagination, maximum 50; 20 is recommended for best results | 分页 |
| `region` | java.lang.String | no | Entity selection | 266,799,48,581 |
| `filter` | java.lang.String | no | Filter parameters, multiple values separated by commas; see the solution introduction for the enumeration | shipInToday,ksCiphertext |
| `sort` | java.lang.String | no | Sorting parameter; for enum values see the solution introduction | {"price":"asc"} |
| `outMemberId` | java.lang.String | no | External user uid | 外部用户uid |
| `priceStart` | java.lang.String | no | Wholesale price start | 10 |
| `priceEnd` | java.lang.String | no | Wholesale price range end | 20 |
| `categoryId` | java.lang.Long | no | Category ID | 类目id |
| `imageAddress` | java.lang.String | no | Image URL; only used in the scenario of querying with a 1688 image link, other cases are not guaranteed to return data | 图片地址 |
| `country` | java.lang.String | yes | Language | 如en-英语，详细枚举请参考开发人员参考菜单 |
| `keyword` | String | no | Search within results | 书本 |
| `auxiliaryText` | String | no | Multimodal image-search copy text | 热卖的 |
| `productCollectionId` | String | no | Xunyuantong workbench assortment ID | 21432232 |
| `keywordTranslate` | Boolean | no | Whether the search term has already been translated; if true, search directly without translating the keyword | false |
| `itemTitle` | String | no | Product title | 商品标题 |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:product.search.imageQuery.result.ResultModelV5](#m-product-search-imagequery-result-resultmodelv5) | yes | Return value | 返回值 |

<a id="m-product-search-imagequery-result-resultmodelv5"></a>
#### product.search.imageQuery.result.ResultModelV5

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `success` | String | yes | Whether successful | 是否成功 |
| `code` | String | yes | code | code |
| `message` | String | yes | message | message |
| `result` | [message:product.search.imageQuery.model.PageInfoV4](#m-product-search-imagequery-model-pageinfov4) | yes | Result | 结果 |

<a id="m-product-search-imagequery-model-pageinfov4"></a>
#### product.search.imageQuery.model.PageInfoV4

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `totalRecords` | Integer | yes | Total quantity | 1 |
| `totalPage` | Integer | yes | Total number of pages | 1 |
| `pageSize` | Integer | yes | Pagination | 1 |
| `currentPage` | Integer | yes | Pagination | 1 |
| `data` | [message:product.search.imageQuery.model.ProductInfoModelV3[]](#m-product-search-imagequery-model-productinfomodelv3[]) | yes | Data | 数据 |
| `picRegionInfo` | [message:com.alibaba.cbu.offer.model.PicRegionInfo](#m-com-alibaba-cbu-offer-model-picregioninfo) | yes | Entity information | {"currentRegion":"265,597,326,764","yoloCropRegion":"265,597,326,764;443,783,154,595"} |

<a id="m-product-search-imagequery-model-productinfomodelv3[]"></a>
#### product.search.imageQuery.model.ProductInfoModelV3[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `imageUrl` | String | yes | Image url | 图片url |
| `subject` | String | yes | Title | 标题 |
| `subjectTrans` | String | yes | Multilingual title | 多语言标题 |
| `priceInfo` | [message:product.search.imageQuery.model.PriceInfoV3](#m-product-search-imagequery-model-priceinfov3) | yes | Price | 价格 |
| `offerId` | Long | yes | Product ID | 商品id |
| `isJxhy` | Boolean | yes | Whether it is curated supply | 是否精选货源 |
| `repurchaseRate` | String | yes | Repurchase rate | 13% |
| `monthSold` | Integer | yes | 30-day sales volume | 1234 |
| `traceInfo` | String | yes | Report tracking data to 1688 | object_id@620201390233^object_type@offer |
| `isOnePsale` | Boolean | yes | Whether it is dropshipping | true |
| `sellerIdentities` | String[] | yes | Merchant identity | super_factory-超级工厂 powerful_merchants-实力商家 tp_member-诚信通会员 |
| `offerIdentities` | String[] | yes | Product tag | yx-严选 |
| `tradeScore` | String | yes | Product transaction rating | 5.0 |
| `promotionModel` | [message:product.search.imageQuery.model.PromotionModelV2](#m-product-search-imagequery-model-promotionmodelv2) | yes | Marketing information | 营销信息 |
| `topCategoryId` | Long | yes | Top-level category | 1 |
| `secondCategoryId` | Long | yes | Secondary category | 2 |
| `thirdCategoryId` | Long | yes | Third-level category | 3 |
| `isPatentProduct` | Boolean | yes | Whether it is a patented product | true |
| `createDate` | String | yes | Product creation time | 2024-04-20 08:00:00 |
| `modifyDate` | String | yes | Product modification time | 2024-04-20 08:00:00 |
| `isSelect` | Boolean | yes | Cross-border select assortment | true |
| `sellerDataInfo` | [message:product.search.imageQuery.model.SellerDataInfoV1](#m-product-search-imagequery-model-sellerdatainfov1) | yes | Product data | 11 |
| `productSimpleShippingInfo` | [message:product.search.imageQuery.model.ProductSimpleShippingInfo](#m-product-search-imagequery-model-productsimpleshippinginfo) | yes | Brief shipping information | 11 |
| `minOrderQuantity` | Integer | yes | Minimum order quantity | 1 |
| `token` | String | yes | Plugin same-store rebate flag | 5L9Z3683O2iVIvwFxNdqnUlMVbKYHWGWcjs |
| `promotionURL` | String | yes | Product promotion link (product detail page) | https://detail.1688.com/offer/710891050473.html?fromkv=refer:HVKTIRCHJVJFKR2RGRLDMTS2KJDUCNCEKNGUUUKHKVMUISKOLJKEYNKRLBATESZXI5CTGVCHJVJFESCBLFCE2T2KKFDTIM2UKE6T |
| `companyName` | String | yes | Shop name | 浙江省一路发发商贸公司 |
| `companyAddress` | String | yes | Store's region | 浙江省杭州市 |
| `offerDataInfo` | [message:com.alibaba.cbu.offer.model.OfferDataInfoV1](#m-com-alibaba-cbu-offer-model-offerdatainfov1) | yes | Information related to the product dimension | xx |
| `sales7d` | String | yes | Sales volume in the last 7 days | 1 |
| `productTradeInfo` | [message:com.alibaba.cbu.offer.model.ProductTradeInfo](#m-com-alibaba-cbu-offer-model-producttradeinfo) | yes | Product sales volume data | 1 |
| `invoiceInfo` | [message:product.search.queryProductDetail.model.InvoiceInfo](#m-product-search-queryproductdetail-model-invoiceinfo) | yes | Invoice information | {         "supportOnlineInvoice": false,         "supportFastInvoice": false,         "invoiceTypes": [           "普票"         ],         "taxpayerType": "一般纳税人"       } |

<a id="m-product-search-imagequery-model-priceinfov3"></a>
#### product.search.imageQuery.model.PriceInfoV3

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `price` | String | yes | Wholesale price | 1 |
| `jxhyPrice` | String | yes | Curated supply price for dropshipping | 1 |
| `pfJxhyPrice` | String | yes | Curated supply wholesale price | 1 |
| `consignPrice` | String | yes | Dropshipping price. When isOnePsale=true, it indicates dropshipping. | 1 |
| `promotionPrice` | String | yes | Marketing price | 1 |

<a id="m-product-search-imagequery-model-promotionmodelv2"></a>
#### product.search.imageQuery.model.PromotionModelV2

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `hasPromotion` | Boolean | yes | Whether there is marketing/promotion | true |
| `promotionType` | String | yes | Marketing type | plus |

<a id="m-product-search-imagequery-model-sellerdatainfov1"></a>
#### product.search.imageQuery.model.SellerDataInfoV1

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

<a id="m-product-search-imagequery-model-productsimpleshippinginfo"></a>
#### product.search.imageQuery.model.ProductSimpleShippingInfo

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `shippingTimeGuarantee` | String | yes | shipIn24Hours - 24-hour shipping shipIn48Hours - 48-hour shipping | shipIn24Hours |
| `perfectFulfillmentRate30d` | String | yes | Perfect fulfillment rate over the last 30 days | 99.12% |
| `pickupWithin24hRate30d` | String | yes | 24-hour pickup rate over the last 30 days | 100% |
| `qualityReturnRate30d` | String | yes | Quality return rate over the last 30 days | 1.23% |
| `perfectFulfillmentRate7d` | String | yes | Perfect fulfillment rate over the last 7 days | 99.12% |
| `pickupWithin24hRate7d` | String | yes | 24-hour pickup rate over the last 7 days | 99.12% |
| `delayedShippingRate7d` | String | yes | Delayed shipment rate over the last 7 days | 12.34% |

<a id="m-com-alibaba-cbu-offer-model-offerdatainfov1"></a>
#### com.alibaba.cbu.offer.model.OfferDataInfoV1

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `offerApplauseRate` | String | yes | Product positive review rate | 100.0% |
| `kjFeaturedServices` | String | yes | Cross-border featured services. When there are multiple services, they are separated by English semicolons. | C |
| `qualityRfdRate30d` | String | yes | Quality refund rate over the last 30 days | 0% |
| `repurchaseRate30d` | String | yes | 30-day repurchase rate | 33.9% |
| `pickupRate24h30d` | String | yes | 24-hour pickup rate | 98% |

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

<a id="m-com-alibaba-cbu-offer-model-picregioninfo"></a>
#### com.alibaba.cbu.offer.model.PicRegionInfo

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `currentRegion` | java.lang.String | yes |  |  |
| `yoloCropRegion` | java.lang.String | yes |  |  |

## Samples

**Input parameter example**

```
{
  "offerQueryParam": {
    "imageId": "图片id",
    "beginPage": 123,
    "pageSize": 123,
    "region": "266,799,48,581",
    "filter": "shipInToday,ksCiphertext",
    "sort": "{\"price\":\"asc\"}",
    "outMemberId": "外部用户uid",
    "priceStart": "10",
    "priceEnd": "20",
    "categoryId": 123,
    "imageAddress": "图片地址",
    "country": "如en-英语，详细枚举请参考开发人员参考菜单",
    "keyword": "书本",
    "auxiliaryText": "热卖的",
    "productCollectionId": "21432232",
    "keywordTranslate": false,
    "itemTitle": "可选传递商品标题，提升图搜效果"
  }
}
```

**Output parameter example**

```
{
  "result": {
    "success": "是否成功",
    "code": "code",
    "message": "message",
    "result": {
      "totalRecords": 1,
      "totalPage": 1,
      "pageSize": 1,
      "currentPage": 1,
      "data": [
        {
          "imageUrl": "图片url",
          "subject": "标题",
          "subjectTrans": "多语言标题",
          "priceInfo": {
            "price": "1",
            "jxhyPrice": "1",
            "pfJxhyPrice": "1",
            "consignPrice": "1",
            "promotionPrice": "1"
          },
          "offerId": 123,
          "isJxhy": true,
          "repurchaseRate": "13%",
          "monthSold": 1234,
          "traceInfo": "object_id@620201390233^object_type@offer",
          "isOnePsale": true,
          "sellerIdentities": [
            "super_factory-超级工厂 powerful_merchants-实力商家 tp_member-诚信通会员"
          ],
          "offerIdentities": [
            "yx-严选"
          ],
          "tradeScore": "5.0",
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
            "perfectFulfillmentRate30d": "99.12%",
            "pickupWithin24hRate30d": "100%",
            "qualityReturnRate30d": "1.23%",
            "perfectFulfillmentRate7d": "99.12%",
            "pickupWithin24hRate7d": "99.12%",
            "delayedShippingRate7d": "12.34%"
          },
          "minOrderQuantity": 1,
          "token": "5L9Z3683O2iVIvwFxNdqnUlMVbKYHWGWcjs",
          "promotionURL": "https://detail.1688.com/offer/710891050473.html?fromkv=refer:HVKTIRCHJVJFKR2RGRLDMTS2KJDUCNCEKNGUUUKHKVMUISKOLJKEYNKRLBATESZXI5CTGVCHJVJFESCBLFCE2T2KKFDTIM2UKE6T",
          "companyName": "浙江省一路发发商贸公司",
          "companyAddress": "浙江省杭州市",
          "offerDataInfo": {
            "offerApplauseRate": "100.0%",
            "kjFeaturedServices": "C",
            "qualityRfdRate30d": "0%",
            "repurchaseRate30d": "33.9%",
            "pickupRate24h30d": "98%"
          },
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
              "[\"普票\",\"专票\"]"
            ],
            "taxpayerType": "一般纳税人 |  小规模纳税人"
          }
        }
      ],
      "picRegionInfo": {
        "currentRegion": "",
        "yoloCropRegion": ""
      }
    }
  }
}
```
