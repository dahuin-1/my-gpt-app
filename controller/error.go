package controller

import "fmt"

type ErrorType struct {
	ErrorCode string `json:"error_code"`
	Message   string `json:"message"`
	Status    string `json:"status"`
}

func (e *ErrorType) Error() string {
	return fmt.Sprintf(`{"error_code":"%s", "message":"%s"}`, e.ErrorCode, e.Message)
}

var (
	ErrInternal     = ErrorType{"0000", "internal server error", "InternalServerError"}
	ErrUserNotFound = ErrorType{"1000", "user does not exist", "Unauthorized"}
)
