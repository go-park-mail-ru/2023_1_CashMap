package middleware

import (
	"depeche/pkg/apperror"
	"depeche/pkg/logs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorMiddleware() gin.HandlerFunc {
	logger := logs.GetLogger()
	return func(ctx *gin.Context) {
		ctx.Next()

		if len(ctx.Errors) == 0 {
			return
		}
		err := ctx.Errors[0].Unwrap()

		serverError, ok := err.(*apperror.ServerError)
		if ok {
			err = serverError.UserErr
			logger.Error(serverError.Unwrap())
		} else {
			logger.Error(err)
		}

		ctx.JSON(Errors[err].Code, gin.H{
			"status":  Errors[err].Code,
			"message": Errors[err].Message,
		})
	}
}

type ErrorResponse struct {
	Code    int    `json:"status" example:"400"`
	Message string `json:"message" example:"Невалидный запрос."`
}

var Errors = map[error]ErrorResponse{
	apperror.UserNotFound: {
		http.StatusNotFound,
		"Пользователь не найден.",
	},
	apperror.PostNotFound: {
		http.StatusNotFound,
		"Запрашиваемый пост не найден",
	},
	apperror.CommunityNotFound: {
		http.StatusNotFound,
		"Запрашиваемое сообщества не найдено",
	},
	apperror.PostEditingNowAllowed: {
		http.StatusForbidden,
		"Редактирование этого поста не разрешено",
	},
	apperror.NoAuth: {
		http.StatusUnauthorized,
		"Нет авторизации.",
	},
	apperror.BadRequest: {
		http.StatusBadRequest,
		"Невалидный запрос.",
	},
	apperror.IncorrectCredentials: {
		http.StatusUnauthorized,
		"Неверный email или пароль.",
	},
	apperror.UserAlreadyExists: {
		http.StatusConflict,
		"Пользователь с таким email уже существует.",
	},
	apperror.Forbidden: {
		http.StatusForbidden,
		"Доступ запрещен.",
	},

	apperror.InternalServerError: {
		http.StatusInternalServerError,
		"Ошибка сервера :(",
	},
	apperror.RepeatedSubscribe: {
		http.StatusConflict,
		"Повторная подписка.",
	},
	apperror.TooLargePayload: {
		http.StatusRequestEntityTooLarge,
		"Превышен допустимый размер файла",
	},

	apperror.AlreadyLiked: {
		http.StatusConflict,
		"Лайк уже поставлен",
	},

	apperror.LikeIsMissing: {
		http.StatusConflict,
		"Нельзя убрать несуществующий лайк",
	},

	apperror.IllegalFileExtensionError: {
		http.StatusBadRequest,
		"Недопустимое расширение файла",
	},
	apperror.GroupNotFound: {
		http.StatusNotFound,
		"Группа не найдена.",
	},
	apperror.GroupAlreadyExists: {
		http.StatusConflict,
		"Сообщество с таким идентификатором уже существует.",
	},
}
