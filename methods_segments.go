package amocrm

import (
	"github.com/gofiber/fiber/v2"
	"github.com/whatcrm/go-amocrm/models"
)

func (c *Get) CustomerSegmentsList(segmentID string) (out models.MainResponse, err error) {
	c.api.log("CustomerSegmentsList request is started...")

	options := makeRequestOptions{
		Method:  fiber.MethodGet,
		BaseURL: customerSegmentsURL,
		In:      nil,
		Out:     &out,
		Params:  nil,
	}

	if segmentID != "" {
		options.BaseURL += "/" + segmentID
	}

	if err = c.api.makeRequest(options); err != nil {
		return
	}

	c.api.log("returning the struct...")
	return
}

//
//func (api *API) CustomersMode(in models.CustomersMode) (out models.CustomersMode, err error) {
//	api.log("CustomersMode request is started...")
//
//	options := makeRequestOptions{
//		Method:  fiber.MethodPatch,
//		BaseURL: customersMode,
//		In:      in,
//		Out:     &out,
//		Params:  nil,
//	}
//
//	if err = api.makeRequest(options); err != nil {
//		return
//	}
//
//	api.log("returning the struct...")
//	return
//}
//
//func (api *API) CustomersMode(in models.CustomersMode) (out models.CustomersMode, err error) {
//	api.log("CustomersMode request is started...")
//
//	options := makeRequestOptions{
//		Method:  fiber.MethodPatch,
//		BaseURL: customersMode,
//		In:      in,
//		Out:     &out,
//		Params:  nil,
//	}
//
//	if err = api.makeRequest(options); err != nil {
//		return
//	}
//
//	api.log("returning the struct...")
//	return
//}
//
//func (api *API) CustomersMode(in models.CustomersMode) (out models.CustomersMode, err error) {
//	api.log("CustomersMode request is started...")
//
//	options := makeRequestOptions{
//		Method:  fiber.MethodPatch,
//		BaseURL: customersMode,
//		In:      in,
//		Out:     &out,
//		Params:  nil,
//	}
//
//	if err = api.makeRequest(options); err != nil {
//		return
//	}
//
//	api.log("returning the struct...")
//	return
//}
