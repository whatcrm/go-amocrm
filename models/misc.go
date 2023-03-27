package models

type CustomFields struct {
	//FieldID   int                  `json:"field_id"`
	FieldCode string               `json:"field_code"`
	Values    []CustomFieldsValues `json:"values"`
}

type CustomFieldsValues struct {
	Value string `json:"value"`
}
