# Recommend products by keyword (cross-border)

Original name: 跨境场景根据关键字推荐商品  
API: `com.alibaba.product:alibaba.product.suggest.crossBorder:1` · Category: Products  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.product:alibaba.product.suggest.crossBorder-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.product/alibaba.product.suggest.crossBorder/{appKey}`  
Requires user authorization (access_token) · Requires signature

Recommend products by keyword and category in cross-border scenarios, sorted by sales volume. Note: this API is rate-limited and is only suitable for manual-association scenarios.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `keyWord` | String | yes | Search keyword for the product, usually the product title | 商品标题 |
| `loginId` | String | no | Seller loginId | alitestforisv01 |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `resultList` | [message:alibaba.search.ProductSearchResultInfo[]](#m-alibaba-search-productsearchresultinfo[]) | yes | Search return results | [] |
| `success` | String | yes | Whether successful | true |
| `errorMsg` | String | yes | Error description |   |
| `errorCode` | String | yes | Error code |   |

<a id="m-alibaba-search-productsearchresultinfo[]"></a>
#### alibaba.search.ProductSearchResultInfo[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `amountOnSale` | Integer | yes | Available quantity for sale | 100 |
| `minPurchaseQuantity` | Double | yes | Minimum order quantity | 3 |
| `picUrl` | String | yes | Product image URL | img/ibank/2018/794/316/9422613497_991974782.jpg |
| `price` | Double | yes | Reference product price | 100 |
| `productID` | Long | yes | Product ID | 1123123331 |
| `bookedCount` | Double | yes | The number of times the product has been sold (by order count) | 2123 |
| `saleQuantity` | Double | yes | How many units of this product have been sold (measured in product units) | 1 |
| `province` | String | yes | Product shipping province code | 浙江 |
| `city` | String | yes | Product shipping city | 杭州 |
| `retailPrice` | Double | yes | Suggested retail price | 100 |
| `subject` | String | yes | Product title | 【原D现货】韩版新秋女裙 时尚方格<font color=red>小</font>立领不规则下摆气质<font color=red>连衣裙</font> |
| `unit` | String | yes | Product unit | 件 |
| `skuList` | [message:alibaba.simple.sku[]](#m-alibaba-simple-sku[]) | yes | SKU information | [] |

<a id="m-alibaba-simple-sku[]"></a>
#### alibaba.simple.sku[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `description` | String | yes | SKU description | 颜色:红色;尺码:L |
| `amountOnSale` | Integer | yes | Available quantity for sale | 100 |
| `skuId` | Long | yes | skuID, globally unique identifier | 3508426014362 |
| `specId` | String | yes | specID, unique within a product; may repeat across different products. | 8d28b045489c250b69870da3b7c71b1d |

## Samples

**Example of returned result**

```
{
  "resultList": [
    {
      "amountOnSale": 241682,
      "city": "义乌市",
      "minPurchaseQuantity": 100,
      "picUrl": "https://cbu01.alicdn.com/img/ibank/2015/112/013/2204310211_1220361846.jpg",
      "price": 0.06,
      "productID": 45636254415,
      "province": "浙江",
      "retailPrice": 0.5,
      "saleQuantity": 159821,
      "subject": "创意儿童戒指批发 糖果色塑料/树脂戒指混批时尚女童女孩戒指批发",
      "unit": "PCS",
      "skuList": [
        {
          "description": "颜色:黄色",
          "amountOnSale": 36866,
          "skuId": 92418963545,
          "specId": "36c58a69219820709b9475e70bd789d4"
        }
      ]
    }
  ],
  "success": "true"
}
```
