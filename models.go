package amocrm

// Handler with main data to work with amoCRM

type API struct {
	ClientID     string
	ClientSecret string
	AccessToken  string
	Domain       string
	RedirectURI  string
	Debug        bool
}

// Struct to send a data to amoCRM in order to get the tokens

type OAuth struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
	RefreshToken string `json:"refresh_token,omitempty"`
	Code         string `json:"code,omitempty"`
	RedirectURI  string `json:"redirect_uri"`
}

// The tokens' struct

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
}

type OauthAccount struct {
	ID             int    `json:"id"`
	Subdomain      string `json:"subdomain"`
	Domain         string `json:"domain"`
	TopLevelDomain string `json:"top_level_domain"`
}

// Account represents amoCRM Account entity json DTO.
type Account struct {
	ID                      int    `json:"id"`
	Name                    string `json:"name"`
	Subdomain               string `json:"subdomain"`
	CreatedAt               int    `json:"created_at"`
	CreatedBy               int    `json:"created_by"`
	UpdatedAt               int    `json:"updated_at"`
	UpdatedBy               int    `json:"updated_by"`
	CurrentUserID           int    `json:"current_user_id"`
	Country                 string `json:"country"`
	Currency                string `json:"currency"`
	CurrencySymbol          string `json:"currency_symbol"`
	CustomersMode           string `json:"customers_mode"`
	IsUnsortedOn            bool   `json:"is_unsorted_on"`
	MobileFeatureVersion    int    `json:"mobile_feature_version"`
	IsLossReasonEnabled     bool   `json:"is_loss_reason_enabled"`
	IsHelpbotEnabled        bool   `json:"is_helpbot_enabled"`
	IsTechnicalAccount      bool   `json:"is_technical_account"`
	ContactNameDisplayOrder int    `json:"contact_name_display_order"`
	Links                   struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
}

type TestRequest struct {
	ID             int    `json:"id"`
	Subdomain      string `json:"subdomain"`
	Domain         string `json:"domain"`
	TopLevelDomain string `json:"top_level_domain"`
}
