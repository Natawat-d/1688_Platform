package pricing

import (
	"strconv"
	"strings"
	"time"

	"marketplace/internal/ali"
)

// Fee rule scopes, narrowest last. The strings are what the fee_rules table
// stores, so they are matched case-insensitively and nothing else matches.
const (
	ScopeGlobal   = "global"
	ScopeCategory = "category"
	ScopeSupplier = "supplier"
	ScopeProduct  = "product"
)

// FeeRule is one row of the fee_rules table: what the platform adds on top of
// the landed cost of a product.
//
// FeeBps is basis points of the per-unit subtotal (1000 = 10%), FeeFixedFen is
// added per unit after that, and MinFeeFen is a per-unit floor. The window is
// half open: a rule applies from From inclusive until To exclusive, and a nil
// To means it never expires.
type FeeRule struct {
	ID          int64
	Scope       string // ScopeGlobal, ScopeCategory, ScopeSupplier, ScopeProduct
	ScopeValue  string // category id / sellerOpenId / offer id, ids in decimal string form
	FeeBps      int
	FeeFixedFen ali.Fen
	MinFeeFen   ali.Fen
	Priority    int
	From        time.Time
	To          *time.Time
}

// scopeRank is the specificity of a scope, used as the tie break after
// Priority. A rule written for one product beats one written for its supplier,
// which beats one for its category, which beats the house default.
func scopeRank(scope string) int {
	switch strings.ToLower(strings.TrimSpace(scope)) {
	case ScopeProduct:
		return 3
	case ScopeSupplier:
		return 2
	case ScopeCategory:
		return 1
	case ScopeGlobal:
		return 0
	}
	return -1 // unknown scope: never matches
}

// active reports whether now falls inside the rule's [From, To) window.
func (r FeeRule) active(now time.Time) bool {
	if now.Before(r.From) {
		return false
	}
	return r.To == nil || now.Before(*r.To)
}

// matches reports whether the rule's scope covers p.
func (r FeeRule) matches(p Product) bool {
	switch scopeRank(r.Scope) {
	case 3:
		return r.ScopeValue == strconv.FormatInt(p.OfferID, 10)
	case 2:
		return r.ScopeValue == p.SellerOpenID
	case 1:
		return r.ScopeValue == strconv.FormatInt(p.CategoryID, 10)
	case 0:
		return true
	}
	return false
}

// SelectFeeRule picks the fee rule that applies to p at now, or nil when none
// does. Candidates are the rules whose window contains now and whose scope
// covers p; the winner is the highest Priority, then the most specific scope
// (product > supplier > category > global), then the highest ID, so that the
// most recently created of two otherwise identical rules wins.
//
// The rules slice is not modified and the choice does not depend on its order.
func SelectFeeRule(rules []FeeRule, p Product, now time.Time) *FeeRule {
	best := -1
	for i := range rules {
		r := rules[i]
		if scopeRank(r.Scope) < 0 || !r.active(now) || !r.matches(p) {
			continue
		}
		if best < 0 || better(r, rules[best]) {
			best = i
		}
	}
	if best < 0 {
		return nil
	}
	out := rules[best]
	return &out
}

// better reports whether a outranks b.
func better(a, b FeeRule) bool {
	if a.Priority != b.Priority {
		return a.Priority > b.Priority
	}
	if ra, rb := scopeRank(a.Scope), scopeRank(b.Scope); ra != rb {
		return ra > rb
	}
	return a.ID > b.ID
}

// Rules adapts a rule set to the callback Quote wants, evaluating the window
// once at now. Snapshot the result per order so a rule edited mid-checkout
// cannot change a price the shopper has already been shown.
func Rules(rules []FeeRule, now time.Time) func(Product) *FeeRule {
	return func(p Product) *FeeRule { return SelectFeeRule(rules, p, now) }
}
