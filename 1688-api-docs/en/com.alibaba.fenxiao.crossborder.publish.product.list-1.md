# Listable product list

Original name: 铺货商品列表  
API: `com.alibaba.fenxiao.crossborder:publish.product.list:1` · Category: Listing (Publishing)  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:publish.product.list-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/publish.product.list/{appKey}`  
Requires user authorization (access_token) · Requires signature

Query the list of products available for listing. Supports paginated queries of product-ID sets by assortment (goods-pool) ID.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `param` | [message:com.alibaba.global1688.silicon.user.api.param.WbPublishableOfferQueryParam](#m-com-alibaba-global1688-silicon-user-api-param-wbpublishableofferqueryparam) | yes | Request parameters for querying listable products | {"palletId":"example","pageNo":1,"pageSize":10,"subjects":[578,571]} |

<a id="m-com-alibaba-global1688-silicon-user-api-param-wbpublishableofferqueryparam"></a>
#### com.alibaba.global1688.silicon.user.api.param.WbPublishableOfferQueryParam

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `pageNo` | java.lang.Integer | yes | pageNo field | 1 |
| `pageSize` | java.lang.Integer | yes | pageSize field | 1 |
| `palletId` | java.lang.String | yes | palletId field | example |
| `subjects` | Integer[] | no | List of WB category subjectIDs, optional; when provided, only products in the corresponding categories are returned. | [578,571] |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:com.alibaba.global1688.silicon.user.api.common.ResultGTuskmls](#m-com-alibaba-global1688-silicon-user-api-common-resultgtuskmls) | yes | Query result | {"success":true,"code":"SUCCESS","result":{"offerIds":[1,2],"pageNo":1,"pageSize":10,"total":100}} |

<a id="m-com-alibaba-global1688-silicon-user-api-common-resultgtuskmls"></a>
#### com.alibaba.global1688.silicon.user.api.common.ResultGTuskmls

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `code` | java.lang.String | yes | code field | example |
| `message` | java.lang.String | yes | message field | example |
| `permissionName` | java.lang.String | yes | permissionName field | example |
| `result` | [message:com.alibaba.global1688.silicon.user.api.result.WbPublishableOfferPageDTO](#m-com-alibaba-global1688-silicon-user-api-result-wbpublishableofferpagedto) | yes | result field | {} |
| `success` | java.lang.Boolean | yes | success field | true |

<a id="m-com-alibaba-global1688-silicon-user-api-result-wbpublishableofferpagedto"></a>
#### com.alibaba.global1688.silicon.user.api.result.WbPublishableOfferPageDTO

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `offerIds` | java.lang.Long[] | yes | offerIds field | 1 |
| `pageNo` | java.lang.Integer | yes | pageNo field | 1 |
| `pageSize` | java.lang.Integer | yes | pageSize field | 1 |
| `total` | java.lang.Integer | yes | total field | 1 |

## Samples

**Input parameter example**

```
{
  "param": {
    "pageNo": 1,
    "pageSize": 1,
    "palletId": "example"
  }
}
```

**Output parameter example**

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
