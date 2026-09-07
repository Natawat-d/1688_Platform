# Query multilingual category by category ID

Original name: 根据类目ID查询多语言类目  
API: `com.alibaba.fenxiao.crossborder:category.translation.getById:1` · Category: Categories  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:category.translation.getById-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/category.translation.getById/{appKey}`  
Requires user authorization (access_token) · Requires signature

Multilingual category query. Returns the category details in the requested language for the given category ID, including the list of its child categories.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `outMemberId` | String | no | User's unique ID within the organization. No more than 64 characters, consisting of digits and letters. | 23423532fwef |
| `language` | java.lang.String | yes | Language. See the enum in the FAQ. | ja |
| `categoryId` | java.lang.Long | yes | Category ID | 0 |
| `parentCateId` | Long | no | Parent ID | 0 |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:category.translation.getById.ResultModel](#m-category-translation-getbyid-resultmodel) | yes | Return result | 如下 |

<a id="m-category-translation-getbyid-resultmodel"></a>
#### category.translation.getById.ResultModel

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `success` | boolean | yes | Whether successful | true |
| `code` | java.lang.String | yes | Error code | S0000 |
| `message` | java.lang.String | yes | Error description | 成功 |
| `result` | [message:category.translation.getById.Category](#m-category-translation-getbyid-category) | yes | Actual result | 如下 |

<a id="m-category-translation-getbyid-category"></a>
#### category.translation.getById.Category

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `categoryId` | java.lang.Long | yes | Category ID | 1031910 |
| `chineseName` | java.lang.String | yes | Category name in Chinese | 连衣裙 |
| `translatedName` | java.lang.String | yes | Translated category name | ワンピース |
| `language` | java.lang.String | yes | Language | ja |
| `leaf` | java.lang.String | yes | Whether it is a leaf category | true |
| `level` | java.lang.String | yes | Category level | 2 |
| `parentCateId` | java.lang.String | yes | Parent category ID | 10166 |
| `fromCache` | java.lang.Boolean | yes | Whether the data comes from cache | true |
| `children` | [message:category.translation.getById.ChildCategory[]](#m-category-translation-getbyid-childcategory[]) | yes | Sub-category data | 如下 |

<a id="m-category-translation-getbyid-childcategory[]"></a>
#### category.translation.getById.ChildCategory[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `categoryId` | java.lang.Long | yes | Category ID | 1031910 |
| `chineseName` | java.lang.String | yes | Category name in Chinese | 连衣裙 |
| `translatedName` | java.lang.String | yes | Translated category name | ワンピース |
| `language` | java.lang.String | yes | Language | ja |
| `leaf` | java.lang.Boolean | yes | Whether it is a leaf category | true |
| `level` | java.lang.String | yes | Category level | 2 |
| `parentCateId` | java.lang.Long | yes | Parent category ID | 10166 |
| `fromCache` | java.lang.Boolean | yes | Whether the data comes from cache | true |

## Samples

**Input parameter example**

```
{
  2212505444921,
  "1",
  "ja",
  1031910
}
```

**Return value example**

```
{
    "result":
    {
        "fromCache": true,
        "children": null,
        "level": "2",
        "chineseName": "连衣裙",
        "language": "ja",
        "leaf": false,
        "translatedName": "ワンピース",
        "categoryId": 1031910,
        "parentCateId": 10166
    },
    "success": true,
    "code": "S0000",
    "message": "成功"
}
```
