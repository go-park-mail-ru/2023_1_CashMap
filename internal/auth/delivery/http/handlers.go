package http

import (
	"depeche/internal/auth/entities"
	"depeche/internal/auth/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type AuthHandler struct {
	Service usecase.AuthService
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		Service: usecase.NewAuthService(),
	}
}

func (handler *AuthHandler) SignUp(context *gin.Context) {
	user := &entities.User{}

	err := context.BindJSON(user)
	if err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// TODO: уточнить код ответа (сейчас летит 400)
	err = handler.Service.RegisterUser(user)
	if err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
	}

}

func (handler *AuthHandler) SignIn(context *gin.Context) {
	userAuth := &entities.Credentials{}

	err := context.BindJSON(userAuth)
	if err != nil {
		context.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	sessionId, err := handler.Service.AuthenticateUser(userAuth)
	if err != nil {
		context.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	sessionCookie := &http.Cookie{
		Name:     "session_id",
		Value:    sessionId,
		Expires:  time.Now().Add(time.Hour * 1000), // ЗАХАРДКОЖЕНО ВРЕМЯ ЭКСПИРАЦИИ
		MaxAge:   0,
		HttpOnly: true,
		Secure:   true,
	}

	http.SetCookie(context.Writer, sessionCookie)
	context.Status(http.StatusOK)
}
