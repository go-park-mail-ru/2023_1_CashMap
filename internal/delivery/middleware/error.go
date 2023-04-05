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

		err := ctx.Errors[0].Unwrap()
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
}
