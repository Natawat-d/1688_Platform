# Buyer submits return-shipment info

Original name: 买家提交退款货信息  
API: `com.alibaba.trade:alibaba.trade.refund.returnGoods:1` · Category: Returns & Refunds  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.trade:alibaba.trade.refund.returnGoods-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.trade/alibaba.trade.refund.returnGoods/{appKey}`  
Requires user authorization (access_token) · Requires signature

Used after the seller approves the buyer's return/refund request, for the buyer to submit the return-shipment information. First call alibaba.logistics.OpQueryLogisticCompanyList.offline to look up logistics companies, and use the logistics-company code it returns.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `refundId` | String | yes | Refund order number, starts with TQ | TQ36706338027991577 |
| `logisticsCompanyNo` | String | yes | Logistics company code, query via the alibaba.logistics.OpQueryLogisticCompanyList.offline API | ZTO |
| `freightBill` | String | yes | Logistics company waybill number. Please fill it in accurately, otherwise the seller has the right to refuse the refund. | 3110044550034338 |
| `description` | String | no | Shipping note; content must be between 2-200 characters | 发货说明 |
| `vouchers` | String[] | no | Voucher image URLs. Must use the “image domain/relative path” returned by the API alibaba.trade.uploadRefundVoucher. Up to 10 images can be uploaded; each image must not exceed 1MB; jpg, gif, jpeg, png, and bmp formats are supported. Please upload vouchers as they are required for a subsequent claim (a claim cannot be filed without them). | [https://cbu01.alicdn.com/img/ibank/2019/901/930/11848039109.jpg] |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.ocean.openplatform.biz.trade.result.RefundReturnGoodsResult](#m-alibaba-ocean-openplatform-biz-trade-result-refundreturngoodsresult) | yes | Return result | {"result":{"errorCode":"4004","errorInfo":"物流公司编号和运单号校验不通过","success":false}} |

<a id="m-alibaba-ocean-openplatform-biz-trade-result-refundreturngoodsresult"></a>
#### alibaba.ocean.openplatform.biz.trade.result.RefundReturnGoodsResult

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `errorCode` | java.lang.String | yes | Error code | 4002 |
| `errorInfo` | java.lang.String | yes | Error description | 物流公司编号和运单号必填 |
| `success` | boolean | yes | Whether the submission succeeded | true |

## Error codes

| Error | Symptom | How to fix |
|---|---|---|
| 4001 | User authorization required | Check whether the token has expired; refresh or obtain a new one |
| 4002 | Logistics company code and waybill number are required | Logistics company code and waybill number are required |
| 4003 | Too many uploaded vouchers/attachments | No more than 10 proof documents allowed |
| 4003 | Voucher URL verification failed | Voucher format does not meet requirements |
| 4004 | Logistics company code and waybill number validation failed | Logistics company code and waybill number validation failed; check and enter the correct waybill number |
| 4004 | Logistics company code and waybill number validation failed | Logistics company code and waybill number validation failed; check and enter the correct waybill number. |
| 4004 | Logistics company code and waybill number validation exception | Logistics company code and waybill number validation exception; please retry later |
| 4005 | Refund order information does not exist | Refund order information does not exist |
| 4006 | No permission to operate on the current refund order | No permission to operate on the current refund order; check whether the user account and the refund order number are correct |
| 4007 | The current refund order status does not support returns | The current refund order status does not support returns; the refund order status is incorrect |
| 4008 | Refund order number information is incorrect; please verify and resubmit | Refund order number information is incorrect; please verify and resubmit |
| 5001 | Exception querying refund order information | Exception querying refund order information |
| 5001 | Exception querying refund order information | Exception querying refund order information |
| 5002 | Refund/return operation failed | Refund/return operation failed |
| 5003 | Refund/return operation exception | Refund/return operation exception |

## Samples

**Description**

```
本API在买家退货场景下使用，需要卖家先同意退货申请，提交时物流公司的编码需要调用https://open.1688.com/api/apidocdetail.htm?aopApiCategory=Logistics_NEW&id=com.alibaba.logistics%3Aalibaba.logistics.OpQueryLogisticCompanyList.offline-1 获取，使用companyNo作为公司编码，运单号要填写正确，凭证上传使用https://open.1688.com/api/apidocdetail.htm?aopApiCategory=Logistics_NEW&id=com.alibaba.trade%3Aalibaba.trade.uploadRefundVoucher-1
```
