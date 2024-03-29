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
	Data []models.UserReg
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

	user.Data = append(user.Data, user_req)

	// creating response data
	res_data := models.RegisterDataResponse{
		UserReg: user_req,
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

func CreateSuccessResponse(data models.RegisterDataResponse) *models.UserResponse {
	return &models.UserResponse{
		Data:    data,
		Status:  true,
		Code:    error_code.Success,
		Err_msg: error_code.SuccessMssg,
	}
}
