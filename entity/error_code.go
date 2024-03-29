package error_code

type ErrorCode string

const (
	Success        ErrorCode = "SUCCESS"
	InvalidRequest ErrorCode = "INVALID_REQUEST"
	InternalError  ErrorCode = "INTERNAL_ERROR"
)

const (
	SuccessMssg    = "success"
	InternalErrMsg = "internal error"
	InvalidErrMsg  = "invalid request "
)
