package models

type Contact struct {
	ID                 int             `json:"id,omitempty"`
	Name               string          `json:"name,omitempty"`
	FirstName          string          `json:"first_name,omitempty"`
	LastName           string          `json:"last_name,omitempty"`
	ResponsibleUserID  int             `json:"responsible_user_id,omitempty"`
	CreatedBy          int             `json:"created_by,omitempty"`
	UpdatedBy          int             `json:"updated_by,omitempty"`
	CreatedAt          int             `json:"created_at,omitempty"`
	UpdatedAt          int             `json:"updated_at,omitempty"`
	CustomFieldsValues []CustomFields  `json:"custom_fields_values,omitempty"`
	Embedded           ContactEmbedded `json:"_embedded,omitempty"`
	Links              struct {
		Self struct {
			Href string `json:"href,omitempty"`
		} `json:"self,omitempty"`
	} `json:"_links,omitempty"`
}

type ContactEmbedded struct {
	//Tags []interface{} `json:"_embedded[tags]"`
	//Tags0     []interface{} `json:"_embedded[tags][0]"`
	//TagsID    int    `json:"_embedded[tags][0][id]"`
	//TagsName  string `json:"_embedded[tags][0][name]"`
	//RequestID string `json:"request_id"`
}
