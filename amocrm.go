package amocrm

import (
	"fmt"
	"github.com/whatcrm/go-amocrm/models"
)

func NewAPI(clientID, clientSecret, redirectURI string) AmoCRM {
	return &API{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURI:  redirectURI,
		Debug:        false,
	}
}

func (api *API) SetOptions(domain, accessToken string, debug bool) error {
	if isRegex(domain) {
		api.Domain = domain
	} else {
		return fmt.Errorf("domain name is not set")
	}

	if accessToken != "" {
		api.AccessToken = accessToken
	} else {
		return fmt.Errorf("accessToken is not set")
	}

	api.Debug = debug

	return nil
}

type AmoCRM interface {

	// SetOptions is main API settings
	SetOptions(domain, token string, debug bool) error
	TokenUPD(oAuth *OAuth) (*Tokens, error)
	GetAccount(params *Params) (acc Account, err error)

	//GetLeads - Leads interfaces
	GetLeads(id string, params *Params) ([]models.Lead, error)
	GetUnsorted(id string, params *Params) (unsorted []models.Unsorted, err error)
	CreateUnsortedSIP(in *[]models.UnsortedSIP) (out *models.MainResponse, err error)

	CreateLeads(lead *[]models.Lead, params *Params) (reps models.RequestResponse, err error)
	CreateLeadsComplex(lead *[]models.Lead) (resp []models.LeadComplexResponse, err error)
	ModifyLead(id string, lead *models.Lead, params *Params) (resp models.LeadModifyResponse, err error)

	// GetContacts - Contact interfaces
	GetContacts(contactID string, params *Params) (out []models.Contact, err error)
	CreateContact(in []models.Contact) (out models.RequestResponse, err error)
	ModifyContacts(contactID string, in []models.Contact) (out []models.Contact, err error)

	// GetContactChats -- FIXME
	GetContactChats(contactID, chatID string) (out models.RequestResponse, err error)
	ConnectChatToContact(in *[]models.Chat) (out models.RequestResponse, err error)

	// GetCompanies - Company interfaces
	GetCompanies(companyID string, params *Params) (out []models.Company, err error)
	CreateCompany(in *[]models.Company) (out []models.Company, err error)
	ModifyCompany(companyID string, in *[]models.Company) (out []models.Company, err error)

	//CreateTag - Tag interfaces
	CreateTag(entity string, tag *[]models.Tag) (out models.RequestResponse, err error)

	// CustomersMode - Customers interfaces
	CustomersMode(in models.CustomersMode) (out models.CustomersMode, err error)
	GetCustomers(customerID string, params *Params) (out []models.Customer, err error)
	CreateCustomer(in *[]models.Customer) (out models.RequestResponse, err error)
	ModifyCustomers(customerID string, in *[]models.Customer) (out models.CustomerResponse, err error)

	// GetTransactions - transactions interfaces
	GetTransactions(customerID, transactionID string, params *Params) (out []models.Transaction, err error)
	SetTransaction(customerID string, t *[]models.Transaction) (out models.MainResponse, err error)
	RemoveTransaction(customerID, transactionID string) (err error)

	SetBonusPoints(customerID string, value *models.BonusPoints) (out models.Points, err error)

	// GetCustomerStatuses - statuses interfaces
	GetCustomerStatuses(statusID string) (out []models.CustomerStatus, err error)
	CreateCustomerStatus(statuses []models.CustomerStatus) (out models.MainResponse, err error)
	ModifyCustomerStatus(statusID string, in *models.CustomerStatus) (out models.CustomerStatus, err error)
	RemoveCustomerStatus(statusID string) (err error)

	// GetCustomerSegments - segments interfaces
	GetCustomerSegments(segmentID string) (out []models.CustomerSegment, err error)
	CreateCustomerSegment(in *models.CustomerSegment) (err error)
	ModifyCustomerSegment(segmentID string, in *models.CustomerSegment) (out models.CustomerSegment, err error)
	RemoveCustomerSegment(segmentID string) (err error)

	// ReplyToWidgetRequest - Salesbot Reply to the widget_request handler
	ReplyToWidgetRequest(data, requestURL string) (err error)
	GetLinks(entity, entityID string, params *Params) (out models.RequestResponse, err error)

	log(message ...interface{})
}
