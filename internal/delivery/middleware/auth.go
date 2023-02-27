package middleware

import (
	"depeche/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthMiddleware struct {
	service usecase.Auth
}

func NewAuthMiddleware(authService usecase.Auth) *AuthMiddleware {
	return &AuthMiddleware{
		service: authService,
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
		session, err := am.service.CheckSession(sessionId)
		if err != nil {
			ctx.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		if session == nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		ctx.Set("email", session.Email)
	}
}
