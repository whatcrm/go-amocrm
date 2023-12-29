package amocrm

import (
	"github.com/gofiber/fiber/v2"
	"github.com/whatcrm/go-amocrm/models"
)

func (c *Get) Companies(companyID string, params *Params) (out []models.Company, err error) {
	c.api.log("Get Companies request is started...")

	options := makeRequestOptions{
		Method:  fiber.MethodGet,
		BaseURL: companiesURL,
		In:      nil,
		Out:     &models.Company{},
		Params:  params,
	}

	if companyID != "" {
		options.BaseURL += "/" + companyID
		if err = c.api.makeRequest(options); err != nil {
			return
		}
		out = []models.Company{*options.Out.(*models.Company)}
	}

	if companyID == "" {
		options.Out = &models.RequestResponse{}
		if err = c.api.makeRequest(options); err != nil {
			return
		}
		out = options.Out.(*models.RequestResponse).Embedded.Companies
	}

	c.api.log("returning the struct...")
	return
}

func (c *Create) Companies(in []models.Company) (out []models.Company, err error) {
	c.api.log("Create Company request is started...")

	options := makeRequestOptions{
		Method:  fiber.MethodPost,
		BaseURL: companiesURL,
		In:      in,
		Out:     &models.Company{},
		Params:  nil,
	}

	options.Out = &models.RequestResponse{}
	if err = c.api.makeRequest(options); err != nil {
		return
	}
	out = options.Out.(*models.RequestResponse).Embedded.Companies

	c.api.log("returning the struct...")
	return
}

func (c *Update) Companies(companyID string, in []models.Company) (out []models.Company, err error) {
	c.api.log("CustomersMode request is started...")

	options := makeRequestOptions{
		Method:  fiber.MethodPatch,
		BaseURL: companiesURL,
		In:      in,
		Out:     &models.Company{},
		Params:  nil,
	}

	if companyID != "" {
		options.BaseURL += "/" + companyID
		options.In = in[0]

		if err = c.api.makeRequest(options); err != nil {
			return
		}
		out = []models.Company{*options.Out.(*models.Company)}
	}

	if companyID == "" {
		options.Out = &models.RequestResponse{}
		if err = c.api.makeRequest(options); err != nil {
			return
		}
		out = options.Out.(*models.RequestResponse).Embedded.Companies
	}

	c.api.log("returning the struct...")
	return
}
