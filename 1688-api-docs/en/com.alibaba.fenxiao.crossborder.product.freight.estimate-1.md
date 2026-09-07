# Estimate domestic (China) shipping fee for a product

Original name: 商品中国国内运费预估  
API: `com.alibaba.fenxiao.crossborder:product.freight.estimate:1` · Category: Logistics  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:product.freight.estimate-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/product.freight.estimate/{appKey}`  
Requires user authorization (access_token) · Requires signature

Estimate a product's shipping fee from the product ID and the province/city/district codes of a delivery address within mainland China.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `productFreightQueryParamsNew` | [message:product.freight.estimate.ProductFreightQueryParamsNew](#m-product-freight-estimate-productfreightqueryparamsnew) | yes | Input parameter | 如下 |

<a id="m-product-freight-estimate-productfreightqueryparamsnew"></a>
#### product.freight.estimate.ProductFreightQueryParamsNew

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `offerId` | Long | yes | Product ID | 111111111 |
| `toProvinceCode` | String | yes | China province code | 如浙江省330000 |
| `toCityCode` | String | yes | China city code | 如杭州市330100 |
| `toCountryCode` | String | yes | China region code | 如滨江区330108 |
| `totalNum` | Long | yes | Purchase quantity | 3 |
| `logisticsSkuNumModels` | [message:product.freight.estimate.LogisticsSkuNumModel[]](#m-product-freight-estimate-logisticsskunummodel[]) | yes | SKU quantity | 1 |

<a id="m-product-freight-estimate-logisticsskunummodel[]"></a>
#### product.freight.estimate.LogisticsSkuNumModel[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `skuId` | String | yes | skuId | 12345 |
| `number` | Long | yes | Quantity | 1 |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:product.freight.estimate.ResultModel](#m-product-freight-estimate-resultmodel) | yes |  |  |

<a id="m-product-freight-estimate-resultmodel"></a>
#### product.freight.estimate.ResultModel

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `success` | boolean | yes | Result | true |
| `code` | java.lang.String | yes | Error code | S0000 |
| `message` | java.lang.String | yes | Error description | 成功 |
| `result` | [message:product.freight.estimate.ProductFreightModel](#m-product-freight-estimate-productfreightmodel) | yes | Internal result | 如下 |

<a id="m-product-freight-estimate-productfreightmodel"></a>
#### product.freight.estimate.ProductFreightModel

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `offerId` | java.lang.Long | yes | Product ID | 111111111 |
| `freight` | java.lang.String | yes | Estimated total shipping fee | 12.3 |
| `templateId` | java.lang.Long | yes | Merchant freight template ID | 2324342 |
| `singleProductWeight` | java.lang.Double | yes | Single product weight, in kilograms | 0.15 |
| `singleProductWidth` | java.lang.Double | yes | Single product width, in centimeters | 10 |
| `singleProductHeight` | java.lang.Double | yes | Height of a single product, in centimeters | 10 |
| `singleProductLength` | java.lang.Double | yes | Single product length, unit: centimeters | 10 |
| `templateType` | java.lang.Integer | yes | Merchant shipping fee template type: 1 seller bears the shipping fee, 2 user-defined template, 3 user-defined official template | 2 |
| `templateName` | java.lang.String | yes | Seller shipping fee template name | 测试 |
| `subTemplateType` | java.lang.Integer | yes | Seller shipping fee sub-template type: 0 express delivery, 1 freight, 2 cash on delivery, 3 official | 0 |
| `subTemplateName` | java.lang.String | yes | Merchant shipping fee sub-template name | 测试 |
| `firstFee` | java.lang.String | yes | First-weight/per-piece fee, in yuan | 2.9 |
| `firstUnit` | java.lang.String | yes | First-weight/piece unit, | 如1代表【1件或者首重1千克是firstFee元】 |
| `nextFee` | java.lang.String | yes | Additional weight/piece fee, in yuan | 2 |
| `nextUnit` | java.lang.String | yes | Additional weight/piece unit | 如1代表【之后的1件或者续重1千克是nextFee元】 |
| `discount` | java.lang.String | yes | Freight discount | 空或者1，代表无折扣 |
| `chargeType` | java.lang.String | yes | Billing type: 0 weight, 1 quantity, 2 volume | 0 |
| `freePostage` | Boolean | yes | Whether shipping is included: true - free shipping, false - not free shipping; if not free shipping, read freight | true |
| `productFreightSkuInfoModels` | [message:product.freight.estimate.ProductFreightSkuInfoModel[]](#m-product-freight-estimate-productfreightskuinfomodel[]) | yes | SKU packaging information | 1 |
| `sizeValueType` | Integer | yes | Value type for package weight/dimensions, 1 - value taken from the outer layer, 2 - value taken from productFreightSkuInfoModels | 2 |

<a id="m-product-freight-estimate-productfreightskuinfomodel[]"></a>
#### product.freight.estimate.ProductFreightSkuInfoModel[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `skuId` | String | yes | skuId | 12345678 |
| `singleSkuWeight` | Double | yes | Weight | 1 |
| `singleSkuWidth` | Double | yes | Width | 1 |
| `singleSkuHeight` | Double | yes | Height | 1 |
| `singleSkuLength` | Double | yes | Length | 1 |
