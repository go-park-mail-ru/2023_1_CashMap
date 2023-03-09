package apperror

import "errors"

var (
	UserNotFound         = errors.New("user not found")
	NoAuth               = errors.New("user is not authenticated")
	BadRequest           = errors.New("bad request")
	IncorrectCredentials = errors.New("incorrect email or password")
	UserAlreadyExists    = errors.New("user already exists")
	Forbidden            = errors.New("credentials are not present")
)

var (
	InternalServerError = errors.New("internal server error")
)
