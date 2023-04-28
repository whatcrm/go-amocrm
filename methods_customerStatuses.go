package amocrm

import (
	"github.com/gofiber/fiber/v2"
	"github.com/whatcrm/go-amocrm/models"
	"log"
)

func (c *Get) CustomerStatuses(statusID string) (out []models.CustomerStatus, err error) {
	c.api.log("GetCustomerStatuses request is started...")

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

	if err = c.api.makeRequest(options); err != nil {
		return
	}

	if statusID == "" {
		out = options.Out.(*models.MainResponse).Embedded.CustomerStatus
	} else {
		out = []models.CustomerStatus{*options.Out.(*models.CustomerStatus)}

	}

	log.Println(out)
	c.api.log("returning the struct...")
	return
}

func (c *Create) CustomerStatus(statuses []models.CustomerStatus) (out models.MainResponse, err error) {
	c.api.log("CreateCustomerStatus request is started...")

	options := makeRequestOptions{
		Method:  fiber.MethodPost,
		BaseURL: customerStatusesURL,
		In:      statuses,
		Out:     &out,
		Params:  nil,
	}

	if err = c.api.makeRequest(options); err != nil {
		return
	}
	c.api.log("returning the struct...")
	return
}

func (c *Update) CustomerStatus(statusID string, in *models.CustomerStatus) (out models.CustomerStatus, err error) {
	c.api.log("ModifyCustomerStatus request is started...")

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

	if err = c.api.makeRequest(options); err != nil {
		return
	}
	c.api.log("returning the struct...")
	return
}

func (c *Delete) CustomerStatus(statusID string) (err error) {
	c.api.log("ModifyCustomerStatus request is started...")

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

	if err = c.api.makeRequest(options); err != nil {
		return
	}
	c.api.log("returning the struct...")
	return
}
