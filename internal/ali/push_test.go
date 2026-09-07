package ali_test

// Push topics against 1688-api-docs/messages-raw: every topic we subscribe to
// has a payload struct whose json tags match the documented message fields in
// both directions, its documented sample decodes into that struct, and every
// archived topic id is topicGroupName + "_" + topicName, which is the rule
// PushMessage.Topic relies on.

import (
	"encoding/json"
	"html"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"testing"

	"marketplace/internal/ali"
)

type msgNode struct {
	Name     string          `json:"name"`
	Desc     string          `json:"desc"`
	Type     string          `json:"type"`
	Sample   json.RawMessage `json:"sample"`
	Required bool            `json:"required"`
	Children []msgNode       `json:"children"`
}

type topicDoc struct {
	TopicID        string    `json:"topicId"`
	TopicName      string    `json:"topicName"`
	TopicGroupName string    `json:"topicGroupName"`
	Desc           string    `json:"desc"`
	MessageDocs    []msgNode `json:"messageDocs"`
	Sample         string    `json:"sample"`
}

func loadTopics(t *testing.T, dir string) map[string]*topicDoc {
	t.Helper()
	files, err := filepath.Glob(filepath.Join(dir, "messages-raw", "*.json"))
	if err != nil || len(files) == 0 {
		t.Skipf("no messages-raw corpus under %s", dir)
	}
	out := map[string]*topicDoc{}
	for _, f := range files {
		b, err := os.ReadFile(f)
		if err != nil {
			t.Fatalf("read %s: %v", f, err)
		}
		var d topicDoc
		if err := json.Unmarshal(b, &d); err != nil {
			t.Fatalf("parse %s: %v", f, err)
		}
		base := strings.TrimSuffix(filepath.Base(f), ".json")
		if d.TopicID != base {
			t.Errorf("%s: topicId %q does not match the file name", f, d.TopicID)
		}
		out[d.TopicID] = &d
	}
	return out
}

// msgPathSet flattens messageDocs into dotted paths. Object[] nodes are arrays
// of their children, which hang directly off the array's path. PascalCase
// wrapper keys such as OrderLogisticsTracingModel are ordinary path segments.
func msgPathSet(nodes []msgNode) map[string]bool {
	out := map[string]bool{}
	var walk func(nodes []msgNode, prefix string)
	walk = func(nodes []msgNode, prefix string) {
		for _, n := range nodes {
			p := prefix + n.Name
			out[p] = true
			if len(n.Children) > 0 {
				walk(n.Children, p+".")
			}
		}
	}
	walk(nodes, "")
	return out
}

func TestMessagePayloadsMatchTopics(t *testing.T) {
	dir := docsDir(t)
	topics := loadTopics(t, dir)

	// Every archived topic, subscribed or not: the id is group_name.
	for id, d := range topics {
		if want := ali.TopicType(d.TopicGroupName, d.TopicName); id != want {
			t.Errorf("%s: topicId != topicGroupName+\"_\"+topicName (%q + %q -> %q)", id, d.TopicGroupName, d.TopicName, want)
		}
	}

	payloads := map[string]reflect.Type{
		ali.TopicLogisticsTrace:  reflect.TypeOf(ali.LogisticsTracePush{}),
		ali.TopicMailNoChange:    reflect.TypeOf(ali.MailNoChangePush{}),
		ali.TopicInventoryChange: reflect.TypeOf(ali.InventoryChangePush{}),
		ali.TopicProductChange:   reflect.TypeOf(ali.ProductChange{}),
	}
	for _, topic := range ali.OrderTopics() {
		payloads[topic] = reflect.TypeOf(ali.OrderEvent{})
	}
	for _, topic := range ali.Topics() {
		if _, ok := payloads[topic]; !ok {
			t.Errorf("%s is in ali.Topics() but has no payload struct in this test", topic)
		}
	}

	var names []string
	for topic := range payloads {
		names = append(names, topic)
	}
	sort.Strings(names)
	for _, topic := range names {
		typ := payloads[topic]
		t.Run(topic, func(t *testing.T) {
			d, ok := topics[topic]
			if !ok {
				t.Fatalf("no messages-raw/%s.json in the archive", topic)
			}
			docSet := msgPathSet(d.MessageDocs)
			goSet := goPathSet(typ)

			var undocumented, unmodelled []string
			for p := range goSet {
				if !docSet[p] {
					undocumented = append(undocumented, p)
				}
			}
			for p := range docSet {
				if !goSet[p] {
					unmodelled = append(unmodelled, p)
				}
			}
			reportPaths(t, topic, "Go json tag not in the documented message fields (our typo?)", undocumented)
			reportPaths(t, topic, "documented message field not modelled by "+typ.String(), unmodelled)

			body := []byte(strings.TrimSpace(html.UnescapeString(d.Sample)))
			if !json.Valid(body) {
				t.Logf("%s: sample is not JSON; skipping the decode", topic)
				return
			}
			payload := reflect.New(typ).Interface()
			msg := ali.PushMessage{Type: topic, Data: body}
			if err := msg.Decode(payload); err != nil {
				t.Fatalf("%s: sample does not decode into %s: %v", topic, typ, err)
			}
			assertTopicSample(t, topic, payload)
		})
	}
}

// assertTopicSample pins one concrete documented value per topic.
func assertTopicSample(t *testing.T, topic string, payload any) {
	t.Helper()
	switch v := payload.(type) {
	case *ali.OrderEvent:
		if v.OrderID != 167539019420540000 {
			t.Errorf("%s: orderId = %d, want 167539019420540000", topic, v.OrderID)
		}
		if want := ali.TopicStatus(topic); v.CurrentStatus != want {
			t.Errorf("%s: sample currentStatus = %q, ali.TopicStatus says %q", topic, v.CurrentStatus, want)
		}
		if v.BuyerMemberID != "b2b-665170100" || v.SellerMemberID != "b2b-1676547900b7bb3" {
			t.Errorf("%s: buyer/seller = %q/%q", topic, v.BuyerMemberID, v.SellerMemberID)
		}
		if ts, err := v.MsgSendTime.Time(); err != nil || ts.Year() != 2018 || ts.Hour() != 19 {
			t.Errorf("%s: msgSendTime %q -> %s (%v), want 2018-05-30 19:xx Shanghai", topic, v.MsgSendTime, ts, err)
		}
	case *ali.LogisticsTracePush:
		if v.Model.StatusChanged != ali.TraceConsign {
			t.Errorf("%s: statusChanged = %q, want CONSIGN", topic, v.Model.StatusChanged)
		}
		if v.Model.LogisticsID != "12345" || v.Model.MailNo != "123456" {
			t.Errorf("%s: logisticsId/mailNo = %q/%q", topic, v.Model.LogisticsID, v.Model.MailNo)
		}
		if len(v.Model.OrderLogsItems) != 1 || v.Model.OrderLogsItems[0].OrderID != 2938624554509662976 || v.Model.OrderLogsItems[0].OrderEntryID != 2958624554509662976 {
			t.Errorf("%s: orderLogsItems = %+v (19-digit ids must survive as int64)", topic, v.Model.OrderLogsItems)
		}
	case *ali.MailNoChangePush:
		if v.Model.OldMailNo != "123" || v.Model.NewMailNo != "1234" || v.Model.NewCpCode != "12345" {
			t.Errorf("%s: %+v", topic, v.Model)
		}
		if len(v.Model.OrderLogsItems) != 1 || v.Model.OrderLogsItems[0].OrderID != 2958424554509662976 {
			t.Errorf("%s: orderLogsItems = %+v", topic, v.Model.OrderLogsItems)
		}
	case *ali.InventoryChangePush:
		if len(v.Changes) != 1 {
			t.Fatalf("%s: %d changes, want 1", topic, len(v.Changes))
		}
		c := v.Changes[0]
		if c.OfferID != 1234567890 || c.OfferOnSale != 100 || c.SkuOnSale != 20 {
			t.Errorf("%s: %+v", topic, c)
		}
		if c.Quantity != -10 {
			t.Errorf("%s: quantity = %d, want -10 (a signed delta)", topic, c.Quantity)
		}
		if c.BizTime != 1564984329147 {
			t.Errorf("%s: bizTime = %d, want 1564984329147 (quoted epoch millis in the sample)", topic, c.BizTime)
		}
		if got := c.Time().Year(); got != 2019 {
			t.Errorf("%s: bizTime year = %d, want 2019", topic, got)
		}
	case *ali.ProductChange:
		if v.ProductIDs != "897022666243,897022666242" {
			t.Errorf("%s: productIds = %q", topic, v.ProductIDs)
		}
		if ids := v.OfferIDs(); len(ids) != 2 || ids[0] != 897022666243 || ids[1] != 897022666242 {
			t.Errorf("%s: OfferIDs() = %v, want [897022666243 897022666242]", topic, ids)
		}
		if v.MemberID == "" || v.MsgSendTime.IsZero() {
			t.Errorf("%s: memberId=%q msgSendTime=%q", topic, v.MemberID, v.MsgSendTime)
		}
	default:
		t.Errorf("%s: no sample assertion for %T", topic, payload)
	}
}
