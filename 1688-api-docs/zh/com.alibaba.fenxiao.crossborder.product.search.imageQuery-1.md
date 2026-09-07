# 多语言图搜

API: `com.alibaba.fenxiao.crossborder:product.search.imageQuery:1` · Category: 商品  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:product.search.imageQuery-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/product.search.imageQuery/{appKey}`  
需要授权 (access_token) · 需要签名

多语言图搜

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `offerQueryParam` | [message:product.search.imageQuery.param.OfferQueryParam](#m-product-search-imagequery-param-offerqueryparam) | 是 | {&quot;beginPage&quot;:1,&quot;country&quot;:&quot;en&quot;,&quot;imageAddress&quot;:&quot;https://cbu01.alicdn.com/img/ibank/O1CN01q5lIoD1ZPh3gN3www_!!2928623187-0-cib.jpg&quot;,&quot;pageSize&quot;:1,&quot;userId&quot;:0} | {"beginPage":1,"country":"en","imageAddress":"https://cbu01.alicdn.com/img/ibank/O1CN01q5lIoD1ZPh3gN3www_!!2928623187-0-cib.jpg","pageSize":1,"userId":0} |

<a id="m-product-search-imagequery-param-offerqueryparam"></a>
#### product.search.imageQuery.param.OfferQueryParam

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `imageId` | java.lang.String | 是 | 图片id，必传 | 图片id |
| `beginPage` | java.lang.Integer | 是 | 分页 | 分页 |
| `pageSize` | java.lang.Integer | 是 | 分页，最大不超过50，建议20效果最佳 | 分页 |
| `region` | java.lang.String | 否 | 主体选择 | 266,799,48,581 |
| `filter` | java.lang.String | 否 | 筛选参数，多个通过英文逗号分隔，枚举参见解决方案介绍 | shipInToday,ksCiphertext |
| `sort` | java.lang.String | 否 | 排序参数，枚举参见解决方案介绍 | {"price":"asc"} |
| `outMemberId` | java.lang.String | 否 | 外部用户uid | 外部用户uid |
| `priceStart` | java.lang.String | 否 | 批发价开始 | 10 |
| `priceEnd` | java.lang.String | 否 | 批发价结束 | 20 |
| `categoryId` | java.lang.Long | 否 | 类目id | 类目id |
| `imageAddress` | java.lang.String | 否 | 图片地址，仅使用1688图片链接查询场景，其他不保证有数据返回 | 图片地址 |
| `country` | java.lang.String | 是 | 语言 | 如en-英语，详细枚举请参考开发人员参考菜单 |
| `keyword` | String | 否 | 在结果中搜索 | 书本 |
| `auxiliaryText` | String | 否 | 多模态图搜文案 | 热卖的 |
| `productCollectionId` | String | 否 | 寻源通工作台货盘ID | 21432232 |
| `keywordTranslate` | Boolean | 否 | 搜索词是否已经翻译，true的话直接搜索，不翻译关键词 | false |
| `itemTitle` | String | 否 | 商品标题 | 商品标题 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:product.search.imageQuery.result.ResultModelV5](#m-product-search-imagequery-result-resultmodelv5) | 是 | 返回值 | 返回值 |

<a id="m-product-search-imagequery-result-resultmodelv5"></a>
#### product.search.imageQuery.result.ResultModelV5

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `success` | String | 是 | 是否成功 | 是否成功 |
| `code` | String | 是 | code | code |
| `message` | String | 是 | message | message |
| `result` | [message:product.search.imageQuery.model.PageInfoV4](#m-product-search-imagequery-model-pageinfov4) | 是 | 结果 | 结果 |

<a id="m-product-search-imagequery-model-pageinfov4"></a>
#### product.search.imageQuery.model.PageInfoV4

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `totalRecords` | Integer | 是 | 总数量 | 1 |
| `totalPage` | Integer | 是 | 总页数 | 1 |
| `pageSize` | Integer | 是 | 分页 | 1 |
| `currentPage` | Integer | 是 | 分页 | 1 |
| `data` | [message:product.search.imageQuery.model.ProductInfoModelV3[]](#m-product-search-imagequery-model-productinfomodelv3[]) | 是 | 数据 | 数据 |
| `picRegionInfo` | [message:com.alibaba.cbu.offer.model.PicRegionInfo](#m-com-alibaba-cbu-offer-model-picregioninfo) | 是 | 主体信息 | {"currentRegion":"265,597,326,764","yoloCropRegion":"265,597,326,764;443,783,154,595"} |

<a id="m-product-search-imagequery-model-productinfomodelv3[]"></a>
#### product.search.imageQuery.model.ProductInfoModelV3[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `imageUrl` | String | 是 | 图片url | 图片url |
| `subject` | String | 是 | 标题 | 标题 |
| `subjectTrans` | String | 是 | 多语言标题 | 多语言标题 |
| `priceInfo` | [message:product.search.imageQuery.model.PriceInfoV3](#m-product-search-imagequery-model-priceinfov3) | 是 | 价格 | 价格 |
| `offerId` | Long | 是 | 商品id | 商品id |
| `isJxhy` | Boolean | 是 | 是否精选货源 | 是否精选货源 |
| `repurchaseRate` | String | 是 | 复购率 | 13% |
| `monthSold` | Integer | 是 | 30天销量 | 1234 |
| `traceInfo` | String | 是 | 向1688上报打点数据 | object_id@620201390233^object_type@offer |
| `isOnePsale` | Boolean | 是 | 是否一件代发 | true |
| `sellerIdentities` | String[] | 是 | 商家身份 | super_factory-超级工厂 powerful_merchants-实力商家 tp_member-诚信通会员 |
| `offerIdentities` | String[] | 是 | 商品标 | yx-严选 |
| `tradeScore` | String | 是 | 商品交易评分 | 5.0 |
| `promotionModel` | [message:product.search.imageQuery.model.PromotionModelV2](#m-product-search-imagequery-model-promotionmodelv2) | 是 | 营销信息 | 营销信息 |
| `topCategoryId` | Long | 是 | 一级类目 | 1 |
| `secondCategoryId` | Long | 是 | 二级类目 | 2 |
| `thirdCategoryId` | Long | 是 | 三级类目 | 3 |
| `isPatentProduct` | Boolean | 是 | 是否为专利商品 | true |
| `createDate` | String | 是 | 商品创建时间 | 2024-04-20 08:00:00 |
| `modifyDate` | String | 是 | 商品修改时间 | 2024-04-20 08:00:00 |
| `isSelect` | Boolean | 是 | 跨境select货盘 | true |
| `sellerDataInfo` | [message:product.search.imageQuery.model.SellerDataInfoV1](#m-product-search-imagequery-model-sellerdatainfov1) | 是 | 商品数据 | 11 |
| `productSimpleShippingInfo` | [message:product.search.imageQuery.model.ProductSimpleShippingInfo](#m-product-search-imagequery-model-productsimpleshippinginfo) | 是 | 简略发货信息 | 11 |
| `minOrderQuantity` | Integer | 是 | 最小起批量 | 1 |
| `token` | String | 是 | 插件同店返佣标识 | 5L9Z3683O2iVIvwFxNdqnUlMVbKYHWGWcjs |
| `promotionURL` | String | 是 | 商品推广链接（商品详情页） | https://detail.1688.com/offer/710891050473.html?fromkv=refer:HVKTIRCHJVJFKR2RGRLDMTS2KJDUCNCEKNGUUUKHKVMUISKOLJKEYNKRLBATESZXI5CTGVCHJVJFESCBLFCE2T2KKFDTIM2UKE6T |
| `companyName` | String | 是 | 店铺名称 | 浙江省一路发发商贸公司 |
| `companyAddress` | String | 是 | 店铺所在地区 | 浙江省杭州市 |
| `offerDataInfo` | [message:com.alibaba.cbu.offer.model.OfferDataInfoV1](#m-com-alibaba-cbu-offer-model-offerdatainfov1) | 是 | 商品维度相关信息 | xx |
| `sales7d` | String | 是 | 最近7天销量 | 1 |
| `productTradeInfo` | [message:com.alibaba.cbu.offer.model.ProductTradeInfo](#m-com-alibaba-cbu-offer-model-producttradeinfo) | 是 | 商品销量数据 | 1 |
| `invoiceInfo` | [message:product.search.queryProductDetail.model.InvoiceInfo](#m-product-search-queryproductdetail-model-invoiceinfo) | 是 | 发票信息 | {         "supportOnlineInvoice": false,         "supportFastInvoice": false,         "invoiceTypes": [           "普票"         ],         "taxpayerType": "一般纳税人"       } |

<a id="m-product-search-imagequery-model-priceinfov3"></a>
#### product.search.imageQuery.model.PriceInfoV3

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `price` | String | 是 | 批发价 | 1 |
| `jxhyPrice` | String | 是 | 代发精选货源价 | 1 |
| `pfJxhyPrice` | String | 是 | 批发精选货源价 | 1 |
| `consignPrice` | String | 是 | 一件代发价，当isOnePsale=true表示是一件发代发 | 1 |
| `promotionPrice` | String | 是 | 营销价 | 1 |

<a id="m-product-search-imagequery-model-promotionmodelv2"></a>
#### product.search.imageQuery.model.PromotionModelV2

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `hasPromotion` | Boolean | 是 | 是否有营销 | true |
| `promotionType` | String | 是 | 营销类型 | plus |

<a id="m-product-search-imagequery-model-sellerdatainfov1"></a>
#### product.search.imageQuery.model.SellerDataInfoV1

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

<a id="m-product-search-imagequery-model-productsimpleshippinginfo"></a>
#### product.search.imageQuery.model.ProductSimpleShippingInfo

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `shippingTimeGuarantee` | String | 是 | shipIn24Hours-24小时发货 shipIn48Hours-48小时发货 | shipIn24Hours |
| `perfectFulfillmentRate30d` | String | 是 | 近30天完美履约率 | 99.12% |
| `pickupWithin24hRate30d` | String | 是 | 近30天24h揽收率 | 100% |
| `qualityReturnRate30d` | String | 是 | 近30天品质退货率 | 1.23% |
| `perfectFulfillmentRate7d` | String | 是 | 近7天完美履约率 | 99.12% |
| `pickupWithin24hRate7d` | String | 是 | 近7天24h揽收率 | 99.12% |
| `delayedShippingRate7d` | String | 是 | 近7天延迟发货率 | 12.34% |

<a id="m-com-alibaba-cbu-offer-model-offerdatainfov1"></a>
#### com.alibaba.cbu.offer.model.OfferDataInfoV1

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `offerApplauseRate` | String | 是 | 商品好评率 | 100.0% |
| `kjFeaturedServices` | String | 是 | 跨境特色服务，如果存在多个服务的时候，将会由英文分号分隔 | C |
| `qualityRfdRate30d` | String | 是 | 近30天品质退款率 | 0% |
| `repurchaseRate30d` | String | 是 | 30天复购率 | 33.9% |
| `pickupRate24h30d` | String | 是 | 24小时揽收率 | 98% |

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

<a id="m-com-alibaba-cbu-offer-model-picregioninfo"></a>
#### com.alibaba.cbu.offer.model.PicRegionInfo

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `currentRegion` | java.lang.String | 是 |  |  |
| `yoloCropRegion` | java.lang.String | 是 |  |  |

## 示例

**入参示例**

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

**出参示例**

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
