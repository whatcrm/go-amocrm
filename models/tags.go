package models

type Tag struct {
	ID        int    `json:"id,omitempty"`
	Name      string `json:"name"`
	Color     string `json:"color,omitempty"`
	RequestID int    `json:"request_id,omitempty"`
}

//type TagResponse struct {
//	TotalItems int `json:"_total_items"`
//	Embedded   struct {
//		Tags []Tags `json:"tags"`
//	} `json:"_embedded"`
//}

type TagResponse struct {
	TotalItems int `json:"_total_items"`
	Embedded   struct {
		Tags []struct {
			ID        int    `json:"id"`
			Name      string `json:"name"`
			Color     string `json:"color"`
			RequestID string `json:"request_id"`
		} `json:"tags"`
	} `json:"_embedded"`
}
