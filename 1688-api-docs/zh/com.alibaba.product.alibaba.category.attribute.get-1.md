# 获取叶子类目属性

API: `com.alibaba.product:alibaba.category.attribute.get:1` · Category: 类目  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.product:alibaba.category.attribute.get-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.product/alibaba.category.attribute.get/{appKey}`  
需要授权 (access_token) · 需要签名

根据叶子类目ID获取类目属性

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `categoryID` | Long | 是 | 类目ID |   |
| `webSite` | String | 是 | 站点信息，指定调用的API是属于国际站（alibaba）还是1688网站（1688） |   |
| `scene` | String | 否 | 场景值，可选值为 空 和 processing，默认为空 |  |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `attributes` | [message:alibaba.category.AttributeInfo[]](#m-alibaba-category-attributeinfo[]) | 是 | 类目属性信息 | [] |
| `levelAttrRelList` | [message:alibaba.category.PostLevelAttrRel[]](#m-alibaba-category-postlevelattrrel[]) | 是 | (废弃)类目属性级联关系，只有1688业务返回返回该字段 | [] |
| `attributeLevelMapStr` | java.util.Map | 是 | 级联信息字符串，可强转成map | {"1811:3289490":"20602,2917380,7001","100000691:46874>7108:21958":"8243"} |
| `errorMsg` | String | 是 | 错误描述 |    |
| `errorCode` | String | 是 | 错误码 | 500_1 |
| `success` | Boolean | 是 | 是否成功 | true |

<a id="m-alibaba-category-attributeinfo[]"></a>
#### alibaba.category.AttributeInfo[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `attrID` | Long | 是 | 属性id | 123 |
| `name` | String | 是 | 名称 | 长度 |
| `required` | Boolean | 是 | 是否必填属性 | true |
| `units` | String[] | 是 | 该属性的单位 | mm |
| `isSKUAttribute` | Boolean | 是 | 该属性能否当成SKU属性 | true |
| `attrValues` | [message:alibaba.category.AttributeValueInfo[]](#m-alibaba-category-attributevalueinfo[]) | 是 | 属性可选的属性值 | [] |
| `inputType` | String | 是 | 输入类型。<br>下拉框:1,<br>多选框:2<br>单选框:3,<br>文本输入框:0,<br>数字输入框:-1,<br>下拉框列表:4,<br>日期：5 | 1 |
| `isSupportCustomizeValue` | Boolean | 是 | 用成SKU属性时，是否支持自定义属性值名称，1688不返回该信息 | true |
| `isSupportCustomizeImage` | Boolean | 是 | 用成SKU属性时，是否支持自定义图片展示，1688不返回该信息 | true |
| `enName` | String | 是 | 英文名称，1688无此属性 | length |
| `parentAttrID` | String | 是 | 父属性ID，如果此值为空或零，则表示此属性为一级属性 | 287 |
| `parentAttrValueID` | String | 是 | 父属性值ID，如果此值为空或零，则表示此属性为一级属性 | 3737061 |
| `aspect` | String | 是 | 产品属性:0,<br>交易属性:3,<br>spu匹配属性:5 | 0 |
| `fieldType` | String | 是 | 类型，int：数字；string:字符串；enum：枚举 | enum |
| `isSpecPicAttr` | Boolean | 是 | 是否图片属性 | false |
| `firstLevel` | Boolean | 是 | 是否为一级属性 | true |
| `attrType` | String | 是 | 专业化类目，属性类型 | 0：产品属性，1：规格属性，2：规格扩展属性 |
| `sort` | Integer | 是 | 类目排序 | 1 |
| `recommendAttr` | Boolean | 是 | 是否推荐属性 | false |

<a id="m-alibaba-category-attributevalueinfo[]"></a>
#### alibaba.category.AttributeValueInfo[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `attrValueID` | Long | 是 | 属性值id |  |
| `name` | String | 是 | 名称 |  |
| `enName` | String | 是 | 英文名称 |  |
| `childAttrs` | Long[] | 是 | 该属性值的子属性id |  |
| `isSKU` | Boolean | 是 | 是否SKU属性值 |  |

<a id="m-alibaba-category-postlevelattrrel[]"></a>
#### alibaba.category.PostLevelAttrRel[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `fid` | int | 是 | 属性id |  |
| `subFids` | int[] | 是 | 子关联属性 |  |
| `attrType` | int | 是 | 0和空都为现货属性层级关系，1为加工属性层级关系，后面其它的可加 |  |

## 错误码

| 错误码 | 现象 | 解决方案 |
|---|---|---|
| 500_2 | 数据准备中，请稍后重试。 | 数据正在后台加载，稍后重试，间隔时间建议1～3s |

## 示例

**返回报文示例**

```
{
    "attributeLevelMapStr":{
        "2489638:9955810":"973"
    },
    "attributes":[
        {
            "attrID":2489638,
            "name":"风格类型",
            "required":true,
            "fieldType":"enum",
            "isSKUAttribute":false,
            "attrValues":[
                {
                    "attrValueID":91043051,
                    "name":"气质通勤"
                }
            ],
            "inputType":"1",
            "aspect":"0;",
            "isSpecPicAttr":false
        }
    ]
}
```

**关于级联属性map**

```
map中保存了该类目所有可能的级联属性，以key-value对形式出现。
其中，key为级联关系产生的条件。冒号前为属性id，冒号后面为属性值id
>表示层级关系
value为满足级联条件后下级需要填写的属性。

如：连衣裙类目的属性层级是：货源类别->是否库存->库存类型，即当"货源类别"属性为"现货"的时候，需要继续填写"是否库存"属性，当"是否库存"属性选择"是"，则需要继续填写"库存类型"属性。
该级联属性的key-value对为 "100000691:46874>7108:21958":"8243"
100000691（货源类别的属性ID）46874（现货的属性值ID）7108（是否库存的属性ID）21958（是的属性值ID）8243（库存类型的属性ID）
```
