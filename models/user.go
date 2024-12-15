package models

import (
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoUser struct {
	Id                        string             `json:"_id"`
	FirstName                 string             `json:"first_name"`
	LastName                  string             `json:"last_name"`
	Email                     string             `json:"email"`
	AgencyID                  primitive.ObjectID `bson:"agency_id"`
	Password                  string             `json:"password"`
	CustomerID                string             `json:"customer_id"`
	FinmontCustomerID         string             `json:"finmont_customer_id"`
	AirwallexCustomerID       string             `json:"airwallex_customer_id"`
	UserType                  string             `json:"user_type"`
	IsTravelAgent             bool               `json:"isTravelAgent"`
	IsSSOUser                 bool               `json:"isSSOUser"`
	IsRestrictedAgent         bool               `json:"isRestrictedAgent"`
	PhoneNumber               string             `json:"phoneNumber"`
	Status                    bool               `json:"status"`
	VerificationStatus        bool               `json:"verificationStatus"`
	IsCreditCard              bool               `json:"isCreditCard"`
	ProfilePhoto              string             `json:"profilePhoto"`
	FileName                  string             `json:"fileName"`
	LoginCount                int                `json:"loginCount"`
	IsRewardWalletEnable      bool               `json:"isRewardWalletEnable"`
	IsSystemGeneratedPassword bool               `json:"isSystemGeneratedPassword"`
	SignedUpThrough           string             `json:"signed_up_through"`
	FacebookID                string             `json:"facebook_id"`
	ClaroUserName             string             `json:"claroUserName"`
	GoogleID                  string             `json:"google_id"`
	IsDealRayUser             bool               `json:"isDealRayUser"`
	ActiveAgentName           string             `json:"activeAgentName"`
	AgentAssociation          []string           `json:"agentAssociation"`
	AssociationActionHistory  []string           `json:"associationActionHistory"`
	AffiliateLink             string             `json:"affiliateLink"`
	IsSmsCampaignAllowed      bool               `json:"isSmsCampaignAllowed"`
	ZohoCRMId                 string             `json:"zohoCRMId"`
}

type CockroachDBUser struct {
	Id                        string    `json:"id"`
	FirstName                 string    `json:"first_name"`
	LastName                  string    `json:"last_name"`
	Email                     string    `json:"email"`
	AgencyID                  *string   `json:"agency_id" db:"agency_id"`
	Password                  *string   `json:"password"`
	UserType                  *string   `json:"user_type"`
	IsSSOUser                 bool      `json:"is_SSO_user" gorm:"column:is_SSO_user"`
	IsRestrictedAgent         bool      `json:"is_restricted_agent"`
	PhoneNumber               *string   `json:"phone_number"`
	Status                    bool      `json:"status"`
	VerificationStatus        bool      `json:"verification_status"`
	ProfilePhoto              *string   `json:"profile_photo"`
	LoginCount                int64     `json:"login_count"`
	IsRewardWalletEnable      bool      `json:"is_reward_wallet_enable"`
	IsSystemGeneratedPassword bool      `json:"is_system_generated_password"`
	SignedUpThrough           *string   `json:"signed_up_through"`
	AffiliateLink             *string   `json:"affiliate_link"`
	CreatedAt                 time.Time `json:"created_at"`
	UpdatedAt                 time.Time `json:"updated_at"`
}

func (CockroachDBUser) TableName() string {
	return "users"
}

func (mongoUser MongoUser) ConvertMongoToCockroachUser() CockroachDBUser {
	idHex := mongoUser.AgencyID.Hex()
	return CockroachDBUser{
		Id:                        mongoUser.Id,
		FirstName:                 mongoUser.FirstName,
		LastName:                  mongoUser.LastName,
		Email:                     mongoUser.Email,
		AgencyID:                  &idHex,
		Password:                  &mongoUser.Password,
		UserType:                  getValidUserType(mongoUser.UserType),
		IsSSOUser:                 mongoUser.IsSSOUser,
		IsRestrictedAgent:         mongoUser.IsRestrictedAgent,
		PhoneNumber:               &mongoUser.PhoneNumber,
		Status:                    mongoUser.Status,
		VerificationStatus:        mongoUser.VerificationStatus,
		ProfilePhoto:              &mongoUser.ProfilePhoto,
		LoginCount:                int64(mongoUser.LoginCount),
		IsRewardWalletEnable:      mongoUser.IsRewardWalletEnable,
		IsSystemGeneratedPassword: mongoUser.IsSystemGeneratedPassword,
		SignedUpThrough:           getSignedUpThrough(mongoUser.SignedUpThrough),
		AffiliateLink:             &mongoUser.AffiliateLink,
		CreatedAt:                 time.Now(),
		UpdatedAt:                 time.Now(),
	}
}

func getValidUserType(userType string) *string {
	validUserTypes := map[string]struct{}{
		"user":         {},
		"agency_owner": {},
		"sales_agent":  {},
		"admin":        {},
	}

	// Convert user type to lowercase for comparison
	lowerUserType := strings.ToLower(userType)
	if _, exists := validUserTypes[lowerUserType]; exists {
		return &lowerUserType
	}
	defaultUserType := "user"
	return &defaultUserType
}

func getSignedUpThrough(signedUpThrough string) *string {
	validSignedUpThrough := map[string]struct{}{
		"agency":   {},
		"google":   {},
		"facebook": {},
		"claro":    {},
	}

	// Convert user type to lowercase for comparison
	lowerSignedUpThrough := strings.ToLower(signedUpThrough)
	if _, exists := validSignedUpThrough[lowerSignedUpThrough]; exists {
		return &lowerSignedUpThrough
	}
	defaultSignedUpThrough := "agency"
	return &defaultSignedUpThrough
}
