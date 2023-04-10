package models

type CustomerStatus struct {
	ID         int           `json:"id,omitempty"`
	RequestID  string        `json:"request_id,omitempty"`
	Name       string        `json:"name,omitempty"`
	Sort       int           `json:"sort,omitempty"`
	IsDefault  bool          `json:"is_default,omitempty"`
	Conditions []Condition   `json:"conditions,omitempty"`
	Color      string        `json:"color,omitempty"`
	Type       int           `json:"type,omitempty"`
	AccountID  int           `json:"account_id,omitempty"`
	Links      *LinkResponse `json:"_links,omitempty"`
}

type Condition []struct {
	Type  string `json:"type"`
	Match struct {
		Value int `json:"value"`
	} `json:"match"`
	Conditions []struct {
		Value int `json:"value"`
	} `json:"conditions"`
}
