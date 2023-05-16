package apperror

import (
	"errors"
	"fmt"
)

var (
	NoAuth               = errors.New("user is not authenticated")
	BadRequest           = errors.New("bad request")
	IncorrectCredentials = errors.New("incorrect email or password")
	Forbidden            = errors.New("credentials are not present")
	TooLargePayload      = errors.New("file is too large")
)

var (
	UserAlreadyExists = errors.New("user already exists")
	UserNotFound      = errors.New("user not found")
	PostNotFound      = errors.New("post with given id not found")
	CommunityNotFound = errors.New("community not found")
)

var (
	PostEditingNowAllowed = errors.New("post editing is not allowed")
	AlreadyLiked          = errors.New("like has already set")
	LikeIsMissing         = errors.New("like on this post doesn't exists")
)

var (
	RepeatedSubscribe = errors.New("already subscribed")
)

var (
	InternalServerError = errors.New("internal server error")
)

var (
	IllegalFileExtensionError = errors.New("illegal file extension")
)

var (
	GroupNotFound      = errors.New("group not found")
	UnableToLoadAvatar = errors.New("unable to load avatar")
	GroupAlreadyExists = errors.New("link is already in use")
	GroupTitleRequired = errors.New("group title required")
)

type ServerError struct {
	UserErr     error
	internalErr error
}

func NewServerError(userErr error, internalErr error) *ServerError {
	return &ServerError{
		UserErr:     userErr,
		internalErr: internalErr,
	}
}

func NewBadRequest() *ServerError {
	return &ServerError{
		UserErr:     BadRequest,
		internalErr: nil,
	}
}

func (se *ServerError) Unwrap() error {
	return se.internalErr
}

func (se *ServerError) Error() string {
	return fmt.Sprintf("error [%s] internal error:  %s", se.UserErr, se.internalErr)
}
