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
	Data []models.User
}

// Other code here...

func NewUserDb() *UserDb {
	return &UserDb{
		Data: make([]models.User, 0),
	}
}

func (user *UserDb) Insert(userReq models.User) *models.UserRegistrationResponse {
	// Append user to userDb
	user.Data = append(user.Data, userReq)

	return CreateRegisterResponse(user.Data)
}

func (user *UserDb) Select() *models.ResponseData {
	// Preparing the response data
	resData := models.UserListResponse{
		User: user.Data,
	}

	return CreateSuccessResponse(models.MetaData{}, resData)
}

func (user *UserDb) SelectPagination(limit int, offset int) []models.User {

	resData := user.Data[offset:limit]

	return resData

}

func CreateFaildResponse(code error_code.ErrorCode, message string) *models.UserResponse {

	return &models.UserResponse{

		Status:  false,
		Code:    code,
		Err_msg: message,
	}
}

func CreateSuccessResponse(metadata models.MetaData, Data interface{}) *models.ResponseData {
	return &models.ResponseData{
		Metadata: metadata,
		Data:     Data,
	}
}

func CreateRegisterResponse(data interface{}) *models.UserRegistrationResponse {
	return &models.UserRegistrationResponse{
		Data: data,
	}
}
