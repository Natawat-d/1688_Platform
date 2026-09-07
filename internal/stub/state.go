package stub

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	"marketplace/internal/ali"
	"marketplace/internal/stub/gen"
)

// Clock is the only source of time in this package. Every deadline, timestamp
// and trace node goes through it, so a test can drive a whole order lifecycle in
// microseconds with no sleeping and no flakiness.
type Clock interface{ Now() time.Time }

type systemClock struct{}

func (systemClock) Now() time.Time { return time.Now() }

// SystemClock is the wall clock, used by the real binary.
func SystemClock() Clock { return systemClock{} }

// TestClock is a Clock a test moves by hand.
type TestClock struct {
	mu sync.Mutex
	t  time.Time
}

// NewTestClock starts a hand-driven clock at t.
func NewTestClock(t time.Time) *TestClock { return &TestClock{t: t} }

func (c *TestClock) Now() time.Time {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.t
}

// Advance moves the clock forward.
func (c *TestClock) Advance(d time.Duration) {
	c.mu.Lock()
	c.t = c.t.Add(d)
	c.mu.Unlock()
}

// The documented order statuses. terminated exists in the docs too but nothing
// in this stub produces it.
const (
	StatusWaitPay      = "waitbuyerpay"
	StatusWaitSend     = "waitsellersend"
	StatusWaitReceive  = "waitbuyerreceive"
	StatusConfirmGoods = "confirm_goods"
	StatusSuccess      = "success"
	StatusCancel       = "cancel"
)

// Store errors, mapped onto documented business codes by the handlers.
var (
	errNotExist = errors.New("stub: order does not exist")
	errStatus   = errors.New("stub: order status does not allow this")
	errTooFast  = errors.New("stub: order closed less than ten seconds after creation")
)

// closeGrace is the documented CLOSE_ORDER_TOO_FAST window.
const closeGrace = 10 * time.Second

// NameValue is one sku attribute as the order APIs render it.
type NameValue struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// Line is one product entry of an order.
type Line struct {
	SubItemID       int64       `json:"subItemId"`
	OfferID         int64       `json:"offerId"`
	SkuID           int64       `json:"skuId"`
	SpecID          string      `json:"specId"`
	Quantity        int64       `json:"quantity"`
	UnitPrice       ali.Fen     `json:"unitPrice"`
	Amount          ali.Fen     `json:"amount"`
	Name            string      `json:"name"`
	NameTrans       string      `json:"nameTrans"`
	Unit            string      `json:"unit"`
	CargoNumber     string      `json:"cargoNumber"`
	Images          []string    `json:"images"`
	SkuAttrs        []NameValue `json:"skuAttrs"`
	Status          string      `json:"status"`
	LogisticsStatus int         `json:"logisticsStatus"`
}

// TraceStep is one logistics tracking node.
type TraceStep struct {
	At     time.Time `json:"at"`
	Status string    `json:"status"` // CONSIGN, ACCEPT, TRANSPORT, DELIVERING, SIGN
	Remark string    `json:"remark"`
}

// Logistics is the waybill created when an order ships.
type Logistics struct {
	LogisticsID string      `json:"logisticsId"`
	MailNo      string      `json:"mailNo"`
	CPCode      string      `json:"cpCode"`
	CompanyName string      `json:"companyName"`
	CompanyID   int64       `json:"companyId"`
	CreatedAt   time.Time   `json:"createdAt"`
	Steps       []TraceStep `json:"steps"`
}

// Address is the shipping address as it was submitted.
type Address struct {
	AddressID    int64  `json:"addressId"`
	FullName     string `json:"fullName"`
	Mobile       string `json:"mobile"`
	Phone        string `json:"phone"`
	PostCode     string `json:"postCode"`
	ProvinceText string `json:"provinceText"`
	CityText     string `json:"cityText"`
	AreaText     string `json:"areaText"`
	TownText     string `json:"townText"`
	Address      string `json:"address"`
	DistrictCode string `json:"districtCode"`
	TownCode     string `json:"townCode"`
}

// Order is one transaction. Field names here are ours, not 1688's: the
// documented spellings appear only where responses are built.
type Order struct {
	ID         int64  `json:"id"`
	OutOrderID string `json:"outOrderId"`
	Status     string `json:"status"`
	Flow       string `json:"flow"`
	TradeType  string `json:"tradeType"`
	Message    string `json:"message"`

	BuyerMemberID string `json:"buyerMemberId"`
	BuyerUserID   int64  `json:"buyerUserId"`
	BuyerLoginID  string `json:"buyerLoginId"`

	SellerMemberID string `json:"sellerMemberId"`
	SellerUserID   int64  `json:"sellerUserId"`
	SellerLoginID  string `json:"sellerLoginId"`
	SellerCompany  string `json:"sellerCompany"`
	SellerProvince string `json:"sellerProvince"`
	SellerCity     string `json:"sellerCity"`
	SellerDistrict string `json:"sellerDistrict"`

	Lines    []Line  `json:"lines"`
	PostFee  ali.Fen `json:"postFee"`
	Discount ali.Fen `json:"discount"`
	Total    ali.Fen `json:"total"`

	Address       Address `json:"address"`
	AlipayTradeID string  `json:"alipayTradeId"`
	PayChannel    string  `json:"payChannel"`

	CreateTime   time.Time `json:"createTime"`
	ModifyTime   time.Time `json:"modifyTime"`
	PayTime      time.Time `json:"payTime"`
	ShipTime     time.Time `json:"shipTime"`
	ReceiveTime  time.Time `json:"receiveTime"`
	CompleteTime time.Time `json:"completeTime"`
	CloseTime    time.Time `json:"closeTime"`

	// NextAt is when the lifecycle ticker should look at this order again. Zero
	// means the order is resting: paid-on-demand, cancelled or finished.
	NextAt   time.Time `json:"nextAt"`
	TraceIdx int       `json:"traceIdx"`

	CloseReason      string `json:"closeReason"`
	CloseOperateType string `json:"closeOperateType"`
	CloseRemark      string `json:"closeRemark"`

	Logistics *Logistics `json:"logistics,omitempty"`
}

// SumProduct is the order total excluding shipping.
func (o *Order) SumProduct() ali.Fen {
	var t ali.Fen
	for _, l := range o.Lines {
		t += l.Amount
	}
	return t
}

// Event is one lifecycle transition, handed to the push machinery. It carries
// everything a message payload needs so nothing has to reach back into the store
// while the lock is held.
type Event struct {
	Topic          string
	OrderID        int64
	SubOrderID     int64
	CurrentStatus  string
	BuyerMemberID  string
	SellerMemberID string
	At             time.Time
	LogisticsID    string
	MailNo         string
	CPCode         string
	StatusChanged  string
}

// Delays are the lifecycle waits before the speed multiplier is applied.
type Delays struct {
	Pay  time.Duration
	Ship time.Duration
	Step time.Duration
	Sign time.Duration
}

// StoreConfig is everything the order store needs.
type StoreConfig struct {
	Delays    Delays
	Speed     float64
	AutoPay   bool
	StateFile string
}

// Store holds every order behind one mutex and snapshots itself to disk so a
// restart mid-demo does not lose the orders someone just placed.
type Store struct {
	mu      sync.Mutex
	clock   Clock
	orders  map[int64]*Order
	byOut   map[string]int64
	seq     []int64 // insertion order, for stable listing
	nextSeq int64
	creates int64

	delays  Delays
	speed   float64
	autoPay bool
	running bool

	path  string
	dirty bool
}

// orderIDBase is chosen so generated ids look like the eighteen-digit ones in
// the documented samples.
const orderIDBase int64 = 105581756010628640

// NewStore builds an empty store and loads the snapshot if there is one.
func NewStore(cfg StoreConfig, clock Clock) *Store {
	if cfg.Speed <= 0 {
		cfg.Speed = 1
	}
	s := &Store{
		clock:   clock,
		orders:  map[int64]*Order{},
		byOut:   map[string]int64{},
		delays:  cfg.Delays,
		speed:   cfg.Speed,
		autoPay: cfg.AutoPay,
		running: true,
		path:    cfg.StateFile,
	}
	s.load()
	return s
}

// scaled applies the speed multiplier: STUB_SPEED=60 makes a ten-minute wait
// last ten seconds.
func (s *Store) scaled(d time.Duration) time.Duration {
	if s.speed <= 0 {
		return d
	}
	return time.Duration(float64(d) / s.speed)
}

// CreateParams is one createCrossOrder request, already validated.
type CreateParams struct {
	OutOrderID string
	Flow       string
	TradeType  string
	Message    string
	Address    Address
	Lines      []Line
	PostFee    ali.Fen
	Discount   ali.Fen
	Seller     gen.Seller
	Buyer      Buyer
	// Split asks for the documented "several orders created at once" path: the
	// lines are dealt across two orders instead of one.
	Split bool
}

// Buyer is the authorised user this gateway pretends to be.
type Buyer struct {
	MemberID string
	UserID   int64
	LoginID  string
	Company  string
}

// Create makes one or two orders. outOrderId is honoured idempotently: the same
// external id always returns the same orders, which is the whole reason the
// field exists.
func (s *Store) Create(p CreateParams) ([]*Order, []Event, bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if p.OutOrderID != "" {
		if id, ok := s.byOut[p.OutOrderID]; ok {
			var out []*Order
			for _, oid := range s.seq {
				o := s.orders[oid]
				if o != nil && o.OutOrderID == p.OutOrderID {
					out = append(out, o)
				}
			}
			if len(out) == 0 {
				if o := s.orders[id]; o != nil {
					out = append(out, o)
				}
			}
			return out, nil, true, nil
		}
	}

	now := s.clock.Now()
	groups := [][]Line{p.Lines}
	if p.Split && len(p.Lines) > 1 {
		half := len(p.Lines) / 2
		groups = [][]Line{p.Lines[:half], p.Lines[half:]}
	} else if p.Split {
		// One line but a split was asked for: duplicate the order, which is what
		// the documented "multiple orders" sample looks like from outside.
		groups = [][]Line{p.Lines, p.Lines}
	}

	var orders []*Order
	var events []Event
	for gi, lines := range groups {
		s.nextSeq++
		id := orderIDBase + s.nextSeq*97
		o := &Order{
			ID:               id,
			OutOrderID:       p.OutOrderID,
			Status:           StatusWaitPay,
			Flow:             p.Flow,
			TradeType:        p.TradeType,
			Message:          p.Message,
			BuyerMemberID:    p.Buyer.MemberID,
			BuyerUserID:      p.Buyer.UserID,
			BuyerLoginID:     p.Buyer.LoginID,
			SellerMemberID:   p.Seller.MemberID,
			SellerUserID:     p.Seller.UserID,
			SellerLoginID:    p.Seller.LoginID,
			SellerCompany:    p.Seller.CompanyName,
			SellerProvince:   p.Seller.Province,
			SellerCity:       p.Seller.City,
			SellerDistrict:   p.Seller.District,
			Address:          p.Address,
			PostFee:          p.PostFee,
			Discount:         p.Discount,
			CreateTime:       now,
			ModifyTime:       now,
			AlipayTradeID:    strconv.FormatInt(2000000000000000+id%1000000000000000, 10),
			PayChannel:       "alipay",
			CloseOperateType: "",
		}
		for i, l := range lines {
			l.SubItemID = id + int64(gi*100+i) + 1
			l.Status = StatusWaitSend
			l.LogisticsStatus = 1
			o.Lines = append(o.Lines, l)
		}
		o.Total = o.SumProduct() + o.PostFee - o.Discount
		if s.autoPay {
			o.NextAt = now.Add(s.scaled(s.delays.Pay))
		}

		s.orders[id] = o
		s.seq = append(s.seq, id)
		if p.OutOrderID != "" {
			s.byOut[p.OutOrderID] = id
		}
		orders = append(orders, o)
		events = append(events, Event{
			Topic:          TopicBuyerMake,
			OrderID:        o.ID,
			CurrentStatus:  StatusWaitPay,
			BuyerMemberID:  o.BuyerMemberID,
			SellerMemberID: o.SellerMemberID,
			At:             now,
		})
	}

	s.creates++
	s.markDirty()
	return orders, events, false, nil
}

// Creates is how many successful create calls have been served. The stub uses it
// to return two orders on every fifth one.
func (s *Store) Creates() int64 {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.creates
}

// Get returns a copy-free pointer to an order. Callers read it under no lock,
// which is safe here because the lifecycle only ever appends and the stub is a
// development tool, not a database.
func (s *Store) Get(id int64) (*Order, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	o, ok := s.orders[id]
	return o, ok
}

// GetByOut resolves an external order id.
func (s *Store) GetByOut(outOrderID string) (*Order, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	id, ok := s.byOut[outOrderID]
	if !ok {
		return nil, false
	}
	o, ok := s.orders[id]
	return o, ok
}

// Filter narrows a listing the way getBuyerOrderList documents.
type Filter struct {
	Status     string
	OrderIDs   []int64
	OutOrderID string
	SellerID   string
	CreatedGTE time.Time
	CreatedLTE time.Time
	ModifiedGTE,
	ModifiedLTE time.Time
}

// List returns the matching orders, newest first, and the total before paging.
func (s *Store) List(f Filter, page, pageSize int) ([]*Order, int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	wanted := map[int64]bool{}
	for _, id := range f.OrderIDs {
		wanted[id] = true
	}

	var all []*Order
	for i := len(s.seq) - 1; i >= 0; i-- {
		o := s.orders[s.seq[i]]
		if o == nil {
			continue
		}
		if f.Status != "" && o.Status != f.Status {
			continue
		}
		if len(wanted) > 0 && !wanted[o.ID] {
			continue
		}
		if f.OutOrderID != "" && o.OutOrderID != f.OutOrderID {
			continue
		}
		if f.SellerID != "" && o.SellerMemberID != f.SellerID {
			continue
		}
		if !f.CreatedGTE.IsZero() && o.CreateTime.Before(f.CreatedGTE) {
			continue
		}
		if !f.CreatedLTE.IsZero() && o.CreateTime.After(f.CreatedLTE) {
			continue
		}
		if !f.ModifiedGTE.IsZero() && o.ModifyTime.Before(f.ModifiedGTE) {
			continue
		}
		if !f.ModifiedLTE.IsZero() && o.ModifyTime.After(f.ModifiedLTE) {
			continue
		}
		all = append(all, o)
	}

	total := len(all)
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 20
	}
	from := (page - 1) * pageSize
	if from >= total {
		return nil, total
	}
	to := from + pageSize
	if to > total {
		to = total
	}
	return all[from:to], total
}

// All returns every order, oldest first, for /_control/state.
func (s *Store) All() []*Order {
	s.mu.Lock()
	defer s.mu.Unlock()
	out := make([]*Order, 0, len(s.seq))
	for _, id := range s.seq {
		if o := s.orders[id]; o != nil {
			out = append(out, o)
		}
	}
	return out
}

// Pay marks an order paid. The cashier page and /_control call it directly;
// autopay reaches it through Tick.
func (s *Store) Pay(id int64) ([]Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	o, ok := s.orders[id]
	if !ok {
		return nil, errNotExist
	}
	if o.Status != StatusWaitPay {
		return nil, errStatus
	}
	return s.pay(o, s.clock.Now()), nil
}

// Cancel closes an unpaid order, enforcing the two documented refusals.
func (s *Store) Cancel(id int64, reason, remark string) ([]Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	o, ok := s.orders[id]
	if !ok {
		return nil, errNotExist
	}
	now := s.clock.Now()
	if now.Sub(o.CreateTime) < closeGrace {
		return nil, errTooFast
	}
	if o.Status != StatusWaitPay {
		return nil, errStatus
	}

	o.Status = StatusCancel
	o.CloseReason = reason
	o.CloseRemark = remark
	o.CloseTime = now
	o.ModifyTime = now
	o.NextAt = time.Time{}
	topic := TopicBuyerClose
	o.CloseOperateType = "CLOSE_TRADE_BY_BUYER"
	if reason == "sellerGoodsLack" {
		topic = TopicSellerClose
		o.CloseOperateType = "CLOSE_TRADE_BY_SELLER"
	}
	for i := range o.Lines {
		o.Lines[i].Status = StatusCancel
	}
	s.markDirty()
	return []Event{s.event(o, topic, StatusCancel, now)}, nil
}

// AdvanceTo forces an order into a status, for demos and for /_control.
func (s *Store) AdvanceTo(id int64, status string) ([]Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	o, ok := s.orders[id]
	if !ok {
		return nil, errNotExist
	}
	now := s.clock.Now()

	order := []string{StatusWaitPay, StatusWaitSend, StatusWaitReceive, StatusConfirmGoods, StatusSuccess}
	idx := func(st string) int {
		for i, v := range order {
			if v == st {
				return i
			}
		}
		return -1
	}
	if status == StatusCancel {
		o.Status = StatusCancel
		o.CloseReason = "other"
		o.CloseOperateType = "CLOSE_TRADE_BY_BOPS"
		o.CloseTime = now
		o.ModifyTime = now
		o.NextAt = time.Time{}
		s.markDirty()
		return []Event{s.event(o, TopicSellerClose, StatusCancel, now)}, nil
	}
	target, cur := idx(status), idx(o.Status)
	if target < 0 {
		return nil, errStatus
	}
	if cur < 0 || target <= cur {
		return nil, errStatus
	}

	var events []Event
	for cur < target {
		events = append(events, s.step(o, now)...)
		cur = idx(o.Status)
		if cur < 0 {
			break
		}
	}
	return events, nil
}

// AddTrace appends a logistics node by hand. The order must have shipped.
func (s *Store) AddTrace(id int64, statusChanged, remark string) ([]Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	o, ok := s.orders[id]
	if !ok {
		return nil, errNotExist
	}
	if o.Logistics == nil {
		return nil, errStatus
	}
	now := s.clock.Now()
	o.Logistics.Steps = append(o.Logistics.Steps, TraceStep{At: now, Status: statusChanged, Remark: remark})
	o.ModifyTime = now
	s.markDirty()
	ev := s.event(o, TopicLogisticsTrace, o.Status, now)
	ev.StatusChanged = statusChanged
	ev.LogisticsID = o.Logistics.LogisticsID
	ev.MailNo = o.Logistics.MailNo
	ev.CPCode = o.Logistics.CPCode
	return []Event{ev}, nil
}

// SetClockRunning starts or stops the lifecycle ticker.
func (s *Store) SetClockRunning(running bool, speed float64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.running = running
	if speed > 0 {
		s.speed = speed
	}
	s.markDirty()
}

// ClockState reports what the ticker is doing.
func (s *Store) ClockState() (running bool, speed float64, autoPay bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.running, s.speed, s.autoPay
}

// Reset empties the store. /_control/reset uses it.
func (s *Store) Reset() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.orders = map[int64]*Order{}
	s.byOut = map[string]int64{}
	s.seq = nil
	s.nextSeq = 0
	s.creates = 0
	s.markDirty()
}

// Tick advances every order whose next transition is due and returns the
// messages that should now be published.
func (s *Store) Tick() []Event {
	s.mu.Lock()
	defer s.mu.Unlock()
	if !s.running {
		return nil
	}
	now := s.clock.Now()
	var events []Event
	for _, id := range s.seq {
		o := s.orders[id]
		if o == nil || o.NextAt.IsZero() || now.Before(o.NextAt) {
			continue
		}
		events = append(events, s.step(o, now)...)
	}
	if len(events) > 0 {
		s.markDirty()
	}
	return events
}

// step performs exactly one transition. Callers hold the lock.
func (s *Store) step(o *Order, now time.Time) []Event {
	switch o.Status {
	case StatusWaitPay:
		return s.pay(o, now)
	case StatusWaitSend:
		return s.ship(o, now)
	case StatusWaitReceive:
		return s.deliver(o, now)
	case StatusConfirmGoods:
		o.Status = StatusSuccess
		o.CompleteTime = now
		o.ModifyTime = now
		o.NextAt = time.Time{}
		return []Event{s.event(o, TopicOrderSuccess, StatusSuccess, now)}
	default:
		o.NextAt = time.Time{}
		return nil
	}
}

func (s *Store) pay(o *Order, now time.Time) []Event {
	o.Status = StatusWaitSend
	o.PayTime = now
	o.ModifyTime = now
	o.NextAt = now.Add(s.scaled(s.delays.Ship))
	for i := range o.Lines {
		o.Lines[i].Status = StatusWaitSend
	}
	s.markDirty()
	return []Event{s.event(o, TopicOrderPay, StatusWaitSend, now)}
}

// ship creates the waybill and the first tracking node.
func (s *Store) ship(o *Order, now time.Time) []Event {
	code, name := gen.Carrier(o.ID)
	lg := &Logistics{
		LogisticsID: "LP" + strconv.FormatInt(100000000000000+o.ID%900000000000000, 10),
		MailNo:      code + strconv.FormatInt(300000000000+o.ID%700000000000, 10),
		CPCode:      code,
		CompanyName: name,
		CompanyID:   1000000 + o.ID%900000,
		CreatedAt:   now,
	}
	o.Logistics = lg
	o.Status = StatusWaitReceive
	o.ShipTime = now
	o.ModifyTime = now
	o.TraceIdx = 0
	for i := range o.Lines {
		o.Lines[i].Status = StatusWaitReceive
		o.Lines[i].LogisticsStatus = 2
	}

	var events []Event
	// A multi-line order ships in two goes, which is what the documented
	// PART_PART_SENDGOODS message describes: still waiting for the seller.
	if len(o.Lines) > 1 {
		events = append(events, s.event(o, TopicPartSendGoods, StatusWaitSend, now))
	}
	events = append(events, s.event(o, TopicAnnounceSendGoods, StatusWaitReceive, now))
	events = append(events, s.traceNode(o, now)...)
	o.NextAt = now.Add(s.scaled(s.delays.Step))
	return events
}

// deliver adds the next tracking node, and confirms receipt after the last one.
func (s *Store) deliver(o *Order, now time.Time) []Event {
	events := s.traceNode(o, now)
	if o.TraceIdx >= len(traceNodes) {
		o.Status = StatusConfirmGoods
		o.ReceiveTime = now
		o.ModifyTime = now
		o.NextAt = now.Add(s.scaled(s.delays.Step))
		for i := range o.Lines {
			o.Lines[i].Status = StatusConfirmGoods
			o.Lines[i].LogisticsStatus = 3
		}
		// The documented currentStatus for this topic is not the order status.
		ev := s.event(o, TopicConfirmReceive, "confirm_goods_and_has_subsidy", now)
		events = append(events, ev)
		return events
	}
	if o.TraceIdx == len(traceNodes)-1 {
		o.NextAt = now.Add(s.scaled(s.delays.Sign))
	} else {
		o.NextAt = now.Add(s.scaled(s.delays.Step))
	}
	return events
}

// traceNodes is the fixed five-node journey. The remarks are the sort of thing
// a Chinese courier actually writes, because a demo full of "step 3" reads as a
// demo.
var traceNodes = []struct {
	Status string
	Remark func(o *Order) string
}{
	{"CONSIGN", func(o *Order) string {
		return "商品已出库，等待" + o.Logistics.CompanyName + "揽收，运单号" + o.Logistics.MailNo
	}},
	{"ACCEPT", func(o *Order) string {
		return "在" + o.SellerProvince + o.SellerCity + o.SellerDistrict + "公司进行揽件扫描，即将发往：" + o.Address.ProvinceText + "分拨中心"
	}},
	{"TRANSPORT", func(o *Order) string {
		return "在分拨中心" + o.Address.ProvinceText + "分拨中心进行卸车扫描，本次转运目的地：" + o.Address.CityText + o.Address.AreaText + "公司"
	}},
	{"DELIVERING", func(o *Order) string {
		return "在" + o.Address.CityText + o.Address.AreaText + "公司进行派件扫描；派送业务员：徐洲；联系电话：1381234****"
	}},
	{"SIGN", func(o *Order) string {
		return "快件已被 本人 签收，感谢使用" + o.Logistics.CompanyName + "，期待再次为您服务"
	}},
}

// traceNode appends the next node in the journey and returns the trace message.
func (s *Store) traceNode(o *Order, now time.Time) []Event {
	if o.Logistics == nil || o.TraceIdx >= len(traceNodes) {
		return nil
	}
	n := traceNodes[o.TraceIdx]
	o.TraceIdx++
	o.Logistics.Steps = append(o.Logistics.Steps, TraceStep{At: now, Status: n.Status, Remark: n.Remark(o)})

	ev := s.event(o, TopicLogisticsTrace, o.Status, now)
	ev.StatusChanged = n.Status
	ev.LogisticsID = o.Logistics.LogisticsID
	ev.MailNo = o.Logistics.MailNo
	ev.CPCode = o.Logistics.CPCode
	if len(o.Lines) > 0 {
		ev.SubOrderID = o.Lines[0].SubItemID
	}
	return []Event{ev}
}

func (s *Store) event(o *Order, topic, status string, now time.Time) Event {
	ev := Event{
		Topic:          topic,
		OrderID:        o.ID,
		CurrentStatus:  status,
		BuyerMemberID:  o.BuyerMemberID,
		SellerMemberID: o.SellerMemberID,
		At:             now,
	}
	if len(o.Lines) > 0 {
		ev.SubOrderID = o.Lines[0].SubItemID
	}
	return ev
}

// ---------------------------------------------------------------- persistence

type snapshot struct {
	Orders  []*Order `json:"orders"`
	NextSeq int64    `json:"nextSeq"`
	Creates int64    `json:"creates"`
	Speed   float64  `json:"speed"`
	Running bool     `json:"running"`
}

func (s *Store) markDirty() { s.dirty = true }

// Flush writes the snapshot at most once per interval, and only when something
// changed. Debounced because a fast demo transitions orders several times a
// second and rewriting the file each time is silly.
func (s *Store) Flush(ctx context.Context, interval time.Duration) {
	if s.path == "" {
		return
	}
	t := time.NewTicker(interval)
	defer t.Stop()
	for {
		select {
		case <-ctx.Done():
			_ = s.Save()
			return
		case <-t.C:
			s.mu.Lock()
			dirty := s.dirty
			s.mu.Unlock()
			if dirty {
				_ = s.Save()
			}
		}
	}
}

// Save writes the snapshot atomically: temp file in the same directory, then
// rename, so a crash mid-write cannot leave a half-parsed state file behind.
func (s *Store) Save() error {
	if s.path == "" {
		return nil
	}
	s.mu.Lock()
	snap := snapshot{NextSeq: s.nextSeq, Creates: s.creates, Speed: s.speed, Running: s.running}
	for _, id := range s.seq {
		if o := s.orders[id]; o != nil {
			snap.Orders = append(snap.Orders, o)
		}
	}
	b, err := json.MarshalIndent(snap, "", "  ")
	s.dirty = false
	s.mu.Unlock()
	if err != nil {
		return err
	}

	dir := filepath.Dir(s.path)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	tmp, err := os.CreateTemp(dir, ".stub-state-*")
	if err != nil {
		return err
	}
	name := tmp.Name()
	if _, err := tmp.Write(b); err != nil {
		tmp.Close()
		os.Remove(name)
		return err
	}
	if err := tmp.Close(); err != nil {
		os.Remove(name)
		return err
	}
	return os.Rename(name, s.path)
}

func (s *Store) load() {
	if s.path == "" {
		return
	}
	b, err := os.ReadFile(s.path)
	if err != nil {
		return
	}
	var snap snapshot
	if err := json.Unmarshal(b, &snap); err != nil {
		return
	}
	for _, o := range snap.Orders {
		s.orders[o.ID] = o
		s.seq = append(s.seq, o.ID)
		if o.OutOrderID != "" {
			s.byOut[o.OutOrderID] = o.ID
		}
	}
	s.nextSeq = snap.NextSeq
	s.creates = snap.Creates
	if snap.Speed > 0 {
		s.speed = snap.Speed
	}
	s.running = snap.Running || len(snap.Orders) == 0
}
