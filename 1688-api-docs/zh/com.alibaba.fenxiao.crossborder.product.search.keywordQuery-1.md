# 多语言关键词搜索

API: `com.alibaba.fenxiao.crossborder:product.search.keywordQuery:1` · Category: 商品  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:product.search.keywordQuery-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/product.search.keywordQuery/{appKey}`  
需要授权 (access_token) · 需要签名

多语言关键词搜索

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `offerQueryParam` | [message:product.search.keywordQuery.param.OfferQueryParam](#m-product-search-keywordquery-param-offerqueryparam) | 是 | 查询参数 | {} |

<a id="m-product-search-keywordquery-param-offerqueryparam"></a>
#### product.search.keywordQuery.param.OfferQueryParam

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `keyword` | java.lang.String | 是 | 关键词 | 饼干 |
| `beginPage` | java.lang.Integer | 是 | 分页 | 1 |
| `pageSize` | java.lang.Integer | 是 | 分页 | 1 |
| `filter` | java.lang.String | 否 | 筛选参数，多个通过英文逗号分隔，枚举参见解决方案介绍 | shipInToday,ksCiphertext |
| `sort` | java.lang.String | 否 | 排序参数，枚举参见解决方案介绍 | {"price":"asc"} |
| `outMemberId` | java.lang.String | 否 | 外部用户id | 123 |
| `priceStart` | java.lang.String | 否 | 批发价开始 | 1 |
| `priceEnd` | java.lang.String | 否 | 批发价结束 | 10 |
| `categoryId` | java.lang.Long | 否 | 类目id | 1 |
| `categoryIdList` | String | 否 | 类目id列表，英文逗号隔开，支持多个类目并集 | 2,45 |
| `country` | java.lang.String | 是 | 语言 | 如en-英语，详细枚举参考开发人员参考菜单 |
| `regionOpp` | String | 否 | 商机 | 枚举值见开发人员参考菜单 |
| `productCollectionId` | String | 否 | 寻源通工作台货盘id | 174316138 |
| `snId` | String | 否 | 搜索导航ID，如978或978:1352 | 978:1352 |
| `keywordTranslate` | Boolean | 否 | 关键词是否已翻译，默认未翻译，true的话跳过关键词翻译 | false |
| `saleFilterList` | [message:com.alibaba.cbu.offer.param.SaleFilterParam[]](#m-com-alibaba-cbu-offer-param-salefilterparam[]) | 否 | 销量筛选参数 | [{"saleType":"sales7","saleStart":"10","saleEnd":"100"}] |

<a id="m-com-alibaba-cbu-offer-param-salefilterparam[]"></a>
#### com.alibaba.cbu.offer.param.SaleFilterParam[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `saleType` | String | 否 | 销量类型 | sales7:近7天销量,sales14:近14天销量,sales30:近30天销量,totalSales:总销量，传入多个取交集 |
| `saleStart` | String | 否 | 最小销量 | 10 |
| `saleEnd` | String | 否 | 最大销量 | 100 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:product.search.keywordQuery.result.ResultModelV3](#m-product-search-keywordquery-result-resultmodelv3) | 是 | 返回信息 | 返回信息 |

<a id="m-product-search-keywordquery-result-resultmodelv3"></a>
#### product.search.keywordQuery.result.ResultModelV3

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `success` | Boolean | 是 | 正否正常 | 正否正常 |
| `code` | String | 是 | 状态码 | 状态码 |
| `message` | String | 是 | 提示 | 提示 |
| `result` | [message:product.search.keywordQuery.model.PageInfoV3](#m-product-search-keywordquery-model-pageinfov3) | 是 | 内容 | 内容 |

<a id="m-product-search-keywordquery-model-pageinfov3"></a>
#### product.search.keywordQuery.model.PageInfoV3

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `totalRecords` | Integer | 是 | 总条数 | 分页 |
| `totalPage` | Integer | 是 | 总页码 | 分页 |
| `pageSize` | Integer | 是 | 分页 | 分页 |
| `currentPage` | Integer | 是 | 分页 | 分页 |
| `data` | [message:product.search.keywordQuery.model.ProductInfoModelV2[]](#m-product-search-keywordquery-model-productinfomodelv2[]) | 是 | 数据 | 数据 |

<a id="m-product-search-keywordquery-model-productinfomodelv2[]"></a>
#### product.search.keywordQuery.model.ProductInfoModelV2[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `imageUrl` | String | 是 | 图片地址 | 图片地址 |
| `aigcImageUrl` | String | 是 | 图片地址-aigc处理翻译后 | https:// |
| `subject` | String | 是 | 中文标题 | 中文标题 |
| `subjectTrans` | String | 是 | 外文标题 | 外文标题 |
| `offerId` | Long | 是 | 商品id | 2 |
| `isJxhy` | Boolean | 是 | 是否精选货源 | true |
| `priceInfo` | [message:product.search.keywordQuery.model.PriceInfoV2](#m-product-search-keywordquery-model-priceinfov2) | 是 | 价格 | 1 |
| `repurchaseRate` | String | 是 | 复购率 | 10% |
| `monthSold` | Integer | 是 | 30天销量 | 1213 |
| `traceInfo` | String | 是 | 向1688上报打点数据 | object_id@620201390233^object_type@offer |
| `isOnePsale` | Boolean | 是 | 是否一件代发 | true |
| `sellerIdentities` | String[] | 是 | 商家身份 | super_factory-超级工厂 powerful_merchants-实力商家 tp_member-诚信通会员 |
| `offerIdentities` | String[] | 是 | 商品标 | yx-严选，select-跨境select |
| `tradeScore` | String | 是 | 商品交易评分 | 5.0 |
| `whiteImage` | String | 是 | 商品白底图 | 商品白底图 |
| `promotionModel` | [message:product.search.keywordQuery.model.PromotionModelV2](#m-product-search-keywordquery-model-promotionmodelv2) | 是 | 是否有营销信息 | 目前只透plus |
| `topCategoryId` | Long | 是 | 一级类目 | 1 |
| `secondCategoryId` | Long | 是 | 二级类目 | 2 |
| `thirdCategoryId` | Long | 是 | 三级类目 | 3 |
| `isPatentProduct` | Boolean | 是 | 是否专利商品 | true |
| `createDate` | String | 是 | 商品上架时间 | 2024-04-20 08:00:00 |
| `modifyDate` | String | 是 | 商品修改时间 | 2024-04-20 08:00:00 |
| `isSelect` | Boolean | 是 | 跨境select货盘 | true |
| `minOrderQuantity` | Integer | 是 | 最小起批量 | 1 |
| `sellerDataInfo` | [message:product.search.keywordQuery.model.SellerDataInfoV1](#m-product-search-keywordquery-model-sellerdatainfov1) | 是 | 商品数据 | 1 |
| `productSimpleShippingInfo` | [message:product.search.keywordQuery.model.ProductSimpleShippingInfo](#m-product-search-keywordquery-model-productsimpleshippinginfo) | 是 | 简略发货信息 | 1 |
| `token` | String | 是 | 插件返佣token | abc |
| `promotionURL` | String | 是 | 具有【AI跨境运营助手】模块的1688商品详情页链接 | 商品详情页链接 |
| `sales7d` | String | 是 | 最近7天销量 | 1 |
| `productTradeInfo` | [message:com.alibaba.cbu.offer.model.ProductTradeInfo](#m-com-alibaba-cbu-offer-model-producttradeinfo) | 是 | 商品交易数据 | 1 |
| `invoiceInfo` | [message:product.search.queryProductDetail.model.InvoiceInfo](#m-product-search-queryproductdetail-model-invoiceinfo) | 是 | 发票信息 | "invoiceInfo": {         "supportOnlineInvoice": false,         "supportFastInvoice": false,         "invoiceTypes": [           "普票"         ],         "taxpayerType": "一般纳税人"       } |

<a id="m-product-search-keywordquery-model-priceinfov2"></a>
#### product.search.keywordQuery.model.PriceInfoV2

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `price` | String | 是 | 批发价 | 10 |
| `jxhyPrice` | String | 是 | 代发精选货源价 | 10 |
| `pfJxhyPrice` | String | 是 | 批发精选货源价 | 10 |
| `consignPrice` | String | 是 | 一件代发价 | 10 |
| `promotionPrice` | String | 是 | 营销价 | 10 |

<a id="m-product-search-keywordquery-model-promotionmodelv2"></a>
#### product.search.keywordQuery.model.PromotionModelV2

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `hasPromotion` | Boolean | 是 | 是否有营销 | true |
| `promotionType` | String | 是 | 营销类型 | plus |

<a id="m-product-search-keywordquery-model-sellerdatainfov1"></a>
#### product.search.keywordQuery.model.SellerDataInfoV1

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `tradeMedalLevel` | String | 是 | 卖家交易勋章等级 | 1 |
| `compositeServiceScore` | String | 是 | 综合服务体验分 | 1 |
| `logisticsExperienceScore` | String | 是 | 物流体验分 | 1 |
| `disputeComplaintScore` | String | 是 | 纠纷投诉处理分 | 1 |
| `offerExperienceScore` | String | 是 | 商品体验分 | 1 |
| `afterSalesExperienceScore` | String | 是 | 售后体验分 | 1 |
| `consultingExperienceScore` | String | 是 | 咨询体验分 | 1 |
| `repeatPurchasePercent` | String | 是 | 重复购买率 | 11 |
| `tpYear` | Integer | 是 | 诚信通年限 | 1 |

<a id="m-product-search-keywordquery-model-productsimpleshippinginfo"></a>
#### product.search.keywordQuery.model.ProductSimpleShippingInfo

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `shippingTimeGuarantee` | String | 是 | shipIn24Hours-24小时发货 shipIn48Hours-48小时发货 | shipIn24Hours |
| `perfectFulfillmentRate30d` | String | 是 | 近30天完美履约率,空值代表0 | 100.00% |
| `pickupWithin24hRate30d` | String | 是 | 近30天24h揽收率,空值代表0 | 99.12% |
| `qualityReturnRate30d` | String | 是 | 近30天品质退货率,,空值代表0 | 0.12% |
| `perfectFulfillmentRate7d` | String | 是 | 近7天完美履约率,空值代表0 | 0.00% |
| `pickupWithin24hRate7d` | String | 是 | 近7天24h揽收率,空值代表0 | 95.12% |
| `delayedShippingRate7d` | String | 是 | 近7天延迟发货率,空值代表0 | 12.34% |

<a id="m-com-alibaba-cbu-offer-model-producttradeinfo"></a>
#### com.alibaba.cbu.offer.model.ProductTradeInfo

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `addCartCount7d` | String | 是 | 最近7天加入购物车次数 | 1 |
| `payBuyerCount7d` | String | 是 | 最近7天支付买家数 | 2 |
| `addCartCount30d` | String | 是 | 最近30天加入购物车次数 | 3 |
| `payBuyerCount30d` | String | 是 | 最近30天支付买家数 | 4 |

<a id="m-product-search-queryproductdetail-model-invoiceinfo"></a>
#### product.search.queryProductDetail.model.InvoiceInfo

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `supportOnlineInvoice` | Boolean | 是 | 是否支持在线开票 | true |
| `supportFastInvoice` | Boolean | 是 | 是否支持极速开票 | false |
| `invoiceTypes` | String[] | 是 | 支持的开票类型列表 | ["普票","专票"] |
| `taxpayerType` | String | 是 | 纳税人类型 | 一般纳税人 \|  小规模纳税人 |

## 示例

**入参示例**

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

**出参示例**

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
