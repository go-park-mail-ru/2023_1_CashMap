package app

import (
	"depeche/authorization_ms/api"
	"depeche/configs"
	"depeche/docs"
	"depeche/internal/delivery/handlers"
	"depeche/internal/delivery/middleware"
	"depeche/internal/delivery/wsPool"
	"depeche/internal/repository/postgres"
	httpserver "depeche/internal/server"
	"depeche/internal/session/client"
	staticDelivery "depeche/internal/static/delivery"
	"depeche/internal/static/repository"
	staticService "depeche/internal/static/service"
	"depeche/internal/usecase/service"
	"depeche/pkg/connector"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

	//redisClient, err := connector.ConnectRedis(&cfg.SessionStorage)
	//if err != nil {
	//	log.Fatal(err)
	//}

	//sessionStorage := redis2.NewRedisStorage(client)
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", cfg.AuthMs.Host, cfg.AuthMs.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	authClient := api.NewAuthServiceClient(conn)
	csrfClient := api.NewCSRFServiceClient(conn)

	userStorage := postgres.NewPostgresUserRepo(db)
	feedStorage := postgres.NewFeedStorage(db)
	postStorage := postgres.NewPostRepository(db)
	messageStorage := postgres.NewMessageRepository(db)
	groupStorage := postgres.NewGroupRepository(db)
	fileStorage := repository.NewFileRepository()

	userService := service.NewUserService(userStorage)
	authorizationService := client.NewAuthService(authClient)
	csrfService := client.NewCSRFService(csrfClient)
	feedService := service.NewFeedService(feedStorage, postStorage)
	fileService := staticService.NewFileUsecase(fileStorage)
	postService := service.NewPostService(postStorage)
	groupService := service.NewGroupService(groupStorage)
	msgService := service.NewMessageService(messageStorage, userStorage)

	staticHandler := staticDelivery.NewFileHandler(fileService)
	userHandler := handlers.NewUserHandler(userService, authorizationService, csrfService)
	feedHandler := handlers.NewFeedHandler(feedService)
	postHandler := handlers.NewPostHandler(postService)
	messageHandler := handlers.NewMessageHandler(msgService)
	groupHandler := handlers.NewGroupHandler(groupService)

	handler := handlers.NewHandler(userHandler, feedHandler, postHandler, staticHandler, messageHandler, groupHandler)

	authMiddleware := middleware.NewAuthMiddleware(authorizationService)

	csrfMiddleware := middleware.NewCSRFMiddleware(csrfService)

	pool := wsPool.NewConnectionPool()

	wsMiddleware := middleware.NewWsMiddleware(pool, msgService)

	router := initRouter(handler, authMiddleware, pool, wsMiddleware, csrfMiddleware)

	initValidator()

	server := httpserver.NewServer(router)
	err = server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func initValidator() {
	// govalidator.TagMap["required"] = govalidator.CustomTypeValidator(func(value interface{}, ctx interface{}) bool {
	//	return repository.IsNil(value)
	// })
}

func initRouter(handler handlers.Handler, authMW *middleware.AuthMiddleware, pool *wsPool.ConnectionPool, wsMiddleware *middleware.WsMiddleware, csrfMiddleware *middleware.CSRFMiddleware) *gin.Engine {
	router := gin.Default()
	// [METRICS]
	m := ginmetrics.GetMonitor()
	m.SetMetricPath("/metrics")
	m.SetSlowTime(5)
	m.SetDuration([]float64{0.1, 0.3, 1.2, 5, 10})
	m.Use(router)
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
	apiEndpointsGroup.Use(csrfMiddleware.Middleware())
	{

		// [FEED]
		apiEndpointsGroup.GET("/feed", handler.GetFeed)

		// [MESSAGE]
		messageEndpointsGroup := apiEndpointsGroup.Group("/im")
		{
			messageEndpointsGroup.GET("/chats", handler.GetChats)
			messageEndpointsGroup.GET("/messages", handler.GetMessagesByChatID)
			messageEndpointsGroup.POST("/chat/create", handler.NewChat)
			messageEndpointsGroup.GET("/chat/check", handler.HasDialog)
			messageEndpointsGroup.POST("/send", handler.Send, wsMiddleware.SendMsg)
		}

		// [POST]
		postEndpoints := apiEndpointsGroup.Group("/posts")
		{
			postEndpoints.GET("/id", handler.GetPostsById)
			postEndpoints.GET("/community", handler.GetPostsByCommunityLink)
			postEndpoints.GET("/profile", handler.GetPostsByUserLink)
			postEndpoints.DELETE("/delete", handler.DeletePost)
			postEndpoints.POST("/create", handler.CreatePost)
			postEndpoints.PATCH("/edit", handler.EditPost)
			postEndpoints.POST("/like/set", handler.LikePost)
			postEndpoints.POST("/like/cancel", handler.CancelPostLike)
		}

		// [STATIC]
		staticEndpointsGroup := router.Group("/static")
		{
			staticEndpointsGroup.POST("/upload", handler.LoadFile)
			staticEndpointsGroup.GET("/download", handler.GetFile)
			staticEndpointsGroup.DELETE("/remove", handler.DeleteFile)
		}

		// [USER]

		userEndpoints := apiEndpointsGroup.Group("/user")
		{
			userEndpoints.GET("/search", handler.GetGlobalSearchResult)
			// [PROFILE]
			profileEndpoints := userEndpoints.Group("/profile")
			{
				profileEndpoints.GET("", handler.Self)
				profileEndpoints.GET("/link/:link", handler.Profile)
				profileEndpoints.PATCH("/edit", handler.EditProfile)
				profileEndpoints.GET("/groups", handler.GetUserGroups)

			}
			userEndpoints.GET("/rand", handler.RandomUsers)
			// [FRIENDS]
			userEndpoints.GET("/friends", handler.Friends)

			// [SUBSCRIBES]
			userEndpoints.GET("/sub", handler.Subscribes)
			userEndpoints.GET("/pending", handler.PendingGroupRequests)

			// [SUBSCRIBE]
			userEndpoints.POST("/sub", handler.Subscribe)

			// [UNSUBSCRIBE]
			userEndpoints.POST("/unsub", handler.Unsubscribe)

			// [REJECT]
			userEndpoints.POST("/reject", handler.Reject)
		}

		// [GROUP]
		groupEndpoints := apiEndpointsGroup.Group("/group")
		{
			// TODO :)
			groupLink := groupEndpoints.Group("/link/:link")
			{
				groupLink.GET("", handler.GetGroup)
				groupLink.PATCH("", handler.UpdateGroup)
				groupLink.DELETE("", handler.DeleteGroup)
				groupLink.GET("/subs", handler.GetSubscribers)
				groupLink.POST("/sub", handler.SubscribeGroup)
				groupLink.POST("/unsub", handler.UnsubscribeGroup)
				groupLink.PATCH("/accept", handler.AcceptRequest)
				groupLink.PUT("/accept", handler.AcceptAllRequests)
				groupLink.PATCH("/decl", handler.DeclineRequest)
				groupLink.GET("/pending", handler.PendingGroupRequests)
			}
			groupEndpoints.GET("/self", handler.GetGroups)
			groupEndpoints.GET("/manage", handler.GetManagedGroups)
			groupEndpoints.GET("/hot", handler.GetPopularGroups)
			groupEndpoints.POST("/create", handler.CreateGroup)

		}
		// [WS]
		apiEndpointsGroup.GET("/ws", pool.Connect)

	}
	return router
}
