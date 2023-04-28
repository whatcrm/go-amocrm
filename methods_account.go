package amocrm

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

func (c *Create) TokenUPD(oAuth *OAuth) (t *Tokens, err error) {
	c.api.log("Refresh Token Updater is entered...")
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

	errs := c.api.makeRequest(options)
	if errs != nil {
		return nil, fmt.Errorf("few errors: %v", errs)
	}

	c.api.log("Refresh Token updated...")
	return
}

func (c *Get) GetAccount(params *Params) (acc Account, err error) {
	// INFO - PARAMETERS: WITH
	c.api.log("GetAccount request is started...")

	options := makeRequestOptions{
		Method:  fiber.MethodGet,
		BaseURL: getAccountURL,
		In:      nil,
		Out:     &acc,
		Params:  params,
	}
	if err = c.api.makeRequest(options); err != nil {
		return
	}

	c.api.log("returning the struct...")
	log.Printf("%+v", acc)
	return
}

func (c *Get) SubdomainFromToken() (acc OauthAccount, err error) {
	c.api.log("OAuthAccount request is started...")

	options := makeRequestOptions{
		Method:  fiber.MethodGet,
		BaseURL: accountURL,
		In:      nil,
		Out:     &acc,
	}
	if err = c.api.makeRequest(options); err != nil {
		return
	}

	c.api.log("returning the struct...")
	log.Printf("%+v", acc)
	return
}
