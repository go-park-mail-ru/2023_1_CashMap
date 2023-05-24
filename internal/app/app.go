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
	"depeche/internal/usecase/service"
	"depeche/pkg/connector"
	middleware2 "depeche/pkg/middleware"
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
	stickerStorage := postgres.NewStickerRepository(db)
	commentStorage := postgres.NewCommentStorage(db)

	userService := service.NewUserService(userStorage)
	authorizationService := client.NewAuthService(authClient)
	csrfService := client.NewCSRFService(csrfClient)
	feedService := service.NewFeedService(feedStorage, postStorage)
	postService := service.NewPostService(postStorage)
	groupService := service.NewGroupService(groupStorage)
	msgService := service.NewMessageService(messageStorage, userStorage)
	stickerService := service.NewStickerService(stickerStorage)
	commentService := service.NewCommentService(commentStorage)

	userHandler := handlers.NewUserHandler(userService, authorizationService, csrfService)
	feedHandler := handlers.NewFeedHandler(feedService)
	postHandler := handlers.NewPostHandler(postService)
	messageHandler := handlers.NewMessageHandler(msgService)
	groupHandler := handlers.NewGroupHandler(groupService)
	stickerHandler := handlers.NewStickerHandler(stickerService)
	commentHandler := handlers.NewCommentHandler(commentService)
	handler := handlers.NewHandler(
		userHandler, feedHandler,
		postHandler,
		messageHandler, groupHandler,
		stickerHandler,
		commentHandler)

	authMiddleware := middleware.NewAuthMiddleware(authorizationService)

	csrfMiddleware := middleware.NewCSRFMiddleware(csrfService)

	pool := wsPool.NewConnectionPool()

	wsMiddleware := middleware.NewWsMiddleware(pool, msgService)

	router := initRouter(handler, authMiddleware, pool, wsMiddleware, csrfMiddleware)

	initValidator()

	server := httpserver.NewServer(router, cfg.Port)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
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
	m.SetDuration([]float64{0.02, 0.08, 0.1, 0.2, 0.5})
	m.Use(router)
	// [MIDDLEWARE]
	router.Use(middleware2.CORS())
	router.Use(middleware2.ErrorMiddleware())

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

		// [COMMENT]
		commentEndpointsGroup := apiEndpointsGroup.Group("comment")
		{
			commentEndpointsGroup.GET("/:id", handler.GetCommentById)
			commentEndpointsGroup.GET("/post/:id", handler.GetCommentByPostId)
			commentEndpointsGroup.POST("/create", handler.CreateComment)
			commentEndpointsGroup.PATCH("/edit", handler.EditComment)
			commentEndpointsGroup.POST("/delete/:id", handler.DeleteComment)
		}

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

		// [USER]

		userEndpoints := apiEndpointsGroup.Group("/user")
		{
			userEndpoints.GET("/search", handler.GetGlobalSearchResult)
			userEndpoints.GET("/status", handler.UserStatus)
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

		// [STICKERS]
		stickerEndpoints := apiEndpointsGroup.Group("/sticker")
		{
			stickerEndpoints.GET("/", handler.GetStickerById)
			stickerPackEndpoints := stickerEndpoints.Group("/pack")
			{
				stickerPackEndpoints.GET("/", handler.GetStickerPack)
				stickerPackEndpoints.GET("/info", handler.GetStickerPackInfo)
				stickerPackEndpoints.GET("/hot", handler.GetNewStickerPacks)
				stickerPackEndpoints.GET("/author", handler.GetStickerPacksByAuthor)
				stickerPackEndpoints.GET("/self", handler.GetUserStickerPacks)

				stickerPackEndpoints.POST("/create", handler.UploadStickerPack)
				stickerPackEndpoints.POST("/add", handler.AddStickerPack)
			}
		}
		// [WS]
		apiEndpointsGroup.GET("/ws", pool.Connect)

	}
	return router
}
