package go_amocrm

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

func (api *API) TokenUPD(oAuth *OAuth) (t *Tokens, err error) {
	api.log("Refresh Token Updater is entered...")
	//a, req := api.getAgent(fiber.MethodPost, accessTokenURL, nil)

	fmt.Printf("%+v\n", oAuth)

	//t := Tokens{}
	options := makeRequestOptions{
		Method:  fiber.MethodPost,
		BaseURL: accessTokenURL,
		In:      oAuth,
		Out:     &t,
		Params:  nil,
	}

	errs := api.makeRequest(options)
	if errs != nil {
		return nil, fmt.Errorf("few errors: %v", errs)
	}

	api.log("Refresh Token updated...")
	return
}

func (api *API) GetAccount(params *Params) (acc Account, err error) {
	// INFO - PARAMETERS: WITH
	api.log("GetAccount request is started...")

	options := makeRequestOptions{
		Method:  fiber.MethodGet,
		BaseURL: getAccountURL,
		In:      nil,
		Out:     &acc,
		Params:  params,
	}
	if err = api.makeRequest(options); err != nil {
		return
	}

	api.log("returning the struct...")
	log.Printf("%+v", acc)
	return
}
