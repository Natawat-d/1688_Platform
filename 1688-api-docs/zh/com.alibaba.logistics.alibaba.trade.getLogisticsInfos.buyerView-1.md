# 获取交易订单的物流信息(买家视角)

API: `com.alibaba.logistics:alibaba.trade.getLogisticsInfos.buyerView:1` · Category: 订单  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.logistics:alibaba.trade.getLogisticsInfos.buyerView-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.logistics/alibaba.trade.getLogisticsInfos.buyerView/{appKey}`  
需要授权 (access_token) · 需要签名

该接口需要获得订单买家的授权，获取买家的订单的物流详情，在采购或者分销场景中，作为买家也有获取物流详情的需求。该接口能查能根据订单号查看物流详情，包括发件人，收件人，所发货物明细等。由于物流单录入的原因，可能跟踪信息的API查询会有延迟。该API需要向开放平台申请权限才能访问。In the procurement or distribution scenario, buyers can ask for obtaining the logistics details. The interface can check the logistics details according to the order ID, including the sender, the recipient, the details of the goods sent, and so on. Depending on the logistics information entry time, there may be a delay in API queries regarding the information tracking.

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `orderId` | Long | 是 | 订单号 | 1221434 |
| `fields` | String | 否 | 需要返回的字段，目前有:company.name,sender,receiver,sendgood。返回的字段要用英文逗号分隔开 | company,name,sender,receiver,sendgood |
| `webSite` | String | 是 | 是1688业务还是icbu业务 | 1688或者alibaba |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:alibaba.logistics.OpenPlatformLogisticsOrder[]](#m-alibaba-logistics-openplatformlogisticsorder[]) | 是 | 返回结果 | {} |
| `errorCode` | String | 是 | 错误码 |   |
| `errorMessage` | String | 是 | 错误信息 |   |
| `success` | Boolean | 是 | 是否成功 | true |

<a id="m-alibaba-logistics-openplatformlogisticsorder[]"></a>
#### alibaba.logistics.OpenPlatformLogisticsOrder[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `logisticsId` | String | 是 | 物流信息ID | 物流信息ID |
| `logisticsBillNo` | String | 是 | 物流单号，运单号 | 物流单号，运单号 |
| `orderEntryIds` | String | 是 | 订单号列表，无子订单的等于主订单编号，否则为对应子订单列表 | 129232515787615400,129232515788615400,129232515789615400,129232515790615400 |
| `status` | String | 是 | 物流状态。WAITACCEPT:未受理;CANCEL:已撤销;ACCEPT:已受理;TRANSPORT:运输中;NOGET:揽件失败;SIGN:已签收;UNSIGN:签收异常 | WAITACCEPT |
| `logisticsCompanyId` | String | 是 | 物流公司ID | 物流公司ID |
| `logisticsCompanyName` | String | 是 | 物流公司编码 | 物流公司编码 |
| `logisticsCompanyNo` | String | 是 | 物流公司编号 | 物流公司编号 |
| `remarks` | String | 是 | 备注 | 备注 |
| `serviceFeature` | String | 是 | serviceFeature | serviceFeature |
| `gmtSystemSend` | String | 是 | gmtSystemSend | gmtSystemSend	 |
| `sendGoods` | [message:alibaba.logistics.OpenPlatformLogisticsSendGood[]](#m-alibaba-logistics-openplatformlogisticssendgood[]) | 是 | 商品信息 | 商品信息 |
| `receiver` | [message:alibaba.logistics.OpenPlatformLogisticsReceiver](#m-alibaba-logistics-openplatformlogisticsreceiver) | 是 | 收件人信息 | 收件人信息 |
| `sender` | [message:alibaba.logistics.OpenPlatformLogisticsSender](#m-alibaba-logistics-openplatformlogisticssender) | 是 | 发件人信息 | 发件人信息 |
| `logisticsOrderGoods` | [message:com.alibaba.ocean.openplatform.biz.logistics.common.model.OpenPlatformLogisticsOrderSendGood[]](#m-com-alibaba-ocean-openplatform-biz-logistics-common-model-openplatformlogisticsordersendgood[]) | 是 | 物流订单跟商品关系模型 | {} |
| `logisticsOrderSendGood` | [message:com.alibaba.ocean.openplatform.biz.logistics.common.model.seller.OpenPlatformLogisticsOrderSendGood[]](#m-com-alibaba-ocean-openplatform-biz-logistics-common-model-seller-openplatformlogisticsordersendgood[]) | 是 | 物流和订单关联模型 | {} |

<a id="m-alibaba-logistics-openplatformlogisticssendgood[]"></a>
#### alibaba.logistics.OpenPlatformLogisticsSendGood[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `goodName` | String | 是 | 商品名 |  |
| `quantity` | String | 是 | 商品数量 |  |
| `unit` | String | 是 | 商品单位 |  |

<a id="m-alibaba-logistics-openplatformlogisticsreceiver"></a>
#### alibaba.logistics.OpenPlatformLogisticsReceiver

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `receiverName` | String | 是 | 收件人名字 |  |
| `receiverPhone` | String | 是 | 收件人电话 |  |
| `receiverMobile` | String | 是 | 收件人电话 |  |
| `encrypt` | String | 是 |  |  |
| `receiverProvinceCode` | String | 是 | 省编码 |  |
| `receiverCityCode` | String | 是 | 市编码 |  |
| `receiverCountyCode` | String | 是 | 国家编码 |  |
| `receiverAddress` | String | 是 | 地址 |  |
| `receiverProvince` | String | 是 | 省份 |  |
| `receiverCity` | String | 是 | 城市 |  |
| `receiverCounty` | String | 是 | 国家 |  |

<a id="m-alibaba-logistics-openplatformlogisticssender"></a>
#### alibaba.logistics.OpenPlatformLogisticsSender

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `senderName` | String | 是 | 发件人姓名 |  |
| `senderPhone` | String | 是 | 发件人电话 |  |
| `senderMobile` | String | 是 | 发件人电话 |  |
| `encrypt` | String | 是 |  |  |
| `senderProvinceCode` | String | 是 | 省编码 |  |
| `senderCityCode` | String | 是 | 城市编码 |  |
| `senderCountyCode` | String | 是 | 国家编码 |  |
| `senderAddress` | String | 是 | 发货人地址 |  |
| `senderProvince` | String | 是 | 省份 |  |
| `senderCity` | String | 是 | 城市 |  |
| `senderCounty` | String | 是 | 国家 |  |

<a id="m-com-alibaba-ocean-openplatform-biz-logistics-common-model-openplatformlogisticsordersendgood[]"></a>
#### com.alibaba.ocean.openplatform.biz.logistics.common.model.OpenPlatformLogisticsOrderSendGood[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `logisticsId` | String | 是 | 物流编号 | LP00616288919385 |
| `tradeOrderId` | Long | 是 | 交易主订单号 | 3657499272179232333 |
| `tradeOrderItemId` | Long | 是 | 交易子订单号 | 3657499272179232333 |
| `description` | String | 是 | sku描述 | 颜色: 黑色; 尺码: M; |
| `quantity` | Double | 是 | 数量 | 1 |
| `unit` | String | 是 | 单位 | 件 |
| `productName` | String | 是 | 商品名称 | 品名 |

<a id="m-com-alibaba-ocean-openplatform-biz-logistics-common-model-seller-openplatformlogisticsordersendgood[]"></a>
#### com.alibaba.ocean.openplatform.biz.logistics.common.model.seller.OpenPlatformLogisticsOrderSendGood[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `logisticsId` | String | 是 | 物流编号 | LP00616288919385 |
| `tradeOrderId` | Long | 是 | 交易主订单号 | 3657499272179232333 |
| `tradeOrderItemId` | Long | 是 | 交易子订单号 | 3657499272179232333 |
| `description` | String | 是 | sku描述 | 颜色: 黑色; 尺码: M; |
| `quantity` | Double | 是 | 数量 | 1 |
| `unit` | String | 是 | 单位 | 件 |
| `productName` | String | 是 | 商品名称 | 品名 |

## 错误码

| 错误码 | 现象 | 解决方案 |
|---|---|---|
| 500_2 | 订单尚未发货，暂无物流详情，请稍候再试。 | 订单尚未发货，建议监听订单发货消息，收到消息后查询订单物流状态。参考：https://open.1688.com/doc/topicDetail.htm?spm=a260s.11630592.0.0.555655ed4QUvy5&topicGroup=ORDER&id=ORDER_BUYER_VIEW_ANNOUNCE_SENDGOODS |

## 示例

**返回结果**

```
{
  "result": [
    {
      "logisticsId": "BX111841674232006",
      "orderEntryIds": "149989279191615400",
      "status": "SIGN",
      "logisticsCompanyId": "8",
      "remarks": "asdasdasdasdsad啊啊啊",
      "sendGoods": [
        {
          "unit": "个",
          "quantity": 2,
          "goodName": "【茂茂回归】12.9公开-sku区间价-淘货源"
        }
      ],
      "receiver": {
        "receiverCountyCode": "310105",
        "receiverCity": "上海市",
        "receiverProvinceCode": "310000",
        "receiverCityCode": "310100",
        "encrypt": "CN",
        "receiverCounty": "长宁区",
        "receiverMobile": "13800138000",
        "receiverProvince": "上海",
        "receiverName": "XXX",
        "receiverAddress": "XXX街道 XX路999",
        "receiverPhone": ""
      },
      "sender": {
        "senderCityCode": "330100",
        "senderAddress": "网商路699号",
        "senderName": "张三",
        "senderProvince": "浙江省",
        "encrypt": "CN",
        "senderCity": "杭州市",
        "senderCountyCode": "330108",
        "senderProvinceCode": "330000",
        "senderMobile": "12345678900",
        "senderPhone": "",
        "senderCounty": "滨江区"
      }
    }
  ]
}
```
