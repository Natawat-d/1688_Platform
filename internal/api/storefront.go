package api

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"marketplace/internal/ali"
	"marketplace/internal/catalog"
	"marketplace/internal/order"
	"marketplace/internal/pricing"
	"marketplace/internal/store"
)

func (s *Server) categories(w http.ResponseWriter, r *http.Request) {
	cats, err := s.DB.TopCategories(r.Context())
	if err != nil {
		s.fail(w, r, err, "Could not load categories.")
		return
	}
	out := make([]CategoryDTO, 0, len(cats))
	for _, c := range cats {
		out = append(out, CategoryDTO{ID: id(c.ID), Name: c.Name, Count: c.Count})
	}
	writeJSON(w, http.StatusOK, out)
}

func (s *Server) searchProducts(w http.ResponseWriter, r *http.Request) {
	q := store.SearchQuery{
		Q:          strings.TrimSpace(r.URL.Query().Get("q")),
		CategoryID: queryInt64(r, "cat", 0),
		MinSatang:  queryInt64(r, "min", 0) * 100, // the UI sends whole baht
		MaxSatang:  queryInt64(r, "max", 0) * 100,
		MaxMOQ:     queryInt(r, "moq", 0),
		InStock:    r.URL.Query().Get("instock") == "1" || r.URL.Query().Get("instock") == "true",
		Sort:       r.URL.Query().Get("sort"),
		Page:       queryInt(r, "page", 1),
		Size:       queryInt(r, "size", 24),
	}

	res, err := s.DB.SearchProducts(r.Context(), q)
	if err != nil {
		s.fail(w, r, err, "Search failed.")
		return
	}

	out := SearchResponse{Total: res.Total, Page: res.Page, Size: res.Size, Items: []ProductCardDTO{}}
	for _, p := range res.Items {
		out.Items = append(out.Items, cardDTO(p))
	}
	writeJSON(w, http.StatusOK, out)
}

func (s *Server) getProduct(w http.ResponseWriter, r *http.Request) {
	offerID, err := pathInt(r, "offerId")
	if err != nil {
		writeErr(w, http.StatusBadRequest, "bad_request", "That product id is not valid.")
		return
	}

	fp, err := s.DB.GetProduct(r.Context(), offerID)
	if s.notFound(w, err, "That product") {
		return
	}
	if err != nil {
		s.fail(w, r, err, "Could not load the product.")
		return
	}

	settings, err := s.DB.Settings(r.Context())
	if err != nil {
		s.fail(w, r, err, "Could not load settings.")
		return
	}
	rules, err := s.DB.FeeRules(r.Context())
	if err != nil {
		s.fail(w, r, err, "Could not load fee rules.")
		return
	}
	ps := catalog.ToPricingSettings(settings)
	ruleFor := pricing.Rules(catalog.ToPricingRules(rules), time.Now())

	p := fp.Product
	dto := ProductDTO{
		OfferID:     id(p.OfferID),
		Title:       p.SubjectTrans,
		TitleZh:     p.Subject,
		Description: p.DescriptionTrans,
		Images:      p.Images,
		Seller:      SellerDTO{OpenID: p.SellerOpenID, Name: p.CompanyName, Score: p.TradeScore},
		Status:      p.Status,
		Sellable:    p.Status == "published" && p.Visible,
		MOQ:         int(p.MinOrderQuantity),
		BatchNumber: int(p.BatchNumber),
		QuoteType:   int(p.QuoteType),
		Unit:        p.UnitTrans,
		Stock:       int(p.AmountOnSale),
		Mix: MixDTO{
			General: p.MixGeneral,
			Amount:  cny(ali.Fen(p.MixAmountFen)),
			Number:  int(p.MixNumber),
			Note:    mixNote(p),
		},
		Shipping: ShippingDTO{
			WeightG:  int(p.WeightG),
			LengthMM: int(p.LengthMM),
			WidthMM:  int(p.WidthMM),
			HeightMM: int(p.HeightMM),
		},
		SyncedAt:     p.SyncedAt,
		SKUs:         []SkuDTO{},
		Tiers:        []TierDTO{},
		CategoryPath: []CategoryDTO{},
	}
	if len(dto.Images) == 0 && p.WhiteImage != "" {
		dto.Images = []string{p.WhiteImage}
	}
	if p.CategoryName != "" {
		dto.CategoryPath = append(dto.CategoryPath, CategoryDTO{ID: id(p.TopCategoryID), Name: p.CategoryName})
	}

	// Every displayed price runs through the same engine checkout uses, quoted at
	// the minimum order quantity, so the shop can never advertise a price it will
	// not honour.
	moq := max(int(p.MinOrderQuantity), 1)
	pp := catalog.ToPricingProduct(fp)
	quoteAt := func(sku store.SKU, qty int) ali.Satang {
		res := pricing.Quote(
			[]pricing.Line{{Product: pp, SKU: catalog.ToPricingSKU(sku), Quantity: qty}},
			map[int64]int{p.OfferID: qty}, ruleFor, ps)
		if len(res.Lines) == 0 {
			return 0
		}
		return res.Lines[0].UnitSatang
	}

	var minPrice, maxPrice ali.Satang
	for _, sk := range fp.SKUs {
		unit := quoteAt(sk, moq)
		dto.SKUs = append(dto.SKUs, SkuDTO{
			SkuID:     id(sk.SkuID),
			SpecID:    sk.SpecID,
			Label:     sk.Label,
			Attrs:     attrsOf(sk.Attrs),
			Stock:     int(sk.AmountOnSale),
			Available: sk.AmountOnSale > 0,
			Price:     thb(unit),
		})
		if unit > 0 && (minPrice == 0 || unit < minPrice) {
			minPrice = unit
		}
		if unit > maxPrice {
			maxPrice = unit
		}
	}
	dto.PriceRange.Min = thb(minPrice)
	dto.PriceRange.Max = thb(maxPrice)

	// Tier prices are quoted on total quantity, so each tier is priced at its own
	// threshold, which is exactly what the shopper would pay to reach it.
	if len(fp.SKUs) > 0 {
		for _, t := range fp.Tiers {
			qty := max(int(t.StartQuantity), moq)
			dto.Tiers = append(dto.Tiers, TierDTO{
				StartQuantity: int(t.StartQuantity),
				Price:         thb(quoteAt(fp.SKUs[0], qty)),
			})
		}
	}

	writeJSON(w, http.StatusOK, dto)
}

func mixNote(p store.Product) string {
	if !p.MixGeneral {
		return ""
	}
	// The rule is satisfied by either measure, never both, so the wording has to
	// say "or" or shoppers will over-buy.
	parts := []string{}
	if p.MixNumber > 0 {
		parts = append(parts, strconv.Itoa(int(p.MixNumber))+" items")
	}
	if p.MixAmountFen > 0 {
		parts = append(parts, ali.Fen(p.MixAmountFen).Text())
	}
	if len(parts) == 0 {
		return ""
	}
	return "This supplier accepts mixed orders of at least " + strings.Join(parts, " or ") + "."
}

// ---------------------------------------------------------------------- cart --

func (s *Server) getCart(w http.ResponseWriter, r *http.Request) {
	cartID, err := s.cartID(w, r)
	if err != nil {
		s.fail(w, r, err, "Could not open your cart.")
		return
	}
	q, err := s.Orders.QuoteCart(r.Context(), cartID)
	if err != nil {
		s.fail(w, r, err, "Could not price your cart.")
		return
	}
	writeJSON(w, http.StatusOK, cartDTO(q))
}

func (s *Server) addCartItem(w http.ResponseWriter, r *http.Request) {
	var in struct {
		OfferID  string `json:"offerId"`
		SkuID    string `json:"skuId"`
		Quantity int    `json:"quantity"`
	}
	if err := decode(r, &in); err != nil {
		writeErr(w, http.StatusBadRequest, "bad_request", "That request could not be read.")
		return
	}
	offerID, err1 := strconv.ParseInt(in.OfferID, 10, 64)
	skuID, err2 := strconv.ParseInt(in.SkuID, 10, 64)
	if err1 != nil || err2 != nil || in.Quantity <= 0 {
		writeErr(w, http.StatusBadRequest, "bad_request", "Choose an option and a quantity first.")
		return
	}

	cartID, err := s.cartID(w, r)
	if err != nil {
		s.fail(w, r, err, "Could not open your cart.")
		return
	}

	fp, err := s.DB.GetProduct(r.Context(), offerID)
	if s.notFound(w, err, "That product") {
		return
	}
	if err != nil {
		s.fail(w, r, err, "Could not load the product.")
		return
	}
	specID := ""
	for _, sk := range fp.SKUs {
		if sk.SkuID == skuID {
			specID = sk.SpecID
		}
	}
	if specID == "" {
		writeErr(w, http.StatusBadRequest, "bad_request", "That option is no longer available.")
		return
	}

	if err := s.DB.AddCartItem(r.Context(), store.CartItem{
		CartID: cartID, OfferID: offerID, SkuID: skuID, SpecID: specID, Quantity: int32(in.Quantity),
	}); err != nil {
		s.fail(w, r, err, "Could not add that to your cart.")
		return
	}
	s.getCart(w, r)
}

func (s *Server) updateCartItem(w http.ResponseWriter, r *http.Request) {
	itemID, err := pathInt(r, "id")
	if err != nil {
		writeErr(w, http.StatusBadRequest, "bad_request", "That cart line is not valid.")
		return
	}
	var in struct {
		Quantity int `json:"quantity"`
	}
	if err := decode(r, &in); err != nil {
		writeErr(w, http.StatusBadRequest, "bad_request", "That request could not be read.")
		return
	}
	cartID, err := s.cartID(w, r)
	if err != nil {
		s.fail(w, r, err, "Could not open your cart.")
		return
	}
	if in.Quantity <= 0 {
		if err := s.DB.DeleteCartItem(r.Context(), cartID, itemID); err != nil {
			s.fail(w, r, err, "Could not update your cart.")
			return
		}
	} else if err := s.DB.SetCartItemQty(r.Context(), cartID, itemID, in.Quantity); err != nil {
		s.fail(w, r, err, "Could not update your cart.")
		return
	}
	s.getCart(w, r)
}

func (s *Server) deleteCartItem(w http.ResponseWriter, r *http.Request) {
	itemID, err := pathInt(r, "id")
	if err != nil {
		writeErr(w, http.StatusBadRequest, "bad_request", "That cart line is not valid.")
		return
	}
	cartID, err := s.cartID(w, r)
	if err != nil {
		s.fail(w, r, err, "Could not open your cart.")
		return
	}
	if err := s.DB.DeleteCartItem(r.Context(), cartID, itemID); err != nil {
		s.fail(w, r, err, "Could not update your cart.")
		return
	}
	s.getCart(w, r)
}

// ------------------------------------------------------------------ checkout --

func (s *Server) checkout(w http.ResponseWriter, r *http.Request) {
	var in struct {
		Email   string            `json:"email"`
		Name    string            `json:"name"`
		Phone   string            `json:"phone"`
		Address map[string]string `json:"address"`
	}
	if err := decode(r, &in); err != nil {
		writeErr(w, http.StatusBadRequest, "bad_request", "That request could not be read.")
		return
	}
	if strings.TrimSpace(in.Email) == "" || strings.TrimSpace(in.Name) == "" {
		writeErr(w, http.StatusBadRequest, "bad_request", "Name and email are required.")
		return
	}

	cartID, err := s.cartID(w, r)
	if err != nil {
		s.fail(w, r, err, "Could not open your cart.")
		return
	}

	// The gateway preview happens inside this call, before any money is taken.
	ctx, cancel := ctxTimeout(r, 60*time.Second)
	defer cancel()

	res, err := s.Orders.Checkout(ctx, order.CheckoutInput{
		Email: in.Email, Name: in.Name, Phone: in.Phone, Address: in.Address, CartID: cartID,
	})
	if errors.Is(err, order.ErrNotCheckoutable) {
		out := ErrorResponse{
			Error:   "not_checkoutable",
			Message: "Some items need attention before you can order.",
		}
		for _, is := range res.Issues {
			out.Issues = append(out.Issues, issueDTO(is))
		}
		writeJSON(w, http.StatusConflict, out)
		return
	}
	if err != nil {
		s.fail(w, r, err, "Checkout failed.")
		return
	}

	o := res.Order
	parcels := 1
	if sos, err := s.DB.SupplierOrders(r.Context(), o.ID); err == nil {
		parcels = len(sos)
	}
	writeJSON(w, http.StatusOK, map[string]any{
		"orderId": o.PublicID,
		"token":   o.AccessToken,
		"total":   thb(ali.Satang(o.TotalSatang)),
		"parcels": parcels,
		"payUrl":  "/order/" + o.PublicID + "?t=" + o.AccessToken,
	})
}

// ------------------------------------------------------------------- orders --

// loadOrder resolves an order from its public id and URL secret. Anonymous
// orders have no account behind them, so without the secret the id alone would
// be a guessable handle on someone else's order.
func (s *Server) loadOrder(w http.ResponseWriter, r *http.Request) (store.Order, bool) {
	publicID := r.PathValue("publicId")
	o, err := s.DB.GetOrderByPublicID(r.Context(), publicID)
	if s.notFound(w, err, "That order") {
		return o, false
	}
	if err != nil {
		s.fail(w, r, err, "Could not load the order.")
		return o, false
	}
	token := r.URL.Query().Get("t")
	if subtleCompare(token, o.AccessToken) != 1 {
		writeErr(w, http.StatusNotFound, "not_found", "That order was not found.")
		return o, false
	}
	return o, true
}

func (s *Server) getOrder(w http.ResponseWriter, r *http.Request) {
	o, ok := s.loadOrder(w, r)
	if !ok {
		return
	}
	dto, err := s.orderView(r, o)
	if err != nil {
		s.fail(w, r, err, "Could not load the order.")
		return
	}
	writeJSON(w, http.StatusOK, dto)
}

func (s *Server) orderView(r *http.Request, o store.Order) (OrderDTO, error) {
	items, err := s.DB.OrderItems(r.Context(), o.ID)
	if err != nil {
		return OrderDTO{}, err
	}
	sos, err := s.DB.SupplierOrders(r.Context(), o.ID)
	if err != nil {
		return OrderDTO{}, err
	}
	ids := make([]int64, 0, len(sos))
	for _, so := range sos {
		ids = append(ids, so.ID)
	}
	ships, err := s.DB.ShipmentsFor(r.Context(), ids)
	if err != nil {
		return OrderDTO{}, err
	}
	shipIDs := make([]int64, 0, len(ships))
	for _, sh := range ships {
		shipIDs = append(shipIDs, sh.ID)
	}
	evs, err := s.DB.TrackingEventsFor(r.Context(), shipIDs)
	if err != nil {
		return OrderDTO{}, err
	}
	return orderDTO(o, items, sos, ships, evs), nil
}

// payOrder stands in for the customer-facing payment provider. Taking the money
// is out of scope here; what matters is that payment is what releases the relay.
func (s *Server) payOrder(w http.ResponseWriter, r *http.Request) {
	o, ok := s.loadOrder(w, r)
	if !ok {
		return
	}
	if o.PaidAt != nil {
		dto, _ := s.orderView(r, o)
		writeJSON(w, http.StatusOK, dto)
		return
	}
	if err := s.Orders.MarkPaid(r.Context(), o); err != nil {
		s.fail(w, r, err, "Could not record your payment.")
		return
	}
	fresh, err := s.DB.GetOrder(r.Context(), o.ID)
	if err != nil {
		s.fail(w, r, err, "Could not load the order.")
		return
	}
	dto, err := s.orderView(r, fresh)
	if err != nil {
		s.fail(w, r, err, "Could not load the order.")
		return
	}
	writeJSON(w, http.StatusOK, dto)
}

func (s *Server) cancelOrder(w http.ResponseWriter, r *http.Request) {
	o, ok := s.loadOrder(w, r)
	if !ok {
		return
	}
	sos, err := s.DB.SupplierOrders(r.Context(), o.ID)
	if err != nil {
		s.fail(w, r, err, "Could not load the order.")
		return
	}
	for _, so := range sos {
		if so.Status == order.StatusCancelled || so.Status == order.StatusAtWarehouse {
			continue
		}
		if err := s.Orders.CancelSupplierOrder(r.Context(), so, "buyerCancel", "cancelled by the buyer"); err != nil {
			s.fail(w, r, err, "Could not cancel the order.")
			return
		}
	}
	writeJSON(w, http.StatusOK, map[string]any{"ok": true, "message": "Cancellation requested."})
}

func subtleCompare(a, b string) int {
	if len(a) != len(b) || len(a) == 0 {
		return 0
	}
	var v byte
	for i := 0; i < len(a); i++ {
		v |= a[i] ^ b[i]
	}
	if v == 0 {
		return 1
	}
	return 0
}
