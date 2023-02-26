package handlers

import (
	"depeche/internal/entities"
	"depeche/internal/usecase"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service usecase.User
}

func NewUserHandler(userService usecase.User) *UserHandler {
	return &UserHandler{
		service: userService,
	}
}

func (uh *UserHandler) SignIn(ctx *gin.Context) {
	var user entities.User

	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	sessionId, err := uh.service.SignIn(&user)
	if err != nil {
		ctx.AbortWithError(http.StatusUnauthorized, err)
		return
	}
	sessionCookie := &http.Cookie{
		Name:     "session_id",
		Value:    sessionId,
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

}
