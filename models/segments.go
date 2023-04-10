package models

type CustomerSegment struct {
	ID                          int            `json:"id,omitempty"`
	CreatedAt                   int            `json:"created_at,omitempty"`
	UpdatedAt                   int            `json:"updated_at,omitempty"`
	AccountID                   int            `json:"account_id,omitempty"`
	Name                        string         `json:"name,omitempty"`
	Color                       string         `json:"color,omitempty"`
	AvailableProductsPriceTypes []int          `json:"available_products_price_types,omitempty"`
	BonusPointsConversionRate   float64        `json:"bonus_points_conversion_rate,omitempty"`
	BonusPointsMaxDiscount      int            `json:"bonus_points_max_discount,omitempty"`
	CustomersCount              int            `json:"customers_count,omitempty"`
	CustomFieldsValues          []CustomFields `json:"custom_fields_values,omitempty"`
	Links                       *LinkResponse  `json:"_links,omitempty,omitempty"`
}
