package models

type Transaction struct {
	ID          int    `json:"id,omitempty"`
	Price       int    `json:"price,omitempty"`
	Comment     string `json:"comment,omitempty"`
	CompletedAt int    `json:"completed_at,omitempty"`
	CustomerID  int    `json:"customer_id,omitempty"`
	CreatedBy   int    `json:"created_by,omitempty"`
	UpdatedBy   int    `json:"updated_by,omitempty"`
	CreatedAt   int    `json:"created_at,omitempty"`
	UpdatedAt   int    `json:"updated_at,omitempty"`
	IsDeleted   bool   `json:"is_deleted,omitempty"`
	AccountID   int    `json:"account_id,omitempty"`
	Embedded    struct {
		Customer Customer `json:"customer,omitempty"`
	} `json:"_embedded,omitempty"`
}

type BonusPoints struct {
	// Both values are not allowed
	Earn   int `json:"earn,omitempty"`
	Redeem int `json:"redeem,omitempty"`
}

type Points struct {
	BonusPoints int `json:"bonus_points"`
}
