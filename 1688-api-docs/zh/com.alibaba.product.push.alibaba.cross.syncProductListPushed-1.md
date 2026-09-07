# 跨境场景下将商品加入铺货列表

API: `com.alibaba.product.push:alibaba.cross.syncProductListPushed:1` · Category: 选品铺货  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.product.push:alibaba.cross.syncProductListPushed-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.product.push/alibaba.cross.syncProductListPushed/{appKey}`  
需要授权 (access_token) · 需要签名

跨境场景专用，将商品加入铺货列表（即：新增铺货关系），单次操作上限20条。加入后才可以通过商品详情查询接口查询商品详情。调用前需要联系跨境小二手动配置权限。

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `productIdList` | Long[] | 是 | 1688的商品ID列表,列表长度不能超过20个 | [123456] |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:alibaba.panama.commonResult](#m-alibaba-panama-commonresult) | 是 | 同步结果 | { "success": false} |

<a id="m-alibaba-panama-commonresult"></a>
#### alibaba.panama.commonResult

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `errorCode` | java.lang.String | 是 | 错误码 |  |
| `errorMsg` | java.lang.String | 是 | 错误信息 |  |
| `success` | boolean | 是 | 是否成功 |  |

## 错误码

| 错误码 | 现象 | 解决方案 |
|---|---|---|
| OFFERIDS_IS_NULL | 商品ID列表参数为空 | 检查入参商品是否存在 |
| APPKEY_ISV_CONFIG_IS_NOT_EXIST | 平台appkey配置不存在 | 联系运营在平台添加配置 |
| ISVUSERS_IS_NULL | ISV端用户帐号列表为空 | 同步Isv账户信息 |

## 示例

**入参示例**

```
[526415838394]
```
