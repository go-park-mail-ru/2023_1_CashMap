package main

import (
	"depeche/authorization_ms/api"
	"depeche/authorization_ms/handler"
	"depeche/authorization_ms/repository/redis"
	"depeche/authorization_ms/service"
	"depeche/configs"
	"depeche/pkg/connector"
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

	srv := grpc.NewServer()
	api.RegisterAuthServiceServer(srv, authHandler)

	// TODO add configs
	listener, err := net.Listen()
}
