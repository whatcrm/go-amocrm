package models

type Task struct {
	ID                int           `json:"id,omitempty"`
	CreatedBy         int           `json:"created_by,omitempty"`
	UpdatedBy         int           `json:"updated_by,omitempty"`
	CreatedAt         int64         `json:"created_at,omitempty"`
	UpdatedAt         int64         `json:"updated_at,omitempty"`
	ResponsibleUserID int           `json:"responsible_user_id,omitempty"`
	GroupID           int           `json:"group_id,omitempty"`
	EntityID          int           `json:"entity_id,omitempty"`
	EntityType        string        `json:"entity_type,omitempty"`
	Duration          int           `json:"duration,omitempty"`
	IsCompleted       bool          `json:"is_completed,omitempty"`
	TaskTypeID        int           `json:"task_type_id,omitempty"`
	Text              string        `json:"text,omitempty"`
	Result            []any         `json:"result,omitempty"`
	CompleteTill      int64         `json:"complete_till,omitempty"`
	AccountID         int           `json:"account_id,omitempty"`
	Links             *LinkResponse `json:"_links,omitempty"`
}
