# 接收外部打印UDF链接

API: `com.alibaba.fenxiao.crossborder:label.url.receive:1` · Category: 物流  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:label.url.receive-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/label.url.receive/{appKey}`  
无需授权 · 需要签名

接收外部打印UDF链接

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 否 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `orderId` | java.lang.String | 是 | 订单ID |  |
| `outOrderId` | java.lang.String | 否 | 外部订单ID |  |
| `printUrls` | java.lang.String | 是 | 面单打印链接 |  |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:label.url.receive.hsf.Result](#m-label-url-receive-hsf-result) | 是 |  |  |

<a id="m-label-url-receive-hsf-result"></a>
#### label.url.receive.hsf.Result

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `success` | java.lang.Boolean | 是 |  |  |
| `data` | java.lang.Boolean | 是 |  |  |
| `msg` | java.lang.String | 是 |  |  |
