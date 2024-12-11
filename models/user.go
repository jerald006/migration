package models

type MongoUser struct {
	_id                       string   `json:"_id"`
	FirstName                 string   `json:"first_name"`
	LastName                  string   `json:"last_name"`
	Email                     string   `json:"email"`
	AgencyID                  string   `json:"agency_id"`
	Password                  string   `json:"password"`
	CustomerID                string   `json:"customer_id"`
	FinmontCustomerID         string   `json:"finmont_customer_id"`
	AirwallexCustomerID       string   `json:"airwallex_customer_id"`
	UserType                  string   `json:"user_type"`
	IsTravelAgent             bool     `json:"isTravelAgent"`
	IsSSOUser                 bool     `json:"isSSOUser"`
	IsRestrictedAgent         bool     `json:"isRestrictedAgent"`
	PhoneNumber               string   `json:"phoneNumber"`
	Status                    bool     `json:"status"`
	VerificationStatus        bool     `json:"verificationStatus"`
	IsCreditCard              bool     `json:"isCreditCard"`
	ProfilePhoto              string   `json:"profilePhoto"`
	FileName                  string   `json:"fileName"`
	LoginCount                int      `json:"loginCount"`
	IsRewardWalletEnable      bool     `json:"isRewardWalletEnable"`
	IsSystemGeneratedPassword bool     `json:"isSystemGeneratedPassword"`
	SignedUpThrough           string   `json:"signed_up_through"`
	FacebookID                string   `json:"facebook_id"`
	ClaroUserName             string   `json:"claroUserName"`
	GoogleID                  string   `json:"google_id"`
	IsDealRayUser             bool     `json:"isDealRayUser"`
	ActiveAgentName           string   `json:"activeAgentName"`
	AgentAssociation          []string `json:"agentAssociation"`
	AssociationActionHistory  []string `json:"associationActionHistory"`
	AffiliateLink             string   `json:"affiliateLink"`
	IsSmsCampaignAllowed      bool     `json:"isSmsCampaignAllowed"`
	ZohoCRMId                 string   `json:"zohoCRMId"`
}

type CockroachDBUser struct {
	_id string `json:"_id"`
}
