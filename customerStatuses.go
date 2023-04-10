package amocrm

import (
	"github.com/gofiber/fiber/v2"
	"github.com/whatcrm/go-amocrm/models"
	"log"
)

func (api *API) GetCustomerStatuses(statusID string) (out []models.CustomerStatus, err error) {
	api.log("GetCustomerStatuses request is started...")

	options := makeRequestOptions{
		Method:  fiber.MethodGet,
		BaseURL: customerStatusesURL + "/" + statusID,
		In:      nil,
		Out:     &models.CustomerStatus{},
		Params:  nil,
	}

	if statusID == "" {
		options.BaseURL = customerStatusesURL
		options.Out = &models.MainResponse{}
	}

	if err = api.makeRequest(options); err != nil {
		return
	}

	if statusID == "" {
		out = options.Out.(*models.MainResponse).Embedded.CustomerStatus
	} else {
		out = []models.CustomerStatus{*options.Out.(*models.CustomerStatus)}

	}

	log.Println(out)
	api.log("returning the struct...")
	return
}

func (api *API) CreateCustomerStatus(statuses []models.CustomerStatus) (out models.MainResponse, err error) {
	api.log("CreateCustomerStatus request is started...")

	options := makeRequestOptions{
		Method:  fiber.MethodPost,
		BaseURL: customerStatusesURL,
		In:      statuses,
		Out:     &out,
		Params:  nil,
	}

	if err = api.makeRequest(options); err != nil {
		return
	}
	api.log("returning the struct...")
	return
}

func (api *API) ModifyCustomerStatus(statusID string, in *models.CustomerStatus) (out models.CustomerStatus, err error) {
	api.log("ModifyCustomerStatus request is started...")

	if statusID == "" {
		err = fiber.ErrBadRequest
		return
	}

	options := makeRequestOptions{
		Method:  fiber.MethodPatch,
		BaseURL: customerStatusesURL + "/" + statusID,
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

func (api *API) RemoveCustomerStatus(statusID string) (err error) {
	api.log("ModifyCustomerStatus request is started...")

	if statusID == "" {
		err = fiber.ErrBadRequest
		return
	}

	options := makeRequestOptions{
		Method:  fiber.MethodDelete,
		BaseURL: customerStatusesURL + "/" + statusID,
		In:      nil,
		Out:     nil,
		Params:  nil,
	}

	if err = api.makeRequest(options); err != nil {
		return
	}
	api.log("returning the struct...")
	return
}
