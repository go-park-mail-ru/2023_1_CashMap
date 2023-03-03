package middleware

import "github.com/gin-gonic/gin"

func ErrorWrapper() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		if len(ctx.Errors) == 0 {
			ctx.Status(200)
			return
		}

		// что делать с кодом ошибки - его нельзя записать в ctx.Error, поэтому и получить - тоже
		err := ctx.Errors[0]
		ctx.AbortWithStatusJSON(500, err.JSON())
	}
}
