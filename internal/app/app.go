package app

import (
	"depeche/cmd/app/docs"
	"depeche/internal/delivery/handlers"
	"depeche/internal/delivery/middleware"
	storage "depeche/internal/repository/local_storage"
	httpserver "depeche/internal/server"
	sessionStorage "depeche/internal/session/repository/local_storage"
	authService "depeche/internal/session/service"
	"depeche/internal/usecase/service"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Run() {
	userStorage := storage.NewUserStorage()
	sessionStorage := sessionStorage.NewMemorySessionStorage()
	feedStorage := storage.NewFeedStorage()

	userService := service.NewUserService(userStorage)
	authService := authService.NewAuthService(sessionStorage)
	feedService := service.NewFeedService(feedStorage)

	userHandler := handlers.NewUserHandler(userService, authService)
	feedHandler := handlers.NewFeedHandler(feedService)

	handler := handlers.NewHandler(userHandler, feedHandler, nil)

	authMiddleware := middleware.NewAuthMiddleware(authService)

	router := initRouter(handler, authMiddleware)
	server := httpserver.NewServer(router)

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func initRouter(handler handlers.Handler, authMW *middleware.AuthMiddleware) *gin.Engine {
	router := gin.Default()

	// // swagger api route
	docs.SwaggerInfo.BasePath = "/api"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// вешаем авторизационную миддлвару на все api
	apiEndpointsGroup := router.Group("/api", authMW.Middleware())

	apiEndpointsGroup.GET("/feed", handler.GetFeed)

	authEndpointsGroup := router.Group("/auth")
	{
		authEndpointsGroup.POST("/sign-in", handler.SignIn)
		authEndpointsGroup.POST("/sign-up", handler.SignUp)
		authEndpointsGroup.POST("/logout", handler.LogOut)
		authEndpointsGroup.GET("/check", handler.CheckAuth)
	}

	return router
}
