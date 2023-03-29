package amocrm

import (
	"github.com/gofiber/fiber/v2"
	"github.com/whatcrm/go-amocrm/models"
)

func (api *API) CustomersMode(in models.CustomersMode) (out models.CustomersMode, err error) {
	api.log("CustomersMode request is started...")

	options := makeRequestOptions{
		Method:  fiber.MethodPatch,
		BaseURL: customersMode,
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

func (api *API) CustomersList(id string, customer *models.Customer, params *Params) (out []models.Customer, err error) {
	api.log("CustomersList request is started")

	options := makeRequestOptions{
		Method:  fiber.MethodGet,
		BaseURL: customersURL,
		In:      customer,
		Out:     &models.Customer{},
		Params:  params,
	}

	if id != "" {
		// ID exists
		options.BaseURL += "/" + id
		if err = api.makeRequest(options); err != nil {
			return
		}
		out = []models.Customer{*options.Out.(*models.Customer)}
		api.log("returning the struct...")
		return
	} else {
		// All customers
		options.Out = &models.RequestResponse{}
		if err = api.makeRequest(options); err != nil {
			return
		}
		out = options.Out.(*models.RequestResponse).Embedded.Customers
		api.log("returning the struct...")
		return
	}
}
