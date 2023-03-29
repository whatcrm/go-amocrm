package models

type Customer struct {
	ID                 int         `json:"id"`
	Name               string      `json:"name"`
	NextPrice          int         `json:"next_price"`
	NextDate           int         `json:"next_date"`
	ResponsibleUserID  int         `json:"responsible_user_id"`
	StatusID           int         `json:"status_id"`
	Periodicity        int         `json:"periodicity"`
	CreatedBy          int         `json:"created_by"`
	UpdatedBy          int         `json:"updated_by"`
	CreatedAt          int         `json:"created_at"`
	UpdatedAt          int         `json:"updated_at"`
	ClosestTaskAt      interface{} `json:"closest_task_at"`
	IsDeleted          bool        `json:"is_deleted"`
	CustomFieldsValues interface{} `json:"custom_fields_values"`
	Ltv                int         `json:"ltv"`
	PurchasesCount     int         `json:"purchases_count"`
	AverageCheck       int         `json:"average_check"`
	AccountID          int         `json:"account_id"`
	Links              struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Embedded struct {
		Segments []struct {
			ID    int `json:"id"`
			Links struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
			} `json:"_links"`
		} `json:"segments"`
		Tags []Tag `json:"tags"`
	} `json:"_embedded"`
}

type CustomersMode struct {
	Mode      string `json:"mode"`
	IsEnabled bool   `json:"is_enabled"`
}
