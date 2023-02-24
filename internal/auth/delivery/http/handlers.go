package http

import (
	"depeche/internal/auth/entities"
	"depeche/internal/auth/usecase"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	useCase usecase.UseCase
}

func (handler *Handler) SignUp(context *gin.Context) {

}

func (handler *Handler) SignIn(context *gin.Context) {
	userAuth := &entities.UserAuth{}

	err := context.BindJSON(userAuth)
	if err != nil {
		context.AbortWithError(http.StatusUnauthorized, err)
		return
	}
	fmt.Println(userAuth.Login)
	fmt.Println(userAuth.Password)
	sessionId, err := handler.useCase.AuthenticateUser(userAuth)
	if err != nil {
		context.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	// надо настроить ее нормально
	context.SetCookie("Authentication", sessionId, 100, "/", "/", true, true)
	context.Status(http.StatusOK)
}
