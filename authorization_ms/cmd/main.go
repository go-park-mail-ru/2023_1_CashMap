package main

import (
	"depeche/authorization_ms/api"
	"depeche/authorization_ms/handler"
	"depeche/authorization_ms/repository/redis"
	"depeche/authorization_ms/service"
	"depeche/configs"
	"depeche/pkg/connector"
	"fmt"
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

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.AuthMs.Port))
	if err != nil {
		log.Fatal(err)
	}
	if err := srv.Serve(listener); err != nil {
		log.Fatal(err)
	}

}
