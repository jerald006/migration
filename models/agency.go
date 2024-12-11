package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoAgency struct {
	Id                       primitive.ObjectID `bson:"_id"`
	Owner                    primitive.ObjectID `json:"owner,omitempty"`
	Subdomain                string             `json:"subdomain" validate:"required,subdomain"`
	AgencyName               string             `json:"agencyName" validate:"required"`
	Domain                   string             `json:"domain,omitempty" validate:"domain"`
	Logo                     string             `json:"logo,omitempty"`
	LogoMobile               string             `json:"logoMobile,omitempty"`
	Favicon                  string             `json:"favicon,omitempty"`
	EINNumber                string             `json:"EINnumber,omitempty"`
	CurrentWebsite           string             `json:"currentWebsite,omitempty"`
	GovernmentID             string             `json:"governmentId,omitempty"`
	PrimaryColor             string             `json:"primaryColor,omitempty" default:"#032C4C"`
	SecondaryColor           string             `json:"secondaryColor,omitempty" default:"#38A896"`
	PrimaryTextColor         string             `json:"primaryTextColor,omitempty" default:"#f2f2f2"`
	SecondaryTextColor       string             `json:"secondaryTextColor,omitempty" default:"#f2f2f2"`
	FooterColor              string             `json:"footerColor,omitempty" default:"#0e5191"`
	FooterTextColor          string             `json:"footerTextColor,omitempty" default:"#ffffff"`
	BackgroundImage          string             `json:"backgroundImage,omitempty"`
	MobileBackgroundImage    string             `json:"mobileBackgroundImage,omitempty"`
	BackgroundOverlay        float64            `json:"backgroundOverlay,omitempty" default:"0.0"`
	Heading                  string             `json:"heading,omitempty"`
	SubHeading               string             `json:"subHeading,omitempty"`
	HeadingColor             string             `json:"headingColor,omitempty" default:"#fff"`
	SubHeadingColor          string             `json:"subHeadingColor,omitempty" default:"#fff"`
	SocialMedia              SocialMedia        `json:"social,omitempty"`
	ContactPhoneNumber       string             `json:"contactPhoneNumber,omitempty"`
	ContactEmailID           string             `json:"contactEmailId" validate:"email,required"`
	PrivacyPolicy            string             `json:"privacyPolicy,omitempty"`
	CookiesPolicy            string             `json:"cookiesPolicy,omitempty"`
	Disclaimer               string             `json:"disclaimer,omitempty"`
	WhitelistedSignupDomains []string           `json:"whitelistedSignupDomains,omitempty"`
	AddressLine1             string             `json:"addressLine1,omitempty"`
	AddressLine2             string             `json:"addressLine2,omitempty"`
	State                    string             `json:"state,omitempty"`
	City                     string             `json:"city,omitempty"`
	OriginCountry            string             `json:"originCountry,omitempty"`
	OriginCountryCode        string             `json:"originCountryCode,omitempty"`
	Country                  string             `json:"country,omitempty"`
	ZipCode                  string             `json:"zipCode,omitempty"`
	IATA                     string             `json:"iata,omitempty"`
	ARC                      string             `json:"arc,omitempty"`
	Status                   string             `json:"status,omitempty" default:"Active"`
	PaymentStatus            string             `json:"paymentStatus,omitempty" default:"Inactive"`
	IsSubscribed             bool               `json:"isSubscribed,omitempty"`
	TermsAcceptedAt          *time.Time         `json:"xeniTermsAndConditionsAcceptedAt,omitempty"`
	BookingOptions           BookingOptions     `json:"bookingMenuItemOptions,omitempty"`
	APIAccess                APIAccess          `json:"apiAccess,omitempty"`
	Analytics                Analytics          `json:"analytics,omitempty"`
	CurrentPaymentGateway    string             `json:"currentPaymentGateway,omitempty"`
	ManageLocation           bool               `json:"manageLocation,omitempty"`
	DefaultDealsLocation     LocationSettings   `json:"defaultDealsLocation,omitempty"`
	DefaultTargetLocation    LocationSettings   `json:"defaultTargetLocation,omitempty"`
	MetaData                 MetaData           `json:"metaData,omitempty"`
	OTPEnabled               bool               `json:"otpSignupEnable,omitempty"`
	HideHeaderForProposal    bool               `json:"hideHeaderforDealProposal,omitempty"`
	ExtranetEnabled          bool               `json:"extranetEnabled,omitempty"`
}

// Nested Structures
type SocialMedia struct {
	Facebook  string `json:"facebook,omitempty"`
	Tiktok    string `json:"tiktok,omitempty"`
	Twitter   string `json:"twitter,omitempty"`
	Instagram string `json:"instagram,omitempty"`
	Linkedin  string `json:"linkedin,omitempty"`
	Youtube   string `json:"youtube,omitempty"`
	Discord   string `json:"discord,omitempty"`
}

type BookingOptions struct {
	Title  string `json:"title"`
	Active bool   `json:"active"`
}

type APIAccess struct {
	IsEnabled bool   `json:"isEnabled"`
	Value     string `json:"value,omitempty"`
}

type Analytics struct {
	IsEnabled           bool   `json:"isEnabled"`
	FacebookAnalyticsID string `json:"facebookAnalyticsId,omitempty"`
	GoogleAnalyticsID   string `json:"googleAnalyticsId,omitempty"`
}

type LocationSettings struct {
	Hotels     bool `json:"hotels"`
	Cars       bool `json:"cars"`
	Activities bool `json:"activities"`
}

type MetaData struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Keywords    []string `json:"keywords"`
}

type CockroachDBAgency struct {
	ID                 string     `json:"id" db:"id"`
	OwnerID            *string    `json:"owner_id" db:"owner_id"`
	Subdomain          string     `json:"subdomain" db:"subdomain"`
	AgencyName         string     `json:"agency_name" db:"agency_name"`
	Domain             *string    `json:"domain" db:"domain"`
	EINNumber          *string    `json:"EIN_number" db:"EIN_number"`
	CurrentWebsite     *string    `json:"current_website" db:"current_website"`
	GovernmentID       *string    `json:"government_id" db:"government_id"`
	ContactPhoneNumber *string    `json:"contact_phone_number" db:"contact_phone_number"`
	ContactEmailID     string     `json:"contact_email_id" db:"contact_email_id"`
	CreatedAt          *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt          *time.Time `json:"updated_at" db:"updated_at"`
}

type CockroachDBAgencyPolicy struct {
	ID                              string     `json:"id" db:"id"`
	AgencyID                        *string    `json:"agency_id" db:"agency_id"`
	SubscriptionID                  *string    `json:"subscription_id" db:"subscription_id"`
	PlanPrice                       *float64   `json:"plan_price" db:"plan_price"`
	PlanType                        *string    `json:"plan_type" db:"plan_type"`
	PaymentID                       *string    `json:"payment_id" db:"payment_id"`
	CancelledAt                     *time.Time `json:"cancelled_at" db:"cancelled_at"`
	CancellationReason              *string    `json:"cancellation_reason" db:"cancellation_reason"`
	Duration                        *string    `json:"duration" db:"duration"`
	IsCancellationRequested         *bool      `json:"is_cancellation_requested" db:"is_cancellation_requested"`
	CouponCode                      *string    `json:"coupon_code" db:"coupon_code"`
	PrivacyPolicy                   string     `json:"privacy_policy" db:"privacy_policy"`
	CookiesPolicy                   string     `json:"cookies_policy" db:"cookies_policy"`
	Disclaimer                      string     `json:"disclaimer" db:"disclaimer"`
	IATA                            *string    `json:"IATA" db:"IATA"`
	ARC                             *string    `json:"ARC" db:"ARC"`
	Status                          *string    `json:"status" db:"status"`
	PaymentStatus                   *string    `json:"payment_status" db:"payment_status"`
	IsSubscribed                    bool       `json:"is_subscribed" db:"is_subscribed"`
	XeniTncAcceptedAt               *time.Time `json:"xeni_tnc_accepted_at" db:"xeni_tnc_accepted_at"`
	AllowSameDayBooking             *string    `json:"allow_same_day_booking" db:"allow_same_day_booking"`
	ShowPublishedPriceToCustomer    *string    `json:"show_published_price_to_customer" db:"show_published_price_to_customer"`
	BookingMenuItemOptions          *string    `json:"booking_menu_item_options" db:"booking_menu_item_options"`
	XenipayAPIKey                   *string    `json:"xenipay_api_key" db:"xenipay_api_key"`
	XenipaySecretKey                *string    `json:"xenipay_secret_key" db:"xenipay_secret_key"`
	XenipayAccountID                *string    `json:"xenipay_account_id" db:"xenipay_account_id"`
	HeardAboutFrom                  *string    `json:"heard_about_from" db:"heard_about_from"`
	HearAboutFromName               *string    `json:"hear_about_from_name" db:"hear_about_from_name"`
	IsSSOEnabled                    *bool      `json:"is_SSO_enabled" db:"is_SSO_enabled"`
	TokenValidateURL                *string    `json:"token_validate_URL" db:"token_validate_URL"`
	TokenKey                        *string    `json:"token_key" db:"token_key"`
	IsDecoupledSignupEnabled        *bool      `json:"is_decoupled_signup_enabled" db:"is_decoupled_signup_enabled"`
	IsDirectLoginAllowedForSSOUsers *bool      `json:"is_direct_login_allowed_for_SSO_users" db:"is_direct_login_allowed_for_SSO_users"`
	IsHide                          bool       `json:"is_hide" db:"is_hide"`
	IsAPIAccessEnabled              *bool      `json:"is_api_access_enabled" db:"is_api_access_enabled"`
	APIKey                          *string    `json:"api_key" db:"api_key"`
	IsFlightTicketingAllowed        bool       `json:"is_flight_ticketing_allowed" db:"is_flight_ticketing_allowed"`
	SalesAgentsLimit                *int64     `json:"sales_agents_limit" db:"sales_agents_limit"`
	CustomersLimit                  *int64     `json:"customers_limit" db:"customers_limit"`
	CurrentPaymentGateway           *string    `json:"current_payment_gateway" db:"current_payment_gateway"`
	ManageLocation                  bool       `json:"manage_location" db:"manage_location"`
	DefaultDealsLocation            *string    `json:"default_deals_location" db:"default_deals_location"`
	MailgunEmailSentDate            *time.Time `json:"mailgun_email_sent_date" db:"mailgun_email_sent_date"`
	UploadIDEmailSentDate           *time.Time `json:"upload_ID_email_sent_date" db:"upload_ID_email_sent_date"`
	OTPSignupEnable                 *bool      `json:"OTP_signup_enable" db:"OTP_signup_enable"`
	HideHeaderForDealProposal       *bool      `json:"hide_header_for_deal_proposal" db:"hide_header_for_deal_proposal"`
	ExtranetEnabled                 *bool      `json:"extranet_enabled" db:"extranet_enabled"`
	Metadata                        *string    `json:"metadata" db:"metadata"`
	CurrentHotelVendor              *string    `json:"current_hotel_vendor" db:"current_hotel_vendor"`
	HQCountry                       *string    `json:"hq_country" db:"hq_country"`
	PreferredLanguageForBusiness    *string    `json:"preferred_language_for_business" db:"preferred_language_for_business"`
	BookingVolume                   *string    `json:"booking_volume" db:"booking_volume"`
	UseForBusiness                  *string    `json:"use_for_business" db:"use_for_business"`
	ProductInterested               *string    `json:"product_interested" db:"product_interested"`
	ThreeDSEnable                   *bool      `json:"three_ds_enable" db:"three_ds_enable"`
	IsHideDealOrProposalOption      *bool      `json:"is_hide_deal_or_proposal_option" db:"is_hide_deal_or_proposal_option"`
	CreatedAt                       *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt                       *time.Time `json:"updated_at" db:"updated_at"`
	ExpiredAt                       *time.Time `json:"expired_at" db:"expired_at"`
	ReturnURL                       *string    `json:"return_url" db:"return_url"`
	DefaultTargetLocation           *string    `json:"default_target_location" db:"default_target_location"`
}
