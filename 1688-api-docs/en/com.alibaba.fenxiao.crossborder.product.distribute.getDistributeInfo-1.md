# Get product selling points for listing

Original name: 获取商品铺货卖点  
API: `com.alibaba.fenxiao.crossborder:product.distribute.getDistributeInfo:1` · Category: Products  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:product.distribute.getDistributeInfo-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/product.distribute.getDistributeInfo/{appKey}`  
No user authorization · Requires signature

For listing scenarios: get the selling-point information needed to list a product.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | no | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `offerId` | java.lang.String | yes | offerId | 1 |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:product.distribute.getDistributeInfo.Result](#m-product-distribute-getdistributeinfo-result) | yes | Return result | {} |

<a id="m-product-distribute-getdistributeinfo-result"></a>
#### product.distribute.getDistributeInfo.Result

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `success` | java.lang.Boolean | yes | Whether the request was successful | true |
| `msg` | java.lang.String | yes | Error message | 系统异常 |
| `data` | [message:product.distribute.getDistributeInfo.ItemAIGCVO](#m-product-distribute-getdistributeinfo-itemaigcvo) | yes | Specific business data | {"keywords":{"ru":[],"en":["Distinctive","Comfortable","Trendy","Quality","Versatile","Unique"],"zh":null},"detailText":[],"manufacturerInfoVO":null,"detailSellerPointVO":{"ru":[],"en":["Distinctive Feature:The rings distinctive feature of droplet glaze adds a unique touch to your style, making you stand out.","Comfortable Fit:Designed to fit comfortably on the index finger, this ring provides a pleasant wearing experience.","Trendy Accessory:As a trendy accessory, this ring adds a touch of coolness to your look, making it ideal for vacations or special occasions.","High-Quality Material:Crafted with high-quality enamel, this ring is not only stylish but also durable and long-lasting.","Versatile Style:With its fresh and retro blue glaze finish, this ring can be paired with a variety of outfits, perfect for both casual and formal occasions.","Unique Design:This ring features a unique design inspired by coconut trees, making it a standout piece in any jewelry collection."],"zh":null},"shortTitle":{"ru":null,"en":["Fashion Cool Vacation Style Index Ring","Retro Blue Enamel Palm Tree Ring"],"zh":null}} |

<a id="m-product-distribute-getdistributeinfo-itemaigcvo"></a>
#### product.distribute.getDistributeInfo.ItemAIGCVO

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `detailText` | java.lang.String[] | yes | Product detail page text | [] |
| `detailSellerPointVO` | [message:product.distribute.getDistributeInfo.SellerPointVO](#m-product-distribute-getdistributeinfo-sellerpointvo) | yes | Selling points on the product detail page | {} |
| `keywords` | [message:product.distribute.getDistributeInfo.SellerPointVO](#m-product-distribute-getdistributeinfo-sellerpointvo) | yes | Selling points on the product detail page | {} |
| `shortTitle` | [message:product.distribute.getDistributeInfo.SellerPointVO](#m-product-distribute-getdistributeinfo-sellerpointvo) | yes | Short title | {} |
| `manufacturerInfoVO` | [message:alibaba.cbu.aigc.dto.ManufacturerInfoVO](#m-alibaba-cbu-aigc-dto-manufacturerinfovo) | yes | Manufacturer information | {} |

<a id="m-product-distribute-getdistributeinfo-sellerpointvo"></a>
#### product.distribute.getDistributeInfo.SellerPointVO

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `zh` | java.lang.String[] | yes | Chinese | 1 |
| `en` | java.lang.String[] | yes | English | 1 |
| `ru` | java.lang.String[] | yes | Russian | 1 |

<a id="m-alibaba-cbu-aigc-dto-manufacturerinfovo"></a>
#### alibaba.cbu.aigc.dto.ManufacturerInfoVO

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `manufacturerNameCn` | java.lang.String | yes | Manufacturer name (Chinese) | 制造商1 |
| `manufacturerNameEn` | java.lang.String | yes | Manufacturer name (English) | 1 |
| `addressCn` | java.lang.String | yes | Address (Chinese) | 北京市 |
| `addressEn` | java.lang.String | yes | Manufacturer name (English) | beijing |
| `phone` | java.lang.String | yes | Phone number | 15678987689 |
| `email` | java.lang.String | yes | Email | 12121212@163.com |
