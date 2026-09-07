# Get leaf-category attributes

Original name: 获取叶子类目属性  
API: `com.alibaba.product:alibaba.category.attribute.get:1` · Category: Categories  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.product:alibaba.category.attribute.get-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.product/alibaba.category.attribute.get/{appKey}`  
Requires user authorization (access_token) · Requires signature

Get category attributes by leaf-category ID.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `categoryID` | Long | yes | Category ID |   |
| `webSite` | String | yes | Site information, specifies whether the called API belongs to the international site (alibaba) or the 1688 website (1688). |   |
| `scene` | String | no | Scene value; optional values are empty and processing, default is empty |  |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `attributes` | [message:alibaba.category.AttributeInfo[]](#m-alibaba-category-attributeinfo[]) | yes | Category attribute info | [] |
| `levelAttrRelList` | [message:alibaba.category.PostLevelAttrRel[]](#m-alibaba-category-postlevelattrrel[]) | yes | (Deprecated) Category attribute cascading relationship; this field is only returned for 1688 business | [] |
| `attributeLevelMapStr` | java.util.Map | yes | Cascading information string, can be cast to a map | {"1811:3289490":"20602,2917380,7001","100000691:46874>7108:21958":"8243"} |
| `errorMsg` | String | yes | Error description |    |
| `errorCode` | String | yes | Error code | 500_1 |
| `success` | Boolean | yes | Whether successful | true |

<a id="m-alibaba-category-attributeinfo[]"></a>
#### alibaba.category.AttributeInfo[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `attrID` | Long | yes | Attribute id | 123 |
| `name` | String | yes | Name | 长度 |
| `required` | Boolean | yes | Whether it is a required attribute | true |
| `units` | String[] | yes | Unit of this attribute | mm |
| `isSKUAttribute` | Boolean | yes | Whether this attribute can be used as a SKU attribute | true |
| `attrValues` | [message:alibaba.category.AttributeValueInfo[]](#m-alibaba-category-attributevalueinfo[]) | yes | Optional attribute values for the attribute | [] |
| `inputType` | String | yes | Input type.<br>Dropdown: 1,<br>Multi-select box: 2<br>Radio button: 3,<br>Text input box: 0,<br>Number input box: -1,<br>Dropdown list: 4,<br>Date: 5 | 1 |
| `isSupportCustomizeValue` | Boolean | yes | Whether custom attribute value names are supported when used as a SKU attribute; 1688 does not return this information. | true |
| `isSupportCustomizeImage` | Boolean | yes | Whether custom image display is supported when used as a SKU attribute; 1688 does not return this information. | true |
| `enName` | String | yes | English name; 1688 does not have this attribute | length |
| `parentAttrID` | String | yes | Parent attribute ID. If this value is empty or zero, it indicates that this attribute is a top-level (first-level) attribute. | 287 |
| `parentAttrValueID` | String | yes | Parent attribute value ID. If this value is empty or zero, it indicates that this attribute is a top-level (first-level) attribute. | 3737061 |
| `aspect` | String | yes | Product attribute: 0,<br>Transaction attribute: 3,<br>SPU matching attribute: 5 | 0 |
| `fieldType` | String | yes | Type. int: number; string: string; enum: enumeration | enum |
| `isSpecPicAttr` | Boolean | yes | Whether it is an image attribute | false |
| `firstLevel` | Boolean | yes | Whether it is a top-level (first-level) attribute | true |
| `attrType` | String | yes | Specialized category, attribute type | 0：产品属性，1：规格属性，2：规格扩展属性 |
| `sort` | Integer | yes | Category sort order | 1 |
| `recommendAttr` | Boolean | yes | Whether it is a recommended attribute | false |

<a id="m-alibaba-category-attributevalueinfo[]"></a>
#### alibaba.category.AttributeValueInfo[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `attrValueID` | Long | yes | Attribute value ID |  |
| `name` | String | yes | Name |  |
| `enName` | String | yes | English name |  |
| `childAttrs` | Long[] | yes | The sub-attribute ID of this attribute value |  |
| `isSKU` | Boolean | yes | Whether it is a SKU attribute value |  |

<a id="m-alibaba-category-postlevelattrrel[]"></a>
#### alibaba.category.PostLevelAttrRel[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `fid` | int | yes | Attribute id |  |
| `subFids` | int[] | yes | Sub-associated attribute |  |
| `attrType` | int | yes | 0 and empty both indicate an in-stock attribute hierarchy relationship, 1 indicates a processing attribute hierarchy relationship; other values may be added later |  |

## Error codes

| Error | Symptom | How to fix |
|---|---|---|
| 500_2 | Data is being prepared, please try again later. | Data is loading in the background; please retry later. A retry interval of 1-3s is recommended. |

## Samples

**Example response message**

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

**About the cascading attribute map**

```
map中保存了该类目所有可能的级联属性，以key-value对形式出现。
其中，key为级联关系产生的条件。冒号前为属性id，冒号后面为属性值id
>表示层级关系
value为满足级联条件后下级需要填写的属性。

如：连衣裙类目的属性层级是：货源类别->是否库存->库存类型，即当"货源类别"属性为"现货"的时候，需要继续填写"是否库存"属性，当"是否库存"属性选择"是"，则需要继续填写"库存类型"属性。
该级联属性的key-value对为 "100000691:46874>7108:21958":"8243"
100000691（货源类别的属性ID）46874（现货的属性值ID）7108（是否库存的属性ID）21958（是的属性值ID）8243（库存类型的属性ID）
```
