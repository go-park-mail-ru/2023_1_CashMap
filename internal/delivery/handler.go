package delivery

import "github.com/gin-gonic/gin"

type UserHandler interface {
	SignIn(ctx *gin.Context)
	SignUp(ctx *gin.Context)
	LogOut(ctx *gin.Context)
}
