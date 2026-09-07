# Multilingual in-store product search

Original name: 多语言商品店搜  
API: `com.alibaba.fenxiao.crossborder:product.search.querySellerOfferList:1` · Category: Products  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:product.search.querySellerOfferList-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/product.search.querySellerOfferList/{appKey}`  
Requires user authorization (access_token) · Requires signature

Multilingual search of the products in a seller's store.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `offerQueryParam` | [message:product.search.querySellerOfferList.param.OfferQueryParam](#m-product-search-querysellerofferlist-param-offerqueryparam) | yes | Request parameters | {} |

<a id="m-product-search-querysellerofferlist-param-offerqueryparam"></a>
#### product.search.querySellerOfferList.param.OfferQueryParam

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `keyword` | String | no | Keyword | 饼干 |
| `beginPage` | Integer | yes | Pagination | 1 |
| `pageSize` | Integer | yes | Pagination | 1 |
| `filter` | String | no | Filter parameters, multiple values separated by commas; see the solution introduction for the enumeration | shipInToday,ksCiphertext |
| `sort` | String | no | Sorting parameter; for enum values see the solution introduction | {"price":"asc"} |
| `outMemberId` | String | no | External user ID | 123 |
| `priceStart` | String | no | Wholesale price start | 1 |
| `priceEnd` | String | no | Wholesale price range end | 1 |
| `categoryId` | Long | no | Category ID | 1 |
| `country` | String | yes | City | japan |
| `sellerOpenId` | String | yes | Masked merchant store ID | 123 |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:product.search.querySellerOfferList.result.ResultModelV3](#m-product-search-querysellerofferlist-result-resultmodelv3) | yes | Return message | 返回信息 |

<a id="m-product-search-querysellerofferlist-result-resultmodelv3"></a>
#### product.search.querySellerOfferList.result.ResultModelV3

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `success` | Boolean | yes | Whether normal | 正否正常 |
| `code` | String | yes | Status code | 状态码 |
| `message` | String | yes | Prompt | 提示 |
| `result` | [message:product.search.querySellerOfferList.model.PageInfoV3](#m-product-search-querysellerofferlist-model-pageinfov3) | yes | Content | 内容 |

<a id="m-product-search-querysellerofferlist-model-pageinfov3"></a>
#### product.search.querySellerOfferList.model.PageInfoV3

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `totalRecords` | Integer | yes | Total count | 分页 |
| `totalPage` | Integer | yes | Total page count | 分页 |
| `pageSize` | Integer | yes | Pagination | 分页 |
| `currentPage` | Integer | yes | Pagination | 分页 |
| `data` | [message:product.search.querySellerOfferList.model.ProductInfoModelV2[]](#m-product-search-querysellerofferlist-model-productinfomodelv2[]) | yes | Data | 数据 |

<a id="m-product-search-querysellerofferlist-model-productinfomodelv2[]"></a>
#### product.search.querySellerOfferList.model.ProductInfoModelV2[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `imageUrl` | String | yes | Image URL | 图片地址 |
| `aigcImageUrl` | String | yes | Image URL - after AIGC processing and translation | https:// |
| `subject` | String | yes | Chinese title | 中文标题 |
| `subjectTrans` | String | yes | Foreign-language title | 外文标题 |
| `offerId` | Long | yes | Product ID | 2 |
| `isJxhy` | Boolean | yes | Whether it is curated supply | true |
| `repurchaseRate` | String | yes | Repurchase rate | 10% |
| `monthSold` | Integer | yes | 30-day sales volume | 1213 |
| `traceInfo` | String | yes | Report tracking data to 1688 | object_id@620201390233^object_type@offer |
| `isOnePsale` | Boolean | yes | Whether it is dropshipping | true |
| `priceInfo` | [message:product.search.querySellerOfferList.model.PriceInfoV2](#m-product-search-querysellerofferlist-model-priceinfov2) | yes | Price | 1 |
| `createDate` | String | yes | Product creation time | 2021-04-08 08:00:00 |
| `modifyDate` | String | yes | Product modification time | 2021-04-08 08:00:00 |
| `isPatentProduct` | Boolean | yes | Whether it is a patented product | true |
| `offerIdentities` | String[] | yes | Product tag | select-跨境select |
| `isSelect` | String | yes | Cross-border select assortment | true |
| `token` | String | yes | Plugin rebate token | abc |
| `promotionURL` | String | yes | Link to the 1688 product detail page with the [AI Cross-Border Operations Assistant] module | 商品详情页链接 |

<a id="m-product-search-querysellerofferlist-model-priceinfov2"></a>
#### product.search.querySellerOfferList.model.PriceInfoV2

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `price` | String | yes | Wholesale price | 10 |
| `jxhyPrice` | String | yes | Curated supply price for dropshipping | 10 |
| `pfJxhyPrice` | String | yes | Curated supply wholesale price | 10 |
| `consignPrice` | String | yes | Dropshipping price | 10 |

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
    "priceEnd": "1",
    "categoryId": 1,
    "country": "japan",
    "sellerOpenId": "123"
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
