# Query leaf-category attribute and value mappings

Original name: 叶子类目属性属性值映射查询  
API: `com.alibaba.fenxiao.crossborder:product.category.getAttrById:1` · Category: Products  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:product.category.getAttrById-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/product.category.getAttrById/{appKey}`  
Requires user authorization (access_token) · Requires signature

Category mapping query.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `categoryId` | java.lang.String | yes |  |  |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `success` | Boolean | yes |  |  |
| `msg` | String | yes |  |  |
| `data` | [message:alibaba.cbu.mapping.dto.PVMappingVO](#m-alibaba-cbu-mapping-dto-pvmappingvo) | yes |  |  |

<a id="m-alibaba-cbu-mapping-dto-pvmappingvo"></a>
#### alibaba.cbu.mapping.dto.PVMappingVO

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `categoryId` | java.lang.String | yes | Category ID | 1032097 |
| `categoryName` | java.lang.String | yes | Category name | 车用吸尘器 |
| `externalCategoryId` | java.lang.String | yes | External category ID | 2012 |
| `externalCategoryName` | java.lang.String | yes | External category name | 吸尘器车用 |
| `attributes` | [message:alibaba.cbu.mapping.dto.PVMappingAttribute[]](#m-alibaba-cbu-mapping-dto-pvmappingattribute[]) | yes | Attribute | 略 |
| `productShippingInfo` | [message:com.alibaba.cbu.mapping.dto.ProductShippingInfo](#m-com-alibaba-cbu-mapping-dto-productshippinginfo) | yes | Shipping information | 如下 |

<a id="m-alibaba-cbu-mapping-dto-pvmappingattribute[]"></a>
#### alibaba.cbu.mapping.dto.PVMappingAttribute[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `attributeId` | java.lang.String | yes | Attribute ID | 1945076 |
| `attributeName` | java.lang.String | yes | Attribute name | 电源线尺寸 |
| `valueId` | java.lang.String | yes | Value ID | 8528111 |
| `valueName` | java.lang.String | yes | Value name | 280cm |
| `externalAttributeId` | java.lang.String | yes | External attribute ID | 85079 |
| `externalValueId` | java.lang.String | yes | External value ID | 12345 |
| `externalAttributeName` | java.lang.String | yes | External attribute name | 电线尺寸 |
| `externalValueName` | String | yes | External value name | 外部值名称 |

<a id="m-com-alibaba-cbu-mapping-dto-productshippinginfo"></a>
#### com.alibaba.cbu.mapping.dto.ProductShippingInfo

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `weight` | String | yes | Weight, in kg | 0.05 |
| `width` | String | yes | Width, in cm | 5 |
| `height` | String | yes | Height, in cm | 5 |
| `length` | String | yes | Length, in cm | 5 |
