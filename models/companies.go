package models

type Company struct {
	ID                int             `json:"id,omitempty"`
	Name              string          `json:"name,omitempty"`
	ResponsibleUserID int             `json:"responsible_user_id,omitempty"`
	CreatedBy         int             `json:"created_by,omitempty"`
	UpdatedBy         int             `json:"updated_by,omitempty"`
	CreatedAt         int             `json:"created_at,omitempty"`
	UpdatedAt         int             `json:"updated_at,omitempty"`
	CustomFieldsValue []CustomFields  `json:"custom_fields_value,omitempty"`
	CompanyEmbedded   CompanyEmbedded `json:"_embedded,omitempty"`
	RequestID         int             `json:"request_id,omitempty"`
	Links             struct {
		Self struct {
			Href string `json:"href,omitempty"`
		} `json:"self,omitempty"`
	} `json:"_links,omitempty"`
}

type CompanyEmbedded struct {
	Tag []Tag `json:"tags,omitempty"`
}
