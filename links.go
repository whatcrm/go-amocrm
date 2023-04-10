package amocrm

import (
	"github.com/gofiber/fiber/v2"
	"github.com/whatcrm/go-amocrm/models"
)

func (api *API) GetLinks(entity, entityID string, params *Params) (out models.RequestResponse, err error) {
	api.log("GetLinks request is started...")

	var baseURL string
	switch entity {
	case "leads":
		baseURL += leadURL + "/" + entityID + linksURL

	}

	options := makeRequestOptions{
		Method:  fiber.MethodGet,
		BaseURL: baseURL,
		In:      nil,
		Out:     &out,
		Params:  params,
	}

	if err = api.makeRequest(options); err != nil {
		return
	}

	api.log("returning the struct...")
	return
}
