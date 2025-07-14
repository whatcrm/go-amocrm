package amocrm

import (
	"github.com/gofiber/fiber/v2"
	"github.com/whatcrm/go-amocrm/models"
)

func (c *Get) Tasks(taskID string, params *Params) (out models.RequestResponse, err error) {
	c.api.log("GetContacts request is started...")

	options := makeRequestOptions{
		Method:  fiber.MethodGet,
		BaseURL: tasksURL,
		In:      nil,
		Out:     &models.Task{},
		Params:  params,
	}

	if taskID != "" {
		options.BaseURL += "/" + taskID
		if err = c.api.makeRequest(options); err != nil {
			return
		}
		out = models.RequestResponse{
			Embedded: &models.ResponseEmbedded{
				Tasks: []models.Task{
					*options.Out.(*models.Task),
				},
			},
		}

		c.api.log("returning the struct...")
		return
	}

	options.Out = &out
	if err = c.api.makeRequest(options); err != nil {
		return
	}

	c.api.log("returning the struct...")
	return
}

func (c *Create) Tasks(in []models.Task) (out models.RequestResponse, err error) {
	c.api.log("CreateContact request is started...")

	options := makeRequestOptions{
		Method:  fiber.MethodPost,
		BaseURL: tasksURL,
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

func (c *Update) Tasks(taskID string, in []models.Task) (out []models.Task, err error) {
	c.api.log("ModifyContacts request is started...")

	options := makeRequestOptions{
		Method:  fiber.MethodPatch,
		BaseURL: tasksURL,
		In:      in,
		Out:     &models.Task{},
		Params:  nil,
	}

	if taskID != "" {
		options.BaseURL += "/" + taskID
		options.In = in[0]
		if err = c.api.makeRequest(options); err != nil {
			return
		}
		out = []models.Task{*options.Out.(*models.Task)}
	}

	if taskID == "" {
		options.Out = &models.RequestResponse{}
		if err = c.api.makeRequest(options); err != nil {
			return
		}
		out = options.Out.(*models.RequestResponse).Embedded.Tasks
	}

	c.api.log("returning the struct...")
	return
}
