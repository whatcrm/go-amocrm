package amocrm

import (
	"github.com/gofiber/fiber/v2"
	"github.com/whatcrm/go-amocrm/models"
)

func (api *API) GetContacts(contactID string, params *Params) (out []models.Contact, err error) {
	api.log("GetContacts request is started...")

	options := makeRequestOptions{
		Method:  fiber.MethodGet,
		BaseURL: contactsURL,
		In:      nil,
		Out:     &models.Contact{},
		Params:  params,
	}

	if contactID != "" {
		options.BaseURL += "/" + contactID
		if err = api.makeRequest(options); err != nil {
			return
		}
		out = []models.Contact{*options.Out.(*models.Contact)}
	}

	if contactID == "" {
		options.Out = &models.RequestResponse{}
		if err = api.makeRequest(options); err != nil {
			return
		}
		out = options.Out.(*models.RequestResponse).Embedded.Contacts
	}

	api.log("returning the struct...")
	return
}

func (api *API) CreateContact(in []models.Contact) (out models.RequestResponse, err error) {
	api.log("CreateContact request is started...")

	options := makeRequestOptions{
		Method:  fiber.MethodPost,
		BaseURL: contactsURL,
		In:      in,
		Out:     &out,
		Params:  nil,
	}

	if err = api.makeRequest(options); err != nil {
		return
	}

	api.log("returning the struct...")
	return
}

func (api *API) ModifyContacts(contactID string, in []models.Contact) (out []models.Contact, err error) {
	api.log("ModifyContacts request is started...")

	options := makeRequestOptions{
		Method:  fiber.MethodPatch,
		BaseURL: contactsURL,
		In:      in,
		Out:     &models.Contact{},
		Params:  nil,
	}

	if contactID != "" {
		options.BaseURL += "/" + contactID
		options.In = in[0]
		if err = api.makeRequest(options); err != nil {
			return
		}
		out = []models.Contact{*options.Out.(*models.Contact)}
	}

	if contactID == "" {
		options.Out = &models.RequestResponse{}
		if err = api.makeRequest(options); err != nil {
			return
		}
		out = options.Out.(*models.RequestResponse).Embedded.Contacts
	}

	api.log("returning the struct...")
	return
}

func (api *API) GetContactChats(contactID, chatID string) (out models.RequestResponse, err error) {
	api.log("GetContactChats request is started...")

	p := &Params{
		ContactID: contactID,
		ChatID:    chatID,
	}

	options := makeRequestOptions{
		Method:  fiber.MethodPatch,
		BaseURL: contactsChatURL,
		In:      nil,
		Out:     &out,
		Params:  p,
	}

	api.getAgent(options.Method, options.BaseURL, options.Params)
	//req.RequestURI()

	api.log("returning the struct...")
	return
}

func (api *API) ConnectChatToContact(in *[]models.Chat) (out models.RequestResponse, err error) {
	api.log("ConnectChatToContact request is started...")

	options := makeRequestOptions{
		Method:  fiber.MethodPatch,
		BaseURL: contactsURL,
		In:      in,
		Out:     &out,
		Params:  nil,
	}
	if err = api.makeRequest(options); err != nil {
		return
	}

	api.log("returning the struct...")
	return
}
