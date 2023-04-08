package app

import (
	"depeche/configs"
	"depeche/docs"
	"depeche/internal/delivery/handlers"
	"depeche/internal/delivery/middleware"
	"depeche/internal/delivery/wsPool"
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
	"github.com/asaskevich/govalidator"
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

	db := connector.GetSqlxConnector(postgresDefault, cfg.DBMSName)
	// TODO: как обработать ошибку в дефере нормаль?...
	defer db.Close()

	client, err := connector.ConnectRedis(&cfg.SessionStorage)
	if err != nil {
		log.Fatal(err)
	}

	//db, err := connector.ConnectPostgres(&cfg.DB)
	//if err != nil {
	//	log.Fatal(err)
	//}

	sessionStorage := redis.NewRedisStorage(client)

	userStorage := postgres.NewPostgresUserRepo(db)
	feedStorage := storage.NewFeedStorage()
	postStorage := postgres.NewPostRepository(db)

	msgStorage := postgres.NewMessageRepo(db)

	userService := service.NewUserService(userStorage)
	authService := authService.NewAuthService(sessionStorage)
	feedService := service.NewFeedService(feedStorage)
	fileService := staticService.NewFileUsecase()
	postService := service.NewPostService(postStorage)
	msgService := service.NewMessageService(msgStorage, userStorage)

	staticHandler := staticDelivery.NewFileHandler(fileService)
	userHandler := handlers.NewUserHandler(userService, authService)
	feedHandler := handlers.NewFeedHandler(feedService)
	postHandler := handlers.NewPostHandler(postService)
	msgHandler := handlers.NewMessageHandler(msgService)

	handler := handlers.NewHandler(userHandler, feedHandler, postHandler, staticHandler, msgHandler)

	authMiddleware := middleware.NewAuthMiddleware(authService)

	pool := wsPool.NewConnectionPool()

	wsMiddleware := middleware.NewWsMiddleware(pool, msgService)

	router := initRouter(handler, authMiddleware, pool, wsMiddleware)
	server := httpserver.NewServer(router)

	err = server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func initValidator() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func initRouter(handler handlers.Handler, authMW *middleware.AuthMiddleware, pool *wsPool.ConnectionPool, wsMiddleware *middleware.WsMiddleware) *gin.Engine {
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

		// [POST]
		apiEndpointsGroup.GET("/posts/id", handler.GetPostsById)
		apiEndpointsGroup.GET("/posts/community", handler.GetPostsByCommunityLink)
		apiEndpointsGroup.GET("/posts/profile", handler.GetPostsByUserLink)
		apiEndpointsGroup.DELETE("/posts/delete", handler.DeletePost)
		apiEndpointsGroup.POST("/posts/create", handler.CreatePost)
		apiEndpointsGroup.PATCH("/posts/edit", handler.EditPost)

		// [STATIC]
		staticEndpointsGroup := apiEndpointsGroup.Group("/static")
		{
			staticEndpointsGroup.POST("/upload", handler.LoadFile)
			staticEndpointsGroup.GET("/download", handler.GetFile)
			staticEndpointsGroup.DELETE("/remove", handler.DeleteFile)
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
				profileEndpoints.GET("/all", handler.AllUsers)
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

		// [MESSAGE]
		messageEndpoints := apiEndpointsGroup.Group("/im")
		{
			messageEndpoints.POST("/send", handler.Send, wsMiddleware.SendMsg)
		}

		//[WS]
		apiEndpointsGroup.GET("/ws", pool.Connect)

	}
	return router
}
