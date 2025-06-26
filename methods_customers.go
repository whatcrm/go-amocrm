package amocrm

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/whatcrm/go-amocrm/models"
)

func (c *Create) CustomersMode(in models.CustomersMode) (out models.CustomersMode, err error) {
	c.api.log("CustomersMode request is started...")

	options := makeRequestOptions{
		Method:  fiber.MethodPatch,
		BaseURL: customersMode,
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

func (c *Get) Customers(customerID string, params *Params) (out *models.RequestResponse, err error) {
	c.api.log("GetCustomers request is started")

	options := makeRequestOptions{
		Method:  fiber.MethodGet,
		BaseURL: customersURL,
		In:      nil,
		Out:     &models.Customer{},
		Params:  params,
	}

	if customerID != "" {
		// ID exists
		options.BaseURL += "/" + customerID
		if err = c.api.makeRequest(options); err != nil {
			return
		}
		out = &models.RequestResponse{
			Embedded: &models.ResponseEmbedded{
				Customers: []models.Customer{
					*options.Out.(*models.Customer),
				},
			},
		}
		c.api.log("returning the struct...")
		return
	}

	// All customers
	options.Out = &out
	if err = c.api.makeRequest(options); err != nil {
		return
	}
	c.api.log("returning the struct...")
	return
}

func (c *Create) Customer(in *[]models.Customer) (out models.RequestResponse, err error) {
	c.api.log("CreateCustomer request is started...")

	options := makeRequestOptions{
		Method:  fiber.MethodPost,
		BaseURL: customersURL,
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

func (c *Update) Customers(customerID string, in *[]models.Customer) (out models.CustomerResponse, err error) {
	c.api.log("ModifyCustomers request is started")

	options := makeRequestOptions{
		Method:  fiber.MethodPatch,
		BaseURL: customersURL,
		In:      in,
		Out:     &models.CustomerResponse{},
		Params:  nil,
	}

	if customerID != "" {
		// ID exists
		options.BaseURL += "/" + customerID
		options.In = (*in)[0]
	}

	if err = c.api.makeRequest(options); err != nil {
		return
	}
	c.api.log("returning the struct...")
	return
}

func (c *Get) Transactions(customerID, transactionID string, params *Params) (out *models.RequestResponse, err error) {
	// NOTE: params only for all transactions
	c.api.log("GetTransactions request is started")

	url := customersURL + transactionsURL
	if customerID != "" {
		url = customersURL + "/" + customerID + transactionsURL
	}
	if transactionID != "" {
		url += "/" + transactionID
	}

	options := makeRequestOptions{
		Method:  fiber.MethodGet,
		BaseURL: url,
		In:      nil,
		Out:     &out,
	}

	if transactionID == "" && customerID == "" {
		// the API does not accept Get parameters
		// unless it's the list of all transactions
		// so you could filter it
		options.Params = params
	}

	if transactionID != "" {
		options.Out = &models.Transaction{}
	} else {
		options.Out = &models.RequestResponse{}
	}

	if err = c.api.makeRequest(options); err != nil {
		return
	}

	if transactionID != "" {
		out = &models.RequestResponse{
			Embedded: &models.ResponseEmbedded{
				Transactions: []models.Transaction{
					*options.Out.(*models.Transaction),
				},
			},
		}
	}

	c.api.log("returning the struct...")
	return
}

func (c *Create) Transaction(customerID string, t *[]models.Transaction) (out models.MainResponse, err error) {
	c.api.log("SetTransaction request is started...")

	options := makeRequestOptions{
		Method:  fiber.MethodPost,
		BaseURL: customersURL + "/" + customerID + transactionsURL,
		In:      t,
		Out:     &out,
		Params:  nil,
	}

	if err = c.api.makeRequest(options); err != nil {
		return
	}

	c.api.log("returning the struct...")
	return
}

func (c *Delete) Transaction(customerID, transactionID string) (err error) {
	c.api.log("RemoveTransaction request is started...")

	options := makeRequestOptions{
		Method:  fiber.MethodDelete,
		BaseURL: customersURL + transactionsURL + "/" + transactionID,
		In:      nil,
		Out:     nil,
		Params:  nil,
	}

	if customerID != "" && transactionID != "" {
		// customerID and transactionID exists
		options.BaseURL = customersURL + "/" + customerID + transactionsURL + "/" + transactionID
	} else if transactionID == "" {
		err = fmt.Errorf("transactionID can not be empty")
		return
	}

	if err = c.api.makeRequest(options); err != nil {
		return
	}
	c.api.log("returning the struct...")
	return
}

func (c *Create) BonusPoints(customerID string, value *models.BonusPoints) (out models.Points, err error) {
	c.api.log("SetBonusPoints request is started...")

	if value.Redeem != 0 && value.Earn != 0 {
		err = fmt.Errorf("both fields are not allowed")
		return
	}

	options := makeRequestOptions{
		Method:  fiber.MethodPost,
		BaseURL: customersURL + "/" + customerID + bonusPointsURL,
		In:      value,
		Out:     &out,
		Params:  nil,
	}

	if err = c.api.makeRequest(options); err != nil {
		return
	}
	c.api.log("returning the struct...")
	return
}
