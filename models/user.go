package models

import error_code "github.com/BoruTamena/UserManagement/entity"

// request
type UserReg struct {
	Id int `json:"id"`
	UserLogIn
}

type UserProfile struct {
	UserReg
	img []string
}

type UserLogIn struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type RefreshTokenReq struct {
	RefreshToken string ` json:"refreshtoken" `
}

// response
type UserResponse struct {
	Data    interface{}          `json:"data"`
	Status  bool                 `json:"status"`
	Code    error_code.ErrorCode `json:"code"`
	Err_msg string               `json:"message"`
}

type RegisterDataResponse struct {
	User UserReg
}

type UserListREsponse struct {
	User map[int]UserReg
}
