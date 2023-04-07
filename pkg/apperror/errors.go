package apperror

import "errors"

var (
	NoAuth               = errors.New("user is not authenticated")
	BadRequest           = errors.New("bad request")
	IncorrectCredentials = errors.New("incorrect email or password")
	Forbidden            = errors.New("credentials are not present")
)

var (
	UserAlreadyExists = errors.New("user already exists")
	UserNotFound      = errors.New("user not found")
)

var (
	RepeatedSubscribe = errors.New("already subscribed")
)

var (
	InternalServerError = errors.New("internal server error")
)
