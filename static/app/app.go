package app

import (
	"depeche/authorization_ms/api"
	"depeche/configs"
	httpserver "depeche/internal/server"
	"depeche/pkg/middleware"
	"depeche/static/delivery"
	"depeche/static/repository"
	"depeche/static/service"
	static_api "depeche/static/static_grpc"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
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

	grpcClient, err := grpc.Dial(fmt.Sprintf("%s:%d", "auth", cfg.AuthMs.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	colorService := service.NewColorService()
	colorHandler := delivery.NewColorHandler(colorService)
	srv := grpc.NewServer()
	static_api.RegisterColorServiceServer(srv, colorHandler)
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.StaticMs.ColorPort))
	if err != nil {
		log.Fatal(err)
	}
	if err := srv.Serve(listener); err != nil {
		log.Fatal(err)
	}

	authServiceClient := api.NewAuthServiceClient(grpcClient)
	authService := service.NewAuthService(authServiceClient)

	router := gin.Default()
	router.Use(middleware.CORS())
	router.Use(middleware.ErrorMiddleware())
	router.Use(delivery.AuthMiddleware(authService))

	// [METRICS]
	metricRouter := gin.Default()
	m := ginmetrics.GetMonitor()
	m.SetMetricPath("/metrics")
	m.SetSlowTime(5)
	m.SetDuration([]float64{0.02, 0.08, 0.1, 0.2, 0.5})
	m.UseWithoutExposingEndpoint(router)
	m.Expose(metricRouter)
	// [STATIC]
	staticEndpointsGroup := router.Group("/static-service")
	{
		staticEndpointsGroup.POST("/upload", staticHandler.LoadFile)
		staticEndpointsGroup.GET("/download", staticHandler.GetFile)
		staticEndpointsGroup.DELETE("/remove", staticHandler.DeleteFile)
	}

	go func() {
		err := metricRouter.Run(":8092")
		if err != nil {
			log.Fatal(err)
		}
	}()
	server := httpserver.NewServer(router, cfg.StaticMs.Port)
	err = server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		return
	}

}
