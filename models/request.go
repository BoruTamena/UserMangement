package models

type Phone string

// signup request model
type User struct {
	Id int `json:"id" `
	UserLogIn
	PhoneNumber string   `json:"phonenumber ,omitempty"`
	Email       string   `json:"email, omitempty" `
	Image       []string `json:"image, omitempty" `
	Address     string   `json:"address, omitempty" `
}

// singin request model
type UserLogin struct {
	UserName string `json:"username, omitempty" `
	Password string `json:"password, omitempty" `
}

// pagination request
type PaginationReq struct {
	Page     string
	PageSize string
}
