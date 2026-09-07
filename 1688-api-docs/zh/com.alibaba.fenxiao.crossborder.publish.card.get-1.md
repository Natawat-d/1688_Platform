# 铺货商品卡片

API: `com.alibaba.fenxiao.crossborder:publish.card.get:1` · Category: 铺货  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:publish.card.get-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/publish.card.get/{appKey}`  
需要授权 (access_token) · 需要签名

获取单个 1688 商品已转换为 WB 格式的卡片数据。

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `param` | [message:alibaba.global1688.silicon.user.api.param.WbProductCardQueryParam](#m-alibaba-global1688-silicon-user-api-param-wbproductcardqueryparam) | 是 | 查询商品卡片的参数 | {"offerId":1234567890,"country":"ru"} |

<a id="m-alibaba-global1688-silicon-user-api-param-wbproductcardqueryparam"></a>
#### alibaba.global1688.silicon.user.api.param.WbProductCardQueryParam

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `country` | java.lang.String | 是 | country 字段 | country_example |
| `offerId` | java.lang.Long | 是 | offerId 字段 | 1 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:alibaba.global1688.silicon.user.api.result.WbProductCardResponseDTO](#m-alibaba-global1688-silicon-user-api-result-wbproductcardresponsedto) | 是 | 商品卡片数据返回结果 | {"success":true,"result":{"offerId":1234567890,"subjectID":12345,"subject":"4#131","status":"published","productRating":4.6,"images":["https://example.com/img1.jpg"],"videos":[],"cards":[{"offerId":577263837457,"colorName":"红色","sizes":[{"wbSize":"48","techSize":"XL","skuID":123456,"amount":100,"priceRanges":[{"startQuantity":1,"price":"10.5"}]}]}]},"code":"200","permissionName":"","message":"success"} |

<a id="m-alibaba-global1688-silicon-user-api-result-wbproductcardresponsedto"></a>
#### alibaba.global1688.silicon.user.api.result.WbProductCardResponseDTO

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `offerId` | java.lang.Long | 是 | 1688商品ID | 577263837457 |
| `subjectID` | java.lang.Integer | 是 | WB商品类目ID | 6352 |
| `subject` | String | 是 | WB类目externalCategoryId，格式为parentID#subjectID，顶级或未知类目退化为纯subjectID | 239#2499 |
| `status` | java.lang.String | 是 | 商品卡片状态 | published |
| `productRating` | java.lang.Float | 是 | 商品评分(0-5) | 4.8 |
| `images` | java.util.List | 是 | 商品图片列表 | ["https://example.com/img1.jpg"] |
| `videos` | java.util.List | 是 | 商品视频列表 | [] |
| `cards` | [message:alibaba.global1688.silicon.user.api.result.WbProductCardDTO[]](#m-alibaba-global1688-silicon-user-api-result-wbproductcarddto[]) | 是 | 商品卡片列表(按颜色变体拆分) | [{"offerId":577263837457,"colorName":"红色","colorCode":"RED"}] |

<a id="m-alibaba-global1688-silicon-user-api-result-wbproductcarddto[]"></a>
#### alibaba.global1688.silicon.user.api.result.WbProductCardDTO[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `brandName` | java.lang.String | 是 | brandName 字段 | brandName_example |
| `characteristics` | [message:alibaba.global1688.silicon.user.api.result.WbProductCardDTO.Characteristic[]](#m-alibaba-global1688-silicon-user-api-result-wbproductcarddto-characteristic[]) | 是 | characteristics 字段 | {} |
| `description` | java.lang.String | 是 | description 字段 | description_example |
| `dimensions` | [message:alibaba.global1688.silicon.user.api.result.WbProductCardDTO.Dimensions](#m-alibaba-global1688-silicon-user-api-result-wbproductcarddto-dimensions) | 是 | dimensions 字段 | {} |
| `isAdult` | java.lang.Boolean | 是 | isAdult 字段 | true |
| `images` | java.lang.String[] | 否 | 商品主图列表 | ["https://img.1688.com/1.jpg"] |
| `videos` | java.lang.String[] | 否 | 商品视频列表 | [] |
| `sellerInfo` | [message:alibaba.global1688.silicon.user.api.result.WbProductCardDTO.SellerInfo](#m-alibaba-global1688-silicon-user-api-result-wbproductcarddto-sellerinfo) | 是 | sellerInfo 字段 | {} |
| `sizes` | [message:alibaba.global1688.silicon.user.api.result.WbProductCardDTO.Size[]](#m-alibaba-global1688-silicon-user-api-result-wbproductcarddto-size[]) | 是 | sizes 字段 | {} |
| `title` | java.lang.String | 是 | title 字段 | title_example |
| `weight` | [message:alibaba.global1688.silicon.user.api.result.WbProductCardDTO.Weight](#m-alibaba-global1688-silicon-user-api-result-wbproductcarddto-weight) | 是 | weight 字段 | {} |
| `wholesale` | [message:alibaba.global1688.silicon.user.api.result.WbProductCardDTO.Wholesale](#m-alibaba-global1688-silicon-user-api-result-wbproductcarddto-wholesale) | 是 | wholesale 字段 | {} |

<a id="m-alibaba-global1688-silicon-user-api-result-wbproductcarddto-characteristic[]"></a>
#### alibaba.global1688.silicon.user.api.result.WbProductCardDTO.Characteristic[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `id` | java.lang.Integer | 是 | id 字段 | 1 |
| `values` | java.lang.String[] | 是 | values 字段 | values_example |

<a id="m-alibaba-global1688-silicon-user-api-result-wbproductcarddto-dimensions"></a>
#### alibaba.global1688.silicon.user.api.result.WbProductCardDTO.Dimensions

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `height` | java.math.BigDecimal | 是 | height 字段 | 1.0 |
| `length` | java.math.BigDecimal | 是 | length 字段 | 1.0 |
| `unit` | java.lang.String | 是 | unit 字段 | unit_example |
| `width` | java.math.BigDecimal | 是 | width 字段 | 1.0 |

<a id="m-alibaba-global1688-silicon-user-api-result-wbproductcarddto-sellerinfo"></a>
#### alibaba.global1688.silicon.user.api.result.WbProductCardDTO.SellerInfo

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `name` | java.lang.String | 是 | name 字段 | name_example |
| `onTimeShipmentRate` | java.lang.Float | 是 | onTimeShipmentRate 字段 | 1.0 |
| `repeatPurchaseRate` | java.lang.Float | 是 | repeatPurchaseRate 字段 | 1.0 |
| `serviceRating` | java.lang.Float | 是 | serviceRating 字段 | 1.0 |

<a id="m-alibaba-global1688-silicon-user-api-result-wbproductcarddto-size[]"></a>
#### alibaba.global1688.silicon.user.api.result.WbProductCardDTO.Size[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `amount` | java.lang.Integer | 是 | amount 字段 | 1 |
| `isDeleted` | java.lang.Boolean | 是 | isDeleted 字段 | true |
| `priceRanges` | [message:alibaba.global1688.silicon.user.api.result.WbProductCardDTO.PriceRange[]](#m-alibaba-global1688-silicon-user-api-result-wbproductcarddto-pricerange[]) | 是 | priceRanges 字段 | {} |
| `skuID` | java.lang.Long | 是 | skuID 字段 | 1 |
| `techSize` | java.lang.String | 是 | techSize 字段 | techSize_example |
| `wbSize` | String | 是 | WB 尺码标记，字符串；无尺码（dimensionless）类目为 null | 48 |

<a id="m-alibaba-global1688-silicon-user-api-result-wbproductcarddto-pricerange[]"></a>
#### alibaba.global1688.silicon.user.api.result.WbProductCardDTO.PriceRange[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `price` | java.lang.String | 是 | price 字段 | price_example |
| `startQuantity` | java.lang.Integer | 是 | startQuantity 字段 | 1 |

<a id="m-alibaba-global1688-silicon-user-api-result-wbproductcarddto-weight"></a>
#### alibaba.global1688.silicon.user.api.result.WbProductCardDTO.Weight

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `gross` | java.math.BigDecimal | 是 | gross 字段 | 1.0 |
| `unit` | java.lang.String | 是 | unit 字段 | unit_example |

<a id="m-alibaba-global1688-silicon-user-api-result-wbproductcarddto-wholesale"></a>
#### alibaba.global1688.silicon.user.api.result.WbProductCardDTO.Wholesale

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `minOrderQuantity` | java.lang.Integer | 是 | minOrderQuantity 字段 | 1 |
| `quantum` | java.lang.Integer | 是 | quantum 字段 | 1 |

## 示例

**入参示例**

```
{
  "param": {
    "country": "country_example",
    "offerId": 1
  }
}
```

**出参示例**

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
