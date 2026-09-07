# 解除关注商品

API: `com.alibaba.product:alibaba.product.unfollow.crossborder:1` · Category: 选品铺货  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.product:alibaba.product.unfollow.crossborder-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.product/alibaba.product.unfollow.crossborder/{appKey}`  
需要授权 (access_token) · 需要签名

解除关注商品

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `productId` | Long | 是 | 商品id | 36143645361 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `code` | int | 是 | code,0表示成功 | 0 |
| `message` | String | 是 | 结果的描述 | success |

## 示例

**出参示例**

```
{
  "code": 0,
  "message": "success"
}
```
