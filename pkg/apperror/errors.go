package apperror

import (
	"errors"
	"fmt"
	"net/http"
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
	TooMuchAttachments    = errors.New("more than 10 attachments passed")
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

var (
	StickerNotFound     = errors.New("sticker not found")
	StickerpackNotFound = errors.New("stickerpack not found")
	TooManyStickers     = errors.New("to many sticker passed to 1 stickerpack")
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

type ErrorResponse struct {
	Code    int    `json:"status" example:"400"`
	Message string `json:"message" example:"Невалидный запрос."`
}

var Errors = map[error]ErrorResponse{
	UserNotFound: {
		http.StatusNotFound,
		"Пользователь не найден.",
	},
	PostNotFound: {
		http.StatusNotFound,
		"Запрашиваемый пост не найден",
	},
	CommunityNotFound: {
		http.StatusNotFound,
		"Запрашиваемое сообщества не найдено",
	},
	PostEditingNowAllowed: {
		http.StatusForbidden,
		"Редактирование этого поста не разрешено",
	},
	NoAuth: {
		http.StatusUnauthorized,
		"Нет авторизации.",
	},
	BadRequest: {
		http.StatusBadRequest,
		"Невалидный запрос.",
	},
	IncorrectCredentials: {
		http.StatusUnauthorized,
		"Неверный email или пароль.",
	},
	UserAlreadyExists: {
		http.StatusConflict,
		"Пользователь с таким email уже существует.",
	},
	Forbidden: {
		http.StatusForbidden,
		"Доступ запрещен.",
	},

	InternalServerError: {
		http.StatusInternalServerError,
		"Ошибка сервера :(",
	},
	RepeatedSubscribe: {
		http.StatusConflict,
		"Повторная подписка.",
	},
	TooLargePayload: {
		http.StatusRequestEntityTooLarge,
		"Превышен допустимый размер файла",
	},

	AlreadyLiked: {
		http.StatusConflict,
		"Лайк уже поставлен",
	},

	LikeIsMissing: {
		http.StatusConflict,
		"Нельзя убрать несуществующий лайк",
	},

	IllegalFileExtensionError: {
		http.StatusBadRequest,
		"Недопустимое расширение файла",
	},
	GroupNotFound: {
		http.StatusNotFound,
		"Группа не найдена.",
	},
	GroupAlreadyExists: {
		http.StatusConflict,
		"Сообщество с таким идентификатором уже существует.",
	},
	TooMuchAttachments: {
		http.StatusRequestEntityTooLarge,
		"Нельзя добавлять более 10 вложений.",
	},
	StickerNotFound: {
		http.StatusNotFound,
		"Стикер не найден.",
	},
	StickerpackNotFound: {
		http.StatusNotFound,
		"Стикерпак не найден",
	},
	TooManyStickers: {
		http.StatusRequestEntityTooLarge,
		"Слишком много стикеров для одно стикерпака (не более 20)",
	},
	GroupTitleRequired: {
		http.StatusBadRequest,
		"Имя сообщества - обязательное поле",
	},
}
