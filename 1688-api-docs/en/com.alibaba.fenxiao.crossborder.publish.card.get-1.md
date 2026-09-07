# Listing product card

Original name: 铺货商品卡片  
API: `com.alibaba.fenxiao.crossborder:publish.card.get:1` · Category: Listing (Publishing)  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:publish.card.get-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/publish.card.get/{appKey}`  
Requires user authorization (access_token) · Requires signature

Get the card data for a single 1688 product, converted to WB format.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `param` | [message:alibaba.global1688.silicon.user.api.param.WbProductCardQueryParam](#m-alibaba-global1688-silicon-user-api-param-wbproductcardqueryparam) | yes | Parameters for querying the product card | {"offerId":1234567890,"country":"ru"} |

<a id="m-alibaba-global1688-silicon-user-api-param-wbproductcardqueryparam"></a>
#### alibaba.global1688.silicon.user.api.param.WbProductCardQueryParam

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `country` | java.lang.String | yes | country field | country_example |
| `offerId` | java.lang.Long | yes | offerId field | 1 |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.global1688.silicon.user.api.result.WbProductCardResponseDTO](#m-alibaba-global1688-silicon-user-api-result-wbproductcardresponsedto) | yes | Product card data return result | {"success":true,"result":{"offerId":1234567890,"subjectID":12345,"subject":"4#131","status":"published","productRating":4.6,"images":["https://example.com/img1.jpg"],"videos":[],"cards":[{"offerId":577263837457,"colorName":"红色","sizes":[{"wbSize":"48","techSize":"XL","skuID":123456,"amount":100,"priceRanges":[{"startQuantity":1,"price":"10.5"}]}]}]},"code":"200","permissionName":"","message":"success"} |

<a id="m-alibaba-global1688-silicon-user-api-result-wbproductcardresponsedto"></a>
#### alibaba.global1688.silicon.user.api.result.WbProductCardResponseDTO

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `offerId` | java.lang.Long | yes | 1688 product ID | 577263837457 |
| `subjectID` | java.lang.Integer | yes | WB product category ID | 6352 |
| `subject` | String | yes | WB category externalCategoryId, in the format parentID#subjectID; top-level or unknown categories degrade to a plain subjectID | 239#2499 |
| `status` | java.lang.String | yes | Product card status | published |
| `productRating` | java.lang.Float | yes | Product rating (0-5) | 4.8 |
| `images` | java.util.List | yes | List of product images | ["https://example.com/img1.jpg"] |
| `videos` | java.util.List | yes | List of product videos | [] |
| `cards` | [message:alibaba.global1688.silicon.user.api.result.WbProductCardDTO[]](#m-alibaba-global1688-silicon-user-api-result-wbproductcarddto[]) | yes | List of product cards (split by color variant) | [{"offerId":577263837457,"colorName":"红色","colorCode":"RED"}] |

<a id="m-alibaba-global1688-silicon-user-api-result-wbproductcarddto[]"></a>
#### alibaba.global1688.silicon.user.api.result.WbProductCardDTO[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `brandName` | java.lang.String | yes | The brandName field | brandName_example |
| `characteristics` | [message:alibaba.global1688.silicon.user.api.result.WbProductCardDTO.Characteristic[]](#m-alibaba-global1688-silicon-user-api-result-wbproductcarddto-characteristic[]) | yes | The characteristics field | {} |
| `description` | java.lang.String | yes | description field | description_example |
| `dimensions` | [message:alibaba.global1688.silicon.user.api.result.WbProductCardDTO.Dimensions](#m-alibaba-global1688-silicon-user-api-result-wbproductcarddto-dimensions) | yes | The dimensions field | {} |
| `isAdult` | java.lang.Boolean | yes | isAdult field | true |
| `images` | java.lang.String[] | no | List of product main images | ["https://img.1688.com/1.jpg"] |
| `videos` | java.lang.String[] | no | List of product videos | [] |
| `sellerInfo` | [message:alibaba.global1688.silicon.user.api.result.WbProductCardDTO.SellerInfo](#m-alibaba-global1688-silicon-user-api-result-wbproductcarddto-sellerinfo) | yes | sellerInfo field | {} |
| `sizes` | [message:alibaba.global1688.silicon.user.api.result.WbProductCardDTO.Size[]](#m-alibaba-global1688-silicon-user-api-result-wbproductcarddto-size[]) | yes | sizes field | {} |
| `title` | java.lang.String | yes | title field | title_example |
| `weight` | [message:alibaba.global1688.silicon.user.api.result.WbProductCardDTO.Weight](#m-alibaba-global1688-silicon-user-api-result-wbproductcarddto-weight) | yes | The weight field | {} |
| `wholesale` | [message:alibaba.global1688.silicon.user.api.result.WbProductCardDTO.Wholesale](#m-alibaba-global1688-silicon-user-api-result-wbproductcarddto-wholesale) | yes | wholesale field | {} |

<a id="m-alibaba-global1688-silicon-user-api-result-wbproductcarddto-characteristic[]"></a>
#### alibaba.global1688.silicon.user.api.result.WbProductCardDTO.Characteristic[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `id` | java.lang.Integer | yes | The id field | 1 |
| `values` | java.lang.String[] | yes | values field | values_example |

<a id="m-alibaba-global1688-silicon-user-api-result-wbproductcarddto-dimensions"></a>
#### alibaba.global1688.silicon.user.api.result.WbProductCardDTO.Dimensions

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `height` | java.math.BigDecimal | yes | The height field | 1.0 |
| `length` | java.math.BigDecimal | yes | length field | 1.0 |
| `unit` | java.lang.String | yes | The unit field | unit_example |
| `width` | java.math.BigDecimal | yes | The width field | 1.0 |

<a id="m-alibaba-global1688-silicon-user-api-result-wbproductcarddto-sellerinfo"></a>
#### alibaba.global1688.silicon.user.api.result.WbProductCardDTO.SellerInfo

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `name` | java.lang.String | yes | The name field | name_example |
| `onTimeShipmentRate` | java.lang.Float | yes | The onTimeShipmentRate field | 1.0 |
| `repeatPurchaseRate` | java.lang.Float | yes | repeatPurchaseRate field | 1.0 |
| `serviceRating` | java.lang.Float | yes | serviceRating field | 1.0 |

<a id="m-alibaba-global1688-silicon-user-api-result-wbproductcarddto-size[]"></a>
#### alibaba.global1688.silicon.user.api.result.WbProductCardDTO.Size[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `amount` | java.lang.Integer | yes | amount field | 1 |
| `isDeleted` | java.lang.Boolean | yes | isDeleted field | true |
| `priceRanges` | [message:alibaba.global1688.silicon.user.api.result.WbProductCardDTO.PriceRange[]](#m-alibaba-global1688-silicon-user-api-result-wbproductcarddto-pricerange[]) | yes | priceRanges field | {} |
| `skuID` | java.lang.Long | yes | The skuID field | 1 |
| `techSize` | java.lang.String | yes | techSize field | techSize_example |
| `wbSize` | String | yes | WB size flag, string; null for sizeless (dimensionless) categories | 48 |

<a id="m-alibaba-global1688-silicon-user-api-result-wbproductcarddto-pricerange[]"></a>
#### alibaba.global1688.silicon.user.api.result.WbProductCardDTO.PriceRange[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `price` | java.lang.String | yes | The price field | price_example |
| `startQuantity` | java.lang.Integer | yes | The startQuantity field | 1 |

<a id="m-alibaba-global1688-silicon-user-api-result-wbproductcarddto-weight"></a>
#### alibaba.global1688.silicon.user.api.result.WbProductCardDTO.Weight

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `gross` | java.math.BigDecimal | yes | gross field | 1.0 |
| `unit` | java.lang.String | yes | The unit field | unit_example |

<a id="m-alibaba-global1688-silicon-user-api-result-wbproductcarddto-wholesale"></a>
#### alibaba.global1688.silicon.user.api.result.WbProductCardDTO.Wholesale

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `minOrderQuantity` | java.lang.Integer | yes | The minOrderQuantity field | 1 |
| `quantum` | java.lang.Integer | yes | quantum field | 1 |

## Samples

**Input parameter example**

```
{
  "param": {
    "country": "country_example",
    "offerId": 1
  }
}
```

**Output parameter example**

```
{
  "result": {
    "code": "code_example",
    "message": "message_example",
    "permissionName": "permissionName_example",
    "result": {
      "brandName": "brandName_example",
      "characteristics": [
        {
          "id": 1,
          "images": [
            "images_example"
          ],
          "values": [
            "values_example"
          ]
        }
      ],
      "description": "description_example",
      "dimensions": {
        "height": "1.0",
        "length": "1.0",
        "unit": "unit_example",
        "width": "1.0"
      },
      "images": [
        "images_example"
      ],
      "isAdult": true,
      "offerId": 1,
      "productRating": 1.0,
      "sellerInfo": [
        {
          "name": "name_example",
          "onTimeShipmentRate": 1.0,
          "positiveReviewRate": 1.0,
          "regPeriod": 1.0,
          "repeatPurchaseRate": 1.0,
          "serviceRating": 1.0,
          "verificationStatus": "verificationStatus_example"
        }
      ],
      "sizes": [
        {
          "amount": 1,
          "isDeleted": true,
          "priceRanges": [
            {
              "price": "price_example",
              "startQuantity": 1
            }
          ],
          "skuID": 1,
          "techSize": "techSize_example",
          "wbSize": 1
        }
      ],
      "status": "status_example",
      "subjectID": 1,
      "title": "title_example",
      "videos": [
        "videos_example"
      ],
      "weight": {
        "gross": "1.0",
        "unit": "unit_example"
      },
      "wholesale": {
        "minOrderQuantity": 1,
        "quantum": 1
      }
    },
    "success": true
  }
}
```
