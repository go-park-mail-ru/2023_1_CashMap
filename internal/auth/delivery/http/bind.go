package http

import (
	"github.com/gin-gonic/gin"
)

func BindAuthEndpoints(router *gin.Engine, authHandler *AuthHandler) {

	authEndpointsGroup := router.Group("/auth")
	{
		authEndpointsGroup.POST("/sign-in", authHandler.SignIn)
		authEndpointsGroup.POST("/sign-up", authHandler.SignUp)
	}
}
