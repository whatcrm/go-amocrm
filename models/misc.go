package models

type CustomFields struct {
	//FieldID   int                  `json:"field_id"`
	FieldCode string               `json:"field_code"`
	Values    []CustomFieldsValues `json:"values"`
}

type CustomFieldsValues struct {
	Value string `json:"value"`
}

type RequestResponse struct {
	Page  int `json:"_page,omitempty"`
	Links struct {
		Self struct {
			Href string `json:"href,omitempty"`
		} `json:"self,omitempty"`
	} `json:"_links,omitempty"`
	Embedded struct {
		Leads     []Lead     `json:"leads,omitempty"`
		Unsorted  []Unsorted `json:"unsorted,omitempty"`
		Customers []Customer `json:"customers,omitempty"`
	} `json:"_embedded,omitempty"`
}
