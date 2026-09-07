# 叶子类目属性属性值映射查询

API: `com.alibaba.fenxiao.crossborder:product.category.getAttrById:1` · Category: 商品  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:product.category.getAttrById-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/product.category.getAttrById/{appKey}`  
需要授权 (access_token) · 需要签名

类目映射查询

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `categoryId` | java.lang.String | 是 |  |  |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `success` | Boolean | 是 |  |  |
| `msg` | String | 是 |  |  |
| `data` | [message:alibaba.cbu.mapping.dto.PVMappingVO](#m-alibaba-cbu-mapping-dto-pvmappingvo) | 是 |  |  |

<a id="m-alibaba-cbu-mapping-dto-pvmappingvo"></a>
#### alibaba.cbu.mapping.dto.PVMappingVO

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `categoryId` | java.lang.String | 是 | 类目ID | 1032097 |
| `categoryName` | java.lang.String | 是 | 类目名称 | 车用吸尘器 |
| `externalCategoryId` | java.lang.String | 是 | 外部类目ID | 2012 |
| `externalCategoryName` | java.lang.String | 是 | 外部类目名称 | 吸尘器车用 |
| `attributes` | [message:alibaba.cbu.mapping.dto.PVMappingAttribute[]](#m-alibaba-cbu-mapping-dto-pvmappingattribute[]) | 是 | 属性 | 略 |
| `productShippingInfo` | [message:com.alibaba.cbu.mapping.dto.ProductShippingInfo](#m-com-alibaba-cbu-mapping-dto-productshippinginfo) | 是 | 发货信息 | 如下 |

<a id="m-alibaba-cbu-mapping-dto-pvmappingattribute[]"></a>
#### alibaba.cbu.mapping.dto.PVMappingAttribute[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `attributeId` | java.lang.String | 是 | 属性ID | 1945076 |
| `attributeName` | java.lang.String | 是 | 属性名称 | 电源线尺寸 |
| `valueId` | java.lang.String | 是 | 值ID | 8528111 |
| `valueName` | java.lang.String | 是 | 值名称 | 280cm |
| `externalAttributeId` | java.lang.String | 是 | 外部属性ID | 85079 |
| `externalValueId` | java.lang.String | 是 | 外部值ID | 12345 |
| `externalAttributeName` | java.lang.String | 是 | 外部属性名称 | 电线尺寸 |
| `externalValueName` | String | 是 | 外部值名称 | 外部值名称 |

<a id="m-com-alibaba-cbu-mapping-dto-productshippinginfo"></a>
#### com.alibaba.cbu.mapping.dto.ProductShippingInfo

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `weight` | String | 是 | 重，单位kg | 0.05 |
| `width` | String | 是 | 宽，单位cm | 5 |
| `height` | String | 是 | 高，单位cm | 5 |
| `length` | String | 是 | 长，单位cm | 5 |
