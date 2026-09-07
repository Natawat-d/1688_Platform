# 跨境场景获取商品列表

API: `com.alibaba.product:alibaba.cross.productList:1` · Category: 选品铺货  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.product:alibaba.cross.productList-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.product/alibaba.cross.productList/{appKey}`  
需要授权 (access_token) · 需要签名

该接口需要校验跨境的铺货关系，跨境业务专用。商品模型V2

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `productIdList` | Long[] | 是 | 商品Id列表 | [574325651942,570027659932] |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `productList` | [message:alibaba.product.ProductInfo[]](#m-alibaba-product-productinfo[]) | 是 | 商品列表 | [] |
| `success` | Boolean | 是 | 是否成功 | true |
| `message` | String | 是 | 返回信息 | 57002765XXX:商品不存在 |

<a id="m-alibaba-product-productinfo[]"></a>
#### alibaba.product.ProductInfo[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `productID` | Long | 是 | 商品ID | 574325651942 |
| `productType` | String | 是 | 商品类型，在线批发商品(wholesale)或者询盘商品(sourcing)，1688网站缺省为wholesale | wholesale |
| `attributes` | [message:alibaba.product.ProductAttribute[]](#m-alibaba-product-productattribute[]) | 是 | 商品属性和属性值 |   |
| `groupID` | Long[] | 是 | 分组ID，确定商品所属分组。1688可传入多个分组ID，国际站同一个商品只能属于一个分组，因此默认只取第一个 |   |
| `status` | String | 是 | 商品状态。published:上网状态;member expired:会员撤销;auto expired:自然过期;expired:过期(包含手动过期与自动过期);member deleted:会员删除;modified:修改;new:新发;deleted:删除;TBD:to be delete;approved:审批通过;auditing:审核中;untread:审核不通过; | published |
| `subject` | String | 是 | 商品标题，最多128个字符 | 吉米兔小时候 宝宝成长纪念册相册配套 使用空白记录彩色加页1 |
| `description` | String | 是 | 商品详情描述，可包含图片中心的图片URL |   |
| `language` | String | 是 | 语种，参见FAQ 语种枚举值，1688网站默认传入CHINESE | CHINESE |
| `periodOfValidity` | Integer | 是 | 信息有效期，按天计算，国际站无此信息 | 3650 |
| `bizType` | Integer | 是 | 业务类型。1：商品，2：加工，3：代理，4：合作，5：商务服务。国际站按默认商品。 | 1 |
| `image` | [message:alibaba.product.ProductImageInfo](#m-alibaba-product-productimageinfo) | 是 | 商品主图 | {} |
| `extendInfos` | [message:alibaba.product.ProductExtendInfo[]](#m-alibaba-product-productextendinfo[]) | 是 | 商品扩展信息 |   |
| `supplierLoginId` | java.lang.String | 是 | 供应商loginId | alitestforisv02 |
| `categoryID` | Long | 是 | 类目ID，标识商品所属类目 | 1048182 |
| `categoryName` | java.lang.String | 是 | 类目名 |   |
| `productCargoNumber` | java.lang.String | 是 | 商品货号，产品属性中的货号 |   |

<a id="m-alibaba-product-productattribute[]"></a>
#### alibaba.product.ProductAttribute[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `attributeID` | Long | 是 | 属性ID | 123456 |
| `attributeName` | String | 是 | 属性名称 | color |
| `valueID` | Long | 是 | 属性值ID | 123456 |
| `value` | String | 是 | 属性值 | grey |
| `isCustom` | Boolean | 是 | 是否为自定义属性，国际站无需关注 | true |

<a id="m-alibaba-product-productimageinfo"></a>
#### alibaba.product.ProductImageInfo

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `images` | String[] | 是 | 主图列表，需先使用图片上传接口上传图片 | ["img/ibank/2018/502/115/9153511205_1606139362.jpg"] |
| `isWatermark` | Boolean | 是 | 是否打水印，是(true)或否(false)，1688无需关注此字段，1688的水印信息在上传图片时处理 |   |
| `isWatermarkFrame` | Boolean | 是 | 水印是否有边框，有边框(true)或者无边框(false)，1688无需关注此字段，1688的水印信息在上传图片时处理 |   |
| `watermarkPosition` | String | 是 | 水印位置，在中间(center)或者在底部(bottom)，1688无需关注此字段，1688的水印信息在上传图片时处理 |   |

<a id="m-alibaba-product-productextendinfo[]"></a>
#### alibaba.product.ProductExtendInfo[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `key` | String | 是 | 扩展结构的key | 代销价格,consignPrice;<br>买家保障,buyerProtection; |
| `value` | String | 是 | 扩展结构的value | 代销价格,key为skuId，value为用户设置的代销价，<br>示例：31151771910:2088.0;31151771909:2088.0;31151771908:2088.0;31152339121:2088.0;<br>买家保障,string数组，value为买保全拼，<br>示例：["psbj","swtwlybt","swtbh","ssbxsfh"] |

## 错误码

| 错误码 | 现象 | 解决方案 |
|---|---|---|
| 商品不存在 | 574325651XXX:商品不存在,57002765XXX:商品不存在 | 检查铺货列表中是否存在该商品 |

## 示例

**入参示例**

```
{
  "productIdList":[574325651942,570027659932]
}
```

**出参示例**

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
