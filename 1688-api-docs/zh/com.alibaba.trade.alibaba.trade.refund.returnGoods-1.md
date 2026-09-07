# 买家提交退款货信息

API: `com.alibaba.trade:alibaba.trade.refund.returnGoods:1` · Category: 退货退款  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.trade.refund.returnGoods-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.trade.refund.returnGoods/{appKey}`  
需要授权 (access_token) · 需要签名

买家申请退货退款时，卖家同意后，买家提交退款货信息使用，需要先调用alibaba.logistics.OpQueryLogisticCompanyList.offline查询物流公司信息，使用接口返回的物流公司编码

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `refundId` | String | 是 | 退款单号，TQ开头 | TQ36706338027991577 |
| `logisticsCompanyNo` | String | 是 | 物流公司编码，调用alibaba.logistics.OpQueryLogisticCompanyList.offline接口查询 | ZTO |
| `freightBill` | String | 是 | 物流公司运单号，请准确填写，否则卖家有权拒绝退款 | 3110044550034338 |
| `description` | String | 否 | 发货说明，内容在2-200个字之间 | 发货说明 |
| `vouchers` | String[] | 否 | 凭证图片URLs，必须使用API alibaba.trade.uploadRefundVoucher返回的“图片域名/相对路径”，最多可上传 10 张图片 ；单张大小不超过1M；支持jpg、gif、jpeg、png、和bmp格式。 请上传凭证，以便以后续赔所需（不上传将无法理赔） | [https://cbu01.alicdn.com/img/ibank/2019/901/930/11848039109.jpg] |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:alibaba.ocean.openplatform.biz.trade.result.RefundReturnGoodsResult](#m-alibaba-ocean-openplatform-biz-trade-result-refundreturngoodsresult) | 是 | 返回结果 | {"result":{"errorCode":"4004","errorInfo":"物流公司编号和运单号校验不通过","success":false}} |

<a id="m-alibaba-ocean-openplatform-biz-trade-result-refundreturngoodsresult"></a>
#### alibaba.ocean.openplatform.biz.trade.result.RefundReturnGoodsResult

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `errorCode` | java.lang.String | 是 | 错误码 | 4002 |
| `errorInfo` | java.lang.String | 是 | 错误描述 | 物流公司编号和运单号必填 |
| `success` | boolean | 是 | 是否提交成功 | true |

## 错误码

| 错误码 | 现象 | 解决方案 |
|---|---|---|
| 4001 | 需要用户授权 | 检查token是否过期，重新刷新或获取 |
| 4002 | 物流公司编号和运单号必填 | 物流公司编号和运单号必填 |
| 4003 | 上传凭证数量过多 | 凭证不能超过10张 |
| 4003 | 凭证URL校验不通过 | 凭证格式不符合要求 |
| 4004 | 物流公司编号和运单号校验失败 | 物流公司编号和运单号校验失败，检查并填写正确的运单号 |
| 4004 | 物流公司编号和运单号校验不通过 | 物流公司编号和运单号校验不通过，检查并填写正确的运单号 |
| 4004 | 物流公司编号和运单号校验异常 | 物流公司编号和运单号校验异常，可稍后重试下 |
| 4005 | 退款单信息不存在 | 退款单信息不存在 |
| 4006 | 无当前退款单操作权限 | 无当前退款单操作权限，检查用户账号是否正确，退款单号是否正确 |
| 4007 | 当前退款单状态不支持退货 | 当前退款单状态不支持退货，退款单状态不对 |
| 4008 | 退款单号信息有误，请核对后重新提交 | 退款单号信息有误，请核对后重新提交 |
| 5001 | 退款单信息查询异常 | 退款单信息查询异常 |
| 5001 | 退款单信息查询异常 | 退款单信息查询异常 |
| 5002 | 退款退货操作失败 | 退款退货操作失败 |
| 5003 | 退款退货操作异常 | 退款退货操作异常 |

## 示例

**说明**

```
本API在买家退货场景下使用，需要卖家先同意退货申请，提交时物流公司的编码需要调用https://open.1688.com/api/apidocdetail.htm?aopApiCategory=Logistics_NEW&id=com.alibaba.logistics%3Aalibaba.logistics.OpQueryLogisticCompanyList.offline-1 获取，使用companyNo作为公司编码，运单号要填写正确，凭证上传使用https://open.1688.com/api/apidocdetail.htm?aopApiCategory=Logistics_NEW&id=com.alibaba.trade%3Aalibaba.trade.uploadRefundVoucher-1
```
