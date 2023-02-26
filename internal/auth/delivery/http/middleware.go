package http

import (
	"depeche/internal/auth/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthMiddleware struct {
	authServices usecase.AuthService
}

func NewAuthMiddleware(authService usecase.AuthService) gin.HandlerFunc {
	return (&AuthMiddleware{
		authServices: authService,
	}).Handle
}

// TODO: auth middleware (нужна дополнительная механика для обновления токена при его экспирации)

func (middleware *AuthMiddleware) Handle(context *gin.Context) {
	sessionId, err := context.Cookie("session_id")

	if err != nil {
		context.AbortWithError(401, err)
		return
	}
	exists, err := middleware.authServices.ValidateSession(sessionId)
	if err != nil {
		context.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	if !exists {
		context.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}
