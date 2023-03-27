package apperror

import (
	"net/http"
)

var (
	UserNotFound         = NewServerError(http.StatusNotFound, "user not found")
	NoAuth               = NewServerError(http.StatusUnauthorized, "user is not authenticated")
	BadRequest           = NewServerError(http.StatusBadRequest, "bad request")
	IncorrectCredentials = NewServerError(http.StatusUnauthorized, "incorrect email or password")
	UserAlreadyExists    = NewServerError(http.StatusConflict, "user already exists")
	Forbidden            = NewServerError(http.StatusForbidden, "credentials are not present")
	TooLargePayload      = NewServerError(http.StatusRequestEntityTooLarge, "file is too large")
)

var (
	InternalServerError = NewServerError(http.StatusInternalServerError, "internal server error")
)

type ServerError struct {
	Code    int
	Message string
}

func (serverError *ServerError) Error() string {
	return serverError.Message
}

func NewServerError(code int, message string) *ServerError {
	return &ServerError{
		code,
		message,
	}
}
