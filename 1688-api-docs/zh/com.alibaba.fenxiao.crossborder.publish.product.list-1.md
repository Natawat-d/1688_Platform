# 铺货商品列表

API: `com.alibaba.fenxiao.crossborder:publish.product.list:1` · Category: 铺货  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:publish.product.list-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/publish.product.list/{appKey}`  
需要授权 (access_token) · 需要签名

查询可铺货的商品列表，支持按货盘 ID 分页查询商品 ID 集合。

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `param` | [message:com.alibaba.global1688.silicon.user.api.param.WbPublishableOfferQueryParam](#m-com-alibaba-global1688-silicon-user-api-param-wbpublishableofferqueryparam) | 是 | 查询可铺货商品的请求参数 | {"palletId":"example","pageNo":1,"pageSize":10,"subjects":[578,571]} |

<a id="m-com-alibaba-global1688-silicon-user-api-param-wbpublishableofferqueryparam"></a>
#### com.alibaba.global1688.silicon.user.api.param.WbPublishableOfferQueryParam

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `pageNo` | java.lang.Integer | 是 | pageNo 字段 | 1 |
| `pageSize` | java.lang.Integer | 是 | pageSize 字段 | 1 |
| `palletId` | java.lang.String | 是 | palletId 字段 | example |
| `subjects` | Integer[] | 否 | WB 类目 subjectID 列表，可选；传入时仅返回对应类目的商品 | [578,571] |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:com.alibaba.global1688.silicon.user.api.common.ResultGTuskmls](#m-com-alibaba-global1688-silicon-user-api-common-resultgtuskmls) | 是 | 查询结果 | {"success":true,"code":"SUCCESS","result":{"offerIds":[1,2],"pageNo":1,"pageSize":10,"total":100}} |

<a id="m-com-alibaba-global1688-silicon-user-api-common-resultgtuskmls"></a>
#### com.alibaba.global1688.silicon.user.api.common.ResultGTuskmls

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `code` | java.lang.String | 是 | code 字段 | example |
| `message` | java.lang.String | 是 | message 字段 | example |
| `permissionName` | java.lang.String | 是 | permissionName 字段 | example |
| `result` | [message:com.alibaba.global1688.silicon.user.api.result.WbPublishableOfferPageDTO](#m-com-alibaba-global1688-silicon-user-api-result-wbpublishableofferpagedto) | 是 | result 字段 | {} |
| `success` | java.lang.Boolean | 是 | success 字段 | true |

<a id="m-com-alibaba-global1688-silicon-user-api-result-wbpublishableofferpagedto"></a>
#### com.alibaba.global1688.silicon.user.api.result.WbPublishableOfferPageDTO

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `offerIds` | java.lang.Long[] | 是 | offerIds 字段 | 1 |
| `pageNo` | java.lang.Integer | 是 | pageNo 字段 | 1 |
| `pageSize` | java.lang.Integer | 是 | pageSize 字段 | 1 |
| `total` | java.lang.Integer | 是 | total 字段 | 1 |

## 示例

**入参示例**

```
{
  "param": {
    "pageNo": 1,
    "pageSize": 1,
    "palletId": "example"
  }
}
```

**出参示例**

```
{
  "result": {
    "code": "example",
    "message": "example",
    "permissionName": "example",
    "result": {
      "offerIds": [
        1
      ],
      "pageNo": 1,
      "pageSize": 1,
      "total": 1
    },
    "success": true
  }
}
```
