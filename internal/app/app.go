package app

import (
	"depeche/internal/delivery/handlers"
	"depeche/internal/delivery/middleware"
	storage "depeche/internal/repository/localStorage"
	httpserver "depeche/internal/server"
	session "depeche/internal/session/localStorage"
	"depeche/internal/usecase/service"

	"github.com/gin-gonic/gin"
)

func Run() {
	userStorage := storage.NewMemoryStorage()
	sessionStorage := session.NewMemorySessionStorage()
	userService := service.NewUserService(userStorage)
	authService := service.NewAuthService(sessionStorage)
	userHandler := handlers.NewUserHandler(userService, authService)
	authMiddleware := middleware.NewAuthMiddleware(authService)
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
		id, _ := context.Cookie("session_id")
		context.Writer.WriteString("TEST: your session_id=" + id)
	})

	authEndpointsGroup := router.Group("/auth")
	{
		authEndpointsGroup.POST("/sign-in", userHandler.SignIn)
		authEndpointsGroup.POST("/sign-up", userHandler.SignUp)
	}

	return router
}
