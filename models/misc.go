package models

type Link struct {
	ToEntityID   int    `json:"to_entity_id"`
	ToEntityType string `json:"to_entity_type"`
	Metadata     struct {
		Quantity    int  `json:"quantity"`
		CatalogID   int  `json:"catalog_id"`
		MainContact bool `json:"main_contact"`
	} `json:"metadata"`
}

type CustomFields struct {
	FieldID   int                  `json:"field_id,omitempty"`
	FieldName string               `json:"field_name,omitempty"`
	FieldCode string               `json:"field_code,omitempty"`
	FieldType string               `json:"field_type,omitempty"`
	Values    []CustomFieldsValues `json:"values,omitempty"`
}

type CustomFieldsValues struct {
	Value    any    `json:"value,omitempty"`
	EnumID   int    `json:"enum_id,omitempty"`
	EnumCode string `json:"enum_code,omitempty"`
}

type RequestResponse struct {
	TotalItems int               `json:"_total_items,omitempty"`
	Page       int               `json:"_page,omitempty"`
	Links      *LinkResponse     `json:"_links,omitempty"`
	Embedded   *ResponseEmbedded `json:"_embedded,omitempty"`
}

type ResponseEmbedded struct {
	Contacts     []Contact     `json:"contacts,omitempty"`
	Companies    []Company     `json:"companies,omitempty"`
	Leads        []Lead        `json:"leads,omitempty"`
	Tags         []Tag         `json:"tags,omitempty"`
	Unsorted     []Unsorted    `json:"unsorted,omitempty"`
	Customers    []Customer    `json:"customers,omitempty"`
	Transactions []Transaction `json:"transactions,omitempty"`
	Chats        []Chat        `json:"chats,omitempty"`
	Links        []Link        `json:"links"`
	Tasks        []Task        `json:"tasks"`
}

type Response struct {
	Title  string `json:"title"`
	Type   string `json:"type"`
	Status int    `json:"status"`
	Detail string `json:"detail"`
}

type MainResponse struct {
	TotalItems int           `json:"_total_items"`
	Links      *LinkResponse `json:"_links"`
	Embedded   struct {
		Contact         []Contact         `json:"contact,omitempty"`
		Unsorted        []Unsorted        `json:"unsorted,omitempty"`
		Transaction     []Transaction     `json:"transactions,omitempty"`
		CustomerStatus  []CustomerStatus  `json:"statuses,omitempty"`
		CustomerSegment []CustomerSegment `json:"segments,omitempty"`
	} `json:"_embedded"`
}

type FallbackResponse struct {
	Errors struct {
		ID      string `json:"id"`
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"errors"`
	Title  string `json:"title"`
	Type   string `json:"type"`
	Status int    `json:"status"`
	Detail string `json:"detail"`
}
