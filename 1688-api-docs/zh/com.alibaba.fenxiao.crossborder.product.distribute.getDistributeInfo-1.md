# 获取商品铺货卖点

API: `com.alibaba.fenxiao.crossborder:product.distribute.getDistributeInfo:1` · Category: 商品  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:product.distribute.getDistributeInfo-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/product.distribute.getDistributeInfo/{appKey}`  
无需授权 · 需要签名

用于铺货场景下，获取铺货需要的卖点信息

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 否 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `offerId` | java.lang.String | 是 | offerId | 1 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:product.distribute.getDistributeInfo.Result](#m-product-distribute-getdistributeinfo-result) | 是 | 返回结果 | {} |

<a id="m-product-distribute-getdistributeinfo-result"></a>
#### product.distribute.getDistributeInfo.Result

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `success` | java.lang.Boolean | 是 | 请求是否成功 | true |
| `msg` | java.lang.String | 是 | 错误信息 | 系统异常 |
| `data` | [message:product.distribute.getDistributeInfo.ItemAIGCVO](#m-product-distribute-getdistributeinfo-itemaigcvo) | 是 | 具体业务数据 | {"keywords":{"ru":[],"en":["Distinctive","Comfortable","Trendy","Quality","Versatile","Unique"],"zh":null},"detailText":[],"manufacturerInfoVO":null,"detailSellerPointVO":{"ru":[],"en":["Distinctive Feature:The rings distinctive feature of droplet glaze adds a unique touch to your style, making you stand out.","Comfortable Fit:Designed to fit comfortably on the index finger, this ring provides a pleasant wearing experience.","Trendy Accessory:As a trendy accessory, this ring adds a touch of coolness to your look, making it ideal for vacations or special occasions.","High-Quality Material:Crafted with high-quality enamel, this ring is not only stylish but also durable and long-lasting.","Versatile Style:With its fresh and retro blue glaze finish, this ring can be paired with a variety of outfits, perfect for both casual and formal occasions.","Unique Design:This ring features a unique design inspired by coconut trees, making it a standout piece in any jewelry collection."],"zh":null},"shortTitle":{"ru":null,"en":["Fashion Cool Vacation Style Index Ring","Retro Blue Enamel Palm Tree Ring"],"zh":null}} |

<a id="m-product-distribute-getdistributeinfo-itemaigcvo"></a>
#### product.distribute.getDistributeInfo.ItemAIGCVO

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `detailText` | java.lang.String[] | 是 | 商品详情页文本 | [] |
| `detailSellerPointVO` | [message:product.distribute.getDistributeInfo.SellerPointVO](#m-product-distribute-getdistributeinfo-sellerpointvo) | 是 | 商品详情页卖点 | {} |
| `keywords` | [message:product.distribute.getDistributeInfo.SellerPointVO](#m-product-distribute-getdistributeinfo-sellerpointvo) | 是 | 商品详情页卖点 | {} |
| `shortTitle` | [message:product.distribute.getDistributeInfo.SellerPointVO](#m-product-distribute-getdistributeinfo-sellerpointvo) | 是 | 短标题 | {} |
| `manufacturerInfoVO` | [message:alibaba.cbu.aigc.dto.ManufacturerInfoVO](#m-alibaba-cbu-aigc-dto-manufacturerinfovo) | 是 | 制造商信息 | {} |

<a id="m-product-distribute-getdistributeinfo-sellerpointvo"></a>
#### product.distribute.getDistributeInfo.SellerPointVO

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `zh` | java.lang.String[] | 是 | 中文 | 1 |
| `en` | java.lang.String[] | 是 | 英文 | 1 |
| `ru` | java.lang.String[] | 是 | 俄语 | 1 |

<a id="m-alibaba-cbu-aigc-dto-manufacturerinfovo"></a>
#### alibaba.cbu.aigc.dto.ManufacturerInfoVO

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `manufacturerNameCn` | java.lang.String | 是 | 制造商名称（中文） | 制造商1 |
| `manufacturerNameEn` | java.lang.String | 是 | 制造商名称（英文） | 1 |
| `addressCn` | java.lang.String | 是 | 地址（中文） | 北京市 |
| `addressEn` | java.lang.String | 是 | 制造商名称（英文） | beijing |
| `phone` | java.lang.String | 是 | 电话 | 15678987689 |
| `email` | java.lang.String | 是 | 邮箱 | 12121212@163.com |
