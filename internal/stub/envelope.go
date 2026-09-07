package stub

import (
	"encoding/json"

	"marketplace/internal/ali"
)

// EVERY response in this package is built as map[string]any with the field names
// written out as literal strings copied from the documentation. No response
// struct from internal/ali is imported, and none ever should be: if both sides
// shared a struct then a misspelling like orderPreviewResuslt would be
// symmetric and therefore invisible. Catching exactly that is why the stub is a
// separate program speaking the real protocol.

// famA builds the Family A envelope:
//
//	{"result":{"success":true,"code":"200","message":"成功","result":<payload>}}
//
// Used by keywordQuery, queryProductDetail, category.translation.getById and
// product.freight.estimate.
func famA(payload any) map[string]any {
	return map[string]any{
		"result": map[string]any{
			"success": true,
			"code":    "200",
			"message": "成功",
			"result":  payload,
		},
	}
}

// famAErr is Family A carrying a refusal instead of a payload.
func famAErr(code, message string) map[string]any {
	return map[string]any{
		"result": map[string]any{
			"success": false,
			"code":    code,
			"message": message,
			"result":  nil,
		},
	}
}

// famASN is the keywordSNQuery variant of Family A: the same shape, but the
// status fields are spelled retCode and retMsg. That is not a typo on our side —
// product.search.keywordSNQuery really does document it that way while its
// siblings use code and message.
func famASN(payload any) map[string]any {
	return map[string]any{
		"result": map[string]any{
			"success": true,
			"retCode": "S0000",
			"retMsg":  "成功",
			"result":  payload,
		},
	}
}

// famB builds the Family B envelope:
//
//	{"result":<payload>,"errorCode":"","errorMessage":"","success":true}
//
// Used by alibaba.account.basic.
func famB(payload any) map[string]any {
	return map[string]any{
		"result":       payload,
		"errorCode":    "",
		"errorMessage": "",
		"success":      true,
	}
}

// famBErr is Family B carrying a refusal.
func famBErr(code, message string) map[string]any {
	return map[string]any{
		"errorCode":    code,
		"errorMessage": message,
		"success":      false,
	}
}

// famBStr is Family B with success as a STRING. alibaba.trade.get.buyerView
// documents success as String while every sibling documents it as Boolean, and
// a client that assumes Boolean everywhere will fail here. On purpose.
func famBStr(payload any) map[string]any {
	return map[string]any{
		"result":       payload,
		"errorCode":    "",
		"errorMessage": "",
		"success":      "true",
	}
}

// famBStrErr is famBStr carrying a refusal, success still a string.
func famBStrErr(code, message string) map[string]any {
	return map[string]any{
		"errorCode":    code,
		"errorMessage": message,
		"success":      "false",
	}
}

// yuan renders a Fen amount as a BARE JSON DECIMAL, e.g. 6.15, which is how the
// order-reading family sends BigDecimal amounts. json.RawMessage is the only way
// to emit a decimal number from Go without a float64 anywhere near the money.
func yuan(f ali.Fen) json.RawMessage { return json.RawMessage(f.String()) }

// price renders a Fen amount as the decimal STRING the catalogue family uses,
// e.g. "18.50".
func price(f ali.Fen) string { return f.String() }

// rawNumber emits an already-formatted decimal as a bare JSON number. Same
// reason as yuan: weights and dimensions are documented as Double, and building
// them out of integers keeps float64 out of the codebase entirely.
func rawNumber(s string) json.RawMessage { return json.RawMessage(s) }
