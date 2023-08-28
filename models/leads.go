package models

type Lead struct {
	// ID сделки
	ID int `json:"id,omitempty"`
	// Название сделки
	Name string `json:"name,omitempty"`
	// Бюджет сделки
	Price int `json:"price,omitempty"`
	// ID статуса, в который добавляется сделка,
	// по-умолчанию – первый этап главной воронки
	StatusID int `json:"status_id,omitempty"`
	// ID воронки, в которую добавляется сделка
	PipelineID int `json:"pipeline_id,omitempty"`
	// ID пользователя, создающий сделку
	CreatedBy int `json:"created_by,omitempty"`
	// ID пользователя, изменяющий сделку
	UpdatedBy int `json:"updated_by,omitempty"`
	// Дата закрытия сделки, передается в Unix Timestamp
	ClosedAt int `json:"closed_at,omitempty"`
	// 	Дата создания сделки, передается в Unix Timestamp
	CreatedAt int `json:"created_at,omitempty"`
	// Дата изменения сделки, передается в Unix Timestamp
	UpdatedAt int `json:"updated_at,omitempty"`
	// ID причины отказа
	//LossReasonID      int    `json:"loss_reason_id"`
	// ID пользователя, ответственного за сделку
	ResponsibleUserID int `json:"responsible_user_id,omitempty"`
	// Массив, содержащий информацию по дополнительным полям,
	// заданным для данной сделки. Поле не является обязательным.
	CustomFieldsValues []CustomFields `json:"custom_fields_values,omitempty"`
	// Metadata - данные о контакте
	Metadata LeadMD `json:"metadata,omitempty"`
	// Данные вложенных сущностей, при создании и редактировании можно
	// передать только теги. Поле не является обязательным
	Embedded LeadEmbedded `json:"_embedded,omitempty"`
	// Links - это ссылка на сделку
	Links struct {
		Self struct {
			Href string `json:"href,omitempty"`
		} `json:"self,omitempty"`
	} `json:"_links,omitempty"`
}

type LeadEmbedded struct {
	Tags      []Tag     `json:"tags,omitempty"`
	Companies []Company `json:"companies,omitempty"`
	Contacts  []Contact `json:"contacts,omitempty"`
}

type LeadComplexResponse struct {
	ID        int      `json:"id,omitempty"`
	ContactID int      `json:"contact_id,omitempty"`
	CompanyID int      `json:"company_id,omitempty"`
	RequestID []string `json:"request_id,omitempty"`
	Merged    bool     `json:"merged,omitempty"`
}

type LeadModifyResponse struct {
	ID        int `json:"id,omitempty"`
	UpdatedAt int `json:"updated_at,omitempty"`
	Links     struct {
		Self struct {
			Href string `json:"href,omitempty"`
		} `json:"self,omitempty"`
	} `json:"_links,omitempty"`
}

type Unsorted struct {
	UID        string `json:"uid,omitempty"`
	SourceUID  string `json:"source_uid,omitempty"`
	SourceName string `json:"source_name,omitempty"`
	Category   string `json:"category,omitempty"`
	PipelineID int    `json:"pipeline_id,omitempty"`
	CreatedAt  int    `json:"created_at,omitempty"`
	Metadata   LeadMD `json:"metadata,omitempty"`
	AccountID  int    `json:"account_id,omitempty"`
	Links      struct {
		Self struct {
			Href string `json:"href,omitempty"`
		} `json:"self,omitempty"`
	} `json:"_links,omitempty"`
	Embedded UnsortedEmbedded `json:"_embedded,omitempty"`
}

type UnsortedSIP struct {
	RequestID  string           `json:"request_id"`
	SourceName string           `json:"source_name,omitempty"`
	SourceUID  string           `json:"source_uid"`
	PipelineID int              `json:"pipeline_id,omitempty"`
	CreatedAt  int              `json:"created_at,omitempty"`
	Embedded   UnsortedEmbedded `json:"_embedded,omitempty"`
	Metadata   SIPmd            `json:"metadata,omitempty"`
}

type UnsortedForms struct {
	RequestID  string           `json:"request_id"`
	SourceName string           `json:"source_name,omitempty"`
	SourceUID  string           `json:"source_uid"`
	PipelineID int              `json:"pipeline_id,omitempty"`
	CreatedAt  int              `json:"created_at,omitempty"`
	Embedded   UnsortedEmbedded `json:"_embedded,omitempty"`
	Metadata   FormMD           `json:"metadata,omitempty"`
}

type UnsortedEmbedded struct {
	Contacts  []Contact `json:"contacts,omitempty"`
	Leads     []Lead    `json:"leads,omitempty"`
	Companies []Company `json:"companies,omitempty"`
}
