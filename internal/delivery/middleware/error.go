package middleware

import (
	"depeche/pkg/apperror"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		if len(ctx.Errors) == 0 {
			return
		}

		err, ok := ctx.Errors[0].Unwrap().(*apperror.ServerError)
		if !ok {
			// TODO internal error
			return
		}
		//TODO log err.Unwrap
		ctx.JSON(Errors[err.UserErr].Code, gin.H{
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
}
