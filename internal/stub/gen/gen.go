package gen

import (
	"encoding/binary"
	"strconv"
	"strings"
	"time"

	"marketplace/internal/ali"
)

// CatalogueBase is the first offer id of the generated catalogue. Real 1688
// offer ids are twelve digits and up; these are deliberately in a band nothing
// real occupies.
const CatalogueBase int64 = 900000000000

// epoch anchors every generated date. time.Now() is never called here: an offer
// generated today and the same offer generated next month are byte-identical.
var epoch = time.Date(2024, 1, 1, 8, 0, 0, 0, time.FixedZone("CST", 8*60*60))

// Catalogue returns n consecutive offer ids starting at CatalogueBase.
func Catalogue(n int) []int64 { return CatalogueFrom(CatalogueBase, n) }

// CatalogueFrom returns n consecutive offer ids starting at base. /_control/reset
// uses it to hand out a different catalogue without changing the generator.
func CatalogueFrom(base int64, n int) []int64 {
	if n < 0 {
		n = 0
	}
	ids := make([]int64, n)
	for i := range ids {
		ids[i] = base + int64(i)
	}
	return ids
}

// SkuAttr is one sales attribute of one SKU.
type SkuAttr struct {
	AttributeID int64
	NameZH      string
	NameEN      string
	ValueZH     string
	ValueEN     string
	ImagePath   string
}

// SKU is one saleable specification of an offer.
type SKU struct {
	SkuID          int64
	SpecID         string
	Stock          int64
	Price          ali.Fen
	PromotionPrice ali.Fen
	CargoNumber    string
	WeightG        int64
	Attrs          []SkuAttr
}

// Tier is one wholesale price break.
type Tier struct {
	StartQuantity  int
	Price          ali.Fen
	PromotionPrice ali.Fen
}

// Attribute is one CPV product attribute (not a sales attribute).
type Attribute struct {
	AttributeID int64
	NameZH      string
	NameEN      string
	ValueZH     string
	ValueEN     string
}

// Product is everything the stub knows about one product. It is a pure function
// of the offer id: see Offer. (Go will not let a type and a function share a
// name, and the function is the entry point every caller uses, so the type is
// the one that got renamed.)
type Product struct {
	ID int64

	Archetype        string
	CategoryID       int64
	TopCategoryID    int64
	SecondCategoryID int64
	ThirdCategoryID  int64
	CategoryZH       string
	CategoryEN       string

	Subject      string
	SubjectTrans string
	Description  string
	SellingZH    []string
	SellingEN    []string

	// Status is the documented offer status. Every 37th offer is not
	// "published", which is what makes the "cannot be ordered" path reachable.
	Status string

	Seller Seller

	MOQ int
	// BatchNumber is the mixed-batch lot size: 0 (no lot) or MOQ.
	BatchNumber int
	// QuoteType cycles 0/1/2 across the catalogue: 0 no SKU quote by quantity,
	// 1 quote per SKU spec, 2 has SKUs but quotes by product quantity.
	QuoteType int

	UnitZH string
	UnitEN string

	SKUs  []SKU
	Tiers []Tier

	Attributes []Attribute

	// ImagePaths are stub-relative, e.g. "/img/900000000000/0.svg". The server
	// prefixes STUB_BASE; the generator has no idea what that is.
	ImagePaths []string

	Stock   int64
	WeightG int64
	LengthMM,
	WidthMM,
	HeightMM int64

	GeneralHunpi bool
	MixAmount    ali.Fen
	MixNumber    int

	SoldOut        int64
	MonthSold      int
	Sales7d        int
	TradeScore     string
	RepurchaseRate string
	IsJxhy         bool
	IsSelect       bool
	IsOnePsale     bool
	HasPromotion   bool

	CreateDate time.Time
	ModifyDate time.Time
}

// MinPrice is the lowest SKU price, which is what the search card shows.
func (o Product) MinPrice() ali.Fen {
	if len(o.SKUs) == 0 {
		return 0
	}
	min := o.SKUs[0].Price
	for _, s := range o.SKUs[1:] {
		if s.Price < min {
			min = s.Price
		}
	}
	return min
}

// SKUBySpec finds a SKU by its specId.
func (o Product) SKUBySpec(specID string) (SKU, bool) {
	for _, s := range o.SKUs {
		if s.SpecID == specID {
			return s, true
		}
	}
	return SKU{}, false
}

// TierPrice is the unit price for a given quantity: the last tier whose
// startQuantity has been reached.
func (o Product) TierPrice(qty int64) ali.Fen {
	price := o.MinPrice()
	for _, t := range o.Tiers {
		if qty >= int64(t.StartQuantity) {
			price = t.Price
		}
	}
	return price
}

// Orderable reports whether the offer may be bought at all. A non-published
// offer maps onto the documented 500_001 refusal.
func (o Product) Orderable() bool { return o.Status == "published" }

// Offer is total and pure: every int64 yields a well-formed product, and the
// same id always yields the same one. All variety comes from an FNV-1a hash of
// the id stretched by splitmix64.
func Offer(id int64) Product {
	r := newRNG(id)

	arch := archetypes[int(uint64(id)%uint64(len(archetypes)))]
	leaf := arch.LeafCategories[r.intn(len(arch.LeafCategories))]
	seller := sellers[int(uint64(hash64(id+7))%uint64(len(sellers)))]

	prefix := arch.Prefix[r.intn(len(arch.Prefix))]
	material := arch.Material[r.intn(len(arch.Material))]
	noun := arch.Noun[r.intn(len(arch.Noun))]

	o := Product{
		ID:               id,
		Archetype:        arch.Key,
		CategoryID:       leaf.ID,
		TopCategoryID:    arch.TopCategoryID,
		SecondCategoryID: arch.SecondCategoryID,
		ThirdCategoryID:  leaf.ID,
		CategoryZH:       leaf.ZH,
		CategoryEN:       leaf.EN,
		Subject:          prefix.ZH + material.ZH + noun.ZH + "厂家直销批发",
		SubjectTrans:     prefix.EN + " " + material.EN + " " + noun.EN + " Factory Direct Wholesale",
		Seller:           seller,
		UnitZH:           arch.Unit.ZH,
		UnitEN:           arch.Unit.EN,
		QuoteType:        int(uint64(id) % 3),
		Status:           "published",
		TradeScore:       []string{"4.2", "4.5", "4.7", "4.8", "5.0"}[r.intn(5)],
		RepurchaseRate:   strconv.Itoa(5+r.intn(45)) + "%",
		IsJxhy:           id%4 == 0,
		IsSelect:         id%6 == 0,
		IsOnePsale:       id%3 != 2,
		HasPromotion:     id%9 == 0,
		MonthSold:        r.intn(4000),
		Sales7d:          r.intn(400),
		CreateDate:       epoch.AddDate(0, 0, r.intn(600)),
	}
	o.ModifyDate = o.CreateDate.AddDate(0, 0, 1+r.intn(120))
	o.Description = o.Subject + "。" + arch.NameZH + "源头厂家，" + seller.City + seller.District + "发货，支持混批与一件代发。"

	for _, w := range arch.Selling {
		o.SellingZH = append(o.SellingZH, w.ZH)
		o.SellingEN = append(o.SellingEN, w.EN)
	}

	// Every 37th offer is not on sale.
	if id%37 == 0 {
		o.Status = statuses[int(uint64(hash64(id))%uint64(len(statuses)))]
	}

	o.MOQ = []int{1, 2, 5, 10, 20}[r.intn(5)]
	if id%2 == 0 {
		o.BatchNumber = o.MOQ
	}
	o.GeneralHunpi = id%5 != 0
	o.MixNumber = 1 + r.intn(4)
	o.MixAmount = ali.Fen(int64(1+r.intn(20)) * 1000)

	o.WeightG = arch.MinG + int64(r.intn(int(arch.MaxG-arch.MinG+1)))
	o.LengthMM = 40 + int64(r.intn(460))
	o.WidthMM = 30 + int64(r.intn(320))
	o.HeightMM = 15 + int64(r.intn(200))

	base := ali.Fen(arch.MinFen + int64(r.intn(int(arch.MaxFen-arch.MinFen+1))))
	base = base - base%10 // keep prices to whole jiao, as real listings tend to be

	o.SKUs = buildSKUs(id, arch, base, &r)
	o.Tiers = buildTiers(base, &r)

	for _, s := range o.SKUs {
		o.Stock += s.Stock
	}
	o.SoldOut = int64(o.MonthSold)*3 + int64(r.intn(900))

	// Product-level CPV attributes: the archetype's own axes flattened, plus
	// two fixed ones that every real listing carries.
	o.Attributes = []Attribute{
		{AttributeID: 100, NameZH: "货号", NameEN: "Item No.", ValueZH: cargoNumber(id, 0), ValueEN: cargoNumber(id, 0)},
		{AttributeID: 101, NameZH: "产地", NameEN: "Place of Origin", ValueZH: seller.Province + seller.City, ValueEN: seller.City},
		{AttributeID: 102, NameZH: "材质", NameEN: "Material", ValueZH: material.ZH, ValueEN: material.EN},
		{AttributeID: 103, NameZH: "风格", NameEN: "Style", ValueZH: prefix.ZH, ValueEN: prefix.EN},
	}

	imgs := 3 + r.intn(3)
	for i := 0; i < imgs; i++ {
		o.ImagePaths = append(o.ImagePaths, ImagePath(id, i))
	}
	return o
}

// buildSKUs makes 1-12 SKUs over 0, 1 or 2 attribute axes. Every 53rd SKU is out
// of stock, and every 101st offer carries a zero-priced first SKU so the
// documented 500_008 refusal is reachable.
func buildSKUs(id int64, arch Archetype, base ali.Fen, r *rng) []SKU {
	axisCount := r.intn(3) // 0, 1 or 2 axes
	if axisCount > len(arch.Axes) {
		axisCount = len(arch.Axes)
	}

	type combo []SkuAttr
	combos := []combo{nil}
	for a := 0; a < axisCount; a++ {
		axis := arch.Axes[a]
		// Take a deterministic slice of the axis values, at least two.
		n := 2 + r.intn(len(axis.Values)-1)
		if n > len(axis.Values) {
			n = len(axis.Values)
		}
		var next []combo
		for _, c := range combos {
			for v := 0; v < n; v++ {
				val := axis.Values[v]
				attr := SkuAttr{
					AttributeID: axis.AttributeID,
					NameZH:      axis.NameZH,
					NameEN:      axis.NameEN,
					ValueZH:     val.ZH,
					ValueEN:     val.EN,
				}
				if a == 0 {
					attr.ImagePath = ImagePath(id, v)
				}
				next = append(next, append(append(combo(nil), c...), attr))
			}
		}
		combos = next
	}
	if len(combos) > 12 {
		combos = combos[:12]
	}

	skus := make([]SKU, 0, len(combos))
	for i, c := range combos {
		price := base + ali.Fen(int64(i)*int64(10+r.intn(90)))
		s := SKU{
			SkuID:       skuID(id, i),
			SpecID:      specID(id, i),
			Stock:       int64(20 + r.intn(9800)),
			Price:       price,
			CargoNumber: cargoNumber(id, i),
			WeightG:     arch.MinG + int64(r.intn(int(arch.MaxG-arch.MinG+1))),
			Attrs:       c,
		}
		if id%9 == 0 {
			s.PromotionPrice = price - price/10
		}
		// Every 53rd SKU across the catalogue has nothing left.
		if (id+int64(i))%53 == 0 {
			s.Stock = 0
		}
		skus = append(skus, s)
	}
	if id%101 == 0 && len(skus) > 0 {
		skus[0].Price = 0
		skus[0].PromotionPrice = 0
	}
	return skus
}

// buildTiers makes 2-4 wholesale price breaks with descending prices.
func buildTiers(base ali.Fen, r *rng) []Tier {
	starts := []int{1, 10, 50, 100}
	n := 2 + r.intn(3) // 2, 3 or 4
	tiers := make([]Tier, 0, n)
	price := base + ali.Fen(int64(20+r.intn(200)))
	for i := 0; i < n; i++ {
		t := Tier{StartQuantity: starts[i], Price: price}
		if r.intn(4) == 0 {
			t.PromotionPrice = price - price/20
		}
		tiers = append(tiers, t)
		drop := price / ali.Fen(8+r.intn(8))
		if drop < 10 {
			drop = 10
		}
		if price-drop < 100 {
			drop = 0
		}
		price -= drop
	}
	return tiers
}

// skuID is a globally unique-looking sku id, 13 digits like the real ones.
func skuID(offerID int64, i int) int64 {
	h := hash64(offerID*131 + int64(i))
	return 3000000000000 + int64(h%1000000000000)
}

// specID is the 32-hex-character spec identifier the trade APIs key on.
func specID(offerID int64, i int) string {
	a := hash64(offerID*7919 + int64(i)*31)
	b := hash64(int64(a) ^ int64(i+1))
	var sb strings.Builder
	sb.Grow(32)
	writeHex16(&sb, a)
	writeHex16(&sb, b)
	return sb.String()
}

func writeHex16(sb *strings.Builder, v uint64) {
	const digits = "0123456789abcdef"
	for shift := 60; shift >= 0; shift -= 4 {
		sb.WriteByte(digits[(v>>uint(shift))&0xf])
	}
}

func cargoNumber(offerID int64, i int) string {
	return "CN" + strconv.FormatInt(offerID%1000000, 10) + "-" + strconv.Itoa(i+1)
}

// ImagePath is the stub-relative URL of one offer image. The server turns it
// into an absolute URL with STUB_BASE, so the demo never touches the network.
func ImagePath(offerID int64, n int) string {
	return "/img/" + strconv.FormatInt(offerID, 10) + "/" + strconv.Itoa(n) + ".svg"
}

// hash64 is FNV-1a over the little-endian bytes of v.
func hash64(v int64) uint64 {
	const offset64 = 14695981039346656037
	const prime64 = 1099511628211
	var b [8]byte
	binary.LittleEndian.PutUint64(b[:], uint64(v))
	h := uint64(offset64)
	for _, c := range b {
		h ^= uint64(c)
		h *= prime64
	}
	return h
}

// rng is splitmix64 seeded from hash64. Deterministic, allocation-free and, most
// importantly, not math/rand: the global source there could be seeded by anyone.
type rng struct{ s uint64 }

func newRNG(id int64) rng { return rng{s: hash64(id)} }

func (r *rng) next() uint64 {
	r.s += 0x9E3779B97F4A7C15
	z := r.s
	z = (z ^ (z >> 30)) * 0xBF58476D1CE4E5B9
	z = (z ^ (z >> 27)) * 0x94D049BB133111EB
	return z ^ (z >> 31)
}

func (r *rng) intn(n int) int {
	if n <= 0 {
		return 0
	}
	return int(r.next() % uint64(n))
}
