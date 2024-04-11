package models

type MetaData struct {
	Page    int `json:"page ,omitempty"`
	PerPage int `json:"per_page ,omitempty"`
}

type ResponseData struct {
	Metadata MetaData    `json:"meta_data,omitempty"`
	Data     interface{} `json:"data,omitempty"`
}

type UserRegistrationResponse struct {
	Data interface{} `json:"data,omitempty"`
}
