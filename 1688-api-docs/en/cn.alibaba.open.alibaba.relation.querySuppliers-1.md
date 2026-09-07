# [Relationship] Distributor: query supplier list

Original name: 【关系】分销商-查询供应商列表  
API: `cn.alibaba.open:alibaba.relation.querySuppliers:1` · Category: Product Selection & Listing  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=cn.alibaba.open:alibaba.relation.querySuppliers-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/cn.alibaba.open/alibaba.relation.querySuppliers/{appKey}`  
Requires user authorization (access_token) · Requires signature

Get the supplier list for a distributor by userID.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `supplierLoginId` | String | no | Supplier login ID; specifying this parameter allows querying the distribution relationship between the authorized user and the specified supplier | 李定国 |
| `currentPage` | Integer | no | Current page number | 1 |
| `pageSize` | Integer | no | Number of items per page | 10 |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.relation.suppliers-result](#m-alibaba-relation-suppliers-result) | yes |  |  |

<a id="m-alibaba-relation-suppliers-result"></a>
#### alibaba.relation.suppliers-result

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `count` | Integer | yes | Total number of items |  |
| `currentPage` | Integer | yes | Current page number |  |
| `pageSize` | Integer | yes | Number of items per page |  |
| `relationModels` | [message:alibaba.relation.supplierModel[]](#m-alibaba-relation-suppliermodel[]) | yes | Result set |  |

<a id="m-alibaba-relation-suppliermodel[]"></a>
#### alibaba.relation.supplierModel[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `consignStatus` | String | yes | Cooperation status |  |
| `consignCreateTime` | Long | yes | Creation time of the consignment cooperation |  |
| `supplierLoginId` | String | yes | Supplier's login ID |  |
| `supplierCompany` | String | yes | Supplier company name |  |
| `lastOrder` | Long | yes | Orders over the last 180 days |  |
| `lastAmount` | Long | yes | Transaction amount over the last 180 days (in cents/fen) |  |
| `distributionNum` | Integer | yes | Quantity already listed |  |
| `memberId` | String | yes | Member's memberId |  |

## Samples

**Example output for a successful access**

```
{
  "relationModels": [
    {
      "supplierLoginId": "yqq001",
      "supplierId": 3636630767,
      "lastOrder": 0,
      "supplierCompany": "中国A&V有限公司",
      "distributionNum": 56,
      "consignCreateTime": "2016-04-07 11:48:15",
      "lastAmount": 0,
      "consignStatus": "aborting"
    },
    {
      "supplierLoginId": "panzeyitest18",
      "supplierId": 3687900876,
      "lastOrder": 0,
      "supplierCompany": "上海大凤姐有限公司",
      "distributionNum": 47,
      "consignCreateTime": "2015-12-02 17:42:47",
      "lastAmount": 0,
      "consignStatus": "normal"
    },
    {
      "supplierLoginId": "chaoqiang002",
      "supplierId": 3680118905,
      "lastOrder": 0,
      "supplierCompany": "上海大凤姐有限公司",
      "distributionNum": 2,
      "consignCreateTime": "2015-08-26 16:18:45",
      "lastAmount": 0,
      "consignStatus": "normal"
    },
    {
      "supplierLoginId": "鹿丹企业账户03",
      "supplierId": 3611768146,
      "lastOrder": 0,
      "supplierCompany": "中国A&V有限公司",
      "distributionNum": 0,
      "consignCreateTime": "2014-08-13 21:02:41",
      "lastAmount": 0,
      "consignStatus": "normal"
    }
  ],
  "count": 4,
  "pageSize": 10,
  "currentPage": 1
}
```
