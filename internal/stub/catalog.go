package stub

import (
	"net/url"
	"strconv"
	"strings"

	"marketplace/internal/ali"
	"marketplace/internal/stub/gen"
)

// The catalogue endpoints. Every field name below is a literal copied out of
// 1688-api-docs/en/, including the two spellings of the same idea that the two
// endpoints use: the search card says promotionURL, the detail says promotionUrl.

// handleAccountBasic serves com.alibaba.account:alibaba.account.basic:1.
// Family B, and note isOfficalLogistics — 1688's spelling, not a typo of ours.
func (s *Server) handleAccountBasic(_ url.Values) any {
	return famB(map[string]any{
		"loginId":            buyer.LoginID,
		"memberId":           buyer.MemberID,
		"userId":             buyer.UserID,
		"maturity":           0,
		"memo":               "",
		"modifyDate":         ali.FormatCompact(s.clock.Now()),
		"createDate":         "20130312160828000+0800",
		"categoryName":       "女装",
		"categoryId":         10166,
		"trustScore":         0,
		"enterpriseAccount":  true,
		"personAccount":      false,
		"communityLevel":     "A2",
		"joinFrom":           "CJ_COMMON_JOIN_M",
		"rateNum":            0,
		"rateSum":            0,
		"gmtPaidJoin":        "20170502000942000+0800",
		"buyRate":            3,
		"saleRate":           0,
		"companyName":        buyer.Company,
		"supplierName":       buyer.Company,
		"sellerName":         buyer.Company,
		"homepageUrl":        s.cfg.Base + "/shop/" + buyer.LoginID,
		"shopUrl":            s.cfg.Base + "/shop/" + buyer.LoginID,
		"saleKeywords":       "女装,男装,女鞋,衬衫",
		"buyKeywords":        "男装,女装,女鞋",
		"tpYear":             0,
		"memberBizType":      "ENTERPRISE",
		"domainInPlatforms":  []any{s.cfg.Base + "/shop/" + buyer.LoginID},
		"icon":               "/cms/upload/2011/116/401/104611_1301427272.png",
		"phoneNo":            "86-0571-81895955",
		"mobileNo":           "15251667788",
		"industry":           "女装",
		"product":            "连衣裙,衬衫",
		"department":         "跨境采购部",
		"addressLocation":    "浙江省杭州市滨江区网商路699号",
		"email":              "buyer@example.test",
		"kuaJingBao":         true,
		"crossBorder":        true,
		"isPm":               false,
		"pm":                 false,
		"fm":                 false,
		"isOfficalLogistics": true,
		"status":             "enabled",
		"openUid":            "sE542234KSls4KSfsthjST",
	})
}

// handleKeywordQuery serves product.search.keywordQuery. Family A, with the
// search-card shape — which is NOT the detail shape.
func (s *Server) handleKeywordQuery(form url.Values) any {
	p, ok := jsonObject(form, "offerQueryParam")
	if !ok {
		return famAErr("400", "offerQueryParam is not a JSON object")
	}
	keyword := strings.TrimSpace(mapStr(p, "keyword"))
	country := mapStr(p, "country")
	if country == "" {
		return famAErr("400", "country is required")
	}
	page := int(mapInt64(p, "beginPage"))
	if page < 1 {
		page = 1
	}
	size := int(mapInt64(p, "pageSize"))
	if size < 1 {
		size = 20
	}
	if size > 50 {
		size = 50
	}
	categoryID := mapInt64(p, "categoryId")
	priceStart, _ := ali.ParseFen(mapStr(p, "priceStart"))
	priceEnd, _ := ali.ParseFen(mapStr(p, "priceEnd"))

	var matched []gen.Product
	for _, id := range s.catalogue() {
		o := s.offer(id)
		if !matchesKeyword(o, keyword) {
			continue
		}
		if categoryID != 0 && o.CategoryID != categoryID && o.SecondCategoryID != categoryID && o.TopCategoryID != categoryID {
			continue
		}
		if priceStart > 0 && o.MinPrice() < priceStart {
			continue
		}
		if priceEnd > 0 && o.MinPrice() > priceEnd {
			continue
		}
		matched = append(matched, o)
	}

	total := len(matched)
	totalPage := (total + size - 1) / size
	from := (page - 1) * size
	if from > total {
		from = total
	}
	to := from + size
	if to > total {
		to = total
	}

	data := make([]any, 0, to-from)
	for _, o := range matched[from:to] {
		data = append(data, s.searchCard(o))
	}

	return famA(map[string]any{
		"totalRecords": total,
		"totalPage":    totalPage,
		"pageSize":     size,
		"currentPage":  page,
		"data":         data,
	})
}

func matchesKeyword(o gen.Product, keyword string) bool {
	if keyword == "" {
		return true
	}
	k := strings.ToLower(keyword)
	for _, hay := range []string{o.Subject, o.SubjectTrans, o.CategoryZH, o.CategoryEN} {
		if strings.Contains(strings.ToLower(hay), k) {
			return true
		}
	}
	return false
}

// searchCard builds one product.search.keywordQuery.model.ProductInfoModelV2.
func (s *Server) searchCard(o gen.Product) map[string]any {
	min := o.MinPrice()
	card := map[string]any{
		"imageUrl":     s.imageURL(o.ImagePaths[0]),
		"aigcImageUrl": s.imageURL(o.ImagePaths[0]),
		"subject":      o.Subject,
		"subjectTrans": o.SubjectTrans,
		"offerId":      o.ID,
		"isJxhy":       o.IsJxhy,
		"priceInfo": map[string]any{
			"price":          price(min),
			"jxhyPrice":      price(min - min/20),
			"pfJxhyPrice":    price(min - min/25),
			"consignPrice":   price(min + min/10),
			"promotionPrice": price(min - min/10),
		},
		"repurchaseRate":   o.RepurchaseRate,
		"monthSold":        o.MonthSold,
		"traceInfo":        "object_id@" + strconv.FormatInt(o.ID, 10) + "^object_type@offer",
		"isOnePsale":       o.IsOnePsale,
		"sellerIdentities": s.chaosList(o.ID, sellerIdentities(o)),
		"offerIdentities":  s.chaosList(o.ID, offerIdentities(o)),
		"tradeScore":       o.TradeScore,
		"whiteImage":       s.imageURL(o.ImagePaths[0]),
		"promotionModel":   map[string]any{"hasPromotion": o.HasPromotion, "promotionType": promotionType(o)},
		"topCategoryId":    o.TopCategoryID,
		"secondCategoryId": o.SecondCategoryID,
		"thirdCategoryId":  o.ThirdCategoryID,
		"isPatentProduct":  o.ID%13 == 0,
		"createDate":       ali.FormatNaive(o.CreateDate),
		"modifyDate":       ali.FormatNaive(o.ModifyDate),
		"isSelect":         o.IsSelect,
		"minOrderQuantity": o.MOQ,
		"sellerDataInfo":   s.sellerDataInfoSearch(o),
		"token":            "tk" + strconv.FormatInt(o.ID%100000000, 10),
		"promotionURL":     s.cfg.Base + "/offer/" + strconv.FormatInt(o.ID, 10) + ".html?cps=1",
		"sales7d":          strconv.Itoa(o.Sales7d),
		"productTradeInfo": productTradeInfo(o),
		"invoiceInfo":      invoiceInfo(o),
		"productSimpleShippingInfo": map[string]any{
			"shippingTimeGuarantee":     shippingGuarantee(o),
			"perfectFulfillmentRate30d": "99.10%",
			"pickupWithin24hRate30d":    "98.20%",
			"qualityReturnRate30d":      "0.12%",
			"perfectFulfillmentRate7d":  "99.80%",
			"pickupWithin24hRate7d":     "97.60%",
			"delayedShippingRate7d":     "0.40%",
		},
	}
	return s.chaosExtra(o.ID, card)
}

func promotionType(o gen.Product) string {
	if o.HasPromotion {
		return "plus"
	}
	return ""
}

func shippingGuarantee(o gen.Product) string {
	if o.ID%2 == 0 {
		return "shipIn24Hours"
	}
	return "shipIn48Hours"
}

func sellerIdentities(o gen.Product) []any {
	var out []any
	if o.Seller.TPYear > 0 {
		out = append(out, "tp_member-诚信通会员")
	}
	if o.Seller.MedalLevel == "5" {
		out = append(out, "powerful_merchants-实力商家")
	}
	if o.ID%7 == 0 {
		out = append(out, "super_factory-超级工厂")
	}
	return out
}

func offerIdentities(o gen.Product) []any {
	var out []any
	if o.IsSelect {
		out = append(out, "select-跨境select")
	}
	if o.IsJxhy {
		out = append(out, "yx-严选")
	}
	return out
}

func productTradeInfo(o gen.Product) map[string]any {
	return map[string]any{
		"addCartCount7d":   strconv.Itoa(o.Sales7d * 3),
		"payBuyerCount7d":  strconv.Itoa(o.Sales7d / 2),
		"addCartCount30d":  strconv.Itoa(o.MonthSold * 3),
		"payBuyerCount30d": strconv.Itoa(o.MonthSold / 2),
	}
}

func invoiceInfo(o gen.Product) map[string]any {
	types := []any{"普票"}
	if o.ID%3 == 0 {
		types = append(types, "专票")
	}
	return map[string]any{
		"supportOnlineInvoice": o.ID%2 == 0,
		"supportFastInvoice":   o.ID%5 == 0,
		"invoiceTypes":         types,
		"taxpayerType":         "一般纳税人",
	}
}

func (s *Server) sellerDataInfoSearch(o gen.Product) map[string]any {
	return map[string]any{
		"tradeMedalLevel":           o.Seller.MedalLevel,
		"compositeServiceScore":     "4.6",
		"logisticsExperienceScore":  "4.5",
		"disputeComplaintScore":     "4.2",
		"offerExperienceScore":      "4.7",
		"afterSalesExperienceScore": "4.3",
		"consultingExperienceScore": "4.8",
		"repeatPurchasePercent":     strings.TrimSuffix(o.RepurchaseRate, "%"),
		"tpYear":                    o.Seller.TPYear,
	}
}

// handleKeywordSNQuery serves product.search.keywordSNQuery. Family A with the
// retCode / retMsg spelling.
func (s *Server) handleKeywordSNQuery(form url.Values) any {
	p, ok := jsonObject(form, "snParams")
	if !ok {
		return famAErr("400", "snParams is not a JSON object")
	}
	keyword := mapStr(p, "keyword")
	if strings.TrimSpace(keyword) == "" {
		return famAErr("400", "keyword is required")
	}

	nav := []any{
		snGroup("973", "风格", "Style", [][3]string{
			{"28105", "韩版", "Korean version"},
			{"28106", "法式", "French"},
			{"28107", "复古", "Retro"},
			{"28108", "简约", "Minimalist"},
		}),
		snGroup("978", "材质", "Material", [][3]string{
			{"1352", "棉麻", "Cotton linen"},
			{"1353", "雪纺", "Chiffon"},
			{"1354", "针织", "Knitted"},
		}),
		snGroup("981", "适用季节", "Season", [][3]string{
			{"6601", "春秋", "Spring & Autumn"},
			{"6602", "夏季", "Summer"},
			{"6603", "冬季", "Winter"},
		}),
	}
	return famASN(nav)
}

func snGroup(id, name, trans string, children [][3]string) map[string]any {
	kids := make([]any, 0, len(children))
	for _, c := range children {
		kids = append(kids, map[string]any{
			"id":            id + ":" + c[0],
			"name":          c[1],
			"translateName": c[2],
		})
	}
	return map[string]any{
		"id":            id,
		"name":          name,
		"translateName": trans,
		"children":      kids,
	}
}

// handleProductDetail serves product.search.queryProductDetail. Family A, detail
// shape: richer than the search card and differently spelled in places.
func (s *Server) handleProductDetail(form url.Values) any {
	p, ok := jsonObject(form, "offerDetailParam")
	if !ok {
		return famAErr("400", "offerDetailParam is not a JSON object")
	}
	offerID := mapInt64(p, "offerId")
	if offerID <= 0 {
		return famAErr("400", "offerId is required")
	}
	if mapStr(p, "country") == "" {
		return famAErr("400", "country is required")
	}
	o := s.offer(offerID)

	skus := make([]any, 0, len(o.SKUs))
	for _, sk := range o.SKUs {
		attrs := make([]any, 0, len(sk.Attrs))
		for _, a := range sk.Attrs {
			img := ""
			if a.ImagePath != "" {
				img = s.imageURL(a.ImagePath)
			}
			attrs = append(attrs, map[string]any{
				"attributeId":        a.AttributeID,
				"attributeName":      a.NameZH,
				"attributeNameTrans": a.NameEN,
				"value":              a.ValueZH,
				"valueTrans":         a.ValueEN,
				"skuImageUrl":        img,
				"skuImageUrlTrans":   img,
			})
		}
		skus = append(skus, map[string]any{
			"amountOnSale":   sk.Stock,
			"price":          price(sk.Price),
			"jxhyPrice":      price(sk.Price - sk.Price/20),
			"skuId":          sk.SkuID,
			"specId":         sk.SpecID,
			"skuAttributes":  s.chaosList(o.ID, attrs),
			"pfJxhyPrice":    price(sk.Price - sk.Price/25),
			"consignPrice":   price(sk.Price + sk.Price/10),
			"cargoNumber":    sk.CargoNumber,
			"promotionPrice": price(sk.PromotionPrice),
			"fenxiaoPriceInfo": map[string]any{
				"onePiecePrice": price(sk.Price + sk.Price/5),
				"offerPrice":    price(sk.Price),
			},
			"retailPrice": price(sk.Price * 2),
		})
	}

	tiers := make([]any, 0, len(o.Tiers))
	for _, t := range o.Tiers {
		tiers = append(tiers, map[string]any{
			"startQuantity":  t.StartQuantity,
			"price":          price(t.Price),
			"promotionPrice": price(t.PromotionPrice),
		})
	}

	attrs := make([]any, 0, len(o.Attributes))
	for _, a := range o.Attributes {
		attrs = append(attrs, map[string]any{
			"attributeId":        strconv.FormatInt(a.AttributeID, 10),
			"attributeName":      a.NameZH,
			"value":              a.ValueZH,
			"attributeNameTrans": a.NameEN,
			"valueTrans":         a.ValueEN,
		})
	}

	skuShipping := make([]any, 0, len(o.SKUs))
	skuShippingDetail := make([]any, 0, len(o.SKUs))
	for _, sk := range o.SKUs {
		skuShipping = append(skuShipping, map[string]any{
			"specId": sk.SpecID,
			"skuId":  sk.SkuID,
			"width":  mm(o.WidthMM),
			"length": mm(o.LengthMM),
			"height": mm(o.HeightMM),
			"weight": sk.WeightG, // grams here; kilograms in skuShippingDetails
		})
		skuShippingDetail = append(skuShippingDetail, map[string]any{
			"skuId":            strconv.FormatInt(sk.SkuID, 10),
			"width":            mm(o.WidthMM),
			"length":           mm(o.LengthMM),
			"height":           mm(o.HeightMM),
			"weight":           kg(sk.WeightG),
			"pkgSizeSource":    "商家自填",
			"officialLength":   mm(o.LengthMM),
			"officialWidth":    mm(o.WidthMM),
			"officialHeight":   mm(o.HeightMM),
			"officialWeight":   kg(sk.WeightG),
			"aiWeight":         kg(sk.WeightG),
			"aiWeightAccuracy": "80%",
		})
	}

	selling := make([]any, 0, len(o.SellingEN))
	for i, en := range o.SellingEN {
		selling = append(selling, o.SellingZH[i]+":"+en)
	}

	detail := map[string]any{
		"offerId":      o.ID,
		"categoryId":   o.CategoryID,
		"categoryName": o.CategoryZH,
		"subject":      o.Subject,
		"subjectTrans": o.SubjectTrans,
		"description":  o.Description,
		"mainVideo":    "",
		"detailVideo":  "",
		"productImage": map[string]any{
			"images":     toAny(s.imageURLs(o.ImagePaths)),
			"whiteImage": s.imageURL(o.ImagePaths[0]),
		},
		"productImageTrans": map[string]any{
			"images":     toAny(s.imageURLs(o.ImagePaths)),
			"whiteImage": s.imageURL(o.ImagePaths[0]),
		},
		"productAttribute": s.chaosList(o.ID, attrs),
		"productSkuInfos":  s.chaosList(o.ID, skus),
		"productSaleInfo": map[string]any{
			"amountOnSale":   o.Stock,
			"priceRangeList": s.chaosList(o.ID, tiers),
			"quoteType":      o.QuoteType,
			"consignPrice":   price(o.MinPrice() + o.MinPrice()/10),
			"jxhyPrice":      price(o.MinPrice() - o.MinPrice()/20),
			"unitInfo":       map[string]any{"unit": o.UnitZH, "transUnit": o.UnitEN},
			"fenxiaoSaleInfo": map[string]any{
				"onePieceFreePostage": o.IsOnePsale,
				"startQuantity":       o.MOQ,
				"onePiecePrice":       price(o.MinPrice() + o.MinPrice()/5),
				"offerPrice":          price(o.MinPrice()),
			},
			"retailPrice": price(o.MinPrice() * 2),
		},
		"productShippingInfo": map[string]any{
			"sendGoodsAddressText":  o.Seller.Province + o.Seller.City + o.Seller.District,
			"weight":                kg(o.WeightG),
			"width":                 mm(o.WidthMM),
			"height":                mm(o.HeightMM),
			"length":                mm(o.LengthMM),
			"skuShippingInfoList":   s.chaosList(o.ID, skuShipping),
			"shippingTimeGuarantee": shippingGuarantee(o),
			"skuShippingDetails":    s.chaosList(o.ID, skuShippingDetail),
			"pkgSizeSource":         "商家自填",
			"officialLength":        mm(o.LengthMM),
			"officialWidth":         mm(o.WidthMM),
			"officialHeight":        mm(o.HeightMM),
			"officialWeight":        kg(o.WeightG),
		},
		"isJxhy":           o.IsJxhy,
		"sellerOpenId":     o.Seller.OpenID,
		"minOrderQuantity": o.MOQ,
		"batchNumber":      o.BatchNumber,
		"status":           o.Status,
		"tagInfoList": []any{
			map[string]any{"key": "isOnePsale", "value": o.IsOnePsale},
			map[string]any{"key": "select", "value": o.IsSelect},
		},
		"traceInfo": "object_id@" + strconv.FormatInt(o.ID, 10) + "^object_type@offer",
		"sellerMixSetting": map[string]any{
			"generalHunpi": o.GeneralHunpi,
			"mixAmount":    int64(o.MixAmount),
			"mixNumber":    o.MixNumber,
		},
		"productCargoNumber": o.Attributes[0].ValueZH,
		"sellerDataInfo":     s.sellerDataInfoDetail(o),
		"soldOut":            strconv.FormatInt(o.SoldOut, 10),
		"channelPrice":       map[string]any{"channelSkuPriceList": s.chaosList(o.ID, channelPrices(o))},
		"promotionModel":     map[string]any{"hasPromotion": o.HasPromotion, "promotionType": promotionType(o)},
		"tradeScore":         o.TradeScore,
		"topCategoryId":      o.TopCategoryID,
		"secondCategoryId":   o.SecondCategoryID,
		"thirdCategoryId":    o.ThirdCategoryID,
		"sellingPoint":       s.chaosList(o.ID, selling),
		"offerIdentities":    s.chaosList(o.ID, offerIdentities(o)),
		"createDate":         ali.FormatNaive(o.CreateDate),
		// isSelect is documented as a String on this endpoint and a Boolean on
		// the search endpoint. Both spellings are reproduced as documented.
		"isSelect":         strconv.FormatBool(o.IsSelect),
		"certificateList":  s.chaosList(o.ID, s.certificates(o)),
		"promotionUrl":     s.cfg.Base + "/offer/" + strconv.FormatInt(o.ID, 10) + ".html",
		"descriptionTrans": o.SubjectTrans + " — " + strings.Join(o.SellingEN, ", "),
		"companyName":      o.Seller.CompanyName,
		"invoiceInfo":      invoiceInfo(o),
	}
	return famA(s.chaosExtra(o.ID, detail))
}

func channelPrices(o gen.Product) []any {
	out := make([]any, 0, len(o.SKUs))
	for _, sk := range o.SKUs {
		out = append(out, map[string]any{
			"skuId":        sk.SkuID,
			"currentPrice": price(sk.Price - sk.Price/50),
		})
	}
	return out
}

func (s *Server) certificates(o gen.Product) []any {
	if o.ID%13 != 0 {
		return nil
	}
	return []any{map[string]any{
		"certificateName":      "外观专利证书或授权书证书",
		"certificateCode":      "ZL 2018 3 " + strconv.FormatInt(o.ID%10000000, 10) + ".7",
		"certificatePhotoList": []any{s.imageURL(o.ImagePaths[0])},
	}}
}

func (s *Server) sellerDataInfoDetail(o gen.Product) map[string]any {
	return map[string]any{
		"tradeMedalLevel":              o.Seller.MedalLevel,
		"compositeServiceScore":        "3.5",
		"logisticsExperienceScore":     "4.5",
		"disputeComplaintScore":        "3.0",
		"offerExperienceScore":         "4.0",
		"consultingExperienceScore":    "5.0",
		"repeatPurchasePercent":        "0.4666",
		"afterSalesExperienceScore":    "3.0",
		"collect30DayWithin48HPercent": "1",
		"qualityRefundWithin30Day":     "2",
	}
}

// handleCategoryByID serves category.translation.getById. Family A.
//
// The documented parameter table types the top-level leaf and parentCateId as
// java.lang.String while the children type them as Boolean and Long, and the
// sample contradicts the table. The table wins here: a client that assumes one
// JSON type for a logical field across one response deserves to find out.
func (s *Server) handleCategoryByID(form url.Values) any {
	language := form.Get("language")
	categoryID, ok := parseInt64(form.Get("categoryId"))
	if !ok {
		return famAErr("400", "categoryId must be a number")
	}

	archs := gen.Archetypes()

	// Root: the top of the tree, whose children are the six archetypes.
	if categoryID == 0 {
		kids := make([]any, 0, len(archs))
		for _, a := range archs {
			kids = append(kids, map[string]any{
				"categoryId":     a.SecondCategoryID,
				"chineseName":    a.NameZH,
				"translatedName": a.NameEN,
				"language":       language,
				"leaf":           false,
				"level":          "1",
				"parentCateId":   int64(0),
				"fromCache":      true,
			})
		}
		return famA(map[string]any{
			"categoryId":     int64(0),
			"chineseName":    "所有类目",
			"translatedName": "All Categories",
			"language":       language,
			"leaf":           "false",
			"level":          "0",
			"parentCateId":   "0",
			"fromCache":      true,
			"children":       kids,
		})
	}

	for _, a := range archs {
		if a.SecondCategoryID == categoryID {
			kids := make([]any, 0)
			for _, l := range a.LeafCategories {
				kids = append(kids, map[string]any{
					"categoryId":     l.ID,
					"chineseName":    l.ZH,
					"translatedName": l.EN,
					"language":       language,
					"leaf":           true,
					"level":          "2",
					"parentCateId":   a.SecondCategoryID,
					"fromCache":      true,
				})
			}
			return famA(map[string]any{
				"categoryId":     a.SecondCategoryID,
				"chineseName":    a.NameZH,
				"translatedName": a.NameEN,
				"language":       language,
				"leaf":           "false",
				"level":          "1",
				"parentCateId":   strconv.FormatInt(a.TopCategoryID, 10),
				"fromCache":      true,
				"children":       kids,
			})
		}
		for _, l := range a.LeafCategories {
			if l.ID != categoryID {
				continue
			}
			return famA(map[string]any{
				"categoryId":     l.ID,
				"chineseName":    l.ZH,
				"translatedName": l.EN,
				"language":       language,
				"leaf":           "true",
				"level":          "2",
				"parentCateId":   strconv.FormatInt(a.SecondCategoryID, 10),
				"fromCache":      true,
				// A leaf has no children. Documented as an array; sent as null,
				// exactly as the documented sample sends it.
				"children": nil,
			})
		}
	}
	return famAErr("404", "category not found: "+strconv.FormatInt(categoryID, 10))
}

// handleFreightEstimate serves product.freight.estimate. Family A.
//
// toCountryCode is the district code, not a country: 1688's field name, kept
// verbatim so nobody "fixes" it later.
func (s *Server) handleFreightEstimate(form url.Values) any {
	p, ok := jsonObject(form, "productFreightQueryParamsNew")
	if !ok {
		return famAErr("400", "productFreightQueryParamsNew is not a JSON object")
	}
	offerID := mapInt64(p, "offerId")
	if offerID <= 0 {
		return famAErr("400", "offerId is required")
	}
	if mapStr(p, "toProvinceCode") == "" || mapStr(p, "toCityCode") == "" || mapStr(p, "toCountryCode") == "" {
		return famAErr("400", "toProvinceCode, toCityCode and toCountryCode are required")
	}
	qty := mapInt64(p, "totalNum")
	if qty < 1 {
		qty = 1
	}

	o := s.offer(offerID)

	// Charge by weight: a first kilogram, then per additional kilogram, rounded
	// up. Entirely integer arithmetic; the grams never become a float.
	firstFee := ali.Fen(500 + o.ID%800)
	nextFee := ali.Fen(200 + o.ID%400)
	totalG := o.WeightG * qty
	kilos := (totalG + 999) / 1000
	if kilos < 1 {
		kilos = 1
	}
	freight := firstFee + nextFee*ali.Fen(kilos-1)
	freePostage := o.ID%11 == 0
	if freePostage {
		freight = 0
	}

	skuInfos := make([]any, 0, len(o.SKUs))
	for _, sk := range o.SKUs {
		skuInfos = append(skuInfos, map[string]any{
			"skuId":           strconv.FormatInt(sk.SkuID, 10),
			"singleSkuWeight": kg(sk.WeightG),
			"singleSkuWidth":  mm(o.WidthMM),
			"singleSkuHeight": mm(o.HeightMM),
			"singleSkuLength": mm(o.LengthMM),
		})
	}

	return famA(map[string]any{
		"offerId":                     o.ID,
		"freight":                     price(freight),
		"templateId":                  2000000 + o.ID%900000,
		"singleProductWeight":         kg(o.WeightG),
		"singleProductWidth":          mm(o.WidthMM),
		"singleProductHeight":         mm(o.HeightMM),
		"singleProductLength":         mm(o.LengthMM),
		"templateType":                2,
		"templateName":                o.Seller.ShopName + "运费模板",
		"subTemplateType":             0,
		"subTemplateName":             "快递",
		"firstFee":                    price(firstFee),
		"firstUnit":                   "1",
		"nextFee":                     price(nextFee),
		"nextUnit":                    "1",
		"discount":                    "1",
		"chargeType":                  "0",
		"freePostage":                 freePostage,
		"productFreightSkuInfoModels": s.chaosList(o.ID, skuInfos),
		"sizeValueType":               2,
	})
}

// mm renders a millimetre dimension as the centimetres the docs use, as a bare
// decimal with one place.
func mm(v int64) any {
	return rawNumber(strconv.FormatInt(v/10, 10) + "." + strconv.FormatInt(v%10, 10))
}

// kg renders grams as the kilograms the docs use, as a bare decimal with three
// places — 0.001 kg is a documented example value.
func kg(g int64) any {
	whole := g / 1000
	frac := g % 1000
	return rawNumber(strconv.FormatInt(whole, 10) + "." + pad3(frac))
}

func pad3(v int64) string {
	s := strconv.FormatInt(v, 10)
	for len(s) < 3 {
		s = "0" + s
	}
	return s
}

func toAny(ss []string) []any {
	out := make([]any, 0, len(ss))
	for _, s := range ss {
		out = append(out, s)
	}
	return out
}
