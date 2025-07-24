package models

type Company struct {
	ID                 int              `json:"id,omitempty"`
	Name               string           `json:"name,omitempty"`
	ResponsibleUserID  int              `json:"responsible_user_id,omitempty"`
	CreatedBy          int              `json:"created_by,omitempty"`
	UpdatedBy          int              `json:"updated_by,omitempty"`
	CreatedAt          int              `json:"created_at,omitempty"`
	UpdatedAt          int              `json:"updated_at,omitempty"`
	CustomFieldsValues []CustomFields   `json:"custom_fields_values,omitempty"`
	CompanyEmbedded    *CompanyEmbedded `json:"_embedded,omitempty"`
	RequestID          string           `json:"request_id,omitempty"`
	Links              *LinkResponse    `json:"_links,omitempty"`
}

type CompanyEmbedded struct {
	ID              int               `json:"id"`
	Tag             []Tag             `json:"tags,omitempty"`
	Contacts        []Contact         `json:"contacts,omitempty"`
	CatalogElements []CatalogElements `json:"catalog_elements,omitempty"`
}
