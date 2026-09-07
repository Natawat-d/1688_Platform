# 【关系】分销商-查询供应商列表

API: `cn.alibaba.open:alibaba.relation.querySuppliers:1` · Category: 选品铺货  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=cn.alibaba.open:alibaba.relation.querySuppliers-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/cn.alibaba.open/alibaba.relation.querySuppliers/{appKey}`  
需要授权 (access_token) · 需要签名

通过分销商 userID 获取 供应商列表

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `supplierLoginId` | String | 否 | 供应商登录ID，指定该参数可以查询授权用户与指定供应商的分销关系 | 李定国 |
| `currentPage` | Integer | 否 | 当前页码 | 1 |
| `pageSize` | Integer | 否 | 每页多少个 | 10 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:alibaba.relation.suppliers-result](#m-alibaba-relation-suppliers-result) | 是 |  |  |

<a id="m-alibaba-relation-suppliers-result"></a>
#### alibaba.relation.suppliers-result

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `count` | Integer | 是 | 总共项数 |  |
| `currentPage` | Integer | 是 | 当前页码 |  |
| `pageSize` | Integer | 是 | 每页多少个 |  |
| `relationModels` | [message:alibaba.relation.supplierModel[]](#m-alibaba-relation-suppliermodel[]) | 是 | 结果集合 |  |

<a id="m-alibaba-relation-suppliermodel[]"></a>
#### alibaba.relation.supplierModel[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `consignStatus` | String | 是 | 合作状态 |  |
| `consignCreateTime` | Long | 是 | 代销合作创建时间 |  |
| `supplierLoginId` | String | 是 | 供应商登录Id |  |
| `supplierCompany` | String | 是 | 供应商公司名称 |  |
| `lastOrder` | Long | 是 | 近180交易订单 |  |
| `lastAmount` | Long | 是 | 近180交易金额（分） |  |
| `distributionNum` | Integer | 是 | 已铺货数量 |  |
| `memberId` | String | 是 | 会员memberId |  |

## 示例

**正确访问输出示例**

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
