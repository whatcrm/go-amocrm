package models

type LeadMD struct {
	From       string `json:"from,omitempty"`
	To         string `json:"to,omitempty"`
	ReceivedAt int    `json:"received_at,omitempty"`
	Service    string `json:"service,omitempty"`
	Client     struct {
		Name   string `json:"name,omitempty"`
		Avatar string `json:"avatar,omitempty"`
	} `json:"client,omitempty"`
	Origin struct {
		ChatID     string      `json:"chat_id,omitempty"`
		Ref        interface{} `json:"ref,omitempty"`
		VisitorUID interface{} `json:"visitor_uid,omitempty"`
	} `json:"origin,omitempty"`
	LastMessageText int    `json:"last_message_text,omitempty"`
	SourceName      string `json:"source_name,omitempty"`
}

type SIPmd struct {
	IsCallEventNeeded bool   `json:"is_call_event_needed"`
	Uniq              string `json:"uniq"`
	Duration          int    `json:"duration"`
	ServiceCode       string `json:"service_code"`
	Link              string `json:"link"`
	Phone             int64  `json:"phone"`
	CalledAt          int    `json:"called_at"`
	From              string `json:"from"`
}

type FormMD struct {
	IP         string `json:"ip"`
	FormID     string `json:"form_id"`
	FormSentAt int    `json:"form_sent_at"`
	FormName   string `json:"form_name"`
	FormPage   string `json:"form_page"`
	Referer    string `json:"referer"`
}
