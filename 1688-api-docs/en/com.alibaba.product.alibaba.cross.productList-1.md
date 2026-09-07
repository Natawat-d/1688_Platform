# Get product list (cross-border)

Original name: 跨境场景获取商品列表  
API: `com.alibaba.product:alibaba.cross.productList:1` · Category: Product Selection & Listing  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.product:alibaba.cross.productList-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.product/alibaba.cross.productList/{appKey}`  
Requires user authorization (access_token) · Requires signature

This interface validates the cross-border listing relationship and is for cross-border business only. Product model V2.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `productIdList` | Long[] | yes | Product Id list | [574325651942,570027659932] |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `productList` | [message:alibaba.product.ProductInfo[]](#m-alibaba-product-productinfo[]) | yes | Product list | [] |
| `success` | Boolean | yes | Whether successful | true |
| `message` | String | yes | Return message | 57002765XXX:商品不存在 |

<a id="m-alibaba-product-productinfo[]"></a>
#### alibaba.product.ProductInfo[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `productID` | Long | yes | Product ID | 574325651942 |
| `productType` | String | yes | Product type: online wholesale product (wholesale) or inquiry product (sourcing). Defaults to wholesale on the 1688 website. | wholesale |
| `attributes` | [message:alibaba.product.ProductAttribute[]](#m-alibaba-product-productattribute[]) | yes | Product attributes and attribute values |   |
| `groupID` | Long[] | yes | Group ID, determines the group the product belongs to. On 1688, multiple group IDs can be passed in; on the international site, a product can belong to only one group, so by default only the first one is taken. |   |
| `status` | String | yes | Product status. published: online status; member expired: revoked by member; auto expired: naturally expired; expired: expired (includes both manually and automatically expired); member deleted: deleted by member; modified: modified; new: newly published; deleted: deleted; TBD: to be delete; approved: approved; auditing: under review; untread: review not passed; | published |
| `subject` | String | yes | Product title, up to 128 characters | 吉米兔小时候 宝宝成长纪念册相册配套 使用空白记录彩色加页1 |
| `description` | String | yes | Product detail description, may include image URLs from the image center |   |
| `language` | String | yes | Language; see the FAQ for language enum values. The 1688 website passes CHINESE by default | CHINESE |
| `periodOfValidity` | Integer | yes | Information validity period, calculated in days; not applicable for the international site | 3650 |
| `bizType` | Integer | yes | Business type. 1: Product, 2: Processing, 3: Agency, 4: Cooperation, 5: Business service. The international site defaults to Product. | 1 |
| `image` | [message:alibaba.product.ProductImageInfo](#m-alibaba-product-productimageinfo) | yes | Product main image | {} |
| `extendInfos` | [message:alibaba.product.ProductExtendInfo[]](#m-alibaba-product-productextendinfo[]) | yes | Product extension info |   |
| `supplierLoginId` | java.lang.String | yes | Supplier loginId | alitestforisv02 |
| `categoryID` | Long | yes | Category ID, identifies the category the product belongs to | 1048182 |
| `categoryName` | java.lang.String | yes | Category name |   |
| `productCargoNumber` | java.lang.String | yes | Product model/article number, the article number in product attributes |   |

<a id="m-alibaba-product-productattribute[]"></a>
#### alibaba.product.ProductAttribute[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `attributeID` | Long | yes | Attribute ID | 123456 |
| `attributeName` | String | yes | Attribute name | color |
| `valueID` | Long | yes | Attribute value ID | 123456 |
| `value` | String | yes | Attribute value | grey |
| `isCustom` | Boolean | yes | Whether it is a custom attribute; not applicable to the international site | true |

<a id="m-alibaba-product-productimageinfo"></a>
#### alibaba.product.ProductImageInfo

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `images` | String[] | yes | List of main images. The image upload API must be used to upload images first. | ["img/ibank/2018/502/115/9153511205_1606139362.jpg"] |
| `isWatermark` | Boolean | yes | Whether to add a watermark, yes (true) or no (false). 1688 does not need to be concerned with this field; 1688's watermark information is handled when the image is uploaded. |   |
| `isWatermarkFrame` | Boolean | yes | Whether the watermark has a border, with border (true) or without border (false). 1688 does not need to worry about this field; 1688's watermark information is processed when the image is uploaded |   |
| `watermarkPosition` | String | yes | Watermark position, either center or bottom. 1688 does not need to use this field; 1688's watermark information is processed when the image is uploaded |   |

<a id="m-alibaba-product-productextendinfo[]"></a>
#### alibaba.product.ProductExtendInfo[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `key` | String | yes | key of the extension structure | 代销价格,consignPrice;<br>买家保障,buyerProtection; |
| `value` | String | yes | value of the extension structure | 代销价格,key为skuId，value为用户设置的代销价，<br>示例：31151771910:2088.0;31151771909:2088.0;31151771908:2088.0;31152339121:2088.0;<br>买家保障,string数组，value为买保全拼，<br>示例：["psbj","swtwlybt","swtbh","ssbxsfh"] |

## Error codes

| Error | Symptom | How to fix |
|---|---|---|
| Product does not exist | 574325651XXX: product does not exist, 57002765XXX: product does not exist | Check whether this product exists in the listing list |

## Samples

**Input parameter example**

```
{
  "productIdList":[574325651942,570027659932]
}
```

**Output parameter example**

```
{
    "productList":[
        {
            "productID":574325651942,
            "productType":"wholesale",
            "status":"published",
            "subject":"吉米兔小时候 宝宝成长纪念册相册配套 使用空白记录彩色加页1",
            "language":"CHINESE",
            "periodOfValidity":3650,
            "bizType":1,
            "image":{
                "images":[
                    "img/ibank/2018/502/115/9153511205_1606139362.jpg",
                    "img/ibank/2018/625/691/9172196526_1606139362.jpg",
                    "img/ibank/2018/731/022/9172220137_1606139362.jpg",
                    "img/ibank/2018/525/691/9172196525_1606139362.jpg"
                ]
            },
            "supplierLoginId":"alitestforisv02"
        },
        {
            "productID":570027659932,
            "productType":"wholesale",
            "status":"published",
            "subject":"雨发给ad特价在特卖特卖特价狂欢价新款波点（送腰带）v",
            "language":"CHINESE",
            "periodOfValidity":3650,
            "bizType":1,
            "image":{
                "images":[
                    "img/ibank/2018/178/658/8855856871_1630100306.jpg",
                    "img/ibank/2018/854/635/8892536458_1630100306.jpg",
                    "img/ibank/2018/638/958/8855859836_1630100306.jpg",
                    "img/ibank/2018/170/298/8855892071_1630100306.jpg",
                    "img/ibank/2018/139/448/8855844931_1630100306.jpg"
                ]
            },
            "supplierLoginId":"alitestforisv02"
        }
    ],
    "success":true
}
```
