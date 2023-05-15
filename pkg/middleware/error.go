package middleware

import (
	"depeche/pkg/apperror"
	"depeche/pkg/logs"
	"github.com/gin-gonic/gin"
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

		ctx.JSON(apperror.Errors[err].Code, gin.H{
			"status":  apperror.Errors[err].Code,
			"message": apperror.Errors[err].Message,
		})
	}
}
