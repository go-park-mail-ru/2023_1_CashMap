package app

import (
	"depeche/authorization_ms/api"
	"depeche/configs"
	httpserver "depeche/internal/server"
	"depeche/pkg/middleware"
	"depeche/static/delivery"
	"depeche/static/repository"
	"depeche/static/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func StartStaticApp() {
	var cfg configs.Config
	err := configs.InitCfg(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	fileRepository := repository.NewFileRepository()
	fileService := service.NewFileUsecase(fileRepository)
	staticHandler := delivery.NewFileHandler(fileService)

	grpcClient, err := grpc.Dial(fmt.Sprintf("%s:%d", cfg.AuthMs.Host, cfg.AuthMs.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	authServiceClient := api.NewAuthServiceClient(grpcClient)
	authService := service.NewAuthService(authServiceClient)

	router := gin.Default()
	router.Use(middleware.CORS())
	router.Use(middleware.ErrorMiddleware())
	router.Use(delivery.AuthMiddleware(authService))
	// [STATIC]
	staticEndpointsGroup := router.Group("/static")
	{
		staticEndpointsGroup.POST("/upload", staticHandler.LoadFile)
		staticEndpointsGroup.GET("/download", staticHandler.GetFile)
		staticEndpointsGroup.DELETE("/remove", staticHandler.DeleteFile)
	}

	server := httpserver.NewServer(router, cfg.StaticMs.Port)
	err = server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		return
	}
}
