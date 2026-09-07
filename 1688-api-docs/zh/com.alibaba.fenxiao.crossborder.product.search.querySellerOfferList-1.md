# 多语言商品店搜

API: `com.alibaba.fenxiao.crossborder:product.search.querySellerOfferList:1` · Category: 商品  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:product.search.querySellerOfferList-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/product.search.querySellerOfferList/{appKey}`  
需要授权 (access_token) · 需要签名

多语言商品店搜

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `offerQueryParam` | [message:product.search.querySellerOfferList.param.OfferQueryParam](#m-product-search-querysellerofferlist-param-offerqueryparam) | 是 | 请求参数 | {} |

<a id="m-product-search-querysellerofferlist-param-offerqueryparam"></a>
#### product.search.querySellerOfferList.param.OfferQueryParam

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `keyword` | String | 否 | 关键词 | 饼干 |
| `beginPage` | Integer | 是 | 分页 | 1 |
| `pageSize` | Integer | 是 | 分页 | 1 |
| `filter` | String | 否 | 筛选参数，多个通过英文逗号分隔，枚举参见解决方案介绍 | shipInToday,ksCiphertext |
| `sort` | String | 否 | 排序参数，枚举参见解决方案介绍 | {"price":"asc"} |
| `outMemberId` | String | 否 | 外部用户id | 123 |
| `priceStart` | String | 否 | 批发价开始 | 1 |
| `priceEnd` | String | 否 | 批发价结束 | 1 |
| `categoryId` | Long | 否 | 类目id | 1 |
| `country` | String | 是 | 城市 | japan |
| `sellerOpenId` | String | 是 | 商家店铺id脱敏 | 123 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:product.search.querySellerOfferList.result.ResultModelV3](#m-product-search-querysellerofferlist-result-resultmodelv3) | 是 | 返回信息 | 返回信息 |

<a id="m-product-search-querysellerofferlist-result-resultmodelv3"></a>
#### product.search.querySellerOfferList.result.ResultModelV3

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `success` | Boolean | 是 | 正否正常 | 正否正常 |
| `code` | String | 是 | 状态码 | 状态码 |
| `message` | String | 是 | 提示 | 提示 |
| `result` | [message:product.search.querySellerOfferList.model.PageInfoV3](#m-product-search-querysellerofferlist-model-pageinfov3) | 是 | 内容 | 内容 |

<a id="m-product-search-querysellerofferlist-model-pageinfov3"></a>
#### product.search.querySellerOfferList.model.PageInfoV3

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `totalRecords` | Integer | 是 | 总条数 | 分页 |
| `totalPage` | Integer | 是 | 总页码 | 分页 |
| `pageSize` | Integer | 是 | 分页 | 分页 |
| `currentPage` | Integer | 是 | 分页 | 分页 |
| `data` | [message:product.search.querySellerOfferList.model.ProductInfoModelV2[]](#m-product-search-querysellerofferlist-model-productinfomodelv2[]) | 是 | 数据 | 数据 |

<a id="m-product-search-querysellerofferlist-model-productinfomodelv2[]"></a>
#### product.search.querySellerOfferList.model.ProductInfoModelV2[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `imageUrl` | String | 是 | 图片地址 | 图片地址 |
| `aigcImageUrl` | String | 是 | 图片地址-aigc处理翻译后 | https:// |
| `subject` | String | 是 | 中文标题 | 中文标题 |
| `subjectTrans` | String | 是 | 外文标题 | 外文标题 |
| `offerId` | Long | 是 | 商品id | 2 |
| `isJxhy` | Boolean | 是 | 是否精选货源 | true |
| `repurchaseRate` | String | 是 | 复购率 | 10% |
| `monthSold` | Integer | 是 | 30天销量 | 1213 |
| `traceInfo` | String | 是 | 向1688上报打点数据 | object_id@620201390233^object_type@offer |
| `isOnePsale` | Boolean | 是 | 是否一件代发 | true |
| `priceInfo` | [message:product.search.querySellerOfferList.model.PriceInfoV2](#m-product-search-querysellerofferlist-model-priceinfov2) | 是 | 价格 | 1 |
| `createDate` | String | 是 | 商品创建时间 | 2021-04-08 08:00:00 |
| `modifyDate` | String | 是 | 商品修改时间 | 2021-04-08 08:00:00 |
| `isPatentProduct` | Boolean | 是 | 是否为专利商品 | true |
| `offerIdentities` | String[] | 是 | 商品标 | select-跨境select |
| `isSelect` | String | 是 | 跨境select货盘 | true |
| `token` | String | 是 | 插件返佣token | abc |
| `promotionURL` | String | 是 | 具有【AI跨境运营助手】模块的1688商品详情页链接 | 商品详情页链接 |

<a id="m-product-search-querysellerofferlist-model-priceinfov2"></a>
#### product.search.querySellerOfferList.model.PriceInfoV2

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `price` | String | 是 | 批发价 | 10 |
| `jxhyPrice` | String | 是 | 代发精选货源价 | 10 |
| `pfJxhyPrice` | String | 是 | 批发精选货源价 | 10 |
| `consignPrice` | String | 是 | 一件代发价 | 10 |

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
    "priceEnd": "1",
    "categoryId": 1,
    "country": "japan",
    "sellerOpenId": "123"
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
          "repurchaseRate": "10%",
          "monthSold": 1213,
          "traceInfo": "object_id@620201390233^object_type@offer",
          "isOnePsale": true,
          "priceInfo": {
            "price": "10",
            "jxhyPrice": "10",
            "pfJxhyPrice": "10",
            "consignPrice": "10"
          },
          "createDate": "2021-04-08 08:00:00",
          "modifyDate": "2021-04-08 08:00:00",
          "isPatentProduct": true,
          "offerIdentities": [
            "select-跨境select"
          ],
          "isSelect": "true",
          "token": "abc",
          "promotionURL": "商品详情页链接"
        }
      ]
    }
  }
}
```
