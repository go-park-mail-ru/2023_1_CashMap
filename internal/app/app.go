package app

import (
	"depeche/configs"
	"depeche/docs"
	"depeche/internal/delivery/handlers"
	"depeche/internal/delivery/middleware"
	storage "depeche/internal/repository/local_storage"
	"depeche/internal/repository/pstgrs"
	httpserver "depeche/internal/server"
	"depeche/internal/session/repository/redis"
	authService "depeche/internal/session/service"
	staticDelivery "depeche/internal/static/delivery"
	staticService "depeche/internal/static/service"
	"depeche/internal/usecase/service"
	"depeche/pkg/connector"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
)

func Run() {
	var cfg configs.Config
	err := configs.InitCfg(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	postgresDefault, err := connector.GetPostgresConnector(&cfg.DB)
	if err != nil {
		log.Fatal(err)
	}

	sqlxConnector := connector.GetSqlxConnector(postgresDefault, cfg.DBMSName)
	// TODO: как обработать ошибку в дефере нормаль?...
	defer sqlxConnector.Close()

	client, err := connector.ConnectRedis(&cfg.SessionStorage)
	if err != nil {
		log.Fatal(err)
	}

	_, err = connector.ConnectPostgres(&cfg.DB)
	if err != nil {
		log.Fatal(err)
	}

	userStorage := storage.NewUserStorage()
	sessionStorage := redis.NewRedisStorage(client)
	feedStorage := storage.NewFeedStorage()
	postStorage := pstgrs.NewPostRepository(sqlxConnector)

	userService := service.NewUserService(userStorage)
	authService := authService.NewAuthService(sessionStorage)
	feedService := service.NewFeedService(feedStorage)
	fileService := staticService.NewFileUsecase()
	postService := service.NewPostService(postStorage)

	staticHandler := staticDelivery.NewFileHandler(fileService)
	userHandler := handlers.NewUserHandler(userService, authService)
	feedHandler := handlers.NewFeedHandler(feedService)
	postHandler := handlers.NewPostHandler(postService)

	handler := handlers.NewHandler(userHandler, feedHandler, postHandler, staticHandler)

	authMiddleware := middleware.NewAuthMiddleware(authService)

	router := initRouter(handler, authMiddleware)
	server := httpserver.NewServer(router)

	err = server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func initRouter(handler handlers.Handler, authMW *middleware.AuthMiddleware) *gin.Engine {
	router := gin.Default()
	router.Use(middleware.CORS())
	router.Use(middleware.ErrorMiddleware())
	// // swagger api route
	docs.SwaggerInfo.BasePath = "/api"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// вешаем авторизационную миддлвару на все api
	//apiEndpointsGroup := router.Group("/api", authMW.Middleware())
	apiEndpointsGroup := router.Group("/api")

	apiEndpointsGroup.GET("/feed", handler.GetFeed)

	staticEndpointsGroup := apiEndpointsGroup.Group("/static")
	{
		staticEndpointsGroup.POST("/upload", handler.LoadFile)
		staticEndpointsGroup.GET("/download", handler.GetFile)
		staticEndpointsGroup.DELETE("/remove", handler.DeleteFile)
	}

	authEndpointsGroup := router.Group("/auth")
	{
		authEndpointsGroup.POST("/sign-in", handler.SignIn)
		authEndpointsGroup.POST("/sign-up", handler.SignUp)
		authEndpointsGroup.POST("/logout", handler.LogOut)
		authEndpointsGroup.GET("/check", handler.CheckAuth)
	}

	return router
}
