package handlers

import (
	"depeche/internal/entities"
	"depeche/internal/usecase"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service     usecase.User
	authService usecase.Auth
}

func NewUserHandler(userService usecase.User, authService usecase.Auth) *UserHandler {
	return &UserHandler{
		service:     userService,
		authService: authService,
	}
}

func (uh *UserHandler) SignIn(ctx *gin.Context) {
	var user entities.User

	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	_, err = uh.service.SignIn(&user)
	if err != nil {
		ctx.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	token, err := uh.authService.Authenticate(&user)
	sessionCookie := &http.Cookie{
		Name:     "session_id",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 1000), // ЗАХАРДКОЖЕНО ВРЕМЯ ЭКСПИРАЦИИ
		MaxAge:   0,
		HttpOnly: true,
		Secure:   false,
		Path:     "/",
	}

	http.SetCookie(ctx.Writer, sessionCookie)
	ctx.Status(http.StatusOK)
}
func (uh *UserHandler) SignUp(ctx *gin.Context) {
	var user entities.User
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// TODO: уточнить код ответа (сейчас летит 400)
	_, err = uh.service.SignUp(&user)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}
}
func (uh *UserHandler) LogOut(ctx *gin.Context) {
	token, err := ctx.Cookie("session_id")
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	err = uh.authService.LogOut(token)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newCookie := &http.Cookie{
		Name:     "session_id",
		Value:    token,
		Expires:  time.Now().Add(-time.Hour), // ЗАХАРДКОЖЕНО ВРЕМЯ ЭКСПИРАЦИИ
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   false,
		Path:     "/",
	}

	http.SetCookie(ctx.Writer, newCookie)
	ctx.Status(http.StatusOK)
}
