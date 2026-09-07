# Query multilingual category by category name

Original name: 根据类目名称查询多语言类目  
API: `com.alibaba.fenxiao.crossborder:category.translation.getByKeyword:1` · Category: Categories  
Doc page: https://open.1688.com/api/apidocdetail.htm?id=com.alibaba.fenxiao.crossborder:category.translation.getByKeyword-1  
Request URL: `https://gw.open.1688.com/openapi/param2/1/com.alibaba.fenxiao.crossborder/category.translation.getByKeyword/{appKey}`  
Requires user authorization (access_token) · Requires signature

Multilingual category query. Returns the list of matching category details in the requested language for the given category name. Child-category data is not included.

## System parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `_aop_timestamp` | String | no | Request timestamp |  |
| `_aop_signature` | String | yes | Request signature |  |
| `access_token` | String | yes | User authorization token |  |

## Request parameters

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `outMemberId` | java.lang.String | no | User's unique ID within the organization. No more than 64 characters, consisting of digits and letters. | 2423523f13tr12412f |
| `language` | java.lang.String | yes | Language. See the enum in the FAQ. | ja |
| `cateName` | java.lang.String | yes | Category name. Supports fuzzy search. | 裙子 |

## Response

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `result` | [message:category.translation.getByKeyword.ResultModel](#m-category-translation-getbykeyword-resultmodel) | yes | Return result | 如下 |

<a id="m-category-translation-getbykeyword-resultmodel"></a>
#### category.translation.getByKeyword.ResultModel

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `success` | boolean | yes | Whether successful | true |
| `code` | java.lang.String | yes | Error code | S0000 |
| `message` | java.lang.String | yes | Error description | 成功 |
| `result` | [message:category.translation.getByKeyword.Category[]](#m-category-translation-getbykeyword-category[]) | yes | Actual result | 如下 |

<a id="m-category-translation-getbykeyword-category[]"></a>
#### category.translation.getByKeyword.Category[]

| Field | Type | Required | Description | Example |
|---|---|---|---|---|
| `categoryId` | java.lang.Long | yes | Category ID | 1031910 |
| `chineseName` | java.lang.String | yes | Category name in Chinese | 连衣裙 |
| `translatedName` | java.lang.String | yes | Translated category name | ワンピース |
| `language` | java.lang.String | yes | Language | ja |
| `leaf` | java.lang.Boolean | yes | Whether it is a leaf category | true |
| `level` | java.lang.String | yes | Category level | 2 |
| `parentCateId` | java.lang.Long | yes | Parent category ID | 10166 |

## Samples

**Input parameter example**

```
{
  1111111111111,
  "1",
  "ja",
  "靴"
}
```

**Return value example**

```
{
    "result":
    {
        "success": true,
        "code": "S0000",
        "message": "成功",
        "result":
        [
            {
                "categoryId": 10170,
                "chineseName": "女式棉鞋",
                "translatedName": "婦人綿靴",
                "language": "ja",
                "leaf": true,
                "level": "3",
                "parentCateId": 125372002
            },
            {
                "categoryId": 10179,
                "chineseName": "制鞋机械",
                "translatedName": "靴作り機械",
                "language": "ja",
                "leaf": false,
                "level": "2",
                "parentCateId": 65
            },
            {
                "categoryId": 1033463,
                "chineseName": "滑雪靴",
                "translatedName": "スキー靴",
                "language": "ja",
                "leaf": true,
                "level": "3",
                "parentCateId": 2013
            },
            {
                "categoryId": 1034057,
                "chineseName": "鞋定制加工",
                "translatedName": "靴カスタム加工",
                "language": "ja",
                "leaf": true,
                "level": "3",
                "parentCateId": 201301601
            },
            {
                "categoryId": 1034077,
                "chineseName": "二手制鞋设备",
                "translatedName": "中古の靴製造設備",
                "language": "ja",
                "leaf": true,
                "level": "2",
                "parentCateId": 2829
            },
            {
                "categoryId": 1034373,
                "chineseName": "成品鞋、鞋件代理加盟",
                "translatedName": "完成品の靴、靴の部品は代理で加盟します。",
                "language": "ja",
                "leaf": true,
                "level": "3",
                "parentCateId": 125748005
            },
            {
                "categoryId": 1034392,
                "chineseName": "鞋材、鞋件加工",
                "translatedName": "靴材、靴部品加工",
                "language": "ja",
                "leaf": true,
                "level": "2",
                "parentCateId": 2805
            },
            {
                "categoryId": 1034394,
                "chineseName": "修鞋设备",
                "translatedName": "靴修理設備",
                "language": "ja",
                "leaf": true,
                "level": "3",
                "parentCateId": 10179
            },
            {
                "categoryId": 1034395,
                "chineseName": "鞋机配件",
                "translatedName": "靴機械の部品",
                "language": "ja",
                "leaf": true,
                "level": "3",
                "parentCateId": 10179
            },
            {
                "categoryId": 1034477,
                "chineseName": "其他制鞋机械",
                "translatedName": "その他の靴製造機械",
                "language": "ja",
                "leaf": true,
                "level": "3",
                "parentCateId": 10179
            },
            {
                "categoryId": 1034486,
                "chineseName": "鞋成型机",
                "translatedName": "靴成形機",
                "language": "ja",
                "leaf": true,
                "level": "3",
                "parentCateId": 10179
            },
            {
                "categoryId": 1036051,
                "chineseName": "缝鞋机",
                "translatedName": "靴縫い機",
                "language": "ja",
                "leaf": true,
                "level": "3",
                "parentCateId": 10179
            },
            {
                "categoryId": 1036832,
                "chineseName": "鞋修饰机",
                "translatedName": "靴トリムマシン",
                "language": "ja",
                "leaf": true,
                "level": "3",
                "parentCateId": 10179
            },
            {
                "categoryId": 1038548,
                "chineseName": "男式商务皮鞋",
                "translatedName": "メンズビジネス革靴",
                "language": "ja",
                "leaf": true,
                "level": "3",
                "parentCateId": 126506004
            },
            {
                "categoryId": 1041748,
                "chineseName": "烘鞋器、干鞋器",
                "translatedName": "靴ドライヤー、靴ドライヤー",
                "language": "ja",
                "leaf": true,
                "level": "3",
                "parentCateId": 1047981
            },
            {
                "categoryId": 1043192,
                "chineseName": "家居棉鞋",
                "translatedName": "家庭用綿靴",
                "language": "ja",
                "leaf": true,
                "level": "3",
                "parentCateId": 125312003
            },
            {
                "categoryId": 1043193,
                "chineseName": "女式板鞋",
                "translatedName": "レディーススケート靴",
                "language": "ja",
                "leaf": true,
                "level": "3",
                "parentCateId": 125372002
            },
            {
                "categoryId": 1043349,
                "chineseName": "袜子",
                "translatedName": "靴下",
                "language": "ja",
                "leaf": false,
                "level": "2",
                "parentCateId": 312
            },
            {
                "categoryId": 1043351,
                "chineseName": "儿童袜",
                "translatedName": "キッズ靴下",
                "language": "ja",
                "leaf": true,
                "level": "2",
                "parentCateId": 311
            },
            {
                "categoryId": 1043956,
                "chineseName": "鞋套、鞋刷、鞋用品",
                "translatedName": "靴カバー、靴ブラシ、靴用品",
                "language": "ja",
                "leaf": true,
                "level": "3",
                "parentCateId": 1556
            },
            {
                "categoryId": 121878002,
                "chineseName": "营地鞋",
                "translatedName": "キャンプ靴",
                "language": "ja",
                "leaf": true,
                "level": "3",
                "parentCateId": 121900002
            },
            {
                "categoryId": 121884002,
                "chineseName": "棒球鞋",
                "translatedName": "野球靴",
                "language": "ja",
                "leaf": true,
                "level": "3",
                "parentCateId": 121876002
            },
            {
                "categoryId": 121886003,
                "chineseName": "高海拔登山靴、攀冰鞋",
                "translatedName": "標高の高い登山靴、スケート靴",
                "language": "ja",
                "leaf": true,
                "level": "3",
                "parentCateId": 121900002
            },
            {
                "categoryId": 121896003,
                "chineseName": "登山鞋",
                "translatedName": "登山靴",
                "language": "ja",
                "leaf": true,
                "level": "3",
                "parentCateId": 121900002
            },
            {
                "categoryId": 122208002,
                "chineseName": "钓鱼鞋",
                "translatedName": "釣り靴",
                "language": "ja",
                "leaf": true,
                "level": "3",
                "parentCateId": 1044658
            },
            {
                "categoryId": 122246002,
                "chineseName": "军靴、战术靴",
                "translatedName": "軍靴、戦術靴",
                "language": "ja",
                "leaf": true,
                "level": "3",
                "parentCateId": 121900002
            },
            {
                "categoryId": 122406006,
                "chineseName": "鞋靴罩、袋",
                "translatedName": "靴カバー、バッグ",
                "language": "ja",
                "leaf": true,
                "level": "3",
                "parentCateId": 1548
            },
            {
                "categoryId": 122430003,
                "chineseName": "其他袜子",
                "translatedName": "その他の靴下",
                "language": "ja",
                "leaf": true,
                "level": "3",
                "parentCateId": 1043349
            },
            {
                "categoryId": 122586001,
                "chineseName": "宠物鞋靴",
                "translatedName": "ペット靴",
                "language": "ja",
                "leaf": true,
                "level": "3",
                "parentCateId": 201305503
            },
            {
                "categoryId": 122988007,
                "chineseName": "鱼嘴袜、露趾袜",
                "translatedName": "魚の口の靴下、足の指の靴下",
                "language": "ja",
                "leaf": true,
                "level": "3",
                "parentCateId": 1043349
            },
            {
                "categoryId": 123616001,
                "chineseName": "中老年妈妈鞋",
                "translatedName": "中高年ママ靴",
                "language": "ja",
                "leaf": true,
                "level": "3",
                "parentCateId": 125372002
            },
            {
                "categoryId": 123648008,
                "chineseName": "女式布鞋",
                "translatedName": "レディース布靴",
                "language": "ja",
                "leaf": true,
                "level": "3",
                "parentCateId": 125372002
            },
            {
                "categoryId": 124158004,
                "chineseName": "绝缘鞋",
                "translatedName": "絶縁靴",
                "language": "ja",
                "leaf": true,
                "level": "3",
                "parentCateId": 124166001
            },
            {
                "categoryId": 124158005,
                "chineseName": "防化鞋",
                "translatedName": "防化靴",
                "language": "ja",
                "leaf": true,
                "level": "3",
                "parentCateId": 124166001
            },
            {
                "categoryId": 124158006,
                "chineseName": "防砸防刺穿鞋",
                "translatedName": "靴を刺すのを防ぐ",
                "language": "ja",
                "leaf": true,
                "level": "3",
                "parentCateId": 124166001
            },
            {
                "categoryId": 124272013,
                "chineseName": "DIY鞋架、鞋柜",
                "translatedName": "DIY靴棚、下駄箱",
                "language": "ja",
                "leaf": true,
                "level": "3",
                "parentCateId": 122512002
            },
            {
                "categoryId": 125070006,
                "chineseName": "一次性鞋用品",
                "translatedName": "使い捨て靴用品",
                "language": "ja",
                "leaf": true,
                "level": "3",
                "parentCateId": 1556
            },
            {
                "categoryId": 125296002,
                "chineseName": "童/青少年鞋",
                "translatedName": "子供/青少年靴",
                "language": "ja",
                "leaf": true,
                "level": "3",
                "parentCateId": 121876002
            },
            {
                "categoryId": 125378002,
                "chineseName": "童特色鞋",
                "translatedName": "子供特有の靴",
                "language": "ja",
                "leaf": true,
                "level": "3",
                "parentCateId": 125386001
            },
            {
                "categoryId": 125386001,
                "chineseName": "童鞋、婴儿鞋",
                "translatedName": "子供靴、ベビーシューズ",
                "language": "ja",
                "leaf": false,
                "level": "2",
                "parentCateId": 1038378
            },
            {
                "categoryId": 125470004,
                "chineseName": "军迷作训鞋",
                "translatedName": "軍事訓練靴",
                "language": "ja",
                "leaf": true,
                "level": "3",
                "parentCateId": 121876002
            },
            {
                "categoryId": 125730003,
                "chineseName": "鞋配件",
                "translatedName": "靴の部品",
                "language": "ja",
                "leaf": false,
                "level": "2",
                "parentCateId": 1038378
            },
            {
                "categoryId": 126390002,
                "chineseName": "鞋服包装",
                "translatedName": "靴服パッケージ",
                "language": "ja",
                "leaf": false,
                "level": "2",
                "parentCateId": 68
            },
            {
                "categoryId": 126442003,
                "chineseName": "男式板鞋",
                "translatedName": "男性用板靴",
                "language": "ja",
                "leaf": true,
                "level": "3",
                "parentCateId": 126506004
            },
            {
                "categoryId": 126476001,
                "chineseName": "鞋靴包装",
                "translatedName": "靴のパッケージ",
                "language": "ja",
                "leaf": true,
                "level": "3",
                "parentCateId": 126390002
            },
            {
                "categoryId": 126478001,
                "chineseName": "袜子包装",
                "translatedName": "靴下の包装",
                "language": "ja",
                "leaf": true,
                "level": "3",
                "parentCateId": 126390002
            },
            {
                "categoryId": 126480003,
                "chineseName": "男式布鞋",
                "translatedName": "メンズ布靴",
                "language": "ja",
                "leaf": true,
                "level": "3",
                "parentCateId": 126506004
            },
            {
                "categoryId": 126506005,
                "chineseName": "女式时装单鞋",
                "translatedName": "レディースファッション靴",
                "language": "ja",
                "leaf": true,
                "level": "3",
                "parentCateId": 125372002
            },
            {
                "categoryId": 127362009,
                "chineseName": "童皮鞋",
                "translatedName": "子供用革靴",
                "language": "ja",
                "leaf": true,
                "level": "3",
                "parentCateId": 125386001
            },
            {
                "categoryId": 127450007,
                "chineseName": "鞋撑/鞋楦",
                "translatedName": "靴サポート/靴型",
                "language": "ja",
                "leaf": true,
                "level": "3",
                "parentCateId": 125730003
            },
            {
                "categoryId": 127450008,
                "chineseName": "鞋带",
                "translatedName": "靴ひも",
                "language": "ja",
                "leaf": true,
                "level": "3",
                "parentCateId": 125730003
            },
            {
                "categoryId": 127462008,
                "chineseName": "鞋拔",
                "translatedName": "靴抜き",
                "language": "ja",
                "leaf": true,
                "level": "3",
                "parentCateId": 125730003
            },
            {
                "categoryId": 201273557,
                "chineseName": "洗鞋机",
                "translatedName": "靴洗い機",
                "language": "ja",
                "leaf": true,
                "level": "3",
                "parentCateId": 652
            },
            {
                "categoryId": 201301601,
                "chineseName": "鞋材",
                "translatedName": "靴材",
                "language": "ja",
                "leaf": false,
                "level": "2",
                "parentCateId": 1038378
            },
            {
                "categoryId": 201301618,
                "chineseName": "防雨鞋套",
                "translatedName": "雨よけ靴カバー",
                "language": "ja",
                "leaf": true,
                "level": "3",
                "parentCateId": 125730003
            },
            {
                "categoryId": 201301811,
                "chineseName": "鞋袜干爽剂",
                "translatedName": "靴下乾燥剤",
                "language": "ja",
                "leaf": true,
                "level": "3",
                "parentCateId": 1556
            },
            {
                "categoryId": 201304725,
                "chineseName": "鞋袜除臭剂",
                "translatedName": "靴ソックス消臭剤",
                "language": "ja",
                "leaf": true,
                "level": "3",
                "parentCateId": 122196011
            },
            {
                "categoryId": 201304806,
                "chineseName": "宠物袜子",
                "translatedName": "ペット用靴下",
                "language": "ja",
                "leaf": true,
                "level": "3",
                "parentCateId": 201305503
            },
            {
                "categoryId": 201306309,
                "chineseName": "亲子鞋",
                "translatedName": "親子靴",
                "language": "ja",
                "leaf": true,
                "level": "3",
                "parentCateId": 125386001
            },
            {
                "categoryId": 201308711,
                "chineseName": "鞋子收纳架/盒",
                "translatedName": "靴収納ラック/ボックス",
                "language": "ja",
                "leaf": true,
                "level": "3",
                "parentCateId": 1548
            },
            {
                "categoryId": 201330420,
                "chineseName": "道鞋",
                "translatedName": "道の靴",
                "language": "ja",
                "leaf": true,
                "level": "3",
                "parentCateId": 201330711
            },
            {
                "categoryId": 201334120,
                "chineseName": "道袜",
                "translatedName": "道の靴下",
                "language": "ja",
                "leaf": true,
                "level": "3",
                "parentCateId": 201330711
            },
            {
                "categoryId": 201340517,
                "chineseName": "僧鞋",
                "translatedName": "僧靴",
                "language": "ja",
                "leaf": true,
                "level": "3",
                "parentCateId": 201330711
            },
            {
                "categoryId": 201359513,
                "chineseName": "寿衣/鞋/帽",
                "translatedName": "寿衣/靴/帽子",
                "language": "ja",
                "leaf": true,
                "level": "3",
                "parentCateId": 1044217
            },
            {
                "categoryId": 201536601,
                "chineseName": "女式制服皮鞋",
                "translatedName": "レディース制服革靴",
                "language": "ja",
                "leaf": true,
                "level": "3",
                "parentCateId": 125372002
            },
            {
                "categoryId": 201548101,
                "chineseName": "鞋油",
                "translatedName": "靴油",
                "language": "ja",
                "leaf": true,
                "level": "3",
                "parentCateId": 122262003
            },
            {
                "categoryId": 201585403,
                "chineseName": "足球袜",
                "translatedName": "サッカー靴下",
                "language": "ja",
                "leaf": true,
                "level": "3",
                "parentCateId": 122226003
            },
            {
                "categoryId": 201649401,
                "chineseName": "家用鞋套机",
                "translatedName": "家庭用靴カバー機",
                "language": "ja",
                "leaf": true,
                "level": "3",
                "parentCateId": 1556
            },
            {
                "categoryId": 201682809,
                "chineseName": "童鞋量脚器",
                "translatedName": "子供靴の足測定器",
                "language": "ja",
                "leaf": true,
                "level": "3",
                "parentCateId": 125386001
            },
            {
                "categoryId": 201809301,
                "chineseName": "洛丽塔袜",
                "translatedName": "ロリータ靴下",
                "language": "ja",
                "leaf": true,
                "level": "3",
                "parentCateId": 1043349
            },
            {
                "categoryId": 201827725,
                "chineseName": "家居婚鞋",
                "translatedName": "ホーム結婚靴",
                "language": "ja",
                "leaf": true,
                "level": "3",
                "parentCateId": 125312003
            },
            {
                "categoryId": 201834503,
                "chineseName": "学生皮鞋",
                "translatedName": "学生用革靴",
                "language": "ja",
                "leaf": true,
                "level": "3",
                "parentCateId": 125386001
            },
            {
                "categoryId": 201834603,
                "chineseName": "水晶/水钻鞋",
                "translatedName": "クリスタル/ウォータードリルの靴",
                "language": "ja",
                "leaf": true,
                "level": "3",
                "parentCateId": 125386001
            },
            {
                "categoryId": 201837318,
                "chineseName": "男式休闲皮鞋",
                "translatedName": "メンズカジュアル革靴",
                "language": "ja",
                "leaf": true,
                "level": "3",
                "parentCateId": 126506004
            },
            {
                "categoryId": 201837319,
                "chineseName": "乐福鞋/豆豆鞋",
                "translatedName": "楽福靴/豆靴",
                "language": "ja",
                "leaf": true,
                "level": "3",
                "parentCateId": 126506004
            },
            {
                "categoryId": 201843201,
                "chineseName": "女士洞洞鞋",
                "translatedName": "女性の穴の靴",
                "language": "ja",
                "leaf": true,
                "level": "3",
                "parentCateId": 125372002
            },
            {
                "categoryId": 201843202,
                "chineseName": "盆底鞋",
                "translatedName": "盆底の靴",
                "language": "ja",
                "leaf": true,
                "level": "3",
                "parentCateId": 125312003
            },
            {
                "categoryId": 201868102,
                "chineseName": "鞋塞",
                "translatedName": "靴の栓",
                "language": "ja",
                "leaf": true,
                "level": "3",
                "parentCateId": 125730003
            },
            {
                "categoryId": 201898903,
                "chineseName": "乐福鞋",
                "translatedName": "ローファー靴",
                "language": "ja",
                "leaf": true,
                "level": "3",
                "parentCateId": 125372002
            },
            {
                "categoryId": 201903001,
                "chineseName": "童鞋礼盒装",
                "translatedName": "子供靴ギフトボックス",
                "language": "ja",
                "leaf": true,
                "level": "3",
                "parentCateId": 125386001
            },
            {
                "categoryId": 201907604,
                "chineseName": "童洞洞鞋",
                "translatedName": "子供の穴の靴",
                "language": "ja",
                "leaf": true,
                "level": "3",
                "parentCateId": 125386001
            }
        ]
    }
}
```
