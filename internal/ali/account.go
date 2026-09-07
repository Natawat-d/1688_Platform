package ali

import "context"

// SimpleAccountInfo is alibaba.account.simpleAccountInfo, the identity behind
// the access token. Only a handful of its forty-odd documented fields are ever
// populated in practice, so everything is optional.
//
// This is the cheapest call in the whole set and needs no application
// parameters, which makes it the natural smoke test for credentials, signing and
// transport all at once.
type SimpleAccountInfo struct {
	LoginID           string    `json:"loginId"`
	MemberID          string    `json:"memberId"`
	UserID            ID        `json:"userId"`
	OpenUID           string    `json:"openUid"`
	CompanyName       string    `json:"companyName"`
	SupplierName      string    `json:"supplierName"`
	SellerName        string    `json:"sellerName"`
	CategoryID        int       `json:"categoryId"`
	CategoryName      string    `json:"categoryName"`
	MemberBizType     string    `json:"memberBizType"`
	Status            string    `json:"status"`
	EnterpriseAccount FlexBool  `json:"enterpriseAccount"`
	PersonAccount     FlexBool  `json:"personAccount"`
	CrossBorder       FlexBool  `json:"crossBorder"`
	KuaJingBao        FlexBool  `json:"kuaJingBao"`
	Email             string    `json:"email"`
	PhoneNo           string    `json:"phoneNo"`
	MobileNo          string    `json:"mobileNo"`
	HomepageURL       string    `json:"homepageUrl"`
	ShopURL           string    `json:"shopUrl"`
	Icon              string    `json:"icon"`
	TrustScore        int       `json:"trustScore"`
	TpYear            int       `json:"tpYear"`
	CreateDate        Timestamp `json:"createDate"`
	ModifyDate        Timestamp `json:"modifyDate"`
	GmtPaidJoin       Timestamp `json:"gmtPaidJoin"`
	JoinFrom          string    `json:"joinFrom"`
	SaleKeywords      string    `json:"saleKeywords"`
	BuyKeywords       string    `json:"buyKeywords"`
	// IsOfficalLogistics keeps the documented spelling, missing L and all.
	IsOfficalLogistics FlexBool `json:"isOfficalLogistics"`
}

// AccountBasic returns the account the access token belongs to.
func (c *Client) AccountBasic(ctx context.Context) (SimpleAccountInfo, error) {
	body, err := c.Call(ctx, AccountBasic, Params{})
	if err != nil {
		return SimpleAccountInfo{}, err
	}
	return DecodeB[SimpleAccountInfo](AccountBasic.Key(), body)
}
