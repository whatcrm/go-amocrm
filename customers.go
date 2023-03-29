package amocrm

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/whatcrm/go-amocrm/models"
)

func (api *API) CustomersMode(in models.CustomersMode) (out models.CustomersMode, err error) {
	api.log("CustomersMode request is started...")

	options := makeRequestOptions{
		Method:  fiber.MethodPatch,
		BaseURL: customersMode,
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

func (api *API) CustomersList(customerID string, params *Params) (out []models.Customer, err error) {
	api.log("CustomersList request is started")

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
		if err = api.makeRequest(options); err != nil {
			return
		}
		out = []models.Customer{*options.Out.(*models.Customer)}
		api.log("returning the struct...")
		return
	} else {
		// All customers
		options.Out = &models.RequestResponse{}
		if err = api.makeRequest(options); err != nil {
			return
		}
		out = options.Out.(*models.RequestResponse).Embedded.Customers
		api.log("returning the struct...")
		return
	}
}

func (api *API) CreateCustomer(in *models.Customer) (out models.RequestResponse, err error) {
	api.log("CreateCustomer request is started...")

	options := makeRequestOptions{
		Method:  fiber.MethodPost,
		BaseURL: customersURL,
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

func (api *API) ModifyCustomers(customerID string, in *[]models.Customer) (out models.RequestResponse, err error) {
	api.log("ModifyCustomers request is started")

	options := makeRequestOptions{
		Method:  fiber.MethodPatch,
		BaseURL: customersURL,
		In:      in,
		Out:     &models.RequestResponse{},
		Params:  nil,
	}

	if customerID != "" {
		// ID exists
		options.BaseURL += "/" + customerID
		options.In = (*in)[0]
	}

	if err = api.makeRequest(options); err != nil {
		return
	}
	api.log("returning the struct...")
	return
}

func (api *API) TransactionsList(customerID string, params *Params) (out models.RequestResponse, err error) {
	api.log("TransactionsList request is started")

	options := makeRequestOptions{
		Method:  fiber.MethodPost,
		BaseURL: customersURL + transactionsURL,
		In:      nil,
		Out:     &out,
		Params:  params,
	}

	if customerID != "" {
		// Customer ID exists
		options.BaseURL = customersURL + "/" + customerID + transactionsURL
	}

	// All the transactions or just customers transactions
	if err = api.makeRequest(options); err != nil {
		return
	}
	api.log("returning the struct...")
	return
}

func (api *API) GetTransaction(customerID, transactionID string) (out models.Transaction, err error) {
	api.log("TransactionsList request is started")

	options := makeRequestOptions{
		Method:  fiber.MethodPost,
		BaseURL: customersURL + transactionsURL,
		In:      nil,
		Out:     &out,
		Params:  nil,
	}

	if customerID != "" && transactionID != "" {
		// customerID and transactionID exists
		options.BaseURL = customersURL + "/" + customerID + transactionsURL + transactionID
	} else if transactionID != "" {
		// Only transactionID exists
		options.BaseURL += transactionID
	} else if transactionID == "" {
		err = fmt.Errorf("transactionID can not be empty")
		return
	}

	// All the transactions or just customers transactions
	if err = api.makeRequest(options); err != nil {
		return
	}

	api.log("returning the struct...")
	return
}

func (api *API) SetTransaction(customerID string) (out models.RequestResponse, err error) {

	options := makeRequestOptions{
		Method:  fiber.MethodPost,
		BaseURL: customersURL + customerID + transactionsURL,
		In:      nil,
		Out:     &out,
		Params:  nil,
	}

	if err = api.makeRequest(options); err != nil {
		return
	}

	api.log("returning the struct...")
	return
}

func (api *API) RemoveTransaction(customerID, transactionID string) (out models.RequestResponse, err error) {

	options := makeRequestOptions{
		Method:  fiber.MethodDelete,
		BaseURL: customersURL + transactionsURL + transactionID,
		In:      nil,
		Out:     &out,
		Params:  nil,
	}

	if customerID != "" && transactionID != "" {
		// customerID and transactionID exists
		options.BaseURL = customersURL + "/" + customerID + transactionsURL + transactionID
	} else if transactionID == "" {
		err = fmt.Errorf("transactionID can not be empty")
		return
	}

	if err = api.makeRequest(options); err != nil {
		return
	}
	api.log("returning the struct...")
	return
}

func (api *API) SetBonusPoints(customerID string, value *models.BonusPoints) (out models.Points, err error) {

	if value.Redeem != 0 && value.Earn != 0 {
		err = fmt.Errorf("both fields are not allowed")
		return
	}

	options := makeRequestOptions{
		Method:  fiber.MethodPost,
		BaseURL: customersURL + customerID + bonusPointsURL,
		In:      value,
		Out:     &out,
		Params:  nil,
	}

	if err = api.makeRequest(options); err != nil {
		return
	}
	api.log("returning the struct...")
	return
}
