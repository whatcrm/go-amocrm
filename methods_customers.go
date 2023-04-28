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

func (api *API) GetCustomers(customerID string, params *Params) (out []models.Customer, err error) {
	api.log("GetCustomers request is started")

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

func (api *API) CreateCustomer(in *[]models.Customer) (out models.RequestResponse, err error) {
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

func (api *API) ModifyCustomers(customerID string, in *[]models.Customer) (out models.CustomerResponse, err error) {
	api.log("ModifyCustomers request is started")

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

	if err = api.makeRequest(options); err != nil {
		return
	}
	api.log("returning the struct...")
	return
}

func (api *API) GetTransactions(customerID, transactionID string, params *Params) (out []models.Transaction, err error) {
	// NOTE: params only for all transactions
	api.log("GetTransactions request is started")

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
		Out:     &models.RequestResponse{},
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

	if err = api.makeRequest(options); err != nil {
		return
	}

	if transactionID != "" {
		out = []models.Transaction{*options.Out.(*models.Transaction)}
	} else {
		out = options.Out.(*models.RequestResponse).Embedded.Transactions
	}
	api.log("returning the struct...")
	return
}

func (api *API) SetTransaction(customerID string, t *[]models.Transaction) (out models.MainResponse, err error) {
	api.log("SetTransaction request is started...")

	options := makeRequestOptions{
		Method:  fiber.MethodPost,
		BaseURL: customersURL + "/" + customerID + transactionsURL,
		In:      t,
		Out:     &out,
		Params:  nil,
	}

	if err = api.makeRequest(options); err != nil {
		return
	}

	api.log("returning the struct...")
	return
}

func (api *API) RemoveTransaction(customerID, transactionID string) (err error) {
	api.log("RemoveTransaction request is started...")

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

	if err = api.makeRequest(options); err != nil {
		return
	}
	api.log("returning the struct...")
	return
}

func (api *API) SetBonusPoints(customerID string, value *models.BonusPoints) (out models.Points, err error) {
	api.log("SetBonusPoints request is started...")

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

	if err = api.makeRequest(options); err != nil {
		return
	}
	api.log("returning the struct...")
	return
}
