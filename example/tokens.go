package example

import (
	"github.com/whatcrm/go-amocrm"
)

func main() {

	// information from your database
	domain := "example.amocrm.ru"
	refreshToken := "Your-Refresh-Token"

	amoCRM := amocrm.NewAPI(clientID, clientSecret, redirectURI)

	oauth := amocrm.OAuth{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		GrantType:    "refresh-token",
		RefreshToken: refreshToken,
		RedirectURI:  redirectURI,
	}

	upd, err := amoCRM.Create().TokenUPD(&oauth)
	if err != nil {
		return
	}

	amoCRM.SetOptions(domain, upd.AccessToken, true)

	/*

		Here's we just updated the tokens, and they are now available for new requests.

	*/
}
