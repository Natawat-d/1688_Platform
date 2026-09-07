# Add buyer-seller distribution relationship

Original name: 买卖家分销关系添加  
API: `com.alibaba.fenxiao:alibaba.fenxiao.relationadd:1` · Category: Tools  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao:alibaba.fenxiao.relationadd-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao/alibaba.fenxiao.relationadd/{appKey}`  
Requires user authorization (access_token) · Requires signature

Add a buyer-seller distribution relationship by product ID.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `offerId` | java.lang.Long | yes | Product ID | 98129931 |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:alibaba.ocean.openplatform.common.addrelation.ResultModel](#m-alibaba-ocean-openplatform-common-addrelation-resultmodel) | yes |  |  |

<a id="m-alibaba-ocean-openplatform-common-addrelation-resultmodel"></a>
#### alibaba.ocean.openplatform.common.addrelation.ResultModel

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `success` | boolean | yes |  |  |
| `code` | java.lang.String | yes |  |  |
| `message` | java.lang.String | yes |  |  |
| `result` | java.lang.Boolean | yes |  |  |
