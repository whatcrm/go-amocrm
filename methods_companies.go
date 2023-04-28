package amocrm

import (
	"github.com/gofiber/fiber/v2"
	"github.com/whatcrm/go-amocrm/models"
)

func (api *API) GetCompanies(companyID string, params *Params) (out []models.Company, err error) {
	api.log("CustomersMode request is started...")

	options := makeRequestOptions{
		Method:  fiber.MethodGet,
		BaseURL: companiesURL,
		In:      nil,
		Out:     &models.Company{},
		Params:  params,
	}

	if companyID != "" {
		options.BaseURL += "/" + companyID
		if err = api.makeRequest(options); err != nil {
			return
		}
		out = []models.Company{*options.Out.(*models.Company)}
	}

	if companyID == "" {
		options.Out = &models.RequestResponse{}
		if err = api.makeRequest(options); err != nil {
			return
		}
		out = options.Out.(*models.RequestResponse).Embedded.Companies
	}

	api.log("returning the struct...")
	return
}

func (api *API) CreateCompany(in *[]models.Company) (out []models.Company, err error) {
	api.log("CustomersMode request is started...")

	options := makeRequestOptions{
		Method:  fiber.MethodPost,
		BaseURL: companiesURL,
		In:      in,
		Out:     &models.Company{},
		Params:  nil,
	}

	if len(*in) <= 1 {
		options.Out = &models.Company{}
		if err = api.makeRequest(options); err != nil {
			return
		}
		out = []models.Company{*options.Out.(*models.Company)}
	} else {
		options.Out = &models.RequestResponse{}
		if err = api.makeRequest(options); err != nil {
			return
		}
		out = options.Out.(*models.RequestResponse).Embedded.Companies
	}

	api.log("returning the struct...")
	return
}

func (api *API) ModifyCompany(companyID string, in *[]models.Company) (out []models.Company, err error) {
	api.log("CustomersMode request is started...")

	options := makeRequestOptions{
		Method:  fiber.MethodPatch,
		BaseURL: companiesURL,
		In:      in,
		Out:     &models.Company{},
		Params:  nil,
	}

	if companyID != "" {
		options.BaseURL += "/" + companyID
		options.In = (*in)[0]

		if err = api.makeRequest(options); err != nil {
			return
		}
		out = []models.Company{*options.Out.(*models.Company)}
	}

	if companyID == "" {
		options.Out = &models.RequestResponse{}
		if err = api.makeRequest(options); err != nil {
			return
		}
		out = options.Out.(*models.RequestResponse).Embedded.Companies
	}

	api.log("returning the struct...")
	return
}
