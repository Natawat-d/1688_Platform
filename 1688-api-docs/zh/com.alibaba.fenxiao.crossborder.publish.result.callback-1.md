# 发布结果回调

API: `com.alibaba.fenxiao.crossborder:publish.result.callback:1` · Category: 铺货  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:publish.result.callback-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/publish.result.callback/{appKey}`  
需要授权 (access_token) · 需要签名

WB-1688 集成 Method 3 发布结果回调接口

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `request` | [message:alibaba.global1688.silicon.user.api.param.WbPublishCallbackRequestDTO](#m-alibaba-global1688-silicon-user-api-param-wbpublishcallbackrequestdto) | 是 | 回调请求参数 | {"offerId":123456789,"imtId":"imt123","status":"SUCCESS","requestId":"req123","cards":[]} |

<a id="m-alibaba-global1688-silicon-user-api-param-wbpublishcallbackrequestdto"></a>
#### alibaba.global1688.silicon.user.api.param.WbPublishCallbackRequestDTO

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `offerId` | java.lang.Long | 是 | 1688商品ID | 123456789 |
| `imtId` | java.lang.String | 是 | WB商品IMT ID | imt123 |
| `status` | java.lang.String | 是 | 发布状态 | SUCCESS |
| `requestId` | java.lang.String | 是 | 请求ID | req123 |
| `cards` | [message:alibaba.global1688.silicon.user.api.param.WbPublishCallbackRequestDTO.WbCardResult[]](#m-alibaba-global1688-silicon-user-api-param-wbpublishcallbackrequestdto-wbcardresult[]) | 是 | 卡片结果列表 | [] |

<a id="m-alibaba-global1688-silicon-user-api-param-wbpublishcallbackrequestdto-wbcardresult[]"></a>
#### alibaba.global1688.silicon.user.api.param.WbPublishCallbackRequestDTO.WbCardResult[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:ResultGT52f0rf](#m-resultgt52f0rf) | 是 | 回调结果 | {"success":true,"duplicate":false,"requestId":"req123"} |

<a id="m-resultgt52f0rf"></a>
#### ResultGT52f0rf

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `success` | java.lang.Boolean | 是 | 是否成功 | true |
| `result` | [message:alibaba.global1688.silicon.user.api.result.WbPublishCallbackResponseDTO](#m-alibaba-global1688-silicon-user-api-result-wbpublishcallbackresponsedto) | 是 | 返回结果 | {} |
| `code` | java.lang.String | 否 | 错误码 |  |
| `permissionName` | java.lang.String | 否 | 权限名称 |  |
| `message` | java.lang.String | 否 | 错误信息 |  |

<a id="m-alibaba-global1688-silicon-user-api-result-wbpublishcallbackresponsedto"></a>
#### alibaba.global1688.silicon.user.api.result.WbPublishCallbackResponseDTO

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `success` | java.lang.Boolean | 是 | 是否成功 | true |
| `duplicate` | java.lang.Boolean | 是 | 是否重复请求 | false |
| `requestId` | java.lang.String | 是 | 请求ID | req123 |

## 示例

**入参示例**

```
{"request":{"offerId":123456789,"imtId":"imt123","status":"SUCCESS","requestId":"req123","cards":[]}}
```

**出参示例**

```
{"result":{"success":true,"duplicate":false,"requestId":"req123"}}
```
