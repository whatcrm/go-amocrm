package example

import (
	"github.com/whatcrm/go-amocrm"
)

const (
	clientID     = "Your-Client-ID"
	clientSecret = "Your-Client-Secret"
	redirectURI  = "Your-Redirect-URL"
)

func main() {

	domain := "example.amocrm.ru"
	accessToken := "Your-Access-Token"

	amoCRM := amocrm.NewAPI(clientID, clientSecret, redirectURI)
	amoCRM.SetOptions(domain, accessToken, true)

	/*

		Now, you can use the *amoCRM* (*API) to make requests to amoCRM/Kommo
		In order for it to work, you have to maintain tokens,
		therefore make sure to save them and retrieve when needed.

		You can update your tokens with TokenUPD method, like in examples/tokens.go

	*/
}
