package main

import (
	"depeche/authorization_ms/api"
	"depeche/authorization_ms/handler"
	"depeche/authorization_ms/metrics"
	"depeche/authorization_ms/repository/redis"
	"depeche/authorization_ms/service"
	"depeche/configs"
	"depeche/pkg/connector"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	var cfg configs.Config
	err := configs.InitCfg(&cfg)
	if err != nil {
		log.Fatal(err)
	}
	client, err := connector.ConnectRedis(&cfg.SessionStorage)
	if err != nil {
		log.Fatal(err)
	}
	authRepository := redis.NewRedisStorage(client)
	authService := service.NewAuthService(authRepository)
	authHandler := handler.NewAuthHandler(authService)

	csrfRepository := redis.NewCSRFStorage(client)
	csrfService := service.NewCSRFService(csrfRepository)
	csrfHandler := handler.NewCSRFHandler(csrfService)

	srv := grpc.NewServer()
	api.RegisterAuthServiceServer(srv, authHandler)
	api.RegisterCSRFServiceServer(srv, csrfHandler)

	prometheus.MustRegister(metrics.RequestCounter)
	prometheus.MustRegister(metrics.DurationHistogram)
	//http.Handle("/metrics", promhttp.Handler())
	//go func() {
	//	err := http.ListenAndServe(":8091", nil)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.AuthMs.Port))
	if err != nil {
		log.Fatal(err)
	}
	if err := srv.Serve(listener); err != nil {
		log.Fatal(err)
	}

}
