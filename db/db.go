package db

import (
	error_code "github.com/BoruTamena/UserManagement/entity"
	"github.com/BoruTamena/UserManagement/models"
)

const (
	InvalidUserNameErrMsg = "Invalid User Name "
	InvalidPasswordMsg    = "Invalid PassWord "
)

type UserDb struct {
	Data map[int]models.UserReg
}

func NewUserDb() *UserDb {
	return &UserDb{
		Data: make(map[int]models.UserReg),
	}
}

func (user *UserDb) Insert(user_req models.UserReg) *models.UserResponse {

	if len(user_req.UserName) == 0 {

		return CreateFaildResponse(error_code.InvalidRequest, InvalidUserNameErrMsg)
	}

	if len(user_req.Password) == 0 {
		return CreateFaildResponse(error_code.InvalidRequest, InvalidPasswordMsg)
	}

	// hashing password

	// append user to userdb

	user.Data[user_req.Id] = user_req

	// creating response data
	res_data := models.RegisterDataResponse{
		User: user_req,
	}

	return CreateSuccessResponse(res_data)

}

func (user *UserDb) Select() *models.UserResponse {

	// prepareing the response data
	res_data := models.UserListREsponse{
		User: user.Data,
	}

	return CreateSuccessResponse(res_data)

}

func CreateFaildResponse(code error_code.ErrorCode, message string) *models.UserResponse {

	return &models.UserResponse{

		Status:  false,
		Code:    code,
		Err_msg: message,
	}
}

func CreateSuccessResponse(data interface{}) *models.UserResponse {
	return &models.UserResponse{
		Data:    data,
		Status:  true,
		Code:    error_code.Success,
		Err_msg: error_code.SuccessMssg,
	}
}
