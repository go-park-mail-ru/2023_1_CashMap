package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowCredentials = true
	corsConfig.AllowMethods = []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Content-Type", "Origin", "X-Csrf-Token",
		"Access-Control-Expose-Headers", "Connection"}
	corsConfig.AllowOrigins = []string{
		"http://95.163.212.121:8000",
		"http://95.163.212.121:8080",
		"http://95.163.212.121:80",
		"http://127.0.0.1:8000",
		"http://127.0.0.1:8080",
		"https://depeche.su",
		"http://95.163.212.121:8082",
		"http://95.163.212.121:443"}
	corsConfig.ExposeHeaders = []string{"X-Csrf-Token"}

	return cors.New(corsConfig)
}
