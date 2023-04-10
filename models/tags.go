package models

type Tag struct {
	ID        int    `json:"id,omitempty"`
	Name      string `json:"name"`
	Color     string `json:"color,omitempty"`
	RequestID int    `json:"request_id,omitempty"`
}
