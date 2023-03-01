package app

import (
	"depeche/internal/delivery"
	"depeche/internal/delivery/handlers"
	"depeche/internal/delivery/middleware"
	storage "depeche/internal/repository/local_storage"
	httpserver "depeche/internal/server"
	session "depeche/internal/session/localStorage"
	"depeche/internal/usecase/service"

	"github.com/gin-gonic/gin"
)

func Run() {
	userStorage := storage.NewUserStorage()
	sessionStorage := session.NewMemorySessionStorage()
	feedStorage := storage.NewFeedStorage()

	userService := service.NewUserService(userStorage)
	authService := service.NewAuthService(sessionStorage)
	feedService := service.NewFeedService(feedStorage)

	userHandler := handlers.NewUserHandler(userService, authService)
	feedHandler := handlers.NewFeedHandler(feedService)

	handler := handlers.NewHandler(userHandler, feedHandler, nil)

	authMiddleware := middleware.NewAuthMiddleware(authService)

	router := initRouter(handler, authMiddleware)
	server := httpserver.NewServer(router)

	err := server.ListenAndServe()
	if err != nil {
		return
	}
}

func initRouter(handler delivery.Handler, authMW *middleware.AuthMiddleware) *gin.Engine {
	router := gin.Default()

	// вешаем авторизационную миддлвару на все api
	apiEndpointsGroup := router.Group("/api", authMW.Middleware())

	apiEndpointsGroup.GET("/feed", handler.GetPosts)

	// тестовый эндпоинт
	apiEndpointsGroup.GET("/test", func(context *gin.Context) {
		id, _ := context.Cookie("session_id")
		context.Writer.WriteString("TEST: your session_id=" + id)
	})

	authEndpointsGroup := router.Group("/auth")
	{
		authEndpointsGroup.POST("/sign-in", handler.SignIn)
		authEndpointsGroup.POST("/sign-up", handler.SignUp)
		authEndpointsGroup.POST("/logout", handler.LogOut)
	}

	return router
}
