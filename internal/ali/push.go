package ali

import (
	"context"
	"encoding/json"
	"strings"
	"time"
)

// The message side of the integration: the envelope 1688 pushes to our HTTPS
// endpoint, the payloads of the topics we subscribe to, and the three replay
// calls that fill the gaps when our endpoint was down.
//
// The push envelope and the replay envelope are the same shape, so one
// PushMessage type serves both the receiver and the gap-fill job.

// PushMessage is one message, as delivered to the push endpoint and as returned
// by push.cursor.messageList and push.query.messageList.
//
// Type is redundant with TopicGroup and TopicName — it is exactly
// TopicGroup + "_" + TopicName — but it is the field that is always present, so
// dispatch on Type and treat the other two as extras.
type PushMessage struct {
	MsgID      ID              `json:"msgId"`
	Type       string          `json:"type"`
	UserInfo   string          `json:"userInfo"`
	Data       json.RawMessage `json:"data"`
	GmtBorn    int64           `json:"gmtBorn"`
	TopicGroup string          `json:"topicGroup"`
	AppKey     string          `json:"appKey"`
	TopicName  string          `json:"topicName"`
}

// Topic is the message's topic id: Type when the gateway sent one, otherwise the
// group and name joined back together.
func (m PushMessage) Topic() string {
	if m.Type != "" {
		return m.Type
	}
	return TopicType(m.TopicGroup, m.TopicName)
}

// Born is the creation time of the message. gmtBorn is epoch milliseconds.
func (m PushMessage) Born() time.Time { return time.UnixMilli(m.GmtBorn).In(CST) }

// Decode unmarshals the message payload into v.
func (m PushMessage) Decode(v any) error {
	if len(m.Data) == 0 {
		return nil
	}
	return json.Unmarshal(m.Data, v)
}

// TopicType builds the type field of a message from its group and name: the
// gateway composes it as topicGroup + "_" + topicName, which is also exactly the
// topic id used when subscribing.
func TopicType(group, name string) string {
	if group == "" {
		return name
	}
	if name == "" {
		return group
	}
	return group + "_" + name
}

// Topic groups.
const (
	GroupOrder     = "ORDER"
	GroupLogistics = "LOGISTICS"
	GroupProduct   = "PRODUCT"
)

// The topics this platform subscribes to.
//
// The eight ORDER_BUYER_VIEW_* topics all carry the same OrderEvent payload;
// so do the order topics we do not subscribe to (BOPS_CLOSE, PRICE_MODIFY,
// STEP_PAY), should they ever be added.
//
// Note the misspelling in ORDER_BUYER_VIEW_ORDER_COMFIRM_RECEIVEGOODS. It is
// 1688's, it is on the wire, and correcting it means never matching the topic.
const (
	TopicOrderBuyerMake      = "ORDER_BUYER_VIEW_BUYER_MAKE"
	TopicOrderPay            = "ORDER_BUYER_VIEW_ORDER_PAY"
	TopicOrderAnnounceSend   = "ORDER_BUYER_VIEW_ANNOUNCE_SENDGOODS"
	TopicOrderPartSend       = "ORDER_BUYER_VIEW_PART_PART_SENDGOODS"
	TopicOrderConfirmReceive = "ORDER_BUYER_VIEW_ORDER_COMFIRM_RECEIVEGOODS"
	TopicOrderSuccess        = "ORDER_BUYER_VIEW_ORDER_SUCCESS"
	TopicOrderBuyerClose     = "ORDER_BUYER_VIEW_ORDER_BUYER_CLOSE"
	TopicOrderSellerClose    = "ORDER_BUYER_VIEW_ORDER_SELLER_CLOSE"

	TopicLogisticsTrace = "LOGISTICS_BUYER_VIEW_TRACE"
	TopicMailNoChange   = "LOGISTICS_MAIL_NO_CHANGE"

	TopicInventoryChange = "PRODUCT_PRODUCT_INVENTORY_CHANGE"
	TopicProductChange   = "PRODUCT_RELATION_VIEW_PRODUCT_CHANGE"
)

// orderTopicStatus is the order status each order topic announces, taken from
// the documented payload of that topic. Both sides use it: the stub stamps
// currentStatus from here when it emits a message, and the receiver uses it to
// sanity-check a message whose payload lost its status.
var orderTopicStatus = map[string]string{
	TopicOrderBuyerMake:      StatusWaitBuyerPay,
	TopicOrderPay:            StatusWaitSellerSend,
	TopicOrderAnnounceSend:   StatusWaitBuyerReceive,
	TopicOrderPartSend:       StatusWaitSellerSend,
	TopicOrderConfirmReceive: "confirm_goods_and_has_subsidy",
	TopicOrderSuccess:        StatusSuccess,
	TopicOrderBuyerClose:     StatusCancel,
	TopicOrderSellerClose:    StatusCancel,
}

// TopicStatus is the order status a topic announces, or "" for a topic that is
// not an order topic.
func TopicStatus(topic string) string { return orderTopicStatus[topic] }

// OrderTopics is every order topic we subscribe to, in lifecycle order.
func OrderTopics() []string {
	return []string{
		TopicOrderBuyerMake,
		TopicOrderPay,
		TopicOrderAnnounceSend,
		TopicOrderPartSend,
		TopicOrderConfirmReceive,
		TopicOrderSuccess,
		TopicOrderBuyerClose,
		TopicOrderSellerClose,
	}
}

// Topics is every topic we subscribe to, order topics first.
func Topics() []string {
	return append(OrderTopics(),
		TopicLogisticsTrace,
		TopicMailNoChange,
		TopicInventoryChange,
		TopicProductChange,
	)
}

// ------------------------------------------------------------ topic payloads

// OrderEvent is the payload of every ORDER_BUYER_VIEW_* topic. The topic says
// what happened; currentStatus says where the order landed.
//
// msgSendTime is the naive "2018-05-30 19:24:18" spelling, meaning Shanghai.
// Messages are asynchronous and arrive out of order, so apply one only when its
// time is newer than the state already held.
type OrderEvent struct {
	OrderID        ID        `json:"orderId"`
	CurrentStatus  string    `json:"currentStatus"`
	MsgSendTime    Timestamp `json:"msgSendTime"`
	BuyerMemberID  string    `json:"buyerMemberId"`
	SellerMemberID string    `json:"sellerMemberId"`
}

// LogisticsTracePush is the payload of LOGISTICS_BUYER_VIEW_TRACE. The single
// member is PascalCase on the wire, unlike every other payload key.
type LogisticsTracePush struct {
	Model LogisticsTracingModel `json:"OrderLogisticsTracingModel"`
}

// LogisticsTracingModel is the waybill whose status changed.
type LogisticsTracingModel struct {
	LogisticsID    string         `json:"logisticsId"`
	CpCode         string         `json:"cpCode"`
	MailNo         string         `json:"mailNo"`
	StatusChanged  string         `json:"statusChanged"`
	ChangeTime     Timestamp      `json:"changeTime"`
	OrderLogsItems []OrderLogItem `json:"orderLogsItems"`
}

// The statusChanged values of LOGISTICS_BUYER_VIEW_TRACE. Each fires once, the
// first time the waybill reaches that node.
const (
	TraceConsign    = "CONSIGN"    // shipped
	TraceAccept     = "ACCEPT"     // picked up
	TraceTransport  = "TRANSPORT"  // in transit
	TraceDelivering = "DELIVERING" // out for delivery
	TraceAgentSign  = "AGENT_SIGN" // waiting at a pickup point
	TraceSign       = "SIGN"       // delivered
	TraceFailed     = "FAILED"     // exception
)

// MailNoChangePush is the payload of LOGISTICS_MAIL_NO_CHANGE: the carrier or
// the tracking number of an existing waybill was corrected.
type MailNoChangePush struct {
	Model MailNoChangeModel `json:"MailNoChangeModel"`
}

// MailNoChangeModel is the before and after of a waybill correction.
type MailNoChangeModel struct {
	LogisticsID    string         `json:"logisticsId"`
	OldCpCode      string         `json:"oldCpCode"`
	NewCpCode      string         `json:"newCpCode"`
	OldMailNo      string         `json:"oldMailNo"`
	NewMailNo      string         `json:"newMailNo"`
	EventTime      Timestamp      `json:"eventTime"`
	OrderLogsItems []OrderLogItem `json:"orderLogsItems"`
}

// OrderLogItem ties a waybill message back to an order and one of its lines.
type OrderLogItem struct {
	OrderID      ID `json:"orderId"`
	OrderEntryID ID `json:"orderEntryId"`
}

// InventoryChangePush is the payload of PRODUCT_PRODUCT_INVENTORY_CHANGE. The
// single member is an ARRAY: one message can carry several sku changes.
//
// Stock messages only arrive for offers we have followed through the
// follow-product API.
type InventoryChangePush struct {
	Changes []InventoryChange `json:"OfferInventoryChangeList"`
}

// InventoryChange is one stock movement. Quantity is a SIGNED DELTA, not the new
// level: -10 means ten fewer. The level is offerOnSale / skuOnSale.
type InventoryChange struct {
	OfferID     ID        `json:"offerId"`
	OfferOnSale int64     `json:"offerOnSale"`
	SkuID       ID        `json:"skuId"`
	SkuOnSale   int64     `json:"skuOnSale"`
	Quantity    int64     `json:"quantity"`
	BizTime     FlexInt64 `json:"bizTime"` // epoch millis, quoted in the sample
}

// Time is when the stock moved.
func (c InventoryChange) Time() time.Time { return time.UnixMilli(int64(c.BizTime)).In(CST) }

// ProductChange is the payload of PRODUCT_RELATION_VIEW_PRODUCT_CHANGE, which
// covers every edit a seller can make to an offer.
type ProductChange struct {
	MemberID    string    `json:"memberId"`
	ProductIDs  string    `json:"productIds"` // comma separated, "id,id"
	Action      string    `json:"action"`
	MsgSendTime Timestamp `json:"msgSendTime"`
}

// OfferIDs splits the comma-separated productIds into ids.
func (p ProductChange) OfferIDs() IDs {
	var out IDs
	for _, s := range strings.Split(p.ProductIDs, ",") {
		s = strings.TrimSpace(s)
		if s == "" {
			continue
		}
		var id ID
		if err := id.UnmarshalJSON([]byte(s)); err != nil || id == 0 {
			continue
		}
		out = append(out, id)
	}
	return out
}

// The action values of PRODUCT_RELATION_VIEW_PRODUCT_CHANGE. Anything that is
// not a delist means re-fetch the offer; a delist means hide it.
const (
	ActionNew           = "new"            // seller published
	ActionModify        = "modify"         // seller edited
	ActionMemberDelete  = "member_delete"  // seller deleted
	ActionMemberExpired = "member_expired" // seller delisted
	ActionRepost        = "repost"         // seller reposted
	ActionAudit         = "audit"          // reviewed by 1688 staff
	ActionUpdate        = "update"         // changed by 1688 operations
	ActionRevised       = "revised"        // corrected
	ActionPublish       = "publish"        // listed
	ActionSkuNew        = "sku_new"
	ActionSkuModify     = "sku_modify"
	ActionSkuDelete     = "sku_delete"
)

// ---------------------------------------------------------------- replay calls

// PushQuery selects which stored messages to replay. Both list calls take the
// same filters; Quantity applies to the cursor call and Page/PageSize to the
// query call.
type PushQuery struct {
	CreateStartTime time.Time
	CreateEndTime   time.Time
	Type            string // a topic id, e.g. TopicOrderPay
	UserInfo        string // the member id the message belongs to
	Quantity        int    // cursor style: 20 to 200, default 20
	Page            int    // query style, from 1
	PageSize        int    // query style: 20 to 200, default 20
}

func (q PushQuery) params() Params {
	p := Params{}
	setTime(p, "createStartTime", q.CreateStartTime)
	setTime(p, "createEndTime", q.CreateEndTime)
	setStr(p, "type", q.Type)
	setStr(p, "userInfo", q.UserInfo)
	return p
}

// PushMessagePage is the paged answer of push.query.messageList.
type PushMessagePage struct {
	Datas      []PushMessage `json:"datas"`
	TotalCount int           `json:"totalCount"`
}

// CursorMessages replays undelivered messages, oldest first.
//
// READING CONSUMES: every message this returns is confirmed automatically, so
// the next call with the same filter returns the next batch and a crash between
// the read and our own commit loses those messages for good. Use it for gap fill
// at startup and persist what comes back before asking for more.
func (c *Client) CursorMessages(ctx context.Context, q PushQuery) ([]PushMessage, error) {
	p := q.params()
	if q.Quantity > 0 {
		p["quantity"] = q.Quantity
	}
	body, err := c.Call(ctx, PushCursorList, p)
	if err != nil {
		return nil, err
	}
	var out struct {
		PushMessageList []PushMessage `json:"pushMessageList"`
	}
	if err := DecodeFlat(PushCursorList.Key(), body, &out); err != nil {
		return nil, err
	}
	return out.PushMessageList, nil
}

// QueryMessages replays messages by page and does NOT confirm them: call
// ConfirmMessages once they are safely stored. Until then they keep coming back,
// and confirming shifts the pagination under the next page.
func (c *Client) QueryMessages(ctx context.Context, q PushQuery) (PushMessagePage, error) {
	p := q.params()
	if q.Page > 0 {
		p["page"] = q.Page
	}
	if q.PageSize > 0 {
		p["pageSize"] = q.PageSize
	}
	body, err := c.Call(ctx, PushQueryList, p)
	if err != nil {
		return PushMessagePage{}, err
	}
	var out struct {
		PushMessagePage PushMessagePage `json:"pushMessagePage"`
	}
	if err := DecodeFlat(PushQueryList.Key(), body, &out); err != nil {
		return PushMessagePage{}, err
	}
	return out.PushMessagePage, nil
}

// ConfirmMessages marks messages as consumed. Only the query-style replay needs
// it; the cursor-style call confirms as it reads.
func (c *Client) ConfirmMessages(ctx context.Context, msgIDs IDs) (bool, error) {
	if len(msgIDs) == 0 {
		return true, nil
	}
	body, err := c.Call(ctx, PushConfirm, Params{"msgIdList": msgIDs})
	if err != nil {
		return false, err
	}
	var out struct {
		IsSuccess FlexBool `json:"isSuccess"`
	}
	if err := DecodeFlat(PushConfirm.Key(), body, &out); err != nil {
		return false, err
	}
	return out.IsSuccess.Bool(), nil
}
