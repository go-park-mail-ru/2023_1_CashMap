package app

import (
	"depeche/configs"
	"depeche/docs"
	"depeche/internal/delivery/handlers"
	"depeche/internal/delivery/middleware"
	storage "depeche/internal/repository/local_storage"
	"depeche/internal/repository/postgres"
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
	client, err := connector.ConnectRedis(&cfg.SessionStorage)
	if err != nil {
		log.Fatal(err)
	}

	db, err := connector.ConnectPostgres(&cfg.DB)
	if err != nil {
		log.Fatal(err)
	}

	userStorage := postgres.NewPostgresUserRepo(db)
	sessionStorage := redis.NewRedisStorage(client)
	feedStorage := storage.NewFeedStorage()

	userService := service.NewUserService(userStorage)
	authService := authService.NewAuthService(sessionStorage)
	feedService := service.NewFeedService(feedStorage)
	fileService := staticService.NewFileUsecase()

	staticHandler := staticDelivery.NewFileHandler(fileService)
	userHandler := handlers.NewUserHandler(userService, authService)
	feedHandler := handlers.NewFeedHandler(feedService)

	handler := handlers.NewHandler(userHandler, feedHandler, nil, staticHandler)

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
	// [MIDDLEWARE]
	router.Use(middleware.CORS())
	router.Use(middleware.ErrorMiddleware())

	// [SWAGGER]
	docs.SwaggerInfo.BasePath = "/api"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// [AUTH]
	authEndpointsGroup := router.Group("/auth")
	{
		authEndpointsGroup.POST("/sign-in", handler.SignIn)
		authEndpointsGroup.POST("/sign-up", handler.SignUp)
		authEndpointsGroup.POST("/logout", handler.LogOut)
		authEndpointsGroup.GET("/check", handler.CheckAuth)
	}

	// [API]
	apiEndpointsGroup := router.Group("/api")
	apiEndpointsGroup.Use(authMW.Middleware())
	{

		// [FEED]
		apiEndpointsGroup.GET("/feed", handler.GetFeed)

		// [STATIC]
		staticEndpointsGroup := apiEndpointsGroup.Group("/static")
		{
			staticEndpointsGroup.POST("/upload", handler.LoadFile)
		}

		// [USER]
		userEndpoints := apiEndpointsGroup.Group("/user")
		{
			// [PROFILE]
			profileEndpoints := userEndpoints.Group("/profile")
			{
				profileEndpoints.GET("", handler.Self)
				profileEndpoints.GET("/:link", handler.Profile)
				profileEndpoints.PATCH("/edit", handler.EditProfile)
			}
			// [FRIENDS]
			userEndpoints.GET("/friends", handler.Friends)

			// [SUBSCRIBES]
			userEndpoints.GET("/sub", handler.Subscribes)

			// [SUBSCRIBE]
			userEndpoints.POST("/sub", handler.Subscribe)

			// [UNSUBSCRIBE]
			userEndpoints.POST("/unsub", handler.Unsubscribe)

			// [REJECT]
			userEndpoints.POST("/reject", handler.Reject)
		}

	}
	return router
}
