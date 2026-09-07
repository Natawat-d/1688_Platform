# 根据类目Id查询类目

API: `com.alibaba.product:alibaba.category.get:1` · Category: 选品铺货  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.product:alibaba.category.get-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.product/alibaba.category.get/{appKey}`  
无需授权 · 需要签名

类目查询。如果需要获取所有1688类目信息，需要从根类目开始遍历获取整个类目树。即：先传0获取所有一级类目ID，然后在通过获取到的一级类目ID遍历获取所二级类目，最后通过遍历二级类目ID获取三级类目。注意：1688类目仅三级，三级类目即发布商品所需的叶子类目。

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 否 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `categoryID` | Long | 是 | 类目id,必须大于等于0， 如果为0，则查询所有一级类目 | 1031910 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `succes` | String | 是 | 是否成功 | true |
| `categoryInfo` | [message:alibaba.category.CategoryInfo[]](#m-alibaba-category-categoryinfo[]) | 是 | 类目列表 | {} |
| `errorMsg` | String | 是 | 错误信息 |   |
| `errorCode` | String | 是 | 错误码 | 500_1 |

<a id="m-alibaba-category-categoryinfo[]"></a>
#### alibaba.category.CategoryInfo[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `categoryID` | Long | 是 | 类目ID | 123456 |
| `name` | String | 是 | 类目名称 | 连衣裙 |
| `level` | Integer | 是 | 类目层级，1688无此内容 | 0 |
| `isLeaf` | Boolean | 是 | 是否叶子类目（只有叶子类目才能发布商品） | false |
| `parentIDs` | Long[] | 是 | 父类目ID数组,1688只返回一个父id | [23] |
| `childIDs` | Long[] | 是 | 子类目ID数组，1688无此内容 | [] |
| `childCategorys` | [message:alibaba.child.category.info[]](#m-alibaba-child-category-info[]) | 是 | 子类目信息 | "" |
| `minOrderQuantity` | Long | 是 | 最小起订量 | 1 |
| `featureInfos` | [message:alibaba.category.categoryInfo.FeatureInfo[]](#m-alibaba-category-categoryinfo-featureinfo[]) | 是 | 类目特征信息 | [{}] |
| `categoryType` | String | 是 | 类目的类型，1为1688大市场类目，2为1688工业品专业化类目，3为1688主流商品类目 | 1 |
| `isSupportProcessing` | Boolean | 是 | 类目是否支持加工定制 | true |

<a id="m-alibaba-child-category-info[]"></a>
#### alibaba.child.category.info[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `id` | Long | 是 | 子类目ID | 54 |
| `name` | String | 是 | 子类目名称 | 服饰配件、饰品 |
| `isLeaf` | Boolean | 是 | 是否叶子类目（只有叶子类目才能发布商品） | true |
| `categoryType` | String | 是 | 类目的类型：1表示cbu类目，2表示gallop类目 | 1 |

<a id="m-alibaba-category-categoryinfo-featureinfo[]"></a>
#### alibaba.category.categoryInfo.FeatureInfo[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `key` | String | 是 | 名称 | name |
| `value` | String | 是 | 值 | jiagong |
| `status` | Integer | 是 | 状态 | 0 |
| `hierarchy` | Boolean | 是 | 是否继承到子元素上 | true |

## 示例

**返回示例**

```
{}
```
