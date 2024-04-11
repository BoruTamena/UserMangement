package error_code

type ErrorCode string
type ErrorType string

const (
	Success        ErrorCode = "SUCCESS"
	InvalidRequest ErrorCode = "INVALID_REQUEST"
	InternalError  ErrorCode = "INTERNAL_ERROR"
)

const (
	UnableToSave         ErrorType = "UNABLE_TO_SAVE"
	UnableToFindResource ErrorType = "UNABLE_TO_FIND_RESOURCE"
	UnableToRead         ErrorType = " UNABLE_TO_READ"
	Unauthorized         ErrorType = "UNAUTHORIZED"
)

const (
	SuccessMssg    = "success"
	InternalErrMsg = "internal error"
	InvalidErrMsg  = "invalid request "
)
