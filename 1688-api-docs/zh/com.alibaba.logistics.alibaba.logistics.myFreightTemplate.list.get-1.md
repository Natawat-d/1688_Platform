# 获取物流模板详情

API: `com.alibaba.logistics:alibaba.logistics.myFreightTemplate.list.get:1` · Category: 物流  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.logistics:alibaba.logistics.myFreightTemplate.list.get-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.logistics/alibaba.logistics.myFreightTemplate.list.get/{appKey}`  
需要授权 (access_token) · 需要签名

根据物流模版ID获取卖家的物流模板。运费模板ID为0表示运费说明，为1表示卖家承担运费

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `templateId` | java.lang.Long | 否 | 模版id，用于单条查询的场景 | xxx |
| `querySubTemplate` | java.lang.Boolean | 否 | 是否查询子模板 | false |
| `queryRate` | java.lang.Boolean | 否 | 是否查询子模板费率 | false |

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:alibaba.logistics.FreightTemplate[]](#m-alibaba-logistics-freighttemplate[]) | 是 | 返回结果 | [] |
| `errorCode` | java.lang.String | 是 | 错误码 | 错误码 |
| `errorMsg` | java.lang.String | 是 | 错误描述 | 错误描述 |

<a id="m-alibaba-logistics-freighttemplate[]"></a>
#### alibaba.logistics.FreightTemplate[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `addressCodeText` | java.lang.String | 是 | 地址信息 | xxx |
| `fromAreaCode` | java.lang.String | 是 | 发货地址地区码 | xxx |
| `id` | java.lang.Long | 是 | 地址ID | xxx |
| `memberId` | java.lang.String | 是 | 会员ID | xxx |
| `name` | java.lang.String | 是 | 名称 | xxx |
| `remark` | java.lang.String | 是 | 备注 | xxx |
| `status` | java.lang.Integer | 是 | 状态 | xxx |
| `expressSubTemplate` | [message:alibaba.openplatform.logistics.DeliverySubTemplateDetailDTO](#m-alibaba-openplatform-logistics-deliverysubtemplatedetaildto) | 是 | 快递子模版 | xxx |
| `logisticsSubTemplate` | [message:alibaba.openplatform.logistics.DeliverySubTemplateDetailDTO](#m-alibaba-openplatform-logistics-deliverysubtemplatedetaildto) | 是 | 货运子模版 | xxx |
| `codSubTemplate` | [message:alibaba.openplatform.logistics.DeliverySubTemplateDetailDTO](#m-alibaba-openplatform-logistics-deliverysubtemplatedetaildto) | 是 | 货到付款子模版 | xxx |
| `type` | String | 是 | 类型 3-官方物流模板，2或无值时-用户物流模板 | 3 |

<a id="m-alibaba-openplatform-logistics-deliverysubtemplatedetaildto"></a>
#### alibaba.openplatform.logistics.DeliverySubTemplateDetailDTO

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `subTemplateDTO` | [message:alibaba.openplatform.logistics.DeliverySubTemplateDTO](#m-alibaba-openplatform-logistics-deliverysubtemplatedto) | 是 | 子模板 |  |
| `rateList` | [message:alibaba.openplatform.logistics.DeliveryRateDetailDTO[]](#m-alibaba-openplatform-logistics-deliveryratedetaildto[]) | 是 | 费率 |  |

<a id="m-alibaba-openplatform-logistics-deliverysubtemplatedto"></a>
#### alibaba.openplatform.logistics.DeliverySubTemplateDTO

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `chargeType` | java.lang.Integer | 是 | 计件类型。0:重量 1:件数 2:体积 |  |
| `isSysTemplate` | java.lang.Boolean | 是 | 是否系统模板 |  |
| `serviceChargeType` | java.lang.Integer | 是 | 运费承担类型 卖家承担：0；买家承担：1。 |  |
| `serviceType` | java.lang.Integer | 是 | 服务类型。0:快递 1:货运 2:货到付款 |  |
| `type` | java.lang.Integer | 是 | 子模板类型 0基准 1增值。默认0。 |  |

<a id="m-alibaba-openplatform-logistics-deliveryratedetaildto[]"></a>
#### alibaba.openplatform.logistics.DeliveryRateDetailDTO[]

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `isSysRate` | boolean | 是 | 是否系统模板 |  |
| `toAreaCodeText` | java.lang.String | 是 | 地址编码文本，用顿号隔开。例如：上海、福建省、广东省 |  |
| `rateDTO` | [message:alibaba.openplatform.logistics.DeliveryRateDTO](#m-alibaba-openplatform-logistics-deliveryratedto) | 是 | 普通子模板费率 |  |
| `sysRateDTO` | [message:alibaba.openplatform.logistics.DeliverySysRateDTO](#m-alibaba-openplatform-logistics-deliverysysratedto) | 是 | 系统子模板费率 |  |

<a id="m-alibaba-openplatform-logistics-deliveryratedto"></a>
#### alibaba.openplatform.logistics.DeliveryRateDTO

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `firstUnit` | java.lang.Long | 是 | 首重（单位：克）或首件（单位：件） |  |
| `firstUnitFee` | java.lang.Long | 是 | 首重或首件的价格 |  |
| `leastExpenses` | java.lang.Long | 是 | 最低一票 |  |

<a id="m-alibaba-openplatform-logistics-deliverysysratedto"></a>
#### alibaba.openplatform.logistics.DeliverySysRateDTO

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `firstUnit` | java.lang.Long | 是 | 首重（单位：克）或首件（单位：件） |  |
| `firstUnitFee` | java.lang.Long | 是 | 首重或首件的价格 |  |
| `leastExpenses` | java.lang.Long | 是 | 最低一票 |  |
| `nextUnit` | java.lang.Long | 是 | 续重（单位：克）或续件（单位：件）单位 |  |

## 示例

**出参示例**

```
      
{
    "result":[
        {
            "addressCodeText":"湖北省 荆门市 东宝区",
            "fromAreaCode":"420802",
            "id":11864709,
            "memberId":"b2b-1624961198",
            "name":"test168",
            "remark":"",
            "status":1
        }
    ]
}
```
