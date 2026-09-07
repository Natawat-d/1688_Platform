# 根据类目ID查询多语言类目

API: `com.alibaba.fenxiao.crossborder:category.translation.getById:1` · Category: 类目  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:category.translation.getById-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/category.translation.getById/{appKey}`  
需要授权 (access_token) · 需要签名

多语言类目查询接口。根据当前语种和类目ID查询对应语种的类目详情，包含当前类目的下级类目列表数据。

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `outMemberId` | String | 否 | 用户在机构的唯一ID。不超过64位，由数字和字母组成。 | 23423532fwef |
| `language` | java.lang.String | 是 | 语种。见常见问题中的枚举。 | ja |
| `categoryId` | java.lang.Long | 是 | 类目ID | 0 |
| `parentCateId` | Long | 否 | 父id | 0 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:category.translation.getById.ResultModel](#m-category-translation-getbyid-resultmodel) | 是 | 返回结果 | 如下 |

<a id="m-category-translation-getbyid-resultmodel"></a>
#### category.translation.getById.ResultModel

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `success` | boolean | 是 | 是否成功 | true |
| `code` | java.lang.String | 是 | 错误码 | S0000 |
| `message` | java.lang.String | 是 | 错误描述 | 成功 |
| `result` | [message:category.translation.getById.Category](#m-category-translation-getbyid-category) | 是 | 实际结果 | 如下 |

<a id="m-category-translation-getbyid-category"></a>
#### category.translation.getById.Category

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `categoryId` | java.lang.Long | 是 | 类目ID | 1031910 |
| `chineseName` | java.lang.String | 是 | 类目中文名称 | 连衣裙 |
| `translatedName` | java.lang.String | 是 | 类目翻译名称 | ワンピース |
| `language` | java.lang.String | 是 | 语种 | ja |
| `leaf` | java.lang.String | 是 | 是否叶子类目 | true |
| `level` | java.lang.String | 是 | 类目层级 | 2 |
| `parentCateId` | java.lang.String | 是 | 上层类目ID | 10166 |
| `fromCache` | java.lang.Boolean | 是 | 数据是否来自缓存 | true |
| `children` | [message:category.translation.getById.ChildCategory[]](#m-category-translation-getbyid-childcategory[]) | 是 | 子类目数据 | 如下 |

<a id="m-category-translation-getbyid-childcategory[]"></a>
#### category.translation.getById.ChildCategory[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `categoryId` | java.lang.Long | 是 | 类目ID | 1031910 |
| `chineseName` | java.lang.String | 是 | 类目中文名称 | 连衣裙 |
| `translatedName` | java.lang.String | 是 | 类目翻译名称 | ワンピース |
| `language` | java.lang.String | 是 | 语种 | ja |
| `leaf` | java.lang.Boolean | 是 | 是否叶子类目 | true |
| `level` | java.lang.String | 是 | 类目层级 | 2 |
| `parentCateId` | java.lang.Long | 是 | 上层类目ID | 10166 |
| `fromCache` | java.lang.Boolean | 是 | 数据是否来自缓存 | true |

## 示例

**入参示例**

```
{
  2212505444921,
  "1",
  "ja",
  1031910
}
```

**返回值示例**

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
