package amocrm

const (
	testURL         = "oauth2/account/subdomain"
	accessTokenURL  = "oauth2/access_token"
	getAccountURL   = "api/v4/account"
	httpURL         = "https://"
	amocrmURL       = "www.amocrm.ru"
	noEntityURL     = "api/v4/"
	leadURL         = "api/v4/leads"
	leadUnsorted    = "api/v4/leads/unsorted"
	leadComplexURL  = "api/v4/leads/complex"
	leadUnsortedSIP = "api/v4/leads/unsorted/sip"
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
	// Params is an URL Parameters
	Params *Params
}

type Params struct {
	// The struct contains parameters for requests params, i.e. https://example.com/api/v4/leads?query=
	With   string
	Page   string
	Limit  string
	Query  string
	Filter string
	Order  string
}
