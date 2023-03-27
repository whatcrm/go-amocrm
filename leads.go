package amocrm

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"whatcrm_notify/pkg/amocrm/models"
)

// TODO Lead Struct
// TODO GET Parameters in Requests - https://www.amocrm.ru/developers/content/crm_platform/leads-api#leads-list

func (api *API) CreateUnsortedSIP(in *[]models.UnsortedSIP) (out *models.UnsortedSIPResponse, err error) {
	api.log("CreateUnsortedSIP request is started...")

	options := makeRequestOptions{
		Method:  fiber.MethodPost,
		BaseURL: leadUnsortedSIP,
		In:      in,
		Out:     &out,
		Params:  nil,
	}

	err = api.makeRequest(options)
	if err != nil {
		return
	}
	return
}

func (api *API) GetUnsorted(id string, params *Params) (unsorted []models.Unsorted, err error) {
	api.log("GetUnsorted request is started...")

	options := makeRequestOptions{
		Method:  fiber.MethodGet,
		BaseURL: leadUnsorted,
		In:      nil,
		Out:     &models.Unsorted{},
		Params:  params,
	}

	if id != "" {
		// ID exists
		options.BaseURL += "/" + id
		if err = api.makeRequest(options); err != nil {
			return
		}
		unsorted = []models.Unsorted{*options.Out.(*models.Unsorted)}
		api.log("returning the struct...")
		return
	} else {
		// All unsorted leads
		options.Out = &models.LeadResponse{}
		if err = api.makeRequest(options); err != nil {
			return
		}
		unsorted = options.Out.(*models.LeadResponse).Embedded.Unsorted
		api.log("returning the struct...")
		return
	}
}

func (api *API) GetLeads(id string, params *Params) (lead []models.Lead, err error) {
	api.log("GetLead request is started...")

	options := makeRequestOptions{
		Method:  fiber.MethodGet,
		BaseURL: leadURL,
		In:      nil,
		Out:     &models.Lead{},
		Params:  params,
	}

	if id != "" {
		// ID exists
		options.BaseURL += "/" + id
		err = api.makeRequest(options)
		if err != nil {
			return
		}
		lead = []models.Lead{*options.Out.(*models.Lead)}
		api.log("returning the struct...")
		return
	} else {
		// All leads
		options.Out = &models.LeadResponse{}
		err = api.makeRequest(options)
		if err != nil {
			return
		}
		lead = options.Out.(*models.LeadResponse).Embedded.Leads
		api.log("returning the struct...")
		return
	}
}

func (api *API) CreateLeads(lead *[]models.Lead, params *Params) (resp models.LeadResponse, err error) {
	api.log("CreateLeads request is started...")

	options := makeRequestOptions{
		Method:  fiber.MethodPost,
		BaseURL: leadURL,
		In:      lead,
		Out:     &resp,
		Params:  params,
	}
	err = api.makeRequest(options)
	if err != nil {
		return
	}

	api.log("returning the struct...")
	log.Println(lead)
	return
}

func (api *API) CreateLeadsComplex(lead *[]models.Lead) (resp []models.LeadComplexResponse, err error) {
	api.log("CreateLeads request is started...")

	options := makeRequestOptions{
		Method:  fiber.MethodPost,
		BaseURL: leadComplexURL,
		In:      lead,
		Out:     &resp,
		Params:  nil,
	}
	err = api.makeRequest(options)
	if err != nil {
		return
	}

	api.log("returning the struct...")
	return
}

func (api *API) ModifyLead(id string, lead *models.Lead, params *Params) (resp models.LeadModifyResponse, err error) {
	api.log("ModifyLead request started...")

	options := makeRequestOptions{
		Method:  fiber.MethodPatch,
		BaseURL: leadURL + "/" + id,
		In:      lead,
		Out:     &resp,
		Params:  params,
	}
	err = api.makeRequest(options)
	if err != nil {
		return
	}

	api.log("returning the struct...")
	return
}
