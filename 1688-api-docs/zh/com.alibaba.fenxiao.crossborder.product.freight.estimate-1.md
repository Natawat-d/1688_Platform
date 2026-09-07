# 商品中国国内运费预估

API: `com.alibaba.fenxiao.crossborder:product.freight.estimate:1` · Category: 物流  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:product.freight.estimate-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/product.freight.estimate/{appKey}`  
需要授权 (access_token) · 需要签名

根据商品ID、中国国内收货地址的省市区编码，预估商品的运费。

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `productFreightQueryParamsNew` | [message:product.freight.estimate.ProductFreightQueryParamsNew](#m-product-freight-estimate-productfreightqueryparamsnew) | 是 | 入参 | 如下 |

<a id="m-product-freight-estimate-productfreightqueryparamsnew"></a>
#### product.freight.estimate.ProductFreightQueryParamsNew

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `offerId` | Long | 是 | 商品ID | 111111111 |
| `toProvinceCode` | String | 是 | 中国省份编码 | 如浙江省330000 |
| `toCityCode` | String | 是 | 中国城市编码 | 如杭州市330100 |
| `toCountryCode` | String | 是 | 中国地区编码 | 如滨江区330108 |
| `totalNum` | Long | 是 | 购买件数 | 3 |
| `logisticsSkuNumModels` | [message:product.freight.estimate.LogisticsSkuNumModel[]](#m-product-freight-estimate-logisticsskunummodel[]) | 是 | sku件数 | 1 |

<a id="m-product-freight-estimate-logisticsskunummodel[]"></a>
#### product.freight.estimate.LogisticsSkuNumModel[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `skuId` | String | 是 | skuId | 12345 |
| `number` | Long | 是 | 数量 | 1 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:product.freight.estimate.ResultModel](#m-product-freight-estimate-resultmodel) | 是 |  |  |

<a id="m-product-freight-estimate-resultmodel"></a>
#### product.freight.estimate.ResultModel

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `success` | boolean | 是 | 结果 | true |
| `code` | java.lang.String | 是 | 错误码 | S0000 |
| `message` | java.lang.String | 是 | 错误描述 | 成功 |
| `result` | [message:product.freight.estimate.ProductFreightModel](#m-product-freight-estimate-productfreightmodel) | 是 | 内部结果 | 如下 |

<a id="m-product-freight-estimate-productfreightmodel"></a>
#### product.freight.estimate.ProductFreightModel

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `offerId` | java.lang.Long | 是 | 商品ID | 111111111 |
| `freight` | java.lang.String | 是 | 预估总运费 | 12.3 |
| `templateId` | java.lang.Long | 是 | 商家运费模板ID | 2324342 |
| `singleProductWeight` | java.lang.Double | 是 | 单商品重量，单位千克 | 0.15 |
| `singleProductWidth` | java.lang.Double | 是 | 单商品宽度，单位厘米 | 10 |
| `singleProductHeight` | java.lang.Double | 是 | 单商品高度，单位厘米 | 10 |
| `singleProductLength` | java.lang.Double | 是 | 单商品长度，单位厘米 | 10 |
| `templateType` | java.lang.Integer | 是 | 商家运费模板类型，1卖家承担运费，2用户自定义模板，3用户自定义官方模版 | 2 |
| `templateName` | java.lang.String | 是 | 商家运费模板名称 | 测试 |
| `subTemplateType` | java.lang.Integer | 是 | 商家运费子模板类型，0快递，1货运，2货到付款，3官方 | 0 |
| `subTemplateName` | java.lang.String | 是 | 商家运费子模板名称 | 测试 |
| `firstFee` | java.lang.String | 是 | 首重/件费用，单位元 | 2.9 |
| `firstUnit` | java.lang.String | 是 | 首重/件单位， | 如1代表【1件或者首重1千克是firstFee元】 |
| `nextFee` | java.lang.String | 是 | 续重/件费用，单位元 | 2 |
| `nextUnit` | java.lang.String | 是 | 续重/件单位 | 如1代表【之后的1件或者续重1千克是nextFee元】 |
| `discount` | java.lang.String | 是 | 运费折扣 | 空或者1，代表无折扣 |
| `chargeType` | java.lang.String | 是 | 计费类型，0重量，1件数，2体积 | 0 |
| `freePostage` | Boolean | 是 | 是否包邮，true-包邮 false-不包邮，不包邮freight读取 | true |
| `productFreightSkuInfoModels` | [message:product.freight.estimate.ProductFreightSkuInfoModel[]](#m-product-freight-estimate-productfreightskuinfomodel[]) | 是 | sku包装信息 | 1 |
| `sizeValueType` | Integer | 是 | 件重尺取值类型，1-为外层取值，2-为从productFreightSkuInfoModels中取值 | 2 |

<a id="m-product-freight-estimate-productfreightskuinfomodel[]"></a>
#### product.freight.estimate.ProductFreightSkuInfoModel[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `skuId` | String | 是 | skuId | 12345678 |
| `singleSkuWeight` | Double | 是 | 重量 | 1 |
| `singleSkuWidth` | Double | 是 | 宽 | 1 |
| `singleSkuHeight` | Double | 是 | 高 | 1 |
| `singleSkuLength` | Double | 是 | 长 | 1 |
