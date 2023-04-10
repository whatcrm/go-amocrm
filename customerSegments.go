package amocrm

import (
	"github.com/gofiber/fiber/v2"
	"github.com/whatcrm/go-amocrm/models"
)

func (api *API) GetCustomerSegments(segmentID string) (out []models.CustomerSegment, err error) {
	api.log("CustomerSegmentsList request is started...")

	options := makeRequestOptions{
		Method:  fiber.MethodGet,
		BaseURL: customerSegmentsURL,
		In:      nil,
		Out:     &models.CustomerSegment{},
		Params:  nil,
	}

	if segmentID != "" {
		options.BaseURL += "/" + segmentID
	} else {
		options.Out = &models.MainResponse{}
	}

	if err = api.makeRequest(options); err != nil {
		return
	}

	if segmentID == "" {
		out = options.Out.(*models.MainResponse).Embedded.CustomerSegment
	} else {
		out = []models.CustomerSegment{*options.Out.(*models.CustomerSegment)}

	}

	api.log("returning the struct...")
	return
}

func (api *API) CreateCustomerSegment(in *models.CustomerSegment) (err error) {
	api.log("CreateCustomerSegment request is started...")

	options := makeRequestOptions{
		Method:  fiber.MethodPost,
		BaseURL: customerSegmentsURL,
		In:      in,
		Out:     nil,
		Params:  nil,
	}

	if err = api.makeRequest(options); err != nil {
		return
	}

	api.log("returning the struct...")
	return
}

func (api *API) ModifyCustomerSegment(segmentID string, in *models.CustomerSegment) (out models.CustomerSegment, err error) {
	api.log("ModifyCustomerSegment request is started...")

	if segmentID == "" {
		err = fiber.ErrBadRequest
		return
	}

	options := makeRequestOptions{
		Method:  fiber.MethodPatch,
		BaseURL: customerSegmentsURL + "/" + segmentID,
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

func (api *API) RemoveCustomerSegment(segmentID string) (err error) {
	api.log("CustomersMode request is started...")

	options := makeRequestOptions{
		Method:  fiber.MethodDelete,
		BaseURL: customerSegmentsURL + "/" + segmentID,
		In:      nil,
		Out:     nil,
	}

	if err = api.makeRequest(options); err != nil {
		return
	}

	api.log("returning the struct...")
	return
}
