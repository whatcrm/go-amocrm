package models

type Chat struct {
	ChatID    string `json:"chat_id"`
	ContactID int    `json:"contact_id"`
	ID        int    `json:"id"`
	RequestID string `json:"request_id"`
}
