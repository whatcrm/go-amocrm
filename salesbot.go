package amocrm

import (
	"github.com/gofiber/fiber/v2"
)

type WidgetRequestResponse struct {
	Data            WidgetRequestData `json:"data,omitempty"`
	ExecuteHandlers *ExecuteHandlers  `json:"execute_handlers,omitempty"`
}

type WidgetRequestData struct {
	Status string `json:"status,omitempty"`
}

type ExecuteHandlers struct {
	Handler string         `json:"handler,omitempty"`
	Params  *HandlerParams `json:"params,omitempty"`
}

type HandlerParams struct {
	Type    string   `json:"type,omitempty"`
	Value   string   `json:"value,omitempty"`
	Buttons []string `json:"buttons,omitempty"`
}

func (api *API) ReplyToWidgetRequest(data, requestURL string) (err error) {
	api.log("WidgetRequestReply request is started...")

	wrp := WidgetRequestResponse{
		Data: WidgetRequestData{
			Status: data,
		},
	}

	options := makeRequestOptions{
		Method:  fiber.MethodPost,
		BaseURL: requestURL,
		In:      wrp,
	}

	err = api.makeRequest(options)
	if err != nil {
		return err
	}

	api.log("WidgetRequestReply request is finished...")
	return
}
