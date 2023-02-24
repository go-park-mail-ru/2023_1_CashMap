package app

import (
	"depeche/internal/auth/delivery/http"
	httpserver "depeche/internal/server"
	"github.com/gin-gonic/gin"
)

func Run() {
	router := initRouter()

	server := httpserver.NewServer(router)

	err := server.ListenAndServe()
	if err != nil {
		return
	}
}

func initRouter() *gin.Engine {
	router := gin.Default()

	authHandler := http.Handler{}
	authEndpoints := router.Group("/auth")
	{
		authEndpoints.POST("/signin", authHandler.SignIn)
		authEndpoints.POST("/signup", authHandler.SignUp)
	}

	return router
}
