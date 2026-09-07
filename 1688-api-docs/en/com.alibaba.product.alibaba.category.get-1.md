# Query category by category ID

Original name: 根据类目Id查询类目  
API: `com.alibaba.product:alibaba.category.get:1` · Category: Product Selection & Listing  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.product:alibaba.category.get-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.product/alibaba.category.get/{appKey}`  
No user authorization · Requires signature

Category query. To retrieve all 1688 categories, traverse the whole category tree starting from the root: first pass 0 to get all level-1 category IDs, then iterate the level-1 IDs to get all level-2 categories, and finally iterate the level-2 IDs to get the level-3 categories. Note: 1688 categories have only three levels; level-3 categories are the leaf categories required for publishing products.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | no | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `categoryID` | Long | yes | Category ID, must be greater than or equal to 0. If 0, all first-level categories are queried | 1031910 |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `succes` | String | yes | Whether successful | true |
| `categoryInfo` | [message:alibaba.category.CategoryInfo[]](#m-alibaba-category-categoryinfo[]) | yes | List of categories | {} |
| `errorMsg` | String | yes | Error message |   |
| `errorCode` | String | yes | Error code | 500_1 |

<a id="m-alibaba-category-categoryinfo[]"></a>
#### alibaba.category.CategoryInfo[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `categoryID` | Long | yes | Category ID | 123456 |
| `name` | String | yes | Category name | 连衣裙 |
| `level` | Integer | yes | Category level; not applicable on 1688. | 0 |
| `isLeaf` | Boolean | yes | Whether it is a leaf category (only leaf categories can publish products) | false |
| `parentIDs` | Long[] | yes | Array of parent category IDs; 1688 only returns one parent ID | [23] |
| `childIDs` | Long[] | yes | Array of sub-category IDs; not applicable to 1688 | [] |
| `childCategorys` | [message:alibaba.child.category.info[]](#m-alibaba-child-category-info[]) | yes | Sub-category information | "" |
| `minOrderQuantity` | Long | yes | Minimum order quantity | 1 |
| `featureInfos` | [message:alibaba.category.categoryInfo.FeatureInfo[]](#m-alibaba-category-categoryinfo-featureinfo[]) | yes | Category feature information | [{}] |
| `categoryType` | String | yes | Category type. 1 is 1688 open marketplace category, 2 is 1688 industrial products specialized category, 3 is 1688 mainstream product category | 1 |
| `isSupportProcessing` | Boolean | yes | Whether the category supports processing/customization | true |

<a id="m-alibaba-child-category-info[]"></a>
#### alibaba.child.category.info[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `id` | Long | yes | Sub-category ID | 54 |
| `name` | String | yes | Sub-category name | 服饰配件、饰品 |
| `isLeaf` | Boolean | yes | Whether it is a leaf category (only leaf categories can publish products) | true |
| `categoryType` | String | yes | Category type: 1 indicates CBU category, 2 indicates Gallop category | 1 |

<a id="m-alibaba-category-categoryinfo-featureinfo[]"></a>
#### alibaba.category.categoryInfo.FeatureInfo[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `key` | String | yes | Name | name |
| `value` | String | yes | Value | jiagong |
| `status` | Integer | yes | Status | 0 |
| `hierarchy` | Boolean | yes | Whether it is inherited to child elements | true |

## Samples

**Return example**

```
{}
```
