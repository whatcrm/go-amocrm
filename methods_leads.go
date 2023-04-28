package amocrm

import (
	"github.com/gofiber/fiber/v2"
	"github.com/whatcrm/go-amocrm/models"
	"log"
)

// TODO GET Parameters in Requests - https://www.amocrm.ru/developers/content/crm_platform/leads-api#leads-list

func (c *Create) UnsortedSIP(in *[]models.UnsortedSIP) (out *models.MainResponse, err error) {
	c.api.log("CreateUnsortedSIP request is started...")

	options := makeRequestOptions{
		Method:  fiber.MethodPost,
		BaseURL: leadUnsortedSIP,
		In:      in,
		Out:     &out,
		Params:  nil,
	}

	err = c.api.makeRequest(options)
	if err != nil {
		return
	}
	return
}

func (c *Get) Unsorted(id string, params *Params) (unsorted []models.Unsorted, err error) {
	c.api.log("GetUnsorted request is started...")

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
		if err = c.api.makeRequest(options); err != nil {
			return
		}
		unsorted = []models.Unsorted{*options.Out.(*models.Unsorted)}
		c.api.log("returning the struct...")
		return
	} else {
		// All unsorted leads
		options.Out = &models.RequestResponse{}
		if err = c.api.makeRequest(options); err != nil {
			return
		}
		unsorted = options.Out.(*models.RequestResponse).Embedded.Unsorted
		c.api.log("returning the struct...")
		return
	}
}

func (c *Get) Leads(id string, params *Params) (lead []models.Lead, err error) {
	c.api.log("GetLead request is started...")

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
		err = c.api.makeRequest(options)
		if err != nil {
			return
		}
		lead = []models.Lead{*options.Out.(*models.Lead)}
		c.api.log("returning the struct...")
		return
	} else {
		// All leads
		options.Out = &models.RequestResponse{}
		err = c.api.makeRequest(options)
		if err != nil {
			return
		}
		lead = options.Out.(*models.RequestResponse).Embedded.Leads
		c.api.log("returning the struct...")
		return
	}
}

func (c *Create) Leads(lead *[]models.Lead, params *Params) (resp models.RequestResponse, err error) {
	c.api.log("CreateLeads request is started...")

	options := makeRequestOptions{
		Method:  fiber.MethodPost,
		BaseURL: leadURL,
		In:      lead,
		Out:     &resp,
		Params:  params,
	}
	err = c.api.makeRequest(options)
	if err != nil {
		return
	}

	c.api.log("returning the struct...")
	log.Println(lead)
	return
}

func (c *Create) LeadsComplex(lead *[]models.Lead) (resp []models.LeadComplexResponse, err error) {
	c.api.log("CreateLeads request is started...")

	options := makeRequestOptions{
		Method:  fiber.MethodPost,
		BaseURL: leadComplexURL,
		In:      lead,
		Out:     &resp,
		Params:  nil,
	}
	err = c.api.makeRequest(options)
	if err != nil {
		return
	}

	c.api.log("returning the struct...")
	return
}

func (c *Update) Lead(id string, lead *models.Lead, params *Params) (resp models.LeadModifyResponse, err error) {
	c.api.log("ModifyLead request started...")

	options := makeRequestOptions{
		Method:  fiber.MethodPatch,
		BaseURL: leadURL + "/" + id,
		In:      lead,
		Out:     &resp,
		Params:  params,
	}
	err = c.api.makeRequest(options)
	if err != nil {
		return
	}

	c.api.log("returning the struct...")
	return
}
