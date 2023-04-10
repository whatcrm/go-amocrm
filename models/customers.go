package models

type Customer struct {
	ID                 int         `json:"id,omitempty"`
	RequestID          string      `json:"request_id,omitempty"`
	Name               string      `json:"name,omitempty"`
	NextPrice          int         `json:"next_price,omitempty"`
	NextDate           int         `json:"next_date,omitempty"`
	ResponsibleUserID  int         `json:"responsible_user_idc,omitempty'"`
	StatusID           int         `json:"status_id,omitempty"`
	Periodicity        int         `json:"periodicity,omitempty"`
	CreatedBy          int         `json:"created_by,omitempty"`
	UpdatedBy          int         `json:"updated_by,omitempty"`
	CreatedAt          int         `json:"created_at,omitempty"`
	UpdatedAt          int         `json:"updated_at,omitempty"`
	ClosestTaskAt      interface{} `json:"closest_task_at,omitempty"`
	IsDeleted          bool        `json:"is_deleted,omitempty"`
	CustomFieldsValues interface{} `json:"custom_fields_values,omitempty"`
	Ltv                int         `json:"ltv,omitempty"`
	PurchasesCount     int         `json:"purchases_count,omitempty"`
	AverageCheck       int         `json:"average_check,omitempty"`
	AccountID          int         `json:"account_id,omitempty"`
	Links              struct {
		Self struct {
			Href string `json:"href,omitempty"`
		} `json:"self,omitempty"`
	} `json:"_links,omitempty"`
	Embedded CustomersEmbedded `json:"_embedded,omitempty"`
}

type CustomersEmbedded struct {
	Segments        []CustomerSegment `json:"segments,omitempty"`
	Tags            []Tag             `json:"tags,omitempty"`
	CatalogElements []CatalogElements `json:"catalog_elements"`
	Contacts        []Contact         `json:"contacts,omitempty"`
	Companies       []Company         `json:"companies,omitempty"`
}

type CatalogElements struct {
	ID        int         `json:"id,omitempty"`
	Metadata  interface{} `json:"metadata"`
	Quantity  int         `json:"quantity"`
	CatalogID int         `json:"catalog_id"`
	PriceID   int         `json:"price_id"`
}

type CustomersMode struct {
	Mode      string `json:"mode"`
	IsEnabled bool   `json:"is_enabled"`
}

type CustomerResponse struct {
	ID        int           `json:"id"`
	Name      string        `json:"name"`
	StatusID  int           `json:"status_id"`
	CreatedBy int           `json:"created_by"`
	UpdatedBy int           `json:"updated_by"`
	CreatedAt int           `json:"created_at"`
	UpdatedAt int           `json:"updated_at"`
	AccountID int           `json:"account_id"`
	RequestID string        `json:"request_id"`
	Links     *LinkResponse `json:"_links,omitempty"`
}

type LinkResponse struct {
	Self struct {
		Href string `json:"href,omitempty"`
	} `json:"self,omitempty"`
}
