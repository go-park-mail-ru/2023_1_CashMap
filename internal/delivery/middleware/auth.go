package middleware

import (
	authService "depeche/authorization_ms/service"
	"depeche/pkg/apperror"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthMiddleware struct {
	service authService.Auth
}

func NewAuthMiddleware(authService authService.Auth) *AuthMiddleware {
	return &AuthMiddleware{
		service: authService,
	}
}

func (am *AuthMiddleware) Middleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sessionId, err := ctx.Cookie("session_id")
		if err != nil {
			err = apperror.NoAuth
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  apperror.Errors[err].Code,
				"message": apperror.Errors[err].Message,
			})
			return
		}
		session, err := am.service.CheckSession(sessionId)
		if err != nil {
			err = apperror.NoAuth
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  apperror.Errors[err].Code,
				"message": apperror.Errors[err].Message,
			})
			return
		}

		ctx.Set("email", session.Email)
	}
}
