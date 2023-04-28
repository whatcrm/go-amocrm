package amocrm

import (
	"github.com/gofiber/fiber/v2"
	"github.com/whatcrm/go-amocrm/models"
)

func (c *Get) Links(entity, entityID string, params *Params) (out models.RequestResponse, err error) {
	c.api.log("GetLinks request is started...")

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

	if err = c.api.makeRequest(options); err != nil {
		return
	}

	c.api.log("returning the struct...")
	return
}
