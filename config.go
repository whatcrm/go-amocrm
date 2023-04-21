package amocrm

const (
	accountURL     = "oauth2/account/subdomain"
	accessTokenURL = "oauth2/access_token"
	getAccountURL  = "api/v4/account"
	httpURL        = "https://"
	amocrmURL      = "www.amocrm.ru"
	noEntityURL    = "api/v4/"
)

// Leads Constants
const (
	leadURL         = "api/v4/leads"
	leadUnsorted    = "api/v4/leads/unsorted"
	leadComplexURL  = "api/v4/leads/complex"
	leadUnsortedSIP = "api/v4/leads/unsorted/sip"
)

// Contacts Constants
const (
	contactsURL     = "api/v4/contacts"
	contactsChatURL = "api/v4/contacts/chats"
)

// Company Constants
const (
	companiesURL = "api/v4/companies"
)

// EntityConstants
const (
	tagsURL  = "/tags"
	linksURL = "/links"
)

// Customers Constants
const (
	customersURL    = "api/v4/customers"
	customersMode   = "api/v4/customers/mode"
	transactionsURL = "/transactions"
	bonusPointsURL  = "/bonus_points"

	customerSegmentsURL = "api/v4/customers/segments"
	customerStatusesURL = "api/v4/customers/statuses"

	CustomerModeSegments    = "segments"
	CustomerModePeriodicity = "periodicity"
)

type makeRequestOptions struct {
	// Method is a request's method
	Method string
	// BaseURL is a url from constants above.
	BaseURL string
	// In is a struct, which will be marshalled to Request Body
	In interface{}
	// Out is a struct, which will be unmarshalled
	Out interface{}
	// Params is a URL Parameters
	Params *Params
}

type Params struct {
	With      With
	Page      string
	Limit     string
	Query     string
	Filter    string
	Order     string
	ContactID string
	ChatID    string
}

type With struct {
	CatalogElements bool
	Contacts        bool
	Leads           bool
	Companies       bool
}
