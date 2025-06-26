package amocrm

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/whatcrm/go-amocrm/models"
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

func (c *Get) Unsorted(id string, params *Params) (out models.RequestResponse, err error) {
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

		out = models.RequestResponse{
			Embedded: &models.ResponseEmbedded{
				Unsorted: []models.Unsorted{
					*options.Out.(*models.Unsorted),
				},
			},
		}
		c.api.log("returning the struct...")
		return
	}

	// All unsorted leads
	options.Out = &out
	if err = c.api.makeRequest(options); err != nil {
		return
	}
	c.api.log("returning the struct...")
	return
}

func (c *Get) Leads(id string, params *Params) (out models.RequestResponse, err error) {
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

		out = models.RequestResponse{
			Embedded: &models.ResponseEmbedded{
				Leads: []models.Lead{
					*options.Out.(*models.Lead),
				},
			},
		}
		c.api.log("returning the struct...")
		return
	}

	// All leads
	options.Out = &out
	err = c.api.makeRequest(options)
	if err != nil {
		return
	}

	c.api.log("returning the struct...")
	return

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
