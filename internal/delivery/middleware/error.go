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

			ctx.Status(200)
			return
		}

		err := ctx.Errors[0].Unwrap()
		ctx.JSON(Errors[err].statusCode, gin.H{
			"status":  Errors[err].statusCode,
			"message": Errors[err].message,
		})
	}
}

var Errors = map[error]struct {
	statusCode int
	message    string
}{
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
