package models

import error_code "github.com/BoruTamena/UserManagement/entity"

// request
type UserReg struct {
	Id       int    `json:"id"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

// response
type UserResponse struct {
	Data    interface{}          `json:"data"`
	Status  bool                 `json:"status"`
	Code    error_code.ErrorCode `json:"code"`
	Err_msg string               `json:"message"`
}

type RegisterDataResponse struct {
	UserReg
}
