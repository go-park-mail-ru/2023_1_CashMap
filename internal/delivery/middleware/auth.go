package middleware

import (
	"depeche/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthMiddleware struct {
	service usecase.User
}

func NewAuthMiddleware(service usecase.User) *AuthMiddleware {
	return &AuthMiddleware{
		service: service,
	}
}

func (am *AuthMiddleware) Middleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// TODO захэндлить ошибки
		sessionId, err := ctx.Cookie("session_id")

		if err != nil {
			ctx.AbortWithError(401, err)
			return
		}
		exists, err := am.service.CheckSession(sessionId)
		if err != nil {
			ctx.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		if !exists {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}
