// Package gen builds a deterministic 1688-shaped catalogue out of nothing but an
// offer id. It reads no files, calls no clock and uses no math/rand, so the same
// binary always serves the same catalogue and a demo recorded today replays
// identically tomorrow.
package gen

// Word is one vocabulary item with its Chinese and translated spellings. The
// stub speaks both because the catalogue APIs return subject (Chinese) and
// subjectTrans (translated) side by side.
type Word struct{ ZH, EN string }

// Axis is one SKU attribute axis, e.g. colour or size. AttributeID is the
// documented attributeId of the axis; the values carry their own ids implicitly
// by position.
type Axis struct {
	AttributeID int64
	NameZH      string
	NameEN      string
	Values      []Word
}

// Archetype is one product family. Every offer belongs to exactly one, and the
// archetype decides its vocabulary, its price band, its weight band and which
// attribute axes its SKUs vary over.
type Archetype struct {
	Key              string
	NameZH           string
	NameEN           string
	TopCategoryID    int64
	SecondCategoryID int64
	LeafCategories   []leafCategory

	Prefix   []Word
	Material []Word
	Noun     []Word
	Selling  []Word
	Unit     Word

	// Price band in fen and weight band in grams, both inclusive.
	MinFen, MaxFen int64
	MinG, MaxG     int64

	Axes []Axis
}

type leafCategory struct {
	ID int64
	Word
}

// Seller is one supplier. Twelve of them, fixed, so an order always has a
// plausible counterparty and the multi-seller rejection in the preview API has
// something real to trip over.
type Seller struct {
	MemberID    string
	UserID      int64
	LoginID     string
	CompanyName string
	ShopName    string
	Province    string
	City        string
	District    string
	TPYear      int
	MedalLevel  string
	OpenID      string
}

// sellers is the fixed supplier table.
var sellers = []Seller{
	{"b2b-2248564064", 2248564064, "alitestforisv02", "杭州宜路发发电子商务有限公司", "宜路发发服饰专营店", "浙江省", "杭州市", "滨江区", 7, "5", "sE542234KSls4KSfsthjST"},
	{"b2b-1624747073", 1624747073, "alitestforisv04", "义乌市锦程针织品有限公司", "锦程针织旗舰店", "浙江省", "金华市", "义乌市", 3, "4", "TQ8812LKsjfa9KLdjfKQb"},
	{"b2b-2248544159", 2248544159, "alitestforisv06", "广州市白云区嘉禾皮具厂", "嘉禾皮具工厂店", "广东省", "广州市", "白云区", 11, "5", "PW1029Jsdkf83KDlsjfQ2"},
	{"b2b-1676547900", 1676547900, "alitestforisv08", "深圳市南山区数联电子有限公司", "数联数码配件专营", "广东省", "深圳市", "南山区", 5, "4", "ZX7741Ksldj39FKdlsjA1"},
	{"b2b-1623492085", 1623492085, "alitestforisv10", "汕头市澄海区玩具家居用品厂", "澄海家居优选", "广东省", "汕头市", "澄海区", 2, "3", "NM3390Ksdlf83JKlsdfQ9"},
	{"b2b-2211084557", 2211084557, "alitestforisv12", "泉州市晋江鞋服贸易有限公司", "晋江鞋服批发", "福建省", "泉州市", "晋江市", 9, "5", "BV5527Klsdjf93KLsdfP4"},
	{"b2b-1739284018", 1739284018, "alitestforisv14", "苏州市吴江区丝绸纺织厂", "吴江丝绸源头厂", "江苏省", "苏州市", "吴江区", 6, "4", "LK9983Ksldjf83KLsdjQ0"},
	{"b2b-2093847561", 2093847561, "alitestforisv16", "广州市花都区美妆日化有限公司", "花都美妆日化", "广东省", "广州市", "花都区", 4, "4", "GH2218Kslfj93KLsdjfR7"},
	{"b2b-1884752013", 1884752013, "alitestforisv18", "温州市瓯海区文具制造厂", "瓯海文具直营", "浙江省", "温州市", "瓯海区", 8, "5", "RT6674Ksldjf93KLsdjF3"},
	{"b2b-2130987441", 2130987441, "alitestforisv20", "东莞市长安镇精密五金厂", "长安精密五金", "广东省", "东莞市", "长安镇", 10, "5", "YU4432Ksldjf83KLsdjW8"},
	{"b2b-1955620877", 1955620877, "alitestforisv22", "宁波市海曙区家纺有限公司", "海曙家纺工厂", "浙江省", "宁波市", "海曙区", 1, "3", "IO8856Kslfj93KLsdjfE5"},
	{"b2b-2074318890", 2074318890, "alitestforisv24", "石家庄市裕华区箱包制造厂", "裕华箱包直销", "河北省", "石家庄市", "裕华区", 12, "5", "QA1163Kslfj93KLsdjfT6"},
}

// Sellers exposes the fixed supplier table, so the stub can answer
// account.basic and the control endpoints without re-deriving it.
func Sellers() []Seller { return append([]Seller(nil), sellers...) }

var (
	colourAxis = Axis{
		AttributeID: 3216596,
		NameZH:      "颜色",
		NameEN:      "Color",
		Values: []Word{
			{"黑色", "Black"}, {"白色", "White"}, {"米杏色", "Beige"}, {"藏青色", "Navy"},
			{"雾霾蓝", "Haze Blue"}, {"豆沙粉", "Dusty Pink"}, {"军绿色", "Army Green"}, {"焦糖棕", "Caramel"},
		},
	}
	clothingSizeAxis = Axis{
		AttributeID: 450,
		NameZH:      "尺码",
		NameEN:      "Size",
		Values:      []Word{{"S", "S"}, {"M", "M"}, {"L", "L"}, {"XL", "XL"}, {"2XL", "2XL"}},
	}
	capacityAxis = Axis{
		AttributeID: 1797,
		NameZH:      "容量",
		NameEN:      "Capacity",
		Values:      []Word{{"小号", "Small"}, {"中号", "Medium"}, {"大号", "Large"}},
	}
	specAxis = Axis{
		AttributeID: 219,
		NameZH:      "规格",
		NameEN:      "Specification",
		Values:      []Word{{"单个装", "Single"}, {"两个装", "Pack of 2"}, {"五个装", "Pack of 5"}, {"十个装", "Pack of 10"}},
	}
	plugAxis = Axis{
		AttributeID: 100154,
		NameZH:      "接口类型",
		NameEN:      "Connector",
		Values:      []Word{{"Type-C", "Type-C"}, {"Lightning", "Lightning"}, {"Micro USB", "Micro USB"}},
	}
	scentAxis = Axis{
		AttributeID: 7062,
		NameZH:      "香型",
		NameEN:      "Scent",
		Values:      []Word{{"白茶", "White Tea"}, {"鼠尾草", "Sage"}, {"雪松", "Cedar"}, {"无香", "Unscented"}},
	}
	nibAxis = Axis{
		AttributeID: 3350,
		NameZH:      "笔尖",
		NameEN:      "Nib",
		Values:      []Word{{"0.38mm", "0.38mm"}, {"0.5mm", "0.5mm"}, {"0.7mm", "0.7mm"}},
	}
)

// archetypes is the fixed set of six product families.
var archetypes = []Archetype{
	{
		Key: "apparel", NameZH: "女装", NameEN: "Women's Apparel",
		TopCategoryID: 1, SecondCategoryID: 10166,
		LeafCategories: []leafCategory{
			{1031910, Word{"连衣裙", "Dress"}},
			{1031911, Word{"衬衫", "Blouse"}},
			{1031912, Word{"针织衫", "Knitwear"}},
		},
		Prefix:   []Word{{"韩版", "Korean-style"}, {"法式", "French"}, {"复古", "Retro"}, {"新中式", "New Chinese"}, {"通勤", "Commuter"}, {"设计感", "Designer"}},
		Material: []Word{{"雪纺", "Chiffon"}, {"棉麻", "Cotton Linen"}, {"针织", "Knitted"}, {"醋酸", "Acetate"}, {"莫代尔", "Modal"}},
		Noun:     []Word{{"连衣裙", "Dress"}, {"衬衫", "Shirt"}, {"半身裙", "Skirt"}, {"针织开衫", "Cardigan"}, {"阔腿裤", "Wide-leg Trousers"}, {"西装外套", "Blazer"}},
		Selling:  []Word{{"一件代发", "Dropshipping supported"}, {"现货速发", "In stock, ships fast"}, {"厂家直销", "Factory direct"}},
		Unit:     Word{"件", "Piece"},
		MinFen:   1800, MaxFen: 26800, MinG: 180, MaxG: 900,
		Axes: []Axis{colourAxis, clothingSizeAxis},
	},
	{
		Key: "bags", NameZH: "箱包皮具", NameEN: "Bags & Leather Goods",
		TopCategoryID: 2, SecondCategoryID: 122152,
		LeafCategories: []leafCategory{
			{1032010, Word{"单肩包", "Shoulder Bag"}},
			{1032011, Word{"双肩包", "Backpack"}},
			{1032012, Word{"钱包", "Wallet"}},
		},
		Prefix:   []Word{{"大容量", "Large-capacity"}, {"轻奢", "Affordable-luxury"}, {"防泼水", "Water-repellent"}, {"多层", "Multi-layer"}, {"简约", "Minimalist"}, {"通勤", "Commuter"}},
		Material: []Word{{"牛皮", "Cowhide"}, {"帆布", "Canvas"}, {"尼龙", "Nylon"}, {"PU", "PU Leather"}, {"牛津布", "Oxford Cloth"}},
		Noun:     []Word{{"单肩包", "Shoulder Bag"}, {"斜挎包", "Crossbody Bag"}, {"双肩背包", "Backpack"}, {"手提包", "Handbag"}, {"钱包", "Wallet"}, {"行李箱", "Suitcase"}},
		Selling:  []Word{{"支持定制", "Customisation available"}, {"工厂现货", "Factory stock"}, {"支持一件代发", "Dropshipping supported"}},
		Unit:     Word{"个", "Piece"},
		MinFen:   3500, MaxFen: 49800, MinG: 320, MaxG: 3200,
		Axes: []Axis{colourAxis, capacityAxis},
	},
	{
		Key: "home", NameZH: "家居日用", NameEN: "Home & Living",
		TopCategoryID: 3, SecondCategoryID: 133409,
		LeafCategories: []leafCategory{
			{1033110, Word{"收纳盒", "Storage Box"}},
			{1033111, Word{"床上用品", "Bedding"}},
			{1033112, Word{"厨房用品", "Kitchenware"}},
		},
		Prefix:   []Word{{"加厚", "Thickened"}, {"日式", "Japanese-style"}, {"北欧", "Nordic"}, {"折叠", "Foldable"}, {"防尘", "Dustproof"}, {"多功能", "Multifunctional"}},
		Material: []Word{{"PP塑料", "PP Plastic"}, {"竹纤维", "Bamboo Fibre"}, {"304不锈钢", "304 Stainless Steel"}, {"全棉", "Pure Cotton"}, {"硅胶", "Silicone"}},
		Noun:     []Word{{"收纳盒", "Storage Box"}, {"抽纸盒", "Tissue Box"}, {"沥水篮", "Drain Basket"}, {"四件套", "Bedding Set"}, {"挂钩", "Wall Hook"}, {"保鲜盒", "Food Container"}},
		Selling:  []Word{{"源头工厂", "Source factory"}, {"量大从优", "Bulk discount"}, {"48小时发货", "Ships within 48h"}},
		Unit:     Word{"套", "Set"},
		MinFen:   600, MaxFen: 12800, MinG: 120, MaxG: 5400,
		Axes: []Axis{colourAxis, specAxis},
	},
	{
		Key: "electronics", NameZH: "数码配件", NameEN: "Electronics Accessories",
		TopCategoryID: 4, SecondCategoryID: 144001,
		LeafCategories: []leafCategory{
			{1034410, Word{"数据线", "Charging Cable"}},
			{1034411, Word{"手机壳", "Phone Case"}},
			{1034412, Word{"充电器", "Charger"}},
		},
		Prefix:   []Word{{"快充", "Fast-charging"}, {"编织", "Braided"}, {"磁吸", "Magnetic"}, {"便携", "Portable"}, {"无线", "Wireless"}, {"车载", "In-car"}},
		Material: []Word{{"铝合金", "Aluminium Alloy"}, {"TPU", "TPU"}, {"液态硅胶", "Liquid Silicone"}, {"尼龙编织", "Nylon Braid"}, {"ABS", "ABS"}},
		Noun:     []Word{{"数据线", "Charging Cable"}, {"手机壳", "Phone Case"}, {"充电头", "Charging Head"}, {"支架", "Phone Stand"}, {"移动电源", "Power Bank"}, {"蓝牙耳机", "Bluetooth Earbuds"}},
		Selling:  []Word{{"通过CE认证", "CE certified"}, {"一年质保", "One-year warranty"}, {"支持贴牌", "OEM available"}},
		Unit:     Word{"条", "Piece"},
		MinFen:   290, MaxFen: 9900, MinG: 30, MaxG: 480,
		Axes: []Axis{colourAxis, plugAxis},
	},
	{
		Key: "beauty", NameZH: "美妆个护", NameEN: "Beauty & Personal Care",
		TopCategoryID: 5, SecondCategoryID: 155123,
		LeafCategories: []leafCategory{
			{1035510, Word{"洗护套装", "Hair Care Set"}},
			{1035511, Word{"香氛", "Fragrance"}},
			{1035512, Word{"美妆工具", "Makeup Tools"}},
		},
		Prefix:   []Word{{"氨基酸", "Amino-acid"}, {"温和", "Gentle"}, {"控油", "Oil-control"}, {"保湿", "Moisturising"}, {"便携装", "Travel-size"}, {"沙龙级", "Salon-grade"}},
		Material: []Word{{"植萃", "Botanical"}, {"玻尿酸", "Hyaluronic Acid"}, {"神经酰胺", "Ceramide"}, {"茶树", "Tea Tree"}, {"角鲨烷", "Squalane"}},
		Noun:     []Word{{"洗发水", "Shampoo"}, {"护手霜", "Hand Cream"}, {"身体乳", "Body Lotion"}, {"香薰蜡烛", "Scented Candle"}, {"洁面巾", "Facial Towel"}, {"化妆刷", "Makeup Brush"}},
		Selling:  []Word{{"备案齐全", "Fully registered"}, {"可代加工", "OEM/ODM"}, {"支持一件代发", "Dropshipping supported"}},
		Unit:     Word{"瓶", "Bottle"},
		MinFen:   890, MaxFen: 15800, MinG: 90, MaxG: 1200,
		Axes: []Axis{scentAxis, capacityAxis},
	},
	{
		Key: "stationery", NameZH: "文具办公", NameEN: "Stationery & Office",
		TopCategoryID: 6, SecondCategoryID: 166188,
		LeafCategories: []leafCategory{
			{1036610, Word{"笔类", "Pens"}},
			{1036611, Word{"本册", "Notebooks"}},
			{1036612, Word{"办公收纳", "Desk Organisers"}},
		},
		Prefix:   []Word{{"速干", "Quick-dry"}, {"学生", "Student"}, {"ins风", "Ins-style"}, {"考试专用", "Exam-grade"}, {"大容量", "High-capacity"}, {"简约", "Minimalist"}},
		Material: []Word{{"树脂", "Resin"}, {"再生纸", "Recycled Paper"}, {"PVC", "PVC"}, {"金属", "Metal"}, {"牛皮纸", "Kraft Paper"}},
		Noun:     []Word{{"中性笔", "Gel Pen"}, {"笔记本", "Notebook"}, {"文件袋", "Document Pouch"}, {"便利贴", "Sticky Notes"}, {"笔袋", "Pencil Case"}, {"桌面收纳架", "Desk Organiser"}},
		Selling:  []Word{{"工厂直发", "Ships from factory"}, {"支持印LOGO", "Logo printing available"}, {"混批不限量", "Unlimited mixed batch"}},
		Unit:     Word{"支", "Piece"},
		MinFen:   120, MaxFen: 6800, MinG: 15, MaxG: 900,
		Axes: []Axis{colourAxis, nibAxis},
	},
}

// Archetypes exposes the fixed families, for the category APIs and for
// /_control/catalogue.
func Archetypes() []Archetype { return append([]Archetype(nil), archetypes...) }

// statuses are the documented offer statuses. Anything other than "published"
// means the offer cannot be bought, which is what makes 500_001 reachable.
var statuses = []string{"expired", "member expired", "auto expired", "member deleted", "auditing", "untread", "new"}

// carriers is the small courier table used when an order ships.
var carriers = []struct{ Code, Name string }{
	{"ZTO", "中通快递"},
	{"YTO", "圆通速递"},
	{"STO", "申通快递"},
	{"YUNDA", "韵达速递"},
	{"SF", "顺丰速运"},
	{"JD", "京东物流"},
}

// Carrier returns one courier from the fixed table, chosen deterministically.
func Carrier(n int64) (code, name string) {
	c := carriers[int(uint64(n)%uint64(len(carriers)))]
	return c.Code, c.Name
}
