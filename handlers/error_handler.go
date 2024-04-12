package handlers

import (
	"log"
	"net/http"
)

type ErrorT string

type ErrCode int

type Error struct {
	ErrorCode int
	ErrorType string
	ErrorMsg  string
}

func NewError(err error, errType ErrorT, code ErrCode) *Error {
	return &Error{
		ErrorCode: int(code),
		ErrorType: string(errType),
		ErrorMsg:  err.Error(),
	}
}

func (er Error) HandleError(w http.ResponseWriter, r *http.Request) {

	// reading error from context

	ctx := r.Context()

	err := ctx.Value("err")

	log.Println("err from  context", err)

	// create a appropriate response
	if err != nil {
		if Err, ok := err.(*Error); ok {
			http.Error(w, Err.ErrorMsg, Err.ErrorCode)
			return
		}
	}

}
