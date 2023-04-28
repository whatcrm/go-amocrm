package amocrm

type Create struct {
	api *API
}
type Get struct {
	api *API
}
type Update struct {
	api *API
}
type Delete struct {
	api *API
}

func (api *API) Create() *Create {
	return &Create{api}
}

func (api *API) Get() *Get {
	return &Get{api}
}
func (api *API) Update() *Update {
	return &Update{api}
}
func (api *API) Delete() *Delete {
	return &Delete{api}
}
