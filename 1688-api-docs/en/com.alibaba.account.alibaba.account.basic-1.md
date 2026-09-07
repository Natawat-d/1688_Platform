# Get basic info of the authorized user

Original name: 获取授权用户的基本信息  
API: `com.alibaba.account:alibaba.account.basic:1` · Category: Membership & Accounts  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.account:alibaba.account.basic-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.account/alibaba.account.basic/{appKey}`  
Requires user authorization (access_token) · Requires signature

Get the basic information of the authorized user.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.account.simpleAccountInfo](#m-alibaba-account-simpleaccountinfo) | yes | Member information | {} |
| `errorCode` | String | yes | Error code |   |
| `errorMessage` | String | yes | Error message |   |
| `success` | Boolean | yes | Whether successful | true |

<a id="m-alibaba-account-simpleaccountinfo"></a>
#### alibaba.account.simpleAccountInfo

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `loginId` | java.lang.String | yes | Login name |  |
| `saleRate` | java.lang.Long | yes | Seller level |  |
| `maturity` | int | yes | Member maturity level |  |
| `memo` | java.lang.String | yes | Remark |  |
| `modifyDate` | java.util.Date | yes | Modification date |  |
| `categoryName` | java.lang.String | yes | Category name of the main industry |  |
| `trustScore` | int | yes | Chengxintong (TrustPass) index |  |
| `userId` | java.lang.Long | yes | User ID |  |
| `enterpriseAccount` | boolean | yes | Whether it is a company member |  |
| `createDate` | java.util.Date | yes | Creation time |  |
| `communityLevel` | java.lang.String | yes | TradeManager (Maoyitong) user identifier |  |
| `joinFrom` | java.lang.String | yes | Registration source |  |
| `rateNum` | java.lang.Long | yes | Buyer's credit score (Taobao) |  |
| `gmtPaidJoin` | java.util.Date | yes | Chengxintong (Trust Pass) join date |  |
| `buyRate` | java.lang.Long | yes | Buyer level |  |
| `categoryId` | int | yes | Category ID of the main industry |  |
| `companyName` | java.lang.String | yes | Company name |  |
| `personAccount` | boolean | yes | Whether it is a personal account |  |
| `homepageUrl` | java.lang.String | yes | Company homepage |  |
| `saleKeywords` | java.lang.String | yes | Sale keyword |  |
| `tpYear` | int | yes | Chengxintong membership years |  |
| `buyKeywords` | java.lang.String | yes | Buying request keyword |  |
| `memberBizType` | java.lang.String | yes | Business type<br> Enterprise ENTERPRISE(&quot;5&quot;),<br>Self-employed SELF_EMPLOYED(&quot;6&quot;),<br> Public institution or social organization ORGANIZATION(&quot;7&quot;),<br> Individual PERSONAL(&quot;8&quot;); |  |
| `rateSum` | java.lang.Long | yes | Seller's credit score (Taobao) |  |
| `domainInPlatforms` | String[] | yes | Shop address on 1688 |  |
| `memberId` | java.lang.String | yes | memberId |  |
| `shopUrl` | java.lang.String | yes | WangPu (shop) homepage URL |  |
| `supplierName` | java.lang.String | yes | Supplier name |  |
| `icon` | java.lang.String | yes | User avatar |  |
| `phoneNo` | java.lang.String | yes | Landline phone number |  |
| `industry` | java.lang.String | yes | Main industry |  |
| `product` | java.lang.String | yes | Main products |  |
| `department` | java.lang.String | yes | Department |  |
| `mobileNo` | java.lang.String | yes | Business contact mobile number (edited in member basic information)	<br>Business contact mobile number |  |
| `addressLocation` | java.lang.String | yes | Contact address |  |
| `email` | java.lang.String | yes | Business contact email |  |
| `sellerName` | java.lang.String | yes | Seller name |  |
| `kuaJingBao` | Boolean | yes | Whether it is Kuajingbao (Cross-Border Pay) | true |
| `crossBorder` | Boolean | yes | Whether it is Kuajingbao (Cross-Border Pay) | true |
| `isPm` | Boolean | yes | Whether it is a verified merchant | true |
| `pm` | Boolean | yes | Whether it is a verified merchant | true |
| `fm` | Boolean | yes | Whether it is a factory-finding merchant | true |
| `isOfficalLogistics` | Boolean | yes | Whether official logistics is supported | true |
| `status` | String | yes | Sub-account status | enabled |
| `openUid` | String | yes | Open platform encrypted user ID | sE542234KSls4KSfsthjST |

## Samples

**Output parameter example**

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
