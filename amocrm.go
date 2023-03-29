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
	CreateUnsortedSIP(in *[]models.UnsortedSIP) (out *models.UnsortedSIPResponse, err error)

	CreateLeads(lead *[]models.Lead, params *Params) (reps models.RequestResponse, err error)
	CreateLeadsComplex(lead *[]models.Lead) (resp []models.LeadComplexResponse, err error)
	ModifyLead(id string, lead *models.Lead, params *Params) (resp models.LeadModifyResponse, err error)

	//CreateTag - Tag interfaces

	//CreateTag(entity string, tag *[]models.Tag) (*models.TagResponse, error)

	log(message ...string)
}
