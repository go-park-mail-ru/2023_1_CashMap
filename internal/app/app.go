package app

import (
	"depeche/internal/delivery/handlers"
	"depeche/internal/delivery/middleware"
	storage "depeche/internal/repository/localStorage"
	httpserver "depeche/internal/server"
	session "depeche/internal/session/localStorage"
	"depeche/internal/usecase"

	"github.com/gin-gonic/gin"
)

func Run() {
	data := storage.NewMemoryStorage()
	sessionStorage := session.NewMemorySessionStorage()
	userService := usecase.NewUserService(data, sessionStorage)
	userHandler := handlers.NewUserHandler(userService)
	authMiddleware := middleware.NewAuthMiddleware(userService)
	router := initRouter(userHandler, authMiddleware)
	server := httpserver.NewServer(router)

	err := server.ListenAndServe()
	if err != nil {
		return
	}
}

func initRouter(userHandler *handlers.UserHandler, authMW *middleware.AuthMiddleware) *gin.Engine {
	router := gin.Default()

	// вешаем авторизационную миддлвару на все api
	apiEndpointsGroup := router.Group("/api", authMW.Middleware())

	// тестовый эндпоинт
	apiEndpointsGroup.GET("/test", func(context *gin.Context) {
		context.Writer.WriteString("TEST")
	})

	authEndpointsGroup := router.Group("/auth")
	{
		authEndpointsGroup.POST("/sign-in", userHandler.SignIn)
		authEndpointsGroup.POST("/sign-up", userHandler.SignUp)
	}

	return router
}
