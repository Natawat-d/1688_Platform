# 跨境场景根据关键字推荐商品

API: `com.alibaba.product:alibaba.product.suggest.crossBorder:1` · Category: 商品  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.product:alibaba.product.suggest.crossBorder-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.product/alibaba.product.suggest.crossBorder/{appKey}`  
需要授权 (access_token) · 需要签名

跨境场景根据关键字及类目推荐商品，按销量排序。注意：该API有流量控制，仅适用于手动关联场景

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `keyWord` | String | 是 | 商品的搜索关键字，通常是商品的标题 | 商品标题 |
| `loginId` | String | 否 | 卖家loginId | alitestforisv01 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `resultList` | [message:alibaba.search.ProductSearchResultInfo[]](#m-alibaba-search-productsearchresultinfo[]) | 是 | 搜索的返回结果 | [] |
| `success` | String | 是 | 是否成功 | true |
| `errorMsg` | String | 是 | 错误描述 |   |
| `errorCode` | String | 是 | 错误码 |   |

<a id="m-alibaba-search-productsearchresultinfo[]"></a>
#### alibaba.search.ProductSearchResultInfo[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `amountOnSale` | Integer | 是 | 可售数量 | 100 |
| `minPurchaseQuantity` | Double | 是 | 最小起批量 | 3 |
| `picUrl` | String | 是 | 产品的图片地址 | img/ibank/2018/794/316/9422613497_991974782.jpg |
| `price` | Double | 是 | 参考商品价格 | 100 |
| `productID` | Long | 是 | 商品ID | 1123123331 |
| `bookedCount` | Double | 是 | 商品售卖了多少笔(以订单为维度) | 2123 |
| `saleQuantity` | Double | 是 | 该商品售卖了多少件(以商品单位为维度) | 1 |
| `province` | String | 是 | 商品发货省份码 | 浙江 |
| `city` | String | 是 | 商品发货城市 | 杭州 |
| `retailPrice` | Double | 是 | 建议零售价 | 100 |
| `subject` | String | 是 | 商品标题 | 【原D现货】韩版新秋女裙 时尚方格<font color=red>小</font>立领不规则下摆气质<font color=red>连衣裙</font> |
| `unit` | String | 是 | 商品单位 | 件 |
| `skuList` | [message:alibaba.simple.sku[]](#m-alibaba-simple-sku[]) | 是 | sku信息 | [] |

<a id="m-alibaba-simple-sku[]"></a>
#### alibaba.simple.sku[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `description` | String | 是 | 规格描述 | 颜色:红色;尺码:L |
| `amountOnSale` | Integer | 是 | 可售数量 | 100 |
| `skuId` | Long | 是 | skuID，全局唯一标示 | 3508426014362 |
| `specId` | String | 是 | specID，商品内唯一，不同商品间可能重复 | 8d28b045489c250b69870da3b7c71b1d |

## 示例

**返回结果示例**

```
{
  "resultList": [
    {
      "amountOnSale": 241682,
      "city": "义乌市",
      "minPurchaseQuantity": 100,
      "picUrl": "https://cbu01.alicdn.com/img/ibank/2015/112/013/2204310211_1220361846.jpg",
      "price": 0.06,
      "productID": 45636254415,
      "province": "浙江",
      "retailPrice": 0.5,
      "saleQuantity": 159821,
      "subject": "创意儿童戒指批发 糖果色塑料/树脂戒指混批时尚女童女孩戒指批发",
      "unit": "PCS",
      "skuList": [
        {
          "description": "颜色:黄色",
          "amountOnSale": 36866,
          "skuId": 92418963545,
          "specId": "36c58a69219820709b9475e70bd789d4"
        }
      ]
    }
  ],
  "success": "true"
}
```
