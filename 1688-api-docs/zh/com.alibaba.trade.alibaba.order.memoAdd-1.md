# 修改订单备忘

API: `com.alibaba.trade:alibaba.order.memoAdd:1` · Category: 订单  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.order.memoAdd-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.order.memoAdd/{appKey}`  
需要授权 (access_token) · 需要签名

授权用户为卖家修改卖家备忘，授权用户为买家修改买家备忘
注意：该接口可重复调用，备注内容将覆盖前一次调用

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `orderId` | java.lang.Long | 是 | 订单ID | 1234567 |
| `memo` | java.lang.String | 是 | 备忘信息 | 订单备忘详情 |
| `remarkIcon` | String | 是 | 备忘图标，目前仅支持数字。1位红色图标，2为蓝色图标，3为绿色图标，4为黄色图标 | 2 |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `success` | java.lang.Boolean | 是 | 是否成功 | true |
| `errorCode` | java.lang.String | 是 | 错误编码 | 400_4 |
| `errorMsg` | java.lang.String | 是 | 错误信息 | 修改失败 |

## 示例

**订单不存在出参示例**

```
{
  "errorCode": "400_4",
  "errorMsg": "修改备注信息出错[ORDER_NOT_EXIST]，order[1198263322216969811]is not exsits",
  "success": true
}
```
