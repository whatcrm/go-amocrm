package example

import (
	"github.com/whatcrm/go-amocrm"
	"github.com/whatcrm/go-amocrm/models"
	"log"
	"time"
)

func main() {

	// information from your database
	domain := "example.amocrm.ru"
	accessToken := "Your-Acess-Token"

	// connection to Rest API
	amoCRM := amocrm.NewAPI(clientID, clientSecret, redirectURI)

	// Setting the options for the request
	amoCRM.SetOptions(domain, accessToken, true)

	// we will need the models from go-amocrm/models
	// in order to create an object
	myLeads := []models.Lead{
		{
			Name:      "My Lead",
			Price:     10000,
			CreatedAt: int(time.Now().Unix()),
		},
		{
			Name:      "My Lead 2",
			Price:     15000,
			CreatedAt: int(time.Now().Unix()),
		},
	}

	leads, err := amoCRM.Create().Leads(&myLeads, nil)
	if err != nil {
		return
	}

	log.Println(leads)
	/*
		models.RequestResponse in "leads" variable contains your leads,
		make sure to handle the error properly, this is just an example
	*/
}
