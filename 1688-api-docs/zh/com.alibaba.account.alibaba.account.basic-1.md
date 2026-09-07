# 获取授权用户的基本信息

API: `com.alibaba.account:alibaba.account.basic:1` · Category: 会员  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.account:alibaba.account.basic-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.account/alibaba.account.basic/{appKey}`  
需要授权 (access_token) · 需要签名

获取授权用户的基本信息

## 系统级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `_aop_timestamp` | String | 否 | 请求时间戳 |  |
| `_aop_signature` | String | 是 | 请求签名 |  |
| `access_token` | String | 是 | 用户授权令牌 |  |

## 应用级参数

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|

## 返回结果

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `result` | [message:alibaba.account.simpleAccountInfo](#m-alibaba-account-simpleaccountinfo) | 是 | 会员信息 | {} |
| `errorCode` | String | 是 | 错误码 |   |
| `errorMessage` | String | 是 | 错误信息 |   |
| `success` | Boolean | 是 | 是否成功 | true |

<a id="m-alibaba-account-simpleaccountinfo"></a>
#### alibaba.account.simpleAccountInfo

| 字段 | 类型 | 必填 | 描述 | 示例 |
|---|---|---|---|---|
| `loginId` | java.lang.String | 是 | 登录名 |  |
| `saleRate` | java.lang.Long | 是 | 卖家等级 |  |
| `maturity` | int | 是 | 会员成熟度 |  |
| `memo` | java.lang.String | 是 | 备注 |  |
| `modifyDate` | java.util.Date | 是 | 修改日期 |  |
| `categoryName` | java.lang.String | 是 | 主营行业的类目名称 |  |
| `trustScore` | int | 是 | 诚信通指数 |  |
| `userId` | java.lang.Long | 是 | 用户ID |  |
| `enterpriseAccount` | boolean | 是 | 是否是公司会员 |  |
| `createDate` | java.util.Date | 是 | 创建时间 |  |
| `communityLevel` | java.lang.String | 是 | 贸易通用户标识 |  |
| `joinFrom` | java.lang.String | 是 | 注册来源 |  |
| `rateNum` | java.lang.Long | 是 | 买家的信用分值(淘宝) |  |
| `gmtPaidJoin` | java.util.Date | 是 | 诚信通加入时间 |  |
| `buyRate` | java.lang.Long | 是 | 买家等级 |  |
| `categoryId` | int | 是 | 主营行业的类目id |  |
| `companyName` | java.lang.String | 是 | 公司名 |  |
| `personAccount` | boolean | 是 | 是否是个人账号 |  |
| `homepageUrl` | java.lang.String | 是 | 公司主页 |  |
| `saleKeywords` | java.lang.String | 是 | 出售关键字 |  |
| `tpYear` | int | 是 | 诚信通年限 |  |
| `buyKeywords` | java.lang.String | 是 | 求购关键字 |  |
| `memberBizType` | java.lang.String | 是 | 业务类型<br> 企业单位 ENTERPRISE(&quot;5&quot;),<br>个体经营 SELF_EMPLOYED(&quot;6&quot;),<br> 事业单位或社会团体 ORGANIZATION(&quot;7&quot;),<br> 个人 PERSONAL(&quot;8&quot;); |  |
| `rateSum` | java.lang.Long | 是 | 卖家的信用分值(淘宝) |  |
| `domainInPlatforms` | String[] | 是 | 1688上店铺地址 |  |
| `memberId` | java.lang.String | 是 | memberId |  |
| `shopUrl` | java.lang.String | 是 | 旺铺首页地址 |  |
| `supplierName` | java.lang.String | 是 | 供应商名称 |  |
| `icon` | java.lang.String | 是 | 用户头像 |  |
| `phoneNo` | java.lang.String | 是 | 固定电话 |  |
| `industry` | java.lang.String | 是 | 主营行业 |  |
| `product` | java.lang.String | 是 | 主营产品 |  |
| `department` | java.lang.String | 是 | 部门 |  |
| `mobileNo` | java.lang.String | 是 | 业务联系手机（会员基本信息中编辑）	<br>业务联系手机 |  |
| `addressLocation` | java.lang.String | 是 | 联系地址 |  |
| `email` | java.lang.String | 是 | 业务联系邮箱 |  |
| `sellerName` | java.lang.String | 是 | 卖家名称 |  |
| `kuaJingBao` | Boolean | 是 | 是否跨境宝 | true |
| `crossBorder` | Boolean | 是 | 是否跨境宝 | true |
| `isPm` | Boolean | 是 | 是否实商 | true |
| `pm` | Boolean | 是 | 是否实商 | true |
| `fm` | Boolean | 是 | 是否找工厂商家 | true |
| `isOfficalLogistics` | Boolean | 是 | 是否支持官方物流 | true |
| `status` | String | 是 | 子账号状态 | enabled |
| `openUid` | String | 是 | 开放平台加密用户Id | sE542234KSls4KSfsthjST |

## 示例

**出参示例**

```
{
  "result": {
    "loginId": "alitestforisv01",
    "maturity": 0,
    "categoryName": "女装",
    "trustScore": 0,
    "userId": 1623492085,
    "enterpriseAccount": true,
    "createDate": "20130312160828000+0800",
    "joinFrom": "CJ_COMMON_JOIN_M",
    "rateNum": 0,
    "gmtPaidJoin": "20170502000942000+0800",
    "buyRate": 3,
    "categoryId": 10166,
    "companyName": "AOP对外测试账号01",
    "personAccount": false,
    "saleKeywords": "女装,男装,女鞋,衬衫",
    "tpYear": 0,
    "buyKeywords": "男装,女装,女鞋",
    "memberBizType": "PERSONAL",
    "rateSum": 0,
    "memberId": "b2b-1623492085",
    "supplierName": "AOP对外测试账号01",
    "icon": "/cms/upload/2011/116/401/104611_1301427272.png",
    "fm": false
  }
}
```
