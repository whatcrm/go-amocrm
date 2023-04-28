package amocrm

import (
	"github.com/gofiber/fiber/v2"
	"github.com/whatcrm/go-amocrm/models"
)

func (c *Create) Tag(entity string, tag *[]models.Tag) (out models.RequestResponse, err error) {
	c.api.log("CreateTag request started...")

	if entity == "" {
		err = fiber.ErrBadRequest
		return
	}

	options := makeRequestOptions{
		Method:  "",
		BaseURL: noEntityURL + entity + tagsURL,
		In:      tag,
		Out:     &out,
		Params:  nil,
	}

	if err = c.api.makeRequest(options); err != nil {
		return
	}

	c.api.log("returning the struct...")
	return
}
