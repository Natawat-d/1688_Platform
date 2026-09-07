# Receive external print-label (UDF) URL

Original name: 接收外部打印UDF链接  
API: `com.alibaba.fenxiao.crossborder:label.url.receive:1` · Category: Logistics  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:label.url.receive-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/label.url.receive/{appKey}`  
No user authorization · Requires signature

Receive an external print-label (UDF) URL.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | no | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `orderId` | java.lang.String | yes | Order ID |  |
| `outOrderId` | java.lang.String | no | External order ID |  |
| `printUrls` | java.lang.String | yes | Shipping label print link |  |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:label.url.receive.hsf.Result](#m-label-url-receive-hsf-result) | yes |  |  |

<a id="m-label-url-receive-hsf-result"></a>
#### label.url.receive.hsf.Result

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `success` | java.lang.Boolean | yes |  |  |
| `data` | java.lang.Boolean | yes |  |  |
| `msg` | java.lang.String | yes |  |  |
